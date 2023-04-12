# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Primo Dialler LLC 249K

Tested At: 12 Apr 23 21:42 UTC\
Initial Validity Period: 45 day(s)\
Remaining Validity Period: 15 day(s)\
Subject: CN=SHAKEN Primo Dialler LLC 249K, OU=Service Provider, O=Primo Dialler LLC, ST=Wyoming, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://app.connexcs.com/api/stir-shaken/cert/45.crt

[View certificate details](https://understandingwebpki.com/?cert=MIIC4jCCAoigAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkZL8wCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDMxMzA2Mzk1MFoXDTIzMDQyNzA2Mzk1MFowfjELMAkGA1UEBhMCVVMxEDAOBgNVBAgMB1d5b21pbmcxGjAYBgNVBAoMEVByaW1vIERpYWxsZXIgTExDMRkwFwYDVQQLDBBTZXJ2aWNlIFByb3ZpZGVyMSYwJAYDVQQDDB1TSEFLRU4gUHJpbW8gRGlhbGxlciBMTEMgMjQ5SzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABBIz0KsI91%2Bi8SaDwi0dW1fNpm1KyLI2V9yRq2ZQ48%2F3LRkSXW9WRBi21W30HDBOGd%2BxfqYf%2F5PUFvEm7F%2BKIVqjgdswgdgwFgYIKwYBBQUHARoECjAIoAYWBDI0OUswFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMB0GA1UdDgQWBBSAyc%2Bis0SMvx%2FHSOAXczRBKJ46BzAfBgNVHSMEGDAWgBSs05P1Q0PMCr5FWBcTfZJ83MMBRjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSAAwRQIgRJqESjhzTc8B4H3WgqQOXoJssCGFdSKvq3p7z0KHaAUCIQDTs%2Fsz8XRjT0jB%2FP6r0OOZjntM9Xxwu%2BrT6kNwnbhvJw%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 249K' |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 12 Apr 23 22:02 UTC