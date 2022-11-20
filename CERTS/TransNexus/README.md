# STIR/SHAKEN CA Ecosystem Compliance

## TransNexus

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 464 certificates were included in the corpus being tested
- 1 certificates in the corpus were skipped because they are duplicates
- 419 certificates in the corpus were skipped because they are expired
- 1 certificates in the corpus were skipped because they are not currently trusted
- 43 certificates being tested against the remaining rules
- 2.26 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 65.12% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 2.33% of certificates are too old to be assessed against currently enforced expectations
- 110 days is the average remaining validity for the certificates in the corpus
- 109 days is the average initial validity for the certificates in the corpus
- 26 certificates expire in the next 30 days

| Instances | Test | Source |
|-----------|------|--------|
| 3 | [e_atis_certificate_policies](ISSUES/e_atis_certificate_policies/README.md) | ATIS1000080 |
| 20 | [e_incorrect_ku_encoding](ISSUES/e_incorrect_ku_encoding/README.md) | RFC5280 |
| 3 | [e_us_cp_ambiguous_identifier](ISSUES/e_us_cp_ambiguous_identifier/README.md) | US_SHAKEN_CP |
| 43 | [e_us_cp_subject_sn](ISSUES/e_us_cp_subject_sn/README.md) | US_SHAKEN_CP |
| 28 | [w_shaken_subject_rdn_unknown](ISSUES/w_shaken_subject_rdn_unknown/README.md) | SHAKEN_PKI_BEST_PRACTICES |

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
- 5474 days is the average remaining validity for the certificates in the corpus
- 5113 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days

| Instances | Test | Source |
|-----------|------|--------|
| 3 | [e_incorrect_ku_encoding](ISSUES/e_incorrect_ku_encoding/README.md) | RFC5280 |
| 3 | [e_us_cp_ca_key_usage_crl_sign](ISSUES/e_us_cp_ca_key_usage_crl_sign/README.md) | US_SHAKEN_CP |
| 3 | [w_shaken_ca_subject_rdn_unknown](ISSUES/w_shaken_ca_subject_rdn_unknown/README.md) | SHAKEN_PKI_BEST_PRACTICES |

### Details

#### Leaf Certificates

