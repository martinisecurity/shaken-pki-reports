# STIR/SHAKEN CA Ecosystem Compliance

## Comcast

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 79 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 51 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 28 certificates being tested against the remaining rules
- 6.00 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 100.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 30 days is the average remaining validity for the certificates in the corpus
- 30 days is the average initial validity for the certificates in the corpus
- 28 certificates expire in the next 30 days
- 14.00 average number of unexpired certificates per OCN observed
- 2 unique OCNs observed in unexpired and valid certificate corpus

| Instances | Test | Source |
|-----------|------|--------|
| 28 | [e_atis_certificate_policies](ISSUES/e_atis_certificate_policies/README.md) | ATIS1000080 |
| 28 | [e_atis_subject_cn](ISSUES/e_atis_subject_cn/README.md) | ATIS1000080 |
| 28 | [e_atis_subject_key_identifier](ISSUES/e_atis_subject_key_identifier/README.md) | ATIS1000080 |
| 28 | [e_us_cp_ambiguous_identifier](ISSUES/e_us_cp_ambiguous_identifier/README.md) | US_SHAKEN_CP |
| 28 | [e_us_cp_subject_sn](ISSUES/e_us_cp_subject_sn/README.md) | US_SHAKEN_CP |
| 28 | [w_ext_subject_key_identifier_missing_sub_cert](ISSUES/w_ext_subject_key_identifier_missing_sub_cert/README.md) | RFC5280 |

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
- 6871 days is the average remaining validity for the certificates in the corpus
- 7150 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days

No error, warning, or notice level issues were found

### Details

#### Leaf Certificates

