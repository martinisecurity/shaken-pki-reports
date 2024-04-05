# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 345J

Tested At: 05 Apr 24 18:45 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -396 day(s)\
Subject: CN=SHAKEN 345J, O=Ooma Inc, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA4, O=TransNexus\\, Inc., C=US\
Link: https://certificates.transnexus.com/345J/a43df45d-6e8c-4968-a8fd-ab97612aeeac.pem

[View certificate details](https://x509.io/?cert=MIICyDCCAm6gAwIBAgIQQV1wj29jCelyrzySn491FTAKBggqhkjOPQQDAjBWMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEsMCoGA1UEAxMjVHJhbnNOZXh1cywgSW5jLiBTSEFLRU4gSXNzdWluZyBDQTQwHhcNMjMwMjA0MDQzMDEzWhcNMjMwMzA2MDQzMDEyWjA2MQswCQYDVQQGEwJVUzERMA8GA1UEChMIT29tYSBJbmMxFDASBgNVBAMTC1NIQUtFTiAzNDVKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEWxekLaNYBcsx3NsSllxAG3EMghyIgv8GEqDN6spjLzWMOOdQvDZcxqoL7m6C%2Br%2Byp4ooChAaW4K3qiPVWAtvMaOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMB0GA1UdDgQWBBTkYxDyD%2BOk5OQEJl6d4uhMh%2BYPkzAfBgNVHSMEGDAWgBQw9fXyt%2BFLCw8QdX1IpJDxPYsoKjAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQzNDVKMAoGCCqGSM49BAMCA0gAMEUCIDryUuwkFS6ALtxPr04DXqxuVTJ0gZeLH4Lp1uWKluBnAiEA272AnZ35VxYnS0yYTo1%2FfCXagKikd81BKQn6RwgbdNs%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 05 Apr 24 19:04 UTC