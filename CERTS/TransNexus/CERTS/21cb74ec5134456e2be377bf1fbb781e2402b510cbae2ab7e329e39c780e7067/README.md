# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 849J

Tested At: 22 Aug 24 15:22 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -661 day(s)\
Subject: CN=SHAKEN 849J, OU=SHAKEN, O=Fuse.Cloud, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/86e241b8-9c8e-4431-b35f-4d92844a1da9/5964d027be7565d20bd77769e0ddc97b.pem

[View certificate details](https://x509.io/?cert=MIIC6zCCApKgAwIBAgIQc6CudKyTCujPYX%2Fos8G8PDAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMjMyMDIxNDdaFw0yMjEwMzAyMDIxNDZaMEkxCzAJBgNVBAYTAlVTMRMwEQYDVQQKEwpGdXNlLkNsb3VkMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiA4NDlKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAElk9cj2LRp3oJ%2FArF%2Ba2Pt2BAJHPpILXPRgV1udEwXIE6S%2BVtE72P4MMVcVwewJyYb0eIkaDExC9ry4kqBxpDY6OCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBQC2Spkr5dFbEEqj0nif%2BRl6z5NNTAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQ4NDlKMAoGCCqGSM49BAMCA0cAMEQCIAxRGg3IPMqmznOgsGOi3yuR%2B2hEhOpGi5L6GG5%2B0WTkAiAjLg3bwp%2BT86hu%2FPdORz85ZuT9JNMolz2nbyypjzP3pw%3D%3D)

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


Generated: 22 Aug 24 15:44 UTC