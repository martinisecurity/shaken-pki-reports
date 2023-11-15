package main

import (
	"fmt"
	"time"

	"github.com/martinisecurity/shaken-pki-reports/cmd/internal"
	uLint "github.com/martinisecurity/shaken-pki-reports/lint"
	"github.com/zmap/zlint/v3/lint"
)

type CertificateGroupReport struct {
	Amount                    int
	Items                     []*LintCommandItem
	RepositoryErrorAmount     int
	RepositoryWarnAmount      int
	RepositoryNoticeAmount    int
	TestedAmount              int
	SkippedExpiredAmount      int
	SkippedUntrustedAmount    int
	SkippedRepositoriesAmount int
	ErrorAmount               int
	WarnAmount                int
	NoticeAmount              int
	NotEffectiveAmount        int
	ExpiresSoon               int
	averageRemainingValidity  []int
	averageInitialValidity    []int
	Skipped                   []*LintCommandItem
}

func average(sum int, amount int) float64 {
	if amount == 0 {
		return 0
	}
	return float64(sum) / float64(amount)
}

func averagePercent(sum int, amount int) float64 {
	return average(sum, amount) * 100
}

func (t *CertificateGroupReport) Join(c *CertificateGroupReport) {
	t.Items = append(t.Items, c.Items...)
	t.Amount += c.Amount
	t.RepositoryErrorAmount += c.RepositoryErrorAmount
	t.RepositoryWarnAmount += c.RepositoryWarnAmount
	t.RepositoryNoticeAmount += c.RepositoryNoticeAmount
	t.TestedAmount += c.TestedAmount
	t.SkippedExpiredAmount += c.SkippedExpiredAmount
	t.SkippedUntrustedAmount += c.SkippedUntrustedAmount
	t.SkippedRepositoriesAmount += c.SkippedRepositoriesAmount
	t.ErrorAmount += c.ErrorAmount
	t.WarnAmount += c.WarnAmount
	t.NoticeAmount += c.NoticeAmount
	t.NotEffectiveAmount += c.NotEffectiveAmount
	t.ExpiresSoon += c.ExpiresSoon
	t.averageRemainingValidity = append(t.averageRemainingValidity, c.averageRemainingValidity...)
	t.averageInitialValidity = append(t.averageInitialValidity, c.averageInitialValidity...)
	t.Skipped = append(t.Skipped, c.Skipped...)
}

func (t *CertificateGroupReport) AverageCertificatesWithProblems() float64 {
	sum := 0
	count := 0
	for _, v := range t.Items {
		if v.CertificateResult.HasProblems() {
			count += 1
			sum += len(v.CertificateResult.Problems)
		}
	}

	return average(sum, count)
}

func (t *CertificateGroupReport) AverageRemainingValidity() float64 {
	sum := 0
	for _, v := range t.averageRemainingValidity {
		sum += v
	}

	return average(sum, len(t.averageRemainingValidity))
}

func (t *CertificateGroupReport) AverageInitialValidity() float64 {
	sum := 0
	for _, v := range t.averageInitialValidity {
		sum += v
	}

	return average(sum, len(t.averageInitialValidity))
}

func (t *CertificateGroupReport) AverageRepositoryErrors() float64 {
	return averagePercent(t.RepositoryErrorAmount, t.TestedAmount)
}

func (t *CertificateGroupReport) AverageRepositoryWarns() float64 {
	return averagePercent(t.RepositoryWarnAmount, t.TestedAmount)
}

func (t *CertificateGroupReport) AverageRepositoryNotices() float64 {
	return averagePercent(t.RepositoryNoticeAmount, t.TestedAmount)
}

func (t *CertificateGroupReport) AverageErrors() float64 {
	return averagePercent(t.ErrorAmount, t.TestedAmount)
}

func (t *CertificateGroupReport) AverageWarns() float64 {
	return averagePercent(t.WarnAmount, t.TestedAmount)
}

func (t *CertificateGroupReport) AverageNotices() float64 {
	return averagePercent(t.NoticeAmount, t.TestedAmount)
}

func (t *CertificateGroupReport) AverageNotEffective() float64 {
	return averagePercent(t.NotEffectiveAmount, t.TestedAmount)
}

