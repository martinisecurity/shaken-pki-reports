# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN ALD Telecom 780J

Tested At: 04 Nov 22 01:10 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 342 day(s)\
Subject: CN=SHAKEN ALD Telecom 780J, OU=ALD_Telecom, O=ALD Telecom, ST=Arizona, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/ALD_Telecom_780J

[View certificate details](https://understandingwebpki.com/?cert=MIIC0TCCAnegAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkTwkwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAxMTE3MDg0M1oXDTIzMTAxMTE3MDg0M1owbTELMAkGA1UEBhMCVVMxEDAOBgNVBAgMB0FyaXpvbmExFDASBgNVBAoMC0FMRCBUZWxlY29tMRQwEgYDVQQLDAtBTERfVGVsZWNvbTEgMB4GA1UEAwwXU0hBS0VOIEFMRCBUZWxlY29tIDc4MEowWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATrCf7p2dfITkSwC6B1rTvmhygSeQJPCk%2F%2BrBOtuvbh5Xq4ZxNiagKMfjIcD3ttnBAUoy%2BprtYjRyGMsg0zaLXIo4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQ3ODBKMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQU5D6yJDqBy8u7UBJwCsnFyuynU18wHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0gAMEUCIGCAha%2FoKlqwoEUEXitI977nsKa1Sp3XOapSLkKeNzOUAiEAzSRGlMcIq%2By8WKxK8YEeYpeuJBp5IS7gpRJCsdkQDys%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 780J' |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 04 Nov 22 01:11 UTC