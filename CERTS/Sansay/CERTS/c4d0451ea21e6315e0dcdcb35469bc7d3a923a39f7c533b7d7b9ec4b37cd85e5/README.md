# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN NTC International, INC 016K

Tested At: 07 Dec 22 18:47 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -26 day(s)\
Subject: emailAddress=billing@ntcinternationalinc.com, CN=SHAKEN NTC International\\, INC 016K, OU=Operations, O=NTC International\\, INC, ST=Delaware, C=US, emailAddress=billing@ntcinternationalinc.com\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/016K/order/63_016K_115

[View certificate details](https://understandingwebpki.com/?cert=MIIDxTCCA2ygAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkT18wCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAxMjA3MjcyMloXDTIyMTExMTA3MjcyMlowgbMxCzAJBgNVBAYTAlVTMREwDwYDVQQIDAhEZWxhd2FyZTEfMB0GA1UECgwWTlRDIEludGVybmF0aW9uYWwsIElOQzETMBEGA1UECwwKT3BlcmF0aW9uczErMCkGA1UEAwwiU0hBS0VOIE5UQyBJbnRlcm5hdGlvbmFsLCBJTkMgMDE2SzEuMCwGCSqGSIb3DQEJARYfYmlsbGluZ0BudGNpbnRlcm5hdGlvbmFsaW5jLmNvbTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABM51nAzTItLw89uHu4KL0XTDPUzsmH0Ds0Hbi5F2DwFVaiKK5RvFSodn8j5IjXI49giYgMA0uPagzDbZoRkuPbmjggGIMIIBhDAWBggrBgEFBQcBGgQKMAigBhYEMDE2SzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFLzKcgLcCcHz52zFgMN6Ws%2FpccehMIHKBgNVHSMEgcIwgb%2BAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGoYGQpIGNMIGKMQswCQYDVQQGEwJVUzETMBEGA1UECAwKQ2FsaWZvcm5pYTESMBAGA1UEBwwJU2FuIERpZWdvMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEhMB8GA1UEAwwYU0hBS0VOIFNhbnNheSBSb290IENBIFVTghQUtV84BfXwexO1%2BLDe7SDyUXA%2BnDBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDRwAwRAIgRTfRc%2BxcSycX0Yc4%2Fxjgqr0CBJ3st1MctQyrGe7Bv6wCICOvdQYmr%2FyT%2BpM7t%2F5AAQxv1x9R%2Fkid%2FiI4LNpfvUzp)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [w_cp_1_3_subject_email](../../ISSUES/w_cp_1_3_subject_email/README.md) | warn | US_SHAKEN_CP | Email addresses are not allowed as the CP does not specify how to validate them |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 016K' |


Generated: 07 Dec 22 18:54 UTC