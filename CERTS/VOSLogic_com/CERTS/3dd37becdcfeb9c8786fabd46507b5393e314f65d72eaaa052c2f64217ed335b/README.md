# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN

Tested At: 02 Jun 23 01:00 UTC\
Initial Validity Period: 825 day(s)\
Remaining Validity Period: 711 day(s)\
Subject: CN=SHAKEN, OU=VOIP, O=VOSLogic.com, ST=FL, C=US\
Issuer: emailAddress=support@voslogic.com, CN=*.voslogic.com, O=VOSLogic.com, ST=FL, C=US, emailAddress=support@voslogic.com\
Link: https://cdn.cnxcdn.com/shaken/59.crt

[View certificate details](https://understandingwebpki.com/?cert=MIIBuzCCAWECFEgQlvaYq3zPfx%2B%2B1fkcSzQxYNN%2BMAoGCCqGSM49BAMCMG8xCzAJBgNVBAYTAlVTMQswCQYDVQQIDAJGTDEVMBMGA1UECgwMVk9TTG9naWMuY29tMRcwFQYDVQQDDA4qLnZvc2xvZ2ljLmNvbTEjMCEGCSqGSIb3DQEJARYUc3VwcG9ydEB2b3Nsb2dpYy5jb20wHhcNMjMwMjA4MDAwMTI1WhcNMjUwNTEzMDAwMTI1WjBRMQswCQYDVQQGEwJVUzELMAkGA1UECAwCRkwxFTATBgNVBAoMDFZPU0xvZ2ljLmNvbTENMAsGA1UECwwEVk9JUDEPMA0GA1UEAwwGU0hBS0VOMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAENqxcGaPrejadvh1ndA9Dss%2BvZ6qGY%2FMUzI%2Fof0oln3cvdEmKK1LuybccepQ7dgNFO%2FFD9G2UVYjW8W5WzzgRmzAKBggqhkjOPQQDAgNIADBFAiAavZ2ZtBQcH%2BvQTka%2BgqfQZY2Nf3KsVcqc3uPG9iuiwAIhAKC5IRdGgBmLVp%2FHI44rZ45YdcXKCWUxwHbCsPRj%2FDEp)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_basic_constraints](../../ISSUES/e_atis_basic_constraints/README.md) | error | ATIS1000080 | STI certificates shall contain a BasicConstraints extension marked critical |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |
| [e_atis_authority_key_identifier](../../ISSUES/e_atis_authority_key_identifier/README.md) | error | ATIS1000080 | STI certificates shall contain an Authority Key Identifier extension |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | STI certificate shall contain TNAuthorizationList extension |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Cannot get SPC value from the TNAuthList extension, STI certificate shall contain TNAuthorizationList extension |
| [e_atis_version](../../ISSUES/e_atis_version/README.md) | error | ATIS1000080 | STI certificates shall contain Version field specifying version 3 |
| [e_atis_crl_distribution](../../ISSUES/e_atis_crl_distribution/README.md) | error | ATIS1000080 | STI End-Entity certificates shall contain a CRL Distribution Points extension |
| [w_ext_subject_key_identifier_missing_sub_cert](../../ISSUES/w_ext_subject_key_identifier_missing_sub_cert/README.md) | warn | RFC5280 |  |
| [e_atis_tn_auth_list](../../ISSUES/e_atis_tn_auth_list/README.md) | error | ATIS1000080 | STI certificate shall contain TNAuthorizationList extension |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_ext_authority_key_identifier_missing](../../ISSUES/e_ext_authority_key_identifier_missing/README.md) | error | RFC5280 |  |
| [e_ext_authority_key_identifier_no_key_identifier](../../ISSUES/e_ext_authority_key_identifier_no_key_identifier/README.md) | error | RFC5280 |  |
| [e_atis_subject_key_identifier](../../ISSUES/e_atis_subject_key_identifier/README.md) | error | ATIS1000080 | STI certificates shall contain a Subject Key Identifier extension |
| [e_atis_key_usage](../../ISSUES/e_atis_key_usage/README.md) | error | ATIS1000080 | STI certificates shall contain a Key Usage extension marked as critical |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 02 Jun 23 01:12 UTC