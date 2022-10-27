package main

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"net/url"
	"path"
	"sort"
	"strings"
	"time"

	"github.com/martinisecurity/shaken-pki-reports/cmd/internal"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
)

// PrintFindingList prints the list of findings
func PrintFindingList(w io.Writer, f *Findings) {
	if f.LeafCertificates > 0 {
		fmt.Fprintf(w, "- Average validity span as of leaf certificates %d days\n", f.ValidityDays/int(f.LeafCertificates))
		fmt.Fprintf(w, "- Percentage of leaf certificates expiring in the next 30 days is %0.2f%%\n", percent(f.SoonExpiredCertificates, f.LeafCertificates))
	}
}

// PrintTotalReport prints the total report
func PrintTotalReport(w io.Writer, r *LintTotalResult) {
	fmt.Fprintln(w, "# STIR/SHAKEN CA Ecosystem Compliance")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "[Approved Certificate Authorities](https://authenticate.iconectiv.com/approved-certification-authorities) in the STIR/SHAKEN ecosystem are required to meet technical requirements from [ATIS-1000080](https://access.atis.org/apps/group_public/document.php?document_id=62163) and policy requirements from the supporting CA ecosystem’s [Certificate Policy](https://authenticate.iconectiv.com/documents-authenticate).")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "This report is generated using [Zlint](https://github.com/zmap/zlint) a tool commonly used to asses CA ecosystem compliance with such requirements. The tests used to generate this report are currently not part of the main Zlint distribution but can be found [here](https://github.com/martinisecurity/zlint).")

	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "## Summary")
	fmt.Fprintln(w, "")
	if r.LeafCertificates.Amount > 0 {
		fmt.Fprintln(w, "### Leaf Certificates")
		fmt.Fprintln(w, "")
		PrintFindingList(w, r.LeafCertificates.Findings)
		PrintOrganizationsTable(w, r.LeafCertificates, "leaf-certificates")
		fmt.Fprintln(w, "")
		fmt.Fprintln(w, "\\* The percent of certificates per issuer is calculated against total certificates from all issuers.\\")
		fmt.Fprintln(w, "\\*\\* The percent of errors, warnings and notices is calculated against total observed certificates from the specified issuer.\\")
		fmt.Fprintln(w, "\\*\\*\\* Tests use the ATIS 1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.\\")
		fmt.Fprintln(w, "\\*\\*\\*\\* For information on failed certificate repository retrievals see this [report](URL/README.md).\\")
		fmt.Fprintf(w, "\\*\\*\\*\\*\\* %d certificates skipped because they are currently expired.\\\n", r.LeafCertificates.ExpiredCertificates)
		fmt.Fprintf(w, "\\*\\*\\*\\*\\*\\* %d certificates skipped because they are not currently trusted by the STI-PA.\n", r.LeafCertificates.UntrustedCertificates)
	}
	if r.CaCertificates.Amount > 0 {
		fmt.Fprintln(w, "")
		fmt.Fprintln(w, "### CA Certificates")
		fmt.Fprintln(w, "")
		PrintOrganizationsTable(w, r.CaCertificates, "ca-certificates")
		fmt.Fprintln(w, "")
		fmt.Fprintln(w, "\\* The percent of certificates per issuer is calculated against total certificates from all issuers.\\")
		fmt.Fprintln(w, "\\*\\* The percent of errors, warnings and notices is calculated against total observed certificates from the specified issuer.\\")
		fmt.Fprintf(w, "\\*\\*\\* %d certificates skipped because they are currently expired.\\\n", r.CaCertificates.ExpiredCertificates)
		fmt.Fprintf(w, "\\*\\*\\*\\* %d certificates skipped because they are not currently trusted by the STI-PA.\n", r.CaCertificates.UntrustedCertificates)
	}

	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "## Key")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "| Type | Description |")
	fmt.Fprintln(w, "|------|-------------|")
	fmt.Fprintln(w, "| Error | Tests in which the specifications are unambiguous on what the expected behavior must be. |")
	fmt.Fprintln(w, "| Warning	| Tests in which the specifications are ambiguous or are provide only a recommendation. |")
	fmt.Fprintln(w, "| Notice | Tests in which industry best practices are not followed. |")
	fmt.Fprintln(w, "| Not Effective	| Tests that exist in the current specifications but were not in effect at the time of issuance. |")

	PrintFooter(w)
}

