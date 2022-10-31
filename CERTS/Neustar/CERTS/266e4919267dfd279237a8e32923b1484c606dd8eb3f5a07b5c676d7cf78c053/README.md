# STIR/SHAKEN CA Ecosystem Compliance

## Certificate CenturyLink.com

Tested At: 31 Oct 22 20:24 UTC\
Initial Validity Period: 1096 day(s)\
Remaining Validity Period: 71 day(s)\
Subject: C=US, ST=CA, L=Monroe, O=Support, OU=Communications, CN=CenturyLink.com\
Issuer: C=US, O=Neustar Information Services Inc, OU=www.ccid.neustar, CN=Neustar Certified Caller ID CA-1\
Link: https://cr.ccid.neustar.biz/ccid/authn/v2/certs/11001.10012

View: [Click to view](https://understandingwebpki.com/?cert=MIID%2BzCCAuOgAwIBAgIUGFko9jLiaggPzUHJ9QICtqumBG8wDQYJKoZIhvcNAQELBQAwfjEpMCcGA1UEAwwgTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIENBLTExGTAXBgNVBAsMEHd3dy5jY2lkLm5ldXN0YXIxKTAnBgNVBAoMIE5ldXN0YXIgSW5mb3JtYXRpb24gU2VydmljZXMgSW5jMQswCQYDVQQGEwJVUzAeFw0yMDAxMTAyMDA5MDNaFw0yMzAxMTAyMDA5MDNaMHAxGDAWBgNVBAMMD0NlbnR1cnlMaW5rLmNvbTEXMBUGA1UECwwOQ29tbXVuaWNhdGlvbnMxEDAOBgNVBAoMB1N1cHBvcnQxDzANBgNVBAcMBk1vbnJvZTELMAkGA1UECAwCQ0ExCzAJBgNVBAYTAlVTMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEEPiNlGRnSNvkauQ6SxPjLUpRMmgRxRtmIHbaBhCVdPXCEsxzeWBpgY7ND%2Bju4m7YB7PQ%2BGb6IOb%2FOxs0FUVfSaOCAUgwggFEMAwGA1UdEwEB%2FwQCMAAwHwYDVR0jBBgwFoAUO7lcyzEXnEjP1Npm6422qNXhCfcwgYEGCCsGAQUFBwEBBHUwczBHBggrBgEFBQcwAoY7aHR0cDovL2NhY2VydHMuY2NpZC5uZXVzdGFyL05ldXN0YXJDZXJ0aWZpZWRDYWxsZXJJZENBMS5jcnQwKAYIKwYBBQUHMAGGHGh0dHA6Ly9vY3NwLWNhMS5jY2lkLm5ldXN0YXIwSAYDVR0fBEEwPzA9oDugOYY3aHR0cDovL2NybC5jY2lkLm5ldXN0YXIvTmV1c3RhckNlcnRpZmllZENhbGxlcklkQ0ExLmNybDAdBgNVHQ4EFgQUAWBuLN%2Fb0nn97j1pD1KRe5Eul%2BAwDgYDVR0PAQH%2FBAQDAgeAMBYGCCsGAQUFBwEaBAowCKAGFgQ4ODI2MA0GCSqGSIb3DQEBCwUAA4IBAQAF%2BLdvjhXSawCgp3Xn82nr%2BEuOHv%2F2Qpgr2GjYJZiA83iNkYIZZ6C%2BEGu%2FIQvPQoAvXs7pJJzEuSGbniH5jRhgMqPZGz64uguwmH8u8wZjDTnhdW7dnlOrP3EMj9jJoncnyGPQrIRFGbGhDPWxsl8X1U8rvitUd%2BkOXRgUUdvKq7lEHFLLxoTTixbp%2FtT9dMLEU0UsjJa5bUeWalqpVR%2BkvEK1WvyD3AaK456wlMj1nM2XEND2jt24G%2BnOB9MjhZljQgl1yRTZtKviqUZrwq5hEUdHIYvGWXaiuNK7MZkkmb%2F7i5hirrFPLFJ5K4NLBnsJsYLvmPlG3CCGoXDYtJXo)


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


Generated: 31/10/2022 at 20:32:42