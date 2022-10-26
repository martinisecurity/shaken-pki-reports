package lints_test

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/martinisecurity/shaken-pki-reports/lint"
	"github.com/martinisecurity/shaken-pki-reports/lints"
)

func Test_atisCacheHeader_Execute(t *testing.T) {
	type args struct {
		data *lint.LintData
	}
	tests := []struct {
		name string
		args args
		want *lint.LintResult
	}{
		{
			name: "correct cache headers",
			args: args{
				data: &lint.LintData{
					Response: &http.Response{
						StatusCode: 200,
						Header: http.Header{
							"Date": []string{
								"Thu, 20 Oct 2022 19:52:23 GMT",
							},
							"Expires": []string{
								"Wed, 21 Oct 2022 19:52:23 GMT",
							},
							"Cache-Control": []string{
								"public, max-age=86400",
							},
						},
					},
				},
			},
			want: &lint.LintResult{
				Status: lint.Pass,
			},
		},
		{
			name: "incorrect Pragma",
			args: args{
				data: &lint.LintData{
					Response: &http.Response{
						StatusCode: 200,
						Header: http.Header{
							"Pragma": []string{
								"no-cache",
							},
						},
					},
				},
			},
			want: &lint.LintResult{
				Status:  lint.Error,
				Details: "The STI-VS shall implement the cache behavior. The Pragma header contains 'no-cache'",
			},
		},
		{
			name: "incorrect Date format",
			args: args{
				data: &lint.LintData{
					Response: &http.Response{
						StatusCode: 200,
						Header: http.Header{
							"Date": []string{
								"incorrect date format",
							},
							"Expires": []string{
								"incorrect date format",
							},
						},
					},
				},
			},
			want: &lint.LintResult{
				Status:  lint.Error,
				Details: "The STI-VS shall implement the cache behavior. Cannot parse Date or Expires headers",
			},
		},
		{
			name: "incorrect Expires duration",
			args: args{
				data: &lint.LintData{
					Response: &http.Response{
						StatusCode: 200,
						Header: http.Header{
							"Date": []string{
								"Thu, 20 Oct 2022 19:52:23 GMT",
							},
							"Expires": []string{
								"Wed, 21 Oct 2022 19:52:22 GMT",
							},
						},
					},
				},
			},
			want: &lint.LintResult{
				Status:  lint.Error,
				Details: "The STI-VS shall implement the cache behavior. The expiration time is less than 24 hours",
			},
		},
		{
			name: "cache-control is missed",
			args: args{
				data: &lint.LintData{
					Response: &http.Response{
						StatusCode: 200,
						Header: http.Header{
							"Date": []string{
								"Thu, 20 Oct 2022 19:52:23 GMT",
							},
							"Expires": []string{
								"Wed, 21 Oct 2022 19:52:23 GMT",
							},
						},
					},
				},
			},
			want: &lint.LintResult{
				Status:  lint.Error,
				Details: "The STI-VS shall implement the cache behavior. The Cache-Control header is missed",
			},
		},
		{
			name: "cache-control has private tag",
			args: args{
				data: &lint.LintData{
					Response: &http.Response{
						StatusCode: 200,
						Header: http.Header{
							"Date": []string{
								"Thu, 20 Oct 2022 19:52:23 GMT",
							},
							"Expires": []string{
								"Wed, 21 Oct 2022 19:52:23 GMT",
							},
							"Cache-Control": []string{
								"private",
							},
						},
					},
				},
			},
			want: &lint.LintResult{
				Status:  lint.Error,
				Details: "The STI-VS shall implement the cache behavior. The Cache-Control header doesn't have the 'private' tag",
			},
		},
		{
			name: "cache-control doesn't have max-age directive",
			args: args{
				data: &lint.LintData{
					Response: &http.Response{
						StatusCode: 200,
						Header: http.Header{
							"Date": []string{
								"Thu, 20 Oct 2022 19:52:23 GMT",
							},
							"Expires": []string{
								"Wed, 21 Oct 2022 19:52:23 GMT",
							},
							"Cache-Control": []string{
								"public",
							},
						},
					},
				},
			},
			want: &lint.LintResult{
				Status:  lint.Error,
				Details: "The STI-VS shall implement the cache behavior. The Cache-Control header doesn't have 'max-age' directive",
			},
		},
		{
			name: "cache-control doesn't have max-age directive",
			args: args{
				data: &lint.LintData{
					Response: &http.Response{
						StatusCode: 200,
						Header: http.Header{
							"Date": []string{
								"Thu, 20 Oct 2022 19:52:23 GMT",
							},
							"Expires": []string{
								"Wed, 21 Oct 2022 19:52:23 GMT",
							},
							"Cache-Control": []string{
								"public, max-age=3600",
							},
						},
					},
				},
			},
			want: &lint.LintResult{
				Status:  lint.Error,
				Details: "The STI-VS shall implement the cache behavior. The Cache-Control header has 'max-age' directive but it's value is less than 24 hours",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := lints.NewAtisCacheHeader()
			if got := h.Execute(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("atisCacheHeader.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
