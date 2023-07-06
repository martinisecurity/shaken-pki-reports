# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Cyberlynk Network, LLC 086K

Tested At: 06 Jul 23 13:59 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 0 day(s)\
Subject: emailAddress=software@cyberlynk.net, CN=SHAKEN Cyberlynk Network\\, LLC 086K, OU=NOC, O=Cyberlynk Network\\, LLC, ST=Winsconsin, C=US, emailAddress=software@cyberlynk.net\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/086K/order/25_086K_91

[View certificate details](https://understandingwebpki.com/?cert=MIIDCjCCArCgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkc5wwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDYwNjAxMDQwMFoXDTIzMDcwNjAxMDQwMFowgaUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApXaW5zY29uc2luMR8wHQYDVQQKDBZDeWJlcmx5bmsgTmV0d29yaywgTExDMQwwCgYDVQQLDANOT0MxKzApBgNVBAMMIlNIQUtFTiBDeWJlcmx5bmsgTmV0d29yaywgTExDIDA4NksxJTAjBgkqhkiG9w0BCQEWFnNvZnR3YXJlQGN5YmVybHluay5uZXQwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATpkCFA8LvfWu1asIDP6BOFqM1qDdJaPFG22XbgahglidWLpWI5TFVs9ve8O2SFujBOe8D%2FL3v33OCGvktdD0iPo4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQwODZLMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzAdBgNVHQ4EFgQULPPQSRAOl8C1R02jCqCfhikI2SswHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0gAMEUCIQCEB5IqPK8gEv5215aua%2BTmGx4ujLGHqW37148DWaz4vwIgGaP28Mo8JP5%2FRqhy%2BFkZMEEy1BemAaQ08NT2CYIsLgU%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [w_cp_1_3_subject_email](../../ISSUES/w_cp_1_3_subject_email/README.md) | warn | US_SHAKEN_CP | Email addresses are not allowed as the CP does not specify how to validate them |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 086K' |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 06 Jul 23 14:08 UTC