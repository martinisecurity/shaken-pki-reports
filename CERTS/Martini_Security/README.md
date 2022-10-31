# STIR/SHAKEN CA Ecosystem Compliance

## Martini Security

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

7 certificates were included in the corpus being tested\
0 certificates in the corpus were skipped because they were expired\
0 certificates in the corpus were skipped because they are not currently trusted\
0.00% of certificates contain one or more Error level issue\
0.00% of certificates contain one or more Warning level issue\
57.14% of certificates contain one or more Notice level issue\
0.00% of certificates are too old to be assessed against currently enforced expectations\
155 days is the average remaining validity for the certificates in the corpus\
129 days is the average initial validity for the certificates in the corpus\
0 certificates expire in the next 30 days

| Instances | Test Status | Source |
|-----------|-------------|--------|
| 1 | [n_sti_certificate_policy_critical](ISSUES/n_sti_certificate_policy_critical/README.md) | ATIS-1000080 |

#### CA Certificates

3 certificates were included in the corpus being tested\
0 certificates in the corpus were skipped because they were expired\
0 certificates in the corpus were skipped because they are not currently trusted\
0.00% of certificates contain one or more Error level issue\
100.00% of certificates contain one or more Warning level issue\
33.33% of certificates contain one or more Notice level issue\
0.00% of certificates are too old to be assessed against currently enforced expectations\
5433 days is the average remaining validity for the certificates in the corpus\
4261 days is the average initial validity for the certificates in the corpus\
0 certificates expire in the next 30 days

| Instances | Test Status | Source |
|-----------|-------------|--------|
| 1 | [n_sti_ca_certificate_policy_critical](ISSUES/n_sti_ca_certificate_policy_critical/README.md) | ATIS-1000080 |
| 1 | [w_pki_ca_subject_rdn_unknown](ISSUES/w_pki_ca_subject_rdn_unknown/README.md) | SHAKEN PKI Best Practice |

### Details

#### Leaf Certificates

| Created at | Subject | Problems | Link |
|------------|---------|----------|------|
| 25 Oct 22 20:57 UTC | SHAKEN 073K | false | [view](CERTS/cf4ee8b8d56521071e2935ac9cb358738902b531589b8e19afcc43cd73c0f01f/README.md) |
| 21 Oct 22 19:21 UTC | SHAKEN 709J | false | [view](CERTS/254c9be0cd95dfe372f0a4117e166da225d468720851f54e3c3edef44480a3d6/README.md) |
| 21 Oct 22 17:38 UTC | SHAKEN 709J | false | [view](CERTS/09e45414349c71ce7b7b92101e2de0607ebb989e096ccd861b4e8d37e4e72c2f/README.md) |
| 28 Sep 22 17:54 UTC | SHAKEN 709J | true | [view](CERTS/7a6b614242beb541f6ce04ce89734a5b601571cd298075c157ef9adc3efcc49c/README.md) |
| 29 Sep 22 18:20 UTC | SHAKEN 073K | true | [view](CERTS/72976965ef94346f0f682ba480111061fac2910dacab06407d65a26db8dd6f06/README.md) |
| 28 Sep 22 21:31 UTC | SHAKEN 709J | true | [view](CERTS/200d370a7a5109b711f29fcfe86b70288592ee0d634045c034faba784acb6e23/README.md) |
| 28 Sep 22 18:19 UTC | SHAKEN 709J | true | [view](CERTS/8862209bef596987c13ab19a89a9fc62018dc2a4e8c9cb927827aadf1c458eee/README.md) |

#### CA Certificates

| Created at | Subject | Problems | Link |
|------------|---------|----------|------|
| 02 Oct 22 10:40 UTC | Martini Security SHAKEN G2 | true | [view](CERTS/bf818ddbd3ae492e4a85331b85b52f4d2cdef8287bf910b59e247b6c132fa7fd/README.md) |
| 03 May 22 15:41 UTC | Martini Security SHAKEN G1 | true | [view](CERTS/46e98b1599688c83928b66bcdd49e723ee4207128b5ba4488046edf2879970ef/README.md) |
| 03 May 22 15:31 UTC | Martini Security SHAKEN R1 | true | [view](CERTS/6077e368b0a0e4b6076eaa07ce67d6652ef310c4757776b76af84a6a9e003cdf/README.md) |


Generated: 31/10/2022 at 16:43:22