# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 4427

Tested At: 01 Nov 22 22:49 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 204 day(s)\
Subject: CN=SHAKEN 4427, O=Ziply Fiber, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN CA-1, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US\
Link: https://cr.ccid.neustar.biz/ccid/authn/v2/certs/11410.10158.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIC%2FDCCAqKgAwIBAgIUeocyyqbZvZn1UGL0fakvwt74AdswCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0xMB4XDTIyMDUyNDE2MTgwOVoXDTIzMDUyNDE2MTgwOVowOTELMAkGA1UEBhMCVVMxFDASBgNVBAoMC1ppcGx5IEZpYmVyMRQwEgYDVQQDDAtTSEFLRU4gNDQyNzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABF6hG%2F7%2BPteJyOqGnWAjcYr5viBbyQ%2BIVDM%2Fg0kdOe9fLektGY4ZMm6rADDbVxTJ1gMqKgCfL03AIajB%2BGkJ4UGjggE5MIIBNTAWBggrBgEFBQcBGgQKMAigBhYENDQyNzAMBgNVHRMBAf8EAjAAMB8GA1UdIwQYMBaAFK%2FRyMLuckyD%2FD%2FtGadtHZCyB%2FA6MFsGCCsGAQUFBwEBBE8wTTBLBggrBgEFBQcwAoY%2FaHR0cDovL2NhY2VydHMtdXMuY2NpZC5uZXVzdGFyL05ldXN0YXJDZXJ0aWZpZWRDYWxsZXJJZENBMS5jcnQgMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwHQYDVR0OBBYEFJJ6jE%2FFAWeZh7zxVq%2B1NxexweQSMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNIADBFAiEAz5pLICBPLLelAhZGxYfVXjno3wkIEas90tM36loIVmUCIBMEDd5jtdqFUWq33INzM6uEJq7YgAp2Fz9l5OjiLHVw)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 01/11/2022 at 22:50:57