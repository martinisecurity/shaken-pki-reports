# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Broadband Dynamics LLC 583j

Tested At: 04 Oct 24 16:06 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -259 day(s)\
Subject: CN=SHAKEN Broadband Dynamics LLC 583j, OU=Operations, O=Broadband Dynamics LLC, ST=Arizona, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/583j_BROADBAND_DYNAMICS_STIR_SHAKEN.crt

[View certificate details](https://x509.io/?cert=MIIDlDCCAzugAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkioMwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMTIxOTE4NTYxN1oXDTI0MDExODE4NTYxN1owgYIxCzAJBgNVBAYTAlVTMRAwDgYDVQQIDAdBcml6b25hMR8wHQYDVQQKDBZCcm9hZGJhbmQgRHluYW1pY3MgTExDMRMwEQYDVQQLDApPcGVyYXRpb25zMSswKQYDVQQDDCJTSEFLRU4gQnJvYWRiYW5kIER5bmFtaWNzIExMQyA1ODNqMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEGvpK77E6nNmb5Ib3%2B0atBc1rvOUU0j%2FHa5SXcvxc%2F1iNaQpDhxr6Cu8bDw4U2WQdGr4vPdmitNornKsYZ0qtjaOCAYgwggGEMBYGCCsGAQUFBwEaBAowCKAGFgQ1ODNqMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQUiYHRBUDlqk8R8LsC1JaQPodW5t8wgcoGA1UdIwSBwjCBv4AUrNOT9UNDzAq%2BRVgXE32SfNzDAUahgZCkgY0wgYoxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRIwEAYDVQQHDAlTYW4gRGllZ28xGzAZBgNVBAoMElNhbnNheSBDb3Jwb3JhdGlvbjESMBAGA1UECwwJU2Fuc2F5IENBMSEwHwYDVQQDDBhTSEFLRU4gU2Fuc2F5IFJvb3QgQ0EgVVOCFBS1XzgF9fB7E7X4sN7tIPJRcD6cMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNHADBEAiBtKtQIYROnsjcPl8uYEMyF9WUwJ8ZVQR9R31F%2BFkXASAIgNb3JKnIYNJJJcMkK%2B8Lv89nLjLL3cW4lks5dlNXQ69I%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension contains an invalid OID value: 2.16.840.1.114569.1.1.1. Available OIDs: 2.16.840.1.114569.1.1.3, 2.16.840.1.114569.1.1.4 |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 583j', but common name is 'SHAKEN Broadband Dynamics LLC 583j' |
| [e_atis_tn_auth_list_spc_format](../../ISSUES/e_atis_tn_auth_list_spc_format/README.md) | error | ATIS1000080 | the SPC value '583j' contains characters other than uppercase letters and numbers |


Generated: 04 Oct 24 16:29 UTC