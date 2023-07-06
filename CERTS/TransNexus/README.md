# STIR/SHAKEN CA Ecosystem Compliance

## TransNexus

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 3206 certificates were included in the corpus being tested
- 2 certificates in the corpus were skipped because they are duplicates
- 3102 certificates in the corpus were skipped because they are expired
- 1 certificates in the corpus were skipped because they are not currently trusted
- 101 certificates being tested against the remaining rules
- 1.08 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 3.96% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 60 days is the average remaining validity for the certificates in the corpus
- 61 days is the average initial validity for the certificates in the corpus
- 86 certificates expire in the next 30 days
- 1.58 average number of unexpired certificates per OCN observed
- 64 unique OCNs observed in unexpired and valid certificate corpus

| Instances | Test | Source |
|-----------|------|--------|
| 4 | [e_incorrect_ku_encoding](ISSUES/e_incorrect_ku_encoding/README.md) | RFC5280 |
| 101 | [e_us_cp_subject_sn](ISSUES/e_us_cp_subject_sn/README.md) | US_SHAKEN_CP |
| 4 | [w_shaken_subject_rdn_unknown](ISSUES/w_shaken_subject_rdn_unknown/README.md) | SHAKEN_PKI_BEST_PRACTICES |

#### CA Certificates

- 6 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 0 certificates in the corpus were skipped because they are expired
- 1 certificates in the corpus were skipped because they are not currently trusted
- 5 certificates being tested against the remaining rules
- 3.00 issues on average found in unexpired, trusted, and non-compliant certificates
- 60.00% of certificates contain one or more Error level issue
- 60.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 60.00% of certificates are too old to be assessed against currently enforced expectations
- 5436 days is the average remaining validity for the certificates in the corpus
- 5113 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days

| Instances | Test | Source |
|-----------|------|--------|
| 3 | [e_incorrect_ku_encoding](ISSUES/e_incorrect_ku_encoding/README.md) | RFC5280 |
| 3 | [e_us_cp_ca_key_usage_crl_sign](ISSUES/e_us_cp_ca_key_usage_crl_sign/README.md) | US_SHAKEN_CP |
| 3 | [w_shaken_ca_subject_rdn_unknown](ISSUES/w_shaken_ca_subject_rdn_unknown/README.md) | SHAKEN_PKI_BEST_PRACTICES |

### Details

