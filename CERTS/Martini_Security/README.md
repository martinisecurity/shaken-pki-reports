# STIR/SHAKEN CA Ecosystem Compliance

## Martini Security

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 8 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 0 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 8 certificates being tested against the remaining rules
- 1.00 issues on average found in unexpired, trusted, and non-compliant certificates
- 0.00% of certificates contain one or more Error level issue
- 0.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 147 days is the average remaining validity for the certificates in the corpus
- 124 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days

| Instances | Test | Source |
|-----------|------|--------|
| 4 | [n_atis_certificate_policy_critical](ISSUES/n_atis_certificate_policy_critical/README.md) | ATIS1000080 |

#### CA Certificates

- 3 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 0 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 3 certificates being tested against the remaining rules
- 1.00 issues on average found in unexpired, trusted, and non-compliant certificates
- 0.00% of certificates contain one or more Error level issue
- 0.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 5433 days is the average remaining validity for the certificates in the corpus
- 4261 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days

| Instances | Test | Source |
|-----------|------|--------|
| 1 | [n_atis_ca_certificate_policy_critical](ISSUES/n_atis_ca_certificate_policy_critical/README.md) | ATIS1000080 |

### Details

#### Leaf Certificates

| Created at | Subject | Problems | Link |
|------------|---------|----------|------|
| 28 Sep 22 17:54 UTC | SHAKEN 709J | true | [view](CERTS/7a6b614242beb541f6ce04ce89734a5b601571cd298075c157ef9adc3efcc49c/README.md) |
| 28 Sep 22 18:19 UTC | SHAKEN 709J | true | [view](CERTS/8862209bef596987c13ab19a89a9fc62018dc2a4e8c9cb927827aadf1c458eee/README.md) |
| 28 Sep 22 21:31 UTC | SHAKEN 709J | true | [view](CERTS/200d370a7a5109b711f29fcfe86b70288592ee0d634045c034faba784acb6e23/README.md) |
| 29 Sep 22 18:20 UTC | SHAKEN 073K | true | [view](CERTS/72976965ef94346f0f682ba480111061fac2910dacab06407d65a26db8dd6f06/README.md) |
| 21 Oct 22 17:38 UTC | SHAKEN 709J | false | [view](CERTS/09e45414349c71ce7b7b92101e2de0607ebb989e096ccd861b4e8d37e4e72c2f/README.md) |
| 21 Oct 22 19:21 UTC | SHAKEN 709J | false | [view](CERTS/254c9be0cd95dfe372f0a4117e166da225d468720851f54e3c3edef44480a3d6/README.md) |
| 25 Oct 22 20:57 UTC | SHAKEN 073K | false | [view](CERTS/cf4ee8b8d56521071e2935ac9cb358738902b531589b8e19afcc43cd73c0f01f/README.md) |
| 01 Nov 22 22:01 UTC | SHAKEN 223K | false | [view](CERTS/643004ffd41ca62fdd29de3af56544e8d9aa1fc6194c9160617a050136e69858/README.md) |

#### CA Certificates

| Created at | Subject | Problems | Link |
|------------|---------|----------|------|
| 03 May 22 15:31 UTC | Martini Security SHAKEN R1 | false | [view](CERTS/6077e368b0a0e4b6076eaa07ce67d6652ef310c4757776b76af84a6a9e003cdf/README.md) |
| 03 May 22 15:41 UTC | Martini Security SHAKEN G1 | true | [view](CERTS/46e98b1599688c83928b66bcdd49e723ee4207128b5ba4488046edf2879970ef/README.md) |
| 02 Oct 22 10:40 UTC | Martini Security SHAKEN G2 | false | [view](CERTS/bf818ddbd3ae492e4a85331b85b52f4d2cdef8287bf910b59e247b6c132fa7fd/README.md) |


Generated: 02 Nov 22 15:41 UTC