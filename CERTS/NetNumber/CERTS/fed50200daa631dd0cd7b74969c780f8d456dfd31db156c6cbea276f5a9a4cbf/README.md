# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Plivo Inc

Tested At: 21 Aug 23 20:01 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 270 day(s)\
Subject: L=Austin, ST=Texas, O=Plivo Inc, C=US, CN=Plivo Inc\
Issuer: L=Lowell, ST=Massachusettes, OU=US, O=NetNumber Inc, C=US, CN=NetNumber SHAKEN Root Intermediate CA 1\
Link: https://certificate.zt.plivo.com/cert20240517.crt

[View certificate details](https://understandingwebpki.com/?cert=MIIC2jCCAmCgAwIBAgIIDl5HHwawdMMwCgYIKoZIzj0EAwIwgY4xMDAuBgNVBAMMJ05ldE51bWJlciBTSEFLRU4gUm9vdCBJbnRlcm1lZGlhdGUgQ0EgMTELMAkGA1UEBhMCVVMxFjAUBgNVBAoMDU5ldE51bWJlciBJbmMxCzAJBgNVBAsMAlVTMRcwFQYDVQQIDA5NYXNzYWNodXNldHRlczEPMA0GA1UEBwwGTG93ZWxsMB4XDTIzMDUxODAwMDAwMFoXDTI0MDUxNzAwMDAwMFowVjESMBAGA1UEAwwJUGxpdm8gSW5jMQswCQYDVQQGEwJVUzESMBAGA1UECgwJUGxpdm8gSW5jMQ4wDAYDVQQIDAVUZXhhczEPMA0GA1UEBwwGQXVzdGluMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAER%2BDMfKslaXg6AONm1uDq0SvPgtKQlt1WjVkmSVrr%2FRJ6N2NNwWSmXYI2BIb0UH0vTXyJnq5aYAz1qKFJe3tasqOB3jCB2zAWBggrBgEFBQcBGgQKMAigBhYEODAwSjAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAfBgNVHSMEGDAWgBRxL8iC3KjgIuPfoGj5%2BF5chN7lvTAdBgNVHQ4EFgQUtkYGfjB06qkxy1FF7dsOhIGQd5swRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMBoGA1UdIAEB%2FwQQMA4wDAYKYIZIAYb%2FCQEBATAKBggqhkjOPQQDAgNoADBlAjEA195s%2Fg3NNX%2FjuQZIRVN33B8uygUIbzyRLiab42hWoRyUBYWUdAEv1%2FV%2FHkcUCRSLAjB9K864t5SRc%2FbIJfxzpZXfdiRxfpJqT6YJ5vXCNtT40O%2FQuW9jbjrOTaMO7ZRaAeY%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 800J' |
| [n_atis_certificate_policy_critical](../../ISSUES/n_atis_certificate_policy_critical/README.md) | notice | ATIS1000080 | STI certificates should contain a CertificatePolicies extension marked uncritical |


Generated: 21 Aug 23 20:18 UTC