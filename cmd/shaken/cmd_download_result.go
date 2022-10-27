package main

import (
	"net/url"

	uLint "github.com/martinisecurity/shaken-pki-reports/lint"
)

// LintUrlSummaryResult keeps summary for all URL tests
type LintUrlSummaryResult struct {
	Organizations map[string]*LintUrlOrgResult
	Clients       map[string]*LintUrlOrgResult
	LintResult
}

// NewLintUrlSummaryResult create new instance of NewLintUrlSummaryResult
func NewLintUrlSummaryResult() *LintUrlSummaryResult {
	return &LintUrlSummaryResult{
		Organizations: map[string]*LintUrlOrgResult{},
		Clients:       map[string]*LintUrlOrgResult{},
	}
}

// AppendLink adds URL test result into the specified organization and updates statuses
func (t *LintUrlSummaryResult) AppendLink(name string, l *uLint.LintResultSet) {
	// create or reuse the organization
	org := t.Organizations[name]
	if org == nil {
		org = NewLintUrlOrgResult()
		org.Name = name
		t.Organizations[name] = org
	}

	org.AppendLink(l)

	// get host
	host := "Unknown"
	u, err := url.Parse(l.Url)
	if err == nil {
		host = u.Host
	}
	client := t.Clients[host]
	if client == nil {
		client = NewLintUrlOrgResult()
		client.Name = host
		t.Clients[name] = client
	}
	client.AppendLink(l)

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
