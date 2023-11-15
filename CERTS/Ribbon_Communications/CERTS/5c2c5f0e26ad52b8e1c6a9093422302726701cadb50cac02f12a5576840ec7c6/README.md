# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Eastex SHAKEN 2068

Tested At: 15 Nov 23 16:37 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 63 day(s)\
Subject: CN=Eastex SHAKEN 2068, OU=Voice, O=Eastex, C=US\
Issuer: CN=SHAKEN Ribbon Issuing CA, OU=Certification Authorities, O=Ribbon Communications, C=US\
Link: https://prod001-prod011-cr.rbbnidhub.com/T0ezuXpngz/STI-2023-2068-3d0b382600afa29c04bd4fc3edfe2038

[View certificate details](https://understandingwebpki.com/?cert=MIIC4zCCAoigAwIBAgIQPQs4JgCvopwEvU%2FD7f4gODAKBggqhkjOPQQDAjB0MQswCQYDVQQGEwJVUzEeMBwGA1UEChMVUmliYm9uIENvbW11bmljYXRpb25zMSIwIAYDVQQLExlDZXJ0aWZpY2F0aW9uIEF1dGhvcml0aWVzMSEwHwYDVQQDExhTSEFLRU4gUmliYm9uIElzc3VpbmcgQ0EwHhcNMjMwMTE2MjAyNzIzWhcNMjQwMTE2MjAyNzIyWjBLMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGRWFzdGV4MQ4wDAYDVQQLDAVWb2ljZTEbMBkGA1UEAwwSRWFzdGV4IFNIQUtFTiAyMDY4MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE%2Fa8Y3LIYGyUYfb8ZmdBQ05X7dF3cD89Rl2HBeYKnE8wZwK5kwiribqjmphBnuQYewy7XJAbz2%2FXkzrn5KN7I8KOCASMwggEfMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMB0GA1UdDgQWBBSAkJ9fELIGRzg4l%2BEuDCmOVAYXaDAZBgNVHSAEEjAQMA4GCmCGSAGG%2FwkBAQEwADBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwQwYIKwYBBQUHAQEENzA1MDMGCCsGAQUFBzAChidodHRwOi8vc3RpY3IucmJibmlkaHViLmNvbS9yYmJuX2ljYS5jcnQwHwYDVR0jBBgwFoAUj9%2BXO0VyeUL6rz3v3FI1QGCMpNgwFgYIKwYBBQUHARoECjAIoAYWBDIwNjgwCgYIKoZIzj0EAwIDSQAwRgIhAPsB8XYaZ9RMcgD4HevussuxXiUbNni3P%2BQIjgGdJRUBAiEAhJ%2BT6kaY0y5QtO1Xq509yZjTWzrZKD%2BLT1jSOpzBsLI%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 15 Nov 23 17:17 UTC