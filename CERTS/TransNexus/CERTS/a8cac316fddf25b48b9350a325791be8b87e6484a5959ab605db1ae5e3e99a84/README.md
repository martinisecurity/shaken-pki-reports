# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 755J

Tested At: 01 Nov 22 20:29 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 233 day(s)\
Subject: CN=SHAKEN 755J, OU=SHAKEN, O=CMS Internet, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.transnexus.com/755J/5e526fc6-7b0a-43a5-bb13-263d7edc74e3.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIC7jCCApSgAwIBAgIQZ%2F2MUM%2FuIeN8yqu71YlNMjAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjA2MjIxODQ2NTFaFw0yMzA2MjIxODQ2NTBaMEsxCzAJBgNVBAYTAlVTMRUwEwYDVQQKEwxDTVMgSW50ZXJuZXQxDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDc1NUowWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQKRMmzaG%2BKEECcLHzI%2F%2B18q2iYphjTou8b5yJcJ0SgpkFZLB1GAbRhKZ9PC2tEp%2BaGKq3n0eiaA%2B9s%2FN5RLbtRo4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCAIAwHQYDVR0OBBYEFBZeJ%2BWcrYXRFORCTASxE5%2FNKJrYMB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDc1NUowCgYIKoZIzj0EAwIDSAAwRQIhAOqWdFrBLMBsbZfIiauExtp8HkO%2BjGpuifJ1fQBNkS5cAiAYyAg4pl0%2BHLK2OtwClxWP8UdHtUI2lI19PHJmLYJCRg%3D%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |


Generated: 01/11/2022 at 20:31:14