# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Apeiron Systems 012J

Tested At: 26 Aug 24 18:14 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -25 day(s)\
Subject: CN=SHAKEN Apeiron Systems 012J, OU=NOC, O=Apeiron Systems, ST=California, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/012J/429C7C70711E3820F0B8E1DEAE6FF3262264A1FE.pem

[View certificate details](https://x509.io/?cert=MIIC1DCCAnqgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkof4wCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTI0MDcwMjE2NTA0NFoXDTI0MDgwMTE2NTA0NFowcDELMAkGA1UEBhMCVVMxEzARBgNVBAgMCkNhbGlmb3JuaWExGDAWBgNVBAoMD0FwZWlyb24gU3lzdGVtczEMMAoGA1UECwwDTk9DMSQwIgYDVQQDDBtTSEFLRU4gQXBlaXJvbiBTeXN0ZW1zIDAxMkowWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATvM7voHFBJOHVADYo1gKHK2wU0lgJXrz%2FJPCTX0AKFzXtzFfZF0xEGkCL70IlycCO4vTtcloIGOIJTSN8ZSPxko4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQwMTJKMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQUfHBaMnt53HaA60SExkm4zJE0Wl4wHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0gAMEUCIQDCR6zelHCZoWYGdYVV%2BXKcZXdh74eYe6UP5CdL9PixRQIgJk0m%2FMnQgTLxkOQsYTfFL0N7RI%2B9qQiVB5uFrWkFvEw%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 012J', but common name is 'SHAKEN Apeiron Systems 012J' |
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension contains an invalid OID value: 2.16.840.1.114569.1.1.1. Available OIDs: 2.16.840.1.114569.1.1.3, 2.16.840.1.114569.1.1.4 |


Generated: 26 Aug 24 18:49 UTC