# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN CIMA Telecom, Inc 313K

Tested At: 22 Aug 24 15:36 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -82 day(s)\
Subject: CN=SHAKEN CIMA Telecom\\, Inc 313K, O=CIMA Telecom\\, Inc, ST=Florida, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/313K/429C7C70711E3820F0B8E1DEAE6FF32622649B09.pem

[View certificate details](https://x509.io/?cert=MIICxzCCAm2gAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkmwkwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTI0MDUwMjE1MjQ1MVoXDTI0MDYwMTE1MjQ1MVowYzELMAkGA1UEBhMCVVMxEDAOBgNVBAgMB0Zsb3JpZGExGjAYBgNVBAoMEUNJTUEgVGVsZWNvbSwgSW5jMSYwJAYDVQQDDB1TSEFLRU4gQ0lNQSBUZWxlY29tLCBJbmMgMzEzSzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABFmcWIlCBgPkAnk9Pe3lHPPJWz%2FUYri3JuFS7R6yEJOfv7pfdZam2gSKAj4oM3z5Ioh04bAcrdvO7S5mQ1QsYR6jgdswgdgwFgYIKwYBBQUHARoECjAIoAYWBDMxM0swFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEEMB0GA1UdDgQWBBRtr%2FQM77fTyZx4t0EtbrcASrAb4jAfBgNVHSMEGDAWgBSs05P1Q0PMCr5FWBcTfZJ83MMBRjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSAAwRQIgGcDayN%2FFH%2Frf68rXVz90MiUM5X1uzVOwNkYZoiWuNBECIQC%2BJpD9QAraUsMlWqBRmetT8vw2xsAMOb0lOwjJKt%2Bd4w%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 313K', but common name is 'SHAKEN CIMA Telecom, Inc 313K' |
| [e_atis_ext_crl_distribution_struct](../../ISSUES/e_atis_ext_crl_distribution_struct/README.md) | error | ATIS1000080 | CRL Distribution Point shall contain a CRLIssuer field |


Generated: 22 Aug 24 16:06 UTC