# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 073H

Tested At: 15 Nov 23 16:15 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: -96 day(s)\
Subject: CN=SHAKEN 073H, OU=SHAKEN, O=Telnyx LCC, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.transnexus.com/073H/d89fe29a-1d12-4fe3-b002-f0dc1c09904d.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC7DCCApKgAwIBAgIQUZ9QFZ6SovNu5d7lPCToMTAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjA4MTAxODExMTFaFw0yMzA4MTAxODExMTBaMEkxCzAJBgNVBAYTAlVTMRMwEQYDVQQKEwpUZWxueXggTENDMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiAwNzNIMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEDYPyrL0cXocoE1RuTrGX7SI0ebm2j%2BnkoQxuWGQqFPLOhsPPc%2BC7S1exDRuDxvyr0ixIGHA2BvANOt5nI8jQt6OCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBSwFpIoIaoBf1TQegaynlJYh0pjljAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQwNzNIMAoGCCqGSM49BAMCA0gAMEUCIAbwAXykWmYMjbLlQy4cM5OukX6D5W%2Fn%2BwajGMiVBjIGAiEAhK5g6wRaMIh6IPY6%2Bm14P7%2FVNc98EiAq6Tz%2FFigCf7A%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 15 Nov 23 17:17 UTC