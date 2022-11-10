# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN MagicJack 324E

Tested At: 10 Nov 22 23:21 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 8 day(s)\
Subject: CN=SHAKEN MagicJack 324E, OU=NOC, O=MagicJack, ST=Florida, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://sns.magicjack.com/324E.chain.crt.pem

[View certificate details](https://understandingwebpki.com/?cert=MIICxTCCAmugAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkUdUwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAxOTE2MTYzNFoXDTIyMTExODE2MTYzNFowYTELMAkGA1UEBhMCVVMxEDAOBgNVBAgMB0Zsb3JpZGExEjAQBgNVBAoMCU1hZ2ljSmFjazEMMAoGA1UECwwDTk9DMR4wHAYDVQQDDBVTSEFLRU4gTWFnaWNKYWNrIDMyNEUwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASPxYeO7gUtKwcxTmxvuIA9qpRvfNeHrlNku7%2F1%2FAey0nHwIv8jDX1PlNEk8rH6jq5ZVxcVFxOai1mUrWxUJh%2Fvo4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQzMjRFMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQUpsl%2B2DMRMF2F6vY%2BWaw9ZgCWdRMwHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0gAMEUCIFlsySWSavRzxorFHvPTBxMyUp%2FowIZz2ziMcNbruE7aAiEA56Gf%2Bq7ROzGAqvqO3LCGWOIoyHvvBrnI2zQWgQXXAP0%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 324E' |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |


Generated: 10 Nov 22 23:30 UTC