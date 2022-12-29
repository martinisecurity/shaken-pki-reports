# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 7453

Tested At: 29 Dec 22 07:35 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -59 day(s)\
Subject: CN=SHAKEN 7453, OU=SHAKEN, O=TPx, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/2a26bd25-8973-44ba-9b94-a8e4716b3888/4c9cfd7c1be4440cb2d95dc133ef4e14.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC5DCCAougAwIBAgIQVJ0HRikLU16m101qMiHnZTAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMjMyMDE4NDZaFw0yMjEwMzAyMDE4NDVaMEIxCzAJBgNVBAYTAlVTMQwwCgYDVQQKEwNUUHgxDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDc0NTMwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAR01T0tegzT8zxVSggxukhDVP7DNU05ETQ64V4em6wXBmu9iy48RsRQHDOWOKzU0IV6Htz60dPaXCZzoOLDlEW5o4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCAIAwHQYDVR0OBBYEFOWD0POj0a30%2FZ%2FYvvjOp5Ek2JqyMB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDc0NTMwCgYIKoZIzj0EAwIDRwAwRAIgSfmRd1pjr3A4ug%2Bjp1Vt8QctAa0OaxvisewhfvbxP%2F8CIE9gGF%2BJrmas0OG0h3JQZZTSaA%2BKw8MghiNKejPnu7qW)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 29 Dec 22 07:47 UTC