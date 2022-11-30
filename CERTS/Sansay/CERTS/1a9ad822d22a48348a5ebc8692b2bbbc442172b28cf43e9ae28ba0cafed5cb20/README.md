# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Nobelbiz, Inc. 596J

Tested At: 30 Nov 22 15:53 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 16 day(s)\
Subject: CN=SHAKEN Nobelbiz\\, Inc. 596J, OU=NOC, O=Nobelbiz\\, Inc., ST=California, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/NobelBiz_596J

[View certificate details](https://understandingwebpki.com/?cert=MIIC0zCCAnigAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkVY0wCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTExNTIxMTA0MVoXDTIyMTIxNTIxMTA0MVowbjELMAkGA1UEBhMCVVMxEzARBgNVBAgMCkNhbGlmb3JuaWExFzAVBgNVBAoMDk5vYmVsYml6LCBJbmMuMQwwCgYDVQQLDANOT0MxIzAhBgNVBAMMGlNIQUtFTiBOb2JlbGJpeiwgSW5jLiA1OTZKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE7OQCe1Y3BvkHYtREe1szJLq%2BVq8zqxgCxdGcg8OLJaM803%2BsjXxzlaepdHYiOJBt31qvYvRBSsvCI%2BqCXvI6E6OB2zCB2DAWBggrBgEFBQcBGgQKMAigBhYENTk2SjAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFLdgZtiViQnXFZDWR2Mbmmm5hKLDMB8GA1UdIwQYMBaAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNJADBGAiEAw5db%2BHiRJwPv3eHuRWjfArlYIMI1jtfk89L1wudH82ECIQDoX9yHmiT8iHlqRo8uFiWzzP9Frq1JY4raepgH7vwQsw%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 596J' |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |


Generated: 30 Nov 22 16:07 UTC