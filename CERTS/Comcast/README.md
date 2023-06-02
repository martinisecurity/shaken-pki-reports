# STIR/SHAKEN CA Ecosystem Compliance

## Comcast

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 122 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 70 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 52 certificates being tested against the remaining rules
- 6.00 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 100.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 29 days is the average remaining validity for the certificates in the corpus
- 30 days is the average initial validity for the certificates in the corpus
- 52 certificates expire in the next 30 days
- 52.00 average number of unexpired certificates per OCN observed
- 1 unique OCNs observed in unexpired and valid certificate corpus

| Instances | Test | Source |
|-----------|------|--------|
| 52 | [e_atis_certificate_policies](ISSUES/e_atis_certificate_policies/README.md) | ATIS1000080 |
| 52 | [e_atis_subject_cn](ISSUES/e_atis_subject_cn/README.md) | ATIS1000080 |
| 52 | [e_atis_subject_key_identifier](ISSUES/e_atis_subject_key_identifier/README.md) | ATIS1000080 |
| 52 | [e_us_cp_ambiguous_identifier](ISSUES/e_us_cp_ambiguous_identifier/README.md) | US_SHAKEN_CP |
| 52 | [e_us_cp_subject_sn](ISSUES/e_us_cp_subject_sn/README.md) | US_SHAKEN_CP |
| 52 | [w_ext_subject_key_identifier_missing_sub_cert](ISSUES/w_ext_subject_key_identifier_missing_sub_cert/README.md) | RFC5280 |

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
- 6810 days is the average remaining validity for the certificates in the corpus
- 7150 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days

No error, warning, or notice level issues were found

### Details

