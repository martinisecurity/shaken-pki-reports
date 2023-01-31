# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 726J

Tested At: 31 Jan 23 21:39 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -105 day(s)\
Subject: CN=SHAKEN 726J, OU=SHAKEN, O=Votacall Inc, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/88dbb645-62bf-4d40-960c-4e8e3223880f/2161b4be0d61ff1a6e74f551cd1b8049.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC7jCCApSgAwIBAgIQfG0dSMlKVqDxnqPprdSmATAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMTEyMDIwMjZaFw0yMjEwMTgyMDIwMjVaMEsxCzAJBgNVBAYTAlVTMRUwEwYDVQQKEwxWb3RhY2FsbCBJbmMxDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDcyNkowWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQW1tCrdsar4ujvsRdlrtgKplkbgLnP5Mn%2F771TO25dEjLvt9tius2z0ZNeqUGY%2FhDl8XEliaEnn%2BLpNdm4%2Be3Ko4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCAIAwHQYDVR0OBBYEFE%2FlJwk%2FCQxU8JF6wjtTfWzBK0g1MB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDcyNkowCgYIKoZIzj0EAwIDSAAwRQIgLpiJbzJ4CJ8WBEpmUljnFb9n5JmeYgD%2FYHK8V19PfaQCIQC6o2NBg5muFyVy4m7YOwiJmziUPdIxrrbE0T0hf5maqg%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |


Generated: 31 Jan 23 21:50 UTC