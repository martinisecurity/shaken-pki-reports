# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Phone.com, Inc. 633J

Tested At: 22 Aug 24 15:31 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -69 day(s)\
Subject: CN=SHAKEN Phone.com\\, Inc. 633J, OU=voipteam, O=Phone.com\\, Inc., ST=New Jersey, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/633J/429C7C70711E3820F0B8E1DEAE6FF32622649C5C.pem

[View certificate details](https://x509.io/?cert=MIIC2TCCAn%2BgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJknFwwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTI0MDUxNTE0MTQ0OFoXDTI0MDYxNDE0MTQ0OFowdTELMAkGA1UEBhMCVVMxEzARBgNVBAgMCk5ldyBKZXJzZXkxGDAWBgNVBAoMD1Bob25lLmNvbSwgSW5jLjERMA8GA1UECwwIdm9pcHRlYW0xJDAiBgNVBAMMG1NIQUtFTiBQaG9uZS5jb20sIEluYy4gNjMzSjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABJTLc%2BXG1A5e1KyUnjizH7mo5f%2F%2B26dXea1HU0AeWM5RZHqKQFPEVumydDsK204FUayisWfuxa8XQh2AlfbaMV2jgdswgdgwFgYIKwYBBQUHARoECjAIoAYWBDYzM0owFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMB0GA1UdDgQWBBTWuKTd71%2FET2TtuoPYJdJ4kF9f3TAfBgNVHSMEGDAWgBSs05P1Q0PMCr5FWBcTfZJ83MMBRjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSAAwRQIgM5AipaTtkdTlao7n6YYWF57whweDr43LMvQp%2Fxy8XvACIQCi9jmXQ9IpeBtUJ6zabtvjer0uCtCQFb6FUOt72ULFRQ%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension contains an invalid OID value: 2.16.840.1.114569.1.1.1. Available OIDs: 2.16.840.1.114569.1.1.3, 2.16.840.1.114569.1.1.4 |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 633J', but common name is 'SHAKEN Phone.com, Inc. 633J' |


Generated: 22 Aug 24 15:44 UTC