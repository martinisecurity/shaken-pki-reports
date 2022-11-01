package main

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/url"
	"path"
	"sort"
	"strings"
	"time"

	"github.com/martinisecurity/shaken-pki-reports/cmd/internal"
	uLint "github.com/martinisecurity/shaken-pki-reports/lint"
	"github.com/zmap/zlint/v3/lint"
)

const (
	HEADER_CERTS = "# STIR/SHAKEN CA Ecosystem Compliance"
	HEADER_REPOS = "# STIR/SHAKEN Certificate Repository Compliance"
)

func PrintCertificateSummaryReport(w io.Writer, r *CertificateSummaryReport) {
	fmt.Fprintln(w, HEADER_CERTS)
	fmt.Fprintln(w)

	fmt.Fprintln(w, "[Approved Certificate Authorities](https://ecosystemcompliance.martinisecurity.com/#:~:text=Approved%20Certificate%20Authorities) in the STIR/SHAKEN ecosystem are required to meet technical requirements from [ATIS-1000080](https://access.atis.org/apps/group_public/document.php?document_id=62163) and policy requirements from the supporting CA ecosystem’s [Certificate Policy](https://authenticate.iconectiv.com/documents-authenticate).")
	fmt.Fprintln(w)

	fmt.Fprintln(w, "This report is broken int two parts:")
	fmt.Fprintln(w, "1. One generated using [Zlint](https://github.com/zmap/zlint) a tool commonly used to asses CA ecosystem compliance with such requirements. The tests used to generate this report are currently not part of the main [Zlint](https://github.com/martinisecurity/zlint) distribution but can be found here.")
	fmt.Fprintln(w, "2. One generated with a custom script that eumerates the known STIR/SHAKEN certificates and asses each repository against the current rule set . The source for this test can be found here while the report itself can be found [here](REPOS/README.md).")
	fmt.Fprintln(w)

	fmt.Fprintln(w, "## Summary")
	fmt.Fprintln(w)

	fmt.Fprintln(w, "### Leaf Certificates")
	fmt.Fprintln(w)
	PrintCertificateFindings(w, &r.Leaf.CertificateGroupReport)
	fmt.Fprintln(w)
	fmt.Fprintln(w, "### CA Certificates")
	fmt.Fprintln(w)
	PrintCertificateFindings(w, &r.CA.CertificateGroupReport)
	fmt.Fprintln(w)

	fmt.Fprintln(w, "## Certificate Repository")
	fmt.Fprintln(w)
	fmt.Fprintf(w, "- %0.2f%% of certificate repositories contain one or more Error level issue\n", r.Leaf.AverageRepositoryErrors())
	fmt.Fprintf(w, "- %0.2f%% of certificates repositories contain one or more Warning level issue\n", r.Leaf.AverageRepositoryWarns())
	fmt.Fprintf(w, "- %0.2f%% of certificates repositories contain one or more Notice level issue\n", r.Leaf.AverageRepositoryNotices())
	fmt.Fprintln(w)

	fmt.Fprintln(w, "## Details")
	fmt.Fprintln(w)
	fmt.Fprintln(w, "\\* The percent of certificates per issuer is calculated against total certificates from all issuers.\\")
	fmt.Fprintln(w, "\\*\\* The percent of errors, warnings and notices is calculated against total observed certificates from the specified issuer.\\")
	fmt.Fprintln(w, "\\*\\*\\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.")
	fmt.Fprintln(w)
	fmt.Fprintln(w, "### Leaf Certificates")
	fmt.Fprintln(w)
	PrintCertificateSummaryIssuers(w, r.Leaf, "leaf-certificates")
	fmt.Fprintln(w)

	fmt.Fprintln(w, "### CA Certificates")
	fmt.Fprintln(w)
	PrintCertificateSummaryIssuers(w, r.CA, "ca-certificates")
	fmt.Fprintln(w)

	fmt.Fprintln(w, "### Key")
	fmt.Fprintln(w)
	fmt.Fprintln(w, "| Type | Description |")
	fmt.Fprintln(w, "|------|-------------|")
	fmt.Fprintln(w, "| Errors | Tests in which the specifications are unambiguous on what the expected behavior must be. |")
	fmt.Fprintln(w, "| Warnings | Tests in which the specifications are ambiguous or are provide only a recommendation. |")
	fmt.Fprintln(w, "| Notices | Tests in which industry best practices are not followed. |")
	fmt.Fprintln(w, "| Not Effective | Tests that exist in the current specifications but were not in effect at the time of issuance. |")
	fmt.Fprintln(w)

	PrintFooter(w)
}

