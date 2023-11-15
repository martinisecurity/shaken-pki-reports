# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Identidad Advertising Development LLC 617K

Tested At: 15 Nov 23 18:03 UTC\
Initial Validity Period: 180 day(s)\
Remaining Validity Period: 59 day(s)\
Subject: CN=SHAKEN Identidad Advertising Development LLC 617K, OU=IT-VOIP, O=Identidad Advertising Development LLC, ST=Florida, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/Identidad_Advertising_Development_LLC_617K

[View certificate details](https://understandingwebpki.com/?cert=MIIDAzCCAqigAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkefIwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDcxNzE0MjA1NloXDTI0MDExMzE0MjA1NlowgZ0xCzAJBgNVBAYTAlVTMRAwDgYDVQQIDAdGbG9yaWRhMS4wLAYDVQQKDCVJZGVudGlkYWQgQWR2ZXJ0aXNpbmcgRGV2ZWxvcG1lbnQgTExDMRAwDgYDVQQLDAdJVC1WT0lQMTowOAYDVQQDDDFTSEFLRU4gSWRlbnRpZGFkIEFkdmVydGlzaW5nIERldmVsb3BtZW50IExMQyA2MTdLMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE38lc0VdyS9SAgOByiNdsjzX%2B2rxie9pW9%2BfOJ76IU%2FW6QcOJAQZI%2F9My1mkfgKBfRpg13mHKZ%2Fe2shms0La2vaOB2zCB2DAWBggrBgEFBQcBGgQKMAigBhYENjE3SzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwHQYDVR0OBBYEFFCsgO7C3rXZy478os5ZMR%2FYS2KRMB8GA1UdIwQYMBaAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNJADBGAiEApl6lNiogYVjEqJE7gNEau7WI4s8XF5F%2FXLvvYg2cPCICIQCtSz1IgrEuX1yqv9VX1cYel5LIMUkAVnvL54ppDW%2BJZQ%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 617K' |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |


Generated: 15 Nov 23 18:10 UTC