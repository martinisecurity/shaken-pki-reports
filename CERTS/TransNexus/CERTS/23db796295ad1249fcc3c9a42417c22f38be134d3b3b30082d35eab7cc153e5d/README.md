# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 737J

Tested At: 22 Aug 24 16:01 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: -398 day(s)\
Subject: CN=SHAKEN 737J, OU=SHAKEN, O=US Internet Corp, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://sti.ravon.net/c/737J_2022-07-20

[View certificate details](https://x509.io/?cert=MIIC8jCCApigAwIBAgIQdpUZTniQemnsKYAFkMKZuDAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjA3MjAxNzAwMzhaFw0yMzA3MjAxNzAwMzdaME8xCzAJBgNVBAYTAlVTMRkwFwYDVQQKExBVUyBJbnRlcm5ldCBDb3JwMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiA3MzdKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEz0u2mFRDToS9elHZ1YmgSbZLcRrgm0EAgBKx4EBC%2BqbY1oWPvhaGYEgnRWcE6HhJtQB94ifKLkzcCitCPIDtmqOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBRC3uLWcfXMCuVf91n90OzRsf8W3jAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQ3MzdKMAoGCCqGSM49BAMCA0gAMEUCIQC8RC52aqcokdrjr8CQ3RWTCswnFdDo0PDQrCk3aypo7wIgdqwJUWbtN15fjoSnSRIEU7%2BS44DFZRUH%2F%2B5XQsZeCHM%3D)

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


Generated: 22 Aug 24 16:06 UTC