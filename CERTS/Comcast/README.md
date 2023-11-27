# STIR/SHAKEN CA Ecosystem Compliance

## Comcast

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 157 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 137 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 20 certificates being tested against the remaining rules
- 7.00 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 100.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 29 days is the average remaining validity for the certificates in the corpus
- 30 days is the average initial validity for the certificates in the corpus
- 20 certificates expire in the next 30 days
- 20.00 average number of unexpired certificates per OCN observed
- 1 unique OCNs observed in unexpired and valid certificate corpus

| Instances | Test | Source |
|-----------|------|--------|
| 20 | [e_atis_ext_certificate_policies](ISSUES/e_atis_ext_certificate_policies/README.md) | ATIS1000080 |
| 20 | [e_atis_serial_number_size](ISSUES/e_atis_serial_number_size/README.md) | ATIS1000080 |
| 20 | [e_atis_subject_cn_spc](ISSUES/e_atis_subject_cn_spc/README.md) | ATIS1000080 |
| 20 | [e_atis_subject_key_identifier](ISSUES/e_atis_subject_key_identifier/README.md) | ATIS1000080 |
| 20 | [e_atis_subject_key_identifier_size](ISSUES/e_atis_subject_key_identifier_size/README.md) | ATIS1000080 |
| 20 | [e_shaken_certificate_policies_id](ISSUES/e_shaken_certificate_policies_id/README.md) | US_SHAKEN_CP |
| 20 | [w_ext_subject_key_identifier_missing_sub_cert](ISSUES/w_ext_subject_key_identifier_missing_sub_cert/README.md) | RFC5280 |

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
- 6750 days is the average remaining validity for the certificates in the corpus
- 7150 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days

No error, warning, or notice level issues were found

### Details

