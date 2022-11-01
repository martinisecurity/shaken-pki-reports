# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 735J

Tested At: 01 Nov 22 22:10 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: 1 day(s)\
Subject: CN=SHAKEN 735J, OU=SHAKEN, O=AVOXI Inc., C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/11124177-79f3-48b2-867a-386d4dc61c99/467ce6407e421d4db6b024d99ba3d4ea.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIC7DCCApKgAwIBAgIQR0qb8MA4LZH1uoh8BYLVLzAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMjYyMDE4NTJaFw0yMjExMDIyMDE4NTFaMEkxCzAJBgNVBAYTAlVTMRMwEQYDVQQKEwpBVk9YSSBJbmMuMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiA3MzVKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEoqnxnxPsJKJDB8UYDnd4F%2FK98v5CAD%2B%2Fa8OhQoZ10FiRj8gLHAEFAdVXXgDHbFYHVdB0vGBKWBeie%2FoBVBCeV6OCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBQAi3eVodpbqnkVRsShqsk2QiBJKjAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQ3MzVKMAoGCCqGSM49BAMCA0gAMEUCIQDYdHuOrMHnUKv2%2Bpmlm09LxE1lBFpuk%2FWkHZipK6a9CQIgfQCKXjZI%2BrzGYKIhZJDwVm%2FLxhDfhBQ9d1uleMsAgRA%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 01/11/2022 at 22:19:34