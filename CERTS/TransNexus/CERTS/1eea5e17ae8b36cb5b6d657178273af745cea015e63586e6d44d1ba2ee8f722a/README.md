# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 706J

Tested At: 12 Feb 24 18:55 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: -152 day(s)\
Subject: CN=SHAKEN 706J, OU=SHAKEN, O=TCN, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.transnexus.com/706J/fdb13f3d-0f05-47c8-a047-af2f9e2f7af5.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC5TCCAougAwIBAgIQX4iz5rOuXyo%2BWVXh6OUhkTAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjA5MTMxNzI0NDlaFw0yMzA5MTMxNzI0NDhaMEIxCzAJBgNVBAYTAlVTMQwwCgYDVQQKEwNUQ04xDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDcwNkowWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQx%2BKckfjDFK87CpTfCrmwX%2F4zrIt2dW4thkEjU9iESG2AusWrZyEkqtuZp3PpEnOu2ltogB49m134%2F3cLhMDUAo4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCAIAwHQYDVR0OBBYEFKewD5Maow5dvMl2kNna%2BBaJqvnoMB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDcwNkowCgYIKoZIzj0EAwIDSAAwRQIgC8%2F3oFUZ%2FygRr52l5y4AbJahiJ9r5o0290iHCDaLOYgCIQDT9Uz4XQIiMlQAZ97s14HaFdKRiHZ9epw2DtYH1yRkCw%3D%3D)

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