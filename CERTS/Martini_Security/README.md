# STIR/SHAKEN CA Ecosystem Compliance

## Martini Security

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 92 certificates were included in the corpus being tested
- 32 certificates in the corpus were skipped because they are duplicates
- 30 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 30 certificates being tested against the remaining rules
- 1.00 issues on average found in unexpired, trusted, and non-compliant certificates
- 0.00% of certificates contain one or more Error level issue
- 0.00% of certificates contain one or more Warning level issue
- 3.33% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 86 days is the average remaining validity for the certificates in the corpus
- 89 days is the average initial validity for the certificates in the corpus
- 9 certificates expire in the next 30 days
- 2.00 average number of unexpired certificates per OCN observed
- 15 unique OCNs observed in unexpired and valid certificate corpus

| Instances | Test | Source |
|-----------|------|--------|
| 1 | [n_atis_certificate_policy_critical](ISSUES/n_atis_certificate_policy_critical/README.md) | ATIS1000080 |

#### CA Certificates

- 3 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 0 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 3 certificates being tested against the remaining rules
- 1.00 issues on average found in unexpired, trusted, and non-compliant certificates
- 0.00% of certificates contain one or more Error level issue
- 0.00% of certificates contain one or more Warning level issue
- 33.33% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 5399 days is the average remaining validity for the certificates in the corpus
- 4261 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days

| Instances | Test | Source |
|-----------|------|--------|
| 1 | [n_atis_ca_certificate_policy_critical](ISSUES/n_atis_ca_certificate_policy_critical/README.md) | ATIS1000080 |

### Details

