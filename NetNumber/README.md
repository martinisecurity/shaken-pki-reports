# STIR/SHAKEN CA Ecosystem Compliance

## NetNumber

### Leaf Certificates

- Average validity span as of leaf certificates 113 days
- Percentage of leaf certificates expiring in the next 30 days is 75.00%
- Certificates with Errors: 4
- Certificates with Warnings: 1
- Certificates with Notices: 4
- Certificates with tests not executed as the requirements were Not Effective at issuance time: 0

| Status | Code | Source | Instances |
|--------|------|--------|-----------|
| notice | [n_sti_certificate_policy_critical](ISSUES/n_sti_certificate_policy_critical/README.md#leaf-certificates) | ATIS-1000080 | 4 |
| error | [e_cp1_3_subject_sn](ISSUES/e_cp1_3_subject_sn/README.md#leaf-certificates) | United States SHAKEN CP | 4 |
| error | [e_sti_certificate_policies](ISSUES/e_sti_certificate_policies/README.md#leaf-certificates) | ATIS-1000080 | 4 |
| error | [e_sti_subject_cn](ISSUES/e_sti_subject_cn/README.md#leaf-certificates) | ATIS-1000080 | 4 |
| warn | [w_cp1_3_subject_rdn_unknown](ISSUES/w_cp1_3_subject_rdn_unknown/README.md#leaf-certificates) | United States SHAKEN CP | 1 |
| error | [e_sti_crl_distribution](ISSUES/e_sti_crl_distribution/README.md#leaf-certificates) | ATIS-1000080 | 1 |
| error | [e_cp1_3_ambiguous_identifier](ISSUES/e_cp1_3_ambiguous_identifier/README.md#leaf-certificates) | United States SHAKEN CP | 4 |

\* Tests use the ATIS 1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.\
\*\* 1 certificates skipped because they are currently expired.\
\*\*\* 0 certificates skipped because they are not currently trusted by the STI-PA.

#### Issued certificates

| Created at | Name | Problems | Link |
|------------|------|----------|------|
| 10 Jun 22 00:00 UTC | Plivo Inc | true | [view](CERTIFICATES/0ebf970365dbca8232b80e72c6da7e05bb43d33a/README.md) |
| 02 Oct 22 01:01 UTC | Google SHAKEN cert 969H | true | [view](CERTIFICATES/ff1d02a3c6ad3f781b7f0c1ed0d1ce118ccf17fd/README.md) |
| 25 Oct 22 01:08 UTC | Google SHAKEN cert 969H | true | [view](CERTIFICATES/b217a06e40b398741fb71143b36092f26d2bfee5/README.md) |
| 26 Oct 22 15:00 UTC | Baltimore-Washington Telephone Company SHAKEN cert 8697 | true | [view](CERTIFICATES/58237e16980e6098eaf1ed3897e29619ffb27965/README.md) |

\* For issues relating to this CAs certificate repositories see this [report](URL/README.md).

### CA Certificates

- Certificates with Errors: 2
- Certificates with Warnings: 0
- Certificates with Notices: 0
- Certificates with tests not executed as the requirements were Not Effective at issuance time: 3

| Status | Code | Source | Instances |
|--------|------|--------|-----------|
| not effective | [e_sti_ca_crl_distribution](ISSUES/e_sti_ca_crl_distribution/README.md#ca-certificates) | ATIS-1000080 | 1 |
| not effective | [n_sti_ca_certificate_policy_critical](ISSUES/n_sti_ca_certificate_policy_critical/README.md#ca-certificates) | ATIS-1000080 | 1 |
| not effective | [w_cp1_3_ca_subject_rdn_unknown](ISSUES/w_cp1_3_ca_subject_rdn_unknown/README.md#ca-certificates) | United States SHAKEN CP | 3 |
| not effective | [e_sti_ca_subject_cn](ISSUES/e_sti_ca_subject_cn/README.md#ca-certificates) | ATIS-1000080 | 3 |
| not effective | [e_cp1_3_ca_key_usage_crl_sign](ISSUES/e_cp1_3_ca_key_usage_crl_sign/README.md#ca-certificates) | United States SHAKEN CP | 3 |
| not effective | [e_sti_ca_serial_number](ISSUES/e_sti_ca_serial_number/README.md#ca-certificates) | ATIS-1000080 | 3 |
| error | [e_sti_ca_signature_algorithm](ISSUES/e_sti_ca_signature_algorithm/README.md#ca-certificates) | ATIS-1000080 | 2 |
| not effective | [e_sti_root_extension_unknown](ISSUES/e_sti_root_extension_unknown/README.md#ca-certificates) | ATIS-1000080 | 2 |
| not effective | [e_sti_root_certificate_policies](ISSUES/e_sti_root_certificate_policies/README.md#ca-certificates) | ATIS-1000080 | 2 |
| error | [e_sti_ca_subject_public_key](ISSUES/e_sti_ca_subject_public_key/README.md#ca-certificates) | ATIS-1000080 | 2 |
| not effective | [e_sti_ca_extension_unknown](ISSUES/e_sti_ca_extension_unknown/README.md#ca-certificates) | ATIS-1000080 | 1 |
| not effective | [e_sti_ca_certificate_policies](ISSUES/e_sti_ca_certificate_policies/README.md#ca-certificates) | ATIS-1000080 | 1 |

\* Tests use the ATIS 1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.\
\*\* 0 certificates skipped because they are currently expired.\
\*\*\* 0 certificates skipped because they are not currently trusted by the STI-PA.

#### Issued certificates

| Created at | Name | Problems | Link |
|------------|------|----------|------|
| 12 Jul 21 23:25 UTC | NetNumber SHAKEN Root CA | false | [view](CERTIFICATES/f5e317e9218445de21deca1f67f25452db6f4242/README.md) |
| 27 Sep 21 19:45 UTC | NetNumber SHAKEN Root CA 1 | true | [view](CERTIFICATES/83319d7352105c9f04a6abbe72052c929cbdf6e2/README.md) |
| 29 Sep 21 13:22 UTC | NetNumber SHAKEN Root Intermediate CA 1 | true | [view](CERTIFICATES/563c5f6e1e683f67ab636b71bf04a9d49e4ccbae/README.md) |

Generated: 27/10/2022 at 21:42:52