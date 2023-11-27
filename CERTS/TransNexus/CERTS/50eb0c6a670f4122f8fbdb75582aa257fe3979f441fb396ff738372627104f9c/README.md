# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 722J

Tested At: 27 Nov 23 23:16 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -368 day(s)\
Subject: CN=SHAKEN 722J, OU=SHAKEN, O=EvolveIP LLC, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.transnexus.com/722J/f75ff50a-1fd7-4087-a0a1-de57acb1afc5.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC7jCCApSgAwIBAgIQc%2BKyCHzJSNGtnXxVDfceGjAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMjUxNTMzMDhaFw0yMjExMjQxNTMzMDdaMEsxCzAJBgNVBAYTAlVTMRUwEwYDVQQKEwxFdm9sdmVJUCBMTEMxDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDcyMkowWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQ8f0w80%2F9ZFljqzq0go21KcrjLZ7qmhEyg45SEWPAG%2F8Fo2Jn%2FYAenoRoCzw7IuqU8P6QUG5Fq4YXyjxMk5otUo4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCAIAwHQYDVR0OBBYEFAI7mT3J%2FhAQ7DXaxwt9jbd%2BiZnqMB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDcyMkowCgYIKoZIzj0EAwIDSAAwRQIhANIfcvfuWaXSLlHKydwmC%2Be%2Bq5R1X2vvW5S2B4GH8wCGAiBMsL5Lxw4VDpM7oxcdjw8%2BO%2Fu8lHlpgBF8BEIl3%2Bd0DA%3D%3D)

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


Generated: 27 Nov 23 23:28 UTC