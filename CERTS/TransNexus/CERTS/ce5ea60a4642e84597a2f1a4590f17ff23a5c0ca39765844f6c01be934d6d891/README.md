# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 297K

Tested At: 30 Jan 23 23:00 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -128 day(s)\
Subject: CN=SHAKEN 297K, OU=SHAKEN, O=Clarity Voice, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/859299ca-4c86-4dfd-9785-6c758bee1b37/fc801558e11d85f018c7783bb92429e8.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC7zCCApWgAwIBAgIQasR6AsAYWxLIsx1XsQGFRjAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjA5MTcxOTUyMzZaFw0yMjA5MjQxOTUyMzVaMEwxCzAJBgNVBAYTAlVTMRYwFAYDVQQKEw1DbGFyaXR5IFZvaWNlMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiAyOTdLMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEKrHK6vXqArQLTcfSpbpjc31Tv8TEK6c1BOuya0LE%2B3QUO0ENMAas5tTXrWhwo%2FvRvETjMXiEyjSUM0jMIaHvLqOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBRqYO9OAeLdmmh1PA7JEFsBCOEPZDAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQyOTdLMAoGCCqGSM49BAMCA0gAMEUCICkaPyUh15Bx%2BkLRAJ41uvOYX%2BmjErSDnLcktwmrqt1LAiEAsRHatp4NCW1aW5tyKdnq1WTXcAvF38J4mXUPQBTYZ4E%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |


Generated: 30 Jan 23 23:10 UTC