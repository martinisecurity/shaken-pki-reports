# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 177K

Tested At: 29 Dec 22 07:35 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -45 day(s)\
Subject: CN=SHAKEN 177K, OU=SHAKEN, O=Cytracom\\, LLC, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/8d118994-4994-4735-ac71-42c0bbb7848f/86828f36b776fa3cf2b88baecaddac8a.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC7zCCApWgAwIBAgIQay%2BO3jyEFnCGBY1UMS9xxzAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjExMDYxNzM4NTBaFw0yMjExMTMxNzM4NDlaMEwxCzAJBgNVBAYTAlVTMRYwFAYDVQQKEw1DeXRyYWNvbSwgTExDMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiAxNzdLMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE%2B6iWH340AqnUm1d%2BS52fTp7QB1o74ilQZ%2BHWjJplw1cjwYtr9FpEWD1lxz7coSeEH3aETALQIhDpz0Q%2FLDP%2Fd6OCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMB0GA1UdDgQWBBRpoVcfiPfduWRmSXc08d6J2uH13jAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQxNzdLMAoGCCqGSM49BAMCA0gAMEUCIQDiNkX2M8THZnRaYviRkpNHgwbUl28k11ykh8kAspzW0AIgdcnTls%2BheEvSS700A9yCdnBTuhzGlpavSwOyKwFgA3Y%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 29 Dec 22 07:47 UTC