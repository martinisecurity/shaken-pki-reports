# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 033H

Tested At: 02 Dec 22 07:27 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 193 day(s)\
Subject: CN=SHAKEN 033H, O=GetGo Communications LLC, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN CA-1, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US\
Link: https://cr.ccid.neustar.biz/ccid/authn/v2/certs/11478.10162.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIDCTCCAq%2BgAwIBAgIUTHhRCKAFE8he2njl52hAy3HMU4EwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0xMB4XDTIyMDYxMjE1MTgwMloXDTIzMDYxMjE1MTgwMlowRjELMAkGA1UEBhMCVVMxITAfBgNVBAoMGEdldEdvIENvbW11bmljYXRpb25zIExMQzEUMBIGA1UEAwwLU0hBS0VOIDAzM0gwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATgAsKZZnH8lGin53ftAE%2BX3SCtoDiEDz2U%2FsnL09wF4gUNLynAXVPAc2Tbjo%2BHh7yzGmMMDSkIM2ZP6twYDghbo4IBOTCCATUwFgYIKwYBBQUHARoECjAIoAYWBDAzM0gwDAYDVR0TAQH%2FBAIwADAfBgNVHSMEGDAWgBSv0cjC7nJMg%2Fw%2F7RmnbR2QsgfwOjBbBggrBgEFBQcBAQRPME0wSwYIKwYBBQUHMAKGP2h0dHA6Ly9jYWNlcnRzLXVzLmNjaWQubmV1c3Rhci9OZXVzdGFyQ2VydGlmaWVkQ2FsbGVySWRDQTEuY3J0IDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMB0GA1UdDgQWBBSQhllkuP0JkmOlbg9koTraoCA4XTAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSAAwRQIgVkff5mKVJG4OeqzBoQOJE8a%2FZYZANaIPmFcmdrZ2EN8CIQD1CKjjvPH5GfI0cKkcZIs7o0zu0z44HDuOCY5b3f%2FZ1w%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |


Generated: 02 Dec 22 07:30 UTC