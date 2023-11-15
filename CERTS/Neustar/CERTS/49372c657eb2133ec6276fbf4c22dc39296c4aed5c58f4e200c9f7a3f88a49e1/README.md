# STIR/SHAKEN CA Ecosystem Compliance

## Certificate ATT SHAKEN 4036

Tested At: 15 Nov 23 16:09 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 269 day(s)\
Subject: CN=ATT SHAKEN 4036, OU=AT&T Services\\, Inc., O=ATT SHAKEN E-E, L=Dallas, ST=Texas, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN CA-1, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US\
Link: https://cert.sticr.att.net:8443/certs/att/0e6ca793-8797-40fe-9fde-c25af249bc12

[View certificate details](https://understandingwebpki.com/?cert=MIIDQTCCAuegAwIBAgIUFOEIoISlbCbV0REGnuz5zdWF2JIwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0xMB4XDTIzMDgxMTE1MTUzNloXDTI0MDgxMDE1MTUzNlowfzELMAkGA1UEBhMCVVMxDjAMBgNVBAgMBVRleGFzMQ8wDQYDVQQHDAZEYWxsYXMxFzAVBgNVBAoMDkFUVCBTSEFLRU4gRS1FMRwwGgYDVQQLDBNBVCZUIFNlcnZpY2VzLCBJbmMuMRgwFgYDVQQDDA9BVFQgU0hBS0VOIDQwMzYwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQEPk1cplKQyakNspg6sd3r98kNiLYR0NTJPEj9jubQk6yDTLbyiNt6Fw%2BFuTDXcugxy8mNym266%2BFLXJVWtGo1o4IBODCCATQwFgYIKwYBBQUHARoECjAIoAYWBDQwMzYwDAYDVR0TAQH%2FBAIwADAfBgNVHSMEGDAWgBSv0cjC7nJMg%2Fw%2F7RmnbR2QsgfwOjBaBggrBgEFBQcBAQROMEwwSgYIKwYBBQUHMAKGPmh0dHA6Ly9jYWNlcnRzLXVzLmNjaWQubmV1c3Rhci9OZXVzdGFyQ2VydGlmaWVkQ2FsbGVySWRDQTEuY3J0MBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwHQYDVR0OBBYEFOe8z72bQKiEL%2BQXOziVsTMhe7AhMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNIADBFAiB9oHIeYvqP%2FwjMJu2QNx2Rnq3Gt%2BwJTL5C5Ir%2BI%2FEoZwIhAIKR8josMVB5BFHDV%2FIZrmEqmtBlFzppd7Gq4b%2FsuhfC)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 15 Nov 23 17:17 UTC