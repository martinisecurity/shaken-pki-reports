# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN

Tested At: 24 Nov 23 11:00 UTC\
Initial Validity Period: 825 day(s)\
Remaining Validity Period: 742 day(s)\
Subject: CN=SHAKEN, OU=VOIP, O=AcmeTelecom\\, Inc., L=Somewhere, ST=VA, C=US\
Issuer: CN=voiptalk, OU=Voice, O=Voip Talk, L=Sheridan, ST=WY, C=US\
Link: https://cdn.cnxcdn.com/shaken/d20f2bf9ad.crt

[View certificate details](https://understandingwebpki.com/?cert=MIIB6jCCAZCgAwIBAgIUVQGy6uR5TC3odP02ea21UWLxt8IwCgYIKoZIzj0EAwIwZDELMAkGA1UEBhMCVVMxCzAJBgNVBAgMAldZMREwDwYDVQQHDAhTaGVyaWRhbjESMBAGA1UECgwJVm9pcCBUYWxrMQ4wDAYDVQQLDAVWb2ljZTERMA8GA1UEAwwIdm9pcHRhbGswHhcNMjMwOTAyMDcyODQ5WhcNMjUxMjA1MDcyODQ5WjBqMQswCQYDVQQGEwJVUzELMAkGA1UECAwCVkExEjAQBgNVBAcMCVNvbWV3aGVyZTEaMBgGA1UECgwRQWNtZVRlbGVjb20sIEluYy4xDTALBgNVBAsMBFZPSVAxDzANBgNVBAMMBlNIQUtFTjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABLgoB2l1eo%2B98fdgBpWWrqhTcb663lNTEh8thLh%2FB%2FeOxi%2Fc8QNFrU6%2FMiUIZvPzg79m4fsXIHZPySiGQqOiZTajGjAYMBYGCCsGAQUFBwEaBAowCKAGFgQxMDAxMAoGCCqGSM49BAMCA0gAMEUCIFqqayjTvpGUfDzlVEm6Dx3st9c50PBjeBidx2sFmwyXAiEA%2BjPHibj1PV4toDo%2BL9HYfqQcA9dAUBtzlb9bknW8S6A%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_ext_crl_distribution](../../ISSUES/e_atis_ext_crl_distribution/README.md) | error | ATIS1000080 | STI End-Entity certificates shall contain a CRL Distribution Points extension |
| [e_atis_subject_key_identifier_size](../../ISSUES/e_atis_subject_key_identifier_size/README.md) | error | ATIS1000080 | Subject Key Identifier extension not found |
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension does not contain a single OID value that identifies the SHAKEN Certificate Policy established by the STI-PA |
| [e_ext_authority_key_identifier_missing](../../ISSUES/e_ext_authority_key_identifier_missing/README.md) | error | RFC5280 |  |
| [e_atis_ext_certificate_policies](../../ISSUES/e_atis_ext_certificate_policies/README.md) | error | ATIS1000080 | The Certificate Policies extension is not present |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 1001', but common name is 'SHAKEN' |
| [e_atis_ext_authority_key_identifier](../../ISSUES/e_atis_ext_authority_key_identifier/README.md) | error | ATIS1000080 | STI certificates shall contain an Authority Key Identifier extension |
| [e_atis_ext_key_usage](../../ISSUES/e_atis_ext_key_usage/README.md) | error | ATIS1000080 | Key Usage extension not found |
| [e_atis_ext_basic_constraints](../../ISSUES/e_atis_ext_basic_constraints/README.md) | error | ATIS1000080 | BasicConstraints extension not found |
| [e_atis_subject_key_identifier](../../ISSUES/e_atis_subject_key_identifier/README.md) | error | ATIS1000080 | Subject Key Identifier extension not found |
| [w_ext_subject_key_identifier_missing_sub_cert](../../ISSUES/w_ext_subject_key_identifier_missing_sub_cert/README.md) | warn | RFC5280 |  |
| [e_ext_authority_key_identifier_no_key_identifier](../../ISSUES/e_ext_authority_key_identifier_no_key_identifier/README.md) | error | RFC5280 |  |


Generated: 24 Nov 23 11:17 UTC