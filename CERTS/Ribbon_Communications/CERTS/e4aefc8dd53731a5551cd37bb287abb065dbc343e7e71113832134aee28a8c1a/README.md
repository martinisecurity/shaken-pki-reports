# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Veracity SHAKEN 716D

Tested At: 15 Nov 23 16:12 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 189 day(s)\
Subject: CN=Veracity SHAKEN 716D, OU=Voice, O=Veracity Networks LLC, C=US\
Issuer: CN=SHAKEN Ribbon Issuing CA, OU=Certification Authorities, O=Ribbon Communications, C=US\
Link: https://prod001-prod011-cr.rbbnidhub.com/VaMEsVjGRz/0-6f633bc2813830d744b6236fe4d3b741

[View certificate details](https://understandingwebpki.com/?cert=MIIC8zCCApmgAwIBAgIQb2M7woE4MNdEtiNv5NO3QTAKBggqhkjOPQQDAjB0MQswCQYDVQQGEwJVUzEeMBwGA1UEChMVUmliYm9uIENvbW11bmljYXRpb25zMSIwIAYDVQQLExlDZXJ0aWZpY2F0aW9uIEF1dGhvcml0aWVzMSEwHwYDVQQDExhTSEFLRU4gUmliYm9uIElzc3VpbmcgQ0EwHhcNMjMwNTIzMDYwMzIzWhcNMjQwNTIyMDYwMzIyWjBcMQswCQYDVQQGEwJVUzEeMBwGA1UECgwVVmVyYWNpdHkgTmV0d29ya3MgTExDMQ4wDAYDVQQLDAVWb2ljZTEdMBsGA1UEAwwUVmVyYWNpdHkgU0hBS0VOIDcxNkQwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQvOhqV%2FrsIRCAUucf4GdN2mAYxBeUcTz1tTcniJpNP2792JW08rAo5IsC96%2FRnBfbvyZ2Mlqa3FvR8cEyINHNBo4IBIzCCAR8wDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwHQYDVR0OBBYEFE9mnwLwMJU%2FRJ6Jm3tU7DihQD6LMBkGA1UdIAQSMBAwDgYKYIZIAYb%2FCQEBATAAMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDBDBggrBgEFBQcBAQQ3MDUwMwYIKwYBBQUHMAKGJ2h0dHA6Ly9zdGljci5yYmJuaWRodWIuY29tL3JiYm5faWNhLmNydDAfBgNVHSMEGDAWgBSP35c7RXJ5QvqvPe%2FcUjVAYIyk2DAWBggrBgEFBQcBGgQKMAigBhYENzE2RDAKBggqhkjOPQQDAgNIADBFAiEAgdjhtCuffoJysKSU2CUV9Vxo0wJoBDtMx45Oe8GHNTECIEAffRQUxv3qfqH7SepydbDBPLj%2BD2SPkWfZv5tW6IEo)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 15 Nov 23 16:51 UTC