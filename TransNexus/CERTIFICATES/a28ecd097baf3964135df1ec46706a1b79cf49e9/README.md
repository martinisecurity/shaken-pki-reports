# STIR/SHAKEN CA Ecosystem Compliance
## TransNexus

### Certificate a28ecd097baf3964135df1ec46706a1b79cf49e9
Tested At: 2022-10-28 18:22:40 +0000 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: 5 day(s)\
Subject: CN=SHAKEN 056J, OU=SHAKEN, O=Telco Experts, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US

Link: https://certificates.clearip.com/384afe84-56e7-4f0a-a8f8-9be6e100ab3e/bd57ebc311aedc3a3d6e9e257b5f542c.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIC7zCCApWgAwIBAgIQdnCcx0CgkLWEg3tAWoezMDAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMjYxMzQ1MzZaFw0yMjExMDIxMzQ1MzVaMEwxCzAJBgNVBAYTAlVTMRYwFAYDVQQKEw1UZWxjbyBFeHBlcnRzMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiAwNTZKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEQ3rLCF6VzcaxdPMbgC6bLimjaNzXeKMePVsDqsbW2%2F1ZHaCkkMY1IjjeIQFUt92ff%2FWWWgYF7l%2FV7%2BPuacB1aqOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBTUTo%2FL8gjsLWWCeqGB8eJvYovvLDAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQwNTZKMAoGCCqGSM49BAMCA0gAMEUCIQDWuCU%2BpGl9%2FfLQq1W1qhNZrJCGiMLoyxXPsAxcm%2FcY6gIgfJft4HyTRwQc3c1jBSbJzq5Jurf4Wu0VewWlb3Enmm8%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_incorrect_ku_encoding | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| w_pki_subject_rdn_unknown | warn | SHAKEN PKI Best Practice | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| e_cp1_3_subject_sn | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |

Generated: 28/10/2022 at 18:22:55