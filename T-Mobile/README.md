# STIR/SHAKEN CA Ecosystem Compliance

## T-Mobile

### Leaf Certificates

- Average validity span as of leaf certificates 366 days
- Percentage of leaf certificates expiring in the next 30 days is 0.00%
- Certificates with Errors: 1
- Certificates with Warnings: 1
- Certificates with Notices: 0
- Certificates with tests not executed as the requirements were Not Effective at issuance time: 0

| Status | Code | Source | Instances |
|--------|------|--------|-----------|
| error | [e_sti_subject_cn](ISSUES/e_sti_subject_cn/README.md#leaf-certificates) | ATIS-1000080 | 1 |
| error | [e_sti_basic_constraints](ISSUES/e_sti_basic_constraints/README.md#leaf-certificates) | ATIS-1000080 | 1 |
| warn | [w_pki_subject_rdn_unknown](ISSUES/w_pki_subject_rdn_unknown/README.md#leaf-certificates) | SHAKEN PKI Best Practice | 1 |
| error | [e_cp1_3_ambiguous_identifier](ISSUES/e_cp1_3_ambiguous_identifier/README.md#leaf-certificates) | United States SHAKEN CP | 1 |
| error | [e_cp1_3_subject_sn](ISSUES/e_cp1_3_subject_sn/README.md#leaf-certificates) | United States SHAKEN CP | 1 |
| error | [e_sti_certificate_policies](ISSUES/e_sti_certificate_policies/README.md#leaf-certificates) | ATIS-1000080 | 1 |
| error | [e_sti_extension_unknown](ISSUES/e_sti_extension_unknown/README.md#leaf-certificates) | ATIS-1000080 | 1 |

\* Tests use the ATIS 1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.\
\*\* 0 certificates skipped because they are currently expired.\
\*\*\* 0 certificates skipped because they are not currently trusted by the STI-PA.

#### Issued certificates

| Created at | Name | Problems | Link |
|------------|------|----------|------|
| 14 Jul 22 21:05 UTC | cert.stir.t-mobile.com | true | [view](CERTIFICATES/9ac0def10180cb22aad9e8885715c9409031697a/README.md) |

\* For issues relating to this CAs certificate repositories see this [report](URL/README.md).

### CA Certificates

- Certificates with Errors: 0
- Certificates with Warnings: 0
- Certificates with Notices: 2
- Certificates with tests not executed as the requirements were Not Effective at issuance time: 2

| Status | Code | Source | Instances |
|--------|------|--------|-----------|
| not effective | [e_sti_ca_signature_algorithm](ISSUES/e_sti_ca_signature_algorithm/README.md#ca-certificates) | ATIS-1000080 | 2 |
| not effective | [e_sti_ca_version](ISSUES/e_sti_ca_version/README.md#ca-certificates) | ATIS-1000080 | 2 |
| not effective | [e_sti_ca_subject_key_identifier](ISSUES/e_sti_ca_subject_key_identifier/README.md#ca-certificates) | ATIS-1000080 | 2 |
| not effective | [e_sti_ca_extension_unknown](ISSUES/e_sti_ca_extension_unknown/README.md#ca-certificates) | ATIS-1000080 | 1 |
| not effective | [e_sti_ca_certificate_policies](ISSUES/e_sti_ca_certificate_policies/README.md#ca-certificates) | ATIS-1000080 | 1 |
| not effective | [e_sti_ca_subject](ISSUES/e_sti_ca_subject/README.md#ca-certificates) | ATIS-1000080 | 2 |
| not effective | [e_sti_ca_key_usage](ISSUES/e_sti_ca_key_usage/README.md#ca-certificates) | ATIS-1000080 | 2 |
| not effective | [e_sti_root_authority_key_identifier](ISSUES/e_sti_root_authority_key_identifier/README.md#ca-certificates) | ATIS-1000080 | 1 |
| not effective | [e_sti_root_certificate_policies](ISSUES/e_sti_root_certificate_policies/README.md#ca-certificates) | ATIS-1000080 | 1 |
| not effective | [e_sti_ca_crl_distribution](ISSUES/e_sti_ca_crl_distribution/README.md#ca-certificates) | ATIS-1000080 | 1 |
| not effective | [e_sti_ca_serial_number](ISSUES/e_sti_ca_serial_number/README.md#ca-certificates) | ATIS-1000080 | 2 |
| not effective | [e_cp1_3_ca_key_usage_crl_sign](ISSUES/e_cp1_3_ca_key_usage_crl_sign/README.md#ca-certificates) | United States SHAKEN CP | 2 |
| not effective | [e_sti_ca_issuer_dn](ISSUES/e_sti_ca_issuer_dn/README.md#ca-certificates) | ATIS-1000080 | 2 |
| not effective | [e_sti_root_extension_unknown](ISSUES/e_sti_root_extension_unknown/README.md#ca-certificates) | ATIS-1000080 | 1 |
| not effective | [n_sti_ca_certificate_policy_critical](ISSUES/n_sti_ca_certificate_policy_critical/README.md#ca-certificates) | ATIS-1000080 | 1 |
| not effective | [e_sti_ca_authority_key_identifier](ISSUES/e_sti_ca_authority_key_identifier/README.md#ca-certificates) | ATIS-1000080 | 1 |
| not effective | [e_sti_ca_subject_public_key](ISSUES/e_sti_ca_subject_public_key/README.md#ca-certificates) | ATIS-1000080 | 2 |
| not effective | [e_sti_basic_constraints](ISSUES/e_sti_basic_constraints/README.md#ca-certificates) | ATIS-1000080 | 2 |
| not effective | [e_sti_ca_subject_cn](ISSUES/e_sti_ca_subject_cn/README.md#ca-certificates) | ATIS-1000080 | 2 |
| notice | [n_pki_ca_key_usage](ISSUES/n_pki_ca_key_usage/README.md#ca-certificates) | SHAKEN PKI Best Practice | 2 |

\* Tests use the ATIS 1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.\
\*\* 0 certificates skipped because they are currently expired.\
\*\*\* 0 certificates skipped because they are not currently trusted by the STI-PA.

#### Issued certificates

| Created at | Name | Problems | Link |
|------------|------|----------|------|
| 19 Sep 19 20:12 UTC | TMOBILE-PROD-ROOT-STIRSHAKEN-EC | false | [view](CERTIFICATES/7dddc0874c3665ba6d3e5fce061c3e5ad7761511/README.md) |
| 27 Sep 19 17:12 UTC | TMOBILE-PROD-SUB-STIRSHAKEN-EC | false | [view](CERTIFICATES/45f4213a0916f509c15c42441aa6811f71c047ba/README.md) |

Generated: 28/10/2022 at 18:22:55