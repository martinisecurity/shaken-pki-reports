# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Siskiyou SHAKEN 2339

Tested At: 24 Nov 23 11:11 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: -93 day(s)\
Subject: CN=Siskiyou SHAKEN 2339, OU=STI, O=Siskiyou Telephone Company, C=US\
Issuer: CN=SHAKEN Ribbon Issuing CA, OU=Certification Authorities, O=Ribbon Communications, C=US\
Link: https://prod001-prod011-cr.rbbnidhub.com/3MCvDamnRz/Aug222023-6e684a540e092145d414fcce5decbbfd

[View certificate details](https://understandingwebpki.com/?cert=MIIC9TCCApygAwIBAgIQbmhKVA4JIUXUFPzOXey7%2FTAKBggqhkjOPQQDAjB0MQswCQYDVQQGEwJVUzEeMBwGA1UEChMVUmliYm9uIENvbW11bmljYXRpb25zMSIwIAYDVQQLExlDZXJ0aWZpY2F0aW9uIEF1dGhvcml0aWVzMSEwHwYDVQQDExhTSEFLRU4gUmliYm9uIElzc3VpbmcgQ0EwHhcNMjIwODIyMTYwNDExWhcNMjMwODIyMTYwNDEwWjBfMQswCQYDVQQGEwJVUzEjMCEGA1UECgwaU2lza2l5b3UgVGVsZXBob25lIENvbXBhbnkxDDAKBgNVBAsMA1NUSTEdMBsGA1UEAwwUU2lza2l5b3UgU0hBS0VOIDIzMzkwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASSgcHTe9tr54Tpb%2BDt17%2FQWYxmJ1erv1Qe4GSjduL5Xjzf5ydbzM4v7te0RcYvj1wm6sUrmn5TZ6gAyWnIePA3o4IBIzCCAR8wDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwHQYDVR0OBBYEFMIiXAGlrfac95dd2oc%2BQ0JTw5EPMBkGA1UdIAQSMBAwDgYKYIZIAYb%2FCQEBATAAMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDBDBggrBgEFBQcBAQQ3MDUwMwYIKwYBBQUHMAKGJ2h0dHA6Ly9zdGljci5yYmJuaWRodWIuY29tL3JiYm5faWNhLmNydDAfBgNVHSMEGDAWgBSP35c7RXJ5QvqvPe%2FcUjVAYIyk2DAWBggrBgEFBQcBGgQKMAigBhYEMjMzOTAKBggqhkjOPQQDAgNHADBEAiA1VD7KAvZd%2BPCnsC7e1tiE0hEoLYR%2BUcCyLVXSPbQLKgIgSDfNa1AKhrGXpTlymDS%2BUDaBYOUh%2B6B6%2FoOswX5NI9I%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension contains an invalid OID value: 2.16.840.1.114569.1.1.1. Available OIDs: 2.16.840.1.114569.1.1.3, 2.16.840.1.114569.1.1.4 |
| [e_atis_ext_crl_distribution_struct](../../ISSUES/e_atis_ext_crl_distribution_struct/README.md) | error | ATIS1000080 | CRL Distribution Point shall contain a CRLIssuer field |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 24 Nov 23 11:17 UTC