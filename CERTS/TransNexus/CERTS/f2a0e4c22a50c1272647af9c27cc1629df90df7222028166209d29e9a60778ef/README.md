# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 111K

Tested At: 04 Oct 24 15:43 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -727 day(s)\
Subject: CN=SHAKEN 111K, OU=SHAKEN, O=Cloud Voice Tel, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/ae8a8a64-668f-4934-823d-42ebc587f82a/af862ae5ed3fbfa249b1c23dfb7b7ec6.pem

[View certificate details](https://x509.io/?cert=MIIC8DCCApegAwIBAgIQX32alcYKJpfQUcyBb8wqpjAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMDExNTQzNTBaFw0yMjEwMDgxNTQzNDlaME4xCzAJBgNVBAYTAlVTMRgwFgYDVQQKEw9DbG91ZCBWb2ljZSBUZWwxDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDExMUswWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAARBm7C3%2BV6jMsruP4m2t%2BxncAjJUQpDnEqxz%2Bj6a2aRNZXue%2BKkNBWt6lkKihJIRmd1SXIaA%2F9dbRz1LGfBqw2No4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCAIAwHQYDVR0OBBYEFMG1vxAAJsIsabc1VtvoOx2wZq1wMB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDExMUswCgYIKoZIzj0EAwIDRwAwRAIgDy6mUNyOYNMEJhb%2F4ZIhRAlTHvXpYN0FTe0OLfyrgdYCIEcFnyL8x9Dhnhcc7B9KaR7YQEYsLl%2BLfLsXG%2Fe%2FTSB7)

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


Generated: 04 Oct 24 16:29 UTC