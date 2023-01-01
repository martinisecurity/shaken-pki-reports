# STIR/SHAKEN CA Ecosystem Compliance

## TransNexus

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 926 certificates were included in the corpus being tested
- 1 certificates in the corpus were skipped because they are duplicates
- 877 certificates in the corpus were skipped because they are expired
- 1 certificates in the corpus were skipped because they are not currently trusted
- 47 certificates being tested against the remaining rules
- 1.72 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 29.79% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 2.13% of certificates are too old to be assessed against currently enforced expectations
- 102 days is the average remaining validity for the certificates in the corpus
- 103 days is the average initial validity for the certificates in the corpus
- 35 certificates expire in the next 30 days
- 1.18 average number of unexpired certificates per OCN observed
- 40 unique OCNs observed in unexpired and valid certificate corpus

| Instances | Test | Source |
|-----------|------|--------|
| 3 | [e_atis_certificate_policies](ISSUES/e_atis_certificate_policies/README.md) | ATIS1000080 |
| 14 | [e_incorrect_ku_encoding](ISSUES/e_incorrect_ku_encoding/README.md) | RFC5280 |
| 3 | [e_us_cp_ambiguous_identifier](ISSUES/e_us_cp_ambiguous_identifier/README.md) | US_SHAKEN_CP |
| 47 | [e_us_cp_subject_sn](ISSUES/e_us_cp_subject_sn/README.md) | US_SHAKEN_CP |
| 14 | [w_shaken_subject_rdn_unknown](ISSUES/w_shaken_subject_rdn_unknown/README.md) | SHAKEN_PKI_BEST_PRACTICES |

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
- 5467 days is the average remaining validity for the certificates in the corpus
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
| 12&#160;Oct&#160;22&#160;19:39&#160;UTC | SHAKEN 815G | 10&#160;Jan&#160;23&#160;19:39&#160;UTC | true | [view](CERTS/a79f50d4fcd3e680518379f12a86a4b9a21c5e27dc5422a8439525edf350c4a7/README.md) |
| 20&#160;Oct&#160;22&#160;15:48&#160;UTC | SHAKEN 622J | 18&#160;Jan&#160;23&#160;15:48&#160;UTC | true | [view](CERTS/2e34c765f110b197f2ec1052b1960bc25ce320dfce7d7035cc532cd1eb6dbd57/README.md) |
| 23&#160;Nov&#160;22&#160;17:31&#160;UTC | SHAKEN 4036 | 10&#160;May&#160;23&#160;17:31&#160;UTC | true | [view](CERTS/e03ff4259e6b456f1bb19d011f005932e83b0a35b520730c97cd77ad8868071c/README.md) |
| 30&#160;Nov&#160;22&#160;12:32&#160;UTC | SHAKEN 193E | 29&#160;Jan&#160;23&#160;12:32&#160;UTC | true | [view](CERTS/b1085c4ed6a32a87dc773d65e4e7aec591f7625ad80f63ee690d56b74b7a1430/README.md) |
| 03&#160;Dec&#160;22&#160;04:11&#160;UTC | SHAKEN 345J | 02&#160;Jan&#160;23&#160;04:11&#160;UTC | true | [view](CERTS/8991e6ec3e951d1b91f0e0f63b5bc9bdc929b5c76a73fea6c5af7a6401d0aa6e/README.md) |
| 03&#160;Dec&#160;22&#160;05:00&#160;UTC | SHAKEN 345J | 02&#160;Jan&#160;23&#160;05:00&#160;UTC | true | [view](CERTS/df273ff8f12fbf2fda88e527381fba4ac91927b0be97ec38a2d67993a5e0c074/README.md) |
| 05&#160;Dec&#160;22&#160;15:23&#160;UTC | SHAKEN 8526 | 04&#160;Jan&#160;23&#160;15:23&#160;UTC | true | [view](CERTS/29453875467b9cbda7eccf1ee68baaa883140c03b133004d5f21864cec03a029/README.md) |
| 12&#160;Dec&#160;22&#160;15:43&#160;UTC | SHAKEN 722J | 11&#160;Jan&#160;23&#160;15:43&#160;UTC | true | [view](CERTS/234bb11af3db6a37bc8a1807ba5458331e17031ec7772a9d790809a0c2856761/README.md) |
| 12&#160;Dec&#160;22&#160;18:12&#160;UTC | SHAKEN 952J | 11&#160;Jan&#160;23&#160;18:12&#160;UTC | true | [view](CERTS/060316d5062c63a01850cffc57dd7d976e37d02e17d9f77d165ba3a16c1172e0/README.md) |
| 22&#160;Dec&#160;22&#160;17:30&#160;UTC | SHAKEN 505J | 22&#160;Mar&#160;23&#160;17:30&#160;UTC | true | [view](CERTS/7ff43923df80c3a7e2d854b42f33c69069af4625f5267df0d7ff7b2004ad1c1b/README.md) |
| 24&#160;Dec&#160;22&#160;04:21&#160;UTC | SHAKEN 345J | 23&#160;Jan&#160;23&#160;04:21&#160;UTC | true | [view](CERTS/4fef66289eb1ff86ee5182af55c3cb2a6d4050e31fd08e250807b47e7d2b45ea/README.md) |
| 24&#160;Dec&#160;22&#160;05:00&#160;UTC | SHAKEN 345J | 23&#160;Jan&#160;23&#160;05:00&#160;UTC | true | [view](CERTS/1046866f335f424f2678b093d3af73a2a75c9cf37302e75a94954054f092fc91/README.md) |
| 26&#160;Dec&#160;22&#160;12:17&#160;UTC | SHAKEN 130B | 02&#160;Jan&#160;23&#160;12:17&#160;UTC | true | [view](CERTS/0fdde59ab18f3c1c61f0b461aac512241f01c43657a8b65c60cee3192564c27b/README.md) |
| 26&#160;Dec&#160;22&#160;14:30&#160;UTC | SHAKEN 841J | 09&#160;Jan&#160;23&#160;14:30&#160;UTC | true | [view](CERTS/b06c04c6087013d23450759d1cb3836b75f437037e0ac6ee2a32cd8e05ee33cb/README.md) |
| 26&#160;Dec&#160;22&#160;16:38&#160;UTC | SHAKEN 291K | 02&#160;Jan&#160;23&#160;16:38&#160;UTC | true | [view](CERTS/595617ce853002da0cf0892715a88ae7d8c42a124becf9885d33a084067d9ad6/README.md) |
| 26&#160;Dec&#160;22&#160;20:20&#160;UTC | SHAKEN 983J | 02&#160;Jan&#160;23&#160;20:20&#160;UTC | true | [view](CERTS/343c0c0d68edc5dad41046794f405bc20b7dfb77a5a1e96c928812e47d63d6da/README.md) |
| 26&#160;Dec&#160;22&#160;20:30&#160;UTC | SHAKEN 738J | 04&#160;Jan&#160;23&#160;20:30&#160;UTC | true | [view](CERTS/c1adec3417362012ea155b8ccbf3b1128f59e82c6148f02f9e194c7c16a08a9e/README.md) |
| 27&#160;Dec&#160;22&#160;01:45&#160;UTC | SHAKEN 278K | 03&#160;Jan&#160;23&#160;01:45&#160;UTC | true | [view](CERTS/fa4ab69b5b611bd41e5e6f606e650c0e3de2c1fb7f1e1d70fce92875cb625a33/README.md) |
| 27&#160;Dec&#160;22&#160;14:38&#160;UTC | SHAKEN 2550 | 03&#160;Jan&#160;23&#160;14:38&#160;UTC | true | [view](CERTS/43ec2ef05470a560bb602d6b166e026bdeefa12d40f0b7e140d225435f09fb01/README.md) |
| 27&#160;Dec&#160;22&#160;20:17&#160;UTC | SHAKEN 177K | 03&#160;Jan&#160;23&#160;20:17&#160;UTC | true | [view](CERTS/39d9fc6d189d2c923322191e229ed0801b47142d25ec531ab1a07acfbdabdb21/README.md) |
| 28&#160;Dec&#160;22&#160;14:46&#160;UTC | SHAKEN 186K | 04&#160;Jan&#160;23&#160;14:46&#160;UTC | true | [view](CERTS/95698af34583f5d3751faaf361c67d01663882cea9cfff3adf371995502a7cb9/README.md) |
| 28&#160;Dec&#160;22&#160;17:29&#160;UTC | SHAKEN 107K | 04&#160;Jan&#160;23&#160;17:29&#160;UTC | true | [view](CERTS/9b0f6d060380a282a1f2786f5c8a69f56b85c7b9cc526cbb675f6994cce974f4/README.md) |
| 28&#160;Dec&#160;22&#160;18:41&#160;UTC | SHAKEN 056K | 04&#160;Jan&#160;23&#160;18:41&#160;UTC | true | [view](CERTS/452d11b5a5a6afa74d3be2570c6f5046bb8540662d568e37e282509aa414ef53/README.md) |
| 28&#160;Dec&#160;22&#160;20:07&#160;UTC | SHAKEN 366G | 04&#160;Jan&#160;23&#160;20:07&#160;UTC | true | [view](CERTS/ca330edd2ca2c02d7dc74d9a9b70511c887358a717d39d97343421a449a913da/README.md) |
| 28&#160;Dec&#160;22&#160;20:22&#160;UTC | SHAKEN 674J | 04&#160;Jan&#160;23&#160;20:22&#160;UTC | true | [view](CERTS/68643c36899359b24a311581366fc952b2510c21bb7d76d4b4e28bfe993de89e/README.md) |
| 28&#160;Dec&#160;22&#160;20:24&#160;UTC | SHAKEN 738J | 04&#160;Jan&#160;23&#160;20:24&#160;UTC | true | [view](CERTS/21f5861a17a7b08a332ce1c26fa72e6fdcce7feb0b080330eb2a9956e6882ef0/README.md) |
| 28&#160;Dec&#160;22&#160;20:25&#160;UTC | SHAKEN 733J | 04&#160;Jan&#160;23&#160;20:25&#160;UTC | true | [view](CERTS/f7f106957c38628e0627e93487b4b09aa3fe316b2474b8eac076f64650032979/README.md) |
| 28&#160;Dec&#160;22&#160;20:26&#160;UTC | SHAKEN 819H | 04&#160;Jan&#160;23&#160;20:26&#160;UTC | true | [view](CERTS/b06beb64fddbbbda770bf848d29ceab3e3675150b626d78a9f8a797ff1a1a47b/README.md) |
| 28&#160;Dec&#160;22&#160;20:29&#160;UTC | SHAKEN 469A | 04&#160;Jan&#160;23&#160;20:29&#160;UTC | true | [view](CERTS/68c61fde816a89a35aeb37c468384a5b532de5cc8eea413e793f57392ac1cc2d/README.md) |
| 28&#160;Dec&#160;22&#160;20:29&#160;UTC | SHAKEN 849J | 04&#160;Jan&#160;23&#160;20:29&#160;UTC | true | [view](CERTS/c1d860b5ff45d7f29eab5f88ecb6e9e60cd026171719ed8abeeacb0d39034582/README.md) |
| 28&#160;Dec&#160;22&#160;20:31&#160;UTC | SHAKEN 738J | 04&#160;Jan&#160;23&#160;20:31&#160;UTC | true | [view](CERTS/3772f73ee499298f29c4c982ee3ea09a774177e7316557c30e25fbcd1d4592ef/README.md) |
| 28&#160;Dec&#160;22&#160;20:31&#160;UTC | SHAKEN 459J | 04&#160;Jan&#160;23&#160;20:31&#160;UTC | true | [view](CERTS/b4e1fdeac43850cafe6d8991c5275004c649909c68715fbe968af78671c76be2/README.md) |
| 28&#160;Dec&#160;22&#160;20:32&#160;UTC | SHAKEN 366G | 04&#160;Jan&#160;23&#160;20:32&#160;UTC | true | [view](CERTS/3677c43f9acabd5df69d04af47d5ff0818881ef063310211f1d23a2348595643/README.md) |
| 28&#160;Dec&#160;22&#160;20:32&#160;UTC | SHAKEN 738J | 04&#160;Jan&#160;23&#160;20:32&#160;UTC | true | [view](CERTS/0d69a5bf8f75286522a694e20cad3eccb012c5d82c7ae27337064dac62107fa2/README.md) |
| 29&#160;Dec&#160;22&#160;15:44&#160;UTC | SHAKEN 551G | 05&#160;Jan&#160;23&#160;15:44&#160;UTC | true | [view](CERTS/87541e5ae8c444b412cd50cf9324bce446b5cf0e57ec4e132a2921f4526bab6b/README.md) |

