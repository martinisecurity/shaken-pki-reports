# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 738J

Tested At: 26 Aug 24 18:11 UTC\
Initial Validity Period: 9 day(s)\
Remaining Validity Period: -683 day(s)\
Subject: CN=SHAKEN 738J, OU=SHAKEN, O=SkySwitch, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.transnexus.com/738J/1d80ee50-712d-4ae3-89f2-d794dff47108.pem

[View certificate details](https://x509.io/?cert=MIIC6zCCApGgAwIBAgIQQkj50y%2FhctwASBEGGPgtSTAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMDMyMDE5MDVaFw0yMjEwMTIyMDE5MDRaMEgxCzAJBgNVBAYTAlVTMRIwEAYDVQQKEwlTa3lTd2l0Y2gxDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDczOEowWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAARK7aFIr2iW516KfUAXR14FYvqeEWMC9mSlY0xg7iqe1cNTZ2TxVB3Scr9jyIUdB6b8Kp3CNLH%2F9wnLMAwgpf6Po4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCAIAwHQYDVR0OBBYEFDztVqIkK7bXurMRc2XA3ohsRYs8MB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDczOEowCgYIKoZIzj0EAwIDSAAwRQIgBMP5D%2B21uyy%2BBYgdexLh%2BDRhhMt2dvJ2ue6SAeYsl7ECIQCxLg8x8%2Fivdvpr4K8Q5GV4doPjxgdosQcfzUYTowxgZg%3D%3D)

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


Generated: 26 Aug 24 18:49 UTC