func (t *CertificateGroupReport) Append(i *LintCommandItem) bool {
	if i == nil || i.Certificate == nil || i.CertificateResult == nil {
		return false
	}

	t.Amount += 1

	i.UpdateStatuses()

	skipped := false
	if i.IsDuplicateRepository {
		t.SkippedRepositoriesAmount += 1
		skipped = true
	} else if i.IsUntrusted {
		t.SkippedUntrustedAmount += 1
		skipped = true
	} else if i.IsExpired {
		t.SkippedExpiredAmount += 1
		skipped = true
	}
	if skipped {
		// Append to skipped
		t.Skipped = append(t.Skipped, i)
		return true
	}

	t.Items = append(t.Items, i)
	t.TestedAmount += 1

	// update dates
	//   initial
	initialValidity := internal.GetValidityDays(i.Certificate)
	t.averageInitialValidity = append(t.averageInitialValidity, initialValidity)
	//   remaining
	remainingValidity := internal.GetRemainingDays(i.Certificate, time.Now())
	t.averageRemainingValidity = append(t.averageInitialValidity, remainingValidity)

	if time.Now().AddDate(0, 0, 30).After(i.Certificate.NotAfter) {
		i.IsExpireSoon = true
		t.ExpiresSoon += 1
	}

	// refresh status counters
	cr := i.CertificateResult

	if ur := i.UrlResult; ur != nil {
		if ur.HasErrors {
			t.RepositoryErrorAmount += 1
		}
		if ur.HasWarnings {
			t.RepositoryWarnAmount += 1
		}
		if ur.HasNotices {
			t.RepositoryNoticeAmount += 1
		}
	}

	if len(cr.Error) > 0 {
		t.ErrorAmount += 1
	}
	if len(cr.Warn) > 0 {
		t.WarnAmount += 1
	}
	if len(cr.Notice) > 0 {
		t.NoticeAmount += 1
	}
	if len(cr.NotEffective) > 0 {
		t.NotEffectiveAmount += 1
	}

	return true
}

type Problem struct {
	Name   string
	Items  []*LintCommandItem
	Source string
}

type CertificateIssuerReport struct {
	Name       string
	Problems   map[string]*Problem
	UniqueOCNs map[string][]*LintCommandItem
	CertificateGroupReport
}

func NewCertificateIssuerReport() *CertificateIssuerReport {
	return &CertificateIssuerReport{
		Problems:   map[string]*Problem{},
		UniqueOCNs: map[string][]*LintCommandItem{},
	}
}

func (t *CertificateIssuerReport) GetAverageUniqueOCN() float64 {
	uniqueSum := 0
	uniqueAmount := 0
	for _, v := range t.UniqueOCNs {
		uniqueSum += len(v)
		uniqueAmount += 1
	}

	if uniqueAmount == 0 {
		return 0
	}

	return float64(uniqueSum) / float64(uniqueAmount)
}

func (t *CertificateIssuerReport) Append(i *LintCommandItem) bool {
	if !t.CertificateGroupReport.Append(i) {
		return false
	}

	if i.IsSkipped() {
		return true
	}

	// append problems
	problemsCounter := 0
	for code := range i.CertificateResult.Problems {
		problem := t.Problems[code]
		if problem == nil {
			l := lint.GlobalRegistry().ByName(code)
			problem = &Problem{
				Name:   code,
				Source: string(l.Source),
			}
			t.Problems[code] = problem
		}
		problem.Items = append(problem.Items, i)
		problemsCounter += 1
	}

	// Update Unique OCNs
	uniqueName := internal.GetUniqueOCN(i.Certificate)
	t.UniqueOCNs[uniqueName] = append(t.UniqueOCNs[uniqueName], i)

	return true
}

type CertificateIssuersReport struct {
	Issuers    map[string]*CertificateIssuerReport
	UniqueOCNs map[string]*CertificateIssuerReport
	CertificateGroupReport
}

func NewCertificateIssuersReport() *CertificateIssuersReport {
	return &CertificateIssuersReport{
		Issuers:    map[string]*CertificateIssuerReport{},
		UniqueOCNs: map[string]*CertificateIssuerReport{},
	}
}

func (t *CertificateIssuersReport) GetAverageUniqueOCN() float64 {
	uniqueSum := 0
	uniqueAmount := 0
	for _, v := range t.UniqueOCNs {
		uniqueSum += v.Amount
		uniqueAmount += 1
	}

	if uniqueAmount == 0 {
		return 0
	}

	return float64(uniqueSum) / float64(uniqueAmount)
}

