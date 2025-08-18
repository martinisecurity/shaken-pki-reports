# STIR/SHAKEN CA Ecosystem Compliance

## Certificate www.voicecarrier.com

Tested At: 18 Aug 25 21:10 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 126 day(s)\
Subject: CN=www.voicecarrier.com, OU=voip, O=VoiceCarrier, L=TX, ST=TX, C=US\
Issuer: CN=www.voicecarrier.com, OU=voip, O=VoiceCarrier, L=TX, ST=TX, C=US\
Link: https://www.voicecarrier.com/vc_shaken.pem

[View certificate details](https://x509.io/?cert=MIIByDCCAW4CCQCzveOrMZFtlDAKBggqhkjOPQQDAjBsMQswCQYDVQQGEwJVUzELMAkGA1UECAwCVFgxCzAJBgNVBAcMAlRYMRUwEwYDVQQKDAxWb2ljZUNhcnJpZXIxDTALBgNVBAsMBHZvaXAxHTAbBgNVBAMMFHd3dy52b2ljZWNhcnJpZXIuY29tMB4XDTI0MTIyMjE1MjAxN1oXDTI1MTIyMjE1MjAxN1owbDELMAkGA1UEBhMCVVMxCzAJBgNVBAgMAlRYMQswCQYDVQQHDAJUWDEVMBMGA1UECgwMVm9pY2VDYXJyaWVyMQ0wCwYDVQQLDAR2b2lwMR0wGwYDVQQDDBR3d3cudm9pY2VjYXJyaWVyLmNvbTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABM73O8t0DbujgZ3ZD3WgooJbnnCK%2F5QBBCsn19m%2BiBYyYElIXzMG1pQouUu2SbY3LVoFxCbwG3SKy0IUGbhX33UwCgYIKoZIzj0EAwIDSAAwRQIhAOQ%2FCjHOj23zcLxF2Ql%2BcBE7zz3RTFzMj%2BO%2BFZmAUMvyAiB2%2FQRRI0q1kW4gyiSGtuG9SU5l9n%2FwE2OWb4gpWZ6sxw%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_ext_authority_key_identifier_no_key_identifier](../../ISSUES/e_ext_authority_key_identifier_no_key_identifier/README.md) | error | RFC5280 |  |
| [e_atis_ext_authority_key_identifier](../../ISSUES/e_atis_ext_authority_key_identifier/README.md) | error | ATIS1000080 | STI certificates shall contain an Authority Key Identifier extension |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common Name attribute 'www.voicecarrier.com' does not contain 'SHAKEN' |
| [e_atis_version](../../ISSUES/e_atis_version/README.md) | error | ATIS1000080 | STI certificates shall contain Version field specifying version 3 |
| [e_atis_ext_basic_constraints](../../ISSUES/e_atis_ext_basic_constraints/README.md) | error | ATIS1000080 | BasicConstraints extension not found |
| [e_atis_subject_key_identifier](../../ISSUES/e_atis_subject_key_identifier/README.md) | error | ATIS1000080 | Subject Key Identifier extension not found |
| [e_atis_ext_key_usage](../../ISSUES/e_atis_ext_key_usage/README.md) | error | ATIS1000080 | Key Usage extension not found |
| [e_atis_ext_crl_distribution](../../ISSUES/e_atis_ext_crl_distribution/README.md) | error | ATIS1000080 | STI End-Entity certificates shall contain a CRL Distribution Points extension |
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension does not contain a single OID value that identifies the SHAKEN Certificate Policy established by the STI-PA |
| [e_atis_tn_auth_list](../../ISSUES/e_atis_tn_auth_list/README.md) | error | ATIS1000080 | the TNAuthList extension is not present |
| [e_atis_ext_certificate_policies](../../ISSUES/e_atis_ext_certificate_policies/README.md) | error | ATIS1000080 | The Certificate Policies extension is not present |
| [w_ext_subject_key_identifier_missing_sub_cert](../../ISSUES/w_ext_subject_key_identifier_missing_sub_cert/README.md) | warn | RFC5280 |  |
| [e_atis_tn_auth_list_spc_format](../../ISSUES/e_atis_tn_auth_list_spc_format/README.md) | error | ATIS1000080 | STI certificate shall contain TNAuthorizationList extension |
| [e_atis_subject_key_identifier_size](../../ISSUES/e_atis_subject_key_identifier_size/README.md) | error | ATIS1000080 | Subject Key Identifier extension not found |


Generated: 18 Aug 25 21:13 UTC