#### Leaf Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 03&#160;May&#160;23&#160;10:11&#160;UTC | SHAKEN | 02&#160;Jun&#160;23&#160;10:11&#160;UTC | true | [view](CERTS/c723596b6833f0d87c495530e420d8713b20670904d49ef5f65edb2e614a274d/README.md) |
| 03&#160;May&#160;23&#160;12:30&#160;UTC | SHAKEN | 02&#160;Jun&#160;23&#160;12:30&#160;UTC | true | [view](CERTS/680317aa8eb7d278a7a971d9174c7b54247d55bb3ee4c16ffc1d45254dbc2fd2/README.md) |
| 04&#160;May&#160;23&#160;10:11&#160;UTC | SHAKEN | 03&#160;Jun&#160;23&#160;10:11&#160;UTC | true | [view](CERTS/36c0029140ca8f683bb983dd51d4df33541b4ca7b7868c3a9929845658bad341/README.md) |
| 04&#160;May&#160;23&#160;12:30&#160;UTC | SHAKEN | 03&#160;Jun&#160;23&#160;12:30&#160;UTC | true | [view](CERTS/879f77edb83d9d27c41de978011d59461873a20407cd5d8c771dea868c518df8/README.md) |
| 04&#160;May&#160;23&#160;16:37&#160;UTC | SHAKEN | 03&#160;Jun&#160;23&#160;16:37&#160;UTC | true | [view](CERTS/ffe6e12d7108e5cd6dd454df403dd8777d3b7d0c79794195c6d0e2ba1d2b7cbc/README.md) |
| 05&#160;May&#160;23&#160;10:11&#160;UTC | SHAKEN | 04&#160;Jun&#160;23&#160;10:11&#160;UTC | true | [view](CERTS/34223627acb7439c76d388b992aa1733bd716d138736a4e1cee0a143a07b12e7/README.md) |
| 05&#160;May&#160;23&#160;12:30&#160;UTC | SHAKEN | 04&#160;Jun&#160;23&#160;12:30&#160;UTC | true | [view](CERTS/9e5f58003aff4179faeda9fcec8ff597179394efc2a45c5d4f17bb89be00727b/README.md) |
| 06&#160;May&#160;23&#160;10:11&#160;UTC | SHAKEN | 05&#160;Jun&#160;23&#160;10:11&#160;UTC | true | [view](CERTS/a9ca7aab0d31d7bd10a2d09e0862f1b5655f8c861c1fc77b32eca471c5bc87f8/README.md) |
| 06&#160;May&#160;23&#160;12:30&#160;UTC | SHAKEN | 05&#160;Jun&#160;23&#160;12:30&#160;UTC | true | [view](CERTS/362e10a06212791f9e8eeb4669a82900a084f178e7799aaf9c1e91da0a7e3505/README.md) |
| 07&#160;May&#160;23&#160;12:30&#160;UTC | SHAKEN | 06&#160;Jun&#160;23&#160;12:30&#160;UTC | true | [view](CERTS/8fe3b7b818b7df0c645dc1cce660b0a35118ee37aba3325e0f39bb04e4c71bd2/README.md) |
| 08&#160;May&#160;23&#160;10:11&#160;UTC | SHAKEN | 07&#160;Jun&#160;23&#160;10:11&#160;UTC | true | [view](CERTS/9c73ac02ca7402a181048801249a333ab38a764eae331291a651a2449dbf6b1a/README.md) |
| 08&#160;May&#160;23&#160;12:30&#160;UTC | SHAKEN | 07&#160;Jun&#160;23&#160;12:30&#160;UTC | true | [view](CERTS/6c0f4314bb7cddcbf3a50a4202a1a81722bb31e0c7032ea816665aaaffe7ee25/README.md) |
| 09&#160;May&#160;23&#160;10:11&#160;UTC | SHAKEN | 08&#160;Jun&#160;23&#160;10:11&#160;UTC | true | [view](CERTS/1ad7123f3d6148d7b2ae599000f76e82b7c67e6f12d09faca41510423ff81f51/README.md) |
| 09&#160;May&#160;23&#160;12:30&#160;UTC | SHAKEN | 08&#160;Jun&#160;23&#160;12:30&#160;UTC | true | [view](CERTS/61f913d774c0476f6f5e9ca91a8286b5bea162d46a1eccbd7cfbaf22fc20d689/README.md) |
| 10&#160;May&#160;23&#160;10:11&#160;UTC | SHAKEN | 09&#160;Jun&#160;23&#160;10:11&#160;UTC | true | [view](CERTS/c7508fc3e15f1126574d3daf8b5b90f7c13f823a8bcd2463861178e9be9a265d/README.md) |
| 10&#160;May&#160;23&#160;12:30&#160;UTC | SHAKEN | 09&#160;Jun&#160;23&#160;12:30&#160;UTC | true | [view](CERTS/86074b64ddc8d759bfcbebf5f90f534808c31107468fbf3291f43a5af71b930b/README.md) |
| 11&#160;May&#160;23&#160;10:11&#160;UTC | SHAKEN | 10&#160;Jun&#160;23&#160;10:11&#160;UTC | true | [view](CERTS/699cbd58c54e3a64987e7154b7681948529994783dd4975cd406c5a95b7313b3/README.md) |
| 11&#160;May&#160;23&#160;12:30&#160;UTC | SHAKEN | 10&#160;Jun&#160;23&#160;12:30&#160;UTC | true | [view](CERTS/1855f4aec01ee2438b400cb7c2807d60c9027a8789997205d7de002c17bb77ff/README.md) |
| 12&#160;May&#160;23&#160;10:11&#160;UTC | SHAKEN | 11&#160;Jun&#160;23&#160;10:11&#160;UTC | true | [view](CERTS/13d914092c68bfa21cec1db5ed40808d9ac9941f03e05e3f71c8d85da386e52f/README.md) |
| 12&#160;May&#160;23&#160;12:30&#160;UTC | SHAKEN | 11&#160;Jun&#160;23&#160;12:30&#160;UTC | true | [view](CERTS/3a0dfd172435b9633a6954c0936860a76dec95f9e033c5418218c97a821c77bb/README.md) |
| 13&#160;May&#160;23&#160;10:11&#160;UTC | SHAKEN | 12&#160;Jun&#160;23&#160;10:11&#160;UTC | true | [view](CERTS/1f2083f50f74eccc81bfdcec345eb25c2671ae36cab20cdc0224815c028ed2d6/README.md) |
| 13&#160;May&#160;23&#160;12:30&#160;UTC | SHAKEN | 12&#160;Jun&#160;23&#160;12:30&#160;UTC | true | [view](CERTS/9498e12d5c74d8fe39e4a44910591fbdd4da3cba7f12bc2a1ebc7e58c390e584/README.md) |
| 14&#160;May&#160;23&#160;12:30&#160;UTC | SHAKEN | 13&#160;Jun&#160;23&#160;12:30&#160;UTC | true | [view](CERTS/6d116141de1ca9421453177e7cd6a0f880206c114ee2ec7ddcc683a7b27d65a8/README.md) |
| 15&#160;May&#160;23&#160;10:11&#160;UTC | SHAKEN | 14&#160;Jun&#160;23&#160;10:11&#160;UTC | true | [view](CERTS/7a40b72910f3e9799c3e387a8d2284c84338abecdc1429e21997f05cadef9427/README.md) |
| 15&#160;May&#160;23&#160;12:30&#160;UTC | SHAKEN | 14&#160;Jun&#160;23&#160;12:30&#160;UTC | true | [view](CERTS/1ef5dbaa009e70e4166ba9b7b9e5b1f082d863b7647f9a099c89aadb04c34c0d/README.md) |
| 16&#160;May&#160;23&#160;12:30&#160;UTC | SHAKEN | 15&#160;Jun&#160;23&#160;12:30&#160;UTC | true | [view](CERTS/fac9791978b8a3cabba7cf5a7fd81c96aeca561ca3d181d7c3e3afb2eaa35486/README.md) |
| 17&#160;May&#160;23&#160;10:11&#160;UTC | SHAKEN | 16&#160;Jun&#160;23&#160;10:11&#160;UTC | true | [view](CERTS/db76aeb7f49e2d16a011195aa10c3af8b33cb96c6e329ccdcf68d81e8aec6466/README.md) |
| 17&#160;May&#160;23&#160;12:30&#160;UTC | SHAKEN | 16&#160;Jun&#160;23&#160;12:30&#160;UTC | true | [view](CERTS/c0cb0aed45c9f4a48d7cf2237215e244cd42d39f8283c90439e6cb3f92511e64/README.md) |
| 18&#160;May&#160;23&#160;10:11&#160;UTC | SHAKEN | 17&#160;Jun&#160;23&#160;10:11&#160;UTC | true | [view](CERTS/fa91eb5c84feec0d95be86d8f25e63c1af5ac8f2f05ff2d2fa60737067a3ef49/README.md) |
| 18&#160;May&#160;23&#160;12:30&#160;UTC | SHAKEN | 17&#160;Jun&#160;23&#160;12:30&#160;UTC | true | [view](CERTS/52f45794c6f42cc7059ff7c10f7a22b089bafdad9cbbbfb395431133e43e07e2/README.md) |
| 19&#160;May&#160;23&#160;10:11&#160;UTC | SHAKEN | 18&#160;Jun&#160;23&#160;10:11&#160;UTC | true | [view](CERTS/c5dcd8aff9441fdc0113e083ddd9f983dafd0cb1ed1a645167c344227efb8efc/README.md) |
| 19&#160;May&#160;23&#160;12:30&#160;UTC | SHAKEN | 18&#160;Jun&#160;23&#160;12:30&#160;UTC | true | [view](CERTS/97b34d720cd2b27a5b2958a1bd15df8e52df691cc93dba240c1a2e5d013f2a2a/README.md) |
| 20&#160;May&#160;23&#160;10:11&#160;UTC | SHAKEN | 19&#160;Jun&#160;23&#160;10:11&#160;UTC | true | [view](CERTS/714fb06fc5a734fef70fde557a03863fb6589eda6691807e9dd02bb47e9fbc11/README.md) |
| 20&#160;May&#160;23&#160;12:30&#160;UTC | SHAKEN | 19&#160;Jun&#160;23&#160;12:30&#160;UTC | true | [view](CERTS/36ca3df1e783d4365e3db4af526e8f46d8c6637475c3e6a0acffee2d303f5b97/README.md) |
| 21&#160;May&#160;23&#160;12:30&#160;UTC | SHAKEN | 20&#160;Jun&#160;23&#160;12:30&#160;UTC | true | [view](CERTS/8b0087ec8ce0a7b8bd5e422ea9f1209a3361ce04c055d8e71742c8bc866c04a4/README.md) |
| 22&#160;May&#160;23&#160;10:11&#160;UTC | SHAKEN | 21&#160;Jun&#160;23&#160;10:11&#160;UTC | true | [view](CERTS/290c12eee4bfe13bd70e5a1ff3547f09508db2ba00b0e0cb5ea701a1bbb55ea6/README.md) |
| 22&#160;May&#160;23&#160;12:30&#160;UTC | SHAKEN | 21&#160;Jun&#160;23&#160;12:30&#160;UTC | true | [view](CERTS/26cca1744a2be73d5b9d4dcd5835e641cf5d25f0c06135d643432defbeb5c0ef/README.md) |
| 23&#160;May&#160;23&#160;10:11&#160;UTC | SHAKEN | 22&#160;Jun&#160;23&#160;10:11&#160;UTC | true | [view](CERTS/c5862c021264f0947f8dbc114566666d81dc9c07292b19288c1fcdb54c0cdba7/README.md) |
| 23&#160;May&#160;23&#160;12:30&#160;UTC | SHAKEN | 22&#160;Jun&#160;23&#160;12:30&#160;UTC | true | [view](CERTS/ea3102db093a64f582782e1bf5964063689428745f58b8d6b9e40558f5eaf141/README.md) |
| 24&#160;May&#160;23&#160;10:11&#160;UTC | SHAKEN | 23&#160;Jun&#160;23&#160;10:11&#160;UTC | true | [view](CERTS/eddfcaf3f1ab56994dbc05316d20c4712ba3c2a8b0cc853d582cd9220296590f/README.md) |
| 24&#160;May&#160;23&#160;12:30&#160;UTC | SHAKEN | 23&#160;Jun&#160;23&#160;12:30&#160;UTC | true | [view](CERTS/0d0df0b547cd1d562d745a7f73239f51584cebd4692697690e372eb32b18c933/README.md) |
| 25&#160;May&#160;23&#160;10:11&#160;UTC | SHAKEN | 24&#160;Jun&#160;23&#160;10:11&#160;UTC | true | [view](CERTS/3389c1b863cc0e9bbe2967ab8a7fbb4a36413cbc01a8ac7d78b49c46a71cadb6/README.md) |
| 25&#160;May&#160;23&#160;12:30&#160;UTC | SHAKEN | 24&#160;Jun&#160;23&#160;12:30&#160;UTC | true | [view](CERTS/b68fb21d138c9490f1576889b1b9e7372af7f0784157b35b3736cafa7c130169/README.md) |
| 26&#160;May&#160;23&#160;10:11&#160;UTC | SHAKEN | 25&#160;Jun&#160;23&#160;10:11&#160;UTC | true | [view](CERTS/291c9b7c6b7c3b8206a1f1e686adf4b6e14effacd4ac90d380931564a1b63bdb/README.md) |
| 26&#160;May&#160;23&#160;12:30&#160;UTC | SHAKEN | 25&#160;Jun&#160;23&#160;12:30&#160;UTC | true | [view](CERTS/e646fa05f0479dc828d26248d09f77a91bc9750a9e7fca23fc4ed6c9c027f531/README.md) |
| 27&#160;May&#160;23&#160;10:11&#160;UTC | SHAKEN | 26&#160;Jun&#160;23&#160;10:11&#160;UTC | true | [view](CERTS/498fc801b62b25c0c2b572a6881b60d497ffe374511a7f861cca17f69a863d10/README.md) |
| 27&#160;May&#160;23&#160;12:30&#160;UTC | SHAKEN | 26&#160;Jun&#160;23&#160;12:30&#160;UTC | true | [view](CERTS/e2fff7bf5ee495890fae5ff095d7e01867489d97800526e00f7806564c89c06c/README.md) |
| 28&#160;May&#160;23&#160;12:30&#160;UTC | SHAKEN | 27&#160;Jun&#160;23&#160;12:30&#160;UTC | true | [view](CERTS/fb59ac51427075af933d0b2d329f6db50ccff1eec7f6ff33b66bed70a1478183/README.md) |
| 29&#160;May&#160;23&#160;12:30&#160;UTC | SHAKEN | 28&#160;Jun&#160;23&#160;12:30&#160;UTC | true | [view](CERTS/894a42e6ba4f844c6a994c8fa48ce0114b72069d89f1ac20e55f300e36300047/README.md) |
| 30&#160;May&#160;23&#160;10:11&#160;UTC | SHAKEN | 29&#160;Jun&#160;23&#160;10:11&#160;UTC | true | [view](CERTS/4ee050aed50cc4080150492d9c0db72c4c8ffe2b56b5398ec52c3b3e3dd5bbd9/README.md) |
| 30&#160;May&#160;23&#160;12:30&#160;UTC | SHAKEN | 29&#160;Jun&#160;23&#160;12:30&#160;UTC | true | [view](CERTS/8fbd163a5e2b1d39cebc8c0c8fa6bbb3b3b0efc860dff32ef346ec9908eddceb/README.md) |
| 31&#160;May&#160;23&#160;12:30&#160;UTC | SHAKEN | 30&#160;Jun&#160;23&#160;12:30&#160;UTC | true | [view](CERTS/bc7758333fc47c86488331b6189339b3a04bb6b5f61d3f77fcdcb3a04a2549a1/README.md) |

#### CA Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 17&#160;Mar&#160;20&#160;14:45&#160;UTC | Comcast SHAKEN Root CA | 12&#160;Mar&#160;40&#160;14:45&#160;UTC | false | [view](CERTS/b1132c5f12c3ca4d2ff119f2df99544336eb1703512ac99cc42d596e25125bbd/README.md) |
| 06&#160;Apr&#160;20&#160;13:48&#160;UTC | Comcast SHAKEN Intermediate CA | 06&#160;Jun&#160;39&#160;13:48&#160;UTC | false | [view](CERTS/2f3abdfe711377f4d59f31d941962797c61f021c3924af31de99031f9ac54f77/README.md) |


Generated: 02 Jun 23 01:12 UTC