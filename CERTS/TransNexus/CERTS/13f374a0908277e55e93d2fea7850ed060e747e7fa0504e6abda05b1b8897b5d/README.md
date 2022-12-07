# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 107K

Tested At: 07 Dec 22 18:45 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -29 day(s)\
Subject: CN=SHAKEN 107K, OU=SHAKEN, O=Atheral, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/4ba6f296-2d7c-42c2-8fe4-12b74f5076df/cb4f0e786038a1221b086198e6449705.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC6TCCAo%2BgAwIBAgIQRRiYQrtBPkovpR77VubvJzAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjExMDExNzI0NTRaFw0yMjExMDgxNzI0NTNaMEYxCzAJBgNVBAYTAlVTMRAwDgYDVQQKEwdBdGhlcmFsMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiAxMDdLMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE0erf%2FYswCwnq2TJEAXN6KIG62vv49Kt2vW4vhXwvGe2TONyr6uDClHP%2B%2FKT6Rpk7FVnITwqGInC1chxTNFQp2KOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMB0GA1UdDgQWBBSOVbnMcdh3vRk0X728irqZIDltvDAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQxMDdLMAoGCCqGSM49BAMCA0gAMEUCIE902wnTY6fgrhMDGgqBai5MumnbNDD%2FlE%2Bzc35Uq9W3AiEAw5lVq0w%2B6vM98hmtDqtDn2yxaxABmhv3Pt%2FV0cmEWvc%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 07 Dec 22 18:54 UTC