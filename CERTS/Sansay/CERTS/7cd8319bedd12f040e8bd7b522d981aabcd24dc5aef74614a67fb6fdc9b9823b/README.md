# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN DLS Internet Services 815J

Tested At: 21 Aug 23 20:14 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 66 day(s)\
Subject: CN=SHAKEN DLS Internet Services 815J, OU=sansaynss01.dls.net, O=DLS Internet Services, ST=Illinois, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/DLS_Internet_Services_815J

[View certificate details](https://understandingwebpki.com/?cert=MIIC7zCCApWgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkU0YwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAyNjE5NDE1MloXDTIzMTAyNjE5NDE1MlowgYoxCzAJBgNVBAYTAlVTMREwDwYDVQQIDAhJbGxpbm9pczEeMBwGA1UECgwVRExTIEludGVybmV0IFNlcnZpY2VzMRwwGgYDVQQLDBNzYW5zYXluc3MwMS5kbHMubmV0MSowKAYDVQQDDCFTSEFLRU4gRExTIEludGVybmV0IFNlcnZpY2VzIDgxNUowWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQ%2FL5g4yQSabbyTek%2FaTQM3Dzu%2F6oFmgPLKn9y%2FG%2FzDwzO5M9ob4U2nuciqD9IEUephnlJ7Y2jFO5AY%2B69Dtwkso4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQ4MTVKMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQU0wyRK%2BVLlV11Hq0JUj2f2Mu1gIEwHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0gAMEUCIBEY28RbukA4A9mPoqGrcREWxPrAJt0a5b2bJISDg9rTAiEAteHow2uYWQdh2uTjvnnv25A3qiWn%2B8%2BXNRosHWjFbyo%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 815J' |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |


Generated: 21 Aug 23 20:18 UTC