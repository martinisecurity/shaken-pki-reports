# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN NTC International, INC 016K

Tested At: 21 Aug 23 20:09 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -124 day(s)\
Subject: emailAddress=billing@ntcinternationalinc.com, CN=SHAKEN NTC International\\, INC 016K, OU=Operations, O=NTC International\\, INC, ST=Delaware, C=US, emailAddress=billing@ntcinternationalinc.com\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/016K/order/222_016K_115

[View certificate details](https://understandingwebpki.com/?cert=MIIDGDCCAr6gAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkZfowCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDMyMDE4MDcxMFoXDTIzMDQxOTE4MDcxMFowgbMxCzAJBgNVBAYTAlVTMREwDwYDVQQIDAhEZWxhd2FyZTEfMB0GA1UECgwWTlRDIEludGVybmF0aW9uYWwsIElOQzETMBEGA1UECwwKT3BlcmF0aW9uczErMCkGA1UEAwwiU0hBS0VOIE5UQyBJbnRlcm5hdGlvbmFsLCBJTkMgMDE2SzEuMCwGCSqGSIb3DQEJARYfYmlsbGluZ0BudGNpbnRlcm5hdGlvbmFsaW5jLmNvbTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABM51nAzTItLw89uHu4KL0XTDPUzsmH0Ds0Hbi5F2DwFVaiKK5RvFSodn8j5IjXI49giYgMA0uPagzDbZoRkuPbmjgdswgdgwFgYIKwYBBQUHARoECjAIoAYWBDAxNkswFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMB0GA1UdDgQWBBS8ynIC3AnB8%2BdsxYDDelrP6XHHoTAfBgNVHSMEGDAWgBSs05P1Q0PMCr5FWBcTfZJ83MMBRjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSAAwRQIhAJ4GxpEPoib4FSfdkcTJT4T3RvTsy1nImVDuM9CMQiYDAiATQeLv5coZIJNKl%2B3I9CtgmIBZipoUzhHI4jq%2BJW6XZA%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [w_cp_1_3_subject_email](../../ISSUES/w_cp_1_3_subject_email/README.md) | warn | US_SHAKEN_CP | Email addresses are not allowed as the CP does not specify how to validate them |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 016K' |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 21 Aug 23 20:18 UTC