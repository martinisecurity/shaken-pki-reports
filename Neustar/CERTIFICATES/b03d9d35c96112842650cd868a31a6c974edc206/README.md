# STIR/SHAKEN CA Ecosystem Compliance
## Neustar

### Certificate b03d9d35c96112842650cd868a31a6c974edc206
Tested At: 2022-10-27 18:24:47 +0000 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 198 day(s)\
Subject: CN=SHAKEN 473G, O=Telengy LLC, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN CA-1, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US

Link: https://cr.ccid.neustar.biz/ccid/authn/v2/certs/11538.10154

View: [Click to view](https://understandingwebpki.com/?cert=MIIC%2BzCCAqKgAwIBAgIUPxg63r7A8hlM1Frs4yIQ4LJC4tYwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0xMB4XDTIyMDUxMzE0MzcxNloXDTIzMDUxMzE0MzcxNlowOTELMAkGA1UEBhMCVVMxFDASBgNVBAoMC1RlbGVuZ3kgTExDMRQwEgYDVQQDDAtTSEFLRU4gNDczRzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABP%2BvaIotA5DrTeCb0d5w5POPbI47ZZ2z9r75eDdXAQUJL1x9gNdqQvjcH2Qo5qa8dXdUoSd7ltAvdtmkPVjstDSjggE5MIIBNTAWBggrBgEFBQcBGgQKMAigBhYENDczRzAMBgNVHRMBAf8EAjAAMB8GA1UdIwQYMBaAFK%2FRyMLuckyD%2FD%2FtGadtHZCyB%2FA6MFsGCCsGAQUFBwEBBE8wTTBLBggrBgEFBQcwAoY%2FaHR0cDovL2NhY2VydHMtdXMuY2NpZC5uZXVzdGFyL05ldXN0YXJDZXJ0aWZpZWRDYWxsZXJJZENBMS5jcnQgMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwHQYDVR0OBBYEFFII0HK1h%2B1iwzePJ20Iqf6QWSZHMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNHADBEAiAtDU6%2BaeJJeTeMIn1AYj6xljJjk8dLvabWKQ5UC5R9%2FAIgfPceLtNxUFxWxF0sCYzWI3NzoSgxPgmVOiH6nsGbns0%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_cp1_3_subject_sn | error | CPv1.3 | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| e_sti_extension_unknown | error | ATIS-1000080v4 | STI certificate shall not include extensions that are not specified |
| e_sti_certificate_policies | error | ATIS-1000080v4 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| e_sti_crl_distribution | error | ATIS-1000080v4 | CRL Distribution Point shall be reachable if the requesting IP address within the program ACLs |

Generated: 27/10/2022 at 18:24:52