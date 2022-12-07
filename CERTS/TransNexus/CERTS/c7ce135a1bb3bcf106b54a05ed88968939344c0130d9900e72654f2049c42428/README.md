# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 819H

Tested At: 07 Dec 22 18:45 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -22 day(s)\
Subject: CN=SHAKEN 819H, OU=SHAKEN, O=BluIP Inc., C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/3c4ce448-386b-4d47-a276-7fe32e380a83/4b4864132e0ee2eed5e9ad420537cfb6.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC7DCCApKgAwIBAgIQWFAaigVAzRvotHb5T06PvDAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjExMDcyMDIxMTdaFw0yMjExMTQyMDIxMTZaMEkxCzAJBgNVBAYTAlVTMRMwEQYDVQQKEwpCbHVJUCBJbmMuMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiA4MTlIMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEALHSEO9wKwM7v6AznuQ0HFXJiYiw0JSs7ZLvUpdDNm%2Bt2XwBhR7j6z%2BeAAYhNyseHDlzzjJ6M7%2BabLVVzpz9iaOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMB0GA1UdDgQWBBQZB4O0SdrjwjVBVjHWZAw8rzLlUTAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQ4MTlIMAoGCCqGSM49BAMCA0gAMEUCIDPSmk9cV8TSB3GOJpAci7ilt9RCsvvrw7UGVjKVPnXxAiEA0ZM89u8L89hMIOC%2F8GHIVmUzYlOje7fh3CpdZD%2BI4%2FE%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 07 Dec 22 18:54 UTC