# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN MagicJack 324E

Tested At: 28 Nov 23 19:56 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 2 day(s)\
Subject: CN=SHAKEN MagicJack 324E, OU=NOC, O=MagicJack, ST=Florida, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr-east1.sansay.com/MagicJack_324E.crt

[View certificate details](https://understandingwebpki.com/?cert=MIICxDCCAmugAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkVrkwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTEyOTIyMDQyN1oXDTIzMTEyOTIyMDQyN1owYTELMAkGA1UEBhMCVVMxEDAOBgNVBAgMB0Zsb3JpZGExEjAQBgNVBAoMCU1hZ2ljSmFjazEMMAoGA1UECwwDTk9DMR4wHAYDVQQDDBVTSEFLRU4gTWFnaWNKYWNrIDMyNEUwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASPxYeO7gUtKwcxTmxvuIA9qpRvfNeHrlNku7%2F1%2FAey0nHwIv8jDX1PlNEk8rH6jq5ZVxcVFxOai1mUrWxUJh%2Fvo4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQzMjRFMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQUpsl%2B2DMRMF2F6vY%2BWaw9ZgCWdRMwHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0cAMEQCIGYi4%2BQgdhxwP9u2SGSw1gCyBdzh2TEachLllWJZJA5YAiBaOpFm09TGVB7kaJe2bhiPj58tLsWQxdQQhSYQR%2BWjjw%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 324E', but common name is 'SHAKEN MagicJack 324E' |
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension contains an invalid OID value: 2.16.840.1.114569.1.1.1. Available OIDs: 2.16.840.1.114569.1.1.3, 2.16.840.1.114569.1.1.4 |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 28 Nov 23 20:21 UTC