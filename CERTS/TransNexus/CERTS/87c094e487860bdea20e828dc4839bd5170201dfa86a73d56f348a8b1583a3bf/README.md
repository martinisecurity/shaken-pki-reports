# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 177K

Tested At: 12 Apr 23 01:03 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -182 day(s)\
Subject: CN=SHAKEN 177K, OU=SHAKEN, O=Cytracom\\, LLC, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/8d118994-4994-4735-ac71-42c0bbb7848f/3c51bbddeccc4e1d7c27af971c824a68.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC7jCCApWgAwIBAgIQX48l8tasNRzTPL5IZMUNUTAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMDQxNzM2NTNaFw0yMjEwMTExNzM2NTJaMEwxCzAJBgNVBAYTAlVTMRYwFAYDVQQKEw1DeXRyYWNvbSwgTExDMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiAxNzdLMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEfLTWzBioHFA3t8EgQ072mWVBmjpbsSuxKCHk8NBC3hpC3Kx8ymstzkeumvw2DOZVSiffPcRzyYrNyrW4cuM3YKOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBTy29yiBcRVzfIfolTulFMgH6jf9DAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQxNzdLMAoGCCqGSM49BAMCA0cAMEQCIENcI1m%2FA3CtGIZQeuhdxigi4Ekv5X9HJovSmib7Y2prAiA8ybnQElIYP85f%2Bnuto1VMFx0I4IF3gsilXGTcFY4Fug%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 12 Apr 23 01:46 UTC