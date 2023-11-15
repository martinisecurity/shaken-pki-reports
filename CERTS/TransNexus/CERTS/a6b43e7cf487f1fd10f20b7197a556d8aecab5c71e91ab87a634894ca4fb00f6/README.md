# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 031K

Tested At: 15 Nov 23 15:59 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -370 day(s)\
Subject: CN=SHAKEN 031K, OU=SHAKEN, O=TISD, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.transnexus.com/031K/ba0288df-daa5-44b5-a1a3-3270140129ed.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC5zCCAoygAwIBAgIQZ4%2FLfYIaElEsdRhdJkTawDAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMTAxODM2NTlaFw0yMjExMDkxODM2NThaMEMxCzAJBgNVBAYTAlVTMQ0wCwYDVQQKEwRUSVNEMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiAwMzFLMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEA2otyUR7bCPIGK2ozboGh9D700pABVHHHb%2Fa0bY%2FZ6Tpy4DGW6VTnjEvdRx2pIUHZQ2s524J17DDDsKmwKbk4KOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBQB74WG0xTnvIunT5aNICTQWcrMTzAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQwMzFLMAoGCCqGSM49BAMCA0kAMEYCIQDa%2BO8TeVTQsXv9diKJ4oconj25X35D4W68k5XcMB0dsgIhAP9%2FpQpVu01L3oUQj%2BI%2BXf%2BgPPaATbtCLGcuNOBu5XvQ)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 15 Nov 23 16:51 UTC