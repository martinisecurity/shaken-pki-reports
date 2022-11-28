# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Primo Dialler LLC 249K

Tested At: 28 Nov 22 20:37 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 11 day(s)\
Subject: CN=SHAKEN Primo Dialler LLC 249K, OU=Service Provider, O=Primo Dialler LLC, ST=Wyoming, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://app.connexcs.com/api/stir-shaken/cert/45.crt

[View certificate details](https://understandingwebpki.com/?cert=MIIC4TCCAoigAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkVPwwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTEwOTE4NDkwNFoXDTIyMTIwOTE4NDkwNFowfjELMAkGA1UEBhMCVVMxEDAOBgNVBAgMB1d5b21pbmcxGjAYBgNVBAoMEVByaW1vIERpYWxsZXIgTExDMRkwFwYDVQQLDBBTZXJ2aWNlIFByb3ZpZGVyMSYwJAYDVQQDDB1TSEFLRU4gUHJpbW8gRGlhbGxlciBMTEMgMjQ5SzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABLTF9FomrzoaGS9Ie%2BWG9dl%2By1l%2BvHywMGC9BtD%2FuCR6wOsW6zoTJFgQrVpqy7zeilLFkdVUC%2B0PRWnWV1RfBqmjgdswgdgwFgYIKwYBBQUHARoECjAIoAYWBDI0OUswFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMB0GA1UdDgQWBBS5kQR%2BOKKbE2Izi9U7u%2F%2BiTPmKqzAfBgNVHSMEGDAWgBSs05P1Q0PMCr5FWBcTfZJ83MMBRjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDRwAwRAIgLvsl6UtSLhIER8EQjMj%2Fx1rqDCzRzkjGAHf1MVrEiaQCIGUkkFB5RS5osNlnQFzSOUU9rVbTVSooBnxdYg8zxV1N)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 249K' |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |


Generated: 28 Nov 22 20:41 UTC