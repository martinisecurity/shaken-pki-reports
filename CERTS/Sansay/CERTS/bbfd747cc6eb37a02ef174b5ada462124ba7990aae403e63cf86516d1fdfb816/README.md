# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Mercury Access Solutions 634K

Tested At: 15 Nov 23 18:03 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 230 day(s)\
Subject: CN=SHAKEN Mercury Access Solutions 634K, OU=Mercury Access Solutions, O=Mercury Access Solutions, ST=Kansas, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/MercuryBroadband_634K

[View certificate details](https://understandingwebpki.com/?cert=MIIC%2BDCCAp6gAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkd4swCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDcwMzE2MzkzMloXDTI0MDcwMjE2MzkzMlowgZMxCzAJBgNVBAYTAlVTMQ8wDQYDVQQIDAZLYW5zYXMxITAfBgNVBAoMGE1lcmN1cnkgQWNjZXNzIFNvbHV0aW9uczEhMB8GA1UECwwYTWVyY3VyeSBBY2Nlc3MgU29sdXRpb25zMS0wKwYDVQQDDCRTSEFLRU4gTWVyY3VyeSBBY2Nlc3MgU29sdXRpb25zIDYzNEswWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATV9yc89D0X32A36kZJxXSlWGqT%2BEAuDP1cecycLtEQGAy5Ww3rk7xdBmMFQVYrctYngyE5jgyWL6YgZGF0QiA%2Bo4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQ2MzRLMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzAdBgNVHQ4EFgQUa8jbUhbOdokSIf2VxE8v%2B6R0NSQwHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0gAMEUCIQCSpjSkfmtqkKmT3wIfpVFJ6fZpVYN0H0oueGGBzDEhzQIgS9BrT%2FGyxlLMlz2h0nFzFIRXAOSkHfPoocfS0Q5Y7XI%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 634K' |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 15 Nov 23 18:10 UTC