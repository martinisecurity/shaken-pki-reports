# STIR/SHAKEN CA Ecosystem Compliance

## Certificate prod SHAKEN 811J

Tested At: 21 Nov 23 01:22 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 81 day(s)\
Subject: CN=prod SHAKEN 811J, OU=Alianza, O=CaptionCall, L=Salt Lake City, ST=Utah, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN CA-2, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US\
Link: https://api.alianza.com/v2/stir-shaken/certs/2dd72c30-5e06-49a5-bbec-fff3cdf9444b/key.crt

[View certificate details](https://understandingwebpki.com/?cert=MIIDPjCCAuSgAwIBAgIUMk0uW8uETc5nb%2FxPyJJqm%2BIhzPQwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0yMB4XDTIzMDIwOTE5MzYxN1oXDTI0MDIwOTE5MzYxN1oweDELMAkGA1UEBhMCVVMxDTALBgNVBAgMBFV0YWgxFzAVBgNVBAcMDlNhbHQgTGFrZSBDaXR5MRQwEgYDVQQKDAtDYXB0aW9uQ2FsbDEQMA4GA1UECwwHQWxpYW56YTEZMBcGA1UEAwwQcHJvZCBTSEFLRU4gODExSjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABKiSBxMGSmBV0kVlim5ASMUMC8ww3VGrM1T4hzbLezbd2f8JsMxdLNj60LgvNFLez6Y2QnZIiroBuiXS1Gku9g6jggE8MIIBODAWBggrBgEFBQcBGgQKMAigBhYEODExSjAMBgNVHRMBAf8EAjAAMB8GA1UdIwQYMBaAFIJOFf%2F%2Bn2pnUeTIl8dtdMP8ZUV4MBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwHQYDVR0OBBYEFNK8qKfNXPjfgBaTD19iyhv1o2vMMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNIADBFAiEA3nALH0pjgN5zIxAjcnL5nap%2BxeEfOeSphqiEObZk6s4CIAowdnzUFFU52jDIWGJrfQ6vbZMo%2FQsHNWTQHdJOhGgz)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 21 Nov 23 01:55 UTC