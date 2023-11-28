# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN CIMA Telecom, Inc 313K

Tested At: 28 Nov 23 16:04 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -53 day(s)\
Subject: emailAddress=jmedina@cimagroup.com, CN=SHAKEN CIMA Telecom\\, Inc 313K, OU=Operations, O=CIMA Telecom\\, Inc, ST=Florida, C=US, emailAddress=jmedina@cimagroup.com\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/313K/order/144_313K_152

[View certificate details](https://understandingwebpki.com/?cert=MIIDAzCCAqmgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkf5YwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDkwNTIxMDAzN1oXDTIzMTAwNTIxMDAzN1owgZ4xCzAJBgNVBAYTAlVTMRAwDgYDVQQIDAdGbG9yaWRhMRowGAYDVQQKDBFDSU1BIFRlbGVjb20sIEluYzETMBEGA1UECwwKT3BlcmF0aW9uczEmMCQGA1UEAwwdU0hBS0VOIENJTUEgVGVsZWNvbSwgSW5jIDMxM0sxJDAiBgkqhkiG9w0BCQEWFWptZWRpbmFAY2ltYWdyb3VwLmNvbTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABFmcWIlCBgPkAnk9Pe3lHPPJWz%2FUYri3JuFS7R6yEJOfv7pfdZam2gSKAj4oM3z5Ioh04bAcrdvO7S5mQ1QsYR6jgdswgdgwFgYIKwYBBQUHARoECjAIoAYWBDMxM0swFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMB0GA1UdDgQWBBRtr%2FQM77fTyZx4t0EtbrcASrAb4jAfBgNVHSMEGDAWgBSs05P1Q0PMCr5FWBcTfZJ83MMBRjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSAAwRQIhAMQxFJpelIpVj8xBCvWHixrTzrpaQhlgLLO6581keyABAiB17luRmMYCH7lxPaxzYp5fjo2aV6ALxnrILuuzepBbPg%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 313K', but common name is 'SHAKEN CIMA Telecom, Inc 313K' |
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |


Generated: 28 Nov 23 16:15 UTC