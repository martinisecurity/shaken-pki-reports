# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Midwest Telecom of America 919A

Tested At: 22 Aug 24 15:54 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: -90 day(s)\
Subject: CN=SHAKEN Midwest Telecom of America 919A, OU=Midwest, O=Midwest Telecom of America, ST=Indiana, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/Midwest_Telecom_919A

[View certificate details](https://x509.io/?cert=MIIC7DCCApKgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkcgswCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDUyNTE1MDA1N1oXDTI0MDUyNDE1MDA1N1owgYcxCzAJBgNVBAYTAlVTMRAwDgYDVQQIDAdJbmRpYW5hMSMwIQYDVQQKDBpNaWR3ZXN0IFRlbGVjb20gb2YgQW1lcmljYTEQMA4GA1UECwwHTWlkd2VzdDEvMC0GA1UEAwwmU0hBS0VOIE1pZHdlc3QgVGVsZWNvbSBvZiBBbWVyaWNhIDkxOUEwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASc%2B7rcxOTmQX%2B5Tvc%2B1en%2FE%2BjTUgUlZyNA1rZNk6o1ccOkbSthOZiQDEAAs0gsvgWaUKvpJdHx0o2HE8AbIBv9o4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQ5MTlBMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzAdBgNVHQ4EFgQU%2FVdAJhF8isLqTiXILq6yRXPpjWowHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0gAMEUCIQDLY3twq3%2B4BDRXd0ggEtBhKDLm%2FmpQZBemludc%2B098%2FwIgB7%2BJ%2BqxIrRmHpF4%2FYUgtMFDyyotHY46OiVjNUn%2BlEG4%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 919A', but common name is 'SHAKEN Midwest Telecom of America 919A' |


Generated: 22 Aug 24 16:06 UTC