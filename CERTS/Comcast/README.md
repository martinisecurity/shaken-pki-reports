# STIR/SHAKEN CA Ecosystem Compliance

## Comcast

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 51 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 9 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 42 certificates being tested against the remaining rules
- 6.00 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 100.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 29 days is the average remaining validity for the certificates in the corpus
- 30 days is the average initial validity for the certificates in the corpus
- 42 certificates expire in the next 30 days
- 42.00 average number of unexpired certificates per OCN observed
- 1 unique OCNs observed in unexpired and valid certificate corpus

| Instances | Test | Source |
|-----------|------|--------|
| 42 | [e_atis_certificate_policies](ISSUES/e_atis_certificate_policies/README.md) | ATIS1000080 |
| 42 | [e_atis_subject_cn](ISSUES/e_atis_subject_cn/README.md) | ATIS1000080 |
| 42 | [e_atis_subject_key_identifier](ISSUES/e_atis_subject_key_identifier/README.md) | ATIS1000080 |
| 42 | [e_us_cp_ambiguous_identifier](ISSUES/e_us_cp_ambiguous_identifier/README.md) | US_SHAKEN_CP |
| 42 | [e_us_cp_subject_sn](ISSUES/e_us_cp_subject_sn/README.md) | US_SHAKEN_CP |
| 42 | [w_ext_subject_key_identifier_missing_sub_cert](ISSUES/w_ext_subject_key_identifier_missing_sub_cert/README.md) | RFC5280 |

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
- 6821 days is the average remaining validity for the certificates in the corpus
- 7150 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days

No error, warning, or notice level issues were found

### Details

