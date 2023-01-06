# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 0172

Tested At: 06 Jan 23 02:53 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -71 day(s)\
Subject: CN=SHAKEN 0172, OU=SHAKEN, O=CallTower, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.transnexus.com/0172/d1be3f78-bfc4-426b-8bd8-28bfd3a114db.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC6zCCApGgAwIBAgIQZ38der00d1ciQgnk%2FHurITAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMTkxNzIyMjRaFw0yMjEwMjYxNzIyMjNaMEgxCzAJBgNVBAYTAlVTMRIwEAYDVQQKEwlDYWxsVG93ZXIxDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDAxNzIwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAR%2FwZoeUMY6FmXe8R5URZ9W0TpH6TMjUMHI%2BMANofUTfU2gxZrlft9QTbGUFKWr8Pk6CisBcbUThOsW%2Fr9SCJx7o4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCAIAwHQYDVR0OBBYEFGmdd1S2A8EOdzVP9I5axwYosfLEMB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDAxNzIwCgYIKoZIzj0EAwIDSAAwRQIgDpu8Z1SSqxZVzqJ4YNX5mJ9sQLLa5%2F4AtBhbS7N%2FdqsCIQCRR1FY8pJof%2FNM5VFlfGXL8W8NACpll0gxiGS7zxQI9A%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 06 Jan 23 03:03 UTC