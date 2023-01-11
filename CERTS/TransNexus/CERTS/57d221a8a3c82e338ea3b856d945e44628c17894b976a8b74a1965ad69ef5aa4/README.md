# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 983J

Tested At: 11 Jan 23 21:48 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -63 day(s)\
Subject: CN=SHAKEN 983J, OU=SHAKEN, O=ESI, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/59ebb7c1-25bd-4dfc-9794-fcb104b2f66a/4edcb44e0360ed67fbd30405e73b6e21.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC5DCCAougAwIBAgIQVKOB%2BYc5TRRXgrIL%2FO12uzAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjExMDIyMDE1MDRaFw0yMjExMDkyMDE1MDNaMEIxCzAJBgNVBAYTAlVTMQwwCgYDVQQKEwNFU0kxDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDk4M0owWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAARZooxa%2Fi21lT8cra33HdAruAYiWGHDHE3AHbkpAAss6vsyCLBshGzq0sSbAoW%2FM%2Feg8C%2Fqj4w8MZzmd7Sc6vvNo4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwHQYDVR0OBBYEFKsbYw30IumipyTqS77v88PTJi%2BSMB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDk4M0owCgYIKoZIzj0EAwIDRwAwRAIgNRLyALrykxbvM0gbj74Qe7Z3MoR6kUMYK1X8pHZDaqICIBZmS%2F6jk49PaUWb4N12UpPpyQB3NekbKDwJzkZnmtkT)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 11 Jan 23 21:59 UTC