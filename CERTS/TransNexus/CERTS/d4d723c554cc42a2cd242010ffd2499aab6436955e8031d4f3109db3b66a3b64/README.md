# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 166K

Tested At: 26 Aug 24 18:11 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -656 day(s)\
Subject: CN=SHAKEN 166K, OU=SHAKEN, O=Gonthier Inc, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.transnexus.com/166K/b4f563ea-1dd9-40bc-be12-ef0190b82bad.pem

[View certificate details](https://x509.io/?cert=MIIC7jCCApSgAwIBAgIQQA5yBjrcSUOikgjvm5tkOjAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMTAxODI3MThaFw0yMjExMDkxODI3MTdaMEsxCzAJBgNVBAYTAlVTMRUwEwYDVQQKEwxHb250aGllciBJbmMxDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDE2NkswWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAARDofGwS0vLXTX3s6ksrD7fKjcpM%2BI8MdJIEXbbWyqTJZKgIrUUd9UsuLxvPVL3zWyw8jAyMGPgi8VrZg9DhZfHo4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCAIAwHQYDVR0OBBYEFF8%2FJVlK1vjNhLBaLkXh6iJ9zecxMB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDE2NkswCgYIKoZIzj0EAwIDSAAwRQIgQo66q%2FMgdaOtqMRuf4I1rX3uRKgVuI0RQEfalCBUi6YCIQDDuvY87XFIwuywC2Iin7CUDnPvBFbuwYFZJ6b6ml%2BlxA%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 26 Aug 24 18:49 UTC