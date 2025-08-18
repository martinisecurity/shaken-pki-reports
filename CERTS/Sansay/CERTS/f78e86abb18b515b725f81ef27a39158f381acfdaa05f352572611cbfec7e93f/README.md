# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Dynalink Communications Inc 991D

Tested At: 18 Aug 25 20:48 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: -425 day(s)\
Subject: CN=SHAKEN Dynalink Communications Inc 991D, OU=Voice Engineering, O=Dynalink Communications Inc, ST=New York, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/991D_20230620.pem

[View certificate details](https://x509.io/?cert=MIIC%2BTCCAp%2BgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkdf0wCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDYyMDIwNTY0MVoXDTI0MDYxOTIwNTY0MVowgZQxCzAJBgNVBAYTAlVTMREwDwYDVQQIDAhOZXcgWW9yazEkMCIGA1UECgwbRHluYWxpbmsgQ29tbXVuaWNhdGlvbnMgSW5jMRowGAYDVQQLDBFWb2ljZSBFbmdpbmVlcmluZzEwMC4GA1UEAwwnU0hBS0VOIER5bmFsaW5rIENvbW11bmljYXRpb25zIEluYyA5OTFEMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAELM9B2uwVuWhkBR0N6NPyO7yMWi%2BQy3eV7eQpxTfVdY%2Ba3c2a%2FSJqvVArLubxyRhbwhqafPKub49I5lUUt2OIcqOB2zCB2DAWBggrBgEFBQcBGgQKMAigBhYEOTkxRDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwHQYDVR0OBBYEFLIVT8g3YeS0VOTblotwTvRHoJ6bMB8GA1UdIwQYMBaAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNIADBFAiEA6LRLUJQfGeVbWkMI%2FbxS8URNUlTi7Roncbq%2FInqfAd4CIBE%2BvwFvBhDnwHOJmb3V2nrSEy56gMiX%2B5mOkT6LqWjp)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 991D', but common name is 'SHAKEN Dynalink Communications Inc 991D' |
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |


Generated: 18 Aug 25 21:13 UTC