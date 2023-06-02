# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 706J

Tested At: 02 Jun 23 01:03 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 104 day(s)\
Subject: CN=SHAKEN 706J, OU=SHAKEN, O=TCN, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.transnexus.com/706J/fdb13f3d-0f05-47c8-a047-af2f9e2f7af5.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC5TCCAougAwIBAgIQX4iz5rOuXyo%2BWVXh6OUhkTAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjA5MTMxNzI0NDlaFw0yMzA5MTMxNzI0NDhaMEIxCzAJBgNVBAYTAlVTMQwwCgYDVQQKEwNUQ04xDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDcwNkowWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQx%2BKckfjDFK87CpTfCrmwX%2F4zrIt2dW4thkEjU9iESG2AusWrZyEkqtuZp3PpEnOu2ltogB49m134%2F3cLhMDUAo4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCAIAwHQYDVR0OBBYEFKewD5Maow5dvMl2kNna%2BBaJqvnoMB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDcwNkowCgYIKoZIzj0EAwIDSAAwRQIgC8%2F3oFUZ%2FygRr52l5y4AbJahiJ9r5o0290iHCDaLOYgCIQDT9Uz4XQIiMlQAZ97s14HaFdKRiHZ9epw2DtYH1yRkCw%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 02 Jun 23 01:12 UTC