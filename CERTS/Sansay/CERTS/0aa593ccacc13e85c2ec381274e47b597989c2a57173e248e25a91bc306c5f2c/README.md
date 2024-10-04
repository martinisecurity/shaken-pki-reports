# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Cloud Connect LLC 455K

Tested At: 04 Oct 24 16:23 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: -140 day(s)\
Subject: CN=SHAKEN Cloud Connect LLC 455K, OU=1234, O=Cloud Connect LLC, ST=New York, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://sproxy1.teleserosuite.com/Cloud_Connect_LLC_2023

[View certificate details](https://x509.io/?cert=MIIC1jCCAn2gAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkcPwwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDUxODAxMTIzMVoXDTI0MDUxNzAxMTIzMVowczELMAkGA1UEBhMCVVMxETAPBgNVBAgMCE5ldyBZb3JrMRowGAYDVQQKDBFDbG91ZCBDb25uZWN0IExMQzENMAsGA1UECwwEMTIzNDEmMCQGA1UEAwwdU0hBS0VOIENsb3VkIENvbm5lY3QgTExDIDQ1NUswWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAT2288nk%2FkcDnWoCRklTfaBmEIz8gnYjQ7npZd6psN%2FYCoar5XKmnqvc%2Ba36b%2Bk9O2eEupnOALCOGcx8OV%2BKWato4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQ0NTVLMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzAdBgNVHQ4EFgQUd8dSGcSzvg9BAm0e5WA9yy576wkwHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0cAMEQCIACx7qEBj0BNlEboEYxhxDm03oe07RXmjLmBXmN8UggTAiA%2Fiti1LIPySvdNzADQO4U15twPM2wcJEyPSwj99NvejA%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 455K', but common name is 'SHAKEN Cloud Connect LLC 455K' |


Generated: 04 Oct 24 16:29 UTC