# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Bulk Solutions, LLC 644J

Tested At: 12 Apr 23 01:01 UTC\
Initial Validity Period: 350 day(s)\
Remaining Validity Period: 275 day(s)\
Subject: CN=SHAKEN Bulk Solutions\\, LLC 644J, OU=608070, O=Bulk Solutions\\, LLC, ST=Delaware, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: http://mango.voipplus.net/cert.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC3TCCAoOgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkXxUwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDEyNjE0MjYzNloXDTI0MDExMTE0MjYzNloweTELMAkGA1UEBhMCVVMxETAPBgNVBAgMCERlbGF3YXJlMRwwGgYDVQQKDBNCdWxrIFNvbHV0aW9ucywgTExDMQ8wDQYDVQQLDAY2MDgwNzAxKDAmBgNVBAMMH1NIQUtFTiBCdWxrIFNvbHV0aW9ucywgTExDIDY0NEowWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASP9%2FuRZZfL5uYK8tgNbuVy5w7aBavi1nQxPYEFVcAGS8UbiWGIoRWDTEs9eS528mF7P09D5ez3JKwfobYlR1Hto4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQ2NDRKMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzAdBgNVHQ4EFgQUa%2FicET9ftsA5%2FWackOfEjPntGjMwHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0gAMEUCIQD5UWy4h172GHtXeAwUcDSz4j%2Fkb6D8cYte6AZuaI1JKgIgJMFfvHV%2F9HjQPx%2FxExg5CK9CzrEvba%2BWJFHzAvpIoV0%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 644J' |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 12 Apr 23 01:46 UTC