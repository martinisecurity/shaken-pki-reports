# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Consolidated Communications 5113

Tested At: 21 Aug 23 20:14 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 24 day(s)\
Subject: CN=SHAKEN Consolidated Communications 5113, OU=Operations, O=Consolidated Communications, ST=New Hampshire, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/Consolidated_5113.crt

[View certificate details](https://understandingwebpki.com/?cert=MIIDpTCCA0ugAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkfQkwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDgxNTEzMjgzNFoXDTIzMDkxNDEzMjgzNFowgZIxCzAJBgNVBAYTAlVTMRYwFAYDVQQIDA1OZXcgSGFtcHNoaXJlMSQwIgYDVQQKDBtDb25zb2xpZGF0ZWQgQ29tbXVuaWNhdGlvbnMxEzARBgNVBAsMCk9wZXJhdGlvbnMxMDAuBgNVBAMMJ1NIQUtFTiBDb25zb2xpZGF0ZWQgQ29tbXVuaWNhdGlvbnMgNTExMzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABCJN0D3bum0fgrJUbj3xuV2tyVkYhjoNUh%2B9aafaYb2kdEPqLeK%2FutlailErROW8w%2FNHVR8obT%2B6N1%2Bf%2B7SO602jggGIMIIBhDAWBggrBgEFBQcBGgQKMAigBhYENTExMzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFFlr7SUu%2BKsDd4PxHNGZiB6cc8ksMIHKBgNVHSMEgcIwgb%2BAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGoYGQpIGNMIGKMQswCQYDVQQGEwJVUzETMBEGA1UECAwKQ2FsaWZvcm5pYTESMBAGA1UEBwwJU2FuIERpZWdvMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEhMB8GA1UEAwwYU0hBS0VOIFNhbnNheSBSb290IENBIFVTghQUtV84BfXwexO1%2BLDe7SDyUXA%2BnDBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSAAwRQIhANCjziJMm1vxmQLprq65nN9NELHWMLZIx66Ue8qhYBpbAiBJDgqj%2Bsy3no75jEIRCdm9EvH1zwS2pEHEQVU2kfQQ1g%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 5113' |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |


Generated: 21 Aug 23 20:18 UTC