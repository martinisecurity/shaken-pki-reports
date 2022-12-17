# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN MagicJack 324E

Tested At: 17 Dec 22 12:13 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 348 day(s)\
Subject: CN=SHAKEN MagicJack 324E, OU=NOC, O=MagicJack, ST=Florida, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr-east1.sansay.com/MagicJack_324E.crt

[View certificate details](https://understandingwebpki.com/?cert=MIICxDCCAmugAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkVrkwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTEyOTIyMDQyN1oXDTIzMTEyOTIyMDQyN1owYTELMAkGA1UEBhMCVVMxEDAOBgNVBAgMB0Zsb3JpZGExEjAQBgNVBAoMCU1hZ2ljSmFjazEMMAoGA1UECwwDTk9DMR4wHAYDVQQDDBVTSEFLRU4gTWFnaWNKYWNrIDMyNEUwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASPxYeO7gUtKwcxTmxvuIA9qpRvfNeHrlNku7%2F1%2FAey0nHwIv8jDX1PlNEk8rH6jq5ZVxcVFxOai1mUrWxUJh%2Fvo4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQzMjRFMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQUpsl%2B2DMRMF2F6vY%2BWaw9ZgCWdRMwHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0cAMEQCIGYi4%2BQgdhxwP9u2SGSw1gCyBdzh2TEachLllWJZJA5YAiBaOpFm09TGVB7kaJe2bhiPj58tLsWQxdQQhSYQR%2BWjjw%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 324E' |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |


Generated: 17 Dec 22 12:22 UTC