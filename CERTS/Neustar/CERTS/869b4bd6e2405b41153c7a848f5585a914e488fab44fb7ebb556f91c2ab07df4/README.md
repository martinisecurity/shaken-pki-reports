# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 712J

Tested At: 01 Nov 22 20:26 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 232 day(s)\
Subject: CN=SHAKEN 712J, O=ANPI Business LLC, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN CA-1, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US\
Link: https://cr-partner.ccid.neustar.biz/ccid/authn/v2/certs/87.166

View: [Click to view](https://understandingwebpki.com/?cert=MIIDAjCCAqigAwIBAgIUJffMY%2BZZIlMcBtiydPRk67k0GPAwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0xMB4XDTIyMDYyMTE4Mzk0OFoXDTIzMDYyMTE4Mzk0OFowPzELMAkGA1UEBhMCVVMxGjAYBgNVBAoMEUFOUEkgQnVzaW5lc3MgTExDMRQwEgYDVQQDDAtTSEFLRU4gNzEySjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABLIhkrsHewIbYIWvBZnAdpgPTi%2FsZAi75huIHnYfoR5ABTZzrh7XyXbJ7HRjLJ5scEpDQcuKVc89yqQO5mrKhdajggE5MIIBNTAWBggrBgEFBQcBGgQKMAigBhYENzEySjAMBgNVHRMBAf8EAjAAMB8GA1UdIwQYMBaAFK%2FRyMLuckyD%2FD%2FtGadtHZCyB%2FA6MFsGCCsGAQUFBwEBBE8wTTBLBggrBgEFBQcwAoY%2FaHR0cDovL2NhY2VydHMtdXMuY2NpZC5uZXVzdGFyL05ldXN0YXJDZXJ0aWZpZWRDYWxsZXJJZENBMS5jcnQgMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwHQYDVR0OBBYEFGamodhrj%2FuI0vNtcfs5lmloofHTMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNIADBFAiEAniq24eaqDz3qBvLFNM%2F1TscF82JuWW7VxePelseaZxoCIBLJMEaDwKh%2FmtbF%2FxA15kvWzWzSDOG3nA6u2b7M642o)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |


Generated: 01/11/2022 at 20:34:21