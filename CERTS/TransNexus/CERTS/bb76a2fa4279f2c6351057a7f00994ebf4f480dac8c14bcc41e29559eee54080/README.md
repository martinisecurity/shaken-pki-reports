# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 495J

Tested At: 11 Jan 23 23:07 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -109 day(s)\
Subject: CN=SHAKEN 495J, OU=SHAKEN, O=Stratus, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/94c136bd-6438-40a6-b740-7765d5ead597/8b3ee25d58998b94749600fcd5ec8b75.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC6TCCAo%2BgAwIBAgIQYu%2B%2BAXM8sZM9yhczyGJqdzAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjA5MTcyMDE3NTBaFw0yMjA5MjQyMDE3NDlaMEYxCzAJBgNVBAYTAlVTMRAwDgYDVQQKEwdTdHJhdHVzMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiA0OTVKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEw6WNBNlcWFKKz4AqnZ2usZtgBgCL8wZKeGYoItxgde1D116drYbd%2FZukqzX1JsK%2FOmA%2FvkGie5b31nTMbmkeCKOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBQPhgmSedjwgkJg%2FdRJy3nvY9wd5TAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQ0OTVKMAoGCCqGSM49BAMCA0gAMEUCIEPGKFvy7g2D1YHh6DvWv%2Fk02QJVNhJDRVR%2FNBW54MBzAiEAoorzXbkP8O4JfHqYjgNmUGgiYsmBJcE6MzaUmhZt2VI%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 11 Jan 23 23:18 UTC