# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 030J

Tested At: 11 Jan 23 23:08 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 84 day(s)\
Subject: CN=SHAKEN 030J, O=ANI Networks, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN CA-1, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US\
Link: https://certs.iverify-aninetworks.net/aninetworks_20220405.crt

[View certificate details](https://understandingwebpki.com/?cert=MIIC%2FjCCAqOgAwIBAgIUFTwaYq3wIlBGnP48PLdJDPyXqeYwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0xMB4XDTIyMDQwNTEyNTM1NVoXDTIzMDQwNTEyNTM1NVowOjELMAkGA1UEBhMCVVMxFTATBgNVBAoMDEFOSSBOZXR3b3JrczEUMBIGA1UEAwwLU0hBS0VOIDAzMEowWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAARqE6smv0hxD%2BfbFBON%2FaAjnT02NixG%2B7CMBWOFGgOm4WbLTHzpOqSEq88R63xH7BRbR17VRIE2bz0vYqtnr82Co4IBOTCCATUwFgYIKwYBBQUHARoECjAIoAYWBDAzMEowDAYDVR0TAQH%2FBAIwADAfBgNVHSMEGDAWgBSv0cjC7nJMg%2Fw%2F7RmnbR2QsgfwOjBbBggrBgEFBQcBAQRPME0wSwYIKwYBBQUHMAKGP2h0dHA6Ly9jYWNlcnRzLXVzLmNjaWQubmV1c3Rhci9OZXVzdGFyQ2VydGlmaWVkQ2FsbGVySWRDQTEuY3J0IDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMB0GA1UdDgQWBBRPFApR9vH1ZjU08pBcQtupADgnujAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSQAwRgIhAOlglgoUl%2F5BDAKzOs9Ujt2PGoJ%2FMU2%2BhLpk7SLmbFLOAiEA3kk0OelLylR0tAPmzatjgvYeWXXbqavvokDKccAvUgo%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |


Generated: 11 Jan 23 23:18 UTC