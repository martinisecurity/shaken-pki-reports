# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 841J

Tested At: 11 Jan 23 21:48 UTC\
Initial Validity Period: 14 day(s)\
Remaining Validity Period: -94 day(s)\
Subject: CN=SHAKEN 841J, OU=SHAKEN, O=Securus Technologies LLC, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.transnexus.com/841J/9d7f66f7-57ce-45d7-8287-9e09ae7b3383.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC%2BjCCAqCgAwIBAgIQRUXxclbGAbwYMReMp1iYPjAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjA5MjUwNDE0NThaFw0yMjEwMDkwNDE0NTdaMFcxCzAJBgNVBAYTAlVTMSEwHwYDVQQKExhTZWN1cnVzIFRlY2hub2xvZ2llcyBMTEMxDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDg0MUowWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAARFgGVro1z%2FK68trl%2Fk24YebnXslYevHBkdTWEt0y%2Fnp%2BbvbmXYWgySt9K8V7JQzkGKH2VJE7aEG1l5lCdUvBpmo4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCAIAwHQYDVR0OBBYEFLtMRK%2BlyCqvrnwsat%2FWIXYwf3VxMB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDg0MUowCgYIKoZIzj0EAwIDSAAwRQIhALx9fvAb2W3kgh%2FmQ9YKhCK8gpbt%2FvP8MU8XFjrfbClDAiA5pw4AuBBtKDwXafPQ31NohGMYs8pF%2BTvGa1P9kDg%2FLg%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 11 Jan 23 21:59 UTC