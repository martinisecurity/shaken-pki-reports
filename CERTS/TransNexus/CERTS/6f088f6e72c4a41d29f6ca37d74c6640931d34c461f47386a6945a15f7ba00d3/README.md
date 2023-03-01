# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 1821

Tested At: 01 Mar 23 18:12 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -157 day(s)\
Subject: CN=SHAKEN 1821, OU=SHAKEN, O=InPhonex.com\\, LLC., C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/4eed3a39-8f41-4acb-9b86-42a558b3455a/3985ea9b54d2309e13a5e8996d4c846a.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC9DCCApqgAwIBAgIQbSiK8z%2BjoFpqbRi1hotgUDAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjA5MTcyMTAzMzNaFw0yMjA5MjQyMTAzMzJaMFExCzAJBgNVBAYTAlVTMRswGQYDVQQKExJJblBob25leC5jb20sIExMQy4xDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDE4MjEwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAARJ8z0I5Z%2B%2FQaeLnt%2FqqQpRZF%2FFZZsEWkXERUVQvkVMeAmPQ9me84O7tTRg9FPXQ477A7R5HyENLdvda7O44dvNo4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCAIAwHQYDVR0OBBYEFAKrIM6XohXe%2Fek1oMSojW61dXgUMB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDE4MjEwCgYIKoZIzj0EAwIDSAAwRQIgZbGrYJtsBUgoTiTnd2X1iTMj11M1Nd3rYBf45D7dXuACIQCHJ84DZA9tBQDy0xiYYukhjSOb9ALg2RR7lTxHZ5BY9Q%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 01 Mar 23 18:22 UTC