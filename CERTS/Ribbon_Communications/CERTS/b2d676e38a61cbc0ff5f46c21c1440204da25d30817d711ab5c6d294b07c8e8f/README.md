# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Thacker SHAKEN 0419

Tested At: 15 Nov 23 16:37 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 182 day(s)\
Subject: CN=Thacker SHAKEN 0419, OU=STI, O=Thacker-Grigsby Telephone Company\\, INC., C=US\
Issuer: CN=SHAKEN Ribbon Issuing CA, OU=Certification Authorities, O=Ribbon Communications, C=US\
Link: https://prod001-prod011-cr.rbbnidhub.com/RAdOrzknRz/STI-2305-0419-78a3190c5023734f0ee557f8c04b4450

[View certificate details](https://understandingwebpki.com/?cert=MIIDAjCCAqigAwIBAgIQeKMZDFAjc08O5Vf4wEtEUDAKBggqhkjOPQQDAjB0MQswCQYDVQQGEwJVUzEeMBwGA1UEChMVUmliYm9uIENvbW11bmljYXRpb25zMSIwIAYDVQQLExlDZXJ0aWZpY2F0aW9uIEF1dGhvcml0aWVzMSEwHwYDVQQDExhTSEFLRU4gUmliYm9uIElzc3VpbmcgQ0EwHhcNMjMwNTE1MjAxMzE1WhcNMjQwNTE0MjAxMzE0WjBrMQswCQYDVQQGEwJVUzEwMC4GA1UECgwnVGhhY2tlci1Hcmlnc2J5IFRlbGVwaG9uZSBDb21wYW55LCBJTkMuMQwwCgYDVQQLDANTVEkxHDAaBgNVBAMME1RoYWNrZXIgU0hBS0VOIDA0MTkwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAARNdqDJPJBmzSZR%2Fc3B2ya%2FOBirfelx3C1s%2Bl54nyHi7LJOJLAN%2BoUut7S1mof5IiIr9jsUgkWKdw6kn9xPPklUo4IBIzCCAR8wDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwHQYDVR0OBBYEFKiOgGUu2lZoAatYhS69zuz1y%2Bv5MBkGA1UdIAQSMBAwDgYKYIZIAYb%2FCQEBATAAMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDBDBggrBgEFBQcBAQQ3MDUwMwYIKwYBBQUHMAKGJ2h0dHA6Ly9zdGljci5yYmJuaWRodWIuY29tL3JiYm5faWNhLmNydDAfBgNVHSMEGDAWgBSP35c7RXJ5QvqvPe%2FcUjVAYIyk2DAWBggrBgEFBQcBGgQKMAigBhYEMDQxOTAKBggqhkjOPQQDAgNIADBFAiASB5jE5aG7t9WyaFihUHFB%2BCkzKszVXvinE6FSawxanAIhAOZGES6oVe%2BQZk1GNEN8Nirjat7Z9GsVvSv1tfAmpKg8)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |


Generated: 15 Nov 23 17:17 UTC