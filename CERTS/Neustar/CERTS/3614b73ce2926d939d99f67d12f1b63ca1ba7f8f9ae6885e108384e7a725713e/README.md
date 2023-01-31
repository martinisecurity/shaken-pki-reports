# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 732J

Tested At: 31 Jan 23 21:40 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 113 day(s)\
Subject: CN=SHAKEN 732J, O=NUSO LLC, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN CA-1, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US\
Link: https://cr.ccid.neustar.biz/ccid/authn/v2/certs/11376.10159.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC%2BTCCAp%2BgAwIBAgIUbcmhQmCAMj%2Frc3sYmNMvx%2FmdTbQwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0xMB4XDTIyMDUyNDE4NDcxM1oXDTIzMDUyNDE4NDcxM1owNjELMAkGA1UEBhMCVVMxETAPBgNVBAoMCE5VU08gTExDMRQwEgYDVQQDDAtTSEFLRU4gNzMySjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABJGaFtHk7oTJVHNgGAB50sybe3LKjSdde8SCg6ZfYz9sh3MG3o9e4I%2FF2s%2Fzj%2FnOLc3aorZBxECJ2C6gzGsh0zWjggE5MIIBNTAWBggrBgEFBQcBGgQKMAigBhYENzMySjAMBgNVHRMBAf8EAjAAMB8GA1UdIwQYMBaAFK%2FRyMLuckyD%2FD%2FtGadtHZCyB%2FA6MFsGCCsGAQUFBwEBBE8wTTBLBggrBgEFBQcwAoY%2FaHR0cDovL2NhY2VydHMtdXMuY2NpZC5uZXVzdGFyL05ldXN0YXJDZXJ0aWZpZWRDYWxsZXJJZENBMS5jcnQgMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwHQYDVR0OBBYEFP%2F6vz%2B3xngxkxMvgHGwWFVVXWUjMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNIADBFAiBhSPNUBXlvwXuEjn4tXPHgHEPIr%2BBBlGU5Ch6syhn6qwIhAKPUTLt0e9hlxqgN70%2BTMTLeY5bvQD25AYvvf0h%2FnjkH)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 31 Jan 23 21:50 UTC