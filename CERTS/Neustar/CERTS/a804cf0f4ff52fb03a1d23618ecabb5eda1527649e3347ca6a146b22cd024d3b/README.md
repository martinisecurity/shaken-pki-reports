# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 553J

Tested At: 07 Dec 22 18:46 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 198 day(s)\
Subject: CN=SHAKEN 553J, OU=VOIP, O=Whitesky Communications LLC, L=Northport, ST=AL, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN CA-1, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US\
Link: https://certs.sip.wtsky.net/prod/shaken-6-2023.cer

[View certificate details](https://understandingwebpki.com/?cert=MIIDPDCCAuKgAwIBAgIURaI2P9Jsx7RD6C1kjmV5TWZsFJ8wCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0xMB4XDTIyMDYyMzE2MzYxOFoXDTIzMDYyMzE2MzYxOFoweTELMAkGA1UEBhMCVVMxCzAJBgNVBAgMAkFMMRIwEAYDVQQHDAlOb3J0aHBvcnQxJDAiBgNVBAoMG1doaXRlc2t5IENvbW11bmljYXRpb25zIExMQzENMAsGA1UECwwEVk9JUDEUMBIGA1UEAwwLU0hBS0VOIDU1M0owWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQZ5o0jeSry66%2BhV2UrYHIR4MCsjypnjQ3RiIQiDuTD92R0LkalQ%2BixgE1qYPglLbBHB%2BT5ae7LlQsG6fsd8S1So4IBOTCCATUwFgYIKwYBBQUHARoECjAIoAYWBDU1M0owDAYDVR0TAQH%2FBAIwADAfBgNVHSMEGDAWgBSv0cjC7nJMg%2Fw%2F7RmnbR2QsgfwOjBbBggrBgEFBQcBAQRPME0wSwYIKwYBBQUHMAKGP2h0dHA6Ly9jYWNlcnRzLXVzLmNjaWQubmV1c3Rhci9OZXVzdGFyQ2VydGlmaWVkQ2FsbGVySWRDQTEuY3J0IDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMB0GA1UdDgQWBBRW12yTlqYszKS%2FeuugmMqnaHBD7DAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSAAwRQIgU3TF0DWAEd7kuNXeexP0JO8bL3fDCb30lAQbLiUl7AECIQDSuknhR%2FXIZ9OVZVIUiYStu2phAQjA1U0mA8MIflv7XA%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |


Generated: 07 Dec 22 18:54 UTC