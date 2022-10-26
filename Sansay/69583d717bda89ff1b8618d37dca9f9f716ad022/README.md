# STIR/SHAKEN CA Ecosystem Compliance
## Sansay

### Certificate 69583d717bda89ff1b8618d37dca9f9f716ad022
Tested At: 2022-10-26 21:14:16 +0000 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 16 day(s)\
Subject: CN=SHAKEN Matrix 7379, OU=Operations, O=Matrix, ST=Texas, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US

Link: https://cr.sansay.com/Matrix-7379

View: [Click to view](https://understandingwebpki.com/?cert=MIICxDCCAmqgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkT1gwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAxMTIyMTcxNloXDTIyMTExMDIyMTcxNlowYDELMAkGA1UEBhMCVVMxDjAMBgNVBAgMBVRleGFzMQ8wDQYDVQQKDAZNYXRyaXgxEzARBgNVBAsMCk9wZXJhdGlvbnMxGzAZBgNVBAMMElNIQUtFTiBNYXRyaXggNzM3OTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABE4GmAGYBCOAjL6rb9BWCSH32ogQfdkjAeHFQqwI%2BV0kBB1pSP%2FpqKSG6LrIf2SEMB8QZu3GjwfBz4B5ZR1PzpCjgdswgdgwFgYIKwYBBQUHARoECjAIoAYWBDczNzkwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMB0GA1UdDgQWBBRw4gWxXeqIxrPFYCmD8l3g2uZo4zAfBgNVHSMEGDAWgBSs05P1Q0PMCr5FWBcTfZJ83MMBRjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSAAwRQIhAJGz77mSX7hq3ldxHQHrG42Ja%2BcBO5eOG0Kzha5mILvNAiBDlYTkYsnfTNVGvXJQCAZiuFQGwwy78RBk%2Fqaabc7Vbg%3D%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_sti_certificate_policies | error | ATIS-1000080v4 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| e_cp1_3_subject_sn | error | CPv1.3 | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| w_cp1_3_subject_rdn_unknown | warn | CPv1.3 | Only CN, C, O, and SERIALNUMBER can be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| e_sti_subject_cn | error | ATIS-1000080v4 | Common name shall contain the text string 'SHAKEN 7379' |
| e_sti_crl_distribution | error | ATIS-1000080v4 | CRL Distribution Point shall be reachable if the requesting IP address within the program ACLs |
| e_cp1_3_ambiguous_identifier | error | CPv1.3 | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |

Generated: 26/10/2022 at 21:14:23