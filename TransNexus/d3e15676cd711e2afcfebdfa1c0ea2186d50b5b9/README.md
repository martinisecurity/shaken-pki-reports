# STIR/SHAKEN CA Ecosystem Compliance
## TransNexus

### Certificate d3e15676cd711e2afcfebdfa1c0ea2186d50b5b9
Tested At: 2022-10-26 20:22:00 +0000 UTC\
Initial Validity Period: 90 day(s)\
Remaining Validity Period: 60 day(s)\
Subject: CN=SHAKEN 505J, OU=SHAKEN, O=HFA Services LLC dba Call48, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US

Link: https://certificates.transnexus.com/505J/ec7d1cb5-b714-4373-ad07-ef0e95a1e098.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIC%2FjCCAqOgAwIBAgIQVP4AElRAE1jGc3mM8cfK0TAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjA5MjYxODQyMjRaFw0yMjEyMjUxODQyMjNaMFoxCzAJBgNVBAYTAlVTMSQwIgYDVQQKExtIRkEgU2VydmljZXMgTExDIGRiYSBDYWxsNDgxDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDUwNUowWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAT%2BimYfRZpQKygpE%2F9uNyexL1bopg5hDmewEnAs0Y2Res6KZWQyClNyRniQHAuPLxm6ivsM3t3CTqTPDVc0rwnQo4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCAIAwHQYDVR0OBBYEFOTdOlOigmkla3IyIAeE87BgDTWCMB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDUwNUowCgYIKoZIzj0EAwIDSQAwRgIhAML7I%2BVvSnMZZomE4xo7IzK3rSQrf0OpbJUTg4SsZrBYAiEAoJn0P8Ed1U14nULOZIigYBcbQuo33LIu6bqkYkBpYEw%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| w_cp1_3_subject_rdn_unknown | warn | CPv1.3 | Only CN, C, O, and SERIALNUMBER can be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| e_sti_crl_distribution | error | ATIS-1000080v4 | CRL Distribution Point shall be reachable if the requesting IP address within the program ACLs |
| e_incorrect_ku_encoding | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| e_cp1_3_subject_sn | error | CPv1.3 | STI certificate shall include a ‘serialNumber’ attribute along with the CN |

Generated: 26/10/2022 at 20:22:11