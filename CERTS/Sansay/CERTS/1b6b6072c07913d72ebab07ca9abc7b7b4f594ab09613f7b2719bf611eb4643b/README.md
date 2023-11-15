# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN CIMA Telecom, Inc 313K

Tested At: 15 Nov 23 18:03 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -13 day(s)\
Subject: emailAddress=jmedina@cimagroup.com, CN=SHAKEN CIMA Telecom\\, Inc 313K, OU=Operations, O=CIMA Telecom\\, Inc, ST=Florida, C=US, emailAddress=jmedina@cimagroup.com\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/313K/order/168_313K_152

[View certificate details](https://understandingwebpki.com/?cert=MIIDAzCCAqmgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkgqIwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMTAwMjE4NDUzNloXDTIzMTEwMTE4NDUzNlowgZ4xCzAJBgNVBAYTAlVTMRAwDgYDVQQIDAdGbG9yaWRhMRowGAYDVQQKDBFDSU1BIFRlbGVjb20sIEluYzETMBEGA1UECwwKT3BlcmF0aW9uczEmMCQGA1UEAwwdU0hBS0VOIENJTUEgVGVsZWNvbSwgSW5jIDMxM0sxJDAiBgkqhkiG9w0BCQEWFWptZWRpbmFAY2ltYWdyb3VwLmNvbTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABFmcWIlCBgPkAnk9Pe3lHPPJWz%2FUYri3JuFS7R6yEJOfv7pfdZam2gSKAj4oM3z5Ioh04bAcrdvO7S5mQ1QsYR6jgdswgdgwFgYIKwYBBQUHARoECjAIoAYWBDMxM0swFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMB0GA1UdDgQWBBRtr%2FQM77fTyZx4t0EtbrcASrAb4jAfBgNVHSMEGDAWgBSs05P1Q0PMCr5FWBcTfZJ83MMBRjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSAAwRQIgMkd7uJW6K5%2Btnj6sTpUQD3GkBMeayfakGVaTMKBjYNsCIQDruYAYBbFRIrjZgdkQtbgFomVibd0VnbwSgMUofwq2TA%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 313K' |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [w_cp_1_3_subject_email](../../ISSUES/w_cp_1_3_subject_email/README.md) | warn | US_SHAKEN_CP | Email addresses are not allowed as the CP does not specify how to validate them |


Generated: 15 Nov 23 18:10 UTC