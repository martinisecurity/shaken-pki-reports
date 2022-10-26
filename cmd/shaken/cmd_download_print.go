package main

import (
	"fmt"
	"io"
	"net/url"
	"path"
	"sort"

	uLint "github.com/martinisecurity/shaken-pki-reports/lint"
)

// PrintUrlIssue prints a report for specified rule
func PrintUrlIssue(w io.Writer, code string, r *LintUrlOrgResult) {
	// header
	fmt.Fprintln(w, "# STIR/SHAKEN Certificate Repository Compliance")
	fmt.Fprintln(w, "")

	// organization name
	fmt.Fprintf(w, "## %s\n", r.Name)
	fmt.Fprintln(w, "")

	// rule details
	fmt.Fprintf(w, "Code: %s\\\n", code)
	rule := uLint.FindRuleByName(code)
	if rule != nil {
		fmt.Fprintf(w, "Source: %s\\\n", rule.Source)
		fmt.Fprintf(w, "Description: %s\n", rule.Description)
	}

	for _, link := range r.Links {
		if !link.HasProblem(code) {
			continue
		}

		PrintUrlDetails(w, link)
	}

	// footer
	PrintFooter(w)
}

// PrintUrlDetails prints detailed information for the URL
func PrintUrlDetails(w io.Writer, r *uLint.LintResultSet) {
	fmt.Fprintf(w, "### %s\n", r.Url)
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "| Code | Status | Source | Details |")
	fmt.Fprintln(w, "|------|--------|--------|---------|")
	counter := 0
	for code, l := range r.Results {
		if l.Status != uLint.Pass {
			counter += 1
			ruleInfo := uLint.FindRuleByName(code)
			fmt.Fprintf(w, "| %s | %s | %s | %s |\n", code, l.Status, ruleInfo.Source, l.Details)
		}
	}
	fmt.Fprintln(w, "")
	if counter == 0 {
		fmt.Fprintf(w, "%d tests were ran and no warning or error level issues were found\n", len(r.Results))
		fmt.Fprintln(w, "")
	}
}

// PrintUrlOrg prints URL report for the organization
func PrintUrlOrg(w io.Writer, r *LintUrlOrgResult) {
	fmt.Fprintln(w, "# STIR/SHAKEN Certificate Repository Compliance")
	fmt.Fprintln(w, "")
	fmt.Fprintf(w, "## %s\n", r.Name)
	fmt.Fprintln(w, "")

	fmt.Fprintln(w, "| Code | Source | Instances |")
	fmt.Fprintln(w, "|------|--------|-----------|")
	for code, instances := range r.Problems {
		rule := uLint.FindRuleByName(code)
		codeLink := fmt.Sprintf("[%s](ISSUES/%s.%s)", code, code, linkExt)
		fmt.Fprintf(w, "| %s | %s | %d |\n", codeLink, rule.Source, instances)
	}
	fmt.Fprintln(w, "")

	for _, v := range r.Links {
		fmt.Fprintf(w, "### %s\n", v.Url)
		fmt.Fprintln(w, "")
		fmt.Fprintln(w, "| Code | Status | Source | Details |")
		fmt.Fprintln(w, "|------|--------|--------|---------|")
		counter := 0
		for code, l := range v.Results {
			if l.Status != uLint.Pass {
				counter += 1
				ruleInfo := uLint.FindRuleByName(code)
				fmt.Fprintf(w, "| %s | %s | %s | %s |\n", code, l.Status, ruleInfo.Source, l.Details)
			}
		}
		fmt.Fprintln(w, "")
		if counter == 0 {
			fmt.Fprintf(w, "%d tests were ran and no warning or error level issues were found\n", len(v.Results))
			fmt.Fprintln(w, "")
		}
	}

	PrintFooter(w)
}

// PrintUrlSummary prints summary report
func PrintUrlSummary(w io.Writer, r *LintUrlSummaryResult) {
	fmt.Fprintln(w, "# STIR/SHAKEN Certificate Repository Compliance")
	fmt.Fprintln(w, "")

	fmt.Fprintln(w, "Participants in the STIR/SHAKEN ecosystem are required to publish the certificates they use for signing calls so that other participants can retrieve these certificates and use them for validating signatures.")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "To ease meeting this requirement many of the certificate authorities publish the certificates for their customers and the majority of the STIR/SHAKEN deployments use these CA-provided repositories. There are still some cases where customers choose to host the certificate on their own.")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "This report looks at what errors and compliance issues relying parties may experience when trying to use these certificate repositories.")
	fmt.Fprintln(w, "")

	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "| Issuers | Links | Errors | Warnings | Notices |")
	fmt.Fprintln(w, "|---------|-------|--------|----------|---------|")

	// order r.Issuers keys
	keys := make([]string, 0, len(r.Organizations))
	for k := range r.Organizations {
		keys = append(keys, k)
	}
	sort.Slice(keys[:], func(a int, b int) bool {
		return keys[a] < keys[b]
	})

	for _, key := range keys {
		issuer := r.Organizations[key]
		issuerNameLink := fmt.Sprintf("[%s](%s)", key, url.PathEscape(path.Join(key, fmt.Sprintf("URL.%s", linkExt))))
		fmt.Fprintf(w, "| %s | %d (%0.2f%%) | %d (%0.2f%%) | %d (%0.2f%%) | %d (%0.2f%%) |\n", issuerNameLink, issuer.Amount, percent(issuer.Amount, r.Amount), issuer.Errors, percent(issuer.Errors, issuer.Amount), issuer.Warnings, percent(issuer.Warnings, issuer.Amount), issuer.Notices, percent(issuer.Notices, issuer.Amount))
	}
	fmt.Fprintf(w, "| **Total** | %d (100%%) | %d (%0.2f%%) | %d (%0.2f%%) | %d (%0.2f%%) |\n", r.Amount, r.Errors, percent(r.Errors, r.Amount), r.Warnings, percent(r.Warnings, r.Amount), r.Notices, percent(r.Notices, r.Amount))

	PrintFooter(w)
}
