# STIR/SHAKEN CA Ecosystem Compliance

## TransNexus

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 582 certificates were included in the corpus being tested
- 1 certificates in the corpus were skipped because they are duplicates
- 504 certificates in the corpus were skipped because they are expired
- 1 certificates in the corpus were skipped because they are not currently trusted
- 76 certificates being tested against the remaining rules
- 1.57 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 27.63% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 1.32% of certificates are too old to be assessed against currently enforced expectations
- 68 days is the average remaining validity for the certificates in the corpus
- 67 days is the average initial validity for the certificates in the corpus
- 60 certificates expire in the next 30 days
- 1.36 average number of unexpired certificates per OCN observed
- 56 unique OCNs observed in unexpired and valid certificate corpus

| Instances | Test | Source |
|-----------|------|--------|
| 3 | [e_atis_certificate_policies](ISSUES/e_atis_certificate_policies/README.md) | ATIS1000080 |
| 16 | [e_incorrect_ku_encoding](ISSUES/e_incorrect_ku_encoding/README.md) | RFC5280 |
| 3 | [e_us_cp_ambiguous_identifier](ISSUES/e_us_cp_ambiguous_identifier/README.md) | US_SHAKEN_CP |
| 76 | [e_us_cp_subject_sn](ISSUES/e_us_cp_subject_sn/README.md) | US_SHAKEN_CP |
| 21 | [w_shaken_subject_rdn_unknown](ISSUES/w_shaken_subject_rdn_unknown/README.md) | SHAKEN_PKI_BEST_PRACTICES |

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
- 5472 days is the average remaining validity for the certificates in the corpus
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
| 23 Aug 22 04:14 UTC | SHAKEN 866J | true | [view](CERTS/4c4fdb320c51c8582d595c2d307cd770d409daf533ff573bcc73afaed83f6b7d/README.md) |
| 01 Sep 22 03:25 UTC | SHAKEN 6628 | true | [view](CERTS/37b9a67bb8e3272048330f269b5dbc285ac0278df4670ab7e291b92b6548fa2d/README.md) |
| 13 Sep 22 17:24 UTC | SHAKEN 706J | true | [view](CERTS/1eea5e17ae8b36cb5b6d657178273af745cea015e63586e6d44d1ba2ee8f722a/README.md) |
| 26 Sep 22 18:40 UTC | SHAKEN 505J | true | [view](CERTS/b186c959554e652b69824a2e45ac08ec135105b0e01c9d2eadeb3cf46c130670/README.md) |
| 26 Sep 22 18:42 UTC | SHAKEN 505J | true | [view](CERTS/1c0003ac10eeaa04229e507c15f71ee018c3902c3e8c20fbe42e533b8682ba8c/README.md) |
| 12 Oct 22 19:39 UTC | SHAKEN 815G | true | [view](CERTS/a79f50d4fcd3e680518379f12a86a4b9a21c5e27dc5422a8439525edf350c4a7/README.md) |
| 20 Oct 22 15:48 UTC | SHAKEN 622J | true | [view](CERTS/2e34c765f110b197f2ec1052b1960bc25ce320dfce7d7035cc532cd1eb6dbd57/README.md) |
| 31 Oct 22 12:32 UTC | SHAKEN 193E | true | [view](CERTS/0de5e0e6787b9aa19c02c014821dd7a2e62e438f4eb9fbbb761eb86df8c69ff7/README.md) |
| 31 Oct 22 18:04 UTC | SHAKEN 952J | true | [view](CERTS/93040a7e0cb4876c378dd10737f7dec5d054a6a744ebec63e17f8adda605a6ca/README.md) |
| 01 Nov 22 11:14 UTC | SHAKEN 807J | true | [view](CERTS/0d97796c10ef21ba02a0f7954d82ace5e41d7ea66e4ea9eb16801c3a1a1d9f03/README.md) |
| 12 Nov 22 04:05 UTC | SHAKEN 345J | true | [view](CERTS/ac38b9036e27ff2b753a0727e8ccdfd800d459dfdf595454c6a9dc7e9cd8ad74/README.md) |
| 12 Nov 22 04:54 UTC | SHAKEN 345J | true | [view](CERTS/0a4891c263ce12ebb5116b43048b2b427314bd6133848201f8186b62f7efc68a/README.md) |
| 18 Nov 22 15:42 UTC | SHAKEN 722J | true | [view](CERTS/4111bb56ba547a2f493d1e9c2baefa33ce9bd117904133477d063d6fa0d6b9b2/README.md) |
| 18 Nov 22 18:04 UTC | SHAKEN 952J | true | [view](CERTS/b05bb158c054b6e579fe20750c83ba5524896545d629b96b51d4c8d9ab38dfd3/README.md) |
| 21 Nov 22 20:25 UTC | SHAKEN 738J | true | [view](CERTS/6b8346b320b4956f1b2726c057062818646170e7ae6f65614dc498cf80947979/README.md) |
| 22 Nov 22 23:52 UTC | SHAKEN 841J | true | [view](CERTS/21ac6662af2208c0149298f8f075e00220d1518fa1a839cb770d834f878afc00/README.md) |
| 23 Nov 22 17:31 UTC | SHAKEN 4036 | true | [view](CERTS/e03ff4259e6b456f1bb19d011f005932e83b0a35b520730c97cd77ad8868071c/README.md) |
| 23 Nov 22 20:16 UTC | SHAKEN 983J | true | [view](CERTS/a31233b83663326a8d5caa09924ed4e0e021cc9e59911ae20667557b2a37c745/README.md) |
| 24 Nov 22 14:36 UTC | SHAKEN 2550 | true | [view](CERTS/bd435a93d616a1848c46e8a8364ad7ad4491537c9a2cd500e61592dfba64a106/README.md) |
| 25 Nov 22 08:02 UTC | SHAKEN 551G | true | [view](CERTS/d7525ebf1cf60ce77f0390b661c1125dc30886f532523e25e394e7cd5b02551a/README.md) |
| 25 Nov 22 17:26 UTC | SHAKEN 107K | true | [view](CERTS/d05e96555e640a58d48cafb95fa9bcb0f8310dd5244b31573639c41841aaa6a0/README.md) |
| 25 Nov 22 19:57 UTC | SHAKEN 297K | true | [view](CERTS/dc29343cb4842a14b4c425cbc0c68e347e6d3e7fc5b2473fae7c18025ecf38e6/README.md) |
| 25 Nov 22 20:20 UTC | SHAKEN 735J | true | [view](CERTS/f2b2ec0d3422c98fc85e48d519d399716b0caca1236df5fdba0c9cf496c349e5/README.md) |
| 25 Nov 22 20:20 UTC | SHAKEN 738J | true | [view](CERTS/9bb764757d413207075209248b8f35df0d3a056609e08d1f047b8eec64b37d9e/README.md) |
| 25 Nov 22 20:20 UTC | SHAKEN 733J | true | [view](CERTS/97694297d993ee44cc7bbd9daaf48646fa9c89e0f6369fd4ad3055fe276a3a81/README.md) |
| 25 Nov 22 20:24 UTC | SHAKEN 849J | true | [view](CERTS/08a7f05d6988980c3b8145eec141bad858c56bc68804167f803d05382ea45acf/README.md) |
| 25 Nov 22 20:24 UTC | SHAKEN 691A | true | [view](CERTS/317f684b3eab18c64c9a9ee018cb1ce58c9b8a7284e77d792f83fff3425d607a/README.md) |
| 25 Nov 22 20:24 UTC | SHAKEN 469A | true | [view](CERTS/36eea868c67ac1e482d9107028c8f2aca59ddd3d19152485442fcf18eca6e1c6/README.md) |
| 25 Nov 22 20:25 UTC | SHAKEN 738J | true | [view](CERTS/6b212d2bc8e603f90ff5337b46267a351999bea8e0a33680604d585664ba04cd/README.md) |
| 25 Nov 22 20:25 UTC | SHAKEN 459J | true | [view](CERTS/4ec87ab18fe914eed7ff1b6ef744e02b058d4f3f65077b69811a91d64c2b099e/README.md) |
| 25 Nov 22 20:25 UTC | SHAKEN 672B | true | [view](CERTS/0d5b5096b1cb9f4f4132c6d6ce4a98ab76d4b31a7927a635da40cd7ea4154842/README.md) |
| 25 Nov 22 20:26 UTC | SHAKEN 366G | true | [view](CERTS/6191a25456dabca43312150a58545151c39ad6d38261c09bedd62782c1fa7e1e/README.md) |
| 25 Nov 22 20:27 UTC | SHAKEN 0226 | true | [view](CERTS/c50f237291f814bbc29b8294ccf88c421554e661636844a8dba77d6840f5cdad/README.md) |
| 25 Nov 22 20:27 UTC | SHAKEN 738J | true | [view](CERTS/7dd772716339dbfa1a387102aae81796a254702e0512a1052c2159bd5a2b6e74/README.md) |
| 25 Nov 22 21:02 UTC | SHAKEN 518J | true | [view](CERTS/eae143d71eb0e9d1e2567e9ed60a74262056af67e00e222e9bcd4c4967ca1d24/README.md) |
| 26 Nov 22 12:15 UTC | SHAKEN 130B | true | [view](CERTS/75f49c92c06174c0cd5c0be0a00e95e6231c82b73e8346ca874b1157a2ac53dc/README.md) |
| 26 Nov 22 13:21 UTC | SHAKEN 0172 | true | [view](CERTS/ca6aa82c57dd3759d91bface6744ba1eb1f66250d9903d1569a8ce4987f1d055/README.md) |
| 26 Nov 22 19:10 UTC | SHAKEN 749J | true | [view](CERTS/7c27a9d1b94b73f550d9ab13f422b606046d40b51aeef0928bc89783b109263b/README.md) |
| 26 Nov 22 20:16 UTC | SHAKEN 983J | true | [view](CERTS/c0dd50071168f308c2b17fb8817b79cce8ebb550dcf2bb5aa09fea3f6b0fee4c/README.md) |
| 26 Nov 22 21:44 UTC | SHAKEN 606F | true | [view](CERTS/24a04f3c49b09c28f1771afd8619837e3866a80ffbd7c2a7c301825440f5ffe6/README.md) |
| 27 Nov 22 01:42 UTC | SHAKEN 278K | true | [view](CERTS/04a072d6405a0c1e1e7e6893f7467ec9b7cb4aac0bd2e931039e472e7d8b407a/README.md) |
| 27 Nov 22 14:36 UTC | SHAKEN 2550 | true | [view](CERTS/3f37e0eb5fb459beae9c4695866e96b30b0951c8b69ce00c8dfe53c6babae06d/README.md) |
| 27 Nov 22 18:48 UTC | SHAKEN 140K | true | [view](CERTS/66d303a259302adc9fcc8552ee79b88513c8f8fcc17727c3c2736024ba331305/README.md) |
| 27 Nov 22 20:14 UTC | SHAKEN 177K | true | [view](CERTS/76d8c73358c5e3473bbd12634cb431ccd0e1e35eea838113c0479bedd86e0952/README.md) |
| 28 Nov 22 14:44 UTC | SHAKEN 012K | true | [view](CERTS/1a8df3282e53e1e16e2f8fd0c010236f9eb73feecf4820a45d09689377c21a48/README.md) |
| 28 Nov 22 14:45 UTC | SHAKEN 747J | true | [view](CERTS/51aba108c9adb06a0cbf3169f52def2494ae0e34809db5e8be805b2dd6b10fa0/README.md) |
| 28 Nov 22 17:27 UTC | SHAKEN 107K | true | [view](CERTS/6b3fa92a1b668430d01c7d5feb4d1437309bd965d325946f52bbb9baa9799cc7/README.md) |
| 28 Nov 22 18:39 UTC | SHAKEN 056K | true | [view](CERTS/d479d62f9ad73674b85b522676a943bdae2b98bbe84bdf2311f67be99928f9ba/README.md) |
| 28 Nov 22 19:57 UTC | SHAKEN 297K | true | [view](CERTS/0c8d278fefe556489a1e55827741152fe5367e9789f23d68c016c7cf433f94d8/README.md) |
| 28 Nov 22 20:03 UTC | SHAKEN 366G | true | [view](CERTS/bc2302e39f66a086aff6bc7ff6a1f6fa3179bc2594a661b3adc7f887c2ca75bd/README.md) |
| 28 Nov 22 20:19 UTC | SHAKEN 674J | true | [view](CERTS/f4aca516210d1898cf1235c07c0ae7240fbcce66de7da84ec74d1c0678c4b439/README.md) |
| 28 Nov 22 20:21 UTC | SHAKEN 700H | true | [view](CERTS/6c12cec9743eca93fedcde12ece8ca39102fd2f509d403537ac672c52bbd0e29/README.md) |
| 28 Nov 22 20:22 UTC | SHAKEN 819H | true | [view](CERTS/8820b9c41789203bacd4f4d8f1e615600b90b1a04d0394119d6a1462c3e9598e/README.md) |
| 28 Nov 22 20:23 UTC | SHAKEN 769J | true | [view](CERTS/2189c0e66d59dfe3a38d7b81bab0c002f31f828204b03983611086b5d7065c7b/README.md) |
| 28 Nov 22 20:23 UTC | SHAKEN 589J | true | [view](CERTS/a2ae8ed2d98786e7540b2372026ca9b434330a6643ad523fce30a1ef28c3e46a/README.md) |
| 28 Nov 22 20:24 UTC | SHAKEN 849J | true | [view](CERTS/e9a606b3c321c52a53561a49e8181ab30d9481abd89ef2cbfc34623275efac03/README.md) |
| 28 Nov 22 20:25 UTC | SHAKEN 469A | true | [view](CERTS/d1abdb1d3c297b034dcb54efb78ec58ff7806c28ae4d0fbcf632d35729d15d10/README.md) |
| 28 Nov 22 20:25 UTC | SHAKEN 495J | true | [view](CERTS/b595d811e8b22f27492993fcce6b1b22c327b71845bcf6607cb70f63a352dab8/README.md) |
| 28 Nov 22 20:25 UTC | SHAKEN 790J | true | [view](CERTS/c473c8d369026351484772a5b550f938ad4d270663b27ef9ddfe28d12cb5385b/README.md) |
| 28 Nov 22 20:26 UTC | SHAKEN 738J | true | [view](CERTS/e40280dfe9a62708996198fc786e18c7578f37d284ff1e2c9d2cf24e94b878ee/README.md) |
| 28 Nov 22 20:26 UTC | SHAKEN 1680 | true | [view](CERTS/a0f8ded4427f795b60b93b3c7d2eb401be0e255137a7d7c1e8375d6d054e321b/README.md) |
| 28 Nov 22 20:26 UTC | SHAKEN 738J | true | [view](CERTS/b6078f94356cbe08f80c2e43db13b38115962d906b6bf2767a06897070f338a3/README.md) |
| 28 Nov 22 20:26 UTC | SHAKEN 459J | true | [view](CERTS/e98dfa73bf9179dc6cd12f1ac7dc2cfdb406697e3c805096ae3e07f71faac92b/README.md) |
| 28 Nov 22 20:27 UTC | SHAKEN 366G | true | [view](CERTS/67363bbcda7164eb64222794b1c59fdce0151272522557d225fbcf0993c24f16/README.md) |
| 28 Nov 22 20:27 UTC | SHAKEN 0226 | true | [view](CERTS/f6877a0d9a658648d53c76be801797cab0f542e5a5a39f925e4eff01ae8621e3/README.md) |
| 28 Nov 22 20:27 UTC | SHAKEN 738J | true | [view](CERTS/50bc2517929b0699dc7d4589e12156a35e353b1cc14e942b4b2868451a8cdf7a/README.md) |
| 28 Nov 22 21:02 UTC | SHAKEN 518J | true | [view](CERTS/26e80150c6353e2f8372c7b8649a6a7699bb72d61c8254c156791299b2bd8dac/README.md) |

