# STIR/SHAKEN CA Ecosystem Compliance

## TransNexus

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 526 certificates were included in the corpus being tested
- 1 certificates in the corpus were skipped because they are duplicates
- 433 certificates in the corpus were skipped because they are expired
- 1 certificates in the corpus were skipped because they are not currently trusted
- 91 certificates being tested against the remaining rules
- 1.55 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 27.47% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 1.10% of certificates are too old to be assessed against currently enforced expectations
- 58 days is the average remaining validity for the certificates in the corpus
- 57 days is the average initial validity for the certificates in the corpus
- 73 certificates expire in the next 30 days
- 1.69 average number of unexpired certificates per OCN observed
- 54 unique OCNs observed in unexpired and valid certificate corpus

| Instances | Test | Source |
|-----------|------|--------|
| 3 | [e_atis_certificate_policies](ISSUES/e_atis_certificate_policies/README.md) | ATIS1000080 |
| 19 | [e_incorrect_ku_encoding](ISSUES/e_incorrect_ku_encoding/README.md) | RFC5280 |
| 3 | [e_us_cp_ambiguous_identifier](ISSUES/e_us_cp_ambiguous_identifier/README.md) | US_SHAKEN_CP |
| 91 | [e_us_cp_subject_sn](ISSUES/e_us_cp_subject_sn/README.md) | US_SHAKEN_CP |
| 25 | [w_shaken_subject_rdn_unknown](ISSUES/w_shaken_subject_rdn_unknown/README.md) | SHAKEN_PKI_BEST_PRACTICES |

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
| 23 Aug 22 04:14 UTC | SHAKEN 866J | true | [view](CERTS/4c4fdb320c51c8582d595c2d307cd770d409daf533ff573bcc73afaed83f6b7d/README.md) |
| 01 Sep 22 03:25 UTC | SHAKEN 6628 | true | [view](CERTS/37b9a67bb8e3272048330f269b5dbc285ac0278df4670ab7e291b92b6548fa2d/README.md) |
| 13 Sep 22 17:24 UTC | SHAKEN 706J | true | [view](CERTS/1eea5e17ae8b36cb5b6d657178273af745cea015e63586e6d44d1ba2ee8f722a/README.md) |
| 26 Sep 22 18:40 UTC | SHAKEN 505J | true | [view](CERTS/b186c959554e652b69824a2e45ac08ec135105b0e01c9d2eadeb3cf46c130670/README.md) |
| 26 Sep 22 18:42 UTC | SHAKEN 505J | true | [view](CERTS/1c0003ac10eeaa04229e507c15f71ee018c3902c3e8c20fbe42e533b8682ba8c/README.md) |
| 01 Oct 22 12:31 UTC | SHAKEN 193E | true | [view](CERTS/2a2a297de4aa7620b8f38d5266dd4f76fbd6b5c79a87411b589c2226518ce475/README.md) |
| 12 Oct 22 19:39 UTC | SHAKEN 815G | true | [view](CERTS/a79f50d4fcd3e680518379f12a86a4b9a21c5e27dc5422a8439525edf350c4a7/README.md) |
| 20 Oct 22 15:48 UTC | SHAKEN 622J | true | [view](CERTS/2e34c765f110b197f2ec1052b1960bc25ce320dfce7d7035cc532cd1eb6dbd57/README.md) |
| 25 Oct 22 15:33 UTC | SHAKEN 722J | true | [view](CERTS/50eb0c6a670f4122f8fbdb75582aa257fe3979f441fb396ff738372627104f9c/README.md) |
| 26 Oct 22 10:43 UTC | SHAKEN 8526 | true | [view](CERTS/b68dadcfa6267b8e9dd012e42154b55cab6d3694f4543e46af44bab1e4ba971e/README.md) |
| 31 Oct 22 12:32 UTC | SHAKEN 193E | true | [view](CERTS/0de5e0e6787b9aa19c02c014821dd7a2e62e438f4eb9fbbb761eb86df8c69ff7/README.md) |
| 31 Oct 22 18:04 UTC | SHAKEN 952J | true | [view](CERTS/93040a7e0cb4876c378dd10737f7dec5d054a6a744ebec63e17f8adda605a6ca/README.md) |
| 01 Nov 22 11:14 UTC | SHAKEN 807J | true | [view](CERTS/0d97796c10ef21ba02a0f7954d82ace5e41d7ea66e4ea9eb16801c3a1a1d9f03/README.md) |
| 12 Nov 22 04:05 UTC | SHAKEN 345J | true | [view](CERTS/ac38b9036e27ff2b753a0727e8ccdfd800d459dfdf595454c6a9dc7e9cd8ad74/README.md) |
| 12 Nov 22 04:54 UTC | SHAKEN 345J | true | [view](CERTS/0a4891c263ce12ebb5116b43048b2b427314bd6133848201f8186b62f7efc68a/README.md) |
| 14 Nov 22 14:10 UTC | SHAKEN 841J | true | [view](CERTS/51cf29551891eb1f1935e83cc86a50c52807f84c49b3b7c05461a2d89e4d8a75/README.md) |
| 14 Nov 22 20:24 UTC | SHAKEN 738J | true | [view](CERTS/aa1140caf05a17dfee3e6713e4e26a30cd29442e0dc0d5ffc29a2522cef3408a/README.md) |
| 16 Nov 22 18:14 UTC | SHAKEN 060K | true | [view](CERTS/6de8554dccdf250ab0e48a41aab008b35030816a6703b56815b5f7955959909a/README.md) |
| 16 Nov 22 18:38 UTC | SHAKEN 089K | true | [view](CERTS/ebd3acd3501eeb9d7f40b69df65fc52b9e3b568097b6fb45bc7495e6751bcaa9/README.md) |
| 16 Nov 22 20:02 UTC | SHAKEN 366G | true | [view](CERTS/e79aed8b72cca256fc18f44c5715b6a4f17145e3f50bf2f5305c9ecd1be3c168/README.md) |
| 16 Nov 22 20:19 UTC | SHAKEN 674J | true | [view](CERTS/b49b8f73bdbba3a16a67e541dd0c162a5247d409776f91855d2eb8cf7ed99e9f/README.md) |
| 16 Nov 22 20:20 UTC | SHAKEN 738J | true | [view](CERTS/b42e7ab46fb071123047e9bc872f760710c6185e3b74d2e969f70d29bba199c4/README.md) |
| 16 Nov 22 20:20 UTC | SHAKEN 700H | true | [view](CERTS/8178b89ab448b6f83ce420e547db2f49cfaadd93f61c86c23f87331f43c9456f/README.md) |
| 16 Nov 22 20:20 UTC | SHAKEN 733J | true | [view](CERTS/64f63e86573afe42f68bc4b9dddaf321ffee6a4239863cfa0efefe42aa2ec2ef/README.md) |
| 16 Nov 22 20:21 UTC | SHAKEN 819H | true | [view](CERTS/402e1f214004756fc9215e7c36b1452926575ba01cec3a14cc32397f583ff0b4/README.md) |
| 16 Nov 22 20:24 UTC | SHAKEN 469A | true | [view](CERTS/4d0e7f231fc63329b7c2b094d9cf3c2ddcda153d2e61eddca8bf8a79554de26b/README.md) |
| 16 Nov 22 20:24 UTC | SHAKEN 790J | true | [view](CERTS/ba89497bc114157e5b728ca175dd8dd388c2a4ef6e8543f7ee693747658697c2/README.md) |
| 16 Nov 22 20:25 UTC | SHAKEN 738J | true | [view](CERTS/51c22d8787b9404a94ecb566f5c62e8506cfbc5231a97da4e60bcb2ac0da1a0a/README.md) |
| 16 Nov 22 20:25 UTC | SHAKEN 459J | true | [view](CERTS/ba4b65ae4bf383df4832cf251a733dc341cf9c47d7c5eaa232a51aaf8f48bf9a/README.md) |
| 16 Nov 22 20:26 UTC | SHAKEN 366G | true | [view](CERTS/c633e6ebaa2ecb45d6b4d724d88a6de6d695706a83e31c05429484c3df57f0d3/README.md) |
| 16 Nov 22 20:26 UTC | SHAKEN 0226 | true | [view](CERTS/d20520e1bb7f2b7a2585d017ac75918130fda7e5881432d7a4c4c6e44408f718/README.md) |
| 16 Nov 22 20:26 UTC | SHAKEN 738J | true | [view](CERTS/4feb69bdbe9ef2b93c4652f1a803c27c4f252d539e4e3000eb6f0d1d9760fe6e/README.md) |
| 16 Nov 22 20:54 UTC | SHAKEN 518J | true | [view](CERTS/eca5ff6b372d6b84cdb861d2381170e29630dad6302d75a274d4843b54f22287/README.md) |
| 17 Nov 22 20:15 UTC | SHAKEN 983J | true | [view](CERTS/a9004f0431018736dc8756dbdf2228059db10dadf538356dd7ae83892a54365d/README.md) |
| 18 Nov 22 03:32 UTC | SHAKEN 0172 | true | [view](CERTS/7c2b110e349057c416c2f43fb77f672d34b773fd11ec020bf5862aa6921c76fd/README.md) |
| 18 Nov 22 14:35 UTC | SHAKEN 2550 | true | [view](CERTS/9c167daffac6d99a821bdf0aa3710d1fb5fa23f8307746358286a2c0a0ffae68/README.md) |
| 18 Nov 22 15:42 UTC | SHAKEN 722J | true | [view](CERTS/4111bb56ba547a2f493d1e9c2baefa33ce9bd117904133477d063d6fa0d6b9b2/README.md) |
| 18 Nov 22 18:04 UTC | SHAKEN 952J | true | [view](CERTS/b05bb158c054b6e579fe20750c83ba5524896545d629b96b51d4c8d9ab38dfd3/README.md) |
| 18 Nov 22 20:08 UTC | SHAKEN 864J | true | [view](CERTS/0f7f7697354353bd9d2ab9242343d42f4a12e66045709b0fb7e50469e313f8b3/README.md) |
| 18 Nov 22 20:13 UTC | SHAKEN 177K | true | [view](CERTS/e1e3958056b15dd9d82913101142469c028a264dae2033dfe50ab5bebc4514f2/README.md) |
| 19 Nov 22 18:14 UTC | SHAKEN 060K | true | [view](CERTS/93ebcbb2eaf5438d17ba174426b8753138757772644885642cfe5797bd32966a/README.md) |
| 19 Nov 22 18:38 UTC | SHAKEN 089K | true | [view](CERTS/ec707d95be99d5a7110a080bfd963df45407aba37c0e514c806fa9a59af8b5bd/README.md) |
| 19 Nov 22 19:57 UTC | SHAKEN 297K | true | [view](CERTS/22f9e18959d3e37a3e53dd31c0c8e95153ad52e480f2cb785085829b0ebef93f/README.md) |
| 19 Nov 22 20:19 UTC | SHAKEN 674J | true | [view](CERTS/612c4923305c08bcf0d2a8ec10c1828c1a9ebd96a75ed540f5e46ad2f70cfe1b/README.md) |
| 19 Nov 22 20:20 UTC | SHAKEN 738J | true | [view](CERTS/2372aac07c78cfbaf7c5d62d5c2ef1d700a953f26e3950ffb6f6f97f54601739/README.md) |
| 19 Nov 22 20:20 UTC | SHAKEN 700H | true | [view](CERTS/55f45d52865d21e186e1b32be6636444b944c0d69677b5bb5249ecaaebf6c582/README.md) |
| 19 Nov 22 20:22 UTC | SHAKEN 819H | true | [view](CERTS/e22ad1fbfb8d7dc6f452c347c559aaa6021f287c49e93bc4b288fe7a7c260239/README.md) |
| 19 Nov 22 20:23 UTC | SHAKEN 806J | true | [view](CERTS/6bf9a2f9487a04751a54131d47afea78ccbb65e5e95b00c13a24a68ee28521ed/README.md) |
| 19 Nov 22 20:24 UTC | SHAKEN 691J | true | [view](CERTS/3d3f042d846ae38a651bba04a841e214dee00c01afcb1627de0aac8778ba4635/README.md) |
| 19 Nov 22 20:24 UTC | SHAKEN 726J | true | [view](CERTS/a685af3578a1c94a3c5f85aca999c8197cd2113fec3f691137712e9ea0f3e96e/README.md) |
| 19 Nov 22 20:24 UTC | SHAKEN 495J | true | [view](CERTS/5a786330ca32d3e94d8f0333125cb42429676b07d04a72f356d12554381f8138/README.md) |
| 19 Nov 22 20:24 UTC | SHAKEN 790J | true | [view](CERTS/cde75e2a3f2720907cb77afd73879f37a085953d313299975ce56e12a6f18fd6/README.md) |
| 19 Nov 22 20:25 UTC | SHAKEN 738J | true | [view](CERTS/5bdb5d35c062ad58fa747e9a3a1bd2c3fe58366b9e509710b523db2720b53231/README.md) |
| 19 Nov 22 20:25 UTC | SHAKEN 459J | true | [view](CERTS/fc0eab6847d3c3f2783caca3fe90f96bf2165fbe9736753164ee2868dd1969d5/README.md) |
| 19 Nov 22 20:25 UTC | SHAKEN 672B | true | [view](CERTS/c3d60fb50497aeb57ccfc9c5844ac807506bd942734636e629ce31174381522f/README.md) |
| 19 Nov 22 20:26 UTC | SHAKEN 366G | true | [view](CERTS/6959e68ac60d19da9f2031cc0489cc5849c9d8d755a38ebe1a1cf66170ba2ff6/README.md) |
| 19 Nov 22 20:26 UTC | SHAKEN 0226 | true | [view](CERTS/70bb19ec974868f587136d40d3c5559b7ec67cbb4e39a17d6a306a9d4eb4ff86/README.md) |
| 19 Nov 22 20:26 UTC | SHAKEN 738J | true | [view](CERTS/f3719859a2b0aa88dc120730712688b7ec3287985b0e0b90e7d09c0e14b28624/README.md) |
| 19 Nov 22 20:58 UTC | SHAKEN 518J | true | [view](CERTS/5f54fea211d14038d589b536909589a813349e138d083f8b1c2a01274b8c6f8c/README.md) |
| 20 Nov 22 10:20 UTC | SHAKEN 551G | true | [view](CERTS/ffbf5e3bbd21c49711456c6a87c75d9ff73c1184b20b4fcaee97fe0668213a86/README.md) |
| 20 Nov 22 20:16 UTC | SHAKEN 983J | true | [view](CERTS/f8ecfa390e369dc4bbef9c63ef0bce41506eb8c4a078fa733cb037c32e1f7b0a/README.md) |
| 20 Nov 22 21:43 UTC | SHAKEN 606F | true | [view](CERTS/2741d28c11d00d8d599100c6d152146558102a819f47e84441f7b6f49d2d13b4/README.md) |
| 21 Nov 22 01:42 UTC | SHAKEN 278K | true | [view](CERTS/d504432099f8d63aaac06453f97dd3b27aafaf55c8b70e4c553959bc9b7fcdf7/README.md) |
| 21 Nov 22 13:25 UTC | SHAKEN 734J | true | [view](CERTS/0d9b79e63404f78d8ad52b427cd601f9e6cae2e1460b1515fd10119783d66c88/README.md) |
| 21 Nov 22 14:35 UTC | SHAKEN 2550 | true | [view](CERTS/1d51152d79e331af7ad3d277a44def14467f1c7d0e15539d157bafe49dbc500d/README.md) |
| 21 Nov 22 20:08 UTC | SHAKEN 864J | true | [view](CERTS/d66aa6cb0158655c064525336a4522b64a07365b7ea85945f6a0f5afd4c02bea/README.md) |
| 21 Nov 22 20:13 UTC | SHAKEN 177K | true | [view](CERTS/b489e40e3be0bc47a9bdb906ee1e183ac25162ec0f1fb04f1a1e3b55d7eef9c9/README.md) |
| 21 Nov 22 20:25 UTC | SHAKEN 738J | true | [view](CERTS/6b8346b320b4956f1b2726c057062818646170e7ae6f65614dc498cf80947979/README.md) |
| 22 Nov 22 08:25 UTC | SHAKEN 0172 | true | [view](CERTS/5d20fcdf67720cc55f4e300a8bd6b6a018a972d91b720b65f54d86f45438be60/README.md) |
| 22 Nov 22 14:44 UTC | SHAKEN 747J | true | [view](CERTS/a097de8ecc292748ee5cf1cbfef385d339de8d8bef975cd408ca96d9322b17b6/README.md) |
| 22 Nov 22 17:26 UTC | SHAKEN 107K | true | [view](CERTS/410b25f8f41389263f1ea83e4bf2d1c7e6e86c392617e8fef442899ef4ea6d93/README.md) |
| 22 Nov 22 20:23 UTC | SHAKEN 769J | true | [view](CERTS/d9094586312e2d292698be349176431cff52ee10f4137e44ec38d11f04b03d68/README.md) |
| 22 Nov 22 20:24 UTC | SHAKEN 849J | true | [view](CERTS/7f5f53b3f18ae533f64eb305561a8f471674b84111e5070a8843f8027beed0f5/README.md) |
| 22 Nov 22 20:24 UTC | SHAKEN 469A | true | [view](CERTS/e15e4277350f71a25bdd3359776d86442ce0a8baba6c4eb477125435ca8e0235/README.md) |
| 22 Nov 22 20:24 UTC | SHAKEN 495J | true | [view](CERTS/f33d3beceb14cf9a49b7e36aa2e17dd5f044bc1cc3653a58bd26ac5cacccab3a/README.md) |
| 22 Nov 22 20:25 UTC | SHAKEN 738J | true | [view](CERTS/43726b2e4bd4727221019b1e21dd7ab7b63f518b68c6df324cdf603aebbc6897/README.md) |
| 22 Nov 22 20:25 UTC | SHAKEN 459J | true | [view](CERTS/1a361d59ac849f7a4f4a81de0ce7a1639d76f1927627aa34a47c4facd30bb946/README.md) |
| 22 Nov 22 20:25 UTC | SHAKEN 672B | true | [view](CERTS/2045db35890ddee03a3c22ed000a5f2980cd1ac4e2a5c097b6dfa63887f3a03e/README.md) |
| 22 Nov 22 20:26 UTC | SHAKEN 366G | true | [view](CERTS/53643480e568b34f20877fc9318c12a85b381b6fdddeeaa96c1ffdf85e9c64f0/README.md) |
| 22 Nov 22 20:27 UTC | SHAKEN 0226 | true | [view](CERTS/da6fb8416e494f291048be17118cea94ca8ea31238e68eebd4d7886522c09aca/README.md) |
| 22 Nov 22 20:27 UTC | SHAKEN 738J | true | [view](CERTS/88c0d10b066f1887d3d2d08ed6a27524629dc6d5f1509f7686ad39c959eafdef/README.md) |
| 22 Nov 22 21:02 UTC | SHAKEN 518J | true | [view](CERTS/d82800c61eeef2a756c8da2df2fca2a7f46b4921177fc292a8fce996ec9c26b2/README.md) |

