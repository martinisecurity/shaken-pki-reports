# STIR/SHAKEN CA Ecosystem Compliance

## Certificate vonage

Tested At: 31 Oct 22 16:40 UTC\
Initial Validity Period: 1096 day(s)\
Remaining Validity Period: 49 day(s)\
Subject: C=US, ST=NJ, L=Holmdel, O=Vonage Eng, OU=Vonage Eng, CN=vonage\
Issuer: C=US, O=Neustar Information Services Inc, OU=www.ccid.neustar, CN=Neustar Certified Caller ID CA-1\
Link: https://cr.ccid.neustar.biz/ccid/authn/v2/certs/11063.10194.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIID8jCCAtqgAwIBAgIUT%2BaXkf0amRRF7kW02Mqg6nFobmkwDQYJKoZIhvcNAQELBQAwfjEpMCcGA1UEAwwgTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIENBLTExGTAXBgNVBAsMEHd3dy5jY2lkLm5ldXN0YXIxKTAnBgNVBAoMIE5ldXN0YXIgSW5mb3JtYXRpb24gU2VydmljZXMgSW5jMQswCQYDVQQGEwJVUzAeFw0xOTEyMTgyMDA4MjVaFw0yMjEyMTgyMDA4MjVaMGcxDzANBgNVBAMMBnZvbmFnZTETMBEGA1UECwwKVm9uYWdlIEVuZzETMBEGA1UECgwKVm9uYWdlIEVuZzEQMA4GA1UEBwwHSG9sbWRlbDELMAkGA1UECAwCTkoxCzAJBgNVBAYTAlVTMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE77qasdW3CLativLGa%2FrsJBzVem0LGkSV7EtGDu5IoAAS18srDmq%2B8baIqoRJ7Kk41tCWfp8VybhIP3NJwEBwBKOCAUgwggFEMAwGA1UdEwEB%2FwQCMAAwHwYDVR0jBBgwFoAUO7lcyzEXnEjP1Npm6422qNXhCfcwgYEGCCsGAQUFBwEBBHUwczBHBggrBgEFBQcwAoY7aHR0cDovL2NhY2VydHMuY2NpZC5uZXVzdGFyL05ldXN0YXJDZXJ0aWZpZWRDYWxsZXJJZENBMS5jcnQwKAYIKwYBBQUHMAGGHGh0dHA6Ly9vY3NwLWNhMS5jY2lkLm5ldXN0YXIwSAYDVR0fBEEwPzA9oDugOYY3aHR0cDovL2NybC5jY2lkLm5ldXN0YXIvTmV1c3RhckNlcnRpZmllZENhbGxlcklkQ0ExLmNybDAdBgNVHQ4EFgQU%2FXpyIcR3tQS3RsQGKgGHqn8IP5gwDgYDVR0PAQH%2FBAQDAgeAMBYGCCsGAQUFBwEaBAowCKAGFgQxOTdEMA0GCSqGSIb3DQEBCwUAA4IBAQBnyM%2B7c78wj3K7iZtWZtCtkUcky3eHOv9FJRCtG%2BXLv8OPXEBdUXSJBXtT4X7XDLFyn5Vgjqolkb7Mia5iqnOUtPKmLqnXyrOTFbbFbXSCazZf6reg5BMOZxvJTEOESbn36AsT8FoDFvXmmW97leSQt%2F2PMFtYV6CyhUyqY227zt41QmGeW6E3U7CAT7XUH4g4lwTbq%2FK8T1VhQgO58QPjjhIlF0n6eE3Wk2F0SE7Vc79WRsEPfnt9apwfO9TI9fDFnKZ34ElNV1BFW74GQANGKCq4NfySQJ7nS2nYLlaUKioJKQ%2Bw5XftmNLSBGINlrCnz%2BuijZiVOTOaQGzq%2BnAu)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_pki_subject_rdn_unknown](../../ISSUES/w_pki_subject_rdn_unknown/README.md) | warn | SHAKEN PKI Best Practice | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


### Not Effective

- e_cp1_3_ambiguous_identifier
- e_cp1_3_subject_sn
- e_sti_authority_key_identifier
- e_sti_basic_constraints
- e_sti_certificate_policies
- e_sti_crl_distribution
- e_sti_crl_distribution_not_reachable
- e_sti_extension_unknown
- e_sti_issuer_dn
- e_sti_key_usage
- e_sti_serial_number
- e_sti_signature_algorithm
- e_sti_subject
- e_sti_subject_cn
- e_sti_subject_key_identifier
- e_sti_subject_public_key
- e_sti_tn_auth_list
- e_sti_version
- w_cp_1_3_subject_email

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 31/10/2022 at 16:43:22