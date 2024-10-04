# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Matrix 3058

Tested At: 04 Oct 24 15:58 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -90 day(s)\
Subject: CN=SHAKEN Matrix 3058, OU=NOC, O=Matrix, ST=Texas, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/3058/429C7C70711E3820F0B8E1DEAE6FF32622649F00.pem

[View certificate details](https://x509.io/?cert=MIICvjCCAmOgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJknwAwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTI0MDYwNTIzMTYwOFoXDTI0MDcwNTIzMTYwOFowWTELMAkGA1UEBhMCVVMxDjAMBgNVBAgMBVRleGFzMQ8wDQYDVQQKDAZNYXRyaXgxDDAKBgNVBAsMA05PQzEbMBkGA1UEAwwSU0hBS0VOIE1hdHJpeCAzMDU4MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEzUotEqptZc4M0%2Fx0MfF15yWz8%2Fntg4VDtUOPInSqs1AZUeQf2VfshGFf2LsxKq6gTQRS2LkM%2B79lJYpEztDLyKOB2zCB2DAWBggrBgEFBQcBGgQKMAigBhYEMzA1ODAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFAJFOlo0TAnsTivHURmXKZS0qDDWMB8GA1UdIwQYMBaAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNJADBGAiEAohhr0%2BYjXm0F12YQjBs1HCLMkk6GiK59Ff%2Bcwk%2B6VIwCIQCuXaanoPQfUE8ZQpsprUb%2Bz7niOO4OWYlL%2FqafZB8C7Q%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 3058', but common name is 'SHAKEN Matrix 3058' |
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension contains an invalid OID value: 2.16.840.1.114569.1.1.1. Available OIDs: 2.16.840.1.114569.1.1.3, 2.16.840.1.114569.1.1.4 |


Generated: 04 Oct 24 16:29 UTC