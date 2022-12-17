# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 962J

Tested At: 17 Dec 22 00:03 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 209 day(s)\
Subject: CN=SHAKEN 962J, O=Versatel LLC, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN CA-1, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US\
Link: https://cr-partner.ccid.neustar.biz/ccid/authn/v2/certs/176.230

[View certificate details](https://understandingwebpki.com/?cert=MIIC%2FjCCAqOgAwIBAgIULK4A2AyxALSQ2%2BwDDZ6wsJWVBMgwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0xMB4XDTIyMDcxMzE3MzA1MVoXDTIzMDcxMzE3MzA1MVowOjELMAkGA1UEBhMCVVMxFTATBgNVBAoMDFZlcnNhdGVsIExMQzEUMBIGA1UEAwwLU0hBS0VOIDk2MkowWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATOmw8BE7JkeVhOOf7GlDZrxr8sxBfK%2FYJWcQ4qYGioaEgzv72tVH5oPjE2iJo%2BkE%2FWPIJ2RKZg3FLXenh41IrLo4IBOTCCATUwFgYIKwYBBQUHARoECjAIoAYWBDk2MkowDAYDVR0TAQH%2FBAIwADAfBgNVHSMEGDAWgBSv0cjC7nJMg%2Fw%2F7RmnbR2QsgfwOjBbBggrBgEFBQcBAQRPME0wSwYIKwYBBQUHMAKGP2h0dHA6Ly9jYWNlcnRzLXVzLmNjaWQubmV1c3Rhci9OZXVzdGFyQ2VydGlmaWVkQ2FsbGVySWRDQTEuY3J0IDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMB0GA1UdDgQWBBS2WvfM3zlxVOAtms%2BNYmqQiicaETAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSQAwRgIhAJpME4YK120irmO6vAoBIU%2BOg9KxAzlUlK14hmABlrtGAiEAjrVZe2gXVSVKL%2FJMNMzM8uYUPC5HV762BcWkhJAzt%2Bw%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |


Generated: 17 Dec 22 00:12 UTC