# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Veracity SHAKEN 716D

Tested At: 31 Oct 22 16:42 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 201 day(s)\
Subject: CN=Veracity SHAKEN 716D, OU=Voice, O=Veracity Networks LLC, C=US\
Issuer: CN=SHAKEN Ribbon Issuing CA, OU=Certification Authorities, O=Ribbon Communications, C=US\
Link: https://prod001-cr.rbbnidhub.com/VaMEsVjGRz/May20222023-7f9cc6dca255376339d3370091b2fdde

View: [Click to view](https://understandingwebpki.com/?cert=MIIC8jCCApmgAwIBAgIQf5zG3KJVN2M50zcAkbL93jAKBggqhkjOPQQDAjB0MQswCQYDVQQGEwJVUzEeMBwGA1UEChMVUmliYm9uIENvbW11bmljYXRpb25zMSIwIAYDVQQLExlDZXJ0aWZpY2F0aW9uIEF1dGhvcml0aWVzMSEwHwYDVQQDExhTSEFLRU4gUmliYm9uIElzc3VpbmcgQ0EwHhcNMjIwNTIwMTQzMjUwWhcNMjMwNTIwMTQzMjQ5WjBcMQswCQYDVQQGEwJVUzEeMBwGA1UECgwVVmVyYWNpdHkgTmV0d29ya3MgTExDMQ4wDAYDVQQLDAVWb2ljZTEdMBsGA1UEAwwUVmVyYWNpdHkgU0hBS0VOIDcxNkQwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATHqObgDONxVFhyinyqKiQ%2BoHZ6AeZxVCueo8QLkCE0Kd16SK5ojXLFpFUR8vyb%2FM3iDynoZ1yN2YjyS0EcOH4co4IBIzCCAR8wDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwHQYDVR0OBBYEFI4O0WCF2a5K2WzylSAlIL3yKgtRMBkGA1UdIAQSMBAwDgYKYIZIAYb%2FCQEBATAAMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDBDBggrBgEFBQcBAQQ3MDUwMwYIKwYBBQUHMAKGJ2h0dHA6Ly9zdGljci5yYmJuaWRodWIuY29tL3JiYm5faWNhLmNydDAfBgNVHSMEGDAWgBSP35c7RXJ5QvqvPe%2FcUjVAYIyk2DAWBggrBgEFBQcBGgQKMAigBhYENzE2RDAKBggqhkjOPQQDAgNHADBEAiAHM9%2BcSBjmO2UUhlILGJYDt4CPQJtPwFcEOapBO6Zq6AIgWx8UL%2BsxJSnxbR1HGrFIiJYKav3KSp3GAAerqCiPdSI%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_cp1_3_subject_sn](../../ISSUES/e_cp1_3_subject_sn/README.md) | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_sti_certificate_policies](../../ISSUES/e_sti_certificate_policies/README.md) | error | ATIS&#x2011;1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_sti_extension_unknown](../../ISSUES/e_sti_extension_unknown/README.md) | error | ATIS&#x2011;1000080 | STI certificate shall not include extensions that are not specified |
| [w_pki_subject_rdn_unknown](../../ISSUES/w_pki_subject_rdn_unknown/README.md) | warn | SHAKEN PKI Best Practice | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_cp1_3_ambiguous_identifier](../../ISSUES/e_cp1_3_ambiguous_identifier/README.md) | error | United States SHAKEN CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |


Generated: 31/10/2022 at 16:43:22