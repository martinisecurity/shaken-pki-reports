# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 819H

Tested At: 28 Nov 23 19:53 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -403 day(s)\
Subject: CN=SHAKEN 819H, OU=SHAKEN, O=BluIP Inc., C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/3c4ce448-386b-4d47-a276-7fe32e380a83/7a4d69f0a0ff2d376e851bab1a16f9e5.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC7DCCApKgAwIBAgIQSnWKoMkHzoygrt1t9TNrJzAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMTQyMDE5MThaFw0yMjEwMjEyMDE5MTdaMEkxCzAJBgNVBAYTAlVTMRMwEQYDVQQKEwpCbHVJUCBJbmMuMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiA4MTlIMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEkIaXyBTetJHkLYA7GORN27SfcRggMXd5u3jcFxTXnYmSSGvqN4qbIWqK49c7JjFkBK%2Bx3g2m29vY6%2FKg3EANC6OCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBQAwjIMz1gRFcLzaFEEoeqL9kvVwjAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQ4MTlIMAoGCCqGSM49BAMCA0gAMEUCIQCXTHn9cR%2FsSNBCPDOvYgaBl81jWdhoCiJa9WmMaYI6mQIgVo%2B9imQmcYcsO3sU6GHRasN9xWTwoOpOJT8PCe8GNyU%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 28 Nov 23 20:21 UTC