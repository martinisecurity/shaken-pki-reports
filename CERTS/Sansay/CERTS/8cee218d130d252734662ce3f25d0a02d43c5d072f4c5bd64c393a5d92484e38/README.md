# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN PNG Telecommunications Inc 3395

Tested At: 18 Aug 25 20:27 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -382 day(s)\
Subject: CN=SHAKEN PNG Telecommunications Inc 3395, OU=NOC, O=PNG Telecommunications Inc, ST=Ohio, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/3395/429C7C70711E3820F0B8E1DEAE6FF3262264A210.pem

[View certificate details](https://x509.io/?cert=MIIC5jCCAougAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkohAwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTI0MDcwMjE3MTcyMFoXDTI0MDgwMTE3MTcyMFowgYAxCzAJBgNVBAYTAlVTMQ0wCwYDVQQIDARPaGlvMSMwIQYDVQQKDBpQTkcgVGVsZWNvbW11bmljYXRpb25zIEluYzEMMAoGA1UECwwDTk9DMS8wLQYDVQQDDCZTSEFLRU4gUE5HIFRlbGVjb21tdW5pY2F0aW9ucyBJbmMgMzM5NTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABMwNyaEeWA0vGW0AzW6J40Tgty3ddjhlKHaR5Fbaaw3hgWwULA51dFB5gSMg2n4dJJVznCBP85MHY%2BZzFAPHVQKjgdswgdgwFgYIKwYBBQUHARoECjAIoAYWBDMzOTUwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMB0GA1UdDgQWBBT6etr4H3xed5Sv%2FS9%2FyUk43joO4TAfBgNVHSMEGDAWgBSs05P1Q0PMCr5FWBcTfZJ83MMBRjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSQAwRgIhALY4rHbiPkXqdwLwdC%2FOkQxMaKmyVVxlbYvjDxR79DtOAiEA4cjpyjTL51f5sJ9Fa6tCGv2gbndJZGCvD1XV9JbGC9I%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 3395', but common name is 'SHAKEN PNG Telecommunications Inc 3395' |
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension contains an invalid OID value: 2.16.840.1.114569.1.1.1. Available OIDs: 2.16.840.1.114569.1.1.3, 2.16.840.1.114569.1.1.4 |


Generated: 18 Aug 25 21:13 UTC