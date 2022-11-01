# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 738J

Tested At: 01 Nov 22 22:48 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 239 day(s)\
Subject: CN=SHAKEN 738J, O=BCM One Inc, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN CA-1, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US\
Link: https://cr-partner.ccid.neustar.biz/ccid/authn/v2/certs/165.181

View: [Click to view](https://understandingwebpki.com/?cert=MIIC%2FTCCAqKgAwIBAgIUWFKILznP3SDUCE79BqJwyoiPPOcwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0xMB4XDTIyMDYyODIwMTUwNFoXDTIzMDYyODIwMTUwNFowOTELMAkGA1UEBhMCVVMxFDASBgNVBAoMC0JDTSBPbmUgSW5jMRQwEgYDVQQDDAtTSEFLRU4gNzM4SjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABPSlSzLo0Yc3Mb1QEXvi1Dtu4pCN9c0KtvyGPbgezr%2Bq4HMUealh8LOsCIX7U69auw7BOrpY5CA3m3dld%2BUljjejggE5MIIBNTAWBggrBgEFBQcBGgQKMAigBhYENzM4SjAMBgNVHRMBAf8EAjAAMB8GA1UdIwQYMBaAFK%2FRyMLuckyD%2FD%2FtGadtHZCyB%2FA6MFsGCCsGAQUFBwEBBE8wTTBLBggrBgEFBQcwAoY%2FaHR0cDovL2NhY2VydHMtdXMuY2NpZC5uZXVzdGFyL05ldXN0YXJDZXJ0aWZpZWRDYWxsZXJJZENBMS5jcnQgMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwHQYDVR0OBBYEFEJ0gel8N1sVw5J0CxccstCRu4nWMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNJADBGAiEAjncS6Z8A1kygo5xHZO5sabo%2FNeNtl%2FVoaTZC%2FaUnkOACIQCZmKvtF%2B0NokDikdN9mgagDZTN7TV9MsqSLRHc2mHebg%3D%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 01/11/2022 at 22:50:57