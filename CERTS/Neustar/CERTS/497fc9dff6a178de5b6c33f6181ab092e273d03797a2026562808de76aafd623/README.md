# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN-925C

Tested At: 08 Feb 23 19:34 UTC\
Initial Validity Period: 1096 day(s)\
Remaining Validity Period: 412 day(s)\
Subject: C=US, ST=MN, L=Brainerd, O=Consolidated Telephone Company, OU=CTC VOIP, CN=SHAKEN-925C\
Issuer: C=US, O=Neustar Information Services Inc, OU=www.ccid.neustar, CN=Neustar Certified Caller ID CA-1\
Link: https://cr-partner.ccid.neustar.biz/ccid/authn/v2/certs/25.196

[View certificate details](https://understandingwebpki.com/?cert=MIIECjCCAvKgAwIBAgIUDuz%2BRZb3BkBSa%2FgF5q2D3EMQTOMwDQYJKoZIhvcNAQELBQAwfjEpMCcGA1UEAwwgTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIENBLTExGTAXBgNVBAsMEHd3dy5jY2lkLm5ldXN0YXIxKTAnBgNVBAoMIE5ldXN0YXIgSW5mb3JtYXRpb24gU2VydmljZXMgSW5jMQswCQYDVQQGEwJVUzAeFw0yMTAzMjYxMjA1NTlaFw0yNDAzMjYxMjA1NTlaMH8xFDASBgNVBAMMC1NIQUtFTi05MjVDMREwDwYDVQQLDAhDVEMgVk9JUDEnMCUGA1UECgweQ29uc29saWRhdGVkIFRlbGVwaG9uZSBDb21wYW55MREwDwYDVQQHDAhCcmFpbmVyZDELMAkGA1UECAwCTU4xCzAJBgNVBAYTAlVTMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEyhY%2FcpH6EPLxsptu79GJQaVUvUg3Cnv98RUW%2B0cu99LRmW3KdwqW%2BI%2FM00PGqPBQ0MzFKVlMQy%2FNFLMmRMy6yaOCAUgwggFEMAwGA1UdEwEB%2FwQCMAAwHwYDVR0jBBgwFoAUO7lcyzEXnEjP1Npm6422qNXhCfcwgYEGCCsGAQUFBwEBBHUwczBHBggrBgEFBQcwAoY7aHR0cDovL2NhY2VydHMuY2NpZC5uZXVzdGFyL05ldXN0YXJDZXJ0aWZpZWRDYWxsZXJJZENBMS5jcnQwKAYIKwYBBQUHMAGGHGh0dHA6Ly9vY3NwLWNhMS5jY2lkLm5ldXN0YXIwSAYDVR0fBEEwPzA9oDugOYY3aHR0cDovL2NybC5jY2lkLm5ldXN0YXIvTmV1c3RhckNlcnRpZmllZENhbGxlcklkQ0ExLmNybDAdBgNVHQ4EFgQUNQPJ%2B0YCNrPPYrnlqDAAFxLzWI4wDgYDVR0PAQH%2FBAQDAgeAMBYGCCsGAQUFBwEaBAowCKAGFgQ5MjVDMA0GCSqGSIb3DQEBCwUAA4IBAQCtnOfCnVxNsnn8BNNzkoWXMvbXkSn8ejSvI0whmlCMc3IFtn8ZhbmJ2XuciPIL7UmNOjlV7ao1taMLceM46LYisNdFJf19VhDNLHfDTvn4oq%2F9CQhY%2F7j075uTwcWM3%2F84dwHGZx3ueQY9s1vp7B%2FTa7jC3PYamLtE64bsqUD%2FOOyb5AhTJ8s5j3HAoUPMa%2FDIyAgsoXlOEKtUoK%2BBgPuyTJ%2F1reo%2Fne2Cdz3QlG7tb8hEcKSRTtEmtqTa1WKxZrxaq06icA0QZTcO4Pp6drPBoRUY6HqdToMUVYfBf8FjnIkDGro%2BYKpzzfAedyVfVP9obL2zKjqtgFs5yZnZyQDs)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |

### Not Effective

- e_atis_extension_unknown
- e_atis_serial_number
- e_atis_signature_algorithm
- e_atis_subject_cn

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 08 Feb 23 19:45 UTC