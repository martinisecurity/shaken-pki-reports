# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN MagicJack 324E

Tested At: 22 Aug 24 15:59 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 97 day(s)\
Subject: CN=SHAKEN MagicJack 324E, OU=NOC, O=MagicJack, ST=Florida, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://sns.magicjack.com/MagicJack_2411_324E.crt

[View certificate details](https://x509.io/?cert=MIICxTCCAmugAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkh6gwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMTEyNzE4MDQyM1oXDTI0MTEyNjE4MDQyM1owYTELMAkGA1UEBhMCVVMxEDAOBgNVBAgMB0Zsb3JpZGExEjAQBgNVBAoMCU1hZ2ljSmFjazEMMAoGA1UECwwDTk9DMR4wHAYDVQQDDBVTSEFLRU4gTWFnaWNKYWNrIDMyNEUwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASPxYeO7gUtKwcxTmxvuIA9qpRvfNeHrlNku7%2F1%2FAey0nHwIv8jDX1PlNEk8rH6jq5ZVxcVFxOai1mUrWxUJh%2Fvo4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQzMjRFMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQUpsl%2B2DMRMF2F6vY%2BWaw9ZgCWdRMwHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0gAMEUCIB8Lnj8S6udQXH0XQNmnT9Ou71eKq8dwIZ4W%2B04cOnUtAiEA60cnQNzCcw5U%2FIDt9VadUVSrcX97l7kYGCoFL1BuhQs%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension contains an invalid OID value: 2.16.840.1.114569.1.1.1. Available OIDs: 2.16.840.1.114569.1.1.3, 2.16.840.1.114569.1.1.4 |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 324E', but common name is 'SHAKEN MagicJack 324E' |


Generated: 22 Aug 24 16:06 UTC