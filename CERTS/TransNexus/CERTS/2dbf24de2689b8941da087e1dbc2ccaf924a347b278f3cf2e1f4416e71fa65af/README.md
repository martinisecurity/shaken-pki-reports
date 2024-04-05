# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 606F

Tested At: 05 Apr 24 18:42 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -536 day(s)\
Subject: CN=SHAKEN 606F, OU=SHAKEN, O=Global Data Systems Inc., C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/756cb700-f9d2-4a05-850e-c9dfe3e22de4/6410808e3d363b24e911e9dd36c6ef8d.pem

[View certificate details](https://x509.io/?cert=MIIC%2BzCCAqCgAwIBAgIQWIVOh67onr3K1DgQFMQNszAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMDkyMTM5NTlaFw0yMjEwMTYyMTM5NThaMFcxCzAJBgNVBAYTAlVTMSEwHwYDVQQKExhHbG9iYWwgRGF0YSBTeXN0ZW1zIEluYy4xDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDYwNkYwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATbcgWK1g7Lpwegx2IKeFd8CEnvGnffJryk8%2FWkXYhnE%2FCGvMFx%2FbbgNG5NoatgQUPHaFWXkudYMyY0Z5WoFa4eo4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCAIAwHQYDVR0OBBYEFAIGsoRNUkr8TZhX9dbIl3%2FtMKmnMB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDYwNkYwCgYIKoZIzj0EAwIDSQAwRgIhAOw18Z3KkRhe54dtV49JbtdF7%2BKvQSovj%2FrN0OGv1YE0AiEAt0Xku27HtVw6mzfn0DICyOb8zseeXcO85%2FlXgCZbxsA%3D)

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


Generated: 05 Apr 24 19:04 UTC