# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 674J

Tested At: 21 Nov 23 16:42 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -392 day(s)\
Subject: CN=SHAKEN 674J, OU=SHAKEN, O=Panterra Networks Inc., C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/0b98ecf4-0a43-4b89-8fc0-8e029d8258fa/e1722e7aa08be1907fea5bfc847bb891.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC%2BDCCAp6gAwIBAgIQXj5l7lE2sjNx30bwKW8PwjAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMTcyMDE2NTNaFw0yMjEwMjQyMDE2NTJaMFUxCzAJBgNVBAYTAlVTMR8wHQYDVQQKExZQYW50ZXJyYSBOZXR3b3JrcyBJbmMuMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiA2NzRKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE8By7l8gDz2hMsgwRfUBZdiwfCzzt%2BIE6dz8AGBUeIEk65Q27maLqqjN30LFCgNzkO4IdmMnSxPuw0mIoPUMcbqOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBTrXlL2EQQGOWpUhaWE2h9jicJDTjAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQ2NzRKMAoGCCqGSM49BAMCA0gAMEUCIQDJMc3b%2BM%2ByxXFPV9k%2BSqLgNFdI7Mu1rnt%2BJo2iIHAbJQIgdLFhyDffjDIZUM6Si7TPzFI52htq6RNVQlM%2FocnFUKw%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 21 Nov 23 17:16 UTC