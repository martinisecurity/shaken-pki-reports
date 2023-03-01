# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 793J

Tested At: 01 Mar 23 18:12 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -106 day(s)\
Subject: CN=SHAKEN 793J, OU=SHAKEN, O=UniTel Voice, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/45c216c2-fbb9-4f33-a3cf-6a64aad19357/e9b3255d75c9fe1c90b40cc1a9776011.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC7TCCApSgAwIBAgIQa0jMrFPOcXCvEA9zRYPdKjAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjExMDgxNTQ1MTlaFw0yMjExMTUxNTQ1MThaMEsxCzAJBgNVBAYTAlVTMRUwEwYDVQQKEwxVbmlUZWwgVm9pY2UxDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDc5M0owWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATmxT4kkSqJfxOxSIs7oxJFfd63bSf2SoKJ3qfeHMsDpBbtmmNh5NQqtpkDuhYvfoHLLzIbPA9UdK9heJvk6urxo4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwHQYDVR0OBBYEFD6nzDL5x2d%2Fj1yU9W1hNbl7osuSMB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDc5M0owCgYIKoZIzj0EAwIDRwAwRAIgcPfiAMEf9u69PSZ0yf74jeHBG2IN%2B2sA3QoPr%2BxwRDsCIHYUpDDIG2RLq6oo0DolMDFXjDAJ5Md8LdVMoPb5otiB)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 01 Mar 23 18:22 UTC