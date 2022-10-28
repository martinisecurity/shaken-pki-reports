package main

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/martinisecurity/shaken-pki-reports/cmd/internal"
	uLint "github.com/martinisecurity/shaken-pki-reports/lint"
)

// LintUrlOrgGroupResult
type LintUrlOrgGroupResult struct {
	Name  string
	Items map[string]*LintUrlOrgResult
	LintResult
}

func NewLintUrlOrgGroupResult() *LintUrlOrgGroupResult {
	return &LintUrlOrgGroupResult{
		Items: map[string]*LintUrlOrgResult{},
	}
}

// AppendLink adds URL test result into the specified organization and updates statuses
func (t *LintUrlOrgGroupResult) AppendLink(name string, l *uLint.LintResultSet) {
	// create or reuse the organization
	org := t.Items[name]
	if org == nil {
		org = NewLintUrlOrgResult()
		org.Name = name
		t.Items[name] = org
	}

	org.AppendLink(l)

	// update statuses
	t.Amount += 1
	if l.HasErrors {
		t.Errors += 1
	}
	if l.HasWarnings {
		t.Warnings += 1
	}
	if l.HasNotices {
		t.Notices += 1
	}
}

// LintUrlSummaryResult keeps summary for all URL tests
type LintUrlSummaryResult struct {
	Groups map[string]*LintUrlOrgGroupResult
	LintResult
}

// NewLintUrlSummaryResult create new instance of NewLintUrlSummaryResult
func NewLintUrlSummaryResult() *LintUrlSummaryResult {
	return &LintUrlSummaryResult{
		Groups: map[string]*LintUrlOrgGroupResult{},
	}
}

var wellknownCaDomains = map[string]string{
	"certificates.clearip.com":         "TransNexus",
	"certificates.peeringhub.io":       "Peeringhub",
	"certificates.transnexus.com":      "TransNexus",
	"cr-partner.ccid.neustar.biz":      "Neustar",
	"cr.ccid.neustar.biz":              "Neustar",
	"cr.sansay.com":                    "Sansay",
	"p.mtsec.me":                       "Martini Security",
	"prod001-cr.rbbnidhub.com":         "Ribbon Communications",
	"prod001-prod011-cr.rbbnidhub.com": "Ribbon Communications",
	"sticr.stir.comcast.com":           "Comcast",
	"t-mobile-sticr.fosrvt.com":        "T-Mobile",
	"sti-cr.cgah.tnsi.com":             "Metaswitch",
}

const (
	groupName_CA              = "CA"
	groupName_ServiceProvider = "SP"
)

func (t *LintUrlSummaryResult) GetOrgName(l *uLint.LintResultSet) (string, string) {
	var caName string
	u, _ := url.Parse(l.Url)
	if u == nil {
		return "Unknown", groupName_ServiceProvider
	} else {
		if caName = wellknownCaDomains[u.Hostname()]; len(caName) > 0 {
			// CA
			return caName, groupName_CA
		}

		// Service Provider
		var spName string
		if l.StatusCode == 200 {
			if certs := internal.ParseCertificates(l.Body); len(certs) > 0 {
				cert := certs[0].Certificate
				if len(cert.Subject.Organization) > 0 {
					// use O name from the Subject
					return cert.Subject.Organization[0], groupName_ServiceProvider
				} else if len(cert.Issuer.Organization) > 0 {
					// use O name from the Issuer
					return cert.Issuer.Organization[1], groupName_ServiceProvider
				} else if spc, err := internal.GetTNEntrySPC(cert); err == nil {
					// use SPC value from the TNAuthList extension
					return fmt.Sprintf("SHAKEN %s", spc), groupName_ServiceProvider
				}
			}
		}

		// use Host name from the URL
		parts := strings.Split(u.Hostname(), ".")
		if _, err := strconv.Atoi(parts[len(parts)-1]); err == nil {
			// IP address
			spName = u.Hostname()
		} else {
			// domain
			spName = fmt.Sprintf("%s.%s", parts[len(parts)-2], parts[len(parts)-1])
		}

		return spName, groupName_ServiceProvider
	}
}

// AppendLink adds URL test result into the specified organization group and updates statuses
func (t *LintUrlSummaryResult) AppendLink(l *uLint.LintResultSet) {
	orgName, groupName := t.GetOrgName(l)

	// create or reuse the group
	group := t.Groups[groupName]
	if group == nil {
		group = NewLintUrlOrgGroupResult()
		group.Name = groupName
		t.Groups[groupName] = group
	}

	group.AppendLink(orgName, l)

	// update statuses
	t.Amount += 1
	if l.HasErrors {
		t.Errors += 1
	}
	if l.HasWarnings {
		t.Warnings += 1
	}
	if l.HasNotices {
		t.Notices += 1
	}
}

// LintUrlOrgResult contains test result for the organization
type LintUrlOrgResult struct {
	// Name of the organization
	Name string
	// Links keeps test URL result
	Links []*uLint.LintResultSet
	// Problems keeps map of problem
	Problems map[string]int
	LintResult
}

// NewLintUrlOrgResult creates new instance of LintUrlOrgResult
func NewLintUrlOrgResult() *LintUrlOrgResult {
	return &LintUrlOrgResult{
		Problems: map[string]int{},
	}
}

// AppendLink adds URL test result lint and updates statuses
func (t *LintUrlOrgResult) AppendLink(l *uLint.LintResultSet) {
	// add test result into the list
	t.Links = append(t.Links, l)

	// add each problem into the map
	for code, r := range l.Results {
		if r.Status != uLint.Pass {
			t.Problems[code] += 1
		}
	}

	// update statuses
	t.Amount += 1
	if l.HasErrors {
		t.Errors += 1
	}
	if l.HasWarnings {
		t.Warnings += 1
	}
	if l.HasNotices {
		t.Notices += 1
	}
}

// var wellknownDomainNames = map[string]string{}

// func getClientName(domain string) string {
// 	if name := wellknownDomainNames[domain]; len(name) > 0 {
// 		return name
// 	}

// 	return "Unknown"
// }
