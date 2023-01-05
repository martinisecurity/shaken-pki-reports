# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 166K

Tested At: 05 Jan 23 18:07 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -57 day(s)\
Subject: CN=SHAKEN 166K, OU=SHAKEN, O=Gonthier Inc, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.transnexus.com/166K/b4f563ea-1dd9-40bc-be12-ef0190b82bad.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC7jCCApSgAwIBAgIQQA5yBjrcSUOikgjvm5tkOjAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMTAxODI3MThaFw0yMjExMDkxODI3MTdaMEsxCzAJBgNVBAYTAlVTMRUwEwYDVQQKEwxHb250aGllciBJbmMxDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDE2NkswWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAARDofGwS0vLXTX3s6ksrD7fKjcpM%2BI8MdJIEXbbWyqTJZKgIrUUd9UsuLxvPVL3zWyw8jAyMGPgi8VrZg9DhZfHo4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCAIAwHQYDVR0OBBYEFF8%2FJVlK1vjNhLBaLkXh6iJ9zecxMB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDE2NkswCgYIKoZIzj0EAwIDSAAwRQIgQo66q%2FMgdaOtqMRuf4I1rX3uRKgVuI0RQEfalCBUi6YCIQDDuvY87XFIwuywC2Iin7CUDnPvBFbuwYFZJ6b6ml%2BlxA%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 05 Jan 23 18:35 UTC