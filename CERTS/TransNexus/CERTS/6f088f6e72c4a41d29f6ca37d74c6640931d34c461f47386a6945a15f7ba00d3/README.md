# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 1821

Tested At: 12 Feb 24 18:54 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -505 day(s)\
Subject: CN=SHAKEN 1821, OU=SHAKEN, O=InPhonex.com\\, LLC., C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/4eed3a39-8f41-4acb-9b86-42a558b3455a/3985ea9b54d2309e13a5e8996d4c846a.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC9DCCApqgAwIBAgIQbSiK8z%2BjoFpqbRi1hotgUDAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjA5MTcyMTAzMzNaFw0yMjA5MjQyMTAzMzJaMFExCzAJBgNVBAYTAlVTMRswGQYDVQQKExJJblBob25leC5jb20sIExMQy4xDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDE4MjEwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAARJ8z0I5Z%2B%2FQaeLnt%2FqqQpRZF%2FFZZsEWkXERUVQvkVMeAmPQ9me84O7tTRg9FPXQ477A7R5HyENLdvda7O44dvNo4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCAIAwHQYDVR0OBBYEFAKrIM6XohXe%2Fek1oMSojW61dXgUMB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDE4MjEwCgYIKoZIzj0EAwIDSAAwRQIgZbGrYJtsBUgoTiTnd2X1iTMj11M1Nd3rYBf45D7dXuACIQCHJ84DZA9tBQDy0xiYYukhjSOb9ALg2RR7lTxHZ5BY9Q%3D%3D)

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


Generated: 12 Feb 24 19:26 UTC