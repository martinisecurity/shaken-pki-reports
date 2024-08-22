# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN ACS Technologies 488K

Tested At: 22 Aug 24 15:41 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 159 day(s)\
Subject: CN=SHAKEN ACS Technologies 488K, OU=ACS Technologies, O=ACS Technologies, ST=South Carolina, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/488K/429C7C70711E3820F0B8E1DEAE6FF32622648FD4.pem

[View certificate details](https://x509.io/?cert=MIIC6TCCAo6gAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkj9QwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTI0MDEyOTEzMjMwN1oXDTI1MDEyODEzMjMwN1owgYMxCzAJBgNVBAYTAlVTMRcwFQYDVQQIDA5Tb3V0aCBDYXJvbGluYTEZMBcGA1UECgwQQUNTIFRlY2hub2xvZ2llczEZMBcGA1UECwwQQUNTIFRlY2hub2xvZ2llczElMCMGA1UEAwwcU0hBS0VOIEFDUyBUZWNobm9sb2dpZXMgNDg4SzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABFMe9YwMofpHj%2F50MuZFl7yLQATMdgdAhuqq2qwLR19fK2UNrCkFqpxTLfbOBgBLlNAVLCyyePfpFk07nFEgbSOjgdswgdgwFgYIKwYBBQUHARoECjAIoAYWBDQ4OEswFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMB0GA1UdDgQWBBTbQUsyrlcmqIln2ps%2B8zHK2r0EpTAfBgNVHSMEGDAWgBSs05P1Q0PMCr5FWBcTfZJ83MMBRjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSQAwRgIhAPFB9NaDuqT4TPrmToBHGKtZvOG4Ltf%2Bk2OYtHCb8zTtAiEAoHs5pPWITZUVhy%2Fo6uKnRq9FEIcww4bRh5jXYzeeyuM%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 488K', but common name is 'SHAKEN ACS Technologies 488K' |


Generated: 22 Aug 24 16:06 UTC