# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Quality Voice & Data Inc. 548J

Tested At: 31 Oct 22 16:41 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 23 day(s)\
Subject: emailAddress=kelsey@qualityvoicedata.com, CN=SHAKEN Quality Voice & Data Inc. 548J, OU=Quality Voice & Data, O=Quality Voice & Data Inc., ST=Nebraska, C=US, emailAddress=kelsey@qualityvoicedata.com\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/548J/order/165_548J_67

View: [Click to view](https://understandingwebpki.com/?cert=MIID0jCCA3igAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkUtAwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAyNDE2MjUyNVoXDTIyMTEyMzE2MjUyNVowgb8xCzAJBgNVBAYTAlVTMREwDwYDVQQIDAhOZWJyYXNrYTEiMCAGA1UECgwZUXVhbGl0eSBWb2ljZSAmIERhdGEgSW5jLjEdMBsGA1UECwwUUXVhbGl0eSBWb2ljZSAmIERhdGExLjAsBgNVBAMMJVNIQUtFTiBRdWFsaXR5IFZvaWNlICYgRGF0YSBJbmMuIDU0OEoxKjAoBgkqhkiG9w0BCQEWG2tlbHNleUBxdWFsaXR5dm9pY2VkYXRhLmNvbTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABHxIpvtIpathfVQA%2Bf9h5YTuYtpurXcE4Fo%2BWfOrs4qTvs97AEUcl9TF5yLWG%2BWt%2FOsXbViKVjuykX3%2F5qk0fMSjggGIMIIBhDAWBggrBgEFBQcBGgQKMAigBhYENTQ4SjAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFMreULfINsK1cltcvHQlGBGCmgS5MIHKBgNVHSMEgcIwgb%2BAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGoYGQpIGNMIGKMQswCQYDVQQGEwJVUzETMBEGA1UECAwKQ2FsaWZvcm5pYTESMBAGA1UEBwwJU2FuIERpZWdvMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEhMB8GA1UEAwwYU0hBS0VOIFNhbnNheSBSb290IENBIFVTghQUtV84BfXwexO1%2BLDe7SDyUXA%2BnDBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSAAwRQIgNsqeXltarIkoGm%2BMZ0yzBt2bfRlulrtrBx8LNyJWppcCIQDRaX%2B7FnWoWP8v1gUbnAsRPoOiBPsq624jvO%2F6UKF11Q%3D%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_cp_1_3_subject_email](../../ISSUES/w_cp_1_3_subject_email/README.md) | warn | United States SHAKEN CP | Email addresses are not allowed as the CP does not specify how to validate them |
| [w_pki_subject_rdn_unknown](../../ISSUES/w_pki_subject_rdn_unknown/README.md) | warn | SHAKEN PKI Best Practice | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_cp1_3_ambiguous_identifier](../../ISSUES/e_cp1_3_ambiguous_identifier/README.md) | error | United States SHAKEN CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_sti_certificate_policies](../../ISSUES/e_sti_certificate_policies/README.md) | error | ATIS&#x2011;1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_cp1_3_subject_sn](../../ISSUES/e_cp1_3_subject_sn/README.md) | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_sti_subject_cn](../../ISSUES/e_sti_subject_cn/README.md) | error | ATIS&#x2011;1000080 | Common name shall contain the text string 'SHAKEN 548J' |


Generated: 31/10/2022 at 16:43:22