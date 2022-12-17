# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 854D

Tested At: 17 Dec 22 16:58 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -41 day(s)\
Subject: CN=SHAKEN 854D, OU=SHAKEN, O=Xtel, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/627e4253-759f-447e-b1a0-c9cda9a6d597/0975d43d6640bf49c8eb3d523bcac470.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC5TCCAoygAwIBAgIQdEFPm5wPo9u5rBO9eumhtjAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMjkyMDIxNTlaFw0yMjExMDUyMDIxNThaMEMxCzAJBgNVBAYTAlVTMQ0wCwYDVQQKEwRYdGVsMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiA4NTREMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE5WPfGiYmzBTtqsJY5eUEeVPlqmyVRe09kabMftVxudJXYS9pafd%2Be13hTByXFC7JZYValkXhNYAujbrdN4mgMqOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMB0GA1UdDgQWBBR2z0phpWyERWmxFi%2BJc4rQxfd5cDAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQ4NTREMAoGCCqGSM49BAMCA0cAMEQCIFDmPpr%2BD8w3%2F6H54loRKEP1ubtPbf5U14d%2BGQHZk06IAiB%2FwtT3JSD0pGgZwhxPtjMk5e5f7YJL6POgUc0HUOVPiw%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 17 Dec 22 17:07 UTC