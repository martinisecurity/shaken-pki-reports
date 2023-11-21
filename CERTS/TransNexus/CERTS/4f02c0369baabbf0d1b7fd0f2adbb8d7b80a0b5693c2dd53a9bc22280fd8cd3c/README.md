# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 722J

Tested At: 21 Nov 23 16:46 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -386 day(s)\
Subject: CN=SHAKEN 722J, OU=SHAKEN, O=EvolveIP LLC, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.transnexus.com/722J/a547c920-fe4c-4086-a4cf-a744f0a1a4ce.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC7jCCApSgAwIBAgIQUmuUufynvg9JFyDpuvleRTAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMDExNTIzMjhaFw0yMjEwMzExNTIzMjdaMEsxCzAJBgNVBAYTAlVTMRUwEwYDVQQKEwxFdm9sdmVJUCBMTEMxDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDcyMkowWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATXvcrxfdNgF8U6D9i1dI9aAYkV0jgjUS0bVyeZslC55nU%2BSSH4zQqXb28uqv7mrkyFzVr57fO%2FIH1gkozCkE9Co4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCAIAwHQYDVR0OBBYEFIRxVEuov2K4AkFbiPiIyVolPHe2MB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDcyMkowCgYIKoZIzj0EAwIDSAAwRQIgFLZwgo%2FYdr9BVnPgbppYnQi9uj8oSZ6202kCILbrBF0CIQDXUlrPK%2Fvh00BrRsdo5nTARHVHC8RpEI1%2FtBceNBJ36w%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 21 Nov 23 17:16 UTC