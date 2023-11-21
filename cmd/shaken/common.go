package main

import (
	"crypto"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"net/url"
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
	Name     string
	Amount   uint
	Errors   uint
	Warnings uint
	Notices  uint
	NE       uint
}

// Mkdir creates a new dictionary with the specified name.
// It doesn't do anything if directory already exists. Returns error if os.Mkdir has problems.
func Mkdir(name string) error {
	name = escapeMdLink(name)
	if _, err := os.Stat(name); os.IsNotExist(err) {
		if err := os.Mkdir(name, os.ModePerm); err != nil {
			return fmt.Errorf("cannot create directory %s, %s", name, err.Error())
		}
	}

	return nil
}

// CreateReport creates a directory with README.md file in it and returns point ot the file.
func CreateReport(name string) (*os.File, error) {
	name = escapeMdLink(name)

	err := Mkdir(name)
	if err != nil {
		return nil, fmt.Errorf("cannot create the report directory, %w", err)
	}

	file, err := os.Create(path.Join(name, "README.md"))
	if err != nil {
		return nil, fmt.Errorf("cannot create the report file, %w", err)
	}

	return file, nil
}

func CreateCSV(name string) (*os.File, error) {
	file, err := os.Create(fmt.Sprintf("%s.csv", name))
	if err != nil {
		return nil, fmt.Errorf("cannot create the CSV file, %w", err)
	}

	return file, nil
}

// PrintFooter prints common footer
func PrintFooter(w io.Writer) {
	now := time.Now()
	fmt.Fprintln(w, "")
	fmt.Fprintf(w, "Generated: %s", now.Format(time.RFC822))
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
		p.AddCert(cert)
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
	"Telonium":                         "Telonium",
	"Telonium Communications LLC":      "Telonium",
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

func escapeMdLink(link string) string {
	// replace all spaces with underscores
	link = strings.Replace(link, " ", "_", -1)
	return link
}

func escapeName(name string) string {
	// replace all spaces with underscores
	name = strings.Replace(name, " ", "_", -1)
	// replace all dots with underscores
	name = strings.Replace(name, ".", "_", -1)
	// replace all commas with underscores
	name = strings.Replace(name, ",", "_", -1)
	// replace all slashes with underscores
	name = strings.Replace(name, "/", "", -1)
	return name
}

func getCertificateId(c *x509.Certificate) string {
	return hex.EncodeToString(c.FingerprintSHA256)
}

func getRepositoryId(u *url.URL) string {
	res := []byte{}
	digest := crypto.SHA1.New()
	digest.Write([]byte(u.String()))

	return hex.EncodeToString(digest.Sum(res))
}
