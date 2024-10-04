# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Global Net Holdings Inc 306K

Tested At: 04 Oct 24 16:15 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 109 day(s)\
Subject: CN=SHAKEN Global Net Holdings Inc 306K, OU=oks, O=Global Net Holdings Inc, ST=Pennsylvania, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/Global_Net_306k

[View certificate details](https://x509.io/?cert=MIIC5jCCAo2gAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkjuswCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTI0MDEyMjE0MzMxN1oXDTI1MDEyMTE0MzMxN1owgYIxCzAJBgNVBAYTAlVTMRUwEwYDVQQIDAxQZW5uc3lsdmFuaWExIDAeBgNVBAoMF0dsb2JhbCBOZXQgSG9sZGluZ3MgSW5jMQwwCgYDVQQLDANva3MxLDAqBgNVBAMMI1NIQUtFTiBHbG9iYWwgTmV0IEhvbGRpbmdzIEluYyAzMDZLMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEcfIdmShBwehOLg2HYsxDHu9pdm4A7EC16b74yPZZ%2FzcJiZwmhMnmg1%2F%2BozcqHgVi9cirgJOZVf%2BMlSV7cPfAfaOB2zCB2DAWBggrBgEFBQcBGgQKMAigBhYEMzA2SzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwHQYDVR0OBBYEFFh0Z7n3RBunTvKJpwL8CWMyhLK3MB8GA1UdIwQYMBaAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNHADBEAiEA3J1Fi%2FL7acAz%2BNSFM3Vq7Qa3lTt4dHxkDDfohFwL6LYCHxGDAoA%2Bu61hoL%2BuBOdQrPP536KCfjqVL%2Fsw6RxagIQ%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 306K', but common name is 'SHAKEN Global Net Holdings Inc 306K' |


Generated: 04 Oct 24 16:29 UTC