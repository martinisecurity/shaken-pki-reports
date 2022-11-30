# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Global Net Holdings Inc 306K

Tested At: 30 Nov 22 18:16 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 22 day(s)\
Subject: CN=SHAKEN Global Net Holdings Inc 306K, OU=Engineering\\ , O=Global Net Holdings Inc, ST=Pennsylvania, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/Global_Net_Holdings_Inc_306K

[View certificate details](https://understandingwebpki.com/?cert=MIIC8DCCApagAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkVhQwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTEyMTIyMTUyNFoXDTIyMTIyMTIyMTUyNFowgYsxCzAJBgNVBAYTAlVTMRUwEwYDVQQIDAxQZW5uc3lsdmFuaWExIDAeBgNVBAoMF0dsb2JhbCBOZXQgSG9sZGluZ3MgSW5jMRUwEwYDVQQLDAxFbmdpbmVlcmluZyAxLDAqBgNVBAMMI1NIQUtFTiBHbG9iYWwgTmV0IEhvbGRpbmdzIEluYyAzMDZLMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAESUnsw4b%2B4nxRSmnL62P09Ri85iZGM%2BZDK%2Fy%2BHQLLReuLg1ys7mswR%2F2TRuYHFHNxfV8PLWoz3KVyvOy38dnUQ6OB2zCB2DAWBggrBgEFBQcBGgQKMAigBhYEMzA2SzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFC9NS%2BhB16D3qCFOUgMD9S8Dw6M%2BMB8GA1UdIwQYMBaAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNIADBFAiEAyfEe5qmhxeLXOcgTTSiyRz8QCAJr6AJDYuAZWG1%2BdRwCIA8gxAwxu6LblPPHMQbD7DdJM76lESwAcBLReZkr0y3i)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 306K' |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 30 Nov 22 18:29 UTC