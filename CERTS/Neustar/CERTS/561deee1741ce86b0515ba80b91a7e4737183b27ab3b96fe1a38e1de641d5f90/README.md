# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 1591

Tested At: 26 Aug 24 17:46 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: -193 day(s)\
Subject: CN=SHAKEN 1591, O=Southeast Nebraska Communications Inc, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN CA-2, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US\
Link: https://cr-partner.ccid.neustar.biz/ccid/authn/v2/certs/199.263

[View certificate details](https://x509.io/?cert=MIIDGDCCAr%2BgAwIBAgIUC2zFcvq8R5E0HtoSDZxp5YNwZokwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0yMB4XDTIzMDIxNTE3MDYwMFoXDTI0MDIxNTE3MDYwMFowUzELMAkGA1UEBhMCVVMxLjAsBgNVBAoMJVNvdXRoZWFzdCBOZWJyYXNrYSBDb21tdW5pY2F0aW9ucyBJbmMxFDASBgNVBAMMC1NIQUtFTiAxNTkxMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEtD9LUw2Iz1sa49cvK8POONhyT6fS8Sv3vHMExyx6hGGCMjE248o0Im0lTqQypyniKqjvXLkMzM8AGGAX2vC1aqOCATwwggE4MBYGCCsGAQUFBwEaBAowCKAGFgQxNTkxMAwGA1UdEwEB%2FwQCMAAwHwYDVR0jBBgwFoAUgk4V%2F%2F6famdR5MiXx210w%2FxlRXgwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAdBgNVHQ4EFgQUX%2BEspsQI3%2FNKmkztIYBh3or7U5gwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0cAMEQCIFgnEcMAGvWnGnFOFfz6It3lVUID%2BCBY3xy5096gMwG4AiBqYtOLTmv7w1NyGJEutor3pfXVOvpQXzDFJ2A8zXic9w%3D%3D)

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