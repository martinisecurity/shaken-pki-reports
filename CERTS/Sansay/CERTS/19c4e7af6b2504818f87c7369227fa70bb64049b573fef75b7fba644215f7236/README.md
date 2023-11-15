# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN IP Link Telecom Inc. 902J

Tested At: 15 Nov 23 16:34 UTC\
Initial Validity Period: 90 day(s)\
Remaining Validity Period: -36 day(s)\
Subject: emailAddress=ops@iplinktelecom.com, CN=SHAKEN IP Link Telecom Inc. 902J, OU=IP Link Telecom Inc., O=IP Link Telecom Inc., ST=Oregon, C=US, emailAddress=ops@iplinktelecom.com\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/902J/order/52_902J_68

[View certificate details](https://understandingwebpki.com/?cert=MIIDEzCCArigAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkeXwwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDcxMjE0NTgxNFoXDTIzMTAxMDE0NTgxNFowga0xCzAJBgNVBAYTAlVTMQ8wDQYDVQQIDAZPcmVnb24xHTAbBgNVBAoMFElQIExpbmsgVGVsZWNvbSBJbmMuMR0wGwYDVQQLDBRJUCBMaW5rIFRlbGVjb20gSW5jLjEpMCcGA1UEAwwgU0hBS0VOIElQIExpbmsgVGVsZWNvbSBJbmMuIDkwMkoxJDAiBgkqhkiG9w0BCQEWFW9wc0BpcGxpbmt0ZWxlY29tLmNvbTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABKywISh8J2PUUdpZsF%2BM2%2BRWcVP1XuuMhDmThS%2B7iFdGWGDS1C%2FxwZCao%2Fb88w2FpB8RNUXmEnm7lVCSDpE5sfajgdswgdgwFgYIKwYBBQUHARoECjAIoAYWBDkwMkowFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMB0GA1UdDgQWBBQkWHVUI6i7xW0DodBk2Qo9I1yMeTAfBgNVHSMEGDAWgBSs05P1Q0PMCr5FWBcTfZJ83MMBRjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSQAwRgIhAMTQWBQykqkHKo4SkNt07z874jJCRwI5U%2FJIk%2BHf8QkkAiEAl6H29VW7lRm5Ak8z9e12FcFuvdiFuTjuUnvnO6zNhFc%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 902J' |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_cp_1_3_subject_email](../../ISSUES/w_cp_1_3_subject_email/README.md) | warn | US_SHAKEN_CP | Email addresses are not allowed as the CP does not specify how to validate them |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 15 Nov 23 17:17 UTC