# STIR/SHAKEN CA Ecosystem Compliance
## Sansay

### Certificate 992de6f528e6b90ab11e952b9d0ef618e8086cbc
Tested At: 2022-10-27 21:42:37 +0000 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 30 day(s)\
Subject: CN=SHAKEN PNG Telecommunications Inc 3395, OU=NOC, O=PNG Telecommunications Inc, ST=Ohio, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US

Link: https://cr.sansay.com/PowerNetGlobal_3395

View: [Click to view](https://understandingwebpki.com/?cert=MIIC5DCCAougAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkU4wwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAyNzAwNDM0NFoXDTIyMTEyNjAwNDM0NFowgYAxCzAJBgNVBAYTAlVTMQ0wCwYDVQQIDARPaGlvMSMwIQYDVQQKDBpQTkcgVGVsZWNvbW11bmljYXRpb25zIEluYzEMMAoGA1UECwwDTk9DMS8wLQYDVQQDDCZTSEFLRU4gUE5HIFRlbGVjb21tdW5pY2F0aW9ucyBJbmMgMzM5NTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABMwNyaEeWA0vGW0AzW6J40Tgty3ddjhlKHaR5Fbaaw3hgWwULA51dFB5gSMg2n4dJJVznCBP85MHY%2BZzFAPHVQKjgdswgdgwFgYIKwYBBQUHARoECjAIoAYWBDMzOTUwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMB0GA1UdDgQWBBT6etr4H3xed5Sv%2FS9%2FyUk43joO4TAfBgNVHSMEGDAWgBSs05P1Q0PMCr5FWBcTfZJ83MMBRjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDRwAwRAIgTT6wBKwoA02tPuDq1q1%2BlMR0Kp5h1W65lMzY7qYElvgCIGAb1DCVL87vkVVlJR6i7O4iVra3TbO1gQeg2Kja8Iyj)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_sti_certificate_policies | error | ATIS-1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| w_cp1_3_subject_rdn_unknown | warn | United States SHAKEN CP | Only CN, C, O, and SERIALNUMBER can be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| e_cp1_3_ambiguous_identifier | error | United States SHAKEN CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| e_sti_subject_cn | error | ATIS-1000080 | Common name shall contain the text string 'SHAKEN 3395' |
| e_cp1_3_subject_sn | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |

Generated: 27/10/2022 at 21:42:52