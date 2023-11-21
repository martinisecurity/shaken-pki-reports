package main

import (
	"encoding/pem"
	"fmt"
	"io"
	"net/http"

	"github.com/zmap/zcrypto/encoding/asn1"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zcrypto/x509/pkix"
)

type RevocationStatus int

const (
	RevocationStatusUnknown RevocationStatus = iota
	RevocationStatusGood
	RevocationStatusRevoked
)

func CheckChainRevocation(chain []*x509.Certificate) ([]RevocationStatus, error) {
	statuses := []RevocationStatus{}
	for _, cert := range chain {
		status, err := CheckCertRevocation(cert)
		if err != nil {
			return nil, err
		}
		statuses = append(statuses, status)
	}

	return statuses, nil
}

func CheckCertRevocation(cert *x509.Certificate) (RevocationStatus, error) {
	// check CRL
	for _, crlUrl := range cert.CRLDistributionPoints {
		crl, err := GetCRL(crlUrl)
		if err != nil {
			return RevocationStatusUnknown, err
		}

		status := CheckRevocationStatus(cert, crl)
		if status != RevocationStatusUnknown {
			return status, nil
		}
	}

	return RevocationStatusUnknown, nil
}

func GetCRL(crlUrl string) (*pkix.CertificateList, error) {
	// Set proxy URL
	proxyUrl := "https://wfe.dev.martinisecurity.com/v1/stipa/proxy?url=" + crlUrl

	// Send direct request
	crl, err := sendRequest(crlUrl)
	if err != nil {
		// Send request via proxy if direct request failed
		crl, err = sendRequest(proxyUrl)
		if err != nil {
			return nil, err
		}
	}

	return crl, nil
}

func sendRequest(url string) (*pkix.CertificateList, error) {
	// Create request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request to %s, %s", url, err.Error())
	}

	// Send request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request to %s, %s", url, err.Error())
	}

	// Check response status code
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error response status code %d", resp.StatusCode)
	}

	// Read response body
	var crl *pkix.CertificateList
	rawData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body, %s", err.Error())
	}
	pemBlock, _ := pem.Decode(rawData)
	if pemBlock == nil {
		crl, err = x509.ParseCRL(rawData)
		if err != nil {
			return nil, fmt.Errorf("error parsing CRL from DER, %s", err.Error())
		}
	} else {
		crl, err = x509.ParseCRL(pemBlock.Bytes)
		if err != nil {
			return nil, fmt.Errorf("error parsing CRL from PEM, %s", err.Error())
		}
	}

	return crl, nil
}

var OidExtensionDeltaCRLIndicator = asn1.ObjectIdentifier{2, 5, 29, 27}

func FindDeltaCrl(crl *pkix.CertificateList) (*pkix.CertificateList, error) {
	var deltaCrlUrl string
	for _, ext := range crl.TBSCertList.Extensions {
		if ext.Id.Equal(OidExtensionDeltaCRLIndicator) {
			deltaCrlUrl = string(ext.Value)
			break
		}
	}

	if deltaCrlUrl == "" {
		return nil, fmt.Errorf("no delta CRL found")
	}

	return GetCRL(deltaCrlUrl)
}

func HasDeltaCrl(crl *pkix.CertificateList) bool {
	var deltaExt *pkix.Extension
	for _, ext := range crl.TBSCertList.Extensions {
		if ext.Id.Equal(OidExtensionDeltaCRLIndicator) {
			deltaExt = &ext
			break
		}
	}

	return deltaExt != nil
}

func CheckRevocationStatus(cert *x509.Certificate, crl *pkix.CertificateList) RevocationStatus {
	if crl == nil {
		return RevocationStatusUnknown
	}

	revokedCerts := crl.TBSCertList.RevokedCertificates
	for _, revokedCert := range revokedCerts {
		if revokedCert.SerialNumber.Cmp(cert.SerialNumber) == 0 {
			return RevocationStatusRevoked
		}
	}

	return RevocationStatusGood
}
