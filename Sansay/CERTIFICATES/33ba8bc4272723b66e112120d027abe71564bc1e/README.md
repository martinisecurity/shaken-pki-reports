# STIR/SHAKEN CA Ecosystem Compliance
## Sansay

### Certificate 33ba8bc4272723b66e112120d027abe71564bc1e
Tested At: 2022-10-28 16:27:10 +0000 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 22 day(s)\
Subject: emailAddress=elevitt@1pcom.com, CN=SHAKEN 1stPoint Communications\\, LLC 463G, OU=NOC, O=1stPoint Communications\\, LLC, ST=New York, C=US, emailAddress=elevitt@1pcom.com\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US

Link: https://cr.sansay.com/463G/order/448_463G_53

View: [Click to view](https://understandingwebpki.com/?cert=MIIDvTCCA2OgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkUdYwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAxOTE2MzUxMFoXDTIyMTExODE2MzUxMFowgaoxCzAJBgNVBAYTAlVTMREwDwYDVQQIDAhOZXcgWW9yazElMCMGA1UECgwcMXN0UG9pbnQgQ29tbXVuaWNhdGlvbnMsIExMQzEMMAoGA1UECwwDTk9DMTEwLwYDVQQDDChTSEFLRU4gMXN0UG9pbnQgQ29tbXVuaWNhdGlvbnMsIExMQyA0NjNHMSAwHgYJKoZIhvcNAQkBFhFlbGV2aXR0QDFwY29tLmNvbTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABC7Os0cKRZ4dUV6B4QP8X%2BK3yo23fTX6t3lLNg%2FDTv2neRCPqx2o%2BZM8JVwSxj4LgyXCBlAuJFSvSoNR2r8CMEajggGIMIIBhDAWBggrBgEFBQcBGgQKMAigBhYENDYzRzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFM81qK%2Bd66zq6IUVZdoP8XnclRGbMIHKBgNVHSMEgcIwgb%2BAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGoYGQpIGNMIGKMQswCQYDVQQGEwJVUzETMBEGA1UECAwKQ2FsaWZvcm5pYTESMBAGA1UEBwwJU2FuIERpZWdvMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEhMB8GA1UEAwwYU0hBS0VOIFNhbnNheSBSb290IENBIFVTghQUtV84BfXwexO1%2BLDe7SDyUXA%2BnDBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSAAwRQIgVWt2FuOU0pRenDZqeGDSth881%2FH1gqqpIdhFjTBcE%2B8CIQDEWGTOUlNWjwSXjvvxGJ9u2CaFZB1AalVsSEgHkAGlWg%3D%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_sti_certificate_policies | error | ATIS-1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| w_cp_1_3_subject_email | warn | United States SHAKEN CP | Email addresses are not allowed as the CP does not specify how to validate them |
| e_cp1_3_subject_sn | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| e_cp1_3_ambiguous_identifier | error | United States SHAKEN CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| e_sti_subject_cn | error | ATIS-1000080 | Common name shall contain the text string 'SHAKEN 463G' |
| w_pki_subject_rdn_unknown | warn | SHAKEN PKI Best Practice | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |

Generated: 28/10/2022 at 16:28:22