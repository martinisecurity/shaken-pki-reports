# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Miracle Telecom 496K

Tested At: 15 Nov 23 18:03 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 304 day(s)\
Subject: CN=SHAKEN Miracle Telecom 496K, OU=Operations, O=Miracle Telecom, ST=District Of Columbia, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/Miracle_Telecom_496K

[View certificate details](https://understandingwebpki.com/?cert=MIIC5TCCAoygAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkgKMwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDkxNTE0NDAwNFoXDTI0MDkxNDE0NDAwNFowgYExCzAJBgNVBAYTAlVTMR0wGwYDVQQIDBREaXN0cmljdCBPZiBDb2x1bWJpYTEYMBYGA1UECgwPTWlyYWNsZSBUZWxlY29tMRMwEQYDVQQLDApPcGVyYXRpb25zMSQwIgYDVQQDDBtTSEFLRU4gTWlyYWNsZSBUZWxlY29tIDQ5NkswWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAT6k8Ww7a3vICnZ%2FlyCoUJ30YNkz3bML6OK9n8q7shZQ9HAfgaMK8BVNbztD8SN8E023%2B67zENQa0LQ6Gq6n7%2FUo4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQ0OTZLMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzAdBgNVHQ4EFgQUGJr0vvBRp4oW4jLE8nkA26u%2BAAIwHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0cAMEQCIGlcfmNzfJun2ONpuKZ%2B%2BhbpLSvakw9p8EWegh5grHgcAiA0ZFDgkHXtIUK4qUzBANPZGU%2FPSoEAsizklDJxKHPwaA%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 496K' |


Generated: 15 Nov 23 18:10 UTC