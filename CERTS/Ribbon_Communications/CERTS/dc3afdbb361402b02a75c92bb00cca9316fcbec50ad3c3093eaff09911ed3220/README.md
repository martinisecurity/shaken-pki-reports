# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Siskiyou SHAKEN 2339

Tested At: 05 Jan 23 18:26 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 229 day(s)\
Subject: CN=Siskiyou SHAKEN 2339, OU=STI, O=Siskiyou Telephone Company, C=US\
Issuer: CN=SHAKEN Ribbon Issuing CA, OU=Certification Authorities, O=Ribbon Communications, C=US\
Link: https://prod001-prod011-cr.rbbnidhub.com/3MCvDamnRz/Aug222023-6e684a540e092145d414fcce5decbbfd

[View certificate details](https://understandingwebpki.com/?cert=MIIC9TCCApygAwIBAgIQbmhKVA4JIUXUFPzOXey7%2FTAKBggqhkjOPQQDAjB0MQswCQYDVQQGEwJVUzEeMBwGA1UEChMVUmliYm9uIENvbW11bmljYXRpb25zMSIwIAYDVQQLExlDZXJ0aWZpY2F0aW9uIEF1dGhvcml0aWVzMSEwHwYDVQQDExhTSEFLRU4gUmliYm9uIElzc3VpbmcgQ0EwHhcNMjIwODIyMTYwNDExWhcNMjMwODIyMTYwNDEwWjBfMQswCQYDVQQGEwJVUzEjMCEGA1UECgwaU2lza2l5b3UgVGVsZXBob25lIENvbXBhbnkxDDAKBgNVBAsMA1NUSTEdMBsGA1UEAwwUU2lza2l5b3UgU0hBS0VOIDIzMzkwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASSgcHTe9tr54Tpb%2BDt17%2FQWYxmJ1erv1Qe4GSjduL5Xjzf5ydbzM4v7te0RcYvj1wm6sUrmn5TZ6gAyWnIePA3o4IBIzCCAR8wDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwHQYDVR0OBBYEFMIiXAGlrfac95dd2oc%2BQ0JTw5EPMBkGA1UdIAQSMBAwDgYKYIZIAYb%2FCQEBATAAMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDBDBggrBgEFBQcBAQQ3MDUwMwYIKwYBBQUHMAKGJ2h0dHA6Ly9zdGljci5yYmJuaWRodWIuY29tL3JiYm5faWNhLmNydDAfBgNVHSMEGDAWgBSP35c7RXJ5QvqvPe%2FcUjVAYIyk2DAWBggrBgEFBQcBGgQKMAigBhYEMjMzOTAKBggqhkjOPQQDAgNHADBEAiA1VD7KAvZd%2BPCnsC7e1tiE0hEoLYR%2BUcCyLVXSPbQLKgIgSDfNa1AKhrGXpTlymDS%2BUDaBYOUh%2B6B6%2FoOswX5NI9I%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 05 Jan 23 18:35 UTC