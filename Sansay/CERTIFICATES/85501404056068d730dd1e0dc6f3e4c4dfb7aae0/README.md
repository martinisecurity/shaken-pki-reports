# STIR/SHAKEN CA Ecosystem Compliance
## Sansay

### Certificate 85501404056068d730dd1e0dc6f3e4c4dfb7aae0
Tested At: 2022-10-27 18:56:24 +0000 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 30 day(s)\
Subject: CN=SHAKEN NETRIO LLC 020K, OU=NOC, O=NETRIO LLC, ST=Texas, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US

Link: https://cr.sansay.com/NETRIO_LLC_020K

View: [Click to view](https://understandingwebpki.com/?cert=MIICxDCCAmugAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkU30wCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAyNzAwMzgwOVoXDTIyMTEyNjAwMzgwOVowYTELMAkGA1UEBhMCVVMxDjAMBgNVBAgMBVRleGFzMRMwEQYDVQQKDApORVRSSU8gTExDMQwwCgYDVQQLDANOT0MxHzAdBgNVBAMMFlNIQUtFTiBORVRSSU8gTExDIDAyMEswWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATIwd%2FKwmHggCNfmSkcdHSrEQBSTDC5pLX1W6gA2%2BBf3sHszti9%2BYZTMf8XgJbHRLMylFri3hkPLYm6sSHCB%2FnUo4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQwMjBLMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQUkZ%2BnrUwU%2BUBQuA3%2FS9Co2pJh6qcwHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0cAMEQCIGwZsPMgTCdga9iOFjWwmmC%2FfY6E413b6ipgV7S6PeOWAiB2DXKeBLc5p9r%2FEg0IgZvTgqdJWFLRW0VRcmeBSLEA5g%3D%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_cp1_3_subject_sn | error | CPv1.3 | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| e_sti_subject_cn | error | ATIS-1000080v4 | Common name shall contain the text string 'SHAKEN 020K' |
| e_cp1_3_ambiguous_identifier | error | CPv1.3 | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| e_sti_crl_distribution | error | ATIS-1000080v4 | CRL Distribution Point shall be reachable if the requesting IP address within the program ACLs |
| w_cp1_3_subject_rdn_unknown | warn | CPv1.3 | Only CN, C, O, and SERIALNUMBER can be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| e_sti_certificate_policies | error | ATIS-1000080v4 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |

Generated: 27/10/2022 at 18:57:27