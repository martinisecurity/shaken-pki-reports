# STIR/SHAKEN CA Ecosystem Compliance
## Sansay

### Certificate 96c46e788ace45c3011eb790bf471fd8460f9a7e
Tested At: 2022-10-26 21:00:22 +0000 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 15 day(s)\
Subject: CN=SHAKEN PNG Telecommunications Inc 3395, OU=NOC, O=PNG Telecommunications Inc, ST=Ohio, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US

Link: https://cr.sansay.com/PowerNetGlobal_3395

View: [Click to view](https://understandingwebpki.com/?cert=MIIC5TCCAougAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkTz4wCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAxMTE3MjMyMVoXDTIyMTExMDE3MjMyMVowgYAxCzAJBgNVBAYTAlVTMQ0wCwYDVQQIDARPaGlvMSMwIQYDVQQKDBpQTkcgVGVsZWNvbW11bmljYXRpb25zIEluYzEMMAoGA1UECwwDTk9DMS8wLQYDVQQDDCZTSEFLRU4gUE5HIFRlbGVjb21tdW5pY2F0aW9ucyBJbmMgMzM5NTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABMwNyaEeWA0vGW0AzW6J40Tgty3ddjhlKHaR5Fbaaw3hgWwULA51dFB5gSMg2n4dJJVznCBP85MHY%2BZzFAPHVQKjgdswgdgwFgYIKwYBBQUHARoECjAIoAYWBDMzOTUwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMB0GA1UdDgQWBBT6etr4H3xed5Sv%2FS9%2FyUk43joO4TAfBgNVHSMEGDAWgBSs05P1Q0PMCr5FWBcTfZJ83MMBRjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSAAwRQIhAOB0pSukuScCSLPGTmugUCeEGNZ%2FSof6DSuKx3Im1ocUAiA7xqVWvf%2BV1I876E8IbxODI6ZARt6swMhBGoHJRH7muQ%3D%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_sti_certificate_policies | error | ATIS-1000080v4 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| e_cp1_3_ambiguous_identifier | error | CPv1.3 | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| e_sti_crl_distribution | error | ATIS-1000080v4 | CRL Distribution Point shall be reachable if the requesting IP address within the program ACLs |
| e_sti_subject_cn | error | ATIS-1000080v4 | Common name shall contain the text string 'SHAKEN 3395' |
| e_cp1_3_subject_sn | error | CPv1.3 | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| w_cp1_3_subject_rdn_unknown | warn | CPv1.3 | Only CN, C, O, and SERIALNUMBER can be included. Additional RNDs may introduce ambiguity and may not be verifiable |

Generated: 26/10/2022 at 21:01:13