#### Leaf Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 20&#160;Jul&#160;22&#160;17:00&#160;UTC | SHAKEN 737J | 20&#160;Jul&#160;23&#160;17:00&#160;UTC | true | [view](CERTS/23db796295ad1249fcc3c9a42417c22f38be134d3b3b30082d35eab7cc153e5d/README.md) |
| 25&#160;Jul&#160;22&#160;18:36&#160;UTC | SHAKEN 578J | 25&#160;Jul&#160;23&#160;18:36&#160;UTC | true | [view](CERTS/f4af5a28131d680ef2682ae86181b64f0c919864d7dbaef00c0598d314b537a1/README.md) |
| 10&#160;Aug&#160;22&#160;18:11&#160;UTC | SHAKEN 073H | 10&#160;Aug&#160;23&#160;18:11&#160;UTC | true | [view](CERTS/87dabb91a689a6ac540b5ffbf5eb89326fada83fb822d2bb9ec31b68251b579e/README.md) |
| 13&#160;Sep&#160;22&#160;17:24&#160;UTC | SHAKEN 706J | 13&#160;Sep&#160;23&#160;17:24&#160;UTC | true | [view](CERTS/1eea5e17ae8b36cb5b6d657178273af745cea015e63586e6d44d1ba2ee8f722a/README.md) |
| 07&#160;Dec&#160;22&#160;00:45&#160;UTC | SHAKEN 299K | 07&#160;Dec&#160;23&#160;00:45&#160;UTC | true | [view](CERTS/7b5e04ff01ace2ee384df3507886684204901bf23851866eb6b3891ecdd96a92/README.md) |
| 31&#160;Dec&#160;22&#160;03:33&#160;UTC | SHAKEN 8526 | 31&#160;Dec&#160;23&#160;03:33&#160;UTC | true | [view](CERTS/96672cb8cc7c71e38d9e6350eac6c4f5c430c1eefeabe00cc095527e20a20c9b/README.md) |
| 16&#160;Jan&#160;23&#160;17:54&#160;UTC | SHAKEN 622J | 15&#160;Jul&#160;23&#160;17:54&#160;UTC | true | [view](CERTS/4c659921cd1ec39c239ab31d4c1ca813017c01d51512ea11a1fe0572f034c534/README.md) |
| 25&#160;Jan&#160;23&#160;14:28&#160;UTC | SHAKEN 873J | 25&#160;Jan&#160;24&#160;14:28&#160;UTC | true | [view](CERTS/36e9ba32d337ad13909bee0b62125c9769a049c4ae10d749687d75af4e49387d/README.md) |
| 27&#160;Mar&#160;23&#160;17:18&#160;UTC | SHAKEN 2720 | 26&#160;Mar&#160;24&#160;17:18&#160;UTC | true | [view](CERTS/7282c0683ebff27881b1ccd1f66664c3102340ff657d9363fcecb9c1b7924c12/README.md) |
| 10&#160;Apr&#160;23&#160;12:54&#160;UTC | SHAKEN 815G | 09&#160;Jul&#160;23&#160;12:54&#160;UTC | true | [view](CERTS/71206dbb241c8dbf9de187fca7b7e515fff61fbedcd1baacdc4a6a399ca2f887/README.md) |
| 28&#160;Apr&#160;23&#160;17:50&#160;UTC | SHAKEN 663J | 27&#160;Apr&#160;24&#160;17:50&#160;UTC | true | [view](CERTS/bd944d8e1acfeaaa520a8c87826b19a8509b52b083525a940e9ee4db8af1a99b/README.md) |
| 08&#160;May&#160;23&#160;17:37&#160;UTC | SHAKEN 4036 | 04&#160;Nov&#160;23&#160;17:37&#160;UTC | true | [view](CERTS/1f7c5f89be3bc2a774c0b657a342d25a23ec7f07e38c8e3ed1d45b720641d919/README.md) |
| 24&#160;May&#160;23&#160;08:31&#160;UTC | SHAKEN 736J | 23&#160;May&#160;24&#160;08:31&#160;UTC | true | [view](CERTS/d273815c57e441d2066144a02d029b73c3d839f5d31cbaeb72a7b6bded62da53/README.md) |
| 06&#160;Jun&#160;23&#160;13:41&#160;UTC | SHAKEN 6628 | 03&#160;Dec&#160;23&#160;13:41&#160;UTC | true | [view](CERTS/1892b1588772ce058cea6d06c2e82ea66fae4902994a60ab326d96de2635e22d/README.md) |
| 06&#160;Jun&#160;23&#160;17:02&#160;UTC | SHAKEN 551G | 06&#160;Jul&#160;23&#160;17:02&#160;UTC | true | [view](CERTS/367a17ac6576b27bc6fc09c5f3d224cf0cea55bba50e9d08e57e5c73f052d2e4/README.md) |
| 06&#160;Jun&#160;23&#160;19:10&#160;UTC | SHAKEN 597J | 05&#160;Jun&#160;24&#160;19:10&#160;UTC | true | [view](CERTS/7f15a14bb0b8ad72d450e673344a747703cdbe15741d2696c371c83c5ca8e148/README.md) |
| 07&#160;Jun&#160;23&#160;19:34&#160;UTC | SHAKEN 577F | 07&#160;Jul&#160;23&#160;19:34&#160;UTC | true | [view](CERTS/c8079556354af7299281f80299c1c43da13fcda99b23bbb4bbe5da392d2da510/README.md) |
| 10&#160;Jun&#160;23&#160;04:54&#160;UTC | SHAKEN 345J | 10&#160;Jul&#160;23&#160;04:54&#160;UTC | true | [view](CERTS/3f99c6aa042be2a8d29b2528437dd2e0c8164e072f799dbc533f6465e34d261e/README.md) |
| 10&#160;Jun&#160;23&#160;05:31&#160;UTC | SHAKEN 345J | 10&#160;Jul&#160;23&#160;05:31&#160;UTC | true | [view](CERTS/15a957bb919f112ff481c7bd1dc5b6d36879b001b424d90cc6d85174a7c3578a/README.md) |
| 10&#160;Jun&#160;23&#160;19:04&#160;UTC | SHAKEN 952J | 10&#160;Jul&#160;23&#160;19:04&#160;UTC | true | [view](CERTS/7162e69a8140b434a3543cb2440619683e8156b87931944c1e86f8a178640365/README.md) |
| 10&#160;Jun&#160;23&#160;21:08&#160;UTC | SHAKEN 193E | 09&#160;Aug&#160;23&#160;21:08&#160;UTC | true | [view](CERTS/5bce05428a294fc2a9639e4f657c2ce3484e4b3fbbd2a29eb02780edb41cf608/README.md) |
| 15&#160;Jun&#160;23&#160;14:33&#160;UTC | SHAKEN 366G | 15&#160;Jul&#160;23&#160;14:33&#160;UTC | true | [view](CERTS/6e4100783b07c6edaa44e32920e7772ada456c4faae54661024b52b49fe3711c/README.md) |
| 15&#160;Jun&#160;23&#160;20:10&#160;UTC | SHAKEN 505J | 13&#160;Sep&#160;23&#160;20:10&#160;UTC | true | [view](CERTS/81d039b45fcc19c802cbe6760a256f6b17c2ce2ebecefbf0120e6bc14093bba4/README.md) |
| 15&#160;Jun&#160;23&#160;23:31&#160;UTC | SHAKEN 4632 | 15&#160;Jul&#160;23&#160;23:31&#160;UTC | true | [view](CERTS/96de1ec51cab748ac3e7651428780ec6d0263467a83741d5a6808a93c2cfee92/README.md) |
| 16&#160;Jun&#160;23&#160;18:51&#160;UTC | SHAKEN 857J | 15&#160;Jun&#160;24&#160;18:51&#160;UTC | true | [view](CERTS/97834968950fa42accd7ab2d9f43d12eb3dd919b5d5a72a153275419104aa534/README.md) |
| 22&#160;Jun&#160;23&#160;16:43&#160;UTC | SHAKEN 722J | 22&#160;Jul&#160;23&#160;16:43&#160;UTC | true | [view](CERTS/3a4197f478ead0901e00ef6b6357303672520e0806a7aae865d38d94ba42950c/README.md) |
| 25&#160;Jun&#160;23&#160;13:15&#160;UTC | SHAKEN 807J | 24&#160;Aug&#160;23&#160;13:15&#160;UTC | true | [view](CERTS/e35d163bfe2822eb6b9914458ea238ebc9c114c1d69533a90a72b3e9a98832ba/README.md) |
| 25&#160;Jun&#160;23&#160;19:41&#160;UTC | SHAKEN 577F | 25&#160;Jul&#160;23&#160;19:41&#160;UTC | true | [view](CERTS/167d37eb343deb2923d597d60e3e0d48220f034e566903030524c211405a5200/README.md) |
| 27&#160;Jun&#160;23&#160;17:04&#160;UTC | SHAKEN 551G | 27&#160;Jul&#160;23&#160;17:04&#160;UTC | true | [view](CERTS/6faab7cc1017c017eb573361376b302f764bb14d06c6cd81bc51f05d0e200b91/README.md) |
| 28&#160;Jun&#160;23&#160;19:10&#160;UTC | SHAKEN 952J | 28&#160;Jul&#160;23&#160;19:10&#160;UTC | true | [view](CERTS/9ef0498dbb3d2e81e66887f66cf1ac7097e2d9a09c5f3c2443d9b0a0e8ef9fbd/README.md) |
| 29&#160;Jun&#160;23&#160;11:12&#160;UTC | SHAKEN 841J | 13&#160;Jul&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/610ca037c1e58df40e2381e6e2a553920e206421d26b5f967cc6106c45b9fced/README.md) |
| 29&#160;Jun&#160;23&#160;15:00&#160;UTC | SHAKEN 012K | 06&#160;Jul&#160;23&#160;15:00&#160;UTC | true | [view](CERTS/33f9566617cbf51c76b99a8065bce2b37a7fa89c388d71582e224ab1921964df/README.md) |
| 29&#160;Jun&#160;23&#160;15:03&#160;UTC | SHAKEN 056J | 06&#160;Jul&#160;23&#160;15:03&#160;UTC | true | [view](CERTS/02768b604e1d5831be722ceefb911331babedcfc0674f8bea8f64e0057c4b72d/README.md) |
| 29&#160;Jun&#160;23&#160;17:47&#160;UTC | SHAKEN 107K | 06&#160;Jul&#160;23&#160;17:47&#160;UTC | true | [view](CERTS/1f4d98785c6bf14b8169842a89eed6765218db8260208b3b6d3ab7e09ca94f3f/README.md) |
| 29&#160;Jun&#160;23&#160;18:59&#160;UTC | SHAKEN 089K | 06&#160;Jul&#160;23&#160;18:59&#160;UTC | true | [view](CERTS/088782651883f8ddfb2b619f181dca9e793266f977147948565d62733c25c486/README.md) |
| 29&#160;Jun&#160;23&#160;19:49&#160;UTC | SHAKEN 604K | 06&#160;Jul&#160;23&#160;19:49&#160;UTC | true | [view](CERTS/fae359e267898718865bba0f3443e2dd3a4504dd95f22af22556cea0725e1be0/README.md) |
| 29&#160;Jun&#160;23&#160;20:40&#160;UTC | SHAKEN 674J | 06&#160;Jul&#160;23&#160;20:40&#160;UTC | true | [view](CERTS/f562203607f88310021f58aa09c59395f94678d08ad74692ef7e8023b0649b44/README.md) |
| 29&#160;Jun&#160;23&#160;20:43&#160;UTC | SHAKEN 738J | 06&#160;Jul&#160;23&#160;20:42&#160;UTC | true | [view](CERTS/d65b7ded82c7fa8e66c4ea57966073a9341414c8012001959e882ea305efa2ca/README.md) |
| 29&#160;Jun&#160;23&#160;20:44&#160;UTC | SHAKEN 700H | 06&#160;Jul&#160;23&#160;20:44&#160;UTC | true | [view](CERTS/24885ee9b6a46fa564e9b3f55de111cb04d76ca83360ceff0642e53ace54fee0/README.md) |
| 29&#160;Jun&#160;23&#160;20:44&#160;UTC | SHAKEN 733J | 06&#160;Jul&#160;23&#160;20:44&#160;UTC | true | [view](CERTS/eae317c213a4632d0defb9c608d68b55629e80e24986938305149e2385547a9a/README.md) |
| 29&#160;Jun&#160;23&#160;20:46&#160;UTC | SHAKEN 769J | 06&#160;Jul&#160;23&#160;20:46&#160;UTC | true | [view](CERTS/32bf8b52092d89ce1c00adefd7b2f4308f3dba06f7d7fdda6616cff2b6693ac8/README.md) |
| 29&#160;Jun&#160;23&#160;20:49&#160;UTC | SHAKEN 469A | 06&#160;Jul&#160;23&#160;20:49&#160;UTC | true | [view](CERTS/98cc07428d3ffd79f4658621be33fb1159cba331f9a7058569527e30e8493a52/README.md) |
| 29&#160;Jun&#160;23&#160;20:49&#160;UTC | SHAKEN 849J | 06&#160;Jul&#160;23&#160;20:49&#160;UTC | true | [view](CERTS/23101b8bc3c4df35c58192c692939f9580647956e58403f5726e668762e22c4e/README.md) |
| 29&#160;Jun&#160;23&#160;20:49&#160;UTC | SHAKEN 790J | 06&#160;Jul&#160;23&#160;20:49&#160;UTC | true | [view](CERTS/f57b3e66ffe1da7ae9672071aa146fc4f1e7719f72351e9a064ea824124908f3/README.md) |
| 29&#160;Jun&#160;23&#160;20:50&#160;UTC | SHAKEN 625J | 06&#160;Jul&#160;23&#160;20:50&#160;UTC | true | [view](CERTS/16f022c715af8b9778448bcba2b5c982ac57d2b50cf283ca5c8a29f2a72eeb56/README.md) |
| 29&#160;Jun&#160;23&#160;20:50&#160;UTC | SHAKEN 738J | 06&#160;Jul&#160;23&#160;20:50&#160;UTC | true | [view](CERTS/637bbb4b54fd211d2cb579e91de0da005d78b5ff54301cdeef2918a5bf930779/README.md) |
| 29&#160;Jun&#160;23&#160;20:50&#160;UTC | SHAKEN 459J | 06&#160;Jul&#160;23&#160;20:50&#160;UTC | true | [view](CERTS/7f0ae616f00494ec3ddba992688a9325a7d9d6e8c5b9863b38e3856997e5081a/README.md) |
| 29&#160;Jun&#160;23&#160;20:51&#160;UTC | SHAKEN 366G | 06&#160;Jul&#160;23&#160;20:51&#160;UTC | true | [view](CERTS/074ac35ce3199d1153285c0432e8c63e275d2cf702337709b2cf9f4f84879878/README.md) |
| 29&#160;Jun&#160;23&#160;20:51&#160;UTC | SHAKEN 738J | 06&#160;Jul&#160;23&#160;20:51&#160;UTC | true | [view](CERTS/910325f876013d80a5d8a01f69960dbfac74a3a0b4ae304e1e89186080a8f1af/README.md) |
| 30&#160;Jun&#160;23&#160;13:52&#160;UTC | SHAKEN 1373 | 07&#160;Jul&#160;23&#160;13:52&#160;UTC | true | [view](CERTS/f4c3b12328fc5c0f7c1c332e9ed0e57e1da8624b527eeae2e99d16777e7e01b2/README.md) |
| 30&#160;Jun&#160;23&#160;19:26&#160;UTC | SHAKEN 749J | 07&#160;Jul&#160;23&#160;19:26&#160;UTC | true | [view](CERTS/56017b2ebb5f647e3e544d4aa2b101bf1665a6d673f7004fc15dd4b6415dc89d/README.md) |
| 30&#160;Jun&#160;23&#160;22:03&#160;UTC | SHAKEN 606F | 07&#160;Jul&#160;23&#160;22:03&#160;UTC | true | [view](CERTS/18f492f63b1103a6bd8f4ea774c10de51c92c3f603e1b47da24171f55da0f17f/README.md) |
| 01&#160;Jul&#160;23&#160;05:00&#160;UTC | SHAKEN 345J | 31&#160;Jul&#160;23&#160;05:00&#160;UTC | true | [view](CERTS/5fe65f7ce04a1b757d41bc3c0be904bf9082f0e23d5d325399bb3207836c4434/README.md) |
| 01&#160;Jul&#160;23&#160;05:31&#160;UTC | SHAKEN 345J | 31&#160;Jul&#160;23&#160;05:31&#160;UTC | true | [view](CERTS/0d1d61101358786277c695342a2362b2ae2284ed2d1a1fe0def8a0da0ff1b1a2/README.md) |
| 01&#160;Jul&#160;23&#160;11:47&#160;UTC | SHAKEN 1577 | 08&#160;Jul&#160;23&#160;11:47&#160;UTC | true | [view](CERTS/455698d2138280b6563d36eab00d49849a07ce8bff2da0f41cf43fc158783c62/README.md) |
| 01&#160;Jul&#160;23&#160;11:48&#160;UTC | SHAKEN 1577 | 08&#160;Jul&#160;23&#160;11:48&#160;UTC | true | [view](CERTS/9a02df720ee9a0b78f26812b73be273a9a7bb8b9edf5fe8b301c00d86621eb2a/README.md) |
| 01&#160;Jul&#160;23&#160;14:54&#160;UTC | SHAKEN 2550 | 08&#160;Jul&#160;23&#160;14:54&#160;UTC | true | [view](CERTS/35878226b518704613e6956d91cd1b18cd7661166c5ace68f02b79ec42a08577/README.md) |
| 01&#160;Jul&#160;23&#160;16:36&#160;UTC | SHAKEN 753J | 08&#160;Jul&#160;23&#160;16:36&#160;UTC | true | [view](CERTS/9aa91f27168b8462c1817c98b4212337afe5373091fe9d1a32ebaf6831cfaf90/README.md) |
| 01&#160;Jul&#160;23&#160;20:34&#160;UTC | SHAKEN 177K | 08&#160;Jul&#160;23&#160;20:34&#160;UTC | true | [view](CERTS/7acba1e8e2069b169cca50055f92cc24a8c9c680a6af26ae963c5a4eccc52324/README.md) |
| 02&#160;Jul&#160;23&#160;17:47&#160;UTC | SHAKEN 107K | 09&#160;Jul&#160;23&#160;17:47&#160;UTC | true | [view](CERTS/b30dfae7530aa7a3252b24f515aa51df7d801deff41add7aab17c8ad7aee384c/README.md) |
| 02&#160;Jul&#160;23&#160;18:59&#160;UTC | SHAKEN 089K | 09&#160;Jul&#160;23&#160;18:59&#160;UTC | true | [view](CERTS/22061a5c51103db2acee2e8529f45036461078f06721cf5ccc22b8c554089113/README.md) |
| 02&#160;Jul&#160;23&#160;19:49&#160;UTC | SHAKEN 604K | 09&#160;Jul&#160;23&#160;19:49&#160;UTC | true | [view](CERTS/002d465ba7d6cbd8ed8bfaac763521b7ee898086383108ab8d8f221488678856/README.md) |
| 02&#160;Jul&#160;23&#160;20:18&#160;UTC | SHAKEN 297K | 09&#160;Jul&#160;23&#160;20:18&#160;UTC | true | [view](CERTS/f8c5ffe32161f3e25728cf4dc2474308c304cdd6a90af8886b5b7c90bed8264e/README.md) |
| 02&#160;Jul&#160;23&#160;20:41&#160;UTC | SHAKEN 674J | 09&#160;Jul&#160;23&#160;20:41&#160;UTC | true | [view](CERTS/c5e432ec9d6baf90eb78c4187618d1ead5ff0d894bcf3d7f91f75213613b7c56/README.md) |
| 02&#160;Jul&#160;23&#160;20:42&#160;UTC | SHAKEN 735J | 09&#160;Jul&#160;23&#160;20:42&#160;UTC | true | [view](CERTS/cb398a10b03fbe54ead94c3255cb7734d9305270fd763abe8e8072b1b2ba74d8/README.md) |
| 02&#160;Jul&#160;23&#160;20:43&#160;UTC | SHAKEN 738J | 09&#160;Jul&#160;23&#160;20:43&#160;UTC | true | [view](CERTS/bc41a9558f2a336a752c7cb7592dde8d102c44aabc500bcaee362eba8e3a6adf/README.md) |
| 02&#160;Jul&#160;23&#160;20:44&#160;UTC | SHAKEN 700H | 09&#160;Jul&#160;23&#160;20:44&#160;UTC | true | [view](CERTS/e49d7dfdfc0175db800186eb21ad4d6a8c717dcbe6212e12ab99839ce9c58141/README.md) |
| 02&#160;Jul&#160;23&#160;20:44&#160;UTC | SHAKEN 856H | 09&#160;Jul&#160;23&#160;20:44&#160;UTC | true | [view](CERTS/78d2effddc1842007fd9f85f2662f2741937a8ecf9fe62c0a7783a96b9babf00/README.md) |
| 02&#160;Jul&#160;23&#160;20:45&#160;UTC | SHAKEN 819H | 09&#160;Jul&#160;23&#160;20:45&#160;UTC | true | [view](CERTS/60256e81306298f86cfece555381bf580aacf3905b32ba529524f221891b0e88/README.md) |
| 02&#160;Jul&#160;23&#160;20:46&#160;UTC | SHAKEN 0753 | 09&#160;Jul&#160;23&#160;20:46&#160;UTC | true | [view](CERTS/f74a2539bcca32b1fbc3106965255b7e32f286d746d898e2caf1ef52298ab10f/README.md) |
| 02&#160;Jul&#160;23&#160;20:46&#160;UTC | SHAKEN 769J | 09&#160;Jul&#160;23&#160;20:46&#160;UTC | true | [view](CERTS/43caa32db46498e85df20e7cb2aca0b9d7b00bbd87ee37eeb50bdf6672bebb14/README.md) |
| 02&#160;Jul&#160;23&#160;20:50&#160;UTC | SHAKEN 849J | 09&#160;Jul&#160;23&#160;20:50&#160;UTC | true | [view](CERTS/48644caec3114dc24f0f9348680801963c444eda88550ef5462c109cd173e5dd/README.md) |
| 02&#160;Jul&#160;23&#160;20:50&#160;UTC | SHAKEN 469A | 09&#160;Jul&#160;23&#160;20:50&#160;UTC | true | [view](CERTS/031df76aa9cb96f0d01ce40fe399cb5adf38a1284f362cbc6c9ad49074eee015/README.md) |
| 02&#160;Jul&#160;23&#160;20:50&#160;UTC | SHAKEN 790J | 09&#160;Jul&#160;23&#160;20:50&#160;UTC | true | [view](CERTS/12c1686f9a381ee72dfc9a307c12d7247c34246b06c22802209e0b0964810f9e/README.md) |
| 02&#160;Jul&#160;23&#160;20:51&#160;UTC | SHAKEN 738J | 09&#160;Jul&#160;23&#160;20:51&#160;UTC | true | [view](CERTS/efa6c0802df8a205f318e0ca6b5622e9725805ebfd948d27892361e3c1f12774/README.md) |
| 02&#160;Jul&#160;23&#160;20:51&#160;UTC | SHAKEN 459J | 09&#160;Jul&#160;23&#160;20:51&#160;UTC | true | [view](CERTS/d346ebfbb9a70044824a20fbbd40f5abe19fe5c26904073198a95ba8898ddeb9/README.md) |
| 02&#160;Jul&#160;23&#160;20:52&#160;UTC | SHAKEN 366G | 09&#160;Jul&#160;23&#160;20:52&#160;UTC | true | [view](CERTS/a63d4e9deb2e98d80d3f7a539315de1f9959ea279e14a137f52d8851f23c99bf/README.md) |
| 02&#160;Jul&#160;23&#160;20:52&#160;UTC | SHAKEN 0226 | 09&#160;Jul&#160;23&#160;20:52&#160;UTC | true | [view](CERTS/4a551252956ef5880f4891da1127bdd9839874118a3b2c1353ae9e3ed2ead08d/README.md) |
| 02&#160;Jul&#160;23&#160;20:52&#160;UTC | SHAKEN 738J | 09&#160;Jul&#160;23&#160;20:52&#160;UTC | true | [view](CERTS/ca5cbcf4e96a29b9df492492370ae540c5284f818277c5cb5d1ada4ada990091/README.md) |
| 02&#160;Jul&#160;23&#160;21:14&#160;UTC | SHAKEN 1772 | 09&#160;Jul&#160;23&#160;21:14&#160;UTC | true | [view](CERTS/ad7aa19b303264cf34b3fc7d8a628373f46d23ef5689428e10d694c2d12a5dd8/README.md) |
| 03&#160;Jul&#160;23&#160;01:53&#160;UTC | SHAKEN 0172 | 10&#160;Jul&#160;23&#160;01:53&#160;UTC | true | [view](CERTS/0491c2cd6fcb77e28c8bec69679af96760f9d31ca29f20f0c903716e66170b23/README.md) |
| 03&#160;Jul&#160;23&#160;13:52&#160;UTC | SHAKEN 1373 | 10&#160;Jul&#160;23&#160;13:52&#160;UTC | true | [view](CERTS/97b650606de77a0a8a87f323af1f5e39d98ce4bcbe71d94e26cc348ab25e432d/README.md) |
| 03&#160;Jul&#160;23&#160;15:02&#160;UTC | SHAKEN 159H | 10&#160;Jul&#160;23&#160;15:02&#160;UTC | true | [view](CERTS/4e64814466a0cf54c108c1e5c74aea2716786a96953dc3f9dc0e9db48608106e/README.md) |
| 03&#160;Jul&#160;23&#160;19:26&#160;UTC | SHAKEN 749J | 10&#160;Jul&#160;23&#160;19:26&#160;UTC | true | [view](CERTS/f35e414bb8607975b1990aa23b881d5011be3628974565ee89d5c5da35eb252c/README.md) |
| 03&#160;Jul&#160;23&#160;20:36&#160;UTC | SHAKEN 983J | 10&#160;Jul&#160;23&#160;20:36&#160;UTC | true | [view](CERTS/120962a06c7c0f9d07a1bd48e4df8bde6cf98ab01cca9f62e3ed32e16ae9dfa5/README.md) |
| 03&#160;Jul&#160;23&#160;20:44&#160;UTC | SHAKEN 2252 | 10&#160;Jul&#160;23&#160;20:44&#160;UTC | true | [view](CERTS/ff615d103726daab320a4c162c7cad033de2d3b7052cd9082f3bb8db4e74a941/README.md) |
| 03&#160;Jul&#160;23&#160;20:56&#160;UTC | SHAKEN 738J | 12&#160;Jul&#160;23&#160;20:56&#160;UTC | true | [view](CERTS/a2c7b090bb0bf9c0beba49f56db7a42ff0069e31e1cbd936997da1db1535e5ec/README.md) |
| 03&#160;Jul&#160;23&#160;21:42&#160;UTC | SHAKEN 2251 | 10&#160;Jul&#160;23&#160;21:42&#160;UTC | true | [view](CERTS/aa9066ce0df90420d2102ab7cfcbf477c84b5990731c98dd8dbab132f61b1cf6/README.md) |
| 03&#160;Jul&#160;23&#160;23:44&#160;UTC | SHAKEN 4632 | 02&#160;Aug&#160;23&#160;23:44&#160;UTC | true | [view](CERTS/3f1fcb730728c4820b12ee73278e2b86a2fa6b3fcd920e52fe200c04ac6e68c6/README.md) |
| 04&#160;Jul&#160;23&#160;02:00&#160;UTC | SHAKEN 278K | 11&#160;Jul&#160;23&#160;02:00&#160;UTC | true | [view](CERTS/c8d8e07b48d962e12e0c6e52a3371568d52b0df7701757ab55d1bb2695568b94/README.md) |
| 04&#160;Jul&#160;23&#160;11:47&#160;UTC | SHAKEN 1577 | 11&#160;Jul&#160;23&#160;11:47&#160;UTC | true | [view](CERTS/2bc529a7a3a80f5378aecfee6595db20b046f121f4d39a40351bb3fc007dfecb/README.md) |
| 04&#160;Jul&#160;23&#160;14:54&#160;UTC | SHAKEN 2550 | 11&#160;Jul&#160;23&#160;14:54&#160;UTC | true | [view](CERTS/2019e9e3473274a75720f4b84d9497c73d816fcab97a537a21a3c5b8929fcce1/README.md) |
| 04&#160;Jul&#160;23&#160;15:41&#160;UTC | SHAKEN 284K | 11&#160;Jul&#160;23&#160;15:41&#160;UTC | true | [view](CERTS/ea37acfa3b53648d40180e501bc66c7091c520c0a1c732abb97e9f456c522ed1/README.md) |
| 04&#160;Jul&#160;23&#160;20:34&#160;UTC | SHAKEN 177K | 11&#160;Jul&#160;23&#160;20:34&#160;UTC | true | [view](CERTS/e692a78b6e023f928441f3365eb5ed6e28e230c0c2e9cd1dd413ce17206c21d4/README.md) |
| 05&#160;Jul&#160;23&#160;15:00&#160;UTC | SHAKEN 012K | 12&#160;Jul&#160;23&#160;15:00&#160;UTC | true | [view](CERTS/b2a7be26ab469a0853a7d4243b24316bc1a2f6b4908830e856c53bda56084ad5/README.md) |
| 05&#160;Jul&#160;23&#160;17:47&#160;UTC | SHAKEN 107K | 12&#160;Jul&#160;23&#160;17:47&#160;UTC | true | [view](CERTS/2c207400f15a3d52c41d3dcabab860ed03b38cf144faa024a52363f739cce597/README.md) |
| 05&#160;Jul&#160;23&#160;18:59&#160;UTC | SHAKEN 089K | 12&#160;Jul&#160;23&#160;18:59&#160;UTC | true | [view](CERTS/3dd9f69d48bf30a4fa4a3772bcf39e06a14cdc028a9470eae84bce99d1bf0391/README.md) |
| 05&#160;Jul&#160;23&#160;19:00&#160;UTC | SHAKEN 056K | 12&#160;Jul&#160;23&#160;19:00&#160;UTC | true | [view](CERTS/462b50f9c8b6a4db657bd52e3a76818c8370c6fb5e9ae405ac03248c93cf2311/README.md) |
| 05&#160;Jul&#160;23&#160;20:41&#160;UTC | SHAKEN 674J | 12&#160;Jul&#160;23&#160;20:41&#160;UTC | true | [view](CERTS/2dd99cefa78238b15d8f9f606b01982630e298e0948c272439190bdf403b5ef6/README.md) |
| 05&#160;Jul&#160;23&#160;20:43&#160;UTC | SHAKEN 738J | 12&#160;Jul&#160;23&#160;20:43&#160;UTC | true | [view](CERTS/cb5f732d807c8f6427db638210b5403765a077faa3526255b5f5e46ac529342a/README.md) |
| 05&#160;Jul&#160;23&#160;20:52&#160;UTC | SHAKEN 738J | 12&#160;Jul&#160;23&#160;20:52&#160;UTC | true | [view](CERTS/a713896f5ae903cc46e78312635ded22fd6dc7e4ccd214c6b41e4c5ff7784114/README.md) |

