# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 2550

Tested At: 22 Aug 24 15:20 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -654 day(s)\
Subject: CN=SHAKEN 2550, OU=SHAKEN, O=123.Net\\, Inc, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/2e3276e5-a5c5-43c7-ae11-3a226bf1b5ed/5cc23b30bf323860b4a4f48bccfd670f.pem

[View certificate details](https://x509.io/?cert=MIIC7zCCApSgAwIBAgIQUDODU3Jk7pKn5OyryXcobjAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMzExNDMzNTBaFw0yMjExMDcxNDMzNDlaMEsxCzAJBgNVBAYTAlVTMRUwEwYDVQQKEwwxMjMuTmV0LCBJbmMxDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDI1NTAwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAARXyW2YPp%2BaDdY4Nwn7DoJu3t83ww9xySHCC7NFpl3HC4ltmN%2Bwsl1sOCTUgjYchjW%2FK9Y04s5XDiRjlW1AS%2BRJo4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwHQYDVR0OBBYEFIljAYn1PqS%2FQMRGOZR9ykgAHwfiMB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDI1NTAwCgYIKoZIzj0EAwIDSQAwRgIhAP4M7xXSb1hZ1oMmYMV0%2BoNiSOzvBmpgdIJncU3mwsm6AiEAjc4iIWomPkMk7rCghYD89uxuG4u36i0X3Qwc08aFFTY%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 22 Aug 24 16:06 UTC