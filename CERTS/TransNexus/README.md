# STIR/SHAKEN CA Ecosystem Compliance

## TransNexus

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 794 certificates were included in the corpus being tested
- 1 certificates in the corpus were skipped because they are duplicates
- 687 certificates in the corpus were skipped because they are expired
- 1 certificates in the corpus were skipped because they are not currently trusted
- 105 certificates being tested against the remaining rules
- 1.38 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 17.14% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 0.95% of certificates are too old to be assessed against currently enforced expectations
- 52 days is the average remaining validity for the certificates in the corpus
- 52 days is the average initial validity for the certificates in the corpus
- 92 certificates expire in the next 30 days
- 1.44 average number of unexpired certificates per OCN observed
- 73 unique OCNs observed in unexpired and valid certificate corpus

| Instances | Test | Source |
|-----------|------|--------|
| 3 | [e_atis_certificate_policies](ISSUES/e_atis_certificate_policies/README.md) | ATIS1000080 |
| 16 | [e_incorrect_ku_encoding](ISSUES/e_incorrect_ku_encoding/README.md) | RFC5280 |
| 3 | [e_us_cp_ambiguous_identifier](ISSUES/e_us_cp_ambiguous_identifier/README.md) | US_SHAKEN_CP |
| 105 | [e_us_cp_subject_sn](ISSUES/e_us_cp_subject_sn/README.md) | US_SHAKEN_CP |
| 18 | [w_shaken_subject_rdn_unknown](ISSUES/w_shaken_subject_rdn_unknown/README.md) | SHAKEN_PKI_BEST_PRACTICES |

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
- 5470 days is the average remaining validity for the certificates in the corpus
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
| 13&#160;Jan&#160;22&#160;14:17&#160;UTC | MobileSphere SHAKEN 873J | 13&#160;Jan&#160;23&#160;14:17&#160;UTC | true | [view](CERTS/8f74ddd5e6964042b119d5b3f17d24694ec78c8688e0ffc96d13fdb0e8df26d4/README.md) |
| 30&#160;Mar&#160;22&#160;12:54&#160;UTC | Fusion Connect SHAKEN 2720 | 30&#160;Mar&#160;23&#160;12:54&#160;UTC | true | [view](CERTS/03a010b294807e90cc1309d7466fd2ddc591158a69be6950ff4d0eab32de0799/README.md) |
| 27&#160;Apr&#160;22&#160;18:22&#160;UTC | CCI SHAKEN 663J | 27&#160;Apr&#160;23&#160;18:22&#160;UTC | true | [view](CERTS/9cfd0b9974acdc9b5fce3888cf2f6cf99173e7955adeebd37a1e60dfc444a0b3/README.md) |
| 10&#160;Jun&#160;22&#160;14:00&#160;UTC | SHAKEN 597J | 10&#160;Jun&#160;23&#160;14:00&#160;UTC | true | [view](CERTS/4baa838eb96ce358aa1565d0add8b75072258efd5d21a5bbb0eec878026a67be/README.md) |
| 20&#160;Jun&#160;22&#160;20:25&#160;UTC | SHAKEN 857J | 20&#160;Jun&#160;23&#160;20:25&#160;UTC | true | [view](CERTS/e1e0a4851fb7b51d617959c00fb6a17e7b23d34e1555ba70ab01c9182a2f85f8/README.md) |
| 22&#160;Jun&#160;22&#160;18:46&#160;UTC | SHAKEN 755J | 22&#160;Jun&#160;23&#160;18:46&#160;UTC | true | [view](CERTS/a8cac316fddf25b48b9350a325791be8b87e6484a5959ab605db1ae5e3e99a84/README.md) |
| 29&#160;Jun&#160;22&#160;20:24&#160;UTC | SHAKEN 736J | 29&#160;Jun&#160;23&#160;20:24&#160;UTC | true | [view](CERTS/71b25a480077816998e48e6a76f272e70aa91ecb19b0e773cd8777823cdead32/README.md) |
| 25&#160;Jul&#160;22&#160;18:36&#160;UTC | SHAKEN 578J | 25&#160;Jul&#160;23&#160;18:36&#160;UTC | true | [view](CERTS/f4af5a28131d680ef2682ae86181b64f0c919864d7dbaef00c0598d314b537a1/README.md) |
| 10&#160;Aug&#160;22&#160;18:11&#160;UTC | SHAKEN 073H | 10&#160;Aug&#160;23&#160;18:11&#160;UTC | true | [view](CERTS/87dabb91a689a6ac540b5ffbf5eb89326fada83fb822d2bb9ec31b68251b579e/README.md) |
| 23&#160;Aug&#160;22&#160;04:14&#160;UTC | SHAKEN 866J | 19&#160;Feb&#160;23&#160;04:14&#160;UTC | true | [view](CERTS/4c4fdb320c51c8582d595c2d307cd770d409daf533ff573bcc73afaed83f6b7d/README.md) |
| 01&#160;Sep&#160;22&#160;03:25&#160;UTC | SHAKEN 6628 | 03&#160;Jan&#160;23&#160;03:25&#160;UTC | true | [view](CERTS/37b9a67bb8e3272048330f269b5dbc285ac0278df4670ab7e291b92b6548fa2d/README.md) |
| 13&#160;Sep&#160;22&#160;17:24&#160;UTC | SHAKEN 706J | 13&#160;Sep&#160;23&#160;17:24&#160;UTC | true | [view](CERTS/1eea5e17ae8b36cb5b6d657178273af745cea015e63586e6d44d1ba2ee8f722a/README.md) |
| 26&#160;Sep&#160;22&#160;18:40&#160;UTC | SHAKEN 505J | 25&#160;Dec&#160;22&#160;18:40&#160;UTC | true | [view](CERTS/b186c959554e652b69824a2e45ac08ec135105b0e01c9d2eadeb3cf46c130670/README.md) |
| 26&#160;Sep&#160;22&#160;18:42&#160;UTC | SHAKEN 505J | 25&#160;Dec&#160;22&#160;18:42&#160;UTC | true | [view](CERTS/1c0003ac10eeaa04229e507c15f71ee018c3902c3e8c20fbe42e533b8682ba8c/README.md) |
| 12&#160;Oct&#160;22&#160;19:39&#160;UTC | SHAKEN 815G | 10&#160;Jan&#160;23&#160;19:39&#160;UTC | true | [view](CERTS/a79f50d4fcd3e680518379f12a86a4b9a21c5e27dc5422a8439525edf350c4a7/README.md) |
| 20&#160;Oct&#160;22&#160;15:48&#160;UTC | SHAKEN 622J | 18&#160;Jan&#160;23&#160;15:48&#160;UTC | true | [view](CERTS/2e34c765f110b197f2ec1052b1960bc25ce320dfce7d7035cc532cd1eb6dbd57/README.md) |
| 31&#160;Oct&#160;22&#160;12:32&#160;UTC | SHAKEN 193E | 30&#160;Dec&#160;22&#160;12:32&#160;UTC | true | [view](CERTS/0de5e0e6787b9aa19c02c014821dd7a2e62e438f4eb9fbbb761eb86df8c69ff7/README.md) |
| 01&#160;Nov&#160;22&#160;11:14&#160;UTC | SHAKEN 807J | 31&#160;Dec&#160;22&#160;11:14&#160;UTC | true | [view](CERTS/0d97796c10ef21ba02a0f7954d82ace5e41d7ea66e4ea9eb16801c3a1a1d9f03/README.md) |
| 18&#160;Nov&#160;22&#160;15:42&#160;UTC | SHAKEN 722J | 18&#160;Dec&#160;22&#160;15:42&#160;UTC | true | [view](CERTS/4111bb56ba547a2f493d1e9c2baefa33ce9bd117904133477d063d6fa0d6b9b2/README.md) |
| 18&#160;Nov&#160;22&#160;18:04&#160;UTC | SHAKEN 952J | 18&#160;Dec&#160;22&#160;18:04&#160;UTC | true | [view](CERTS/b05bb158c054b6e579fe20750c83ba5524896545d629b96b51d4c8d9ab38dfd3/README.md) |
| 23&#160;Nov&#160;22&#160;17:31&#160;UTC | SHAKEN 4036 | 10&#160;May&#160;23&#160;17:31&#160;UTC | true | [view](CERTS/e03ff4259e6b456f1bb19d011f005932e83b0a35b520730c97cd77ad8868071c/README.md) |
| 30&#160;Nov&#160;22&#160;12:32&#160;UTC | SHAKEN 193E | 29&#160;Jan&#160;23&#160;12:32&#160;UTC | true | [view](CERTS/b1085c4ed6a32a87dc773d65e4e7aec591f7625ad80f63ee690d56b74b7a1430/README.md) |
| 03&#160;Dec&#160;22&#160;04:11&#160;UTC | SHAKEN 345J | 02&#160;Jan&#160;23&#160;04:11&#160;UTC | true | [view](CERTS/8991e6ec3e951d1b91f0e0f63b5bc9bdc929b5c76a73fea6c5af7a6401d0aa6e/README.md) |
| 03&#160;Dec&#160;22&#160;05:00&#160;UTC | SHAKEN 345J | 02&#160;Jan&#160;23&#160;05:00&#160;UTC | true | [view](CERTS/df273ff8f12fbf2fda88e527381fba4ac91927b0be97ec38a2d67993a5e0c074/README.md) |
| 05&#160;Dec&#160;22&#160;15:23&#160;UTC | SHAKEN 8526 | 04&#160;Jan&#160;23&#160;15:23&#160;UTC | true | [view](CERTS/29453875467b9cbda7eccf1ee68baaa883140c03b133004d5f21864cec03a029/README.md) |
| 09&#160;Dec&#160;22&#160;19:11&#160;UTC | SHAKEN 841J | 23&#160;Dec&#160;22&#160;19:11&#160;UTC | true | [view](CERTS/af11f73f26eda12928f6b5c61c8bcecc0467dd3924267f7fb0fb8d6bd9820cf9/README.md) |
| 10&#160;Dec&#160;22&#160;17:27&#160;UTC | SHAKEN 107K | 17&#160;Dec&#160;22&#160;17:27&#160;UTC | true | [view](CERTS/126451a6a680727d9468ca324c603bff3aa2703f87c585dff68277f4025b5b66/README.md) |
| 10&#160;Dec&#160;22&#160;18:40&#160;UTC | SHAKEN 089K | 17&#160;Dec&#160;22&#160;18:40&#160;UTC | true | [view](CERTS/354e72a4eb4a9cc96af2f1664c75a4a9cf7bac770a30a42a4d600cb630b844e3/README.md) |
| 10&#160;Dec&#160;22&#160;19:59&#160;UTC | SHAKEN 297K | 17&#160;Dec&#160;22&#160;19:59&#160;UTC | true | [view](CERTS/74e4079abfce7562746b817a8db78a4622b49d5ba09cbc4356b25fe6c916c81f/README.md) |
| 10&#160;Dec&#160;22&#160;20:05&#160;UTC | SHAKEN 366G | 17&#160;Dec&#160;22&#160;20:05&#160;UTC | true | [view](CERTS/52ac4bca14ab304a804678a214df11c7ddebad029ed1fe1fc6dc23f019dc2922/README.md) |
| 10&#160;Dec&#160;22&#160;20:20&#160;UTC | SHAKEN 674J | 17&#160;Dec&#160;22&#160;20:20&#160;UTC | true | [view](CERTS/5c5ff79c800c3b25ba803b8ecb8b7a0cb764400fd298c414c31b438852b26a95/README.md) |
| 10&#160;Dec&#160;22&#160;20:22&#160;UTC | SHAKEN 594J | 17&#160;Dec&#160;22&#160;20:21&#160;UTC | true | [view](CERTS/1184f51d0289faabb98ec8843a83aa26658fa76a2a9b8317e264f165453b2e17/README.md) |
| 10&#160;Dec&#160;22&#160;20:22&#160;UTC | SHAKEN 738J | 17&#160;Dec&#160;22&#160;20:22&#160;UTC | true | [view](CERTS/24c1aa497de706a085015aef6ad44b9aa5da9d655fa83296d73e0bb8187d39bf/README.md) |
| 10&#160;Dec&#160;22&#160;20:23&#160;UTC | SHAKEN 733J | 17&#160;Dec&#160;22&#160;20:23&#160;UTC | true | [view](CERTS/c37cfa08b5f3e888d9224d31c2f6288962501263e4e2482cc544af8ca6a3387d/README.md) |
| 10&#160;Dec&#160;22&#160;20:25&#160;UTC | SHAKEN 769J | 17&#160;Dec&#160;22&#160;20:25&#160;UTC | true | [view](CERTS/fac5bfdbb82f0328634f947accb811f90d97cdff7237f746b4c47b20f3ff1838/README.md) |
| 10&#160;Dec&#160;22&#160;20:25&#160;UTC | SHAKEN 916J | 17&#160;Dec&#160;22&#160;20:25&#160;UTC | true | [view](CERTS/2892f53850bad1b32c56dc27f04f2f8c2a68dbf21953f735b644e52edb06463c/README.md) |
| 10&#160;Dec&#160;22&#160;20:26&#160;UTC | SHAKEN 854D | 17&#160;Dec&#160;22&#160;20:26&#160;UTC | true | [view](CERTS/84d39a066d7b8cbb4c52bfa2cd8845ed95a83805cb8b8f9ecd2d7df1bba9aea4/README.md) |
| 10&#160;Dec&#160;22&#160;20:27&#160;UTC | SHAKEN 469A | 17&#160;Dec&#160;22&#160;20:27&#160;UTC | true | [view](CERTS/8f671a9cd4f526e0159d7630c926b7ca6cb8c012d8609293e46db0b1be7f217d/README.md) |
| 10&#160;Dec&#160;22&#160;20:27&#160;UTC | SHAKEN 726J | 17&#160;Dec&#160;22&#160;20:27&#160;UTC | true | [view](CERTS/4124d728e694d86020281ec3d2cbc38df0a1fa0104eb1049abbc5f77283d0cf8/README.md) |
| 10&#160;Dec&#160;22&#160;20:27&#160;UTC | SHAKEN 495J | 17&#160;Dec&#160;22&#160;20:27&#160;UTC | true | [view](CERTS/590f09a11aca23bb4b624995ab925819bf7bc58a951663401d142007cd7e592e/README.md) |
| 10&#160;Dec&#160;22&#160;20:27&#160;UTC | SHAKEN 790J | 17&#160;Dec&#160;22&#160;20:27&#160;UTC | true | [view](CERTS/7b1e4e99647893cb3d6774b5ff13fc15e30ed61a6c8af049f75b528778da6b5f/README.md) |
| 10&#160;Dec&#160;22&#160;20:28&#160;UTC | SHAKEN 738J | 17&#160;Dec&#160;22&#160;20:28&#160;UTC | true | [view](CERTS/6ff6ba3e2275be1b5e5fa4349a316ab80ca9d5b36c93b65114d1e8d47d0de1e2/README.md) |
| 10&#160;Dec&#160;22&#160;20:28&#160;UTC | SHAKEN 459J | 17&#160;Dec&#160;22&#160;20:28&#160;UTC | true | [view](CERTS/d3fcb07b828ae15c1c9d9ac00c44f2771cf72557635791dadc31cfe3d3f3ebb1/README.md) |
| 10&#160;Dec&#160;22&#160;20:28&#160;UTC | SHAKEN 672B | 17&#160;Dec&#160;22&#160;20:28&#160;UTC | true | [view](CERTS/298fb233d449b2056af6fe2b602f0edba9e287ebab5716be504a34b9040663f3/README.md) |
| 10&#160;Dec&#160;22&#160;20:29&#160;UTC | SHAKEN 366G | 17&#160;Dec&#160;22&#160;20:29&#160;UTC | true | [view](CERTS/1efd59ce1b57bccdacec422adb4d52d5d66184d8a8cb3086afceaaf92479b3f5/README.md) |
| 10&#160;Dec&#160;22&#160;20:29&#160;UTC | SHAKEN 0226 | 17&#160;Dec&#160;22&#160;20:29&#160;UTC | true | [view](CERTS/e162ab2e8de82f27b19772dfe4024cd59fa9884e61b1192602d6073040c8edb9/README.md) |
| 10&#160;Dec&#160;22&#160;20:29&#160;UTC | SHAKEN 738J | 17&#160;Dec&#160;22&#160;20:29&#160;UTC | true | [view](CERTS/d7c02e574eb1f64712ae7891b67cc3195cc51836619f6a7868f30b1e8cd95a92/README.md) |
| 10&#160;Dec&#160;22&#160;21:02&#160;UTC | SHAKEN 518J | 17&#160;Dec&#160;22&#160;21:02&#160;UTC | true | [view](CERTS/35f9e8af3f166a65df27c6eadaf5d65a7977958b3ee9176fabe10ab353863882/README.md) |
| 11&#160;Dec&#160;22&#160;10:13&#160;UTC | SHAKEN 625J | 18&#160;Dec&#160;22&#160;10:13&#160;UTC | true | [view](CERTS/5085bc09d10f83d317771c8c1aeff1c7ff2b8cacea7c1b4bc32c05c7c5116203/README.md) |
| 11&#160;Dec&#160;22&#160;12:16&#160;UTC | SHAKEN 130B | 18&#160;Dec&#160;22&#160;12:16&#160;UTC | true | [view](CERTS/a8695d7445ef815c754f532a3a98a4bb1e0d00de2a692944edafcd4dcc5caad1/README.md) |
| 11&#160;Dec&#160;22&#160;18:06&#160;UTC | SHAKEN 406K | 18&#160;Dec&#160;22&#160;18:06&#160;UTC | true | [view](CERTS/b50f0ec6c761edd451697552d595edcf1894f916d87d3b5d6ffe750d1cd0753f/README.md) |
| 11&#160;Dec&#160;22&#160;19:55&#160;UTC | SHAKEN 314K | 18&#160;Dec&#160;22&#160;19:55&#160;UTC | true | [view](CERTS/8a8dbbccc565907a1e5377b539d9492b0c8ebce47bba8a6616c1967657eeb0da/README.md) |
| 11&#160;Dec&#160;22&#160;20:18&#160;UTC | SHAKEN 983J | 18&#160;Dec&#160;22&#160;20:18&#160;UTC | true | [view](CERTS/3d8d1ab66f127d204d94dd17537d725a6b9d9924dfd2e2b6084d3cc40ea96cfa/README.md) |
| 11&#160;Dec&#160;22&#160;21:45&#160;UTC | SHAKEN 606F | 18&#160;Dec&#160;22&#160;21:45&#160;UTC | true | [view](CERTS/4807a347ac25d603fc1d4851813375602c4efde6ecef01676d09cda953b7cfb2/README.md) |
| 11&#160;Dec&#160;22&#160;22:28&#160;UTC | SHAKEN 120K | 18&#160;Dec&#160;22&#160;22:28&#160;UTC | true | [view](CERTS/608ee867cf61116ebe11c1c16f03629aa5283bf1ba2a7f860b07e2881cf4b255/README.md) |
| 12&#160;Dec&#160;22&#160;01:44&#160;UTC | SHAKEN 278K | 19&#160;Dec&#160;22&#160;01:44&#160;UTC | true | [view](CERTS/779318d9559d6d699e61eda9216d72cbbe47707983cf0ae5be1de534febf8dd4/README.md) |
| 12&#160;Dec&#160;22&#160;13:27&#160;UTC | SHAKEN 734J | 19&#160;Dec&#160;22&#160;13:27&#160;UTC | true | [view](CERTS/d2d56af39e29d187f109d03026f2408ac9e76d85e0ca1ff7d350b96de0966cf1/README.md) |
| 12&#160;Dec&#160;22&#160;14:37&#160;UTC | SHAKEN 2550 | 19&#160;Dec&#160;22&#160;14:37&#160;UTC | true | [view](CERTS/3ce840842c570db19c5fff8fd8730932505ea512edde018b141ced93a9bed8a1/README.md) |
| 12&#160;Dec&#160;22&#160;15:43&#160;UTC | SHAKEN 722J | 11&#160;Jan&#160;23&#160;15:43&#160;UTC | true | [view](CERTS/234bb11af3db6a37bc8a1807ba5458331e17031ec7772a9d790809a0c2856761/README.md) |
| 12&#160;Dec&#160;22&#160;18:12&#160;UTC | SHAKEN 952J | 11&#160;Jan&#160;23&#160;18:12&#160;UTC | true | [view](CERTS/060316d5062c63a01850cffc57dd7d976e37d02e17d9f77d165ba3a16c1172e0/README.md) |
| 12&#160;Dec&#160;22&#160;20:09&#160;UTC | SHAKEN 864J | 19&#160;Dec&#160;22&#160;20:09&#160;UTC | true | [view](CERTS/a8a6015b1d7fcb9ea592b56ecddc7a898109cfd958280206ce9867601eb2b888/README.md) |
| 12&#160;Dec&#160;22&#160;20:15&#160;UTC | SHAKEN 177K | 19&#160;Dec&#160;22&#160;20:15&#160;UTC | true | [view](CERTS/9f8feff65799d57da24c5e21a465294c735fc241babc414cd032ecafaba98dbd/README.md) |
| 12&#160;Dec&#160;22&#160;20:27&#160;UTC | SHAKEN 848J | 19&#160;Dec&#160;22&#160;20:27&#160;UTC | true | [view](CERTS/0fa6767e8ab704048bcfb0fd46ae496e2a5fd201e42dead0eb3462f8edb1c8c9/README.md) |
| 12&#160;Dec&#160;22&#160;20:28&#160;UTC | SHAKEN 738J | 21&#160;Dec&#160;22&#160;20:28&#160;UTC | true | [view](CERTS/759ad048688565b86768369621e35e96a6ad187ed59bbfa9067ff73f7d6596e4/README.md) |
| 13&#160;Dec&#160;22&#160;08:55&#160;UTC | SHAKEN 0172 | 20&#160;Dec&#160;22&#160;08:55&#160;UTC | true | [view](CERTS/c30e722e841c6eaf0d6f1544635d8a88d32dfee5c29ee89200530c99fa0c77a0/README.md) |
| 13&#160;Dec&#160;22&#160;14:46&#160;UTC | SHAKEN 186K | 20&#160;Dec&#160;22&#160;14:46&#160;UTC | true | [view](CERTS/123e7d93cd7c4873706653a44eb9439b034f4140faa567b49a1df101be96dc09/README.md) |
| 13&#160;Dec&#160;22&#160;16:31&#160;UTC | SHAKEN 763J | 20&#160;Dec&#160;22&#160;16:31&#160;UTC | true | [view](CERTS/90dd7fb67833d5239a622121d43d5f497cce1ef7517cdb16ae5547907ecd394e/README.md) |
| 13&#160;Dec&#160;22&#160;17:28&#160;UTC | SHAKEN 107K | 20&#160;Dec&#160;22&#160;17:28&#160;UTC | true | [view](CERTS/670010feecf6fd8f0830039f898d75758154f0a2a04745b7e6a342726f7d975c/README.md) |
| 13&#160;Dec&#160;22&#160;18:16&#160;UTC | SHAKEN 060K | 20&#160;Dec&#160;22&#160;18:16&#160;UTC | true | [view](CERTS/fbf0035fdd1b2358d442ee2165262dda0c831d6436697ae04092bfd1b354be0c/README.md) |
| 13&#160;Dec&#160;22&#160;18:40&#160;UTC | SHAKEN 089K | 20&#160;Dec&#160;22&#160;18:40&#160;UTC | true | [view](CERTS/3eef7eb4f60fe71b677f0ecd3c342e47f7ed277803d1b091e5ee808adddc7d14/README.md) |
| 13&#160;Dec&#160;22&#160;18:40&#160;UTC | SHAKEN 056K | 20&#160;Dec&#160;22&#160;18:40&#160;UTC | true | [view](CERTS/b5960c01004092d5e37b2464608a9ac5ca4ce99ada955070f4473a06b9d2a429/README.md) |
| 13&#160;Dec&#160;22&#160;18:43&#160;UTC | SHAKEN 9714 | 20&#160;Dec&#160;22&#160;18:42&#160;UTC | true | [view](CERTS/d41b6b296545811b2a10f7a61c4f3ace7226e6050fc6c8f5577b3f082d6bdaae/README.md) |
| 13&#160;Dec&#160;22&#160;19:59&#160;UTC | SHAKEN 297K | 20&#160;Dec&#160;22&#160;19:59&#160;UTC | true | [view](CERTS/ae49e63fc9720ae18d1ca1ba7ca3e8a2beaf2d5ecdf609323dca7d5023a997c3/README.md) |
| 13&#160;Dec&#160;22&#160;20:05&#160;UTC | SHAKEN 366G | 20&#160;Dec&#160;22&#160;20:05&#160;UTC | true | [view](CERTS/59f68da9caf90177e8fc406d0c75025fdc43611d58a6dc915f1cfd9c1899125f/README.md) |
| 13&#160;Dec&#160;22&#160;20:21&#160;UTC | SHAKEN 674J | 20&#160;Dec&#160;22&#160;20:21&#160;UTC | true | [view](CERTS/78166b554a99f8bd7115cee1cd721d234fd5eb67bcf058a57722fd8753e54fa7/README.md) |
| 13&#160;Dec&#160;22&#160;20:22&#160;UTC | SHAKEN 5512 | 20&#160;Dec&#160;22&#160;20:22&#160;UTC | true | [view](CERTS/cf5212472009dcb60fb7c49ca4f5f9c4af0ba8fb296cef02eba7b3eb96ee11d1/README.md) |
| 13&#160;Dec&#160;22&#160;20:22&#160;UTC | SHAKEN 738J | 20&#160;Dec&#160;22&#160;20:22&#160;UTC | true | [view](CERTS/8c4dc652f9c4b2425305adb530447494ee8b90248092762ebb76bf8cc6b62e21/README.md) |
| 13&#160;Dec&#160;22&#160;20:23&#160;UTC | SHAKEN 856H | 20&#160;Dec&#160;22&#160;20:23&#160;UTC | true | [view](CERTS/ad1a59fba9c4fd894a3935bddeba5fddbde08cc985e5ce7103e70b776995718b/README.md) |
| 13&#160;Dec&#160;22&#160;20:26&#160;UTC | SHAKEN 727J | 20&#160;Dec&#160;22&#160;20:26&#160;UTC | true | [view](CERTS/c9386d7db6903e1b52907146b7edab2e8dd3c4b65ac0167ae1620823b72893e3/README.md) |
| 13&#160;Dec&#160;22&#160;20:28&#160;UTC | SHAKEN 849J | 20&#160;Dec&#160;22&#160;20:28&#160;UTC | true | [view](CERTS/8a265f11cdb5c249a0f264ab8f9d8c9b10da7e6bb6e039253e9def2e258b1b58/README.md) |
| 13&#160;Dec&#160;22&#160;20:28&#160;UTC | SHAKEN 726J | 20&#160;Dec&#160;22&#160;20:28&#160;UTC | true | [view](CERTS/b3848108f9a1d8ca956a775f0c6daf526a6af507c849681fc9713b3bd5a3bd0f/README.md) |
| 13&#160;Dec&#160;22&#160;20:28&#160;UTC | SHAKEN 495J | 20&#160;Dec&#160;22&#160;20:28&#160;UTC | true | [view](CERTS/791563bc525a96b19c8e2161180600cc93d7b44e08002b905eb40a86bebef712/README.md) |
| 13&#160;Dec&#160;22&#160;20:29&#160;UTC | SHAKEN 625J | 20&#160;Dec&#160;22&#160;20:29&#160;UTC | true | [view](CERTS/38eb0e30ee45e28f21488370cf2fdd297501c5f661300d4a48b7581c0b402014/README.md) |
| 13&#160;Dec&#160;22&#160;20:29&#160;UTC | SHAKEN 738J | 20&#160;Dec&#160;22&#160;20:29&#160;UTC | true | [view](CERTS/f6220fc6833861a35637e8d356711bf7dc334ecab1b20f9549743a148070b1fa/README.md) |
| 13&#160;Dec&#160;22&#160;20:29&#160;UTC | SHAKEN 459J | 20&#160;Dec&#160;22&#160;20:29&#160;UTC | true | [view](CERTS/144e584f4e552a53f00bc4404ddc60b91431b6112668938c6c9e84dbc31682c8/README.md) |
| 13&#160;Dec&#160;22&#160;20:29&#160;UTC | SHAKEN 672B | 20&#160;Dec&#160;22&#160;20:29&#160;UTC | true | [view](CERTS/8135982e3516113f0652cbbcfd6f150a2c78deb3b0f2eeece4d4532942409532/README.md) |
| 13&#160;Dec&#160;22&#160;20:30&#160;UTC | SHAKEN 366G | 20&#160;Dec&#160;22&#160;20:30&#160;UTC | true | [view](CERTS/1eaacc73b22298f268f76159a7f10f25978a10a4097791a8dd88a7f8981cb990/README.md) |
| 13&#160;Dec&#160;22&#160;20:30&#160;UTC | SHAKEN 0226 | 20&#160;Dec&#160;22&#160;20:30&#160;UTC | true | [view](CERTS/836f8587c74c13f0ffb5aa11bd44ae92011c4982ca8a3688560ae80d6768bace/README.md) |
| 13&#160;Dec&#160;22&#160;20:30&#160;UTC | SHAKEN 738J | 20&#160;Dec&#160;22&#160;20:30&#160;UTC | true | [view](CERTS/2ccf1992cacc575a933e4e218620eca930cc59eda7352b5393943628ea3dd023/README.md) |
| 13&#160;Dec&#160;22&#160;21:02&#160;UTC | SHAKEN 518J | 20&#160;Dec&#160;22&#160;21:02&#160;UTC | true | [view](CERTS/5544a31744d05629239d8c1b08ca2c5dcf4e7890d8001e7dc51d8a18313e773e/README.md) |
| 14&#160;Dec&#160;22&#160;12:16&#160;UTC | SHAKEN 130B | 21&#160;Dec&#160;22&#160;12:16&#160;UTC | true | [view](CERTS/75d9eee369772c6b8610cfe6e8e0c9ce84537f8bc8850f5f5a3a94ea3c08a042/README.md) |
| 14&#160;Dec&#160;22&#160;15:13&#160;UTC | SHAKEN 6922 | 21&#160;Dec&#160;22&#160;15:13&#160;UTC | true | [view](CERTS/4542751c679ff03352c0109b15ce582b48b7e6001a6d33bb775dfd14f44a5e1b/README.md) |
| 14&#160;Dec&#160;22&#160;15:47&#160;UTC | SHAKEN 793J | 21&#160;Dec&#160;22&#160;15:47&#160;UTC | true | [view](CERTS/05468f8b9ae66fdddadf653600bcc5efa8015053ba25e4244c1de72313a5a0a7/README.md) |
| 14&#160;Dec&#160;22&#160;16:43&#160;UTC | SHAKEN 967J | 21&#160;Dec&#160;22&#160;16:43&#160;UTC | true | [view](CERTS/70433a5c2a329cbe01b8074a6ee5e7ff86bfd74a85089f8a829747f4aa9462f6/README.md) |
| 14&#160;Dec&#160;22&#160;19:11&#160;UTC | SHAKEN 749J | 21&#160;Dec&#160;22&#160;19:11&#160;UTC | true | [view](CERTS/727a3454008f768e859dd10e19aa89acd81ba4dabb537c6991f939cd635b4908/README.md) |
| 14&#160;Dec&#160;22&#160;20:19&#160;UTC | SHAKEN 983J | 21&#160;Dec&#160;22&#160;20:19&#160;UTC | true | [view](CERTS/d71dbc63a97b93a13f4f2fb45162fcf6bf0378a0cd3e66f5e56667aca8a05593/README.md) |
| 14&#160;Dec&#160;22&#160;22:42&#160;UTC | SHAKEN 551G | 21&#160;Dec&#160;22&#160;22:42&#160;UTC | true | [view](CERTS/97618d6cdc78e1adfeb765eb4a6319d1b5afe07aef4c64a7b8416747df4a444c/README.md) |
| 15&#160;Dec&#160;22&#160;01:44&#160;UTC | SHAKEN 278K | 22&#160;Dec&#160;22&#160;01:44&#160;UTC | true | [view](CERTS/6fea96c7d133a4e45902257382bd4b1ea323c3b5ef50eca3c1988a7e0886c45f/README.md) |
| 15&#160;Dec&#160;22&#160;14:37&#160;UTC | SHAKEN 2550 | 22&#160;Dec&#160;22&#160;14:37&#160;UTC | true | [view](CERTS/5d50f71a74935e71636aad857bd486cb6335c9a90640fa8fb2c355e886d244c8/README.md) |
| 15&#160;Dec&#160;22&#160;16:00&#160;UTC | SHAKEN 311H | 22&#160;Dec&#160;22&#160;16:00&#160;UTC | true | [view](CERTS/9e30ca392121a83c88a1d85e23ee6a0ed676277a17e38022a5650b690716a788/README.md) |
| 15&#160;Dec&#160;22&#160;18:49&#160;UTC | SHAKEN 140K | 22&#160;Dec&#160;22&#160;18:49&#160;UTC | true | [view](CERTS/630563fd9ea909aa41453ad9c0bd61d118a1ae6b047adafa2976988c2452d8e8/README.md) |
| 15&#160;Dec&#160;22&#160;20:09&#160;UTC | SHAKEN 864J | 22&#160;Dec&#160;22&#160;20:09&#160;UTC | true | [view](CERTS/6cd3a2889fce1acd4b66c7156f058f94761e22e0f1f1896be8660987f76218ee/README.md) |
| 15&#160;Dec&#160;22&#160;20:15&#160;UTC | SHAKEN 177K | 22&#160;Dec&#160;22&#160;20:15&#160;UTC | true | [view](CERTS/745e89c28295d8d6780dbfebf9ff9eaff838973b678eda5c4e0252596bb74c71/README.md) |
| 16&#160;Dec&#160;22&#160;14:46&#160;UTC | SHAKEN 186K | 23&#160;Dec&#160;22&#160;14:46&#160;UTC | true | [view](CERTS/4b425db0358f880c8eb8fbe651715acf744c1551549aabe2996f9009bfd5178d/README.md) |
| 16&#160;Dec&#160;22&#160;17:04&#160;UTC | SHAKEN 660C | 23&#160;Dec&#160;22&#160;17:04&#160;UTC | true | [view](CERTS/40c55e201869292c3f64c46a9f396c8cab8e1d5b2cf6a73f430971f577d78a65/README.md) |

