# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 1stPoint Communications, LLC 463G

Tested At: 04 Oct 24 16:02 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -575 day(s)\
Subject: emailAddress=elevitt@1pcom.com, CN=SHAKEN 1stPoint Communications\\, LLC 463G, OU=NOC, O=1stPoint Communications\\, LLC, ST=New York, C=US, emailAddress=elevitt@1pcom.com\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/463G/order/558_463G_53

[View certificate details](https://x509.io/?cert=MIIDEDCCArWgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkYKkwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDIwNzA3MjAwOVoXDTIzMDMwOTA3MjAwOVowgaoxCzAJBgNVBAYTAlVTMREwDwYDVQQIDAhOZXcgWW9yazElMCMGA1UECgwcMXN0UG9pbnQgQ29tbXVuaWNhdGlvbnMsIExMQzEMMAoGA1UECwwDTk9DMTEwLwYDVQQDDChTSEFLRU4gMXN0UG9pbnQgQ29tbXVuaWNhdGlvbnMsIExMQyA0NjNHMSAwHgYJKoZIhvcNAQkBFhFlbGV2aXR0QDFwY29tLmNvbTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABC7Os0cKRZ4dUV6B4QP8X%2BK3yo23fTX6t3lLNg%2FDTv2neRCPqx2o%2BZM8JVwSxj4LgyXCBlAuJFSvSoNR2r8CMEajgdswgdgwFgYIKwYBBQUHARoECjAIoAYWBDQ2M0cwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMB0GA1UdDgQWBBTPNaivneus6uiFFWXaD%2FF53JURmzAfBgNVHSMEGDAWgBSs05P1Q0PMCr5FWBcTfZJ83MMBRjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSQAwRgIhAJXH9xTyQCK9gUaT5mw7gDgrtm%2FDcCnjkCWEX7nAuXCmAiEAiPNAZ3oJWXXCnLhQbzi02L%2Bk9dD%2FHwo6xXgS1rwyRHo%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 463G', but common name is 'SHAKEN 1stPoint Communications, LLC 463G' |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 04 Oct 24 16:29 UTC