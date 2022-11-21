# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN PNG Telecommunications Inc 3395

Tested At: 21 Nov 22 20:45 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 25 day(s)\
Subject: CN=SHAKEN PNG Telecommunications Inc 3395, OU=NOC, O=PNG Telecommunications Inc, ST=Ohio, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/PowerNetGlobal_3395

[View certificate details](https://understandingwebpki.com/?cert=MIIC5TCCAougAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkVZUwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTExNTIxMTU0NVoXDTIyMTIxNTIxMTU0NVowgYAxCzAJBgNVBAYTAlVTMQ0wCwYDVQQIDARPaGlvMSMwIQYDVQQKDBpQTkcgVGVsZWNvbW11bmljYXRpb25zIEluYzEMMAoGA1UECwwDTk9DMS8wLQYDVQQDDCZTSEFLRU4gUE5HIFRlbGVjb21tdW5pY2F0aW9ucyBJbmMgMzM5NTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABMwNyaEeWA0vGW0AzW6J40Tgty3ddjhlKHaR5Fbaaw3hgWwULA51dFB5gSMg2n4dJJVznCBP85MHY%2BZzFAPHVQKjgdswgdgwFgYIKwYBBQUHARoECjAIoAYWBDMzOTUwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMB0GA1UdDgQWBBT6etr4H3xed5Sv%2FS9%2FyUk43joO4TAfBgNVHSMEGDAWgBSs05P1Q0PMCr5FWBcTfZJ83MMBRjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSAAwRQIhALyBH6jsfNRmuNmKH%2FRyn2iTDJgev7r4isdYHK%2B8dA7eAiBQbpgMYn0oRMsoPXKGN1uv8PzNJP%2BY01KxEx%2FxNvmgRg%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 3395' |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |


Generated: 21 Nov 22 20:55 UTC