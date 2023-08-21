# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Godaddy 463K

Tested At: 21 Aug 23 20:14 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 25 day(s)\
Subject: CN=SHAKEN Godaddy 463K, OU=ENM-OPS, O=Godaddy, ST=Arizona, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/Godaddy_463K

[View certificate details](https://understandingwebpki.com/?cert=MIICxjCCAmugAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkfS4wCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDgxNTIxMTQ1M1oXDTIzMDkxNDIxMTQ1M1owYTELMAkGA1UEBhMCVVMxEDAOBgNVBAgMB0FyaXpvbmExEDAOBgNVBAoMB0dvZGFkZHkxEDAOBgNVBAsMB0VOTS1PUFMxHDAaBgNVBAMME1NIQUtFTiBHb2RhZGR5IDQ2M0swWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAS5T5DHueAT3RIkDrGXNdFtORYQmRDJcXUxo%2FJpGAbTbqOdgnwAKo3S8cbvYGUyeXAYqmcu5UMkNzaA7Gb8tMzpo4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQ0NjNLMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzAdBgNVHQ4EFgQUneZjdQJxFhEGrrYYv3p4k1GLKh4wHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0kAMEYCIQDl%2BEA3hIBN9P%2BI0ng%2FuuQtCBrJPMmqIqpMAhKN60%2BH6AIhAI%2Fw2293jkamVWi7i1zGNUDpl56yUgzgShIm0V0inQKx)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 463K' |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 21 Aug 23 20:18 UTC