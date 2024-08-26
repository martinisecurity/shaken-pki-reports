# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 7575

Tested At: 26 Aug 24 17:47 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: -305 day(s)\
Subject: CN=SHAKEN 7575, O=Lumen, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN CA-2, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US\
Link: https://cr.ccid.neustar.biz/ccid/authn/v2/certs/11563.10209.pem

[View certificate details](https://x509.io/?cert=MIIC%2BTCCAp%2BgAwIBAgIUdGsJGGqju%2BrxFiWCyrVusFix%2Fv0wCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0yMB4XDTIyMTAyNjE2NDAyNVoXDTIzMTAyNjE2NDAyNVowMzELMAkGA1UEBhMCVVMxDjAMBgNVBAoMBUx1bWVuMRQwEgYDVQQDDAtTSEFLRU4gNzU3NTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABM7HztY3E3h3cCQlhrjt3x9nXiJC95o3grxLdXivEhZXHto5xFxucQ9PJErkebeF4nBO65%2FHGFMD3BkRcV7aqG6jggE8MIIBODAWBggrBgEFBQcBGgQKMAigBhYENzU3NTAMBgNVHRMBAf8EAjAAMB8GA1UdIwQYMBaAFIJOFf%2F%2Bn2pnUeTIl8dtdMP8ZUV4MBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwHQYDVR0OBBYEFGylkAvkUJcNv1QFYHNyq%2Bcg8nDaMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNIADBFAiEA4aCtiRSKi2IqJ2DKVmaGqSshVt435Wbvx64xS8CGyo0CIHEFmkTdfcvovuQ%2FFrLff9j1NX2%2B7ZW3%2BkP%2FgYMPTFpc)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 26 Aug 24 18:03 UTC