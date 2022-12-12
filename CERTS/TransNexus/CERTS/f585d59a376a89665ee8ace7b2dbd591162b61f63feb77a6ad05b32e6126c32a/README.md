# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 672B

Tested At: 12 Dec 22 23:36 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -31 day(s)\
Subject: CN=SHAKEN 672B, OU=SHAKEN, O=Clear Rate, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/dcef9a97-b864-43ac-830b-2290a8d0f002/6cbc48b81e9f047a5eabc057c2b439d8.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC7TCCApKgAwIBAgIQf1Fmyt3Qy%2FScnthfy2upXjAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjExMDQyMDI0NDVaFw0yMjExMTEyMDI0NDRaMEkxCzAJBgNVBAYTAlVTMRMwEQYDVQQKEwpDbGVhciBSYXRlMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiA2NzJCMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEJJ4rWTfyJwdyLqT4XQ1cKY2sDQx60GElQDX1w0wlUZhcSJ3wt0UvapGZRFUqrmsWbyiSsj4uqCNTpxOe65ErB6OCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMB0GA1UdDgQWBBSLBVahh9G1pS2nlOKhHbTvwVbUozAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQ2NzJCMAoGCCqGSM49BAMCA0kAMEYCIQCBd7Bstkr88POwA%2BcDssZERWUOz9eA9AeqFEO5huf%2FPQIhAOsm9l9fYHrUvo80BqJDxThta9Un7TLEE%2BVRI5KqVkQl)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 12 Dec 22 23:45 UTC