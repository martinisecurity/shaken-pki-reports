# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 005K

Tested At: 05 Jan 23 18:26 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 125 day(s)\
Subject: CN=SHAKEN 005K, O=PrimeVOX Communications, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN CA-1, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US\
Link: https://pvx1.s3.us-east-2.amazonaws.com/stirshaken/8448cc7eb8424d6ad5d2e6d71bcf6629.cer

[View certificate details](https://understandingwebpki.com/?cert=MIIDCTCCAq6gAwIBAgIUcCEni9NwdKcAwyrefaLU9XewzXwwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0xMB4XDTIyMDUxMDEwNDIwMloXDTIzMDUxMDEwNDIwMlowRTELMAkGA1UEBhMCVVMxIDAeBgNVBAoMF1ByaW1lVk9YIENvbW11bmljYXRpb25zMRQwEgYDVQQDDAtTSEFLRU4gMDA1SzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABKa%2Fi8f2iek7Ns%2FwBqCbbGQSTbBqGZRakBlqlXiHQSJZ%2BvdkGfE2BXeJSIeehQ%2FsrN9GW9F4M5FWDenNpLWTkbSjggE5MIIBNTAWBggrBgEFBQcBGgQKMAigBhYEMDA1SzAMBgNVHRMBAf8EAjAAMB8GA1UdIwQYMBaAFK%2FRyMLuckyD%2FD%2FtGadtHZCyB%2FA6MFsGCCsGAQUFBwEBBE8wTTBLBggrBgEFBQcwAoY%2FaHR0cDovL2NhY2VydHMtdXMuY2NpZC5uZXVzdGFyL05ldXN0YXJDZXJ0aWZpZWRDYWxsZXJJZENBMS5jcnQgMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwHQYDVR0OBBYEFC305lvxTvitXsmQxYq0EbrJEF4sMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNJADBGAiEAhIXaWh34WeHexLZ8zyvf4cQ1uEmgoMV0tHaqvyjQxB4CIQCc6BwcgbAIUDwa3U%2BqU9YrvF276FGyt8ylhle8gTopqA%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |


Generated: 05 Jan 23 18:35 UTC