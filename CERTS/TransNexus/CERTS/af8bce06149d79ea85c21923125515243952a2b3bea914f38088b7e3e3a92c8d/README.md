# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 726J

Tested At: 26 Aug 24 17:44 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -677 day(s)\
Subject: CN=SHAKEN 726J, OU=SHAKEN, O=Votacall Inc, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/88dbb645-62bf-4d40-960c-4e8e3223880f/2161b4be0d61ff1a6e74f551cd1b8049.pem

[View certificate details](https://x509.io/?cert=MIIC7jCCApSgAwIBAgIQfG0dSMlKVqDxnqPprdSmATAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMTEyMDIwMjZaFw0yMjEwMTgyMDIwMjVaMEsxCzAJBgNVBAYTAlVTMRUwEwYDVQQKEwxWb3RhY2FsbCBJbmMxDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDcyNkowWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQW1tCrdsar4ujvsRdlrtgKplkbgLnP5Mn%2F771TO25dEjLvt9tius2z0ZNeqUGY%2FhDl8XEliaEnn%2BLpNdm4%2Be3Ko4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCAIAwHQYDVR0OBBYEFE%2FlJwk%2FCQxU8JF6wjtTfWzBK0g1MB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDcyNkowCgYIKoZIzj0EAwIDSAAwRQIgLpiJbzJ4CJ8WBEpmUljnFb9n5JmeYgD%2FYHK8V19PfaQCIQC6o2NBg5muFyVy4m7YOwiJmziUPdIxrrbE0T0hf5maqg%3D%3D)

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


Generated: 26 Aug 24 18:03 UTC