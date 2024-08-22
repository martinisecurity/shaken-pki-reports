package main

import (
	"encoding/pem"
	"errors"
	"reflect"
	"testing"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zcrypto/x509/pkix"
)

// MockFetcher is a mock implementation of the CRLFetcher interface for testing purposes.
type MockFetcher struct {
	CRLs map[string]*pkix.CertificateList
	Err  error
}

// GetCRL returns a pre-defined CRL or an error if set.
func (m *MockFetcher) GetCRL(crlUrl string) (*pkix.CertificateList, error) {
	if m.Err != nil {
		return nil, m.Err
	}
	crl, ok := m.CRLs[crlUrl]
	if !ok {
		return nil, errors.New("CRL not found")
	}
	return crl, nil
}

func TestCheckCertRevocation(t *testing.T) {
	certRaw := `-----BEGIN CERTIFICATE-----
MIIB/TCCAaOgAwIBAgIRAKWysCPZ9qYxP2vgg86yrOowCgYIKoZIzj0EAwIwPDEZ
MBcGA1UEChMQTWFydGluaSBTZWN1cml0eTEfMB0GA1UEAxMWTWFydGluaSBTZWN1
cml0eSBTSEFLRTAeFw0yNDA4MjIxMjUyNDdaFw0yNTA4MjIxMjUyNDdaMDwxGTAX
BgNVBAoTEE1hcnRpbmkgU2VjdXJpdHkxHzAdBgNVBAMTFk1hcnRpbmkgU2VjdXJp
dHkgU0hBS0UwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATxyxVkxk1/DxmRapK8
87oFxxo630FIkQnEWFMq/5/Pk6qjeRm2ZZR1EJnJUO07Mcm8EU5eURelCmsstA3M
Guzjo4GFMIGCMA4GA1UdDwEB/wQEAwIBhjATBgNVHSUEDDAKBggrBgEFBQcDATAP
BgNVHRMBAf8EBTADAQH/MB0GA1UdDgQWBBQPSC5FUCpRK9Jpdr0cbT9xGUJWrTAr
BgNVHR8EJDAiMCCgHqAchhpodHRwOi8vZXhhbXBsZS5jb20vY3JsLnBlbTAKBggq
hkjOPQQDAgNIADBFAiB5wpaQXMhFQgl4QUanFXqgO+A86G+2LwOlwNYZVb6wygIh
AJ42TiB5Dt66wpXyQeotEcqaUOgctR0o7mz6H9+ylkEy
-----END CERTIFICATE-----`

	crlNotRevokedRaw := `-----BEGIN X509 CRL-----
MIH1MIGcAgEBMAoGCCqGSM49BAMCMDwxGTAXBgNVBAoTEE1hcnRpbmkgU2VjdXJp
dHkxHzAdBgNVBAMTFk1hcnRpbmkgU2VjdXJpdHkgU0hBS0UXDTI0MDgyMjEyNTI0
N1oXDTI1MDgyMjEyNTI0N1qgLzAtMB8GA1UdIwQYMBaAFA9ILkVQKlEr0ml2vRxt
P3EZQlatMAoGA1UdFAQDAgEBMAoGCCqGSM49BAMCA0gAMEUCIQD00oJC5abiVt2e
asD8/E1HzF2b8tdtKMQq8zfc66w+BgIgVMm2h8b5LwG9taQc3y/UxKQr9vrrYsZD
3hcKa4AjYIU=
-----END X509 CRL-----`

	crlRevokedRaw := `-----BEGIN X509 CRL-----
MIIBHDCBwgIBATAKBggqhkjOPQQDAjA8MRkwFwYDVQQKExBNYXJ0aW5pIFNlY3Vy
aXR5MR8wHQYDVQQDExZNYXJ0aW5pIFNlY3VyaXR5IFNIQUtFFw0yNDA4MjIxMjUy
NDdaFw0yNTA4MjIxMjUyNDdaMCQwIgIRAKWysCPZ9qYxP2vgg86yrOoXDTI0MDgy
MjEyNTI0N1qgLzAtMB8GA1UdIwQYMBaAFA9ILkVQKlEr0ml2vRxtP3EZQlatMAoG
A1UdFAQDAgECMAoGCCqGSM49BAMCA0kAMEYCIQCss48hcvQryoXTGDxeN9K1hDIJ
tOvYeVPGGcXQkK63ywIhAKFoLzRRFRii1wCegca0kZBr+TNIzfh7Fqdu5RHFao0T
-----END X509 CRL-----`

	crlWithDeltaURLNotRevokedRaw := `-----BEGIN X509 CRL-----
MIIBIjCByQIBATAKBggqhkjOPQQDAjA8MRkwFwYDVQQKExBNYXJ0aW5pIFNlY3Vy
aXR5MR8wHQYDVQQDExZNYXJ0aW5pIFNlY3VyaXR5IFNIQUtFFw0yNDA4MjIxMjUy
NDdaFw0yNTA4MjIxMjUyNDdaoFwwWjAfBgNVHSMEGDAWgBQPSC5FUCpRK9Jpdr0c
bT9xGUJWrTAKBgNVHRQEAwIBAzArBgNVHS4EJDAiEyBodHRwOi8vZXhhbXBsZS5j
b20vY3JsL2RlbHRhLnBlbTAKBggqhkjOPQQDAgNIADBFAiEA/1f6XfDyJmdRwgMu
kyITE0l2eq5AAGKzd0YAppl9ccQCIH+u853gM0RdlGDmlXd/z7JeIj6rSatQFct/
bpS/n4Zd
-----END X509 CRL-----`

	deltaCrlNotRevokedRaw := `-----BEGIN X509 CRL-----
MIH0MIGcAgEBMAoGCCqGSM49BAMCMDwxGTAXBgNVBAoTEE1hcnRpbmkgU2VjdXJp
dHkxHzAdBgNVBAMTFk1hcnRpbmkgU2VjdXJpdHkgU0hBS0UXDTI0MDgyMjEyNTI0
N1oXDTI1MDgyMjEyNTI0N1qgLzAtMB8GA1UdIwQYMBaAFA9ILkVQKlEr0ml2vRxt
P3EZQlatMAoGA1UdFAQDAgEEMAoGCCqGSM49BAMCA0cAMEQCIF0+0Qq51WkVpLHB
mmfFiyqH0YpPaHWznvFdZpXCR9HeAiBwHgl2toHgN8a8mf2fR4B1bSx1gAZlm7j3
u8DqHzMerg==
-----END X509 CRL-----`

	deltaCrlRevokedRaw := `-----BEGIN X509 CRL-----
MIIBGzCBwgIBATAKBggqhkjOPQQDAjA8MRkwFwYDVQQKExBNYXJ0aW5pIFNlY3Vy
aXR5MR8wHQYDVQQDExZNYXJ0aW5pIFNlY3VyaXR5IFNIQUtFFw0yNDA4MjIxMjUy
NDdaFw0yNTA4MjIxMjUyNDdaMCQwIgIRAKWysCPZ9qYxP2vgg86yrOoXDTI0MDgy
MjEyNTI0N1qgLzAtMB8GA1UdIwQYMBaAFA9ILkVQKlEr0ml2vRxtP3EZQlatMAoG
A1UdFAQDAgEFMAoGCCqGSM49BAMCA0gAMEUCIHqjUTmrH8L7hXPrxK5/rdpFR4s2
+RmsJA24Z1Kq3x+RAiEA/RO2cgMZlPbB8BwHth7WVwfvWCT9sjpOo1zSQ7DzDkI=
-----END X509 CRL-----`

	pemBlock, _ := pem.Decode([]byte(certRaw))
	cert, _ := x509.ParseCertificate(pemBlock.Bytes)

	tests := []struct {
		name    string
		crlUrls map[string]string
		want    RevocationStatus
		wantErr bool
	}{
		{
			name: "Certificate not revoked",
			crlUrls: map[string]string{
				"http://example.com/crl.pem": crlNotRevokedRaw,
			},
			want:    RevocationStatusGood,
			wantErr: false,
		},
		{
			name: "Certificate revoked",
			crlUrls: map[string]string{
				"http://example.com/crl.pem": crlRevokedRaw,
			},
			want:    RevocationStatusRevoked,
			wantErr: false,
		},
		{
			name: "Delta CRL revoked",
			crlUrls: map[string]string{
				"http://example.com/crl.pem":       crlWithDeltaURLNotRevokedRaw,
				"http://example.com/crl/delta.pem": deltaCrlRevokedRaw,
			},
			want:    RevocationStatusRevoked,
			wantErr: false,
		},
		{
			name: "Delta CRL not revoked",
			crlUrls: map[string]string{
				"http://example.com/crl.pem":       crlWithDeltaURLNotRevokedRaw,
				"http://example.com/crl/delta.pem": deltaCrlNotRevokedRaw,
			},
			want:    RevocationStatusGood,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			crls := make(map[string]*pkix.CertificateList)
			for url, crlRaw := range tt.crlUrls {
				crl, _ := x509.ParseCRL([]byte(crlRaw))
				crls[url] = crl
			}
			mockFetcher := &MockFetcher{
				CRLs: crls,
			}

			got, err := CheckCertRevocationWithFetcher(cert, mockFetcher)
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
