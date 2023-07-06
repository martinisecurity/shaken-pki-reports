# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Triangle SHAKEN 2257

Tested At: 06 Jul 23 14:06 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 239 day(s)\
Subject: CN=Triangle SHAKEN 2257, OU=STI, O=Triangle, C=US\
Issuer: CN=SHAKEN Ribbon Issuing CA, OU=Certification Authorities, O=Ribbon Communications, C=US\
Link: https://prod001-prod011-cr.rbbnidhub.com/fh2Fl4xVgz/10000-457e9bfe46560b334453a78b38eb7826

[View certificate details](https://understandingwebpki.com/?cert=MIIC4zCCAoqgAwIBAgIQRX6b%2FkZWCzNEU6eLOOt4JjAKBggqhkjOPQQDAjB0MQswCQYDVQQGEwJVUzEeMBwGA1UEChMVUmliYm9uIENvbW11bmljYXRpb25zMSIwIAYDVQQLExlDZXJ0aWZpY2F0aW9uIEF1dGhvcml0aWVzMSEwHwYDVQQDExhTSEFLRU4gUmliYm9uIElzc3VpbmcgQ0EwHhcNMjMwMzAxMTkzMjQyWhcNMjQwMjI5MTkzMjQxWjBNMQswCQYDVQQGEwJVUzERMA8GA1UECgwIVHJpYW5nbGUxDDAKBgNVBAsMA1NUSTEdMBsGA1UEAwwUVHJpYW5nbGUgU0hBS0VOIDIyNTcwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAARlVWffeFOXoHusCxgevcSQ5Jg%2Bqbfx00xOVNevPdtkmT4v%2FPm1xpN6LMZuXI9DEUSLldyqje4tdx2qW7ruFW%2FMo4IBIzCCAR8wDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwHQYDVR0OBBYEFG3F4F2gxYnNVHGpD2ueBvnGDBNTMBkGA1UdIAQSMBAwDgYKYIZIAYb%2FCQEBATAAMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDBDBggrBgEFBQcBAQQ3MDUwMwYIKwYBBQUHMAKGJ2h0dHA6Ly9zdGljci5yYmJuaWRodWIuY29tL3JiYm5faWNhLmNydDAfBgNVHSMEGDAWgBSP35c7RXJ5QvqvPe%2FcUjVAYIyk2DAWBggrBgEFBQcBGgQKMAigBhYEMjI1NzAKBggqhkjOPQQDAgNHADBEAiBQ1oJm8cLSg3iUwzmzMTSs0I45Tle2p%2BWLw86Ozq3aIgIgcR8%2BXae6%2BTP%2BQTGRgQ2R%2B7gFy1UFvxEORQVfswp0Ql8%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |


Generated: 06 Jul 23 14:08 UTC