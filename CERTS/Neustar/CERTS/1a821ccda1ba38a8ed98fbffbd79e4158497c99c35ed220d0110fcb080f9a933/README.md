# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 295K

Tested At: 11 Jan 23 23:08 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 217 day(s)\
Subject: CN=SHAKEN 295K, O=LOGIN LLC, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN CA-1, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US\
Link: https://cr-partner.ccid.neustar.biz/ccid/authn/v2/certs/180.235

[View certificate details](https://understandingwebpki.com/?cert=MIIC%2BTCCAqCgAwIBAgIUSIkdAcR3T%2BjGcC%2FTp%2FsClTOccvEwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0xMB4XDTIyMDgxNjIwMDcxOFoXDTIzMDgxNjIwMDcxOFowNzELMAkGA1UEBhMCVVMxEjAQBgNVBAoMCUxPR0lOIExMQzEUMBIGA1UEAwwLU0hBS0VOIDI5NUswWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAARKT2AD0HQ%2Ft7wbbuZOwvNl%2F4DISAwbLGEquTJEcG09X19FgQtS04KVu1qOlv8RdPBjdgwGK%2BJdhS3BJaJsTDBGo4IBOTCCATUwFgYIKwYBBQUHARoECjAIoAYWBDI5NUswDAYDVR0TAQH%2FBAIwADAfBgNVHSMEGDAWgBSv0cjC7nJMg%2Fw%2F7RmnbR2QsgfwOjBbBggrBgEFBQcBAQRPME0wSwYIKwYBBQUHMAKGP2h0dHA6Ly9jYWNlcnRzLXVzLmNjaWQubmV1c3Rhci9OZXVzdGFyQ2VydGlmaWVkQ2FsbGVySWRDQTEuY3J0IDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMB0GA1UdDgQWBBRol0MDkJuSOfxhxwCfR%2FfKk2e2iDAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDRwAwRAIgAwzq2wheBi41Pp9KYUoiRvIo7%2B4gZ5yoT2AXPaxo4dsCIGsrBljNXjOk%2FSRYNtHfGTJlUOE%2FIAHWatZeeaMVuChX)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 11 Jan 23 23:18 UTC