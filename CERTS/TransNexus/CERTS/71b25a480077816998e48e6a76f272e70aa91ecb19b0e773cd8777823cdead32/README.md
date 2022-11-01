# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 736J

Tested At: 01 Nov 22 16:28 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 241 day(s)\
Subject: CN=SHAKEN 736J, OU=SHAKEN, O=Masergy Communications, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.transnexus.com/736J/27cf1c16-f0f0-41fa-a0c1-6c167396fe34.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIC%2BDCCAp6gAwIBAgIQWRQ2JbMmoCW2ZxrIW1Oa2zAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjA2MjkyMDI0MzhaFw0yMzA2MjkyMDI0MzdaMFUxCzAJBgNVBAYTAlVTMR8wHQYDVQQKExZNYXNlcmd5IENvbW11bmljYXRpb25zMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiA3MzZKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEYKlJkB%2BmZz1ZOnkBLs9aVrtmGG7BJm0QKG6NbbjOQd0RzU7PF6UGW7T5oSbyFNHoM%2B0n%2Bc8GFCSUEiyMbFGjM6OCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBQcKm6v0ilnb5qwHMq%2BkpidObWlPjAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQ3MzZKMAoGCCqGSM49BAMCA0gAMEUCIQCUvrdfgejPV9Kr6huNG7OByuX8CO%2FYVYmEhrHjqE84cgIgOjoQsGgc7%2Bqfnqlf2C5cr6f145EUxDzoOTyujQ9%2BX%2F0%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN PKI Best Practice | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |


Generated: 01/11/2022 at 16:30:07