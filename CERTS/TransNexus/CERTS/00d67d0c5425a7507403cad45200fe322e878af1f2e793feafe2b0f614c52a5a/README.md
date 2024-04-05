# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 012K

Tested At: 05 Apr 24 18:39 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -403 day(s)\
Subject: CN=SHAKEN 012K, O=CallCurrent\\, Inc., C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA4, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/06ebd24a-1f0a-46d5-8a2f-a7ae49be8ed6/0faa0714f6d21c2b9b69c6030f05e9af.pem

[View certificate details](https://x509.io/?cert=MIIC0jCCAnegAwIBAgIQbny%2Bse9a3M3uYbfO%2B4FSozAKBggqhkjOPQQDAjBWMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEsMCoGA1UEAxMjVHJhbnNOZXh1cywgSW5jLiBTSEFLRU4gSXNzdWluZyBDQTQwHhcNMjMwMjIwMTQ1MDI1WhcNMjMwMjI3MTQ1MDI0WjA%2FMQswCQYDVQQGEwJVUzEaMBgGA1UEChMRQ2FsbEN1cnJlbnQsIEluYy4xFDASBgNVBAMTC1NIQUtFTiAwMTJLMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEHTbfs9DBjvD4fUKOlI9FyZtzt9VabR8Xr1AB1TsZI3A5aYQEz%2Fq2C1JQYrDD5LK4RrdVaLOmHvGZS8JSamyL96OCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMB0GA1UdDgQWBBSa4KUIprKuZlQKo%2Fgg9TU1tZKPcjAfBgNVHSMEGDAWgBQw9fXyt%2BFLCw8QdX1IpJDxPYsoKjAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQwMTJLMAoGCCqGSM49BAMCA0kAMEYCIQDiDV%2F4RWSLAK9FbkEb6jFzfr1h2uTNBeOeAeRAGrpVdQIhAPF8toNMrghHU4koK34rHc5N5AvAh67UzABV7tqSZmHq)

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