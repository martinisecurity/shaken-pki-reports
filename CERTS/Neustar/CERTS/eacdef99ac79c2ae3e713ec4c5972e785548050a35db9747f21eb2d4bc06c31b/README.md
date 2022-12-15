# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 139K

Tested At: 15 Dec 22 18:11 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 198 day(s)\
Subject: CN=SHAKEN 139K, O=FaxLogic LLC, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN CA-1, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US\
Link: https://cdn.pgxn.net/sti/20230701.cer

[View certificate details](https://understandingwebpki.com/?cert=MIIC%2FTCCAqOgAwIBAgIUWD6CzSZWU722h0mPyOB8WLfSe%2BswCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0xMB4XDTIyMDcwMTE2MjEzN1oXDTIzMDcwMTE2MjEzN1owOjELMAkGA1UEBhMCVVMxFTATBgNVBAoMDEZheExvZ2ljIExMQzEUMBIGA1UEAwwLU0hBS0VOIDEzOUswWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQsZ9MDAZSlLv%2Fc6SYpTLpJKMrcz4VJeymwT7UWadWVR6tvg%2FYX0IY1tb3C2F6oq9nbmWTFXS6EDPwcNDlpMxaKo4IBOTCCATUwFgYIKwYBBQUHARoECjAIoAYWBDEzOUswDAYDVR0TAQH%2FBAIwADAfBgNVHSMEGDAWgBSv0cjC7nJMg%2Fw%2F7RmnbR2QsgfwOjBbBggrBgEFBQcBAQRPME0wSwYIKwYBBQUHMAKGP2h0dHA6Ly9jYWNlcnRzLXVzLmNjaWQubmV1c3Rhci9OZXVzdGFyQ2VydGlmaWVkQ2FsbGVySWRDQTEuY3J0IDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMB0GA1UdDgQWBBTJJPJ7VsTLcs%2FqSthF4Ov3iMLA3DAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSAAwRQIgYu%2F1kJ0zNtjYGxdqCBD%2BJLVUE%2BTgTa1qzxupimMSpTACIQDFPCFjBwOCbUD9dohXBqP%2B%2BKMYgNHAakU46UvF1BFqrQ%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |


Generated: 15 Dec 22 18:35 UTC