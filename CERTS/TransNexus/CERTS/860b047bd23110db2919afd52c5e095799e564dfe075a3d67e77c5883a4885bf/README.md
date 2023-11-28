# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 056J

Tested At: 28 Nov 23 16:00 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -391 day(s)\
Subject: CN=SHAKEN 056J, OU=SHAKEN, O=Telco Experts, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/384afe84-56e7-4f0a-a8f8-9be6e100ab3e/bd57ebc311aedc3a3d6e9e257b5f542c.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC7zCCApWgAwIBAgIQdnCcx0CgkLWEg3tAWoezMDAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMjYxMzQ1MzZaFw0yMjExMDIxMzQ1MzVaMEwxCzAJBgNVBAYTAlVTMRYwFAYDVQQKEw1UZWxjbyBFeHBlcnRzMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiAwNTZKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEQ3rLCF6VzcaxdPMbgC6bLimjaNzXeKMePVsDqsbW2%2F1ZHaCkkMY1IjjeIQFUt92ff%2FWWWgYF7l%2FV7%2BPuacB1aqOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBTUTo%2FL8gjsLWWCeqGB8eJvYovvLDAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQwNTZKMAoGCCqGSM49BAMCA0gAMEUCIQDWuCU%2BpGl9%2FfLQq1W1qhNZrJCGiMLoyxXPsAxcm%2FcY6gIgfJft4HyTRwQc3c1jBSbJzq5Jurf4Wu0VewWlb3Enmm8%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 28 Nov 23 16:15 UTC