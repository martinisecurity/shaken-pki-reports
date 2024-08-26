# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Swift Telco LLC 452K

Tested At: 26 Aug 24 17:42 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -487 day(s)\
Subject: CN=SHAKEN Swift Telco LLC 452K, OU=ST124, O=Swift Telco LLC, ST=South Carolina, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cdn.cnxcdn.com/shaken/60.crt

[View certificate details](https://x509.io/?cert=MIIC2TCCAoCgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkZ3IwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDMyNzIzMTExM1oXDTIzMDQyNjIzMTExM1owdjELMAkGA1UEBhMCVVMxFzAVBgNVBAgMDlNvdXRoIENhcm9saW5hMRgwFgYDVQQKDA9Td2lmdCBUZWxjbyBMTEMxDjAMBgNVBAsMBVNUMTI0MSQwIgYDVQQDDBtTSEFLRU4gU3dpZnQgVGVsY28gTExDIDQ1MkswWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAARLheqeb8xtlwm0FGHKko9AH9or5fYrENzE2fwFBbZteQiYu8xRusLSgYdMWXLovVuvafBvmekMWsYQ6PcsvPCeo4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQ0NTJLMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzAdBgNVHQ4EFgQU2TizrjviN0vAbh5aPrzCLLqT2EUwHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0cAMEQCIHINllKwRY%2FXEjjt%2FBWNI5tF46Mpl%2FkLd10xNiSSoZ5hAiBFDSE9o1pJe49fTeVXIZiiMHykE6ZrsNC7JmWKH5lKBw%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 452K', but common name is 'SHAKEN Swift Telco LLC 452K' |
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |


Generated: 26 Aug 24 18:03 UTC