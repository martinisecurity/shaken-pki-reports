# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Momentum Telecom 9157

Tested At: 04 Oct 24 16:13 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -90 day(s)\
Subject: CN=SHAKEN Momentum Telecom 9157, OU=NOC, O=Momentum Telecom, ST=Georgia, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/9157/429C7C70711E3820F0B8E1DEAE6FF32622649EFD.pem

[View certificate details](https://x509.io/?cert=MIIC1DCCAnmgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJknv0wCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTI0MDYwNTIzMTQ0OFoXDTI0MDcwNTIzMTQ0OFowbzELMAkGA1UEBhMCVVMxEDAOBgNVBAgMB0dlb3JnaWExGTAXBgNVBAoMEE1vbWVudHVtIFRlbGVjb20xDDAKBgNVBAsMA05PQzElMCMGA1UEAwwcU0hBS0VOIE1vbWVudHVtIFRlbGVjb20gOTE1NzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABMAoR2hhlAdFCEs0FU4ijCdcMKOr8L38LGFGNqTxndGpUVfQTjvjgl1cFmyFxp2fb1vOAVoMqT4Q5Tz8Br3O%2BQejgdswgdgwFgYIKwYBBQUHARoECjAIoAYWBDkxNTcwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMB0GA1UdDgQWBBQ%2Fof%2BXyLnO4yitFukJpfvLLFMnQzAfBgNVHSMEGDAWgBSs05P1Q0PMCr5FWBcTfZJ83MMBRjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSQAwRgIhAMTz5z4jYEcg0KtEJXTDSLW%2BsnMTCjyi459Tx%2FIRDWg7AiEAy0CtxlNTtAwuF2fYlY97XDI2A6Fos2v92J1dLCfUua8%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 9157', but common name is 'SHAKEN Momentum Telecom 9157' |
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension contains an invalid OID value: 2.16.840.1.114569.1.1.1. Available OIDs: 2.16.840.1.114569.1.1.3, 2.16.840.1.114569.1.1.4 |


Generated: 04 Oct 24 16:29 UTC