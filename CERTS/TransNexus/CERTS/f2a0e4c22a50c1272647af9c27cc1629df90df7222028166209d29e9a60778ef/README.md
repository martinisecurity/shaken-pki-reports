# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 111K

Tested At: 07 Dec 22 18:46 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -60 day(s)\
Subject: CN=SHAKEN 111K, OU=SHAKEN, O=Cloud Voice Tel, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/ae8a8a64-668f-4934-823d-42ebc587f82a/af862ae5ed3fbfa249b1c23dfb7b7ec6.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC8DCCApegAwIBAgIQX32alcYKJpfQUcyBb8wqpjAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMDExNTQzNTBaFw0yMjEwMDgxNTQzNDlaME4xCzAJBgNVBAYTAlVTMRgwFgYDVQQKEw9DbG91ZCBWb2ljZSBUZWwxDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDExMUswWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAARBm7C3%2BV6jMsruP4m2t%2BxncAjJUQpDnEqxz%2Bj6a2aRNZXue%2BKkNBWt6lkKihJIRmd1SXIaA%2F9dbRz1LGfBqw2No4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCAIAwHQYDVR0OBBYEFMG1vxAAJsIsabc1VtvoOx2wZq1wMB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDExMUswCgYIKoZIzj0EAwIDRwAwRAIgDy6mUNyOYNMEJhb%2F4ZIhRAlTHvXpYN0FTe0OLfyrgdYCIEcFnyL8x9Dhnhcc7B9KaR7YQEYsLl%2BLfLsXG%2Fe%2FTSB7)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |


Generated: 07 Dec 22 18:54 UTC