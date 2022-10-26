# STIR/SHAKEN CA Ecosystem Compliance
## TransNexus

### Certificate dcfb7c5055b8e4f07f3136116eaf7e354dabb82a
Tested At: 2022-10-26 21:00:56 +0000 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: 4 day(s)\
Subject: CN=SHAKEN 738J, OU=SHAKEN, O=SIP.US LLC, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US

Link: https://certificates.clearip.com/bd65ad46-d349-42b3-b28c-24436475d793/b6580ee349b2c3d8a252f43d6a4463d9.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIC7DCCApKgAwIBAgIQYXSnaE3LhFzAVZ4QE28hHTAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMjMyMDIyNThaFw0yMjEwMzAyMDIyNTdaMEkxCzAJBgNVBAYTAlVTMRMwEQYDVQQKEwpTSVAuVVMgTExDMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiA3MzhKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEtFuaaZunYto66hYumKl5GjvjeM0V8JMq1SHawoybXwGFyH%2Fx6qqj27MIvZeRzMSCMNYIuweCdSxy4pBbqJ0IC6OCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBSdv%2BacAc9abb0th5miHUiXSQ33%2BjAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQ3MzhKMAoGCCqGSM49BAMCA0gAMEUCICiUPjgu7RpSp%2FxYbNEjANJMIaxf7KfxOql%2BTp%2BAgv%2FxAiEA2aXR76VcwiBd2LbS%2FfuJnZF9%2FEjuip3BKqcFGMgaEqI%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_cp1_3_subject_sn | error | CPv1.3 | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| e_incorrect_ku_encoding | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| e_sti_crl_distribution | error | ATIS-1000080v4 | CRL Distribution Point shall be reachable if the requesting IP address within the program ACLs |
| w_cp1_3_subject_rdn_unknown | warn | CPv1.3 | Only CN, C, O, and SERIALNUMBER can be included. Additional RNDs may introduce ambiguity and may not be verifiable |

Generated: 26/10/2022 at 21:01:13