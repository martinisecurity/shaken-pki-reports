# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 518J

Tested At: 31 Jan 23 21:39 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -84 day(s)\
Subject: CN=SHAKEN 518J, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/transnexus/fe9c54ed46aa52dc9103cc75a9c0d5cb.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC8jCCApigAwIBAgIQfx9%2Fz%2BJmRuHcY1mesRHs%2FTAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjExMDEyMDQ5MTJaFw0yMjExMDgyMDQ5MTFaME8xCzAJBgNVBAYTAlVTMRkwFwYDVQQKExBUcmFuc05leHVzLCBJbmMuMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiA1MThKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE%2BAMOIZeAJb97jAklvYf4NNE%2FYdQTDpjpfMuVgyhHW7ZnW3d%2BkGqLGJyo6qJzjk2p8Ky7bvYzrWrPCINmySc0h6OCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMB0GA1UdDgQWBBS1kwYo2dfR5HZzDOz%2BGfRtqrA2fDAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQ1MThKMAoGCCqGSM49BAMCA0gAMEUCIQDs1MsqI1wCVBWsMrCVQD1bJkZgB98LSqrod5vWxSCzywIgTUbCW8VfPQQEfQik%2BBCbAy8Fb1D%2B2pAQ1%2FCRycgitTo%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 31 Jan 23 21:50 UTC