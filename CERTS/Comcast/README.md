# STIR/SHAKEN CA Ecosystem Compliance

## Comcast

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 110 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 75 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 35 certificates being tested against the remaining rules
- 6.00 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 100.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 29 days is the average remaining validity for the certificates in the corpus
- 30 days is the average initial validity for the certificates in the corpus
- 35 certificates expire in the next 30 days
- 35.00 average number of unexpired certificates per OCN observed
- 1 unique OCNs observed in unexpired and valid certificate corpus

| Instances | Test | Source |
|-----------|------|--------|
| 35 | [e_atis_certificate_policies](ISSUES/e_atis_certificate_policies/README.md) | ATIS1000080 |
| 35 | [e_atis_subject_cn](ISSUES/e_atis_subject_cn/README.md) | ATIS1000080 |
| 35 | [e_atis_subject_key_identifier](ISSUES/e_atis_subject_key_identifier/README.md) | ATIS1000080 |
| 35 | [e_us_cp_ambiguous_identifier](ISSUES/e_us_cp_ambiguous_identifier/README.md) | US_SHAKEN_CP |
| 35 | [e_us_cp_subject_sn](ISSUES/e_us_cp_subject_sn/README.md) | US_SHAKEN_CP |
| 35 | [w_ext_subject_key_identifier_missing_sub_cert](ISSUES/w_ext_subject_key_identifier_missing_sub_cert/README.md) | RFC5280 |

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
- 6783 days is the average remaining validity for the certificates in the corpus
- 7150 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days

No error, warning, or notice level issues were found

### Details

