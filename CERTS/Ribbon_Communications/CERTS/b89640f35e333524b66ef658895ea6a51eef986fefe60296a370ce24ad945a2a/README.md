# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Haviland SHAKEN 5237

Tested At: 02 Jun 23 01:10 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 195 day(s)\
Subject: CN=Haviland SHAKEN 5237, OU=Support, O=Haviland, C=US\
Issuer: CN=SHAKEN Ribbon Issuing CA, OU=Certification Authorities, O=Ribbon Communications, C=US\
Link: https://prod001-prod011-cr.rbbnidhub.com/9f_1JJO7Rz/STI5237-684e51f5f56b220ae2d2a29a99fc377a

[View certificate details](https://understandingwebpki.com/?cert=MIIC6TCCAo6gAwIBAgIQaE5R9fVrIgri0qKamfw3ejAKBggqhkjOPQQDAjB0MQswCQYDVQQGEwJVUzEeMBwGA1UEChMVUmliYm9uIENvbW11bmljYXRpb25zMSIwIAYDVQQLExlDZXJ0aWZpY2F0aW9uIEF1dGhvcml0aWVzMSEwHwYDVQQDExhTSEFLRU4gUmliYm9uIElzc3VpbmcgQ0EwHhcNMjIxMjEzMTE0MDQ4WhcNMjMxMjEzMTE0MDQ3WjBRMQswCQYDVQQGEwJVUzERMA8GA1UECgwISGF2aWxhbmQxEDAOBgNVBAsMB1N1cHBvcnQxHTAbBgNVBAMMFEhhdmlsYW5kIFNIQUtFTiA1MjM3MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEN3%2FHCubLoECzjEr1DdbSxMRD39XYAWMOscZ2GtTdNi2CX0UqeTqteFRv3ji4slLGmXIJsK8spknp8ekD7jPKP6OCASMwggEfMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMB0GA1UdDgQWBBTSQP%2Fh5LP0NIi9WLKBYBFtnCCO%2BTAZBgNVHSAEEjAQMA4GCmCGSAGG%2FwkBAQEwADBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwQwYIKwYBBQUHAQEENzA1MDMGCCsGAQUFBzAChidodHRwOi8vc3RpY3IucmJibmlkaHViLmNvbS9yYmJuX2ljYS5jcnQwHwYDVR0jBBgwFoAUj9%2BXO0VyeUL6rz3v3FI1QGCMpNgwFgYIKwYBBQUHARoECjAIoAYWBDUyMzcwCgYIKoZIzj0EAwIDSQAwRgIhAOMl49kWiaoS9s3rZRxCvNVgjeE7XSMfkXbWzsnvJFUfAiEAhgYg1Qb9czOOorLRV3EPYWFOIfrV6wLyKvZClYkW%2FJI%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |


Generated: 02 Jun 23 01:12 UTC