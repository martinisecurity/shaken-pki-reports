# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 807J

Tested At: 04 Oct 24 15:50 UTC\
Initial Validity Period: 60 day(s)\
Remaining Validity Period: -760 day(s)\
Subject: CN=SHAKEN 807J, OU=SHAKEN, O=SipPhony LLC, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.transnexus.com/807J/99ee724e-0393-41fa-85a3-cb5dfc0ee4de.pem

[View certificate details](https://x509.io/?cert=MIIC7jCCApSgAwIBAgIQdfXWkjFF2oGnSGHL%2FvZf%2BzAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjA3MDYyMDQyMjJaFw0yMjA5MDQyMDQyMjFaMEsxCzAJBgNVBAYTAlVTMRUwEwYDVQQKEwxTaXBQaG9ueSBMTEMxDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDgwN0owWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQpxp%2FJ97hBEuFcZXkR3y3CrJ0Sv2xvUqU24ce65TwH01SeoJ3xobyUC%2FC%2FjGAbWn5hj6Cz%2BJSDaPqyWo%2B59I3ho4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCAIAwHQYDVR0OBBYEFAQp2s%2BAaDLJaHbvzMnrTCyqg0Z1MB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDgwN0owCgYIKoZIzj0EAwIDSAAwRQIgek05rAga8iNXquTG%2BcV5bNynguCHi%2FZhqUvVeocdoOYCIQCLsdJKkhBvkDphFBkF1d5AUKq8XQ7W%2F4YxCVWJFgpUpg%3D%3D)

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