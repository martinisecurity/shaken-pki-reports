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

	links := strings.Split(string(raw), "\n")
	lintResults := NewLintUrlSummaryResult()
	for _, link := range links {
		// use http links only
		if !strings.HasPrefix(link, "http") {
			continue
		}

		// lint URL
		lintResult := uLinter.LintUrl(link)
		lintResults.AppendLink(lintResult)
		if lintResult.StatusCode != 200 {
			continue
		}

		// parse certificates and put them into the folder
		certs := internal.ParseCertificates(lintResult.Body)
		files := []string{}
		for _, pemCert := range certs {
			cert := pemCert.Certificate

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

	// create ORGS dir
	orgsDir := path.Join(REPORT_DIR_NAME, "URL", "ORGS")
	if err := Mkdir(orgsDir); err != nil {
		return err
	}

	for groupName, group := range lintResults.Groups {
		groupDir := path.Join(orgsDir, groupName)
		if err := Mkdir(groupDir); err != nil {
			return err
		}
		for _, org := range group.Items {
			if err := SaveOrgReport(org, groupDir); err != nil {
				return err
			}
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
	orgFile, err := CreateReport(orgDir)
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
