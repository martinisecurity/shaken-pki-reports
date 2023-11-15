# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Cherry Voice 506K

Tested At: 15 Nov 23 16:34 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 120 day(s)\
Subject: CN=SHAKEN Cherry Voice 506K, OU=Network Operations, O=Cherry Voice Inc, ST=Wyoming, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/Cherry_Voice_506K

[View certificate details](https://understandingwebpki.com/?cert=MIIC3jCCAoSgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkZP8wCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDMxNDIxMzk1N1oXDTI0MDMxMzIxMzk1N1owejELMAkGA1UEBhMCVVMxEDAOBgNVBAgMB1d5b21pbmcxGTAXBgNVBAoMEENoZXJyeSBWb2ljZSBJbmMxGzAZBgNVBAsMEk5ldHdvcmsgT3BlcmF0aW9uczEhMB8GA1UEAwwYU0hBS0VOIENoZXJyeSBWb2ljZSA1MDZLMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEIiHJ47TXwAkSW8Q%2BzMBEBteIePnJxtMftNLyTH1iyt4BljOPfPVI%2FBciDnpQSMp%2FKn8rv7M0wqxKNfvFW7M1CaOB2zCB2DAWBggrBgEFBQcBGgQKMAigBhYENTA2SzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwHQYDVR0OBBYEFOvOoQfrA0WQgA3HKntK2a%2FXRLItMB8GA1UdIwQYMBaAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNIADBFAiAefYHTK3QZp7US%2BobRhS8jyJzfXD4VKNiz6mgVn5RjIAIhAPnzbpNC3E4wy%2BYsVk5ssfzfKBTML0KiPc1EhoTSKFhZ)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 506K' |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 15 Nov 23 17:17 UTC