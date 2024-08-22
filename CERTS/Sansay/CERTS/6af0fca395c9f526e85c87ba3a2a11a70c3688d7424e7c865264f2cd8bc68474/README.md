# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Lightspeed Voice 557F

Tested At: 22 Aug 24 15:32 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -279 day(s)\
Subject: CN=SHAKEN Lightspeed Voice 557F, OU=Engineering, O=Lightspeed Voice, ST=Florida, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/lightspeedvoice

[View certificate details](https://x509.io/?cert=MIIC2zCCAoGgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkhKcwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMTAxODE0MjMxNFoXDTIzMTExNzE0MjMxNFowdzELMAkGA1UEBhMCVVMxEDAOBgNVBAgMB0Zsb3JpZGExGTAXBgNVBAoMEExpZ2h0c3BlZWQgVm9pY2UxFDASBgNVBAsMC0VuZ2luZWVyaW5nMSUwIwYDVQQDDBxTSEFLRU4gTGlnaHRzcGVlZCBWb2ljZSA1NTdGMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEBXl3GCua%2BqMXp31Y8k4f0nN4gH7VMfG4sDv%2FiicinVgnyuVMVKGxE6ogbymRuI31eSsPawjrSItpvtyAgyUGcKOB2zCB2DAWBggrBgEFBQcBGgQKMAigBhYENTU3RjAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFFcihViwyO%2BHvCln%2FCblxsPQ29H9MB8GA1UdIwQYMBaAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNIADBFAiEAoxt2JWiWT9bdrCkHMCXCnby%2BiAvJBNkse1zC6YNrVp8CIFY8%2FMU3pdPUOLFs6QUnHuQ4KhlvBIQBw3rlId%2FVYjzT)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 557F', but common name is 'SHAKEN Lightspeed Voice 557F' |
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension contains an invalid OID value: 2.16.840.1.114569.1.1.1. Available OIDs: 2.16.840.1.114569.1.1.3, 2.16.840.1.114569.1.1.4 |


Generated: 22 Aug 24 15:44 UTC