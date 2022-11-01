# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 857J

Tested At: 01 Nov 22 20:27 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 231 day(s)\
Subject: CN=SHAKEN 857J, OU=SHAKEN, O=IEvolve, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://icomm.i-evolve.net/certs/icomm-shaken-sbc01.crt

View: [Click to view](https://understandingwebpki.com/?cert=MIIC6TCCAo%2BgAwIBAgIQdzrboydYX5iXoWuUdl2cgzAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjA2MjAyMDI1MTBaFw0yMzA2MjAyMDI1MDlaMEYxCzAJBgNVBAYTAlVTMRAwDgYDVQQKEwdJRXZvbHZlMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiA4NTdKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE9FM9t%2FXuNyxHqfNfn1IaSH%2F127YdNSVe78mXSFi5gaAWwE6E3c6rZcbIM67tmZvt3vxNGFFB4k3ZwnPyGX8x8KOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBQto7l0Be2eJ6prBW7uPQdi5Zv%2FRzAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQ4NTdKMAoGCCqGSM49BAMCA0gAMEUCIQDFb%2FOYDPhLwp3sbts9QFlIdUJYB%2F6W%2BoKeHtnsxDAPowIgF2SFHeauc9H66vNFuQMkZklJBqKMjCfj%2BBKozDHtVxU%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |


Generated: 01/11/2022 at 20:34:21