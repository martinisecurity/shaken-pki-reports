# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 821J

Tested At: 10 Nov 22 06:41 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 181 day(s)\
Subject: CN=SHAKEN 821J, O=Bluerock Communications, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN CA-1, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US\
Link: https://cr.ccid.neustar.biz/ccid/authn/v2/certs/11535.10153

[View certificate details](https://understandingwebpki.com/?cert=MIIDBzCCAq6gAwIBAgIUUO3niZ9Xqin26FOVQeO4%2F2k40PcwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0xMB4XDTIyMDUwOTE0NTM1NVoXDTIzMDUwOTE0NTM1NVowRTELMAkGA1UEBhMCVVMxIDAeBgNVBAoMF0JsdWVyb2NrIENvbW11bmljYXRpb25zMRQwEgYDVQQDDAtTSEFLRU4gODIxSjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABJPw5B8K7td3JOJhaES3YXzrJSdwR9E37LJr6ZR2zx4JaYuYL1fxd2%2F3zVHqpTip0U0DDGu3QOtg7tpG02N3RCmjggE5MIIBNTAWBggrBgEFBQcBGgQKMAigBhYEODIxSjAMBgNVHRMBAf8EAjAAMB8GA1UdIwQYMBaAFK%2FRyMLuckyD%2FD%2FtGadtHZCyB%2FA6MFsGCCsGAQUFBwEBBE8wTTBLBggrBgEFBQcwAoY%2FaHR0cDovL2NhY2VydHMtdXMuY2NpZC5uZXVzdGFyL05ldXN0YXJDZXJ0aWZpZWRDYWxsZXJJZENBMS5jcnQgMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwHQYDVR0OBBYEFKteKEbWJt3jEjAYXwvWqgfEAVJ4MA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNHADBEAiBcV1GBYTXZ3CvDW4w0GGnGCEhjADZ30cZsTetEsztpFAIgSmC9jx7tHt4mRPhmUWt7VU4b6vzQATdXnY1NAUrA7E4%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |


Generated: 10 Nov 22 06:43 UTC