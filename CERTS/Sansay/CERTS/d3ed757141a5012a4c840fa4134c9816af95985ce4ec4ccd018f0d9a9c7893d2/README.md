# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN IDT America, Corp 363A

Tested At: 24 Nov 23 11:08 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -143 day(s)\
Subject: emailAddress=thomas.musella@idt.net, CN=SHAKEN IDT America\\, Corp 363A, OU=NOC, O=IDT America\\, Corp, ST=New Jersey, C=US, emailAddress=thomas.musella@idt.net\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/363A/order/145_363A_17

[View certificate details](https://understandingwebpki.com/?cert=MIIDADCCAqagAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkczkwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDYwMzE1NTQ1MFoXDTIzMDcwMzE1NTQ1MFowgZsxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApOZXcgSmVyc2V5MRowGAYDVQQKDBFJRFQgQW1lcmljYSwgQ29ycDEMMAoGA1UECwwDTk9DMSYwJAYDVQQDDB1TSEFLRU4gSURUIEFtZXJpY2EsIENvcnAgMzYzQTElMCMGCSqGSIb3DQEJARYWdGhvbWFzLm11c2VsbGFAaWR0Lm5ldDBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABPQM8yPpWKxm0BJ%2BjGRI%2F%2Fn99wUDOlM0pPkwuiTqWy8TLzNxSa1qR7pJYQDwtTuzur7Db7X4HGB5yZxyIiXZWL6jgdswgdgwFgYIKwYBBQUHARoECjAIoAYWBDM2M0EwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMB0GA1UdDgQWBBRtxUb554hu2Cjs4EgzW1ORcFzRVzAfBgNVHSMEGDAWgBSs05P1Q0PMCr5FWBcTfZJ83MMBRjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSAAwRQIgYdoV3IqD%2Fl5f9dhX%2BIe8lAI3weDov%2FkB8EUXK8fo3ToCIQDaUAYz4J%2FdQt%2F6C60eJJ15QHgnukSZN9lg%2BwqCOxTL8A%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_ext_crl_distribution_struct](../../ISSUES/e_atis_ext_crl_distribution_struct/README.md) | error | ATIS1000080 | CRL Distribution Point shall contain a CRLIssuer field |
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 363A', but common name is 'SHAKEN IDT America, Corp 363A' |


Generated: 24 Nov 23 11:17 UTC