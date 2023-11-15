# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Peoples SHAKEN 2130

Tested At: 15 Nov 23 16:12 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: -151 day(s)\
Subject: CN=Peoples SHAKEN 2130, OU=STI, O=Peoples, C=US\
Issuer: CN=SHAKEN Ribbon Issuing CA, OU=Certification Authorities, O=Ribbon Communications, C=US\
Link: https://prod001-prod011-cr.rbbnidhub.com/R3kZfzj7gz/STI2130-5e62e39ccbcec63dd71e5d503ec51e3e

[View certificate details](https://understandingwebpki.com/?cert=MIIC4TCCAoigAwIBAgIQXmLjnMvOxj3XHl1QPsUePjAKBggqhkjOPQQDAjB0MQswCQYDVQQGEwJVUzEeMBwGA1UEChMVUmliYm9uIENvbW11bmljYXRpb25zMSIwIAYDVQQLExlDZXJ0aWZpY2F0aW9uIEF1dGhvcml0aWVzMSEwHwYDVQQDExhTSEFLRU4gUmliYm9uIElzc3VpbmcgQ0EwHhcNMjIwNjE2MTg1MDMyWhcNMjMwNjE2MTg1MDMxWjBLMQswCQYDVQQGEwJVUzEQMA4GA1UECgwHUGVvcGxlczEMMAoGA1UECwwDU1RJMRwwGgYDVQQDDBNQZW9wbGVzIFNIQUtFTiAyMTMwMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEUjmDuONkeM4cCDxjSi947%2FFWHvVgo%2FLHshorOscXDYTzKuuWeGfTAm47reha%2BZZ8vcLIory%2Bswx0lBCSWDFGH6OCASMwggEfMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMB0GA1UdDgQWBBTGwoBh7zHmS8Wuqg1fst2wXN1yajAZBgNVHSAEEjAQMA4GCmCGSAGG%2FwkBAQEwADBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwQwYIKwYBBQUHAQEENzA1MDMGCCsGAQUFBzAChidodHRwOi8vc3RpY3IucmJibmlkaHViLmNvbS9yYmJuX2ljYS5jcnQwHwYDVR0jBBgwFoAUj9%2BXO0VyeUL6rz3v3FI1QGCMpNgwFgYIKwYBBQUHARoECjAIoAYWBDIxMzAwCgYIKoZIzj0EAwIDRwAwRAIgQIKDPKrgRTGyS5hbrTfEquDT9%2F0bF6fx83l0YkM2f%2FUCIFwMMQiinRz142d9ntqthOuat1hRWhY1rxpmMAGhOxYo)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |


Generated: 15 Nov 23 16:51 UTC