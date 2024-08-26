# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN SouthPoint Communications 205k

Tested At: 26 Aug 24 17:47 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -487 day(s)\
Subject: emailAddress=chris.duncan@soundcurve.com, CN=SHAKEN SouthPoint Communications 205k, OU=Operations, O=SouthPoint Communications, ST=California, C=US, emailAddress=chris.duncan@soundcurve.com\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/205k/order/247_205k_99

[View certificate details](https://x509.io/?cert=MIIDHDCCAsKgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkZ5MwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDMyODE3MDg1OVoXDTIzMDQyNzE3MDg1OVowgbcxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMSIwIAYDVQQKDBlTb3V0aFBvaW50IENvbW11bmljYXRpb25zMRMwEQYDVQQLDApPcGVyYXRpb25zMS4wLAYDVQQDDCVTSEFLRU4gU291dGhQb2ludCBDb21tdW5pY2F0aW9ucyAyMDVrMSowKAYJKoZIhvcNAQkBFhtjaHJpcy5kdW5jYW5Ac291bmRjdXJ2ZS5jb20wWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASfJ3OhB3xnPprzTOKiNwTougqlsCZMeV6M0kcwk8F%2BAkWyorpqFfoEGE65rnxAU2fZJu74W5lCPZtiPEyGdxs1o4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQyMDVrMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzAdBgNVHQ4EFgQU9%2BIQO2rwmus8MaQutkdnPhX%2F4aswHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0gAMEUCIB%2BbxSvWBOrq%2F%2BFCrT%2FfY127EX%2BugO8mYbCmL%2BhO5jHtAiEA%2BQYeAh829kbbBkT7%2FO%2FzBZMXUBiyGKKLSh9eOlELbB0%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 205k', but common name is 'SHAKEN SouthPoint Communications 205k' |
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |
| [e_atis_tn_auth_list_spc_format](../../ISSUES/e_atis_tn_auth_list_spc_format/README.md) | error | ATIS1000080 | the SPC value '205k' contains characters other than uppercase letters and numbers |


Generated: 26 Aug 24 18:03 UTC