# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 674J

Tested At: 01 Nov 22 20:29 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: 1 day(s)\
Subject: CN=SHAKEN 674J, OU=SHAKEN, O=Panterra Networks Inc., C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/0b98ecf4-0a43-4b89-8fc0-8e029d8258fa/bd16518450fdc9b0e55bc59cd934697c.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIC%2BTCCAp6gAwIBAgIQeoReV0y8Yp7uZqauLwSHAjAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMjYyMDE3NTBaFw0yMjExMDIyMDE3NDlaMFUxCzAJBgNVBAYTAlVTMR8wHQYDVQQKExZQYW50ZXJyYSBOZXR3b3JrcyBJbmMuMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiA2NzRKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEB9vqaVqB36X7oLDfYly%2BmQkzKB3dDyFUeZksJ5suEXwSOIdTy6Z%2FW%2Buv%2F2Z3nJkvjTfEOvsiAKWLvkWe9566haOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBQTttZSpUY7qZw%2FnakwNqY5C5u39jAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQ2NzRKMAoGCCqGSM49BAMCA0kAMEYCIQCBTgb1YRyHrM2sMezhSOL4U4HR8tcU4%2FqVWHeGr%2FSYEwIhAO3wJJ9snY%2BmGbLxkdCpMpZF6QRWP0%2BIHfnbzFcaay8a)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |


Generated: 01/11/2022 at 20:31:14