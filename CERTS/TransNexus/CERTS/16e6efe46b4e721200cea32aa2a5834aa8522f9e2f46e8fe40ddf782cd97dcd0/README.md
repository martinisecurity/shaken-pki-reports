# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 819H

Tested At: 21 Nov 23 17:38 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -407 day(s)\
Subject: CN=SHAKEN 819H, OU=SHAKEN, O=BluIP Inc., C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/3c4ce448-386b-4d47-a276-7fe32e380a83/c9df3383702c360f3000ec9968f03d8b.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC7DCCApKgAwIBAgIQYOroziwl1etZ16alc59fAjAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMDIyMDE4MjFaFw0yMjEwMDkyMDE4MjBaMEkxCzAJBgNVBAYTAlVTMRMwEQYDVQQKEwpCbHVJUCBJbmMuMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiA4MTlIMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE2TkjLXoPrikt%2F%2FH0NNpQb1W2epnIJo4lYJidgPDgrcboJjBKMpu5R34DG2ZTLwZgCVRiQ32EB11NF6hggNQel6OCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBSRLcnAg%2FdL0N%2FQ0ktcNdJfUpyHQTAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQ4MTlIMAoGCCqGSM49BAMCA0gAMEUCIBkaC6k9VE6ko4u9PWRrtuM7lVHe8gShmW3XtmHQ6iZJAiEApbraebtUUZBop60by%2B0c4sbNQhlNcyDMe7tT4xXz298%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 21 Nov 23 17:53 UTC