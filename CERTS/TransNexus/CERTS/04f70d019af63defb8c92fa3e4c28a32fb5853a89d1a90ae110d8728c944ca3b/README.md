# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 849J

Tested At: 02 Dec 22 07:44 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -47 day(s)\
Subject: CN=SHAKEN 849J, OU=SHAKEN, O=Fuse.Cloud, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/86e241b8-9c8e-4431-b35f-4d92844a1da9/d976dc1a93579782c232c8eda65e58aa.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC7DCCApKgAwIBAgIQWzdfGhJRdccMw9oReX3QATAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMDgyMDIwMDhaFw0yMjEwMTUyMDIwMDdaMEkxCzAJBgNVBAYTAlVTMRMwEQYDVQQKEwpGdXNlLkNsb3VkMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiA4NDlKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE53aNBGWim9tXFFfTWchDRoZp2P7EnGDVCJfXqNThjcDxWumNbxSsTKGT92xqw84ziQQ7QDSzasMlSDpuu%2BwW86OCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBTcCBp3KyQye0bNa7nTdkMhWqx3hDAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQ4NDlKMAoGCCqGSM49BAMCA0gAMEUCIQDEXG8kdv8KDt%2BR%2FLlXQy%2Ff4iWgjbk7J0XiQ8q4S9cWGAIgFm%2Frgxhxm9BVNcoCKZUqdHxZ3spF7FnruL6ZWsm85LM%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 02 Dec 22 07:46 UTC