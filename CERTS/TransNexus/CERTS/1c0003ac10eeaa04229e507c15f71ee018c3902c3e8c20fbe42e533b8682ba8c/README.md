# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 505J

Tested At: 31 Oct 22 18:32 UTC\
Initial Validity Period: 90 day(s)\
Remaining Validity Period: 56 day(s)\
Subject: CN=SHAKEN 505J, OU=SHAKEN, O=HFA Services LLC dba Call48, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.transnexus.com/505J/ec7d1cb5-b714-4373-ad07-ef0e95a1e098.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIC%2FjCCAqOgAwIBAgIQVP4AElRAE1jGc3mM8cfK0TAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjA5MjYxODQyMjRaFw0yMjEyMjUxODQyMjNaMFoxCzAJBgNVBAYTAlVTMSQwIgYDVQQKExtIRkEgU2VydmljZXMgTExDIGRiYSBDYWxsNDgxDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDUwNUowWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAT%2BimYfRZpQKygpE%2F9uNyexL1bopg5hDmewEnAs0Y2Res6KZWQyClNyRniQHAuPLxm6ivsM3t3CTqTPDVc0rwnQo4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCAIAwHQYDVR0OBBYEFOTdOlOigmkla3IyIAeE87BgDTWCMB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDUwNUowCgYIKoZIzj0EAwIDSQAwRgIhAML7I%2BVvSnMZZomE4xo7IzK3rSQrf0OpbJUTg4SsZrBYAiEAoJn0P8Ed1U14nULOZIigYBcbQuo33LIu6bqkYkBpYEw%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| [w_pki_subject_rdn_unknown](../../ISSUES/w_pki_subject_rdn_unknown/README.md) | warn | SHAKEN PKI Best Practice | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_cp1_3_subject_sn](../../ISSUES/e_cp1_3_subject_sn/README.md) | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 31/10/2022 at 18:34:12