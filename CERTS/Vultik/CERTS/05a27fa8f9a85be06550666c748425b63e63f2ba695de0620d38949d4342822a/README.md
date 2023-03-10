# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN

Tested At: 09 Mar 23 22:57 UTC\
Initial Validity Period: 825 day(s)\
Remaining Validity Period: 787 day(s)\
Subject: CN=SHAKEN, OU=CIA, O=Vultik, L=DesMoines, ST=IA, C=US\
Issuer: emailAddress=support@vultik.com, CN=*.vultik.com, O=Vultik, L=Des Moines, ST=IA, C=US, emailAddress=support@vultik.com\
Link: https://cdn.cnxcdn.com/shaken/56.crt

[View certificate details](https://understandingwebpki.com/?cert=MIIB1DCCAXkCFAq%2B20A6z25bSKqWpAsjsnuco0nbMAoGCCqGSM49BAMCMHoxCzAJBgNVBAYTAlVTMQswCQYDVQQIDAJJQTETMBEGA1UEBwwKRGVzIE1vaW5lczEPMA0GA1UECgwGVnVsdGlrMRUwEwYDVQQDDAwqLnZ1bHRpay5jb20xITAfBgkqhkiG9w0BCQEWEnN1cHBvcnRAdnVsdGlrLmNvbTAeFw0yMzAxMzAwNTU2MDdaFw0yNTA1MDQwNTU2MDdaMF4xCzAJBgNVBAYTAlVTMQswCQYDVQQIDAJJQTESMBAGA1UEBwwJRGVzTW9pbmVzMQ8wDQYDVQQKDAZWdWx0aWsxDDAKBgNVBAsMA0NJQTEPMA0GA1UEAwwGU0hBS0VOMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEGINFI3MykzZykSGNdimr8%2Fj7t881iWZxi66uag4WKCqSG4oeM%2B%2F8rlwiGqDWJnNgwBEmab6XXPBXwcSp7eSdRzAKBggqhkjOPQQDAgNJADBGAiEAxk8FCiWxuZL7LdCg8wL2cYNH9IpjQ6Eo%2FUWD2uxrnmgCIQD9kaC6K8Wmplw70FrDShOp9vyxzMRQ0KBahOB1dlcsLg%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_crl_distribution](../../ISSUES/e_atis_crl_distribution/README.md) | error | ATIS1000080 | STI End-Entity certificates shall contain a CRL Distribution Points extension |
| [e_atis_basic_constraints](../../ISSUES/e_atis_basic_constraints/README.md) | error | ATIS1000080 | STI certificates shall contain a BasicConstraints extension marked critical |
| [w_ext_subject_key_identifier_missing_sub_cert](../../ISSUES/w_ext_subject_key_identifier_missing_sub_cert/README.md) | warn | RFC5280 |  |
| [e_atis_authority_key_identifier](../../ISSUES/e_atis_authority_key_identifier/README.md) | error | ATIS1000080 | STI certificates shall contain an Authority Key Identifier extension |
| [e_atis_version](../../ISSUES/e_atis_version/README.md) | error | ATIS1000080 | STI certificates shall contain Version field specifying version 3 |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Cannot get SPC value from the TNAuthList extension, STI certificate shall contain TNAuthorizationList extension |
| [e_atis_key_usage](../../ISSUES/e_atis_key_usage/README.md) | error | ATIS1000080 | STI certificates shall contain a Key Usage extension marked as critical |
| [e_ext_authority_key_identifier_missing](../../ISSUES/e_ext_authority_key_identifier_missing/README.md) | error | RFC5280 |  |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | STI certificate shall contain TNAuthorizationList extension |
| [e_atis_tn_auth_list](../../ISSUES/e_atis_tn_auth_list/README.md) | error | ATIS1000080 | STI certificate shall contain TNAuthorizationList extension |
| [e_atis_subject_key_identifier](../../ISSUES/e_atis_subject_key_identifier/README.md) | error | ATIS1000080 | STI certificates shall contain a Subject Key Identifier extension |
| [e_ext_authority_key_identifier_no_key_identifier](../../ISSUES/e_ext_authority_key_identifier_no_key_identifier/README.md) | error | RFC5280 |  |


Generated: 10 Mar 23 02:25 UTC