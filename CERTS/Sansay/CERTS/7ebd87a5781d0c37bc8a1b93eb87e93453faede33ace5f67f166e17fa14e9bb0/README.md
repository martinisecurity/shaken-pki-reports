# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Noble Systems Communications LLC 187J

Tested At: 18 Aug 25 20:22 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -408 day(s)\
Subject: CN=SHAKEN Noble Systems Communications LLC 187J, OU=NOC, O=Noble Systems Communications LLC, ST=Georgia, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/187J/429C7C70711E3820F0B8E1DEAE6FF32622649EF7.pem

[View certificate details](https://x509.io/?cert=MIIC9DCCApqgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJknvcwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTI0MDYwNTIzMDM1M1oXDTI0MDcwNTIzMDM1M1owgY8xCzAJBgNVBAYTAlVTMRAwDgYDVQQIDAdHZW9yZ2lhMSkwJwYDVQQKDCBOb2JsZSBTeXN0ZW1zIENvbW11bmljYXRpb25zIExMQzEMMAoGA1UECwwDTk9DMTUwMwYDVQQDDCxTSEFLRU4gTm9ibGUgU3lzdGVtcyBDb21tdW5pY2F0aW9ucyBMTEMgMTg3SjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABP%2FmGlJuj%2FstdPV1ajhIPkNsLxwZQGX6idMO%2Fp8KZhhsDFAajP4jO2H2H5v6dRG3QkxY459oHsT7Ovh3AZPOq46jgdswgdgwFgYIKwYBBQUHARoECjAIoAYWBDE4N0owFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMB0GA1UdDgQWBBROIbMJwncdkQ3fRdiI1rOMaa87hDAfBgNVHSMEGDAWgBSs05P1Q0PMCr5FWBcTfZJ83MMBRjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSAAwRQIgVwVmiBL91vjWD7Pq%2FEOBryWG3hgOjYGPJ0bqNQgJR7wCIQCtLjO6bSnJreV8udny%2Byi2gVcfbYt2J4TMUCgUJTjr2g%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 187J', but common name is 'SHAKEN Noble Systems Communications LLC 187J' |
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension contains an invalid OID value: 2.16.840.1.114569.1.1.1. Available OIDs: 2.16.840.1.114569.1.1.3, 2.16.840.1.114569.1.1.4 |


Generated: 18 Aug 25 21:13 UTC