# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Surrytel STI SHAKEN 0503

Tested At: 15 Nov 23 16:12 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 168 day(s)\
Subject: CN=Surrytel STI SHAKEN 0503, OU=Voice, O=Surrytel, C=US\
Issuer: CN=SHAKEN Ribbon Issuing CA, OU=Certification Authorities, O=Ribbon Communications, C=US\
Link: https://prod001-prod011-cr.rbbnidhub.com/x9lB3qynRz/STI-202305-0503-1a01332c0243d06b0cb7d77293786255

[View certificate details](https://understandingwebpki.com/?cert=MIIC6zCCApCgAwIBAgIQGgEzLAJD0GsMt9dyk3hiVTAKBggqhkjOPQQDAjB0MQswCQYDVQQGEwJVUzEeMBwGA1UEChMVUmliYm9uIENvbW11bmljYXRpb25zMSIwIAYDVQQLExlDZXJ0aWZpY2F0aW9uIEF1dGhvcml0aWVzMSEwHwYDVQQDExhTSEFLRU4gUmliYm9uIElzc3VpbmcgQ0EwHhcNMjMwNTAyMTMzMDEzWhcNMjQwNTAxMTMzMDEyWjBTMQswCQYDVQQGEwJVUzERMA8GA1UECgwIU3Vycnl0ZWwxDjAMBgNVBAsMBVZvaWNlMSEwHwYDVQQDDBhTdXJyeXRlbCBTVEkgU0hBS0VOIDA1MDMwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAARUuYvVQx2E8be82WyxYaxTnqEw12IPEPl2XeXVqRc2y2bvXx1swxVtArzSOxJZAzjzdQXESN5e6xAWZXOCWOE8o4IBIzCCAR8wDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwHQYDVR0OBBYEFIdquUUysgtC6H8NQddogk37%2Fr80MBkGA1UdIAQSMBAwDgYKYIZIAYb%2FCQEBATAAMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDBDBggrBgEFBQcBAQQ3MDUwMwYIKwYBBQUHMAKGJ2h0dHA6Ly9zdGljci5yYmJuaWRodWIuY29tL3JiYm5faWNhLmNydDAfBgNVHSMEGDAWgBSP35c7RXJ5QvqvPe%2FcUjVAYIyk2DAWBggrBgEFBQcBGgQKMAigBhYEMDUwMzAKBggqhkjOPQQDAgNJADBGAiEArahEUJtJpnOSGuL2DuUzlPeSR%2Bq8wZAIvr7I%2B7G9I%2BICIQDMuReBNs6uXpcTzrmREV8svJfirF%2BNtWWkzOLXoWcstA%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 15 Nov 23 16:51 UTC