| Created at | Subject | Problems | Link |
|------------|---------|----------|------|
| 13 Jan 22 14:17 UTC | MobileSphere SHAKEN 873J | true | [view](CERTS/8f74ddd5e6964042b119d5b3f17d24694ec78c8688e0ffc96d13fdb0e8df26d4/README.md) |
| 30 Mar 22 12:54 UTC | Fusion Connect SHAKEN 2720 | true | [view](CERTS/03a010b294807e90cc1309d7466fd2ddc591158a69be6950ff4d0eab32de0799/README.md) |
| 27 Apr 22 18:22 UTC | CCI SHAKEN 663J | true | [view](CERTS/9cfd0b9974acdc9b5fce3888cf2f6cf99173e7955adeebd37a1e60dfc444a0b3/README.md) |
| 10 Jun 22 14:00 UTC | SHAKEN 597J | true | [view](CERTS/4baa838eb96ce358aa1565d0add8b75072258efd5d21a5bbb0eec878026a67be/README.md) |
| 20 Jun 22 20:25 UTC | SHAKEN 857J | true | [view](CERTS/e1e0a4851fb7b51d617959c00fb6a17e7b23d34e1555ba70ab01c9182a2f85f8/README.md) |
| 22 Jun 22 18:46 UTC | SHAKEN 755J | true | [view](CERTS/a8cac316fddf25b48b9350a325791be8b87e6484a5959ab605db1ae5e3e99a84/README.md) |
| 29 Jun 22 20:24 UTC | SHAKEN 736J | true | [view](CERTS/71b25a480077816998e48e6a76f272e70aa91ecb19b0e773cd8777823cdead32/README.md) |
| 25 Jul 22 18:36 UTC | SHAKEN 578J | true | [view](CERTS/f4af5a28131d680ef2682ae86181b64f0c919864d7dbaef00c0598d314b537a1/README.md) |
| 10 Aug 22 18:11 UTC | SHAKEN 073H | true | [view](CERTS/87dabb91a689a6ac540b5ffbf5eb89326fada83fb822d2bb9ec31b68251b579e/README.md) |
| 01 Sep 22 03:25 UTC | SHAKEN 6628 | true | [view](CERTS/37b9a67bb8e3272048330f269b5dbc285ac0278df4670ab7e291b92b6548fa2d/README.md) |
| 13 Sep 22 17:24 UTC | SHAKEN 706J | true | [view](CERTS/1eea5e17ae8b36cb5b6d657178273af745cea015e63586e6d44d1ba2ee8f722a/README.md) |
| 26 Sep 22 18:40 UTC | SHAKEN 505J | true | [view](CERTS/b186c959554e652b69824a2e45ac08ec135105b0e01c9d2eadeb3cf46c130670/README.md) |
| 26 Sep 22 18:42 UTC | SHAKEN 505J | true | [view](CERTS/1c0003ac10eeaa04229e507c15f71ee018c3902c3e8c20fbe42e533b8682ba8c/README.md) |
| 01 Oct 22 12:31 UTC | SHAKEN 193E | true | [view](CERTS/2a2a297de4aa7620b8f38d5266dd4f76fbd6b5c79a87411b589c2226518ce475/README.md) |
| 12 Oct 22 19:39 UTC | SHAKEN 815G | true | [view](CERTS/a79f50d4fcd3e680518379f12a86a4b9a21c5e27dc5422a8439525edf350c4a7/README.md) |
| 20 Oct 22 15:48 UTC | SHAKEN 622J | true | [view](CERTS/2e34c765f110b197f2ec1052b1960bc25ce320dfce7d7035cc532cd1eb6dbd57/README.md) |
| 22 Oct 22 04:04 UTC | SHAKEN 345J | true | [view](CERTS/27c018db560aefbff83f32326db2cbf3f1260051386569d7775cf98fae145bae/README.md) |
| 22 Oct 22 04:44 UTC | SHAKEN 345J | true | [view](CERTS/109a4ac394ec4e6eb6ba42d2ad6bc799f2548629a8bbfb9420b9ec1ca5df6ff9/README.md) |
| 25 Oct 22 15:33 UTC | SHAKEN 722J | true | [view](CERTS/50eb0c6a670f4122f8fbdb75582aa257fe3979f441fb396ff738372627104f9c/README.md) |
| 26 Oct 22 10:43 UTC | SHAKEN 8526 | true | [view](CERTS/b68dadcfa6267b8e9dd012e42154b55cab6d3694f4543e46af44bab1e4ba971e/README.md) |
| 31 Oct 22 12:32 UTC | SHAKEN 193E | true | [view](CERTS/0de5e0e6787b9aa19c02c014821dd7a2e62e438f4eb9fbbb761eb86df8c69ff7/README.md) |
| 31 Oct 22 18:04 UTC | SHAKEN 952J | true | [view](CERTS/93040a7e0cb4876c378dd10737f7dec5d054a6a744ebec63e17f8adda605a6ca/README.md) |
| 01 Nov 22 11:14 UTC | SHAKEN 807J | true | [view](CERTS/0d97796c10ef21ba02a0f7954d82ace5e41d7ea66e4ea9eb16801c3a1a1d9f03/README.md) |
| 12 Nov 22 04:05 UTC | SHAKEN 345J | true | [view](CERTS/ac38b9036e27ff2b753a0727e8ccdfd800d459dfdf595454c6a9dc7e9cd8ad74/README.md) |
| 12 Nov 22 04:54 UTC | SHAKEN 345J | true | [view](CERTS/0a4891c263ce12ebb5116b43048b2b427314bd6133848201f8186b62f7efc68a/README.md) |
| 14 Nov 22 14:10 UTC | SHAKEN 841J | true | [view](CERTS/51cf29551891eb1f1935e83cc86a50c52807f84c49b3b7c05461a2d89e4d8a75/README.md) |
| 14 Nov 22 16:23 UTC | SHAKEN 035K | true | [view](CERTS/767cae91c22bd52fdeeced9de496bbb6bbfeff8dbc54acf18c7f21aa93e0e558/README.md) |
| 14 Nov 22 19:09 UTC | SHAKEN 749J | true | [view](CERTS/6860688fe116dcf110fea72d54bd761b855d90a48d51101b0f88c584974d4b0d/README.md) |
| 14 Nov 22 20:15 UTC | SHAKEN 983J | true | [view](CERTS/e0d9e4f7ddc0c72df726419b95ec96a7bd285bbc981153d07f31e180e0c48f65/README.md) |
| 14 Nov 22 20:24 UTC | SHAKEN 738J | true | [view](CERTS/aa1140caf05a17dfee3e6713e4e26a30cd29442e0dc0d5ffc29a2522cef3408a/README.md) |
| 14 Nov 22 23:52 UTC | SHAKEN 970J | true | [view](CERTS/c31c4a2bcc4fe8bef479ac03fd220b1bdbdda164c4ff7bad02bbb192917e3631/README.md) |
| 15 Nov 22 01:42 UTC | SHAKEN 278K | true | [view](CERTS/aebcc6da84cc77550e791eec626797199a17d392930eeb129dc359a0cc90f7f3/README.md) |
| 15 Nov 22 12:42 UTC | SHAKEN 551G | true | [view](CERTS/4069553710aac18a7170ed4bc2abc5c4f2e8fb28137b9c926da299bb097196a9/README.md) |
| 15 Nov 22 20:08 UTC | SHAKEN 864J | true | [view](CERTS/32df47cc081f6599146db7bf8f967eb7da700d573c926706d65e69d691ae0927/README.md) |
| 16 Nov 22 20:20 UTC | SHAKEN 738J | true | [view](CERTS/b42e7ab46fb071123047e9bc872f760710c6185e3b74d2e969f70d29bba199c4/README.md) |
| 16 Nov 22 20:25 UTC | SHAKEN 738J | true | [view](CERTS/51c22d8787b9404a94ecb566f5c62e8506cfbc5231a97da4e60bcb2ac0da1a0a/README.md) |
| 16 Nov 22 20:26 UTC | SHAKEN 366G | true | [view](CERTS/c633e6ebaa2ecb45d6b4d724d88a6de6d695706a83e31c05429484c3df57f0d3/README.md) |
| 16 Nov 22 20:26 UTC | SHAKEN 738J | true | [view](CERTS/4feb69bdbe9ef2b93c4652f1a803c27c4f252d539e4e3000eb6f0d1d9760fe6e/README.md) |
| 18 Nov 22 14:35 UTC | SHAKEN 2550 | true | [view](CERTS/9c167daffac6d99a821bdf0aa3710d1fb5fa23f8307746358286a2c0a0ffae68/README.md) |
| 18 Nov 22 15:42 UTC | SHAKEN 722J | true | [view](CERTS/4111bb56ba547a2f493d1e9c2baefa33ce9bd117904133477d063d6fa0d6b9b2/README.md) |
| 18 Nov 22 18:04 UTC | SHAKEN 952J | true | [view](CERTS/b05bb158c054b6e579fe20750c83ba5524896545d629b96b51d4c8d9ab38dfd3/README.md) |
| 18 Nov 22 20:08 UTC | SHAKEN 864J | true | [view](CERTS/0f7f7697354353bd9d2ab9242343d42f4a12e66045709b0fb7e50469e313f8b3/README.md) |
| 19 Nov 22 20:20 UTC | SHAKEN 738J | true | [view](CERTS/2372aac07c78cfbaf7c5d62d5c2ef1d700a953f26e3950ffb6f6f97f54601739/README.md) |

