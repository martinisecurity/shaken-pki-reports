package lints

import (
	"net/url"
	"strings"

	"github.com/martinisecurity/shaken-pki-reports/lint"
)

type urlComponent struct{}

func init() {
	lint.RegisterRule(&lint.LintRule{
		Code:        "e_atis_url_component",
		Description: "The STI-VS shall not dereference URLs that contain a userinfo subcomponent, query component, or fragment identifier component",
		Source:      lint.Atis1000074Source,
		Rule:        NewUrlComponent,
	})
}

func NewUrlComponent() lint.LintRuleInterface {
	return &urlComponent{}
}

// CheckApplies implements LintUrlRuleInterface
func (*urlComponent) CheckApplies(data *lint.LintData) bool {
	_, err := url.Parse(data.Url)
	return err != nil
}

// Execute implements lint.LintUrlRuleInterface
func (*urlComponent) Execute(data *lint.LintData) *lint.LintResult {
	url, _ := url.Parse(data.Url)
	if len(url.Fragment) > 0 || strings.HasSuffix(data.Url, "#") ||
		len(url.RawQuery) > 0 || strings.HasSuffix(data.Url, "?") ||
		url.User != nil {
		return &lint.LintResult{
			Status:  lint.Error,
			Details: "The STI-VS shall not dereference URLs that contain a userinfo subcomponent, query component, or fragment identifier component",
		}
	}

	return &lint.LintResult{
		Status: lint.Pass,
	}
}
