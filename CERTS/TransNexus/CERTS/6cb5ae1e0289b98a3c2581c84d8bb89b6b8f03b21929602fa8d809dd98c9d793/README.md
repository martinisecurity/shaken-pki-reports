# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 186K

Tested At: 15 Nov 23 16:14 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -393 day(s)\
Subject: CN=SHAKEN 186K, OU=SHAKEN, O=Go2Uno LLC, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/92df566f-5229-4e9a-b7c7-ec203fc4196d/f6fd06cff383c1babdd4f9e1eca1b778.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC7TCCApKgAwIBAgIQRKdQ6M44R0KbbkzpNp24iTAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMTExNDI2MzdaFw0yMjEwMTgxNDI2MzZaMEkxCzAJBgNVBAYTAlVTMRMwEQYDVQQKEwpHbzJVbm8gTExDMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiAxODZLMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE6EPczLXwmeIZ%2BYk1tNKXzYQszXtwlSONjPdhgb6l0mRYGJOHMUSg6VYaPkQq4YqpzqxQJdEckNr3Y68c1VK6RaOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBSW7TVk07q6lr0DePybMKpKp662JDAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQxODZLMAoGCCqGSM49BAMCA0kAMEYCIQDvtAfaGgfDB1Dnwaxu3mqsnte6i%2B%2FoTu9WiEr%2Bci%2Bo8gIhAJC33OMOYlfnXHWy4NhwJ1%2BMOvsSKC%2F4ZdlctJjyytx0)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |


Generated: 15 Nov 23 17:17 UTC