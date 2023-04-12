# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 727J

Tested At: 12 Apr 23 21:44 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -149 day(s)\
Subject: CN=SHAKEN 727J, OU=SHAKEN, O=SumoFiber, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/60e08b8b-2fd4-455e-9e9a-faa0470820d8/03e56de8fcf7a553e3a54bdb82bd163b.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC6jCCApGgAwIBAgIQa6DrRrqsY3w%2BE1TD3Y7%2B2zAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjExMDcyMDIyMjlaFw0yMjExMTQyMDIyMjhaMEgxCzAJBgNVBAYTAlVTMRIwEAYDVQQKEwlTdW1vRmliZXIxDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDcyN0owWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAARb0x1VQv4OBnXYcwInyaaRDuU1ENUSRlYjsU8fxHjvgB0ttdGf2a0f1xtfnx8GDVo1PcoopUKfV%2Foxiwwdo8mto4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwHQYDVR0OBBYEFA8IYSuiJxyfRNJL%2Fx8zfG8j%2F3l8MB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDcyN0owCgYIKoZIzj0EAwIDRwAwRAIgO2CQqjIfD%2B9vLhSelgsdKcAWvMjyxZIfCmSpFTGc8IkCICVn5mhHZSvkBT5nWRWes4n1CSwZgE4mLzvPxIR%2B%2F1mV)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 12 Apr 23 22:02 UTC