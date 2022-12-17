# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 700H

Tested At: 17 Dec 22 16:58 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 172 day(s)\
Subject: CN=SHAKEN 700H, O=Metro FiberNet LLC, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN CA-1, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US\
Link: https://cr-partner.ccid.neustar.biz/ccid/authn/v2/certs/142.143

[View certificate details](https://understandingwebpki.com/?cert=MIIDBDCCAqmgAwIBAgIUIc0R1p3BiVvw8o6EpUICqy8eI3QwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0xMB4XDTIyMDYwNjE5Mjk0NFoXDTIzMDYwNjE5Mjk0NFowQDELMAkGA1UEBhMCVVMxGzAZBgNVBAoMEk1ldHJvIEZpYmVyTmV0IExMQzEUMBIGA1UEAwwLU0hBS0VOIDcwMEgwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQ5IrO%2FyrUKM4wBbrEAZGSIVjGnw1JynMJ%2Flpkfboqb9wNkDq979HFjR%2FMfGAcng5Wk%2FXX9oIUuzzm3BbFoQ1ySo4IBOTCCATUwFgYIKwYBBQUHARoECjAIoAYWBDcwMEgwDAYDVR0TAQH%2FBAIwADAfBgNVHSMEGDAWgBSv0cjC7nJMg%2Fw%2F7RmnbR2QsgfwOjBbBggrBgEFBQcBAQRPME0wSwYIKwYBBQUHMAKGP2h0dHA6Ly9jYWNlcnRzLXVzLmNjaWQubmV1c3Rhci9OZXVzdGFyQ2VydGlmaWVkQ2FsbGVySWRDQTEuY3J0IDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMB0GA1UdDgQWBBQyJp98u8Xeob9gQ4kVH7yTJ9L5JjAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSQAwRgIhAN7Szz3gdV7mqKaDxIuKHdNRvdRt2mOOu1ox4eITBpb0AiEAhzpCH5QKKz21zdOxTHvY44NKckKmjtkso82g%2FVJ9SzQ%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 17 Dec 22 17:07 UTC