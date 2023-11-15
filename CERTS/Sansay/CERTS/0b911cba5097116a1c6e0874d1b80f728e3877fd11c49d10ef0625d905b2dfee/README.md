# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Cyberlynk Network, LLC 086K

Tested At: 15 Nov 23 16:34 UTC\
Initial Validity Period: 60 day(s)\
Remaining Validity Period: -165 day(s)\
Subject: CN=SHAKEN Cyberlynk Network\\, LLC 086K, OU=NOC, O=Cyberlynk Network\\, LLC, ST=Wisconsin, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/086K/order/210_086K_91

[View certificate details](https://understandingwebpki.com/?cert=MIIC4DCCAoegAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkaOswCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDQwNDE2MjMxMVoXDTIzMDYwMzE2MjMxMVowfTELMAkGA1UEBhMCVVMxEjAQBgNVBAgMCVdpc2NvbnNpbjEfMB0GA1UECgwWQ3liZXJseW5rIE5ldHdvcmssIExMQzEMMAoGA1UECwwDTk9DMSswKQYDVQQDDCJTSEFLRU4gQ3liZXJseW5rIE5ldHdvcmssIExMQyAwODZLMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE6ZAhQPC731rtWrCAz%2BgThajNag3SWjxRttl24GoYJYnVi6ViOUxVbPb3vDtkhbowTnvA%2Fy9799zghr5LXQ9Ij6OB2zCB2DAWBggrBgEFBQcBGgQKMAigBhYEMDg2SzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwHQYDVR0OBBYEFCzz0EkQDpfAtUdNowqgn4YpCNkrMB8GA1UdIwQYMBaAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNHADBEAiBC1VtAIJXZGxV9Q8xildp%2F%2BgBevFYmYNImgkodbuCaeAIgUN%2B%2BKWCr6dB%2BY%2FepfeSps%2FiJbDE9uCIanQFedG9Nt%2Bc%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 086K' |


Generated: 15 Nov 23 17:17 UTC