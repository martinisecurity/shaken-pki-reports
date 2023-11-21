# STIR/SHAKEN CA Ecosystem Compliance

## Certificate PioneerTelephone

Tested At: 21 Nov 23 01:32 UTC\
Initial Validity Period: 1096 day(s)\
Remaining Validity Period: 76 day(s)\
Subject: C=US, ST=OK, L=Kingfisher, O=Central Office Switching, OU=ptci, CN=PioneerTelephone\
Issuer: C=US, O=Neustar Information Services Inc, OU=www.ccid.neustar, CN=Neustar Certified Caller ID CA-1\
Link: https://cr.ccid.neustar.biz/ccid/authn/v2/certs/11144.10173.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIEBzCCAu%2BgAwIBAgIUcLTkCWtEoefA%2B3TEanfxfVpVyz0wDQYJKoZIhvcNAQELBQAwfjEpMCcGA1UEAwwgTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIENBLTExGTAXBgNVBAsMEHd3dy5jY2lkLm5ldXN0YXIxKTAnBgNVBAoMIE5ldXN0YXIgSW5mb3JtYXRpb24gU2VydmljZXMgSW5jMQswCQYDVQQGEwJVUzAeFw0yMTAyMDMxMjQyMjJaFw0yNDAyMDQxMjQyMjJaMHwxGTAXBgNVBAMMEFBpb25lZXJUZWxlcGhvbmUxDTALBgNVBAsMBHB0Y2kxITAfBgNVBAoMGENlbnRyYWwgT2ZmaWNlIFN3aXRjaGluZzETMBEGA1UEBwwKS2luZ2Zpc2hlcjELMAkGA1UECAwCT0sxCzAJBgNVBAYTAlVTMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAED0beIOoCA6SkKdzKIsrMYb0ma%2B0rAnM7NL7P5wbUtgeWt81fMjwz%2BkD6mDze4p5Q4gk1FXSbd%2FxwkmzaIzvqvKOCAUgwggFEMAwGA1UdEwEB%2FwQCMAAwHwYDVR0jBBgwFoAUO7lcyzEXnEjP1Npm6422qNXhCfcwgYEGCCsGAQUFBwEBBHUwczBHBggrBgEFBQcwAoY7aHR0cDovL2NhY2VydHMuY2NpZC5uZXVzdGFyL05ldXN0YXJDZXJ0aWZpZWRDYWxsZXJJZENBMS5jcnQwKAYIKwYBBQUHMAGGHGh0dHA6Ly9vY3NwLWNhMS5jY2lkLm5ldXN0YXIwSAYDVR0fBEEwPzA9oDugOYY3aHR0cDovL2NybC5jY2lkLm5ldXN0YXIvTmV1c3RhckNlcnRpZmllZENhbGxlcklkQ0ExLmNybDAdBgNVHQ4EFgQUXYu%2Fmqflkuucq4URrpJj8Oz06aYwDgYDVR0PAQH%2FBAQDAgeAMBYGCCsGAQUFBwEaBAowCKAGFgQyMDE4MA0GCSqGSIb3DQEBCwUAA4IBAQCOref3ERGsSHOzbEpZFBinnXWwdjVBBAKhZ9wR5Ge19DkrldW2FnqvYP42hZ3cTJqqw3ST3JdArNjNDoxNyBkc%2Bw0QeXmz%2FcYwFZr0hnW30PNDQ8ImYIkkAxFTmuQF9CDP3a%2BZibWthtUmg5%2BQWyYI5x4iVPYO9g40m3VaULv3oQoahUnA1T9CgoeN1Ebrb0gplx2ySndjE2sf57XGqLOkB1qukZvLDV3r%2FyWBZpIwnSlpuhhSrbZgYypEZ5spUCOj8%2BhHNSvIqlNXsa9vtGL%2FtVqV3rVA%2Fc5ujqC%2Fa%2BySZnUwmaNs41859E3HSmlqf5jvHBcNhcOiDeCKHWD9%2B1tz)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_signature_algorithm](../../ISSUES/e_atis_signature_algorithm/README.md) | error | ATIS1000080 | SignatureAlgorithm field is not 'ecdsa-with-SHA256', got 1.2.840.113549.1.1.11 |
| [e_atis_ext_certificate_policies](../../ISSUES/e_atis_ext_certificate_policies/README.md) | error | ATIS1000080 | the Certificate Policies extension is not present |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common Name attribute 'PioneerTelephone' does not contain 'SHAKEN' |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_cn_spc
- e_atis_subject_key_identifier_size
- e_atis_subject_o_required
- e_atis_tn_auth_list_spc_format
- e_shaken_certificate_policies_id

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 21 Nov 23 01:55 UTC