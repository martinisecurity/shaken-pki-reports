# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Solarus SHAKEN 0974

Tested At: 15 Nov 23 16:12 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 215 day(s)\
Subject: CN=Solarus SHAKEN 0974, OU=STI, O=Solarus, C=US\
Issuer: CN=SHAKEN Ribbon Issuing CA, OU=Certification Authorities, O=Ribbon Communications, C=US\
Link: https://prod001-prod011-cr.rbbnidhub.com/z9Ih9bs4Rz/STI-0974-2023-1df96aa7006c73b0c4228c26728e82c7

[View certificate details](https://understandingwebpki.com/?cert=MIIC4zCCAoigAwIBAgIQHflqpwBsc7DEIowmco6CxzAKBggqhkjOPQQDAjB0MQswCQYDVQQGEwJVUzEeMBwGA1UEChMVUmliYm9uIENvbW11bmljYXRpb25zMSIwIAYDVQQLExlDZXJ0aWZpY2F0aW9uIEF1dGhvcml0aWVzMSEwHwYDVQQDExhTSEFLRU4gUmliYm9uIElzc3VpbmcgQ0EwHhcNMjMwNjE4MDQzMzUzWhcNMjQwNjE3MDQzMzUyWjBLMQswCQYDVQQGEwJVUzEQMA4GA1UECgwHU29sYXJ1czEMMAoGA1UECwwDU1RJMRwwGgYDVQQDDBNTb2xhcnVzIFNIQUtFTiAwOTc0MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAElO9TvUaucrTxuXWnCKQWXno9nTHnOOjvHMHrQxHHT2ubP3kVrSRcIEVgi2mtZaVecZGh25QG5bQb8Hc5bSe8BaOCASMwggEfMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMB0GA1UdDgQWBBQSRklY0lZqQzh3uSu3H59phydwtTAZBgNVHSAEEjAQMA4GCmCGSAGG%2FwkBAQEwADBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwQwYIKwYBBQUHAQEENzA1MDMGCCsGAQUFBzAChidodHRwOi8vc3RpY3IucmJibmlkaHViLmNvbS9yYmJuX2ljYS5jcnQwHwYDVR0jBBgwFoAUj9%2BXO0VyeUL6rz3v3FI1QGCMpNgwFgYIKwYBBQUHARoECjAIoAYWBDA5NzQwCgYIKoZIzj0EAwIDSQAwRgIhAMdcPFq9UyznbccJ16qeUs%2BxuIhBidQdeQvhcaoiRSjPAiEAotiSw05E%2Fid6TMnjG%2BpFP3H3ZZim5vQDoLXx9%2BJX2Pg%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |


Generated: 15 Nov 23 16:51 UTC