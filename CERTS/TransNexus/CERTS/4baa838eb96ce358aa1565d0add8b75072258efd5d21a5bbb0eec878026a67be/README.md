# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 597J

Tested At: 01 Dec 22 19:19 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 191 day(s)\
Subject: CN=SHAKEN 597J, OU=SHAKEN, O=Voxtelesys, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.transnexus.com/597J/6572b8e4-01de-48c3-9b11-d21cd134c569.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC7DCCApKgAwIBAgIQZ1EczwN4SdxPwFo%2FrAW29zAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjA2MTAxNDAwMjhaFw0yMzA2MTAxNDAwMjdaMEkxCzAJBgNVBAYTAlVTMRMwEQYDVQQKEwpWb3h0ZWxlc3lzMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiA1OTdKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEWCjpm45qK5%2BK%2B8gdEChXERPnyR1H7H1MTxdS1g5TCg9j%2BO%2Bd2RKEFxwB2nwzhmJJ%2B5AeurXhNf3yv34XkUDSzKOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBQ60T3nX2HAjDadtAf%2B5Ao9STJl7zAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQ1OTdKMAoGCCqGSM49BAMCA0gAMEUCICpo%2FUAeOYa9WYHqC4vCagMyejRzlVqMO7pFj6I88vItAiEAiZ9ww9dZQbis2Ad4NbcLDfJPRfgPv0TxvJP7QLOH8hs%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |


Generated: 01 Dec 22 19:22 UTC