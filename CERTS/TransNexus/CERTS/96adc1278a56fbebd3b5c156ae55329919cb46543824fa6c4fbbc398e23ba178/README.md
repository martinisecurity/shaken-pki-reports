# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 738J

Tested At: 04 Oct 24 15:49 UTC\
Initial Validity Period: 9 day(s)\
Remaining Validity Period: -736 day(s)\
Subject: CN=SHAKEN 738J, OU=SHAKEN, O=SkySwitch, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.transnexus.com/738J/3f3bd1bb-3dc4-47d2-be79-0b2adc6a7fac.pem

[View certificate details](https://x509.io/?cert=MIIC7DCCApGgAwIBAgIQUxBOHuTb%2BIXejXxbC0frgjAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjA5MTkyMDE3MDRaFw0yMjA5MjgyMDE3MDNaMEgxCzAJBgNVBAYTAlVTMRIwEAYDVQQKEwlTa3lTd2l0Y2gxDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDczOEowWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQ8UybgaGEx4mb8CpiLn04%2Bpl3zMi820xXpkTxFdWw6%2Bh%2BPK619pEAVoIrZj9zMz9lXkfJrzCoy0SFInK7S17KUo4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCAIAwHQYDVR0OBBYEFPvJefvboIk12cdMSmuVdvbl2%2BYSMB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDczOEowCgYIKoZIzj0EAwIDSQAwRgIhALwDLQYX54%2FyUXERaoch%2B18F1KcFTgB5vjkgPcqkZ7RVAiEA124K1PZnEw0nz4crRRdDPAK8pInBEIvBU2yRh67l1Yw%3D)

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


Generated: 04 Oct 24 16:29 UTC