# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 674J

Tested At: 10 Nov 22 23:19 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: 1 day(s)\
Subject: CN=SHAKEN 674J, OU=SHAKEN, O=Panterra Networks Inc., C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/0b98ecf4-0a43-4b89-8fc0-8e029d8258fa/7f1d21dc26446452e0caa5d5a5f2accd.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC%2BDCCAp6gAwIBAgIQUwneISngRwd4v8giKZr8LjAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjExMDQyMDE4MTdaFw0yMjExMTEyMDE4MTZaMFUxCzAJBgNVBAYTAlVTMR8wHQYDVQQKExZQYW50ZXJyYSBOZXR3b3JrcyBJbmMuMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiA2NzRKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEmHJZ%2BFkS7xXLbdJB1%2F3dIjDnifYpZDUWwUqm5fTxo97hctXCyuPuOeRXa3MEcIjPkxR70i5lolr1OxWSK6Ux7qOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMB0GA1UdDgQWBBRE9zMRZo1aaq8eDC9jw658LziIDzAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQ2NzRKMAoGCCqGSM49BAMCA0gAMEUCIQCo%2BO61wBz64lsy9M1hKx2R1LU0i1F3nWzBgf%2FXXYLlDwIgVAzXOEe9TCecCvD4b0v7sKUiWQ1Pd%2FbXH3klCO%2FuqhU%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 10 Nov 22 23:30 UTC