# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 983J

Tested At: 31 Jan 23 21:39 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -74 day(s)\
Subject: CN=SHAKEN 983J, OU=SHAKEN, O=ESI, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/59ebb7c1-25bd-4dfc-9794-fcb104b2f66a/dafc6d95003d711a029d6a69af6b67bd.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC5TCCAougAwIBAgIQdBW1raMiIibAECYryDYuSzAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjExMTEyMDE1MTRaFw0yMjExMTgyMDE1MTNaMEIxCzAJBgNVBAYTAlVTMQwwCgYDVQQKEwNFU0kxDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDk4M0owWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQ%2BV%2B99j5XMSdBaiVZfd8%2Fv0DOw25UIL5XN9ScQnY6iZ4AKJENRQSoafWoMqtsBYbOGmYK%2Fj%2F77mwTQGWMd038co4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwHQYDVR0OBBYEFMdE9ygAM54Ry6D%2Fpp3yI6Ta4JQlMB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDk4M0owCgYIKoZIzj0EAwIDSAAwRQIgD8E1jwhjQjccNlXce3inyEAeDdXNTXnKE3rYVGBpLeACIQD41Wjs%2FYwN3rIhKHKTBFaDpGPTS7U6%2FCm951%2BONKpeiQ%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 31 Jan 23 21:50 UTC