# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 150K

Tested At: 01 Nov 22 20:26 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 280 day(s)\
Subject: CN=SHAKEN 150K, O=White Label Communications, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN CA-1, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US\
Link: https://cr-partner.ccid.neustar.biz/ccid/authn/v2/certs/179.233

View: [Click to view](https://understandingwebpki.com/?cert=MIIDDDCCArGgAwIBAgIUGmsL%2FXfOma3CUQvEP4lr41iKOV0wCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0xMB4XDTIyMDgwODEyNTgzMVoXDTIzMDgwODEyNTgzMVowSDELMAkGA1UEBhMCVVMxIzAhBgNVBAoMGldoaXRlIExhYmVsIENvbW11bmljYXRpb25zMRQwEgYDVQQDDAtTSEFLRU4gMTUwSzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABCgd1Xz2Ym7PAUihAjjH9n4xko54ocrbKN5VIC3AnUJ4Y9Gu7NHJupwDx6ByL1vT5jRoWAOomHdUAGF4XuZJFF2jggE5MIIBNTAWBggrBgEFBQcBGgQKMAigBhYEMTUwSzAMBgNVHRMBAf8EAjAAMB8GA1UdIwQYMBaAFK%2FRyMLuckyD%2FD%2FtGadtHZCyB%2FA6MFsGCCsGAQUFBwEBBE8wTTBLBggrBgEFBQcwAoY%2FaHR0cDovL2NhY2VydHMtdXMuY2NpZC5uZXVzdGFyL05ldXN0YXJDZXJ0aWZpZWRDYWxsZXJJZENBMS5jcnQgMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwHQYDVR0OBBYEFEs1NuJpCxsVVePouuPE%2F2JVF1fsMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNJADBGAiEA%2FJJWRJQEu0QhmbiZNawmsQD%2BcLS86q36Y8G4VKX7aa4CIQCr9NLGuR6nwUAqExejwwmDegdF1AMmyHqd3aqdEIYKVQ%3D%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |


Generated: 01/11/2022 at 20:34:21