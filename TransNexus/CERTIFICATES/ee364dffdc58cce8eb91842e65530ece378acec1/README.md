# STIR/SHAKEN CA Ecosystem Compliance
## TransNexus

### Certificate ee364dffdc58cce8eb91842e65530ece378acec1
Tested At: 2022-10-28 16:28:15 +0000 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 238 day(s)\
Subject: CN=SHAKEN 755J, OU=SHAKEN, O=CMS Internet, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US

Link: https://certificates.transnexus.com/755J/5e526fc6-7b0a-43a5-bb13-263d7edc74e3.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIC7jCCApSgAwIBAgIQZ%2F2MUM%2FuIeN8yqu71YlNMjAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjA2MjIxODQ2NTFaFw0yMzA2MjIxODQ2NTBaMEsxCzAJBgNVBAYTAlVTMRUwEwYDVQQKEwxDTVMgSW50ZXJuZXQxDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDc1NUowWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQKRMmzaG%2BKEECcLHzI%2F%2B18q2iYphjTou8b5yJcJ0SgpkFZLB1GAbRhKZ9PC2tEp%2BaGKq3n0eiaA%2B9s%2FN5RLbtRo4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCAIAwHQYDVR0OBBYEFBZeJ%2BWcrYXRFORCTASxE5%2FNKJrYMB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDc1NUowCgYIKoZIzj0EAwIDSAAwRQIhAOqWdFrBLMBsbZfIiauExtp8HkO%2BjGpuifJ1fQBNkS5cAiAYyAg4pl0%2BHLK2OtwClxWP8UdHtUI2lI19PHJmLYJCRg%3D%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_incorrect_ku_encoding | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| e_cp1_3_subject_sn | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| w_pki_subject_rdn_unknown | warn | SHAKEN PKI Best Practice | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |

Generated: 28/10/2022 at 16:28:22