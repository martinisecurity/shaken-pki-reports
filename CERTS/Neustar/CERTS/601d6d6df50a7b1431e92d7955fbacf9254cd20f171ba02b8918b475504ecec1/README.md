# STIR/SHAKEN CA Ecosystem Compliance

## Certificate AGOC

Tested At: 02 Nov 22 21:16 UTC\
Initial Validity Period: 1096 day(s)\
Remaining Validity Period: 215 day(s)\
Subject: C=US, ST=PA, L=BUTLER, O=AGOC, OU=AGOC, CN=AGOC\
Issuer: C=US, O=Neustar Information Services Inc, OU=www.ccid.neustar, CN=Neustar Certified Caller ID CA-1\
Link: https://cr.ccid.neustar.biz/ccid/authn/v2/certs/11114.10192.pem

[View certificate details](https://understandingwebpki.com/?cert=MIID4zCCAsugAwIBAgIUL7v48xo8jl4j8c2Dqbd5JgHakLMwDQYJKoZIhvcNAQELBQAwfjEpMCcGA1UEAwwgTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIENBLTExGTAXBgNVBAsMEHd3dy5jY2lkLm5ldXN0YXIxKTAnBgNVBAoMIE5ldXN0YXIgSW5mb3JtYXRpb24gU2VydmljZXMgSW5jMQswCQYDVQQGEwJVUzAeFw0yMDA2MDQxODM5MzhaFw0yMzA2MDUxODM5MzhaMFgxDTALBgNVBAMMBEFHT0MxDTALBgNVBAsMBEFHT0MxDTALBgNVBAoMBEFHT0MxDzANBgNVBAcMBkJVVExFUjELMAkGA1UECAwCUEExCzAJBgNVBAYTAlVTMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEFQSM3N5FS8lwBiRuRh2wvOXPM%2BMNO%2BRSJWaM9TWhIByerL2FCrpnxXOLU3zTLgKS1WFifLCYUGTq5q%2Fv1rmHMKOCAUgwggFEMAwGA1UdEwEB%2FwQCMAAwHwYDVR0jBBgwFoAUO7lcyzEXnEjP1Npm6422qNXhCfcwgYEGCCsGAQUFBwEBBHUwczBHBggrBgEFBQcwAoY7aHR0cDovL2NhY2VydHMuY2NpZC5uZXVzdGFyL05ldXN0YXJDZXJ0aWZpZWRDYWxsZXJJZENBMS5jcnQwKAYIKwYBBQUHMAGGHGh0dHA6Ly9vY3NwLWNhMS5jY2lkLm5ldXN0YXIwSAYDVR0fBEEwPzA9oDugOYY3aHR0cDovL2NybC5jY2lkLm5ldXN0YXIvTmV1c3RhckNlcnRpZmllZENhbGxlcklkQ0ExLmNybDAdBgNVHQ4EFgQU5lY2qOAIvJ%2B1fkunIsIoVTHvL3swDgYDVR0PAQH%2FBAQDAgeAMBYGCCsGAQUFBwEaBAowCKAGFgQ0MThDMA0GCSqGSIb3DQEBCwUAA4IBAQCHDmf3TM3LhjDPnD5l44UWejx0NcEfvIx5XqMMnf5kqFwP1%2BDPxbJHawr%2BTPzMoMDuJ5%2FYyxdy9iHapGgtbME6RRdpdC2sTFeMQTkkf0gIjqhI7VVDUnviZ%2FFd81p4lduthCKQnSufS7ezQ15i0O4BLgJPxzOFEuO%2Fz7gIGXuh5U30cImUMp4vFKvjmv0EC25bbIPyRq1ehdj%2BRNYTTDLDaJK%2Br4INK62P2e8WemnPkklT%2B3bbh9nxkmDVF7k1kGbgHYbTKU1TEQhnQemt%2BGXoORZT6tGvEn6c9GgNfEKpU0IMHy5ET3Vh8T69dPjvhub6LzVElQCo7M0IDxSFAG7j)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |

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
- e_us_cp_subject_sn
- w_cp_1_3_subject_email

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 02 Nov 22 21:24 UTC