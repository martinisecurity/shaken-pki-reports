# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 763H

Tested At: 29 Dec 22 07:36 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 162 day(s)\
Subject: CN=SHAKEN 763H, O=PS Lightwave, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN CA-1, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US\
Link: https://cr.ccid.neustar.biz/ccid/authn/v2/certs/11358.10161.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC%2FTCCAqOgAwIBAgIUaX1v4LEFljYfshn4H%2BC%2BbRKRfbIwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0xMB4XDTIyMDYwODEzMzUwMloXDTIzMDYwODEzMzUwMlowOjELMAkGA1UEBhMCVVMxFTATBgNVBAoMDFBTIExpZ2h0d2F2ZTEUMBIGA1UEAwwLU0hBS0VOIDc2M0gwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASk9RISI9aENu9ZZdCpgmDlaBQ6UCPg5SwQYTAZw4USpd8%2BIZvbO%2BtTmVFyBnqUpJ0hrO0P6ZNS%2FsJaVfO5wgEvo4IBOTCCATUwFgYIKwYBBQUHARoECjAIoAYWBDc2M0gwDAYDVR0TAQH%2FBAIwADAfBgNVHSMEGDAWgBSv0cjC7nJMg%2Fw%2F7RmnbR2QsgfwOjBbBggrBgEFBQcBAQRPME0wSwYIKwYBBQUHMAKGP2h0dHA6Ly9jYWNlcnRzLXVzLmNjaWQubmV1c3Rhci9OZXVzdGFyQ2VydGlmaWVkQ2FsbGVySWRDQTEuY3J0IDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMB0GA1UdDgQWBBQXRMRy5faC%2FrnVm5GKWBxdAF2izDAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSAAwRQIhAMijM%2Be3%2B4EP9jpM0cWp1gjNJvMuWCC%2FdJSYXZWMdAoyAiApmHb1h47lel6UQPY7rG7ugafg9XlLPbruk3oTnH2Q9A%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |


Generated: 29 Dec 22 07:47 UTC