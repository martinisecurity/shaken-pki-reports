# STIR/SHAKEN CA Ecosystem Compliance

## VoiceCarrier

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 1 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 0 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 1 certificates being tested against the remaining rules
- 14.00 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 100.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 246 days is the average remaining validity for the certificates in the corpus
- 365 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days
- 1.00 average number of unexpired certificates per OCN observed
- 1 unique OCNs observed in unexpired and valid certificate corpus

| Instances | Test | Source |
|-----------|------|--------|
| 1 | [e_atis_ext_authority_key_identifier](ISSUES/e_atis_ext_authority_key_identifier/README.md) | ATIS1000080 |
| 1 | [e_atis_ext_basic_constraints](ISSUES/e_atis_ext_basic_constraints/README.md) | ATIS1000080 |
| 1 | [e_atis_ext_certificate_policies](ISSUES/e_atis_ext_certificate_policies/README.md) | ATIS1000080 |
| 1 | [e_atis_ext_crl_distribution](ISSUES/e_atis_ext_crl_distribution/README.md) | ATIS1000080 |
| 1 | [e_atis_ext_key_usage](ISSUES/e_atis_ext_key_usage/README.md) | ATIS1000080 |
| 1 | [e_atis_subject_cn](ISSUES/e_atis_subject_cn/README.md) | ATIS1000080 |
| 1 | [e_atis_subject_key_identifier](ISSUES/e_atis_subject_key_identifier/README.md) | ATIS1000080 |
| 1 | [e_atis_subject_key_identifier_size](ISSUES/e_atis_subject_key_identifier_size/README.md) | ATIS1000080 |
| 1 | [e_atis_tn_auth_list](ISSUES/e_atis_tn_auth_list/README.md) | ATIS1000080 |
| 1 | [e_atis_tn_auth_list_spc_format](ISSUES/e_atis_tn_auth_list_spc_format/README.md) | ATIS1000080 |
| 1 | [e_atis_version](ISSUES/e_atis_version/README.md) | ATIS1000080 |
| 1 | [e_ext_authority_key_identifier_no_key_identifier](ISSUES/e_ext_authority_key_identifier_no_key_identifier/README.md) | RFC5280 |
| 1 | [e_shaken_certificate_policies_id](ISSUES/e_shaken_certificate_policies_id/README.md) | US_SHAKEN_CP |
| 1 | [w_ext_subject_key_identifier_missing_sub_cert](ISSUES/w_ext_subject_key_identifier_missing_sub_cert/README.md) | RFC5280 |

### Details

#### Leaf Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 22&#160;Dec&#160;24&#160;15:20&#160;UTC | www.voicecarrier.com | 22&#160;Dec&#160;25&#160;15:20&#160;UTC | true | [view](CERTS/1415b6414cfbcad7f1dfc15ede30a9eff415154035d0205d89bd1039ab268af2/README.md) |


Generated: 18 Aug 25 21:13 UTC