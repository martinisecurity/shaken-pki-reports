# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Consolidated Communications 5113

Tested At: 04 Oct 24 16:02 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -112 day(s)\
Subject: CN=SHAKEN Consolidated Communications 5113, OU=Operations, O=Consolidated Communications, ST=New Hampshire, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/5113/429C7C70711E3820F0B8E1DEAE6FF32622649C57.pem

[View certificate details](https://x509.io/?cert=MIIDpTCCA0ugAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJknFcwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTI0MDUxNTE0MDk1NloXDTI0MDYxNDE0MDk1NlowgZIxCzAJBgNVBAYTAlVTMRYwFAYDVQQIDA1OZXcgSGFtcHNoaXJlMSQwIgYDVQQKDBtDb25zb2xpZGF0ZWQgQ29tbXVuaWNhdGlvbnMxEzARBgNVBAsMCk9wZXJhdGlvbnMxMDAuBgNVBAMMJ1NIQUtFTiBDb25zb2xpZGF0ZWQgQ29tbXVuaWNhdGlvbnMgNTExMzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABCJN0D3bum0fgrJUbj3xuV2tyVkYhjoNUh%2B9aafaYb2kdEPqLeK%2FutlailErROW8w%2FNHVR8obT%2B6N1%2Bf%2B7SO602jggGIMIIBhDAWBggrBgEFBQcBGgQKMAigBhYENTExMzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFFlr7SUu%2BKsDd4PxHNGZiB6cc8ksMIHKBgNVHSMEgcIwgb%2BAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGoYGQpIGNMIGKMQswCQYDVQQGEwJVUzETMBEGA1UECAwKQ2FsaWZvcm5pYTESMBAGA1UEBwwJU2FuIERpZWdvMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEhMB8GA1UEAwwYU0hBS0VOIFNhbnNheSBSb290IENBIFVTghQUtV84BfXwexO1%2BLDe7SDyUXA%2BnDBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSAAwRQIgHrXSC6R1sr2bmm%2FFXgd0OC5ypt4fhEhye7HVQSSQJW4CIQDOpPrJ7V%2B%2Bcs3Vloh5NFNtbOOBo2Ev85qMhYp6f0VFbQ%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 5113', but common name is 'SHAKEN Consolidated Communications 5113' |
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension contains an invalid OID value: 2.16.840.1.114569.1.1.1. Available OIDs: 2.16.840.1.114569.1.1.3, 2.16.840.1.114569.1.1.4 |


Generated: 04 Oct 24 16:29 UTC