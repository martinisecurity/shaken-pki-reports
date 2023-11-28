# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 7661

Tested At: 28 Nov 23 16:03 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 91 day(s)\
Subject: CN=SHAKEN 7661, O=Cox Communications Inc, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN CA-2, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US\
Link: https://cr.ccid.neustar.biz/ccid/authn/v2/certs/11126.10191

[View certificate details](https://understandingwebpki.com/?cert=MIIDCzCCArCgAwIBAgIUO6LTfKnrSwCQLl9NJBimUR6V%2F%2F8wCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0yMB4XDTIzMDIyNzE1MzQzOVoXDTI0MDIyNzE1MzQzOVowRDELMAkGA1UEBhMCVVMxHzAdBgNVBAoMFkNveCBDb21tdW5pY2F0aW9ucyBJbmMxFDASBgNVBAMMC1NIQUtFTiA3NjYxMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE943%2FDERhSESz3mfkkGNY9slDWu8jmX1dX8l5dEe5G4DtVKUAwCHxnhqxQq5qZnNhbUEfMl1Miypl%2BvMuBC2KaqOCATwwggE4MBYGCCsGAQUFBwEaBAowCKAGFgQ3NjYxMAwGA1UdEwEB%2FwQCMAAwHwYDVR0jBBgwFoAUgk4V%2F%2F6famdR5MiXx210w%2FxlRXgwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAdBgNVHQ4EFgQUqu0JHCHQn5TH3QYU4773pEWfgYQwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0kAMEYCIQCL2DpC7FHl8cGu6ENa51%2B881bU9oMlp4jM7EWZ9jJO1AIhAINFo0Oz7hCdFbwLx0b3xiWO1o%2FKw1Srd4wz%2FqTIDOR1)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 28 Nov 23 16:15 UTC