func PrintCertificateFindings(w io.Writer, r *CertificateGroupReport) {
	fmt.Fprintf(w, "- %d certificates were included in the corpus being tested\n", r.Amount)
	fmt.Fprintf(w, "- %d repositories in the corpus were skipped because they are duplicates\n", r.SkippedRepositoriesAmount)
	fmt.Fprintf(w, "- %d certificates in the corpus were skipped because they are expired\n", r.SkippedExpiredAmount)
	fmt.Fprintf(w, "- %d certificates in the corpus were skipped because they are not currently trusted\n", r.SkippedUntrustedAmount)
	fmt.Fprintf(w, "- %d certificates being tested against the remaining rules\n", r.TestedAmount)
	fmt.Fprintf(w, "- %0.2f%% of certificates contain one or more Error level issue\n", r.AverageErrors())
	fmt.Fprintf(w, "- %0.2f%% of certificates contain one or more Warning level issue\n", r.AverageWarns())
	fmt.Fprintf(w, "- %0.2f%% of certificates contain one or more Notice level issue\n", r.AverageNotices())
	fmt.Fprintf(w, "- %0.2f%% of certificates are too old to be assessed against currently enforced expectations\n", r.AverageNotEffective())
	fmt.Fprintf(w, "- %0.0f days is the average remaining validity for the certificates in the corpus\n", r.AverageRemainingValidity())
	fmt.Fprintf(w, "- %0.0f days is the average initial validity for the certificates in the corpus\n", r.AverageInitialValidity())
	fmt.Fprintf(w, "- %d certificates expire in the next 30 days\n", r.ExpiresSoon)
}

