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
| error | [e_sti_extension_unknown](ISSUES/e_sti_extension_unknown.html#leaf-certificates) | ATIS-1000080v4 | 3 |
| error | [e_sti_certificate_policies](ISSUES/e_sti_certificate_policies.html#leaf-certificates) | ATIS-1000080v4 | 3 |
| error | [e_sti_crl_distribution](ISSUES/e_sti_crl_distribution.html#leaf-certificates) | ATIS-1000080v4 | 3 |
| warn | [w_cp1_3_subject_rdn_unknown](ISSUES/w_cp1_3_subject_rdn_unknown.html#leaf-certificates) | CPv1.3 | 3 |
| error | [e_cp1_3_subject_sn](ISSUES/e_cp1_3_subject_sn.html#leaf-certificates) | CPv1.3 | 3 |
| error | [e_cp1_3_ambiguous_identifier](ISSUES/e_cp1_3_ambiguous_identifier.html#leaf-certificates) | CPv1.3 | 3 |

\* Tests use the ATIS 1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.\
\*\* 0 certificates skipped because they are currently expired.\
\*\*\* 0 certificates skipped because they are not currently trusted by the STI-PA.

#### Issued certificates

| Created at | Name | Problems | Link |
|------------|------|----------|------|
| 29 Apr 22 17:43 UTC | Nuwave Communications SHAKEN 620J | true | [view](aa06ee0a2f5d5930087d1a2dc67e5d6d649506cf%2Findex.html) |
| 20 May 22 14:32 UTC | Veracity SHAKEN 716D | true | [view](eaa33c30a7def67c3c9acd4e19eaadb6b0f73ed1%2Findex.html) |
| 10 Jun 22 19:00 UTC | Netfortris SHAKEN 8886 | true | [view](917049a2c60a5a3116f69124efa6ec0dc9c119ed%2Findex.html) |

\* For issues relating to this CAs certificate repositories see this [report](URL.html).

### CA Certificates

- Certificates with Errors: 0
- Certificates with Warnings: 0
- Certificates with Notices: 2
- Certificates with tests not executed as the requirements were Not Effective at issuance time: 2

| Status | Code | Source | Instances |
|--------|------|--------|-----------|
| not effective | [e_sti_ca_version](ISSUES/e_sti_ca_version.html#ca-certificates) | ATIS-1000080v4 | 2 |
| not effective | [e_sti_ca_subject](ISSUES/e_sti_ca_subject.html#ca-certificates) | ATIS-1000080v4 | 2 |
| not effective | [e_sti_ca_signature_algorithm](ISSUES/e_sti_ca_signature_algorithm.html#ca-certificates) | ATIS-1000080v4 | 2 |
| not effective | [w_cp1_3_ca_subject_rdn_unknown](ISSUES/w_cp1_3_ca_subject_rdn_unknown.html#ca-certificates) | CPv1.3 | 2 |
| not effective | [e_sti_ca_serial_number](ISSUES/e_sti_ca_serial_number.html#ca-certificates) | ATIS-1000080v4 | 2 |
| not effective | [e_sti_ca_subject_cn](ISSUES/e_sti_ca_subject_cn.html#ca-certificates) | ATIS-1000080v4 | 2 |
| not effective | [e_sti_ca_key_usage](ISSUES/e_sti_ca_key_usage.html#ca-certificates) | ATIS-1000080v4 | 2 |
| not effective | [e_sti_root_certificate_policies](ISSUES/e_sti_root_certificate_policies.html#ca-certificates) | ATIS-1000080v4 | 1 |
| not effective | [e_sti_ca_extension_unknown](ISSUES/e_sti_ca_extension_unknown.html#ca-certificates) | ATIS-1000080v4 | 1 |
| not effective | [e_sti_ca_authority_key_identifier](ISSUES/e_sti_ca_authority_key_identifier.html#ca-certificates) | ATIS-1000080v4 | 1 |
| not effective | [e_sti_ca_crl_distribution](ISSUES/e_sti_ca_crl_distribution.html#ca-certificates) | ATIS-1000080v4 | 1 |
| not effective | [e_cp1_3_ca_key_usage_crl_sign](ISSUES/e_cp1_3_ca_key_usage_crl_sign.html#ca-certificates) | CPv1.3 | 2 |
| notice | [n_pki_ca_key_usage](ISSUES/n_pki_ca_key_usage.html#ca-certificates) | SHAKEN PKI Best Practice | 2 |
| not effective | [e_sti_ca_subject_key_identifier](ISSUES/e_sti_ca_subject_key_identifier.html#ca-certificates) | ATIS-1000080v4 | 2 |
| not effective | [e_sti_ca_issuer_dn](ISSUES/e_sti_ca_issuer_dn.html#ca-certificates) | ATIS-1000080v4 | 2 |
| not effective | [e_sti_root_extension_unknown](ISSUES/e_sti_root_extension_unknown.html#ca-certificates) | ATIS-1000080v4 | 1 |
| not effective | [n_sti_ca_certificate_policy_critical](ISSUES/n_sti_ca_certificate_policy_critical.html#ca-certificates) | ATIS-1000080v4 | 1 |
| not effective | [e_sti_ca_certificate_policies](ISSUES/e_sti_ca_certificate_policies.html#ca-certificates) | ATIS-1000080v4 | 1 |
| not effective | [e_sti_ca_subject_public_key](ISSUES/e_sti_ca_subject_public_key.html#ca-certificates) | ATIS-1000080v4 | 2 |
| not effective | [e_sti_basic_constraints](ISSUES/e_sti_basic_constraints.html#ca-certificates) | ATIS-1000080v4 | 2 |

\* Tests use the ATIS 1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.\
\*\* 0 certificates skipped because they are currently expired.\
\*\*\* 0 certificates skipped because they are not currently trusted by the STI-PA.

#### Issued certificates

| Created at | Name | Problems | Link |
|------------|------|----------|------|
| 13 May 21 00:00 UTC | SHAKEN Ribbon Root CA | false | [view](5e00d019ee3884ec7d8513deb8f4e8671a1db86e%2Findex.html) |
| 13 May 21 00:00 UTC | SHAKEN Ribbon Issuing CA | false | [view](cac26ee453b887be4c545426c19733add7c138c0%2Findex.html) |

Generated: 26/10/2022 at 23:14:41