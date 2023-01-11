# STIR/SHAKEN CA Ecosystem Compliance

## Comcast

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 117 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 92 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 25 certificates being tested against the remaining rules
- 6.00 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 100.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 30 days is the average remaining validity for the certificates in the corpus
- 30 days is the average initial validity for the certificates in the corpus
- 25 certificates expire in the next 30 days
- 25.00 average number of unexpired certificates per OCN observed
- 1 unique OCNs observed in unexpired and valid certificate corpus

| Instances | Test | Source |
|-----------|------|--------|
| 25 | [e_atis_certificate_policies](ISSUES/e_atis_certificate_policies/README.md) | ATIS1000080 |
| 25 | [e_atis_subject_cn](ISSUES/e_atis_subject_cn/README.md) | ATIS1000080 |
| 25 | [e_atis_subject_key_identifier](ISSUES/e_atis_subject_key_identifier/README.md) | ATIS1000080 |
| 25 | [e_us_cp_ambiguous_identifier](ISSUES/e_us_cp_ambiguous_identifier/README.md) | US_SHAKEN_CP |
| 25 | [e_us_cp_subject_sn](ISSUES/e_us_cp_subject_sn/README.md) | US_SHAKEN_CP |
| 25 | [w_ext_subject_key_identifier_missing_sub_cert](ISSUES/w_ext_subject_key_identifier_missing_sub_cert/README.md) | RFC5280 |

#### CA Certificates

- 2 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 0 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 2 certificates being tested against the remaining rules
- 0.00 issues on average found in unexpired, trusted, and non-compliant certificates
- 0.00% of certificates contain one or more Error level issue
- 0.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 100.00% of certificates are too old to be assessed against currently enforced expectations
- 6857 days is the average remaining validity for the certificates in the corpus
- 7150 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days

No error, warning, or notice level issues were found

### Details

