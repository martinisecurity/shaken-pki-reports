# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Rayfield Communications, Inc. 006K

Tested At: 27 Nov 23 22:46 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 365 day(s)\
Subject: CN=SHAKEN Rayfield Communications\\, Inc. 006K, OU=VoIP, O=Rayfield Communications\\, Inc., ST=Missouri, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/Rayfield_Communications_Inc._006K

[View certificate details](https://understandingwebpki.com/?cert=MIIC7zCCApagAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkh6QwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMTEyNzE2MjAyN1oXDTI0MTEyNjE2MjAyN1owgYsxCzAJBgNVBAYTAlVTMREwDwYDVQQIDAhNaXNzb3VyaTEmMCQGA1UECgwdUmF5ZmllbGQgQ29tbXVuaWNhdGlvbnMsIEluYy4xDTALBgNVBAsMBFZvSVAxMjAwBgNVBAMMKVNIQUtFTiBSYXlmaWVsZCBDb21tdW5pY2F0aW9ucywgSW5jLiAwMDZLMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE4hYpO8vVQiNpQnjWHQoE4hVTe1rj2SEhJKqNqLAnoFtLaBh9lMy9a58h3Rz0Tr%2FkFcvyaz%2FuTsiPDDn8lBVZL6OB2zCB2DAWBggrBgEFBQcBGgQKMAigBhYEMDA2SzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFCxjV%2Fl%2BVxoY14LRjiqPwwETJtBQMB8GA1UdIwQYMBaAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNHADBEAiAW6mnUGtsVcBUS0TtD6rBH5TjoazPOi2gHtaRdadr1FwIgdiBG4R8yjonViGdwMKL%2BkKwepXOqKyFeL7QsGJH8xZI%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 006K', but common name is 'SHAKEN Rayfield Communications, Inc. 006K' |
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension contains an invalid OID value: 2.16.840.1.114569.1.1.1. Available OIDs: 2.16.840.1.114569.1.1.3, 2.16.840.1.114569.1.1.4 |


Generated: 27 Nov 23 22:56 UTC