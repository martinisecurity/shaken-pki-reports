# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 7453

Tested At: 28 Nov 23 16:00 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -402 day(s)\
Subject: CN=SHAKEN 7453, OU=SHAKEN, O=TPx, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/2a26bd25-8973-44ba-9b94-a8e4716b3888/eac658a0c2f06f600cf8e58e1359a05a.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC5TCCAougAwIBAgIQfCa2WoEm45WkA0rJq0IqWjAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMTQyMDE4MDZaFw0yMjEwMjEyMDE4MDVaMEIxCzAJBgNVBAYTAlVTMQwwCgYDVQQKEwNUUHgxDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDc0NTMwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASs%2FbvyI0Apo%2F0zJMxxG3TH2RZrQCtw20VjzqGy5DVtPKhUmhrLyiKjdkYAESfyeDTP5bFFTQUW4L6CLxGf5Bzco4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCAIAwHQYDVR0OBBYEFPmJ63oEnWHojnb38bQ36VLEkXN3MB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDc0NTMwCgYIKoZIzj0EAwIDSAAwRQIgITK1eMXc4h0qp0MUP7W%2BgtF0FWVFhLij7IoN%2FqU2MzkCIQCNyaH3xU7eY51rSJ%2F39vmPTCElMiDEMNVn24%2B%2B7dRxMg%3D%3D)

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


Generated: 28 Nov 23 16:15 UTC