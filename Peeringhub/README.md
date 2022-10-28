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
| error | [e_sti_certificate_policies](ISSUES/e_sti_certificate_policies/README.md#leaf-certificates) | ATIS-1000080 | 5 |
| error | [e_cp1_3_subject_sn](ISSUES/e_cp1_3_subject_sn/README.md#leaf-certificates) | United States SHAKEN CP | 5 |
| warn | [w_pki_subject_rdn_unknown](ISSUES/w_pki_subject_rdn_unknown/README.md#leaf-certificates) | SHAKEN PKI Best Practice | 5 |
| error | [e_cp1_3_ambiguous_identifier](ISSUES/e_cp1_3_ambiguous_identifier/README.md#leaf-certificates) | United States SHAKEN CP | 5 |

\* Tests use the ATIS 1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.\
\*\* 2 certificates skipped because they are currently expired.\
\*\*\* 0 certificates skipped because they are not currently trusted by the STI-PA.

#### Issued certificates

| Created at | Name | Problems | Link |
|------------|------|----------|------|
| 29 Jul 22 16:41 UTC | ATI SHAKEN 731J | true | [view](CERTIFICATES/aa988e126a27ca2e888fcafa0b50d3785af8709b/README.md) |
| 31 Jul 22 09:09 UTC | Voiceterm SHAKEN 240K | true | [view](CERTIFICATES/08231a00b8e4b762293c2716b289d4073ce7ed49/README.md) |
| 26 Aug 22 23:31 UTC | Teleinx SHAKEN 744J | true | [view](CERTIFICATES/417122a7bf47730e788f52241d189a83b05cedc9/README.md) |
| 15 Oct 22 00:00 UTC | VOCALTRANSIT SHAKEN 783J | true | [view](CERTIFICATES/3e8edc21d5a50f2efc695825d5c3c7e3e5d34d48/README.md) |
| 19 Oct 22 19:24 UTC | TalkAsiaVoip LLC SHAKEN 198K | true | [view](CERTIFICATES/360f1867f89798d6c73ce738d31f5db88dc645fd/README.md) |

\* For issues relating to this CAs certificate repositories see this [report](URL/README.md).

### CA Certificates

- Certificates with Errors: 1
- Certificates with Warnings: 2
- Certificates with Notices: 2
- Certificates with tests not executed as the requirements were Not Effective at issuance time: 1

| Status | Code | Source | Instances |
|--------|------|--------|-----------|
| not effective | [e_sti_ca_serial_number](ISSUES/e_sti_ca_serial_number/README.md#ca-certificates) | ATIS-1000080 | 1 |
| not effective | [e_sti_root_extension_unknown](ISSUES/e_sti_root_extension_unknown/README.md#ca-certificates) | ATIS-1000080 | 1 |
| notice | [n_pki_ca_key_usage](ISSUES/n_pki_ca_key_usage/README.md#ca-certificates) | SHAKEN PKI Best Practice | 2 |
| not effective | [e_sti_ca_subject_cn](ISSUES/e_sti_ca_subject_cn/README.md#ca-certificates) | ATIS-1000080 | 1 |
| warn | [w_pki_ca_subject_rdn_unknown](ISSUES/w_pki_ca_subject_rdn_unknown/README.md#ca-certificates) | SHAKEN PKI Best Practice | 2 |
| error | [e_sti_ca_certificate_policies](ISSUES/e_sti_ca_certificate_policies/README.md#ca-certificates) | ATIS-1000080 | 1 |
| not effective | [e_sti_root_certificate_policies](ISSUES/e_sti_root_certificate_policies/README.md#ca-certificates) | ATIS-1000080 | 1 |
| not effective | [e_cp1_3_ca_key_usage_crl_sign](ISSUES/e_cp1_3_ca_key_usage_crl_sign/README.md#ca-certificates) | United States SHAKEN CP | 2 |

\* Tests use the ATIS 1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.\
\*\* 0 certificates skipped because they are currently expired.\
\*\*\* 0 certificates skipped because they are not currently trusted by the STI-PA.

#### Issued certificates

| Created at | Name | Problems | Link |
|------------|------|----------|------|
| 17 Dec 20 15:31 UTC | Peeringhub Inc Root CA | true | [view](CERTIFICATES/d2590d58fcb7193f05fdefc76a1897333a060627/README.md) |
| 22 Jun 22 22:45 UTC | Peeringhub Inc SHAKEN Intermediate CA 2 | true | [view](CERTIFICATES/7f08026ddf8b2e37428ba8218541a8437a2ab962/README.md) |

Generated: 28/10/2022 at 18:22:55