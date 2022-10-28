# STIR/SHAKEN CA Ecosystem Compliance
## Sansay

### Certificate 6232068ab8e75570bd010619f5fd471b86fc5f14
Tested At: 2022-10-28 10:32:29 +0000 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 27 day(s)\
Subject: emailAddress=billing@ntcinternationalinc.com, CN=SHAKEN NTC International\\, INC 016K, OU=Operations, O=NTC International\\, INC, ST=Delaware, C=US, emailAddress=billing@ntcinternationalinc.com\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US

Link: https://cr.sansay.com/016K/order/76_016K_115

View: [Click to view](https://understandingwebpki.com/?cert=MIIDxzCCA2ygAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkUuUwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAyNTA2MjIyMVoXDTIyMTEyNDA2MjIyMVowgbMxCzAJBgNVBAYTAlVTMREwDwYDVQQIDAhEZWxhd2FyZTEfMB0GA1UECgwWTlRDIEludGVybmF0aW9uYWwsIElOQzETMBEGA1UECwwKT3BlcmF0aW9uczErMCkGA1UEAwwiU0hBS0VOIE5UQyBJbnRlcm5hdGlvbmFsLCBJTkMgMDE2SzEuMCwGCSqGSIb3DQEJARYfYmlsbGluZ0BudGNpbnRlcm5hdGlvbmFsaW5jLmNvbTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABM51nAzTItLw89uHu4KL0XTDPUzsmH0Ds0Hbi5F2DwFVaiKK5RvFSodn8j5IjXI49giYgMA0uPagzDbZoRkuPbmjggGIMIIBhDAWBggrBgEFBQcBGgQKMAigBhYEMDE2SzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFLzKcgLcCcHz52zFgMN6Ws%2FpccehMIHKBgNVHSMEgcIwgb%2BAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGoYGQpIGNMIGKMQswCQYDVQQGEwJVUzETMBEGA1UECAwKQ2FsaWZvcm5pYTESMBAGA1UEBwwJU2FuIERpZWdvMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEhMB8GA1UEAwwYU0hBS0VOIFNhbnNheSBSb290IENBIFVTghQUtV84BfXwexO1%2BLDe7SDyUXA%2BnDBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSQAwRgIhAPYgeaMA5Lo5%2B9uTNuDP%2BTl%2FRm4a6fMWh%2BgQ4z%2Fo%2Bt8eAiEAjr04gKLegNwPnm%2B7AGXenyWT4gPWnDD8WdyCB5BfidU%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_cp1_3_ambiguous_identifier | error | United States SHAKEN CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| e_sti_certificate_policies | error | ATIS-1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| e_sti_subject_cn | error | ATIS-1000080 | Common name shall contain the text string 'SHAKEN 016K' |
| e_cp1_3_subject_sn | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| w_cp_1_3_subject_email | warn | United States SHAKEN CP | Email addresses are not allowed as the CP does not specify how to validate them |
| w_pki_subject_rdn_unknown | warn | SHAKEN PKI Best Practice | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |

Generated: 28/10/2022 at 10:33:25