#### CA Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 20&#160;Aug&#160;21&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Issuing CA3 | 19&#160;Aug&#160;31&#160;23:59&#160;UTC | true | [view](CERTS/51b51431ffd8055cd7e30c213540f4868258058d4c4c9f23f1ed8bea75101e37/README.md) |
| 20&#160;Aug&#160;21&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Issuing CA1 | 19&#160;Aug&#160;31&#160;23:59&#160;UTC | true | [view](CERTS/037a7c6621814bd5f01fa2080a77a245458694d9d566e0050e068fc853317857/README.md) |
| 20&#160;Aug&#160;21&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Root CA1 | 19&#160;Aug&#160;41&#160;23:59&#160;UTC | true | [view](CERTS/59ab943f4e8c3c17755cc8e4abb5ac65736ce74ec587c62cb744ab0babbbe2fe/README.md) |
| 24&#160;Oct&#160;22&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Issuing CA4 | 23&#160;Oct&#160;32&#160;23:59&#160;UTC | false | [view](CERTS/80dd12d935f46c256c7c1289a0834b21ec8f2f8c35f2864928d1e75f3a280665/README.md) |
| 24&#160;Oct&#160;22&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Root CA2 | 23&#160;Oct&#160;42&#160;23:59&#160;UTC | false | [view](CERTS/a26e04fc786ab70b8085236b2c53f8cfbf5d0c6a5c2c9c3e9f91669fbb8ea4d5/README.md) |


Generated: 06 Jul 23 14:08 UTC