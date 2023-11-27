# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 672B

Tested At: 27 Nov 23 22:25 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -378 day(s)\
Subject: CN=SHAKEN 672B, OU=SHAKEN, O=Clear Rate, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/dcef9a97-b864-43ac-830b-2290a8d0f002/4498653f30a0807792d650aa85959a40.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC7TCCApKgAwIBAgIQfCcwA9okOVscZCcmO2X4%2BzAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjExMDcyMDI0NTVaFw0yMjExMTQyMDI0NTRaMEkxCzAJBgNVBAYTAlVTMRMwEQYDVQQKEwpDbGVhciBSYXRlMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiA2NzJCMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAETFpDThSkY5gCrgBMjIg4nSeGgIiyDAW1T%2FDMxpaX0sdCJcFuE%2Frbji44%2BBtezO%2BPXVKyaRNAJXPVb4gfAtBDMqOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMB0GA1UdDgQWBBTcmfVDZilJytfubqCgI8Ha%2F%2BfSBjAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQ2NzJCMAoGCCqGSM49BAMCA0kAMEYCIQDVslcqIuPsQmBYROMBRkjZGPlwy0Y%2BIIMN5mmCazM8LAIhANJhGHSOhkM%2FbD%2Fz6QDGtjeTbYXwTDVbffE0J1BD8WJu)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 27 Nov 23 22:56 UTC