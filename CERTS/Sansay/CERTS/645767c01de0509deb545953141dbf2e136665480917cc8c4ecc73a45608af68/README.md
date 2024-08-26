# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Go Voip Dialing LLC 704K

Tested At: 26 Aug 24 17:48 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 9 day(s)\
Subject: CN=SHAKEN Go Voip Dialing LLC 704K, OU=main, O=Go Voip Dialing LLC, ST=California, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/govoipdialing_704365

[View certificate details](https://x509.io/?cert=MIIC3jCCAoOgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkf4swCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDkwNTE1NTQyN1oXDTI0MDkwNDE1NTQyN1oweTELMAkGA1UEBhMCVVMxEzARBgNVBAgMCkNhbGlmb3JuaWExHDAaBgNVBAoME0dvIFZvaXAgRGlhbGluZyBMTEMxDTALBgNVBAsMBG1haW4xKDAmBgNVBAMMH1NIQUtFTiBHbyBWb2lwIERpYWxpbmcgTExDIDcwNEswWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAARDAUVMNTi9VBiVBUAYnodnHGoxsfvDNiEDlvVY4cpzxxg1PhMCWco6bKsbQKpSccweFcVwCdFmq%2FrjARQ%2BECrYo4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQ3MDRLMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzAdBgNVHQ4EFgQUb6vXFEsx0T9yd0TRRq8LWy2xMiwwHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0kAMEYCIQCQIHYuyMsYXNWFIti4c67xnRTgDR5%2FcQiWw20dIrhslAIhAJkb5nbSguVSuT%2FVHbOxEMwC9mKGUB0YnzcwFSbk4%2B6z)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 704K', but common name is 'SHAKEN Go Voip Dialing LLC 704K' |
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |


Generated: 26 Aug 24 18:03 UTC