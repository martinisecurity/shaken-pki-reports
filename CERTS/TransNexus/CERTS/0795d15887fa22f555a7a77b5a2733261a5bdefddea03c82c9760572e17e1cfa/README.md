# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 130B

Tested At: 26 Aug 24 18:04 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -533 day(s)\
Subject: CN=SHAKEN 130B, O=Global Telecom, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA4, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/4efae05c-c8cd-405b-ae97-c62a52afd074/ca17abdd4b92c0cfdb778fe7d248d855.pem

[View certificate details](https://x509.io/?cert=MIICzjCCAnSgAwIBAgIQclZYYCDc6R5IqvFseRFpVzAKBggqhkjOPQQDAjBWMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEsMCoGA1UEAxMjVHJhbnNOZXh1cywgSW5jLiBTSEFLRU4gSXNzdWluZyBDQTQwHhcNMjMwMzA1MTIyMTI1WhcNMjMwMzEyMTIyMTI0WjA8MQswCQYDVQQGEwJVUzEXMBUGA1UEChMOR2xvYmFsIFRlbGVjb20xFDASBgNVBAMTC1NIQUtFTiAxMzBCMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEhgqizVe%2B7839AvSAP7OW5QiIB4GvQNp%2Fg3EQSgeiiGH4g%2F7yHDs10NWIAafHIExjUiIKCLpo0hmbtR4KYDDSj6OCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMB0GA1UdDgQWBBRufnLrv2XDCpZ%2BPjMM%2BQQr1dTxGzAfBgNVHSMEGDAWgBQw9fXyt%2BFLCw8QdX1IpJDxPYsoKjAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQxMzBCMAoGCCqGSM49BAMCA0gAMEUCICdeyjCGHF9uK7Pi73hY99c5LXAUKu%2Fpq6oXA%2FzGJZ1TAiEAkbg394PMWW52CDpt22i3AC1oAXQKM1wUTGVOb6vVX8w%3D)

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