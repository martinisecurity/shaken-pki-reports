# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 815G

Tested At: 02 Dec 22 07:44 UTC\
Initial Validity Period: 90 day(s)\
Remaining Validity Period: 40 day(s)\
Subject: CN=SHAKEN 815G, OU=SHAKEN, O=382 Communications, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.transnexus.com/815G/c147c307-645b-4fcd-99d0-2be9a709b265.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC8zCCApqgAwIBAgIQVBzkBS9QkBvFLNA%2BHAwGIDAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMTIxOTM5MjNaFw0yMzAxMTAxOTM5MjJaMFExCzAJBgNVBAYTAlVTMRswGQYDVQQKExIzODIgQ29tbXVuaWNhdGlvbnMxDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDgxNUcwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQ6ZwnWT%2FW9Es%2Fw2sdBjPxX%2Bx2eeIWPmXdnwVDVbBf0nCpSXEuNEyLxQvqSecGlgrQBHEbiE2YfTuh8ME85ZtAso4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCAIAwHQYDVR0OBBYEFFcubvF9SYzWT9Emc5wWGsEWm5tLMB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDgxNUcwCgYIKoZIzj0EAwIDRwAwRAIgQJm%2FCsRXL2ANGGz%2BGYrB%2FzOKNsbd%2B8peo%2FCyFL1eXYwCIAsEBqvjJWLRwaFtJcYpR5%2FM6cPInOXVJzeiGBCE8prA)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 02 Dec 22 07:46 UTC