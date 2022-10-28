# STIR/SHAKEN CA Ecosystem Compliance
## TransNexus

### Certificate 530153d7d36e6a903758e07a9b38a256b381df80
Tested At: 2022-10-28 19:21:52 +0000 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: 6 day(s)\
Subject: CN=SHAKEN 790J, OU=SHAKEN, O=Viirtue, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US

Link: https://certificates.clearip.com/a429df3a-62d2-4851-90c8-fc446859fb08/615bf0d5dc2548fbf2877af69489316b.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIC6jCCAo%2BgAwIBAgIQQSdqdAUSIRCC%2BsJKBOzphzAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMjYyMDIyNDlaFw0yMjExMDIyMDIyNDhaMEYxCzAJBgNVBAYTAlVTMRAwDgYDVQQKEwdWaWlydHVlMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiA3OTBKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAENy%2BzEHTf8yTIIZxG14sKoULmeEF0nEE2My32O86m%2FCQNl0eJ708QmcAq%2BrWzUIr%2FFmBruV%2BhhUgg8SK0JAgD4qOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBQEr94lHqeb4RnLic1cYYMhPJlv2jAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQ3OTBKMAoGCCqGSM49BAMCA0kAMEYCIQDX%2FyG6E01lWbP4G86uY3QiOEZNT2AuuWi8UOtswtWusgIhAJn0Q4aU92vJmoywGRLEvlk4P%2BSqy5f9d5EYPdUWUVKh)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_cp1_3_subject_sn | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| e_incorrect_ku_encoding | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| w_pki_subject_rdn_unknown | warn | SHAKEN PKI Best Practice | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |

Generated: 28/10/2022 at 19:22:10