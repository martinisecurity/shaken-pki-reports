package lints

import "github.com/martinisecurity/shaken-pki-reports/lint"

func init() {
	lint.RegisterRule(&lint.LintRule{
		Code:        "e_bad_url",
		Source:      lint.SystemSource,
		Description: "Subscribers shall provide unrestricted access to its repositories and shall implement logical and physical controls to prevent unauthorized write access to those repositories",
		Rule:        NewNothing,
	})
	lint.RegisterRule(&lint.LintRule{
		Code:        "e_tls_transport",
		Source:      lint.SystemSource,
		Description: "TLS problem on link loading",
		Rule:        NewNothing,
	})
}

type nothing struct{}

func NewNothing() lint.LintRuleInterface {
	return &nothing{}
}

// CheckApplies implements LintUrlRuleInterface
func (*nothing) CheckApplies(data *lint.LintData) bool {
	return false
}

// Execute implements lint.LintUrlRuleInterface
func (*nothing) Execute(data *lint.LintData) *lint.LintResult {
	return &lint.LintResult{
		Status: lint.Pass,
	}
}
