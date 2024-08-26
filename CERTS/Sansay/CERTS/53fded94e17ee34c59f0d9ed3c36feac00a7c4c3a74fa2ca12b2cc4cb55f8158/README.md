# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Doylestown Communications, Inc 0609

Tested At: 26 Aug 24 18:15 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -404 day(s)\
Subject: emailAddress=jclarke@heritagetelephone.com, CN=SHAKEN Doylestown Communications\\, Inc 0609, OU=Operations, O=Doylestown Communications\\, Inc, ST=Ohio, C=US, emailAddress=jclarke@heritagetelephone.com\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/0609/order/426_849C_78

[View certificate details](https://x509.io/?cert=MIIDITCCAsigAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkdcgwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDYxOTAwNTkxMVoXDTIzMDcxOTAwNTkxMVowgb0xCzAJBgNVBAYTAlVTMQ0wCwYDVQQIDARPaGlvMScwJQYDVQQKDB5Eb3lsZXN0b3duIENvbW11bmljYXRpb25zLCBJbmMxEzARBgNVBAsMCk9wZXJhdGlvbnMxMzAxBgNVBAMMKlNIQUtFTiBEb3lsZXN0b3duIENvbW11bmljYXRpb25zLCBJbmMgMDYwOTEsMCoGCSqGSIb3DQEJARYdamNsYXJrZUBoZXJpdGFnZXRlbGVwaG9uZS5jb20wWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQfMY36tt3eLI8%2BixGez1wm%2FlMUa9iBIAG9ydW4Mb2qaTpIPhp7RwLJmS5B89aaMqocFuLsmK7kl0IsAR27D2EMo4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQwNjA5MBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzAdBgNVHQ4EFgQUmRWb%2BLS5QtrYYRnXt6y4XlYE4cMwHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0cAMEQCIDOHzpZ7xnC5%2BFXRFtuOa1kJ0m%2FuhJ0uRacLBmjrV0G4AiBXRZC8prGbrZfaKIbsFSlEg7RLHaEajsZLRcf1zrQFLA%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 0609', but common name is 'SHAKEN Doylestown Communications, Inc 0609' |


Generated: 26 Aug 24 18:49 UTC