# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 755B

Tested At: 09 Mar 23 22:59 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 105 day(s)\
Subject: CN=SHAKEN 755B, O=Hiawatha Broadband Communications, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN CA-1, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US\
Link: https://cr-partner.ccid.neustar.biz/ccid/authn/v2/certs/159.173

[View certificate details](https://understandingwebpki.com/?cert=MIIDETCCArigAwIBAgIUNZCbpfsHcRlWTkONnZvPgG9O1TcwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0xMB4XDTIyMDYyMjIwNTYxMVoXDTIzMDYyMjIwNTYxMVowTzELMAkGA1UEBhMCVVMxKjAoBgNVBAoMIUhpYXdhdGhhIEJyb2FkYmFuZCBDb21tdW5pY2F0aW9uczEUMBIGA1UEAwwLU0hBS0VOIDc1NUIwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATa%2BdNGlbTAk3QiQ6fumR%2F83WI2v9n7HikRba840Qva3qOhfIlV4J5MDRwtvXqgHPLohzgfU99mPFYNXdhuK%2Bf2o4IBOTCCATUwFgYIKwYBBQUHARoECjAIoAYWBDc1NUIwDAYDVR0TAQH%2FBAIwADAfBgNVHSMEGDAWgBSv0cjC7nJMg%2Fw%2F7RmnbR2QsgfwOjBbBggrBgEFBQcBAQRPME0wSwYIKwYBBQUHMAKGP2h0dHA6Ly9jYWNlcnRzLXVzLmNjaWQubmV1c3Rhci9OZXVzdGFyQ2VydGlmaWVkQ2FsbGVySWRDQTEuY3J0IDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMB0GA1UdDgQWBBTLui5dzb%2Bcs2na7kGF4da1qD4VPzAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDRwAwRAIgLD3xoQoA0pZN9pPqM6STSQDZJ%2Bzdt79tDP%2BBJKbdpLMCIF58%2Bm6U%2B75pt8VpilB%2BGATvLFeKngFes4DAqCAHcijL)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |


Generated: 10 Mar 23 02:25 UTC