# STIR/SHAKEN CA Ecosystem Compliance
## TransNexus

### Certificate 3475a0a95ab1463443919df71ceff158ee753348
Tested At: 2022-10-28 19:21:49 +0000 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: 5 day(s)\
Subject: CN=SHAKEN 107K, OU=SHAKEN, O=Atheral, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US

Link: https://certificates.clearip.com/4ba6f296-2d7c-42c2-8fe4-12b74f5076df/8624c57cb5bae15eefe0c18490948d41.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIC6DCCAo%2BgAwIBAgIQcg0SxTWkHk0lKE%2B2OEv54jAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMjYxNzI0MjRaFw0yMjExMDIxNzI0MjNaMEYxCzAJBgNVBAYTAlVTMRAwDgYDVQQKEwdBdGhlcmFsMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiAxMDdLMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEV1qwbnS4vGfhYfZtTHcyq%2BE7F2A5ZMExTzsIfdGbJ7YRXeVjgbsueScBqw%2BZEaKR1tCJHVh5yzGW1qPsUfGsaqOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBQqH%2BVohCWjNcECCGJzvnkTb1huxjAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQxMDdLMAoGCCqGSM49BAMCA0cAMEQCIDho1qbNw9pluHO3UK5e3VNIGCMD3bAIWMPd1xLBviZSAiAChL1kDWAznu2asA04Vg4qMCEeV13%2F%2F7ggrYURa5swnw%3D%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_incorrect_ku_encoding | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| w_pki_subject_rdn_unknown | warn | SHAKEN PKI Best Practice | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| e_cp1_3_subject_sn | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |

Generated: 28/10/2022 at 19:22:10