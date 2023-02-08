# STIR/SHAKEN CA Ecosystem Compliance

## Martini Security

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 72 certificates were included in the corpus being tested
- 28 certificates in the corpus were skipped because they are duplicates
- 23 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 21 certificates being tested against the remaining rules
- 1.00 issues on average found in unexpired, trusted, and non-compliant certificates
- 0.00% of certificates contain one or more Error level issue
- 0.00% of certificates contain one or more Warning level issue
- 4.76% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 102 days is the average remaining validity for the certificates in the corpus
- 103 days is the average initial validity for the certificates in the corpus
- 5 certificates expire in the next 30 days
- 1.75 average number of unexpired certificates per OCN observed
- 12 unique OCNs observed in unexpired and valid certificate corpus

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
- 5408 days is the average remaining validity for the certificates in the corpus
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
| 16&#160;Nov&#160;22&#160;17:26&#160;UTC | SHAKEN 048K | 14&#160;Feb&#160;23&#160;00:00&#160;UTC | false | [view](CERTS/2e29a51cd9c7a6d1939de8a4524c509163a3a4125d86636e2b8178157bbeb6f1/README.md) |
| 18&#160;Nov&#160;22&#160;16:55&#160;UTC | SHAKEN 148K | 16&#160;Feb&#160;23&#160;05:00&#160;UTC | false | [view](CERTS/5d69437cbe50c04f6704e2757d505efb74a203c1e4eed608374f4bcd0670e6ac/README.md) |
| 18&#160;Nov&#160;22&#160;18:56&#160;UTC | SHAKEN 115K | 16&#160;Feb&#160;23&#160;05:00&#160;UTC | false | [view](CERTS/874c25749b8deaeb272ce246241e9b3405c050b32adbf79ed922ff23a675de93/README.md) |
| 21&#160;Nov&#160;22&#160;22:30&#160;UTC | SHAKEN 709J | 19&#160;Feb&#160;23&#160;06:00&#160;UTC | false | [view](CERTS/4570376dae8fcd4c2df89429aac988c7c3f6f8eb8bf4de7f7ead3195fec45c29/README.md) |
| 28&#160;Nov&#160;22&#160;18:45&#160;UTC | SHAKEN 110K | 26&#160;Feb&#160;23&#160;07:00&#160;UTC | false | [view](CERTS/43e4066e1aaa9b942e351948de42e8bb478f16aa499c118b03e8793555feb806/README.md) |
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

#### CA Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 03&#160;May&#160;22&#160;15:31&#160;UTC | Martini Security SHAKEN R1 | 03&#160;May&#160;47&#160;21:31&#160;UTC | false | [view](CERTS/6077e368b0a0e4b6076eaa07ce67d6652ef310c4757776b76af84a6a9e003cdf/README.md) |
| 03&#160;May&#160;22&#160;15:41&#160;UTC | Martini Security SHAKEN G1 | 02&#160;May&#160;27&#160;15:41&#160;UTC | true | [view](CERTS/46e98b1599688c83928b66bcdd49e723ee4207128b5ba4488046edf2879970ef/README.md) |
| 02&#160;Oct&#160;22&#160;10:40&#160;UTC | Martini Security SHAKEN G2 | 01&#160;Oct&#160;27&#160;10:40&#160;UTC | false | [view](CERTS/bf818ddbd3ae492e4a85331b85b52f4d2cdef8287bf910b59e247b6c132fa7fd/README.md) |


Generated: 08 Feb 23 19:45 UTC