// PrintOrganizationReport prints Organization report
func PrintOrganizationReport(w io.Writer, name string, r *LintTotalResult) {
	// header
	fmt.Fprintln(w, "# STIR/SHAKEN CA Ecosystem Compliance")

	// summary
	fmt.Fprintln(w, "")
	fmt.Fprintf(w, "## %s\n", name)

	issuersKeys := []string{
		"Leaf Certificates",
		"CA Certificates",
	}
	issuers := map[string]*LintOrganizationResult{
		issuersKeys[0]: r.LeafCertificates.Issuers[name],
		issuersKeys[1]: r.CaCertificates.Issuers[name],
	}
	for _, issuerType := range issuersKeys {
		issuer := issuers[issuerType]
		if issuer == nil {
			continue
		}
		fmt.Fprintln(w, "")
		fmt.Fprintf(w, "### %s\n", issuerType)
		fmt.Fprintln(w, "")
		PrintFindingList(w, issuer.Findings)
		fmt.Fprintf(w, "- Certificates with Errors: %d\n", issuer.Errors)
		fmt.Fprintf(w, "- Certificates with Warnings: %d\n", issuer.Warnings)
		fmt.Fprintf(w, "- Certificates with Notices: %d\n", issuer.Notices)
		fmt.Fprintf(w, "- Certificates with tests not executed as the requirements were Not Effective at issuance time: %d\n", issuer.NE)
		fmt.Fprintln(w, "")
		fmt.Fprintln(w, "| Status | Code | Source | Instances |")
		fmt.Fprintln(w, "|--------|------|--------|-----------|")
		for code, issue := range issuer.Issues {
			rule := lint.GlobalRegistry().ByName(code)
			codeLink := fmt.Sprintf("[%s](ISSUES/%s/README.md#%s)", code, code, strings.ReplaceAll(strings.ToLower(issuerType), " ", "-"))
			fmt.Fprintf(w, "| %s | %s | %s | %d |\n", statusToString(issue.Type), codeLink, rule.Source, issue.Amount)
		}

		// summery footer
		// TODO don't show for CA
		fmt.Fprintln(w, "")
		fmt.Fprintln(w, "\\* Tests use the ATIS 1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.\\")
		fmt.Fprintf(w, "\\*\\* %d certificates skipped because they are currently expired.\\\n", issuer.ExpiredCertificates)
		fmt.Fprintf(w, "\\*\\*\\* %d certificates skipped because they are not currently trusted by the STI-PA.\n", issuer.UntrustedCertificates)

		// certificates
		fmt.Fprintln(w, "")
		fmt.Fprintln(w, "#### Issued certificates")
		fmt.Fprintln(w, "")
		fmt.Fprintln(w, "| Created at | Name | Problems | Link |")
		fmt.Fprintln(w, "|------------|------|----------|------|")
		sort.Slice(issuer.Certificates[:], func(i, j int) bool {
			return issuer.Certificates[i].Cert.NotBefore.Unix() < issuer.Certificates[j].Cert.NotBefore.Unix()
		})
		for _, certReport := range issuer.Certificates {
			fmt.Fprintf(w, "| %s | %s | %s | %s |\n",
				certReport.Cert.NotBefore.Format(time.RFC822),                                                       // created at
				certReport.Cert.Subject.CommonName,                                                                  // name
				fmt.Sprintf("%t", certReport.Result.ErrorsPresent || certReport.Result.WarningsPresent),             // problems
				fmt.Sprintf("[view](CERTIFICATES/%s)", path.Join(escapeMdLink(certReport.Thumbprint), "README.md")), // link
			)
		}
		if issuerType == "Leaf Certificates" {
			fmt.Fprintln(w, "")
			fmt.Fprintln(w, "\\* For issues relating to this CAs certificate repositories see this [report](URL/README.md).")
		}
	}

	PrintFooter(w)
}

