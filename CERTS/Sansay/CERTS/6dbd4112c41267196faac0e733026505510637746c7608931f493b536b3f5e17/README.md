# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Doylestown Communications, Inc 849C

Tested At: 04 Oct 24 16:10 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -399 day(s)\
Subject: emailAddress=jclarke@heritagetelephone.com, CN=SHAKEN Doylestown Communications\\, Inc 849C, OU=Operations, O=Doylestown Communications\\, Inc, ST=Ohio, C=US, emailAddress=jclarke@heritagetelephone.com\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/849C/order/487_849C_78

[View certificate details](https://x509.io/?cert=MIIDIjCCAsigAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJke48wCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDgwMTIxMjE0MVoXDTIzMDgzMTIxMjE0MVowgb0xCzAJBgNVBAYTAlVTMQ0wCwYDVQQIDARPaGlvMScwJQYDVQQKDB5Eb3lsZXN0b3duIENvbW11bmljYXRpb25zLCBJbmMxEzARBgNVBAsMCk9wZXJhdGlvbnMxMzAxBgNVBAMMKlNIQUtFTiBEb3lsZXN0b3duIENvbW11bmljYXRpb25zLCBJbmMgODQ5QzEsMCoGCSqGSIb3DQEJARYdamNsYXJrZUBoZXJpdGFnZXRlbGVwaG9uZS5jb20wWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAARLv9VxWLlL9%2BMWvP1NNf9qred%2FtfHgjW1Eo7m35POfL4uDMx8AScgnKQnIlqtvcd0Kyt59vWYnUe%2BWghKNF6Z0o4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQ4NDlDMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzAdBgNVHQ4EFgQUywQNnfsX6%2BM4dUWhRO7erlOKfNcwHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0gAMEUCIQC1Ez%2FrJX5GzK0o%2BoVtYzCV1BXw%2BUoskEFzMz4MagiSKwIgK7Cjc5mHtl6FUeZOezGRTjKqrWz8ZVYADQGrHRz3Yl0%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 849C', but common name is 'SHAKEN Doylestown Communications, Inc 849C' |
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |


Generated: 04 Oct 24 16:29 UTC