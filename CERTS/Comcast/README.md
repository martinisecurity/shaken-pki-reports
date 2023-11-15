# STIR/SHAKEN CA Ecosystem Compliance

## Comcast

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 58 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 40 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 18 certificates being tested against the remaining rules
- 6.00 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 100.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 29 days is the average remaining validity for the certificates in the corpus
- 30 days is the average initial validity for the certificates in the corpus
- 18 certificates expire in the next 30 days
- 18.00 average number of unexpired certificates per OCN observed
- 1 unique OCNs observed in unexpired and valid certificate corpus

| Instances | Test | Source |
|-----------|------|--------|
| 18 | [e_atis_certificate_policies](ISSUES/e_atis_certificate_policies/README.md) | ATIS1000080 |
| 18 | [e_atis_subject_cn](ISSUES/e_atis_subject_cn/README.md) | ATIS1000080 |
| 18 | [e_atis_subject_key_identifier](ISSUES/e_atis_subject_key_identifier/README.md) | ATIS1000080 |
| 18 | [e_us_cp_ambiguous_identifier](ISSUES/e_us_cp_ambiguous_identifier/README.md) | US_SHAKEN_CP |
| 18 | [e_us_cp_subject_sn](ISSUES/e_us_cp_subject_sn/README.md) | US_SHAKEN_CP |
| 18 | [w_ext_subject_key_identifier_missing_sub_cert](ISSUES/w_ext_subject_key_identifier_missing_sub_cert/README.md) | RFC5280 |

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
- 6754 days is the average remaining validity for the certificates in the corpus
- 7150 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days

No error, warning, or notice level issues were found

### Details

