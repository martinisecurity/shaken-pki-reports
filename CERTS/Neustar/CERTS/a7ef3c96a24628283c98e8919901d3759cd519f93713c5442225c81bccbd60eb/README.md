# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 844J

Tested At: 26 Aug 24 17:49 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: -185 day(s)\
Subject: CN=SHAKEN 844J, O=Unified Office Inc, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN CA-2, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US\
Link: https://downloads.uotcn.net/certs/uo-shaken-cert-20230222.pem

[View certificate details](https://x509.io/?cert=MIIDBzCCAqygAwIBAgIUHM2wg23Mc%2BRX7MtHZcjyOcpttAkwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0yMB4XDTIzMDIyMjIxNDI1OVoXDTI0MDIyMjIxNDI1OVowQDELMAkGA1UEBhMCVVMxGzAZBgNVBAoMElVuaWZpZWQgT2ZmaWNlIEluYzEUMBIGA1UEAwwLU0hBS0VOIDg0NEowWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASTjCdS%2BlU1U2j1QUSHaWdgf3N8Knn01qsGhUjHEy%2BEMiYNQoO%2FVjxvgVYH8nJ6oRbkubg362V%2FJIYrhSbpPeoCo4IBPDCCATgwFgYIKwYBBQUHARoECjAIoAYWBDg0NEowDAYDVR0TAQH%2FBAIwADAfBgNVHSMEGDAWgBSCThX%2F%2Fp9qZ1HkyJfHbXTD%2FGVFeDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMB0GA1UdDgQWBBRtMf%2BLqCVxr18Tf9pdlxoykckFYjAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSQAwRgIhAIzs4xHWn0Stu0aOXDYq1u%2FSYziKscYVp4%2FKPABhytLMAiEAurb9Xa89W84nbaG88ba7UWaJgVNmQeusEweHA45O9Jo%3D)

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