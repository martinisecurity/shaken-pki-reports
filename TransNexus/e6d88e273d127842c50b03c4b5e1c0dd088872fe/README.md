# STIR/SHAKEN CA Ecosystem Compliance
## TransNexus

### Certificate e6d88e273d127842c50b03c4b5e1c0dd088872fe
Tested At: 2022-10-26 22:31:23 +0000 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: 4 day(s)\
Subject: CN=SHAKEN 7453, OU=SHAKEN, O=TPx, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US

Link: https://certificates.clearip.com/2a26bd25-8973-44ba-9b94-a8e4716b3888/4c9cfd7c1be4440cb2d95dc133ef4e14.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIC5DCCAougAwIBAgIQVJ0HRikLU16m101qMiHnZTAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMjMyMDE4NDZaFw0yMjEwMzAyMDE4NDVaMEIxCzAJBgNVBAYTAlVTMQwwCgYDVQQKEwNUUHgxDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDc0NTMwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAR01T0tegzT8zxVSggxukhDVP7DNU05ETQ64V4em6wXBmu9iy48RsRQHDOWOKzU0IV6Htz60dPaXCZzoOLDlEW5o4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCAIAwHQYDVR0OBBYEFOWD0POj0a30%2FZ%2FYvvjOp5Ek2JqyMB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDc0NTMwCgYIKoZIzj0EAwIDRwAwRAIgSfmRd1pjr3A4ug%2Bjp1Vt8QctAa0OaxvisewhfvbxP%2F8CIE9gGF%2BJrmas0OG0h3JQZZTSaA%2BKw8MghiNKejPnu7qW)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_sti_crl_distribution | error | ATIS-1000080v4 | CRL Distribution Point shall be reachable if the requesting IP address within the program ACLs |
| e_cp1_3_subject_sn | error | CPv1.3 | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| e_incorrect_ku_encoding | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| w_cp1_3_subject_rdn_unknown | warn | CPv1.3 | Only CN, C, O, and SERIALNUMBER can be included. Additional RNDs may introduce ambiguity and may not be verifiable |

Generated: 26/10/2022 at 22:31:35