# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 179K

Tested At: 17 Dec 22 00:02 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -59 day(s)\
Subject: CN=SHAKEN 179K, OU=SHAKEN, O=The Tech Consultants\\, LLC, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/2d93d07b-20ff-4289-a40d-959976cc0595/0630c442a9ee31c6271432091950d170.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC%2BjCCAqGgAwIBAgIQREX63H72pQDeqU1mvfgN6DAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMTExODQ0MDhaFw0yMjEwMTgxODQ0MDdaMFgxCzAJBgNVBAYTAlVTMSIwIAYDVQQKExlUaGUgVGVjaCBDb25zdWx0YW50cywgTExDMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiAxNzlLMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE1eR8WQlR5YWTYZMgsjRNHFlmVOdNUV4dg2l%2Bk5zdChi%2FHNggNh5kwh9EHY6GfAFqxe%2FTBqg4yeXayXH16c%2F56qOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBTNFznRI%2FUbKBrFGfW2gAdGUZhU5DAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQxNzlLMAoGCCqGSM49BAMCA0cAMEQCIHClM%2Bat8rnx4bNp1oT1wQB11pBYZCX7aAk8NusNCiMGAiASo0pkLElAuJFSDe4gTRr3nQRwQCv15jJ6e8Po%2B%2F%2F8FQ%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |


Generated: 17 Dec 22 00:12 UTC