# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Primo Dialler LLC 249K

Tested At: 21 Nov 23 19:08 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -271 day(s)\
Subject: CN=SHAKEN Primo Dialler LLC 249K, OU=Service Provider, O=Primo Dialler LLC, ST=Wyoming, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/Primo_0809

[View certificate details](https://understandingwebpki.com/?cert=MIIC4jCCAoigAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkXnkwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDEyNDA0MjkzOVoXDTIzMDIyMzA0MjkzOVowfjELMAkGA1UEBhMCVVMxEDAOBgNVBAgMB1d5b21pbmcxGjAYBgNVBAoMEVByaW1vIERpYWxsZXIgTExDMRkwFwYDVQQLDBBTZXJ2aWNlIFByb3ZpZGVyMSYwJAYDVQQDDB1TSEFLRU4gUHJpbW8gRGlhbGxlciBMTEMgMjQ5SzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABHlg2xvFwhbA5%2Ff8kG3PdxUc0PmWADerjS0uoXtuD2gCd3e0N5%2FhF9b9%2BM9k3%2FcLpm6Q0i7o%2BnD7O5mgGJhrJzyjgdswgdgwFgYIKwYBBQUHARoECjAIoAYWBDI0OUswFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMB0GA1UdDgQWBBSwuEXxFSv1C20CUEyNJEyr60U8eDAfBgNVHSMEGDAWgBSs05P1Q0PMCr5FWBcTfZJ83MMBRjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSAAwRQIhAIHuZfS%2Bhguun%2Be3xwitWHe5VEt7fjaSpUwYC%2FtqkCJqAiAtcTWviZlgLW%2Btr2sqIX%2Fw2yYllax6UIBx3WF%2F7M21%2Bg%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | the Certificate Policies extension contains an invalid OID value: 2.16.840.1.114569.1.1.1. Available OIDs: [2.16.840.1.114569.1.1.3 2.16.840.1.114569.1.1.4] |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 249K' |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 21 Nov 23 19:18 UTC