# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN VOICE1 LLC 265K

Tested At: 28 Nov 23 16:06 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 325 day(s)\
Subject: CN=SHAKEN VOICE1 LLC 265K, OU=SYSOPS, O=VOICE1 LLC, ST=Oregon, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/VOICE1_LLC_265K

[View certificate details](https://understandingwebpki.com/?cert=MIIDdzCCAx2gAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkhLgwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMTAxODIyMzYxMloXDTI0MTAxNzIyMzYxMlowZTELMAkGA1UEBhMCVVMxDzANBgNVBAgMBk9yZWdvbjETMBEGA1UECgwKVk9JQ0UxIExMQzEPMA0GA1UECwwGU1lTT1BTMR8wHQYDVQQDDBZTSEFLRU4gVk9JQ0UxIExMQyAyNjVLMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEFYowbc3rJlnoFdIRlMt1OPx5vxLgY89hMvdatRp76U878ejv6Jswu%2Fa1UobtvBBbEr8Zg%2FQIzKsA971kib6h6aOCAYgwggGEMBYGCCsGAQUFBwEaBAowCKAGFgQyNjVLMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQUHzY6CEjqtgqJNTkVa9Qisj88RIAwgcoGA1UdIwSBwjCBv4AUrNOT9UNDzAq%2BRVgXE32SfNzDAUahgZCkgY0wgYoxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRIwEAYDVQQHDAlTYW4gRGllZ28xGzAZBgNVBAoMElNhbnNheSBDb3Jwb3JhdGlvbjESMBAGA1UECwwJU2Fuc2F5IENBMSEwHwYDVQQDDBhTSEFLRU4gU2Fuc2F5IFJvb3QgQ0EgVVOCFBS1XzgF9fB7E7X4sN7tIPJRcD6cMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNIADBFAiBFeVchRXWO%2BgLvy3jeR2tuMqRrsjyBDAVn10Qje99DUQIhAKjhNQ4SDePbnc1F7NsGJh4H83G7PN2Y4jlbE4KqYB4k)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 265K', but common name is 'SHAKEN VOICE1 LLC 265K' |
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension contains an invalid OID value: 2.16.840.1.114569.1.1.1. Available OIDs: 2.16.840.1.114569.1.1.3, 2.16.840.1.114569.1.1.4 |


Generated: 28 Nov 23 16:15 UTC