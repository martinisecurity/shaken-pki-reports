package main

import (
	"sort"
	"time"

	"github.com/martinisecurity/shaken-pki-reports/cmd/internal"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3"
	"github.com/zmap/zlint/v3/lint"
)

// LintCertificateResult represents extended ZLint test result
type LintCertificateResult struct {
	Link         string
	Cert         *x509.Certificate
	Thumbprint   string
	Result       *zlint.ResultSet
	Organization string
	IsExpired    bool
	IsUntrusted  bool
}

type Findings struct {
	LeafCertificates        uint
	ValidityDays            int
	SoonExpiredCertificates uint
}

type LintIssue struct {
	Type   lint.LintStatus
	Amount uint
}

type LintOrganizationResult struct {
	Name                  string
	Issues                map[string]*LintIssue
	Certificates          []*LintCertificateResult
	ExpiredCertificates   int
	UntrustedCertificates int
	LintResult
	*Findings
}

// AppendCertificate adds certificate test result and updates statuses
func (t *LintOrganizationResult) AppendCertificate(c *LintCertificateResult) {
	t.Certificates = append(t.Certificates, c)
	nePresents := false

	// Update Issues
	for code, result := range c.Result.Results {
		// filter results by Status
		if !(result.Status == lint.Error ||
			result.Status == lint.Warn ||
			result.Status == lint.Notice ||
			result.Status == lint.NE) {
			continue
		}
		issue := t.Issues[code]
		if issue == nil {
			issue = &LintIssue{
				Type: result.Status,
			}
			t.Issues[code] = issue
		}

		issue.Amount += 1

		// c.Result doesn't have NEPresents
		if !nePresents && result.Status == lint.NE {
			nePresents = true
		}
	}

	// Update counters
	t.Amount += 1
	if c.Result.ErrorsPresent {
		t.Errors += 1
	}
	if c.Result.WarningsPresent {
		t.Warnings += 1
	}
	if c.Result.NoticesPresent {
		t.Notices += 1
	}
	if nePresents {
		t.NE += 1
	}

	// finding
	if !c.Cert.IsCA {
		t.LeafCertificates += 1
		t.ValidityDays += internal.GetValidityDays(c.Cert)

		if time.Now().AddDate(0, 0, 30).After(c.Cert.NotAfter) {
			t.SoonExpiredCertificates += 1
		}
	}
}

type LintCertificatesResult struct {
	Issuers               map[string]*LintOrganizationResult
	ExpiredCertificates   int
	UntrustedCertificates int
	LintResult
	*Findings
}

// NewLintCertificatesResult creates new instance of LintCertificatesResult.
func NewLintCertificatesResult() *LintCertificatesResult {
	return &LintCertificatesResult{
		Issuers:  map[string]*LintOrganizationResult{},
		Findings: &Findings{},
	}
}

// AppendCertificate adds certificate test result and updates statuses
func (t *LintCertificatesResult) AppendCertificate(c *LintCertificateResult) {
	issuer := t.Issuers[c.Organization]
	if issuer == nil {
		issuer = &LintOrganizationResult{
			Name:     c.Organization,
			Issues:   map[string]*LintIssue{},
			Findings: &Findings{},
		}
		t.Issuers[c.Organization] = issuer
	}

	skip := false
	if c.IsExpired {
		issuer.ExpiredCertificates += 1
		t.ExpiredCertificates += 1
		skip = true
	}

	if c.IsUntrusted {
		issuer.UntrustedCertificates += 1
		t.UntrustedCertificates += 1
		skip = true
	}

	if skip {
		// exclude untrusted certificates
		return
	}

	issuer.AppendCertificate(c)

	// Update counters
	t.Amount += 1
	if c.Result.ErrorsPresent {
		t.Errors += 1
	}
	if c.Result.WarningsPresent {
		t.Warnings += 1
	}
	if c.Result.NoticesPresent {
		t.Notices += 1
	}
	for _, issue := range c.Result.Results {
		if issue.Status == lint.NE {
			t.NE += 1
			break
		}
	}

	// finding
	t.LeafCertificates += issuer.LeafCertificates
	t.ValidityDays += issuer.ValidityDays
	t.SoonExpiredCertificates += issuer.SoonExpiredCertificates
}

type LintTotalResult struct {
	Issues           map[string]int
	LeafCertificates *LintCertificatesResult
	CaCertificates   *LintCertificatesResult
}

// NewLintTotalResult creates new instance of LintTotalResult.
func NewLintTotalResult() *LintTotalResult {
	return &LintTotalResult{
		Issues:           map[string]int{},
		LeafCertificates: NewLintCertificatesResult(),
		CaCertificates:   NewLintCertificatesResult(),
	}
}

// AppendCertificate adds certificate test result. If tested certificate is CA it adds it into the CaCertificates, otherwise LeafCertificates.
func (t *LintTotalResult) AppendCertificate(r *LintCertificateResult) {
	// update Issues counter
	for code, result := range r.Result.Results {
		if result.Status == lint.Error ||
			result.Status == lint.Warn ||
			result.Status == lint.Notice ||
			result.Status == lint.NE {
			t.Issues[code] += 1
		}
	}

	if r.Cert.IsCA {
		t.CaCertificates.AppendCertificate(r)
		return
	}

	t.LeafCertificates.AppendCertificate(r)
}

// getOrganizationsNames returns ordered list of unique names fro Leaf and CA issuers
func (t *LintTotalResult) getOrganizationsNames() []string {
	nameMap := map[string]bool{}
	// read all names for Leaf certs
	for n := range t.LeafCertificates.Issuers {
		nameMap[n] = true
	}
	// read all names for CA certs
	for n := range t.CaCertificates.Issuers {
		nameMap[n] = true
	}

	res := []string{}
	for n := range nameMap {
		res = append(res, n)
	}

	// sort names
	sort.Slice(res[:], func(i, j int) bool {
		return res[i] < res[j]
	})

	return res
}
