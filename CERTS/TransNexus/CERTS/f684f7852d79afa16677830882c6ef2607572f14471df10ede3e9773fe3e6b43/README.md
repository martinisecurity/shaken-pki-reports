# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 120K

Tested At: 02 Jun 23 01:02 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -201 day(s)\
Subject: CN=SHAKEN 120K, OU=SHAKEN, O=VC3, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/bcb89748-38e2-45ae-8b54-7efd23525de2/3701279a130218d9cf05b967e8e024e3.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC5DCCAougAwIBAgIQcPoX1FS9pyv07XueVsG5MDAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjExMDUyMjI0MzFaFw0yMjExMTIyMjI0MzBaMEIxCzAJBgNVBAYTAlVTMQwwCgYDVQQKEwNWQzMxDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDEyMEswWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATOz7rc356T96dKSzxr5ecZmZQ6WojYCiS8hCcoJp5Z8qJxo8AkT5SqceV3hiZPmDJtFd%2BC9fKCtRjd8CA5V2%2Bjo4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwHQYDVR0OBBYEFDM75Sh0dP50bKZI5GofuVCxV%2FPQMB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDEyMEswCgYIKoZIzj0EAwIDRwAwRAIgCU%2Fb%2BZFRV%2BXQTKrnedrb9IBvFEVPj%2FGE8buqQGUoFbMCIAl%2FkVzMv7sYLPJkC4mkrAN9%2BjT2PAHOTPPmPGA6zpvU)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 02 Jun 23 01:12 UTC