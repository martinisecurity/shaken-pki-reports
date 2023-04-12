# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 278K

Tested At: 12 Apr 23 01:02 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -201 day(s)\
Subject: CN=SHAKEN 278K, OU=SHAKEN, O=OIT, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/36952c34-78d8-4c32-9499-188a418f31ba/4de9a24de596a626dd4a02acbebbfed7.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC5TCCAougAwIBAgIQfJUmie1NHMhEzeAUlbVPODAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjA5MTYwMTM4MTJaFw0yMjA5MjMwMTM4MTFaMEIxCzAJBgNVBAYTAlVTMQwwCgYDVQQKEwNPSVQxDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDI3OEswWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASQVzx1M%2FPG8NGFuLlAURJrAuqcn6AsvNZUS9E60tpu1g1yoAJJE4LDY84ApsuKj7b1AFKLuU3AwFKaKFPn%2FIAao4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCAIAwHQYDVR0OBBYEFKWV3yWRgECBGCNedszmPzFIwL0OMB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDI3OEswCgYIKoZIzj0EAwIDSAAwRQIhAPN8Bw5m2AElzveJDiXPumOPkiA3xGf6XyD7QCTYc9Y3AiBo2NMfTOFXN3e6yraeXmdlRRFRPaiNFG%2BUf9B8NsgXfw%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |


Generated: 12 Apr 23 01:46 UTC