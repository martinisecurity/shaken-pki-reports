# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 063E

Tested At: 22 Nov 23 03:23 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 15 day(s)\
Subject: CN=SHAKEN 063E, O=Peerless Network, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN CA-2, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US\
Link: https://cr.ccid.neustar.biz/ccid/authn/v2/certs/11032.10214.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIDBDCCAqqgAwIBAgIUYTCTlxQtIe18LLsPvlgefvARxdswCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0yMB4XDTIyMTIwNjE5NDM0M1oXDTIzMTIwNjE5NDM0M1owPjELMAkGA1UEBhMCVVMxGTAXBgNVBAoMEFBlZXJsZXNzIE5ldHdvcmsxFDASBgNVBAMMC1NIQUtFTiAwNjNFMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEbjJCYP8dmbmYa06ZXwJZn9faG%2BF%2BIr7OfS0dDCn0TErI%2BGpFqsuDlBguI0EWxo%2FMaijcmn8ePguOrQBPSdFybqOCATwwggE4MBYGCCsGAQUFBwEaBAowCKAGFgQwNjNFMAwGA1UdEwEB%2FwQCMAAwHwYDVR0jBBgwFoAUgk4V%2F%2F6famdR5MiXx210w%2FxlRXgwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAdBgNVHQ4EFgQUx02JRjhN%2BZkafxzquOZHuSXI4L0wDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0gAMEUCIQCetzGpSQBLwHP%2BKChWuu7YJFa6QKJQMaZw5NO3GkJlIgIgSW5romYNlkhZhBs9U11Emk6jS%2BiPy20CsOJ3a2u10F4%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 22 Nov 23 03:38 UTC