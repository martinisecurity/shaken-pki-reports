package main

import (
	"encoding/pem"
	"reflect"
	"testing"

	"github.com/zmap/zcrypto/x509"
)

func TestCheckChainRevocation(t *testing.T) {
	certRaw1 := `-----BEGIN CERTIFICATE-----
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
-----END CERTIFICATE-----`

	pemBlock, _ := pem.Decode([]byte(certRaw1))
	cert1, _ := x509.ParseCertificate(pemBlock.Bytes)

	type args struct {
		chain []*x509.Certificate
	}
	tests := []struct {
		name    string
		args    args
		want    []RevocationStatus
		wantErr bool
	}{
		{
			name: "Martini Revoked Cert",
			args: args{
				chain: []*x509.Certificate{
					cert1,
				},
			},
			want: []RevocationStatus{
				RevocationStatusRevoked,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CheckChainRevocation(tt.args.chain)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckChainRevocation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CheckChainRevocation() = %v, want %v", got, tt.want)
			}
		})
	}
}
