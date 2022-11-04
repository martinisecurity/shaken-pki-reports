# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN

Tested At: 04 Nov 22 01:09 UTC\
Initial Validity Period: 1096 day(s)\
Remaining Validity Period: 326 day(s)\
Subject: C=US, ST=CA, L=Los Alamitos, O=Piratel LLC, OU=VOIP, CN=SHAKEN\
Issuer: C=US, O=Neustar Information Services Inc, OU=www.ccid.neustar, CN=Neustar Certified Caller ID CA-1\
Link: https://certificates.piratel.com/piratel_20201008_49b10950e1.crt

[View certificate details](https://understandingwebpki.com/?cert=MIID8jCCAtqgAwIBAgIUbLcjKmkavulRrrZCLT%2F7WelCxwEwDQYJKoZIhvcNAQELBQAwfjEpMCcGA1UEAwwgTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIENBLTExGTAXBgNVBAsMEHd3dy5jY2lkLm5ldXN0YXIxKTAnBgNVBAoMIE5ldXN0YXIgSW5mb3JtYXRpb24gU2VydmljZXMgSW5jMQswCQYDVQQGEwJVUzAeFw0yMDA5MjUwMDI1NTFaFw0yMzA5MjYwMDI1NTFaMGcxDzANBgNVBAMMBlNIQUtFTjENMAsGA1UECwwEVk9JUDEUMBIGA1UECgwLUGlyYXRlbCBMTEMxFTATBgNVBAcMDExvcyBBbGFtaXRvczELMAkGA1UECAwCQ0ExCzAJBgNVBAYTAlVTMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEEWKCeNgdBvRN9xfcB9HdrbBvU9GgKPw0KWSU9q%2BEA0ImeF9IwG6rnw6obsgEQCBRgAM19%2B95Msn7MKbCwy7hOaOCAUgwggFEMBYGCCsGAQUFBwEaBAowCKAGFgQ1MDJKMAwGA1UdEwEB%2FwQCMAAwHwYDVR0jBBgwFoAUO7lcyzEXnEjP1Npm6422qNXhCfcwgYEGCCsGAQUFBwEBBHUwczBHBggrBgEFBQcwAoY7aHR0cDovL2NhY2VydHMuY2NpZC5uZXVzdGFyL05ldXN0YXJDZXJ0aWZpZWRDYWxsZXJJZENBMS5jcnQwKAYIKwYBBQUHMAGGHGh0dHA6Ly9vY3NwLWNhMS5jY2lkLm5ldXN0YXIwSAYDVR0fBEEwPzA9oDugOYY3aHR0cDovL2NybC5jY2lkLm5ldXN0YXIvTmV1c3RhckNlcnRpZmllZENhbGxlcklkQ0ExLmNybDAdBgNVHQ4EFgQUPuWWZZbf0ymMR9uqdrIzZSktaEUwDgYDVR0PAQH%2FBAQDAgeAMA0GCSqGSIb3DQEBCwUAA4IBAQA601UKfJz%2B53pNNDH2ViKbp1cASroAVpPbtxsOpThMBA80CI6P7LJMDACt8n46O8Imhaod5w7VKjndywKmEe34if3X6h3LOin%2Frr2M8USCHLp%2B4jskfg5qCx%2BlO56IRfMO1Zsh1Uyq3sERKVwWHVkqwGwasCHEzoL4B5WxvRphaJlyirK6p2WSbamsxGc7VCcAdIozv0TM5IHGSzhdheH2mRHB2RHZQHtcWJPLqc%2FhSMWMaNYMw5UnmgU0TSYSpiMyxPevVoECOb7UNEq6c%2BlrfFnz0%2F3u7xTn1XqMy%2Bop0kcxphAtOwu70OE6u6pn3aHjS8MpuNN6aK9XGF6eAZ5v)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |

### Not Effective

- e_atis_authority_key_identifier
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


Generated: 04 Nov 22 01:11 UTC