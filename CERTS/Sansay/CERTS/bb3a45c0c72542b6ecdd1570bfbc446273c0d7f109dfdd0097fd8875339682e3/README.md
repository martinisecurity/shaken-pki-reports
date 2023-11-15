# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Nemerald 326K

Tested At: 15 Nov 23 16:10 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 134 day(s)\
Subject: CN=SHAKEN Nemerald 326K, OU=Engineering, O=Nemerald, ST=California, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/Nemerald_326K

[View certificate details](https://understandingwebpki.com/?cert=MIICzTCCAnSgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkZ6QwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDMyODIzNTkwMFoXDTI0MDMyNzIzNTkwMFowajELMAkGA1UEBhMCVVMxEzARBgNVBAgMCkNhbGlmb3JuaWExETAPBgNVBAoMCE5lbWVyYWxkMRQwEgYDVQQLDAtFbmdpbmVlcmluZzEdMBsGA1UEAwwUU0hBS0VOIE5lbWVyYWxkIDMyNkswWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATGxUQVw3o7S1S3oakRHEuDonvn%2F%2FPetBUfkNjSWMn%2F1AxOVBiKOGNA7OAHcT8QVLEXY6%2BSlqPnCp9EUkNVNF3go4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQzMjZLMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzAdBgNVHQ4EFgQUfk8hny9bZtK9YTotCHl8RnJFr7EwHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0cAMEQCHz%2F4SKuNleddW%2BV%2FP3rOcDg1dvbQUOp%2Fpx%2BxE3gRyhACIQCMXLIdxRqDubTTn0jA4x8B0JVgAZtfRD7qhuF2KxqSEQ%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 326K' |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 15 Nov 23 16:51 UTC