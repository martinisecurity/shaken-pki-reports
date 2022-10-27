# STIR/SHAKEN CA Ecosystem Compliance
## Sansay

### Certificate d35c00b39b9dc470c5324e76a20695cba26a2bf3
Tested At: 2022-10-27 21:42:45 +0000 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 30 day(s)\
Subject: CN=SHAKEN Telcentris Inc. dba Voxox 696J, OU=NOC, O=Telcentris Inc. dba Voxox, ST=California, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US

Link: https://cr.sansay.com/Voxox_696J

View: [Click to view](https://understandingwebpki.com/?cert=MIIC6TCCAo%2BgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkU4EwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAyNzAwMzk0MloXDTIyMTEyNjAwMzk0MlowgYQxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMSIwIAYDVQQKDBlUZWxjZW50cmlzIEluYy4gZGJhIFZveG94MQwwCgYDVQQLDANOT0MxLjAsBgNVBAMMJVNIQUtFTiBUZWxjZW50cmlzIEluYy4gZGJhIFZveG94IDY5NkowWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATbHNz%2B%2Bg%2B%2F6gYnvUjtzRASJlG0qJ7Mu1NeSe3Wr5t15kO7M%2BARvH7%2BoaYbfls69EoqBa97VOyAjrFOdqDH%2FOuEo4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQ2OTZKMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQU0NP78Nxc95rRRl%2FkhaK3%2FB93ZawwHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0gAMEUCID0jfHGXpOC2%2BAo0OX7gldprZKasP1nF5R2%2BPBWc2eFHAiEAyv3ouC1Fq77P%2F80fpII3juz3JR5ELFOjWVLK31yPQJo%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_cp1_3_ambiguous_identifier | error | United States SHAKEN CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| w_cp1_3_subject_rdn_unknown | warn | United States SHAKEN CP | Only CN, C, O, and SERIALNUMBER can be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| e_cp1_3_subject_sn | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| e_sti_subject_cn | error | ATIS-1000080 | Common name shall contain the text string 'SHAKEN 696J' |
| e_sti_certificate_policies | error | ATIS-1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |

Generated: 27/10/2022 at 21:42:52