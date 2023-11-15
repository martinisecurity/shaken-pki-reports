# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Polar SHAKEN 1630

Tested At: 15 Nov 23 16:12 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 177 day(s)\
Subject: CN=Polar SHAKEN 1630, OU=Voice, O=Polar Communications Mutual Aid Corporation, C=US\
Issuer: CN=SHAKEN Ribbon Issuing CA, OU=Certification Authorities, O=Ribbon Communications, C=US\
Link: https://prod001-prod011-cr.rbbnidhub.com/87rmTaP4gz/10000-5f466faa9cc3dae51dd87191dd1f78b4

[View certificate details](https://understandingwebpki.com/?cert=MIIDBzCCAqygAwIBAgIQX0ZvqpzD2uUd2HGR3R94tDAKBggqhkjOPQQDAjB0MQswCQYDVQQGEwJVUzEeMBwGA1UEChMVUmliYm9uIENvbW11bmljYXRpb25zMSIwIAYDVQQLExlDZXJ0aWZpY2F0aW9uIEF1dGhvcml0aWVzMSEwHwYDVQQDExhTSEFLRU4gUmliYm9uIElzc3VpbmcgQ0EwHhcNMjMwNTEwMTgyNzI0WhcNMjQwNTA5MTgyNzIzWjBvMQswCQYDVQQGEwJVUzE0MDIGA1UECgwrUG9sYXIgQ29tbXVuaWNhdGlvbnMgTXV0dWFsIEFpZCBDb3Jwb3JhdGlvbjEOMAwGA1UECwwFVm9pY2UxGjAYBgNVBAMMEVBvbGFyIFNIQUtFTiAxNjMwMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE3SidbUPk47bbRDg%2FxjuVrEwKHMvP%2FbSCd8N%2FXU13JGIq0NqmLSR88dkwL1ZUhVnH13WJ3EIfrYFsu%2BIEYO5XDKOCASMwggEfMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMB0GA1UdDgQWBBT9Jqm86mXiEWZTyIo4tsdFui0hmDAZBgNVHSAEEjAQMA4GCmCGSAGG%2FwkBAQEwADBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwQwYIKwYBBQUHAQEENzA1MDMGCCsGAQUFBzAChidodHRwOi8vc3RpY3IucmJibmlkaHViLmNvbS9yYmJuX2ljYS5jcnQwHwYDVR0jBBgwFoAUj9%2BXO0VyeUL6rz3v3FI1QGCMpNgwFgYIKwYBBQUHARoECjAIoAYWBDE2MzAwCgYIKoZIzj0EAwIDSQAwRgIhAIw7qDpPNCEjWDY8DZbgoIGxwtwZ6bL1IsSKmcJ86%2FggAiEA6rJs%2BRHh3VVIVI8W%2FeUKqHoaJ4%2BtyPfD%2B1tczZTfd78%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 15 Nov 23 16:51 UTC