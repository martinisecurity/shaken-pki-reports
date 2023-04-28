# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Starlinq PBX Inc. 267K

Tested At: 28 Apr 23 02:09 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 194 day(s)\
Subject: CN=SHAKEN Starlinq PBX Inc. 267K, OU=Engineering, O=Starlinq PBX Inc., ST=California, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/Starlinq_PBX_Inc365._267K

[View certificate details](https://understandingwebpki.com/?cert=MIIC4TCCAoagAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkVNEwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTEwNzIxNTM0NVoXDTIzMTEwNzIxNTM0NVowfDELMAkGA1UEBhMCVVMxEzARBgNVBAgMCkNhbGlmb3JuaWExGjAYBgNVBAoMEVN0YXJsaW5xIFBCWCBJbmMuMRQwEgYDVQQLDAtFbmdpbmVlcmluZzEmMCQGA1UEAwwdU0hBS0VOIFN0YXJsaW5xIFBCWCBJbmMuIDI2N0swWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQFWFAAZBAAtvNzptGUCDf34dndd7tp2hfyz%2BGjScWjudp97Ueter1LmCIBzqBv%2BMIj%2BKPHIQ9ryAzhDHnUsAKio4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQyNjdLMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQUBZq%2BGRvxjIXucD18dxTfTd%2BA2HEwHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0kAMEYCIQCk6ijp1EERWLYCX9DpdS2TmiMzu%2FE%2BopHDZfR9rDcoMAIhAOx%2Fc9v1T8kR2P2sHwgpPPXUcZpNPnmbCWkNOOkW5td5)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 267K' |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |


Generated: 28 Apr 23 02:17 UTC