# STIR/SHAKEN CA Ecosystem Compliance
## Neustar

### Certificate f13529f6095c87fe53eaa93182d490716370d9ef
Tested At: 2022-10-26 21:01:06 +0000 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 229 day(s)\
Subject: CN=SHAKEN 033H, O=GetGo Communications LLC, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN CA-1, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US

Link: https://cr.ccid.neustar.biz/ccid/authn/v2/certs/11478.10162.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIDCTCCAq%2BgAwIBAgIUTHhRCKAFE8he2njl52hAy3HMU4EwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0xMB4XDTIyMDYxMjE1MTgwMloXDTIzMDYxMjE1MTgwMlowRjELMAkGA1UEBhMCVVMxITAfBgNVBAoMGEdldEdvIENvbW11bmljYXRpb25zIExMQzEUMBIGA1UEAwwLU0hBS0VOIDAzM0gwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATgAsKZZnH8lGin53ftAE%2BX3SCtoDiEDz2U%2FsnL09wF4gUNLynAXVPAc2Tbjo%2BHh7yzGmMMDSkIM2ZP6twYDghbo4IBOTCCATUwFgYIKwYBBQUHARoECjAIoAYWBDAzM0gwDAYDVR0TAQH%2FBAIwADAfBgNVHSMEGDAWgBSv0cjC7nJMg%2Fw%2F7RmnbR2QsgfwOjBbBggrBgEFBQcBAQRPME0wSwYIKwYBBQUHMAKGP2h0dHA6Ly9jYWNlcnRzLXVzLmNjaWQubmV1c3Rhci9OZXVzdGFyQ2VydGlmaWVkQ2FsbGVySWRDQTEuY3J0IDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMB0GA1UdDgQWBBSQhllkuP0JkmOlbg9koTraoCA4XTAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSAAwRQIgVkff5mKVJG4OeqzBoQOJE8a%2FZYZANaIPmFcmdrZ2EN8CIQD1CKjjvPH5GfI0cKkcZIs7o0zu0z44HDuOCY5b3f%2FZ1w%3D%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_cp1_3_subject_sn | error | CPv1.3 | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| e_sti_extension_unknown | error | ATIS-1000080v4 | STI certificate shall not include extensions that are not specified |
| e_sti_crl_distribution | error | ATIS-1000080v4 | CRL Distribution Point shall be reachable if the requesting IP address within the program ACLs |
| e_sti_certificate_policies | error | ATIS-1000080v4 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |

Generated: 26/10/2022 at 21:01:13