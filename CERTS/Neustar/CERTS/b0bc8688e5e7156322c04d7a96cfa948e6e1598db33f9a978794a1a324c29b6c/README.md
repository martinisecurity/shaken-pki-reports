# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 128K

Tested At: 07 Jan 23 19:09 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 167 day(s)\
Subject: CN=SHAKEN 128K, O=textPlus, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN CA-1, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US\
Link: https://cr-partner.ccid.neustar.biz/ccid/authn/v2/certs/156.170

[View certificate details](https://understandingwebpki.com/?cert=MIIC%2BTCCAp%2BgAwIBAgIUANQDtgoT7LyyO11i6EC6nWsi8ZQwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0xMB4XDTIyMDYyMjIwNDgxMFoXDTIzMDYyMjIwNDgxMFowNjELMAkGA1UEBhMCVVMxETAPBgNVBAoMCHRleHRQbHVzMRQwEgYDVQQDDAtTSEFLRU4gMTI4SzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABEZfTZMk8aVEioorGLlvhCa4v63r3L3WfqoZtvWWr1Ie54H7bg0QQInQkfYjvT%2BXqrGzAdkVWUWCu7eWliesLFujggE5MIIBNTAWBggrBgEFBQcBGgQKMAigBhYEMTI4SzAMBgNVHRMBAf8EAjAAMB8GA1UdIwQYMBaAFK%2FRyMLuckyD%2FD%2FtGadtHZCyB%2FA6MFsGCCsGAQUFBwEBBE8wTTBLBggrBgEFBQcwAoY%2FaHR0cDovL2NhY2VydHMtdXMuY2NpZC5uZXVzdGFyL05ldXN0YXJDZXJ0aWZpZWRDYWxsZXJJZENBMS5jcnQgMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwHQYDVR0OBBYEFI0FH%2Bjt10qHgLGIPA1tnU8QUegFMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNIADBFAiAEk3e7kHbTjUwSnTKr0D3cohhHgNPXbiBeRUxbYtsjjgIhAJCBjkeD3Y4hqhE9njtojdwlLyTbcK14%2FGpvUR%2FhpeBv)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 07 Jan 23 19:18 UTC