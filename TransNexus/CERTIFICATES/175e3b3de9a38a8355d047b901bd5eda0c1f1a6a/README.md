# STIR/SHAKEN CA Ecosystem Compliance
## TransNexus

### Certificate 175e3b3de9a38a8355d047b901bd5eda0c1f1a6a
Tested At: 2022-10-27 18:55:43 +0000 UTC\
Initial Validity Period: 90 day(s)\
Remaining Validity Period: 83 day(s)\
Subject: CN=SHAKEN 622J, OU=SHAKEN, O=Skye Telecom LLC, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US

Link: https://certificates.transnexus.com/622J/e06fa5f9-05dc-4e5f-894b-372cfe982864.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIC8jCCApigAwIBAgIQZJ7deBaWM8lMhTNSOo9DpDAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMjAxNTQ4MjNaFw0yMzAxMTgxNTQ4MjJaME8xCzAJBgNVBAYTAlVTMRkwFwYDVQQKExBTa3llIFRlbGVjb20gTExDMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiA2MjJKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEGkO7DK50s%2Br5pjVRW5y86tdLQ1KISgzvt%2FTmk8tlEkYthsK6MUj65VSCD%2FXSLp9eUav6Q%2FbWEQ2PqRun1CpAyKOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBQIPv3owt%2FR9tjM7AUQ13aZ9vbJjzAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQ2MjJKMAoGCCqGSM49BAMCA0gAMEUCIQDOq4jZHcCJw8VjxA%2BCVBPCwa0Hxbfa1HkrY1OcG4vc9wIgFkz038UFHt9IVxnhiAj1txvpG%2BApDvwH0PmRvV05v8c%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_incorrect_ku_encoding | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| w_cp1_3_subject_rdn_unknown | warn | CPv1.3 | Only CN, C, O, and SERIALNUMBER can be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| e_sti_crl_distribution | error | ATIS-1000080v4 | CRL Distribution Point shall be reachable if the requesting IP address within the program ACLs |
| e_cp1_3_subject_sn | error | CPv1.3 | STI certificate shall include a ‘serialNumber’ attribute along with the CN |

Generated: 27/10/2022 at 18:57:27