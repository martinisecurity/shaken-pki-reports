# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 066K

Tested At: 31 Oct 22 19:20 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 9 day(s)\
Subject: CN=SHAKEN 066K, OU=SHAKEN, O=U Communications, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.transnexus.com/066K/8d407530-1355-4576-8fa1-d7ad1efe4299.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIC8jCCApigAwIBAgIQSkEy6xm%2Bwg1MBbx5vWk%2FjDAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMTAxODM5NDlaFw0yMjExMDkxODM5NDhaME8xCzAJBgNVBAYTAlVTMRkwFwYDVQQKExBVIENvbW11bmljYXRpb25zMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiAwNjZLMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEaHI9NILAkumPsuuHw8qhBrdQvImETSfGUXzr7nBCSNmbEvqal3CwvoK3aey%2FAEsJsFaXzMnE7JPSMuTZb2FmIaOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBQYI3ACpkZisGzy7ZbRzxTO7lXYFTAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQwNjZLMAoGCCqGSM49BAMCA0gAMEUCIQDUtYq%2BeYbgqD7mCDVlBYaWfTbm2tq0vc%2FK33Qyw4WkkgIgVHClvaGBeBYj9d1uxXFasEtZwxOHWF2b0DcWyDVtRgM%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_cp1_3_subject_sn](../../ISSUES/e_cp1_3_subject_sn/README.md) | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_pki_subject_rdn_unknown](../../ISSUES/w_pki_subject_rdn_unknown/README.md) | warn | SHAKEN PKI Best Practice | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |


Generated: 31/10/2022 at 19:21:49