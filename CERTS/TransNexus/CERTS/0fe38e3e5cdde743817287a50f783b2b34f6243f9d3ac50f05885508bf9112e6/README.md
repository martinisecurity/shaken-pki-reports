# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 5512

Tested At: 26 Aug 24 18:02 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -623 day(s)\
Subject: CN=SHAKEN 5512, O=Andrew Ward Consulting LLC, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA4, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/19aed76a-a067-4715-8a05-1993cc9d939e/d764483293f04cca61ad4e414fdcf180.pem

[View certificate details](https://x509.io/?cert=MIIC2jCCAoCgAwIBAgIQbL%2B8lj1Zq5bE3Loj1guWzjAKBggqhkjOPQQDAjBWMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEsMCoGA1UEAxMjVHJhbnNOZXh1cywgSW5jLiBTSEFLRU4gSXNzdWluZyBDQTQwHhcNMjIxMjA0MjAyMTE3WhcNMjIxMjExMjAyMTE2WjBIMQswCQYDVQQGEwJVUzEjMCEGA1UEChMaQW5kcmV3IFdhcmQgQ29uc3VsdGluZyBMTEMxFDASBgNVBAMTC1NIQUtFTiA1NTEyMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEH0Dn7iiSAzLTvXSR0Pnp%2FQ%2FegTOwVg2mqtwlQ0nOUrPRHiXwQ%2FOy7FJEhsHQMd3rbhYE6JtVAN%2BVDseIaWTvTaOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMB0GA1UdDgQWBBRBDX%2F0lmWTuJCbn82S0KnWyzWrSzAfBgNVHSMEGDAWgBQw9fXyt%2BFLCw8QdX1IpJDxPYsoKjAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQ1NTEyMAoGCCqGSM49BAMCA0gAMEUCIQD09IeQ4XLSRI%2BTq7R5uv0jx0%2Bd5zPu0mNK%2F4bawi5HYAIgDxVUuzK0OK66vhq%2FejeEjgDFNl%2BXcVJzgUbH0JH%2F89E%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 26 Aug 24 18:49 UTC