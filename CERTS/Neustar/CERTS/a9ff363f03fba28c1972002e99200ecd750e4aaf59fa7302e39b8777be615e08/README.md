# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 939H

Tested At: 04 Nov 22 01:11 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 225 day(s)\
Subject: CN=SHAKEN 939H, O=Commio, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN CA-1, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US\
Link: https://storage.googleapis.com/stirshaken/ShakeNBakeCert.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC9zCCAp2gAwIBAgIUIyzFjUCn%2FEZsBaSxTB%2BEnd57cEUwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0xMB4XDTIyMDYxNjE4MjYwM1oXDTIzMDYxNjE4MjYwM1owNDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBkNvbW1pbzEUMBIGA1UEAwwLU0hBS0VOIDkzOUgwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAS7DiqPpwx14jOKTtyEJyv7zv8tImt61oRbFVLHY6ysNU7tW6atOWqSXASyMB356OFBebBIzppnvYV6k%2FYOYM4ro4IBOTCCATUwFgYIKwYBBQUHARoECjAIoAYWBDkzOUgwDAYDVR0TAQH%2FBAIwADAfBgNVHSMEGDAWgBSv0cjC7nJMg%2Fw%2F7RmnbR2QsgfwOjBbBggrBgEFBQcBAQRPME0wSwYIKwYBBQUHMAKGP2h0dHA6Ly9jYWNlcnRzLXVzLmNjaWQubmV1c3Rhci9OZXVzdGFyQ2VydGlmaWVkQ2FsbGVySWRDQTEuY3J0IDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMB0GA1UdDgQWBBR%2BPZs3GTLQHwBDQ283xu8to3RigTAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSAAwRQIhAIW5%2B69VQzf6ToT%2Fv%2FjHwlF1HZbQpeEAzDk2Bnib8uiYAiAwxCOlHR9p3nMIij94bExqpMgS2VQcvCRowtjlT3s3dA%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |


Generated: 04 Nov 22 01:11 UTC