#### CA Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 20&#160;Aug&#160;21&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Issuing CA3 | 19&#160;Aug&#160;31&#160;23:59&#160;UTC | true | [view](CERTS/51b51431ffd8055cd7e30c213540f4868258058d4c4c9f23f1ed8bea75101e37/README.md) |
| 20&#160;Aug&#160;21&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Issuing CA1 | 19&#160;Aug&#160;31&#160;23:59&#160;UTC | true | [view](CERTS/037a7c6621814bd5f01fa2080a77a245458694d9d566e0050e068fc853317857/README.md) |
| 20&#160;Aug&#160;21&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Root CA1 | 19&#160;Aug&#160;41&#160;23:59&#160;UTC | true | [view](CERTS/59ab943f4e8c3c17755cc8e4abb5ac65736ce74ec587c62cb744ab0babbbe2fe/README.md) |
| 24&#160;Oct&#160;22&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Issuing CA4 | 23&#160;Oct&#160;32&#160;23:59&#160;UTC | false | [view](CERTS/80dd12d935f46c256c7c1289a0834b21ec8f2f8c35f2864928d1e75f3a280665/README.md) |
| 24&#160;Oct&#160;22&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Root CA2 | 23&#160;Oct&#160;42&#160;23:59&#160;UTC | false | [view](CERTS/a26e04fc786ab70b8085236b2c53f8cfbf5d0c6a5c2c9c3e9f91669fbb8ea4d5/README.md) |


Generated: 01 Jan 23 23:34 UTC