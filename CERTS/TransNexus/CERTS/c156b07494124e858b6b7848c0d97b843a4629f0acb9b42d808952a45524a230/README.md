# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 815G

Tested At: 26 Aug 24 17:45 UTC\
Initial Validity Period: 90 day(s)\
Remaining Validity Period: -685 day(s)\
Subject: CN=SHAKEN 815G, OU=SHAKEN, O=382 Communications, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.transnexus.com/815G/dc14f4fd-d8a5-4499-a672-069fd27430eb.pem

[View certificate details](https://x509.io/?cert=MIIC8zCCApqgAwIBAgIQW%2F1zcEVpq%2Fi00iA6PhZxQDAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjA3MTMxNTE2MTlaFw0yMjEwMTExNTE2MThaMFExCzAJBgNVBAYTAlVTMRswGQYDVQQKExIzODIgQ29tbXVuaWNhdGlvbnMxDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDgxNUcwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASxvh2wi3PK4yds%2BHCIpLk56UDC3sAy1Z63sEUqOSyXHI6%2BqAzYnZojp8wtHTB0Jc5GCUnbSkQfQzwyNX8FYH1Wo4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCAIAwHQYDVR0OBBYEFN57gSQgP7hzeuel7wJ47ga2ukQoMB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDgxNUcwCgYIKoZIzj0EAwIDRwAwRAIfOHDz8w0oQ3h7p5O6fnDjkdpPOsndDfp5d%2Ftr1CNXTQIhANbezLR7cT%2BcZDqPH5C1lPT6bl0gyp2hl22INfthoGlt)

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