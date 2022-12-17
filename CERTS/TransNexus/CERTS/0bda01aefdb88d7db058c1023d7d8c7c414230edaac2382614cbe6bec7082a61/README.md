# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 469A

Tested At: 17 Dec 22 12:13 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -47 day(s)\
Subject: CN=SHAKEN 469A, OU=SHAKEN, O=T3 Communications Inc., C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/88814589-ad88-4c53-b3c3-47f8334afb98/0edc94c91013b6897f7e8960ef0cedc9.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC9zCCAp6gAwIBAgIQQHq0hqzyNSDR2cEI7wL%2B%2BTAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMjMyMDIxNDhaFw0yMjEwMzAyMDIxNDdaMFUxCzAJBgNVBAYTAlVTMR8wHQYDVQQKExZUMyBDb21tdW5pY2F0aW9ucyBJbmMuMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiA0NjlBMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE%2ByWsFlThCHHEmusg%2BNjfJaq219vXxfUFr9%2FwOFw2McmXyFKHrMLHS5Vty9G3vpYp9i6Pmt%2FYtFPqVz%2BTJHTM36OCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBR4Rw3C6MM1TETAwcc2kk8toNEBfTAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQ0NjlBMAoGCCqGSM49BAMCA0cAMEQCIEWtJGfsTF4SD5p8BCovplZN0ftBdWkyXgH2APtmkOJKAiBW%2FL%2BcTrbvewn6kDXZnsNOQNmgCVGFug0dISdqSZ0jDw%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 17 Dec 22 12:22 UTC