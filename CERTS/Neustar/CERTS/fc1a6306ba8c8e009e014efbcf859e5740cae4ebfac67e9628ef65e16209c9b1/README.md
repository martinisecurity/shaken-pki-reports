# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 219K

Tested At: 11 Jan 23 21:49 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 219 day(s)\
Subject: CN=SHAKEN 219K, O=Avaya Cloud Inc, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN CA-1, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US\
Link: https://cr.ccid.neustar.biz/ccid/authn/v2/certs/11553.10201

[View certificate details](https://understandingwebpki.com/?cert=MIIDADCCAqagAwIBAgIUYw%2Br4VnGB9ph4ojEAURw%2BPSiXmwwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0xMB4XDTIyMDgxODE4MDcyOFoXDTIzMDgxODE4MDcyOFowPTELMAkGA1UEBhMCVVMxGDAWBgNVBAoMD0F2YXlhIENsb3VkIEluYzEUMBIGA1UEAwwLU0hBS0VOIDIxOUswWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQk78qXH%2FUe8zK34CrSlP4k%2BLiaeNuDaP3tV4s8g5Ydbdt0rDEUxA4pGHcsR6ppOMHRE%2BNQAPonyueyk2qOmfRso4IBOTCCATUwFgYIKwYBBQUHARoECjAIoAYWBDIxOUswDAYDVR0TAQH%2FBAIwADAfBgNVHSMEGDAWgBSv0cjC7nJMg%2Fw%2F7RmnbR2QsgfwOjBbBggrBgEFBQcBAQRPME0wSwYIKwYBBQUHMAKGP2h0dHA6Ly9jYWNlcnRzLXVzLmNjaWQubmV1c3Rhci9OZXVzdGFyQ2VydGlmaWVkQ2FsbGVySWRDQTEuY3J0IDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMB0GA1UdDgQWBBTR8S9PHMFntU6G%2BDjZCbOzwLttNzAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSAAwRQIgVasV0J3H8fWnf2ahDnBK1JpZLCbBX%2FZXdrdWIz3otKkCIQCk%2BUO8l2%2BelQgklaIeIGz7ZynWO1KIHv13t88kM2mvrg%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 11 Jan 23 21:59 UTC