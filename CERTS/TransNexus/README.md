# STIR/SHAKEN CA Ecosystem Compliance

## TransNexus

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 2179 certificates were included in the corpus being tested
- 2 certificates in the corpus were skipped because they are duplicates
- 2082 certificates in the corpus were skipped because they are expired
- 1 certificates in the corpus were skipped because they are not currently trusted
- 94 certificates being tested against the remaining rules
- 1.19 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 8.51% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 69 days is the average remaining validity for the certificates in the corpus
- 66 days is the average initial validity for the certificates in the corpus
- 78 certificates expire in the next 30 days
- 1.42 average number of unexpired certificates per OCN observed
- 66 unique OCNs observed in unexpired and valid certificate corpus

| Instances | Test | Source |
|-----------|------|--------|
| 1 | [e_atis_certificate_policies](ISSUES/e_atis_certificate_policies/README.md) | ATIS1000080 |
| 8 | [e_incorrect_ku_encoding](ISSUES/e_incorrect_ku_encoding/README.md) | RFC5280 |
| 1 | [e_us_cp_ambiguous_identifier](ISSUES/e_us_cp_ambiguous_identifier/README.md) | US_SHAKEN_CP |
| 94 | [e_us_cp_subject_sn](ISSUES/e_us_cp_subject_sn/README.md) | US_SHAKEN_CP |
| 8 | [w_shaken_subject_rdn_unknown](ISSUES/w_shaken_subject_rdn_unknown/README.md) | SHAKEN_PKI_BEST_PRACTICES |

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
- 5450 days is the average remaining validity for the certificates in the corpus
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
| 27&#160;Apr&#160;22&#160;18:22&#160;UTC | CCI SHAKEN 663J | 27&#160;Apr&#160;23&#160;18:22&#160;UTC | true | [view](CERTS/9cfd0b9974acdc9b5fce3888cf2f6cf99173e7955adeebd37a1e60dfc444a0b3/README.md) |
| 10&#160;Jun&#160;22&#160;14:00&#160;UTC | SHAKEN 597J | 10&#160;Jun&#160;23&#160;14:00&#160;UTC | true | [view](CERTS/4baa838eb96ce358aa1565d0add8b75072258efd5d21a5bbb0eec878026a67be/README.md) |
| 20&#160;Jun&#160;22&#160;20:25&#160;UTC | SHAKEN 857J | 20&#160;Jun&#160;23&#160;20:25&#160;UTC | true | [view](CERTS/e1e0a4851fb7b51d617959c00fb6a17e7b23d34e1555ba70ab01c9182a2f85f8/README.md) |
| 22&#160;Jun&#160;22&#160;18:46&#160;UTC | SHAKEN 755J | 22&#160;Jun&#160;23&#160;18:46&#160;UTC | true | [view](CERTS/a8cac316fddf25b48b9350a325791be8b87e6484a5959ab605db1ae5e3e99a84/README.md) |
| 29&#160;Jun&#160;22&#160;20:24&#160;UTC | SHAKEN 736J | 29&#160;Jun&#160;23&#160;20:24&#160;UTC | true | [view](CERTS/71b25a480077816998e48e6a76f272e70aa91ecb19b0e773cd8777823cdead32/README.md) |
| 25&#160;Jul&#160;22&#160;18:36&#160;UTC | SHAKEN 578J | 25&#160;Jul&#160;23&#160;18:36&#160;UTC | true | [view](CERTS/f4af5a28131d680ef2682ae86181b64f0c919864d7dbaef00c0598d314b537a1/README.md) |
| 10&#160;Aug&#160;22&#160;18:11&#160;UTC | SHAKEN 073H | 10&#160;Aug&#160;23&#160;18:11&#160;UTC | true | [view](CERTS/87dabb91a689a6ac540b5ffbf5eb89326fada83fb822d2bb9ec31b68251b579e/README.md) |
| 13&#160;Sep&#160;22&#160;17:24&#160;UTC | SHAKEN 706J | 13&#160;Sep&#160;23&#160;17:24&#160;UTC | true | [view](CERTS/1eea5e17ae8b36cb5b6d657178273af745cea015e63586e6d44d1ba2ee8f722a/README.md) |
| 23&#160;Nov&#160;22&#160;17:31&#160;UTC | SHAKEN 4036 | 10&#160;May&#160;23&#160;17:31&#160;UTC | true | [view](CERTS/e03ff4259e6b456f1bb19d011f005932e83b0a35b520730c97cd77ad8868071c/README.md) |
| 07&#160;Dec&#160;22&#160;00:45&#160;UTC | SHAKEN 299K | 07&#160;Dec&#160;23&#160;00:45&#160;UTC | true | [view](CERTS/7b5e04ff01ace2ee384df3507886684204901bf23851866eb6b3891ecdd96a92/README.md) |
| 31&#160;Dec&#160;22&#160;03:33&#160;UTC | SHAKEN 8526 | 31&#160;Dec&#160;23&#160;03:33&#160;UTC | true | [view](CERTS/96672cb8cc7c71e38d9e6350eac6c4f5c430c1eefeabe00cc095527e20a20c9b/README.md) |
| 03&#160;Jan&#160;23&#160;16:07&#160;UTC | SHAKEN 6628 | 03&#160;Jun&#160;23&#160;16:07&#160;UTC | true | [view](CERTS/dec9a2b7e2fce8ca94e2b1313886772d8176d6adcac7c4cb2b295315fa79f5ab/README.md) |
| 12&#160;Jan&#160;23&#160;05:08&#160;UTC | SHAKEN 621J | 12&#160;May&#160;23&#160;05:08&#160;UTC | true | [view](CERTS/017fbb31b0e0c0f3dbb7de67e5e3f95c4ce10e546aaaca32faa97fb61eafcfe6/README.md) |
| 16&#160;Jan&#160;23&#160;17:54&#160;UTC | SHAKEN 622J | 15&#160;Jul&#160;23&#160;17:54&#160;UTC | true | [view](CERTS/4c659921cd1ec39c239ab31d4c1ca813017c01d51512ea11a1fe0572f034c534/README.md) |
| 25&#160;Jan&#160;23&#160;14:28&#160;UTC | SHAKEN 873J | 25&#160;Jan&#160;24&#160;14:28&#160;UTC | true | [view](CERTS/36e9ba32d337ad13909bee0b62125c9769a049c4ae10d749687d75af4e49387d/README.md) |
| 27&#160;Feb&#160;23&#160;01:43&#160;UTC | SHAKEN 807J | 28&#160;Apr&#160;23&#160;01:43&#160;UTC | true | [view](CERTS/ba83be2505aeab5690525902158ec1372b1b8917d4db12e9e43dceee579dadfb/README.md) |
| 09&#160;Mar&#160;23&#160;18:46&#160;UTC | SHAKEN 193E | 08&#160;May&#160;23&#160;18:46&#160;UTC | true | [view](CERTS/34187f020bbbd3596eea83f5f2243b84920396d2c46eed5b8078a0b696c383e5/README.md) |
| 14&#160;Mar&#160;23&#160;16:44&#160;UTC | SHAKEN 551G | 13&#160;Apr&#160;23&#160;16:44&#160;UTC | true | [view](CERTS/59dedf353c817690bfe09fbbdb984cbecf13fd9193ea8c309d750e9113cba84b/README.md) |
| 16&#160;Mar&#160;23&#160;14:05&#160;UTC | SHAKEN 366G | 15&#160;Apr&#160;23&#160;14:05&#160;UTC | true | [view](CERTS/5771548185fd9014849b339406baa6c2460446147db76470c8bbfc338cdb42bc/README.md) |
| 17&#160;Mar&#160;23&#160;23:13&#160;UTC | SHAKEN 4632 | 16&#160;Apr&#160;23&#160;23:13&#160;UTC | true | [view](CERTS/666fd44a3ad17a9dd4a3c65051224143d82cc1a192655872cd4f077f05551ed4/README.md) |
| 18&#160;Mar&#160;23&#160;04:40&#160;UTC | SHAKEN 345J | 17&#160;Apr&#160;23&#160;04:40&#160;UTC | true | [view](CERTS/773eee310dfa2459d96dbadc7e6463b019e2d090b993c9cca043ca2192b158a4/README.md) |
| 18&#160;Mar&#160;23&#160;05:11&#160;UTC | SHAKEN 345J | 17&#160;Apr&#160;23&#160;05:11&#160;UTC | true | [view](CERTS/f647d623a58e4ae8c6c6ca1f78c883810c33105ce8023e36af56822cd8ffd687/README.md) |
| 18&#160;Mar&#160;23&#160;16:04&#160;UTC | SHAKEN 722J | 17&#160;Apr&#160;23&#160;16:04&#160;UTC | true | [view](CERTS/fbb1f2ea714e4b4373ec21e8011fd71a6716c49fbce8ba7d4654147ad66fafd9/README.md) |
| 20&#160;Mar&#160;23&#160;18:53&#160;UTC | SHAKEN 505J | 18&#160;Jun&#160;23&#160;18:53&#160;UTC | true | [view](CERTS/9a4feceff25f99e7d04d1542ef56bba8f060987c5b6734239f185997c5e58cf1/README.md) |
| 27&#160;Mar&#160;23&#160;17:18&#160;UTC | SHAKEN 2720 | 26&#160;Mar&#160;24&#160;17:18&#160;UTC | true | [view](CERTS/7282c0683ebff27881b1ccd1f66664c3102340ff657d9363fcecb9c1b7924c12/README.md) |
| 27&#160;Mar&#160;23&#160;19:12&#160;UTC | SHAKEN 577F | 26&#160;Apr&#160;23&#160;19:12&#160;UTC | true | [view](CERTS/717010b11b69f8050e3a16eb9ced59c5abf1734320d54959ac6d51e3e92846b9/README.md) |
| 30&#160;Mar&#160;23&#160;18:51&#160;UTC | SHAKEN 952J | 29&#160;Apr&#160;23&#160;18:51&#160;UTC | true | [view](CERTS/e21d4b5c92be0e6db90e3a1a986ca049c8edd8c4bad4290f0ca6f4ae42adc537/README.md) |
| 04&#160;Apr&#160;23&#160;16:51&#160;UTC | SHAKEN 551G | 04&#160;May&#160;23&#160;16:50&#160;UTC | true | [view](CERTS/5123718621bfacb494869d8ac857c3387021c70d5e838ee00fadbeaa399bdec2/README.md) |
| 04&#160;Apr&#160;23&#160;23:14&#160;UTC | SHAKEN 4632 | 04&#160;May&#160;23&#160;23:14&#160;UTC | true | [view](CERTS/a89b5761f81a706cc098bc2c966e51a203c45a76bf423d09f379210a17aeae44/README.md) |
| 06&#160;Apr&#160;23&#160;10:33&#160;UTC | SHAKEN 841J | 20&#160;Apr&#160;23&#160;10:33&#160;UTC | true | [view](CERTS/28d8877e8df6975fbbcfb5b371fd03352e4bc7163be96f34dc40f30c4b4cfd79/README.md) |
| 06&#160;Apr&#160;23&#160;14:53&#160;UTC | SHAKEN 012K | 13&#160;Apr&#160;23&#160;14:53&#160;UTC | true | [view](CERTS/2d596961e75afe90da02ec8638347d42b7de8d2eae30498d58317e1329ffd66a/README.md) |
| 06&#160;Apr&#160;23&#160;17:37&#160;UTC | SHAKEN 107K | 13&#160;Apr&#160;23&#160;17:37&#160;UTC | true | [view](CERTS/718cfd989595bc5fce3b4d719ee577d0c072bb1e162156f76cec857bc5b7d7e7/README.md) |
| 06&#160;Apr&#160;23&#160;20:08&#160;UTC | SHAKEN 297K | 13&#160;Apr&#160;23&#160;20:08&#160;UTC | true | [view](CERTS/26fce18bdb5f54ebb90a0f74327c7eaaca2f187840de2181674a6cd2968f3551/README.md) |
| 06&#160;Apr&#160;23&#160;20:31&#160;UTC | SHAKEN 674J | 13&#160;Apr&#160;23&#160;20:31&#160;UTC | true | [view](CERTS/2b41b597564cd95c3c21951fc9dd3a82ad3eb12c834598c25d3b3427e7caeef8/README.md) |
| 06&#160;Apr&#160;23&#160;20:32&#160;UTC | SHAKEN 735J | 13&#160;Apr&#160;23&#160;20:32&#160;UTC | true | [view](CERTS/66491b91e0c7ae4650aaf03f5704da83e59659bcd8d0fd25d77a310a0c91e430/README.md) |
| 06&#160;Apr&#160;23&#160;20:33&#160;UTC | SHAKEN 738J | 13&#160;Apr&#160;23&#160;20:33&#160;UTC | true | [view](CERTS/e59d08bb07d2fb8afca78770e6a458d3bb1e412cf69942171382040eb0c15338/README.md) |
| 06&#160;Apr&#160;23&#160;20:34&#160;UTC | SHAKEN 700H | 13&#160;Apr&#160;23&#160;20:34&#160;UTC | true | [view](CERTS/a98d0c7611842415c42936f1453e72d7cbf51bef54b86bf3f0e46a4fb7b4188b/README.md) |
| 06&#160;Apr&#160;23&#160;20:34&#160;UTC | SHAKEN 856H | 13&#160;Apr&#160;23&#160;20:34&#160;UTC | true | [view](CERTS/96d923d1f6f5076926d7a9e2a55537cd3e3ac71fbd0f8c83cea6c27e41f42de8/README.md) |
| 06&#160;Apr&#160;23&#160;20:35&#160;UTC | SHAKEN 819H | 13&#160;Apr&#160;23&#160;20:35&#160;UTC | true | [view](CERTS/0aefd09afa8df56a67f7d756936000b681e8e63b87c789902da8684f4da2daee/README.md) |
| 06&#160;Apr&#160;23&#160;20:36&#160;UTC | SHAKEN 053G | 13&#160;Apr&#160;23&#160;20:36&#160;UTC | true | [view](CERTS/51c9e2becafd055679a19c29cf6c1e40334642df3f4fd791ab58ff7077c08f56/README.md) |
| 06&#160;Apr&#160;23&#160;20:38&#160;UTC | SHAKEN 469A | 13&#160;Apr&#160;23&#160;20:38&#160;UTC | true | [view](CERTS/9e35c81f897bd2eea7ad15313e1877ca0962baa7fd6958a2e91965c0a1e30ee2/README.md) |
| 06&#160;Apr&#160;23&#160;20:39&#160;UTC | SHAKEN 721J | 13&#160;Apr&#160;23&#160;20:39&#160;UTC | true | [view](CERTS/44d6c3642e1565006a238f44eb96ad2894708d9089776d4f6293253fca2ccd3c/README.md) |
| 06&#160;Apr&#160;23&#160;20:40&#160;UTC | SHAKEN 625J | 13&#160;Apr&#160;23&#160;20:40&#160;UTC | true | [view](CERTS/c5f2a33149e4a730527dbb082cd8365a3d79f8786687a70b32700ce43e0cfa8d/README.md) |
| 06&#160;Apr&#160;23&#160;20:40&#160;UTC | SHAKEN 738J | 13&#160;Apr&#160;23&#160;20:40&#160;UTC | true | [view](CERTS/b329beacd3857aa71dcd3cc9a1ee155437c796c5080f13017a2202f959f4efa3/README.md) |
| 06&#160;Apr&#160;23&#160;20:40&#160;UTC | SHAKEN 672B | 13&#160;Apr&#160;23&#160;20:40&#160;UTC | true | [view](CERTS/5efe2f5ec66f08180476321b1f36daf3f9bd0f971cdfe8beb6929484183ca70b/README.md) |
| 06&#160;Apr&#160;23&#160;20:41&#160;UTC | SHAKEN 366G | 13&#160;Apr&#160;23&#160;20:41&#160;UTC | true | [view](CERTS/ce0a0d2639013799e34de4a7f462a9c32e019f266bfb110f4c9ab97da6cb1564/README.md) |
| 06&#160;Apr&#160;23&#160;20:42&#160;UTC | SHAKEN 738J | 13&#160;Apr&#160;23&#160;20:42&#160;UTC | true | [view](CERTS/f9932ef43d0672dd3d0f0fa44042301ea895a12536b23dcde45e80224f4083d7/README.md) |
| 07&#160;Apr&#160;23&#160;20:29&#160;UTC | SHAKEN 983J | 14&#160;Apr&#160;23&#160;20:29&#160;UTC | true | [view](CERTS/11e5b585deadc506a73090b8ebb7614dcabb8b7603505ba3a0d5b8be31705c9d/README.md) |
| 08&#160;Apr&#160;23&#160;02:54&#160;UTC | SHAKEN 625J | 15&#160;Apr&#160;23&#160;02:54&#160;UTC | true | [view](CERTS/ee744301529a92e8e21ad985357047146fcebe9f9a92d76a7988c57349256cab/README.md) |
| 08&#160;Apr&#160;23&#160;04:42&#160;UTC | SHAKEN 345J | 08&#160;May&#160;23&#160;04:42&#160;UTC | true | [view](CERTS/dd31135745befe8cbca897037627e600c7991cb797c5b6244a61783993e2f2cd/README.md) |
| 08&#160;Apr&#160;23&#160;05:13&#160;UTC | SHAKEN 345J | 08&#160;May&#160;23&#160;05:13&#160;UTC | true | [view](CERTS/ed628fcdb76ed3264990a17371588f8b30d06f282a040b2d612a3bcb7aac835b/README.md) |
| 08&#160;Apr&#160;23&#160;05:54&#160;UTC | SHAKEN 982J | 15&#160;Apr&#160;23&#160;05:54&#160;UTC | true | [view](CERTS/bff1c69df4e7e0406a9808f4068c325becf91b931a07fdc881b39a45d5a2e66c/README.md) |
| 08&#160;Apr&#160;23&#160;14:47&#160;UTC | SHAKEN 2550 | 15&#160;Apr&#160;23&#160;14:47&#160;UTC | true | [view](CERTS/e091c7b45c9f56b9341b698299e68e5533d4d7333ce601ed8674505307837e87/README.md) |
| 08&#160;Apr&#160;23&#160;16:15&#160;UTC | SHAKEN 2277 | 15&#160;Apr&#160;23&#160;16:15&#160;UTC | true | [view](CERTS/642d7173a81e6698be6201942b5e4dd024f5ac810bfbf09f1b71de0cd604573c/README.md) |
| 08&#160;Apr&#160;23&#160;17:39&#160;UTC | SHAKEN 833J | 15&#160;Apr&#160;23&#160;17:39&#160;UTC | true | [view](CERTS/49323e96991da611147edda6de769263afeddfd4e53590bc9941691420503fbf/README.md) |
| 08&#160;Apr&#160;23&#160;18:53&#160;UTC | SHAKEN 193E | 07&#160;Jun&#160;23&#160;18:53&#160;UTC | true | [view](CERTS/8cc7b08b6c144bae17ff71317ebea1877859c2bbdd387c6b9083a130bf7ab18c/README.md) |
| 08&#160;Apr&#160;23&#160;20:25&#160;UTC | SHAKEN 177K | 15&#160;Apr&#160;23&#160;20:25&#160;UTC | true | [view](CERTS/5544b7b9be7d8c989dff3eab17552808bd77062d3862653969ad2dd2332a7fef/README.md) |
| 09&#160;Apr&#160;23&#160;14:53&#160;UTC | SHAKEN 012K | 16&#160;Apr&#160;23&#160;14:53&#160;UTC | true | [view](CERTS/653fdb6373317547170d852056745cc9ecc85b446209a795e2e4786320142493/README.md) |
| 09&#160;Apr&#160;23&#160;14:55&#160;UTC | SHAKEN 056J | 16&#160;Apr&#160;23&#160;14:55&#160;UTC | true | [view](CERTS/30aa90471c874a36d04685e0807e54b4f0e60b6359618ab86bc2c905743ac7ee/README.md) |
| 09&#160;Apr&#160;23&#160;17:38&#160;UTC | SHAKEN 107K | 16&#160;Apr&#160;23&#160;17:38&#160;UTC | true | [view](CERTS/6f6f1cfaebd349c108bb56c07b991d3d4a1ab5b46918b4cff0379c643e64d231/README.md) |
| 09&#160;Apr&#160;23&#160;18:27&#160;UTC | SHAKEN 060K | 16&#160;Apr&#160;23&#160;18:27&#160;UTC | true | [view](CERTS/c652e9aafff311a0d07a567e7e21d101e1e555aa00ee056bd23868bc99d244ef/README.md) |
| 09&#160;Apr&#160;23&#160;18:50&#160;UTC | SHAKEN 089K | 16&#160;Apr&#160;23&#160;18:50&#160;UTC | true | [view](CERTS/45803eef2808bd5ed1d8499d774bfde2377ff24684275caf2f623478d2bd27f9/README.md) |
| 09&#160;Apr&#160;23&#160;18:50&#160;UTC | SHAKEN 056K | 16&#160;Apr&#160;23&#160;18:50&#160;UTC | true | [view](CERTS/091f81fea35e4bbf6927401dfa3a33b00080d3dadc489247f3e145a807efce63/README.md) |
| 09&#160;Apr&#160;23&#160;18:52&#160;UTC | SHAKEN 9714 | 16&#160;Apr&#160;23&#160;18:52&#160;UTC | true | [view](CERTS/44914ef7fcd3ee3879372b82ceaf625a29ee91bf7677f2b319e0d7873751f11b/README.md) |
| 09&#160;Apr&#160;23&#160;20:08&#160;UTC | SHAKEN 297K | 16&#160;Apr&#160;23&#160;20:08&#160;UTC | true | [view](CERTS/f67d78dd4671bbc69f4ea3ddd81cfb22ac4a86eb418d2a104d18d186e689783e/README.md) |
| 09&#160;Apr&#160;23&#160;20:31&#160;UTC | SHAKEN 674J | 16&#160;Apr&#160;23&#160;20:31&#160;UTC | true | [view](CERTS/639c5d9d8fc28f74a8751338638c895fd91a85d8c950873360976512caee61c6/README.md) |
| 09&#160;Apr&#160;23&#160;20:32&#160;UTC | SHAKEN 735J | 16&#160;Apr&#160;23&#160;20:32&#160;UTC | true | [view](CERTS/aa89af9a688f66d26ff0db8540c9bcec9efef935c5de5c4504a77bddd45e8d08/README.md) |
| 09&#160;Apr&#160;23&#160;20:33&#160;UTC | SHAKEN 738J | 16&#160;Apr&#160;23&#160;20:33&#160;UTC | true | [view](CERTS/39a29ef23ebdfacd8fef66d24938b50637b9e3958c96b4e8fe4cc3732f6ff2fa/README.md) |
| 09&#160;Apr&#160;23&#160;20:34&#160;UTC | SHAKEN 700H | 16&#160;Apr&#160;23&#160;20:34&#160;UTC | true | [view](CERTS/acec0d98e9a102b46c78872894c533c59cf3a2d6f39d8ac398882404becf5122/README.md) |
| 09&#160;Apr&#160;23&#160;20:34&#160;UTC | SHAKEN 733J | 16&#160;Apr&#160;23&#160;20:34&#160;UTC | true | [view](CERTS/73286106ca5ef5102550b3e1563c1e01f6764c3152bd89c3d32a92d1421d70f2/README.md) |
| 09&#160;Apr&#160;23&#160;20:34&#160;UTC | SHAKEN 856H | 16&#160;Apr&#160;23&#160;20:34&#160;UTC | true | [view](CERTS/e30a884ff213d929bc50a73ca78a4ddf5082ebed35a1a2710196944189b4507f/README.md) |
| 09&#160;Apr&#160;23&#160;20:35&#160;UTC | SHAKEN 819H | 16&#160;Apr&#160;23&#160;20:35&#160;UTC | true | [view](CERTS/0f1a54f72fe82e0222a7fc4d66b639d60742e9f813fb07827c855e7cb9f55608/README.md) |
| 09&#160;Apr&#160;23&#160;20:36&#160;UTC | SHAKEN 769J | 16&#160;Apr&#160;23&#160;20:36&#160;UTC | true | [view](CERTS/d3ae1bb20f8ddbfa38835c13f217e90e6df597375903a359451628805708eef5/README.md) |
| 09&#160;Apr&#160;23&#160;20:39&#160;UTC | SHAKEN 469A | 16&#160;Apr&#160;23&#160;20:39&#160;UTC | true | [view](CERTS/45ca2ae627195d0de03d5ea9100bf079aa4c77f80a803d463ebb91cdd6a7f82f/README.md) |
| 09&#160;Apr&#160;23&#160;20:39&#160;UTC | SHAKEN 726J | 16&#160;Apr&#160;23&#160;20:39&#160;UTC | true | [view](CERTS/c3cd7cad887589b2528b351d570ca4274df9bebc9bd6596458bd91c9752fd17a/README.md) |
| 09&#160;Apr&#160;23&#160;20:39&#160;UTC | SHAKEN 790J | 16&#160;Apr&#160;23&#160;20:39&#160;UTC | true | [view](CERTS/64451bf1d1b0e517dfa5987f73dfe188d7793c23337a1e456e93e5e3dd3cfdaa/README.md) |
| 09&#160;Apr&#160;23&#160;20:39&#160;UTC | SHAKEN 721J | 16&#160;Apr&#160;23&#160;20:39&#160;UTC | true | [view](CERTS/946435a28fae93ea5b9864ca0422690d1a3c6a13afe6cdebf33835f1dcf124f3/README.md) |
| 09&#160;Apr&#160;23&#160;20:40&#160;UTC | SHAKEN 738J | 16&#160;Apr&#160;23&#160;20:40&#160;UTC | true | [view](CERTS/b35235e21fb104e0431117b8b044e3dd73ecde031a0ee2d3af215d08b6652a0f/README.md) |
| 09&#160;Apr&#160;23&#160;20:40&#160;UTC | SHAKEN 459J | 16&#160;Apr&#160;23&#160;20:40&#160;UTC | true | [view](CERTS/b3c43bde10ed51eb7b3e0a1e0bbc67cfe28cebf8c07f3839c917c77d32eb6bef/README.md) |
| 09&#160;Apr&#160;23&#160;20:41&#160;UTC | SHAKEN 366G | 16&#160;Apr&#160;23&#160;20:41&#160;UTC | true | [view](CERTS/dd71ac4388a7dd5f3f120f7c243cd6a03a5254fa5306f27b19989aafa131dc53/README.md) |
| 09&#160;Apr&#160;23&#160;20:41&#160;UTC | SHAKEN 4156 | 16&#160;Apr&#160;23&#160;20:41&#160;UTC | true | [view](CERTS/16e778ac5dca191d7eb67db68c64aec7141c7ef0c47b34957635eb74fbce9fba/README.md) |
| 09&#160;Apr&#160;23&#160;20:41&#160;UTC | SHAKEN 759J | 16&#160;Apr&#160;23&#160;20:41&#160;UTC | true | [view](CERTS/4ff5858950c7dab2cac329b55a503ac7001669bfdf8c41647c2d947d35e44d58/README.md) |
| 09&#160;Apr&#160;23&#160;20:41&#160;UTC | SHAKEN 0226 | 16&#160;Apr&#160;23&#160;20:41&#160;UTC | true | [view](CERTS/b9e15500d1dff4e4b853c0237179faa5310c776edf87e7fb31bce90703aadea4/README.md) |
| 09&#160;Apr&#160;23&#160;20:42&#160;UTC | SHAKEN 738J | 16&#160;Apr&#160;23&#160;20:42&#160;UTC | true | [view](CERTS/77256243e998b878be95ecad7d1ea82331ad04bf9bab3600e6b22dc8b1749fc9/README.md) |
| 10&#160;Apr&#160;23&#160;00:24&#160;UTC | SHAKEN 0172 | 17&#160;Apr&#160;23&#160;00:24&#160;UTC | true | [view](CERTS/0b276f8ba713126c7008ec3f823f77e821c5547db1e97cdded6a18f06713a5c7/README.md) |
| 10&#160;Apr&#160;23&#160;12:54&#160;UTC | SHAKEN 815G | 09&#160;Jul&#160;23&#160;12:54&#160;UTC | true | [view](CERTS/71206dbb241c8dbf9de187fca7b7e515fff61fbedcd1baacdc4a6a399ca2f887/README.md) |
| 10&#160;Apr&#160;23&#160;19:19&#160;UTC | SHAKEN 749J | 17&#160;Apr&#160;23&#160;19:19&#160;UTC | true | [view](CERTS/8f8cea0b0f0dfb57bcaeaaed229857c0fc2901141a08b20713e025199f777ec6/README.md) |
| 10&#160;Apr&#160;23&#160;20:29&#160;UTC | SHAKEN 983J | 17&#160;Apr&#160;23&#160;20:29&#160;UTC | true | [view](CERTS/9c18db314341badc44c526f940e06cde67fb93b51acd4ad3864dfc0f0cc49ac3/README.md) |
| 10&#160;Apr&#160;23&#160;20:44&#160;UTC | SHAKEN 738J | 19&#160;Apr&#160;23&#160;20:44&#160;UTC | true | [view](CERTS/5f8fd36c961c6420918943cf164899ebd1e5ef25b3975ce13c70845662b498e0/README.md) |
| 10&#160;Apr&#160;23&#160;21:55&#160;UTC | SHAKEN 606F | 17&#160;Apr&#160;23&#160;21:55&#160;UTC | true | [view](CERTS/1d6eab90cc8318f39d7c4dc544e8e89969bdacfc0af21ec323f0b313c5249b02/README.md) |
| 11&#160;Apr&#160;23&#160;01:52&#160;UTC | SHAKEN 278K | 18&#160;Apr&#160;23&#160;01:52&#160;UTC | true | [view](CERTS/74bfd39233ecf8f526dc350e42886d434880bcdb97061b1e6189d9884e2e6c38/README.md) |
| 11&#160;Apr&#160;23&#160;13:36&#160;UTC | SHAKEN 734J | 18&#160;Apr&#160;23&#160;13:36&#160;UTC | true | [view](CERTS/9309aafdb557ee8654be3c3c313b94191b56970ab424d7b025128a07750bf4cf/README.md) |
| 11&#160;Apr&#160;23&#160;14:47&#160;UTC | SHAKEN 2550 | 18&#160;Apr&#160;23&#160;14:47&#160;UTC | true | [view](CERTS/a14e3d2db2bd7a229dab06cf7722041696bdf2765d689d15a82e5d792b00cd40/README.md) |
| 11&#160;Apr&#160;23&#160;16:12&#160;UTC | SHAKEN 722J | 11&#160;May&#160;23&#160;16:12&#160;UTC | true | [view](CERTS/e9b82c67177482f81b00257a1bb3e819f7e5b91b9f3a5e0780411449a321ffb2/README.md) |

