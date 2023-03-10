# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 854D

Tested At: 09 Mar 23 22:58 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -133 day(s)\
Subject: CN=SHAKEN 854D, OU=SHAKEN, O=Xtel, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/627e4253-759f-447e-b1a0-c9cda9a6d597/eef13829c4b0f7d34e0feab46b71acf9.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC5TCCAoygAwIBAgIQRkNzAZLLOYZu0xX4S09AyjAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMjAyMDIxMTBaFw0yMjEwMjcyMDIxMDlaMEMxCzAJBgNVBAYTAlVTMQ0wCwYDVQQKEwRYdGVsMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiA4NTREMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEiuLVU3DpNb0fOWTa4UGLtxONfpAN5VIP1R66PkNcb94o%2Bwcj1TIQre9RHWPnmChp%2BU6gH2K1rbM%2BUKlG2YN4bKOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBSYLUW4i5Gvana0UbBDYkZvsrOrKTAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQ4NTREMAoGCCqGSM49BAMCA0cAMEQCIGzQyAITUDOvLsdw6oeuu%2Btoq7Igvxd%2F9VTehyeP8fiIAiBZQJ9nBoKgrt71mo%2FjJwZvFryV3TC7QxlSVyrnqBhD4Q%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 10 Mar 23 02:25 UTC