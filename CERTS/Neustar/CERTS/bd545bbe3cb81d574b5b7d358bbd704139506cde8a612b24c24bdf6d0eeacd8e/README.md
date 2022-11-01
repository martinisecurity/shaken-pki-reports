# STIR/SHAKEN CA Ecosystem Compliance

## Certificate frontier.com

Tested At: 01 Nov 22 09:57 UTC\
Initial Validity Period: 1096 day(s)\
Remaining Validity Period: 90 day(s)\
Subject: C=US, ST=NY, L=Rochester, O=Support, OU=Communications, CN=frontier.com\
Issuer: C=US, O=Neustar Information Services Inc, OU=www.ccid.neustar, CN=Neustar Certified Caller ID CA-1\
Link: https://cr.ccid.neustar.biz/ccid/authn/v2/certs/11022.10175

View: [Click to view](https://understandingwebpki.com/?cert=MIID%2BzCCAuOgAwIBAgIUKhPu%2B%2BjycK2FLWt6BtgXjlVGtjAwDQYJKoZIhvcNAQELBQAwfjEpMCcGA1UEAwwgTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIENBLTExGTAXBgNVBAsMEHd3dy5jY2lkLm5ldXN0YXIxKTAnBgNVBAoMIE5ldXN0YXIgSW5mb3JtYXRpb24gU2VydmljZXMgSW5jMQswCQYDVQQGEwJVUzAeFw0yMDAxMjkxNTE2NTBaFw0yMzAxMjkxNTE2NTBaMHAxFTATBgNVBAMMDGZyb250aWVyLmNvbTEXMBUGA1UECwwOQ29tbXVuaWNhdGlvbnMxEDAOBgNVBAoMB1N1cHBvcnQxEjAQBgNVBAcMCVJvY2hlc3RlcjELMAkGA1UECAwCTlkxCzAJBgNVBAYTAlVTMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEWzfYktQY%2BDbXluTaOG0ABn3cAYhgF2wjjz5Cx09x9Nmio0s1oXsOrfbQi%2Bbkss0OmL0VCdQY3xRF6qkgL97rS6OCAUgwggFEMAwGA1UdEwEB%2FwQCMAAwHwYDVR0jBBgwFoAUO7lcyzEXnEjP1Npm6422qNXhCfcwgYEGCCsGAQUFBwEBBHUwczBHBggrBgEFBQcwAoY7aHR0cDovL2NhY2VydHMuY2NpZC5uZXVzdGFyL05ldXN0YXJDZXJ0aWZpZWRDYWxsZXJJZENBMS5jcnQwKAYIKwYBBQUHMAGGHGh0dHA6Ly9vY3NwLWNhMS5jY2lkLm5ldXN0YXIwSAYDVR0fBEEwPzA9oDugOYY3aHR0cDovL2NybC5jY2lkLm5ldXN0YXIvTmV1c3RhckNlcnRpZmllZENhbGxlcklkQ0ExLmNybDAdBgNVHQ4EFgQUbhUqnXCfxsaiaxca6dS8gAllsoYwDgYDVR0PAQH%2FBAQDAgeAMBYGCCsGAQUFBwEaBAowCKAGFgQwMDZFMA0GCSqGSIb3DQEBCwUAA4IBAQBfmiaaw8Q%2F2F1bmBW4HYyM9%2BlGB7RgGsRmyIh9ukqj3OBhHvGrsV%2Bs7yP%2F%2FBqfVc1eq9T%2BJvoUr%2BWa%2F851ir8mlOk8fD4rXpB3XbmiS1RkKM8eYbEIqLakhgOYCy4YdG90ELgMVccf667dmiGQExJ%2BQ6L5hSRAlAmIzV4jBBA3GjSYCXNlAC1q8wTQVBJhjBEqQTh5B%2FqY7%2FHxpL9j6c16QMwiMoQJfPprd80NIu1CLozyfxhH6nHH04G5LQiu8jYeERPjoSHF%2FOFEwQG6xlE6nnpEMEDOZJsFJTyqNwUD6QYhjXaxEW2Pe09csXHAGKjo8kMUVXZ57Ps5Q7XHH2%2BP)


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


Generated: 01/11/2022 at 10:05:32