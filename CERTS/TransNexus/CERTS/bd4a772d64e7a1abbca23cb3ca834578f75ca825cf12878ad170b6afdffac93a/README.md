# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 2550

Tested At: 04 Oct 24 15:34 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -709 day(s)\
Subject: CN=SHAKEN 2550, OU=SHAKEN, O=123.Net\\, Inc, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/2e3276e5-a5c5-43c7-ae11-3a226bf1b5ed/2fa92ca193aef4dd820a6def4f41d586.pem

[View certificate details](https://x509.io/?cert=MIIC7jCCApSgAwIBAgIQQuwOTw370bcteDcLQgQ19jAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMTkxNDMzMjZaFw0yMjEwMjYxNDMzMjVaMEsxCzAJBgNVBAYTAlVTMRUwEwYDVQQKEwwxMjMuTmV0LCBJbmMxDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDI1NTAwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAARCWkm7SxnDS8rO0b%2BJvpFX9sx%2BElA5H0YVmAQwvpn%2FIuedXKQGhP6ocILKmYMOEgnvOkMOxpAjHEt1x%2Bk8wTVho4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCAIAwHQYDVR0OBBYEFAEBAhcYuMudhbXxI47imBhFCxcGMB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDI1NTAwCgYIKoZIzj0EAwIDSAAwRQIhAPX%2By4HriuXjv8mcrp27u%2FlXMDNeKsHfL2J%2F2xV0%2FyUUAiBvjV8PvQoDNHhilzOh4QWhTXyZE07BBkZvIMsyxfQDSg%3D%3D)

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