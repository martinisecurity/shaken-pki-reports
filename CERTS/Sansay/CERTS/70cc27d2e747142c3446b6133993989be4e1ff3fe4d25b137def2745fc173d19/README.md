# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Mercury Access Solutions 634K

Tested At: 04 Oct 24 16:07 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 270 day(s)\
Subject: CN=SHAKEN Mercury Access Solutions 634K, OU=Mercury Access Solutions, O=Mercury Access Solutions, ST=Kansas, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/634K/429C7C70711E3820F0B8E1DEAE6FF3262264A1E3.pem

[View certificate details](https://x509.io/?cert=MIIC%2BTCCAp6gAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkoeMwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTI0MDcwMTE2MjEyOFoXDTI1MDcwMTE2MjEyOFowgZMxCzAJBgNVBAYTAlVTMQ8wDQYDVQQIDAZLYW5zYXMxITAfBgNVBAoMGE1lcmN1cnkgQWNjZXNzIFNvbHV0aW9uczEhMB8GA1UECwwYTWVyY3VyeSBBY2Nlc3MgU29sdXRpb25zMS0wKwYDVQQDDCRTSEFLRU4gTWVyY3VyeSBBY2Nlc3MgU29sdXRpb25zIDYzNEswWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATV9yc89D0X32A36kZJxXSlWGqT%2BEAuDP1cecycLtEQGAy5Ww3rk7xdBmMFQVYrctYngyE5jgyWL6YgZGF0QiA%2Bo4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQ2MzRLMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzAdBgNVHQ4EFgQUa8jbUhbOdokSIf2VxE8v%2B6R0NSQwHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0kAMEYCIQDcJgz4p6x64OIovK4JvPT4N6FJujoWSwFcZEHWqyxSBgIhAP4Vir3eOxHvJ6S%2FoU%2BQnHPojJKhoVh9qfs9yrK5%2Bggg)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 634K', but common name is 'SHAKEN Mercury Access Solutions 634K' |
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |


Generated: 04 Oct 24 16:29 UTC