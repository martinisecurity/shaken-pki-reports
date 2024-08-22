# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Nobelbiz, Inc. 596J

Tested At: 22 Aug 24 15:54 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -195 day(s)\
Subject: CN=SHAKEN Nobelbiz\\, Inc. 596J, OU=NOC, O=Nobelbiz\\, Inc., ST=California, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/NobelBiz_596J

[View certificate details](https://x509.io/?cert=MIIC0TCCAnigAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkjVEwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTI0MDEwOTIwNTUyNFoXDTI0MDIwODIwNTUyNFowbjELMAkGA1UEBhMCVVMxEzARBgNVBAgMCkNhbGlmb3JuaWExFzAVBgNVBAoMDk5vYmVsYml6LCBJbmMuMQwwCgYDVQQLDANOT0MxIzAhBgNVBAMMGlNIQUtFTiBOb2JlbGJpeiwgSW5jLiA1OTZKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE7OQCe1Y3BvkHYtREe1szJLq%2BVq8zqxgCxdGcg8OLJaM803%2BsjXxzlaepdHYiOJBt31qvYvRBSsvCI%2BqCXvI6E6OB2zCB2DAWBggrBgEFBQcBGgQKMAigBhYENTk2SjAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFLdgZtiViQnXFZDWR2Mbmmm5hKLDMB8GA1UdIwQYMBaAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNHADBEAiB3c9B2fI5UhShwB7q2lYYI3%2BTIKaBXzT7D0HI%2FOl03%2FwIgIyp0gBYUGIe8g%2FqI6V2xjs%2B7JdKax%2BgoLiTzRlSFRxk%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 596J', but common name is 'SHAKEN Nobelbiz, Inc. 596J' |
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension contains an invalid OID value: 2.16.840.1.114569.1.1.1. Available OIDs: 2.16.840.1.114569.1.1.3, 2.16.840.1.114569.1.1.4 |


Generated: 22 Aug 24 16:06 UTC