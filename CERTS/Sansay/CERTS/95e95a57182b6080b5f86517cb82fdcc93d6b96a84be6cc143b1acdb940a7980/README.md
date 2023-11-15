# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Televergence Solutions Inc 779J

Tested At: 15 Nov 23 16:07 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -118 day(s)\
Subject: emailAddress=ddeutsch@televergence.com, CN=SHAKEN Televergence Solutions Inc 779J, OU=Production, O=Televergence Solutions Inc, ST=Tennessee, C=US, emailAddress=ddeutsch@televergence.com\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/779J/order/334_779J_73

[View certificate details](https://understandingwebpki.com/?cert=MIIDGzCCAsGgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkde4wCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDYyMDExMjkwMloXDTIzMDcyMDExMjkwMlowgbYxCzAJBgNVBAYTAlVTMRIwEAYDVQQIDAlUZW5uZXNzZWUxIzAhBgNVBAoMGlRlbGV2ZXJnZW5jZSBTb2x1dGlvbnMgSW5jMRMwEQYDVQQLDApQcm9kdWN0aW9uMS8wLQYDVQQDDCZTSEFLRU4gVGVsZXZlcmdlbmNlIFNvbHV0aW9ucyBJbmMgNzc5SjEoMCYGCSqGSIb3DQEJARYZZGRldXRzY2hAdGVsZXZlcmdlbmNlLmNvbTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABI7YJeB2zUiDx0TCzLx53mIokjf1Zw8pZBRXCiyTzdvxaZ9YPFSsOsaqudnIF5bZ9e47vkgQvlGQsZKxR%2Fvi60ujgdswgdgwFgYIKwYBBQUHARoECjAIoAYWBDc3OUowFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMB0GA1UdDgQWBBQAf2WVYtjP2AE%2F07vmJrW95Y60TTAfBgNVHSMEGDAWgBSs05P1Q0PMCr5FWBcTfZJ83MMBRjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSAAwRQIhAI9W9WZ%2FXFR0313nOnPTjPC9YM1%2B4ZtO%2FwDyqRCQGSuiAiA9gf0M%2F9ieKHPboJBKavOoilR1IBGW7ZFs1dossUZ8Uw%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_cp_1_3_subject_email](../../ISSUES/w_cp_1_3_subject_email/README.md) | warn | US_SHAKEN_CP | Email addresses are not allowed as the CP does not specify how to validate them |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 779J' |


Generated: 15 Nov 23 16:51 UTC