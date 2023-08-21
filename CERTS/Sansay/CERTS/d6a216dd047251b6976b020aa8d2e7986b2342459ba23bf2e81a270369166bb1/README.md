# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Zito Media Voice 624G

Tested At: 21 Aug 23 20:14 UTC\
Initial Validity Period: 90 day(s)\
Remaining Validity Period: 35 day(s)\
Subject: CN=SHAKEN Zito Media Voice 624G, OU=Voice, O=Zito Media Voice, ST=Pennsylvania, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/Zito_Media_Voice_624G

[View certificate details](https://understandingwebpki.com/?cert=MIIC2jCCAoCgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkdrswCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDYyNzE3NTgwOFoXDTIzMDkyNTE3NTgwOFowdjELMAkGA1UEBhMCVVMxFTATBgNVBAgMDFBlbm5zeWx2YW5pYTEZMBcGA1UECgwQWml0byBNZWRpYSBWb2ljZTEOMAwGA1UECwwFVm9pY2UxJTAjBgNVBAMMHFNIQUtFTiBaaXRvIE1lZGlhIFZvaWNlIDYyNEcwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASYlk9T3pbwNeumtCcdbKBXB2tj1CqKVkr8gnLTwLovGeo2XQKhIbsp0vTniJm97L1tJUo0%2FadvYl%2B4fAu2y%2B%2BRo4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQ2MjRHMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzAdBgNVHQ4EFgQUC4yVGADLhziWdz163VqjXaKxeUwwHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0gAMEUCIEPaEbrgyLS1SpOM%2F3hy3goAJ%2FoeXvLVBdxuS6JuJWRWAiEA3PT7Rg5ULtFAJnA%2BQijGAdOowpDL3DZk3ECKZN1ij%2Fk%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 624G' |


Generated: 21 Aug 23 20:18 UTC