#### CA Certificates

| Created at | Subject | Problems | Link |
|------------|---------|----------|------|
| 20 Aug 21 00:00 UTC | TransNexus, Inc. SHAKEN Issuing CA3 | true | [view](CERTS/51b51431ffd8055cd7e30c213540f4868258058d4c4c9f23f1ed8bea75101e37/README.md) |
| 20 Aug 21 00:00 UTC | TransNexus, Inc. SHAKEN Issuing CA1 | true | [view](CERTS/037a7c6621814bd5f01fa2080a77a245458694d9d566e0050e068fc853317857/README.md) |
| 20 Aug 21 00:00 UTC | TransNexus, Inc. SHAKEN Root CA1 | true | [view](CERTS/59ab943f4e8c3c17755cc8e4abb5ac65736ce74ec587c62cb744ab0babbbe2fe/README.md) |
| 24 Oct 22 00:00 UTC | TransNexus, Inc. SHAKEN Issuing CA4 | false | [view](CERTS/80dd12d935f46c256c7c1289a0834b21ec8f2f8c35f2864928d1e75f3a280665/README.md) |
| 24 Oct 22 00:00 UTC | TransNexus, Inc. SHAKEN Root CA2 | false | [view](CERTS/a26e04fc786ab70b8085236b2c53f8cfbf5d0c6a5c2c9c3e9f91669fbb8ea4d5/README.md) |


Generated: 20 Nov 22 22:57 UTC