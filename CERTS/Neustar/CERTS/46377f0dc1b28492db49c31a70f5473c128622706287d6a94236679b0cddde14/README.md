# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 506J

Tested At: 04 Oct 24 15:51 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: -317 day(s)\
Subject: CN=SHAKEN 506J, O=Twilio International, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN CA-2, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US\
Link: https://cr.ccid.neustar.biz/ccid/authn/v2/certs/11095.10213.pem

[View certificate details](https://x509.io/?cert=MIIDCTCCAq6gAwIBAgIUIFl%2Bidm6KHDJI8ZAMgYHEcI2nQ8wCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0yMB4XDTIyMTEyMTE2MzIyNloXDTIzMTEyMTE2MzIyNlowQjELMAkGA1UEBhMCVVMxHTAbBgNVBAoMFFR3aWxpbyBJbnRlcm5hdGlvbmFsMRQwEgYDVQQDDAtTSEFLRU4gNTA2SjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABFaeTH5FMEYyncUxr%2FlkM9RYOkEoq5evh93oKbdwJG5%2F68YCnUB0XgkLIjFv16C7RyZ2GuC4g0LyreJhPE1SC1yjggE8MIIBODAWBggrBgEFBQcBGgQKMAigBhYENTA2SjAMBgNVHRMBAf8EAjAAMB8GA1UdIwQYMBaAFIJOFf%2F%2Bn2pnUeTIl8dtdMP8ZUV4MBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwHQYDVR0OBBYEFG%2BrydAGZQt6TsIHk%2F9vNdFcj57jMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNJADBGAiEAkkUaOM51QFzBFH9V4zoVLxYNLFZ%2FfHgsF6BZ7qah17gCIQCNFdv2T%2B9mbcjkD9Y0mxZnHOWvSrybHE8MQB4CKEJhDg%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 04 Oct 24 16:29 UTC