# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN IPBTel 535K

Tested At: 04 Oct 24 16:15 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: -76 day(s)\
Subject: CN=SHAKEN IPBTel 535K, OU=IPBTel, O=IPBTel, ST=Texas, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/IPBTel_535K

[View certificate details](https://x509.io/?cert=MIICvzCCAmagAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkekMwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDcyMDE4NTIyMFoXDTI0MDcxOTE4NTIyMFowXDELMAkGA1UEBhMCVVMxDjAMBgNVBAgMBVRleGFzMQ8wDQYDVQQKDAZJUEJUZWwxDzANBgNVBAsMBklQQlRlbDEbMBkGA1UEAwwSU0hBS0VOIElQQlRlbCA1MzVLMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAENbIrO%2FvZ9vr08kz1o5adMnLrgrSqbl2s77%2FeBRNcAuoyf0WtxBeAA3sfrQ%2BmF5Z33ejARPk4MF9C0lqjPBYYLqOB2zCB2DAWBggrBgEFBQcBGgQKMAigBhYENTM1SzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwHQYDVR0OBBYEFDdJ3tDjLcy%2FE28NQGM3Hl5xNwjXMB8GA1UdIwQYMBaAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNHADBEAiBLT6vWPPe1esygWzZpexYy8RDvJ1Ws%2B0QSGu5wsD7jTAIgYGx5fDNjP5gru%2BAMcaxqe064RFTJirj%2FUVsl9kB5sC8%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 535K', but common name is 'SHAKEN IPBTel 535K' |


Generated: 04 Oct 24 16:29 UTC