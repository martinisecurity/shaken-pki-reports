# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 691J

Tested At: 04 Oct 24 15:39 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -719 day(s)\
Subject: CN=SHAKEN 691J, OU=SHAKEN, O=Toly Digital Networks Inc, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/84c98183-b28e-415a-8dc4-ba02b9ca5418/d99780f1c4a7a1ad60a27b0d675c73f5.pem

[View certificate details](https://x509.io/?cert=MIIC%2BzCCAqGgAwIBAgIQfS9OJHhCirjDGxQfWhoJBjAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMDgyMDIwMDdaFw0yMjEwMTUyMDIwMDZaMFgxCzAJBgNVBAYTAlVTMSIwIAYDVQQKExlUb2x5IERpZ2l0YWwgTmV0d29ya3MgSW5jMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiA2OTFKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEyU38c2IecoBSiXNJ2R%2FNIJzMp4V6PfSeyMjkFsvAndU3GF3z4xvnVOBKI8RmBqRUL5G%2B0YF%2FZb2FJXpws%2BLnqKOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBQxcZE20mpo6MoIx2MaeLepoeXcwzAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQ2OTFKMAoGCCqGSM49BAMCA0gAMEUCIECg%2FUUBD8UXHSAtVkxNUaDu%2BeXmhKuogvWGNsHiWVOIAiEAza7hQzIi4pC5hV8sbCtxbk7sGbhBCRDIaQuVg2g7MmE%3D)

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


Generated: 04 Oct 24 16:29 UTC