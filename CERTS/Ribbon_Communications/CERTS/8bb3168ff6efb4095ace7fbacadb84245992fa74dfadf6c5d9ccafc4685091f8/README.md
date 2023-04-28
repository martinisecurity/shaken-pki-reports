# STIR/SHAKEN CA Ecosystem Compliance

## Certificate UniVoIP SHAKEN 629J

Tested At: 28 Apr 23 02:11 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 50 day(s)\
Subject: CN=UniVoIP SHAKEN 629J, OU=STI, O=UniVoIP, C=US\
Issuer: CN=SHAKEN Ribbon Issuing CA, OU=Certification Authorities, O=Ribbon Communications, C=US\
Link: https://prod001-prod011-cr.rbbnidhub.com/W_HMwhNngz/629J-33da08c43c8599ca98230bb276bac5a2

[View certificate details](https://understandingwebpki.com/?cert=MIIC4TCCAoigAwIBAgIQM9oIxDyFmcqYIwuydrrFojAKBggqhkjOPQQDAjB0MQswCQYDVQQGEwJVUzEeMBwGA1UEChMVUmliYm9uIENvbW11bmljYXRpb25zMSIwIAYDVQQLExlDZXJ0aWZpY2F0aW9uIEF1dGhvcml0aWVzMSEwHwYDVQQDExhTSEFLRU4gUmliYm9uIElzc3VpbmcgQ0EwHhcNMjIwNjE2MTkzMDQyWhcNMjMwNjE2MTkzMDQxWjBLMQswCQYDVQQGEwJVUzEQMA4GA1UECgwHVW5pVm9JUDEMMAoGA1UECwwDU1RJMRwwGgYDVQQDDBNVbmlWb0lQIFNIQUtFTiA2MjlKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE8d0WtoA8%2FAJhY%2BZIaUb3QJSKmSpOGgcWqZc1QnaTu9TeDgeRdJwBLNmrCz0YJhnY1I0gMNWyAUsEygMQ%2FhXMsaOCASMwggEfMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMB0GA1UdDgQWBBRiAaEnDRdMqrg75VohIdLfznTaPDAZBgNVHSAEEjAQMA4GCmCGSAGG%2FwkBAQEwADBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwQwYIKwYBBQUHAQEENzA1MDMGCCsGAQUFBzAChidodHRwOi8vc3RpY3IucmJibmlkaHViLmNvbS9yYmJuX2ljYS5jcnQwHwYDVR0jBBgwFoAUj9%2BXO0VyeUL6rz3v3FI1QGCMpNgwFgYIKwYBBQUHARoECjAIoAYWBDYyOUowCgYIKoZIzj0EAwIDRwAwRAIgMo%2B%2F57KC9Nzs4EvR3nHRHL4ZZyotDaLLXZ8Oo3xTM0sCIC6qURs%2FG9soNwB4d6lN0lMpGwK02CGw%2BfcE1UGN%2FgGn)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |


Generated: 28 Apr 23 02:17 UTC