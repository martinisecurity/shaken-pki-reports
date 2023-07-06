# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 1stPoint Communications, LLC 463G

Tested At: 06 Jul 23 14:01 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -85 day(s)\
Subject: emailAddress=elevitt@1pcom.com, CN=SHAKEN 1stPoint Communications\\, LLC 463G, OU=NOC, O=1stPoint Communications\\, LLC, ST=New York, C=US, emailAddress=elevitt@1pcom.com\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/463G/order/589_463G_53

[View certificate details](https://understandingwebpki.com/?cert=MIIDEDCCArWgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkZL0wCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDMxMzA0MzAwOFoXDTIzMDQxMjA0MzAwOFowgaoxCzAJBgNVBAYTAlVTMREwDwYDVQQIDAhOZXcgWW9yazElMCMGA1UECgwcMXN0UG9pbnQgQ29tbXVuaWNhdGlvbnMsIExMQzEMMAoGA1UECwwDTk9DMTEwLwYDVQQDDChTSEFLRU4gMXN0UG9pbnQgQ29tbXVuaWNhdGlvbnMsIExMQyA0NjNHMSAwHgYJKoZIhvcNAQkBFhFlbGV2aXR0QDFwY29tLmNvbTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABC7Os0cKRZ4dUV6B4QP8X%2BK3yo23fTX6t3lLNg%2FDTv2neRCPqx2o%2BZM8JVwSxj4LgyXCBlAuJFSvSoNR2r8CMEajgdswgdgwFgYIKwYBBQUHARoECjAIoAYWBDQ2M0cwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMB0GA1UdDgQWBBTPNaivneus6uiFFWXaD%2FF53JURmzAfBgNVHSMEGDAWgBSs05P1Q0PMCr5FWBcTfZJ83MMBRjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSQAwRgIhAM%2BHCDdZoLtbsodRWuuxlgdOlAKagToaWRIqy%2FDZjA9WAiEAoZaclDcOELQ756UyXGKLSV3aHn8t9Xz327oGJF2Kj5A%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 463G' |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_cp_1_3_subject_email](../../ISSUES/w_cp_1_3_subject_email/README.md) | warn | US_SHAKEN_CP | Email addresses are not allowed as the CP does not specify how to validate them |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 06 Jul 23 14:08 UTC