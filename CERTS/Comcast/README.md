# STIR/SHAKEN CA Ecosystem Compliance

## Comcast

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 92 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 61 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 31 certificates being tested against the remaining rules
- 6.00 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 100.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 30 days is the average remaining validity for the certificates in the corpus
- 30 days is the average initial validity for the certificates in the corpus
- 31 certificates expire in the next 30 days
- 31.00 average number of unexpired certificates per OCN observed
- 1 unique OCNs observed in unexpired and valid certificate corpus

| Instances | Test | Source |
|-----------|------|--------|
| 31 | [e_atis_certificate_policies](ISSUES/e_atis_certificate_policies/README.md) | ATIS1000080 |
| 31 | [e_atis_subject_cn](ISSUES/e_atis_subject_cn/README.md) | ATIS1000080 |
| 31 | [e_atis_subject_key_identifier](ISSUES/e_atis_subject_key_identifier/README.md) | ATIS1000080 |
| 31 | [e_us_cp_ambiguous_identifier](ISSUES/e_us_cp_ambiguous_identifier/README.md) | US_SHAKEN_CP |
| 31 | [e_us_cp_subject_sn](ISSUES/e_us_cp_subject_sn/README.md) | US_SHAKEN_CP |
| 31 | [w_ext_subject_key_identifier_missing_sub_cert](ISSUES/w_ext_subject_key_identifier_missing_sub_cert/README.md) | RFC5280 |

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
- 6867 days is the average remaining validity for the certificates in the corpus
- 7150 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days

No error, warning, or notice level issues were found

### Details

#### Leaf Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 13&#160;Nov&#160;22&#160;11:12&#160;UTC | SHAKEN | 13&#160;Dec&#160;22&#160;11:12&#160;UTC | true | [view](CERTS/74f040438a5061043205b98f00a22ee2c2814ff85ef2d5ab7e7ae1aaf8ca3ebb/README.md) |
| 14&#160;Nov&#160;22&#160;11:12&#160;UTC | SHAKEN | 14&#160;Dec&#160;22&#160;11:12&#160;UTC | true | [view](CERTS/2c2824b8daf6fbde4b27b9266f7b890421ae246c6ec7638e665ea8509b534931/README.md) |
| 15&#160;Nov&#160;22&#160;11:12&#160;UTC | SHAKEN | 15&#160;Dec&#160;22&#160;11:12&#160;UTC | true | [view](CERTS/222bd4b1b8e20bc0713597bc3a2f615a03799176423a62b5484202404d2f1c5f/README.md) |
| 15&#160;Nov&#160;22&#160;21:42&#160;UTC | SHAKEN | 15&#160;Dec&#160;22&#160;21:42&#160;UTC | true | [view](CERTS/2ecfec86f125af457101adfb389f06977c207b1a0ed599423962b2b56c155c32/README.md) |
| 16&#160;Nov&#160;22&#160;11:12&#160;UTC | SHAKEN | 16&#160;Dec&#160;22&#160;11:12&#160;UTC | true | [view](CERTS/d7f89ec7a89359811ba51ea7e5f485468bf5cfa9d4f2d1dbec91acc8eba913f9/README.md) |
| 17&#160;Nov&#160;22&#160;11:12&#160;UTC | SHAKEN | 17&#160;Dec&#160;22&#160;11:12&#160;UTC | true | [view](CERTS/3f701bbc82532d6d519e377113baf63e7f1509a0bc063d93f8a41d0c2a874819/README.md) |
| 18&#160;Nov&#160;22&#160;11:12&#160;UTC | SHAKEN | 18&#160;Dec&#160;22&#160;11:12&#160;UTC | true | [view](CERTS/941bbea2a5f8141a04f09050a1c0970ac42aaf4dff931b73154972159d6bbdd7/README.md) |
| 19&#160;Nov&#160;22&#160;11:12&#160;UTC | SHAKEN | 19&#160;Dec&#160;22&#160;11:12&#160;UTC | true | [view](CERTS/672738adfd2648702ffb5f1a5e024ceeabc6468a2c7978be7e14f3d39652d289/README.md) |
| 20&#160;Nov&#160;22&#160;11:12&#160;UTC | SHAKEN | 20&#160;Dec&#160;22&#160;11:12&#160;UTC | true | [view](CERTS/e77e6751fe93bb57ee66f175493ac702d872365ac83245add394068a3fd5e56b/README.md) |
| 21&#160;Nov&#160;22&#160;11:12&#160;UTC | SHAKEN | 21&#160;Dec&#160;22&#160;11:12&#160;UTC | true | [view](CERTS/0d7c61872461c62b6488c0986012e3fbe60af386df382246bb9643a9f88c03b6/README.md) |
| 22&#160;Nov&#160;22&#160;11:12&#160;UTC | SHAKEN | 22&#160;Dec&#160;22&#160;11:12&#160;UTC | true | [view](CERTS/2f6ca39ec95f8261d33d6b4b3d924f40cdfd1345433f10c2a31a83d5458770e4/README.md) |
| 23&#160;Nov&#160;22&#160;11:12&#160;UTC | SHAKEN | 23&#160;Dec&#160;22&#160;11:12&#160;UTC | true | [view](CERTS/602c9fd6bfc4231002ae884534c56987b8dd66a7afa504670a7631d19b7a07d9/README.md) |
| 24&#160;Nov&#160;22&#160;11:12&#160;UTC | SHAKEN | 24&#160;Dec&#160;22&#160;11:12&#160;UTC | true | [view](CERTS/10bd31097a4a104040283ce19cfb771db731bdcd5fa3ba614bc0633740fe49c8/README.md) |
| 25&#160;Nov&#160;22&#160;11:12&#160;UTC | SHAKEN | 25&#160;Dec&#160;22&#160;11:12&#160;UTC | true | [view](CERTS/04238255c6ec89f4694bf0d49d10cfb3b802201c79f87527bb0f9f6112cb35c9/README.md) |
| 26&#160;Nov&#160;22&#160;11:12&#160;UTC | SHAKEN | 26&#160;Dec&#160;22&#160;11:12&#160;UTC | true | [view](CERTS/122bd732a52eafd21e1a328ed7d94c7dabbe164a3f5c60e62fef8f9e85b5653c/README.md) |
| 27&#160;Nov&#160;22&#160;11:12&#160;UTC | SHAKEN | 27&#160;Dec&#160;22&#160;11:12&#160;UTC | true | [view](CERTS/786b58767652fc5f159d5cc86fc0671426ce45ae8197663c6e369ea9e725129e/README.md) |
| 28&#160;Nov&#160;22&#160;11:12&#160;UTC | SHAKEN | 28&#160;Dec&#160;22&#160;11:12&#160;UTC | true | [view](CERTS/c0715f592e3969eda8bd6311cf412a73173fa9577e2f281c8f161cbbc42ed89f/README.md) |
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

#### CA Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 17&#160;Mar&#160;20&#160;14:45&#160;UTC | Comcast SHAKEN Root CA | 12&#160;Mar&#160;40&#160;14:45&#160;UTC | false | [view](CERTS/b1132c5f12c3ca4d2ff119f2df99544336eb1703512ac99cc42d596e25125bbd/README.md) |
| 06&#160;Apr&#160;20&#160;13:48&#160;UTC | Comcast SHAKEN Intermediate CA | 06&#160;Jun&#160;39&#160;13:48&#160;UTC | false | [view](CERTS/2f3abdfe711377f4d59f31d941962797c61f021c3924af31de99031f9ac54f77/README.md) |


Generated: 12 Dec 22 23:45 UTC