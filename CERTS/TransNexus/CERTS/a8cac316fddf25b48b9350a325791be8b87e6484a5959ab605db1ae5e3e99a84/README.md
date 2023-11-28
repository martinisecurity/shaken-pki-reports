# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 755J

Tested At: 28 Nov 23 10:27 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: -158 day(s)\
Subject: CN=SHAKEN 755J, OU=SHAKEN, O=CMS Internet, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.transnexus.com/755J/5e526fc6-7b0a-43a5-bb13-263d7edc74e3.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC7jCCApSgAwIBAgIQZ%2F2MUM%2FuIeN8yqu71YlNMjAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjA2MjIxODQ2NTFaFw0yMzA2MjIxODQ2NTBaMEsxCzAJBgNVBAYTAlVTMRUwEwYDVQQKEwxDTVMgSW50ZXJuZXQxDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDc1NUowWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQKRMmzaG%2BKEECcLHzI%2F%2B18q2iYphjTou8b5yJcJ0SgpkFZLB1GAbRhKZ9PC2tEp%2BaGKq3n0eiaA%2B9s%2FN5RLbtRo4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCAIAwHQYDVR0OBBYEFBZeJ%2BWcrYXRFORCTASxE5%2FNKJrYMB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDc1NUowCgYIKoZIzj0EAwIDSAAwRQIhAOqWdFrBLMBsbZfIiauExtp8HkO%2BjGpuifJ1fQBNkS5cAiAYyAg4pl0%2BHLK2OtwClxWP8UdHtUI2lI19PHJmLYJCRg%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 28 Nov 23 10:53 UTC