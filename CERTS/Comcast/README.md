# STIR/SHAKEN CA Ecosystem Compliance

## Comcast

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 60 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 35 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 25 certificates being tested against the remaining rules
- 6.00 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 100.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 29 days is the average remaining validity for the certificates in the corpus
- 30 days is the average initial validity for the certificates in the corpus
- 25 certificates expire in the next 30 days

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
- 6875 days is the average remaining validity for the certificates in the corpus
- 7150 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days

No error, warning, or notice level issues were found

### Details

#### Leaf Certificates

| Created at | Subject | Problems | Link |
|------------|---------|----------|------|
| 19 Oct 22 13:29 UTC | SHAKEN | true | [view](CERTS/735539bc495f743e25fea39ecef9566c17bedaf976cf0057d2bb3569ab448d02/README.md) |
| 20 Oct 22 11:12 UTC | SHAKEN | true | [view](CERTS/e01c63f891eca98375b1583ba29104c72618fbe3109cd7b558c37529af2cb29e/README.md) |
| 20 Oct 22 13:29 UTC | SHAKEN | true | [view](CERTS/56ac7c9635383e869ee9b816df361d00c4f17afe60188f234cc1e2906a913781/README.md) |
| 21 Oct 22 11:12 UTC | SHAKEN | true | [view](CERTS/8fb4a8e814f1c87cf2456b1ecf1e5ce491a5eb399b2d37de0d1a34b4ddb637db/README.md) |
| 21 Oct 22 13:29 UTC | SHAKEN | true | [view](CERTS/fc473a65f1226915a1bddec7b99fa1e53f0e1acc2f4e10106b07b18fc0054b4d/README.md) |
| 22 Oct 22 13:30 UTC | SHAKEN | true | [view](CERTS/8b259dd4c8adee7126f6d77d613a73d8e123a155ea3455a9531f86ef8572aa0e/README.md) |
| 24 Oct 22 11:12 UTC | SHAKEN | true | [view](CERTS/7dd1ed0b11d7d96b811a11562effd14cedc04552b643303507df78f89983e24c/README.md) |
| 25 Oct 22 11:12 UTC | SHAKEN | true | [view](CERTS/5eaa41cb4d2988f9518bc9d580a9c548dd41a6d7e12ba89c919fc302ca3ebaa4/README.md) |
| 26 Oct 22 11:12 UTC | SHAKEN | true | [view](CERTS/753a9492590251d2621c7d0ca1debbbb708d93027dca45ea69f3dfa251819eda/README.md) |
| 27 Oct 22 11:12 UTC | SHAKEN | true | [view](CERTS/4359cd2f3d758f495356d11669974761f5671c7512a2278340d0927a0a390fb4/README.md) |
| 27 Oct 22 13:29 UTC | SHAKEN | true | [view](CERTS/7fb833a2e3672fc5bcabcf05829f5d5336abba7e88577ad444e5179dd2df5ebb/README.md) |
| 28 Oct 22 11:12 UTC | SHAKEN | true | [view](CERTS/67a979476b3e4fb995be4075053f9a805b3df4b18dc09f532c2c239b3dadc13d/README.md) |
| 28 Oct 22 13:29 UTC | SHAKEN | true | [view](CERTS/748774b40fe43fdc9b29363b0bdf107884b3447e2110f8429cc6b893993c1bd2/README.md) |
| 29 Oct 22 13:29 UTC | SHAKEN | true | [view](CERTS/af6785002cd605e314824b90e6e06e3e39bb849078865200001a86382c0dd370/README.md) |
| 31 Oct 22 11:12 UTC | SHAKEN | true | [view](CERTS/cf1f5b80441df2551cfa967ecd684687eb9ea9dd5ce30a4035c8646cdb966c59/README.md) |
| 31 Oct 22 13:30 UTC | SHAKEN | true | [view](CERTS/3d44d0a8d9179445f0ac646cceb3f6a3566f63e2bed4ec08b58a017b063fa36d/README.md) |
| 01 Nov 22 11:12 UTC | SHAKEN | true | [view](CERTS/672077a75a58fb6c5d6bb8f067812b721dbad767c682832c2fa646e43ffd36b6/README.md) |
| 01 Nov 22 13:29 UTC | SHAKEN | true | [view](CERTS/a47b02b168a680b892b4ce263af30b19632ac8b8d62b1ef29c055c6f20f4ed01/README.md) |
| 04 Nov 22 11:12 UTC | SHAKEN | true | [view](CERTS/e548f5c32d7bc30c138b95a4f6c1d833f72a348cccedfbc93e51c6016ca40c66/README.md) |
| 05 Nov 22 11:12 UTC | SHAKEN | true | [view](CERTS/fc76258aa402845efe5fc520c3423ae08c34da4b35581008f74b69944989e863/README.md) |
| 06 Nov 22 11:12 UTC | SHAKEN | true | [view](CERTS/65428888f38cc12f0b7f665b53fedbfe1d34e6514267677e776403f7e1a19723/README.md) |
| 07 Nov 22 11:12 UTC | SHAKEN | true | [view](CERTS/cdb68bffe8949ee6681bfe003d1674b4f3b3b41733fd460cadfa147063691083/README.md) |
| 08 Nov 22 11:12 UTC | SHAKEN | true | [view](CERTS/3e9171d58db76444586aeb4cc519fe853c088f7b473cb9aeb5daa24ba63dcbbd/README.md) |
| 09 Nov 22 11:12 UTC | SHAKEN | true | [view](CERTS/9d1362a16150ed27893d2fee57985d5e1a26cd49d88afa28467cdba9f2a630e6/README.md) |
| 10 Nov 22 11:12 UTC | SHAKEN | true | [view](CERTS/ca52606f16d25ae30baaf64ad5fe45985f479ccc9edaf41447212c1ae7f10798/README.md) |

#### CA Certificates

| Created at | Subject | Problems | Link |
|------------|---------|----------|------|
| 17 Mar 20 14:45 UTC | Comcast SHAKEN Root CA | false | [view](CERTS/b1132c5f12c3ca4d2ff119f2df99544336eb1703512ac99cc42d596e25125bbd/README.md) |
| 06 Apr 20 13:48 UTC | Comcast SHAKEN Intermediate CA | false | [view](CERTS/2f3abdfe711377f4d59f31d941962797c61f021c3924af31de99031f9ac54f77/README.md) |


Generated: 17 Nov 22 19:20 UTC