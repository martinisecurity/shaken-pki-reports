# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Longlines SHAKEN 1260

Tested At: 15 Nov 23 16:12 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 217 day(s)\
Subject: CN=Longlines SHAKEN 1260, OU=STI, O=Long Lines Broadband, C=US\
Issuer: CN=SHAKEN Ribbon Issuing CA, OU=Certification Authorities, O=Ribbon Communications, C=US\
Link: https://prod001-prod011-cr.rbbnidhub.com/qWFquhQ4gz/STI-202306-1260-37d5fc075821511b582497c396489b57

[View certificate details](https://understandingwebpki.com/?cert=MIIC8TCCApegAwIBAgIQN9X8B1ghURtYJJfDlkibVzAKBggqhkjOPQQDAjB0MQswCQYDVQQGEwJVUzEeMBwGA1UEChMVUmliYm9uIENvbW11bmljYXRpb25zMSIwIAYDVQQLExlDZXJ0aWZpY2F0aW9uIEF1dGhvcml0aWVzMSEwHwYDVQQDExhTSEFLRU4gUmliYm9uIElzc3VpbmcgQ0EwHhcNMjMwNjE5MjAxMzU1WhcNMjQwNjE4MjAxMzU0WjBaMQswCQYDVQQGEwJVUzEdMBsGA1UECgwUTG9uZyBMaW5lcyBCcm9hZGJhbmQxDDAKBgNVBAsMA1NUSTEeMBwGA1UEAwwVTG9uZ2xpbmVzIFNIQUtFTiAxMjYwMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE5CouNnxkZPBsEdcsZSOpn9kNWBv64GIOqdh7IIisyCDD2HbxWhe2r2huGg40WNKDFGWoEQJ%2Bkjmro1%2BiLlcN0aOCASMwggEfMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMB0GA1UdDgQWBBSQuzN5j%2FH4NaYPA2lQvWJVYHvSmzAZBgNVHSAEEjAQMA4GCmCGSAGG%2FwkBAQEwADBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwQwYIKwYBBQUHAQEENzA1MDMGCCsGAQUFBzAChidodHRwOi8vc3RpY3IucmJibmlkaHViLmNvbS9yYmJuX2ljYS5jcnQwHwYDVR0jBBgwFoAUj9%2BXO0VyeUL6rz3v3FI1QGCMpNgwFgYIKwYBBQUHARoECjAIoAYWBDEyNjAwCgYIKoZIzj0EAwIDSAAwRQIhAPhRmebRPhGMm%2Fs7u%2FbhJk2%2FTrH72SmadaB79C4LMx7WAiBcM5L1EBjZT0A9DDEvvfRXwYD13BV%2BI%2BuqgU1fvz1jeg%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |


Generated: 15 Nov 23 16:51 UTC