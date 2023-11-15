# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Telcentris Inc. dba Voxox 696J

Tested At: 15 Nov 23 16:34 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 2 day(s)\
Subject: CN=SHAKEN Telcentris Inc. dba Voxox 696J, OU=NOC, O=Telcentris Inc. dba Voxox, ST=California, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/Voxox_696J_1

[View certificate details](https://understandingwebpki.com/?cert=MIIC6TCCAo%2BgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkhHMwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMTAxNzIyMzIyNFoXDTIzMTExNjIyMzIyNFowgYQxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMSIwIAYDVQQKDBlUZWxjZW50cmlzIEluYy4gZGJhIFZveG94MQwwCgYDVQQLDANOT0MxLjAsBgNVBAMMJVNIQUtFTiBUZWxjZW50cmlzIEluYy4gZGJhIFZveG94IDY5NkowWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATbHNz%2B%2Bg%2B%2F6gYnvUjtzRASJlG0qJ7Mu1NeSe3Wr5t15kO7M%2BARvH7%2BoaYbfls69EoqBa97VOyAjrFOdqDH%2FOuEo4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQ2OTZKMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQU0NP78Nxc95rRRl%2FkhaK3%2FB93ZawwHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0gAMEUCIQC%2FOvrHqo7Q02QStx1O3e%2Fe3tq9Yx4kaH7bymKXP7IMaQIgTLWdwoYo2%2FH%2B8CNQK4JbJXp57%2FWZQ33D31FugflFI1o%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 696J' |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 15 Nov 23 17:17 UTC