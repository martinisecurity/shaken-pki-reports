# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN BareTelecom 864J

Tested At: 15 Nov 23 16:09 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -30 day(s)\
Subject: emailAddress=chris@baretelecom.com, CN=SHAKEN BareTelecom 864J, OU=Operations, O=BareTelecom, ST=California, C=US, emailAddress=chris@baretelecom.com\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/864J/order/178_864J_159

[View certificate details](https://understandingwebpki.com/?cert=MIIC%2BzCCAqCgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkgKUwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDkxNTE2NTMzOVoXDTIzMTAxNTE2NTMzOVowgZUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRQwEgYDVQQKDAtCYXJlVGVsZWNvbTETMBEGA1UECwwKT3BlcmF0aW9uczEgMB4GA1UEAwwXU0hBS0VOIEJhcmVUZWxlY29tIDg2NEoxJDAiBgkqhkiG9w0BCQEWFWNocmlzQGJhcmV0ZWxlY29tLmNvbTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABDDS7l0vaaxRo5PuD%2BPqVhVve6XeHWGMrTrmV563iRvazHazHmoMABAUKMD0FgTQtsh5TKpUQdvJvnJwByVoCaGjgdswgdgwFgYIKwYBBQUHARoECjAIoAYWBDg2NEowFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMB0GA1UdDgQWBBRIdUFl6KvT9XyIaPknmcYnWIKCbzAfBgNVHSMEGDAWgBSs05P1Q0PMCr5FWBcTfZJ83MMBRjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSQAwRgIhANPX61WEF8UWCEGVQOAFNv2fpOdlMVcdJhIcN3PRe2omAiEA5mxdmAo5IbUroZCj7fhlN2HmH6Jt5Z2kB3jAZifyRlM%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 864J' |
| [w_cp_1_3_subject_email](../../ISSUES/w_cp_1_3_subject_email/README.md) | warn | US_SHAKEN_CP | Email addresses are not allowed as the CP does not specify how to validate them |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 15 Nov 23 16:51 UTC