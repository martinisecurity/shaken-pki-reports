# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 841J

Tested At: 26 Aug 24 17:45 UTC\
Initial Validity Period: 14 day(s)\
Remaining Validity Period: -679 day(s)\
Subject: CN=SHAKEN 841J, OU=SHAKEN, O=Securus Technologies LLC, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.transnexus.com/841J/439a753e-d8b0-488e-a771-bd146baafa6c.pem

[View certificate details](https://x509.io/?cert=MIIC%2BTCCAqCgAwIBAgIQY59qyNZ%2FuDLN9A06RbeshTAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMDMxMzU0MTdaFw0yMjEwMTcxMzU0MTZaMFcxCzAJBgNVBAYTAlVTMSEwHwYDVQQKExhTZWN1cnVzIFRlY2hub2xvZ2llcyBMTEMxDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDg0MUowWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATQ8yX6rTUEVLdHaxZx%2BOWIyM7CRAsT9lM38mfziVKCeUD34r7LaV3N%2F4p3vA%2FESJhasIO6QBmxxf7whP5tvp29o4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCAIAwHQYDVR0OBBYEFIlU9Pdm8hTX7X8YWIIuLYbEtmh%2BMB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDg0MUowCgYIKoZIzj0EAwIDRwAwRAIgKV4aqfNf%2FrbpYp8DV4YjRv5X3vxmWWs5S17p%2FNP1JkICIH33k6NrFIuCUndT8rcD6kwZBkZxQtRD1q81XqtCGqpx)

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


Generated: 26 Aug 24 18:03 UTC