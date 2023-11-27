# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 804J

Tested At: 27 Nov 23 23:20 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 66 day(s)\
Subject: CN=SHAKEN 804J, O=QuestBlue Systems Inc, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN CA-2, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US\
Link: https://customer.questblue.com/assets/questblue_shaken.cer

[View certificate details](https://understandingwebpki.com/?cert=MIIDCDCCAq%2BgAwIBAgIUCFfomTOnq8K2rEYMP47tAvZ1PDQwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0yMB4XDTIzMDIwMTIwMzYzNVoXDTI0MDIwMTIwMzYzNVowQzELMAkGA1UEBhMCVVMxHjAcBgNVBAoMFVF1ZXN0Qmx1ZSBTeXN0ZW1zIEluYzEUMBIGA1UEAwwLU0hBS0VOIDgwNEowWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAT6Ttc%2Bw4ZtoCxD4AGf1VOHTqazkgGtQLHiMvCGz9h%2BGW9mBwQaaheP3%2B%2F7MWwiPHIaJ%2BitmUKdp7BRr83H1nRJo4IBPDCCATgwFgYIKwYBBQUHARoECjAIoAYWBDgwNEowDAYDVR0TAQH%2FBAIwADAfBgNVHSMEGDAWgBSCThX%2F%2Fp9qZ1HkyJfHbXTD%2FGVFeDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMB0GA1UdDgQWBBSm0XsIGDnSAWMiSestnsmN9wIJgDAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDRwAwRAIgbY55ofW8gtqHviEtTWaNX6SXoGxj%2BH5Lu0iMBqBcYs4CICc7XiOQlIOyRDSPMGJFgiLfaJNLdYU4RQLNaSEtSBXT)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 27 Nov 23 23:28 UTC