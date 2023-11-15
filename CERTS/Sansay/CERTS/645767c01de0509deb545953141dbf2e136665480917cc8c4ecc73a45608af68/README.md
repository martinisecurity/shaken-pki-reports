# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Go Voip Dialing LLC 704K

Tested At: 15 Nov 23 16:10 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 294 day(s)\
Subject: CN=SHAKEN Go Voip Dialing LLC 704K, OU=main, O=Go Voip Dialing LLC, ST=California, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/govoipdialing_704365

[View certificate details](https://understandingwebpki.com/?cert=MIIC3jCCAoOgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkf4swCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDkwNTE1NTQyN1oXDTI0MDkwNDE1NTQyN1oweTELMAkGA1UEBhMCVVMxEzARBgNVBAgMCkNhbGlmb3JuaWExHDAaBgNVBAoME0dvIFZvaXAgRGlhbGluZyBMTEMxDTALBgNVBAsMBG1haW4xKDAmBgNVBAMMH1NIQUtFTiBHbyBWb2lwIERpYWxpbmcgTExDIDcwNEswWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAARDAUVMNTi9VBiVBUAYnodnHGoxsfvDNiEDlvVY4cpzxxg1PhMCWco6bKsbQKpSccweFcVwCdFmq%2FrjARQ%2BECrYo4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQ3MDRLMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzAdBgNVHQ4EFgQUb6vXFEsx0T9yd0TRRq8LWy2xMiwwHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0kAMEYCIQCQIHYuyMsYXNWFIti4c67xnRTgDR5%2FcQiWw20dIrhslAIhAJkb5nbSguVSuT%2FVHbOxEMwC9mKGUB0YnzcwFSbk4%2B6z)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 704K' |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 15 Nov 23 16:51 UTC