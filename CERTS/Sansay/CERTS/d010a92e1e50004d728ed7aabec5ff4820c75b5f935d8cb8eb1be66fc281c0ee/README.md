# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN BareTelecom 864J

Tested At: 04 Oct 24 16:11 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -47 day(s)\
Subject: CN=SHAKEN BareTelecom 864J, O=BareTelecom, ST=California, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/864J/429C7C70711E3820F0B8E1DEAE6FF3262264A3EB.pem

[View certificate details](https://x509.io/?cert=MIICvjCCAmSgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJko%2BswCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTI0MDcxOTE0Mjk0MVoXDTI0MDgxODE0Mjk0MVowWjELMAkGA1UEBhMCVVMxEzARBgNVBAgMCkNhbGlmb3JuaWExFDASBgNVBAoMC0JhcmVUZWxlY29tMSAwHgYDVQQDDBdTSEFLRU4gQmFyZVRlbGVjb20gODY0SjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABDDS7l0vaaxRo5PuD%2BPqVhVve6XeHWGMrTrmV563iRvazHazHmoMABAUKMD0FgTQtsh5TKpUQdvJvnJwByVoCaGjgdswgdgwFgYIKwYBBQUHARoECjAIoAYWBDg2NEowFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEEMB0GA1UdDgQWBBRIdUFl6KvT9XyIaPknmcYnWIKCbzAfBgNVHSMEGDAWgBSs05P1Q0PMCr5FWBcTfZJ83MMBRjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSAAwRQIhAO7AgFVVwtLw1hHnEJR9lQ23t1VooTcfXDR4q0X66K5RAiAfBOzkWSQCAsR8gFZ37h7wVm2%2Bu4V2dI%2FrnFWobz7Dag%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 864J', but common name is 'SHAKEN BareTelecom 864J' |
| [e_atis_ext_crl_distribution_struct](../../ISSUES/e_atis_ext_crl_distribution_struct/README.md) | error | ATIS1000080 | CRL Distribution Point shall contain a CRLIssuer field |


Generated: 04 Oct 24 16:29 UTC