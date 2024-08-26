# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 849J

Tested At: 26 Aug 24 17:44 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -551 day(s)\
Subject: CN=SHAKEN 849J, O=Fuse.Cloud, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA4, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/86e241b8-9c8e-4431-b35f-4d92844a1da9/d70e45adbbf1dd9ff4666def08cdd98c.pem

[View certificate details](https://x509.io/?cert=MIICyTCCAnCgAwIBAgIQb9rhqBK2FyVOhNwOep7lZzAKBggqhkjOPQQDAjBWMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEsMCoGA1UEAxMjVHJhbnNOZXh1cywgSW5jLiBTSEFLRU4gSXNzdWluZyBDQTQwHhcNMjMwMjE0MjAzNDI1WhcNMjMwMjIxMjAzNDI0WjA4MQswCQYDVQQGEwJVUzETMBEGA1UEChMKRnVzZS5DbG91ZDEUMBIGA1UEAxMLU0hBS0VOIDg0OUowWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAR9VjddLTmPacIRXjiD3dY6uz87muYvZL%2BIqDgOqfqgm2OZ0AZBANnDJR0IW4%2B8b5IskPiqi7qf3ATboOam6EYCo4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwHQYDVR0OBBYEFIxwFjGT1xVI0u42ZmKWNiUcfSaCMB8GA1UdIwQYMBaAFDD19fK34UsLDxB1fUikkPE9iygqMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDg0OUowCgYIKoZIzj0EAwIDRwAwRAIgVjhac2RGgfKXQfd5UALc8mRwtAAhGutUKdpofqzgXXMCIEcQhtonuDUCuOTGFeO5ZNzkrXOoQcpt7RvTku2bywuz)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 26 Aug 24 18:03 UTC