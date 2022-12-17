# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 869J

Tested At: 17 Dec 22 17:00 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 151 day(s)\
Subject: CN=SHAKEN 869J, O=Sipnex Telecom LLC, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN CA-1, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US\
Link: https://ssc.getsipnav.com/certs/fef3962eaffeb276be4bf92fa16a666186c83733

[View certificate details](https://understandingwebpki.com/?cert=MIIDAzCCAqmgAwIBAgIUbCPyl9mbNXRPq8DzE65mjU7m2%2BAwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0xMB4XDTIyMDUxNzEyMTAzMloXDTIzMDUxNzEyMTAzMlowQDELMAkGA1UEBhMCVVMxGzAZBgNVBAoMElNpcG5leCBUZWxlY29tIExMQzEUMBIGA1UEAwwLU0hBS0VOIDg2OUowWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQNJVmzaHy2P8nOlJKyQe%2F%2BXx0U9wlNRbGwrxNfU9H%2B0XKYIFhk5p27LPeIAxVm8U8B0RqNpdOLWHv8VRPAEqQ4o4IBOTCCATUwFgYIKwYBBQUHARoECjAIoAYWBDg2OUowDAYDVR0TAQH%2FBAIwADAfBgNVHSMEGDAWgBSv0cjC7nJMg%2Fw%2F7RmnbR2QsgfwOjBbBggrBgEFBQcBAQRPME0wSwYIKwYBBQUHMAKGP2h0dHA6Ly9jYWNlcnRzLXVzLmNjaWQubmV1c3Rhci9OZXVzdGFyQ2VydGlmaWVkQ2FsbGVySWRDQTEuY3J0IDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMB0GA1UdDgQWBBQrVu8iJlBSD95SUHGbs%2BEYaugyOzAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSAAwRQIhAJFc1ZdwZlPRZ5xjtR43RwDWElsM4O%2Fm0AGnkDGdbutlAiAUxDpk3zcfH0QZCdty5vHK1gKYLQ4c2WLwNHznAxi4ig%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |


Generated: 17 Dec 22 17:07 UTC