| Created at | Subject | Problems | Link |
|------------|---------|----------|------|
| 01 Nov 22 11:12 UTC | SHAKEN | true | [view](CERTS/672077a75a58fb6c5d6bb8f067812b721dbad767c682832c2fa646e43ffd36b6/README.md) |
| 01 Nov 22 13:29 UTC | SHAKEN | true | [view](CERTS/a47b02b168a680b892b4ce263af30b19632ac8b8d62b1ef29c055c6f20f4ed01/README.md) |
| 04 Nov 22 11:12 UTC | SHAKEN | true | [view](CERTS/e548f5c32d7bc30c138b95a4f6c1d833f72a348cccedfbc93e51c6016ca40c66/README.md) |
| 05 Nov 22 11:12 UTC | SHAKEN | true | [view](CERTS/fc76258aa402845efe5fc520c3423ae08c34da4b35581008f74b69944989e863/README.md) |
| 06 Nov 22 11:12 UTC | SHAKEN | true | [view](CERTS/65428888f38cc12f0b7f665b53fedbfe1d34e6514267677e776403f7e1a19723/README.md) |
| 07 Nov 22 11:12 UTC | SHAKEN | true | [view](CERTS/cdb68bffe8949ee6681bfe003d1674b4f3b3b41733fd460cadfa147063691083/README.md) |
| 08 Nov 22 11:12 UTC | SHAKEN | true | [view](CERTS/3e9171d58db76444586aeb4cc519fe853c088f7b473cb9aeb5daa24ba63dcbbd/README.md) |
| 09 Nov 22 11:12 UTC | SHAKEN | true | [view](CERTS/9d1362a16150ed27893d2fee57985d5e1a26cd49d88afa28467cdba9f2a630e6/README.md) |
| 10 Nov 22 11:12 UTC | SHAKEN | true | [view](CERTS/ca52606f16d25ae30baaf64ad5fe45985f479ccc9edaf41447212c1ae7f10798/README.md) |
| 11 Nov 22 11:12 UTC | SHAKEN | true | [view](CERTS/8f4edf7b390186bcef76eb1ed3fd648bf478ab8e2bed96d7fdc7a0b7bbe83524/README.md) |
| 13 Nov 22 11:12 UTC | SHAKEN | true | [view](CERTS/74f040438a5061043205b98f00a22ee2c2814ff85ef2d5ab7e7ae1aaf8ca3ebb/README.md) |
| 14 Nov 22 11:12 UTC | SHAKEN | true | [view](CERTS/2c2824b8daf6fbde4b27b9266f7b890421ae246c6ec7638e665ea8509b534931/README.md) |
| 15 Nov 22 11:12 UTC | SHAKEN | true | [view](CERTS/222bd4b1b8e20bc0713597bc3a2f615a03799176423a62b5484202404d2f1c5f/README.md) |
| 15 Nov 22 21:42 UTC | SHAKEN | true | [view](CERTS/2ecfec86f125af457101adfb389f06977c207b1a0ed599423962b2b56c155c32/README.md) |
| 16 Nov 22 11:12 UTC | SHAKEN | true | [view](CERTS/d7f89ec7a89359811ba51ea7e5f485468bf5cfa9d4f2d1dbec91acc8eba913f9/README.md) |
| 17 Nov 22 11:12 UTC | SHAKEN | true | [view](CERTS/3f701bbc82532d6d519e377113baf63e7f1509a0bc063d93f8a41d0c2a874819/README.md) |
| 18 Nov 22 11:12 UTC | SHAKEN | true | [view](CERTS/941bbea2a5f8141a04f09050a1c0970ac42aaf4dff931b73154972159d6bbdd7/README.md) |
| 19 Nov 22 11:12 UTC | SHAKEN | true | [view](CERTS/672738adfd2648702ffb5f1a5e024ceeabc6468a2c7978be7e14f3d39652d289/README.md) |
| 20 Nov 22 11:12 UTC | SHAKEN | true | [view](CERTS/e77e6751fe93bb57ee66f175493ac702d872365ac83245add394068a3fd5e56b/README.md) |
| 21 Nov 22 11:12 UTC | SHAKEN | true | [view](CERTS/0d7c61872461c62b6488c0986012e3fbe60af386df382246bb9643a9f88c03b6/README.md) |
| 22 Nov 22 11:12 UTC | SHAKEN | true | [view](CERTS/2f6ca39ec95f8261d33d6b4b3d924f40cdfd1345433f10c2a31a83d5458770e4/README.md) |
| 23 Nov 22 11:12 UTC | SHAKEN | true | [view](CERTS/602c9fd6bfc4231002ae884534c56987b8dd66a7afa504670a7631d19b7a07d9/README.md) |
| 24 Nov 22 11:12 UTC | SHAKEN | true | [view](CERTS/10bd31097a4a104040283ce19cfb771db731bdcd5fa3ba614bc0633740fe49c8/README.md) |
| 25 Nov 22 11:12 UTC | SHAKEN | true | [view](CERTS/04238255c6ec89f4694bf0d49d10cfb3b802201c79f87527bb0f9f6112cb35c9/README.md) |
| 26 Nov 22 11:12 UTC | SHAKEN | true | [view](CERTS/122bd732a52eafd21e1a328ed7d94c7dabbe164a3f5c60e62fef8f9e85b5653c/README.md) |
| 27 Nov 22 11:12 UTC | SHAKEN | true | [view](CERTS/786b58767652fc5f159d5cc86fc0671426ce45ae8197663c6e369ea9e725129e/README.md) |
| 28 Nov 22 11:12 UTC | SHAKEN | true | [view](CERTS/c0715f592e3969eda8bd6311cf412a73173fa9577e2f281c8f161cbbc42ed89f/README.md) |
| 29 Nov 22 11:12 UTC | SHAKEN | true | [view](CERTS/f3459e563ad32fa87fe02780f443052ecacc98eabb64a0eeebbf262520bb2d48/README.md) |

#### CA Certificates

| Created at | Subject | Problems | Link |
|------------|---------|----------|------|
| 17 Mar 20 14:45 UTC | Comcast SHAKEN Root CA | false | [view](CERTS/b1132c5f12c3ca4d2ff119f2df99544336eb1703512ac99cc42d596e25125bbd/README.md) |
| 06 Apr 20 13:48 UTC | Comcast SHAKEN Intermediate CA | false | [view](CERTS/2f3abdfe711377f4d59f31d941962797c61f021c3924af31de99031f9ac54f77/README.md) |


Generated: 30 Nov 22 17:39 UTC