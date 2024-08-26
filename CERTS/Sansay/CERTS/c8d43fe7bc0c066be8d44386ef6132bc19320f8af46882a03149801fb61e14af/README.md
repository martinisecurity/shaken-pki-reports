# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN BareTelecom 864J

Tested At: 26 Aug 24 18:34 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -446 day(s)\
Subject: emailAddress=chris@baretelecom.com, CN=SHAKEN BareTelecom 864J, OU=Operations, O=BareTelecom, ST=California, C=US, emailAddress=chris@baretelecom.com\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/864J/order/53_864J_159

[View certificate details](https://x509.io/?cert=MIIC%2BjCCAqCgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkbqIwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDUwODA4Mjc1NFoXDTIzMDYwNzA4Mjc1NFowgZUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRQwEgYDVQQKDAtCYXJlVGVsZWNvbTETMBEGA1UECwwKT3BlcmF0aW9uczEgMB4GA1UEAwwXU0hBS0VOIEJhcmVUZWxlY29tIDg2NEoxJDAiBgkqhkiG9w0BCQEWFWNocmlzQGJhcmV0ZWxlY29tLmNvbTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABDDS7l0vaaxRo5PuD%2BPqVhVve6XeHWGMrTrmV563iRvazHazHmoMABAUKMD0FgTQtsh5TKpUQdvJvnJwByVoCaGjgdswgdgwFgYIKwYBBQUHARoECjAIoAYWBDg2NEowFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMB0GA1UdDgQWBBRIdUFl6KvT9XyIaPknmcYnWIKCbzAfBgNVHSMEGDAWgBSs05P1Q0PMCr5FWBcTfZJ83MMBRjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSAAwRQIgc383iV%2Fx4xsNCrX5joMrmCTuifQEcbwXLKayG%2FoObCACIQDv1k8Wiua4P6vae0X%2FtgGn4eKZrWTC%2FSXR7obA3rII%2FQ%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 864J', but common name is 'SHAKEN BareTelecom 864J' |
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |


Generated: 26 Aug 24 18:49 UTC