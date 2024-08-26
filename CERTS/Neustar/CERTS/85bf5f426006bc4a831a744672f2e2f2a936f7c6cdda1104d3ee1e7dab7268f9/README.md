# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 745J

Tested At: 26 Aug 24 18:01 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: -305 day(s)\
Subject: CN=SHAKEN 745J, O=TextMe Incorporated, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN CA-2, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US\
Link: https://textme-stirshaken.s3.us-west-2.amazonaws.com/textme-bundle_10-26-2023.cer

[View certificate details](https://x509.io/?cert=MIIDBjCCAq2gAwIBAgIUEv0qYHzi%2ByYfFiXMMx0bVAHwZ5IwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0yMB4XDTIyMTAyNjE1NDkzMVoXDTIzMTAyNjE1NDkzMVowQTELMAkGA1UEBhMCVVMxHDAaBgNVBAoME1RleHRNZSBJbmNvcnBvcmF0ZWQxFDASBgNVBAMMC1NIQUtFTiA3NDVKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEOPQqOE33HfHc1H4DxFBQeh14SBcYpYfEbjTVG0%2BBQT15Ym%2BKDgGZ%2FQcnjinEm8fEGYptp7ED7pjteQsyKRU7D6OCATwwggE4MBYGCCsGAQUFBwEaBAowCKAGFgQ3NDVKMAwGA1UdEwEB%2FwQCMAAwHwYDVR0jBBgwFoAUgk4V%2F%2F6famdR5MiXx210w%2FxlRXgwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAdBgNVHQ4EFgQUdv03z%2FYw7NjKY9eaIxvNTwTvhMYwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0cAMEQCIBNNEY9RpPHOwqRoW21houpzcm%2FHIxSYmgjKoYFEFcv3AiBpWR422GRdUGe8o75GtVhV7KHeEynSkpwyBhmFYHTFMQ%3D%3D)

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