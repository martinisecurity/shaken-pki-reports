# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Global Net Holdings Inc 306K

Tested At: 06 Jul 23 14:04 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -73 day(s)\
Subject: CN=SHAKEN Global Net Holdings Inc 306K, OU=Global Net Holdings Inc, O=Global Net Holdings Inc, ST=Pennsylvania, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/globalnetholdings

[View certificate details](https://understandingwebpki.com/?cert=MIIC%2BjCCAqGgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkZtcwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDMyNDE2MDEzOVoXDTIzMDQyMzE2MDEzOVowgZYxCzAJBgNVBAYTAlVTMRUwEwYDVQQIDAxQZW5uc3lsdmFuaWExIDAeBgNVBAoMF0dsb2JhbCBOZXQgSG9sZGluZ3MgSW5jMSAwHgYDVQQLDBdHbG9iYWwgTmV0IEhvbGRpbmdzIEluYzEsMCoGA1UEAwwjU0hBS0VOIEdsb2JhbCBOZXQgSG9sZGluZ3MgSW5jIDMwNkswWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATmxeB13E%2BZH%2F%2BI2j0rfV9mR%2BEFlQmiIgcrtBGrCdmB4HscOMxgGgvyBY5LUoNtNOWF8jQX8iOQN1Me%2BbSMhxJOo4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQzMDZLMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzAdBgNVHQ4EFgQUULF7tYMRQGMTI8E9qotiYc6u0mIwHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0cAMEQCID0zz50s%2Bfa1mSO%2FtuPexSC1jLoQKoGpQsEFzqS5%2FaWXAiBGxlJYP%2BRjYDzIYLTO7KGegIOIPeNqp1lgIslHgWX85Q%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 306K' |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 06 Jul 23 14:08 UTC