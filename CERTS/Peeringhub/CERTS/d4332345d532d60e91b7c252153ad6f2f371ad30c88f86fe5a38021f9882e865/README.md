# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Celerity Telecom SHAKEN 469K

Tested At: 02 Jun 23 01:10 UTC\
Initial Validity Period: 34 day(s)\
Remaining Validity Period: 0 day(s)\
Subject: CN=Celerity Telecom SHAKEN 469K, O=Celerity Telecom, L=SouthWest Ranches, ST=FL, C=US\
Issuer: CN=Peeringhub Inc SHAKEN Intermediate CA 2, OU=Certification Authorities, O=Peeringhub Inc, C=US\
Link: https://s3.amazonaws.com/certificates.peeringhub.io/469K/469K.crt

[View certificate details](https://understandingwebpki.com/?cert=MIIDMDCCAtagAwIBAgIQJGil9W7KTmslpC9AnxvWfDAKBggqhkjOPQQDAjB8MQswCQYDVQQGEwJVUzEXMBUGA1UECgwOUGVlcmluZ2h1YiBJbmMxIjAgBgNVBAsMGUNlcnRpZmljYXRpb24gQXV0aG9yaXRpZXMxMDAuBgNVBAMMJ1BlZXJpbmdodWIgSW5jIFNIQUtFTiBJbnRlcm1lZGlhdGUgQ0EgMjAeFw0yMzA0MjkxNDQ4MzVaFw0yMzA2MDExNTU5MTNaMHgxCzAJBgNVBAYTAlVTMQswCQYDVQQIDAJGTDEaMBgGA1UEBwwRU291dGhXZXN0IFJhbmNoZXMxGTAXBgNVBAoMEENlbGVyaXR5IFRlbGVjb20xJTAjBgNVBAMMHENlbGVyaXR5IFRlbGVjb20gU0hBS0VOIDQ2OUswWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATpeBI%2FfnjDwWYM%2BcME2Qv0%2B2H6alcnmdTJeCySlFZdb%2FzwwLdJ82pVzWzYEgWhUHc7dTRIYwzqldTDSqYy3H1Zo4IBPDCCATgwDgYDVR0PAQH%2FBAQDAgeAMAwGA1UdEwEB%2FwQCMAAwHQYDVR0OBBYEFPnS%2BfjpsgsyH%2BTtO8fc8lUgCyeoMB8GA1UdIwQYMBaAFK6hc1GIKVcRygyp9LEKbk64S00HMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAWBggrBgEFBQcBGgQKMAigBhYENDY5SzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwCgYIKoZIzj0EAwIDSAAwRQIhAKcPpigatWdPZI16bdbhzOd73%2FCDzt6%2F%2F41ktxVc24inAiBJaJFGw5o38WMXFkO72CUSX%2F9qEHgc4Amnov5BhsC8Nw%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |


Generated: 02 Jun 23 01:12 UTC