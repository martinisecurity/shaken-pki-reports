# STIR/SHAKEN CA Ecosystem Compliance
## TransNexus

### Certificate b6d72ada27949e72427f323a03e4ddc7e6d2e803
Tested At: 2022-10-27 22:32:35 +0000 UTC\
Initial Validity Period: 60 day(s)\
Remaining Validity Period: 4 day(s)\
Subject: CN=SHAKEN 193E, OU=SHAKEN, O=8x8 Inc, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US

Link: https://certificates.transnexus.com/193E/7e5d014c-ed0c-4a12-8c38-df64aa8be83d.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIC6jCCAo%2BgAwIBAgIQXEvA1ecxt9LJyf5dNFgEuzAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjA5MDExMjMxMzdaFw0yMjEwMzExMjMxMzZaMEYxCzAJBgNVBAYTAlVTMRAwDgYDVQQKEwc4eDggSW5jMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiAxOTNFMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEDU5RGm%2BOaw8KEmUno25tRoySPa1sqNaZ2fXmZS1We9yS%2FYgH9flnVFq%2FVtWrDmwlgn493wpIrczpa%2B3sn1AwxaOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBSaZCW60ckudWr48I0tQb9Q0wxf7zAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQxOTNFMAoGCCqGSM49BAMCA0kAMEYCIQDQqjn1MRq3JoQ%2Bd2exVceRQwLPSwQVUG45dAVuYSn%2FgwIhAMNf5NWiKaWIo9TbMCtAKgSQdq1DnDSJdrSvqYaZFF%2Fq)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_incorrect_ku_encoding | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| w_cp1_3_subject_rdn_unknown | warn | United States SHAKEN CP | Only CN, C, O, and SERIALNUMBER can be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| e_cp1_3_subject_sn | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |

Generated: 27/10/2022 at 22:33:03