# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Swift Telco LLC 452K

Tested At: 26 Aug 24 17:43 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: -144 day(s)\
Subject: CN=SHAKEN Swift Telco LLC 452K, OU=ST124new, O=Swift Telco LLC, ST=South Carolina, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cdn.cnxcdn.com/shaken/c015310d91.crt

[View certificate details](https://x509.io/?cert=MIIC3TCCAoOgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkaRUwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDQwNTE2Mjc1NloXDTI0MDQwNDE2Mjc1NloweTELMAkGA1UEBhMCVVMxFzAVBgNVBAgMDlNvdXRoIENhcm9saW5hMRgwFgYDVQQKDA9Td2lmdCBUZWxjbyBMTEMxETAPBgNVBAsMCFNUMTI0bmV3MSQwIgYDVQQDDBtTSEFLRU4gU3dpZnQgVGVsY28gTExDIDQ1MkswWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAARXEhAbNp4ks0J%2FCB7AWt%2FlFXdBPRaoJ7CWpbbncrBW7dAa3YBSByH%2BD49pjUblhnet4bIw%2FZoqh%2FL%2BNN6l%2BL06o4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQ0NTJLMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzAdBgNVHQ4EFgQU3kheTu7W%2B7i%2F6LkHHU%2B2FeL24OAwHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0gAMEUCIBe7Gc7XUvm0fO4BjCBOWx3LWFON0tAnz9qJ6THUFivlAiEA9jDZO8ECpgugl24%2B%2F4tynEWtlYfPpptwGeKTOjoXsTw%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 452K', but common name is 'SHAKEN Swift Telco LLC 452K' |
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |


Generated: 26 Aug 24 18:03 UTC