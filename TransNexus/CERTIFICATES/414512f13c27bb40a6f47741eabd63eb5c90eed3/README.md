# STIR/SHAKEN CA Ecosystem Compliance
## TransNexus

### Certificate 414512f13c27bb40a6f47741eabd63eb5c90eed3
Tested At: 2022-10-28 16:27:14 +0000 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 16 day(s)\
Subject: CN=SHAKEN 952J, OU=SHAKEN, O=Syndeo LLC dba Broadvoice, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US

Link: https://certificates.transnexus.com/952J/b81c5f3e-5577-44d2-a22e-5bcfb144b13e.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIC%2FDCCAqGgAwIBAgIQQaPxdUWJnCfxUe8cP%2FV%2BizAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMTMxODAyMzFaFw0yMjExMTIxODAyMzBaMFgxCzAJBgNVBAYTAlVTMSIwIAYDVQQKExlTeW5kZW8gTExDIGRiYSBCcm9hZHZvaWNlMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiA5NTJKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEebmXtP%2FEfuPT%2FhMJDYWpnJXT%2FaoaL7ayWNh2B6WF2T3IQeVq4woR92DZyTghehNheWCq720r3nwSnSq97fjuFKOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBSN3k5gvLleS37Kd%2FO2fdv3B4vA0jAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQ5NTJKMAoGCCqGSM49BAMCA0kAMEYCIQChl4Rqx%2F4SOKQT8IkMCwI7rNtBYz%2BDdhXmvDvnAVAl6QIhAIbIUkF%2FcEJWgJObrZuuxkGHhYAITlHsnO88nrv%2FQYRM)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_incorrect_ku_encoding | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| w_pki_subject_rdn_unknown | warn | SHAKEN PKI Best Practice | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| e_cp1_3_subject_sn | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |

Generated: 28/10/2022 at 16:28:22