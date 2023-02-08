# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN NETRIO LLC 020K

Tested At: 08 Feb 23 19:35 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 9 day(s)\
Subject: CN=SHAKEN NETRIO LLC 020K, OU=NOC, O=NETRIO LLC, ST=Texas, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/NETRIO_LLC_020K

[View certificate details](https://understandingwebpki.com/?cert=MIICxTCCAmugAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkXcowCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDExNzIxNDYxM1oXDTIzMDIxNjIxNDYxM1owYTELMAkGA1UEBhMCVVMxDjAMBgNVBAgMBVRleGFzMRMwEQYDVQQKDApORVRSSU8gTExDMQwwCgYDVQQLDANOT0MxHzAdBgNVBAMMFlNIQUtFTiBORVRSSU8gTExDIDAyMEswWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATIwd%2FKwmHggCNfmSkcdHSrEQBSTDC5pLX1W6gA2%2BBf3sHszti9%2BYZTMf8XgJbHRLMylFri3hkPLYm6sSHCB%2FnUo4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQwMjBLMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQUkZ%2BnrUwU%2BUBQuA3%2FS9Co2pJh6qcwHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0gAMEUCIEeHOtoZ608Ysw3V6VzclCCETaB50%2B0AHaQ3R%2BQb5VfUAiEA%2BIv9pELu9i4SWYR73XZKvmXn9bMUSL8eSbOGg8MaSFI%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 020K' |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |


Generated: 08 Feb 23 19:45 UTC