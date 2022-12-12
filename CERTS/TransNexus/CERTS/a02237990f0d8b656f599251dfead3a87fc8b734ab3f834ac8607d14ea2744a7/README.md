# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 672B

Tested At: 12 Dec 22 23:36 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -22 day(s)\
Subject: CN=SHAKEN 672B, OU=SHAKEN, O=Clear Rate, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/dcef9a97-b864-43ac-830b-2290a8d0f002/e0d8a224d979aec6eae697266c6e8846.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC6zCCApKgAwIBAgIQfpCXcfGRPNOYOVmMAzBeKDAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjExMTMyMDI1MThaFw0yMjExMjAyMDI1MTdaMEkxCzAJBgNVBAYTAlVTMRMwEQYDVQQKEwpDbGVhciBSYXRlMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiA2NzJCMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEN8kWOVECJt5qUl9V1y9%2BsLQY9Ima%2FXkuIgtWd40HpayBE79FXYwQ36QBO1slkQb6cRfIfWY7IzllIXsCg0R4XKOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMB0GA1UdDgQWBBTXP2HFdFmNW7Txenc7ep%2BZKjq1GTAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQ2NzJCMAoGCCqGSM49BAMCA0cAMEQCIDZVxQHbipQPD9K6gkK0aIQ1Rh0Yva4fFrOj%2BoGy4qiNAiB%2BI%2FL9u0vTVXxHD%2FpLqs1oHp2xGwg%2FWBmOwhFX3q9%2Flg%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 12 Dec 22 23:45 UTC