# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 790J

Tested At: 11 Jan 23 21:48 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -64 day(s)\
Subject: CN=SHAKEN 790J, OU=SHAKEN, O=Viirtue, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/a429df3a-62d2-4851-90c8-fc446859fb08/a6eb57dd1b17f42d3931027cb652bb53.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC6DCCAo%2BgAwIBAgIQX9oU0YVtNoHk%2Bo4iZpZnSTAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjExMDEyMDIzMjlaFw0yMjExMDgyMDIzMjhaMEYxCzAJBgNVBAYTAlVTMRAwDgYDVQQKEwdWaWlydHVlMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiA3OTBKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEs1d1t0fmwaGndltAUHg3shyl5K6NHPch%2Bt52%2B2GUZ0P3zOlxAaEpuZC2Yv1z8tr7Y6aPwX3k7YhkSNm5sdDy4KOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMB0GA1UdDgQWBBSejbOSah0Ouc5UJY1jAeQ18FHTKjAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQ3OTBKMAoGCCqGSM49BAMCA0cAMEQCIEHPM4kJfYk72xMwJrrjq5p%2BO4oe5ND1Adz8AppXnjUpAiAdRm2dUDHW%2BAAjy2SsTUDvVKEDKzz6igl%2BVW8UBknR9A%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 11 Jan 23 21:59 UTC