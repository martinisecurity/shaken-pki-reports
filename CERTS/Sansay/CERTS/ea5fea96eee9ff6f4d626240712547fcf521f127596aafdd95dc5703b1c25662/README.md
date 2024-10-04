# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN NTC International, INC 016K

Tested At: 04 Oct 24 15:54 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -373 day(s)\
Subject: emailAddress=billing@ntcinternationalinc.com, CN=SHAKEN NTC International\\, INC 016K, OU=Operations, O=NTC International\\, INC, ST=Delaware, C=US, emailAddress=billing@ntcinternationalinc.com\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/016K/order/383_016K_115

[View certificate details](https://x509.io/?cert=MIIDGTCCAr6gAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkfmswCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDgyNzIxMjU1OFoXDTIzMDkyNjIxMjU1OFowgbMxCzAJBgNVBAYTAlVTMREwDwYDVQQIDAhEZWxhd2FyZTEfMB0GA1UECgwWTlRDIEludGVybmF0aW9uYWwsIElOQzETMBEGA1UECwwKT3BlcmF0aW9uczErMCkGA1UEAwwiU0hBS0VOIE5UQyBJbnRlcm5hdGlvbmFsLCBJTkMgMDE2SzEuMCwGCSqGSIb3DQEJARYfYmlsbGluZ0BudGNpbnRlcm5hdGlvbmFsaW5jLmNvbTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABM51nAzTItLw89uHu4KL0XTDPUzsmH0Ds0Hbi5F2DwFVaiKK5RvFSodn8j5IjXI49giYgMA0uPagzDbZoRkuPbmjgdswgdgwFgYIKwYBBQUHARoECjAIoAYWBDAxNkswFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMB0GA1UdDgQWBBS8ynIC3AnB8%2BdsxYDDelrP6XHHoTAfBgNVHSMEGDAWgBSs05P1Q0PMCr5FWBcTfZJ83MMBRjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSQAwRgIhANFgyshPrXpeh7rE8hKzqdx3USlLN7xDVZG%2ByNbIq%2BL6AiEA2IEIiuG9r%2BFc8o5ArL%2Fsit1q4A0v41T4rB%2FKfz92JFk%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 016K', but common name is 'SHAKEN NTC International, INC 016K' |
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |


Generated: 04 Oct 24 16:29 UTC