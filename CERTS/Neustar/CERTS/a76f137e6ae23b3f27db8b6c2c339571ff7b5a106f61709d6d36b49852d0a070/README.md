# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN

Tested At: 06 Jul 23 13:53 UTC\
Initial Validity Period: 1096 day(s)\
Remaining Validity Period: 10 day(s)\
Subject: C=US, ST=NC, L=Raleigh, O=Bandwidth.com CLEC LLC, OU=Prod, CN=SHAKEN\
Issuer: C=US, O=Neustar Information Services Inc, OU=www.ccid.neustar, CN=Neustar Certified Caller ID CA-1\
Link: https://bw-shaken-cert-pub.s3.amazonaws.com/bandwidth-shaken-cert_20230716.pem

[View certificate details](https://understandingwebpki.com/?cert=MIID%2BDCCAuCgAwIBAgIUbz8Od22dmDT4D98MU%2BCx8rIqJggwDQYJKoZIhvcNAQELBQAwfjEpMCcGA1UEAwwgTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIENBLTExGTAXBgNVBAsMEHd3dy5jY2lkLm5ldXN0YXIxKTAnBgNVBAoMIE5ldXN0YXIgSW5mb3JtYXRpb24gU2VydmljZXMgSW5jMQswCQYDVQQGEwJVUzAeFw0yMDA3MTUwNDEzNDVaFw0yMzA3MTYwNDEzNDVaMG0xDzANBgNVBAMMBlNIQUtFTjENMAsGA1UECwwEUHJvZDEfMB0GA1UECgwWQmFuZHdpZHRoLmNvbSBDTEVDIExMQzEQMA4GA1UEBwwHUmFsZWlnaDELMAkGA1UECAwCTkMxCzAJBgNVBAYTAlVTMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAER%2F1z6%2FMsEqZLXSOOr8lKFQ4dyueuzbe1%2FjCqgz3obWnfV0ngLaR0bCec2SwTPEWUz0bL3Snv6JH0c9HxX38vRKOCAUgwggFEMBYGCCsGAQUFBwEaBAowCKAGFgQ5OTdFMAwGA1UdEwEB%2FwQCMAAwHwYDVR0jBBgwFoAUO7lcyzEXnEjP1Npm6422qNXhCfcwgYEGCCsGAQUFBwEBBHUwczBHBggrBgEFBQcwAoY7aHR0cDovL2NhY2VydHMuY2NpZC5uZXVzdGFyL05ldXN0YXJDZXJ0aWZpZWRDYWxsZXJJZENBMS5jcnQwKAYIKwYBBQUHMAGGHGh0dHA6Ly9vY3NwLWNhMS5jY2lkLm5ldXN0YXIwSAYDVR0fBEEwPzA9oDugOYY3aHR0cDovL2NybC5jY2lkLm5ldXN0YXIvTmV1c3RhckNlcnRpZmllZENhbGxlcklkQ0ExLmNybDAdBgNVHQ4EFgQU2pSEuxc%2BRzaou%2F8c1wqUX5sdeIwwDgYDVR0PAQH%2FBAQDAgeAMA0GCSqGSIb3DQEBCwUAA4IBAQCMsqvU8z9WI3RtrYNi9hBSV2wIidxXA7JfaIiTlFHW3JrW%2BGecm3L1oi5Vh5hGmVIym9nLGudMsuH9zyJRLDrxOw%2F54G9IXgtlQV4zQjczEt3TqfQySu%2BaLU0d%2Fti9OkNlslusCcLoof9bmIkHVj%2B58KjdIgNg1nzYnJlCJJcTh5Bi9s01ho%2BxIPFoKpNevo7L2N%2FXpqctWOIF%2Ffbju%2FvnoYQo9UZo%2BlKghoEtkOUaz0rUPP%2FT0PBjxffIgZBQKnEvAOJmK74XNVtG%2BHycz6NF0aXp0jEo97jCIa4fyC4tpQAZuAc2UlWd9ADdj0bSzOXnbCWA1hjOg%2BY%2B9vzcD2PI)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
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

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 06 Jul 23 14:08 UTC