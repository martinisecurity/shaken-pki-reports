# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 4036

Tested At: 10 Nov 22 06:40 UTC\
Initial Validity Period: 180 day(s)\
Remaining Validity Period: 4 day(s)\
Subject: CN=SHAKEN 4036, OU=SHAKEN, O=ATT, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://cert.sticr.att.net:8443/certs/att/6d021e67-e6a1-4821-bf7a-19491029a6ca

[View certificate details](https://understandingwebpki.com/?cert=MIIC5jCCAougAwIBAgIQUXLKIoq4jcYv9SEMioYeGDAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjA1MTcxODQ3MjZaFw0yMjExMTMxODQ3MjVaMEIxCzAJBgNVBAYTAlVTMQwwCgYDVQQKEwNBVFQxDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDQwMzYwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAS6jZAWuqTTTNihx3HwoCLW%2BFyYQNnRtKwZj00mCnolHLFMe7%2BNGJkD4D37mWsF4WxpjuZbvZn%2FdqQwyhscs7Q8o4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCAIAwHQYDVR0OBBYEFEt5bm34mhAIpWU0oOxs19nm6kasMB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDQwMzYwCgYIKoZIzj0EAwIDSQAwRgIhAJGjH1RItjzZWcQG566kpy9VB9uoiwn2trtDFflTCvfXAiEApyPhmJzrnOaLLYziBYOVb7ygkqOb97ujfqmFEjOgvrc%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 10 Nov 22 06:43 UTC