# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 473G

Tested At: 02 Nov 22 15:33 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 192 day(s)\
Subject: CN=SHAKEN 473G, O=Telengy LLC, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN CA-1, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US\
Link: https://cr.ccid.neustar.biz/ccid/authn/v2/certs/11538.10154

View: [View certificate details](https://understandingwebpki.com/?cert=MIIC%2BzCCAqKgAwIBAgIUPxg63r7A8hlM1Frs4yIQ4LJC4tYwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0xMB4XDTIyMDUxMzE0MzcxNloXDTIzMDUxMzE0MzcxNlowOTELMAkGA1UEBhMCVVMxFDASBgNVBAoMC1RlbGVuZ3kgTExDMRQwEgYDVQQDDAtTSEFLRU4gNDczRzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABP%2BvaIotA5DrTeCb0d5w5POPbI47ZZ2z9r75eDdXAQUJL1x9gNdqQvjcH2Qo5qa8dXdUoSd7ltAvdtmkPVjstDSjggE5MIIBNTAWBggrBgEFBQcBGgQKMAigBhYENDczRzAMBgNVHRMBAf8EAjAAMB8GA1UdIwQYMBaAFK%2FRyMLuckyD%2FD%2FtGadtHZCyB%2FA6MFsGCCsGAQUFBwEBBE8wTTBLBggrBgEFBQcwAoY%2FaHR0cDovL2NhY2VydHMtdXMuY2NpZC5uZXVzdGFyL05ldXN0YXJDZXJ0aWZpZWRDYWxsZXJJZENBMS5jcnQgMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwHQYDVR0OBBYEFFII0HK1h%2B1iwzePJ20Iqf6QWSZHMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNHADBEAiAtDU6%2BaeJJeTeMIn1AYj6xljJjk8dLvabWKQ5UC5R9%2FAIgfPceLtNxUFxWxF0sCYzWI3NzoSgxPgmVOiH6nsGbns0%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |


Generated: 02 Nov 22 15:41 UTC