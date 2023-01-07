# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 297K

Tested At: 07 Jan 23 19:08 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -68 day(s)\
Subject: CN=SHAKEN 297K, OU=SHAKEN, O=Clarity Voice, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/859299ca-4c86-4dfd-9785-6c758bee1b37/34b0b778ed9b4f95b12b2a1c9fbb50de.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC7jCCApWgAwIBAgIQYACplquDdJK6zdl8jaULBjAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMjMxOTU2MDBaFw0yMjEwMzAxOTU1NTlaMEwxCzAJBgNVBAYTAlVTMRYwFAYDVQQKEw1DbGFyaXR5IFZvaWNlMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiAyOTdLMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAElU%2F5SbP5aDN8lVvuO2OBJ22t4cTeLCLI2pPlxznr2%2FmCsYBYsNCNbD43htlXVMwxoD9zWb1AbZlErAYdFEnO7qOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBS%2Bh0rZJn4XaQIz5aLNsDmjYhpVDjAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQyOTdLMAoGCCqGSM49BAMCA0cAMEQCIH0hrU8vEDPKDX4N63M86CAlKUrpe7LSpGLSby1rbh6XAiAz6tcmncQIo7SxEsARBMgMQHFjxnt0AiDZg36xEZDi2A%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 07 Jan 23 19:18 UTC