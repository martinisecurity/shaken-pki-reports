# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 983J

Tested At: 21 Nov 23 18:48 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -412 day(s)\
Subject: CN=SHAKEN 983J, OU=SHAKEN, O=ESI, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/59ebb7c1-25bd-4dfc-9794-fcb104b2f66a/1c12e6720e37cd8b0160dc8d126bdd22.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC5TCCAougAwIBAgIQXwraWY%2Bdyq5q4k85HwJ5XjAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjA5MjcyMDEwMzhaFw0yMjEwMDQyMDEwMzdaMEIxCzAJBgNVBAYTAlVTMQwwCgYDVQQKEwNFU0kxDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDk4M0owWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATjJStdqj%2BKp6iDvQA6kLp0cCBBD01KEGDIk9iBZT4pDRZKDUBWhFNJXG%2BJkDm%2FsELCim%2FRPxNpJgQtSpmCQ0Geo4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCAIAwHQYDVR0OBBYEFDZ1M8oI6ZVH46h9YBuh7VJ3g96SMB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDk4M0owCgYIKoZIzj0EAwIDSAAwRQIgd62PqatBnuj%2FX5IbVntJvjID6doOXXnmnkHBBGHJJJsCIQCX%2BjH6V4qF9hihiSrWMNnKjEqxO9%2B8Sm6dKbo7iPOy3Q%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 21 Nov 23 19:18 UTC