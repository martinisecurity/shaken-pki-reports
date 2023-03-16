# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 558a

Tested At: 16 Mar 23 19:04 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 99 day(s)\
Subject: CN=SHAKEN 558a, O=D&P Communications, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN CA-1, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US\
Link: https://cr.ccid.neustar.biz/ccid/authn/v2/certs/11569.10207.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIDAjCCAqmgAwIBAgIUIrEmb6Z5XRpGHjX5XyW%2F7IbbBKEwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0xMB4XDTIyMDYyMzE1MDQyMloXDTIzMDYyMzE1MDQyMlowQDELMAkGA1UEBhMCVVMxGzAZBgNVBAoMEkQmUCBDb21tdW5pY2F0aW9uczEUMBIGA1UEAwwLU0hBS0VOIDU1OGEwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASqlNsQF02%2FJbcw3v3Zy8jyhRBDEeGkiFpxRJeD2moIbE%2F%2BkhsANxsYG3RAa0PxFq05r10HBnnMVnn%2BpyG0K6aYo4IBOTCCATUwFgYIKwYBBQUHARoECjAIoAYWBDU1OGEwDAYDVR0TAQH%2FBAIwADAfBgNVHSMEGDAWgBSv0cjC7nJMg%2Fw%2F7RmnbR2QsgfwOjBbBggrBgEFBQcBAQRPME0wSwYIKwYBBQUHMAKGP2h0dHA6Ly9jYWNlcnRzLXVzLmNjaWQubmV1c3Rhci9OZXVzdGFyQ2VydGlmaWVkQ2FsbGVySWRDQTEuY3J0IDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMB0GA1UdDgQWBBSIQr1plt2oCHfJkgrZivE4u%2BZldjAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDRwAwRAIgU3vUOalywvB5Cr3e9Kq2sw253khxnKXlrvnmcG%2BPnHkCIDNnXbWEwwIu4znnW%2Bth0iBsezvJS0PHwpqp7lWNQf5k)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |


Generated: 16 Mar 23 19:18 UTC