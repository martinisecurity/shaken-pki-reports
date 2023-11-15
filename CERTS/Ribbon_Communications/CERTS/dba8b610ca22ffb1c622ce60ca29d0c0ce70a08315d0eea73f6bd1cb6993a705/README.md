# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Peoples SHAKEN 2130

Tested At: 15 Nov 23 18:05 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 211 day(s)\
Subject: CN=Peoples SHAKEN 2130, OU=STI, O=Peoples, C=US\
Issuer: CN=SHAKEN Ribbon Issuing CA, OU=Certification Authorities, O=Ribbon Communications, C=US\
Link: https://prod001-prod011-cr.rbbnidhub.com/R3kZfzj7gz/STI-202306-2130-3e01a84e6a320c2e0d9baf5cbeb11f38

[View certificate details](https://understandingwebpki.com/?cert=MIIC4jCCAoigAwIBAgIQPgGoTmoyDC4Nm69cvrEfODAKBggqhkjOPQQDAjB0MQswCQYDVQQGEwJVUzEeMBwGA1UEChMVUmliYm9uIENvbW11bmljYXRpb25zMSIwIAYDVQQLExlDZXJ0aWZpY2F0aW9uIEF1dGhvcml0aWVzMSEwHwYDVQQDExhTSEFLRU4gUmliYm9uIElzc3VpbmcgQ0EwHhcNMjMwNjE0MTcwMjM3WhcNMjQwNjEzMTcwMjM2WjBLMQswCQYDVQQGEwJVUzEQMA4GA1UECgwHUGVvcGxlczEMMAoGA1UECwwDU1RJMRwwGgYDVQQDDBNQZW9wbGVzIFNIQUtFTiAyMTMwMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEb4r5%2B80r6JBYXUbdVNbnLvR6h2sTCgNk7q3GxWv9Pa%2Fymn%2B0Os%2F0a7zPv%2BM3zCFwyFs6ymCjNCE6sYaEq%2F79S6OCASMwggEfMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMB0GA1UdDgQWBBSee8VC2ZXQvwzcuy%2BqtnwFWPil3DAZBgNVHSAEEjAQMA4GCmCGSAGG%2FwkBAQEwADBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwQwYIKwYBBQUHAQEENzA1MDMGCCsGAQUFBzAChidodHRwOi8vc3RpY3IucmJibmlkaHViLmNvbS9yYmJuX2ljYS5jcnQwHwYDVR0jBBgwFoAUj9%2BXO0VyeUL6rz3v3FI1QGCMpNgwFgYIKwYBBQUHARoECjAIoAYWBDIxMzAwCgYIKoZIzj0EAwIDSAAwRQIgVc3FqxgMP3j0U9fWXvYVuT7AGFN%2B1UZXxuDJUEi5h%2FsCIQC5x6aDKSPUOyv16EpZyWGlxCJ6U8v77tXx5MsZf%2F8yDg%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |


Generated: 15 Nov 23 18:10 UTC