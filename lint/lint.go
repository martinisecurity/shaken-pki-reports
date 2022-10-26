package lint

import (
	"net/http"
	"time"
)

// LintRule contains information about the rule.
type LintRule struct {
	// Code is an unique identifier of the rule.
	Code string
	// Description keeps an information about the rule.
	Description string
	Source      LintSource
	// Rule keeps constructor function of the rule.
	Rule func() LintRuleInterface
}

// LintData contains common data for testing.
type LintData struct {
	//  Url keeps URL link.
	Url string
	// Response keeps HTTP response.
	Response *http.Response
	Error    error
	// Content of the HTTP response
	Body []byte
}

// HasResponse returns true if Response field is empty, otherwise false.
func (t *LintData) HasResponse() bool {
	return t.Response != nil
}

// LintRuleInterface declares common methods for the rule.
type LintRuleInterface interface {
	// CheckApplies returns true if the rule should be applied for the testing, otherwise false.
	CheckApplies(data *LintData) bool
	// Execute runs test for the incoming test data.
	Execute(data *LintData) *LintResult
}

// LintRuleSet map of rules.
type LintRuleSet map[string]*LintRule

// LintResult contains result of the rule test.
type LintResult struct {
	Status  LintStatus
	Details string
}

// LintResultSet contains a list of rule tests and common information about the test.
type LintResultSet struct {
	Timestamp time.Time
	// Url keeps URL link.
	Url string
	// StatusCode keeps HTTP status code.
	StatusCode int
	// Body keeps HTTP response content.
	Body        []byte
	Results     map[string]*LintResult
	HasErrors   bool
	HasWarnings bool
	HasNotices  bool
}

// Append adds LintResult into the set and updates statuses.
func (t *LintResultSet) Append(code string, item *LintResult) {
	if item != nil {
		t.Results[code] = item

		// update statuses
		if item.Status == Error {
			t.HasErrors = true
		}
		if item.Status == Warn {
			t.HasWarnings = true
		}
		if item.Status == Notice {
			t.HasNotices = true
		}
	}
}

// HasProblem returns true if the set contains some tests with Error, Warn, or Notice statuses.
func (t *LintResultSet) HasProblem(code string) bool {
	for c, r := range t.Results {
		if c == code && (r.Status == Error || r.Status == Warn || r.Status == Notice) {
			return true
		}
	}

	return false
}
