# STIR/SHAKEN CA Ecosystem Compliance

## Metaswitch

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 36 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 0 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 36 certificates being tested against the remaining rules
- 4.14 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 2.78% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 91.67% of certificates are too old to be assessed against currently enforced expectations
- 1077 days is the average remaining validity for the certificates in the corpus
- 1095 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days
- 1.00 average number of unexpired certificates per OCN observed
- 36 unique OCNs observed in unexpired and valid certificate corpus

| Instances | Test | Source |
|-----------|------|--------|
| 36 | [e_atis_issuer_dn](ISSUES/e_atis_issuer_dn/README.md) | ATIS1000080 |
| 36 | [e_atis_key_usage](ISSUES/e_atis_key_usage/README.md) | ATIS1000080 |
| 1 | [e_atis_subject](ISSUES/e_atis_subject/README.md) | ATIS1000080 |
| 3 | [e_atis_subject_cn](ISSUES/e_atis_subject_cn/README.md) | ATIS1000080 |
| 36 | [e_us_cp_ambiguous_identifier](ISSUES/e_us_cp_ambiguous_identifier/README.md) | US_SHAKEN_CP |
| 36 | [e_us_cp_subject_sn](ISSUES/e_us_cp_subject_sn/README.md) | US_SHAKEN_CP |
| 1 | [w_shaken_subject_rdn_unknown](ISSUES/w_shaken_subject_rdn_unknown/README.md) | SHAKEN_PKI_BEST_PRACTICES |

#### CA Certificates

- 2 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 0 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 2 certificates being tested against the remaining rules
- 1.00 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 0.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 100.00% of certificates are too old to be assessed against currently enforced expectations
- 6076 days is the average remaining validity for the certificates in the corpus
- 5840 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days

| Instances | Test | Source |
|-----------|------|--------|
| 2 | [e_atis_ca_subject](ISSUES/e_atis_ca_subject/README.md) | ATIS1000080 |

### Details