#### CA Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 20&#160;Aug&#160;21&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Issuing CA3 | 19&#160;Aug&#160;31&#160;23:59&#160;UTC | true | [view](CERTS/51b51431ffd8055cd7e30c213540f4868258058d4c4c9f23f1ed8bea75101e37/README.md) |
| 20&#160;Aug&#160;21&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Issuing CA1 | 19&#160;Aug&#160;31&#160;23:59&#160;UTC | true | [view](CERTS/037a7c6621814bd5f01fa2080a77a245458694d9d566e0050e068fc853317857/README.md) |
| 20&#160;Aug&#160;21&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Root CA1 | 19&#160;Aug&#160;41&#160;23:59&#160;UTC | true | [view](CERTS/59ab943f4e8c3c17755cc8e4abb5ac65736ce74ec587c62cb744ab0babbbe2fe/README.md) |
| 24&#160;Oct&#160;22&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Issuing CA4 | 23&#160;Oct&#160;32&#160;23:59&#160;UTC | false | [view](CERTS/80dd12d935f46c256c7c1289a0834b21ec8f2f8c35f2864928d1e75f3a280665/README.md) |
| 24&#160;Oct&#160;22&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Root CA2 | 23&#160;Oct&#160;42&#160;23:59&#160;UTC | false | [view](CERTS/a26e04fc786ab70b8085236b2c53f8cfbf5d0c6a5c2c9c3e9f91669fbb8ea4d5/README.md) |


Generated: 12 Apr 23 22:02 UTC