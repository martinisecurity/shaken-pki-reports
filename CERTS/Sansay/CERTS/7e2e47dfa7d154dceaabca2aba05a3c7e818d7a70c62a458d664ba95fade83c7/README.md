# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Televergence Solutions Inc 779J

Tested At: 04 Oct 24 16:07 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -630 day(s)\
Subject: emailAddress=ddeutsch@televergence.com, CN=SHAKEN Televergence Solutions Inc 779J, OU=Production, O=Televergence Solutions Inc, ST=Tennessee, C=US, emailAddress=ddeutsch@televergence.com\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/779J/order/172_779J_73

[View certificate details](https://x509.io/?cert=MIIDGjCCAsGgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkWIIwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTIxNDE1MjMzMFoXDTIzMDExMzE1MjMzMFowgbYxCzAJBgNVBAYTAlVTMRIwEAYDVQQIDAlUZW5uZXNzZWUxIzAhBgNVBAoMGlRlbGV2ZXJnZW5jZSBTb2x1dGlvbnMgSW5jMRMwEQYDVQQLDApQcm9kdWN0aW9uMS8wLQYDVQQDDCZTSEFLRU4gVGVsZXZlcmdlbmNlIFNvbHV0aW9ucyBJbmMgNzc5SjEoMCYGCSqGSIb3DQEJARYZZGRldXRzY2hAdGVsZXZlcmdlbmNlLmNvbTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABI7YJeB2zUiDx0TCzLx53mIokjf1Zw8pZBRXCiyTzdvxaZ9YPFSsOsaqudnIF5bZ9e47vkgQvlGQsZKxR%2Fvi60ujgdswgdgwFgYIKwYBBQUHARoECjAIoAYWBDc3OUowFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMB0GA1UdDgQWBBQAf2WVYtjP2AE%2F07vmJrW95Y60TTAfBgNVHSMEGDAWgBSs05P1Q0PMCr5FWBcTfZJ83MMBRjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDRwAwRAIgaH73aeqOgNbQVV9FYTEatQqd05HLNE7IJ5h6OGTVSTwCIFzuuE3IYeCwJVwS6vH9PlEmKXduQXBKhc%2FCgoiRnPDl)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension contains an invalid OID value: 2.16.840.1.114569.1.1.1. Available OIDs: 2.16.840.1.114569.1.1.3, 2.16.840.1.114569.1.1.4 |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 779J', but common name is 'SHAKEN Televergence Solutions Inc 779J' |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 04 Oct 24 16:29 UTC