#### CA Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 20&#160;Aug&#160;21&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Issuing CA3 | 19&#160;Aug&#160;31&#160;23:59&#160;UTC | true | [view](CERTS/51b51431ffd8055cd7e30c213540f4868258058d4c4c9f23f1ed8bea75101e37/README.md) |
| 20&#160;Aug&#160;21&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Issuing CA1 | 19&#160;Aug&#160;31&#160;23:59&#160;UTC | true | [view](CERTS/037a7c6621814bd5f01fa2080a77a245458694d9d566e0050e068fc853317857/README.md) |
| 20&#160;Aug&#160;21&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Root CA1 | 19&#160;Aug&#160;41&#160;23:59&#160;UTC | true | [view](CERTS/59ab943f4e8c3c17755cc8e4abb5ac65736ce74ec587c62cb744ab0babbbe2fe/README.md) |
| 24&#160;Oct&#160;22&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Issuing CA4 | 23&#160;Oct&#160;32&#160;23:59&#160;UTC | false | [view](CERTS/80dd12d935f46c256c7c1289a0834b21ec8f2f8c35f2864928d1e75f3a280665/README.md) |
| 24&#160;Oct&#160;22&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Root CA2 | 23&#160;Oct&#160;42&#160;23:59&#160;UTC | false | [view](CERTS/a26e04fc786ab70b8085236b2c53f8cfbf5d0c6a5c2c9c3e9f91669fbb8ea4d5/README.md) |


Generated: 17 Dec 22 17:07 UTC