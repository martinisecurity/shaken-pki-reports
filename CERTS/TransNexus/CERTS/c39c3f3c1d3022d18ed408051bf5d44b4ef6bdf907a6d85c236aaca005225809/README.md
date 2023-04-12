# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 594J

Tested At: 12 Apr 23 01:02 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -148 day(s)\
Subject: CN=SHAKEN 594J, OU=SHAKEN, O=Carolina Digital, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/14652271-f2ba-4197-87c0-704f8c618e4b/b300b62578044f83350c02da994e07c6.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC8jCCApigAwIBAgIQUC9%2BJq1TIpE5%2F2ZIvEj3UTAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjExMDcyMDE5NDRaFw0yMjExMTQyMDE5NDNaME8xCzAJBgNVBAYTAlVTMRkwFwYDVQQKExBDYXJvbGluYSBEaWdpdGFsMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiA1OTRKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEdCMvQrZv7aNHAF5Jb27J%2BdSg1z61oV2Esf3RAfTwE%2BKMG8VKA7YGb0JwWKWx66%2Fa0jnINSOVSl7Gms2XyaKk66OCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMB0GA1UdDgQWBBS%2BYkhLMUTkKTFt5HO472rWOqLWMDAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQ1OTRKMAoGCCqGSM49BAMCA0gAMEUCIBYVgRN9DKfiTW8PZspHlhD%2FMky0HC95ruaOyKhP0ar7AiEA1OptSfVotQPiTQi7JTaMRqxilWB1gICZeVTkC%2FhZdWc%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 12 Apr 23 01:46 UTC