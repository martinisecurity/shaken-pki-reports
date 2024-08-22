# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Pierce SHAKEN 1581

Tested At: 22 Aug 24 15:59 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 71 day(s)\
Subject: CN=Pierce SHAKEN 1581, OU=STI, O=Pierce Telephone Company\\, Inc, C=US\
Issuer: CN=SHAKEN Ribbon Issuing CA, OU=Certification Authorities, O=Ribbon Communications, C=US\
Link: https://prod001-prod011-cr.rbbnidhub.com/tG4-037Sgz/STI-202311-1581-78042709ead9b2bf0eabd0ba5353df92

[View certificate details](https://x509.io/?cert=MIIC%2BDCCAp2gAwIBAgIQeAQnCerZsr8Oq9C6U1PfkjAKBggqhkjOPQQDAjB0MQswCQYDVQQGEwJVUzEeMBwGA1UEChMVUmliYm9uIENvbW11bmljYXRpb25zMSIwIAYDVQQLExlDZXJ0aWZpY2F0aW9uIEF1dGhvcml0aWVzMSEwHwYDVQQDExhTSEFLRU4gUmliYm9uIElzc3VpbmcgQ0EwHhcNMjMxMTAyMTMzNzMxWhcNMjQxMTAxMTMzNzMwWjBgMQswCQYDVQQGEwJVUzEmMCQGA1UECgwdUGllcmNlIFRlbGVwaG9uZSBDb21wYW55LCBJbmMxDDAKBgNVBAsMA1NUSTEbMBkGA1UEAwwSUGllcmNlIFNIQUtFTiAxNTgxMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE5hncrhtY6HJjpv6K%2BTWuM8IcWUFLLHr3%2FhllFQTF9mK3Z0untX9ox94Af9PaV7tE6UGrCczHXZ4MdcDVA5QobKOCASMwggEfMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMB0GA1UdDgQWBBSY1eSHU4FxLbte92zOGLa%2FSHiB%2FDAZBgNVHSAEEjAQMA4GCmCGSAGG%2FwkBAQEwADBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwQwYIKwYBBQUHAQEENzA1MDMGCCsGAQUFBzAChidodHRwOi8vc3RpY3IucmJibmlkaHViLmNvbS9yYmJuX2ljYS5jcnQwHwYDVR0jBBgwFoAUj9%2BXO0VyeUL6rz3v3FI1QGCMpNgwFgYIKwYBBQUHARoECjAIoAYWBDE1ODEwCgYIKoZIzj0EAwIDSQAwRgIhAIi%2Bo4FadRkgkaDGwnBwtqeYFmGBpQQjU44DHTHqRLNMAiEAieOZBJrgbqzIGm8xW%2BV%2FemGj79n%2BzX56o6jvTY0r4zI%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_ext_not_specified](../../ISSUES/e_atis_ext_not_specified/README.md) | error | ATIS1000080 | Certificate contains extensions that are not specified: 1.3.6.1.5.5.7.1.1 |
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension contains an invalid OID value: 2.16.840.1.114569.1.1.1. Available OIDs: 2.16.840.1.114569.1.1.3, 2.16.840.1.114569.1.1.4 |


Generated: 22 Aug 24 16:06 UTC