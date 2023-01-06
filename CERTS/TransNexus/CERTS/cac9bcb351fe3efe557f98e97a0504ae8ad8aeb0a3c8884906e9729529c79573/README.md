# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 2550

Tested At: 06 Jan 23 02:53 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -89 day(s)\
Subject: CN=SHAKEN 2550, OU=SHAKEN, O=123.Net\\, Inc, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/2e3276e5-a5c5-43c7-ae11-3a226bf1b5ed/ccae7beeb2fafe38badbfe4104d488d2.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC7jCCApSgAwIBAgIQaNW3CHcC7CeRQMsw1LB9GTAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMDExNDMwMTlaFw0yMjEwMDgxNDMwMThaMEsxCzAJBgNVBAYTAlVTMRUwEwYDVQQKEwwxMjMuTmV0LCBJbmMxDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDI1NTAwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQ8QQXLS4fJcyr8F57MFjKFsaM7quBB99O9EJr5OOtBMJxM3kq8pt8N1OtyqgcjLKE%2BcALK0v6KIrjWxFJFIhjUo4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCAIAwHQYDVR0OBBYEFGT8b2NhsbGLg19VjwXdRDPmUB2IMB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDI1NTAwCgYIKoZIzj0EAwIDSAAwRQIhALGlePeRG%2B92yGgeTBqQEnj2yZaFDrOXu8oCrrPhlhQnAiBnU%2Bgo5Jtw4lxOEC1nBVSNHxFKhb2LxAIEw25OSDO1Ag%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |


Generated: 06 Jan 23 03:03 UTC