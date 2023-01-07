# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 722J

Tested At: 07 Jan 23 19:08 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -92 day(s)\
Subject: CN=SHAKEN 722J, OU=SHAKEN, O=EvolveIP LLC, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.transnexus.com/722J/2902cf3c-1e73-4c05-ab0e-e702541b2ccc.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC7jCCApSgAwIBAgIQS8OslDKIzyNTejZm%2BvWGQjAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjA5MDcxNTIyNDBaFw0yMjEwMDcxNTIyMzlaMEsxCzAJBgNVBAYTAlVTMRUwEwYDVQQKEwxFdm9sdmVJUCBMTEMxDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDcyMkowWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATYK9DkIX7675fgGmD%2FFiSBAe3duw3tk3JjkJlTNMqLQjZJJrF7tnBM2BYxS50iraidPOtrbhmvPmNX7yhEt0neo4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCAIAwHQYDVR0OBBYEFMUwvwc9xGV%2FKAVVUc8efqgipEgvMB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDcyMkowCgYIKoZIzj0EAwIDSAAwRQIgIKk5SyYaVBbDySG4OWQl%2BJlwTthFBiI6wKXwuJ8gSMsCIQDlk0VNSDtftmiJ0s5mj9zyuD6cckA58vFPDXgR0xbbww%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |


Generated: 07 Jan 23 19:18 UTC