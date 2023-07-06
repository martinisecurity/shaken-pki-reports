# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN IPSBS Managed Services LLC 828J

Tested At: 06 Jul 23 14:02 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -13 day(s)\
Subject: emailAddress=iconectiv-hmc@hostmycalls.com, CN=SHAKEN IPSBS Managed Services LLC 828J, OU=NOC, O=IPSBS Managed Services LLC, ST=Tennessee, C=US, emailAddress=iconectiv-hmc@hostmycalls.com\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/828J/order/311_828J_86

[View certificate details](https://understandingwebpki.com/?cert=MIIDFzCCAr6gAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkcZcwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDUyNDA0MzYzMFoXDTIzMDYyMzA0MzYzMFowgbMxCzAJBgNVBAYTAlVTMRIwEAYDVQQIDAlUZW5uZXNzZWUxIzAhBgNVBAoMGklQU0JTIE1hbmFnZWQgU2VydmljZXMgTExDMQwwCgYDVQQLDANOT0MxLzAtBgNVBAMMJlNIQUtFTiBJUFNCUyBNYW5hZ2VkIFNlcnZpY2VzIExMQyA4MjhKMSwwKgYJKoZIhvcNAQkBFh1pY29uZWN0aXYtaG1jQGhvc3RteWNhbGxzLmNvbTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABA0oRMeHJ2d9TAMrJ28TK2kM3c4mxC%2FnlbB%2BBtN91rEkHvGYc3HSWy5ngDZYSqbxot%2BKwGJXeG8fU2HeFUO3MT%2BjgdswgdgwFgYIKwYBBQUHARoECjAIoAYWBDgyOEowFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMB0GA1UdDgQWBBRC8rHWbSvVLaSU37nCkTHvCmHvFzAfBgNVHSMEGDAWgBSs05P1Q0PMCr5FWBcTfZJ83MMBRjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDRwAwRAIgUY9Mlx9Tcn7yqyQVvpw%2B9xUPHNLsPZLVInLzzAe4VbwCIB1SPhC6oT1EpJFpUQLeV1Q7nwV1BlrdQlq1jIu5q63L)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [w_cp_1_3_subject_email](../../ISSUES/w_cp_1_3_subject_email/README.md) | warn | US_SHAKEN_CP | Email addresses are not allowed as the CP does not specify how to validate them |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 828J' |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 06 Jul 23 14:08 UTC