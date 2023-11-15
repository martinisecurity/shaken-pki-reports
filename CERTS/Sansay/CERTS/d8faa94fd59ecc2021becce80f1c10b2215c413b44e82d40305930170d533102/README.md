# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN IP Link Telecom Inc. 902J

Tested At: 15 Nov 23 18:03 UTC\
Initial Validity Period: 90 day(s)\
Remaining Validity Period: -138 day(s)\
Subject: emailAddress=ops@iplinktelecom.com, CN=SHAKEN IP Link Telecom Inc. 902J, OU=IP Link Telecom Inc., O=IP Link Telecom Inc., ST=Oregon, C=US, emailAddress=ops@iplinktelecom.com\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/902J/order/49_902J_68

[View certificate details](https://understandingwebpki.com/?cert=MIIDEjCCArigAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkaEIwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDQwMTAwMDAyOFoXDTIzMDYzMDAwMDAyOFowga0xCzAJBgNVBAYTAlVTMQ8wDQYDVQQIDAZPcmVnb24xHTAbBgNVBAoMFElQIExpbmsgVGVsZWNvbSBJbmMuMR0wGwYDVQQLDBRJUCBMaW5rIFRlbGVjb20gSW5jLjEpMCcGA1UEAwwgU0hBS0VOIElQIExpbmsgVGVsZWNvbSBJbmMuIDkwMkoxJDAiBgkqhkiG9w0BCQEWFW9wc0BpcGxpbmt0ZWxlY29tLmNvbTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABKQuuhkJ9q43tJzmgt8U1v2%2B%2B6BCvtDco%2B8eYUFhygN6JC4pQW4ad7R4CbQuX5bdM%2FBbVydkYMtyBcUXNNvSw6WjgdswgdgwFgYIKwYBBQUHARoECjAIoAYWBDkwMkowFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMB0GA1UdDgQWBBQnIAYZe0S8BrGKhEwR2XaLw9AbjzAfBgNVHSMEGDAWgBSs05P1Q0PMCr5FWBcTfZJ83MMBRjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSAAwRQIgKJgFUF%2BuJSvkscnqWYPsM2jH%2Ba7Odl2JtinLnCwJ4scCIQDu83%2FbwUR%2F2WHYB8UDviwrvzDR7EMiNTiX3BoXitx0OA%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_cp_1_3_subject_email](../../ISSUES/w_cp_1_3_subject_email/README.md) | warn | US_SHAKEN_CP | Email addresses are not allowed as the CP does not specify how to validate them |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 902J' |


Generated: 15 Nov 23 18:10 UTC