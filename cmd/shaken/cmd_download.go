package main

import (
	"crypto"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/martinisecurity/shaken-pki-reports/cmd/internal"
	uLinter "github.com/martinisecurity/shaken-pki-reports/linter"
	"github.com/zmap/zcrypto/x509"
)

// RunDownloadCommand runs the download command. If there is any problem it returns an error.
func RunDownloadCommand(listPath string, outDir string, includeCa bool) error {
	// check params
	if listPath == "" {
		return fmt.Errorf("cannot get file path, variable 'listPath' is empty")
	}

	// read file with URL list
	raw, err := os.ReadFile(listPath)
	if err != nil {
		return fmt.Errorf("cannot read file %s, %s", listPath, err.Error())
	}

	// create folder for certs uploading
	if err := Mkdir(outDir); err != nil {
		return err
	}

	// get root certs
	rootPool, err := ReadRootCertificates(CA_TRUST_LIST)
	if err != nil {
		return err
	}

	links := strings.Split(string(raw), "\n")
	lintResults := NewLintUrlSummaryResult()
	for _, link := range links {
		// use http links only
		if !strings.HasPrefix(link, "http") {
			continue
		}

		// lint URL
		lintResult := uLinter.LintUrl(link)
		if lintResult.StatusCode != 200 {
			lintResults.AppendLink("Unknown", lintResult)
			continue
		}

		// parse certificates and put them into the folder
		certs := internal.ParseCertificates(lintResult.Body)
		files := []string{}
		intermediatePool := x509.NewCertPool()
		var leafCert *x509.Certificate
		for _, pemCert := range certs {
			cert := pemCert.Certificate
			if cert.IsCA {
				intermediatePool.AddCert(cert)
				if !includeCa {
					// skip CA cert if includeCa is false
					continue
				}
			} else {
				leafCert = cert
			}

			// compute cert sha1 thumbprint
			sha1 := crypto.SHA1.New()
			sha1.Write(cert.Raw)

			fileName := fmt.Sprintf("%s.pem", hex.EncodeToString(sha1.Sum(nil)))
			filePath := path.Join(outDir, fileName)
			if _, err := os.Stat(filePath); err == nil {
				// file exists, skip it
				files = append(files, fmt.Sprintf("exists %s", filePath))
				continue
			}

			// write cert into pem file
			certBlock := pem.Block{
				Type: "CERTIFICATE",
				Headers: map[string]string{
					"Link":    link,
					"Subject": cert.Subject.String(),
					"Issuer":  cert.Issuer.String(),
				},
				Bytes: cert.Raw,
			}
			pemCert := pem.EncodeToMemory(&certBlock)
			err = os.WriteFile(filePath, pemCert, 0644)
			if err != nil {
				fmt.Printf("%-7s%s %s\n", "ERROR", link, err.Error())
				continue
			}
			files = append(files, fmt.Sprintf("write  %s", filePath))
		}

		// get Org name or use default Unknown
		orgName := "Unknown"
		if leafCert != nil {
			orgName = getOrganizationName(leafCert, &x509.VerifyOptions{
				Intermediates: intermediatePool,
				Roots:         rootPool,
			})
		}
		lintResults.AppendLink(orgName, lintResult)

		// log status
		fmt.Printf("%-7s%s\n", "OK", link)
		for _, file := range files {
			fmt.Printf("  %s\n", file)
		}
	}

	// make report dir
	if err := Mkdir(REPORT_DIR_NAME); err != nil {
		return err
	}

	// write summary URL report
	file, err := CreateReport(path.Join(REPORT_DIR_NAME, "URL"))
	if err != nil {
		return err
	}
	defer file.Close()
	PrintUrlSummary(file, lintResults)

	for _, lintOrg := range lintResults.Organizations {
		if err := SaveOrgReport(lintOrg, REPORT_DIR_NAME); err != nil {
			return err
		}
	}

	clientsDir := path.Join(REPORT_DIR_NAME, "URL", "CLIENTS")
	if err := Mkdir(clientsDir); err != nil {
		return err
	}
	for _, lintClient := range lintResults.Clients {
		if err := SaveOrgReport(lintClient, clientsDir); err != nil {
			return err
		}
	}

	return nil
}

func SaveOrgReport(r *LintUrlOrgResult, basePath string) error {
	// create org dir
	orgDir := path.Join(basePath, r.Name)
	err := Mkdir(orgDir)
	if err != nil {
		return err
	}

	// create org report file
	orgFile, err := CreateReport(path.Join(orgDir, "URL"))
	if err != nil {
		return fmt.Errorf("cannot create organization report, %s", err.Error())
	}
	defer orgFile.Close()
	PrintUrlOrg(orgFile, r)

	// create issue report file
	issuesDir := path.Join(orgDir, "ISSUES")
	err = Mkdir(issuesDir)
	if err != nil {
		return fmt.Errorf("cannot create issues report, %s", err.Error())
	}
	for code := range r.Problems {
		issueFile, err := CreateReport(path.Join(issuesDir, code))
		if err != nil {
			return fmt.Errorf("cannot create issues report, %s", err.Error())
		}
		defer issueFile.Close()
		PrintUrlIssue(issueFile, code, r)
	}

	return nil
}
