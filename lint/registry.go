package lint

import "fmt"

// LintRegistry contains a list of rules.
type LintRegistry struct {
	Rules LintRuleSet
}

// RegisterRule adds rule into the global registry. It panics ff the rule with the same code already exists/
func RegisterRule(rule *LintRule) {
	if rule != nil {
		if registry.Rules[rule.Code] != nil {
			panic(fmt.Sprintf("Cannot append %s the URL linter. This code is already in use", rule.Code))
		}
		registry.Rules[rule.Code] = rule
	}
}

// FindRuleByName returns rule for specified name from the global registry. If rule is not found returns nil/
func FindRuleByName(name string) *LintRule {
	return registry.Rules[name]
}

// global registry
var registry LintRegistry = LintRegistry{
	Rules: LintRuleSet{},
}

// GetGlobalRegistry returns the global registry.
func GetGlobalRegistry() *LintRegistry {
	return &registry
}
