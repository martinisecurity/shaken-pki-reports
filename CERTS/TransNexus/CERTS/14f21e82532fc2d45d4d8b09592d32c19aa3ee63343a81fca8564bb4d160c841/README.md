# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 622J

Tested At: 11 Jan 23 20:38 UTC\
Initial Validity Period: 90 day(s)\
Remaining Validity Period: -80 day(s)\
Subject: CN=SHAKEN 622J, OU=SHAKEN, O=Skye Telecom LLC, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.transnexus.com/622J/03cbcfec-3614-4fd4-8c38-495fbae7f25d.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC8jCCApigAwIBAgIQTStGIi6we467NwRrAS%2FoeTAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjA3MjUxNzQzMjhaFw0yMjEwMjMxNzQzMjdaME8xCzAJBgNVBAYTAlVTMRkwFwYDVQQKExBTa3llIFRlbGVjb20gTExDMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiA2MjJKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEKPuteqkw0rR%2BZ4jqyimO3AviGM6ck0cP6VgdgE%2FGpHbYPd9kCm4sunsy%2B2n4zSsuTJ%2F5HJ7amm17czcGFjhI%2FKOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBSGP5A%2BWz8FFHAjLlNPCzq9cgwk0jAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQ2MjJKMAoGCCqGSM49BAMCA0gAMEUCIEe7QMSJ%2FV1vJXQq7t5S1qvRSxZCJpM6A6Pk0JUfRVJeAiEAzOnTzq%2F96NApVLtsv%2FaceMSCgZbBVenqNeXJ%2FP%2F6MBA%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |


Generated: 11 Jan 23 21:04 UTC