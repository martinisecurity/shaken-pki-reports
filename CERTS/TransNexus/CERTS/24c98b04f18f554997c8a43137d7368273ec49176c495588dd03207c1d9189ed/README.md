# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 735J

Tested At: 31 Jan 23 21:39 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -114 day(s)\
Subject: CN=SHAKEN 735J, OU=SHAKEN, O=AVOXI Inc., C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/11124177-79f3-48b2-867a-386d4dc61c99/60797bd734fa6f7e6aa28e23adf7e4d0.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC7DCCApKgAwIBAgIQSayep0mEyMUz%2BI0r3ezfhDAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMDIyMDE2NDZaFw0yMjEwMDkyMDE2NDVaMEkxCzAJBgNVBAYTAlVTMRMwEQYDVQQKEwpBVk9YSSBJbmMuMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiA3MzVKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEw18ZpsBoNr482NR6nvq9eHlnGyW1muw2Gs86yMtv6n00xph0%2B4B%2F8KbNn8BsbgxyZ7sjA60BCFHNL%2B9Fm7GfG6OCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBQfMNSw55SnoleljyHkZhJFV8y9zzAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQ3MzVKMAoGCCqGSM49BAMCA0gAMEUCIFATjZMBPXVp7SIS8pAMX9rY%2B8gJtTKP7Ue8pfJNW7a5AiEAr3WnTcXE0xA8J94FjS9zT6kYYcBu%2BksPs2%2FHv33QuCo%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 31 Jan 23 21:50 UTC