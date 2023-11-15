# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Etex STI SHAKEN 3196

Tested At: 15 Nov 23 18:05 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 140 day(s)\
Subject: CN=Etex STI SHAKEN 3196, OU=Etex, O=Etex, C=US\
Issuer: CN=SHAKEN Ribbon Issuing CA, OU=Certification Authorities, O=Ribbon Communications, C=US\
Link: https://prod001-prod011-cr.rbbnidhub.com/DdpJt757Rz/STI-202304-3196-17c4d8150bc3db2f710070278713bc57

[View certificate details](https://understandingwebpki.com/?cert=MIIC4TCCAoegAwIBAgIQF8TYFQvD2y9xAHAnhxO8VzAKBggqhkjOPQQDAjB0MQswCQYDVQQGEwJVUzEeMBwGA1UEChMVUmliYm9uIENvbW11bmljYXRpb25zMSIwIAYDVQQLExlDZXJ0aWZpY2F0aW9uIEF1dGhvcml0aWVzMSEwHwYDVQQDExhTSEFLRU4gUmliYm9uIElzc3VpbmcgQ0EwHhcNMjMwNDAzMjEwODA2WhcNMjQwNDAyMjEwODA1WjBKMQswCQYDVQQGEwJVUzENMAsGA1UECgwERXRleDENMAsGA1UECwwERXRleDEdMBsGA1UEAwwURXRleCBTVEkgU0hBS0VOIDMxOTYwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAARWnncOqNZklMGkbLnSl%2FfquixLJ69KxyT4giJA5AtkSkYC6xmpzUXLPKwazOY9XrKEDwL1r26%2B7pcLWJSPjU3lo4IBIzCCAR8wDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwHQYDVR0OBBYEFBgSl6tFmjcQM%2B5Jdo%2FVnJ3YATr5MBkGA1UdIAQSMBAwDgYKYIZIAYb%2FCQEBATAAMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDBDBggrBgEFBQcBAQQ3MDUwMwYIKwYBBQUHMAKGJ2h0dHA6Ly9zdGljci5yYmJuaWRodWIuY29tL3JiYm5faWNhLmNydDAfBgNVHSMEGDAWgBSP35c7RXJ5QvqvPe%2FcUjVAYIyk2DAWBggrBgEFBQcBGgQKMAigBhYEMzE5NjAKBggqhkjOPQQDAgNIADBFAiB00RiXywW%2BDROgz3ag59CQ5hmpZqEJkfe0ao%2FbYpYqLwIhAI8v68yr95IQCMs10KwuuEHTCOxy8SgqxkQHo6%2F2OZSX)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 15 Nov 23 18:10 UTC