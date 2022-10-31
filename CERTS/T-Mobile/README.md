# STIR/SHAKEN CA Ecosystem Compliance

## T-Mobile

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

1 certificates were included in the corpus being tested\
0 certificates in the corpus were skipped because they were expired\
0 certificates in the corpus were skipped because they are not currently trusted\
100.00% of certificates contain one or more Error level issue\
100.00% of certificates contain one or more Warning level issue\
0.00% of certificates contain one or more Notice level issue\
0.00% of certificates are too old to be assessed against currently enforced expectations\
312 days is the average remaining validity for the certificates in the corpus\
366 days is the average initial validity for the certificates in the corpus\
0 certificates expire in the next 30 days

| Instances | Test Status | Source |
|-----------|-------------|--------|
| 1 | [e_cp1_3_ambiguous_identifier](ISSUES/e_cp1_3_ambiguous_identifier/README.md) | United States SHAKEN CP |
| 1 | [e_cp1_3_subject_sn](ISSUES/e_cp1_3_subject_sn/README.md) | United States SHAKEN CP |
| 1 | [e_sti_basic_constraints](ISSUES/e_sti_basic_constraints/README.md) | ATIS-1000080 |
| 1 | [e_sti_certificate_policies](ISSUES/e_sti_certificate_policies/README.md) | ATIS-1000080 |
| 1 | [e_sti_extension_unknown](ISSUES/e_sti_extension_unknown/README.md) | ATIS-1000080 |
| 1 | [e_sti_subject_cn](ISSUES/e_sti_subject_cn/README.md) | ATIS-1000080 |
| 1 | [w_pki_subject_rdn_unknown](ISSUES/w_pki_subject_rdn_unknown/README.md) | SHAKEN PKI Best Practice |

#### CA Certificates

2 certificates were included in the corpus being tested\
0 certificates in the corpus were skipped because they were expired\
0 certificates in the corpus were skipped because they are not currently trusted\
0.00% of certificates contain one or more Error level issue\
0.00% of certificates contain one or more Warning level issue\
100.00% of certificates contain one or more Notice level issue\
100.00% of certificates are too old to be assessed against currently enforced expectations\
6317 days is the average remaining validity for the certificates in the corpus\
5478 days is the average initial validity for the certificates in the corpus\
0 certificates expire in the next 30 days

| Instances | Test Status | Source |
|-----------|-------------|--------|
| 2 | [e_cp1_3_ca_key_usage_crl_sign](ISSUES/e_cp1_3_ca_key_usage_crl_sign/README.md) | United States SHAKEN CP |
| 2 | [e_sti_basic_constraints](ISSUES/e_sti_basic_constraints/README.md) | ATIS-1000080 |
| 1 | [e_sti_ca_authority_key_identifier](ISSUES/e_sti_ca_authority_key_identifier/README.md) | ATIS-1000080 |
| 1 | [e_sti_ca_certificate_policies](ISSUES/e_sti_ca_certificate_policies/README.md) | ATIS-1000080 |
| 1 | [e_sti_ca_crl_distribution](ISSUES/e_sti_ca_crl_distribution/README.md) | ATIS-1000080 |
| 1 | [e_sti_ca_extension_unknown](ISSUES/e_sti_ca_extension_unknown/README.md) | ATIS-1000080 |
| 2 | [e_sti_ca_issuer_dn](ISSUES/e_sti_ca_issuer_dn/README.md) | ATIS-1000080 |
| 2 | [e_sti_ca_key_usage](ISSUES/e_sti_ca_key_usage/README.md) | ATIS-1000080 |
| 2 | [e_sti_ca_serial_number](ISSUES/e_sti_ca_serial_number/README.md) | ATIS-1000080 |
| 2 | [e_sti_ca_signature_algorithm](ISSUES/e_sti_ca_signature_algorithm/README.md) | ATIS-1000080 |
| 2 | [e_sti_ca_subject](ISSUES/e_sti_ca_subject/README.md) | ATIS-1000080 |
| 2 | [e_sti_ca_subject_cn](ISSUES/e_sti_ca_subject_cn/README.md) | ATIS-1000080 |
| 2 | [e_sti_ca_subject_key_identifier](ISSUES/e_sti_ca_subject_key_identifier/README.md) | ATIS-1000080 |
| 2 | [e_sti_ca_subject_public_key](ISSUES/e_sti_ca_subject_public_key/README.md) | ATIS-1000080 |
| 2 | [e_sti_ca_version](ISSUES/e_sti_ca_version/README.md) | ATIS-1000080 |
| 1 | [e_sti_root_authority_key_identifier](ISSUES/e_sti_root_authority_key_identifier/README.md) | ATIS-1000080 |
| 1 | [e_sti_root_certificate_policies](ISSUES/e_sti_root_certificate_policies/README.md) | ATIS-1000080 |
| 1 | [e_sti_root_extension_unknown](ISSUES/e_sti_root_extension_unknown/README.md) | ATIS-1000080 |
| 2 | [n_pki_ca_key_usage](ISSUES/n_pki_ca_key_usage/README.md) | SHAKEN PKI Best Practice |
| 1 | [n_sti_ca_certificate_policy_critical](ISSUES/n_sti_ca_certificate_policy_critical/README.md) | ATIS-1000080 |

### Details

#### Leaf Certificates

| Created at | Subject | Problems | Link |
|------------|---------|----------|------|
| 14 Jul 22 21:05 UTC | cert.stir.t-mobile.com | true | [view](CERTS/7f653e15453416082531011acd1d7dad4f664ddf5124f73e27d841138f4a89f8/README.md) |

#### CA Certificates

| Created at | Subject | Problems | Link |
|------------|---------|----------|------|
| 27 Sep 19 17:12 UTC | TMOBILE-PROD-SUB-STIRSHAKEN-EC | true | [view](CERTS/4378a3f136510465232246b4d6a72d600db7e11117ac0e3d59c65528f47c9c4f/README.md) |
| 19 Sep 19 20:12 UTC | TMOBILE-PROD-ROOT-STIRSHAKEN-EC | true | [view](CERTS/d54b8c44268da3eaee9c5483c289652d1bd7f82420891114475470adebf8bf1e/README.md) |


Generated: 31/10/2022 at 18:34:12