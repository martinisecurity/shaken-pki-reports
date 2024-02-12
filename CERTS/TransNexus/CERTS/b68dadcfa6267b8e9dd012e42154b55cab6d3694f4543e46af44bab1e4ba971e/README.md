# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 8526

Tested At: 12 Feb 24 18:56 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -444 day(s)\
Subject: CN=SHAKEN 8526, OU=SHAKEN, O=MetTel, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.transnexus.com/8526/f6978975-50bc-40df-b4a2-a1d7794b321f.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC5zCCAo6gAwIBAgIQW%2Ba5TEL9yKXSEDF6IWId9DAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMjYxMDQzMjFaFw0yMjExMjUxMDQzMjBaMEUxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKEwZNZXRUZWwxDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDg1MjYwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAARdu84I1Kt53t2DZkiw5Ms%2BMvqNcatkYJ0tUnqXV5O7BBkt2INbtNbTpHqoZQm7O6Rw8G6SPo4NUj6U3rPUmuNto4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCAIAwHQYDVR0OBBYEFFTD%2FOLSAW6XBysz47V8Sm8XUxFUMB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDg1MjYwCgYIKoZIzj0EAwIDRwAwRAIgRfQG8OySuAsY2kVYuZ1CPxF7c2ckaMWDMbyTEZVhjYcCIFwU7UIwtPfNZlPCQC7MF4XPa83%2FoMBk2zu5nu0myO5H)

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