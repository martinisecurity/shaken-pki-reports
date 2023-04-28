# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 0347

Tested At: 28 Apr 23 02:05 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 117 day(s)\
Subject: CN=SHAKEN 0347, O=Brantley Telephone Company Inc, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN CA-1, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US\
Link: https://cr-partner.ccid.neustar.biz/ccid/authn/v2/certs/181.237

[View certificate details](https://understandingwebpki.com/?cert=MIIDEDCCArWgAwIBAgIUGs2GUC36GGgUgOLY8IQSCUAkQWMwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0xMB4XDTIyMDgyMjIwMDY0N1oXDTIzMDgyMjIwMDY0N1owTDELMAkGA1UEBhMCVVMxJzAlBgNVBAoMHkJyYW50bGV5IFRlbGVwaG9uZSBDb21wYW55IEluYzEUMBIGA1UEAwwLU0hBS0VOIDAzNDcwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATvxpJHPxUpdYHcd7tYPVjwwE6GHK0bia3J2mjFIEFpGV%2BBEgFECsSoXEUASjeqgKWqMssiYxUoqgKfEwINQaRDo4IBOTCCATUwFgYIKwYBBQUHARoECjAIoAYWBDAzNDcwDAYDVR0TAQH%2FBAIwADAfBgNVHSMEGDAWgBSv0cjC7nJMg%2Fw%2F7RmnbR2QsgfwOjBbBggrBgEFBQcBAQRPME0wSwYIKwYBBQUHMAKGP2h0dHA6Ly9jYWNlcnRzLXVzLmNjaWQubmV1c3Rhci9OZXVzdGFyQ2VydGlmaWVkQ2FsbGVySWRDQTEuY3J0IDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMB0GA1UdDgQWBBSghrqDeRtLkmmgOH3BhBxR1MssEDAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSQAwRgIhAJLT1IzPgShFP0Cp7ztH91fyjwjD%2BjqEVwJFCYqcCPBmAiEAhesRway4JF4PdrASifHcag2NOkeSUlYbbEUcnC32JYQ%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |


Generated: 28 Apr 23 02:17 UTC