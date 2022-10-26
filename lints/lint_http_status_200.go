package lints

import (
	"fmt"

	"github.com/martinisecurity/shaken-pki-reports/lint"
)

type httpStatus200 struct{}

const httpStatus200_details = "HTTP response shall have StatusCode 200"

func init() {
	lint.RegisterRule(&lint.LintRule{
		Code:        "e_http_status_200",
		Description: httpStatus200_details,
		Source:      lint.HttpSource,
		Rule:        NewHttpStatus200,
	})
}

func NewHttpStatus200() lint.LintRuleInterface {
	return &httpStatus200{}
}

// CheckApplies implements LintUrlRuleInterface
func (*httpStatus200) CheckApplies(data *lint.LintData) bool {
	return true
}

// Execute implements lint.LintUrlRuleInterface
func (*httpStatus200) Execute(data *lint.LintData) *lint.LintResult {
	if data.Response.StatusCode != 200 {
		return &lint.LintResult{
			Status:  lint.Error,
			Details: fmt.Sprintf("%s, but it is %s", httpStatus200_details, data.Response.Status),
		}
	}

	return &lint.LintResult{
		Status: lint.Pass,
	}
}
