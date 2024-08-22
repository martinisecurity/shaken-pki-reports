# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 130B

Tested At: 22 Aug 24 15:21 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -613 day(s)\
Subject: CN=SHAKEN 130B, O=Global Telecom, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA4, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/4efae05c-c8cd-405b-ae97-c62a52afd074/c77e96e427f49afbcdd42f48d4458726.pem

[View certificate details](https://x509.io/?cert=MIICzzCCAnSgAwIBAgIQSEg2E0CaOLf4xk9gwry57jAKBggqhkjOPQQDAjBWMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEsMCoGA1UEAxMjVHJhbnNOZXh1cywgSW5jLiBTSEFLRU4gSXNzdWluZyBDQTQwHhcNMjIxMjExMTIxNjEzWhcNMjIxMjE4MTIxNjEyWjA8MQswCQYDVQQGEwJVUzEXMBUGA1UEChMOR2xvYmFsIFRlbGVjb20xFDASBgNVBAMTC1NIQUtFTiAxMzBCMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE9ra%2B3zMDgF7eagwbXhtj0lByg07Ez7hsUXmqUGxfjlymYpb1yfj4xzzcNRkbSv67zQzK1sCYSPzkyLH6cYAW%2FqOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMB0GA1UdDgQWBBRa86YMYxDdWs8O8e2qCYjwACTbGDAfBgNVHSMEGDAWgBQw9fXyt%2BFLCw8QdX1IpJDxPYsoKjAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQxMzBCMAoGCCqGSM49BAMCA0kAMEYCIQCPeQ1aOOpK52Cua4ASYeEL69uwjLvmKQgxDAKULN2VtwIhAKLjsZ3gL2qnYLaysFxtZi8UPyhfrOKQsZAosf6CUgRH)

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