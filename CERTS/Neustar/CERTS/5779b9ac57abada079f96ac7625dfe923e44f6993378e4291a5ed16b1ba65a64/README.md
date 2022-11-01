# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 845J

Tested At: 01 Nov 22 16:28 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 204 day(s)\
Subject: CN=SHAKEN 845J, O=Dialect\\, LLC, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN CA-1, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US\
Link: https://app.batchdialer.com/shakenv2.crt

View: [Click to view](https://understandingwebpki.com/?cert=MIIC%2FDCCAqOgAwIBAgIUbxh1jvB7EtE7IZhsXcEe04LNulUwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0xMB4XDTIyMDUyNDE1NDQzOVoXDTIzMDUyNDE1NDQzOVowOjELMAkGA1UEBhMCVVMxFTATBgNVBAoMDERpYWxlY3QsIExMQzEUMBIGA1UEAwwLU0hBS0VOIDg0NUowWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATxIbcWCrHnMzuz2FBl%2FeRmif9fL07o%2BTdN7aJM%2B7RQ3HufCSLxbphpBwa7qtZJfcU3WpQgSiWgghsCW8GmIEYjo4IBOTCCATUwFgYIKwYBBQUHARoECjAIoAYWBDg0NUowDAYDVR0TAQH%2FBAIwADAfBgNVHSMEGDAWgBSv0cjC7nJMg%2Fw%2F7RmnbR2QsgfwOjBbBggrBgEFBQcBAQRPME0wSwYIKwYBBQUHMAKGP2h0dHA6Ly9jYWNlcnRzLXVzLmNjaWQubmV1c3Rhci9OZXVzdGFyQ2VydGlmaWVkQ2FsbGVySWRDQTEuY3J0IDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMB0GA1UdDgQWBBSrX4ABoWAEsllgk%2F5B%2FnzaYqJYJjAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDRwAwRAIgBlAxpQCyqFbCw4ajCdbbHcsnJZxFq9rJFsCQSU6umDwCIEgs9DLIwAvtgtxuJ3Dq5ttnwcjhQ8%2FE6zTxJfz76md9)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS&#x2011;1000080 | STI certificate shall not include extensions that are not specified |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS&#x2011;1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |


Generated: 01/11/2022 at 16:30:07