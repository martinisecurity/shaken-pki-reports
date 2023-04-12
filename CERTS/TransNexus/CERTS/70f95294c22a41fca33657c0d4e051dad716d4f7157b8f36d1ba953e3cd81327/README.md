# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 278K

Tested At: 12 Apr 23 21:44 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -168 day(s)\
Subject: CN=SHAKEN 278K, OU=SHAKEN, O=OIT, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/36952c34-78d8-4c32-9499-188a418f31ba/cb042f315bec56fc32f101c7cb686278.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC5DCCAougAwIBAgIQeNQZeXrURq%2Brc7MUuW4m6DAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMTkwMTQxMTJaFw0yMjEwMjYwMTQxMTFaMEIxCzAJBgNVBAYTAlVTMQwwCgYDVQQKEwNPSVQxDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDI3OEswWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAARicbH%2BHbg2hCp1RRSRgdo%2BBHME5bYWzU4OKOSNRq%2FxNUYIXFToUZNYX9DL3WeJVCcpm4ix1rWw7XPP8hHstAmCo4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCAIAwHQYDVR0OBBYEFPg610wgV%2BEDCDtobMYU3wKNB6BuMB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDI3OEswCgYIKoZIzj0EAwIDRwAwRAIgbjn5dPrMAXBbqyXF9h%2Ftc1MULPSBjhoOML92LojS9FMCIAJ18FYorYPbMYP%2BZsoRwv9duGPsWrGET1V4gQuWGD6j)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 12 Apr 23 22:02 UTC