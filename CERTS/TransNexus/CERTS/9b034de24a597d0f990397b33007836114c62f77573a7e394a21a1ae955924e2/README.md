# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 278K

Tested At: 31 Jan 23 21:39 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -112 day(s)\
Subject: CN=SHAKEN 278K, OU=SHAKEN, O=OIT, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/36952c34-78d8-4c32-9499-188a418f31ba/1db4b7629e187513dd8c82fe505a86e3.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC5jCCAougAwIBAgIQTlgks3ZvPVGLGxXLJ5WGOzAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMDQwMTM5MTlaFw0yMjEwMTEwMTM5MThaMEIxCzAJBgNVBAYTAlVTMQwwCgYDVQQKEwNPSVQxDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDI3OEswWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAStoztOrvbSQ5u7pvALECt69hCBGj6Yxb8FI%2Bo6gur90bmvKq8iIbUSKurRY%2FMYQHpKKV1lOo3DazjhuPlVWTywo4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCAIAwHQYDVR0OBBYEFBD22Az9AF3FD7JSR%2FzkKIat10hSMB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDI3OEswCgYIKoZIzj0EAwIDSQAwRgIhAKw1J5XHsIUAFEJoiynXFMawsNwrQtac2u0hBmDTosmsAiEAvGWGtdbE%2FJ%2FPw9oKq5d5zUOHLA6Cbu3%2FADedoIdTT9I%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |


Generated: 31 Jan 23 21:50 UTC