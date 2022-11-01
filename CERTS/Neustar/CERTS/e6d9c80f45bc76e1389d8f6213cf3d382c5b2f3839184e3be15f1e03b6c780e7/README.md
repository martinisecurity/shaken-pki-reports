# STIR/SHAKEN CA Ecosystem Compliance

## Certificate vonage

Tested At: 01 Nov 22 20:26 UTC\
Initial Validity Period: 1096 day(s)\
Remaining Validity Period: 47 day(s)\
Subject: C=US, ST=NJ, L=Holmdel, O=Vonage Eng, OU=Vonage Eng, CN=vonage\
Issuer: C=US, O=Neustar Information Services Inc, OU=www.ccid.neustar, CN=Neustar Certified Caller ID CA-1\
Link: https://cr.ccid.neustar.biz/ccid/authn/v2/certs/11053.10194

View: [Click to view](https://understandingwebpki.com/?cert=MIID8jCCAtqgAwIBAgIUT%2BaXkf0amRRF7kW02Mqg6nFobmkwDQYJKoZIhvcNAQELBQAwfjEpMCcGA1UEAwwgTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIENBLTExGTAXBgNVBAsMEHd3dy5jY2lkLm5ldXN0YXIxKTAnBgNVBAoMIE5ldXN0YXIgSW5mb3JtYXRpb24gU2VydmljZXMgSW5jMQswCQYDVQQGEwJVUzAeFw0xOTEyMTgyMDA4MjVaFw0yMjEyMTgyMDA4MjVaMGcxDzANBgNVBAMMBnZvbmFnZTETMBEGA1UECwwKVm9uYWdlIEVuZzETMBEGA1UECgwKVm9uYWdlIEVuZzEQMA4GA1UEBwwHSG9sbWRlbDELMAkGA1UECAwCTkoxCzAJBgNVBAYTAlVTMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE77qasdW3CLativLGa%2FrsJBzVem0LGkSV7EtGDu5IoAAS18srDmq%2B8baIqoRJ7Kk41tCWfp8VybhIP3NJwEBwBKOCAUgwggFEMAwGA1UdEwEB%2FwQCMAAwHwYDVR0jBBgwFoAUO7lcyzEXnEjP1Npm6422qNXhCfcwgYEGCCsGAQUFBwEBBHUwczBHBggrBgEFBQcwAoY7aHR0cDovL2NhY2VydHMuY2NpZC5uZXVzdGFyL05ldXN0YXJDZXJ0aWZpZWRDYWxsZXJJZENBMS5jcnQwKAYIKwYBBQUHMAGGHGh0dHA6Ly9vY3NwLWNhMS5jY2lkLm5ldXN0YXIwSAYDVR0fBEEwPzA9oDugOYY3aHR0cDovL2NybC5jY2lkLm5ldXN0YXIvTmV1c3RhckNlcnRpZmllZENhbGxlcklkQ0ExLmNybDAdBgNVHQ4EFgQU%2FXpyIcR3tQS3RsQGKgGHqn8IP5gwDgYDVR0PAQH%2FBAQDAgeAMBYGCCsGAQUFBwEaBAowCKAGFgQxOTdEMA0GCSqGSIb3DQEBCwUAA4IBAQBnyM%2B7c78wj3K7iZtWZtCtkUcky3eHOv9FJRCtG%2BXLv8OPXEBdUXSJBXtT4X7XDLFyn5Vgjqolkb7Mia5iqnOUtPKmLqnXyrOTFbbFbXSCazZf6reg5BMOZxvJTEOESbn36AsT8FoDFvXmmW97leSQt%2F2PMFtYV6CyhUyqY227zt41QmGeW6E3U7CAT7XUH4g4lwTbq%2FK8T1VhQgO58QPjjhIlF0n6eE3Wk2F0SE7Vc79WRsEPfnt9apwfO9TI9fDFnKZ34ElNV1BFW74GQANGKCq4NfySQJ7nS2nYLlaUKioJKQ%2Bw5XftmNLSBGINlrCnz%2BuijZiVOTOaQGzq%2BnAu)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


### Not Effective

- e_atis_authority_key_identifier
- e_atis_basic_constraints
- e_atis_certificate_policies
- e_atis_crl_distribution
- e_atis_extension_unknown
- e_atis_issuer_dn
- e_atis_key_usage
- e_atis_serial_number
- e_atis_signature_algorithm
- e_atis_subject
- e_atis_subject_cn
- e_atis_subject_key_identifier
- e_atis_subject_public_key
- e_atis_tn_auth_list
- e_atis_version
- e_us_cp_ambiguous_identifier
- e_us_cp_subject_sn
- w_cp_1_3_subject_email

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 01/11/2022 at 20:34:21