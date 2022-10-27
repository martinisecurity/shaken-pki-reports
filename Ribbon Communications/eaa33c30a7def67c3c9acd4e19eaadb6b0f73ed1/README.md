# STIR/SHAKEN CA Ecosystem Compliance
## Ribbon Communications

### Certificate eaa33c30a7def67c3c9acd4e19eaadb6b0f73ed1
Tested At: 2022-10-27 00:06:56 +0000 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 206 day(s)\
Subject: CN=Veracity SHAKEN 716D, OU=Voice, O=Veracity Networks LLC, C=US\
Issuer: CN=SHAKEN Ribbon Issuing CA, OU=Certification Authorities, O=Ribbon Communications, C=US

Link: https://prod001-cr.rbbnidhub.com/VaMEsVjGRz/May20222023-7f9cc6dca255376339d3370091b2fdde

View: [Click to view](https://understandingwebpki.com/?cert=MIIC8jCCApmgAwIBAgIQf5zG3KJVN2M50zcAkbL93jAKBggqhkjOPQQDAjB0MQswCQYDVQQGEwJVUzEeMBwGA1UEChMVUmliYm9uIENvbW11bmljYXRpb25zMSIwIAYDVQQLExlDZXJ0aWZpY2F0aW9uIEF1dGhvcml0aWVzMSEwHwYDVQQDExhTSEFLRU4gUmliYm9uIElzc3VpbmcgQ0EwHhcNMjIwNTIwMTQzMjUwWhcNMjMwNTIwMTQzMjQ5WjBcMQswCQYDVQQGEwJVUzEeMBwGA1UECgwVVmVyYWNpdHkgTmV0d29ya3MgTExDMQ4wDAYDVQQLDAVWb2ljZTEdMBsGA1UEAwwUVmVyYWNpdHkgU0hBS0VOIDcxNkQwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATHqObgDONxVFhyinyqKiQ%2BoHZ6AeZxVCueo8QLkCE0Kd16SK5ojXLFpFUR8vyb%2FM3iDynoZ1yN2YjyS0EcOH4co4IBIzCCAR8wDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwHQYDVR0OBBYEFI4O0WCF2a5K2WzylSAlIL3yKgtRMBkGA1UdIAQSMBAwDgYKYIZIAYb%2FCQEBATAAMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDBDBggrBgEFBQcBAQQ3MDUwMwYIKwYBBQUHMAKGJ2h0dHA6Ly9zdGljci5yYmJuaWRodWIuY29tL3JiYm5faWNhLmNydDAfBgNVHSMEGDAWgBSP35c7RXJ5QvqvPe%2FcUjVAYIyk2DAWBggrBgEFBQcBGgQKMAigBhYENzE2RDAKBggqhkjOPQQDAgNHADBEAiAHM9%2BcSBjmO2UUhlILGJYDt4CPQJtPwFcEOapBO6Zq6AIgWx8UL%2BsxJSnxbR1HGrFIiJYKav3KSp3GAAerqCiPdSI%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_sti_extension_unknown | error | ATIS-1000080v4 | STI certificate shall not include extensions that are not specified |
| e_sti_crl_distribution | error | ATIS-1000080v4 | CRL Distribution Point shall be reachable if the requesting IP address within the program ACLs |
| e_sti_certificate_policies | error | ATIS-1000080v4 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| e_cp1_3_subject_sn | error | CPv1.3 | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| w_cp1_3_subject_rdn_unknown | warn | CPv1.3 | Only CN, C, O, and SERIALNUMBER can be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| e_cp1_3_ambiguous_identifier | error | CPv1.3 | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |

Generated: 27/10/2022 at 00:07:07