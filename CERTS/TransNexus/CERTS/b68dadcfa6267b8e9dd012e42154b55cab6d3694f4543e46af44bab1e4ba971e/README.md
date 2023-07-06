# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 8526

Tested At: 06 Jul 23 13:57 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -223 day(s)\
Subject: CN=SHAKEN 8526, OU=SHAKEN, O=MetTel, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.transnexus.com/8526/f6978975-50bc-40df-b4a2-a1d7794b321f.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC5zCCAo6gAwIBAgIQW%2Ba5TEL9yKXSEDF6IWId9DAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMjYxMDQzMjFaFw0yMjExMjUxMDQzMjBaMEUxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKEwZNZXRUZWwxDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDg1MjYwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAARdu84I1Kt53t2DZkiw5Ms%2BMvqNcatkYJ0tUnqXV5O7BBkt2INbtNbTpHqoZQm7O6Rw8G6SPo4NUj6U3rPUmuNto4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCAIAwHQYDVR0OBBYEFFTD%2FOLSAW6XBysz47V8Sm8XUxFUMB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDg1MjYwCgYIKoZIzj0EAwIDRwAwRAIgRfQG8OySuAsY2kVYuZ1CPxF7c2ckaMWDMbyTEZVhjYcCIFwU7UIwtPfNZlPCQC7MF4XPa83%2FoMBk2zu5nu0myO5H)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |


Generated: 06 Jul 23 14:08 UTC