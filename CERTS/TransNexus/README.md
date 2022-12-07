# STIR/SHAKEN CA Ecosystem Compliance

## TransNexus

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 670 certificates were included in the corpus being tested
- 1 certificates in the corpus were skipped because they are duplicates
- 578 certificates in the corpus were skipped because they are expired
- 1 certificates in the corpus were skipped because they are not currently trusted
- 90 certificates being tested against the remaining rules
- 1.47 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 22.22% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 1.11% of certificates are too old to be assessed against currently enforced expectations
- 60 days is the average remaining validity for the certificates in the corpus
- 59 days is the average initial validity for the certificates in the corpus
- 75 certificates expire in the next 30 days
- 1.43 average number of unexpired certificates per OCN observed
- 63 unique OCNs observed in unexpired and valid certificate corpus

| Instances | Test | Source |
|-----------|------|--------|
| 3 | [e_atis_certificate_policies](ISSUES/e_atis_certificate_policies/README.md) | ATIS1000080 |
| 16 | [e_incorrect_ku_encoding](ISSUES/e_incorrect_ku_encoding/README.md) | RFC5280 |
| 3 | [e_us_cp_ambiguous_identifier](ISSUES/e_us_cp_ambiguous_identifier/README.md) | US_SHAKEN_CP |
| 90 | [e_us_cp_subject_sn](ISSUES/e_us_cp_subject_sn/README.md) | US_SHAKEN_CP |
| 20 | [w_shaken_subject_rdn_unknown](ISSUES/w_shaken_subject_rdn_unknown/README.md) | SHAKEN_PKI_BEST_PRACTICES |

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
- 5471 days is the average remaining validity for the certificates in the corpus
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
| 12&#160;Nov&#160;22&#160;04:05&#160;UTC | SHAKEN 345J | 12&#160;Dec&#160;22&#160;04:05&#160;UTC | true | [view](CERTS/ac38b9036e27ff2b753a0727e8ccdfd800d459dfdf595454c6a9dc7e9cd8ad74/README.md) |
| 12&#160;Nov&#160;22&#160;04:54&#160;UTC | SHAKEN 345J | 12&#160;Dec&#160;22&#160;04:54&#160;UTC | true | [view](CERTS/0a4891c263ce12ebb5116b43048b2b427314bd6133848201f8186b62f7efc68a/README.md) |
| 18&#160;Nov&#160;22&#160;15:42&#160;UTC | SHAKEN 722J | 18&#160;Dec&#160;22&#160;15:42&#160;UTC | true | [view](CERTS/4111bb56ba547a2f493d1e9c2baefa33ce9bd117904133477d063d6fa0d6b9b2/README.md) |
| 18&#160;Nov&#160;22&#160;18:04&#160;UTC | SHAKEN 952J | 18&#160;Dec&#160;22&#160;18:04&#160;UTC | true | [view](CERTS/b05bb158c054b6e579fe20750c83ba5524896545d629b96b51d4c8d9ab38dfd3/README.md) |
| 23&#160;Nov&#160;22&#160;17:31&#160;UTC | SHAKEN 4036 | 10&#160;May&#160;23&#160;17:31&#160;UTC | true | [view](CERTS/e03ff4259e6b456f1bb19d011f005932e83b0a35b520730c97cd77ad8868071c/README.md) |
| 28&#160;Nov&#160;22&#160;20:26&#160;UTC | SHAKEN 738J | 07&#160;Dec&#160;22&#160;20:26&#160;UTC | true | [view](CERTS/e40280dfe9a62708996198fc786e18c7578f37d284ff1e2c9d2cf24e94b878ee/README.md) |
| 30&#160;Nov&#160;22&#160;12:32&#160;UTC | SHAKEN 193E | 29&#160;Jan&#160;23&#160;12:32&#160;UTC | true | [view](CERTS/b1085c4ed6a32a87dc773d65e4e7aec591f7625ad80f63ee690d56b74b7a1430/README.md) |
| 30&#160;Nov&#160;22&#160;20:08&#160;UTC | SHAKEN 864J | 07&#160;Dec&#160;22&#160;20:08&#160;UTC | true | [view](CERTS/22f97699716da9cfef879f54b79bbb387bd18c9e69c8d76ed610937e99feab15/README.md) |
| 30&#160;Nov&#160;22&#160;20:14&#160;UTC | SHAKEN 177K | 07&#160;Dec&#160;22&#160;20:14&#160;UTC | true | [view](CERTS/958094f268a2789dcd29fcd45b06253bea4f97e519ef5201ec05e321c0d9245e/README.md) |
| 01&#160;Dec&#160;22&#160;09:33&#160;UTC | SHAKEN 841J | 15&#160;Dec&#160;22&#160;09:33&#160;UTC | true | [view](CERTS/f98bb6bf4f32b42eda9a2113135539ab9422a7cb2d958e3e8c3fce20e1cbbd6d/README.md) |
| 01&#160;Dec&#160;22&#160;17:27&#160;UTC | SHAKEN 107K | 08&#160;Dec&#160;22&#160;17:27&#160;UTC | true | [view](CERTS/56a106b7e8859d4e44b9640a18df0673305e5276dc34afab402e951d54195461/README.md) |
| 01&#160;Dec&#160;22&#160;18:40&#160;UTC | SHAKEN 089K | 08&#160;Dec&#160;22&#160;18:40&#160;UTC | true | [view](CERTS/fe116bea6ea69c5c432f0529162bd943820a04753c8012877caa4ce123bd283a/README.md) |
| 01&#160;Dec&#160;22&#160;20:03&#160;UTC | SHAKEN 366G | 08&#160;Dec&#160;22&#160;20:03&#160;UTC | true | [view](CERTS/ab5d4a629359cf2e4f75bc3260762a365999ae2b521a9e432b0a6278c64d40b7/README.md) |
| 01&#160;Dec&#160;22&#160;20:19&#160;UTC | SHAKEN 674J | 08&#160;Dec&#160;22&#160;20:19&#160;UTC | true | [view](CERTS/9d5da9e74157eb8574e772e70953f0a00b35d6fd6aab8f5431cea6dceef55424/README.md) |
| 01&#160;Dec&#160;22&#160;20:21&#160;UTC | SHAKEN 7453 | 08&#160;Dec&#160;22&#160;20:21&#160;UTC | true | [view](CERTS/08c0c54ef185820df850ec1f89f1484da8dcdbfa24528c95faf8296c2ddc0e9f/README.md) |
| 01&#160;Dec&#160;22&#160;20:22&#160;UTC | SHAKEN 819H | 08&#160;Dec&#160;22&#160;20:22&#160;UTC | true | [view](CERTS/8a69aaff13a3f0e03ff01782a15f8c53cedc53ce6d30f8b87b8c914756260997/README.md) |
| 01&#160;Dec&#160;22&#160;20:25&#160;UTC | SHAKEN 469A | 08&#160;Dec&#160;22&#160;20:25&#160;UTC | true | [view](CERTS/6ac7f6c67e6327b2721afde8d6a21fce5ee13c9e015efdb8f23dae4478ae1a98/README.md) |
| 01&#160;Dec&#160;22&#160;20:25&#160;UTC | SHAKEN 849J | 08&#160;Dec&#160;22&#160;20:25&#160;UTC | true | [view](CERTS/afb08f175aabd7e1a061fee924868d39ec509c11b5a15612aa35c1c29d2dab5e/README.md) |
| 01&#160;Dec&#160;22&#160;20:25&#160;UTC | SHAKEN 790J | 08&#160;Dec&#160;22&#160;20:25&#160;UTC | true | [view](CERTS/7e1cf83fd6fcfb0ba60baccc12adb85f042277c3a8cd472c5eda749ab2b52390/README.md) |
| 01&#160;Dec&#160;22&#160;20:26&#160;UTC | SHAKEN 625J | 08&#160;Dec&#160;22&#160;20:26&#160;UTC | true | [view](CERTS/129b8b9fd3602d5b6376b1f85886c415dec6ba23eb1513f0153e5a2b55bd22a5/README.md) |
| 01&#160;Dec&#160;22&#160;20:26&#160;UTC | SHAKEN 738J | 08&#160;Dec&#160;22&#160;20:26&#160;UTC | true | [view](CERTS/bfe0ab52f2e9ea4e8b2cb775fae53b712c2c6fc1616273b2cc2489c62de477c7/README.md) |
| 01&#160;Dec&#160;22&#160;20:26&#160;UTC | SHAKEN 459J | 08&#160;Dec&#160;22&#160;20:26&#160;UTC | true | [view](CERTS/a894f7b6708baa4dbba5fb77eed84001f4b929a8302af0d0f9abe55313be3084/README.md) |
| 01&#160;Dec&#160;22&#160;20:27&#160;UTC | SHAKEN 366G | 08&#160;Dec&#160;22&#160;20:27&#160;UTC | true | [view](CERTS/a79e959d7ad69c58715d2fa1bae760975f7be10e222505167a840262820c9594/README.md) |
| 01&#160;Dec&#160;22&#160;20:27&#160;UTC | SHAKEN 738J | 08&#160;Dec&#160;22&#160;20:27&#160;UTC | true | [view](CERTS/326f2c545af1f4a65858cf7979af5ec00b04099620a1ada98c19bac3e820d278/README.md) |
| 01&#160;Dec&#160;22&#160;21:02&#160;UTC | SHAKEN 518J | 08&#160;Dec&#160;22&#160;21:02&#160;UTC | true | [view](CERTS/8cc69257b74e4aec9b9983643f816724d801409cc0fe9326c78805aa3bf30d5e/README.md) |
| 02&#160;Dec&#160;22&#160;20:17&#160;UTC | SHAKEN 983J | 09&#160;Dec&#160;22&#160;20:17&#160;UTC | true | [view](CERTS/6b970a333a2c1c53111994e0ad38d7fe60ca3649825742ec30d44f5917428644/README.md) |
| 02&#160;Dec&#160;22&#160;22:27&#160;UTC | SHAKEN 120K | 09&#160;Dec&#160;22&#160;22:27&#160;UTC | true | [view](CERTS/aee530f565babf584dbc38d8a979f2ff79b3c0f86032a7818dc6ba2a94efada0/README.md) |
| 03&#160;Dec&#160;22&#160;01:43&#160;UTC | SHAKEN 278K | 10&#160;Dec&#160;22&#160;01:43&#160;UTC | true | [view](CERTS/5ebcb57fa7b0c362da12c57eeeb0f8ebcafacb2dad3abf46f22ed3c2a943a2b4/README.md) |
| 03&#160;Dec&#160;22&#160;04:11&#160;UTC | SHAKEN 345J | 02&#160;Jan&#160;23&#160;04:11&#160;UTC | true | [view](CERTS/8991e6ec3e951d1b91f0e0f63b5bc9bdc929b5c76a73fea6c5af7a6401d0aa6e/README.md) |
| 03&#160;Dec&#160;22&#160;05:00&#160;UTC | SHAKEN 345J | 02&#160;Jan&#160;23&#160;05:00&#160;UTC | true | [view](CERTS/df273ff8f12fbf2fda88e527381fba4ac91927b0be97ec38a2d67993a5e0c074/README.md) |
| 03&#160;Dec&#160;22&#160;14:37&#160;UTC | SHAKEN 2550 | 10&#160;Dec&#160;22&#160;14:36&#160;UTC | true | [view](CERTS/8ebdd742f2eefa83792cbda3dcdbfc8e1ebb2b1fa23dfd9f9cabd6723a888310/README.md) |
| 03&#160;Dec&#160;22&#160;20:08&#160;UTC | SHAKEN 864J | 10&#160;Dec&#160;22&#160;20:08&#160;UTC | true | [view](CERTS/ec443f729ccaeed1896afe885a42f768773a34c458cddd1a91f80c4e45f5cd06/README.md) |
| 03&#160;Dec&#160;22&#160;20:15&#160;UTC | SHAKEN 177K | 10&#160;Dec&#160;22&#160;20:15&#160;UTC | true | [view](CERTS/be2abbeee87f1685b35a306af22d0260141173a60383d8ef3abd3aa4ab9287f7/README.md) |
| 04&#160;Dec&#160;22&#160;14:44&#160;UTC | SHAKEN 012K | 11&#160;Dec&#160;22&#160;14:44&#160;UTC | true | [view](CERTS/3dba8e687f5a8cdfa59df2d54c88c553bdb9b6fd2716c252765a92a7f48d080c/README.md) |
| 04&#160;Dec&#160;22&#160;16:31&#160;UTC | SHAKEN 763J | 11&#160;Dec&#160;22&#160;16:31&#160;UTC | true | [view](CERTS/329942657d1ad3e084dc8869ea0e773c44b9f16d174bb6c828497be2dbac1f1a/README.md) |
| 04&#160;Dec&#160;22&#160;17:27&#160;UTC | SHAKEN 107K | 11&#160;Dec&#160;22&#160;17:27&#160;UTC | true | [view](CERTS/68a4579d69df1d1b3a747afc600c58d23b22c40e35e0000c366ab0710e8010f5/README.md) |
| 04&#160;Dec&#160;22&#160;18:40&#160;UTC | SHAKEN 056K | 11&#160;Dec&#160;22&#160;18:40&#160;UTC | true | [view](CERTS/44c99c4155fecc2d573fbee37a6477bd99ee153241abb02a8c05877b103d31aa/README.md) |
| 04&#160;Dec&#160;22&#160;18:42&#160;UTC | SHAKEN 9714 | 11&#160;Dec&#160;22&#160;18:42&#160;UTC | true | [view](CERTS/24763dc02e5a9d039bc1f84dea63ac3538b1a40016d472389bdd4d055ec27041/README.md) |
| 04&#160;Dec&#160;22&#160;20:20&#160;UTC | SHAKEN 3013 | 11&#160;Dec&#160;22&#160;20:20&#160;UTC | true | [view](CERTS/307e4d516b546a4f64288de71398b176249a78eff8b25aba7e97a779924a60c4/README.md) |
| 04&#160;Dec&#160;22&#160;20:20&#160;UTC | SHAKEN 674J | 11&#160;Dec&#160;22&#160;20:20&#160;UTC | true | [view](CERTS/740c1659fe8d69cb5e8485b5f4429cec38bb08ee2c2104803b2d8c9b0f976cd1/README.md) |
| 04&#160;Dec&#160;22&#160;20:21&#160;UTC | SHAKEN 5512 | 11&#160;Dec&#160;22&#160;20:21&#160;UTC | true | [view](CERTS/0fe38e3e5cdde743817287a50f783b2b34f6243f9d3ac50f05885508bf9112e6/README.md) |
| 04&#160;Dec&#160;22&#160;20:21&#160;UTC | SHAKEN 065J | 11&#160;Dec&#160;22&#160;20:21&#160;UTC | true | [view](CERTS/aee58a94de35949189bb119a085e844d5c93f20284edd28d66e214a4a1d8c63e/README.md) |
| 04&#160;Dec&#160;22&#160;20:21&#160;UTC | SHAKEN 1817 | 11&#160;Dec&#160;22&#160;20:21&#160;UTC | true | [view](CERTS/0def95da35e07a018ab3dfca3406b13dd264c3243b1dc305fde639f983a9a821/README.md) |
| 04&#160;Dec&#160;22&#160;20:21&#160;UTC | SHAKEN 738J | 11&#160;Dec&#160;22&#160;20:21&#160;UTC | true | [view](CERTS/e259841ae56a9c5e1e5ef7c1cd59b72731217a5e82a1d5248efaa389458813e1/README.md) |
| 04&#160;Dec&#160;22&#160;20:22&#160;UTC | SHAKEN 700H | 11&#160;Dec&#160;22&#160;20:22&#160;UTC | true | [view](CERTS/1cf256b062a087d1f52e5bf8bf56610f9b7a6925eba1fb7bad86e99d6aebc941/README.md) |
| 04&#160;Dec&#160;22&#160;20:22&#160;UTC | SHAKEN 733J | 11&#160;Dec&#160;22&#160;20:22&#160;UTC | true | [view](CERTS/717a55472dbbbe7cc1c9917beb956585f53f1222e64c2e5716ad1552f58d8cd0/README.md) |
| 04&#160;Dec&#160;22&#160;20:22&#160;UTC | SHAKEN 7453 | 11&#160;Dec&#160;22&#160;20:22&#160;UTC | true | [view](CERTS/5e05cd2e715f89261fce067a40c279dbb85737f64f3128e86b919c496736f309/README.md) |
| 04&#160;Dec&#160;22&#160;20:23&#160;UTC | SHAKEN 819H | 11&#160;Dec&#160;22&#160;20:23&#160;UTC | true | [view](CERTS/354d47845a07effa8896b14aeac3bbd5235c8f6a6d659fa4d98a1f4cc8e1c115/README.md) |
| 04&#160;Dec&#160;22&#160;20:24&#160;UTC | SHAKEN 769J | 11&#160;Dec&#160;22&#160;20:24&#160;UTC | true | [view](CERTS/9a1ee4abca8f84cbcc47f7ddd0c18f07f689b6f3e137d885176500196ad41787/README.md) |
| 04&#160;Dec&#160;22&#160;20:24&#160;UTC | SHAKEN 766J | 11&#160;Dec&#160;22&#160;20:24&#160;UTC | true | [view](CERTS/301033ad0bf8e19b7460620abf96f6288d31019cac00cc37cdde394f2f4a22e1/README.md) |
| 04&#160;Dec&#160;22&#160;20:26&#160;UTC | SHAKEN 469A | 11&#160;Dec&#160;22&#160;20:25&#160;UTC | true | [view](CERTS/6ba784d221622857dfa86e4d045af391e3684d5d8ed44af0c9d5ff54a0407a1f/README.md) |
| 04&#160;Dec&#160;22&#160;20:26&#160;UTC | SHAKEN 495J | 11&#160;Dec&#160;22&#160;20:26&#160;UTC | true | [view](CERTS/1734566b90d33b6bc489b66c1809216380b1fff0529b88db6eb885e4c4ffb9f9/README.md) |
| 04&#160;Dec&#160;22&#160;20:26&#160;UTC | SHAKEN 790J | 11&#160;Dec&#160;22&#160;20:26&#160;UTC | true | [view](CERTS/dfa8c68232bea46d52e123312bf5702eeafc56ace4d612be7787711f9755dde1/README.md) |
| 04&#160;Dec&#160;22&#160;20:27&#160;UTC | SHAKEN 738J | 11&#160;Dec&#160;22&#160;20:27&#160;UTC | true | [view](CERTS/7e2e2e2e62cae0e6c1372cbb6309f5e3cc9e87e25557b94558b20a11a0dec2a9/README.md) |
| 04&#160;Dec&#160;22&#160;20:27&#160;UTC | SHAKEN 459J | 11&#160;Dec&#160;22&#160;20:27&#160;UTC | true | [view](CERTS/8bc07fb325d33ac7103ed93c34aa9b2874ffadf382fd607421304f7756d03ee2/README.md) |
| 04&#160;Dec&#160;22&#160;20:27&#160;UTC | SHAKEN 672B | 11&#160;Dec&#160;22&#160;20:27&#160;UTC | true | [view](CERTS/a72a71039948e41f9a111441fa172669a129d31ddb9862eeedc0f3e5561bbbed/README.md) |
| 04&#160;Dec&#160;22&#160;20:28&#160;UTC | SHAKEN 366G | 11&#160;Dec&#160;22&#160;20:28&#160;UTC | true | [view](CERTS/59f7e89beaa47edce949170d2851448c3e950fe9da862de5b8857744d72c3366/README.md) |
| 04&#160;Dec&#160;22&#160;20:28&#160;UTC | SHAKEN 0226 | 11&#160;Dec&#160;22&#160;20:28&#160;UTC | true | [view](CERTS/b7cb2cf1df5aa4154669204b1da979b5437ff66364e0d374bc217c207823d511/README.md) |
| 04&#160;Dec&#160;22&#160;20:28&#160;UTC | SHAKEN 738J | 11&#160;Dec&#160;22&#160;20:28&#160;UTC | true | [view](CERTS/7557652b364e8d17793cd65d13ef3792ab776d6972bc675cbb9f37d8695d185e/README.md) |
| 04&#160;Dec&#160;22&#160;21:02&#160;UTC | SHAKEN 518J | 11&#160;Dec&#160;22&#160;21:02&#160;UTC | true | [view](CERTS/82d808f1bdf998ef8ada811d024bc98a70a01d47e472f65732364147f6cf5f1b/README.md) |
| 04&#160;Dec&#160;22&#160;23:05&#160;UTC | SHAKEN 0172 | 11&#160;Dec&#160;22&#160;23:05&#160;UTC | true | [view](CERTS/bec6fd78f5e7b9177763996fd66b4a1ad2ead15756a41647568fb4166865b3cb/README.md) |
| 05&#160;Dec&#160;22&#160;03:21&#160;UTC | SHAKEN 551G | 12&#160;Dec&#160;22&#160;03:21&#160;UTC | true | [view](CERTS/ca276aec245b75c6e5f5af9d996ef84e92d2efd912db1d3dd33c2d3560f70806/README.md) |
| 05&#160;Dec&#160;22&#160;12:15&#160;UTC | SHAKEN 130B | 12&#160;Dec&#160;22&#160;12:15&#160;UTC | true | [view](CERTS/1902f2ae5a9965a868c477b55213f70aa110c049d9896229bfb55897da4e2bf0/README.md) |
| 05&#160;Dec&#160;22&#160;15:23&#160;UTC | SHAKEN 8526 | 04&#160;Jan&#160;23&#160;15:23&#160;UTC | true | [view](CERTS/29453875467b9cbda7eccf1ee68baaa883140c03b133004d5f21864cec03a029/README.md) |
| 05&#160;Dec&#160;22&#160;16:36&#160;UTC | SHAKEN 291K | 12&#160;Dec&#160;22&#160;16:36&#160;UTC | true | [view](CERTS/781f5211641bef2dcdcc3e5c7c5f523064a5217e33035848b78b97861820917f/README.md) |
| 05&#160;Dec&#160;22&#160;18:06&#160;UTC | SHAKEN 406K | 12&#160;Dec&#160;22&#160;18:06&#160;UTC | true | [view](CERTS/d66e7f3827854150919e2c7ae95f52ce42ad5526b94e169530c24ea99a8a5aaf/README.md) |
| 05&#160;Dec&#160;22&#160;20:18&#160;UTC | SHAKEN 983J | 12&#160;Dec&#160;22&#160;20:18&#160;UTC | true | [view](CERTS/86c4a780e9432e5f36653795761aee9504ae623e2568ad5c8dfbb8c4c6ba5604/README.md) |
| 05&#160;Dec&#160;22&#160;20:27&#160;UTC | SHAKEN 738J | 14&#160;Dec&#160;22&#160;20:27&#160;UTC | true | [view](CERTS/652a1a6e9b711bc5eba31cca81b066a5a8de44758bcc4e4b860efd27d1b42b15/README.md) |
| 06&#160;Dec&#160;22&#160;01:43&#160;UTC | SHAKEN 278K | 13&#160;Dec&#160;22&#160;01:43&#160;UTC | true | [view](CERTS/c5ff7f48274a7a8f03851226d7dd566ae43e54c22f5884794e36d894b8f6cc59/README.md) |
| 06&#160;Dec&#160;22&#160;14:37&#160;UTC | SHAKEN 2550 | 13&#160;Dec&#160;22&#160;14:37&#160;UTC | true | [view](CERTS/a268d78ca7a2ba41cf5deb6dac12c5621f6f73c3f68416e4366be9fe5026e10a/README.md) |
| 06&#160;Dec&#160;22&#160;18:48&#160;UTC | SHAKEN 140K | 13&#160;Dec&#160;22&#160;18:48&#160;UTC | true | [view](CERTS/0fd2fa51646825281f95779e5032774a2db27d3205fc63de78f1a9563af94ee2/README.md) |
| 06&#160;Dec&#160;22&#160;20:09&#160;UTC | SHAKEN 864J | 13&#160;Dec&#160;22&#160;20:09&#160;UTC | true | [view](CERTS/09baaed6752794dbeb4f94bf75d92f6b9ba1cb8ba1e2a9def93c0a8c586a6a61/README.md) |

