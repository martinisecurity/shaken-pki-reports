# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 056J

Tested At: 28 Apr 23 02:04 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -200 day(s)\
Subject: CN=SHAKEN 056J, OU=SHAKEN, O=Telco Experts, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/384afe84-56e7-4f0a-a8f8-9be6e100ab3e/86041b4ca313a03fab1401d13a820e26.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC7jCCApWgAwIBAgIQSq68td6E%2FBAePXOkdAUPOTAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMDIxMzQyNTNaFw0yMjEwMDkxMzQyNTJaMEwxCzAJBgNVBAYTAlVTMRYwFAYDVQQKEw1UZWxjbyBFeHBlcnRzMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiAwNTZKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEPs9tTG9Wuwl9SCKt8x8jj42XX035kQfR8W4YkYGZbSpnf1bZ%2BtMPCPu1nESybEa29HkHuAjHsw6jA5lispFUiKOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBRPXGhns3NYh2fZ%2BEC4%2B%2F5TsyschjAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQwNTZKMAoGCCqGSM49BAMCA0cAMEQCIF7xzC3XTNfg7mQcgmSqVqNY6vCT%2Fuvn%2FDunKjAQO9fvAiBhSLqOQ2mBBuoaYOmZIwfCZIm5wxxr7A2IqufcLYlVvQ%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |


Generated: 28 Apr 23 02:17 UTC