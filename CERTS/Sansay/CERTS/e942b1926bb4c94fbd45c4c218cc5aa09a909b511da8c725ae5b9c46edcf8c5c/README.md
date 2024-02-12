# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN IPSBS Managed Services LLC 828J

Tested At: 12 Feb 24 16:47 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -394 day(s)\
Subject: emailAddress=iconectiv-hmc@hostmycalls.com, CN=SHAKEN IPSBS Managed Services LLC 828J, OU=NOC, O=IPSBS Managed Services LLC, ST=Tennessee, C=US, emailAddress=iconectiv-hmc@hostmycalls.com\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/828J/order/151_828J_86

[View certificate details](https://understandingwebpki.com/?cert=MIIDGDCCAr6gAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkWJkwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTIxNDE3NTkwNloXDTIzMDExMzE3NTkwNlowgbMxCzAJBgNVBAYTAlVTMRIwEAYDVQQIDAlUZW5uZXNzZWUxIzAhBgNVBAoMGklQU0JTIE1hbmFnZWQgU2VydmljZXMgTExDMQwwCgYDVQQLDANOT0MxLzAtBgNVBAMMJlNIQUtFTiBJUFNCUyBNYW5hZ2VkIFNlcnZpY2VzIExMQyA4MjhKMSwwKgYJKoZIhvcNAQkBFh1pY29uZWN0aXYtaG1jQGhvc3RteWNhbGxzLmNvbTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABA0oRMeHJ2d9TAMrJ28TK2kM3c4mxC%2FnlbB%2BBtN91rEkHvGYc3HSWy5ngDZYSqbxot%2BKwGJXeG8fU2HeFUO3MT%2BjgdswgdgwFgYIKwYBBQUHARoECjAIoAYWBDgyOEowFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMB0GA1UdDgQWBBRC8rHWbSvVLaSU37nCkTHvCmHvFzAfBgNVHSMEGDAWgBSs05P1Q0PMCr5FWBcTfZJ83MMBRjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSAAwRQIgXkJT3zikFhQHWBvMsa%2BthJtlRP1%2FVYvdYp9XocbzQtwCIQCcRPIN2oZKEiSK3QjoqmQShqQGBz8lPr%2B4gS7jBNSe%2Fg%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 828J', but common name is 'SHAKEN IPSBS Managed Services LLC 828J' |
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension contains an invalid OID value: 2.16.840.1.114569.1.1.1. Available OIDs: 2.16.840.1.114569.1.1.3, 2.16.840.1.114569.1.1.4 |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 12 Feb 24 17:02 UTC