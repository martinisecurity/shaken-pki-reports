# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN KW Corporation, inc. 206k

Tested At: 18 Aug 25 20:48 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: -278 day(s)\
Subject: CN=SHAKEN KW Corporation\\, inc. 206k, OU=ITDEPT, O=KW Corporation\\, inc., ST=Michigan, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/KW_Corporation_inc._206k

[View certificate details](https://x509.io/?cert=MIIC3zCCAoWgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkhj8wCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMTExNDEzMzY1MVoXDTI0MTExMzEzMzY1MVowezELMAkGA1UEBhMCVVMxETAPBgNVBAgMCE1pY2hpZ2FuMR0wGwYDVQQKDBRLVyBDb3Jwb3JhdGlvbiwgaW5jLjEPMA0GA1UECwwGSVRERVBUMSkwJwYDVQQDDCBTSEFLRU4gS1cgQ29ycG9yYXRpb24sIGluYy4gMjA2azBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABEnDGEYDl%2BKoqka6szbknjuoSRME08fKIJ7G63N4zuqOLArpVqUDvScDP1DaQXNVNzFKI%2FYoEhPcEyr1hGgPXvGjgdswgdgwFgYIKwYBBQUHARoECjAIoAYWBDIwNmswFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMB0GA1UdDgQWBBTzVOhLQ19w0uTRV4oBQ7ygxpO6pjAfBgNVHSMEGDAWgBSs05P1Q0PMCr5FWBcTfZJ83MMBRjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSAAwRQIgU2sVIqJbnIm4hv6pjV047%2Fx3Lutt0b9V8K827ECjyY8CIQCdmx2T0a9zrdfX6%2FoJk%2B7JzOISrCRjRfmhQkquC61oew%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_tn_auth_list_spc_format](../../ISSUES/e_atis_tn_auth_list_spc_format/README.md) | error | ATIS1000080 | the SPC value '206k' contains characters other than uppercase letters and numbers |
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension contains an invalid OID value: 2.16.840.1.114569.1.1.1. Available OIDs: 2.16.840.1.114569.1.1.3, 2.16.840.1.114569.1.1.4 |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 206k', but common name is 'SHAKEN KW Corporation, inc. 206k' |


Generated: 18 Aug 25 21:13 UTC