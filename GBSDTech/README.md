# STIR/SHAKEN CA Ecosystem Compliance

## GBSDTech

### CA Certificates

- Certificates with Errors: 1
- Certificates with Warnings: 0
- Certificates with Notices: 0
- Certificates with tests not executed as the requirements were Not Effective at issuance time: 2

| Status | Code | Source | Instances |
|--------|------|--------|-----------|
| not effective | [e_sti_ca_subject_key_identifier](ISSUES/e_sti_ca_subject_key_identifier.md#ca-certificates) | ATIS-1000080v4 | 2 |
| not effective | [w_cp1_3_ca_subject_rdn_unknown](ISSUES/w_cp1_3_ca_subject_rdn_unknown.md#ca-certificates) | CPv1.3 | 2 |
| not effective | [e_sti_ca_authority_key_identifier](ISSUES/e_sti_ca_authority_key_identifier.md#ca-certificates) | ATIS-1000080v4 | 1 |
| not effective | [e_sti_basic_constraints](ISSUES/e_sti_basic_constraints.md#ca-certificates) | ATIS-1000080v4 | 2 |
| not effective | [e_sti_ca_issuer_dn](ISSUES/e_sti_ca_issuer_dn.md#ca-certificates) | ATIS-1000080v4 | 2 |
| not effective | [e_sti_ca_serial_number](ISSUES/e_sti_ca_serial_number.md#ca-certificates) | ATIS-1000080v4 | 2 |
| not effective | [e_sti_ca_subject_cn](ISSUES/e_sti_ca_subject_cn.md#ca-certificates) | ATIS-1000080v4 | 2 |
| error | [e_ext_authority_key_identifier_missing](ISSUES/e_ext_authority_key_identifier_missing.md#ca-certificates) | RFC5280 | 1 |
| not effective | [e_sti_ca_subject](ISSUES/e_sti_ca_subject.md#ca-certificates) | ATIS-1000080v4 | 2 |
| not effective | [e_sti_ca_version](ISSUES/e_sti_ca_version.md#ca-certificates) | ATIS-1000080v4 | 2 |
| not effective | [e_sti_root_extension_unknown](ISSUES/e_sti_root_extension_unknown.md#ca-certificates) | ATIS-1000080v4 | 1 |
| not effective | [e_sti_ca_signature_algorithm](ISSUES/e_sti_ca_signature_algorithm.md#ca-certificates) | ATIS-1000080v4 | 2 |
| not effective | [e_sti_ca_crl_distribution](ISSUES/e_sti_ca_crl_distribution.md#ca-certificates) | ATIS-1000080v4 | 1 |
| not effective | [e_cp1_3_ca_key_usage_crl_sign](ISSUES/e_cp1_3_ca_key_usage_crl_sign.md#ca-certificates) | CPv1.3 | 2 |
| not effective | [e_sti_root_certificate_policies](ISSUES/e_sti_root_certificate_policies.md#ca-certificates) | ATIS-1000080v4 | 1 |
| not effective | [e_sti_ca_key_usage](ISSUES/e_sti_ca_key_usage.md#ca-certificates) | ATIS-1000080v4 | 2 |
| not effective | [e_sti_ca_extension_unknown](ISSUES/e_sti_ca_extension_unknown.md#ca-certificates) | ATIS-1000080v4 | 1 |
| error | [e_ext_authority_key_identifier_no_key_identifier](ISSUES/e_ext_authority_key_identifier_no_key_identifier.md#ca-certificates) | RFC5280 | 1 |
| not effective | [e_sti_ca_subject_public_key](ISSUES/e_sti_ca_subject_public_key.md#ca-certificates) | ATIS-1000080v4 | 2 |
| not effective | [e_sti_ca_certificate_policies](ISSUES/e_sti_ca_certificate_policies.md#ca-certificates) | ATIS-1000080v4 | 1 |

\* Tests use the ATIS 1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.\
\*\* 0 certificates skipped because they are currently expired.

#### Issued certificates

| Created at | Name | Problems | Link |
|------------|------|----------|------|
| 05 May 21 19:05 UTC | GBSDTech SHAKEN Root CA | false | [view](c6beb88bee7544f012d9579a8002bf774a717ef0%2FREADME.md) |
| 05 May 21 20:22 UTC | 1RouteGroup SHAKEN Intermediate CA | true | [view](b34acd5cf741f6c98726c200f39517c4bd02d4cd%2FREADME.md) |

Generated: 26/10/2022 at 20:21:29