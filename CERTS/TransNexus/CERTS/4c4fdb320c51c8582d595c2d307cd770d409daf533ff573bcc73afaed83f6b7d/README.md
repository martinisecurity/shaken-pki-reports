# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 866J

Tested At: 21 Aug 23 20:07 UTC\
Initial Validity Period: 180 day(s)\
Remaining Validity Period: -183 day(s)\
Subject: CN=SHAKEN 866J, OU=SHAKEN, O=VocalIP, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.transnexus.com/866J/6442b809-7dbd-4be7-9b98-a8ec96d005ea.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC6TCCAo%2BgAwIBAgIQTRPJ8rR5%2Fa2wT6EvKQBPTzAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjA4MjMwNDE0MDdaFw0yMzAyMTkwNDE0MDZaMEYxCzAJBgNVBAYTAlVTMRAwDgYDVQQKEwdWb2NhbElQMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiA4NjZKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEeSbyMt2Os6iqVProvWl8jDa1tF7PJ5dFKps1N9OLD%2Fg4MdWBU7Li130Yg2JgHjnT%2F%2BP7OH1%2BEPAohHCS1ucB%2BKOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBQfxuvFTmEGo9dkNk%2BZfQXJsSzU5zAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQ4NjZKMAoGCCqGSM49BAMCA0gAMEUCIQCatIuK9%2B1MGGb7XCjBDRj%2FMSfo8LqvKg85%2Frf%2BiunB8wIgEury5Zd%2FLZhBzvMmnJ%2FCq9CwtgBLM9RWXEs1%2BAQsmDA%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |


Generated: 21 Aug 23 20:18 UTC