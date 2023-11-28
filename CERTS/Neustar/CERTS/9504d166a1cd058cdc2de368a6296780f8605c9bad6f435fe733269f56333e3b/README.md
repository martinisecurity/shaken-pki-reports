# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 418c

Tested At: 28 Nov 23 10:28 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 105 day(s)\
Subject: CN=SHAKEN 418c, O=Armstrong Telecommunications Inc, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN CA-2, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US\
Link: https://cr.ccid.neustar.biz/ccid/authn/v2/certs/11112.10192.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIDFDCCArqgAwIBAgIUQhqymUy4%2F1alQmP0nFcoS2O9MMwwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0yMB4XDTIzMDMxMzAxMjUzM1oXDTI0MDMxMjAxMjUzM1owTjELMAkGA1UEBhMCVVMxKTAnBgNVBAoMIEFybXN0cm9uZyBUZWxlY29tbXVuaWNhdGlvbnMgSW5jMRQwEgYDVQQDDAtTSEFLRU4gNDE4YzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABJf8R067IdBVv783RMsqy4HknmbKpzPbWJfb7yV6vhk5vzIU2t1fxPqDHHuRxGPnfzOgzUp08PqZyeUhyFtX8YCjggE8MIIBODAWBggrBgEFBQcBGgQKMAigBhYENDE4YzAMBgNVHRMBAf8EAjAAMB8GA1UdIwQYMBaAFIJOFf%2F%2Bn2pnUeTIl8dtdMP8ZUV4MBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwHQYDVR0OBBYEFBqu7GZcNPk6i5wqjW6O8iaqhjzeMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNIADBFAiEA44AKpUSH%2F%2Foih67tEcN1q49wbZ6VRPSGlYTzpNIr%2FDECIFo3aSTGVvTS8Ug7vVd6dt9HGjaraQo2VgzAr2ARhGUA)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_tn_auth_list_spc_format](../../ISSUES/e_atis_tn_auth_list_spc_format/README.md) | error | ATIS1000080 | the SPC value '418c' contains characters other than uppercase letters and numbers |
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |


Generated: 28 Nov 23 10:53 UTC