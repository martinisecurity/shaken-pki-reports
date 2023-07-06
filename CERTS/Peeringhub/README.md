# STIR/SHAKEN CA Ecosystem Compliance

## Peeringhub

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 15 certificates were included in the corpus being tested
- 2 certificates in the corpus were skipped because they are duplicates
- 2 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 11 certificates being tested against the remaining rules
- 3.09 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 9.09% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 224 days is the average remaining validity for the certificates in the corpus
- 242 days is the average initial validity for the certificates in the corpus
- 5 certificates expire in the next 30 days
- 1.00 average number of unexpired certificates per OCN observed
- 11 unique OCNs observed in unexpired and valid certificate corpus

| Instances | Test | Source |
|-----------|------|--------|
| 11 | [e_atis_certificate_policies](ISSUES/e_atis_certificate_policies/README.md) | ATIS1000080 |
| 11 | [e_us_cp_ambiguous_identifier](ISSUES/e_us_cp_ambiguous_identifier/README.md) | US_SHAKEN_CP |
| 11 | [e_us_cp_subject_sn](ISSUES/e_us_cp_subject_sn/README.md) | US_SHAKEN_CP |
| 1 | [w_shaken_subject_rdn_unknown](ISSUES/w_shaken_subject_rdn_unknown/README.md) | SHAKEN_PKI_BEST_PRACTICES |

#### CA Certificates

- 2 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 0 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 2 certificates being tested against the remaining rules
- 2.50 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 100.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 50.00% of certificates are too old to be assessed against currently enforced expectations
- 5773 days is the average remaining validity for the certificates in the corpus
- 5475 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days

| Instances | Test | Source |
|-----------|------|--------|
| 1 | [e_atis_ca_certificate_policies](ISSUES/e_atis_ca_certificate_policies/README.md) | ATIS1000080 |
| 2 | [e_us_cp_ca_key_usage_crl_sign](ISSUES/e_us_cp_ca_key_usage_crl_sign/README.md) | US_SHAKEN_CP |
| 2 | [w_shaken_ca_subject_rdn_unknown](ISSUES/w_shaken_ca_subject_rdn_unknown/README.md) | SHAKEN_PKI_BEST_PRACTICES |

### Details

#### Leaf Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 29&#160;Jul&#160;22&#160;16:41&#160;UTC | ATI SHAKEN 731J | 29&#160;Jul&#160;23&#160;16:41&#160;UTC | true | [view](CERTS/bbc0a32743f1e659eff17172cbef96ee8cc3000aeb1dda85e1a47d47391ec3ab/README.md) |
| 31&#160;Jul&#160;22&#160;09:09&#160;UTC | Voiceterm SHAKEN 240K | 31&#160;Jul&#160;23&#160;09:09&#160;UTC | true | [view](CERTS/84819934eb8b7f347d7365133dba4376d162d283fa9d09858aa1706ec88487c6/README.md) |
| 26&#160;Aug&#160;22&#160;23:31&#160;UTC | Teleinx SHAKEN 744J | 26&#160;Aug&#160;23&#160;23:31&#160;UTC | true | [view](CERTS/743032377136fc18e443399c5fc57e36a5706188f141a522438a806143997925/README.md) |
| 02&#160;Nov&#160;22&#160;07:51&#160;UTC | Apex Telecom LLC SHAKEN 288K | 02&#160;Nov&#160;23&#160;00:54&#160;UTC | true | [view](CERTS/e64593f4d9a81236af33c40c227d728e49418be50e05be2a8b65ec549a3b9275/README.md) |
| 16&#160;Nov&#160;22&#160;00:03&#160;UTC | Televoip SHAKEN 438K | 11&#160;Nov&#160;23&#160;00:08&#160;UTC | true | [view](CERTS/2e22094a03b8d7e2d2f37cb2ea5d5778d122fdbc3d0c83e3f8f984c8cf03a2e3/README.md) |
| 25&#160;May&#160;23&#160;20:39&#160;UTC | SHAKEN 088K 2023-05 | 24&#160;May&#160;24&#160;20:39&#160;UTC | true | [view](CERTS/1d7fe7341654ec5e2f3f89be1d2e410b4b1d979acb0b0b6b3ff6db0c9382eccb/README.md) |
| 01&#160;Jun&#160;23&#160;00:00&#160;UTC | VOCALTRANSIT SHAKEN 783J v3 | 15&#160;Nov&#160;23&#160;00:00&#160;UTC | true | [view](CERTS/7dc3f9db0b51eed136cea1a1de68f558a7efbac78d7381a341b68e7838db3814/README.md) |
| 21&#160;Jun&#160;23&#160;00:42&#160;UTC | Meta-lynk Telecom SHAKEN 442K | 30&#160;Jan&#160;24&#160;00:00&#160;UTC | true | [view](CERTS/723af9321b9721ed8c7efcecfc7c6dbb59b3de3b957d707fb030fa6006f29b8c/README.md) |
| 26&#160;Jun&#160;23&#160;07:42&#160;UTC | TalkAsiaVoip LLC SHAKEN 198K2 | 26&#160;Jul&#160;23&#160;12:36&#160;UTC | true | [view](CERTS/fab8826aac935fc398130ed821b89a7f1d091958cc2931569289fbd2b6658d1a/README.md) |
| 03&#160;Jul&#160;23&#160;16:58&#160;UTC | MAQS SHAKEN 255K | 02&#160;Aug&#160;23&#160;20:13&#160;UTC | true | [view](CERTS/5b4e11c8aecd5517e617f9352ad191a43987150a881dd8b039fcdefe072bbfed/README.md) |
| 05&#160;Jul&#160;23&#160;09:27&#160;UTC | Celerity Telecom SHAKEN 469K | 01&#160;Aug&#160;23&#160;15:59&#160;UTC | true | [view](CERTS/2b7502048fccddc59e02d8e03960b3c27c99d6207d3cd2cdf26db14a6dc6f13a/README.md) |

#### CA Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 17&#160;Dec&#160;20&#160;15:31&#160;UTC | Peeringhub Inc Root CA | 12&#160;Dec&#160;40&#160;15:31&#160;UTC | true | [view](CERTS/0ad4adc0b4d93fdb0e628c577020c73b8a5caff750e7e499f80ee2ab362a3f6a/README.md) |
| 22&#160;Jun&#160;22&#160;22:45&#160;UTC | Peeringhub Inc SHAKEN Intermediate CA 2 | 19&#160;Jun&#160;32&#160;22:45&#160;UTC | true | [view](CERTS/f00871963a40b04269c4b019968e42f9f40964cbfb512ff5342307e9942874ce/README.md) |


Generated: 06 Jul 23 14:08 UTC