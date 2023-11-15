# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Doylestown Communications, Inc 0609

Tested At: 15 Nov 23 16:34 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -119 day(s)\
Subject: emailAddress=jclarke@heritagetelephone.com, CN=SHAKEN Doylestown Communications\\, Inc 0609, OU=Operations, O=Doylestown Communications\\, Inc, ST=Ohio, C=US, emailAddress=jclarke@heritagetelephone.com\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/0609/order/426_849C_78

[View certificate details](https://understandingwebpki.com/?cert=MIIDITCCAsigAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkdcgwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDYxOTAwNTkxMVoXDTIzMDcxOTAwNTkxMVowgb0xCzAJBgNVBAYTAlVTMQ0wCwYDVQQIDARPaGlvMScwJQYDVQQKDB5Eb3lsZXN0b3duIENvbW11bmljYXRpb25zLCBJbmMxEzARBgNVBAsMCk9wZXJhdGlvbnMxMzAxBgNVBAMMKlNIQUtFTiBEb3lsZXN0b3duIENvbW11bmljYXRpb25zLCBJbmMgMDYwOTEsMCoGCSqGSIb3DQEJARYdamNsYXJrZUBoZXJpdGFnZXRlbGVwaG9uZS5jb20wWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQfMY36tt3eLI8%2BixGez1wm%2FlMUa9iBIAG9ydW4Mb2qaTpIPhp7RwLJmS5B89aaMqocFuLsmK7kl0IsAR27D2EMo4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQwNjA5MBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzAdBgNVHQ4EFgQUmRWb%2BLS5QtrYYRnXt6y4XlYE4cMwHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0cAMEQCIDOHzpZ7xnC5%2BFXRFtuOa1kJ0m%2FuhJ0uRacLBmjrV0G4AiBXRZC8prGbrZfaKIbsFSlEg7RLHaEajsZLRcf1zrQFLA%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_cp_1_3_subject_email](../../ISSUES/w_cp_1_3_subject_email/README.md) | warn | US_SHAKEN_CP | Email addresses are not allowed as the CP does not specify how to validate them |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 0609' |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |


Generated: 15 Nov 23 17:17 UTC