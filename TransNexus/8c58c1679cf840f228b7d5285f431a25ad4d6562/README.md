# STIR/SHAKEN CA Ecosystem Compliance
## TransNexus

### Certificate 8c58c1679cf840f228b7d5285f431a25ad4d6562
Tested At: 2022-10-26 21:00:18 +0000 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 322 day(s)\
Subject: CN=SHAKEN 706J, OU=SHAKEN, O=TCN, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US

Link: https://certificates.transnexus.com/706J/fdb13f3d-0f05-47c8-a047-af2f9e2f7af5.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIC5TCCAougAwIBAgIQX4iz5rOuXyo%2BWVXh6OUhkTAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjA5MTMxNzI0NDlaFw0yMzA5MTMxNzI0NDhaMEIxCzAJBgNVBAYTAlVTMQwwCgYDVQQKEwNUQ04xDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDcwNkowWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQx%2BKckfjDFK87CpTfCrmwX%2F4zrIt2dW4thkEjU9iESG2AusWrZyEkqtuZp3PpEnOu2ltogB49m134%2F3cLhMDUAo4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCAIAwHQYDVR0OBBYEFKewD5Maow5dvMl2kNna%2BBaJqvnoMB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDcwNkowCgYIKoZIzj0EAwIDSAAwRQIgC8%2F3oFUZ%2FygRr52l5y4AbJahiJ9r5o0290iHCDaLOYgCIQDT9Uz4XQIiMlQAZ97s14HaFdKRiHZ9epw2DtYH1yRkCw%3D%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_cp1_3_subject_sn | error | CPv1.3 | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| w_cp1_3_subject_rdn_unknown | warn | CPv1.3 | Only CN, C, O, and SERIALNUMBER can be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| e_incorrect_ku_encoding | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| e_sti_crl_distribution | error | ATIS-1000080v4 | CRL Distribution Point shall be reachable if the requesting IP address within the program ACLs |

Generated: 26/10/2022 at 21:01:13