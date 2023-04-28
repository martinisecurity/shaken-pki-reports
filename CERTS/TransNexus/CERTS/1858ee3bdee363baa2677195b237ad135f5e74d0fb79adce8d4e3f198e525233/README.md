# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 518J

Tested At: 28 Apr 23 02:05 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -188 day(s)\
Subject: CN=SHAKEN 518J, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/transnexus/028fea0e03c28db631b53e6ebffddbe2.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC8zCCApigAwIBAgIQay2ftvxH4JXwB815NeObwTAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMTQyMDQ4MzlaFw0yMjEwMjEyMDQ4MzhaME8xCzAJBgNVBAYTAlVTMRkwFwYDVQQKExBUcmFuc05leHVzLCBJbmMuMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiA1MThKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEKu31Ys%2F%2BDxW7GFIU033kyQu8XHyD3%2BD%2BhbpQye2F2Nkjqa2x42Ls0D52MxlXxDJoOr9uwyUmyaB2Q2x1qATKOaOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBQMNWg%2FXxZpas%2Fj7eFhfYv3qNK%2F8DAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQ1MThKMAoGCCqGSM49BAMCA0kAMEYCIQCe88puSJhFUesvA88jUrRxgzIVswb49Qmu%2FQ5J8AY3DAIhALL5vs9xtYGsmI9We34rbyVmEfibJFI6tR8UGlG%2FCWQf)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 28 Apr 23 02:17 UTC