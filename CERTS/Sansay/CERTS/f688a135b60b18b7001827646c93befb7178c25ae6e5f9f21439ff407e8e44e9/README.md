# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN IP Link Telecom Inc. 902J

Tested At: 28 Nov 23 10:42 UTC\
Initial Validity Period: 90 day(s)\
Remaining Validity Period: -236 day(s)\
Subject: emailAddress=ops@iplinktelecom.com, CN=SHAKEN IP Link Telecom Inc. 902J, OU=IP Link Telecom Inc., O=IP Link Telecom Inc., ST=Oregon, C=US, emailAddress=ops@iplinktelecom.com\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/902J/order/46_902J_68

[View certificate details](https://understandingwebpki.com/?cert=MIIDEjCCArigAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkWxcwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDEwNTIyNDgzOFoXDTIzMDQwNTIyNDgzOFowga0xCzAJBgNVBAYTAlVTMQ8wDQYDVQQIDAZPcmVnb24xHTAbBgNVBAoMFElQIExpbmsgVGVsZWNvbSBJbmMuMR0wGwYDVQQLDBRJUCBMaW5rIFRlbGVjb20gSW5jLjEpMCcGA1UEAwwgU0hBS0VOIElQIExpbmsgVGVsZWNvbSBJbmMuIDkwMkoxJDAiBgkqhkiG9w0BCQEWFW9wc0BpcGxpbmt0ZWxlY29tLmNvbTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABL6w%2BsRjK9RdJbKQ%2BOomb3GbohUU8NOMYdwk2Y7CNjoxoM%2BsiLVPAGbeHQ4P4VyqtdOeC5PQkhtlYjHzM01ZEO%2BjgdswgdgwFgYIKwYBBQUHARoECjAIoAYWBDkwMkowFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMB0GA1UdDgQWBBSxLw4TUOmdUNn5AKb5bj5RA6qnWDAfBgNVHSMEGDAWgBSs05P1Q0PMCr5FWBcTfZJ83MMBRjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSAAwRQIhAOTzGcbBZNnJ8CssNJdP5ygt8p3v8ZQP65ElnC%2BCImhvAiBZZczGHwU%2FNIbwtqo4gk0IROFsHSIA0Q46fHUzXUP%2FoA%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension contains an invalid OID value: 2.16.840.1.114569.1.1.1. Available OIDs: 2.16.840.1.114569.1.1.3, 2.16.840.1.114569.1.1.4 |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 902J', but common name is 'SHAKEN IP Link Telecom Inc. 902J' |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 28 Nov 23 10:53 UTC