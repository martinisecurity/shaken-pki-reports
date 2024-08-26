# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 345J

Tested At: 26 Aug 24 18:11 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -560 day(s)\
Subject: CN=SHAKEN 345J, O=Ooma Inc, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA4, O=TransNexus\\, Inc., C=US\
Link: https://certificates.transnexus.com/345J/8aa4dfa2-a9d6-4524-be6b-14b417b5345e.pem

[View certificate details](https://x509.io/?cert=MIICyTCCAm6gAwIBAgIQQib8TanTOOFozuiBPcHJ4TAKBggqhkjOPQQDAjBWMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEsMCoGA1UEAxMjVHJhbnNOZXh1cywgSW5jLiBTSEFLRU4gSXNzdWluZyBDQTQwHhcNMjMwMTE0MDUwMTMxWhcNMjMwMjEzMDUwMTMwWjA2MQswCQYDVQQGEwJVUzERMA8GA1UEChMIT29tYSBJbmMxFDASBgNVBAMTC1NIQUtFTiAzNDVKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEzeDux3ekMhoBTaA9dMk%2F5GM0m8ENdwC8u80GaESEhz2qblQrj4c7kNp0n3iZEevGSQ%2Br5bq1RebsZdRZDW0mdqOCATwwggE4MAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMB0GA1UdDgQWBBSVmNKSHlQwuRPl8pbPG1FZvcboGTAfBgNVHSMEGDAWgBQw9fXyt%2BFLCw8QdX1IpJDxPYsoKjAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMBYGCCsGAQUFBwEaBAowCKAGFgQzNDVKMAoGCCqGSM49BAMCA0kAMEYCIQC44Tl1nkHIhFCbC%2BXZ%2BBHOrTDxzB2LWaWjz2ybn09HDQIhANRq2L5OqpP9yzl%2FpP%2FvCKDb60lvkKan6xOF5R04akpK)

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