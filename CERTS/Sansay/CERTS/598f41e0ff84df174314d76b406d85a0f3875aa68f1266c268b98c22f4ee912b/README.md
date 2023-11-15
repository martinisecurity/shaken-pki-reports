# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Call Hub Inc. 688K

Tested At: 15 Nov 23 16:09 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 269 day(s)\
Subject: CN=SHAKEN Call Hub Inc. 688K, OU=Technical, O=Call Hub Inc., ST=Wyoming, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://callhub.46labs.com/callhub.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC0zCCAnmgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkfJswCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDgxMTExNDUwNFoXDTI0MDgxMDExNDUwNFowbzELMAkGA1UEBhMCVVMxEDAOBgNVBAgMB1d5b21pbmcxFjAUBgNVBAoMDUNhbGwgSHViIEluYy4xEjAQBgNVBAsMCVRlY2huaWNhbDEiMCAGA1UEAwwZU0hBS0VOIENhbGwgSHViIEluYy4gNjg4SzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABAEyrWu2rndHWsjL5IvpNvX6xguDiXtaRh6J3SkpfpnfWhOGZXliTwC6QjUoVVADOkE4wqIYVwebUw7RsNAhM%2BWjgdswgdgwFgYIKwYBBQUHARoECjAIoAYWBDY4OEswFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMB0GA1UdDgQWBBRDY3gAtPF5qAmgVtWW%2F6JG58PQ8zAfBgNVHSMEGDAWgBSs05P1Q0PMCr5FWBcTfZJ83MMBRjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSAAwRQIgcctrzh8pfloIitBZh8YVahhpUeA2pOxHw2W1YUPMLOcCIQCsZbcAbACxBxJsjnaF3nX%2FveBLMkfTkPRDhsY1Ef0E7w%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 688K' |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 15 Nov 23 17:17 UTC