# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 7453

Tested At: 07 Jan 23 19:08 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -74 day(s)\
Subject: CN=SHAKEN 7453, OU=SHAKEN, O=TPx, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/2a26bd25-8973-44ba-9b94-a8e4716b3888/355303e5eba4a3c0d8078939b31e7c5a.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC5TCCAougAwIBAgIQSEj2iiUwE4vJXPttZAGUJTAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMTcyMDE4MjdaFw0yMjEwMjQyMDE4MjZaMEIxCzAJBgNVBAYTAlVTMQwwCgYDVQQKEwNUUHgxDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDc0NTMwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAARK%2BbhqGuwIRP2ID4CKCEeHxG1RiZSgYKjCYxyOXRae9VEpi%2F%2BEVhJG%2FrfLxHBk%2FZaNaFd7K3ZIloln5HUejZxBo4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCAIAwHQYDVR0OBBYEFOPv6Z%2BYrVwstIydy306reGTB7R1MB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDc0NTMwCgYIKoZIzj0EAwIDSAAwRQIhANNaKyexaMLznkIJq5JOEw3%2BwdVXgZXvbrED0ul20wcYAiA33y%2BMCPs%2Fa7vrcRFdEJle627%2BJi5Fm53fmwFVWeX9GA%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 07 Jan 23 19:18 UTC