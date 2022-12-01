# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Systemverse, LLC. 194K

Tested At: 01 Dec 22 19:07 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 330 day(s)\
Subject: CN=SHAKEN Systemverse\\, LLC. 194K, OU=VOIP, O=Systemverse\\, LLC., ST=Texas, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/Systemverse_LLC._194K

[View certificate details](https://understandingwebpki.com/?cert=MIIC1TCCAnqgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkU0swCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAyNjE5NDMzN1oXDTIzMTAyNjE5NDMzN1owcDELMAkGA1UEBhMCVVMxDjAMBgNVBAgMBVRleGFzMRowGAYDVQQKDBFTeXN0ZW12ZXJzZSwgTExDLjENMAsGA1UECwwEVk9JUDEmMCQGA1UEAwwdU0hBS0VOIFN5c3RlbXZlcnNlLCBMTEMuIDE5NEswWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATS9VEH1XyfpnU0oev95kAlLIRP8pCfp%2BRipM820nLiXMA6YBH5WLMVLagqtoNwwz1ewfSkQmWE3%2B2nrxMFjPnlo4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQxOTRLMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQUFR38DtBYPplHpfNIuHgGjskMUo0wHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0kAMEYCIQCpaknCnX0Oeqiq%2FVhKcvoA8Ys1DPBIBpvzucU%2BjC5zgwIhAJQsbBrVyPCGsVFH3%2BvuXH9exdB4UpBSoRZZ9KVlIgm7)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 194K' |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |


Generated: 01 Dec 22 19:10 UTC