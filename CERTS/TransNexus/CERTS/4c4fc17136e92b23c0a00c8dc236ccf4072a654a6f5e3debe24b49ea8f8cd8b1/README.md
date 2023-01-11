# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 297K

Tested At: 11 Jan 23 21:48 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -52 day(s)\
Subject: CN=SHAKEN 297K, OU=SHAKEN, O=Clarity Voice, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/859299ca-4c86-4dfd-9785-6c758bee1b37/fe12391f1a85a76c2a5f78dc986d9e14.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC7zCCApWgAwIBAgIQdOM92o%2ByN7d7QemCMZ%2FcVDAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjExMTMxOTU3MzJaFw0yMjExMjAxOTU3MzFaMEwxCzAJBgNVBAYTAlVTMRYwFAYDVQQKEw1DbGFyaXR5IFZvaWNlMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiAyOTdLMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAETUZjicP74%2F8XxhiSTGkp1JE9iX%2FyU%2FYJd8gfJjNihKiXGJMz6Uzz6QKihbleqoHfNCwCw0OxR%2B0NbrLSkDcBOaOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMB0GA1UdDgQWBBS%2FVV%2BiLXspL5c88kBjTPESCU5KQjAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQyOTdLMAoGCCqGSM49BAMCA0gAMEUCIBFYij5Dz3n5C8vZG0NR9e1M4Es7zYRDUR%2F0PvyIGB%2BJAiEAq35BavefQhozlWkHtp%2F3tIU5bHF8SAm04d8imarB2JU%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 11 Jan 23 21:59 UTC