# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 518J

Tested At: 17 Dec 22 00:03 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -53 day(s)\
Subject: CN=SHAKEN 518J, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/transnexus/9c0c31ac7c17c03439a63065c55862aa.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC8jCCApigAwIBAgIQVxD7fZ%2F41rt48HNfUic4XjAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMTcyMDQ4NDFaFw0yMjEwMjQyMDQ4NDBaME8xCzAJBgNVBAYTAlVTMRkwFwYDVQQKExBUcmFuc05leHVzLCBJbmMuMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiA1MThKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAETEnwEQfkkeJGWVj4hFrJNvtZT%2FOLwiRzF46q2zux5YObWKpayyImfgIbejX5v61ItguTS1BBJoXRb1UMWKIJ5aOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBREPpleNEZGwSgr4hbTCzhWZjkdiTAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQ1MThKMAoGCCqGSM49BAMCA0gAMEUCIBDFsUJVfjkG15RJSzQXTHh12EJ8Goqhp4iYZ8en2s1gAiEA8qdOyuLbP95uENEf8%2Fvdx8gFgRqS4GVk1b70oFOXckk%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 17 Dec 22 00:12 UTC