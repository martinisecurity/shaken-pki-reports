# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN-925C

Tested At: 18 Aug 25 20:17 UTC\
Initial Validity Period: 1096 day(s)\
Remaining Validity Period: -510 day(s)\
Subject: C=US, ST=MN, L=Brainerd, O=Consolidated Telephone Company, OU=CTC VOIP, CN=SHAKEN-925C\
Issuer: C=US, O=Neustar Information Services Inc, OU=www.ccid.neustar, CN=Neustar Certified Caller ID CA-1\
Link: https://cr-partner.ccid.neustar.biz/ccid/authn/v2/certs/25.196

[View certificate details](https://x509.io/?cert=MIIECjCCAvKgAwIBAgIUDuz%2BRZb3BkBSa%2FgF5q2D3EMQTOMwDQYJKoZIhvcNAQELBQAwfjEpMCcGA1UEAwwgTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIENBLTExGTAXBgNVBAsMEHd3dy5jY2lkLm5ldXN0YXIxKTAnBgNVBAoMIE5ldXN0YXIgSW5mb3JtYXRpb24gU2VydmljZXMgSW5jMQswCQYDVQQGEwJVUzAeFw0yMTAzMjYxMjA1NTlaFw0yNDAzMjYxMjA1NTlaMH8xFDASBgNVBAMMC1NIQUtFTi05MjVDMREwDwYDVQQLDAhDVEMgVk9JUDEnMCUGA1UECgweQ29uc29saWRhdGVkIFRlbGVwaG9uZSBDb21wYW55MREwDwYDVQQHDAhCcmFpbmVyZDELMAkGA1UECAwCTU4xCzAJBgNVBAYTAlVTMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEyhY%2FcpH6EPLxsptu79GJQaVUvUg3Cnv98RUW%2B0cu99LRmW3KdwqW%2BI%2FM00PGqPBQ0MzFKVlMQy%2FNFLMmRMy6yaOCAUgwggFEMAwGA1UdEwEB%2FwQCMAAwHwYDVR0jBBgwFoAUO7lcyzEXnEjP1Npm6422qNXhCfcwgYEGCCsGAQUFBwEBBHUwczBHBggrBgEFBQcwAoY7aHR0cDovL2NhY2VydHMuY2NpZC5uZXVzdGFyL05ldXN0YXJDZXJ0aWZpZWRDYWxsZXJJZENBMS5jcnQwKAYIKwYBBQUHMAGGHGh0dHA6Ly9vY3NwLWNhMS5jY2lkLm5ldXN0YXIwSAYDVR0fBEEwPzA9oDugOYY3aHR0cDovL2NybC5jY2lkLm5ldXN0YXIvTmV1c3RhckNlcnRpZmllZENhbGxlcklkQ0ExLmNybDAdBgNVHQ4EFgQUNQPJ%2B0YCNrPPYrnlqDAAFxLzWI4wDgYDVR0PAQH%2FBAQDAgeAMBYGCCsGAQUFBwEaBAowCKAGFgQ5MjVDMA0GCSqGSIb3DQEBCwUAA4IBAQCtnOfCnVxNsnn8BNNzkoWXMvbXkSn8ejSvI0whmlCMc3IFtn8ZhbmJ2XuciPIL7UmNOjlV7ao1taMLceM46LYisNdFJf19VhDNLHfDTvn4oq%2F9CQhY%2F7j075uTwcWM3%2F84dwHGZx3ueQY9s1vp7B%2FTa7jC3PYamLtE64bsqUD%2FOOyb5AhTJ8s5j3HAoUPMa%2FDIyAgsoXlOEKtUoK%2BBgPuyTJ%2F1reo%2Fne2Cdz3QlG7tb8hEcKSRTtEmtqTa1WKxZrxaq06icA0QZTcO4Pp6drPBoRUY6HqdToMUVYfBf8FjnIkDGro%2BYKpzzfAedyVfVP9obL2zKjqtgFs5yZnZyQDs)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_signature_algorithm](../../ISSUES/e_atis_signature_algorithm/README.md) | error | ATIS1000080 | SignatureAlgorithm field is not 'ecdsa-with-SHA256', got 1.2.840.113549.1.1.11 |
| [e_atis_ext_certificate_policies](../../ISSUES/e_atis_ext_certificate_policies/README.md) | error | ATIS1000080 | The Certificate Policies extension is not present |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_subject_cn_spc
- e_atis_subject_key_identifier_size
- e_atis_subject_o_required
- e_atis_tn_auth_list_spc_format
- e_shaken_certificate_policies_id

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 18 Aug 25 21:13 UTC