func (t *CertificateIssuersReport) Append(i *LintCommandItem) bool {
	if !t.CertificateGroupReport.Append(i) {
		return false
	}

	// issuerName
	issuerName := internal.GetOrganizationName(i.Certificate)
	if len(i.Chain) > 0 {
		issuerName = internal.GetOrganizationName(i.Chain[len(i.Chain)-1])
	}
	if len(wellknownIssueNames[issuerName]) != 0 {
		issuerName = wellknownIssueNames[issuerName]
	}

	name := t.Issuers[issuerName]
	if name == nil {
		name = NewCertificateIssuerReport()
		name.Name = issuerName
		t.Issuers[issuerName] = name
	}
	name.Append(i)

	// Update Unique OCNs
	uniqueName := internal.GetUniqueOCN(i.Certificate)
	uniqueOCN := t.UniqueOCNs[uniqueName]
	if uniqueOCN == nil {
		uniqueOCN = NewCertificateIssuerReport()
		uniqueOCN.Name = uniqueName
		t.UniqueOCNs[uniqueName] = uniqueOCN
	}
	uniqueOCN.Append(i)

	return true
}

type CertificateSummaryReport struct {
	Leaf *CertificateIssuersReport
	CA   *CertificateIssuersReport
	CertificateGroupReport
}

func NewCertificateSummaryReport() *CertificateSummaryReport {
	return &CertificateSummaryReport{
		Leaf: NewCertificateIssuersReport(),
		CA:   NewCertificateIssuersReport(),
	}
}

func (t *CertificateSummaryReport) Append(i *LintCommandItem) bool {
	if !t.CertificateGroupReport.Append(i) {
		return false
	}

	issuers := t.Leaf
	if i.Certificate.IsCA {
		issuers = t.CA
	}
	issuers.Append(i)

	return true
}

func (t *CertificateSummaryReport) GetIssuers() map[string]*CertificateIssuerJoin {
	names := map[string]bool{}

	groups := []*CertificateIssuersReport{
		t.Leaf,
		t.CA,
	}
	for _, group := range groups {
		for name := range group.Issuers {
			names[name] = false
		}
	}

	res := map[string]*CertificateIssuerJoin{}
	for name := range names {
		leaf := t.Leaf.Issuers[name]
		ca := t.CA.Issuers[name]

		issuerJoin := &CertificateIssuerJoin{
			Name:     name,
			Leaf:     leaf,
			CA:       ca,
			Problems: map[string]*Problem{},
		}
		res[name] = issuerJoin
		groups := []*CertificateIssuerReport{
			leaf,
			ca,
		}
		for _, group := range groups {
			if group == nil {
				continue
			}
			issuerJoin.Join(group)
		}
	}

	return res
}

type CertificateIssuerJoin struct {
	Name     string
	Leaf     *CertificateIssuerReport
	CA       *CertificateIssuerReport
	Problems map[string]*Problem
	CertificateGroupReport
}

func (t *CertificateIssuerJoin) Join(r *CertificateIssuerReport) {
	t.CertificateGroupReport.Join(&r.CertificateGroupReport)

	for k, v := range r.Problems {
		problem := t.Problems[k]
		if problem == nil {
			l := lint.GlobalRegistry().ByName(k)
			problem = &Problem{
				Name:   v.Name,
				Source: string(l.Source),
			}
			t.Problems[k] = problem
		}
		problem.Items = append(problem.Items, v.Items...)
	}
}

type RepositoryGroupReport struct {
	Items         []*LintCommandItem
	Amount        int
	TestedAmount  int
	SkippedAmount int
	ErrorAmount   int
	WarnAmount    int
	NoticeAmount  int
	averageTime   []int
}

func (t *RepositoryGroupReport) AverageUrlsWithProblems() float64 {
	sum := 0
	count := 0
	for _, v := range t.Items {
		if v.HasUrlProblems() {
			count += 1
			sum += len(v.UrlProblems)
		}
	}

	return average(sum, count)
}

func (t *RepositoryGroupReport) Append(v *LintCommandItem) bool {
	if v.Url == nil || v.UrlResult == nil {
		return false
	}

	t.Amount += 1

	if v.IsDuplicateRepository {
		t.SkippedAmount += 1
	} else {
		t.Items = append(t.Items, v)
		t.TestedAmount += 1

		if v.UrlResult.HasErrors {
			t.ErrorAmount += 1
		}
		if v.UrlResult.HasWarnings {
			t.WarnAmount += 1
		}
		if v.UrlResult.HasNotices {
			t.NoticeAmount += 1
		}

		t.averageTime = append(t.averageTime, v.UrlResult.Time)
	}

	return true
}

