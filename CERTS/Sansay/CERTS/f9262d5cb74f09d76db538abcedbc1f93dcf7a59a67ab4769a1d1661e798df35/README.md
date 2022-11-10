# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN VoIP Innovations 597F

Tested At: 10 Nov 22 23:20 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 16 day(s)\
Subject: CN=SHAKEN VoIP Innovations 597F, OU=NOC, O=VoIP Innovations, ST=Pennsylvania, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/VoIP_Innovations_597F

[View certificate details](https://understandingwebpki.com/?cert=MIIC1zCCAn6gAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkU34wCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAyNzAwMzgzNFoXDTIyMTEyNjAwMzgzNFowdDELMAkGA1UEBhMCVVMxFTATBgNVBAgMDFBlbm5zeWx2YW5pYTEZMBcGA1UECgwQVm9JUCBJbm5vdmF0aW9uczEMMAoGA1UECwwDTk9DMSUwIwYDVQQDDBxTSEFLRU4gVm9JUCBJbm5vdmF0aW9ucyA1OTdGMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEZxb5rNCMqQTKIowEPxI78IPaRDMoC1gJoQtw%2FiJ%2BjTAuI2%2F4nDDGBz7ACAo3eIerjz3WANlpeFOQbeXW03U8jaOB2zCB2DAWBggrBgEFBQcBGgQKMAigBhYENTk3RjAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFOKfbEv6IldN3qqnsd3NpPk2C9PwMB8GA1UdIwQYMBaAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNHADBEAiAzDugP3zkX8MLNFfk%2Filzd1ZJ5qmWHOOVuiLR0ZShdrwIgNo5Mf5mznru9BdJ7SA42eKmEGJbQYq%2F26xRcgCm1i4I%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 597F' |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |


Generated: 10 Nov 22 23:30 UTC