# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Randolph SHAKEN 0496

Tested At: 15 Nov 23 18:05 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 174 day(s)\
Subject: CN=Randolph SHAKEN 0496, OU=STI, O=Randolph Communications, C=US\
Issuer: CN=SHAKEN Ribbon Issuing CA, OU=Certification Authorities, O=Ribbon Communications, C=US\
Link: https://prod001-prod011-cr.rbbnidhub.com/y7KhPDB4Rz/STI-202305-0496-4d00c291e24d5a5d5c757e8035320452

[View certificate details](https://understandingwebpki.com/?cert=MIIC8TCCApmgAwIBAgIQTQDCkeJNWl1cdX6ANTIEUjAKBggqhkjOPQQDAjB0MQswCQYDVQQGEwJVUzEeMBwGA1UEChMVUmliYm9uIENvbW11bmljYXRpb25zMSIwIAYDVQQLExlDZXJ0aWZpY2F0aW9uIEF1dGhvcml0aWVzMSEwHwYDVQQDExhTSEFLRU4gUmliYm9uIElzc3VpbmcgQ0EwHhcNMjMwNTA4MTcxMzEyWhcNMjQwNTA3MTcxMzExWjBcMQswCQYDVQQGEwJVUzEgMB4GA1UECgwXUmFuZG9scGggQ29tbXVuaWNhdGlvbnMxDDAKBgNVBAsMA1NUSTEdMBsGA1UEAwwUUmFuZG9scGggU0hBS0VOIDA0OTYwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQj5uGBVpWaQ%2FctBn9YxiJ2QNFQxVSsotBoOYm2YKEa4zPrHXT18Guel3HbgjHsplCOOzs4XA0nTRoCsC26Vydxo4IBIzCCAR8wDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwHQYDVR0OBBYEFC9c3b0G55kDGd3trhCLrsY3F7HNMBkGA1UdIAQSMBAwDgYKYIZIAYb%2FCQEBATAAMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDBDBggrBgEFBQcBAQQ3MDUwMwYIKwYBBQUHMAKGJ2h0dHA6Ly9zdGljci5yYmJuaWRodWIuY29tL3JiYm5faWNhLmNydDAfBgNVHSMEGDAWgBSP35c7RXJ5QvqvPe%2FcUjVAYIyk2DAWBggrBgEFBQcBGgQKMAigBhYEMDQ5NjAKBggqhkjOPQQDAgNGADBDAh9bORteLSVSpVFUDsgz5JnlDwaMunrecljdjFIdL40YAiBIFWTPjV8kN3O7ok%2F%2FI95hRSfBXrgrzyPV0Tr0pn9KJQ%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 15 Nov 23 18:10 UTC