# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 7914

Tested At: 08 Feb 23 19:34 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 129 day(s)\
Subject: CN=SHAKEN 7914, O=Easton Telecom Services LLC, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN CA-1, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US\
Link: https://cr-partner.ccid.neustar.biz/ccid/authn/v2/certs/81.164

[View certificate details](https://understandingwebpki.com/?cert=MIIDCzCCArKgAwIBAgIUEVd75u5wmPkZ0nYbUfecX1mjV%2FEwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0xMB4XDTIyMDYxNzEzNTAwMloXDTIzMDYxNzEzNTAwMlowSTELMAkGA1UEBhMCVVMxJDAiBgNVBAoMG0Vhc3RvbiBUZWxlY29tIFNlcnZpY2VzIExMQzEUMBIGA1UEAwwLU0hBS0VOIDc5MTQwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQt%2BdKbDSwGcsBv%2Fx%2BiKwHERf%2BSs7BJndDfWqlCtHHCFHpz50BBTqjLsHRc1c7ZTWp8bV0L7PureEuCozPbouOdo4IBOTCCATUwFgYIKwYBBQUHARoECjAIoAYWBDc5MTQwDAYDVR0TAQH%2FBAIwADAfBgNVHSMEGDAWgBSv0cjC7nJMg%2Fw%2F7RmnbR2QsgfwOjBbBggrBgEFBQcBAQRPME0wSwYIKwYBBQUHMAKGP2h0dHA6Ly9jYWNlcnRzLXVzLmNjaWQubmV1c3Rhci9OZXVzdGFyQ2VydGlmaWVkQ2FsbGVySWRDQTEuY3J0IDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMB0GA1UdDgQWBBSxYb7ozp9aDtpCMuAvSqJvEyReHzAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDRwAwRAIga0Cqjzfp4UarxLqKrkxFWIEy8vN37AFmtsDvnjN6wlYCIDYB3mmDvYzk%2FUluCHfJY%2BjXSY5Nvm6n0dE1DWaAgxuF)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |


Generated: 08 Feb 23 19:45 UTC