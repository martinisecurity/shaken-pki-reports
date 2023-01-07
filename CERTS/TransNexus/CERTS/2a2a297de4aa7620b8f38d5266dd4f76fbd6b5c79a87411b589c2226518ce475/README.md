# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 193E

Tested At: 07 Jan 23 19:08 UTC\
Initial Validity Period: 60 day(s)\
Remaining Validity Period: -38 day(s)\
Subject: CN=SHAKEN 193E, OU=SHAKEN, O=8x8 Inc, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.transnexus.com/193E/88d4a5ec-9d92-462a-9149-735bef3556fc.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC6TCCAo%2BgAwIBAgIQQGi41FpVBWylaeou%2BTfJfDAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMDExMjMxNThaFw0yMjExMzAxMjMxNTdaMEYxCzAJBgNVBAYTAlVTMRAwDgYDVQQKEwc4eDggSW5jMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiAxOTNFMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE1nOQaFwqQTT2vKXEGQX7gGM2%2BIN86VKo%2F4kwDeN7OIvWvCECxOu%2F8SmgCI4yfZcUGPeb9IByruzsMZpGIoSlPqOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBTIuwL07DSauBmVJBaOJIszvAmqVzAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQxOTNFMAoGCCqGSM49BAMCA0gAMEUCIFJMv1Clevk2T8XIc%2FIJil4TU96%2FOgt2TxCTUFkPEs7RAiEAqrDK1hYtECDq3adBvUmOqVzouoGPmcqv3q3J3DyNcMo%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 07 Jan 23 19:18 UTC