#### CA Certificates

| Created at | Subject | Problems | Link |
|------------|---------|----------|------|
| 20 Aug 21 00:00 UTC | TransNexus, Inc. SHAKEN Issuing CA3 | true | [view](CERTS/51b51431ffd8055cd7e30c213540f4868258058d4c4c9f23f1ed8bea75101e37/README.md) |
| 20 Aug 21 00:00 UTC | TransNexus, Inc. SHAKEN Issuing CA1 | true | [view](CERTS/037a7c6621814bd5f01fa2080a77a245458694d9d566e0050e068fc853317857/README.md) |
| 20 Aug 21 00:00 UTC | TransNexus, Inc. SHAKEN Root CA1 | true | [view](CERTS/59ab943f4e8c3c17755cc8e4abb5ac65736ce74ec587c62cb744ab0babbbe2fe/README.md) |
| 24 Oct 22 00:00 UTC | TransNexus, Inc. SHAKEN Issuing CA4 | false | [view](CERTS/80dd12d935f46c256c7c1289a0834b21ec8f2f8c35f2864928d1e75f3a280665/README.md) |
| 24 Oct 22 00:00 UTC | TransNexus, Inc. SHAKEN Root CA2 | false | [view](CERTS/a26e04fc786ab70b8085236b2c53f8cfbf5d0c6a5c2c9c3e9f91669fbb8ea4d5/README.md) |


Generated: 23 Nov 22 18:09 UTC