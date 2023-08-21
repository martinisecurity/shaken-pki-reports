# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Ironton SHAKEN 1234 0175

Tested At: 21 Aug 23 20:16 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 268 day(s)\
Subject: CN=Ironton SHAKEN 1234 0175, OU=Central Office, O=Ironton Telephone, C=US\
Issuer: CN=SHAKEN Ribbon Issuing CA, OU=Certification Authorities, O=Ribbon Communications, C=US\
Link: https://prod001-prod011-cr.rbbnidhub.com/rOzx_GsVRz/0175-529f0eaf99604f407acaafd164cd7e9d

[View certificate details](https://understandingwebpki.com/?cert=MIIC%2FDCCAqKgAwIBAgIQUp8Or5lgT0B6yq%2FRZM1%2BnTAKBggqhkjOPQQDAjB0MQswCQYDVQQGEwJVUzEeMBwGA1UEChMVUmliYm9uIENvbW11bmljYXRpb25zMSIwIAYDVQQLExlDZXJ0aWZpY2F0aW9uIEF1dGhvcml0aWVzMSEwHwYDVQQDExhTSEFLRU4gUmliYm9uIElzc3VpbmcgQ0EwHhcNMjMwNTE2MTkzNTE0WhcNMjQwNTE1MTkzNTEzWjBlMQswCQYDVQQGEwJVUzEaMBgGA1UECgwRSXJvbnRvbiBUZWxlcGhvbmUxFzAVBgNVBAsMDkNlbnRyYWwgT2ZmaWNlMSEwHwYDVQQDDBhJcm9udG9uIFNIQUtFTiAxMjM0IDAxNzUwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAT5713psuaZBbv1Co100GlMulyZeXtgCXOBWWfcs6LphW%2Biw8nc2lGF3D%2FGvT3qIWbWW5Y%2Fh0xUOwRZpKiQrOCCo4IBIzCCAR8wDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwHQYDVR0OBBYEFIWZr34AIP9Xf7cMgmSFWPr0zY5zMBkGA1UdIAQSMBAwDgYKYIZIAYb%2FCQEBATAAMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDBDBggrBgEFBQcBAQQ3MDUwMwYIKwYBBQUHMAKGJ2h0dHA6Ly9zdGljci5yYmJuaWRodWIuY29tL3JiYm5faWNhLmNydDAfBgNVHSMEGDAWgBSP35c7RXJ5QvqvPe%2FcUjVAYIyk2DAWBggrBgEFBQcBGgQKMAigBhYEMDE3NTAKBggqhkjOPQQDAgNIADBFAiA6G7Fuy4ldTElaL%2BB3gSmuoG2xC0LaQKRkU7lQ9U2mZgIhAPKic7s2DnxzgDPoEWTP2O%2B%2F2tYzEKQRinwYWkYD1ech)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 0175' |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |


Generated: 21 Aug 23 20:18 UTC