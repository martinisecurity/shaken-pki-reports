# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 518J

Tested At: 21 Aug 23 20:06 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -312 day(s)\
Subject: CN=SHAKEN 518J, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/transnexus/25cdaea7fd01687e813b32406861e194.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC8jCCApigAwIBAgIQa6o%2BfqA6BBJ0R4snM7jUyjAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMDUyMDQ4MzJaFw0yMjEwMTIyMDQ4MzFaME8xCzAJBgNVBAYTAlVTMRkwFwYDVQQKExBUcmFuc05leHVzLCBJbmMuMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiA1MThKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEccoZTCvLR%2B6z8XKSC8hsmEX1Q8CmbVqNIE7eA5jYFpUulpa1yFJiv3iPRe4UXLLi5B3%2FvyJIOhvr2d6ciUYCOKOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBSO1G1rUUZaqzcd4BbBvlvy81UoUDAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQ1MThKMAoGCCqGSM49BAMCA0gAMEUCIQDPelsMDfBh5djNjCD5AZWz5hqGLXAd6q%2BIdDpmVTWvhwIgFlYPXD8jokbmM9aXI2DAsWGy9rHc8wtk25C73xzo5oo%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 21 Aug 23 20:18 UTC