# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Etex STI SHAKEN 3196

Tested At: 12 Apr 23 01:39 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: -5 day(s)\
Subject: CN=Etex STI SHAKEN 3196, O=Etex, C=US\
Issuer: CN=SHAKEN Ribbon Issuing CA, OU=Certification Authorities, O=Ribbon Communications, C=US\
Link: https://prod001-cr.rbbnidhub.com/DdpJt757Rz/3196-2545b56cf6ba8ebe16fd7e049730421b

[View certificate details](https://understandingwebpki.com/?cert=MIIC0TCCAnigAwIBAgIQJUW1bPa6jr4W%2FX4ElzBCGzAKBggqhkjOPQQDAjB0MQswCQYDVQQGEwJVUzEeMBwGA1UEChMVUmliYm9uIENvbW11bmljYXRpb25zMSIwIAYDVQQLExlDZXJ0aWZpY2F0aW9uIEF1dGhvcml0aWVzMSEwHwYDVQQDExhTSEFLRU4gUmliYm9uIElzc3VpbmcgQ0EwHhcNMjIwNDA2MTgyMDM2WhcNMjMwNDA2MTgyMDM1WjA7MQswCQYDVQQGEwJVUzENMAsGA1UECgwERXRleDEdMBsGA1UEAwwURXRleCBTVEkgU0hBS0VOIDMxOTYwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAR8F47f3f99gZgMjHa6mcn1j%2BXwP%2FMePAoZ3myA7xpkYRc4JM38V4V4sG4cb8l2ggNs7WPLNAq2x8CrV9EkYEfWo4IBIzCCAR8wDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwHQYDVR0OBBYEFKKsLdOZzGYrtrFmejWsdFD4g17pMBkGA1UdIAQSMBAwDgYKYIZIAYb%2FCQEBATAAMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDBDBggrBgEFBQcBAQQ3MDUwMwYIKwYBBQUHMAKGJ2h0dHA6Ly9zdGljci5yYmJuaWRodWIuY29tL3JiYm5faWNhLmNydDAfBgNVHSMEGDAWgBSP35c7RXJ5QvqvPe%2FcUjVAYIyk2DAWBggrBgEFBQcBGgQKMAigBhYEMzE5NjAKBggqhkjOPQQDAgNHADBEAiBNjENknwPD7OzbNaNlgiW8Kav9%2FRgQ%2FQ4RLC1xWl1XoAIgY7FCO7mMq6W2eRPGe4Cxq4hfOECz8ez1PN7P4j7rcQY%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |


Generated: 12 Apr 23 01:46 UTC