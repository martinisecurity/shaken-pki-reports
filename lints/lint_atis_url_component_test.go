package lints_test

import (
	"net/http"
	"net/url"
	"reflect"
	"testing"

	"github.com/martinisecurity/shaken-pki-reports/lint"
	"github.com/martinisecurity/shaken-pki-reports/lints"
)

func Test_queryParams_Execute(t *testing.T) {
	type args struct {
		StatusCode int
		Url        string
	}
	tests := []struct {
		name string
		args args
		want *lint.LintResult
	}{
		{
			name: "correct url",
			args: args{
				StatusCode: 200,
				Url:        "https://appreg.telcoportal.com/mobileapps/neustar/Microtalk-Shaken.pem",
			},
			want: &lint.LintResult{
				Status: lint.Pass,
			},
		},
		{
			name: "url with query params",
			args: args{
				StatusCode: 200,
				Url:        "https://appreg.telcoportal.com/mobileapps/neustar/Microtalk-Shaken.pem?param=1",
			},
			want: &lint.LintResult{
				Status:  lint.Error,
				Details: "The STI-VS shall not dereference URLs that contain a userinfo subcomponent, query component, or fragment identifier component",
			},
		},
		{
			name: "url with empty query params",
			args: args{
				StatusCode: 200,
				Url:        "https://appreg.telcoportal.com/mobileapps/neustar/Microtalk-Shaken.pem?",
			},
			want: &lint.LintResult{
				Status:  lint.Error,
				Details: "The STI-VS shall not dereference URLs that contain a userinfo subcomponent, query component, or fragment identifier component",
			},
		},
		{
			name: "url with fragment",
			args: args{
				StatusCode: 200,
				Url:        "https://appreg.telcoportal.com/mobileapps/neustar/Microtalk-Shaken.pem#some",
			},
			want: &lint.LintResult{
				Status:  lint.Error,
				Details: "The STI-VS shall not dereference URLs that contain a userinfo subcomponent, query component, or fragment identifier component",
			},
		},
		{
			name: "url with user info",
			args: args{
				StatusCode: 200,
				Url:        "https://userinfo@telcoportal.com/mobileapps/neustar/Microtalk-Shaken.pem",
			},
			want: &lint.LintResult{
				Status:  lint.Error,
				Details: "The STI-VS shall not dereference URLs that contain a userinfo subcomponent, query component, or fragment identifier component",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := lints.NewUrlComponent()
			u, _ := url.Parse(tt.args.Url)
			d := &lint.LintData{
				Url: u,
				Response: &http.Response{
					StatusCode: tt.args.StatusCode,
				},
			}
			if got := h.Execute(d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("queryParams.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
