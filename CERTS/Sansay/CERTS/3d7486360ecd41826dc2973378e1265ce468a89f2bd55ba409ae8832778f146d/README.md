# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Townes Telecommunications 0335

Tested At: 28 Nov 23 10:30 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -142 day(s)\
Subject: emailAddress=bshafter@townes.net, CN=SHAKEN Townes Telecommunications 0335, OU=Operations, O=Townes Telecommunications, ST=Arkansas, C=US, emailAddress=bshafter@townes.net\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/0335/order/196_0335_135

[View certificate details](https://understandingwebpki.com/?cert=MIIDETCCArigAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkdFEwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDYwOTA4NTIxOVoXDTIzMDcwOTA4NTIxOVowga0xCzAJBgNVBAYTAlVTMREwDwYDVQQIDAhBcmthbnNhczEiMCAGA1UECgwZVG93bmVzIFRlbGVjb21tdW5pY2F0aW9uczETMBEGA1UECwwKT3BlcmF0aW9uczEuMCwGA1UEAwwlU0hBS0VOIFRvd25lcyBUZWxlY29tbXVuaWNhdGlvbnMgMDMzNTEiMCAGCSqGSIb3DQEJARYTYnNoYWZ0ZXJAdG93bmVzLm5ldDBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABGF1ZxIGC4VjyaPyDsPJTswq3ZbVUN75r3tb%2F0xBGM7RNopsh74mnydsVvBmIpJyUYj3aP%2BcjyDjYWdJZcsN4ZKjgdswgdgwFgYIKwYBBQUHARoECjAIoAYWBDAzMzUwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMB0GA1UdDgQWBBSOmwmrPWR5yxize6vXhgQ5Fu4InzAfBgNVHSMEGDAWgBSs05P1Q0PMCr5FWBcTfZJ83MMBRjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDRwAwRAIgVv1vWsN4h5Wpt2ELQTEg5t5wx6BEkdqK2VgvhaLFI%2FYCIA1aW%2FG9nCVmF5mqbKOSkVeLZFCr1DgcFc67xLwNpylG)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 0335', but common name is 'SHAKEN Townes Telecommunications 0335' |
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |


Generated: 28 Nov 23 10:53 UTC