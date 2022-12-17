# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 177K

Tested At: 17 Dec 22 16:58 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -51 day(s)\
Subject: CN=SHAKEN 177K, OU=SHAKEN, O=Cytracom\\, LLC, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/8d118994-4994-4735-ac71-42c0bbb7848f/f18ba4316996eead719f043a340663da.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC7zCCApWgAwIBAgIQc595HOP3atx6ko%2F7GlQmIDAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMTkxNzM4MTBaFw0yMjEwMjYxNzM4MDlaMEwxCzAJBgNVBAYTAlVTMRYwFAYDVQQKEw1DeXRyYWNvbSwgTExDMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiAxNzdLMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEKkNPm2d%2FMTtQtqmX00LHleWDtN2cYenLMgmRXD3STXxabT4Oth0fwuFj%2BmAkuIs5HLk8S1IHbGl8SesqXr18GaOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBThk0Nf0U1NmgGDvdP%2FZJ4E2fLOcDAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQxNzdLMAoGCCqGSM49BAMCA0gAMEUCIGmf6kPTfPkvLYFewYW8uU%2BIkbZ9Zl6toPJ4DFG05yhdAiEAppmgonytQxcX4MxEVaTXbSjbr99fbdrmPdW4i6ZzrTI%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 17 Dec 22 17:07 UTC