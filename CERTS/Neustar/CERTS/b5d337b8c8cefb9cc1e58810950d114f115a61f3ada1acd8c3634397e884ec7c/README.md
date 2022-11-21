# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 7076

Tested At: 21 Nov 22 20:44 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 64 day(s)\
Subject: CN=SHAKEN 7076, O=Midco, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN CA-1, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US\
Link: https://cr.ccid.neustar.biz/ccid/authn/v2/certs/11158.10130.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC9jCCApygAwIBAgIUYc5DXHTgNpT4Muy3299umvG86okwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0xMB4XDTIyMDEyNDIwMDMwMFoXDTIzMDEyNDIwMDMwMFowMzELMAkGA1UEBhMCVVMxDjAMBgNVBAoMBU1pZGNvMRQwEgYDVQQDDAtTSEFLRU4gNzA3NjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABMLaRFmCo2DNmtRDef3WIa0NLcUTvCB06%2FKO5xOAR8ocVpFoBf2qarGtPpso4sbvHyPykSeWatkXloq88IRvtgejggE5MIIBNTAWBggrBgEFBQcBGgQKMAigBhYENzA3NjAMBgNVHRMBAf8EAjAAMB8GA1UdIwQYMBaAFK%2FRyMLuckyD%2FD%2FtGadtHZCyB%2FA6MFsGCCsGAQUFBwEBBE8wTTBLBggrBgEFBQcwAoY%2FaHR0cDovL2NhY2VydHMtdXMuY2NpZC5uZXVzdGFyL05ldXN0YXJDZXJ0aWZpZWRDYWxsZXJJZENBMS5jcnQgMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwHQYDVR0OBBYEFFQsDj3TLyk78kVMQZTd87x%2BO%2F8zMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNIADBFAiBmZpHS8%2Bnpzmo%2BBM5%2FWyhl4ckLeDmpQYvlNO6GjTp5hgIhAO5gi8m29WQtoA115KJ%2FaPlr8cBmpgsy%2BGFt%2FijxVDZY)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |


Generated: 21 Nov 22 20:55 UTC