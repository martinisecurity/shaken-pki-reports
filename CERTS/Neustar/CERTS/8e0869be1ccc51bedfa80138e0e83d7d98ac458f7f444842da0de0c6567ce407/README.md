# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 1980

Tested At: 02 Jun 23 01:04 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 13 day(s)\
Subject: CN=SHAKEN 1980, O=Chickasaw Telephone Company, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN CA-1, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US\
Link: https://cr.ccid.neustar.biz/ccid/authn/v2/certs/11498.10164.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIDDTCCArKgAwIBAgIUYWiMXyC%2Fu8SnLQHboDZOpGpOEtIwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0xMB4XDTIyMDYxNDE4NDQzMloXDTIzMDYxNDE4NDQzMlowSTELMAkGA1UEBhMCVVMxJDAiBgNVBAoMG0NoaWNrYXNhdyBUZWxlcGhvbmUgQ29tcGFueTEUMBIGA1UEAwwLU0hBS0VOIDE5ODAwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAToviW3bTK7thekKpfWqFpCcv%2FV9wS%2FQAlBD3HDyHihi07rp3S1GYOI9VT44Gl9I9lVkUivGVIsZAluxaPO1lL9o4IBOTCCATUwFgYIKwYBBQUHARoECjAIoAYWBDE5ODAwDAYDVR0TAQH%2FBAIwADAfBgNVHSMEGDAWgBSv0cjC7nJMg%2Fw%2F7RmnbR2QsgfwOjBbBggrBgEFBQcBAQRPME0wSwYIKwYBBQUHMAKGP2h0dHA6Ly9jYWNlcnRzLXVzLmNjaWQubmV1c3Rhci9OZXVzdGFyQ2VydGlmaWVkQ2FsbGVySWRDQTEuY3J0IDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMB0GA1UdDgQWBBTqUUgnmaejep3cXpFb4c9LGNnSSTAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSQAwRgIhAPmj8ddZIK9%2FBLSxKiXU%2Bxbg8cRckLUEZMNvBnBk16ytAiEAxdRP72UNThNYhPM3MREVeKMpBA4giU8LWfyZnSv8HI4%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 02 Jun 23 01:12 UTC