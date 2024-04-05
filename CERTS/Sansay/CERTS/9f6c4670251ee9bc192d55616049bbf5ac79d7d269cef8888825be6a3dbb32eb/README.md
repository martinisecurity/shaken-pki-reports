# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN i3 Broadband 5800

Tested At: 05 Apr 24 18:51 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -307 day(s)\
Subject: emailAddress=doug@i3broadband.com, CN=SHAKEN i3 Broadband 5800, OU=NOC, O=i3 Broadband, ST=Illinois, C=US, emailAddress=doug@i3broadband.com\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/5800/order/10_5800_165

[View certificate details](https://x509.io/?cert=MIIC8zCCApigAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkbdkwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDUwNDE0MjM0NFoXDTIzMDYwMzE0MjM0NFowgY0xCzAJBgNVBAYTAlVTMREwDwYDVQQIDAhJbGxpbm9pczEVMBMGA1UECgwMaTMgQnJvYWRiYW5kMQwwCgYDVQQLDANOT0MxITAfBgNVBAMMGFNIQUtFTiBpMyBCcm9hZGJhbmQgNTgwMDEjMCEGCSqGSIb3DQEJARYUZG91Z0BpM2Jyb2FkYmFuZC5jb20wWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQG9vxbou7w97hFuz8q6f4tNsy4FUDmwCuhd6EF%2Fq0q2UDZvDoOzgLq%2BO2GPlwwM0%2BFRSOKP3%2Fma%2BQokb%2BN6%2FVuo4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQ1ODAwMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzAdBgNVHQ4EFgQUPT4UG9fn9y%2BXTwtELPd3isHRMSEwHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0kAMEYCIQCHUccWcp%2F46olTuhIf4DaMY%2F1N5fNZTUyFl8X%2F6GKrCAIhAMNlLrKIef9be%2FmzbiO07VQTIr%2FBPSrIErWYrEwIrScU)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 5800', but common name is 'SHAKEN i3 Broadband 5800' |


Generated: 05 Apr 24 19:04 UTC