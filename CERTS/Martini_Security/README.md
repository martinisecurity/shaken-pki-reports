# STIR/SHAKEN CA Ecosystem Compliance

## Martini Security

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 49 certificates were included in the corpus being tested
- 20 certificates in the corpus were skipped because they are duplicates
- 5 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 24 certificates being tested against the remaining rules
- 1.00 issues on average found in unexpired, trusted, and non-compliant certificates
- 0.00% of certificates contain one or more Error level issue
- 0.00% of certificates contain one or more Warning level issue
- 4.17% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 98 days is the average remaining validity for the certificates in the corpus
- 101 days is the average initial validity for the certificates in the corpus
- 18 certificates expire in the next 30 days
- 2.00 average number of unexpired certificates per OCN observed
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
- 5415 days is the average remaining validity for the certificates in the corpus
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
| 21&#160;Oct&#160;22&#160;17:38&#160;UTC | SHAKEN 709J | 19&#160;Jan&#160;23&#160;06:00&#160;UTC | false | [view](CERTS/09e45414349c71ce7b7b92101e2de0607ebb989e096ccd861b4e8d37e4e72c2f/README.md) |
| 21&#160;Oct&#160;22&#160;19:21&#160;UTC | SHAKEN 709J | 19&#160;Jan&#160;23&#160;06:00&#160;UTC | false | [view](CERTS/254c9be0cd95dfe372f0a4117e166da225d468720851f54e3c3edef44480a3d6/README.md) |
| 25&#160;Oct&#160;22&#160;20:57&#160;UTC | SHAKEN 073K | 23&#160;Jan&#160;23&#160;07:00&#160;UTC | false | [view](CERTS/cf4ee8b8d56521071e2935ac9cb358738902b531589b8e19afcc43cd73c0f01f/README.md) |
| 01&#160;Nov&#160;22&#160;22:01&#160;UTC | SHAKEN 223K | 30&#160;Jan&#160;23&#160;08:00&#160;UTC | false | [view](CERTS/643004ffd41ca62fdd29de3af56544e8d9aa1fc6194c9160617a050136e69858/README.md) |
| 02&#160;Nov&#160;22&#160;19:01&#160;UTC | SHAKEN 039K | 31&#160;Jan&#160;23&#160;06:00&#160;UTC | false | [view](CERTS/6ffaa261e05bf2c924bb7e0e7f98cc4059d45d0df695ee04d1b686053b124c09/README.md) |
| 02&#160;Nov&#160;22&#160;19:10&#160;UTC | SHAKEN 148K | 31&#160;Jan&#160;23&#160;05:00&#160;UTC | false | [view](CERTS/954cbc64af18f8489faf382126bf131364409bf8590445893e23f62a4227fd6a/README.md) |
| 02&#160;Nov&#160;22&#160;19:19&#160;UTC | SHAKEN 076K | 31&#160;Jan&#160;23&#160;05:00&#160;UTC | false | [view](CERTS/b863f7e1b6fa414b2bc3d6eb3781467c6550a95855ead4407d758a0dbbc5eb23/README.md) |
| 02&#160;Nov&#160;22&#160;19:37&#160;UTC | SHAKEN 031K | 31&#160;Jan&#160;23&#160;06:00&#160;UTC | false | [view](CERTS/f6631a3e9a2dd3a2482faf48c34167789d1da4618f67dd13979e3332177f4ec2/README.md) |
| 02&#160;Nov&#160;22&#160;19:42&#160;UTC | SHAKEN 148K | 31&#160;Jan&#160;23&#160;05:00&#160;UTC | false | [view](CERTS/493046b32c772c2b13b64ef2d3c24e846c5f0958991e5c7231bf570ea347d54f/README.md) |
| 07&#160;Nov&#160;22&#160;22:48&#160;UTC | SHAKEN 042K | 05&#160;Feb&#160;23&#160;05:00&#160;UTC | false | [view](CERTS/b19d4331d0bc3102e5998c08af429d4f26e9e1885e3bdc7149dd1ce84520d778/README.md) |
| 09&#160;Nov&#160;22&#160;20:20&#160;UTC | SHAKEN 066K | 07&#160;Feb&#160;23&#160;06:00&#160;UTC | false | [view](CERTS/18b89a5058feb112abe94515607c30f7b8f678b48b3fd23ea95326b388d9eb4d/README.md) |
| 09&#160;Nov&#160;22&#160;21:40&#160;UTC | SHAKEN 709J | 07&#160;Feb&#160;23&#160;06:00&#160;UTC | false | [view](CERTS/7a47c5bb90d19cfe67def3c4ad34299c0147253316172894a6fe8b69a3ca7318/README.md) |
| 09&#160;Nov&#160;22&#160;22:16&#160;UTC | SHAKEN 709J | 07&#160;Feb&#160;23&#160;06:00&#160;UTC | false | [view](CERTS/4c55c830a412e53a29e5d57497ccadb00362c84ae1a4ae7af23c5b10ddd662fc/README.md) |
| 09&#160;Nov&#160;22&#160;22:18&#160;UTC | SHAKEN 709J | 07&#160;Feb&#160;23&#160;06:00&#160;UTC | false | [view](CERTS/fe6221f8c6e387bed8d817ee0229bdd379457544e316e2dc11b494dfda4002a7/README.md) |
| 09&#160;Nov&#160;22&#160;22:30&#160;UTC | SHAKEN 709J | 07&#160;Feb&#160;23&#160;06:00&#160;UTC | false | [view](CERTS/eb45b48a2734d4e92468b1e0049e8a38b27471293701129c3f6ec4ce53751693/README.md) |
| 09&#160;Nov&#160;22&#160;22:34&#160;UTC | SHAKEN 709J | 07&#160;Feb&#160;23&#160;06:00&#160;UTC | false | [view](CERTS/8affdf2bce5e7de61e762d9e6599f4318813c44f6a28432e36a7a68aa692b8f5/README.md) |
| 09&#160;Nov&#160;22&#160;22:40&#160;UTC | SHAKEN 709J | 07&#160;Feb&#160;23&#160;06:00&#160;UTC | false | [view](CERTS/a906441173496d0d84b6a5bad4378cd55c34937cfa15acd5325d9839c4f08c87/README.md) |
| 09&#160;Nov&#160;22&#160;22:42&#160;UTC | SHAKEN 709J | 07&#160;Feb&#160;23&#160;06:00&#160;UTC | false | [view](CERTS/ebb954e4df5cce02a29f17cad6692c826b5b32522339424d689625b1ae679d91/README.md) |
| 16&#160;Nov&#160;22&#160;17:26&#160;UTC | SHAKEN 048K | 14&#160;Feb&#160;23&#160;00:00&#160;UTC | false | [view](CERTS/2e29a51cd9c7a6d1939de8a4524c509163a3a4125d86636e2b8178157bbeb6f1/README.md) |
| 18&#160;Nov&#160;22&#160;16:55&#160;UTC | SHAKEN 148K | 16&#160;Feb&#160;23&#160;05:00&#160;UTC | false | [view](CERTS/5d69437cbe50c04f6704e2757d505efb74a203c1e4eed608374f4bcd0670e6ac/README.md) |
| 18&#160;Nov&#160;22&#160;18:56&#160;UTC | SHAKEN 115K | 16&#160;Feb&#160;23&#160;05:00&#160;UTC | false | [view](CERTS/874c25749b8deaeb272ce246241e9b3405c050b32adbf79ed922ff23a675de93/README.md) |
| 21&#160;Nov&#160;22&#160;22:30&#160;UTC | SHAKEN 709J | 19&#160;Feb&#160;23&#160;06:00&#160;UTC | false | [view](CERTS/4570376dae8fcd4c2df89429aac988c7c3f6f8eb8bf4de7f7ead3195fec45c29/README.md) |
| 28&#160;Nov&#160;22&#160;18:45&#160;UTC | SHAKEN 110K | 26&#160;Feb&#160;23&#160;07:00&#160;UTC | false | [view](CERTS/43e4066e1aaa9b942e351948de42e8bb478f16aa499c118b03e8793555feb806/README.md) |

#### CA Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 03&#160;May&#160;22&#160;15:31&#160;UTC | Martini Security SHAKEN R1 | 03&#160;May&#160;47&#160;21:31&#160;UTC | false | [view](CERTS/6077e368b0a0e4b6076eaa07ce67d6652ef310c4757776b76af84a6a9e003cdf/README.md) |
| 03&#160;May&#160;22&#160;15:41&#160;UTC | Martini Security SHAKEN G1 | 02&#160;May&#160;27&#160;15:41&#160;UTC | true | [view](CERTS/46e98b1599688c83928b66bcdd49e723ee4207128b5ba4488046edf2879970ef/README.md) |
| 02&#160;Oct&#160;22&#160;10:40&#160;UTC | Martini Security SHAKEN G2 | 01&#160;Oct&#160;27&#160;10:40&#160;UTC | false | [view](CERTS/bf818ddbd3ae492e4a85331b85b52f4d2cdef8287bf910b59e247b6c132fa7fd/README.md) |


Generated: 11 Jan 23 23:18 UTC