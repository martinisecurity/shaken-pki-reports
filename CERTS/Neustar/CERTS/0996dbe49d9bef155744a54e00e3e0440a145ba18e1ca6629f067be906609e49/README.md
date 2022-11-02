# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 030K

Tested At: 02 Nov 22 21:15 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 196 day(s)\
Subject: CN=SHAKEN 030K, O=G12 COMMUNICATIONS, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN CA-1, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US\
Link: https://cr-partner.ccid.neustar.biz/ccid/authn/v2/certs/131.132

[View certificate details](https://understandingwebpki.com/?cert=MIIDAzCCAqmgAwIBAgIUGgiTWGRxyM%2B8IJpTWfgB3cW1uKMwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0xMB4XDTIyMDUxNzIxMTcyM1oXDTIzMDUxNzIxMTcyM1owQDELMAkGA1UEBhMCVVMxGzAZBgNVBAoMEkcxMiBDT01NVU5JQ0FUSU9OUzEUMBIGA1UEAwwLU0hBS0VOIDAzMEswWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAT%2BqmmjYb3Qs9FmKgaglHxIfOvKJYE47SHtzW7al2JfUv%2FZNonB21z1E%2BDbPtQg%2BtjooKwQnM9W0KNIeKr8m%2FdSo4IBOTCCATUwFgYIKwYBBQUHARoECjAIoAYWBDAzMEswDAYDVR0TAQH%2FBAIwADAfBgNVHSMEGDAWgBSv0cjC7nJMg%2Fw%2F7RmnbR2QsgfwOjBbBggrBgEFBQcBAQRPME0wSwYIKwYBBQUHMAKGP2h0dHA6Ly9jYWNlcnRzLXVzLmNjaWQubmV1c3Rhci9OZXVzdGFyQ2VydGlmaWVkQ2FsbGVySWRDQTEuY3J0IDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMB0GA1UdDgQWBBTqctnY3KMB1LH7J84g6a3pYn0F7DAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSAAwRQIgI1hyS2zRmn6Q3dJQZi0nIJT0tjsRG%2BMihdFuweKNYO4CIQDigRtOIqaItdGbUKTkcsRReWLDZQvXkp%2FI%2BOIt%2BAnwcQ%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |


Generated: 02 Nov 22 21:24 UTC