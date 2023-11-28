# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN BareTelecom 864J

Tested At: 28 Nov 23 16:05 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -80 day(s)\
Subject: emailAddress=chris@baretelecom.com, CN=SHAKEN BareTelecom 864J, OU=Operations, O=BareTelecom, ST=California, C=US, emailAddress=chris@baretelecom.com\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/864J/order/143_864J_159

[View certificate details](https://understandingwebpki.com/?cert=MIIC%2BjCCAqCgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkfGswCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDgwOTE5NTgzOVoXDTIzMDkwODE5NTgzOVowgZUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRQwEgYDVQQKDAtCYXJlVGVsZWNvbTETMBEGA1UECwwKT3BlcmF0aW9uczEgMB4GA1UEAwwXU0hBS0VOIEJhcmVUZWxlY29tIDg2NEoxJDAiBgkqhkiG9w0BCQEWFWNocmlzQGJhcmV0ZWxlY29tLmNvbTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABDDS7l0vaaxRo5PuD%2BPqVhVve6XeHWGMrTrmV563iRvazHazHmoMABAUKMD0FgTQtsh5TKpUQdvJvnJwByVoCaGjgdswgdgwFgYIKwYBBQUHARoECjAIoAYWBDg2NEowFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMB0GA1UdDgQWBBRIdUFl6KvT9XyIaPknmcYnWIKCbzAfBgNVHSMEGDAWgBSs05P1Q0PMCr5FWBcTfZJ83MMBRjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSAAwRQIhALUnk3T%2BKxs6Ax4f90K%2BxNI%2B3oUmP5a4YKQbh3fSBOnlAiBhj4Q0vWNNEPPxdG1c6QZk5sQtsxfmMTNxCMKj3KhU5Q%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 864J', but common name is 'SHAKEN BareTelecom 864J' |


Generated: 28 Nov 23 16:15 UTC