#### Leaf Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 11&#160;Dec&#160;20&#160;00:14&#160;UTC | Fidelity Communications SHAKEN Cert 1882 | 11&#160;Dec&#160;23&#160;00:14&#160;UTC | true | [view](CERTS/baaf8e58db0f02327fc5b9b614a5633c7b505ca0b291b606d32a191ee73a05e5/README.md) |
| 17&#160;Dec&#160;20&#160;15:19&#160;UTC | TDS Telecom SHAKEN Cert 7804 | 17&#160;Dec&#160;23&#160;15:19&#160;UTC | true | [view](CERTS/a04a669738b79ff55c2b2197f72a12a112b731906e2e6a925d37ccee2fa00a11/README.md) |
| 07&#160;Jan&#160;21&#160;16:17&#160;UTC | Call48 SHAKEN Cert 505J | 07&#160;Jan&#160;24&#160;16:17&#160;UTC | true | [view](CERTS/5bb4516a62167e6c55e9704e4a49b39ce9ec480e678aa6c02fb6985fc9594997/README.md) |
| 27&#160;Jan&#160;21&#160;17:16&#160;UTC | Telesystem SHAKEN Cert 786E | 27&#160;Jan&#160;24&#160;17:16&#160;UTC | true | [view](CERTS/2d9aca0895c94291596161363091718089a6e7c19dfa57329ae548432533860f/README.md) |
| 16&#160;Feb&#160;21&#160;17:30&#160;UTC | Verizon SHAKEN cert 5807 | 16&#160;Feb&#160;24&#160;17:30&#160;UTC | true | [view](CERTS/d7b413267be2d050d516af4f4a864ffdc2feacc774a1a6264b9cfe68c796f43f/README.md) |
| 10&#160;Mar&#160;21&#160;20:40&#160;UTC | Allstream SHAKEN cert 4130 | 09&#160;Mar&#160;24&#160;20:40&#160;UTC | true | [view](CERTS/a9eec584cf66eade89fa84e043e0eb05900dc40653ea963a96c079ef92d9e2ad/README.md) |
| 10&#160;Mar&#160;21&#160;20:50&#160;UTC | Northeast Communications of Wisconsin SHAKEN Cert 6692 | 09&#160;Mar&#160;24&#160;20:50&#160;UTC | true | [view](CERTS/90042ab31d5de7abb64b89073a0d931ce22d864b89df13a9372887cb5db45f49/README.md) |
| 17&#160;Mar&#160;21&#160;17:06&#160;UTC | RCN SHAKEN Cert 7615 | 16&#160;Mar&#160;24&#160;17:06&#160;UTC | true | [view](CERTS/bbdec20ad80f4a2a8ed097204a9299566beca170460fb648c81a51d195d9b6f1/README.md) |
| 17&#160;Mar&#160;21&#160;17:12&#160;UTC | GCI SHAKEN Cert 7785 | 16&#160;Mar&#160;24&#160;17:12&#160;UTC | true | [view](CERTS/312e58dffa682b464f9867a7c373f9881d092b834767dcabe5baf8c7245e937c/README.md) |
| 17&#160;Mar&#160;21&#160;17:19&#160;UTC | South Central Rural Telecommunications Coop SHAKEN Cert 0418 | 16&#160;Mar&#160;24&#160;17:19&#160;UTC | true | [view](CERTS/6e988e3a74dbb6c1862118841d78e438f345c1ef7e79a96c2087328b5de146c1/README.md) |
| 17&#160;Mar&#160;21&#160;17:25&#160;UTC | Yelcot SHAKEN Cert 1733 | 16&#160;Mar&#160;24&#160;17:25&#160;UTC | true | [view](CERTS/f7e9897313ee276a419725d0aa81886e8f3636ad1cd1e9aad623166f56e6b141/README.md) |
| 24&#160;Mar&#160;21&#160;23:52&#160;UTC | Cspire SHAKEN Cert 6581 | 23&#160;Mar&#160;24&#160;23:52&#160;UTC | true | [view](CERTS/09ed5b3292b5bfc7ac80b1027a827138b9503aa8053a61431a8dc851ecad04f2/README.md) |
| 01&#160;Apr&#160;21&#160;15:28&#160;UTC | USCellular SHAKEN Cert 6349 | 31&#160;Mar&#160;24&#160;15:28&#160;UTC | true | [view](CERTS/c75937e7c843b0054b1a51f6149f47661696b3cfcbcf6bab4dbc708c4453244f/README.md) |
| 07&#160;Apr&#160;21&#160;16:44&#160;UTC | U. S. Telepacific Corp SHAKEN 7453 | 06&#160;Apr&#160;24&#160;16:44&#160;UTC | true | [view](CERTS/9ed03dac797a5a27d52aa5209a4caa6a3ec9c3943d55a2cbfb69416480787da0/README.md) |
| 16&#160;Apr&#160;21&#160;15:27&#160;UTC | CBTS Technology Solutions SHAKEN Cert 600F | 15&#160;Apr&#160;24&#160;15:27&#160;UTC | true | [view](CERTS/3d02021a2da14f1ebfe588256a419be9ebc03c0d1fccc51cc29fa9d4a625c6bf/README.md) |
| 30&#160;Apr&#160;21&#160;17:05&#160;UTC | Hunter Communications Shaken Cert 660C | 29&#160;Apr&#160;24&#160;17:05&#160;UTC | true | [view](CERTS/0d2022504ffa5407f662990a785786cb0da72ce014838e2cffdcd95cb70c6f64/README.md) |
| 05&#160;May&#160;21&#160;21:02&#160;UTC | American Broadband SHAKEN Cert 355D | 04&#160;May&#160;24&#160;21:02&#160;UTC | true | [view](CERTS/f2fe275fba183def77918369f90d81bc080f1ef58c5422620c3dc45140b2ae75/README.md) |
| 18&#160;May&#160;21&#160;13:14&#160;UTC | Segra SHAKEN Cert 1784 | 17&#160;May&#160;24&#160;13:14&#160;UTC | true | [view](CERTS/f802b8b879d063b87665d8b2a67f6d3ba78a94aa35782f664a6a40afc7586f56/README.md) |
| 18&#160;May&#160;21&#160;13:22&#160;UTC | Viaero Wireless SHAKEN Cert 6874 | 17&#160;May&#160;24&#160;13:22&#160;UTC | true | [view](CERTS/fe826427006707276b684368c81c7b2039d4ce72b25c5cf0a0c039993b026573/README.md) |
| 20&#160;May&#160;21&#160;22:04&#160;UTC | Appalachian Wireless SHAKEN Cert 6940 | 19&#160;May&#160;24&#160;22:04&#160;UTC | true | [view](CERTS/e14170c681e75c37d0ca45e304c09cc0d148246bd7d72e96f91f7a8fe27339fa/README.md) |
| 20&#160;May&#160;21&#160;22:13&#160;UTC | Nemont SHAKEN Cert 2247 | 19&#160;May&#160;24&#160;22:13&#160;UTC | true | [view](CERTS/cd4fc6aff73ae9a3e16063ce1bda6bc6e265584b0b930d24ef615174fa6bed20/README.md) |
| 20&#160;May&#160;21&#160;22:20&#160;UTC | Everstream SHAKEN Cert 472C  | 19&#160;May&#160;24&#160;22:20&#160;UTC | true | [view](CERTS/8710bb38debebd39698fb1c273409b173951cca1fab53a6d4c4aca91e61e06df/README.md) |
| 26&#160;May&#160;21&#160;15:37&#160;UTC | WOW Internet Cable and Phone SHAKEN Cert 665E | 25&#160;May&#160;24&#160;15:37&#160;UTC | true | [view](CERTS/9f86e038b4f6c3bf7fabfe8d4008b071f81bd33a80c55ea66a1aadd79eb9b989/README.md) |
| 27&#160;May&#160;21&#160;13:28&#160;UTC | Mediacom 846F | 26&#160;May&#160;24&#160;13:28&#160;UTC | true | [view](CERTS/e6c9e9fd411d8174b3ffe1af4d569c6919f4b98a5d0c6e429cd3682d82284e7e/README.md) |
| 11&#160;Jun&#160;21&#160;16:01&#160;UTC | Clearwave SHAKEN Cert 9915 | 10&#160;Jun&#160;24&#160;16:01&#160;UTC | true | [view](CERTS/de0c13b76c24e2ef34c5de58e88fbfeef9c8adb44175b42f3c69771db6586491/README.md) |
| 11&#160;Jun&#160;21&#160;16:14&#160;UTC | Carolina West Wireless SHAKEN Cert 5932 | 10&#160;Jun&#160;24&#160;16:14&#160;UTC | true | [view](CERTS/e85dec068c9fa23af68b345257d58269303c925d10fc9101b226d1e0e7d62a9d/README.md) |
| 15&#160;Jun&#160;21&#160;13:56&#160;UTC | ENA SHAKEN cert 521F | 14&#160;Jun&#160;24&#160;13:56&#160;UTC | true | [view](CERTS/a5293c77ed92fdeef1aeff04f6bed7932a6083abc55d03d74a3ec5b817fdba2b/README.md) |
| 21&#160;Jun&#160;21&#160;15:05&#160;UTC | Buckeye SHAKEN Cert 7608 | 20&#160;Jun&#160;24&#160;15:05&#160;UTC | true | [view](CERTS/61a18f942b4df978a21d6e9cf1df2f0cea0dbd9609b1fcfec8c5f97d5c6dcd7e/README.md) |
| 30&#160;Jun&#160;21&#160;16:55&#160;UTC | Sonic Telecom SHAKEN cert 433E | 29&#160;Jun&#160;24&#160;16:55&#160;UTC | true | [view](CERTS/a5082b808b3bf200a634e273b7988a617099c9c897a9c596858e6d4ffe4fd352/README.md) |
| 30&#160;Jun&#160;21&#160;16:59&#160;UTC | Utility SHAKEN Cert 9262 | 29&#160;Jun&#160;24&#160;16:59&#160;UTC | true | [view](CERTS/97acbe1f0c7b2ceb99bec93a9da9ff2c1fd55e9a4cd78a32160237751f720236/README.md) |
| 20&#160;Jul&#160;21&#160;19:54&#160;UTC | New Horizon SHAKEN Cert 127E | 19&#160;Jul&#160;24&#160;19:54&#160;UTC | true | [view](CERTS/5b2842e49ecc543187018171fa660e32bdba390a7977c0de97a00aef35b8ae01/README.md) |
| 30&#160;Jul&#160;21&#160;17:16&#160;UTC | Syringa Networks SHAKEN Cert 5869 | 29&#160;Jul&#160;24&#160;17:16&#160;UTC | true | [view](CERTS/5b9b091228790e797adf24925cd1f2508b8372b744611f843a8026a703993d5a/README.md) |
| 16&#160;Sep&#160;21&#160;13:09&#160;UTC | CTS Telecom, Inc SHAKEN Cert 8331 | 15&#160;Sep&#160;24&#160;13:09&#160;UTC | true | [view](CERTS/5c9f6132a96a05a58d02d845ee70e914659682623fb98acbbecd9f8e383dcbec/README.md) |
| 05&#160;May&#160;22&#160;14:14&#160;UTC | Kaplan Telephone SHAKEN cert 0432 | 04&#160;May&#160;25&#160;14:14&#160;UTC | true | [view](CERTS/24c1b7c4dc4aeda21b5d1c5ff7f059903693c4a8b67ecc054ef05542a2ff4c35/README.md) |
| 07&#160;Jun&#160;22&#160;12:24&#160;UTC | Avid Communication SHAKEN Cert 742D | 06&#160;Jun&#160;25&#160;12:24&#160;UTC | true | [view](CERTS/b63d54026dfcdfd16495ad6fdda8993de182c86b4aa870784177c38c53842cba/README.md) |
| 12&#160;Oct&#160;22&#160;17:52&#160;UTC | Northeast Oklahoma Electric Cooperative SHAKEN Cert 945H | 11&#160;Oct&#160;25&#160;17:52&#160;UTC | true | [view](CERTS/b399b86f53e35dfa37c4cd7b28ee0132d934ef73354f564636d8edee42d58ccd/README.md) |

#### CA Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 25&#160;Nov&#160;20&#160;11:21&#160;UTC | Metaswitch STI-CA SHAKEN Root | 20&#160;Nov&#160;40&#160;11:21&#160;UTC | true | [view](CERTS/c27184f75f81fc119b85d51d416477bf635040e5d66ce6051799b37b9aa17485/README.md) |
| 25&#160;Nov&#160;20&#160;11:57&#160;UTC | Metaswitch STI-CA SHAKEN Issuing 1 | 22&#160;Nov&#160;32&#160;11:57&#160;UTC | true | [view](CERTS/b91a9874fbefc3feda9d5f9bd336e8b999c9b15b25aae7fe3c61d87373a5d1a1/README.md) |


Generated: 17 Dec 22 12:22 UTC