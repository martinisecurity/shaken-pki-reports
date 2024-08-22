# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN i3 Broadband 5800

Tested At: 22 Aug 24 15:45 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -55 day(s)\
Subject: CN=SHAKEN i3 Broadband 5800, O=i3 Broadband, ST=Illinois, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/5800/429C7C70711E3820F0B8E1DEAE6FF32622649DFD.pem

[View certificate details](https://x509.io/?cert=MIICvzCCAmSgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJknf0wCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTI0MDUyOTAxNTg1NVoXDTI0MDYyODAxNTg1NVowWjELMAkGA1UEBhMCVVMxETAPBgNVBAgMCElsbGlub2lzMRUwEwYDVQQKDAxpMyBCcm9hZGJhbmQxITAfBgNVBAMMGFNIQUtFTiBpMyBCcm9hZGJhbmQgNTgwMDBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABAb2%2FFui7vD3uEW7Pyrp%2Fi02zLgVQObAK6F3oQX%2BrSrZQNm8Og7OAur47YY%2BXDAzT4VFI4o%2Ff%2BZr5CiRv43r9W6jgdswgdgwFgYIKwYBBQUHARoECjAIoAYWBDU4MDAwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEEMB0GA1UdDgQWBBQ9PhQb1%2Bf3L5dPC0Qs93eKwdExITAfBgNVHSMEGDAWgBSs05P1Q0PMCr5FWBcTfZJ83MMBRjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSQAwRgIhALFdodAk4VwDw5BKSiXmr7SjsO7L9j%2BaAJfKziGyfBusAiEA9lymBjGuvZ2eXWqjPwhrQUE2IAnMEIl%2FMllIOXPEUUQ%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 5800', but common name is 'SHAKEN i3 Broadband 5800' |
| [e_atis_ext_crl_distribution_struct](../../ISSUES/e_atis_ext_crl_distribution_struct/README.md) | error | ATIS1000080 | CRL Distribution Point shall contain a CRLIssuer field |


Generated: 22 Aug 24 16:06 UTC