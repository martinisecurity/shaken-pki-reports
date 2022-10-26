package lints

import (
	"encoding/pem"

	"github.com/martinisecurity/shaken-pki-reports/lint"
	"github.com/zmap/zcrypto/x509"
)

type pemCertificateChain struct{}

const pemCertificateChain_details = "ATIS-1000080 separately indicates that the mime type should be application/pem-certificate-chain"

func init() {
	lint.RegisterRule(&lint.LintRule{
		Code:        "w_aits_pem_certificate_chain",
		Description: pemCertificateChain_details,
		Source:      lint.Atis1000080Source,
		Rule:        NewPemCertificateChain,
	})
}

func NewPemCertificateChain() lint.LintRuleInterface {
	return &pemCertificateChain{}
}

// CheckApplies implements LintUrlRuleInterface
func (*pemCertificateChain) CheckApplies(data *lint.LintData) bool {
	return true
}

// Execute implements lint.LintUrlRuleInterface
func (*pemCertificateChain) Execute(data *lint.LintData) *lint.LintResult {
	var block *pem.Block
	rest := data.Body

	first := true
	for {
		block, rest = pem.Decode(rest)
		if block == nil {
			if first {
				return &lint.LintResult{
					Status:  lint.Warn,
					Details: "HTTP response body should be PEM certificate chain. Response body is not PEM encoded",
				}
			}
			break
		}
		first = false

		// try parse the certificate
		_, err := x509.ParseCertificate(block.Bytes)
		if err != nil {
			return &lint.LintResult{
				Status:  lint.Warn,
				Details: "HTTP response body should be PEM certificate chain. PEM is not a certificate",
			}
		}
	}

	return &lint.LintResult{
		Status: lint.Pass,
	}
}