#### Leaf Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 17&#160;Oct&#160;23&#160;02:38&#160;UTC | SHAKEN | 16&#160;Nov&#160;23&#160;02:38&#160;UTC | true | [view](CERTS/7b7d2bf40c221f446888d1f24d6d38660423f5cca016ffa57de958060515c055/README.md) |
| 18&#160;Oct&#160;23&#160;02:38&#160;UTC | SHAKEN | 17&#160;Nov&#160;23&#160;02:38&#160;UTC | true | [view](CERTS/cd1098e3f4cd30c05dc09c47fe1d6191b70a0995d6f6223884e938f9b46fa59b/README.md) |
| 19&#160;Oct&#160;23&#160;02:38&#160;UTC | SHAKEN | 18&#160;Nov&#160;23&#160;02:38&#160;UTC | true | [view](CERTS/ca67428e6ebc19f49888f6f0de884ca5a7653258ea9afb969de1094bdf7630e0/README.md) |
| 20&#160;Oct&#160;23&#160;02:39&#160;UTC | SHAKEN | 19&#160;Nov&#160;23&#160;02:39&#160;UTC | true | [view](CERTS/edd0e45179e78dc22eb2bc98ad9fb3f5299a1a3a2fa9524dbb7f74601fd9dce3/README.md) |
| 21&#160;Oct&#160;23&#160;02:38&#160;UTC | SHAKEN | 20&#160;Nov&#160;23&#160;02:38&#160;UTC | true | [view](CERTS/1b198c3c61f44c99416f994914a8744329acbe07322b9c15e0caa5e4279a7e68/README.md) |
| 22&#160;Oct&#160;23&#160;02:39&#160;UTC | SHAKEN | 21&#160;Nov&#160;23&#160;02:39&#160;UTC | true | [view](CERTS/d7668eb47d684508189db2f5613d8c17059699ca844db1898708de3953cf6673/README.md) |
| 23&#160;Oct&#160;23&#160;02:38&#160;UTC | SHAKEN | 22&#160;Nov&#160;23&#160;02:38&#160;UTC | true | [view](CERTS/212b0797178fa925cf3bac2f025a1d4c6ed6d5bfe38bda2876331e00a456dd44/README.md) |
| 24&#160;Oct&#160;23&#160;02:38&#160;UTC | SHAKEN | 23&#160;Nov&#160;23&#160;02:38&#160;UTC | true | [view](CERTS/27138bc9434b23cc66c21ed9ca2108eda200576abfae2dc78705b0d92684f111/README.md) |
| 25&#160;Oct&#160;23&#160;02:38&#160;UTC | SHAKEN | 24&#160;Nov&#160;23&#160;02:38&#160;UTC | true | [view](CERTS/feef689bf1cc8679ba98ba192848f7a66d498d98f5a95ef17ceea88b468d0152/README.md) |
| 26&#160;Oct&#160;23&#160;02:39&#160;UTC | SHAKEN | 25&#160;Nov&#160;23&#160;02:39&#160;UTC | true | [view](CERTS/8da7d29c58b5bec64c1c66fdc95b41d77afb06a22f1a91366e3e10341d250558/README.md) |
| 27&#160;Oct&#160;23&#160;02:38&#160;UTC | SHAKEN | 26&#160;Nov&#160;23&#160;02:38&#160;UTC | true | [view](CERTS/35aa06e8a904ac94a056685aab8c313c7c14c2588fc98345dc12fcb7ea2966b1/README.md) |
| 30&#160;Oct&#160;23&#160;02:38&#160;UTC | SHAKEN | 29&#160;Nov&#160;23&#160;02:38&#160;UTC | true | [view](CERTS/b925a07389b773443ddfaf032fe8f36349aa8fa2fb9b8f5ef57c7d1823352c08/README.md) |
| 31&#160;Oct&#160;23&#160;02:38&#160;UTC | SHAKEN | 30&#160;Nov&#160;23&#160;02:38&#160;UTC | true | [view](CERTS/f743f0172b23343ed0e7f5ec154bbafcb9ae57ae86562aeaad84cbf0607d11ed/README.md) |
| 01&#160;Nov&#160;23&#160;02:38&#160;UTC | SHAKEN | 01&#160;Dec&#160;23&#160;02:38&#160;UTC | true | [view](CERTS/160e7e9069366204de585183cc7d551f387a86f3f24159c7c5ae60775c3b7642/README.md) |
| 02&#160;Nov&#160;23&#160;02:39&#160;UTC | SHAKEN | 02&#160;Dec&#160;23&#160;02:39&#160;UTC | true | [view](CERTS/e9b00ef140c2d03210099c3a46cbbea4a81864dfd75523aa7b048fb55036bafb/README.md) |
| 03&#160;Nov&#160;23&#160;02:38&#160;UTC | SHAKEN | 03&#160;Dec&#160;23&#160;02:38&#160;UTC | true | [view](CERTS/fbdf6f212f5d3f401e5a2d0f1beda961b5e4092b682aff97f7f6899df1a260ed/README.md) |
| 06&#160;Nov&#160;23&#160;02:39&#160;UTC | SHAKEN | 06&#160;Dec&#160;23&#160;02:39&#160;UTC | true | [view](CERTS/20bd9caf838f46d12dbfed0e7556f753331413e9c7672896a1f9af19760bbb50/README.md) |
| 07&#160;Nov&#160;23&#160;02:38&#160;UTC | SHAKEN | 07&#160;Dec&#160;23&#160;02:38&#160;UTC | true | [view](CERTS/dfb5e692fde572d2e39724b7d0a84e452b33e57b1629b50ae66d6014b93a2d78/README.md) |

#### CA Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 17&#160;Mar&#160;20&#160;14:45&#160;UTC | Comcast SHAKEN Root CA | 12&#160;Mar&#160;40&#160;14:45&#160;UTC | false | [view](CERTS/b1132c5f12c3ca4d2ff119f2df99544336eb1703512ac99cc42d596e25125bbd/README.md) |
| 06&#160;Apr&#160;20&#160;13:48&#160;UTC | Comcast SHAKEN Intermediate CA | 06&#160;Jun&#160;39&#160;13:48&#160;UTC | false | [view](CERTS/2f3abdfe711377f4d59f31d941962797c61f021c3924af31de99031f9ac54f77/README.md) |


Generated: 15 Nov 23 17:17 UTC