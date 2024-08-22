# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 5512

Tested At: 22 Aug 24 15:18 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -631 day(s)\
Subject: CN=SHAKEN 5512, O=Andrew Ward Consulting LLC, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA4, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/19aed76a-a067-4715-8a05-1993cc9d939e/c50040adab7ff3bf0cef7d7123658954.pem

[View certificate details](https://x509.io/?cert=MIIC2jCCAoCgAwIBAgIQcICoMmalyPEzE3AtYor5%2BTAKBggqhkjOPQQDAjBWMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEsMCoGA1UEAxMjVHJhbnNOZXh1cywgSW5jLiBTSEFLRU4gSXNzdWluZyBDQTQwHhcNMjIxMTIyMjAyMDMxWhcNMjIxMTI5MjAyMDMwWjBIMQswCQYDVQQGEwJVUzEjMCEGA1UEChMaQW5kcmV3IFdhcmQgQ29uc3VsdGluZyBMTEMxFDASBgNVBAMTC1NIQUtFTiA1NTEyMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEY6NvRghP%2FVdAS0sN6edNhIhGKQ1evTooANdJxnRm8B%2BJVbeUpK7d%2BirMq5MDPJZ65lBveB9hviexhmmIJaDQiKOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMB0GA1UdDgQWBBQhfqisF7pBRINCCYa7muMBU0VytTAfBgNVHSMEGDAWgBQw9fXyt%2BFLCw8QdX1IpJDxPYsoKjAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQ1NTEyMAoGCCqGSM49BAMCA0gAMEUCIQD%2Bf2QX8%2BA8WA8mTVXDqPdPqLdesAc1Y7Mi6TUinYUgigIgRNVeTcvyQNRpPOy54naiKeyoxtDVCJJ5NF6ci2lkk%2B8%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 22 Aug 24 15:44 UTC