#### Leaf Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 13&#160;Dec&#160;22&#160;11:12&#160;UTC | SHAKEN | 12&#160;Jan&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/ec20a19dd15adcf3e2f248554be2aaa211a314666566ebff305a7ce8f2177b60/README.md) |
| 14&#160;Dec&#160;22&#160;11:12&#160;UTC | SHAKEN | 13&#160;Jan&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/9af1e9b933422d895f48e7bded49d4cb5a45627270363b404b25460949bbd80a/README.md) |
| 15&#160;Dec&#160;22&#160;11:12&#160;UTC | SHAKEN | 14&#160;Jan&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/c6b0e6e0f16748c7cb79e59d992e4a7f9a5b604db0626c29d087283fc6ea9bd7/README.md) |
| 16&#160;Dec&#160;22&#160;11:12&#160;UTC | SHAKEN | 15&#160;Jan&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/9dd978e65bb964040c395891f6566df6d0c42bce6fb0c59fdfbae4dbb9937be8/README.md) |
| 17&#160;Dec&#160;22&#160;11:12&#160;UTC | SHAKEN | 16&#160;Jan&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/caa1ffa55138d9175e4cc68d4cce4209180b57842ec4503031da17e569dacb1e/README.md) |
| 18&#160;Dec&#160;22&#160;11:12&#160;UTC | SHAKEN | 17&#160;Jan&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/f3c54d0023456edb8ff9770c68c521420c42a39e19c343ed2fa286ed9916f7ff/README.md) |
| 19&#160;Dec&#160;22&#160;11:12&#160;UTC | SHAKEN | 18&#160;Jan&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/3c1feadb782422f8492b83eb4a0fdcbcd230291f94534042646f961f05418cc6/README.md) |
| 20&#160;Dec&#160;22&#160;11:12&#160;UTC | SHAKEN | 19&#160;Jan&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/1c7b199b02aa9141d7cd1818566871a3f270042ce87d44635284e3f7c6207dd4/README.md) |
| 21&#160;Dec&#160;22&#160;11:12&#160;UTC | SHAKEN | 20&#160;Jan&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/babdded586b1f1d3868a0d440c1ae126a27ce77e518c01f7669632081497f9c7/README.md) |
| 22&#160;Dec&#160;22&#160;11:12&#160;UTC | SHAKEN | 21&#160;Jan&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/8a5ac304b6e3a9f79712a464dab5352d52c405abd392c4e43a9a61c43ba7ea52/README.md) |
| 23&#160;Dec&#160;22&#160;11:12&#160;UTC | SHAKEN | 22&#160;Jan&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/89a4fc1d7d8a2a4b27ae4795f57b07d64546f2371f974f6f2c9bf69dcf77da95/README.md) |
| 24&#160;Dec&#160;22&#160;11:12&#160;UTC | SHAKEN | 23&#160;Jan&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/717979f72696355125fbf59fa02f91a40ebc15a1d913e1ac793e75c802f65c6f/README.md) |
| 25&#160;Dec&#160;22&#160;11:12&#160;UTC | SHAKEN | 24&#160;Jan&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/de1ea73d46bf96f00714603919b38804d1177f61fc3a14eb7dbbd3706e1d1872/README.md) |
| 26&#160;Dec&#160;22&#160;11:12&#160;UTC | SHAKEN | 25&#160;Jan&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/6ac462bb6d62f76394804f9da35ca418287bdb23e7e5a9de199ed30461199b86/README.md) |
| 27&#160;Dec&#160;22&#160;11:12&#160;UTC | SHAKEN | 26&#160;Jan&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/c1ce63969a3c9f1c54dd2cc11e722964b70f95548d4dc7dd87c9656e2e885d95/README.md) |
| 28&#160;Dec&#160;22&#160;11:12&#160;UTC | SHAKEN | 27&#160;Jan&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/f96b2daf35b2ef435617f6d54682890046e64ed68fe0be08b6bad4e9456345e9/README.md) |
| 29&#160;Dec&#160;22&#160;11:12&#160;UTC | SHAKEN | 28&#160;Jan&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/181b303cc678bc49aea4b194221ef7ebd6eb913a7edd403eac378abe6973e10e/README.md) |
| 30&#160;Dec&#160;22&#160;11:12&#160;UTC | SHAKEN | 29&#160;Jan&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/ab0353798a61f2d212ffe38c9a33d2fa0c803844aca24da76114b5fa62755a6b/README.md) |
| 31&#160;Dec&#160;22&#160;11:12&#160;UTC | SHAKEN | 30&#160;Jan&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/3138b9e46e3935bf9864546ee6d293467a65ffc9176c83029a10698e213b7df3/README.md) |
| 01&#160;Jan&#160;23&#160;11:12&#160;UTC | SHAKEN | 31&#160;Jan&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/d1856c43a51de1a134f66f52b814e02f74337ad5b333a535998b0c902fea1732/README.md) |
| 02&#160;Jan&#160;23&#160;11:12&#160;UTC | SHAKEN | 01&#160;Feb&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/719acb92673a27a95a56884222e14e9c325ad7e00a496f51ee2ca0b81377c1b6/README.md) |
| 03&#160;Jan&#160;23&#160;11:12&#160;UTC | SHAKEN | 02&#160;Feb&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/ab5ef23f58167423e7ea6cefb7a355d68063da131be72a0464e4dd848866499e/README.md) |
| 04&#160;Jan&#160;23&#160;11:12&#160;UTC | SHAKEN | 03&#160;Feb&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/d03657e109d20dda9ca046a59bd1b6381a611304f9039b7b0bb1e5d547f9d975/README.md) |
| 05&#160;Jan&#160;23&#160;11:12&#160;UTC | SHAKEN | 04&#160;Feb&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/6f49f93c37dfd0a797cbb76e1bb375187bc3dd81e290d71f09062e146ffbc5a5/README.md) |
| 06&#160;Jan&#160;23&#160;11:12&#160;UTC | SHAKEN | 05&#160;Feb&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/6bb26adc661a2c1677f5d179b39730ef80de3915f5e173c8266d9b60859df014/README.md) |

#### CA Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 17&#160;Mar&#160;20&#160;14:45&#160;UTC | Comcast SHAKEN Root CA | 12&#160;Mar&#160;40&#160;14:45&#160;UTC | false | [view](CERTS/b1132c5f12c3ca4d2ff119f2df99544336eb1703512ac99cc42d596e25125bbd/README.md) |
| 06&#160;Apr&#160;20&#160;13:48&#160;UTC | Comcast SHAKEN Intermediate CA | 06&#160;Jun&#160;39&#160;13:48&#160;UTC | false | [view](CERTS/2f3abdfe711377f4d59f31d941962797c61f021c3924af31de99031f9ac54f77/README.md) |


Generated: 11 Jan 23 23:18 UTC