# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 418c

Tested At: 05 Apr 24 18:47 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 306 day(s)\
Subject: CN=SHAKEN 418c, O=Armstrong Telecommunications Inc, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN CA-2, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US\
Link: https://cr.ccid.neustar.biz/ccid/authn/v2/certs/11112.10192.pem

[View certificate details](https://x509.io/?cert=MIIDFTCCArqgAwIBAgIUBcXCJZWniPUcjv2IWr0kCWvRV%2BgwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0yMB4XDTI0MDIwNjE2MDczMFoXDTI1MDIwNTE2MDczMFowTjELMAkGA1UEBhMCVVMxKTAnBgNVBAoMIEFybXN0cm9uZyBUZWxlY29tbXVuaWNhdGlvbnMgSW5jMRQwEgYDVQQDDAtTSEFLRU4gNDE4YzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABIS7%2B17fig3kkGhX9zWhUfcw%2FwzVSi2XYKJu6kkzHiWDZyRD7WcP56MRuA1Q1Qn7IEnwA6FmaBwk9uRmDruyTPmjggE8MIIBODAWBggrBgEFBQcBGgQKMAigBhYENDE4YzAMBgNVHRMBAf8EAjAAMB8GA1UdIwQYMBaAFIJOFf%2F%2Bn2pnUeTIl8dtdMP8ZUV4MBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwHQYDVR0OBBYEFEda1IZVrAF4D2OtWM98enl3IPLgMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNJADBGAiEAyiifq5bDV9LvafBJzpNq1LCbcNVlUYTTh2sbYdPj2VUCIQDT5WunccCA8tVXoeqE4%2BNWfKz7Zu586m38uMHPkjf3iw%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_tn_auth_list_spc_format](../../ISSUES/e_atis_tn_auth_list_spc_format/README.md) | error | ATIS1000080 | the SPC value '418c' contains characters other than uppercase letters and numbers |
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |


Generated: 05 Apr 24 19:04 UTC