#### Leaf Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 23&#160;Jul&#160;23&#160;10:11&#160;UTC | SHAKEN | 22&#160;Aug&#160;23&#160;10:11&#160;UTC | true | [view](CERTS/980f4a4f8e87ac7a1f75df41683fa6b72ecd8064433a8d06d262d9e6ae5f48ba/README.md) |
| 23&#160;Jul&#160;23&#160;13:55&#160;UTC | SHAKEN | 22&#160;Aug&#160;23&#160;13:55&#160;UTC | true | [view](CERTS/a442df43f70c3d6068767e6c528fffbb101fb8236292b4f052b8ac5b0710fa0c/README.md) |
| 24&#160;Jul&#160;23&#160;10:11&#160;UTC | SHAKEN | 23&#160;Aug&#160;23&#160;10:11&#160;UTC | true | [view](CERTS/67128ccf2571b155dfe649a361ab80f43354a72a163ed7d381de9bdda262c2e8/README.md) |
| 24&#160;Jul&#160;23&#160;13:54&#160;UTC | SHAKEN | 23&#160;Aug&#160;23&#160;13:54&#160;UTC | true | [view](CERTS/20c6503d16f45d2af4afc17f14312092a0b9346ed3ee7bf4062da4305e60a043/README.md) |
| 25&#160;Jul&#160;23&#160;10:11&#160;UTC | SHAKEN | 24&#160;Aug&#160;23&#160;10:11&#160;UTC | true | [view](CERTS/62924f4ee8d178d3ee462bbf810bd32b3d44a82612d00f04dc9e86f3c30797f7/README.md) |
| 25&#160;Jul&#160;23&#160;13:55&#160;UTC | SHAKEN | 24&#160;Aug&#160;23&#160;13:55&#160;UTC | true | [view](CERTS/49aaf872a7b5486eab82d98fc687e744415214d9c5ed0e2de5d2a8fe476ff8c5/README.md) |
| 26&#160;Jul&#160;23&#160;10:11&#160;UTC | SHAKEN | 25&#160;Aug&#160;23&#160;10:11&#160;UTC | true | [view](CERTS/374ad71dc4ecb67e1a5eac7dc8bb964d6f0fd67636e7d7730f1bf2c79a23fbe8/README.md) |
| 26&#160;Jul&#160;23&#160;13:54&#160;UTC | SHAKEN | 25&#160;Aug&#160;23&#160;13:54&#160;UTC | true | [view](CERTS/98d754b6364386656e152728fbc49510890b0cd1d848261b7d6d7c244b629a5b/README.md) |
| 27&#160;Jul&#160;23&#160;10:11&#160;UTC | SHAKEN | 26&#160;Aug&#160;23&#160;10:11&#160;UTC | true | [view](CERTS/bc4c0ca0f9c148dd0f6c4638af25aeb811f17caf0289c78daa48f0e0750e4e1d/README.md) |
| 27&#160;Jul&#160;23&#160;13:55&#160;UTC | SHAKEN | 26&#160;Aug&#160;23&#160;13:55&#160;UTC | true | [view](CERTS/5ab8d528426cfaa59ce2fd0c521dd9127ef805a2f775e3ae4dd5a23994c58315/README.md) |
| 28&#160;Jul&#160;23&#160;10:11&#160;UTC | SHAKEN | 27&#160;Aug&#160;23&#160;10:11&#160;UTC | true | [view](CERTS/932517d08801dc174839a784487dedc720f0e49275ee959b5685764592bbb61d/README.md) |
| 28&#160;Jul&#160;23&#160;13:54&#160;UTC | SHAKEN | 27&#160;Aug&#160;23&#160;13:54&#160;UTC | true | [view](CERTS/b4a693e68b21b6a10ec4ee017cb005ed3e2e319ee10998c5f6f62b57f4435ca9/README.md) |
| 29&#160;Jul&#160;23&#160;10:11&#160;UTC | SHAKEN | 28&#160;Aug&#160;23&#160;10:11&#160;UTC | true | [view](CERTS/b1073b4721044a7a3e19734eb5e0ed122f2aadfc8b967ae1baf055d44ded9ecb/README.md) |
| 29&#160;Jul&#160;23&#160;13:55&#160;UTC | SHAKEN | 28&#160;Aug&#160;23&#160;13:55&#160;UTC | true | [view](CERTS/7ad2cc4ae938387a836603c34526e30246fccbb08ecc6fe7eb935d0334852945/README.md) |
| 30&#160;Jul&#160;23&#160;13:54&#160;UTC | SHAKEN | 29&#160;Aug&#160;23&#160;13:54&#160;UTC | true | [view](CERTS/698005abc1b5c4d1f706cfcfe59fcfbac0e5fe7d0dae417a5ddbd6b7a75eb9f1/README.md) |
| 31&#160;Jul&#160;23&#160;10:11&#160;UTC | SHAKEN | 30&#160;Aug&#160;23&#160;10:11&#160;UTC | true | [view](CERTS/c5641784d141c7f7f48cb13bee94534a838e8251bb8441f338898f753a8774e2/README.md) |
| 31&#160;Jul&#160;23&#160;13:55&#160;UTC | SHAKEN | 30&#160;Aug&#160;23&#160;13:55&#160;UTC | true | [view](CERTS/7386a81752cc1d24ae09cf3d6b76230eba1dc9fa2d7fc14587c6281ec8151095/README.md) |
| 01&#160;Aug&#160;23&#160;10:11&#160;UTC | SHAKEN | 31&#160;Aug&#160;23&#160;10:11&#160;UTC | true | [view](CERTS/6c7b979e8efaf5992a510e0ca7c65d0087b800b6c56c8dfaa2770c7677ff5859/README.md) |
| 01&#160;Aug&#160;23&#160;13:54&#160;UTC | SHAKEN | 31&#160;Aug&#160;23&#160;13:54&#160;UTC | true | [view](CERTS/5b1746396a641bc07802a52786a56fe2b0a8ddfaa06e0e18b4d7b7e5e28ff759/README.md) |
| 02&#160;Aug&#160;23&#160;10:11&#160;UTC | SHAKEN | 01&#160;Sep&#160;23&#160;10:11&#160;UTC | true | [view](CERTS/09990e6194cb43b1ed9c4143717810cdfcabbe6ffa684aabe128713ff8246b3e/README.md) |
| 02&#160;Aug&#160;23&#160;13:55&#160;UTC | SHAKEN | 01&#160;Sep&#160;23&#160;13:55&#160;UTC | true | [view](CERTS/4ba1a6d63f598d1027036af3e6fe82e409af0aae7d22d4c9e5f816440f27a808/README.md) |
| 03&#160;Aug&#160;23&#160;10:11&#160;UTC | SHAKEN | 02&#160;Sep&#160;23&#160;10:11&#160;UTC | true | [view](CERTS/a0c767f9d60c81fa8ae0abf904cfae3dc6b08e5e451c54d0ac72f83c51459e34/README.md) |
| 03&#160;Aug&#160;23&#160;13:54&#160;UTC | SHAKEN | 02&#160;Sep&#160;23&#160;13:54&#160;UTC | true | [view](CERTS/8bdfe629deee3a676e2e3781b01fa38fe424ad66805e18dd84f992394660d237/README.md) |
| 04&#160;Aug&#160;23&#160;10:11&#160;UTC | SHAKEN | 03&#160;Sep&#160;23&#160;10:11&#160;UTC | true | [view](CERTS/0b5aa88fc21985146fae95323c8b6288d01e02197ca485c3e2176e8b82256608/README.md) |
| 04&#160;Aug&#160;23&#160;13:55&#160;UTC | SHAKEN | 03&#160;Sep&#160;23&#160;13:55&#160;UTC | true | [view](CERTS/3ba5e86a905dfaec58eb035ae0b379405a04e5f8d5aea1b7fa5fa8e53106e436/README.md) |
| 05&#160;Aug&#160;23&#160;10:11&#160;UTC | SHAKEN | 04&#160;Sep&#160;23&#160;10:11&#160;UTC | true | [view](CERTS/09aaf209d1eaa479b424fb9c2f017113221aefcc6a1fce9949c39e1afa1b5721/README.md) |
| 05&#160;Aug&#160;23&#160;13:54&#160;UTC | SHAKEN | 04&#160;Sep&#160;23&#160;13:54&#160;UTC | true | [view](CERTS/50af360c86e7d8cad206260ce2068500b840ccdbcb57535c7491ff0adcf5fec7/README.md) |
| 06&#160;Aug&#160;23&#160;10:11&#160;UTC | SHAKEN | 05&#160;Sep&#160;23&#160;10:11&#160;UTC | true | [view](CERTS/5e3efb72ea8b44597905668871d9fa63b1b2d27d7b8bad871b6a545a7444070a/README.md) |
| 06&#160;Aug&#160;23&#160;13:55&#160;UTC | SHAKEN | 05&#160;Sep&#160;23&#160;13:55&#160;UTC | true | [view](CERTS/05418a12615f736a528b46abb53dd105f7069657d388017211f53675cdc8fc1b/README.md) |
| 07&#160;Aug&#160;23&#160;10:11&#160;UTC | SHAKEN | 06&#160;Sep&#160;23&#160;10:11&#160;UTC | true | [view](CERTS/9104ad234301d74163432ffbbe417c81ace0ff25689bdb9e460c8acf4813f647/README.md) |
| 07&#160;Aug&#160;23&#160;13:54&#160;UTC | SHAKEN | 06&#160;Sep&#160;23&#160;13:54&#160;UTC | true | [view](CERTS/8eb8f8e449e5fd228d92c4ab81476998f35f2d11c45e600c14f376489ff967d7/README.md) |
| 08&#160;Aug&#160;23&#160;10:11&#160;UTC | SHAKEN | 07&#160;Sep&#160;23&#160;10:11&#160;UTC | true | [view](CERTS/d672622408d932f34521d82ee3fe015485c401ee7cf876646c83a0a7fd368d2d/README.md) |
| 08&#160;Aug&#160;23&#160;13:54&#160;UTC | SHAKEN | 07&#160;Sep&#160;23&#160;13:54&#160;UTC | true | [view](CERTS/d2885cfd2e474b28b969046b5accaa20226940754fbc341f020187044f45b2f9/README.md) |
| 09&#160;Aug&#160;23&#160;10:11&#160;UTC | SHAKEN | 08&#160;Sep&#160;23&#160;10:11&#160;UTC | true | [view](CERTS/f31985b10f10792a224e99c47e7879b9c09e3683e14bcaf2b3f6c9240dcb51d0/README.md) |
| 09&#160;Aug&#160;23&#160;13:54&#160;UTC | SHAKEN | 08&#160;Sep&#160;23&#160;13:54&#160;UTC | true | [view](CERTS/69f6338a80bc9d974154477343ad1bfed6546497ec4a16634b52cec8e5de581f/README.md) |

#### CA Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 17&#160;Mar&#160;20&#160;14:45&#160;UTC | Comcast SHAKEN Root CA | 12&#160;Mar&#160;40&#160;14:45&#160;UTC | false | [view](CERTS/b1132c5f12c3ca4d2ff119f2df99544336eb1703512ac99cc42d596e25125bbd/README.md) |
| 06&#160;Apr&#160;20&#160;13:48&#160;UTC | Comcast SHAKEN Intermediate CA | 06&#160;Jun&#160;39&#160;13:48&#160;UTC | false | [view](CERTS/2f3abdfe711377f4d59f31d941962797c61f021c3924af31de99031f9ac54f77/README.md) |


Generated: 21 Aug 23 20:18 UTC