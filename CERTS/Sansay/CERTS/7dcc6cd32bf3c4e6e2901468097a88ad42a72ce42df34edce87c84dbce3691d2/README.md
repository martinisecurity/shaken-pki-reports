# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Bulk Solutions, LLC 644J

Tested At: 28 Nov 23 19:52 UTC\
Initial Validity Period: 350 day(s)\
Remaining Validity Period: 44 day(s)\
Subject: CN=SHAKEN Bulk Solutions\\, LLC 644J, OU=608070, O=Bulk Solutions\\, LLC, ST=Delaware, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: http://mango.voipplus.net/cert.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC3TCCAoOgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkXxUwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDEyNjE0MjYzNloXDTI0MDExMTE0MjYzNloweTELMAkGA1UEBhMCVVMxETAPBgNVBAgMCERlbGF3YXJlMRwwGgYDVQQKDBNCdWxrIFNvbHV0aW9ucywgTExDMQ8wDQYDVQQLDAY2MDgwNzAxKDAmBgNVBAMMH1NIQUtFTiBCdWxrIFNvbHV0aW9ucywgTExDIDY0NEowWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASP9%2FuRZZfL5uYK8tgNbuVy5w7aBavi1nQxPYEFVcAGS8UbiWGIoRWDTEs9eS528mF7P09D5ez3JKwfobYlR1Hto4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQ2NDRKMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzAdBgNVHQ4EFgQUa%2FicET9ftsA5%2FWackOfEjPntGjMwHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0gAMEUCIQD5UWy4h172GHtXeAwUcDSz4j%2Fkb6D8cYte6AZuaI1JKgIgJMFfvHV%2F9HjQPx%2FxExg5CK9CzrEvba%2BWJFHzAvpIoV0%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 644J', but common name is 'SHAKEN Bulk Solutions, LLC 644J' |
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 28 Nov 23 20:21 UTC