# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN VOIP OFFICE.COM LLC 389K

Tested At: 18 Aug 25 20:50 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 103 day(s)\
Subject: CN=SHAKEN VOIP OFFICE.COM LLC 389K, OU=Communications, O=VOIP OFFICE, ST=Michigan, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://nguc.voipoffice.com/stirshaken/389K.chain.crt.pem

[View certificate details](https://x509.io/?cert=MIIC3TCCAoOgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJktGwwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTI0MTEyOTE4MTAwNFoXDTI1MTEyOTE4MTAwNFoweTELMAkGA1UEBhMCVVMxETAPBgNVBAgMCE1pY2hpZ2FuMRQwEgYDVQQKDAtWT0lQIE9GRklDRTEXMBUGA1UECwwOQ29tbXVuaWNhdGlvbnMxKDAmBgNVBAMMH1NIQUtFTiBWT0lQIE9GRklDRS5DT00gTExDIDM4OUswWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASyNJX%2FUt8BC24kxUNOea06PnB5syn2R%2FNwkXPl1D6C1RrD5Y1QmvcytXYJGLBp9g8A7NrZr6YDRFOGY9Jrk9J4o4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQzODlLMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQUpoZG1j2PdPoEsNQ5EWxpuSYUQ1YwHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0gAMEUCIQDe7rgwFPG215Cq6V5578Idr8zuNFaZkOLLPqCbgJebfgIgF0HCZcN77pqP8jpuZ86S355cVejUxyge2yTdlKVd1v8%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 389K', but common name is 'SHAKEN VOIP OFFICE.COM LLC 389K' |
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension contains an invalid OID value: 2.16.840.1.114569.1.1.1. Available OIDs: 2.16.840.1.114569.1.1.3, 2.16.840.1.114569.1.1.4 |


Generated: 18 Aug 25 21:13 UTC