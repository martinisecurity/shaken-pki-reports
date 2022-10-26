package lints_test

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/martinisecurity/shaken-pki-reports/lint"
	"github.com/martinisecurity/shaken-pki-reports/lints"
)

func Test_queryParams_Execute(t *testing.T) {
	type args struct {
		data *lint.LintData
	}
	tests := []struct {
		name string
		args args
		want *lint.LintResult
	}{
		{
			name: "correct url",
			args: args{
				data: &lint.LintData{
					Response: &http.Response{
						StatusCode: 200,
					},
					Url: "https://appreg.telcoportal.com/mobileapps/neustar/Microtalk-Shaken.pem",
				},
			},
			want: &lint.LintResult{
				Status: lint.Pass,
			},
		},
		{
			name: "url with query params",
			args: args{
				data: &lint.LintData{
					Response: &http.Response{
						StatusCode: 200,
					},
					Url: "https://appreg.telcoportal.com/mobileapps/neustar/Microtalk-Shaken.pem?param=1",
				},
			},
			want: &lint.LintResult{
				Status:  lint.Error,
				Details: "The STI-VS shall not dereference URLs that contain a userinfo subcomponent, query component, or fragment identifier component",
			},
		},
		{
			name: "url with empty query params",
			args: args{
				data: &lint.LintData{
					Response: &http.Response{
						StatusCode: 200,
					},
					Url: "https://appreg.telcoportal.com/mobileapps/neustar/Microtalk-Shaken.pem?",
				},
			},
			want: &lint.LintResult{
				Status:  lint.Error,
				Details: "The STI-VS shall not dereference URLs that contain a userinfo subcomponent, query component, or fragment identifier component",
			},
		},
		{
			name: "url with fragment",
			args: args{
				data: &lint.LintData{
					Response: &http.Response{
						StatusCode: 200,
					},
					Url: "https://appreg.telcoportal.com/mobileapps/neustar/Microtalk-Shaken.pem#some",
				},
			},
			want: &lint.LintResult{
				Status:  lint.Error,
				Details: "The STI-VS shall not dereference URLs that contain a userinfo subcomponent, query component, or fragment identifier component",
			},
		},
		{
			name: "url with empty fragment",
			args: args{
				data: &lint.LintData{
					Response: &http.Response{
						StatusCode: 200,
					},
					Url: "https://appreg.telcoportal.com/mobileapps/neustar/Microtalk-Shaken.pem#",
				},
			},
			want: &lint.LintResult{
				Status:  lint.Error,
				Details: "The STI-VS shall not dereference URLs that contain a userinfo subcomponent, query component, or fragment identifier component",
			},
		},
		{
			name: "url with user info",
			args: args{
				data: &lint.LintData{
					Response: &http.Response{
						StatusCode: 200,
					},
					Url: "https://userinfo@telcoportal.com/mobileapps/neustar/Microtalk-Shaken.pem",
				},
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
			if got := h.Execute(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("queryParams.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
