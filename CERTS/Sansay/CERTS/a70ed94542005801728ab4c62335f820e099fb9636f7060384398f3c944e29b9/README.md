# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Momentum Telecom 1417

Tested At: 28 Nov 23 10:43 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 10 day(s)\
Subject: CN=SHAKEN Momentum Telecom 1417, OU=NOC, O=Momentum Telecom, ST=Georgia, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/MomentumTelecom_1417

[View certificate details](https://understandingwebpki.com/?cert=MIIC0jCCAnmgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkhgUwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMTEwNzE2NDcyNloXDTIzMTIwNzE2NDcyNlowbzELMAkGA1UEBhMCVVMxEDAOBgNVBAgMB0dlb3JnaWExGTAXBgNVBAoMEE1vbWVudHVtIFRlbGVjb20xDDAKBgNVBAsMA05PQzElMCMGA1UEAwwcU0hBS0VOIE1vbWVudHVtIFRlbGVjb20gMTQxNzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABII7HUTol1f%2FqSE%2B%2Faw3uBpttsdkbcNLNe4GTT6QNc1p5xhfUjCm8EJ1xtaeemvakN4DOVzQkrxtJiYyjdmnnLSjgdswgdgwFgYIKwYBBQUHARoECjAIoAYWBDE0MTcwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMB0GA1UdDgQWBBT5vvlSdrpNc3W4WbWUJC6Zwu8MZDAfBgNVHSMEGDAWgBSs05P1Q0PMCr5FWBcTfZJ83MMBRjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDRwAwRAIgNj33wE0kldqWgNaQNHQUeyPjNBHayArGIm4eZ%2F8ol58CIES6%2FXD3iTlVrhBCFox%2BS7EMTpoUonPDj9KbXJcbwT33)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 1417', but common name is 'SHAKEN Momentum Telecom 1417' |
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension contains an invalid OID value: 2.16.840.1.114569.1.1.1. Available OIDs: 2.16.840.1.114569.1.1.3, 2.16.840.1.114569.1.1.4 |


Generated: 28 Nov 23 10:53 UTC