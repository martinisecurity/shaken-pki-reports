# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Ribbon SHAKEN 2086

Tested At: 21 Aug 23 20:16 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 290 day(s)\
Subject: CN=Ribbon SHAKEN 2086, OU=Network Operations, O=Hill Country Telephone Cooperative, C=US\
Issuer: CN=SHAKEN Ribbon Issuing CA, OU=Certification Authorities, O=Ribbon Communications, C=US\
Link: https://prod001-prod011-cr.rbbnidhub.com/JQApas6MRz/10000-54a5c7395fcf8ebacb0d352d1d375852

[View certificate details](https://understandingwebpki.com/?cert=MIIDCjCCArGgAwIBAgIQVKXHOV%2FPjrrLDTUtHTdYUjAKBggqhkjOPQQDAjB0MQswCQYDVQQGEwJVUzEeMBwGA1UEChMVUmliYm9uIENvbW11bmljYXRpb25zMSIwIAYDVQQLExlDZXJ0aWZpY2F0aW9uIEF1dGhvcml0aWVzMSEwHwYDVQQDExhTSEFLRU4gUmliYm9uIElzc3VpbmcgQ0EwHhcNMjMwNjA3MjAxMTQwWhcNMjQwNjA2MjAxMTM5WjB0MQswCQYDVQQGEwJVUzErMCkGA1UECgwiSGlsbCBDb3VudHJ5IFRlbGVwaG9uZSBDb29wZXJhdGl2ZTEbMBkGA1UECwwSTmV0d29yayBPcGVyYXRpb25zMRswGQYDVQQDDBJSaWJib24gU0hBS0VOIDIwODYwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAARGl3takMoP0Yudj00Xbbf3CPO%2BuoIZG%2B1EMQ31KAL85pVHJ%2Bf%2FiefUhkUodyg3tEnrc%2F6fG2AMR4hTXaFIbMLwo4IBIzCCAR8wDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwHQYDVR0OBBYEFCpvzFdmbbF4pZ9x8MCaaRuGXDOgMBkGA1UdIAQSMBAwDgYKYIZIAYb%2FCQEBATAAMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDBDBggrBgEFBQcBAQQ3MDUwMwYIKwYBBQUHMAKGJ2h0dHA6Ly9zdGljci5yYmJuaWRodWIuY29tL3JiYm5faWNhLmNydDAfBgNVHSMEGDAWgBSP35c7RXJ5QvqvPe%2FcUjVAYIyk2DAWBggrBgEFBQcBGgQKMAigBhYEMjA4NjAKBggqhkjOPQQDAgNHADBEAiBQ2yOlYIWNH%2B1m0vExeYN%2FdgUwc9y6lwiHu80rn%2FFPxgIgeVrXJPLlOgn82cQ%2BubCvl3CKARAaDXdt6p8mrJvKjCE%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |


Generated: 21 Aug 23 20:18 UTC