#### Leaf Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 29&#160;Mar&#160;23&#160;10:11&#160;UTC | SHAKEN | 28&#160;Apr&#160;23&#160;10:11&#160;UTC | true | [view](CERTS/c2f7c9f91a55f1e606dfd7486c03c2cb2f16b1583f7e3a80035c0b394d02e808/README.md) |
| 30&#160;Mar&#160;23&#160;10:11&#160;UTC | SHAKEN | 29&#160;Apr&#160;23&#160;10:11&#160;UTC | true | [view](CERTS/46ea8ccb0b1cc8bc6cc9e6871567176f9c07a35309d27906c1a4e5c5813e916d/README.md) |
| 30&#160;Mar&#160;23&#160;12:30&#160;UTC | SHAKEN | 29&#160;Apr&#160;23&#160;12:30&#160;UTC | true | [view](CERTS/a0c4a8c8043a0ccc5441d8419a1bb805cb027fbebf102647539ad029b7057191/README.md) |
| 31&#160;Mar&#160;23&#160;10:11&#160;UTC | SHAKEN | 30&#160;Apr&#160;23&#160;10:11&#160;UTC | true | [view](CERTS/a63118b262f25fac9155350303d1531a0f0df27d5c994d92be9238d970d5be5d/README.md) |
| 31&#160;Mar&#160;23&#160;12:30&#160;UTC | SHAKEN | 30&#160;Apr&#160;23&#160;12:30&#160;UTC | true | [view](CERTS/382849805b2fdfc94792bf56c6597f13c629f668e1dc2e64d83fbe3d73726764/README.md) |
| 01&#160;Apr&#160;23&#160;10:11&#160;UTC | SHAKEN | 01&#160;May&#160;23&#160;10:11&#160;UTC | true | [view](CERTS/d1ecd1ce59c62b30e18951599f74febd321b17147b5503ab54cea56bc6867b50/README.md) |
| 01&#160;Apr&#160;23&#160;12:30&#160;UTC | SHAKEN | 01&#160;May&#160;23&#160;12:30&#160;UTC | true | [view](CERTS/94708bf7d991feb3623085170ddcca56efc4f2cf7813304ee9de3997cfcfdc14/README.md) |
| 02&#160;Apr&#160;23&#160;12:30&#160;UTC | SHAKEN | 02&#160;May&#160;23&#160;12:30&#160;UTC | true | [view](CERTS/34c9cd2534d037b19a17e7b54097ce88fbec9731810c20c55cee01af95b9e94f/README.md) |
| 03&#160;Apr&#160;23&#160;10:11&#160;UTC | SHAKEN | 03&#160;May&#160;23&#160;10:11&#160;UTC | true | [view](CERTS/131e5f7abf6546780e45e66841ef3c398cdd09222e5180803a2c695d144d7676/README.md) |
| 03&#160;Apr&#160;23&#160;12:30&#160;UTC | SHAKEN | 03&#160;May&#160;23&#160;12:30&#160;UTC | true | [view](CERTS/3b21f20bcd5f3f46c90aa5513d78a0989bf70c2807e1a392d4873b47d70e9747/README.md) |
| 04&#160;Apr&#160;23&#160;12:30&#160;UTC | SHAKEN | 04&#160;May&#160;23&#160;12:30&#160;UTC | true | [view](CERTS/677468b77b8c985d593b7b3b62cfec8782ce93875d0a42fb22ef334ff8a2c54d/README.md) |
| 05&#160;Apr&#160;23&#160;10:11&#160;UTC | SHAKEN | 05&#160;May&#160;23&#160;10:11&#160;UTC | true | [view](CERTS/222f878adec6208bd161e15b40cbb1377ec3f073fdc5816097209624ebcdda77/README.md) |
| 05&#160;Apr&#160;23&#160;12:30&#160;UTC | SHAKEN | 05&#160;May&#160;23&#160;12:30&#160;UTC | true | [view](CERTS/08ee4c981eb5a4df63b42efb281d85e908a911cce0fd632a15d247ff3d92cb1b/README.md) |
| 06&#160;Apr&#160;23&#160;10:11&#160;UTC | SHAKEN | 06&#160;May&#160;23&#160;10:11&#160;UTC | true | [view](CERTS/7eb92cf1ed3c073430d5ff5dac808d93f5b8c58d30928d6a89b9514f6f9c55c9/README.md) |
| 06&#160;Apr&#160;23&#160;12:30&#160;UTC | SHAKEN | 06&#160;May&#160;23&#160;12:30&#160;UTC | true | [view](CERTS/dab739013d1e4ac82ea146e9de599ac8143f5e639f8811e54d8d7712ead76e4a/README.md) |
| 07&#160;Apr&#160;23&#160;10:11&#160;UTC | SHAKEN | 07&#160;May&#160;23&#160;10:11&#160;UTC | true | [view](CERTS/6494734c67cba2a4ec417fbf45b7f5e916aa98ef8d6b2b251252d58642d313a9/README.md) |
| 07&#160;Apr&#160;23&#160;12:30&#160;UTC | SHAKEN | 07&#160;May&#160;23&#160;12:30&#160;UTC | true | [view](CERTS/973799927e179f39fb7647e146319f7715ac46b2c77c072a2d248d52e74cacc8/README.md) |
| 08&#160;Apr&#160;23&#160;10:11&#160;UTC | SHAKEN | 08&#160;May&#160;23&#160;10:11&#160;UTC | true | [view](CERTS/b5fe0382c4f26dd9c0148a3dae96525e44224a5d226651ff337da44cc545cdef/README.md) |
| 08&#160;Apr&#160;23&#160;12:30&#160;UTC | SHAKEN | 08&#160;May&#160;23&#160;12:30&#160;UTC | true | [view](CERTS/23d76d0fd8d8c0971b09c4aebadd0f010afbeb46bbe06e6cb5705fe8523fb496/README.md) |
| 09&#160;Apr&#160;23&#160;10:11&#160;UTC | SHAKEN | 09&#160;May&#160;23&#160;10:11&#160;UTC | true | [view](CERTS/c571fecc3c9ac20a5defee2f61a940307891db55df7086654e0de7ede243ae5e/README.md) |
| 09&#160;Apr&#160;23&#160;12:30&#160;UTC | SHAKEN | 09&#160;May&#160;23&#160;12:30&#160;UTC | true | [view](CERTS/2bdb3ed0455f1843176e5da0957e6c0032c411aa0cbab632a6cb4e73e590b063/README.md) |
| 10&#160;Apr&#160;23&#160;10:11&#160;UTC | SHAKEN | 10&#160;May&#160;23&#160;10:11&#160;UTC | true | [view](CERTS/123a8deff74ac503647253d50abc355b8bf977a6194a588526ddd16c38833685/README.md) |
| 10&#160;Apr&#160;23&#160;12:30&#160;UTC | SHAKEN | 10&#160;May&#160;23&#160;12:30&#160;UTC | true | [view](CERTS/7691fca146741a402eb0572d69202e8f516dfa2ce22d28463b6ecae514c4db1b/README.md) |
| 11&#160;Apr&#160;23&#160;10:11&#160;UTC | SHAKEN | 11&#160;May&#160;23&#160;10:11&#160;UTC | true | [view](CERTS/77f720a88a81237803ee826ac5ed137a65d377aeb72307e70c087418abc74b25/README.md) |
| 11&#160;Apr&#160;23&#160;12:30&#160;UTC | SHAKEN | 11&#160;May&#160;23&#160;12:30&#160;UTC | true | [view](CERTS/fa80cc45d72970657ecf731a60c80f62451aa74201de31d8cfea33762583d3c3/README.md) |
| 12&#160;Apr&#160;23&#160;10:11&#160;UTC | SHAKEN | 12&#160;May&#160;23&#160;10:11&#160;UTC | true | [view](CERTS/a89dec6b7a7bb5e663a46945b8c0cfead92b0008f3cf3773f05e02466c3b93e5/README.md) |
| 12&#160;Apr&#160;23&#160;12:30&#160;UTC | SHAKEN | 12&#160;May&#160;23&#160;12:30&#160;UTC | true | [view](CERTS/97336812019ac501322fac2af3c9719288ebbd26c2395779edefe04eb52df614/README.md) |
| 13&#160;Apr&#160;23&#160;10:11&#160;UTC | SHAKEN | 13&#160;May&#160;23&#160;10:11&#160;UTC | true | [view](CERTS/163284d022153265ef8d613910fa55c8e6c1ab8150480a4b3e17d5095332ae97/README.md) |
| 13&#160;Apr&#160;23&#160;12:30&#160;UTC | SHAKEN | 13&#160;May&#160;23&#160;12:30&#160;UTC | true | [view](CERTS/17a6e8cd4ec92d809b710f025d19b1470e1b8016f07c65bbf5cbe1e36b5debc9/README.md) |
| 14&#160;Apr&#160;23&#160;10:11&#160;UTC | SHAKEN | 14&#160;May&#160;23&#160;10:11&#160;UTC | true | [view](CERTS/a919223fe20a244ff625709f3cd98bb86aece8223ff01436c2ff161774b801d6/README.md) |
| 14&#160;Apr&#160;23&#160;12:30&#160;UTC | SHAKEN | 14&#160;May&#160;23&#160;12:30&#160;UTC | true | [view](CERTS/08f3320a1d68bccf268efc45231925c6f6ddbf66c77d6b58ae313f42d948a308/README.md) |
| 15&#160;Apr&#160;23&#160;12:30&#160;UTC | SHAKEN | 15&#160;May&#160;23&#160;12:30&#160;UTC | true | [view](CERTS/17cd326359395a3ea78af14335ebaee024a578824707671e9ce0ea965ad5b538/README.md) |
| 16&#160;Apr&#160;23&#160;12:30&#160;UTC | SHAKEN | 16&#160;May&#160;23&#160;12:30&#160;UTC | true | [view](CERTS/75b130b488b97cb9e065a6081ba1362632bc6f6d17a608bcfe1c3811371f5a0c/README.md) |
| 17&#160;Apr&#160;23&#160;10:11&#160;UTC | SHAKEN | 17&#160;May&#160;23&#160;10:11&#160;UTC | true | [view](CERTS/7b8dc8e707e4eebda86b28fa80fff934df4776ffb2ee5e6d089617c0299b3465/README.md) |
| 17&#160;Apr&#160;23&#160;12:30&#160;UTC | SHAKEN | 17&#160;May&#160;23&#160;12:30&#160;UTC | true | [view](CERTS/226f67238d881df8e6f396e7cbbf1e4df859c28c94caed8ca66ca66164d38db0/README.md) |
| 18&#160;Apr&#160;23&#160;10:11&#160;UTC | SHAKEN | 18&#160;May&#160;23&#160;10:11&#160;UTC | true | [view](CERTS/74db5ab295b087c692288a69b6e92c4d24cbff0af30b2fac73419245a0b876c0/README.md) |
| 18&#160;Apr&#160;23&#160;12:30&#160;UTC | SHAKEN | 18&#160;May&#160;23&#160;12:30&#160;UTC | true | [view](CERTS/b03f2b113caef137babe292456a133893e1ccdd419c1462e8cc7719c2279ee3b/README.md) |
| 19&#160;Apr&#160;23&#160;10:11&#160;UTC | SHAKEN | 19&#160;May&#160;23&#160;10:11&#160;UTC | true | [view](CERTS/55f783c90921286879f4616a0d079a2744a8e7c3af0f3a696ca94f58f6d07342/README.md) |
| 19&#160;Apr&#160;23&#160;12:30&#160;UTC | SHAKEN | 19&#160;May&#160;23&#160;12:30&#160;UTC | true | [view](CERTS/a4f72e677fde33b1488c3c7971ae5b5a797ee1101637198e4fcc239927e5b436/README.md) |
| 20&#160;Apr&#160;23&#160;10:11&#160;UTC | SHAKEN | 20&#160;May&#160;23&#160;10:11&#160;UTC | true | [view](CERTS/94bd9a582de21d85eb1b346e7564044b3fcea8d2c9483c5b767ce82c116d5060/README.md) |
| 20&#160;Apr&#160;23&#160;12:30&#160;UTC | SHAKEN | 20&#160;May&#160;23&#160;12:30&#160;UTC | true | [view](CERTS/20993a9ac47ae87a483a48b0cabe2c2b1a4eb52723f8db3f6fbd83f4c3c8d436/README.md) |
| 21&#160;Apr&#160;23&#160;12:30&#160;UTC | SHAKEN | 21&#160;May&#160;23&#160;12:30&#160;UTC | true | [view](CERTS/4d4d573741dfeb2c382e8c44799f8ef673343bd1f42172afb533694c48dd4732/README.md) |

#### CA Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 17&#160;Mar&#160;20&#160;14:45&#160;UTC | Comcast SHAKEN Root CA | 12&#160;Mar&#160;40&#160;14:45&#160;UTC | false | [view](CERTS/b1132c5f12c3ca4d2ff119f2df99544336eb1703512ac99cc42d596e25125bbd/README.md) |
| 06&#160;Apr&#160;20&#160;13:48&#160;UTC | Comcast SHAKEN Intermediate CA | 06&#160;Jun&#160;39&#160;13:48&#160;UTC | false | [view](CERTS/2f3abdfe711377f4d59f31d941962797c61f021c3924af31de99031f9ac54f77/README.md) |


Generated: 28 Apr 23 02:17 UTC