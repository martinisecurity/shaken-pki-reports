# STIR/SHAKEN CA Ecosystem Compliance

## Comcast

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 101 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 78 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 23 certificates being tested against the remaining rules
- 6.00 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 100.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 29 days is the average remaining validity for the certificates in the corpus
- 30 days is the average initial validity for the certificates in the corpus
- 23 certificates expire in the next 30 days
- 23.00 average number of unexpired certificates per OCN observed
- 1 unique OCNs observed in unexpired and valid certificate corpus

| Instances | Test | Source |
|-----------|------|--------|
| 23 | [e_atis_certificate_policies](ISSUES/e_atis_certificate_policies/README.md) | ATIS1000080 |
| 23 | [e_atis_subject_cn](ISSUES/e_atis_subject_cn/README.md) | ATIS1000080 |
| 23 | [e_atis_subject_key_identifier](ISSUES/e_atis_subject_key_identifier/README.md) | ATIS1000080 |
| 23 | [e_us_cp_ambiguous_identifier](ISSUES/e_us_cp_ambiguous_identifier/README.md) | US_SHAKEN_CP |
| 23 | [e_us_cp_subject_sn](ISSUES/e_us_cp_subject_sn/README.md) | US_SHAKEN_CP |
| 23 | [w_ext_subject_key_identifier_missing_sub_cert](ISSUES/w_ext_subject_key_identifier_missing_sub_cert/README.md) | RFC5280 |

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
- 6861 days is the average remaining validity for the certificates in the corpus
- 7150 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days

No error, warning, or notice level issues were found

### Details