func (t *RepositoryGroupReport) AverageErrors() float64 {
	return averagePercent(t.ErrorAmount, t.TestedAmount)
}

func (t *RepositoryGroupReport) AverageWarns() float64 {
	return averagePercent(t.WarnAmount, t.TestedAmount)
}

func (t *RepositoryGroupReport) AverageNotices() float64 {
	return averagePercent(t.NoticeAmount, t.TestedAmount)
}

func (t *RepositoryGroupReport) AverageTime() float64 {
	sum := 0
	for _, v := range t.averageTime {
		sum += v
	}

	return average(sum, len(t.averageTime))
}

type RepositoryIssuerReport struct {
	Name     string
	Problems map[string]*Problem
	RepositoryGroupReport
}

func NewRepositoryIssuerReport() *RepositoryIssuerReport {
	return &RepositoryIssuerReport{
		Problems: map[string]*Problem{},
	}
}

func (t *RepositoryIssuerReport) Append(v *LintCommandItem) bool {
	if !t.RepositoryGroupReport.Append(v) {
		return false
	}

	if !v.IsDuplicateRepository {
		for c, r := range v.UrlResult.Results {
			if r.Status == uLint.Error ||
				r.Status == uLint.Warn ||
				r.Status == uLint.Notice {
				p := t.Problems[c]
				if p == nil {
					l := uLint.FindRuleByName(c)
					if l == nil {
						panic(fmt.Sprintf("rule %s not found", c))
					}
					p = &Problem{
						Name:   c,
						Source: string(l.Source),
					}
					t.Problems[c] = p
				}
				p.Items = append(p.Items, v)
			}
		}
	}

	return true
}

type RepositoryIssuersType int

const (
	CA RepositoryIssuersType = 0
	SP RepositoryIssuersType = 1
)

type RepositoryIssuers map[string]*RepositoryIssuerReport

type RepositoryIssuersReport struct {
	Type    RepositoryIssuersType
	Issuers RepositoryIssuers
	RepositoryGroupReport
}

func NewRepositoryIssuersReport(t RepositoryIssuersType) *RepositoryIssuersReport {
	return &RepositoryIssuersReport{
		Type:    t,
		Issuers: map[string]*RepositoryIssuerReport{},
	}
}

func (t *RepositoryIssuersReport) Append(v *LintCommandItem) bool {
	if !t.RepositoryGroupReport.Append(v) {
		return false
	}

	// get name
	name := "Unknown"
	if t.Type == CA {
		name = wellknownCaDomains[v.Url.Hostname()]
	} else {
		if v.UrlResult.StatusCode == 200 {
			if certs := internal.ParseCertificates(v.UrlResult.Body); len(certs) > 0 {
				cert := certs[0]
				if len(cert.Subject.Organization) > 0 {
					// use O name from the Subject
					name = cert.Subject.Organization[0]
				} else if len(cert.Issuer.Organization) > 0 {
					// use O name from the Issuer
					name = cert.Issuer.Organization[0]
				} else if spc, err := internal.GetTNEntrySPC(cert); err == nil {
					// use SPC value from the TNAuthList extension
					name = fmt.Sprintf("SHAKEN %s", spc)
				}
			}
		}
	}

	// get or create issuer
	issuer := t.Issuers[name]
	if issuer == nil {
		issuer = NewRepositoryIssuerReport()
		issuer.Name = name
		t.Issuers[name] = issuer
	}
	issuer.Append(v)

	return true
}

type RepositorySummaryReport struct {
	// CA contains summary for Certificate Authorities
	CA *RepositoryIssuersReport
	// SP contains summary for Service Providers
	SP *RepositoryIssuersReport
	RepositoryGroupReport
}

func NewRepositorySummaryReport() *RepositorySummaryReport {
	return &RepositorySummaryReport{
		CA: NewRepositoryIssuersReport(CA),
		SP: NewRepositoryIssuersReport(SP),
	}
}

func (t *RepositorySummaryReport) Append(v *LintCommandItem) bool {
	if !t.RepositoryGroupReport.Append(v) {
		return false
	}

	if len(wellknownCaDomains[v.Url.Hostname()]) > 0 {
		// CA
		t.CA.Append(v)
	} else {
		// SP
		t.SP.Append(v)
	}

	return true
}
