# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 495J

Tested At: 21 Aug 23 20:04 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -309 day(s)\
Subject: CN=SHAKEN 495J, OU=SHAKEN, O=Stratus, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/94c136bd-6438-40a6-b740-7765d5ead597/c554fd909a16573071e40fbac14f35bd.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC6TCCAo%2BgAwIBAgIQXknWFmw2dlRl8TCQV2b5lDAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMDgyMDIwMTJaFw0yMjEwMTUyMDIwMTFaMEYxCzAJBgNVBAYTAlVTMRAwDgYDVQQKEwdTdHJhdHVzMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiA0OTVKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEKuD0pUcCH%2FUGw7DfG2VDJysdrNK%2BaDTJMwnEQUAPq%2FTsJgBfkICUXqdORRp4uBeaXZUKtMC7yIhV288cwvE7hqOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBT%2FoUVKMhRX8R4gZDcvBRIqQ3CnHzAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQ0OTVKMAoGCCqGSM49BAMCA0gAMEUCIHoRCsc37vizOdlA57QjFh%2B5IqMBrBE%2BWXmTexFalchJAiEAxvbZTIvWihFY9eFIbGtxNQGbH9plgu8%2FfM0CzguTWXs%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 21 Aug 23 20:18 UTC