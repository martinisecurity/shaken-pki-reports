# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 606F

Tested At: 06 Jul 23 13:55 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -262 day(s)\
Subject: CN=SHAKEN 606F, OU=SHAKEN, O=Global Data Systems Inc., C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/756cb700-f9d2-4a05-850e-c9dfe3e22de4/6410808e3d363b24e911e9dd36c6ef8d.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC%2BzCCAqCgAwIBAgIQWIVOh67onr3K1DgQFMQNszAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMDkyMTM5NTlaFw0yMjEwMTYyMTM5NThaMFcxCzAJBgNVBAYTAlVTMSEwHwYDVQQKExhHbG9iYWwgRGF0YSBTeXN0ZW1zIEluYy4xDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDYwNkYwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATbcgWK1g7Lpwegx2IKeFd8CEnvGnffJryk8%2FWkXYhnE%2FCGvMFx%2FbbgNG5NoatgQUPHaFWXkudYMyY0Z5WoFa4eo4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCAIAwHQYDVR0OBBYEFAIGsoRNUkr8TZhX9dbIl3%2FtMKmnMB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDYwNkYwCgYIKoZIzj0EAwIDSQAwRgIhAOw18Z3KkRhe54dtV49JbtdF7%2BKvQSovj%2FrN0OGv1YE0AiEAt0Xku27HtVw6mzfn0DICyOb8zseeXcO85%2FlXgCZbxsA%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 06 Jul 23 14:08 UTC