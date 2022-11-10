# STIR/SHAKEN CA Ecosystem Compliance

## Martini Security

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 38 certificates were included in the corpus being tested
- 14 certificates in the corpus were skipped because they are duplicates
- 0 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 24 certificates being tested against the remaining rules
- 1.00 issues on average found in unexpired, trusted, and non-compliant certificates
- 0.00% of certificates contain one or more Error level issue
- 0.00% of certificates contain one or more Warning level issue
- 20.83% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 97 days is the average remaining validity for the certificates in the corpus
- 99 days is the average initial validity for the certificates in the corpus
- 1 certificates expire in the next 30 days

| Instances | Test | Source |
|-----------|------|--------|
| 5 | [n_atis_certificate_policy_critical](ISSUES/n_atis_certificate_policy_critical/README.md) | ATIS1000080 |

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
- 5430 days is the average remaining validity for the certificates in the corpus
- 4261 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days

| Instances | Test | Source |
|-----------|------|--------|
| 1 | [n_atis_ca_certificate_policy_critical](ISSUES/n_atis_ca_certificate_policy_critical/README.md) | ATIS1000080 |

### Details

#### Leaf Certificates

| Created at | Subject | Problems | Link |
|------------|---------|----------|------|
| 27 Sep 22 20:36 UTC | SHAKEN 709J | true | [view](CERTS/7f2f18c940b627651629abb54ca0da9fb3500228f9a95c9292840b764f3ef491/README.md) |
| 28 Sep 22 17:54 UTC | SHAKEN 709J | true | [view](CERTS/7a6b614242beb541f6ce04ce89734a5b601571cd298075c157ef9adc3efcc49c/README.md) |
| 28 Sep 22 18:19 UTC | SHAKEN 709J | true | [view](CERTS/8862209bef596987c13ab19a89a9fc62018dc2a4e8c9cb927827aadf1c458eee/README.md) |
| 28 Sep 22 21:31 UTC | SHAKEN 709J | true | [view](CERTS/200d370a7a5109b711f29fcfe86b70288592ee0d634045c034faba784acb6e23/README.md) |
| 29 Sep 22 18:20 UTC | SHAKEN 073K | true | [view](CERTS/72976965ef94346f0f682ba480111061fac2910dacab06407d65a26db8dd6f06/README.md) |
| 21 Oct 22 17:38 UTC | SHAKEN 709J | false | [view](CERTS/09e45414349c71ce7b7b92101e2de0607ebb989e096ccd861b4e8d37e4e72c2f/README.md) |
| 21 Oct 22 19:21 UTC | SHAKEN 709J | false | [view](CERTS/254c9be0cd95dfe372f0a4117e166da225d468720851f54e3c3edef44480a3d6/README.md) |
| 25 Oct 22 20:57 UTC | SHAKEN 073K | false | [view](CERTS/cf4ee8b8d56521071e2935ac9cb358738902b531589b8e19afcc43cd73c0f01f/README.md) |
| 31 Oct 22 20:37 UTC | SHAKEN 709J | false | [view](CERTS/13e067c23d545b549f88c12775f0c0283a70e8405146ac61972dcfd32787a731/README.md) |
| 01 Nov 22 22:01 UTC | SHAKEN 223K | false | [view](CERTS/643004ffd41ca62fdd29de3af56544e8d9aa1fc6194c9160617a050136e69858/README.md) |
| 02 Nov 22 19:01 UTC | SHAKEN 039K | false | [view](CERTS/6ffaa261e05bf2c924bb7e0e7f98cc4059d45d0df695ee04d1b686053b124c09/README.md) |
| 02 Nov 22 19:10 UTC | SHAKEN 148K | false | [view](CERTS/954cbc64af18f8489faf382126bf131364409bf8590445893e23f62a4227fd6a/README.md) |
| 02 Nov 22 19:19 UTC | SHAKEN 076K | false | [view](CERTS/b863f7e1b6fa414b2bc3d6eb3781467c6550a95855ead4407d758a0dbbc5eb23/README.md) |
| 02 Nov 22 19:37 UTC | SHAKEN 031K | false | [view](CERTS/f6631a3e9a2dd3a2482faf48c34167789d1da4618f67dd13979e3332177f4ec2/README.md) |
| 02 Nov 22 19:42 UTC | SHAKEN 148K | false | [view](CERTS/493046b32c772c2b13b64ef2d3c24e846c5f0958991e5c7231bf570ea347d54f/README.md) |
| 07 Nov 22 22:48 UTC | SHAKEN 042K | false | [view](CERTS/b19d4331d0bc3102e5998c08af429d4f26e9e1885e3bdc7149dd1ce84520d778/README.md) |
| 09 Nov 22 20:20 UTC | SHAKEN 066K | false | [view](CERTS/18b89a5058feb112abe94515607c30f7b8f678b48b3fd23ea95326b388d9eb4d/README.md) |
| 09 Nov 22 21:40 UTC | SHAKEN 709J | false | [view](CERTS/7a47c5bb90d19cfe67def3c4ad34299c0147253316172894a6fe8b69a3ca7318/README.md) |
| 09 Nov 22 22:16 UTC | SHAKEN 709J | false | [view](CERTS/4c55c830a412e53a29e5d57497ccadb00362c84ae1a4ae7af23c5b10ddd662fc/README.md) |
| 09 Nov 22 22:18 UTC | SHAKEN 709J | false | [view](CERTS/fe6221f8c6e387bed8d817ee0229bdd379457544e316e2dc11b494dfda4002a7/README.md) |
| 09 Nov 22 22:30 UTC | SHAKEN 709J | false | [view](CERTS/eb45b48a2734d4e92468b1e0049e8a38b27471293701129c3f6ec4ce53751693/README.md) |
| 09 Nov 22 22:34 UTC | SHAKEN 709J | false | [view](CERTS/8affdf2bce5e7de61e762d9e6599f4318813c44f6a28432e36a7a68aa692b8f5/README.md) |
| 09 Nov 22 22:40 UTC | SHAKEN 709J | false | [view](CERTS/a906441173496d0d84b6a5bad4378cd55c34937cfa15acd5325d9839c4f08c87/README.md) |
| 09 Nov 22 22:42 UTC | SHAKEN 709J | false | [view](CERTS/ebb954e4df5cce02a29f17cad6692c826b5b32522339424d689625b1ae679d91/README.md) |

#### CA Certificates

| Created at | Subject | Problems | Link |
|------------|---------|----------|------|
| 03 May 22 15:31 UTC | Martini Security SHAKEN R1 | false | [view](CERTS/6077e368b0a0e4b6076eaa07ce67d6652ef310c4757776b76af84a6a9e003cdf/README.md) |
| 03 May 22 15:41 UTC | Martini Security SHAKEN G1 | true | [view](CERTS/46e98b1599688c83928b66bcdd49e723ee4207128b5ba4488046edf2879970ef/README.md) |
| 02 Oct 22 10:40 UTC | Martini Security SHAKEN G2 | false | [view](CERTS/bf818ddbd3ae492e4a85331b85b52f4d2cdef8287bf910b59e247b6c132fa7fd/README.md) |


Generated: 10 Nov 22 23:30 UTC