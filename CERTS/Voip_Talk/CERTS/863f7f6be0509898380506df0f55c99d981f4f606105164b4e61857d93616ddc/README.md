# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN

Tested At: 15 Nov 23 15:51 UTC\
Initial Validity Period: 825 day(s)\
Remaining Validity Period: 751 day(s)\
Subject: CN=SHAKEN, OU=VOIP, O=AcmeTelecom\\, Inc., L=Somewhere, ST=VA, C=US\
Issuer: CN=voiptalk, OU=Voice, O=Voip Talk, L=Sheridan, ST=WY, C=US\
Link: https://cdn.cnxcdn.com/shaken/d20f2bf9ad.crt

[View certificate details](https://understandingwebpki.com/?cert=MIIB6jCCAZCgAwIBAgIUVQGy6uR5TC3odP02ea21UWLxt8IwCgYIKoZIzj0EAwIwZDELMAkGA1UEBhMCVVMxCzAJBgNVBAgMAldZMREwDwYDVQQHDAhTaGVyaWRhbjESMBAGA1UECgwJVm9pcCBUYWxrMQ4wDAYDVQQLDAVWb2ljZTERMA8GA1UEAwwIdm9pcHRhbGswHhcNMjMwOTAyMDcyODQ5WhcNMjUxMjA1MDcyODQ5WjBqMQswCQYDVQQGEwJVUzELMAkGA1UECAwCVkExEjAQBgNVBAcMCVNvbWV3aGVyZTEaMBgGA1UECgwRQWNtZVRlbGVjb20sIEluYy4xDTALBgNVBAsMBFZPSVAxDzANBgNVBAMMBlNIQUtFTjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABLgoB2l1eo%2B98fdgBpWWrqhTcb663lNTEh8thLh%2FB%2FeOxi%2Fc8QNFrU6%2FMiUIZvPzg79m4fsXIHZPySiGQqOiZTajGjAYMBYGCCsGAQUFBwEaBAowCKAGFgQxMDAxMAoGCCqGSM49BAMCA0gAMEUCIFqqayjTvpGUfDzlVEm6Dx3st9c50PBjeBidx2sFmwyXAiEA%2BjPHibj1PV4toDo%2BL9HYfqQcA9dAUBtzlb9bknW8S6A%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_ext_subject_key_identifier_missing_sub_cert](../../ISSUES/w_ext_subject_key_identifier_missing_sub_cert/README.md) | warn | RFC5280 |  |
| [e_atis_key_usage](../../ISSUES/e_atis_key_usage/README.md) | error | ATIS1000080 | STI certificates shall contain a Key Usage extension marked as critical |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |
| [e_ext_authority_key_identifier_no_key_identifier](../../ISSUES/e_ext_authority_key_identifier_no_key_identifier/README.md) | error | RFC5280 |  |
| [e_atis_authority_key_identifier](../../ISSUES/e_atis_authority_key_identifier/README.md) | error | ATIS1000080 | STI certificates shall contain an Authority Key Identifier extension |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_atis_basic_constraints](../../ISSUES/e_atis_basic_constraints/README.md) | error | ATIS1000080 | STI certificates shall contain a BasicConstraints extension marked critical |
| [e_ext_authority_key_identifier_missing](../../ISSUES/e_ext_authority_key_identifier_missing/README.md) | error | RFC5280 |  |
| [e_atis_crl_distribution](../../ISSUES/e_atis_crl_distribution/README.md) | error | ATIS1000080 | STI End-Entity certificates shall contain a CRL Distribution Points extension |
| [e_atis_subject_key_identifier](../../ISSUES/e_atis_subject_key_identifier/README.md) | error | ATIS1000080 | STI certificates shall contain a Subject Key Identifier extension |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 1001' |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |


Generated: 15 Nov 23 16:51 UTC