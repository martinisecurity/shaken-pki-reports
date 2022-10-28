# STIR/SHAKEN CA Ecosystem Compliance

## Martini Security

### Leaf Certificates

- Average validity span as of leaf certificates 129 days
- Percentage of leaf certificates expiring in the next 30 days is 0.00%
- Certificates with Errors: 0
- Certificates with Warnings: 0
- Certificates with Notices: 4
- Certificates with tests not executed as the requirements were Not Effective at issuance time: 0

| Status | Code | Source | Instances |
|--------|------|--------|-----------|
| notice | [n_sti_certificate_policy_critical](ISSUES/n_sti_certificate_policy_critical/README.md#leaf-certificates) | ATIS-1000080 | 4 |

\* Tests use the ATIS 1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.\
\*\* 0 certificates skipped because they are currently expired.\
\*\*\* 0 certificates skipped because they are not currently trusted by the STI-PA.

#### Issued certificates

| Created at | Name | Problems | Link |
|------------|------|----------|------|
| 28 Sep 22 17:54 UTC | SHAKEN 709J | false | [view](CERTIFICATES/9747ae9f98177798503a9a3a39d2bdd7d7fc83ff/README.md) |
| 28 Sep 22 18:19 UTC | SHAKEN 709J | false | [view](CERTIFICATES/e0a6d57c161fdaf3bf3e66127a55faed687f369f/README.md) |
| 28 Sep 22 21:31 UTC | SHAKEN 709J | false | [view](CERTIFICATES/2d4251600f5c82cdd43a3417e193174b012fa87d/README.md) |
| 29 Sep 22 18:20 UTC | SHAKEN 073K | false | [view](CERTIFICATES/40f3c020740bc812d269c19b112caf6bd0bbf666/README.md) |
| 21 Oct 22 17:38 UTC | SHAKEN 709J | false | [view](CERTIFICATES/fc9f83c53f5ddeae62cbc537b766317a9fd2833b/README.md) |
| 21 Oct 22 19:21 UTC | SHAKEN 709J | false | [view](CERTIFICATES/24e9291180b8802ed47977791c8cc42c70506452/README.md) |
| 25 Oct 22 20:57 UTC | SHAKEN 073K | false | [view](CERTIFICATES/8b08d98c9f8361f2f52a874484a17039d02b04f3/README.md) |

\* For issues relating to this CAs certificate repositories see this [report](URL/README.md).

### CA Certificates

- Certificates with Errors: 0
- Certificates with Warnings: 3
- Certificates with Notices: 1
- Certificates with tests not executed as the requirements were Not Effective at issuance time: 0

| Status | Code | Source | Instances |
|--------|------|--------|-----------|
| warn | [w_pki_ca_subject_rdn_unknown](ISSUES/w_pki_ca_subject_rdn_unknown/README.md#ca-certificates) | SHAKEN PKI Best Practice | 3 |
| notice | [n_sti_ca_certificate_policy_critical](ISSUES/n_sti_ca_certificate_policy_critical/README.md#ca-certificates) | ATIS-1000080 | 1 |

\* Tests use the ATIS 1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.\
\*\* 0 certificates skipped because they are currently expired.\
\*\*\* 0 certificates skipped because they are not currently trusted by the STI-PA.

#### Issued certificates

| Created at | Name | Problems | Link |
|------------|------|----------|------|
| 03 May 22 15:31 UTC | Martini Security SHAKEN R1 | true | [view](CERTIFICATES/0ffc3156067f94ee63a0ff6f33fce9cff5f51d39/README.md) |
| 03 May 22 15:41 UTC | Martini Security SHAKEN G1 | true | [view](CERTIFICATES/0fe086f321e93ca9ae08a19a89bf9049b7625fcf/README.md) |
| 02 Oct 22 10:40 UTC | Martini Security SHAKEN G2 | true | [view](CERTIFICATES/6af848efe4c680e44114ee32350826704eca8135/README.md) |

Generated: 28/10/2022 at 16:28:22