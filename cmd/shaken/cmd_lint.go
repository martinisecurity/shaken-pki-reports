package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/martinisecurity/shaken-pki-reports/cmd/internal"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3"
	"github.com/zmap/zlint/v3/lint"
)

func RunLintCommand(certPath string, summary bool) error {
	if certPath == "" {
		return fmt.Errorf("cannot get certificate path, variable 'certPath' is empty")
	}

	// This returns an *os.FileInfo type
	file, err := os.Open(certPath)
	if err != nil {
		return fmt.Errorf("cannot open %s. %s", certPath, err.Error())
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return fmt.Errorf("cannot get stat info from the %s. %w", certPath, err)
	}

	if fileInfo.IsDir() {
		// load root certs
		rootPool, err := ReadRootCertificates(CA_TRUST_LIST)
		if err != nil {
			return fmt.Errorf("cannot load root certificates, %w", err)
		}
		checkCerts := []*internal.PemCertificate{}
		intermediatePool := x509.NewCertPool()
		rootCerts := rootPool.Certificates()
		for _, cert := range rootCerts {
			checkCerts = append(checkCerts, &internal.PemCertificate{
				Certificate: cert,
			})
		}

		files, err := ioutil.ReadDir(certPath)
		if err != nil {
			return fmt.Errorf("cannot read the directory %s, %w", certPath, err)
		}

		for _, file := range files {
			if !file.IsDir() {
				certPath := path.Join(certPath, file.Name())
				certRaw, err := os.ReadFile(certPath)
				if err != nil {
					continue
				}
				certs := internal.ParseCertificates(certRaw)
				for _, cert := range certs {
					if cert.Certificate.IsCA {
						// add intermediate certs into pool
						intermediatePool.AddCert(cert.Certificate)
					}
					// add leaf and intermediate certs into check list
					checkCerts = append(checkCerts, cert)
				}
			}
		}

		r := LintCertificates(checkCerts, &x509.VerifyOptions{
			Roots:         rootPool,
			Intermediates: intermediatePool,
		})

		if err := Mkdir(REPORT_DIR_NAME); err != nil {
			return err
		}

		err = SaveTotalReport(r, REPORT_DIR_NAME)
		if err != nil {
			return err
		}
		err = SaveOrganizationReport(r, REPORT_DIR_NAME)
		if err != nil {
			return err
		}
		err = SaveCertificatesReport(r, REPORT_DIR_NAME)
		if err != nil {
			return err
		}
	} else {
		// file is not a directory

		raw, err := os.ReadFile(certPath)
		if err != nil {
			return fmt.Errorf("cannot read the file %s, %s", certPath, err.Error())
		}

		certs := internal.ParseCertificates(raw)
		if len(certs) == 0 {
			return fmt.Errorf("cannot read the certificate from the file %s, %s", certPath, err.Error())
		}

		result, err := LintCertificate(certs[0], nil)
		if err != nil {
			return fmt.Errorf("cannot lint the certificate, %w", err)
		}

		counter := map[lint.LintSource]int{}
		for code, result := range result.Result.Results {
			if result.Status == lint.Pass || result.Status == lint.Error || result.Status == lint.Warn || result.Status == lint.Notice || result.Status == lint.NE {
				rule := lint.GlobalRegistry().ByName(code)
				counter[rule.Source] += 1
			}
		}
		for source, v := range counter {
			fmt.Printf("%s: %d\n", source, v)
		}

		PrintCertificateReport(os.Stdout, result)
	}

	return nil
}

func ReadCertificatesDir(dirPath string) ([]*internal.PemCertificate, error) {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, fmt.Errorf("cannot read the '%s' directory, %w", dirPath, err)
	}

	res := []*internal.PemCertificate{}

	for _, file := range files {
		if !file.IsDir() {
			certPath := path.Join(dirPath, file.Name())
			certRaw, err := os.ReadFile(certPath)
			if err != nil {
				continue
			}

			certs := internal.ParseCertificates(certRaw)
			res = append(res, certs...)
		}
	}

	return res, nil
}

func LintCertificates(certs []*internal.PemCertificate, options *x509.VerifyOptions) *LintTotalResult {
	res := NewLintTotalResult()

	for _, cert := range certs {
		certRes, err := LintCertificate(cert, options)
		if err != nil {
			continue
		}

		res.AppendCertificate(certRes)
	}

	return res
}

