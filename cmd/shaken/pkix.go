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

type CRLFetcher interface {
	GetCRL(crlUrl string) (*pkix.CertificateList, error)
}

type DefaultCRLFetcher struct{}

func (f *DefaultCRLFetcher) GetCRL(crlUrl string) (*pkix.CertificateList, error) {
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

// CertificateRevocationStatus represents the revocation status of a certificate.
type CertificateRevocationStatus struct {
	Certificate *x509.Certificate
	Status      RevocationStatus
}

// CheckChainRevocation checks the revocation status of a certificate chain.
// It takes a chain of x509 certificates and a CRLFetcher interface as input.
// The function returns a slice of CertificateRevocationStatus and an error.
// Each CertificateRevocationStatus in the slice represents the revocation status of a certificate in the chain.
// The CRLFetcher interface is used to fetch the Certificate Revocation List (CRL) for each certificate in the chain.
// If any error occurs during the process, the function returns nil for the statuses slice and the error.
func CheckChainRevocation(chain []*x509.Certificate, fetcher CRLFetcher) ([]CertificateRevocationStatus, error) {
	var statuses []CertificateRevocationStatus

	for _, cert := range chain {
		status, err := CheckCertRevocationWithFetcher(cert, fetcher)
		if err != nil {
			return nil, err
		}

		statuses = append(statuses, CertificateRevocationStatus{
			Certificate: cert,
			Status:      status,
		})
	}

	return statuses, nil
}

func CheckCertRevocation(cert *x509.Certificate) (RevocationStatus, error) {
	return CheckCertRevocationWithFetcher(cert, nil)
}

func CheckCertRevocationWithFetcher(cert *x509.Certificate, fetcher CRLFetcher) (RevocationStatus, error) {
	if fetcher == nil {
		fetcher = &DefaultCRLFetcher{}
	}

	// check CRL
	for _, crlUrl := range cert.CRLDistributionPoints {
		crl, err := fetcher.GetCRL(crlUrl)
		if err != nil {
			return RevocationStatusUnknown, err
		}

		status := CheckRevocationStatus(cert, crl)
		if status == RevocationStatusRevoked {
			return status, nil
		}

		// Check for delta CRL
		if HasDeltaCrl(crl) {
			deltaCrl, err := FindDeltaCrl(crl, fetcher)
			if err != nil {
				return RevocationStatusUnknown, err
			}
			status = CheckRevocationStatus(cert, deltaCrl)
			if status == RevocationStatusRevoked {
				return status, nil
			}
		}
	}

	return RevocationStatusGood, nil
}

// OidExtensionFreshestCRL is the OID for the freshest CRL extension.
var OidExtensionFreshestCRL = asn1.ObjectIdentifier{2, 5, 29, 46}

// FindDeltaCrl finds and fetches the delta CRL for a given CRL using the provided CRLFetcher.
func FindDeltaCrl(crl *pkix.CertificateList, fetcher CRLFetcher) (*pkix.CertificateList, error) {
	var deltaCrlUrls []string
	for _, ext := range crl.TBSCertList.Extensions {
		if ext.Id.Equal(OidExtensionFreshestCRL) {
			_, err := asn1.Unmarshal(ext.Value, &deltaCrlUrls)
			if err != nil {
				return nil, fmt.Errorf("failed to unmarshal delta CRL URL: %v", err)
			}
			break
		}
	}

	if len(deltaCrlUrls) == 0 {
		return nil, fmt.Errorf("no delta CRL found")
	}

	// Assuming the first URL is the one we need
	deltaCrlUrl := deltaCrlUrls[0]
	return fetcher.GetCRL(deltaCrlUrl)
}

// HasDeltaCrl checks if a given CRL has a delta CRL extension.
func HasDeltaCrl(crl *pkix.CertificateList) bool {
	for _, ext := range crl.TBSCertList.Extensions {
		if ext.Id.Equal(OidExtensionFreshestCRL) {
			return true
		}
	}
	return false
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
