# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 726J

Tested At: 29 Dec 22 07:35 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -44 day(s)\
Subject: CN=SHAKEN 726J, OU=SHAKEN, O=Votacall Inc, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/88dbb645-62bf-4d40-960c-4e8e3223880f/69b0d26ffdfea695ce7cfe3464135ee0.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC7TCCApSgAwIBAgIQY3V%2B3cKq0POoB1J%2FNr0sXzAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjExMDcyMDIzMzhaFw0yMjExMTQyMDIzMzdaMEsxCzAJBgNVBAYTAlVTMRUwEwYDVQQKEwxWb3RhY2FsbCBJbmMxDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDcyNkowWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQTwi9CbiLYCi1E8iRPgTwuc7qoJ8TOOcWR8Zf%2FxbzzBMAKtLbk4r2ZmFZ2FuMpmDH150sZCMrYIt8qJfii9Q85o4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwHQYDVR0OBBYEFCw%2BFgESYGGNoc2v4QST4PS%2B3QH2MB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDcyNkowCgYIKoZIzj0EAwIDRwAwRAIgUNJW8vzzUrDTb7eVT%2FxSFWrj3AKd6GzM2QfYYci1brcCIBr7M7TgkO70fXrsRVJ9KFvIChVA6lbrw8300kLF5FTj)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 29 Dec 22 07:47 UTC