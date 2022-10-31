package internal

import (
	"encoding/asn1"
	"encoding/pem"
	"fmt"
	"math"
	"time"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/util"
)

type PemCertificate struct {
	Headers     map[string]string
	Certificate *x509.Certificate
}

// ParseCertificates parses certificates from the DER or PEM source
func ParseCertificates(source []byte) []*x509.Certificate {
	res := []*x509.Certificate{}

	// PEM
	// pem may contain multiple certificates
	for {
		pem, restBytes := pem.Decode(source)
		if pem == nil {
			break
		}
		cert, _ := x509.ParseCertificate(pem.Bytes)
		if cert != nil {
			res = append(res, cert)
		}
		source = restBytes
	}

	// DER
	cert, _ := x509.ParseCertificate(source)
	if cert != nil {
		res = append(res, cert)
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

// TNAuthorizationList represents the ASN.1 structure of the same name. See RFC 8226
type TNAuthorizationList = []TNEntry

// TNEntry represents the ASN.1 structure of the same name. See RFC 8226
type TNEntry struct {
	SPC string
	// Range TelephoneNumberRange `asn1:"tag:1,optional,explicit"`
	// One   TelephoneNumber      `asn1:"tag:2,optional,explicit"`
}

// type ServiceProviderCode = string `asn1:"ia5string"`

func ParseTNAuthorizationList(raw []byte) (TNAuthorizationList, error) {
	var seq asn1.RawValue
	if _, err := asn1.Unmarshal(raw, &seq); err != nil {
		return nil, fmt.Errorf("bad TNAuthorizationList ASN.1 raw, %w", err)
	}
	if !seq.IsCompound || seq.Tag != 16 || seq.Class != 0 {
		return nil, asn1.StructuralError{Msg: "bad TNAuthorizationList sequence"}
	}

	res := make(TNAuthorizationList, 0)

	rest := seq.Bytes
	for len(rest) > 0 {
		var v asn1.RawValue
		var err error
		rest, err = asn1.Unmarshal(rest, &v)
		if err != nil {
			return nil, fmt.Errorf("bad TNEntry ASN.1 raw, %w", err)
		}
		switch v.Tag {
		case 0:
			var spc string
			if _, err := asn1.UnmarshalWithParams(v.Bytes, &spc, "ia5"); err != nil {
				return nil, fmt.Errorf("bad spc ASN.1 raw, %w", err)
			}
			res = append(res, TNEntry{
				SPC: spc,
			})
		}
	}

	return res, nil
}

func GetTNEntrySPC(c *x509.Certificate) (string, error) {
	ext := util.GetExtFromCert(c, util.TNAuthListOID)
	if ext != nil {
		tnList, err := ParseTNAuthorizationList(ext.Value)
		if err != nil {
			return "", fmt.Errorf("bad TNAuthorizationList, %w", err)
		}

		if len(tnList) != 1 {
			return "", fmt.Errorf("TNAuthorizationList shall have only one TN Entry")
		}

		spc := tnList[0].SPC
		if len(spc) == 0 {
			return "", fmt.Errorf("TN Entry shall contain a SPC value")
		}

		return spc, nil
	}

	return "", fmt.Errorf("STI certificate shall contain TNAuthorizationList extension")
}
