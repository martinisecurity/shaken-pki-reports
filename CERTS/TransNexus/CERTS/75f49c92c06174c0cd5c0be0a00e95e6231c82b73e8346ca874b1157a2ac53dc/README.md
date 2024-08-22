# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 130B

Tested At: 22 Aug 24 15:21 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -628 day(s)\
Subject: CN=SHAKEN 130B, O=Global Telecom, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA4, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/4efae05c-c8cd-405b-ae97-c62a52afd074/95dcb4fb492ca712a4f54d03522f372e.pem

[View certificate details](https://x509.io/?cert=MIICzzCCAnSgAwIBAgIQSy2cpNat5LXcOfDy75s%2BwDAKBggqhkjOPQQDAjBWMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEsMCoGA1UEAxMjVHJhbnNOZXh1cywgSW5jLiBTSEFLRU4gSXNzdWluZyBDQTQwHhcNMjIxMTI2MTIxNTE0WhcNMjIxMjAzMTIxNTEzWjA8MQswCQYDVQQGEwJVUzEXMBUGA1UEChMOR2xvYmFsIFRlbGVjb20xFDASBgNVBAMTC1NIQUtFTiAxMzBCMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE7lQauaJPTdaY1aotGefgCBURfeWP%2B5NIkDItzkEiIYEQXtqKqHBthNjzlLf%2Fz93hEcmkvgTiSDeWvldukLkbIaOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMB0GA1UdDgQWBBQcr0JWtE%2FA1W1tpLtfJNKw72VOxTAfBgNVHSMEGDAWgBQw9fXyt%2BFLCw8QdX1IpJDxPYsoKjAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQxMzBCMAoGCCqGSM49BAMCA0kAMEYCIQDcDhSgWEAh0p78pfwhHM63222xqh3me9RhwLqjCvNwcwIhAMeryzHdpZATXsZEpvqjSnns6um2KDDpUvVIxFMKu6s3)

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