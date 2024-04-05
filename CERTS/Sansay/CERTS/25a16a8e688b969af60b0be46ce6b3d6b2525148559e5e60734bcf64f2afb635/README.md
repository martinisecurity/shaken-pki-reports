# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Xchange Telecom LLC 325B

Tested At: 05 Apr 24 18:49 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 223 day(s)\
Subject: CN=SHAKEN Xchange Telecom LLC 325B, OU=Bulk Solutions STI-AS, O=Xchange Telecom LLC, ST=New York, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/325B_20211101.pem

[View certificate details](https://x509.io/?cert=MIIC7DCCApOgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkhlowCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMTExNTE4NDI0MFoXDTI0MTExNDE4NDI0MFowgYgxCzAJBgNVBAYTAlVTMREwDwYDVQQIDAhOZXcgWW9yazEcMBoGA1UECgwTWGNoYW5nZSBUZWxlY29tIExMQzEeMBwGA1UECwwVQnVsayBTb2x1dGlvbnMgU1RJLUFTMSgwJgYDVQQDDB9TSEFLRU4gWGNoYW5nZSBUZWxlY29tIExMQyAzMjVCMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE0ufYApoQnxM2ww8zQ%2Be%2BEyvMooxbmK9ISQFoyTcGN6WoPxmvIXWBGBu1zv3aKBksXVrKOdAsE5j3ONV%2BWqorXqOB2zCB2DAWBggrBgEFBQcBGgQKMAigBhYEMzI1QjAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFMQWXI5nFolUBJpN5j%2BbQWR3TIohMB8GA1UdIwQYMBaAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNHADBEAiBPoepnVEDUvali%2BgxETQ6%2Brp%2BS4eN7zpq%2BAiMUTgxX7wIgQroEgokhgCmYMVWCYv1il7iJQ9JJbmou4JivFZBAnuQ%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 325B', but common name is 'SHAKEN Xchange Telecom LLC 325B' |
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension contains an invalid OID value: 2.16.840.1.114569.1.1.1. Available OIDs: 2.16.840.1.114569.1.1.3, 2.16.840.1.114569.1.1.4 |


Generated: 05 Apr 24 19:04 UTC