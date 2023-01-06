# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 715J

Tested At: 06 Jan 23 02:54 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 194 day(s)\
Subject: CN=SHAKEN 715J, O=Coredial LLC, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN CA-1, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US\
Link: https://cr.ccid.neustar.biz/ccid/authn/v2/certs/11430.10198

[View certificate details](https://understandingwebpki.com/?cert=MIIC%2FTCCAqOgAwIBAgIUZ6cljPNm1ckMc6CRD15655HDQmIwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0xMB4XDTIyMDcxODE3MzMxNVoXDTIzMDcxODE3MzMxNVowOjELMAkGA1UEBhMCVVMxFTATBgNVBAoMDENvcmVkaWFsIExMQzEUMBIGA1UEAwwLU0hBS0VOIDcxNUowWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQ0NjMcIM0A9ZRXWsOcextsE24Q5xO%2F8fjfQyI97%2FmH0c4NGgK8u98WD11GPRpkfJwDmyv1opCgFakt0nx2DAZZo4IBOTCCATUwFgYIKwYBBQUHARoECjAIoAYWBDcxNUowDAYDVR0TAQH%2FBAIwADAfBgNVHSMEGDAWgBSv0cjC7nJMg%2Fw%2F7RmnbR2QsgfwOjBbBggrBgEFBQcBAQRPME0wSwYIKwYBBQUHMAKGP2h0dHA6Ly9jYWNlcnRzLXVzLmNjaWQubmV1c3Rhci9OZXVzdGFyQ2VydGlmaWVkQ2FsbGVySWRDQTEuY3J0IDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMB0GA1UdDgQWBBS9tTa1yJmpPDbi2AvAp0F%2BrA9rZTAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSAAwRQIgV7Kcrjimz8%2FrhHTozJXqXCzBhPXhGrur62EcVMbTg5ICIQCxB5Canu9vs9LavVB4WiHDxivDfCKYXXv7bfm9r97zlQ%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |


Generated: 06 Jan 23 03:03 UTC