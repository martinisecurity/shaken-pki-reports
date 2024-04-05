# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Star2Star Communications, LLC 590J

Tested At: 05 Apr 24 18:53 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: -135 day(s)\
Subject: CN=SHAKEN Star2Star Communications\\, LLC 590J, OU=BVPROD, O=Star2Star Communications\\, LLC, ST=Florida, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/star2star_stirshaken_20220922

[View certificate details](https://x509.io/?cert=MIIC8TCCApegAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkVhMwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTEyMTIxMTUzMVoXDTIzMTEyMTIxMTUzMVowgYwxCzAJBgNVBAYTAlVTMRAwDgYDVQQIDAdGbG9yaWRhMSYwJAYDVQQKDB1TdGFyMlN0YXIgQ29tbXVuaWNhdGlvbnMsIExMQzEPMA0GA1UECwwGQlZQUk9EMTIwMAYDVQQDDClTSEFLRU4gU3RhcjJTdGFyIENvbW11bmljYXRpb25zLCBMTEMgNTkwSjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABDmML9Ds%2Bs4rHgp6h2fVFdvO%2FmXDpLLYprajtfIT8Mjl0kEsB4G1QmZPJaI6Y7J9L00kbNcLB3vo3ZaSuQqu3ImjgdswgdgwFgYIKwYBBQUHARoECjAIoAYWBDU5MEowFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMB0GA1UdDgQWBBS9rZbJdjb%2BC%2BmKmakNHXUxGYtUSzAfBgNVHSMEGDAWgBSs05P1Q0PMCr5FWBcTfZJ83MMBRjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSAAwRQIhAIwyeWZvMjv3hzEsQv8LIVbB27Sn0ZWw%2F8d8zC%2FHXc3pAiAPjBNA7LNsDlsQlbg0eGK5vNgR2c0DX6GIFY7eH%2F30dQ%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 590J', but common name is 'SHAKEN Star2Star Communications, LLC 590J' |
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension contains an invalid OID value: 2.16.840.1.114569.1.1.1. Available OIDs: 2.16.840.1.114569.1.1.3, 2.16.840.1.114569.1.1.4 |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 05 Apr 24 19:04 UTC