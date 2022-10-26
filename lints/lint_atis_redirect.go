package lints

import "github.com/martinisecurity/shaken-pki-reports/lint"

type redirect struct{}

func init() {
	lint.RegisterRule(&lint.LintRule{
		Code:        "e_atis_redirect",
		Description: "The STI-VS shall not follow HTTP redirections",
		Source:      lint.Atis1000074Source,
		Rule:        NewRedirect,
	})
}

func NewRedirect() lint.LintRuleInterface {
	return &redirect{}
}

// CheckApplies implements LintUrlRuleInterface
func (*redirect) CheckApplies(data *lint.LintData) bool {
	return false
}

// Execute implements lint.LintUrlRuleInterface
func (*redirect) Execute(data *lint.LintData) *lint.LintResult {
	// The real test see in LintUrl method
	return &lint.LintResult{
		Status: lint.Pass,
	}
}
