# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Ace Innovative Networks, Inc. 040K

Tested At: 21 Nov 23 18:51 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -312 day(s)\
Subject: emailAddress=elevitt@1pcom.com, CN=SHAKEN Ace Innovative Networks\\, Inc. 040K, OU=NOC, O=Ace Innovative Networks\\, Inc., ST=New York, C=US, emailAddress=elevitt@1pcom.com\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/040K/order/190_040K_84

[View certificate details](https://understandingwebpki.com/?cert=MIIDETCCAregAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkWHcwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTIxNDA0MTk0M1oXDTIzMDExMzA0MTk0M1owgawxCzAJBgNVBAYTAlVTMREwDwYDVQQIDAhOZXcgWW9yazEmMCQGA1UECgwdQWNlIElubm92YXRpdmUgTmV0d29ya3MsIEluYy4xDDAKBgNVBAsMA05PQzEyMDAGA1UEAwwpU0hBS0VOIEFjZSBJbm5vdmF0aXZlIE5ldHdvcmtzLCBJbmMuIDA0MEsxIDAeBgkqhkiG9w0BCQEWEWVsZXZpdHRAMXBjb20uY29tMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEOqyYRuppi4NWGh0Dk0LShTDAr4sxeglOkve2zQwZPQJfpBYQOwwROCoWfqq%2BJEldZwIV9zMtR2PhzXqB%2B22djKOB2zCB2DAWBggrBgEFBQcBGgQKMAigBhYEMDQwSzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFI4Gh%2FAp%2BjtuoawM59l0d1Wq6oMoMB8GA1UdIwQYMBaAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNIADBFAiBvJjj85%2FTVEm7BJpVbHwLme02AnBwGbWP%2FeNesQ4LxCQIhAMJDsysLUiiKrmF7dZscdczZVhM3BiNgjLjm3KVue76c)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | the Certificate Policies extension contains an invalid OID value: 2.16.840.1.114569.1.1.1. Available OIDs: [2.16.840.1.114569.1.1.3 2.16.840.1.114569.1.1.4] |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 040K' |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 21 Nov 23 19:18 UTC