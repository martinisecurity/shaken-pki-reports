# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Godaddy 463K

Tested At: 18 Aug 25 20:48 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: -278 day(s)\
Subject: CN=SHAKEN Godaddy 463K, OU=ENM-OPS, O=Godaddy, ST=Arizona, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/Godaddy_463K

[View certificate details](https://x509.io/?cert=MIICxTCCAmugAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkhj4wCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMTExNDEzMzYwOVoXDTI0MTExMzEzMzYwOVowYTELMAkGA1UEBhMCVVMxEDAOBgNVBAgMB0FyaXpvbmExEDAOBgNVBAoMB0dvZGFkZHkxEDAOBgNVBAsMB0VOTS1PUFMxHDAaBgNVBAMME1NIQUtFTiBHb2RhZGR5IDQ2M0swWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAS5T5DHueAT3RIkDrGXNdFtORYQmRDJcXUxo%2FJpGAbTbqOdgnwAKo3S8cbvYGUyeXAYqmcu5UMkNzaA7Gb8tMzpo4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQ0NjNLMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzAdBgNVHQ4EFgQUneZjdQJxFhEGrrYYv3p4k1GLKh4wHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0gAMEUCIQCQxHlTZxQco0ud8HJz1ttVcI7RYs6d%2FRztTeBg0U%2B0rwIgJlT%2FPJQgpKzaT%2BcJOAUSaFb2KU1aTHxpZKlMAE%2BrUCs%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 463K', but common name is 'SHAKEN Godaddy 463K' |
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |


Generated: 18 Aug 25 21:13 UTC