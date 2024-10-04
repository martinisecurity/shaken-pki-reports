# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 983J

Tested At: 04 Oct 24 15:38 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -682 day(s)\
Subject: CN=SHAKEN 983J, O=ESI, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/59ebb7c1-25bd-4dfc-9794-fcb104b2f66a/4eaa641b4d155d8ec6fd489a165d051a.pem

[View certificate details](https://x509.io/?cert=MIIC1TCCAnqgAwIBAgIQRYcVDNGWYOEVJNGV1aUqujAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjExMTQyMDE1MzRaFw0yMjExMjEyMDE1MzNaMDExCzAJBgNVBAYTAlVTMQwwCgYDVQQKEwNFU0kxFDASBgNVBAMTC1NIQUtFTiA5ODNKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEnxYZq7GhKGFhgUU1VIss8GIZupYdvgxiyv9FAbH%2FNTu9zhzxczmdCTuN%2BAfInC%2BN3wSBBEzB7UvTIdOO0vWAUaOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMB0GA1UdDgQWBBSXuPhzUpLdw7Q9QdTDZ%2BLmAwT%2FejAfBgNVHSMEGDAWgBS7lt4xEs3TlpmEpDYwYDzXUoF9JzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQ5ODNKMAoGCCqGSM49BAMCA0kAMEYCIQD%2B7TmQcx0TIXgOei5Nczh%2FcnNvnUiUFq2hup%2FPvsgAxwIhAJw9uW%2FiCqL9yJPyXk1DWqzjAGXpV6viwBdGc8DiHxnO)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 04 Oct 24 16:29 UTC