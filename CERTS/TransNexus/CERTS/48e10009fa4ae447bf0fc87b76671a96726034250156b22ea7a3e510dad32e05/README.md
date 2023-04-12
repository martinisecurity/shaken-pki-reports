# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 469A

Tested At: 12 Apr 23 01:02 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -145 day(s)\
Subject: CN=SHAKEN 469A, OU=SHAKEN, O=T3 Communications Inc., C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/88814589-ad88-4c53-b3c3-47f8334afb98/293c1d0ab822f89a973e14854644e325.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC%2BTCCAp6gAwIBAgIQdW%2BBH%2F1NiMH8shvDQ5cosTAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjExMTAyMDIzNThaFw0yMjExMTcyMDIzNTdaMFUxCzAJBgNVBAYTAlVTMR8wHQYDVQQKExZUMyBDb21tdW5pY2F0aW9ucyBJbmMuMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiA0NjlBMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEdxcxapScSjWGrozXEUeR%2BqWFR2ZAIeMNB0Ma754BR800zwbxfBBzNO3lSKo%2Fiupp6DbDu7aphPjpS8M8muuiWKOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMB0GA1UdDgQWBBTPejZ7vatXescduYmVTa6cF4NFODAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQ0NjlBMAoGCCqGSM49BAMCA0kAMEYCIQCkIH2bRlRMqfswcS1yKtky5dJyL1hq%2B%2BXHcdl%2FBrZohAIhAJRQrYPkwN1GGljMy%2BXGBqy%2FeS9pkToP4Zgg2LmRZirB)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 12 Apr 23 01:46 UTC