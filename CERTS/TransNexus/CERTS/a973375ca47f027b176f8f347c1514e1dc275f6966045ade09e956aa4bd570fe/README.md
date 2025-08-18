# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 130B

Tested At: 18 Aug 25 20:08 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -896 day(s)\
Subject: CN=SHAKEN 130B, O=Global Telecom, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA4, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/4efae05c-c8cd-405b-ae97-c62a52afd074/9ec4f0ba618653e481d95a5e7a3984ea.pem

[View certificate details](https://x509.io/?cert=MIICzzCCAnSgAwIBAgIQb%2BpDAZMBrCHmPHfxBOBbgjAKBggqhkjOPQQDAjBWMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEsMCoGA1UEAxMjVHJhbnNOZXh1cywgSW5jLiBTSEFLRU4gSXNzdWluZyBDQTQwHhcNMjMwMjI3MTIyMDU0WhcNMjMwMzA2MTIyMDUzWjA8MQswCQYDVQQGEwJVUzEXMBUGA1UEChMOR2xvYmFsIFRlbGVjb20xFDASBgNVBAMTC1NIQUtFTiAxMzBCMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEOnsQx6X43XZPoOFkPz%2Bj9iTj0Wo3%2BKGLkCdDpypmjHQ9rJoRZjWpHIPz8Zd%2FIvgcab%2FREi1UV5rXrrGh%2BN4bR6OCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMB0GA1UdDgQWBBSNZtzFC%2BFhxJr8F%2F0z%2B1xmO4GYujAfBgNVHSMEGDAWgBQw9fXyt%2BFLCw8QdX1IpJDxPYsoKjAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQxMzBCMAoGCCqGSM49BAMCA0kAMEYCIQDp6SkodtXyslBphtEWEGSR4%2BGJ8Yg00AJx7kdkAPOkagIhAIYAHdcyjg%2FV5T51%2FHb5BoJlK%2FNuFxrZ8mNiQ3Mvvo9H)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 18 Aug 25 21:13 UTC