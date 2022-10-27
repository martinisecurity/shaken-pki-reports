# STIR/SHAKEN CA Ecosystem Compliance
## TransNexus

### Certificate ff10d41f263188f0218fa8d9f25b21d42e3fe1f7
Tested At: 2022-10-27 18:57:26 +0000 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: 4 day(s)\
Subject: CN=SHAKEN 674J, OU=SHAKEN, O=Panterra Networks Inc., C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US

Link: https://certificates.clearip.com/0b98ecf4-0a43-4b89-8fc0-8e029d8258fa/c965b971d4b3e1371cd7155cfb0b07d5.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIC%2BDCCAp6gAwIBAgIQVo5RBQP6AI7e6JcVe7eeaDAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMjMyMDE3MjVaFw0yMjEwMzAyMDE3MjRaMFUxCzAJBgNVBAYTAlVTMR8wHQYDVQQKExZQYW50ZXJyYSBOZXR3b3JrcyBJbmMuMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiA2NzRKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE%2Bdc5i5WnarW8mFITuHFZkFUFebdA38roNgXT2J8UmzDqiM00TNAsa6E15IL7TEKQMIPkccdUy7SfQBXtC5aQvKOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBQJt8CV36hdV2VVzYoNamtMJdtIATAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQ2NzRKMAoGCCqGSM49BAMCA0gAMEUCIE495wcqiSXdhHns9wUuDjaGRWtr1tp2w%2Fawca5Eut%2F8AiEA1ZlHzc62BaCcv0M6lEyc8%2FYkY8%2B3UsAB9mEGrW3obxQ%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_sti_crl_distribution | error | ATIS-1000080v4 | CRL Distribution Point shall be reachable if the requesting IP address within the program ACLs |
| e_cp1_3_subject_sn | error | CPv1.3 | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| e_incorrect_ku_encoding | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| w_cp1_3_subject_rdn_unknown | warn | CPv1.3 | Only CN, C, O, and SERIALNUMBER can be included. Additional RNDs may introduce ambiguity and may not be verifiable |

Generated: 27/10/2022 at 18:57:27