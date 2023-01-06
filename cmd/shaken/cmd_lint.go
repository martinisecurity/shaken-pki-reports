package main

import (
	"encoding/hex"
	"fmt"
	"math/rand"
	"net/url"
	"os"
	"path"
	"strings"
	"time"

	"github.com/martinisecurity/shaken-pki-reports/cmd/internal"
	"github.com/martinisecurity/shaken-pki-reports/repository"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3"
	"github.com/zmap/zlint/v3/lint"

	uLint "github.com/martinisecurity/shaken-pki-reports/lint"
	uLinter "github.com/martinisecurity/shaken-pki-reports/linter"
)

type LintCommandArgs struct {
	OutDir string
	Root   string
	Url    string
}

type CaCrAvailableType int

const (
	Unknown CaCrAvailableType = 0
	Yes     CaCrAvailableType = 1
	No      CaCrAvailableType = 2
	Varies  CaCrAvailableType = 3
)

type CertificateResultSet map[string]*lint.LintResult
type LintCertificateResult struct {
	Timestamp    time.Time
	All          CertificateResultSet
	Problems     CertificateResultSet
	Error        CertificateResultSet
	Warn         CertificateResultSet
	Notice       CertificateResultSet
	NotEffective CertificateResultSet
	Pass         CertificateResultSet
	Others       CertificateResultSet
}

func NewLintCertificateResult() *LintCertificateResult {
	return &LintCertificateResult{
		All:          CertificateResultSet{},
		Problems:     CertificateResultSet{},
		Error:        CertificateResultSet{},
		Warn:         CertificateResultSet{},
		Notice:       CertificateResultSet{},
		NotEffective: CertificateResultSet{},
		Pass:         CertificateResultSet{},
		Others:       CertificateResultSet{},
	}
}

func (t *LintCertificateResult) HasProblems() bool {
	return len(t.Problems) > 0
}

func (t *LintCertificateResult) Append(c string, p *lint.LintResult) {
	t.All[c] = p

	switch p.Status {
	case lint.Error:
		{
			t.Problems[c] = p
			t.Error[c] = p

			break
		}
	case lint.Warn:
		{
			t.Problems[c] = p
			t.Warn[c] = p

			break
		}
	case lint.Notice:
		{
			t.Problems[c] = p
			t.Notice[c] = p

			break
		}
	case lint.NE:
		{
			t.NotEffective[c] = p

			break
		}
	case lint.Pass:
		{
			t.Pass[c] = p

			break
		}
	default:
		{
			t.Others[c] = p

			break
		}
	}
}

type LintCommandItem struct {
	Roots                 *x509.CertPool
	Intermediates         *x509.CertPool
	Url                   *url.URL
	UrlProblems           map[string]*uLint.LintResult
	UrlResult             *uLint.LintResultSet
	Certificate           *x509.Certificate
	CertificateResult     *LintCertificateResult
	IsExpired             bool
	IsUntrusted           bool
	IsDuplicateRepository bool
	Chain                 []*x509.Certificate
	certStatusUpdated     bool
	CaCrAvailable         CaCrAvailableType
	IsExpireSoon          bool
	id                    string
}

type LintCommandItems map[string]*LintCommandItem

func (t *LintCommandItem) SetCertificateResult(cr *zlint.ResultSet) {
	t.CertificateResult = NewLintCertificateResult()
	t.CertificateResult.Timestamp = time.Unix(cr.Timestamp, 0)

	for c, p := range cr.Results {
		t.CertificateResult.Append(c, p)
	}
}

func (t *LintCommandItem) SetUrlResult(ur *uLint.LintResultSet) {
	t.UrlResult = ur
	if t.UrlProblems == nil {
		t.UrlProblems = map[string]*uLint.LintResult{}
	}
	for c, p := range ur.Results {
		if p.Status == uLint.Error ||
			p.Status == uLint.Warn ||
			p.Status == uLint.Notice {
			t.UrlProblems[c] = p
		}
	}
}

func (t *LintCommandItem) Id() string {
	if len(t.id) == 0 {
		randId := make([]byte, 20)
		rand.Read(randId)

		t.id = hex.EncodeToString(randId)
		if t.Url != nil {
			t.id = getRepositoryId(t.Url)
		}
		if t.Certificate != nil {
			t.id = getCertificateId(t.Certificate)
		}
	}

	return t.id
}

func (t *LintCommandItem) IsSkipped() bool {
	return t.IsExpired || t.IsUntrusted || t.IsDuplicateRepository
}

func (t *LintCommandItem) HasUrlProblems() bool {
	return len(t.UrlProblems) > 0
}