// PrintCertificateReport prints the certificate report
func PrintCertificateReport(w io.Writer, r *LintCertificateResult) {
	fmt.Fprintf(w, "### Certificate %s\n", r.Thumbprint)
	fmt.Fprintf(w, "Tested At: %s\\\n", time.Unix(r.Result.Timestamp, 0).String())
	fmt.Fprintf(w, "Initial Validity Period: %d day(s)\\\n", internal.GetValidityDays(r.Cert))
	remainingDays := internal.GetRemainingDays(r.Cert, time.Now())
	remaining := fmt.Sprintf("%d day(s)", remainingDays)
	if remainingDays < 1 {
		remaining = "Expired"
	}
	fmt.Fprintf(w, "Remaining Validity Period: %s\\\n", remaining)
	fmt.Fprintf(w, "Subject: %s\\\n", strings.ReplaceAll(r.Cert.Subject.String(), "\\", "\\\\"))
	fmt.Fprintf(w, "Issuer: %s\n\n", strings.ReplaceAll(r.Cert.Issuer.String(), "\\", "\\\\"))
	fmt.Fprintf(w, "Link: %s\n\n", r.Link)
	fmt.Fprintf(w, "View: [Click to view](https://understandingwebpki.com/?cert=%s)\n\n", url.QueryEscape(base64.StdEncoding.EncodeToString(r.Cert.Raw)))
	fmt.Fprintln(w, "")
	fmt.Fprintf(w, "| Code | Type | Source | Details |\n")
	fmt.Fprintf(w, "|------|------|--------|---------|\n")
	for code, result := range r.Result.Results {
		if result.Status == lint.Error ||
			result.Status == lint.Warn ||
			result.Status == lint.Notice {
			rule := lint.GlobalRegistry().ByName(code)
			fmt.Fprintf(w, "| %s | %s | %s | %s |\n", code, statusToString(result.Status), rule.Source, result.Details)
		}
	}

	if !r.Result.ErrorsPresent && !r.Result.WarningsPresent {
		fmt.Fprintln(w, "")
		fmt.Fprintf(w, "%d tests were ran and no warning or error level issues were found\n", len(r.Result.Results))
	}

	neHeader := false
	for code, result := range r.Result.Results {
		if result.Status == lint.NE {
			if !neHeader {
				// Print header only once
				fmt.Fprintln(w, "")
				fmt.Fprintln(w, "### Not Effective")
				fmt.Fprintln(w, "")
				neHeader = true
			}

			fmt.Fprintf(w, "- %s\n", code)
		}
	}

	if neHeader {
		// Issue footer
		fmt.Fprintln(w, "")
		fmt.Fprintln(w, "\\* Tests use the ATIS 1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.")
	}
}

// PrintCertificateReportForIssuer prints the certificate report for specified issuer.
func PrintCertificateReportForIssuer(w io.Writer, issuerName string, r *LintCertificateResult) {
	fmt.Fprintln(w, "# STIR/SHAKEN CA Ecosystem Compliance")
	fmt.Fprintf(w, "## %s\n", issuerName)
	fmt.Fprintln(w, "")
	PrintCertificateReport(w, r)

	PrintFooter(w)
}

// PrintOrganizationsTable prints table with the list of organizations.
func PrintOrganizationsTable(w io.Writer, r *LintCertificatesResult, anchor string) {
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "| Issuers | Certificates | Errors | Warnings | Notices | Not Effective |")
	fmt.Fprintln(w, "|---------|--------------|--------|----------|---------|---------------|")

	// order r.Issuers keys
	keys := make([]string, 0, len(r.Issuers))
	for k := range r.Issuers {
		keys = append(keys, k)
	}
	sort.Slice(keys[:], func(a int, b int) bool {
		return keys[a] < keys[b]
	})

	for _, key := range keys {
		issuer := r.Issuers[key]
		if issuer.Amount == 0 {
			continue
		}
		issuerNameLink := fmt.Sprintf("[%s](%s/README.md#%s)", key, escapeMdLink(key), anchor)
		fmt.Fprintf(w, "| %s | %d (%0.2f%%) | %d (%0.2f%%) | %d (%0.2f%%) | %d (%0.2f%%) | %d (%0.2f%%) |\n", issuerNameLink, issuer.Amount, percent(issuer.Amount, r.Amount), issuer.Errors, percent(issuer.Errors, issuer.Amount), issuer.Warnings, percent(issuer.Warnings, issuer.Amount), issuer.Notices, percent(issuer.Notices, issuer.Amount), issuer.NE, percent(issuer.NE, issuer.Amount))
	}
	fmt.Fprintf(w, "| **Total** | %d (100%%) | %d (%0.2f%%) | %d (%0.2f%%) | %d (%0.2f%%) | %d (%0.2f%%) |\n", r.Amount, r.Errors, percent(r.Errors, r.Amount), r.Warnings, percent(r.Warnings, r.Amount), r.Notices, percent(r.Notices, r.Amount), r.NE, percent(r.NE, r.Amount))
	// TODO add a footnote
	// fmt.Fprintln(w, "")
	// fmt.Fprintf(w, "* 140 tests are included in this test suite")
}

