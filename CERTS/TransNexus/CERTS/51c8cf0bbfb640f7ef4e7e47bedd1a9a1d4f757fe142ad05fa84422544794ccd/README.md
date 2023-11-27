# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 066K

Tested At: 27 Nov 23 22:26 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -383 day(s)\
Subject: CN=SHAKEN 066K, OU=SHAKEN, O=U Communications, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.transnexus.com/066K/8d407530-1355-4576-8fa1-d7ad1efe4299.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC8jCCApigAwIBAgIQSkEy6xm%2Bwg1MBbx5vWk%2FjDAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMTAxODM5NDlaFw0yMjExMDkxODM5NDhaME8xCzAJBgNVBAYTAlVTMRkwFwYDVQQKExBVIENvbW11bmljYXRpb25zMQ8wDQYDVQQLEwZTSEFLRU4xFDASBgNVBAMTC1NIQUtFTiAwNjZLMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEaHI9NILAkumPsuuHw8qhBrdQvImETSfGUXzr7nBCSNmbEvqal3CwvoK3aey%2FAEsJsFaXzMnE7JPSMuTZb2FmIaOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgCAMB0GA1UdDgQWBBQYI3ACpkZisGzy7ZbRzxTO7lXYFTAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQwNjZLMAoGCCqGSM49BAMCA0gAMEUCIQDUtYq%2BeYbgqD7mCDVlBYaWfTbm2tq0vc%2FK33Qyw4WkkgIgVHClvaGBeBYj9d1uxXFasEtZwxOHWF2b0DcWyDVtRgM%3D)

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


Generated: 27 Nov 23 22:56 UTC