#### Leaf Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 30&#160;Oct&#160;23&#160;02:38&#160;UTC | SHAKEN | 29&#160;Nov&#160;23&#160;02:38&#160;UTC | true | [view](CERTS/b925a07389b773443ddfaf032fe8f36349aa8fa2fb9b8f5ef57c7d1823352c08/README.md) |
| 31&#160;Oct&#160;23&#160;02:38&#160;UTC | SHAKEN | 30&#160;Nov&#160;23&#160;02:38&#160;UTC | true | [view](CERTS/f743f0172b23343ed0e7f5ec154bbafcb9ae57ae86562aeaad84cbf0607d11ed/README.md) |
| 01&#160;Nov&#160;23&#160;02:38&#160;UTC | SHAKEN | 01&#160;Dec&#160;23&#160;02:38&#160;UTC | true | [view](CERTS/160e7e9069366204de585183cc7d551f387a86f3f24159c7c5ae60775c3b7642/README.md) |
| 02&#160;Nov&#160;23&#160;02:39&#160;UTC | SHAKEN | 02&#160;Dec&#160;23&#160;02:39&#160;UTC | true | [view](CERTS/e9b00ef140c2d03210099c3a46cbbea4a81864dfd75523aa7b048fb55036bafb/README.md) |
| 03&#160;Nov&#160;23&#160;02:38&#160;UTC | SHAKEN | 03&#160;Dec&#160;23&#160;02:38&#160;UTC | true | [view](CERTS/fbdf6f212f5d3f401e5a2d0f1beda961b5e4092b682aff97f7f6899df1a260ed/README.md) |
| 06&#160;Nov&#160;23&#160;02:39&#160;UTC | SHAKEN | 06&#160;Dec&#160;23&#160;02:39&#160;UTC | true | [view](CERTS/20bd9caf838f46d12dbfed0e7556f753331413e9c7672896a1f9af19760bbb50/README.md) |
| 07&#160;Nov&#160;23&#160;02:38&#160;UTC | SHAKEN | 07&#160;Dec&#160;23&#160;02:38&#160;UTC | true | [view](CERTS/dfb5e692fde572d2e39724b7d0a84e452b33e57b1629b50ae66d6014b93a2d78/README.md) |
| 08&#160;Nov&#160;23&#160;02:38&#160;UTC | SHAKEN | 08&#160;Dec&#160;23&#160;02:38&#160;UTC | true | [view](CERTS/314a7c0f98b9112ab7fe95ebc1c6fdd5ca94e2779708d544bbc871f86972a1ac/README.md) |
| 09&#160;Nov&#160;23&#160;02:38&#160;UTC | SHAKEN | 09&#160;Dec&#160;23&#160;02:38&#160;UTC | true | [view](CERTS/5d7c66ebdd6e98721511ef1fedd48a357bdbd57ac282df580c86f2a22aca49b7/README.md) |
| 10&#160;Nov&#160;23&#160;02:39&#160;UTC | SHAKEN | 10&#160;Dec&#160;23&#160;02:39&#160;UTC | true | [view](CERTS/fba6c3ce319e1a360d1bfaf310ecb05f3a6ea78af40986c85a8eb0b99d9e11ae/README.md) |
| 11&#160;Nov&#160;23&#160;02:38&#160;UTC | SHAKEN | 11&#160;Dec&#160;23&#160;02:38&#160;UTC | true | [view](CERTS/0c9ba6095a8184fbaf5cf7e9b22f7b1840a071b570e30ba2d2cfb9cd33c96577/README.md) |
| 12&#160;Nov&#160;23&#160;02:39&#160;UTC | SHAKEN | 12&#160;Dec&#160;23&#160;02:39&#160;UTC | true | [view](CERTS/d4beb43e23da01746d80c8d05f10503de5947b6bb72fac502ddee5a47a2b5d23/README.md) |
| 13&#160;Nov&#160;23&#160;02:38&#160;UTC | SHAKEN | 13&#160;Dec&#160;23&#160;02:38&#160;UTC | true | [view](CERTS/0268cf16a382d8ad7faa4d0484583fd4b1d04d5a2c4957c853f7a6ff80a6e8d0/README.md) |
| 14&#160;Nov&#160;23&#160;02:39&#160;UTC | SHAKEN | 14&#160;Dec&#160;23&#160;02:39&#160;UTC | true | [view](CERTS/7dd1a2da240e0bf0891878831ca3a68687ce956ddf7541a4fd69b131c4e1fdf1/README.md) |
| 15&#160;Nov&#160;23&#160;02:38&#160;UTC | SHAKEN | 15&#160;Dec&#160;23&#160;02:38&#160;UTC | true | [view](CERTS/7752d4df6a574e73e764cf1a692db8155b579210455ed6ab1faa6992a5fd3d95/README.md) |
| 16&#160;Nov&#160;23&#160;02:39&#160;UTC | SHAKEN | 16&#160;Dec&#160;23&#160;02:39&#160;UTC | true | [view](CERTS/d4384fb980d32402ed4b38df254a5a589a72a72f9ddd070a3114fd6acf3570a4/README.md) |
| 17&#160;Nov&#160;23&#160;02:38&#160;UTC | SHAKEN | 17&#160;Dec&#160;23&#160;02:38&#160;UTC | true | [view](CERTS/0ccbc22db458b7fe8172ce42fd2aabed02abf5d368c8a5bc0fc49086708af179/README.md) |
| 18&#160;Nov&#160;23&#160;02:39&#160;UTC | SHAKEN | 18&#160;Dec&#160;23&#160;02:39&#160;UTC | true | [view](CERTS/28c32051ab4b5da1dca4ed57432c8570cfc535349e29fae3db77e3128f14c9de/README.md) |
| 20&#160;Nov&#160;23&#160;02:38&#160;UTC | SHAKEN | 20&#160;Dec&#160;23&#160;02:38&#160;UTC | true | [view](CERTS/5cf26f20d4db827f0455dc1334bcdee6befa22506530b49a1b45ca1505a37119/README.md) |
| 21&#160;Nov&#160;23&#160;02:38&#160;UTC | SHAKEN | 21&#160;Dec&#160;23&#160;02:38&#160;UTC | true | [view](CERTS/697b8cd776e16f2d218e9bee8790521b231a8648f234d4c5a11dcae896b42e45/README.md) |

#### CA Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 17&#160;Mar&#160;20&#160;14:45&#160;UTC | Comcast SHAKEN Root CA | 12&#160;Mar&#160;40&#160;14:45&#160;UTC | false | [view](CERTS/b1132c5f12c3ca4d2ff119f2df99544336eb1703512ac99cc42d596e25125bbd/README.md) |
| 06&#160;Apr&#160;20&#160;13:48&#160;UTC | Comcast SHAKEN Intermediate CA | 06&#160;Jun&#160;39&#160;13:48&#160;UTC | false | [view](CERTS/2f3abdfe711377f4d59f31d941962797c61f021c3924af31de99031f9ac54f77/README.md) |


Generated: 27 Nov 23 23:28 UTC