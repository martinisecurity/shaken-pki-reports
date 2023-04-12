# STIR/SHAKEN CA Ecosystem Compliance

## Certificate prod SHAKEN 811J

Tested At: 12 Apr 23 01:01 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 304 day(s)\
Subject: CN=prod SHAKEN 811J, OU=Alianza, O=CaptionCall, L=Salt Lake City, ST=Utah, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN CA-2, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US\
Link: https://api.alianza.com/v2/stir-shaken/certs/2dd72c30-5e06-49a5-bbec-fff3cdf9444b/key.crt

[View certificate details](https://understandingwebpki.com/?cert=MIIDPjCCAuSgAwIBAgIUMk0uW8uETc5nb%2FxPyJJqm%2BIhzPQwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0yMB4XDTIzMDIwOTE5MzYxN1oXDTI0MDIwOTE5MzYxN1oweDELMAkGA1UEBhMCVVMxDTALBgNVBAgMBFV0YWgxFzAVBgNVBAcMDlNhbHQgTGFrZSBDaXR5MRQwEgYDVQQKDAtDYXB0aW9uQ2FsbDEQMA4GA1UECwwHQWxpYW56YTEZMBcGA1UEAwwQcHJvZCBTSEFLRU4gODExSjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABKiSBxMGSmBV0kVlim5ASMUMC8ww3VGrM1T4hzbLezbd2f8JsMxdLNj60LgvNFLez6Y2QnZIiroBuiXS1Gku9g6jggE8MIIBODAWBggrBgEFBQcBGgQKMAigBhYEODExSjAMBgNVHRMBAf8EAjAAMB8GA1UdIwQYMBaAFIJOFf%2F%2Bn2pnUeTIl8dtdMP8ZUV4MBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwHQYDVR0OBBYEFNK8qKfNXPjfgBaTD19iyhv1o2vMMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNIADBFAiEA3nALH0pjgN5zIxAjcnL5nap%2BxeEfOeSphqiEObZk6s4CIAowdnzUFFU52jDIWGJrfQ6vbZMo%2FQsHNWTQHdJOhGgz)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 12 Apr 23 01:46 UTC