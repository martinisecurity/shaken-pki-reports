# STIR/SHAKEN CA Ecosystem Compliance

## Peeringhub

### Leaf Certificates

- Average validity span as of leaf certificates 255 days
- Percentage of leaf certificates expiring in the next 30 days is 20.00%
- Certificates with Errors: 5
- Certificates with Warnings: 5
- Certificates with Notices: 0
- Certificates with tests not executed as the requirements were Not Effective at issuance time: 0

| Status | Code | Source | Instances |
|--------|------|--------|-----------|
| error | [e_cp1_3_subject_sn](ISSUES/e_cp1_3_subject_sn.md#leaf-certificates) | CPv1.3 | 5 |
| error | [e_sti_certificate_policies](ISSUES/e_sti_certificate_policies.md#leaf-certificates) | ATIS-1000080v4 | 5 |
| error | [e_sti_crl_distribution](ISSUES/e_sti_crl_distribution.md#leaf-certificates) | ATIS-1000080v4 | 5 |
| error | [e_cp1_3_ambiguous_identifier](ISSUES/e_cp1_3_ambiguous_identifier.md#leaf-certificates) | CPv1.3 | 5 |
| warn | [w_cp1_3_subject_rdn_unknown](ISSUES/w_cp1_3_subject_rdn_unknown.md#leaf-certificates) | CPv1.3 | 5 |

\* Tests use the ATIS 1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.\
\*\* 2 certificates skipped because they are currently expired.

#### Issued certificates

| Created at | Name | Problems | Link |
|------------|------|----------|------|
| 29 Jul 22 16:41 UTC | ATI SHAKEN 731J | true | [view](aa988e126a27ca2e888fcafa0b50d3785af8709b%2FREADME.md) |
| 31 Jul 22 09:09 UTC | Voiceterm SHAKEN 240K | true | [view](08231a00b8e4b762293c2716b289d4073ce7ed49%2FREADME.md) |
| 26 Aug 22 23:31 UTC | Teleinx SHAKEN 744J | true | [view](417122a7bf47730e788f52241d189a83b05cedc9%2FREADME.md) |
| 15 Oct 22 00:00 UTC | VOCALTRANSIT SHAKEN 783J | true | [view](3e8edc21d5a50f2efc695825d5c3c7e3e5d34d48%2FREADME.md) |
| 19 Oct 22 19:24 UTC | TalkAsiaVoip LLC SHAKEN 198K | true | [view](360f1867f89798d6c73ce738d31f5db88dc645fd%2FREADME.md) |

\* For issues relating to this CAs certificate repositories see this [report](URL.md).

### CA Certificates

- Certificates with Errors: 1
- Certificates with Warnings: 0
- Certificates with Notices: 1
- Certificates with tests not executed as the requirements were Not Effective at issuance time: 1

| Status | Code | Source | Instances |
|--------|------|--------|-----------|
| not effective | [e_cp1_3_ca_key_usage_crl_sign](ISSUES/e_cp1_3_ca_key_usage_crl_sign.md#ca-certificates) | CPv1.3 | 1 |
| error | [e_sti_ca_crl_distribution](ISSUES/e_sti_ca_crl_distribution.md#ca-certificates) | ATIS-1000080v4 | 1 |
| not effective | [w_cp1_3_ca_subject_rdn_unknown](ISSUES/w_cp1_3_ca_subject_rdn_unknown.md#ca-certificates) | CPv1.3 | 1 |
| error | [e_sti_ca_certificate_policies](ISSUES/e_sti_ca_certificate_policies.md#ca-certificates) | ATIS-1000080v4 | 1 |
| notice | [n_pki_ca_key_usage](ISSUES/n_pki_ca_key_usage.md#ca-certificates) | SHAKEN PKI Best Practice | 1 |

\* Tests use the ATIS 1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.\
\*\* 0 certificates skipped because they are currently expired.

#### Issued certificates

| Created at | Name | Problems | Link |
|------------|------|----------|------|
| 22 Jun 22 22:45 UTC | Peeringhub Inc SHAKEN Intermediate CA 2 | true | [view](7f08026ddf8b2e37428ba8218541a8437a2ab962%2FREADME.md) |

Generated: 26/10/2022 at 20:22:11