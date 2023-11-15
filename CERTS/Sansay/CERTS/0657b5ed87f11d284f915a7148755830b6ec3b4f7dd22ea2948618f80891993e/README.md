# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN IDT America, Corp 363A

Tested At: 15 Nov 23 18:03 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -74 day(s)\
Subject: emailAddress=thomas.musella@idt.net, CN=SHAKEN IDT America\\, Corp 363A, OU=NOC, O=IDT America\\, Corp, ST=New Jersey, C=US, emailAddress=thomas.musella@idt.net\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/363A/order/571_363A_17

[View certificate details](https://understandingwebpki.com/?cert=MIIDADCCAqagAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJke7kwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDgwMzA2NDkwOVoXDTIzMDkwMjA2NDkwOVowgZsxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApOZXcgSmVyc2V5MRowGAYDVQQKDBFJRFQgQW1lcmljYSwgQ29ycDEMMAoGA1UECwwDTk9DMSYwJAYDVQQDDB1TSEFLRU4gSURUIEFtZXJpY2EsIENvcnAgMzYzQTElMCMGCSqGSIb3DQEJARYWdGhvbWFzLm11c2VsbGFAaWR0Lm5ldDBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABPQM8yPpWKxm0BJ%2BjGRI%2F%2Fn99wUDOlM0pPkwuiTqWy8TLzNxSa1qR7pJYQDwtTuzur7Db7X4HGB5yZxyIiXZWL6jgdswgdgwFgYIKwYBBQUHARoECjAIoAYWBDM2M0EwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMB0GA1UdDgQWBBRtxUb554hu2Cjs4EgzW1ORcFzRVzAfBgNVHSMEGDAWgBSs05P1Q0PMCr5FWBcTfZJ83MMBRjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSAAwRQIhAJodoah%2F0I5tbudVFw7MslkX89BTFkr5sSJw1HYchLlIAiAGRSsAEEqccucUEf7reiArGyUlsml3tfMdW5EbEV0Oow%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 363A' |
| [w_cp_1_3_subject_email](../../ISSUES/w_cp_1_3_subject_email/README.md) | warn | US_SHAKEN_CP | Email addresses are not allowed as the CP does not specify how to validate them |


Generated: 15 Nov 23 18:10 UTC