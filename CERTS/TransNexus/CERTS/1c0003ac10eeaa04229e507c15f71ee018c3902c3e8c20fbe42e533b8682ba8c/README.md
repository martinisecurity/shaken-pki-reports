# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 505J

Tested At: 27 Nov 23 23:16 UTC\
Initial Validity Period: 90 day(s)\
Remaining Validity Period: -337 day(s)\
Subject: CN=SHAKEN 505J, OU=SHAKEN, O=HFA Services LLC dba Call48, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.transnexus.com/505J/ec7d1cb5-b714-4373-ad07-ef0e95a1e098.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC%2FjCCAqOgAwIBAgIQVP4AElRAE1jGc3mM8cfK0TAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjA5MjYxODQyMjRaFw0yMjEyMjUxODQyMjNaMFoxCzAJBgNVBAYTAlVTMSQwIgYDVQQKExtIRkEgU2VydmljZXMgTExDIGRiYSBDYWxsNDgxDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDUwNUowWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAT%2BimYfRZpQKygpE%2F9uNyexL1bopg5hDmewEnAs0Y2Res6KZWQyClNyRniQHAuPLxm6ivsM3t3CTqTPDVc0rwnQo4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCAIAwHQYDVR0OBBYEFOTdOlOigmkla3IyIAeE87BgDTWCMB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDUwNUowCgYIKoZIzj0EAwIDSQAwRgIhAML7I%2BVvSnMZZomE4xo7IzK3rSQrf0OpbJUTg4SsZrBYAiEAoJn0P8Ed1U14nULOZIigYBcbQuo33LIu6bqkYkBpYEw%3D)

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


Generated: 27 Nov 23 23:28 UTC