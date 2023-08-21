# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Empire SHAKEN 087H

Tested At: 21 Aug 23 20:16 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 115 day(s)\
Subject: CN=Empire SHAKEN 087H, OU=Network, O=Empire Telephone, C=US\
Issuer: CN=SHAKEN Ribbon Issuing CA, OU=Certification Authorities, O=Ribbon Communications, C=US\
Link: https://prod001-prod011-cr.rbbnidhub.com/Wlp54ZO4gz/087HNOV2023-7a28b8ee6335f2abf0116dbac19be755

[View certificate details](https://understandingwebpki.com/?cert=MIIC7zCCApSgAwIBAgIQeii47mM18qvwEW26wZvnVTAKBggqhkjOPQQDAjB0MQswCQYDVQQGEwJVUzEeMBwGA1UEChMVUmliYm9uIENvbW11bmljYXRpb25zMSIwIAYDVQQLExlDZXJ0aWZpY2F0aW9uIEF1dGhvcml0aWVzMSEwHwYDVQQDExhTSEFLRU4gUmliYm9uIElzc3VpbmcgQ0EwHhcNMjIxMjE0MTQxMzI5WhcNMjMxMjE0MTQxMzI4WjBXMQswCQYDVQQGEwJVUzEZMBcGA1UECgwQRW1waXJlIFRlbGVwaG9uZTEQMA4GA1UECwwHTmV0d29yazEbMBkGA1UEAwwSRW1waXJlIFNIQUtFTiAwODdIMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEBYT3%2BBfgL8uUtsp1Ou21bSpk7BL2OdhY35vyy2bjI8fN0%2FVp4qM9uszNdT%2BFOwXitHcYXiGRui%2BhOBV8Ia6FXKOCASMwggEfMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMB0GA1UdDgQWBBQHeMaqeE7cxDf%2BjpcYMzlCwGI3dTAZBgNVHSAEEjAQMA4GCmCGSAGG%2FwkBAQEwADBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwQwYIKwYBBQUHAQEENzA1MDMGCCsGAQUFBzAChidodHRwOi8vc3RpY3IucmJibmlkaHViLmNvbS9yYmJuX2ljYS5jcnQwHwYDVR0jBBgwFoAUj9%2BXO0VyeUL6rz3v3FI1QGCMpNgwFgYIKwYBBQUHARoECjAIoAYWBDA4N0gwCgYIKoZIzj0EAwIDSQAwRgIhAIfqvLrps1PMGy0l605mN1Rep2%2Bs6rcjLCgWfHWJQdkxAiEAn41QEcj34fAaPAfbicVz7iIUR5IonKb4%2Fc5vZvAEULY%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |


Generated: 21 Aug 23 20:18 UTC