// PrintIssueGroupReport prints group report for for the specified issue and organization
func PrintIssueGroupReport(w io.Writer, orgName string, issueName string, r *LintTotalResult) {
	// header
	fmt.Fprintln(w, "# STIR/SHAKEN CA Ecosystem Compliance")
	fmt.Fprintln(w, "")
	fmt.Fprintf(w, "## %s", orgName)
	fmt.Fprintln(w, "")
	issueInfo := lint.GlobalRegistry().ByName(issueName)
	fmt.Fprintf(w, "Name: %s\\\n", issueInfo.Name)
	fmt.Fprintf(w, "Source: %s\\\n", issueInfo.Source)
	fmt.Fprintf(w, "Citation: %s\\\n", issueInfo.Citation)
	fmt.Fprintf(w, "Effective Date: %s\\\n", issueInfo.EffectiveDate.Format(time.RFC822))
	fmt.Fprintf(w, "Description: %s\n", issueInfo.Description)
	fmt.Fprintln(w, "")

	leafIssuer := r.LeafCertificates.Issuers[orgName]
	if leafIssuer != nil {
		fmt.Fprintln(w, "### Leaf Certificates")
		fmt.Fprintln(w, "")
		PrintIssueGroupCertificateTable(w, issueName, leafIssuer)
	}

	caIssuer := r.CaCertificates.Issuers[orgName]
	if caIssuer != nil {
		fmt.Fprintln(w, "### CA Certificates")
		fmt.Fprintln(w, "")
		PrintIssueGroupCertificateTable(w, issueName, caIssuer)
	}

	PrintFooter(w)
}

// PrintIssueGroupCertificateTable prints table with the list of certificates for specified issuer and organization
func PrintIssueGroupCertificateTable(w io.Writer, issueName string, org *LintOrganizationResult) {
	fmt.Fprintln(w, "| Status | Subject | Link | Details |")
	fmt.Fprintln(w, "|--------|---------|------|---------|")
	counter := 0
	for _, cert := range org.Certificates {
		issue := cert.Result.Results[issueName]
		if issue == nil || issue.Status == lint.Pass || issue.Status == lint.Fatal || issue.Status == lint.NA || issue.Status == lint.Reserved {
			continue
		}
		subject := strings.ReplaceAll(cert.Cert.Subject.String(), "\\", "\\\\")
		link := fmt.Sprintf("[view](%s)", path.Join("..", "..", "CERTIFICATES", path.Join(computeCertThumbprint(cert.Cert), "README.md")))
		fmt.Fprintf(w, "| %s | %s | %s | %s |\n", statusToString(issue.Status), subject, link, issue.Details)
		counter += 1
	}

	fmt.Fprintln(w, "")
	if counter == 0 {
		fmt.Fprintln(w, "no warning, or error, or not effective date level issues were found")
		fmt.Fprintln(w, "")
	}
}

// Helpers

func computeCertThumbprint(c *x509.Certificate) string {
	thumbprint := sha1.New()
	thumbprint.Write(c.Raw)
	return hex.EncodeToString(thumbprint.Sum(nil))
}

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

func percent(a uint, b uint) float64 {
	return float64(a) / float64(b) * 100
}
