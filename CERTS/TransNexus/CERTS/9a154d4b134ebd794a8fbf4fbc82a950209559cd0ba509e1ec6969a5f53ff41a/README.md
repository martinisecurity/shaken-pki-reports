# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 5512

Tested At: 04 Oct 24 15:33 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -626 day(s)\
Subject: CN=SHAKEN 5512, O=Andrew Ward Consulting LLC, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA4, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/19aed76a-a067-4715-8a05-1993cc9d939e/3acbf5c1f7876be2477ca72021eb11a8.pem

[View certificate details](https://x509.io/?cert=MIIC2jCCAoCgAwIBAgIQRSyRS226Rp%2F0sFdWFrrorzAKBggqhkjOPQQDAjBWMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEsMCoGA1UEAxMjVHJhbnNOZXh1cywgSW5jLiBTSEFLRU4gSXNzdWluZyBDQTQwHhcNMjMwMTA5MjAyNTEyWhcNMjMwMTE2MjAyNTExWjBIMQswCQYDVQQGEwJVUzEjMCEGA1UEChMaQW5kcmV3IFdhcmQgQ29uc3VsdGluZyBMTEMxFDASBgNVBAMTC1NIQUtFTiA1NTEyMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEfqy2%2BDa24GKhDoPpUQPIcbbT6VXLUJycdljWiVwCb7yCqZXWWcYONeGTsl7xs3wMtfkzOavsMqCC9iHNNLtGfaOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMB0GA1UdDgQWBBQSV6Wu0xiNgnHw%2FGNHjScPDYr7AjAfBgNVHSMEGDAWgBQw9fXyt%2BFLCw8QdX1IpJDxPYsoKjAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQ1NTEyMAoGCCqGSM49BAMCA0gAMEUCIBojELcI4KzsQeRipBcM5S9UEeuRgOO6ycdWIZL4uDUwAiEA8JnhSEuT0ng1cEf2pQA33nPaFNHg9sYnhitx3HdN460%3D)

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