package lints_test

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/martinisecurity/shaken-pki-reports/lint"
	"github.com/martinisecurity/shaken-pki-reports/lints"
)

func Test_httpStatus200_Execute(t *testing.T) {
	type args struct {
		data *lint.LintData
	}
	tests := []struct {
		name string
		args args
		want *lint.LintResult
	}{
		{
			name: "status 200",
			args: args{
				data: &lint.LintData{
					Response: &http.Response{
						StatusCode: 200,
					},
				},
			},
			want: &lint.LintResult{
				Status: lint.Pass,
			},
		},
		{
			name: "status is not 200",
			args: args{
				data: &lint.LintData{
					Response: &http.Response{
						StatusCode: 500,
						Status:     "500 Internal",
					},
				},
			},
			want: &lint.LintResult{
				Status:  lint.Error,
				Details: "HTTP response shall have StatusCode 200, but it is 500 Internal",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := lints.NewHttpStatus200()
			if got := h.Execute(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("httpStatus200.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
