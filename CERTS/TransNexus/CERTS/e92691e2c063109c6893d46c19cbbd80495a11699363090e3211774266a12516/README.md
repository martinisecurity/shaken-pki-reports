# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 5512

Tested At: 04 Oct 24 15:33 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -632 day(s)\
Subject: CN=SHAKEN 5512, O=Andrew Ward Consulting LLC, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA4, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/19aed76a-a067-4715-8a05-1993cc9d939e/754e42a4204582aefe4baef31b6763d4.pem

[View certificate details](https://x509.io/?cert=MIIC2jCCAoCgAwIBAgIQbUknyjq0clf3CvxBkhUeRjAKBggqhkjOPQQDAjBWMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEsMCoGA1UEAxMjVHJhbnNOZXh1cywgSW5jLiBTSEFLRU4gSXNzdWluZyBDQTQwHhcNMjMwMTAzMjAyNDQ4WhcNMjMwMTEwMjAyNDQ3WjBIMQswCQYDVQQGEwJVUzEjMCEGA1UEChMaQW5kcmV3IFdhcmQgQ29uc3VsdGluZyBMTEMxFDASBgNVBAMTC1NIQUtFTiA1NTEyMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEC3D3Yp5m6uSy3464jWxKqIT78LuTdKrA3QTIGgEQ8iJDgwpB6GGkhM4IJnIIlHiTqSNP2Z5mYKpr%2BpiZpl1yRaOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMB0GA1UdDgQWBBQvAATo5mcDL1Pe%2F1w2IiAVjeZ%2BFDAfBgNVHSMEGDAWgBQw9fXyt%2BFLCw8QdX1IpJDxPYsoKjAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQ1NTEyMAoGCCqGSM49BAMCA0gAMEUCIBfwNK66MIQf%2Fi2wFJOMjifRU9DwC27IudHfmeQvRNs4AiEArn9pasZ1g%2Bv2KTV89IrtL0zJB3q%2BvjgcRYWayk0gqfU%3D)

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