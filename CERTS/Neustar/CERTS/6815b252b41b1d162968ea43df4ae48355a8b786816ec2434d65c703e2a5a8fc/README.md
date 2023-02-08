# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 261H

Tested At: 08 Feb 23 19:36 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 124 day(s)\
Subject: CN=SHAKEN 261H, O=Voyce, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN CA-1, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US\
Link: https://ssc.getsipnav.com/certs/b00badc474bbf0d965554e422647a4fcc426eb0c

[View certificate details](https://understandingwebpki.com/?cert=MIIC9jCCApygAwIBAgIUSFZASEJVrrV2G6eXUexaVWeAiCcwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0xMB4XDTIyMDYxMjE1MTI0MloXDTIzMDYxMjE1MTI0MlowMzELMAkGA1UEBhMCVVMxDjAMBgNVBAoMBVZveWNlMRQwEgYDVQQDDAtTSEFLRU4gMjYxSDBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABIPUEzYTJsAU%2BrBOuhANh02zh73liai3vPjYGMxK%2BjjRAlYdgoS0WUfcZlLaSnrkvhGSqkqbUPZMRUBj048kRvmjggE5MIIBNTAWBggrBgEFBQcBGgQKMAigBhYEMjYxSDAMBgNVHRMBAf8EAjAAMB8GA1UdIwQYMBaAFK%2FRyMLuckyD%2FD%2FtGadtHZCyB%2FA6MFsGCCsGAQUFBwEBBE8wTTBLBggrBgEFBQcwAoY%2FaHR0cDovL2NhY2VydHMtdXMuY2NpZC5uZXVzdGFyL05ldXN0YXJDZXJ0aWZpZWRDYWxsZXJJZENBMS5jcnQgMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwHQYDVR0OBBYEFMQ%2F3y%2F%2FwU8r%2FStiLi%2B01joV2v%2BPMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNIADBFAiAQAvvPvoJ%2B1Jjxs8pXiNeaej5Cm%2B0RYH%2FSLhHvovrrfgIhAKvU7YufEyUPbORTL5kbmwVv0Tl4O%2BFr%2FZJhccoLFjYL)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |


Generated: 08 Feb 23 19:45 UTC