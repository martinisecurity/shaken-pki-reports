# STIR/SHAKEN CA Ecosystem Compliance
## Sansay

### Certificate 8d5e33ba314af37c9f8f04476e78a09325e9148f
Tested At: 2022-10-26 20:21:41 +0000 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 15 day(s)\
Subject: CN=SHAKEN Telcentris Inc. dba Voxox 696J, OU=NOC, O=Telcentris Inc. dba Voxox, ST=California, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US

Link: https://cr.sansay.com/Voxox_696J

View: [Click to view](https://understandingwebpki.com/?cert=MIIC6TCCAo%2BgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkTxAwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAxMTE3MTMzOFoXDTIyMTExMDE3MTMzOFowgYQxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMSIwIAYDVQQKDBlUZWxjZW50cmlzIEluYy4gZGJhIFZveG94MQwwCgYDVQQLDANOT0MxLjAsBgNVBAMMJVNIQUtFTiBUZWxjZW50cmlzIEluYy4gZGJhIFZveG94IDY5NkowWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATbHNz%2B%2Bg%2B%2F6gYnvUjtzRASJlG0qJ7Mu1NeSe3Wr5t15kO7M%2BARvH7%2BoaYbfls69EoqBa97VOyAjrFOdqDH%2FOuEo4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQ2OTZKMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQU0NP78Nxc95rRRl%2FkhaK3%2FB93ZawwHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0gAMEUCIFHNztbAnHsg15VXMWs7yoifH1RX108HejhH9zgOCmYEAiEAtzHdeZCoeOWDqj4aG2ycScP181EZvo3PT5E2N1IZXK0%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_sti_crl_distribution | error | ATIS-1000080v4 | CRL Distribution Point shall be reachable if the requesting IP address within the program ACLs |
| e_cp1_3_subject_sn | error | CPv1.3 | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| e_cp1_3_ambiguous_identifier | error | CPv1.3 | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| e_sti_subject_cn | error | ATIS-1000080v4 | Common name shall contain the text string 'SHAKEN 696J' |
| w_cp1_3_subject_rdn_unknown | warn | CPv1.3 | Only CN, C, O, and SERIALNUMBER can be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| e_sti_certificate_policies | error | ATIS-1000080v4 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |

Generated: 26/10/2022 at 20:22:11