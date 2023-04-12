# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 854D

Tested At: 12 Apr 23 21:44 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -155 day(s)\
Subject: CN=SHAKEN 854D, OU=SHAKEN, O=Xtel, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/627e4253-759f-447e-b1a0-c9cda9a6d597/b0066404e0aac0b7ab4cd171159c982d.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC5jCCAoygAwIBAgIQQ9OsXKGxeTayK8gOIYL32zAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjExMDEyMDIyMTNaFw0yMjExMDgyMDIyMTJaMEMxCzAJBgNVBAYTAlVTMQ0wCwYDVQQKEwRYdGVsMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiA4NTREMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEIRPgMMBKLbqiwSF0FUELI7Bqj8jeWRoO0wH00aQadk7gFqJXXpFCBW9LAPcKy2gmAEassoWE0D4m2BbvcoBb9qOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMB0GA1UdDgQWBBSND3pomfXOYFnGTHHC1o5F03GaczAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQ4NTREMAoGCCqGSM49BAMCA0gAMEUCIEi34uzarhwsiSZ3N05qauJ5my3SDDi0cCnWxzvDWCPVAiEA881SrnNNqmg8Rk3AjbvNxQDVJEVRGrdFFZ6Y2kDgFf0%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 12 Apr 23 22:02 UTC