func PrintCertificateSummaryIssuers(w io.Writer, r *CertificateIssuersReport, anchor string) {
	fmt.Fprintln(w, "| Issuers | Certificates | Errors | Warnings | Notices | Not Effective |")
	fmt.Fprintln(w, "|---------|--------------|--------|----------|---------|---------------|")
	print := func(amount int) string {
		return fmt.Sprintf("%d (%0.2f%%)", amount, percent(amount, r.TestedAmount))
	}

	// sort issuer names
	keys := make([]string, 0, len(r.Issuers))
	for k := range r.Issuers {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, key := range keys {
		issuer := r.Issuers[key]
		if issuer.TestedAmount == 0 {
			continue
		}
		fmt.Fprintf(w, "| %s | %s | %s | %s | %s | %s |\n",
			fmt.Sprintf("[%s](%s/%s/README.md#%s)", issuer.Name, DIR_CERTS, escapeMdLink(issuer.Name), anchor),
			print(issuer.TestedAmount),
			print(issuer.ErrorAmount),
			print(issuer.WarnAmount),
			print(issuer.NoticeAmount),
			print(issuer.NotEffectiveAmount),
		)
	}
	fmt.Fprintf(w, "| **Total** | %s | %s | %s | %s | %s |\n",
		print(r.TestedAmount),
		print(r.ErrorAmount),
		print(r.WarnAmount),
		print(r.NoticeAmount),
		print(r.NotEffectiveAmount),
	)
}

func PrintCertificateIssuerReport(w io.Writer, r *CertificateIssuerJoin) {
	fmt.Fprintln(w, HEADER_CERTS)
	fmt.Fprintln(w)

	fmt.Fprintf(w, "## %s\n", r.Name)
	fmt.Fprintln(w)

	fmt.Fprintln(w, "### Summary")
	fmt.Fprintln(w)
	fmt.Fprintln(w, "\\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\\")
	fmt.Fprintln(w, "\\*\\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.")
	fmt.Fprintln(w)

	if r.Leaf != nil {
		fmt.Fprintln(w, "#### Leaf Certificates")
		fmt.Fprintln(w)
		PrintCertificateFindings(w, &r.Leaf.CertificateGroupReport)
		fmt.Fprintln(w)
		PrintProblems(w, r.Leaf.Problems)
		fmt.Fprintln(w)
	}
	if r.CA != nil {
		fmt.Fprintln(w, "#### CA Certificates")
		fmt.Fprintln(w)
		PrintCertificateFindings(w, &r.CA.CertificateGroupReport)
		fmt.Fprintln(w)
		PrintProblems(w, r.CA.Problems)
		fmt.Fprintln(w)
	}

	fmt.Fprintln(w, "### Details")
	fmt.Fprintln(w)

	if r.Leaf != nil {
		fmt.Fprintln(w, "#### Leaf Certificates")
		fmt.Fprintln(w)
		PrintCertificates(w, r.Leaf.Items, DIR_CERTS)
		fmt.Fprintln(w)
	}
	if r.CA != nil {
		fmt.Fprintln(w, "#### CA Certificates")
		fmt.Fprintln(w)
		PrintCertificates(w, r.CA.Items, DIR_CERTS)
		fmt.Fprintln(w)
	}

	PrintFooter(w)
}

func PrintCertificateReport(w io.Writer, r *LintCommandItem) {
	fmt.Fprintln(w, HEADER_CERTS)
	fmt.Fprintln(w)

	fmt.Fprintf(w, "## Certificate %s\n", r.Certificate.Subject.CommonName)
	fmt.Fprintln(w)
	fmt.Fprintf(w, "Tested At: %s\\\n", time.Unix(r.CertificateResult.Timestamp, 0).Format(time.RFC822))
	fmt.Fprintf(w, "Initial Validity Period: %d day(s)\\\n", internal.GetValidityDays(r.Certificate))
	fmt.Fprintf(w, "Remaining Validity Period: %d day(s)\\\n", internal.GetRemainingDays(r.Certificate, time.Now()))
	fmt.Fprintf(w, "Subject: %s\\\n", strings.ReplaceAll(r.Certificate.Subject.String(), "\\", "\\\\"))
	fmt.Fprintf(w, "Issuer: %s", strings.ReplaceAll(r.Certificate.Issuer.String(), "\\", "\\\\"))
	if r.Url != nil {
		fmt.Fprintln(w, "\\")
		fmt.Fprintf(w, "Link: %s\n", r.Url.String())
	} else {
		fmt.Fprintln(w)
	}
	fmt.Fprintln(w)
	fmt.Fprintf(w, "View: [Click to view](https://understandingwebpki.com/?cert=%s)\n\n", url.QueryEscape(base64.StdEncoding.EncodeToString(r.Certificate.Raw)))
	fmt.Fprintln(w)
	first := true
	for code, result := range r.CertificateResult.Results {
		if result.Status == lint.Error ||
			result.Status == lint.Warn ||
			result.Status == lint.Notice {
			if first {
				fmt.Fprintf(w, "| Code | Type | Source | Details |\n")
				fmt.Fprintf(w, "|------|------|--------|---------|\n")
				first = false
			}
			rule := lint.GlobalRegistry().ByName(code)
			fmt.Fprintf(w, "| %s | %s | %s | %s |\n",
				fmt.Sprintf("[%s](%s)", code, path.Join("..", "..", DIR_ISSUES, code, "README.md")),
				statusToString(result.Status),
				strings.ReplaceAll(string(rule.Source), "-", "&#x2011;"),
				result.Details,
			)
		}
	}

	if first {
		fmt.Fprintln(w)
		fmt.Fprintf(w, "%d tests were ran and no warning or error level issues were found\n", len(r.CertificateResult.Results))
	}
	fmt.Fprintln(w)

	neHeader := false
	codeKeys := []string{}
	for key := range r.CertificateResult.Results {
		codeKeys = append(codeKeys, key)
	}
	sort.Strings(codeKeys)
	for _, key := range codeKeys {
		result := r.CertificateResult.Results[key]
		if result.Status == lint.NE {
			if !neHeader {
				// Print header only once
				fmt.Fprintln(w)
				fmt.Fprintln(w, "### Not Effective")
				fmt.Fprintln(w)
				neHeader = true
			}

			fmt.Fprintf(w, "- %s\n", key)
		}
	}

	if neHeader {
		// Issue footer
		fmt.Fprintln(w)
		fmt.Fprintln(w, "\\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.")
		fmt.Fprintln(w)
	}

	PrintFooter(w)
}

func PrintCertificateIssueReport(w io.Writer, c string, r *CertificateIssuerJoin) {
	fmt.Fprintln(w, HEADER_CERTS)
	fmt.Fprintln(w)

	fmt.Fprintf(w, "## %s\n", r.Name)
	fmt.Fprintln(w)

	issueInfo := lint.GlobalRegistry().ByName(c)
	fmt.Fprintf(w, "Name: %s\\\n", issueInfo.Name)
	fmt.Fprintf(w, "Source: %s\\\n", issueInfo.Source)
	fmt.Fprintf(w, "Citation: %s\\\n", issueInfo.Citation)
	fmt.Fprintf(w, "Effective Date: %s\\\n", issueInfo.EffectiveDate.Format(time.RFC822))
	fmt.Fprintf(w, "Description: %s\n", issueInfo.Description)
	fmt.Fprintln(w)

	certsDir := path.Join("..", "..", DIR_CERTS)
	if r.Leaf != nil {
		fmt.Fprintln(w, "### Leaf Certificates")
		fmt.Fprintln(w)
		PrintIssueCertificates(w, c, r.Leaf.Problems[c], certsDir)
		fmt.Fprintln(w)
	}

	if r.CA != nil {
		fmt.Fprintln(w, "### CA Certificates")
		fmt.Fprintln(w)
		PrintIssueCertificates(w, c, r.CA.Problems[c], certsDir)
		fmt.Fprintln(w)
	}

	PrintFooter(w)
}

func PrintIssueCertificates(w io.Writer, c string, r *Problem, b string) {
	if r != nil {
		fmt.Fprintln(w, "| Status | Subject | Link | Details |")
		fmt.Fprintln(w, "|--------|---------|------|---------|")
		for _, v := range r.Items {
			issue := v.CertificateResult.Results[c]
			if issue.Status == lint.Error ||
				issue.Status == lint.Warn ||
				issue.Status == lint.Notice ||
				issue.Status == lint.NE {
				fmt.Fprintf(w, "| %s | %s | %s | %s |\n",
					statusToString(issue.Status),
					v.Certificate.Subject.CommonName,
					fmt.Sprintf("[view](%s)", path.Join(b, path.Join(getCertificateId(v.Certificate), "README.md"))),
					issue.Details,
				)
			}
		}
	} else {
		fmt.Fprintln(w, "no warning, or error, or not effective date level issues were found")
	}
}

func PrintCertificates(w io.Writer, r []*LintCommandItem, basePath string) {
	fmt.Fprintln(w, "| Created at | Subject | Problems | Link |")
	fmt.Fprintln(w, "|------------|---------|----------|------|")
	// order by notBefore
	sort.Slice(r[:], func(i, j int) bool {
		return r[i].Certificate.NotBefore.Before(r[j].Certificate.NotBefore)
	})
	for _, v := range r {
		fmt.Fprintf(w, "| %s | %s | %t | %s |\n",
			v.Certificate.NotBefore.Format(time.RFC822),
			v.Certificate.Subject.CommonName,
			v.HasCertificateProblems(),
			fmt.Sprintf("[view](%s)", path.Join(basePath, getCertificateId(v.Certificate), "README.md")),
		)
	}
}

func PrintProblems(w io.Writer, r map[string]*Problem) {
	fmt.Fprintln(w, "| Instances | Test Status | Source |")
	fmt.Fprintln(w, "|-----------|-------------|--------|")

	// sort problems by code names
	keys := make([]string, 0, len(r))
	for k := range r {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, key := range keys {
		p := r[key]
		fmt.Fprintf(w, "| %d | %s | %s |\n",
			len(p.Items),
			fmt.Sprintf("[%s](%s)", key, path.Join(DIR_ISSUES, key, "README.md")),
			p.Source,
		)
	}
}

func PrintRepositorySummaryReport(w io.Writer, r *RepositorySummaryReport) {
	fmt.Fprintln(w, HEADER_REPOS)
	fmt.Fprintln(w)

	fmt.Fprintln(w, "## Summary")
	fmt.Fprintln(w)

	if r.CA.TestedAmount > 0 && r.SP.TestedAmount > 0 {
		fmt.Fprintln(w, "### All Repositories")
		fmt.Fprintln(w)
		PrintRepositoryFindings(w, &r.RepositoryGroupReport)
		fmt.Fprintln(w)
	}

	if r.CA.TestedAmount > 0 {
		fmt.Fprintln(w, "### CA Operated Repositories")
		fmt.Fprintln(w)
		PrintRepositoryFindings(w, &r.CA.RepositoryGroupReport)
		fmt.Fprintln(w)
	}

	if r.SP.TestedAmount > 0 {
		fmt.Fprintln(w, "### Service Provider Operated Repositories")
		fmt.Fprintln(w)
		PrintRepositoryFindings(w, &r.SP.RepositoryGroupReport)
		fmt.Fprintln(w)
	}

	fmt.Fprintln(w, "## Details")
	fmt.Fprintln(w)

	if r.CA.TestedAmount > 0 {
		fmt.Fprintln(w, "### CA Operated Repositories")
		fmt.Fprintln(w)
		PrintSummaryDetails(w, r.CA)
		fmt.Fprintln(w)
	}

	if r.SP.TestedAmount > 0 {
		fmt.Fprintln(w, "### Service Provider Operated Repositories")
		fmt.Fprintln(w)
		PrintSummaryDetails(w, r.SP)
		fmt.Fprintln(w)
	}

	fmt.Fprintln(w, "### Key")
	fmt.Fprintln(w)
	fmt.Fprintln(w, "| Type | Description |")
	fmt.Fprintln(w, "|------|-------------|")
	fmt.Fprintln(w, "| Errors | Tests in which the specifications are unambiguous on what the expected behavior must be. |")
	fmt.Fprintln(w, "| Warnings | Tests in which the specifications are ambiguous or are provide only a recommendation. |")
	fmt.Fprintln(w, "| Notices | Tests in which industry best practices are not followed. |")
	fmt.Fprintln(w)

	PrintFooter(w)
}

func PrintRepositoryIssuerReport(w io.Writer, r *RepositoryIssuerReport) {
	fmt.Fprintln(w, HEADER_REPOS)
	fmt.Fprintln(w)

	fmt.Fprintf(w, "## %s\n", r.Name)
	fmt.Fprintln(w)

	PrintRepositoryFindings(w, &r.RepositoryGroupReport)
	fmt.Fprintln(w)
	if len(r.Problems) > 0 {
		PrintProblems(w, r.Problems)
	} else {
		fmt.Fprintln(w, "No issues found")
	}
	fmt.Fprintln(w)
	PrintRepositories(w, r.Items, DIR_REPOS)
	fmt.Fprintln(w)

	PrintFooter(w)
}

func PrintRepositoryIssuerUrlReport(w io.Writer, issuer *RepositoryIssuerReport, item *LintCommandItem) {
	fmt.Fprintln(w, HEADER_REPOS)
	fmt.Fprintln(w)

	fmt.Fprintf(w, "## %s\n", issuer.Name)
	fmt.Fprintln(w)
	fmt.Fprintf(w, "Name: `%s`\\\n", item.Url)
	fmt.Fprintf(w, "Tested At: %s\\\n", item.UrlResult.Timestamp.Format(time.RFC822))
	fmt.Fprintf(w, "Time: %dms\n", item.UrlResult.Time)
	fmt.Fprintln(w)

	// list problems
	fmt.Fprintln(w, "### Issues")
	fmt.Fprintln(w)
	r := item.UrlResult.Results

	keys := []string{}
	for key := range r {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	first := true
	for _, key := range keys {
		issue := r[key]
		if issue.Status == uLint.Error ||
			issue.Status == uLint.Warn ||
			issue.Status == uLint.Notice {
			if first {
				fmt.Fprintln(w, "| Code | Type | Source | Details |")
				fmt.Fprintln(w, "|------|------|--------|---------|")
				first = false
			}

			l := uLint.FindRuleByName(key)
			fmt.Fprintf(w, "| %s | %s | %s | %s |\n",
				fmt.Sprintf("[%s](%s)", key, path.Join("..", "..", DIR_ISSUES, key, "README.md")),
				issue.Status,
				l.Source,
				issue.Details,
			)
		}
	}

	if first {
		fmt.Fprintln(w, "no warning, or error, or notice date level issues were found")
	}

	PrintFooter(w)
}

func PrintRepositoryIssuerProblemReport(w io.Writer, issuer *RepositoryIssuerReport, problem *Problem) {
	fmt.Fprintln(w, HEADER_REPOS)
	fmt.Fprintln(w)

	fmt.Fprintf(w, "## %s\n", issuer.Name)
	fmt.Fprintln(w)
	fmt.Fprintf(w, "Name: %s\\\n", problem.Name)
	l := uLint.FindRuleByName(problem.Name)
	fmt.Fprintf(w, "Source: %s\\\n", l.Source)
	fmt.Fprintf(w, "Description: %s\n", l.Description)

	// list repositories
	fmt.Fprintln(w, "### Repositories")
	fmt.Fprintln(w)
	PrintRepositories(w, problem.Items, path.Join("..", "..", DIR_REPOS))
	fmt.Fprintln(w)

	PrintFooter(w)
}

func PrintRepositories(w io.Writer, r []*LintCommandItem, basePath string) {
	fmt.Fprintln(w, "| Repository | Problems | Link |")
	fmt.Fprintln(w, "|------------|----------|------|")
	sort.Slice(r[:], func(i, j int) bool {
		return r[i].Url.String() < r[j].Url.String()
	})
	for _, v := range r {
		fmt.Fprintf(w, "| %s | %t | %s |\n",
			fmt.Sprintf("`%s`", v.Url),
			v.UrlResult.HasErrors || v.UrlResult.HasWarnings || v.UrlResult.HasNotices,
			fmt.Sprintf("[view](%s)", path.Join(basePath, getRepositoryId(v.Url), "README.md")),
		)
	}
}

func PrintRepositoryFindings(w io.Writer, r *RepositoryGroupReport) {
	fmt.Fprintf(w, "- %d repositories were included in the corpus being tested\n", r.Amount)
	fmt.Fprintf(w, "- %d repositories in the corpus were skipped because they were duplicated\n", r.SkippedAmount)
	fmt.Fprintf(w, "- %d repositories being tested against the remaining rules\n", r.TestedAmount)
	fmt.Fprintf(w, "- %0.2f%% of repositories contain one or more Error level issue\n", r.AverageErrors())
	fmt.Fprintf(w, "- %0.2f%% of repositories contain one or more Warning level issue\n", r.AverageWarns())
	fmt.Fprintf(w, "- %0.2f%% of repositories contain one or more Notice level issue\n", r.AverageNotices())
	fmt.Fprintf(w, "- %0.0fms average time it took to download each certificate\n", r.AverageTime())
}

func PrintSummaryDetails(w io.Writer, r *RepositoryIssuersReport) {
	name := "Issuers"
	dir := DIR_CA
	if r.Type == SP {
		name = "Providers"
		dir = DIR_SP
	}
	fmt.Fprintf(w, "| %s | Certificates | Errors | Warnings | Notices |\n", name)
	fmt.Fprintln(w, "|----|--------------|--------|----------|---------|")

	print := func(amount int) string {
		return fmt.Sprintf("%d (%0.2f%%)", amount, percent(amount, r.TestedAmount))
	}

	keys := []string{}
	for key := range r.Issuers {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		issuer := r.Issuers[key]
		fmt.Fprintf(w, "| %s | %s | %s | %s | %s |\n",
			fmt.Sprintf("[%s](%s)", issuer.Name, path.Join(DIR_ISSUERS, dir, escapeMdLink(issuer.Name), "README.md")),
			print(issuer.TestedAmount),
			print(issuer.ErrorAmount),
			print(issuer.WarnAmount),
			print(issuer.NoticeAmount),
		)
	}
	fmt.Fprintf(w, "| **Total** | %s | %s | %s | %s |\n",
		print(r.TestedAmount),
		print(r.ErrorAmount),
		print(r.WarnAmount),
		print(r.NoticeAmount),
	)
}

// Helpers

func statusToString(s lint.LintStatus) string {
	switch s {
	case lint.Error:
		return "error"
	case lint.Warn:
		return "warn"
	case lint.Notice:
		return "notice"
	case lint.NE:
		return "not effective"
	default:
		return s.String()
	}
}

func percent(a int, b int) float64 {
	return float64(a) / float64(b) * 100
}
