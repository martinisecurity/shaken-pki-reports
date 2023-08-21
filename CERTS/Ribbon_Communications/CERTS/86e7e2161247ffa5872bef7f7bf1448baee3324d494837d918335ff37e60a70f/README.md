# STIR/SHAKEN CA Ecosystem Compliance

## Certificate ColoradoValley SHAKEN 2059

Tested At: 21 Aug 23 20:16 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 220 day(s)\
Subject: CN=ColoradoValley SHAKEN 2059, OU=STI, O=Colorado Valley Communications, C=US\
Issuer: CN=SHAKEN Ribbon Issuing CA, OU=Certification Authorities, O=Ribbon Communications, C=US\
Link: https://prod001-prod011-cr.rbbnidhub.com/4R6YvcB4gz/STI-202303-2059-574780b87e516b53aa6a2bce87b791a6

[View certificate details](https://understandingwebpki.com/?cert=MIIDADCCAqagAwIBAgIQV0eAuH5Ra1OqaivOh7eRpjAKBggqhkjOPQQDAjB0MQswCQYDVQQGEwJVUzEeMBwGA1UEChMVUmliYm9uIENvbW11bmljYXRpb25zMSIwIAYDVQQLExlDZXJ0aWZpY2F0aW9uIEF1dGhvcml0aWVzMSEwHwYDVQQDExhTSEFLRU4gUmliYm9uIElzc3VpbmcgQ0EwHhcNMjMwMzI4MjIwNTU1WhcNMjQwMzI3MjIwNTU0WjBpMQswCQYDVQQGEwJVUzEnMCUGA1UECgweQ29sb3JhZG8gVmFsbGV5IENvbW11bmljYXRpb25zMQwwCgYDVQQLDANTVEkxIzAhBgNVBAMMGkNvbG9yYWRvVmFsbGV5IFNIQUtFTiAyMDU5MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEJCfC8cqO5nMXt2tTrnm4s6vPnxvAqyoJKaK7AIcyGaHYVcyzhTBMJG9ViQUUD7JUpFdByvmSjLX6G94TBLkVOqOCASMwggEfMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMB0GA1UdDgQWBBTMbAKqCe7a%2BXwj6nLC2y1c0i9BJTAZBgNVHSAEEjAQMA4GCmCGSAGG%2FwkBAQEwADBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwQwYIKwYBBQUHAQEENzA1MDMGCCsGAQUFBzAChidodHRwOi8vc3RpY3IucmJibmlkaHViLmNvbS9yYmJuX2ljYS5jcnQwHwYDVR0jBBgwFoAUj9%2BXO0VyeUL6rz3v3FI1QGCMpNgwFgYIKwYBBQUHARoECjAIoAYWBDIwNTkwCgYIKoZIzj0EAwIDSAAwRQIgZHhy5vvvET92IXcZWMAluW%2Frb%2F11k28fqftF9F8UIkUCIQDFzRmpMo4QIu1LWM%2FziOM%2B7wkyoi0qszRj6dEjJ86uEQ%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |


Generated: 21 Aug 23 20:18 UTC