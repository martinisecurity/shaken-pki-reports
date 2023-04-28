# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 0172

Tested At: 28 Apr 23 02:05 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -179 day(s)\
Subject: CN=SHAKEN 0172, OU=SHAKEN, O=CallTower, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.transnexus.com/0172/cbcf3408-6748-4bd7-a9ff-8886d21b6824.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC7DCCApGgAwIBAgIQTwyEDnfxYXtOnJ2hNucHxTAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMjMyMjExNDhaFw0yMjEwMzAyMjExNDdaMEgxCzAJBgNVBAYTAlVTMRIwEAYDVQQKEwlDYWxsVG93ZXIxDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDAxNzIwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQczMdtLq2oVr0XSyO2MY7b8Gtz2%2ByU48if1gVP9ifR3%2BJfMWPR0tCIrCFZG2GnUPV1Wj%2FlkIGgqoYRmqMcwKIJo4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCAIAwHQYDVR0OBBYEFIBmhe%2FZE9H%2BMC6K41X9wLJhWKqKMB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDAxNzIwCgYIKoZIzj0EAwIDSQAwRgIhALjnC%2F86PieQ8NS%2Bzl4%2BZPWSEH1OfjDWKjA8ZGgCi3kxAiEA54AtPHL%2Bgo9X%2BfdBccPS4EvWqQW2g%2FV0yFIKyGfX0W8%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |


Generated: 28 Apr 23 02:17 UTC