#### Leaf Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 28&#160;Sep&#160;22&#160;18:19&#160;UTC | SHAKEN 709J | 28&#160;Sep&#160;23&#160;17:45&#160;UTC | true | [view](CERTS/8862209bef596987c13ab19a89a9fc62018dc2a4e8c9cb927827aadf1c458eee/README.md) |
| 31&#160;Jan&#160;23&#160;18:52&#160;UTC | SHAKEN 709J | 01&#160;May&#160;23&#160;18:52&#160;UTC | false | [view](CERTS/d3bf5a5dad064261173fe2923056c3aba1cb3ce5508fa575f99f589a5d704349/README.md) |
| 31&#160;Jan&#160;23&#160;19:29&#160;UTC | SHAKEN 709J | 01&#160;May&#160;23&#160;05:00&#160;UTC | false | [view](CERTS/d6202ce43b2df4158350d416329cea5815443bc97ab39de634401471b93b8088/README.md) |
| 31&#160;Jan&#160;23&#160;20:36&#160;UTC | SHAKEN 048K | 01&#160;May&#160;23&#160;00:00&#160;UTC | false | [view](CERTS/1fcd880d330e863e0516710789bf4d1b0099ce9381989f639471c1ceedff38b3/README.md) |
| 31&#160;Jan&#160;23&#160;21:06&#160;UTC | SHAKEN 066K | 01&#160;May&#160;23&#160;05:00&#160;UTC | false | [view](CERTS/ab2084daeab358ff89a0d1baaed5727484060f95ce8e155738975479f2121de6/README.md) |
| 31&#160;Jan&#160;23&#160;21:08&#160;UTC | SHAKEN 031K | 01&#160;May&#160;23&#160;05:00&#160;UTC | false | [view](CERTS/e04b56e8c6bb113ad12826c3b3b78402773e973d6f30173f7b5efe5fcc812979/README.md) |
| 31&#160;Jan&#160;23&#160;21:08&#160;UTC | SHAKEN 073K | 01&#160;May&#160;23&#160;06:00&#160;UTC | false | [view](CERTS/a88278c37acb51a17633db4fedf4f1c7619e99aa1a1d7e1e0f7a161943311cd7/README.md) |
| 31&#160;Jan&#160;23&#160;21:11&#160;UTC | SHAKEN 039K | 01&#160;May&#160;23&#160;05:00&#160;UTC | false | [view](CERTS/fdc2d3e8763a11407e6ae4c8f01332ff89a7aa4c816488ba7a2dff60635f2e60/README.md) |
| 31&#160;Jan&#160;23&#160;21:12&#160;UTC | SHAKEN 115K | 01&#160;May&#160;23&#160;04:00&#160;UTC | false | [view](CERTS/1c94a9ae1da39188b3983b464e982f4c57d0c5d47285b9268c3d196899145ff6/README.md) |
| 31&#160;Jan&#160;23&#160;21:12&#160;UTC | SHAKEN 223K | 01&#160;May&#160;23&#160;07:00&#160;UTC | false | [view](CERTS/60240b7a6a3875e160d13641e098741d5bf076d1216e60c5b0dfb769d2add842/README.md) |
| 31&#160;Jan&#160;23&#160;21:13&#160;UTC | SHAKEN 110K | 01&#160;May&#160;23&#160;06:00&#160;UTC | false | [view](CERTS/fcc4e94fa7c8370f7cf64a8699ae51311adfccb6ab9fb35625376ab929499043/README.md) |
| 31&#160;Jan&#160;23&#160;21:15&#160;UTC | SHAKEN 042K | 01&#160;May&#160;23&#160;04:00&#160;UTC | false | [view](CERTS/cfe65f76c82a57cdeb0d9a00370bdddb5d2bb85617b9bae455875158030dda9d/README.md) |
| 31&#160;Jan&#160;23&#160;21:17&#160;UTC | SHAKEN 076K | 01&#160;May&#160;23&#160;04:00&#160;UTC | false | [view](CERTS/74fc503a1da1fc5930cb44284ab9929012b3e88ae3433b11c10b9301252a791b/README.md) |
| 31&#160;Jan&#160;23&#160;21:17&#160;UTC | SHAKEN 148K | 01&#160;May&#160;23&#160;04:00&#160;UTC | false | [view](CERTS/9f1fdb0887f8cd927231e48adcd1ea684f307a06941572f52fde4ec8fe11e3d9/README.md) |
| 31&#160;Jan&#160;23&#160;21:20&#160;UTC | SHAKEN 148K | 01&#160;May&#160;23&#160;04:00&#160;UTC | false | [view](CERTS/74064fd2430363b7f65b09748b86b525b7d66db7fa3d6ffd8305ceb1b2692764/README.md) |
| 31&#160;Jan&#160;23&#160;21:28&#160;UTC | SHAKEN 148K | 01&#160;May&#160;23&#160;04:00&#160;UTC | false | [view](CERTS/5858460e42c64f682f9b076bf6469b5a6c658d4556f50bec70966895e9e9e115/README.md) |
| 02&#160;Mar&#160;23&#160;20:06&#160;UTC | SHAKEN 709J | 31&#160;May&#160;23&#160;20:06&#160;UTC | false | [view](CERTS/63d43f088ad50edf5b3aabe07bc9ee08a1a90a21544159afb80f783dc6d018b6/README.md) |
| 02&#160;Mar&#160;23&#160;20:12&#160;UTC | SHAKEN 709J | 21&#160;Oct&#160;23&#160;17:35&#160;UTC | false | [view](CERTS/99305af5923a22b7d885f5b666f926b69a26825a390e94a318f13ef60e4d465e/README.md) |
| 02&#160;Mar&#160;23&#160;20:43&#160;UTC | SHAKEN 709J | 01&#160;Apr&#160;23&#160;20:43&#160;UTC | false | [view](CERTS/62372e5504a9a82c8171d2531b0ab45ba372b66ecaaa154d18dcbf5f3a9b1ef5/README.md) |
| 03&#160;Mar&#160;23&#160;15:17&#160;UTC | SHAKEN 473K | 02&#160;Mar&#160;24&#160;15:17&#160;UTC | false | [view](CERTS/0eb7d483fd17cb0f1c65b68302dcfe25acbfa599174977835e04fc37784795bb/README.md) |
| 08&#160;Mar&#160;23&#160;19:20&#160;UTC | SHAKEN 709J | 06&#160;Jun&#160;23&#160;19:20&#160;UTC | false | [view](CERTS/f384b0350255145848e221abe01dc77e00ee96e584cd02f60285f88e9bd4ad04/README.md) |
| 08&#160;Mar&#160;23&#160;19:24&#160;UTC | SHAKEN 709J | 05&#160;Jun&#160;23&#160;19:24&#160;UTC | false | [view](CERTS/ef8ab01efc6230b7dba6be2f0b1199651123deab527dfa69b1c352a473d7e3d4/README.md) |
| 08&#160;Mar&#160;23&#160;19:51&#160;UTC | SHAKEN 738J | 18&#160;Mar&#160;23&#160;19:51&#160;UTC | false | [view](CERTS/a9010e12de11c165573464c20ca27ba0f5bf9b5d0867f334061e988a9616e9c0/README.md) |
| 08&#160;Mar&#160;23&#160;20:53&#160;UTC | SHAKEN 683G | 18&#160;Mar&#160;23&#160;20:53&#160;UTC | false | [view](CERTS/588610e4e175bd477e6b941888f0ebd3007505d9df10b27b51e406d416cb2b3b/README.md) |
| 16&#160;Mar&#160;23&#160;18:28&#160;UTC | SHAKEN 683G | 22&#160;Mar&#160;23&#160;19:51&#160;UTC | false | [view](CERTS/72337b27cddbb8a3b2a418362a625ed612e38687f05b74e9a563800400c8294b/README.md) |
| 16&#160;Mar&#160;23&#160;18:34&#160;UTC | SHAKEN 683G | 22&#160;Mar&#160;23&#160;19:51&#160;UTC | false | [view](CERTS/506c6a85818a232882991a1e8a29ee69f5ce88421f75629bd3af78fa7dc7d680/README.md) |
| 16&#160;Mar&#160;23&#160;18:35&#160;UTC | SHAKEN 683G | 22&#160;Mar&#160;23&#160;19:51&#160;UTC | false | [view](CERTS/0a3bb88550d14f1d76c56029a7d322c9ef6a8dfd6aff6483cf7a2b5abbd8554f/README.md) |
| 16&#160;Mar&#160;23&#160;18:35&#160;UTC | SHAKEN 683G | 22&#160;Mar&#160;23&#160;19:51&#160;UTC | false | [view](CERTS/d9eeb91761af4b775b5207bc35d46fef4a8ecac81993188c500ee8dc331fa74a/README.md) |
| 16&#160;Mar&#160;23&#160;18:39&#160;UTC | SHAKEN 683G | 22&#160;Mar&#160;23&#160;19:51&#160;UTC | false | [view](CERTS/350596e7d7fcdfbd3e9a7e2c394f60ba25ccf97ca53471d4cd18314156850cdd/README.md) |
| 16&#160;Mar&#160;23&#160;18:40&#160;UTC | SHAKEN 683G | 22&#160;Mar&#160;23&#160;19:51&#160;UTC | false | [view](CERTS/92e8ddbe070ade9b579615ea03d64b97301ef541c2a98b9a8949485a63a9a5da/README.md) |

#### CA Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 03&#160;May&#160;22&#160;15:31&#160;UTC | Martini Security SHAKEN R1 | 03&#160;May&#160;47&#160;21:31&#160;UTC | false | [view](CERTS/6077e368b0a0e4b6076eaa07ce67d6652ef310c4757776b76af84a6a9e003cdf/README.md) |
| 03&#160;May&#160;22&#160;15:41&#160;UTC | Martini Security SHAKEN G1 | 02&#160;May&#160;27&#160;15:41&#160;UTC | true | [view](CERTS/46e98b1599688c83928b66bcdd49e723ee4207128b5ba4488046edf2879970ef/README.md) |
| 02&#160;Oct&#160;22&#160;10:40&#160;UTC | Martini Security SHAKEN G2 | 01&#160;Oct&#160;27&#160;10:40&#160;UTC | false | [view](CERTS/bf818ddbd3ae492e4a85331b85b52f4d2cdef8287bf910b59e247b6c132fa7fd/README.md) |


Generated: 16 Mar 23 19:18 UTC