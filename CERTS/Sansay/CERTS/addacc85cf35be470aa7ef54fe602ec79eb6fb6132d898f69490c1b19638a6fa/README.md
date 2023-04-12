# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Doylestown Communications, Inc 849C

Tested At: 12 Apr 23 01:37 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 5 day(s)\
Subject: emailAddress=jclarke@heritagetelephone.com, CN=SHAKEN Doylestown Communications\\, Inc 849C, OU=Operations, O=Doylestown Communications\\, Inc, ST=Ohio, C=US, emailAddress=jclarke@heritagetelephone.com\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/849C/order/300_849C_78

[View certificate details](https://understandingwebpki.com/?cert=MIIDITCCAsigAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkZVswCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDMxNzA4NDY1MloXDTIzMDQxNjA4NDY1Mlowgb0xCzAJBgNVBAYTAlVTMQ0wCwYDVQQIDARPaGlvMScwJQYDVQQKDB5Eb3lsZXN0b3duIENvbW11bmljYXRpb25zLCBJbmMxEzARBgNVBAsMCk9wZXJhdGlvbnMxMzAxBgNVBAMMKlNIQUtFTiBEb3lsZXN0b3duIENvbW11bmljYXRpb25zLCBJbmMgODQ5QzEsMCoGCSqGSIb3DQEJARYdamNsYXJrZUBoZXJpdGFnZXRlbGVwaG9uZS5jb20wWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAARLv9VxWLlL9%2BMWvP1NNf9qred%2FtfHgjW1Eo7m35POfL4uDMx8AScgnKQnIlqtvcd0Kyt59vWYnUe%2BWghKNF6Z0o4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQ4NDlDMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzAdBgNVHQ4EFgQUywQNnfsX6%2BM4dUWhRO7erlOKfNcwHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0cAMEQCIEpBVET5cw7S4tgYcK93Xye2Jq9ZB3bduL%2FFMbHRj59yAiBX2DK%2FjPln8vice9rYhPU%2BDsthS8NK7e5sZB5AgDtJlA%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 849C' |
| [w_cp_1_3_subject_email](../../ISSUES/w_cp_1_3_subject_email/README.md) | warn | US_SHAKEN_CP | Email addresses are not allowed as the CP does not specify how to validate them |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 12 Apr 23 01:46 UTC