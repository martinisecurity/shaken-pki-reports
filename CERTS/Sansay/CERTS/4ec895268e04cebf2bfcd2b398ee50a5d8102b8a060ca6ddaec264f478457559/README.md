# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Momentum Telecom 9157

Tested At: 27 Nov 23 23:20 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 10 day(s)\
Subject: CN=SHAKEN Momentum Telecom 9157, OU=NOC, O=Momentum Telecom, ST=Georgia, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/MomentumTelecom_9157

[View certificate details](https://understandingwebpki.com/?cert=MIIC0jCCAnmgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkhgYwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMTEwNzE2NDc0N1oXDTIzMTIwNzE2NDc0N1owbzELMAkGA1UEBhMCVVMxEDAOBgNVBAgMB0dlb3JnaWExGTAXBgNVBAoMEE1vbWVudHVtIFRlbGVjb20xDDAKBgNVBAsMA05PQzElMCMGA1UEAwwcU0hBS0VOIE1vbWVudHVtIFRlbGVjb20gOTE1NzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABMAoR2hhlAdFCEs0FU4ijCdcMKOr8L38LGFGNqTxndGpUVfQTjvjgl1cFmyFxp2fb1vOAVoMqT4Q5Tz8Br3O%2BQejgdswgdgwFgYIKwYBBQUHARoECjAIoAYWBDkxNTcwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMB0GA1UdDgQWBBQ%2Fof%2BXyLnO4yitFukJpfvLLFMnQzAfBgNVHSMEGDAWgBSs05P1Q0PMCr5FWBcTfZJ83MMBRjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDRwAwRAIgCh2GDtUq6wJrMdN0sS1VQnukCYnuikohiCdM7tXHkUACIEk7MYz8KrA5TKJRtu%2BLq4eASfCdD8z4UFg%2FiBJfx0vj)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 9157', but common name is 'SHAKEN Momentum Telecom 9157' |
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension contains an invalid OID value: 2.16.840.1.114569.1.1.1. Available OIDs: 2.16.840.1.114569.1.1.3, 2.16.840.1.114569.1.1.4 |


Generated: 27 Nov 23 23:28 UTC