func (t *LintCommandItem) UpdateStatuses() {
	// update certificate statuses
	if !t.certStatusUpdated && t.Certificate != nil {
		// build chain
		ch1, ch2, ch3, err := t.Certificate.Verify(x509.VerifyOptions{
			Roots:         t.Roots,
			Intermediates: t.Intermediates,
			CurrentTime:   time.Now(),
		})
		if len(ch1) > 0 {
			t.Chain = ch1[0]
		} else if len(ch2) > 0 {
			t.Chain = ch2[0]
		} else if len(ch3) > 0 {
			t.Chain = ch3[0]
		}

		// check untrusted
		if _, ok := err.(x509.UnknownAuthorityError); ok {
			t.IsUntrusted = true
		}

		// check expired
		if t.Certificate.NotAfter.Before(time.Now()) {
			t.IsExpired = true
		}

		t.certStatusUpdated = true
	}
}

func RunLintCommand(args *LintCommandArgs) error {
	// get urls
	urls, err := ReadURLs(args.Url)
	if err != nil {
		return err
	}

	providers := repository.Providers{
		&repository.MartiniSecurityProvider{},
	}
	for _, provider := range providers {
		repoUrls, err := provider.GetURLs()
		if err != nil {
			fmt.Printf("failed to URLs from the repository provider, %s\n", err.Error())
		}
		urls = append(urls, repoUrls...)
	}

	// get root certs
	rootPool, err := ReadRootCertificates(args.Root)
	if err != nil {
		return err
	}

	icaPool := x509.NewCertPool()
	leafPool := x509.NewCertPool()

	items := []*LintCommandItem{}
	for _, u := range urls {
		item := &LintCommandItem{
			Roots:         rootPool,
			Intermediates: icaPool,
		}
		items = append(items, item)
		item.Url = u

		// lint URL
		timeStart := time.Now()
		fmt.Println("Lint URL")
		fmt.Printf("  URL: %s\n", u.String())
		item.SetUrlResult(uLinter.LintUrl(u))
		fmt.Printf("  Status: %d\n", item.UrlResult.StatusCode)
		timeEnd := time.Now()
		fmt.Printf("  Time: %dms\n", int(timeEnd.Sub(timeStart).Milliseconds()))
		if item.UrlResult.StatusCode != 200 {
			continue
		}

		certs := internal.ParseCertificates(item.UrlResult.Body)
		if len(certs) == 0 {
			continue
		}
		for _, c := range certs {
			if c.IsCA {
				if !c.SelfSigned {
					// ICA
					icaPool.AddCert(c)
				}
			} else {
				// Leaf
				leafPool.AddCert(c)
				item.Certificate = c
				cr, _ := LintCertificate(c)
				if cr != nil {
					// test the leaf certificate
					item.SetCertificateResult(cr)
				}
			}
		}
	}

	// run tests for ICAs
	for _, c := range icaPool.Certificates() {
		cr, _ := LintCertificate(c)
		if cr != nil {
			// test the ICA certificate
			item := &LintCommandItem{
				Roots:         rootPool,
				Intermediates: icaPool,
				Certificate:   c,
			}
			item.SetCertificateResult(cr)
			items = append(items, item)
		}
	}

	// run tests for Roots
	for _, c := range rootPool.Certificates() {
		cr, _ := LintCertificate(c)
		if cr != nil {
			// test the ICA certificate
			item := &LintCommandItem{
				Roots:         rootPool,
				Intermediates: icaPool,
				Certificate:   c,
			}
			item.SetCertificateResult(cr)
			items = append(items, item)
		}
	}

	report := MakeReport(items)

	if err := report.Save(args.OutDir); err != nil {
		return err
	}

	return nil
}

type Report struct {
	Certificates *CertificateSummaryReport
	Repositories *RepositorySummaryReport
}

func NewReport() *Report {
	return &Report{
		Certificates: NewCertificateSummaryReport(),
		Repositories: NewRepositorySummaryReport(),
	}
}

const (
	DIR_CERTS     = "CERTS"
	DIR_ISSUES    = "ISSUES"
	DIR_ISSUERS   = "ISSUERS"
	DIR_REPOS     = "REPOS"
	DIR_CA        = "CA"
	DIR_SP        = "SP"
	DIR_OCN       = "OCN"
	DIR_SKIPPED   = "SKIPPED"
	DIR_UNTRUSTED = "UNTRUSTED"
)

