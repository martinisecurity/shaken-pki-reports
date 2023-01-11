# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Quality Voice & Data Inc. 548J

Tested At: 11 Jan 23 21:50 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -60 day(s)\
Subject: emailAddress=kelsey@qualityvoicedata.com, CN=SHAKEN Quality Voice & Data Inc. 548J, OU=Quality Voice & Data, O=Quality Voice & Data Inc., ST=Nebraska, C=US, emailAddress=kelsey@qualityvoicedata.com\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/548J/order/154_548J_67

[View certificate details](https://understandingwebpki.com/?cert=MIID0jCCA3igAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkUAAwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAxMzE3MjAxM1oXDTIyMTExMjE3MjAxM1owgb8xCzAJBgNVBAYTAlVTMREwDwYDVQQIDAhOZWJyYXNrYTEiMCAGA1UECgwZUXVhbGl0eSBWb2ljZSAmIERhdGEgSW5jLjEdMBsGA1UECwwUUXVhbGl0eSBWb2ljZSAmIERhdGExLjAsBgNVBAMMJVNIQUtFTiBRdWFsaXR5IFZvaWNlICYgRGF0YSBJbmMuIDU0OEoxKjAoBgkqhkiG9w0BCQEWG2tlbHNleUBxdWFsaXR5dm9pY2VkYXRhLmNvbTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABHxIpvtIpathfVQA%2Bf9h5YTuYtpurXcE4Fo%2BWfOrs4qTvs97AEUcl9TF5yLWG%2BWt%2FOsXbViKVjuykX3%2F5qk0fMSjggGIMIIBhDAWBggrBgEFBQcBGgQKMAigBhYENTQ4SjAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFMreULfINsK1cltcvHQlGBGCmgS5MIHKBgNVHSMEgcIwgb%2BAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGoYGQpIGNMIGKMQswCQYDVQQGEwJVUzETMBEGA1UECAwKQ2FsaWZvcm5pYTESMBAGA1UEBwwJU2FuIERpZWdvMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEhMB8GA1UEAwwYU0hBS0VOIFNhbnNheSBSb290IENBIFVTghQUtV84BfXwexO1%2BLDe7SDyUXA%2BnDBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSAAwRQIhAOvzYGmM9O3f280B5J3t9lPVKnTHGJjn6K612lmc8%2F42AiBntNQ3iVTMsDVBzxMZ4NnUehM1igM7VvscgJ5ZXpOchQ%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 548J' |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [w_cp_1_3_subject_email](../../ISSUES/w_cp_1_3_subject_email/README.md) | warn | US_SHAKEN_CP | Email addresses are not allowed as the CP does not specify how to validate them |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |


Generated: 11 Jan 23 21:59 UTC