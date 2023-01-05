# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 767J

Tested At: 05 Jan 23 18:07 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 125 day(s)\
Subject: CN=SHAKEN 767J, O=Newmax LLC dba Intermax Networks, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN CA-1, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US\
Link: https://cr-partner.ccid.neustar.biz/ccid/authn/v2/certs/127.128

[View certificate details](https://understandingwebpki.com/?cert=MIIDEjCCAregAwIBAgIUOsPGEqk%2BnPuUxYuMYI13OInRQFowCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0xMB4XDTIyMDUwOTE4NDQzNVoXDTIzMDUwOTE4NDQzNVowTjELMAkGA1UEBhMCVVMxKTAnBgNVBAoMIE5ld21heCBMTEMgZGJhIEludGVybWF4IE5ldHdvcmtzMRQwEgYDVQQDDAtTSEFLRU4gNzY3SjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABEKC%2Bs8MCrHYNhsGoE5L1yFAz0QvZOA6RuqVfuReUX4uhTlrnH%2Fi7uHiXgPTocwNYPAD7gVljNOAko7ekHmu6%2FqjggE5MIIBNTAWBggrBgEFBQcBGgQKMAigBhYENzY3SjAMBgNVHRMBAf8EAjAAMB8GA1UdIwQYMBaAFK%2FRyMLuckyD%2FD%2FtGadtHZCyB%2FA6MFsGCCsGAQUFBwEBBE8wTTBLBggrBgEFBQcwAoY%2FaHR0cDovL2NhY2VydHMtdXMuY2NpZC5uZXVzdGFyL05ldXN0YXJDZXJ0aWZpZWRDYWxsZXJJZENBMS5jcnQgMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwHQYDVR0OBBYEFIH2glkDzWMqFoXDzn%2FffEzJnYx6MA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNJADBGAiEAn069F3VlBWqRQAS4GI7%2FkZizNPBXizxTQfNgd5831MMCIQDziupAsR0KdP7ZPOT8%2Fq4zst%2BiFzWzDGgPbNuA%2Fd0s4Q%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |


Generated: 05 Jan 23 18:35 UTC