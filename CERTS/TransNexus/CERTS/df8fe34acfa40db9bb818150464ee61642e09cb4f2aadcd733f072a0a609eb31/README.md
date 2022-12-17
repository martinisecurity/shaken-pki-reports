# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 849J

Tested At: 17 Dec 22 16:58 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -56 day(s)\
Subject: CN=SHAKEN 849J, OU=SHAKEN, O=Fuse.Cloud, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/86e241b8-9c8e-4431-b35f-4d92844a1da9/b15fd46f0c04f43108e7920a6c521f54.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC6zCCApKgAwIBAgIQRpg4CR6xxorA7Hqvf%2Bj2DTAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMTQyMDIwMzlaFw0yMjEwMjEyMDIwMzhaMEkxCzAJBgNVBAYTAlVTMRMwEQYDVQQKEwpGdXNlLkNsb3VkMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiA4NDlKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEmSm4pD7oIPeDFFa%2FyrcgVeWt9W4dZL7fdJY32AbYRmhL1968pVFZeBT2u91G6%2BSHWMkHqJIN9YoG1cT3clrcoqOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBRIBp%2Fc5Rl3aRPhQwJJ%2BSxIuAfCnTAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQ4NDlKMAoGCCqGSM49BAMCA0cAMEQCIBQBjZGLAnI4scutsjEtegsbpBlXqXDwHU6%2FaE9J%2BZSKAiAPdeLWNkF7j76bg8fjza15ZR34F%2Bvo3j7%2Bv25KcwPlCQ%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |


Generated: 17 Dec 22 17:07 UTC