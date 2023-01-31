# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 952J

Tested At: 31 Jan 23 21:39 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -116 day(s)\
Subject: CN=SHAKEN 952J, OU=SHAKEN, O=Syndeo LLC dba Broadvoice, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.transnexus.com/952J/187853bb-a8b0-460f-a3f1-7fcd1d30b142.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC%2BjCCAqGgAwIBAgIQSEElA9TaurJnathh6jLKoDAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjA5MDcxNzU0MjdaFw0yMjEwMDcxNzU0MjZaMFgxCzAJBgNVBAYTAlVTMSIwIAYDVQQKExlTeW5kZW8gTExDIGRiYSBCcm9hZHZvaWNlMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiA5NTJKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE3Gmx9mB1ceQ0flaVd0e2%2B3vVYJohDjfMcBQ78BkfnrN%2F6k3bhK%2Bp0yoobdrTycbuEW%2Bc2jJfxu0wu%2BT8wXZ9iqOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBRU9us831WSAowTNRG1wGrn7nHkCjAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQ5NTJKMAoGCCqGSM49BAMCA0cAMEQCIEwfX%2Fis%2Fe%2BCp0UqQSL9VfZhAHtJuyQcU6M0AXS0WCvJAiBN41hQN8%2FOJ23KrUepJehnyFDJHQpySXggG%2FOVk5MVSw%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |


Generated: 31 Jan 23 21:50 UTC