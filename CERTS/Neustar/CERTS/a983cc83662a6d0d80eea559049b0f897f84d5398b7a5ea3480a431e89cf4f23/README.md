# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 938H

Tested At: 15 Nov 23 16:00 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: -125 day(s)\
Subject: CN=SHAKEN 938H, O=MCC Network Services LLC, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN CA-1, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US\
Link: https://cr-partner.ccid.neustar.biz/ccid/authn/v2/certs/103.228

[View certificate details](https://understandingwebpki.com/?cert=MIIDCTCCAq%2BgAwIBAgIUV6nebvn0qbPoLBiRV%2FLjoknB5gcwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0xMB4XDTIyMDcxMjE5NTAyM1oXDTIzMDcxMjE5NTAyM1owRjELMAkGA1UEBhMCVVMxITAfBgNVBAoMGE1DQyBOZXR3b3JrIFNlcnZpY2VzIExMQzEUMBIGA1UEAwwLU0hBS0VOIDkzOEgwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATQfvem6Q66Vev5NOPbJKkHuvhtmEjVk9VsKt4e97OENfqOGbLINL4tCtWcJmI69snibdpaQzbWbwDm8AmR8XB7o4IBOTCCATUwFgYIKwYBBQUHARoECjAIoAYWBDkzOEgwDAYDVR0TAQH%2FBAIwADAfBgNVHSMEGDAWgBSv0cjC7nJMg%2Fw%2F7RmnbR2QsgfwOjBbBggrBgEFBQcBAQRPME0wSwYIKwYBBQUHMAKGP2h0dHA6Ly9jYWNlcnRzLXVzLmNjaWQubmV1c3Rhci9OZXVzdGFyQ2VydGlmaWVkQ2FsbGVySWRDQTEuY3J0IDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMB0GA1UdDgQWBBSN%2FwyPwGmfDNHqTGXMvJs6H%2BoYWjAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSAAwRQIgV2D2UmGP2PDYKMKxIBRMdIVdN41PwKlcSq6rAvn%2BlEECIQCiD63iZvRhyUeON%2F3Pvv1ggUVfvxPgpKlSlylhDRefmQ%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |


Generated: 15 Nov 23 16:51 UTC