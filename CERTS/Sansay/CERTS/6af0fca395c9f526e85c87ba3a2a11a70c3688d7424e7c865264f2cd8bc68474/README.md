# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Lightspeed Voice 557F

Tested At: 15 Nov 23 18:03 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 2 day(s)\
Subject: CN=SHAKEN Lightspeed Voice 557F, OU=Engineering, O=Lightspeed Voice, ST=Florida, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/lightspeedvoice

[View certificate details](https://understandingwebpki.com/?cert=MIIC2zCCAoGgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkhKcwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMTAxODE0MjMxNFoXDTIzMTExNzE0MjMxNFowdzELMAkGA1UEBhMCVVMxEDAOBgNVBAgMB0Zsb3JpZGExGTAXBgNVBAoMEExpZ2h0c3BlZWQgVm9pY2UxFDASBgNVBAsMC0VuZ2luZWVyaW5nMSUwIwYDVQQDDBxTSEFLRU4gTGlnaHRzcGVlZCBWb2ljZSA1NTdGMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEBXl3GCua%2BqMXp31Y8k4f0nN4gH7VMfG4sDv%2FiicinVgnyuVMVKGxE6ogbymRuI31eSsPawjrSItpvtyAgyUGcKOB2zCB2DAWBggrBgEFBQcBGgQKMAigBhYENTU3RjAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFFcihViwyO%2BHvCln%2FCblxsPQ29H9MB8GA1UdIwQYMBaAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNIADBFAiEAoxt2JWiWT9bdrCkHMCXCnby%2BiAvJBNkse1zC6YNrVp8CIFY8%2FMU3pdPUOLFs6QUnHuQ4KhlvBIQBw3rlId%2FVYjzT)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 557F' |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 15 Nov 23 18:10 UTC