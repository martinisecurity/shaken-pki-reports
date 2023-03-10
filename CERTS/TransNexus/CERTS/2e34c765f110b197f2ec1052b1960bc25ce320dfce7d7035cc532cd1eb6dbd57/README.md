# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 622J

Tested At: 09 Mar 23 22:59 UTC\
Initial Validity Period: 90 day(s)\
Remaining Validity Period: -50 day(s)\
Subject: CN=SHAKEN 622J, OU=SHAKEN, O=Skye Telecom LLC, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.transnexus.com/622J/e06fa5f9-05dc-4e5f-894b-372cfe982864.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC8jCCApigAwIBAgIQZJ7deBaWM8lMhTNSOo9DpDAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMjAxNTQ4MjNaFw0yMzAxMTgxNTQ4MjJaME8xCzAJBgNVBAYTAlVTMRkwFwYDVQQKExBTa3llIFRlbGVjb20gTExDMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiA2MjJKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEGkO7DK50s%2Br5pjVRW5y86tdLQ1KISgzvt%2FTmk8tlEkYthsK6MUj65VSCD%2FXSLp9eUav6Q%2FbWEQ2PqRun1CpAyKOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBQIPv3owt%2FR9tjM7AUQ13aZ9vbJjzAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQ2MjJKMAoGCCqGSM49BAMCA0gAMEUCIQDOq4jZHcCJw8VjxA%2BCVBPCwa0Hxbfa1HkrY1OcG4vc9wIgFkz038UFHt9IVxnhiAj1txvpG%2BApDvwH0PmRvV05v8c%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 10 Mar 23 02:25 UTC