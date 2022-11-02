# STIR/SHAKEN CA Ecosystem Compliance

## Certificate CenturyLink.com

Tested At: 02 Nov 22 17:23 UTC\
Initial Validity Period: 1096 day(s)\
Remaining Validity Period: 70 day(s)\
Subject: C=US, ST=CA, L=Monroe, O=Support, OU=Communications, CN=CenturyLink.com\
Issuer: C=US, O=Neustar Information Services Inc, OU=www.ccid.neustar, CN=Neustar Certified Caller ID CA-1\
Link: https://cr.ccid.neustar.biz/ccid/authn/v2/certs/11001.10012

[View certificate details](https://understandingwebpki.com/?cert=MIID%2BzCCAuOgAwIBAgIUGFko9jLiaggPzUHJ9QICtqumBG8wDQYJKoZIhvcNAQELBQAwfjEpMCcGA1UEAwwgTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIENBLTExGTAXBgNVBAsMEHd3dy5jY2lkLm5ldXN0YXIxKTAnBgNVBAoMIE5ldXN0YXIgSW5mb3JtYXRpb24gU2VydmljZXMgSW5jMQswCQYDVQQGEwJVUzAeFw0yMDAxMTAyMDA5MDNaFw0yMzAxMTAyMDA5MDNaMHAxGDAWBgNVBAMMD0NlbnR1cnlMaW5rLmNvbTEXMBUGA1UECwwOQ29tbXVuaWNhdGlvbnMxEDAOBgNVBAoMB1N1cHBvcnQxDzANBgNVBAcMBk1vbnJvZTELMAkGA1UECAwCQ0ExCzAJBgNVBAYTAlVTMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEEPiNlGRnSNvkauQ6SxPjLUpRMmgRxRtmIHbaBhCVdPXCEsxzeWBpgY7ND%2Bju4m7YB7PQ%2BGb6IOb%2FOxs0FUVfSaOCAUgwggFEMAwGA1UdEwEB%2FwQCMAAwHwYDVR0jBBgwFoAUO7lcyzEXnEjP1Npm6422qNXhCfcwgYEGCCsGAQUFBwEBBHUwczBHBggrBgEFBQcwAoY7aHR0cDovL2NhY2VydHMuY2NpZC5uZXVzdGFyL05ldXN0YXJDZXJ0aWZpZWRDYWxsZXJJZENBMS5jcnQwKAYIKwYBBQUHMAGGHGh0dHA6Ly9vY3NwLWNhMS5jY2lkLm5ldXN0YXIwSAYDVR0fBEEwPzA9oDugOYY3aHR0cDovL2NybC5jY2lkLm5ldXN0YXIvTmV1c3RhckNlcnRpZmllZENhbGxlcklkQ0ExLmNybDAdBgNVHQ4EFgQUAWBuLN%2Fb0nn97j1pD1KRe5Eul%2BAwDgYDVR0PAQH%2FBAQDAgeAMBYGCCsGAQUFBwEaBAowCKAGFgQ4ODI2MA0GCSqGSIb3DQEBCwUAA4IBAQAF%2BLdvjhXSawCgp3Xn82nr%2BEuOHv%2F2Qpgr2GjYJZiA83iNkYIZZ6C%2BEGu%2FIQvPQoAvXs7pJJzEuSGbniH5jRhgMqPZGz64uguwmH8u8wZjDTnhdW7dnlOrP3EMj9jJoncnyGPQrIRFGbGhDPWxsl8X1U8rvitUd%2BkOXRgUUdvKq7lEHFLLxoTTixbp%2FtT9dMLEU0UsjJa5bUeWalqpVR%2BkvEK1WvyD3AaK456wlMj1nM2XEND2jt24G%2BnOB9MjhZljQgl1yRTZtKviqUZrwq5hEUdHIYvGWXaiuNK7MZkkmb%2F7i5hirrFPLFJ5K4NLBnsJsYLvmPlG3CCGoXDYtJXo)

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


Generated: 02 Nov 22 17:25 UTC