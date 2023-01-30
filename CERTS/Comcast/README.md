# STIR/SHAKEN CA Ecosystem Compliance

## Comcast

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 130 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 111 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 19 certificates being tested against the remaining rules
- 6.00 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 100.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 29 days is the average remaining validity for the certificates in the corpus
- 30 days is the average initial validity for the certificates in the corpus
- 19 certificates expire in the next 30 days
- 19.00 average number of unexpired certificates per OCN observed
- 1 unique OCNs observed in unexpired and valid certificate corpus

| Instances | Test | Source |
|-----------|------|--------|
| 19 | [e_atis_certificate_policies](ISSUES/e_atis_certificate_policies/README.md) | ATIS1000080 |
| 19 | [e_atis_subject_cn](ISSUES/e_atis_subject_cn/README.md) | ATIS1000080 |
| 19 | [e_atis_subject_key_identifier](ISSUES/e_atis_subject_key_identifier/README.md) | ATIS1000080 |
| 19 | [e_us_cp_ambiguous_identifier](ISSUES/e_us_cp_ambiguous_identifier/README.md) | US_SHAKEN_CP |
| 19 | [e_us_cp_subject_sn](ISSUES/e_us_cp_subject_sn/README.md) | US_SHAKEN_CP |
| 19 | [w_ext_subject_key_identifier_missing_sub_cert](ISSUES/w_ext_subject_key_identifier_missing_sub_cert/README.md) | RFC5280 |

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
- 6850 days is the average remaining validity for the certificates in the corpus
- 7150 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days

No error, warning, or notice level issues were found

### Details

#### Leaf Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 01&#160;Jan&#160;23&#160;11:12&#160;UTC | SHAKEN | 31&#160;Jan&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/d1856c43a51de1a134f66f52b814e02f74337ad5b333a535998b0c902fea1732/README.md) |
| 02&#160;Jan&#160;23&#160;11:12&#160;UTC | SHAKEN | 01&#160;Feb&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/719acb92673a27a95a56884222e14e9c325ad7e00a496f51ee2ca0b81377c1b6/README.md) |
| 03&#160;Jan&#160;23&#160;11:12&#160;UTC | SHAKEN | 02&#160;Feb&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/ab5ef23f58167423e7ea6cefb7a355d68063da131be72a0464e4dd848866499e/README.md) |
| 04&#160;Jan&#160;23&#160;11:12&#160;UTC | SHAKEN | 03&#160;Feb&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/d03657e109d20dda9ca046a59bd1b6381a611304f9039b7b0bb1e5d547f9d975/README.md) |
| 05&#160;Jan&#160;23&#160;11:12&#160;UTC | SHAKEN | 04&#160;Feb&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/6f49f93c37dfd0a797cbb76e1bb375187bc3dd81e290d71f09062e146ffbc5a5/README.md) |
| 06&#160;Jan&#160;23&#160;11:12&#160;UTC | SHAKEN | 05&#160;Feb&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/6bb26adc661a2c1677f5d179b39730ef80de3915f5e173c8266d9b60859df014/README.md) |
| 07&#160;Jan&#160;23&#160;11:12&#160;UTC | SHAKEN | 06&#160;Feb&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/8c0dec73d65c8ad65cea126e9e3c2b67bafecc1227bd7c75d9f86d0fc7fbbfd8/README.md) |
| 08&#160;Jan&#160;23&#160;11:12&#160;UTC | SHAKEN | 07&#160;Feb&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/981b1c08477bff0561af30914688b6f7851b1f00f8cdbd5f488fe2fe57bce688/README.md) |
| 09&#160;Jan&#160;23&#160;11:12&#160;UTC | SHAKEN | 08&#160;Feb&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/e4de04a07cfa306e5185151898b46480588a12b371670ec2a6ba9c80f3412670/README.md) |
| 10&#160;Jan&#160;23&#160;11:12&#160;UTC | SHAKEN | 09&#160;Feb&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/f58fe89d75ec3f17d7ba4460e383485003b05a142325615513099b6b2bedf33b/README.md) |
| 11&#160;Jan&#160;23&#160;11:12&#160;UTC | SHAKEN | 10&#160;Feb&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/22a1688d3d5a97030915af1b9270dbb87ed0c2140e0290691a55f9d18a3d40a9/README.md) |
| 12&#160;Jan&#160;23&#160;11:12&#160;UTC | SHAKEN | 11&#160;Feb&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/96bd34ff4f5772191d0f3984a42c1fcffef4c23aea852585b0508b29be288777/README.md) |
| 13&#160;Jan&#160;23&#160;11:12&#160;UTC | SHAKEN | 12&#160;Feb&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/3cce9b32d9f780fe7977b2eeff88d800b49989bdd193ccc1cdd7c25418969986/README.md) |
| 14&#160;Jan&#160;23&#160;11:12&#160;UTC | SHAKEN | 13&#160;Feb&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/07d64b5f2eb6a38227ce82d05fada9cb92fe9842b6625e975bdf02d6c16fa958/README.md) |
| 15&#160;Jan&#160;23&#160;11:12&#160;UTC | SHAKEN | 14&#160;Feb&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/077793c1b3e8535718879c96f848ccf4cf640104111f28c7efeea832bc3b4035/README.md) |
| 16&#160;Jan&#160;23&#160;11:12&#160;UTC | SHAKEN | 15&#160;Feb&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/3a45713fda01a02191819d5a0527d51f6591ed7f55ff467f0d8f62d22aa5a7c4/README.md) |
| 17&#160;Jan&#160;23&#160;11:12&#160;UTC | SHAKEN | 16&#160;Feb&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/d6ea9611636aa315c22f9443284302e3f4b086bfcc7140cf54e4efafb79e3a0f/README.md) |
| 18&#160;Jan&#160;23&#160;11:12&#160;UTC | SHAKEN | 17&#160;Feb&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/65f138de22895f83ea30f71fef09ca3b2b0b9bfb1d593689717946ba03a7aab4/README.md) |
| 19&#160;Jan&#160;23&#160;11:12&#160;UTC | SHAKEN | 18&#160;Feb&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/ef6e6415ea3ddc7b3ad4a28d3acea23c3f1ac063ff2d0f8fe0beba3bb282524a/README.md) |

#### CA Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 17&#160;Mar&#160;20&#160;14:45&#160;UTC | Comcast SHAKEN Root CA | 12&#160;Mar&#160;40&#160;14:45&#160;UTC | false | [view](CERTS/b1132c5f12c3ca4d2ff119f2df99544336eb1703512ac99cc42d596e25125bbd/README.md) |
| 06&#160;Apr&#160;20&#160;13:48&#160;UTC | Comcast SHAKEN Intermediate CA | 06&#160;Jun&#160;39&#160;13:48&#160;UTC | false | [view](CERTS/2f3abdfe711377f4d59f31d941962797c61f021c3924af31de99031f9ac54f77/README.md) |


Generated: 30 Jan 23 23:10 UTC