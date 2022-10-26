package internal

import (
	"encoding/pem"
	"math"
	"time"

	"github.com/zmap/zcrypto/x509"
)

type PemCertificate struct {
	Headers     map[string]string
	Certificate *x509.Certificate
}

// ParseCertificates parses certificates from the DER or PEM source
func ParseCertificates(source []byte) []*PemCertificate {
	res := []*PemCertificate{}

	// PEM
	// pem may contain multiple certificates
	for {
		pem, restBytes := pem.Decode(source)
		if pem == nil {
			break
		}
		cert, _ := x509.ParseCertificate(pem.Bytes)
		if cert != nil {
			res = append(res, &PemCertificate{
				Headers:     pem.Headers,
				Certificate: cert,
			})
		}
		source = restBytes
	}

	// DER
	cert, _ := x509.ParseCertificate(source)
	if cert != nil {
		res = append(res, &PemCertificate{
			Headers:     map[string]string{},
			Certificate: cert,
		})
	}

	return res
}

// GetOrganizationName returns organization name from the certificate
//
// Priority:
//   - Issuer O
//   - Subject O
//   - Subject CN
func GetOrganizationName(c *x509.Certificate) string {
	org := "Unknown"
	if len(c.Issuer.Organization) > 0 {
		org = c.Issuer.Organization[0]
	} else if len(c.Subject.Organization) > 0 {
		org = c.Subject.Organization[0]
	} else if len(c.Subject.CommonName) > 0 {
		org = c.Subject.CommonName
	}

	return org
}

// GetValidityDays returns validity period of the certificate in days
func GetValidityDays(c *x509.Certificate) int {
	return int(math.Ceil(float64(c.ValidityPeriod) / float64(86400)))
}

// GetRemainingDays returns remaining validity period of the certificate in days
func GetRemainingDays(c *x509.Certificate, now time.Time) int {
	return int(math.Ceil(c.NotAfter.Sub(now).Seconds() / 86400))
}
