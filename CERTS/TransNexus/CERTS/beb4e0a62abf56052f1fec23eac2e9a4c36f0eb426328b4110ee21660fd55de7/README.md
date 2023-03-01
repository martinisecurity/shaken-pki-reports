# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 916J

Tested At: 01 Mar 23 18:12 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -133 day(s)\
Subject: CN=SHAKEN 916J, OU=SHAKEN, O=Fusion Networks, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/48c90e4a-a2cf-4e37-bd41-4c9bf99d32ac/f8191017203e97aa7f3bf75438273a72.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC8DCCApegAwIBAgIQU5d%2F13nHJ9zZLuz%2FSbuoGzAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMTEyMDIwMTRaFw0yMjEwMTgyMDIwMTNaME4xCzAJBgNVBAYTAlVTMRgwFgYDVQQKEw9GdXNpb24gTmV0d29ya3MxDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDkxNkowWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASa5xO3TmMuYyNRV4kwBHNlugiWjoNVQHrQMrmRej6eOi4f1t0s8bs%2BogEUtZbHPZYppxU%2FoM1NdJQNjWdsHnL%2Fo4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCAIAwHQYDVR0OBBYEFPIqfY5f2r%2F2CpA%2Fn3GgjJnbcex4MB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDkxNkowCgYIKoZIzj0EAwIDRwAwRAIgJAOHKWe48wzP8QsMKMrEcef9EYf1JmXmE1%2FMC%2FFddBkCIDaUfCVR0cXQ1xkQQS3Pmk%2BHr6NJSIAqsYrkLRePvEsO)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 01 Mar 23 18:22 UTC