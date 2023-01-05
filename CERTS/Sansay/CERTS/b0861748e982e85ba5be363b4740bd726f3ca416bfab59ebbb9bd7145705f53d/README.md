# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN IPSBS Managed Services LLC 828J

Tested At: 05 Jan 23 18:25 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -56 day(s)\
Subject: emailAddress=iconectiv-hmc@hostmycalls.com, CN=SHAKEN IPSBS Managed Services LLC 828J, OU=NOC, O=IPSBS Managed Services LLC, ST=Tennessee, C=US, emailAddress=iconectiv-hmc@hostmycalls.com\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/828J/order/86_828J_86

[View certificate details](https://understandingwebpki.com/?cert=MIIDxTCCA2ygAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkTrkwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAxMDIwNDY1NVoXDTIyMTEwOTIwNDY1NVowgbMxCzAJBgNVBAYTAlVTMRIwEAYDVQQIDAlUZW5uZXNzZWUxIzAhBgNVBAoMGklQU0JTIE1hbmFnZWQgU2VydmljZXMgTExDMQwwCgYDVQQLDANOT0MxLzAtBgNVBAMMJlNIQUtFTiBJUFNCUyBNYW5hZ2VkIFNlcnZpY2VzIExMQyA4MjhKMSwwKgYJKoZIhvcNAQkBFh1pY29uZWN0aXYtaG1jQGhvc3RteWNhbGxzLmNvbTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABA0oRMeHJ2d9TAMrJ28TK2kM3c4mxC%2FnlbB%2BBtN91rEkHvGYc3HSWy5ngDZYSqbxot%2BKwGJXeG8fU2HeFUO3MT%2BjggGIMIIBhDAWBggrBgEFBQcBGgQKMAigBhYEODI4SjAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFELysdZtK9UtpJTfucKRMe8KYe8XMIHKBgNVHSMEgcIwgb%2BAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGoYGQpIGNMIGKMQswCQYDVQQGEwJVUzETMBEGA1UECAwKQ2FsaWZvcm5pYTESMBAGA1UEBwwJU2FuIERpZWdvMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEhMB8GA1UEAwwYU0hBS0VOIFNhbnNheSBSb290IENBIFVTghQUtV84BfXwexO1%2BLDe7SDyUXA%2BnDBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDRwAwRAIgU%2BO8Witoj4w1tvPlf%2BGl0id9I%2FcpFD4bqCaogLVYr%2FACIEEHDrTZvvEMtpU7OibyzLpWxU%2BB%2BiW0z4C5bw9JQY2Q)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 828J' |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [w_cp_1_3_subject_email](../../ISSUES/w_cp_1_3_subject_email/README.md) | warn | US_SHAKEN_CP | Email addresses are not allowed as the CP does not specify how to validate them |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 05 Jan 23 18:35 UTC