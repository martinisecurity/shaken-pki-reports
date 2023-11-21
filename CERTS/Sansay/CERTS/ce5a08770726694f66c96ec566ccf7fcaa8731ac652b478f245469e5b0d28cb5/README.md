# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Global Net Holdings Inc 306K

Tested At: 21 Nov 23 01:47 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -271 day(s)\
Subject: CN=SHAKEN Global Net Holdings Inc 306K, OU=Engineering\\ , O=Global Net Holdings Inc, ST=Pennsylvania, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/Global_Net_Holdings_Inc_306K

[View certificate details](https://understandingwebpki.com/?cert=MIIC7zCCApagAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkXmYwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDEyMzA5MzAxN1oXDTIzMDIyMjA5MzAxN1owgYsxCzAJBgNVBAYTAlVTMRUwEwYDVQQIDAxQZW5uc3lsdmFuaWExIDAeBgNVBAoMF0dsb2JhbCBOZXQgSG9sZGluZ3MgSW5jMRUwEwYDVQQLDAxFbmdpbmVlcmluZyAxLDAqBgNVBAMMI1NIQUtFTiBHbG9iYWwgTmV0IEhvbGRpbmdzIEluYyAzMDZLMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAESUnsw4b%2B4nxRSmnL62P09Ri85iZGM%2BZDK%2Fy%2BHQLLReuLg1ys7mswR%2F2TRuYHFHNxfV8PLWoz3KVyvOy38dnUQ6OB2zCB2DAWBggrBgEFBQcBGgQKMAigBhYEMzA2SzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFC9NS%2BhB16D3qCFOUgMD9S8Dw6M%2BMB8GA1UdIwQYMBaAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNHADBEAiBbaDl5rYIKMIReslFK3Vl3ONlmuPlhH2nVwSbA5ELy0QIgQoYjEtfNkK7LLDrbt1994LDlQy6pBPQ0epY%2BnMlrtj4%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | the Certificate Policies extension contains an invalid OID value: 2.16.840.1.114569.1.1.1. Available OIDs: [2.16.840.1.114569.1.1.3 2.16.840.1.114569.1.1.4] |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 306K' |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 21 Nov 23 01:55 UTC