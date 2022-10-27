# STIR/SHAKEN CA Ecosystem Compliance
## Sansay

### Certificate 1faae68a15496561af4f703e680003cfb4d57a70
Tested At: 2022-10-27 22:31:47 +0000 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 30 day(s)\
Subject: CN=SHAKEN Noble Systems Communications LLC 187J, OU=NOC, O=Noble Systems Communications LLC, ST=Georgia, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US

Link: https://cr.sansay.com/NobleSys_187J

View: [Click to view](https://understandingwebpki.com/?cert=MIIC9TCCApqgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkU4AwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAyNzAwMzkxMloXDTIyMTEyNjAwMzkxMlowgY8xCzAJBgNVBAYTAlVTMRAwDgYDVQQIDAdHZW9yZ2lhMSkwJwYDVQQKDCBOb2JsZSBTeXN0ZW1zIENvbW11bmljYXRpb25zIExMQzEMMAoGA1UECwwDTk9DMTUwMwYDVQQDDCxTSEFLRU4gTm9ibGUgU3lzdGVtcyBDb21tdW5pY2F0aW9ucyBMTEMgMTg3SjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABP%2FmGlJuj%2FstdPV1ajhIPkNsLxwZQGX6idMO%2Fp8KZhhsDFAajP4jO2H2H5v6dRG3QkxY459oHsT7Ovh3AZPOq46jgdswgdgwFgYIKwYBBQUHARoECjAIoAYWBDE4N0owFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMB0GA1UdDgQWBBROIbMJwncdkQ3fRdiI1rOMaa87hDAfBgNVHSMEGDAWgBSs05P1Q0PMCr5FWBcTfZJ83MMBRjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSQAwRgIhAP9NHBgtmyC9RQXWSaw25zvHwqfudApj%2Bru5jf%2BgRES%2FAiEAwC1Tw6jsgpfnWeBt5bdMA7q489U5tiyQ9cz2Hjo%2FKRI%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_cp1_3_subject_sn | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| e_sti_certificate_policies | error | ATIS-1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| e_sti_subject_cn | error | ATIS-1000080 | Common name shall contain the text string 'SHAKEN 187J' |
| w_cp1_3_subject_rdn_unknown | warn | United States SHAKEN CP | Only CN, C, O, and SERIALNUMBER can be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| e_cp1_3_ambiguous_identifier | error | United States SHAKEN CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |

Generated: 27/10/2022 at 22:33:03