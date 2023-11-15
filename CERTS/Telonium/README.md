# STIR/SHAKEN CA Ecosystem Compliance

## Telonium

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 18 certificates were included in the corpus being tested
- 3 certificates in the corpus were skipped because they are duplicates
- 0 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 15 certificates being tested against the remaining rules
- 3.47 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 0.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 440 days is the average remaining validity for the certificates in the corpus
- 430 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days
- 1.36 average number of unexpired certificates per OCN observed
- 11 unique OCNs observed in unexpired and valid certificate corpus

| Instances | Test | Source |
|-----------|------|--------|
| 5 | [e_atis_certificate_policies](ISSUES/e_atis_certificate_policies/README.md) | ATIS1000080 |
| 15 | [e_atis_extension_unknown](ISSUES/e_atis_extension_unknown/README.md) | ATIS1000080 |
| 15 | [e_atis_issuer_dn](ISSUES/e_atis_issuer_dn/README.md) | ATIS1000080 |
| 1 | [e_atis_subject](ISSUES/e_atis_subject/README.md) | ATIS1000080 |
| 5 | [e_atis_subject_cn](ISSUES/e_atis_subject_cn/README.md) | ATIS1000080 |
| 5 | [e_atis_tn_auth_list](ISSUES/e_atis_tn_auth_list/README.md) | ATIS1000080 |
| 5 | [e_us_cp_ambiguous_identifier](ISSUES/e_us_cp_ambiguous_identifier/README.md) | US_SHAKEN_CP |
| 1 | [e_us_cp_subject_sn](ISSUES/e_us_cp_subject_sn/README.md) | US_SHAKEN_CP |

#### CA Certificates

- 3 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 0 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 3 certificates being tested against the remaining rules
- 3.00 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 0.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 4684 days is the average remaining validity for the certificates in the corpus
- 4503 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days

| Instances | Test | Source |
|-----------|------|--------|
| 1 | [e_atis_ca_certificate_policies](ISSUES/e_atis_ca_certificate_policies/README.md) | ATIS1000080 |
| 1 | [e_atis_ca_crl_distribution](ISSUES/e_atis_ca_crl_distribution/README.md) | ATIS1000080 |
| 2 | [e_atis_ca_key_usage](ISSUES/e_atis_ca_key_usage/README.md) | ATIS1000080 |
| 1 | [e_atis_ca_subject](ISSUES/e_atis_ca_subject/README.md) | ATIS1000080 |
| 3 | [e_atis_ca_subject_cn](ISSUES/e_atis_ca_subject_cn/README.md) | ATIS1000080 |
| 1 | [e_us_cp_ca_key_usage_crl_sign](ISSUES/e_us_cp_ca_key_usage_crl_sign/README.md) | US_SHAKEN_CP |

### Details

