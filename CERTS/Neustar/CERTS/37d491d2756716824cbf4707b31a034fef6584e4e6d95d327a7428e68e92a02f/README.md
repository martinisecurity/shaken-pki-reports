# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 049K

Tested At: 09 Mar 23 22:59 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 112 day(s)\
Subject: CN=SHAKEN 049K, O=Dialpad Inc., C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN CA-1, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US\
Link: https://cr-partner.ccid.neustar.biz/ccid/authn/v2/certs/168.185

[View certificate details](https://understandingwebpki.com/?cert=MIIC%2FDCCAqOgAwIBAgIUJtZdFZG%2BPOJnRauJyFQZGZ%2Be%2BQEwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0xMB4XDTIyMDYyOTIyMDAwMFoXDTIzMDYyOTIyMDAwMFowOjELMAkGA1UEBhMCVVMxFTATBgNVBAoMDERpYWxwYWQgSW5jLjEUMBIGA1UEAwwLU0hBS0VOIDA0OUswWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASKbsF18a7tTg52yppgVZetZUFWaXUvKpVj0FkbOOkb%2BrjUgvxRUpzueQbAi3EomV9iZVsq7pwJ3L6%2BjutXRt%2FDo4IBOTCCATUwFgYIKwYBBQUHARoECjAIoAYWBDA0OUswDAYDVR0TAQH%2FBAIwADAfBgNVHSMEGDAWgBSv0cjC7nJMg%2Fw%2F7RmnbR2QsgfwOjBbBggrBgEFBQcBAQRPME0wSwYIKwYBBQUHMAKGP2h0dHA6Ly9jYWNlcnRzLXVzLmNjaWQubmV1c3Rhci9OZXVzdGFyQ2VydGlmaWVkQ2FsbGVySWRDQTEuY3J0IDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMB0GA1UdDgQWBBTrsv0t2Rwc9ETr2LCwTaWXR6ugqDAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDRwAwRAIgYGP0jihGvnVxfuofK3EYR%2BbOTMRfAwvIo8q5tKZXhVsCIGnCzceK0q2Rid%2FKboizBsQzMgpaM8SSFnHtUMmwKgKP)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |


Generated: 10 Mar 23 02:25 UTC