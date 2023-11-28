# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Quality Voice & Data Inc. 548J

Tested At: 28 Nov 23 16:05 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -342 day(s)\
Subject: emailAddress=kelsey@qualityvoicedata.com, CN=SHAKEN Quality Voice & Data Inc. 548J, OU=Quality Voice & Data, O=Quality Voice & Data Inc., ST=Nebraska, C=US, emailAddress=kelsey@qualityvoicedata.com\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/548J/order/193_548J_67

[View certificate details](https://understandingwebpki.com/?cert=MIIDJDCCAsqgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkVgswCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTEyMTE0MDUxNFoXDTIyMTIyMTE0MDUxNFowgb8xCzAJBgNVBAYTAlVTMREwDwYDVQQIDAhOZWJyYXNrYTEiMCAGA1UECgwZUXVhbGl0eSBWb2ljZSAmIERhdGEgSW5jLjEdMBsGA1UECwwUUXVhbGl0eSBWb2ljZSAmIERhdGExLjAsBgNVBAMMJVNIQUtFTiBRdWFsaXR5IFZvaWNlICYgRGF0YSBJbmMuIDU0OEoxKjAoBgkqhkiG9w0BCQEWG2tlbHNleUBxdWFsaXR5dm9pY2VkYXRhLmNvbTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABHxIpvtIpathfVQA%2Bf9h5YTuYtpurXcE4Fo%2BWfOrs4qTvs97AEUcl9TF5yLWG%2BWt%2FOsXbViKVjuykX3%2F5qk0fMSjgdswgdgwFgYIKwYBBQUHARoECjAIoAYWBDU0OEowFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMB0GA1UdDgQWBBTK3lC3yDbCtXJbXLx0JRgRgpoEuTAfBgNVHSMEGDAWgBSs05P1Q0PMCr5FWBcTfZJ83MMBRjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSAAwRQIgEwTDfpVOqJxHbWIh%2BXdag%2FS30cW95FuLFyXkthQw%2BIECIQDevVMKTcGaKDe7n9%2B9xaOr2QEYL6utxMmWFMkuSthkTg%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 548J', but common name is 'SHAKEN Quality Voice & Data Inc. 548J' |
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension contains an invalid OID value: 2.16.840.1.114569.1.1.1. Available OIDs: 2.16.840.1.114569.1.1.3, 2.16.840.1.114569.1.1.4 |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 28 Nov 23 16:15 UTC