#### Leaf Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 15&#160;Jun&#160;23&#160;18:33&#160;UTC | COMMTRUNKS LLC 622K | 15&#160;Jun&#160;24&#160;18:34&#160;UTC | true | [view](CERTS/5b034b208d3cb8a47d37b81c26a517efccb4e6c4035535a25bb75ddb0030d98b/README.md) |
| 22&#160;Jun&#160;23&#160;18:32&#160;UTC | SHAKEN 755J | 22&#160;Jun&#160;24&#160;18:33&#160;UTC | true | [view](CERTS/08f4adc952c91a69124d0ee71487d19be19f15da68284d2a57ae628b8570fc4c/README.md) |
| 27&#160;Jun&#160;23&#160;15:30&#160;UTC | SHAKEN 633K | 27&#160;Jun&#160;24&#160;15:31&#160;UTC | true | [view](CERTS/afbfce2e5623baca8b3348d69ce703430902c5b33dd88cdcf091897bcc918f33/README.md) |
| 29&#160;Jun&#160;23&#160;20:24&#160;UTC | SHAKEN 620K | 29&#160;Jun&#160;24&#160;20:25&#160;UTC | true | [view](CERTS/5ba1b3123d3fb9c6067decf2e77ed58b9fe8bceaed94df2bfd5a3fa71d3202ac/README.md) |
| 30&#160;Jun&#160;23&#160;21:56&#160;UTC | SHAKEN 627K | 30&#160;Jun&#160;24&#160;21:57&#160;UTC | true | [view](CERTS/d5961ea754d88207a4cc8b33b3617880761c5775850e13f5ec4145954cdd8a20/README.md) |
| 10&#160;Jul&#160;23&#160;14:44&#160;UTC | SHAKEN 656K | 28&#160;Jun&#160;25&#160;19:35&#160;UTC | true | [view](CERTS/fd285a02878cd6c995f532bbc9b51b4721f5dd3828aa271194760254eb1cd549/README.md) |
| 10&#160;Jul&#160;23&#160;19:07&#160;UTC | SHAKEN 627K | 30&#160;Jun&#160;24&#160;21:57&#160;UTC | true | [view](CERTS/184df5cc205f9b14ec71d0170495cd323aaa179020d0b49b661a9a95c4927566/README.md) |
| 10&#160;Jul&#160;23&#160;19:11&#160;UTC | SHAKEN 620K | 29&#160;Jun&#160;24&#160;20:25&#160;UTC | true | [view](CERTS/6b84b7d78cbf691b21e9ac2ea3b820a39bf682833a822a52239a570307f29643/README.md) |
| 10&#160;Jul&#160;23&#160;19:40&#160;UTC | SHAKEN 755J | 22&#160;Jun&#160;24&#160;18:33&#160;UTC | true | [view](CERTS/c8ceddb6f8d9225a7450207343f5f769c0c7dbf97f4a478cc47f419466946815/README.md) |
| 11&#160;Jul&#160;23&#160;18:10&#160;UTC | SHAKEN 084E | 28&#160;Jun&#160;24&#160;15:28&#160;UTC | true | [view](CERTS/f990b49e4c21e28f4470ea06c91bbf2c57dbd8fc71b97c78f87d94bc4c3bd487/README.md) |
| 11&#160;Jul&#160;23&#160;18:13&#160;UTC | SHAKEN 633K | 27&#160;Jun&#160;24&#160;15:31&#160;UTC | true | [view](CERTS/3ebd2a8bc24714dc79cc0bc08170c05f9ab5a4b8eeec4033ef252d098b0712cf/README.md) |
| 11&#160;Jul&#160;23&#160;18:15&#160;UTC | SHAKEN 159K | 22&#160;Jun&#160;25&#160;22:42&#160;UTC | true | [view](CERTS/d9b2af8d51b09c75e70a5c6d60de90a663be51c44f141639708ab7a6cf151dfa/README.md) |
| 11&#160;Jul&#160;23&#160;18:45&#160;UTC | SHAKEN 421K | 21&#160;Jun&#160;25&#160;04:00&#160;UTC | true | [view](CERTS/cfaa04fe3747df8c203aa057940d1d7b1b455b15c8da2ba73cf7384797edfad8/README.md) |
| 11&#160;Jul&#160;23&#160;18:48&#160;UTC | SHAKEN 622K | 15&#160;Jun&#160;24&#160;18:34&#160;UTC | true | [view](CERTS/55851a261d59fda3e570cabe18693d44936ef44de75156c69b6b26bb7f0eeb6f/README.md) |
| 13&#160;Jul&#160;23&#160;05:56&#160;UTC | SHAKEN 651K | 13&#160;Jul&#160;24&#160;05:57&#160;UTC | true | [view](CERTS/df89a5a49894ef3f9991e3ec4c98e0b3df7ad85ac81dfbe2e11b5bfc815da364/README.md) |

#### CA Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 08&#160;Mar&#160;23&#160;18:40&#160;UTC | Telonium STI-CA Root1 | 05&#160;Mar&#160;35&#160;18:40&#160;UTC | true | [view](CERTS/96c66865ce5558c2ce3723c0b414538fcacadcd0f3286108fef57dc447f122f9/README.md) |
| 08&#160;Mar&#160;23&#160;18:40&#160;UTC | Telonium STI-CA Root2 | 07&#160;Mar&#160;38&#160;18:40&#160;UTC | true | [view](CERTS/a58b27999411d3d54121d4eadc82aa128be1fef96cda3029b2015677188ea40b/README.md) |
| 09&#160;Mar&#160;23&#160;15:18&#160;UTC | Telonium STI-CA Intermediate CA | 06&#160;Mar&#160;33&#160;15:18&#160;UTC | true | [view](CERTS/7c701216e591c9a3b84550ff46566dd420c7f182eb3cfc5abe5739cdbe271169/README.md) |


Generated: 15 Nov 23 18:10 UTC