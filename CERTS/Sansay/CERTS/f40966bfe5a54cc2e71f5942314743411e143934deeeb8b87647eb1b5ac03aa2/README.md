# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN VOICE1 LLC 265K

Tested At: 15 Nov 23 16:34 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 338 day(s)\
Subject: CN=SHAKEN VOICE1 LLC 265K, OU=SYSOPS, O=VOICE1 LLC, ST=Oregon, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/VOICE1_LLC_265K

[View certificate details](https://understandingwebpki.com/?cert=MIIDdzCCAx2gAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkhLgwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMTAxODIyMzYxMloXDTI0MTAxNzIyMzYxMlowZTELMAkGA1UEBhMCVVMxDzANBgNVBAgMBk9yZWdvbjETMBEGA1UECgwKVk9JQ0UxIExMQzEPMA0GA1UECwwGU1lTT1BTMR8wHQYDVQQDDBZTSEFLRU4gVk9JQ0UxIExMQyAyNjVLMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEFYowbc3rJlnoFdIRlMt1OPx5vxLgY89hMvdatRp76U878ejv6Jswu%2Fa1UobtvBBbEr8Zg%2FQIzKsA971kib6h6aOCAYgwggGEMBYGCCsGAQUFBwEaBAowCKAGFgQyNjVLMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQUHzY6CEjqtgqJNTkVa9Qisj88RIAwgcoGA1UdIwSBwjCBv4AUrNOT9UNDzAq%2BRVgXE32SfNzDAUahgZCkgY0wgYoxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRIwEAYDVQQHDAlTYW4gRGllZ28xGzAZBgNVBAoMElNhbnNheSBDb3Jwb3JhdGlvbjESMBAGA1UECwwJU2Fuc2F5IENBMSEwHwYDVQQDDBhTSEFLRU4gU2Fuc2F5IFJvb3QgQ0EgVVOCFBS1XzgF9fB7E7X4sN7tIPJRcD6cMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNIADBFAiBFeVchRXWO%2BgLvy3jeR2tuMqRrsjyBDAVn10Qje99DUQIhAKjhNQ4SDePbnc1F7NsGJh4H83G7PN2Y4jlbE4KqYB4k)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 265K' |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 15 Nov 23 17:17 UTC