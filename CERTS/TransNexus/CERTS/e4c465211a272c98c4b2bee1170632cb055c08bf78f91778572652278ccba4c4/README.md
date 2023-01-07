# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 8526

Tested At: 07 Jan 23 19:08 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -68 day(s)\
Subject: CN=SHAKEN 8526, OU=SHAKEN, O=MetTel, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.transnexus.com/8526/ce6a0107-6101-4a05-9a81-ae55b304f8f0.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC6DCCAo6gAwIBAgIQWISHy8vY2ntrIoachDGX3jAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjA5MzAyMjM0MTRaFw0yMjEwMzAyMjM0MTNaMEUxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKEwZNZXRUZWwxDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDg1MjYwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATghpWXe1AxTzICKGVTHAF7a5pzKtRm5siPU8dk3NgHuUhKf755ysEoCc50K7EWaDySuNE1F%2Bu4yVxK8EXB1w%2Fdo4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCAIAwHQYDVR0OBBYEFEbSu8EZevVQKHFSatzoNRx%2FeSQOMB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDg1MjYwCgYIKoZIzj0EAwIDSAAwRQIgMzc97y23OnmHa3OAdSGMHcNHrV8TngCqW4ZWcgCu0QgCIQCV6Ke3er8x5hwBhSaosx47HhgQBIHY%2FIPiVIdKpUizOA%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 07 Jan 23 19:18 UTC