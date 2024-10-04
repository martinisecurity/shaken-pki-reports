# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN IPBTel 535K

Tested At: 04 Oct 24 16:02 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 289 day(s)\
Subject: CN=SHAKEN IPBTel 535K, OU=IPBTel, O=IPBTel, ST=Texas, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/535K/429C7C70711E3820F0B8E1DEAE6FF3262264A3F1.pem

[View certificate details](https://x509.io/?cert=MIICwDCCAmagAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJko%2FEwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTI0MDcxOTE3MzUyMVoXDTI1MDcxOTE3MzUyMVowXDELMAkGA1UEBhMCVVMxDjAMBgNVBAgMBVRleGFzMQ8wDQYDVQQKDAZJUEJUZWwxDzANBgNVBAsMBklQQlRlbDEbMBkGA1UEAwwSU0hBS0VOIElQQlRlbCA1MzVLMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAENbIrO%2FvZ9vr08kz1o5adMnLrgrSqbl2s77%2FeBRNcAuoyf0WtxBeAA3sfrQ%2BmF5Z33ejARPk4MF9C0lqjPBYYLqOB2zCB2DAWBggrBgEFBQcBGgQKMAigBhYENTM1SzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwHQYDVR0OBBYEFDdJ3tDjLcy%2FE28NQGM3Hl5xNwjXMB8GA1UdIwQYMBaAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNIADBFAiEAmiQRHBTqlIfKXtmenAKqsEf5k%2FcfofrSRB6%2B1C70%2FmsCIEjr0%2BGBw0Hh8dnO6JZqrruMIG8IEyTGd9qM4VX%2FOy2Z)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 535K', but common name is 'SHAKEN IPBTel 535K' |
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |


Generated: 04 Oct 24 16:29 UTC