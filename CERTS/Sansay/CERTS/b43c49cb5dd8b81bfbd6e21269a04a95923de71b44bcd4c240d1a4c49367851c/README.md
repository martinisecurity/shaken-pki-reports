# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Consolidated Communications 5113

Tested At: 01 Nov 22 09:58 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 25 day(s)\
Subject: CN=SHAKEN Consolidated Communications 5113, OU=Operations, O=Consolidated Communications, ST=New Hampshire, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/Consolidated_5113.crt

View: [Click to view](https://understandingwebpki.com/?cert=MIIDpDCCA0ugAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkU3MwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAyNjIyNDAyNloXDTIyMTEyNTIyNDAyNlowgZIxCzAJBgNVBAYTAlVTMRYwFAYDVQQIDA1OZXcgSGFtcHNoaXJlMSQwIgYDVQQKDBtDb25zb2xpZGF0ZWQgQ29tbXVuaWNhdGlvbnMxEzARBgNVBAsMCk9wZXJhdGlvbnMxMDAuBgNVBAMMJ1NIQUtFTiBDb25zb2xpZGF0ZWQgQ29tbXVuaWNhdGlvbnMgNTExMzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABCJN0D3bum0fgrJUbj3xuV2tyVkYhjoNUh%2B9aafaYb2kdEPqLeK%2FutlailErROW8w%2FNHVR8obT%2B6N1%2Bf%2B7SO602jggGIMIIBhDAWBggrBgEFBQcBGgQKMAigBhYENTExMzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFFlr7SUu%2BKsDd4PxHNGZiB6cc8ksMIHKBgNVHSMEgcIwgb%2BAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGoYGQpIGNMIGKMQswCQYDVQQGEwJVUzETMBEGA1UECAwKQ2FsaWZvcm5pYTESMBAGA1UEBwwJU2FuIERpZWdvMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEhMB8GA1UEAwwYU0hBS0VOIFNhbnNheSBSb290IENBIFVTghQUtV84BfXwexO1%2BLDe7SDyUXA%2BnDBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDRwAwRAIhAIWII%2FFydWJmgXeMc2FKipMYUsurRuKJViSKKwKsyhwlAh8jIubCBIP2oIa3SvDX%2BEIGHZiYbLXq9HTPo8X%2FRlqT)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_cp1_3_ambiguous_identifier](../../ISSUES/e_cp1_3_ambiguous_identifier/README.md) | error | United States SHAKEN CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_sti_certificate_policies](../../ISSUES/e_sti_certificate_policies/README.md) | error | ATIS&#x2011;1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_cp1_3_subject_sn](../../ISSUES/e_cp1_3_subject_sn/README.md) | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_sti_subject_cn](../../ISSUES/e_sti_subject_cn/README.md) | error | ATIS&#x2011;1000080 | Common name shall contain the text string 'SHAKEN 5113' |
| [w_pki_subject_rdn_unknown](../../ISSUES/w_pki_subject_rdn_unknown/README.md) | warn | SHAKEN PKI Best Practice | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 01/11/2022 at 10:05:32