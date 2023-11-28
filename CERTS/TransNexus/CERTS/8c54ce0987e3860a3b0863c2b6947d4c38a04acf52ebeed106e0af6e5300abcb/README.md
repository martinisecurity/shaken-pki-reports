# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 841J

Tested At: 28 Nov 23 10:27 UTC\
Initial Validity Period: 14 day(s)\
Remaining Validity Period: -390 day(s)\
Subject: CN=SHAKEN 841J, OU=SHAKEN, O=Securus Technologies LLC, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.transnexus.com/841J/f3c3ccf0-0018-4b22-bf77-8a30b9ffd68e.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC%2BjCCAqCgAwIBAgIQeDJWfyVNF0mdJI%2F9aElKmzAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMjAwOTEzMThaFw0yMjExMDMwOTEzMTdaMFcxCzAJBgNVBAYTAlVTMSEwHwYDVQQKExhTZWN1cnVzIFRlY2hub2xvZ2llcyBMTEMxDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDg0MUowWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAARNo91X64NUmlHu7XGS8%2FBYC4OpGG6y7B44%2Bkme%2B3jhwCuKrURNfZM8aKLiHJPKzDaG5jPWQz1n2Q%2FFiJfFWKJyo4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCAIAwHQYDVR0OBBYEFC8u42jFW4tuWzeTHvglkF6BXKdZMB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDg0MUowCgYIKoZIzj0EAwIDSAAwRQIhAJvn2PIRofIR%2BU%2Bi7kgGpVCeU3KK86KsdGQbMNGy8sf7AiBfW1hvccc0Ivr%2FUxtBxfXJezRe9E08CSAx07c%2FFR6pdA%3D%3D)

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