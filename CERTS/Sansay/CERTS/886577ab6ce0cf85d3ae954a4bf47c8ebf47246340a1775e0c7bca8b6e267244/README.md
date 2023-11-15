# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN ConvergeTel LLC 388K

Tested At: 15 Nov 23 16:34 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 350 day(s)\
Subject: CN=SHAKEN ConvergeTel LLC 388K, OU=voice, O=ConvergeTel LLC, ST=Delaware, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/ConvergeTel_LLC_388K

[View certificate details](https://understandingwebpki.com/?cert=MIIC0zCCAnqgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkhd4wCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMTAzMTE2MjM0MFoXDTI0MTAzMDE2MjM0MFowcDELMAkGA1UEBhMCVVMxETAPBgNVBAgMCERlbGF3YXJlMRgwFgYDVQQKDA9Db252ZXJnZVRlbCBMTEMxDjAMBgNVBAsMBXZvaWNlMSQwIgYDVQQDDBtTSEFLRU4gQ29udmVyZ2VUZWwgTExDIDM4OEswWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQCOIPK0FJhgE16fsbgF690HA6sk8y2f4oeQb2MtnroLI2LPVE9UbVT%2FrtGJJlUk66SfdXH1kPzKMvRi8WmpZD5o4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQzODhLMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzAdBgNVHQ4EFgQUHNqatw9fTz%2BCzMTcXAVSfTejZ9kwHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0cAMEQCIAm7kHWzuHXIGfOKr%2BCKN4Hb%2Bz%2BLYQp%2FwTOvdG6tqenfAiAySz8XMBXfp8PKYvrt%2Fua5nkxi%2B1C%2FfFJtjTGzpS4WcA%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 388K' |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 15 Nov 23 17:17 UTC