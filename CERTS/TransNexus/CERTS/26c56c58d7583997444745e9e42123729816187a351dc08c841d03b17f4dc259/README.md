# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 8526

Tested At: 21 Aug 23 20:07 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -320 day(s)\
Subject: CN=SHAKEN 8526, OU=SHAKEN, O=MetTel, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.transnexus.com/8526/88a92c6f-973b-4d14-9221-4c577d177131.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC6TCCAo6gAwIBAgIQdFXiFnLCivesmJD%2BSKh2gTAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjA5MDUxMDMyMDlaFw0yMjEwMDUxMDMyMDhaMEUxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKEwZNZXRUZWwxDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDg1MjYwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAARssCpm9Vt86MClwpg2uMPhNyjyPpwRp%2FCyZIDo5i%2F4p3VQ2VGCgYT5e0UIhKQQ785JhhA21JovusPyO3dUJoZPo4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCAIAwHQYDVR0OBBYEFIXvykEVoSFxmpVSKsL2BPN378MHMB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDg1MjYwCgYIKoZIzj0EAwIDSQAwRgIhAMHemDrba%2BvLigSRZPBKpgxaDF0LcVYJNC17RimdwP%2FTAiEAx5AeE3QwDmEZ2swsYvvxRryslQlukw5Rqsp6GLLHi%2Bg%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 21 Aug 23 20:18 UTC