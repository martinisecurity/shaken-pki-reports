# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Matrix 7379

Tested At: 04 Oct 24 16:07 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 6 day(s)\
Subject: CN=SHAKEN Matrix 7379, OU=Operations, O=Matrix, ST=Texas, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/7379/429C7C70711E3820F0B8E1DEAE6FF3262264A9DA.pem

[View certificate details](https://x509.io/?cert=MIICxDCCAmqgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkqdowCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTI0MDkxMDE0MDMzN1oXDTI0MTAxMDE0MDMzN1owYDELMAkGA1UEBhMCVVMxDjAMBgNVBAgMBVRleGFzMQ8wDQYDVQQKDAZNYXRyaXgxEzARBgNVBAsMCk9wZXJhdGlvbnMxGzAZBgNVBAMMElNIQUtFTiBNYXRyaXggNzM3OTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABE4GmAGYBCOAjL6rb9BWCSH32ogQfdkjAeHFQqwI%2BV0kBB1pSP%2FpqKSG6LrIf2SEMB8QZu3GjwfBz4B5ZR1PzpCjgdswgdgwFgYIKwYBBQUHARoECjAIoAYWBDczNzkwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMB0GA1UdDgQWBBRw4gWxXeqIxrPFYCmD8l3g2uZo4zAfBgNVHSMEGDAWgBSs05P1Q0PMCr5FWBcTfZJ83MMBRjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSAAwRQIhAIkDzi8cN7sgI%2FsCR4ywzomPdf76KGAp1I5%2FCSndTXoRAiBgeyKyVgZj%2FNtMLnJLroB164vokfcACVNRsA0QPFtR%2BA%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 7379', but common name is 'SHAKEN Matrix 7379' |
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension contains an invalid OID value: 2.16.840.1.114569.1.1.1. Available OIDs: 2.16.840.1.114569.1.1.3, 2.16.840.1.114569.1.1.4 |


Generated: 04 Oct 24 16:29 UTC