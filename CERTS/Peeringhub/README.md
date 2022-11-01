# STIR/SHAKEN CA Ecosystem Compliance

## Peeringhub

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 7 certificates were included in the corpus being tested
- 0 repositories in the corpus were skipped because they are duplicates
- 2 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 5 certificates being tested against the remaining rules
- 100.00% of certificates contain one or more Error level issue
- 100.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 216 days is the average remaining validity for the certificates in the corpus
- 255 days is the average initial validity for the certificates in the corpus
- 1 certificates expire in the next 30 days

| Instances | Test Status | Source |
|-----------|-------------|--------|
| 5 | [e_cp1_3_ambiguous_identifier](ISSUES/e_cp1_3_ambiguous_identifier/README.md) | United States SHAKEN CP |
| 5 | [e_cp1_3_subject_sn](ISSUES/e_cp1_3_subject_sn/README.md) | United States SHAKEN CP |
| 5 | [e_sti_certificate_policies](ISSUES/e_sti_certificate_policies/README.md) | ATIS-1000080 |
| 5 | [w_pki_subject_rdn_unknown](ISSUES/w_pki_subject_rdn_unknown/README.md) | SHAKEN PKI Best Practice |

#### CA Certificates

- 2 certificates were included in the corpus being tested
- 0 repositories in the corpus were skipped because they are duplicates
- 0 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 2 certificates being tested against the remaining rules
- 50.00% of certificates contain one or more Error level issue
- 100.00% of certificates contain one or more Warning level issue
- 100.00% of certificates contain one or more Notice level issue
- 50.00% of certificates are too old to be assessed against currently enforced expectations
- 5856 days is the average remaining validity for the certificates in the corpus
- 5475 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days

| Instances | Test Status | Source |
|-----------|-------------|--------|
| 2 | [e_cp1_3_ca_key_usage_crl_sign](ISSUES/e_cp1_3_ca_key_usage_crl_sign/README.md) | United States SHAKEN CP |
| 1 | [e_sti_ca_certificate_policies](ISSUES/e_sti_ca_certificate_policies/README.md) | ATIS-1000080 |
| 1 | [e_sti_ca_serial_number](ISSUES/e_sti_ca_serial_number/README.md) | ATIS-1000080 |
| 1 | [e_sti_ca_subject_cn](ISSUES/e_sti_ca_subject_cn/README.md) | ATIS-1000080 |
| 1 | [e_sti_root_certificate_policies](ISSUES/e_sti_root_certificate_policies/README.md) | ATIS-1000080 |
| 1 | [e_sti_root_extension_unknown](ISSUES/e_sti_root_extension_unknown/README.md) | ATIS-1000080 |
| 2 | [n_pki_ca_key_usage](ISSUES/n_pki_ca_key_usage/README.md) | SHAKEN PKI Best Practice |
| 2 | [w_pki_ca_subject_rdn_unknown](ISSUES/w_pki_ca_subject_rdn_unknown/README.md) | SHAKEN PKI Best Practice |

### Details

#### Leaf Certificates

| Created at | Subject | Problems | Link |
|------------|---------|----------|------|
| 29 Jul 22 16:41 UTC | ATI SHAKEN 731J | true | [view](CERTS/bbc0a32743f1e659eff17172cbef96ee8cc3000aeb1dda85e1a47d47391ec3ab/README.md) |
| 31 Jul 22 09:09 UTC | Voiceterm SHAKEN 240K | true | [view](CERTS/84819934eb8b7f347d7365133dba4376d162d283fa9d09858aa1706ec88487c6/README.md) |
| 26 Aug 22 23:31 UTC | Teleinx SHAKEN 744J | true | [view](CERTS/743032377136fc18e443399c5fc57e36a5706188f141a522438a806143997925/README.md) |
| 15 Oct 22 00:00 UTC | VOCALTRANSIT SHAKEN 783J | true | [view](CERTS/81b78fff8a772249d72d4854d97672d7ac69a83c4900beaac699d28d220d8c13/README.md) |
| 19 Oct 22 19:24 UTC | TalkAsiaVoip LLC SHAKEN 198K | true | [view](CERTS/b49964274f962d87bf69c69b8c2efd07cd4de3d20155339e41ab2752477fd19b/README.md) |

#### CA Certificates

| Created at | Subject | Problems | Link |
|------------|---------|----------|------|
| 17 Dec 20 15:31 UTC | Peeringhub Inc Root CA | true | [view](CERTS/0ad4adc0b4d93fdb0e628c577020c73b8a5caff750e7e499f80ee2ab362a3f6a/README.md) |
| 22 Jun 22 22:45 UTC | Peeringhub Inc SHAKEN Intermediate CA 2 | true | [view](CERTS/f00871963a40b04269c4b019968e42f9f40964cbfb512ff5342307e9942874ce/README.md) |


Generated: 01/11/2022 at 07:33:04