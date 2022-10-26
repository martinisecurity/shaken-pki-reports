# STIR/SHAKEN CA Ecosystem Compliance

## 1RouteGroup

### Leaf Certificates

- Average validity span as of leaf certificates 365 days
- Percentage of leaf certificates expiring in the next 30 days is 0.00%
- Certificates with Errors: 1
- Certificates with Warnings: 1
- Certificates with Notices: 0
- Certificates with tests not executed as the requirements were Not Effective at issuance time: 0

| Status | Code | Source | Instances |
|--------|------|--------|-----------|
| error | [e_cp1_3_ambiguous_identifier](ISSUES/e_cp1_3_ambiguous_identifier.md#leaf-certificates) | CPv1.3 | 1 |
| error | [e_cp1_3_subject_sn](ISSUES/e_cp1_3_subject_sn.md#leaf-certificates) | CPv1.3 | 1 |
| error | [e_sti_crl_distribution](ISSUES/e_sti_crl_distribution.md#leaf-certificates) | ATIS-1000080v4 | 1 |
| error | [e_sti_tn_auth_list](ISSUES/e_sti_tn_auth_list.md#leaf-certificates) | ATIS-1000080v4 | 1 |
| error | [e_sti_certificate_policies](ISSUES/e_sti_certificate_policies.md#leaf-certificates) | ATIS-1000080v4 | 1 |
| error | [e_sti_subject_cn](ISSUES/e_sti_subject_cn.md#leaf-certificates) | ATIS-1000080v4 | 1 |
| warn | [w_cp1_3_subject_rdn_unknown](ISSUES/w_cp1_3_subject_rdn_unknown.md#leaf-certificates) | CPv1.3 | 1 |
| error | [e_sti_serial_number](ISSUES/e_sti_serial_number.md#leaf-certificates) | ATIS-1000080v4 | 1 |

\* Tests use the ATIS 1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.\
\*\* 0 certificates skipped because they are currently expired.

#### Issued certificates

| Created at | Name | Problems | Link |
|------------|------|----------|------|
| 28 Jun 22 18:13 UTC | MYPBXManager SHAKEN | true | [view](a3872afd09406d2745d204893b6b52bbf6380f84%2FREADME.md) |

\* For issues relating to this CAs certificate repositories see this [report](URL.md).

Generated: 26/10/2022 at 20:32:17