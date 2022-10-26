# STIR/SHAKEN CA Ecosystem Compliance

## Martini Security

### Leaf Certificates

- Average validity span as of leaf certificates 129 days
- Percentage of leaf certificates expiring in the next 30 days is 0.00%
- Certificates with Errors: 7
- Certificates with Warnings: 0
- Certificates with Notices: 4
- Certificates with tests not executed as the requirements were Not Effective at issuance time: 0

| Status | Code | Source | Instances |
|--------|------|--------|-----------|
| error | [e_sti_crl_distribution](ISSUES/e_sti_crl_distribution.md#leaf-certificates) | ATIS-1000080v4 | 7 |
| notice | [n_sti_certificate_policy_critical](ISSUES/n_sti_certificate_policy_critical.md#leaf-certificates) | ATIS-1000080v4 | 4 |

\* Tests use the ATIS 1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.\
\*\* 0 certificates skipped because they are currently expired.

#### Issued certificates

| Created at | Name | Problems | Link |
|------------|------|----------|------|
| 28 Sep 22 17:54 UTC | SHAKEN 709J | true | [view](9747ae9f98177798503a9a3a39d2bdd7d7fc83ff%2FREADME.md) |
| 28 Sep 22 18:19 UTC | SHAKEN 709J | true | [view](e0a6d57c161fdaf3bf3e66127a55faed687f369f%2FREADME.md) |
| 28 Sep 22 21:31 UTC | SHAKEN 709J | true | [view](2d4251600f5c82cdd43a3417e193174b012fa87d%2FREADME.md) |
| 29 Sep 22 18:20 UTC | SHAKEN 073K | true | [view](40f3c020740bc812d269c19b112caf6bd0bbf666%2FREADME.md) |
| 21 Oct 22 17:38 UTC | SHAKEN 709J | true | [view](fc9f83c53f5ddeae62cbc537b766317a9fd2833b%2FREADME.md) |
| 21 Oct 22 19:21 UTC | SHAKEN 709J | true | [view](24e9291180b8802ed47977791c8cc42c70506452%2FREADME.md) |
| 25 Oct 22 20:57 UTC | SHAKEN 073K | true | [view](8b08d98c9f8361f2f52a874484a17039d02b04f3%2FREADME.md) |

\* For issues relating to this CAs certificate repositories see this [report](URL.md).

### CA Certificates

- Certificates with Errors: 2
- Certificates with Warnings: 1
- Certificates with Notices: 1
- Certificates with tests not executed as the requirements were Not Effective at issuance time: 1

| Status | Code | Source | Instances |
|--------|------|--------|-----------|
| not effective | [w_cp1_3_ca_subject_rdn_unknown](ISSUES/w_cp1_3_ca_subject_rdn_unknown.md#ca-certificates) | CPv1.3 | 2 |
| notice | [n_sti_ca_certificate_policy_critical](ISSUES/n_sti_ca_certificate_policy_critical.md#ca-certificates) | ATIS-1000080v4 | 1 |
| error | [e_sti_ca_crl_distribution](ISSUES/e_sti_ca_crl_distribution.md#ca-certificates) | ATIS-1000080v4 | 2 |
| not effective | [e_cp1_3_ca_key_usage_crl_sign](ISSUES/e_cp1_3_ca_key_usage_crl_sign.md#ca-certificates) | CPv1.3 | 1 |

\* Tests use the ATIS 1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.\
\*\* 0 certificates skipped because they are currently expired.

#### Issued certificates

| Created at | Name | Problems | Link |
|------------|------|----------|------|
| 03 May 22 15:41 UTC | Martini Security SHAKEN G1 | true | [view](0fe086f321e93ca9ae08a19a89bf9049b7625fcf%2FREADME.md) |
| 02 Oct 22 10:40 UTC | Martini Security SHAKEN G2 | true | [view](6af848efe4c680e44114ee32350826704eca8135%2FREADME.md) |

Generated: 26/10/2022 at 20:32:17