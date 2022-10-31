# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 278K

Tested At: 31 Oct 22 20:24 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: 1 day(s)\
Subject: CN=SHAKEN 278K, OU=SHAKEN, O=OIT, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/36952c34-78d8-4c32-9499-188a418f31ba/292c05077b7b6e80288715aea31e3fa1.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIC5jCCAougAwIBAgIQQHqDisH5axAaZ04u2yNlMTAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMjUwMTQxMzJaFw0yMjExMDEwMTQxMzFaMEIxCzAJBgNVBAYTAlVTMQwwCgYDVQQKEwNPSVQxDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDI3OEswWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASIvE8yZvhl9WKkCzLNuIb0b8pVh%2BqIt9c6tiX4OPJnafmg2%2FT%2BuhEJo1HK37ky0fMlWbk0JdWXdFblXuQgk%2B4qo4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCAIAwHQYDVR0OBBYEFKIzXyQ4mRrLy6OpSdf1MyicCTmoMB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDI3OEswCgYIKoZIzj0EAwIDSQAwRgIhAI94pyoWdA0oNihKCevJwRYLy8tQBmjOaUnBYl8P6bo%2BAiEAx2T4Ftv6yAfJJk0PbnnkDc4wZtWjjKp9uYvKqqSEqnA%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_cp1_3_subject_sn](../../ISSUES/e_cp1_3_subject_sn/README.md) | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| [w_pki_subject_rdn_unknown](../../ISSUES/w_pki_subject_rdn_unknown/README.md) | warn | SHAKEN PKI Best Practice | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 31/10/2022 at 20:32:42