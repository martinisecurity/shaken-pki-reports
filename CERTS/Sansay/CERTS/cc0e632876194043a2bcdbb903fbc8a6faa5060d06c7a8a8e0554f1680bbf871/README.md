# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Matrix 7379

Tested At: 04 Nov 22 01:10 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 22 day(s)\
Subject: CN=SHAKEN Matrix 7379, OU=Operations, O=Matrix, ST=Texas, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/Matrix-7379

[View certificate details](https://understandingwebpki.com/?cert=MIICxDCCAmqgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkU4owCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAyNzAwNDMxNloXDTIyMTEyNjAwNDMxNlowYDELMAkGA1UEBhMCVVMxDjAMBgNVBAgMBVRleGFzMQ8wDQYDVQQKDAZNYXRyaXgxEzARBgNVBAsMCk9wZXJhdGlvbnMxGzAZBgNVBAMMElNIQUtFTiBNYXRyaXggNzM3OTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABE4GmAGYBCOAjL6rb9BWCSH32ogQfdkjAeHFQqwI%2BV0kBB1pSP%2FpqKSG6LrIf2SEMB8QZu3GjwfBz4B5ZR1PzpCjgdswgdgwFgYIKwYBBQUHARoECjAIoAYWBDczNzkwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMB0GA1UdDgQWBBRw4gWxXeqIxrPFYCmD8l3g2uZo4zAfBgNVHSMEGDAWgBSs05P1Q0PMCr5FWBcTfZJ83MMBRjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSAAwRQIgDKgiW5IlU96kzZq30e%2F1p%2BNkwrCzsK%2Bh86%2BNbTzN3tECIQCZIeIKy0tAAkm0aIvj5jryJ0x6I%2BixaWFn9ZsTF8D4kw%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 7379' |


Generated: 04 Nov 22 01:11 UTC