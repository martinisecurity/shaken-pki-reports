# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Peachnet LLC 616K

Tested At: 15 Nov 23 18:03 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 219 day(s)\
Subject: CN=SHAKEN Peachnet LLC 616K, OU=Peachnet LLC, O=Peachnet LLC, ST=Georgia, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/Peachnet_616K

[View certificate details](https://understandingwebpki.com/?cert=MIIC1TCCAnqgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkdjMwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDYyMjE0Mjk0M1oXDTI0MDYyMTE0Mjk0M1owcDELMAkGA1UEBhMCVVMxEDAOBgNVBAgMB0dlb3JnaWExFTATBgNVBAoMDFBlYWNobmV0IExMQzEVMBMGA1UECwwMUGVhY2huZXQgTExDMSEwHwYDVQQDDBhTSEFLRU4gUGVhY2huZXQgTExDIDYxNkswWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATXloyfiPhMOn%2BGHZF%2BwDmuIbDEY%2BtfY72in20Z3%2Bx3i8t6FIZVnsR1ReJChmmWLU4efHFtSrECCcRlq%2F5D9mejo4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQ2MTZLMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzAdBgNVHQ4EFgQU%2BMFXVDBazXEUdc1NJMcy0bNs67cwHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0kAMEYCIQDsR5GXOVdmlHbzwIGY4MjSpfdV7PQG%2BJ9XPKJ%2B%2BiUBvAIhAOfN4qr%2F58etolJfTablI6b3FcwTzrIJEjoW%2F0u5nRTb)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 616K' |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |


Generated: 15 Nov 23 18:10 UTC