func LintCertificate(cert *internal.PemCertificate, options *x509.VerifyOptions) (*LintCertificateResult, error) {
	// Initialize lint registry
	registry, err := lint.GlobalRegistry().Filter(lint.FilterOptions{
		// IncludeSources: lint.SourceList{shaken.ShakenPolicy},
		IncludeSources: lint.SourceList{
			lint.RFC5280,
			lint.ATIS1000080,
			lint.CPv1_3,
			lint.ShakenPKI,
		},
		ExcludeNames: []string{
			"w_distribution_point_missing_ldap_or_uri",
		},
	})
	if err != nil {
		return nil, fmt.Errorf("cannot initialize the lint registry, %s", err.Error())
	}

	link := cert.Headers["Link"]
	if len(link) == 0 {
		link = ""
	}

	organization := getOrganizationName(cert.Certificate, options)

	thumbprint := computeCertThumbprint(cert.Certificate)
	fmt.Printf("Lint certificate %s issued by '%s' (%s)\n", thumbprint, organization, link)

	return &LintCertificateResult{
		Link:         link,
		Cert:         cert.Certificate,
		Thumbprint:   thumbprint,
		Result:       zlint.LintCertificateEx(cert.Certificate, registry),
		Organization: organization,
	}, nil
}

// SaveOrganizationReport writes report for each organization
func SaveOrganizationReport(r *LintTotalResult, outDir string) error {

	names := r.getOrganizationsNames()
	for _, name := range names {
		// create folder
		orgDir := path.Join(outDir, name)
		if err := Mkdir(orgDir); err != nil {
			return err
		}

		// create file
		orgFile := path.Join(orgDir, "README.md")
		file, err := os.Create(orgFile)
		if err != nil {
			return fmt.Errorf("cannot create %s, %w", orgFile, err)
		}
		defer file.Close()

		PrintOrganizationReport(file, name, r)

		for code := range r.Issues {
			if hasIssue(name, code, r) {
				SaveIssueGroupReport(name, code, r, orgDir)
			}
		}
	}

	return nil
}

// hasIssue returns true if LeafCertificates or CaCertificates for selected orgName contains issueName and it has problems
func hasIssue(orgName string, issueName string, r *LintTotalResult) bool {
	certsSet := []*LintCertificatesResult{
		r.LeafCertificates,
		r.CaCertificates,
	}
	for _, certs := range certsSet {
		issuer := certs.Issuers[orgName]
		if issuer != nil {
			issue := issuer.Issues[issueName]
			if issue != nil {
				return true
			}
		}
	}

	return false
}

func SaveIssueGroupReport(orgName string, issueName string, r *LintTotalResult, orgDir string) error {
	// check if Leaf and CA has this issue

	issuesDir := path.Join(orgDir, "ISSUES")
	err := Mkdir(issuesDir)
	if err != nil {
		return fmt.Errorf("cannot save IssueGroup report, %w", err)
	}

	reportFile := path.Join(issuesDir, fmt.Sprintf("%s.md", issueName))
	file, err := os.Create(reportFile)
	if err != nil {
		return fmt.Errorf("cannot save IssueGroup report, %w", err)
	}
	defer file.Close()

	PrintIssueGroupReport(file, orgName, issueName, r)

	return nil
}

func SaveTotalReport(r *LintTotalResult, outDir string) error {
	file, err := os.Create(path.Join(outDir, "README.md"))
	if err != nil {
		return fmt.Errorf("cannot save README.md, %w", err)
	}
	defer file.Close()

	PrintTotalReport(file, r)

	return nil
}

func SaveCertificatesReport(r *LintTotalResult, outDir string) error {
	names := r.getOrganizationsNames()
	for _, name := range names {
		issuers := []*LintOrganizationResult{
			r.LeafCertificates.Issuers[name],
			r.CaCertificates.Issuers[name],
		}
		for _, issuer := range issuers {
			if issuer == nil {
				continue
			}
			for _, cert := range issuer.Certificates {
				// create folder
				certDir := path.Join(outDir, name, cert.Thumbprint)
				if err := Mkdir(certDir); err != nil {
					return fmt.Errorf("cannot create directory %s, %s", certDir, err.Error())
				}

				// create file
				certFile := path.Join(path.Join(certDir, "README.md"))
				file, err := os.Create(certFile)
				if err != nil {
					return fmt.Errorf("cannot create %s, %w", certFile, err)
				}
				defer file.Close()

				PrintCertificateReportForIssuer(file, name, cert)
			}
		}
	}

	return nil
}
