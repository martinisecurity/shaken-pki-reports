# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 418c

Tested At: 18 Aug 25 20:19 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 144 day(s)\
Subject: CN=SHAKEN 418c, O=Armstrong Telecommunications Inc, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN CA-2, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US\
Link: https://cr.ccid.neustar.biz/ccid/authn/v2/certs/11112.10192.pem

[View certificate details](https://x509.io/?cert=MIIDEzCCArqgAwIBAgIUa2O4oBHtdWLnuzD%2FapYamK%2FBazQwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0yMB4XDTI1MDEwODIxMzkyNFoXDTI2MDEwODIxMzkyNFowTjELMAkGA1UEBhMCVVMxKTAnBgNVBAoMIEFybXN0cm9uZyBUZWxlY29tbXVuaWNhdGlvbnMgSW5jMRQwEgYDVQQDDAtTSEFLRU4gNDE4YzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABLz6AaMF1d0JAkL0tKpRBzO%2BfstGLzWJvxxw7sNFdPSy04jxrDrUJoAqgYUX6dWv%2FeTy9QSsXiZYD3T7iCGWMCmjggE8MIIBODAWBggrBgEFBQcBGgQKMAigBhYENDE4YzAMBgNVHRMBAf8EAjAAMB8GA1UdIwQYMBaAFIJOFf%2F%2Bn2pnUeTIl8dtdMP8ZUV4MBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwHQYDVR0OBBYEFIRGxkCsuUM9wJ37loV4xT4vPs9CMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNHADBEAiBvoGagVXgXlDTfTBXU2fLPd8EEd9F1vLCNLbB51thtPAIgYu75fFG%2B4cwjop61StBTqngm6LnUsn%2FMVdHaZ6H9Rq4%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_tn_auth_list_spc_format](../../ISSUES/e_atis_tn_auth_list_spc_format/README.md) | error | ATIS1000080 | the SPC value '418c' contains characters other than uppercase letters and numbers |
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |


Generated: 18 Aug 25 21:13 UTC