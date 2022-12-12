# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 691J

Tested At: 12 Dec 22 23:36 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -58 day(s)\
Subject: CN=SHAKEN 691J, OU=SHAKEN, O=Toly Digital Networks Inc, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/84c98183-b28e-415a-8dc4-ba02b9ca5418/d99780f1c4a7a1ad60a27b0d675c73f5.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC%2BzCCAqGgAwIBAgIQfS9OJHhCirjDGxQfWhoJBjAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMDgyMDIwMDdaFw0yMjEwMTUyMDIwMDZaMFgxCzAJBgNVBAYTAlVTMSIwIAYDVQQKExlUb2x5IERpZ2l0YWwgTmV0d29ya3MgSW5jMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiA2OTFKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEyU38c2IecoBSiXNJ2R%2FNIJzMp4V6PfSeyMjkFsvAndU3GF3z4xvnVOBKI8RmBqRUL5G%2B0YF%2FZb2FJXpws%2BLnqKOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBQxcZE20mpo6MoIx2MaeLepoeXcwzAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQ2OTFKMAoGCCqGSM49BAMCA0gAMEUCIECg%2FUUBD8UXHSAtVkxNUaDu%2BeXmhKuogvWGNsHiWVOIAiEAza7hQzIi4pC5hV8sbCtxbk7sGbhBCRDIaQuVg2g7MmE%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 12 Dec 22 23:45 UTC