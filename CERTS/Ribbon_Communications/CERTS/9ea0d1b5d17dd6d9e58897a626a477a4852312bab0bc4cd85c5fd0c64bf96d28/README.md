# STIR/SHAKEN CA Ecosystem Compliance

## Certificate PGTelco SHAKEN 1718

Tested At: 15 Nov 23 16:12 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 136 day(s)\
Subject: CN=PGTelco SHAKEN 1718, OU=Voice, O=PGTelco, C=US\
Issuer: CN=SHAKEN Ribbon Issuing CA, OU=Certification Authorities, O=Ribbon Communications, C=US\
Link: https://prod001-prod011-cr.rbbnidhub.com/dglpr26GRz/10000-6e486e81a61b3067206cd1cc03b9f1d9

[View certificate details](https://understandingwebpki.com/?cert=MIIC5DCCAoqgAwIBAgIQbkhugaYbMGcgbNHMA7nx2TAKBggqhkjOPQQDAjB0MQswCQYDVQQGEwJVUzEeMBwGA1UEChMVUmliYm9uIENvbW11bmljYXRpb25zMSIwIAYDVQQLExlDZXJ0aWZpY2F0aW9uIEF1dGhvcml0aWVzMSEwHwYDVQQDExhTSEFLRU4gUmliYm9uIElzc3VpbmcgQ0EwHhcNMjMwMzMwMjA0MDQ0WhcNMjQwMzI5MjA0MDQzWjBNMQswCQYDVQQGEwJVUzEQMA4GA1UECgwHUEdUZWxjbzEOMAwGA1UECwwFVm9pY2UxHDAaBgNVBAMME1BHVGVsY28gU0hBS0VOIDE3MTgwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAARAzrNUMvk2BSRu0HZfbMs3Eks6JtejLinJXM7ii0CUx5wuobBuTNY0OLen1InJtkUKWOvacnyV4XC2C0huYQFBo4IBIzCCAR8wDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwHQYDVR0OBBYEFHHNdKpu7H%2FUn3odt%2F85zz1lO5vsMBkGA1UdIAQSMBAwDgYKYIZIAYb%2FCQEBATAAMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDBDBggrBgEFBQcBAQQ3MDUwMwYIKwYBBQUHMAKGJ2h0dHA6Ly9zdGljci5yYmJuaWRodWIuY29tL3JiYm5faWNhLmNydDAfBgNVHSMEGDAWgBSP35c7RXJ5QvqvPe%2FcUjVAYIyk2DAWBggrBgEFBQcBGgQKMAigBhYEMTcxODAKBggqhkjOPQQDAgNIADBFAiBHn%2BJFP3Rdr3c%2FVm0OddYaZPs%2FUumEjfCM2PXCpg79ogIhALYm2%2FrSC2koHt8rjr31jDQexIGWlLBovT3pKYVoZQp9)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 15 Nov 23 16:51 UTC