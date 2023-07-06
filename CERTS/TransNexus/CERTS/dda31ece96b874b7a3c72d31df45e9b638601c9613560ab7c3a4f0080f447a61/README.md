# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 107K

Tested At: 06 Jul 23 13:54 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -263 day(s)\
Subject: CN=SHAKEN 107K, OU=SHAKEN, O=Atheral, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/4ba6f296-2d7c-42c2-8fe4-12b74f5076df/bbb2ce089e69ca426c4375841dad4178.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC6TCCAo%2BgAwIBAgIQZExC3R6J6%2BkI12RkewbKvTAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMDgxNzIyMDBaFw0yMjEwMTUxNzIxNTlaMEYxCzAJBgNVBAYTAlVTMRAwDgYDVQQKEwdBdGhlcmFsMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiAxMDdLMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEzlG5Sw063R8VH0J8%2BW1XYdvX%2FGAB3S1CqtW9MlXk8cK6E31p5%2F%2FU3bS7vYb6n3Yk5%2BEsmGfM8TfcZa3dTjPz1aOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBQKUTAdao4ZybvcXcF1GhfWwh9j2DAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQxMDdLMAoGCCqGSM49BAMCA0gAMEUCIF4w%2FLlSqvjqVIfW3zfYJyc159U8nXfMydcfihH8CqjIAiEAv8IsMZ7IQnNdvafzVkYmVN%2BxYAg6A5EEisZtQHhlkts%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 06 Jul 23 14:08 UTC