#### CA Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 20&#160;Aug&#160;21&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Issuing CA3 | 19&#160;Aug&#160;31&#160;23:59&#160;UTC | true | [view](CERTS/51b51431ffd8055cd7e30c213540f4868258058d4c4c9f23f1ed8bea75101e37/README.md) |
| 20&#160;Aug&#160;21&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Issuing CA1 | 19&#160;Aug&#160;31&#160;23:59&#160;UTC | true | [view](CERTS/037a7c6621814bd5f01fa2080a77a245458694d9d566e0050e068fc853317857/README.md) |
| 20&#160;Aug&#160;21&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Root CA1 | 19&#160;Aug&#160;41&#160;23:59&#160;UTC | true | [view](CERTS/59ab943f4e8c3c17755cc8e4abb5ac65736ce74ec587c62cb744ab0babbbe2fe/README.md) |
| 24&#160;Oct&#160;22&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Issuing CA4 | 23&#160;Oct&#160;32&#160;23:59&#160;UTC | false | [view](CERTS/80dd12d935f46c256c7c1289a0834b21ec8f2f8c35f2864928d1e75f3a280665/README.md) |
| 24&#160;Oct&#160;22&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Root CA2 | 23&#160;Oct&#160;42&#160;23:59&#160;UTC | false | [view](CERTS/a26e04fc786ab70b8085236b2c53f8cfbf5d0c6a5c2c9c3e9f91669fbb8ea4d5/README.md) |


Generated: 07 Dec 22 18:54 UTC