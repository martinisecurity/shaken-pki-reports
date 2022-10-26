# STIR/SHAKEN CA Ecosystem Compliance
## TransNexus

### Certificate 54c7bf84f9870adf34d3624c71a4b8224436f994
Tested At: 2022-10-26 21:14:15 +0000 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: 1 day(s)\
Subject: CN=SHAKEN 854D, OU=SHAKEN, O=Xtel, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US

Link: https://certificates.clearip.com/627e4253-759f-447e-b1a0-c9cda9a6d597/eef13829c4b0f7d34e0feab46b71acf9.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIC5TCCAoygAwIBAgIQRkNzAZLLOYZu0xX4S09AyjAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMjAyMDIxMTBaFw0yMjEwMjcyMDIxMDlaMEMxCzAJBgNVBAYTAlVTMQ0wCwYDVQQKEwRYdGVsMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiA4NTREMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEiuLVU3DpNb0fOWTa4UGLtxONfpAN5VIP1R66PkNcb94o%2Bwcj1TIQre9RHWPnmChp%2BU6gH2K1rbM%2BUKlG2YN4bKOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBSYLUW4i5Gvana0UbBDYkZvsrOrKTAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQ4NTREMAoGCCqGSM49BAMCA0cAMEQCIGzQyAITUDOvLsdw6oeuu%2Btoq7Igvxd%2F9VTehyeP8fiIAiBZQJ9nBoKgrt71mo%2FjJwZvFryV3TC7QxlSVyrnqBhD4Q%3D%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_incorrect_ku_encoding | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| e_sti_crl_distribution | error | ATIS-1000080v4 | CRL Distribution Point shall be reachable if the requesting IP address within the program ACLs |
| w_cp1_3_subject_rdn_unknown | warn | CPv1.3 | Only CN, C, O, and SERIALNUMBER can be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| e_cp1_3_subject_sn | error | CPv1.3 | STI certificate shall include a ‘serialNumber’ attribute along with the CN |

Generated: 26/10/2022 at 21:14:23