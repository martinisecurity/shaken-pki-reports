package lints

import (
	"strings"

	"github.com/martinisecurity/shaken-pki-reports/lint"
)

type atisHeaderContentType struct{}

func init() {
	lint.RegisterRule(&lint.LintRule{
		Code:        "w_atis_content_type",
		Description: "ATIS-1000080 separately indicates that the mime type should be application/pem-certificate-chain",
		Source:      lint.Atis1000080Source,
		Rule:        NewAtisHeaderContentType,
	})
}

func NewAtisHeaderContentType() lint.LintRuleInterface {
	return &atisHeaderContentType{}
}

// CheckApplies implements LintUrlRuleInterface
func (*atisHeaderContentType) CheckApplies(data *lint.LintData) bool {
	return true
}

// Execute implements lint.LintUrlRuleInterface
func (*atisHeaderContentType) Execute(data *lint.LintData) *lint.LintResult {
	contentType := data.Response.Header.Get("Content-Type")
	if !strings.Contains(contentType, "application/pem-certificate-chain") {
		return &lint.LintResult{
			Status:  lint.Warn,
			Details: "HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain",
		}
	}

	return &lint.LintResult{
		Status: lint.Pass,
	}
}
