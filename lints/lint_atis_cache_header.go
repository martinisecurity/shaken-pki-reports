package lints

import (
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/martinisecurity/shaken-pki-reports/lint"
)

type atisCacheHeader struct{}

func init() {
	lint.RegisterRule(&lint.LintRule{
		Code:        "e_atis_cache_header",
		Description: "The STI-VS shall implement the cache behavior described in RFC7234. If the HTTP response does not include any recognized caching directives or indicates caching for less than 24 hours, then the STI-VS should cache the HTTP response for 24 hours",
		Source:      lint.Atis1000074Source,
		Rule:        NewAtisCacheHeader,
	})
}

func NewAtisCacheHeader() lint.LintRuleInterface {
	return &atisCacheHeader{}
}

// CheckApplies implements LintUrlRuleInterface
func (*atisCacheHeader) CheckApplies(data *lint.LintData) bool {
	return true
}

func some(values []string, cb func(text string) bool) bool {
	for _, v := range values {
		if cb(v) {
			return true
		}
	}

	return false
}

func first(values []string) *string {
	if len(values) > 0 {
		return &values[0]
	}

	return nil
}

// Execute implements lint.LintUrlRuleInterface
func (*atisCacheHeader) Execute(data *lint.LintData) *lint.LintResult {
	pragma := data.Response.Header["Pragma"]
	date := first(data.Response.Header["Date"])
	expires := first(data.Response.Header["Expires"])
	cacheControl := first(data.Response.Header["Cache-Control"])

	if some(pragma, func(text string) bool { return strings.ToLower(text) == "no-cache" }) {
		return &lint.LintResult{
			Status:  lint.Error,
			Details: "The STI-VS shall implement the cache behavior. The Pragma header contains 'no-cache'",
		}
	}

	if date != nil && expires != nil {
		dateTime, err := time.Parse(time.RFC1123, *date)
		expiresTime, err2 := time.Parse(time.RFC1123, *expires)
		if err != nil || err2 != nil {
			return &lint.LintResult{
				Status:  lint.Error,
				Details: "The STI-VS shall implement the cache behavior. Cannot parse Date or Expires headers",
			}
		}

		controlDate := dateTime.AddDate(0, 0, 1)
		if !(controlDate.Before(expiresTime) || controlDate.Equal(expiresTime)) {
			return &lint.LintResult{
				Status:  lint.Error,
				Details: "The STI-VS shall implement the cache behavior. The expiration time is less than 24 hours",
			}
		}
	}

	// TODO: Determine if RFC 7234 has language about ignoring Expires header if it conflicts with Cache-Control.
	if cacheControl == nil {
		return &lint.LintResult{
			Status:  lint.Error,
			Details: "The STI-VS shall implement the cache behavior. The Cache-Control header is missed",
		}
	}

	if strings.Contains(*cacheControl, "private") {
		return &lint.LintResult{
			Status:  lint.Error,
			Details: "The STI-VS shall implement the cache behavior. The Cache-Control header doesn't have the 'private' tag",
		}
	}

	regex := regexp.MustCompile("max-age=([0-9]+)")
	matches := regex.FindAllStringSubmatch(*cacheControl, -1)
	if len(matches) == 0 || len(matches[0]) == 0 {
		return &lint.LintResult{
			Status:  lint.Error,
			Details: "The STI-VS shall implement the cache behavior. The Cache-Control header doesn't have 'max-age' directive",
		}
	} else if sec, err := strconv.Atoi(matches[0][1]); err != nil || sec < 86400 {
		return &lint.LintResult{
			Status:  lint.Error,
			Details: "The STI-VS shall implement the cache behavior. The Cache-Control header has 'max-age' directive but it's value is less than 24 hours",
		}
	}

	return &lint.LintResult{
		Status: lint.Pass,
	}
}
