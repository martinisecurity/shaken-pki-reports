package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	"github.com/martinisecurity/shaken-pki-reports/cmd/internal"
	"github.com/zmap/zcrypto/x509"
)

const (
	CA_TRUST_LIST   = "https://wfe.prod.martinisecurity.com/v1/stipa/calist"
	REPORT_DIR_NAME = "x_report"
)

// LintResult is common lint result
type LintResult struct {
	Amount   uint
	Errors   uint
	Warnings uint
	Notices  uint
	NE       uint
}

// Mkdir creates a new dictionary with the specified name.
// It doesn't do anything if directory already exists. Returns error if os.Mkdir has problems.
func Mkdir(name string) error {
	if _, err := os.Stat(name); os.IsNotExist(err) {
		if err := os.Mkdir(name, os.ModePerm); err != nil {
			return fmt.Errorf("cannot create directory %s, %s", name, err.Error())
		}
	}

	return nil
}

// CreateReport creates a directory with README.md file in it and returns point ot the file.
func CreateReport(name string) (*os.File, error) {
	err := Mkdir(name)
	if err != nil {
		return nil, fmt.Errorf("cannot create the report file, %w", err)
	}

	file, err := os.Create(path.Join(name, "README.md"))
	if err != nil {
		return nil, fmt.Errorf("cannot create the report file, %w", err)
	}

	return file, nil
}

// PrintFooter prints common footer
func PrintFooter(w io.Writer) {
	now := time.Now()
	fmt.Fprintln(w, "")
	fmt.Fprintf(w, "Generated: %s at %s", now.Format("02/01/2006"), now.Format("15:04:05"))
}

// ReadRootCertificates reads root certificates from the specified file and returns CertPool.
// If successful returns err == nil, otherwise OS error.
func ReadRootCertificates(fileOrUrl string) (*x509.CertPool, error) {
	var rootPem []byte
	var err error
	p := x509.NewCertPool()

	if strings.HasPrefix(fileOrUrl, "http") {
		// URL
		resp, err := http.Get(fileOrUrl)
		if err != nil {
			return nil, fmt.Errorf("cannot download the list of Root certificates, %w", err)
		}
		if resp.StatusCode != 200 {
			return nil, fmt.Errorf("cannot download the list of Root certificates, HTTP status is not 200 OK")
		}
		defer resp.Body.Close()
		rootPem, err = io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("cannot download the list of Root certificates, %w", err)
		}
	} else {
		// File
		rootPem, err = os.ReadFile(fileOrUrl)
		if err != nil {
			return nil, fmt.Errorf("cannot read the list of Root certificates from the file, %w", err)
		}

	}

	rootCerts := internal.ParseCertificates(rootPem)
	for _, cert := range rootCerts {
		p.AddCert(cert.Certificate)
	}

	return p, nil
}

var wellknownIssueNames = map[string]string{
	"Comcast":                          "Comcast",
	"GBSDTech":                         "GBSDTech",
	"Martini Security, LLC":            "Martini Security",
	"Metaswitch STI-CA SHAKEN Root":    "Metaswitch",
	"NetNumber Inc":                    "NetNumber",
	"Neustar Information Services Inc": "Neustar",
	"Peeringhub Inc":                   "Peeringhub",
	"Sansay Corporation":               "Sansay",
	"TMOBILE-USA":                      "T-Mobile",
	"TransNexus, Inc.":                 "TransNexus",
}

// getOrganizationName returns organization name.
// It tries to get name from the certificate chain, otherwise uses name from the certificate
func getOrganizationName(c *x509.Certificate, options *x509.VerifyOptions) string {
	name := internal.GetOrganizationName(c)
	if options != nil {
		current, expired, never, _ := c.Verify(*options)
		if len(current) > 0 {
			chain := current[0]
			name = internal.GetOrganizationName(chain[len(chain)-1])
		} else if len(expired) > 0 {
			chain := expired[0]
			name = internal.GetOrganizationName(chain[len(chain)-1])
		} else if len(never) > 0 {
			chain := never[0]
			name = internal.GetOrganizationName(chain[len(chain)-1])
		}
	}

	if len(wellknownIssueNames[name]) != 0 {
		name = wellknownIssueNames[name]
	}

	return name
}

func escapeMdLink(link string) string {
	link = strings.Replace(link, " ", "%20", -1)
	return link
}
