# STIR/SHAKEN CA Ecosystem Compliance
## TransNexus

### Certificate c612e0f95b5d7ac7c60edf495a4cbb05014f0c85
Tested At: 2022-10-28 18:54:51 +0000 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: 4 day(s)\
Subject: CN=SHAKEN 2550, OU=SHAKEN, O=123.Net\\, Inc, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US

Link: https://certificates.clearip.com/2e3276e5-a5c5-43c7-ae11-3a226bf1b5ed/8cf027bfba0180817348b19dfc538290.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIC7jCCApSgAwIBAgIQffLv4NauAOMIFp6I8z48wTAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMjUxNDMzMzVaFw0yMjExMDExNDMzMzRaMEsxCzAJBgNVBAYTAlVTMRUwEwYDVQQKEwwxMjMuTmV0LCBJbmMxDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDI1NTAwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQNsCLX7lXQOH1AE1cyLsfkc6a%2BrHM4E5eNHeQI34hhVjT94YlZKj3228Fmp9envGuJFBMbdtNXxHje%2FxUHv1Mjo4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCAIAwHQYDVR0OBBYEFNe2p8t9zxWf008U5hHxJt%2Bjc0QdMB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDI1NTAwCgYIKoZIzj0EAwIDSAAwRQIhANcpCtjrxUo6gVvNAUNsRTwF5S5q4NJ2cTJxb%2Fk0wAcpAiBCTZgcokV9QD25ZM%2Fz31%2B2iXhE5LWH0tMQ3nQ9jDGFNA%3D%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| w_pki_subject_rdn_unknown | warn | SHAKEN PKI Best Practice | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| e_incorrect_ku_encoding | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| e_cp1_3_subject_sn | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |

Generated: 28/10/2022 at 18:55:01