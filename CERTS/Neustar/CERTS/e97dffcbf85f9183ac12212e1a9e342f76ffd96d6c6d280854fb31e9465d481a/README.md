# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 710A

Tested At: 04 Nov 22 01:10 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 278 day(s)\
Subject: CN=SHAKEN 710A, O=Allo Communications LLC, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN CA-1, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US\
Link: https://cr.ccid.neustar.biz/ccid/authn/v2/certs/11438.10200.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIDCDCCAq6gAwIBAgIUMC5UASftOuJHOn%2Bttb%2FD7s8aVgMwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0xMB4XDTIyMDgwODE0MzAwNloXDTIzMDgwODE0MzAwNlowRTELMAkGA1UEBhMCVVMxIDAeBgNVBAoMF0FsbG8gQ29tbXVuaWNhdGlvbnMgTExDMRQwEgYDVQQDDAtTSEFLRU4gNzEwQTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABDBiTA5u6zpnl0L4ZksPIANZPA9o6Z6MdWIYobQSuIuDj1giXYGOZeplrlIWW8hCnGrhxAxBqVML5kpssQJF7gqjggE5MIIBNTAWBggrBgEFBQcBGgQKMAigBhYENzEwQTAMBgNVHRMBAf8EAjAAMB8GA1UdIwQYMBaAFK%2FRyMLuckyD%2FD%2FtGadtHZCyB%2FA6MFsGCCsGAQUFBwEBBE8wTTBLBggrBgEFBQcwAoY%2FaHR0cDovL2NhY2VydHMtdXMuY2NpZC5uZXVzdGFyL05ldXN0YXJDZXJ0aWZpZWRDYWxsZXJJZENBMS5jcnQgMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwHQYDVR0OBBYEFDgR9L37QNCdWqVFfn%2FUGglF4uDVMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNIADBFAiANYdZ4XFAJ%2Bnyf0S%2BVKh9NkZBmG%2FDhK%2BeSg7aSPEj14wIhAMqyDCd5crkhw6U6S6WJnFhCZuOgWo7RyL27IPX5tPlM)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |


Generated: 04 Nov 22 01:11 UTC