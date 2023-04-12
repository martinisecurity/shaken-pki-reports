# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 0172

Tested At: 12 Apr 23 01:03 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -142 day(s)\
Subject: CN=SHAKEN 0172, OU=SHAKEN, O=CallTower, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.transnexus.com/0172/c6ebb6b8-5f0d-4656-935f-cb13d73f6128.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC6jCCApGgAwIBAgIQboZ%2FkX3pdjGvMNuFa60jYDAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjExMTMyMjQwNTNaFw0yMjExMjAyMjQwNTJaMEgxCzAJBgNVBAYTAlVTMRIwEAYDVQQKEwlDYWxsVG93ZXIxDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDAxNzIwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQO2cWLWaN2s4oG3FJCyTtHHySTKkbNSsrRcucHw9E4ZMKo6aQqiqrFkC70cJj6vvLfcTy7blWToW4SrDSnOJxSo4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwHQYDVR0OBBYEFM6jbY6XI8By5qdeRO3kug3quIuZMB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDAxNzIwCgYIKoZIzj0EAwIDRwAwRAIgcPDbFe8t0sLO7p8HEt8TJGlKmDFYFU%2Ff4%2BLWOl96qCYCIG9YIlEEcF%2BRN3ohDxfkmknbPR9OLnRg1%2BcpxOcD3ZAB)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 12 Apr 23 01:46 UTC