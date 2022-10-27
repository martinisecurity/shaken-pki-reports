# STIR/SHAKEN CA Ecosystem Compliance

## Ribbon Communications

### Leaf Certificates

- Average validity span as of leaf certificates 365 days
- Percentage of leaf certificates expiring in the next 30 days is 0.00%
- Certificates with Errors: 3
- Certificates with Warnings: 3
- Certificates with Notices: 0
- Certificates with tests not executed as the requirements were Not Effective at issuance time: 0

| Status | Code | Source | Instances |
|--------|------|--------|-----------|
| error | [e_sti_certificate_policies](ISSUES/e_sti_certificate_policies/README.md#leaf-certificates) | ATIS-1000080v4 | 3 |
| error | [e_sti_extension_unknown](ISSUES/e_sti_extension_unknown/README.md#leaf-certificates) | ATIS-1000080v4 | 3 |
| error | [e_cp1_3_subject_sn](ISSUES/e_cp1_3_subject_sn/README.md#leaf-certificates) | CPv1.3 | 3 |
| warn | [w_cp1_3_subject_rdn_unknown](ISSUES/w_cp1_3_subject_rdn_unknown/README.md#leaf-certificates) | CPv1.3 | 3 |
| error | [e_cp1_3_ambiguous_identifier](ISSUES/e_cp1_3_ambiguous_identifier/README.md#leaf-certificates) | CPv1.3 | 3 |
| error | [e_sti_crl_distribution](ISSUES/e_sti_crl_distribution/README.md#leaf-certificates) | ATIS-1000080v4 | 3 |

\* Tests use the ATIS 1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.\
\*\* 0 certificates skipped because they are currently expired.\
\*\*\* 0 certificates skipped because they are not currently trusted by the STI-PA.

#### Issued certificates

| Created at | Name | Problems | Link |
|------------|------|----------|------|
| 29 Apr 22 17:43 UTC | Nuwave Communications SHAKEN 620J | true | [view](CERTIFICATES/aa06ee0a2f5d5930087d1a2dc67e5d6d649506cf/README.md) |
| 20 May 22 14:32 UTC | Veracity SHAKEN 716D | true | [view](CERTIFICATES/eaa33c30a7def67c3c9acd4e19eaadb6b0f73ed1/README.md) |
| 10 Jun 22 19:00 UTC | Netfortris SHAKEN 8886 | true | [view](CERTIFICATES/917049a2c60a5a3116f69124efa6ec0dc9c119ed/README.md) |

\* For issues relating to this CAs certificate repositories see this [report](URL/README.md).

### CA Certificates

- Certificates with Errors: 0
- Certificates with Warnings: 0
- Certificates with Notices: 2
- Certificates with tests not executed as the requirements were Not Effective at issuance time: 2

| Status | Code | Source | Instances |
|--------|------|--------|-----------|
| not effective | [e_sti_root_certificate_policies](ISSUES/e_sti_root_certificate_policies/README.md#ca-certificates) | ATIS-1000080v4 | 1 |
| not effective | [e_sti_ca_subject_key_identifier](ISSUES/e_sti_ca_subject_key_identifier/README.md#ca-certificates) | ATIS-1000080v4 | 2 |
| not effective | [e_sti_ca_certificate_policies](ISSUES/e_sti_ca_certificate_policies/README.md#ca-certificates) | ATIS-1000080v4 | 1 |
| not effective | [e_sti_ca_extension_unknown](ISSUES/e_sti_ca_extension_unknown/README.md#ca-certificates) | ATIS-1000080v4 | 1 |
| not effective | [n_sti_ca_certificate_policy_critical](ISSUES/n_sti_ca_certificate_policy_critical/README.md#ca-certificates) | ATIS-1000080v4 | 1 |
| not effective | [e_sti_ca_serial_number](ISSUES/e_sti_ca_serial_number/README.md#ca-certificates) | ATIS-1000080v4 | 2 |
| not effective | [e_sti_root_extension_unknown](ISSUES/e_sti_root_extension_unknown/README.md#ca-certificates) | ATIS-1000080v4 | 1 |
| not effective | [e_sti_ca_key_usage](ISSUES/e_sti_ca_key_usage/README.md#ca-certificates) | ATIS-1000080v4 | 2 |
| not effective | [e_sti_basic_constraints](ISSUES/e_sti_basic_constraints/README.md#ca-certificates) | ATIS-1000080v4 | 2 |
| not effective | [e_sti_ca_subject_cn](ISSUES/e_sti_ca_subject_cn/README.md#ca-certificates) | ATIS-1000080v4 | 2 |
| not effective | [e_sti_ca_subject_public_key](ISSUES/e_sti_ca_subject_public_key/README.md#ca-certificates) | ATIS-1000080v4 | 2 |
| not effective | [e_sti_ca_signature_algorithm](ISSUES/e_sti_ca_signature_algorithm/README.md#ca-certificates) | ATIS-1000080v4 | 2 |
| notice | [n_pki_ca_key_usage](ISSUES/n_pki_ca_key_usage/README.md#ca-certificates) | SHAKEN PKI Best Practice | 2 |
| not effective | [w_cp1_3_ca_subject_rdn_unknown](ISSUES/w_cp1_3_ca_subject_rdn_unknown/README.md#ca-certificates) | CPv1.3 | 2 |
| not effective | [e_sti_ca_crl_distribution](ISSUES/e_sti_ca_crl_distribution/README.md#ca-certificates) | ATIS-1000080v4 | 1 |
| not effective | [e_sti_ca_subject](ISSUES/e_sti_ca_subject/README.md#ca-certificates) | ATIS-1000080v4 | 2 |
| not effective | [e_sti_ca_issuer_dn](ISSUES/e_sti_ca_issuer_dn/README.md#ca-certificates) | ATIS-1000080v4 | 2 |
| not effective | [e_cp1_3_ca_key_usage_crl_sign](ISSUES/e_cp1_3_ca_key_usage_crl_sign/README.md#ca-certificates) | CPv1.3 | 2 |
| not effective | [e_sti_ca_version](ISSUES/e_sti_ca_version/README.md#ca-certificates) | ATIS-1000080v4 | 2 |
| not effective | [e_sti_ca_authority_key_identifier](ISSUES/e_sti_ca_authority_key_identifier/README.md#ca-certificates) | ATIS-1000080v4 | 1 |

\* Tests use the ATIS 1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.\
\*\* 0 certificates skipped because they are currently expired.\
\*\*\* 0 certificates skipped because they are not currently trusted by the STI-PA.

#### Issued certificates

| Created at | Name | Problems | Link |
|------------|------|----------|------|
| 13 May 21 00:00 UTC | SHAKEN Ribbon Root CA | false | [view](CERTIFICATES/5e00d019ee3884ec7d8513deb8f4e8671a1db86e/README.md) |
| 13 May 21 00:00 UTC | SHAKEN Ribbon Issuing CA | false | [view](CERTIFICATES/cac26ee453b887be4c545426c19733add7c138c0/README.md) |

Generated: 27/10/2022 at 18:24:52