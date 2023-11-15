# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 737J

Tested At: 15 Nov 23 16:49 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: -117 day(s)\
Subject: CN=SHAKEN 737J, OU=SHAKEN, O=US Internet Corp, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://sti.ravon.net/c/737J_2022-07-20

[View certificate details](https://understandingwebpki.com/?cert=MIIC8jCCApigAwIBAgIQdpUZTniQemnsKYAFkMKZuDAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjA3MjAxNzAwMzhaFw0yMzA3MjAxNzAwMzdaME8xCzAJBgNVBAYTAlVTMRkwFwYDVQQKExBVUyBJbnRlcm5ldCBDb3JwMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiA3MzdKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEz0u2mFRDToS9elHZ1YmgSbZLcRrgm0EAgBKx4EBC%2BqbY1oWPvhaGYEgnRWcE6HhJtQB94ifKLkzcCitCPIDtmqOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBRC3uLWcfXMCuVf91n90OzRsf8W3jAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQ3MzdKMAoGCCqGSM49BAMCA0gAMEUCIQC8RC52aqcokdrjr8CQ3RWTCswnFdDo0PDQrCk3aypo7wIgdqwJUWbtN15fjoSnSRIEU7%2BS44DFZRUH%2F%2B5XQsZeCHM%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 15 Nov 23 16:51 UTC