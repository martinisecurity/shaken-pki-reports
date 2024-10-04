# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN TeleVoIPs 138K

Tested At: 04 Oct 24 16:15 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: -3 day(s)\
Subject: CN=SHAKEN TeleVoIPs 138K, OU=Corp, O=TeleVoIPs, ST=Florida, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/TeleVoIPs_138K

[View certificate details](https://x509.io/?cert=MIICxTCCAmygAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkgpgwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMTAwMjE1MjI1OVoXDTI0MTAwMTE1MjI1OVowYjELMAkGA1UEBhMCVVMxEDAOBgNVBAgMB0Zsb3JpZGExEjAQBgNVBAoMCVRlbGVWb0lQczENMAsGA1UECwwEQ29ycDEeMBwGA1UEAwwVU0hBS0VOIFRlbGVWb0lQcyAxMzhLMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEi%2F5ouSX5xZkj3XnR9lxs1TdEZJU5AhaR6PgjPYxACDbSP%2B3w41mwbosq%2BsQBFuMHHOIsTRv0zRZHgstqBc8MUaOB2zCB2DAWBggrBgEFBQcBGgQKMAigBhYEMTM4SzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFHq2l%2BVzU6t%2BSgU8DjMx3a0hLK9xMB8GA1UdIwQYMBaAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNHADBEAiBRNsYKAY0Ra2XAWODUj1wmpWJepeQ%2BD6UH3rCi83tA3gIgbv%2BE%2F3B3161ShXrHLeazEcTnWnwh6CAljIbSzCRNCyY%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 138K', but common name is 'SHAKEN TeleVoIPs 138K' |
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension contains an invalid OID value: 2.16.840.1.114569.1.1.1. Available OIDs: 2.16.840.1.114569.1.1.3, 2.16.840.1.114569.1.1.4 |


Generated: 04 Oct 24 16:29 UTC