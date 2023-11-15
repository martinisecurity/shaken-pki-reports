# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN IPBTel 535K

Tested At: 15 Nov 23 16:34 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 248 day(s)\
Subject: CN=SHAKEN IPBTel 535K, OU=IPBTel, O=IPBTel, ST=Texas, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/IPBTel_535K

[View certificate details](https://understandingwebpki.com/?cert=MIICvzCCAmagAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkekMwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDcyMDE4NTIyMFoXDTI0MDcxOTE4NTIyMFowXDELMAkGA1UEBhMCVVMxDjAMBgNVBAgMBVRleGFzMQ8wDQYDVQQKDAZJUEJUZWwxDzANBgNVBAsMBklQQlRlbDEbMBkGA1UEAwwSU0hBS0VOIElQQlRlbCA1MzVLMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAENbIrO%2FvZ9vr08kz1o5adMnLrgrSqbl2s77%2FeBRNcAuoyf0WtxBeAA3sfrQ%2BmF5Z33ejARPk4MF9C0lqjPBYYLqOB2zCB2DAWBggrBgEFBQcBGgQKMAigBhYENTM1SzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwHQYDVR0OBBYEFDdJ3tDjLcy%2FE28NQGM3Hl5xNwjXMB8GA1UdIwQYMBaAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNHADBEAiBLT6vWPPe1esygWzZpexYy8RDvJ1Ws%2B0QSGu5wsD7jTAIgYGx5fDNjP5gru%2BAMcaxqe064RFTJirj%2FUVsl9kB5sC8%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 535K' |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |


Generated: 15 Nov 23 17:17 UTC