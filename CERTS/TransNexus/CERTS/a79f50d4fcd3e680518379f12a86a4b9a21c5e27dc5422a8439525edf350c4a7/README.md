# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 815G

Tested At: 21 Nov 23 18:48 UTC\
Initial Validity Period: 90 day(s)\
Remaining Validity Period: -314 day(s)\
Subject: CN=SHAKEN 815G, OU=SHAKEN, O=382 Communications, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.transnexus.com/815G/c147c307-645b-4fcd-99d0-2be9a709b265.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC8zCCApqgAwIBAgIQVBzkBS9QkBvFLNA%2BHAwGIDAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMTIxOTM5MjNaFw0yMzAxMTAxOTM5MjJaMFExCzAJBgNVBAYTAlVTMRswGQYDVQQKExIzODIgQ29tbXVuaWNhdGlvbnMxDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDgxNUcwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQ6ZwnWT%2FW9Es%2Fw2sdBjPxX%2Bx2eeIWPmXdnwVDVbBf0nCpSXEuNEyLxQvqSecGlgrQBHEbiE2YfTuh8ME85ZtAso4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCAIAwHQYDVR0OBBYEFFcubvF9SYzWT9Emc5wWGsEWm5tLMB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDgxNUcwCgYIKoZIzj0EAwIDRwAwRAIgQJm%2FCsRXL2ANGGz%2BGYrB%2FzOKNsbd%2B8peo%2FCyFL1eXYwCIAsEBqvjJWLRwaFtJcYpR5%2FM6cPInOXVJzeiGBCE8prA)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 21 Nov 23 19:18 UTC