# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Phone.com, Inc. 633J

Tested At: 30 Nov 22 15:53 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 16 day(s)\
Subject: CN=SHAKEN Phone.com\\, Inc. 633J, OU=voipteam, O=Phone.com\\, Inc., ST=New Jersey, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/Phonedotcom_633J

[View certificate details](https://understandingwebpki.com/?cert=MIIC2jCCAn%2BgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkVYQwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTExNTIxMDQzNFoXDTIyMTIxNTIxMDQzNFowdTELMAkGA1UEBhMCVVMxEzARBgNVBAgMCk5ldyBKZXJzZXkxGDAWBgNVBAoMD1Bob25lLmNvbSwgSW5jLjERMA8GA1UECwwIdm9pcHRlYW0xJDAiBgNVBAMMG1NIQUtFTiBQaG9uZS5jb20sIEluYy4gNjMzSjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABJTLc%2BXG1A5e1KyUnjizH7mo5f%2F%2B26dXea1HU0AeWM5RZHqKQFPEVumydDsK204FUayisWfuxa8XQh2AlfbaMV2jgdswgdgwFgYIKwYBBQUHARoECjAIoAYWBDYzM0owFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMB0GA1UdDgQWBBTWuKTd71%2FET2TtuoPYJdJ4kF9f3TAfBgNVHSMEGDAWgBSs05P1Q0PMCr5FWBcTfZJ83MMBRjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSQAwRgIhAON9WOnrZFdVV3VNnDBr30UPAt2Igb%2FdIXo%2BZ%2Fzu6UoAAiEAjU6ppAMxYJ8%2BABRigkzKHZov8Ld%2Fxvn0FCUueZM9xRU%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 633J' |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |


Generated: 30 Nov 22 16:07 UTC