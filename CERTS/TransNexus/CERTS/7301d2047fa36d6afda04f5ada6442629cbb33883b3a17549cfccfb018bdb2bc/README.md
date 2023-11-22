# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 060K

Tested At: 22 Nov 23 03:17 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -408 day(s)\
Subject: CN=SHAKEN 060K, OU=SHAKEN, O=Telware, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/0c4e235f-3e3e-4dd3-bbfc-2f15badab180/0962266d3fd2cc4e4b006b3ea70b8519.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC6TCCAo%2BgAwIBAgIQWUFCQ2UtU6tPZhRdBFx3VTAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMDIxODExNDZaFw0yMjEwMDkxODExNDVaMEYxCzAJBgNVBAYTAlVTMRAwDgYDVQQKEwdUZWx3YXJlMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiAwNjBLMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEySd5z%2FrxKkta9OCIqZp%2B4RX2u9U5piYuFgwSlGNBMbARn88APXN1R0GPRKgj%2BuJg5yC%2BwnC4kVeYJ6%2B0gYG17qOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBTNXZwNnbD6th8UR%2BNkC3pRjguI3zAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQwNjBLMAoGCCqGSM49BAMCA0gAMEUCIQD5WQXeDWgmDnlCYt2ozotS5BDs%2BaKzcgzisw6tyjfvdgIgbEtD3YTTqg8R7%2BLhxR9PmGek3%2FsuUQSyAZIHJ41Ds7c%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 22 Nov 23 03:38 UTC