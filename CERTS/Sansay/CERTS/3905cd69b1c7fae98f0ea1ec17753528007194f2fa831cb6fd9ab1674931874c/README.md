# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Cyberlynk Network, LLC 086K

Tested At: 31 Jan 23 21:41 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 5 day(s)\
Subject: emailAddress=software@cyberlynk.net, CN=SHAKEN Cyberlynk Network\\, LLC 086K, OU=NOC, O=Cyberlynk Network\\, LLC, ST=Winsconsin, C=US, emailAddress=software@cyberlynk.net\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/086K/order/203_086K_91

[View certificate details](https://understandingwebpki.com/?cert=MIIDCjCCArCgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkW1kwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDEwNjE0NDI1NloXDTIzMDIwNTE0NDI1NlowgaUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApXaW5zY29uc2luMR8wHQYDVQQKDBZDeWJlcmx5bmsgTmV0d29yaywgTExDMQwwCgYDVQQLDANOT0MxKzApBgNVBAMMIlNIQUtFTiBDeWJlcmx5bmsgTmV0d29yaywgTExDIDA4NksxJTAjBgkqhkiG9w0BCQEWFnNvZnR3YXJlQGN5YmVybHluay5uZXQwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATpkCFA8LvfWu1asIDP6BOFqM1qDdJaPFG22XbgahglidWLpWI5TFVs9ve8O2SFujBOe8D%2FL3v33OCGvktdD0iPo4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQwODZLMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQULPPQSRAOl8C1R02jCqCfhikI2SswHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0gAMEUCIGpH57ThScRByfXkeszwThquslDX5QREOv7InQ3ECjH1AiEAn0wM7j9ZDaZHuHB1Cu11wGZ7zchz%2FJuZGZlJW5VGJgI%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 086K' |
| [w_cp_1_3_subject_email](../../ISSUES/w_cp_1_3_subject_email/README.md) | warn | US_SHAKEN_CP | Email addresses are not allowed as the CP does not specify how to validate them |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |


Generated: 31 Jan 23 21:50 UTC