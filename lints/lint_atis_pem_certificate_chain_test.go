package lints_test

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/martinisecurity/shaken-pki-reports/lint"
	"github.com/martinisecurity/shaken-pki-reports/lints"
)

func Test_pemCertificateChain_Execute(t *testing.T) {
	type args struct {
		data *lint.LintData
	}
	tests := []struct {
		name string
		args args
		want *lint.LintResult
	}{
		{
			name: "correct x509 chain",
			args: args{
				data: &lint.LintData{
					Response: &http.Response{
						StatusCode: 200,
					},
					Body: []byte(`-----BEGIN CERTIFICATE-----
MIIDKjCCAtGgAwIBAgIUHANi+8RbpIVe38s/fJkNV/r2hakwCgYIKoZIzj0EAwIw
cTELMAkGA1UEBhMCVVMxCzAJBgNVBAgTAldBMRAwDgYDVQQHEwdTZWF0dGxlMR4w
HAYDVQQKExVNYXJ0aW5pIFNlY3VyaXR5LCBMTEMxIzAhBgNVBAMTGk1hcnRpbmkg
U2VjdXJpdHkgU0hBS0VOIEcxMB4XDTIyMDkyODE3NTQyNVoXDTIyMTIyNzA2MDAw
MFowdzEUMBIGA1UEAxMLU0hBS0VOIDcwOUoxKTAnBgNVBAUTIDQ1NUE4MTY0RkEy
NkVGMUZCM0ExNEM0MzE4ODM0MzhEMQswCQYDVQQGEwJVUzEnMCUGA1UEChMeTE9X
IExBVEVOQ1kgQ09NTVVOSUNBVElPTlMgTExDMFkwEwYHKoZIzj0CAQYIKoZIzj0D
AQcDQgAEspSxrqIB0kl4lLRyHx4sliqf58RNDQJlqz/xhS/lnlJsQTS3qVs6gqeA
fe6rHejfGVosbM8H2Vx775E46ClfYqOCAT8wggE7MA4GA1UdDwEB/wQEAwIHgDAM
BgNVHRMBAf8EAjAAMB0GA1UdDgQWBBRIGyEaQDL4bz7/pdCXgEmoFb/EbTAfBgNV
HSMEGDAWgBQt1ctQCVNMkSR/I7Gp5SWHOJZJgjAWBggrBgEFBQcBGgQKMAigBhYE
NzA5SjCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUt
YXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQH
EwtCcmlkZ2V3YXRlcjELMAkGA1UECBMCTkoxEzARBgNVBAMTClNUSS1QQSBDUkwx
CzAJBgNVBAYTAlVTMQ8wDQYDVQQKEwZTVEktUEEwGgYDVR0gAQH/BBAwDjAMBgpg
hkgBhv8JAQEDMAoGCCqGSM49BAMCA0cAMEQCICb75dQUYMstzzmEtSXSS0y2aXX5
GhN9uagR1iVaScVMAiBEGBMoMxPZZ0mBEkG+aSB9DujyZy6DLZmrl5NkWhxB3w==
-----END CERTIFICATE-----
#   Issuer: C=US, ST=WA, L=Seattle, O=Martini Security, LLC, CN=Martini Security SHAKEN R1
#   Subject: C=US, ST=WA, L=Seattle, O=Martini Security, LLC, CN=Martini Security SHAKEN G1
#   Validity
#       Not Before: May  3 15:41:00 2022 GMT
#       Not After : May  2 15:41:00 2027 GMT
-----BEGIN CERTIFICATE-----
MIIDEjCCArmgAwIBAgIUWRP/eaQF+0VNFxze3w+czApA53owCgYIKoZIzj0EAwIw
cTELMAkGA1UEBhMCVVMxCzAJBgNVBAgTAldBMRAwDgYDVQQHEwdTZWF0dGxlMR4w
HAYDVQQKExVNYXJ0aW5pIFNlY3VyaXR5LCBMTEMxIzAhBgNVBAMTGk1hcnRpbmkg
U2VjdXJpdHkgU0hBS0VOIFIxMB4XDTIyMDUwMzE1NDEwMFoXDTI3MDUwMjE1NDEw
MFowcTELMAkGA1UEBhMCVVMxCzAJBgNVBAgTAldBMRAwDgYDVQQHEwdTZWF0dGxl
MR4wHAYDVQQKExVNYXJ0aW5pIFNlY3VyaXR5LCBMTEMxIzAhBgNVBAMTGk1hcnRp
bmkgU2VjdXJpdHkgU0hBS0VOIEcxMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE
gHMKTCPENlBRYIP/LPOk528hT0mhODS+qQkHt5ujQkKSfgLeossZ1xb/vZFlUIIv
z67C6Ugb5nupkadMzFQWZ6OCAS0wggEpMA4GA1UdDwEB/wQEAwICBDASBgNVHRMB
Af8ECDAGAQH/AgEAMB0GA1UdDgQWBBQt1ctQCVNMkSR/I7Gp5SWHOJZJgjAfBgNV
HSMEGDAWgBSLmgJLXKXv3jruIFZxYMWJNOBIeTCBpgYDVR0fBIGeMIGbMIGYoDqg
OIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxv
YWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwC
TkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZT
VEktUEEwGgYDVR0gAQH/BBAwDjAMBgpghkgBhv8JAQEDMAoGCCqGSM49BAMCA0cA
MEQCICYdBOW7vfq5o01UgzJTmpu6qfjKbZA/thPQEFb68e+wAiAmSMR/RPxKHvbK
Ga4z8ktL5XUwkI2fbbbxe322Vke0HQ==
-----END CERTIFICATE-----`),
				},
			},
			want: &lint.LintResult{
				Status: lint.Pass,
			},
		},
		{
			name: "incorrect PEM",
			args: args{
				data: &lint.LintData{
					Response: &http.Response{
						StatusCode: 200,
					},
					Body: []byte(`some text`),
				},
			},
			want: &lint.LintResult{
				Status:  lint.Warn,
				Details: "HTTP response body should be PEM certificate chain. Response body is not PEM encoded",
			},
		},
		{
			name: "PEM is not a certificate",
			args: args{
				data: &lint.LintData{
					Response: &http.Response{
						StatusCode: 200,
					},
					Body: []byte(`-----BEGIN RSA PUBLIC EXPONENT-----
AQAB
-----END RSA PUBLIC EXPONENT-----`),
				},
			},
			want: &lint.LintResult{
				Status:  lint.Warn,
				Details: "HTTP response body should be PEM certificate chain. PEM is not a certificate",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := lints.NewPemCertificateChain()
			if got := h.Execute(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pemCertificateChain.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
