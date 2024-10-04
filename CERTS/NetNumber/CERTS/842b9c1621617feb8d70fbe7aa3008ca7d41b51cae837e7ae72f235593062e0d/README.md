# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Plivo Inc

Tested At: 04 Oct 24 15:31 UTC\
Initial Validity Period: 364 day(s)\
Remaining Validity Period: 209 day(s)\
Subject: L=Austin, ST=Texas, O=Plivo Inc, C=US, CN=Plivo Inc\
Issuer: L=Lowell, ST=Massachusettes, OU=US, O=NetNumber Inc, C=US, CN=NetNumber SHAKEN Root Intermediate CA 1\
Link: https://certificate.zt.plivo.com/cert20250501.crt

[View certificate details](https://x509.io/?cert=MIIC2zCCAmCgAwIBAgIIZZwUf2RrZJcwCgYIKoZIzj0EAwIwgY4xMDAuBgNVBAMMJ05ldE51bWJlciBTSEFLRU4gUm9vdCBJbnRlcm1lZGlhdGUgQ0EgMTELMAkGA1UEBhMCVVMxFjAUBgNVBAoMDU5ldE51bWJlciBJbmMxCzAJBgNVBAsMAlVTMRcwFQYDVQQIDA5NYXNzYWNodXNldHRlczEPMA0GA1UEBwwGTG93ZWxsMB4XDTI0MDUwMjAwMDAwMFoXDTI1MDUwMTAwMDAwMFowVjESMBAGA1UEAwwJUGxpdm8gSW5jMQswCQYDVQQGEwJVUzESMBAGA1UECgwJUGxpdm8gSW5jMQ4wDAYDVQQIDAVUZXhhczEPMA0GA1UEBwwGQXVzdGluMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEsZ%2F53R%2FoY2bL%2F8%2Fz58ZsWX2%2BwobsCB9wbC7fxIaZTh7HseqKjK4XWWv%2F8NWIH%2BSFzgUCFw9PgVHr0Ora5YR%2BT6OB3jCB2zAWBggrBgEFBQcBGgQKMAigBhYEODAwSjAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAfBgNVHSMEGDAWgBRxL8iC3KjgIuPfoGj5%2BF5chN7lvTAdBgNVHQ4EFgQU5GqHIC7QCosqjUZDFZ7%2BM3DHqa8wRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMBoGA1UdIAEB%2FwQQMA4wDAYKYIZIAYb%2FCQEBATAKBggqhkjOPQQDAgNpADBmAjEA4ThahgEz6gODyzDoqwdrEF3%2FHuuywrsqWt2LlF5ztGGmTIW7J4a6gsYd3hG7OpkkAjEAzdgHgb7xfB23Wpd6qbl1s5maUYVhoNado2Dh2tVhraUOfteRQyPCXqaVByDsEg6A)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 800J', but common name is 'Plivo Inc' |
| [e_atis_ext_certificate_policies](../../ISSUES/e_atis_ext_certificate_policies/README.md) | error | ATIS1000080 | The Certificate Policies extension is marked as critical |
| [e_atis_serial_number_size](../../ISSUES/e_atis_serial_number_size/README.md) | error | ATIS1000080 | STI certificates shall have a serial number that contains at least 64 bits, got 63 |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common Name attribute 'Plivo Inc' does not contain 'SHAKEN' |
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension contains an invalid OID value: 2.16.840.1.114569.1.1.1. Available OIDs: 2.16.840.1.114569.1.1.3, 2.16.840.1.114569.1.1.4 |


Generated: 04 Oct 24 16:29 UTC