func (t *Report) Save(outDir string) error {
	if err := Mkdir(outDir); err != nil {
		return err
	}

	if err := t.saveCertificates(outDir); err != nil {
		return err
	}

	if err := t.saveRepositories(outDir); err != nil {
		return err
	}

	return nil
}

func (t *Report) saveCertificates(outDir string) error {
	summary, err := CreateReport(outDir)
	if err != nil {
		return err
	}
	defer summary.Close()

	PrintCertificateSummaryReport(summary, t.Certificates)

	if err := SaveExpireSoonCSV(outDir, t.Certificates); err != nil {
		return err
	}

	// Unique OCN CSV
	ocnFile, err := CreateCSV(path.Join(outDir, DIR_OCN))
	if err != nil {
		return err
	}
	defer ocnFile.Close()

	PrintSummaryOCN(ocnFile, t.Certificates)

	// Untrusted CSV
	ocnUntrustedFile, err := CreateCSV(path.Join(outDir, DIR_UNTRUSTED))
	if err != nil {
		return err
	}
	defer ocnUntrustedFile.Close()
	PrintSummaryUntrusted(ocnUntrustedFile, t.Certificates)

	// save issuers
	certsDir := path.Join(outDir, DIR_CERTS)
	if err := Mkdir(certsDir); err != nil {
		return err
	}
	issuers := t.Certificates.GetIssuers()
	for _, issuer := range issuers {
		// issuer report
		issuerDir := path.Join(certsDir, escapeMdLink(issuer.Name))
		issuerFile, err := CreateReport(issuerDir)
		if err != nil {
			return err
		}
		defer issuerFile.Close()

		PrintCertificateIssuerReport(issuerFile, issuer)

		// Skipped
		if len(issuer.Skipped) > 0 {
			skippedDir := path.Join(issuerDir, DIR_SKIPPED)
			skippedFile, err := CreateReport(skippedDir)
			if err != nil {
				return err
			}
			defer skippedFile.Close()
			PrintSkippedReport(skippedFile, issuer)
			for _, v := range issuer.Skipped {
				certsDir := path.Join(issuerDir, DIR_CERTS)
				if err := Mkdir(certsDir); err != nil {
					return err
				}

				certDir := path.Join(certsDir, getCertificateId(v.Certificate))
				certFile, err := CreateReport(certDir)
				if err != nil {
					return err
				}
				defer certFile.Close()

				PrintCertificateReport(certFile, v)
			}
		}

		// Unique OCN CSV
		ocnFile, err := CreateCSV(path.Join(issuerDir, DIR_OCN))
		if err != nil {
			return err
		}
		defer ocnFile.Close()

		if issuer.Leaf != nil {
			PrintIssuerOCN(ocnFile, issuer.Leaf)
		}

		// Untrusted CSV
		if issuer.SkippedUntrustedAmount > 0 {
			untrustedFile, err := CreateCSV(path.Join(issuerDir, DIR_UNTRUSTED))
			if err != nil {
				return err
			}
			defer untrustedFile.Close()

			if issuer.Leaf != nil {
				PrintIssuerUntrusted(untrustedFile, issuer.Leaf)
			}
		}

		if issuer.TestedAmount > 0 {
			// certificate reports
			certsDir := path.Join(issuerDir, DIR_CERTS)
			if err := Mkdir(certsDir); err != nil {
				return err
			}
			for _, v := range issuer.Items {
				if v.Certificate == nil || v.IsExpired || v.IsUntrusted {
					continue
				}
				certDir := path.Join(certsDir, getCertificateId(v.Certificate))
				certFile, err := CreateReport(certDir)
				if err != nil {
					return err
				}
				defer certFile.Close()

				PrintCertificateReport(certFile, v)
			}

			// problems reports
			issuesDir := path.Join(issuerDir, DIR_ISSUES)
			if err := Mkdir(issuesDir); err != nil {
				return err
			}
			for _, p := range issuer.Problems {
				issueDir := path.Join(issuesDir, p.Name)
				issueFile, err := CreateReport(issueDir)
				if err != nil {
					return err
				}
				defer issueFile.Close()

				PrintCertificateIssueReport(issueFile, p.Name, issuer)
			}
		}
	}

	return nil
}

