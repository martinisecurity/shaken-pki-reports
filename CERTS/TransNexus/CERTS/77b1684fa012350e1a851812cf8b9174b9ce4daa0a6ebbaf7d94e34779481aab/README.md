# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 5512

Tested At: 26 Aug 24 17:43 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -602 day(s)\
Subject: CN=SHAKEN 5512, O=Andrew Ward Consulting LLC, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA4, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/19aed76a-a067-4715-8a05-1993cc9d939e/d4987e0e24d4e1a82e88207e55278be3.pem

[View certificate details](https://x509.io/?cert=MIIC2zCCAoCgAwIBAgIQYGF1iYUr9NI05Xic77x%2FHTAKBggqhkjOPQQDAjBWMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEsMCoGA1UEAxMjVHJhbnNOZXh1cywgSW5jLiBTSEFLRU4gSXNzdWluZyBDQTQwHhcNMjIxMjI1MjAyMzQyWhcNMjMwMTAxMjAyMzQxWjBIMQswCQYDVQQGEwJVUzEjMCEGA1UEChMaQW5kcmV3IFdhcmQgQ29uc3VsdGluZyBMTEMxFDASBgNVBAMTC1NIQUtFTiA1NTEyMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEBHmsCtdwlH7ADswXMVRyJbEi0LmI6m%2F6eTGJnP%2BAXIwWrEIRVdpXrCcnjhCnQzBIGePRkF5YqbS0rqAIW%2BNqYqOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMB0GA1UdDgQWBBT7X%2FZR8Y5tTpl0GlnDh4vIIBcesDAfBgNVHSMEGDAWgBQw9fXyt%2BFLCw8QdX1IpJDxPYsoKjAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQ1NTEyMAoGCCqGSM49BAMCA0kAMEYCIQDaiVOls%2BIvqfcejb4MJFdtNRhZhldjQyA8R2g3YLHTegIhAJtt2QO02obeMl%2FhMsRaLPBguuCv3wogG2xSRwnQSKqU)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 26 Aug 24 18:03 UTC