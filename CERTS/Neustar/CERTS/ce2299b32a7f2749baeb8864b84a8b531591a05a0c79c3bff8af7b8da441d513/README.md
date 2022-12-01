# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 074K

Tested At: 01 Dec 22 19:19 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 240 day(s)\
Subject: CN=SHAKEN 074K, O=Coztel Carrier, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN CA-1, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US\
Link: http://5.161.95.22/191c4c42dd7fa6115e84100637e42c99.cer

[View certificate details](https://understandingwebpki.com/?cert=MIIC%2FzCCAqWgAwIBAgIUKCdKr2zyjbQBv%2BS4i2jJH8ixacQwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0xMB4XDTIyMDcyODIwNTY0MloXDTIzMDcyODIwNTY0MlowPDELMAkGA1UEBhMCVVMxFzAVBgNVBAoMDkNvenRlbCBDYXJyaWVyMRQwEgYDVQQDDAtTSEFLRU4gMDc0SzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABDCaJkrHlAZLw5PQKHcFCMGU8JbmMTHBMMaHh26xQ1cxNGJUnWQqykZF9sKs2ZPESRjiid2Bs1Se85RYewHXIDajggE5MIIBNTAWBggrBgEFBQcBGgQKMAigBhYEMDc0SzAMBgNVHRMBAf8EAjAAMB8GA1UdIwQYMBaAFK%2FRyMLuckyD%2FD%2FtGadtHZCyB%2FA6MFsGCCsGAQUFBwEBBE8wTTBLBggrBgEFBQcwAoY%2FaHR0cDovL2NhY2VydHMtdXMuY2NpZC5uZXVzdGFyL05ldXN0YXJDZXJ0aWZpZWRDYWxsZXJJZENBMS5jcnQgMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwHQYDVR0OBBYEFBf57JbTLwnBT5IpeWpmNnCybJrhMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNIADBFAiEAobvQGJ2QDiX5GdCHRnbSV4SlXYVC%2Frbw0z5eP4f1VvcCIEJkOcr1nR79wLi21uyJWXiosjGQG4oGIVFvDblTWgGM)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |


Generated: 01 Dec 22 19:22 UTC