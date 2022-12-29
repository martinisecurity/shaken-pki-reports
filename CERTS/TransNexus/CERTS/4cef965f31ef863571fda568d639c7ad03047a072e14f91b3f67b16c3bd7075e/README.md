# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 970J

Tested At: 29 Dec 22 07:35 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -52 day(s)\
Subject: CN=SHAKEN 970J, OU=SHAKEN, O=Procomm Online Inc, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/795acb3b-a8b6-46c4-8ba6-82b33e759e7b/0472866da4fbd975377364ba11030eb9.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC8jCCApqgAwIBAgIQTh89rN%2F35%2BxJMl%2BWKy%2BKuTAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMzAyMzUxNThaFw0yMjExMDYyMzUxNTdaMFExCzAJBgNVBAYTAlVTMRswGQYDVQQKExJQcm9jb21tIE9ubGluZSBJbmMxDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDk3MEowWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQrzOZ%2F6ke5sHMNjkHcYbIt3iE4MQ9OkEw7w%2Fe5ZNVgfRQscaigmjklqc1Y7Z7l8WcxPJ3mgCZpoIhmm9WybxO8o4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwHQYDVR0OBBYEFNYPFj72QIXp7mf2SC0uNraTQn2vMB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDk3MEowCgYIKoZIzj0EAwIDRgAwQwIgbvKz4ce30LEMzLCcqRXBw9h8XBWlkKvwcjJop1W9jwcCHyO9vjdOnQ4I24wS80n0ZSLO4oIzOtgaatOs6ZZU%2Ftg%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 29 Dec 22 07:47 UTC