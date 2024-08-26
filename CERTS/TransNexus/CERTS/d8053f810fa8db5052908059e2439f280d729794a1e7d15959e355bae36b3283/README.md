# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 672B

Tested At: 26 Aug 24 18:09 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -557 day(s)\
Subject: CN=SHAKEN 672B, O=Clear Rate, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA4, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/dcef9a97-b864-43ac-830b-2290a8d0f002/afb7b5e0263ba81fa287230c315cc4dc.pem

[View certificate details](https://x509.io/?cert=MIICyTCCAnCgAwIBAgIQWK9kLF%2BMBp%2FV09qwiSC1rzAKBggqhkjOPQQDAjBWMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEsMCoGA1UEAxMjVHJhbnNOZXh1cywgSW5jLiBTSEFLRU4gSXNzdWluZyBDQTQwHhcNMjMwMjA4MjAzNTE2WhcNMjMwMjE1MjAzNTE1WjA4MQswCQYDVQQGEwJVUzETMBEGA1UEChMKQ2xlYXIgUmF0ZTEUMBIGA1UEAxMLU0hBS0VOIDY3MkIwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATUTo8E6Q2GAB2FQ2jcOKWOdC0wU8yFf29Fin0Sy6NsTfUDUyWcTfEW5GjSaFx54HUI6EYRRmyUDrtrutn64fJFo4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwHQYDVR0OBBYEFLChqAaS7b0UY9%2FgC5mwhONFaEgfMB8GA1UdIwQYMBaAFDD19fK34UsLDxB1fUikkPE9iygqMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDY3MkIwCgYIKoZIzj0EAwIDRwAwRAIgScgRXmR7bXsvEf64KjZESB11m4PRgrmbdHoLhdRh2QECIFF62e3ju7osuGz%2Fv7g5GIDMcBQRcqGe6IDmM1Mi3G5b)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 26 Aug 24 18:49 UTC