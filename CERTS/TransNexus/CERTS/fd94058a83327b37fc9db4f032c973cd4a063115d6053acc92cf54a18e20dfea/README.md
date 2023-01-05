# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 672B

Tested At: 05 Jan 23 20:55 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -49 day(s)\
Subject: CN=SHAKEN 672B, OU=SHAKEN, O=Clear Rate, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/dcef9a97-b864-43ac-830b-2290a8d0f002/8bc01117af6f4610a8a145bb230bb0ad.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC7DCCApKgAwIBAgIQSM3gejklTOr60%2FeVxPhQ7jAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjExMTAyMDI1MTJaFw0yMjExMTcyMDI1MTFaMEkxCzAJBgNVBAYTAlVTMRMwEQYDVQQKEwpDbGVhciBSYXRlMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiA2NzJCMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE0Ody3zngZ2AOl10BGqbVavNW%2BRd4mi6H5qxZE5fuPNix0r8vRD7pixAPQlE3Or6IPOGiCSHvUhKtE%2FZax8xb2aOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMB0GA1UdDgQWBBSsQOSCKGa9CtRS7SerLYBDLYBc6TAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQ2NzJCMAoGCCqGSM49BAMCA0gAMEUCIQC732mYxwSAT37o5uqKTmj2VR3dbjll2qeJQJ%2BX747brQIgEttDV8%2BFwHLf3Bc68AVsGfY%2FrtbfHDXHB5S0Pk8rDtY%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 05 Jan 23 21:05 UTC