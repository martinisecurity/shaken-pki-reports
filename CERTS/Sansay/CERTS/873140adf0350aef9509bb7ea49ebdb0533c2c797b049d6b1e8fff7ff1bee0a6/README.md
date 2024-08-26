# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN NETRIO LLC 020K

Tested At: 26 Aug 24 18:15 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -51 day(s)\
Subject: CN=SHAKEN NETRIO LLC 020K, OU=NOC, O=NETRIO LLC, ST=Texas, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/020K/429C7C70711E3820F0B8E1DEAE6FF32622649EF5.pem

[View certificate details](https://x509.io/?cert=MIICxjCCAmugAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJknvUwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTI0MDYwNTIyNTkzOFoXDTI0MDcwNTIyNTkzOFowYTELMAkGA1UEBhMCVVMxDjAMBgNVBAgMBVRleGFzMRMwEQYDVQQKDApORVRSSU8gTExDMQwwCgYDVQQLDANOT0MxHzAdBgNVBAMMFlNIQUtFTiBORVRSSU8gTExDIDAyMEswWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATIwd%2FKwmHggCNfmSkcdHSrEQBSTDC5pLX1W6gA2%2BBf3sHszti9%2BYZTMf8XgJbHRLMylFri3hkPLYm6sSHCB%2FnUo4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQwMjBLMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQUkZ%2BnrUwU%2BUBQuA3%2FS9Co2pJh6qcwHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0kAMEYCIQC3yNk78sPunw19ubokBp10vyvha%2BNWlhW3kfvqEakyRAIhAIy5f266RZhqC99zrRzxYm6%2FJIirNpLtlOkXPKw7W7ob)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 020K', but common name is 'SHAKEN NETRIO LLC 020K' |
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension contains an invalid OID value: 2.16.840.1.114569.1.1.1. Available OIDs: 2.16.840.1.114569.1.1.3, 2.16.840.1.114569.1.1.4 |


Generated: 26 Aug 24 18:49 UTC