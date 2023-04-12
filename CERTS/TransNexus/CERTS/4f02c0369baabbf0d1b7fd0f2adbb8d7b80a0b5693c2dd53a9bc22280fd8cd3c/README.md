# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 722J

Tested At: 12 Apr 23 21:46 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -163 day(s)\
Subject: CN=SHAKEN 722J, OU=SHAKEN, O=EvolveIP LLC, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.transnexus.com/722J/a547c920-fe4c-4086-a4cf-a744f0a1a4ce.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC7jCCApSgAwIBAgIQUmuUufynvg9JFyDpuvleRTAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMDExNTIzMjhaFw0yMjEwMzExNTIzMjdaMEsxCzAJBgNVBAYTAlVTMRUwEwYDVQQKEwxFdm9sdmVJUCBMTEMxDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDcyMkowWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATXvcrxfdNgF8U6D9i1dI9aAYkV0jgjUS0bVyeZslC55nU%2BSSH4zQqXb28uqv7mrkyFzVr57fO%2FIH1gkozCkE9Co4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCAIAwHQYDVR0OBBYEFIRxVEuov2K4AkFbiPiIyVolPHe2MB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDcyMkowCgYIKoZIzj0EAwIDSAAwRQIgFLZwgo%2FYdr9BVnPgbppYnQi9uj8oSZ6202kCILbrBF0CIQDXUlrPK%2Fvh00BrRsdo5nTARHVHC8RpEI1%2FtBceNBJ36w%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 12 Apr 23 22:02 UTC