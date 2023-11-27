# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Comtalk Telecom 705K

Tested At: 27 Nov 23 23:20 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 330 day(s)\
Subject: CN=SHAKEN Comtalk Telecom 705K, OU=SUPPORT, O=Comtalk Telecom, ST=Florida, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/Comtalk_Telecom_705K

[View certificate details](https://understandingwebpki.com/?cert=MIIC1TCCAnugAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkhU8wCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMTAyMzE1NTQ0MloXDTI0MTAyMjE1NTQ0MlowcTELMAkGA1UEBhMCVVMxEDAOBgNVBAgMB0Zsb3JpZGExGDAWBgNVBAoMD0NvbXRhbGsgVGVsZWNvbTEQMA4GA1UECwwHU1VQUE9SVDEkMCIGA1UEAwwbU0hBS0VOIENvbXRhbGsgVGVsZWNvbSA3MDVLMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEJGaSvJ8eL9DLh2DVcx9DdzztS1dEGq8v6o1oZ3L%2BtYa293RCoLqdj8sgvRrxOtYQ%2FS8Kn65vjMc7QkcaQygPfKOB2zCB2DAWBggrBgEFBQcBGgQKMAigBhYENzA1SzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwHQYDVR0OBBYEFAkOc4OCwZmQ%2Ft0ZThN1eBqUbhpjMB8GA1UdIwQYMBaAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNIADBFAiEAlSaC262YOrSauA91JbHxH6x93I7svOOiNqhzkwL%2FLhQCIAhsMZJv5ZzlOPL1rqeS8HK8HVkZ3cQZ%2B0Iqbz3qfxGg)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 705K', but common name is 'SHAKEN Comtalk Telecom 705K' |
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |


Generated: 27 Nov 23 23:28 UTC