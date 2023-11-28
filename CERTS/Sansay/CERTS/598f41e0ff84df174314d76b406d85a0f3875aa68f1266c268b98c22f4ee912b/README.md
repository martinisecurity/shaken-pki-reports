# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Call Hub Inc. 688K

Tested At: 28 Nov 23 19:52 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 256 day(s)\
Subject: CN=SHAKEN Call Hub Inc. 688K, OU=Technical, O=Call Hub Inc., ST=Wyoming, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://callhub.46labs.com/callhub.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC0zCCAnmgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkfJswCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDgxMTExNDUwNFoXDTI0MDgxMDExNDUwNFowbzELMAkGA1UEBhMCVVMxEDAOBgNVBAgMB1d5b21pbmcxFjAUBgNVBAoMDUNhbGwgSHViIEluYy4xEjAQBgNVBAsMCVRlY2huaWNhbDEiMCAGA1UEAwwZU0hBS0VOIENhbGwgSHViIEluYy4gNjg4SzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABAEyrWu2rndHWsjL5IvpNvX6xguDiXtaRh6J3SkpfpnfWhOGZXliTwC6QjUoVVADOkE4wqIYVwebUw7RsNAhM%2BWjgdswgdgwFgYIKwYBBQUHARoECjAIoAYWBDY4OEswFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMB0GA1UdDgQWBBRDY3gAtPF5qAmgVtWW%2F6JG58PQ8zAfBgNVHSMEGDAWgBSs05P1Q0PMCr5FWBcTfZJ83MMBRjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSAAwRQIgcctrzh8pfloIitBZh8YVahhpUeA2pOxHw2W1YUPMLOcCIQCsZbcAbACxBxJsjnaF3nX%2FveBLMkfTkPRDhsY1Ef0E7w%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 688K', but common name is 'SHAKEN Call Hub Inc. 688K' |


Generated: 28 Nov 23 20:21 UTC