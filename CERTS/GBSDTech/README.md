# STIR/SHAKEN CA Ecosystem Compliance

## GBSDTech

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
303 days is the average remaining validity for the certificates in the corpus\
365 days is the average initial validity for the certificates in the corpus\
0 certificates expire in the next 30 days

| Instances | Test Status | Source |
|-----------|-------------|--------|
| 1 | [e_cp1_3_ambiguous_identifier](ISSUES/e_cp1_3_ambiguous_identifier/README.md) | United States SHAKEN CP |
| 1 | [e_cp1_3_subject_sn](ISSUES/e_cp1_3_subject_sn/README.md) | United States SHAKEN CP |
| 1 | [e_sti_certificate_policies](ISSUES/e_sti_certificate_policies/README.md) | ATIS-1000080 |
| 1 | [e_sti_crl_distribution](ISSUES/e_sti_crl_distribution/README.md) | ATIS-1000080 |
| 1 | [e_sti_serial_number](ISSUES/e_sti_serial_number/README.md) | ATIS-1000080 |
| 1 | [e_sti_subject_cn](ISSUES/e_sti_subject_cn/README.md) | ATIS-1000080 |
| 1 | [e_sti_tn_auth_list](ISSUES/e_sti_tn_auth_list/README.md) | ATIS-1000080 |
| 1 | [w_pki_subject_rdn_unknown](ISSUES/w_pki_subject_rdn_unknown/README.md) | SHAKEN PKI Best Practice |

#### CA Certificates

2 certificates were included in the corpus being tested\
0 certificates in the corpus were skipped because they were expired\
0 certificates in the corpus were skipped because they are not currently trusted\
50.00% of certificates contain one or more Error level issue\
100.00% of certificates contain one or more Warning level issue\
0.00% of certificates contain one or more Notice level issue\
100.00% of certificates are too old to be assessed against currently enforced expectations\
7119 days is the average remaining validity for the certificates in the corpus\
7300 days is the average initial validity for the certificates in the corpus\
0 certificates expire in the next 30 days

| Instances | Test Status | Source |
|-----------|-------------|--------|
| 1 | [e_cp1_3_ca_key_usage_crl_sign](ISSUES/e_cp1_3_ca_key_usage_crl_sign/README.md) | United States SHAKEN CP |
| 1 | [e_ext_authority_key_identifier_missing](ISSUES/e_ext_authority_key_identifier_missing/README.md) | RFC5280 |
| 1 | [e_ext_authority_key_identifier_no_key_identifier](ISSUES/e_ext_authority_key_identifier_no_key_identifier/README.md) | RFC5280 |
| 1 | [e_sti_ca_authority_key_identifier](ISSUES/e_sti_ca_authority_key_identifier/README.md) | ATIS-1000080 |
| 1 | [e_sti_ca_certificate_policies](ISSUES/e_sti_ca_certificate_policies/README.md) | ATIS-1000080 |
| 1 | [e_sti_ca_crl_distribution](ISSUES/e_sti_ca_crl_distribution/README.md) | ATIS-1000080 |
| 1 | [e_sti_ca_extension_unknown](ISSUES/e_sti_ca_extension_unknown/README.md) | ATIS-1000080 |
| 1 | [e_sti_ca_serial_number](ISSUES/e_sti_ca_serial_number/README.md) | ATIS-1000080 |
| 1 | [e_sti_ca_subject_cn](ISSUES/e_sti_ca_subject_cn/README.md) | ATIS-1000080 |
| 1 | [e_sti_root_certificate_policies](ISSUES/e_sti_root_certificate_policies/README.md) | ATIS-1000080 |
| 1 | [e_sti_root_extension_unknown](ISSUES/e_sti_root_extension_unknown/README.md) | ATIS-1000080 |
| 1 | [w_pki_ca_subject_rdn_unknown](ISSUES/w_pki_ca_subject_rdn_unknown/README.md) | SHAKEN PKI Best Practice |

### Details

#### Leaf Certificates

| Created at | Subject | Problems | Link |
|------------|---------|----------|------|
| 28 Jun 22 18:13 UTC | MYPBXManager SHAKEN | true | [view](CERTS/0c7bf2cc1741b8036003876afadd109dfd5a6b0fb7af3662ae4d02e3340ad9ce/README.md) |

#### CA Certificates

| Created at | Subject | Problems | Link |
|------------|---------|----------|------|
| 05 May 21 20:22 UTC | 1RouteGroup SHAKEN Intermediate CA | true | [view](CERTS/99e9a67644a30ebc33ecc9aa8df6335524d49a4691164e357c5d2406b58a578e/README.md) |
| 05 May 21 19:05 UTC | GBSDTech SHAKEN Root CA | true | [view](CERTS/6d2bee73a01c1c9fe92273ff8ba56e0c870b7b901cbebcc9e12226fc109e1af9/README.md) |


Generated: 31/10/2022 at 16:43:22