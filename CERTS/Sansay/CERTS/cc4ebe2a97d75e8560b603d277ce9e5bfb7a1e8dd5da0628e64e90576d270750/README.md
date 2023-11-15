# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Godaddy 463K

Tested At: 15 Nov 23 18:03 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 364 day(s)\
Subject: CN=SHAKEN Godaddy 463K, OU=ENM-OPS, O=Godaddy, ST=Arizona, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/Godaddy_463K

[View certificate details](https://understandingwebpki.com/?cert=MIICxTCCAmugAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkhj4wCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMTExNDEzMzYwOVoXDTI0MTExMzEzMzYwOVowYTELMAkGA1UEBhMCVVMxEDAOBgNVBAgMB0FyaXpvbmExEDAOBgNVBAoMB0dvZGFkZHkxEDAOBgNVBAsMB0VOTS1PUFMxHDAaBgNVBAMME1NIQUtFTiBHb2RhZGR5IDQ2M0swWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAS5T5DHueAT3RIkDrGXNdFtORYQmRDJcXUxo%2FJpGAbTbqOdgnwAKo3S8cbvYGUyeXAYqmcu5UMkNzaA7Gb8tMzpo4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQ0NjNLMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzAdBgNVHQ4EFgQUneZjdQJxFhEGrrYYv3p4k1GLKh4wHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0gAMEUCIQCQxHlTZxQco0ud8HJz1ttVcI7RYs6d%2FRztTeBg0U%2B0rwIgJlT%2FPJQgpKzaT%2BcJOAUSaFb2KU1aTHxpZKlMAE%2BrUCs%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 463K' |


Generated: 15 Nov 23 18:10 UTC