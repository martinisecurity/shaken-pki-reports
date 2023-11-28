# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Ace Innovative Networks, Inc. 040K

Tested At: 28 Nov 23 19:58 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -234 day(s)\
Subject: emailAddress=elevitt@1pcom.com, CN=SHAKEN Ace Innovative Networks\\, Inc. 040K, OU=NOC, O=Ace Innovative Networks\\, Inc., ST=New York, C=US, emailAddress=elevitt@1pcom.com\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/040K/order/274_040K_84

[View certificate details](https://understandingwebpki.com/?cert=MIIDETCCAregAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkZAYwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDMwODIxMTcxOFoXDTIzMDQwNzIxMTcxOFowgawxCzAJBgNVBAYTAlVTMREwDwYDVQQIDAhOZXcgWW9yazEmMCQGA1UECgwdQWNlIElubm92YXRpdmUgTmV0d29ya3MsIEluYy4xDDAKBgNVBAsMA05PQzEyMDAGA1UEAwwpU0hBS0VOIEFjZSBJbm5vdmF0aXZlIE5ldHdvcmtzLCBJbmMuIDA0MEsxIDAeBgkqhkiG9w0BCQEWEWVsZXZpdHRAMXBjb20uY29tMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEOqyYRuppi4NWGh0Dk0LShTDAr4sxeglOkve2zQwZPQJfpBYQOwwROCoWfqq%2BJEldZwIV9zMtR2PhzXqB%2B22djKOB2zCB2DAWBggrBgEFBQcBGgQKMAigBhYEMDQwSzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQMwHQYDVR0OBBYEFI4Gh%2FAp%2BjtuoawM59l0d1Wq6oMoMB8GA1UdIwQYMBaAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNIADBFAiEAjBzhInjZx%2BD%2FRRMEp8cHJhceWXkbhavmMKNMi6Bua90CIGaeHcKB1RdDCkgnX9zS7OXvUsQqmnSMdrstcAgi9cX2)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 040K', but common name is 'SHAKEN Ace Innovative Networks, Inc. 040K' |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 28 Nov 23 20:21 UTC