# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Voice SY LLC 521K

Tested At: 26 Aug 24 17:56 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 113 day(s)\
Subject: CN=SHAKEN Voice SY LLC 521K, OU=Voice SY LLC Info, O=Voice SY LLC Info, ST=Wyoming, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://primodialer2.46labs.com/primodialer2.pem

[View certificate details](https://x509.io/?cert=MIIC3jCCAoSgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkilYwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMTIxODEzMDI1OVoXDTI0MTIxNzEzMDI1OVowejELMAkGA1UEBhMCVVMxEDAOBgNVBAgMB1d5b21pbmcxGjAYBgNVBAoMEVZvaWNlIFNZIExMQyBJbmZvMRowGAYDVQQLDBFWb2ljZSBTWSBMTEMgSW5mbzEhMB8GA1UEAwwYU0hBS0VOIFZvaWNlIFNZIExMQyA1MjFLMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEqdHi%2Fk5qHiOaUECv9lsvbZaDr0NmUDRDC4UAm4nfvG3z6EJq%2BgKzvJ0nD28F%2BzOfC6ZYLvrZI%2BrCdvZQ955i%2BaOB2zCB2DAWBggrBgEFBQcBGgQKMAigBhYENTIxSzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwHQYDVR0OBBYEFMXYxNtYBExIVN%2FZCkC1%2FbWxKcYtMB8GA1UdIwQYMBaAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNIADBFAiA0l%2FcWMhYM54tg1mEC5CVbB7aEuHbdnQ2Yyupw%2Bw435wIhAP70VLpR6dyfR8lgDsQZYkLfXqWkHGK%2ByrA6BbphRCES)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 521K', but common name is 'SHAKEN Voice SY LLC 521K' |


Generated: 26 Aug 24 18:03 UTC