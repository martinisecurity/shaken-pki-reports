package linter_test

import (
	"net/url"
	"reflect"
	"testing"

	"github.com/martinisecurity/shaken-pki-reports/lint"
	"github.com/martinisecurity/shaken-pki-reports/linter"
)

func TestLintUrl_HttpStatus503(t *testing.T) {
	u, _ := url.Parse("https://app.connexcs.com/api/stir-shaken/cert/41.crt")
	res := linter.LintUrl(u)

	want := &lint.LintResult{
		Status:  lint.Error,
		Details: "HTTP response shall have StatusCode 200, but it is 503 Service Unavailable",
	}
	if test := res.Results["e_http_status_200"]; !reflect.DeepEqual(test, want) {
		t.Errorf("lint.LintUrl() = %v, want %v", test, want)
	}
}

func TestLintUrl_AtisContentType(t *testing.T) {
	u, _ := url.Parse("https://cr.sansay.com/548J/order/144_548J_67")
	res := linter.LintUrl(u)

	want := &lint.LintResult{
		Status:  lint.Warn,
		Details: "HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain",
	}
	if test := res.Results["w_atis_content_type"]; !reflect.DeepEqual(test, want) {
		t.Errorf("lint.LintUrl() = %v, want %v", test, want)
	}
}

func TestLintUrl_AtisProtocol(t *testing.T) {
	code := "w_atis_protocol"
	tests := []struct {
		name string
		url  string
		code string
		want *lint.LintResult
	}{
		{
			name: "https url with default port",
			url:  "https://api.alianza.com/v2/stir-shaken/certs/b45a4083-1554-4412-b5fc-bbd2c027091e/key.crt",
			code: code,
			want: &lint.LintResult{
				Status: lint.Pass,
			},
		},
		{
			name: "https url with port 443",
			url:  "https://api.alianza.com:443/v2/stir-shaken/certs/b45a4083-1554-4412-b5fc-bbd2c027091e/key.crt",
			code: code,
			want: &lint.LintResult{
				Status: lint.Pass,
			},
		},
		{
			name: "https url with port 8080",
			url:  "https://187.174.67.118:8080/7075515eb2d150fc98c43e794c07bbca.cer",
			code: code,
			want: &lint.LintResult{
				Status:  lint.Warn,
				Details: "The verifier should not dereference any protocol other than https or a port other than 443 or 8443",
			},
		},
		{
			name: "http url",
			url:  "http://sti.comsapi.com/258k/ca.crt",
			code: code,
			want: &lint.LintResult{
				Status:  lint.Warn,
				Details: "The verifier should not dereference any protocol other than https or a port other than 443 or 8443",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u, _ := url.Parse(tt.url)
			res := linter.LintUrl(u)
			if got := res.Results[tt.code]; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lint.LintUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}
