# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 738J

Tested At: 04 Oct 24 15:48 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -710 day(s)\
Subject: CN=SHAKEN 738J, OU=SHAKEN, O=SIP.US LLC, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/fa6a5cd3-a6c2-4c98-9995-b87249365948/500c018d64467f387ebd88d27dd86ca0.pem

[View certificate details](https://x509.io/?cert=MIIC7DCCApKgAwIBAgIQavCwdB40BZiFQsyyNQMS8zAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMTcyMDIzMzNaFw0yMjEwMjQyMDIzMzJaMEkxCzAJBgNVBAYTAlVTMRMwEQYDVQQKEwpTSVAuVVMgTExDMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiA3MzhKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAESywgFd%2F5%2BDGgHDClGVDm%2B0Nnqt5%2B4T5JXtFAErIHBJLPuakjjCNMLj991WwhhOsK3RsMj%2F8%2BOip36My6aFHp26OCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBQn4%2BrN4gPwFafEZxT5YSfSUkJRUzAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQ3MzhKMAoGCCqGSM49BAMCA0gAMEUCIQDAKy3dpbu5quMYUfGvd0QRTLEVz%2FNZr%2FuB0ZA6I%2FsoLgIgUl9hdESqDU36XV0vXUfBFCvIMhvaADYmBy1QE5yZweg%3D)

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


Generated: 04 Oct 24 16:29 UTC