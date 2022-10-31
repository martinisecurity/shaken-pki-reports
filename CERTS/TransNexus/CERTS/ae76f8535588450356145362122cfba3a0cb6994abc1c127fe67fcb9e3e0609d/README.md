# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 790J

Tested At: 31 Oct 22 20:24 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: 5 day(s)\
Subject: CN=SHAKEN 790J, OU=SHAKEN, O=Viirtue, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/a429df3a-62d2-4851-90c8-fc446859fb08/59d21baf9d538dbde3950d938f2afaf9.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIC6DCCAo%2BgAwIBAgIQVS1SKd4iRyV5DxMpZ8NaojAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMjkyMDIzMTdaFw0yMjExMDUyMDIzMTZaMEYxCzAJBgNVBAYTAlVTMRAwDgYDVQQKEwdWaWlydHVlMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiA3OTBKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEOqTYOw7nLTTJOT4SpwgiS3K6%2Bzqg7ZkGZ5fSZ0jIgpgx3vQoOu%2B6jWUTHlqyouXBB9B9T0uFXa2l8u90Uk73sqOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMB0GA1UdDgQWBBTvC0UcAKcSVXdThorbVc3US4T9azAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQ3OTBKMAoGCCqGSM49BAMCA0cAMEQCIFxVMvDLknbO7QWvqdj4ALvYgb3RkgeFLQV2uwrjyQ%2BKAiAKTtMaXTy6FdTSd3POHjxMYcprnjzU1Xx9Fx%2BNMabFCw%3D%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_cp1_3_subject_sn](../../ISSUES/e_cp1_3_subject_sn/README.md) | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_pki_subject_rdn_unknown](../../ISSUES/w_pki_subject_rdn_unknown/README.md) | warn | SHAKEN PKI Best Practice | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 31/10/2022 at 20:32:42