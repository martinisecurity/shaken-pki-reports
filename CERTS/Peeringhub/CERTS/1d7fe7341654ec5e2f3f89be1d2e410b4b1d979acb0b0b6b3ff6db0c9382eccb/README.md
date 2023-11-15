# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 088K 2023-05

Tested At: 15 Nov 23 18:05 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 192 day(s)\
Subject: CN=SHAKEN 088K 2023-05, O=SmartTel Inc, C=US\
Issuer: CN=Peeringhub Inc SHAKEN Intermediate CA 2, OU=Certification Authorities, O=Peeringhub Inc, C=US\
Link: https://prov.smarttelinc.net/shaken/cert.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC%2BzCCAqCgAwIBAgIQcAh5%2FWFRHlRPTYs3ccJYvzAKBggqhkjOPQQDAjB8MQswCQYDVQQGEwJVUzEXMBUGA1UECgwOUGVlcmluZ2h1YiBJbmMxIjAgBgNVBAsMGUNlcnRpZmljYXRpb24gQXV0aG9yaXRpZXMxMDAuBgNVBAMMJ1BlZXJpbmdodWIgSW5jIFNIQUtFTiBJbnRlcm1lZGlhdGUgQ0EgMjAeFw0yMzA1MjUyMDM5MDRaFw0yNDA1MjQyMDM5MDRaMEIxCzAJBgNVBAYTAlVTMRUwEwYDVQQKDAxTbWFydFRlbCBJbmMxHDAaBgNVBAMME1NIQUtFTiAwODhLIDIwMjMtMDUwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATqVX%2FZbLZYEJfI82bJb9buGzK61rMmDXophoT9ZEsG6P2W6cSi2UHAWVtHm2PggH5eS7YGy%2BvKFdtrctztRKHdo4IBPDCCATgwDgYDVR0PAQH%2FBAQDAgeAMAwGA1UdEwEB%2FwQCMAAwHQYDVR0OBBYEFHSB1rCDdyN7CQl4XwGkaovwFYACMB8GA1UdIwQYMBaAFK6hc1GIKVcRygyp9LEKbk64S00HMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAWBggrBgEFBQcBGgQKMAigBhYEMDg4SzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwCgYIKoZIzj0EAwIDSQAwRgIhAKmJFX%2F%2FwhdapAMwfVIqidDVCthsr%2B2uDMBa3EyRYx7tAiEA2mUX1EX0chm3ugVivsVBdy5ysDQmDVP%2BPGan0SBcxYA%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |


Generated: 15 Nov 23 18:10 UTC