#### Leaf Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 29&#160;Nov&#160;22&#160;11:12&#160;UTC | SHAKEN | 29&#160;Dec&#160;22&#160;11:12&#160;UTC | true | [view](CERTS/f3459e563ad32fa87fe02780f443052ecacc98eabb64a0eeebbf262520bb2d48/README.md) |
| 30&#160;Nov&#160;22&#160;11:12&#160;UTC | SHAKEN | 30&#160;Dec&#160;22&#160;11:12&#160;UTC | true | [view](CERTS/e95f56923916116222f2ab632926f0eac16b07e6467b9645b45b991202dd27b2/README.md) |
| 01&#160;Dec&#160;22&#160;11:12&#160;UTC | SHAKEN | 31&#160;Dec&#160;22&#160;11:12&#160;UTC | true | [view](CERTS/9ef51637cc35b6d4be35e42512e64ba62ceaa58b697a5fbc06efba175ed59850/README.md) |
| 02&#160;Dec&#160;22&#160;11:12&#160;UTC | SHAKEN | 01&#160;Jan&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/b155c82bd48c21b60c160c029a6c909174b5df5120357d08f705c332666161fb/README.md) |
| 03&#160;Dec&#160;22&#160;11:12&#160;UTC | SHAKEN | 02&#160;Jan&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/b95c628e753c0ab8d54f1accab09a8137df508a9971f625fbb9b42560c54eed1/README.md) |
| 04&#160;Dec&#160;22&#160;11:12&#160;UTC | SHAKEN | 03&#160;Jan&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/97d09760d32e3cfb7bdebe613a3b639a7c5972ac35b4470b7fb7c976d153ce06/README.md) |
| 05&#160;Dec&#160;22&#160;11:12&#160;UTC | SHAKEN | 04&#160;Jan&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/6eddf1a1e5f885515b2aff4dac7178318debd013b0b485bc099a1533609d3b78/README.md) |
| 06&#160;Dec&#160;22&#160;11:12&#160;UTC | SHAKEN | 05&#160;Jan&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/401790e35121f99ce136ce4238200ecd4dfe857119d66df6bb456f11cae44dc1/README.md) |
| 07&#160;Dec&#160;22&#160;11:12&#160;UTC | SHAKEN | 06&#160;Jan&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/bdb7cc90224a4e9799e9bedc069143a350d0c121f56257b1409922ed459d3621/README.md) |
| 08&#160;Dec&#160;22&#160;11:12&#160;UTC | SHAKEN | 07&#160;Jan&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/3de3c4421313045688015c5308f738e47f86c82bddf9a0b4accbfd1cb2bcc906/README.md) |
| 09&#160;Dec&#160;22&#160;11:12&#160;UTC | SHAKEN | 08&#160;Jan&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/1cf570715ba60f235d7a6798e5bc2bbc9b5f58b7002340f7023ccebff059a18e/README.md) |
| 10&#160;Dec&#160;22&#160;11:12&#160;UTC | SHAKEN | 09&#160;Jan&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/2b0a5313339250c627ddafb16f10c4b179ab78b410d4fe99a38b0b7ab91890e2/README.md) |
| 11&#160;Dec&#160;22&#160;11:12&#160;UTC | SHAKEN | 10&#160;Jan&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/9fa74ea058b9b342ec1315971e14c65318f78809646e9a20255d2f341be663bc/README.md) |
| 12&#160;Dec&#160;22&#160;11:12&#160;UTC | SHAKEN | 11&#160;Jan&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/21cd33927b518f32fdd147e5259908ea887f1a69eeececb5c3c8a2ff36fca184/README.md) |
| 13&#160;Dec&#160;22&#160;11:12&#160;UTC | SHAKEN | 12&#160;Jan&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/ec20a19dd15adcf3e2f248554be2aaa211a314666566ebff305a7ce8f2177b60/README.md) |
| 14&#160;Dec&#160;22&#160;11:12&#160;UTC | SHAKEN | 13&#160;Jan&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/9af1e9b933422d895f48e7bded49d4cb5a45627270363b404b25460949bbd80a/README.md) |
| 15&#160;Dec&#160;22&#160;11:12&#160;UTC | SHAKEN | 14&#160;Jan&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/c6b0e6e0f16748c7cb79e59d992e4a7f9a5b604db0626c29d087283fc6ea9bd7/README.md) |
| 16&#160;Dec&#160;22&#160;11:12&#160;UTC | SHAKEN | 15&#160;Jan&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/9dd978e65bb964040c395891f6566df6d0c42bce6fb0c59fdfbae4dbb9937be8/README.md) |
| 17&#160;Dec&#160;22&#160;11:12&#160;UTC | SHAKEN | 16&#160;Jan&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/caa1ffa55138d9175e4cc68d4cce4209180b57842ec4503031da17e569dacb1e/README.md) |
| 18&#160;Dec&#160;22&#160;11:12&#160;UTC | SHAKEN | 17&#160;Jan&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/f3c54d0023456edb8ff9770c68c521420c42a39e19c343ed2fa286ed9916f7ff/README.md) |
| 19&#160;Dec&#160;22&#160;11:12&#160;UTC | SHAKEN | 18&#160;Jan&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/3c1feadb782422f8492b83eb4a0fdcbcd230291f94534042646f961f05418cc6/README.md) |
| 20&#160;Dec&#160;22&#160;11:12&#160;UTC | SHAKEN | 19&#160;Jan&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/1c7b199b02aa9141d7cd1818566871a3f270042ce87d44635284e3f7c6207dd4/README.md) |
| 21&#160;Dec&#160;22&#160;11:12&#160;UTC | SHAKEN | 20&#160;Jan&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/babdded586b1f1d3868a0d440c1ae126a27ce77e518c01f7669632081497f9c7/README.md) |

#### CA Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 17&#160;Mar&#160;20&#160;14:45&#160;UTC | Comcast SHAKEN Root CA | 12&#160;Mar&#160;40&#160;14:45&#160;UTC | false | [view](CERTS/b1132c5f12c3ca4d2ff119f2df99544336eb1703512ac99cc42d596e25125bbd/README.md) |
| 06&#160;Apr&#160;20&#160;13:48&#160;UTC | Comcast SHAKEN Intermediate CA | 06&#160;Jun&#160;39&#160;13:48&#160;UTC | false | [view](CERTS/2f3abdfe711377f4d59f31d941962797c61f021c3924af31de99031f9ac54f77/README.md) |


Generated: 29 Dec 22 07:47 UTC