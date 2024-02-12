# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Telxio Networks 492K

Tested At: 12 Feb 24 16:53 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 66 day(s)\
Subject: CN=SHAKEN Telxio Networks 492K, OU=AS-Unit, O=Telxio Networks, ST=Nevada, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/ss-telxio

[View certificate details](https://understandingwebpki.com/?cert=MIIC0zCCAnqgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJka2kwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDQxOTA2MDcxMloXDTI0MDQxODA2MDcxMlowcDELMAkGA1UEBhMCVVMxDzANBgNVBAgMBk5ldmFkYTEYMBYGA1UECgwPVGVseGlvIE5ldHdvcmtzMRAwDgYDVQQLDAdBUy1Vbml0MSQwIgYDVQQDDBtTSEFLRU4gVGVseGlvIE5ldHdvcmtzIDQ5MkswWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAT%2FwsSU4U5cpWCWeme48YYpDI9l310ZbgWStFFNYwDGNmCSy4SJa8ZjOe3ETmJWGHr65o6c5UAtNmhk2Xldhs4wo4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQ0OTJLMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzAdBgNVHQ4EFgQUhxX4cg1nrgmbxZDdOGfi69zo6w0wHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0cAMEQCIEbnogrk%2Frb08HCijTEAPTmS1LnMB9lCrgOyPmrSFaIBAiA7ShZWbINRZHEJvH3Fdojnk%2BD18MxWt3hK89vn3v4jGg%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 492K', but common name is 'SHAKEN Telxio Networks 492K' |


Generated: 12 Feb 24 17:02 UTC