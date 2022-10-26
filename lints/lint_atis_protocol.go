package lints

import "github.com/martinisecurity/shaken-pki-reports/lint"

type atisProtocol struct{}

func init() {
	lint.RegisterRule(&lint.LintRule{
		Code:        "w_atis_protocol",
		Description: "The verifier should not dereference any protocol other than https or a port other than 443 or 8443",
		Source:      lint.Atis1000080Source,
		Rule:        NewAtisProtocol,
	})
}

func NewAtisProtocol() lint.LintRuleInterface {
	return &atisProtocol{}
}

// CheckApplies implements LintUrlRuleInterface
func (*atisProtocol) CheckApplies(data *lint.LintData) bool {
	return true
}

// Execute implements lint.LintUrlRuleInterface
func (*atisProtocol) Execute(data *lint.LintData) *lint.LintResult {
	port := data.Response.Request.URL.Port()
	if !(data.Response.Request.URL.Scheme == "https" &&
		(port == "" || port == "443" || port == "8443")) {
		return &lint.LintResult{
			Status:  lint.Warn,
			Details: "The verifier should not dereference any protocol other than https or a port other than 443 or 8443",
		}
	}

	return &lint.LintResult{
		Status: lint.Pass,
	}
}
