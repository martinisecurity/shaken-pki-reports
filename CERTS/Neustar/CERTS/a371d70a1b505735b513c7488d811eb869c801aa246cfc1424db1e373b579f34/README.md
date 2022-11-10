# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 709J

Tested At: 10 Nov 22 23:19 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 240 day(s)\
Subject: CN=SHAKEN 709J, O=Low Latency Communications LLC, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN CA-1, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US\
Link: https://cr-partner.ccid.neustar.biz/ccid/authn/v2/certs/77.226

[View certificate details](https://understandingwebpki.com/?cert=MIIDDzCCArWgAwIBAgIUAevDGcykSvlz2sWbcL5kI83AaHswCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0xMB4XDTIyMDcwODE3NDkyMFoXDTIzMDcwODE3NDkyMFowTDELMAkGA1UEBhMCVVMxJzAlBgNVBAoMHkxvdyBMYXRlbmN5IENvbW11bmljYXRpb25zIExMQzEUMBIGA1UEAwwLU0hBS0VOIDcwOUowWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAScY98f2UhC6wmzOlgF6xGAY%2Ftdwq%2FqbcFx5CoqR4i8kx%2F9o1DlnDsGTuA20s9smzS8DnnVsmy%2FOGIl5oPTlFPuo4IBOTCCATUwFgYIKwYBBQUHARoECjAIoAYWBDcwOUowDAYDVR0TAQH%2FBAIwADAfBgNVHSMEGDAWgBSv0cjC7nJMg%2Fw%2F7RmnbR2QsgfwOjBbBggrBgEFBQcBAQRPME0wSwYIKwYBBQUHMAKGP2h0dHA6Ly9jYWNlcnRzLXVzLmNjaWQubmV1c3Rhci9OZXVzdGFyQ2VydGlmaWVkQ2FsbGVySWRDQTEuY3J0IDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMB0GA1UdDgQWBBRxSPgSIEHsiMZOe%2BfVWTGjnoOljjAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSAAwRQIhAISPX0vl2jalTTUNPspcFIcqIjqb6XRWfOli2dam9l0uAiBtrQWSR1pHEX7smb1xyAFcC89EoiQQtGVlKQ3VgB8Zdg%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |


Generated: 10 Nov 22 23:30 UTC