func (t *Report) saveRepositories(outDir string) error {
	reposDir := path.Join(outDir, DIR_REPOS)
	repoFile, err := CreateReport(reposDir)
	if err != nil {
		return err
	}
	defer repoFile.Close()
	PrintRepositorySummaryReport(repoFile, t.Repositories)

	issuesDir := path.Join(reposDir, DIR_ISSUERS)
	if err := Mkdir(issuesDir); err != nil {
		return err
	}
	issuerGroups := []*RepositoryIssuersReport{
		t.Repositories.CA,
		t.Repositories.SP,
	}
	for _, issuerGroup := range issuerGroups {
		dir := DIR_CA
		if issuerGroup.Type == SP {
			dir = DIR_SP
		}

		if issuerGroup.TestedAmount > 0 {
			caDir := path.Join(issuesDir, dir)
			if err := Mkdir(caDir); err != nil {
				return err
			}
			for _, issuer := range issuerGroup.Issuers {
				issuerDir := path.Join(caDir, escapeMdLink(issuer.Name))
				issuerFile, err := CreateReport(issuerDir)
				if err != nil {
					return err
				}
				defer issuerFile.Close()

				PrintRepositoryIssuerReport(issuerFile, issuer)

				if issuer.TestedAmount > 0 {
					issuerReposDir := path.Join(issuerDir, DIR_REPOS)
					if err := Mkdir(issuerReposDir); err != nil {
						return err
					}
					for _, item := range issuer.Items {
						repoDir := path.Join(issuerReposDir, getRepositoryId(item.Url))
						repoFile, err := CreateReport(repoDir)
						if err != nil {
							return err
						}
						defer repoFile.Close()
						PrintRepositoryIssuerUrlReport(repoFile, issuer, item)
					}
				}
				if len(issuer.Problems) > 0 {
					issuerProblemsDir := path.Join(issuerDir, DIR_ISSUES)
					if err := Mkdir(issuerProblemsDir); err != nil {
						return err
					}
					for c, problem := range issuer.Problems {
						problemDir := path.Join(issuerProblemsDir, c)
						problemFile, err := CreateReport(problemDir)
						if err != nil {
							return err
						}
						defer problemFile.Close()

						PrintRepositoryIssuerProblemReport(problemFile, issuer, problem)
					}
				}
			}
		}
	}

	return nil
}

func MakeReport(items []*LintCommandItem) *Report {
	r := NewReport()

	mapItems := LintCommandItems{}
	for _, item := range items {
		item.UpdateStatuses()

		id := item.Id()
		if existingItem := mapItems[id]; existingItem != nil {
			item.IsDuplicateRepository = true
		} else {
			mapItems[id] = item
		}

		if item.Certificate != nil && item.CertificateResult != nil {
			r.Certificates.Append(item)
		}
		if item.Url != nil && item.UrlResult != nil {
			r.Repositories.Append(item)
		}
	}

	return r
}

// ReadURLs reads file with HTTP urls
func ReadURLs(path string) ([]*url.URL, error) {
	// read file with URL list
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("cannot read file with URLs, %w", err)
	}

	links := strings.Split(string(raw), "\n")
	var res []*url.URL
	for _, link := range links {
		// use http links only
		if !strings.HasPrefix(link, "http") {
			continue
		}

		u, err := url.Parse(link)
		if err != nil {
			continue
		}
		res = append(res, u)
	}

	return res, nil
}

var registry lint.Registry

func LintCertificate(c *x509.Certificate) (*zlint.ResultSet, error) {
	fmt.Println("Lint certificate")
	fmt.Printf("  Issuer: %s\n", c.Issuer.String())
	fmt.Printf("  Subject: %s\n", c.Subject.String())
	timeStart := time.Now()
	// Initialize lint registry
	var err error
	if registry == nil {
		registry, err = lint.GlobalRegistry().Filter(lint.FilterOptions{
			// IncludeSources: lint.SourceList{shaken.ShakenPolicy},
			IncludeSources: lint.SourceList{
				lint.RFC5280,
				lint.ATIS1000080,
				lint.UnitedStatesSHAKENCP,
				lint.ShakenPKI,
			},
			ExcludeNames: []string{
				// disable CRL tests for speed increasing
				"e_atis_crl_distribution_not_reachable",
				"e_atis_ca_crl_distribution_not_reachable",

				"w_distribution_point_missing_ldap_or_uri",
			},
		})
		if err != nil {
			return nil, fmt.Errorf("cannot initialize the lint registry, %s", err.Error())
		}
	}

	res := zlint.LintCertificateEx(c, registry)
	timeEnd := time.Now()
	fmt.Printf("  Time: %dms\n", int(timeEnd.Sub(timeStart).Milliseconds()))

	return res, nil
}

func SaveExpireSoonCSV(outDir string, r *CertificateSummaryReport) error {
	file, err := CreateCSV(path.Join(outDir, "EXPIRE_SOON"))
	if err != nil {
		return err
	}
	defer file.Close()

	PrintExpireSoon(file, r.Items)

	return nil
}
