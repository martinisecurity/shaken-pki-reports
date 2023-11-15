# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Gearheart SHAKEN 0408

Tested At: 15 Nov 23 16:12 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 163 day(s)\
Subject: CN=Gearheart SHAKEN 0408, OU=Voice, O=Gearheart Communications, C=US\
Issuer: CN=SHAKEN Ribbon Issuing CA, OU=Certification Authorities, O=Ribbon Communications, C=US\
Link: https://prod001-prod011-cr.rbbnidhub.com/DUIDRwfVgz/10000-6fcc9648f6bec955b83d512a6d0830bb

[View certificate details](https://understandingwebpki.com/?cert=MIIC9jCCAp2gAwIBAgIQb8yWSPa%2ByVW4PVEqbQgwuzAKBggqhkjOPQQDAjB0MQswCQYDVQQGEwJVUzEeMBwGA1UEChMVUmliYm9uIENvbW11bmljYXRpb25zMSIwIAYDVQQLExlDZXJ0aWZpY2F0aW9uIEF1dGhvcml0aWVzMSEwHwYDVQQDExhTSEFLRU4gUmliYm9uIElzc3VpbmcgQ0EwHhcNMjMwNDI3MTYyMjQwWhcNMjQwNDI2MTYyMjM5WjBgMQswCQYDVQQGEwJVUzEhMB8GA1UECgwYR2VhcmhlYXJ0IENvbW11bmljYXRpb25zMQ4wDAYDVQQLDAVWb2ljZTEeMBwGA1UEAwwVR2VhcmhlYXJ0IFNIQUtFTiAwNDA4MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEOgh3E9QaQ%2FjTuIxXGPUpiDEFhwdoZgDV6qvQThipLPgoEaZfg8l9cOeIdeciYTn9PLcHAOC8N1m5tFYehNUUjqOCASMwggEfMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMB0GA1UdDgQWBBRJOBgc6M9LOHUy%2BDnqFXNC%2BvnpvzAZBgNVHSAEEjAQMA4GCmCGSAGG%2FwkBAQEwADBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwQwYIKwYBBQUHAQEENzA1MDMGCCsGAQUFBzAChidodHRwOi8vc3RpY3IucmJibmlkaHViLmNvbS9yYmJuX2ljYS5jcnQwHwYDVR0jBBgwFoAUj9%2BXO0VyeUL6rz3v3FI1QGCMpNgwFgYIKwYBBQUHARoECjAIoAYWBDA0MDgwCgYIKoZIzj0EAwIDRwAwRAIgQj35%2BIXsqVm%2F65zYPc8jZMRBsuOoiYcLNrU8HffJoCICIGIXLult49xANDXCrQ8JBTY8EGbJ2n2sCQCJlkIyOYgk)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 15 Nov 23 16:51 UTC