#### CA Certificates

| Created at | Subject | Problems | Link |
|------------|---------|----------|------|
| 20 Aug 21 00:00 UTC | TransNexus, Inc. SHAKEN Issuing CA3 | true | [view](CERTS/51b51431ffd8055cd7e30c213540f4868258058d4c4c9f23f1ed8bea75101e37/README.md) |
| 20 Aug 21 00:00 UTC | TransNexus, Inc. SHAKEN Issuing CA1 | true | [view](CERTS/037a7c6621814bd5f01fa2080a77a245458694d9d566e0050e068fc853317857/README.md) |
| 20 Aug 21 00:00 UTC | TransNexus, Inc. SHAKEN Root CA1 | true | [view](CERTS/59ab943f4e8c3c17755cc8e4abb5ac65736ce74ec587c62cb744ab0babbbe2fe/README.md) |
| 24 Oct 22 00:00 UTC | TransNexus, Inc. SHAKEN Issuing CA4 | false | [view](CERTS/80dd12d935f46c256c7c1289a0834b21ec8f2f8c35f2864928d1e75f3a280665/README.md) |
| 24 Oct 22 00:00 UTC | TransNexus, Inc. SHAKEN Root CA2 | false | [view](CERTS/a26e04fc786ab70b8085236b2c53f8cfbf5d0c6a5c2c9c3e9f91669fbb8ea4d5/README.md) |


Generated: 30 Nov 22 17:39 UTC