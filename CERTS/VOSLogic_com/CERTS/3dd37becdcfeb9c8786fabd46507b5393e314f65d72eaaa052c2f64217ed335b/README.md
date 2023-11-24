# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN

Tested At: 24 Nov 23 11:00 UTC\
Initial Validity Period: 825 day(s)\
Remaining Validity Period: 536 day(s)\
Subject: CN=SHAKEN, OU=VOIP, O=VOSLogic.com, ST=FL, C=US\
Issuer: emailAddress=support@voslogic.com, CN=*.voslogic.com, O=VOSLogic.com, ST=FL, C=US, emailAddress=support@voslogic.com\
Link: https://cdn.cnxcdn.com/shaken/59.crt

[View certificate details](https://understandingwebpki.com/?cert=MIIBuzCCAWECFEgQlvaYq3zPfx%2B%2B1fkcSzQxYNN%2BMAoGCCqGSM49BAMCMG8xCzAJBgNVBAYTAlVTMQswCQYDVQQIDAJGTDEVMBMGA1UECgwMVk9TTG9naWMuY29tMRcwFQYDVQQDDA4qLnZvc2xvZ2ljLmNvbTEjMCEGCSqGSIb3DQEJARYUc3VwcG9ydEB2b3Nsb2dpYy5jb20wHhcNMjMwMjA4MDAwMTI1WhcNMjUwNTEzMDAwMTI1WjBRMQswCQYDVQQGEwJVUzELMAkGA1UECAwCRkwxFTATBgNVBAoMDFZPU0xvZ2ljLmNvbTENMAsGA1UECwwEVk9JUDEPMA0GA1UEAwwGU0hBS0VOMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAENqxcGaPrejadvh1ndA9Dss%2BvZ6qGY%2FMUzI%2Fof0oln3cvdEmKK1LuybccepQ7dgNFO%2FFD9G2UVYjW8W5WzzgRmzAKBggqhkjOPQQDAgNIADBFAiAavZ2ZtBQcH%2BvQTka%2BgqfQZY2Nf3KsVcqc3uPG9iuiwAIhAKC5IRdGgBmLVp%2FHI44rZ45YdcXKCWUxwHbCsPRj%2FDEp)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_ext_authority_key_identifier](../../ISSUES/e_atis_ext_authority_key_identifier/README.md) | error | ATIS1000080 | STI certificates shall contain an Authority Key Identifier extension |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Cannot get SPC value from the TNAuthList extension, STI certificate shall contain TNAuthorizationList extension |
| [w_ext_subject_key_identifier_missing_sub_cert](../../ISSUES/w_ext_subject_key_identifier_missing_sub_cert/README.md) | warn | RFC5280 |  |
| [e_atis_subject_key_identifier](../../ISSUES/e_atis_subject_key_identifier/README.md) | error | ATIS1000080 | Subject Key Identifier extension not found |
| [e_atis_ext_basic_constraints](../../ISSUES/e_atis_ext_basic_constraints/README.md) | error | ATIS1000080 | BasicConstraints extension not found |
| [e_atis_subject_key_identifier_size](../../ISSUES/e_atis_subject_key_identifier_size/README.md) | error | ATIS1000080 | Subject Key Identifier extension not found |
| [e_atis_ext_certificate_policies](../../ISSUES/e_atis_ext_certificate_policies/README.md) | error | ATIS1000080 | The Certificate Policies extension is not present |
| [e_atis_ext_crl_distribution](../../ISSUES/e_atis_ext_crl_distribution/README.md) | error | ATIS1000080 | STI End-Entity certificates shall contain a CRL Distribution Points extension |
| [e_atis_ext_key_usage](../../ISSUES/e_atis_ext_key_usage/README.md) | error | ATIS1000080 | Key Usage extension not found |
| [e_ext_authority_key_identifier_missing](../../ISSUES/e_ext_authority_key_identifier_missing/README.md) | error | RFC5280 |  |
| [e_atis_version](../../ISSUES/e_atis_version/README.md) | error | ATIS1000080 | STI certificates shall contain Version field specifying version 3 |
| [e_atis_tn_auth_list](../../ISSUES/e_atis_tn_auth_list/README.md) | error | ATIS1000080 | the TNAuthList extension is not present |
| [e_ext_authority_key_identifier_no_key_identifier](../../ISSUES/e_ext_authority_key_identifier_no_key_identifier/README.md) | error | RFC5280 |  |
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension does not contain a single OID value that identifies the SHAKEN Certificate Policy established by the STI-PA |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 24 Nov 23 11:17 UTC