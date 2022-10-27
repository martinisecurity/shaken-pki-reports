# STIR/SHAKEN CA Ecosystem Compliance
## TransNexus

### Certificate 3889bee709733ff19f8acc2ae4a3761d288ffdfe
Tested At: 2022-10-27 21:25:37 +0000 UTC\
Initial Validity Period: 14 day(s)\
Remaining Validity Period: 7 day(s)\
Subject: CN=SHAKEN 841J, OU=SHAKEN, O=Securus Technologies LLC, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US

Link: https://certificates.transnexus.com/841J/f3c3ccf0-0018-4b22-bf77-8a30b9ffd68e.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIC%2BjCCAqCgAwIBAgIQeDJWfyVNF0mdJI%2F9aElKmzAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMjAwOTEzMThaFw0yMjExMDMwOTEzMTdaMFcxCzAJBgNVBAYTAlVTMSEwHwYDVQQKExhTZWN1cnVzIFRlY2hub2xvZ2llcyBMTEMxDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDg0MUowWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAARNo91X64NUmlHu7XGS8%2FBYC4OpGG6y7B44%2Bkme%2B3jhwCuKrURNfZM8aKLiHJPKzDaG5jPWQz1n2Q%2FFiJfFWKJyo4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCAIAwHQYDVR0OBBYEFC8u42jFW4tuWzeTHvglkF6BXKdZMB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDg0MUowCgYIKoZIzj0EAwIDSAAwRQIhAJvn2PIRofIR%2BU%2Bi7kgGpVCeU3KK86KsdGQbMNGy8sf7AiBfW1hvccc0Ivr%2FUxtBxfXJezRe9E08CSAx07c%2FFR6pdA%3D%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| w_cp1_3_subject_rdn_unknown | warn | United States SHAKEN CP | Only CN, C, O, and SERIALNUMBER can be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| e_incorrect_ku_encoding | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| e_cp1_3_subject_sn | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |

Generated: 27/10/2022 at 21:27:34