# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 700H

Tested At: 11 Jan 23 21:48 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -55 day(s)\
Subject: CN=SHAKEN 700H, OU=SHAKEN, O=Metro Fibernet LLC, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/22b6cee0-8559-4c73-8092-6eee861c4b49/1906fe208d804599c3af0a06b8702740.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC8zCCApqgAwIBAgIQSIVJcxOxUUX6Xb5ZAeLJjjAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjExMTAyMDIwMTNaFw0yMjExMTcyMDIwMTJaMFExCzAJBgNVBAYTAlVTMRswGQYDVQQKExJNZXRybyBGaWJlcm5ldCBMTEMxDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDcwMEgwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATUQ5mt1b6QLv0rYyREtLqyXiWV5KfuZse8pDLveQiLJxKjQG2iT7xwWrX%2FNr2%2BLGjFIKK03i%2BNfSkKJNeNbdi3o4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwHQYDVR0OBBYEFHsGTLT9SMu2jRZP0e5W2qPdax8TMB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDcwMEgwCgYIKoZIzj0EAwIDRwAwRAIgE5IOyLDhkcuM4vytT5TtygN9zJi9%2B7umtVgHf0K6V3ICIGwtJP689cJOJipIqsavNcEdFMfNwxVojZ%2BaF4JGBwqn)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 11 Jan 23 21:59 UTC