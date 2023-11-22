# STIR/SHAKEN CA Ecosystem Compliance

## TransNexus

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 5025 certificates were included in the corpus being tested
- 3 certificates in the corpus were skipped because they are duplicates
- 4501 certificates in the corpus were skipped because they are expired
- 458 certificates in the corpus were skipped because they are not currently trusted
- 63 certificates being tested against the remaining rules
- 1.00 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 0.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 4.76% of certificates are too old to be assessed against currently enforced expectations
- 100 days is the average remaining validity for the certificates in the corpus
- 98 days is the average initial validity for the certificates in the corpus
- 48 certificates expire in the next 30 days
- 1.21 average number of unexpired certificates per OCN observed
- 52 unique OCNs observed in unexpired and valid certificate corpus

| Instances | Test | Source |
|-----------|------|--------|
| 63 | [e_us_cp_subject_sn_shall](ISSUES/e_us_cp_subject_sn_shall/README.md) | US_SHAKEN_CP |

#### CA Certificates

- 5 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 0 certificates in the corpus were skipped because they are expired
- 3 certificates in the corpus were skipped because they are not currently trusted
- 2 certificates being tested against the remaining rules
- 0.00 issues on average found in unexpired, trusted, and non-compliant certificates
- 0.00% of certificates contain one or more Error level issue
- 0.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 100.00% of certificates are too old to be assessed against currently enforced expectations
- 5956 days is the average remaining validity for the certificates in the corpus
- 5479 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days

No error, warning, or notice level issues were found

### Details

#### Leaf Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 07&#160;Dec&#160;22&#160;00:45&#160;UTC | SHAKEN 299K | 07&#160;Dec&#160;23&#160;00:45&#160;UTC | true | [view](CERTS/7b5e04ff01ace2ee384df3507886684204901bf23851866eb6b3891ecdd96a92/README.md) |
| 31&#160;Dec&#160;22&#160;03:33&#160;UTC | SHAKEN 8526 | 31&#160;Dec&#160;23&#160;03:33&#160;UTC | true | [view](CERTS/96672cb8cc7c71e38d9e6350eac6c4f5c430c1eefeabe00cc095527e20a20c9b/README.md) |
| 25&#160;Jan&#160;23&#160;14:28&#160;UTC | SHAKEN 873J | 25&#160;Jan&#160;24&#160;14:28&#160;UTC | true | [view](CERTS/36e9ba32d337ad13909bee0b62125c9769a049c4ae10d749687d75af4e49387d/README.md) |
| 27&#160;Mar&#160;23&#160;17:18&#160;UTC | SHAKEN 2720 | 26&#160;Mar&#160;24&#160;17:18&#160;UTC | true | [view](CERTS/7282c0683ebff27881b1ccd1f66664c3102340ff657d9363fcecb9c1b7924c12/README.md) |
| 28&#160;Apr&#160;23&#160;17:50&#160;UTC | SHAKEN 663J | 27&#160;Apr&#160;24&#160;17:50&#160;UTC | true | [view](CERTS/bd944d8e1acfeaaa520a8c87826b19a8509b52b083525a940e9ee4db8af1a99b/README.md) |
| 24&#160;May&#160;23&#160;08:31&#160;UTC | SHAKEN 736J | 23&#160;May&#160;24&#160;08:31&#160;UTC | true | [view](CERTS/d273815c57e441d2066144a02d029b73c3d839f5d31cbaeb72a7b6bded62da53/README.md) |
| 06&#160;Jun&#160;23&#160;13:41&#160;UTC | SHAKEN 6628 | 03&#160;Dec&#160;23&#160;13:41&#160;UTC | true | [view](CERTS/1892b1588772ce058cea6d06c2e82ea66fae4902994a60ab326d96de2635e22d/README.md) |
| 06&#160;Jun&#160;23&#160;19:10&#160;UTC | SHAKEN 597J | 05&#160;Jun&#160;24&#160;19:10&#160;UTC | true | [view](CERTS/7f15a14bb0b8ad72d450e673344a747703cdbe15741d2696c371c83c5ca8e148/README.md) |
| 16&#160;Jun&#160;23&#160;18:51&#160;UTC | SHAKEN 857J | 15&#160;Jun&#160;24&#160;18:51&#160;UTC | true | [view](CERTS/97834968950fa42accd7ab2d9f43d12eb3dd919b5d5a72a153275419104aa534/README.md) |
| 08&#160;Jul&#160;23&#160;21:54&#160;UTC | SHAKEN 578J | 07&#160;Jul&#160;24&#160;21:54&#160;UTC | true | [view](CERTS/5f864f10b33ee5194c0f67cd87af157b58755671a4534b5bd9570ca5eacd3f01/README.md) |
| 13&#160;Jul&#160;23&#160;20:37&#160;UTC | SHAKEN 622J | 09&#160;Jan&#160;24&#160;20:37&#160;UTC | true | [view](CERTS/15fbab040bd63fddd486a1fb0981b328f24805edcb6ab3bd4ed3aa7100b8270a/README.md) |
| 19&#160;Jul&#160;23&#160;15:19&#160;UTC | SHAKEN 737J | 18&#160;Jul&#160;24&#160;15:19&#160;UTC | true | [view](CERTS/471e22577c0adbd26a690e87b378c41df242a42d52a06773fc91feb8283a34c0/README.md) |
| 08&#160;Aug&#160;23&#160;15:15&#160;UTC | SHAKEN 073H | 07&#160;Aug&#160;24&#160;15:15&#160;UTC | true | [view](CERTS/a4ec097b80cea72a2d5356c16b28166d113944db1d1e6d23b6d1d1258a66d946/README.md) |
| 11&#160;Sep&#160;23&#160;17:17&#160;UTC | SHAKEN 505J | 10&#160;Dec&#160;23&#160;17:17&#160;UTC | true | [view](CERTS/3089201a3f99caa04f6249012725a3b9ae9d3f35140fc81a369b0e9351d9c922/README.md) |
| 11&#160;Sep&#160;23&#160;17:18&#160;UTC | SHAKEN 505J | 10&#160;Dec&#160;23&#160;17:18&#160;UTC | true | [view](CERTS/30e65a228b3ae012e4d93552e2f1acf5a0359f3a496c3cd91f191f27cd46b228/README.md) |
| 27&#160;Sep&#160;23&#160;01:49&#160;UTC | SHAKEN 706J | 26&#160;Sep&#160;24&#160;01:49&#160;UTC | true | [view](CERTS/456dbbfbbde51c8c677b0b27e2fce5752a7de5c3407b4760d6fe1a306f3d8b94/README.md) |
| 07&#160;Oct&#160;23&#160;20:53&#160;UTC | SHAKEN 815G | 05&#160;Jan&#160;24&#160;20:53&#160;UTC | true | [view](CERTS/5e3f087377972cddbfa28ad890648c64ac3f5826d73bc03376633716bdaf0b43/README.md) |
| 11&#160;Oct&#160;23&#160;07:59&#160;UTC | SHAKEN 193E | 10&#160;Dec&#160;23&#160;07:59&#160;UTC | true | [view](CERTS/24ff7ba8e2a7f3cc30f93197bc836d803e9016fc34fb1bd69c67f9753e1c4613/README.md) |
| 16&#160;Oct&#160;23&#160;15:00&#160;UTC | SHAKEN 4036 | 13&#160;Apr&#160;24&#160;15:00&#160;UTC | true | [view](CERTS/906907e82fc83e435a762a6274c36c5ea6c88d6f933ba0ae03fe167912cc8429/README.md) |
| 21&#160;Oct&#160;23&#160;03:43&#160;UTC | SHAKEN 807J | 20&#160;Dec&#160;23&#160;03:43&#160;UTC | true | [view](CERTS/87b436606e63deb46efe2eb19aefbadb9d8485e0c414b6ee71d9419e745871fd/README.md) |
| 29&#160;Oct&#160;23&#160;20:42&#160;UTC | SHAKEN 577F | 28&#160;Nov&#160;23&#160;20:42&#160;UTC | true | [view](CERTS/21c8e10692f0772297b5e9521d99affce8dbd70a5a86b2d6fca54c2f0b2e4b9f/README.md) |
| 31&#160;Oct&#160;23&#160;17:32&#160;UTC | SHAKEN 551G | 30&#160;Nov&#160;23&#160;17:32&#160;UTC | true | [view](CERTS/747ce012af7cb00c17e5563d6884f6bd7297438b9452de5a92ab26ec7cd6c66d/README.md) |
| 01&#160;Nov&#160;23&#160;20:04&#160;UTC | SHAKEN 952J | 01&#160;Dec&#160;23&#160;20:04&#160;UTC | true | [view](CERTS/a35ea0c896a2916f3eb992cc7854fd4937e66659b1872c1711f27e6513aafd84/README.md) |
| 04&#160;Nov&#160;23&#160;05:33&#160;UTC | SHAKEN 345J | 04&#160;Dec&#160;23&#160;05:33&#160;UTC | true | [view](CERTS/123ab627570a450aceb599243861d6f75ed0c40993f6d4f563db5b80dc88b61d/README.md) |
| 06&#160;Nov&#160;23&#160;09:42&#160;UTC | SHAKEN 8526 | 05&#160;Nov&#160;24&#160;09:42&#160;UTC | true | [view](CERTS/4c86bf33be8b4189a469827d24c257723b4e5e3236981d9666d04de493b5cb6a/README.md) |
| 09&#160;Nov&#160;23&#160;15:11&#160;UTC | SHAKEN 366G | 09&#160;Dec&#160;23&#160;15:11&#160;UTC | true | [view](CERTS/9ab0a046628e9aee2c3b57e580f9deeded2b53c0cba6012dd16f027dbf0f49c2/README.md) |
| 10&#160;Nov&#160;23&#160;22:02&#160;UTC | SHAKEN 841J | 24&#160;Nov&#160;23&#160;22:02&#160;UTC | true | [view](CERTS/11fb88f8f94d0c6a12d9ec5f6121f25bf58b995a484795a6ecf6cb0e26b2d88c/README.md) |
| 13&#160;Nov&#160;23&#160;17:34&#160;UTC | SHAKEN 722J | 13&#160;Dec&#160;23&#160;17:34&#160;UTC | true | [view](CERTS/ffa6831454ed93bdbf626272b821ecf80e15570f68f94f9b1fe0065b5e5f6c54/README.md) |
| 13&#160;Nov&#160;23&#160;21:15&#160;UTC | SHAKEN 738J | 22&#160;Nov&#160;23&#160;21:15&#160;UTC | true | [view](CERTS/9fa8476e822c5cfb24e9f7571b17506f19db433c75d061ad34bfbccf6a60203f/README.md) |
| 15&#160;Nov&#160;23&#160;02:51&#160;UTC | SHAKEN 345J | 15&#160;Dec&#160;23&#160;02:51&#160;UTC | true | [view](CERTS/1a88a3c0462f90cd3362ad4dd4e8c33342e084d2f922abbf8c4deed275e347a8/README.md) |
| 15&#160;Nov&#160;23&#160;18:53&#160;UTC | SHAKEN 1917 | 22&#160;Nov&#160;23&#160;18:53&#160;UTC | true | [view](CERTS/0bc4b0f5354d98e56ca913d951ad52885fd9a08b682306361beb8c1389d2812a/README.md) |
| 15&#160;Nov&#160;23&#160;19:14&#160;UTC | SHAKEN 647K | 22&#160;Nov&#160;23&#160;19:14&#160;UTC | true | [view](CERTS/5ccdf8e2dc0fd22ece9d71994ff2cb3349013ea71924862f3519806b172db187/README.md) |
| 15&#160;Nov&#160;23&#160;20:47&#160;UTC | SHAKEN 983J | 22&#160;Nov&#160;23&#160;20:47&#160;UTC | true | [view](CERTS/9ea4d5ca76acc07a8f3fff1aa1f1f28a8f9c10e9e2693427c05bd3983e171ed5/README.md) |
| 16&#160;Nov&#160;23&#160;15:04&#160;UTC | SHAKEN 2550 | 23&#160;Nov&#160;23&#160;15:04&#160;UTC | true | [view](CERTS/8817f159d626833d9905cb1087120e556b3969b9dbada53680434b0abaa93108/README.md) |
| 16&#160;Nov&#160;23&#160;16:19&#160;UTC | SHAKEN 646K | 23&#160;Nov&#160;23&#160;16:19&#160;UTC | true | [view](CERTS/a2791ee0972cb1be84bbfa150b3a5f164650ce958ab4a394751aefbc1254da91/README.md) |
| 16&#160;Nov&#160;23&#160;17:57&#160;UTC | SHAKEN 833J | 23&#160;Nov&#160;23&#160;17:57&#160;UTC | true | [view](CERTS/a7e8efe7b47e08f564972ad9d2900f5e5e17e4d53e8fd71fceb1d73dd3c4f8d2/README.md) |
| 16&#160;Nov&#160;23&#160;20:47&#160;UTC | SHAKEN 177K | 23&#160;Nov&#160;23&#160;20:47&#160;UTC | true | [view](CERTS/be47991290bbfedf8734ab62d39769ff5a4ecb0713d4c6cd203f9bc38827a6b8/README.md) |
| 17&#160;Nov&#160;23&#160;15:12&#160;UTC | SHAKEN 747J | 24&#160;Nov&#160;23&#160;15:12&#160;UTC | true | [view](CERTS/4a734f7f8e9cae6cc77e65248d381a0ba5cee61af2b6ecbcaa7dbd29e9abb2de/README.md) |
| 17&#160;Nov&#160;23&#160;18:43&#160;UTC | SHAKEN 060K | 24&#160;Nov&#160;23&#160;18:43&#160;UTC | true | [view](CERTS/c5ed7478d05c794a24cedaa52d6688d66dfe8747fb63a76d1f34b9f85770728d/README.md) |
| 17&#160;Nov&#160;23&#160;19:11&#160;UTC | SHAKEN 089K | 24&#160;Nov&#160;23&#160;19:11&#160;UTC | true | [view](CERTS/526f2d5d945550d1533780c0351170d015b5dd90cb07974fef9dc487f3538361/README.md) |
| 17&#160;Nov&#160;23&#160;20:52&#160;UTC | SHAKEN 674J | 24&#160;Nov&#160;23&#160;20:52&#160;UTC | true | [view](CERTS/ab6c338e268b04c3e7355c19197fceec27f68d199f1ad45b359bf3fc3ddad6e2/README.md) |
| 17&#160;Nov&#160;23&#160;20:57&#160;UTC | SHAKEN 856H | 24&#160;Nov&#160;23&#160;20:57&#160;UTC | true | [view](CERTS/ebfa4159ac40da0746963adc2be3ab51b8d58314a6b9312d357bcebf7e096250/README.md) |
| 17&#160;Nov&#160;23&#160;20:59&#160;UTC | SHAKEN 819H | 24&#160;Nov&#160;23&#160;20:59&#160;UTC | true | [view](CERTS/e996e26789dced38dc1238fc1d6bc028b4977237b805e0a22fa3c125cf26e6a2/README.md) |
| 17&#160;Nov&#160;23&#160;21:04&#160;UTC | SHAKEN 769J | 24&#160;Nov&#160;23&#160;21:04&#160;UTC | true | [view](CERTS/01e0f7db18c96c000de137c7b13111cdb2dee980c5a842db8aa07df0f6b36c63/README.md) |
| 17&#160;Nov&#160;23&#160;21:04&#160;UTC | SHAKEN 0127 | 24&#160;Nov&#160;23&#160;21:04&#160;UTC | true | [view](CERTS/931b20937e723e00c65aa3fbc15f9195810a4a0615402860f7159f39e4f97949/README.md) |
| 17&#160;Nov&#160;23&#160;21:04&#160;UTC | SHAKEN 469A | 24&#160;Nov&#160;23&#160;21:04&#160;UTC | true | [view](CERTS/b7e68b5c760d660d5c926cd04cfe3904ba54e6af2200952baf3d85f43dfc33ae/README.md) |
| 17&#160;Nov&#160;23&#160;21:04&#160;UTC | SHAKEN 726J | 24&#160;Nov&#160;23&#160;21:04&#160;UTC | true | [view](CERTS/1617f6c8e358c58e0114662bbbd84ed186bc6408211d5604f7eac2dcd4b9fa06/README.md) |
| 17&#160;Nov&#160;23&#160;21:09&#160;UTC | SHAKEN 459J | 24&#160;Nov&#160;23&#160;21:09&#160;UTC | true | [view](CERTS/c6ec896a4ab78d7c1512658dc14180f9d26adfceddee0de7ee4cd76b46f63a86/README.md) |
| 18&#160;Nov&#160;23&#160;15:47&#160;UTC | SHAKEN 444G | 25&#160;Nov&#160;23&#160;15:47&#160;UTC | true | [view](CERTS/2dc4a166683a3ef62581d1411849eed207d33ac10af3dd7638d3d9ffa314e74c/README.md) |
| 18&#160;Nov&#160;23&#160;20:47&#160;UTC | SHAKEN 983J | 25&#160;Nov&#160;23&#160;20:47&#160;UTC | true | [view](CERTS/7a382e7afcbaf790869894ad06aa878a30cc71e0d9accfab7cf0618554752633/README.md) |
| 18&#160;Nov&#160;23&#160;21:18&#160;UTC | SHAKEN 663G | 25&#160;Nov&#160;23&#160;21:18&#160;UTC | true | [view](CERTS/120b3569c4deb57a938fca15281078c319514ce76ca25492750d23ec90e28af0/README.md) |
| 18&#160;Nov&#160;23&#160;21:25&#160;UTC | SHAKEN 157C | 25&#160;Nov&#160;23&#160;21:25&#160;UTC | true | [view](CERTS/faab0f70220e183790d7702da98f1f2f80546410ef99addaefcf3c742b85bba6/README.md) |
| 19&#160;Nov&#160;23&#160;01:15&#160;UTC | SHAKEN 0172 | 26&#160;Nov&#160;23&#160;01:15&#160;UTC | true | [view](CERTS/acef983688bff43a3c70debe202a906dd05254321ed1edb9ebe3a855cec64709/README.md) |
| 19&#160;Nov&#160;23&#160;05:37&#160;UTC | SHAKEN 0378 | 26&#160;Nov&#160;23&#160;05:36&#160;UTC | true | [view](CERTS/091d5bfe73159cfd2219dd1b022b360e86971a002d91af15dd114ec1ed9eeca7/README.md) |
| 19&#160;Nov&#160;23&#160;14:15&#160;UTC | SHAKEN 495K | 26&#160;Nov&#160;23&#160;14:15&#160;UTC | true | [view](CERTS/4cba9a358545a3caff52f12dd4a79bfd186f4fd988c4bfe19b8646e81ef05aac/README.md) |
| 19&#160;Nov&#160;23&#160;16:19&#160;UTC | SHAKEN 646K | 26&#160;Nov&#160;23&#160;16:19&#160;UTC | true | [view](CERTS/3316fb17bcc9572bf9a5dea23d621529d9a50328b470791881aed044113a1ed4/README.md) |
| 19&#160;Nov&#160;23&#160;17:57&#160;UTC | SHAKEN 833J | 26&#160;Nov&#160;23&#160;17:57&#160;UTC | true | [view](CERTS/1debaeeea40a6498d5022b9d20e7ab53ffd1bccb41a2aff92f89f87af3a3dd23/README.md) |
| 19&#160;Nov&#160;23&#160;20:14&#160;UTC | SHAKEN 952J | 19&#160;Dec&#160;23&#160;20:14&#160;UTC | true | [view](CERTS/f8af5c3d8985594921c4658dd1889e7670d0cd729e292f8e2405e215860ab80e/README.md) |
| 19&#160;Nov&#160;23&#160;20:48&#160;UTC | SHAKEN 177K | 26&#160;Nov&#160;23&#160;20:48&#160;UTC | true | [view](CERTS/0e06ac62890d5bd8d30ba4422bf0e6077c3040a8d502bc9f77a93f81bd7d9a6a/README.md) |
| 20&#160;Nov&#160;23&#160;20:57&#160;UTC | SHAKEN 733J | 27&#160;Nov&#160;23&#160;20:57&#160;UTC | true | [view](CERTS/0dd384941a89f683a8e32468f50cb411eed0ebe92ed0c5c241b60b796292f85e/README.md) |
| 20&#160;Nov&#160;23&#160;21:05&#160;UTC | SHAKEN 469A | 27&#160;Nov&#160;23&#160;21:05&#160;UTC | true | [view](CERTS/5616be66af20424264738071c4048307e362de452e9eafd2f6b131abf10ff8fa/README.md) |
| 20&#160;Nov&#160;23&#160;21:16&#160;UTC | SHAKEN 738J | 29&#160;Nov&#160;23&#160;21:16&#160;UTC | true | [view](CERTS/1484fa0270b00bc45d75eb840aa6388deb78d595556be04da305af5b76eb8dc8/README.md) |
| 21&#160;Nov&#160;23&#160;17:34&#160;UTC | SHAKEN 551G | 21&#160;Dec&#160;23&#160;17:34&#160;UTC | true | [view](CERTS/07666e96f8549a2f5852bd5dd8766670d08d66a332370a25619b15319b603811/README.md) |

#### CA Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 24&#160;Oct&#160;22&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Issuing CA4 | 23&#160;Oct&#160;32&#160;23:59&#160;UTC | false | [view](CERTS/80dd12d935f46c256c7c1289a0834b21ec8f2f8c35f2864928d1e75f3a280665/README.md) |
| 24&#160;Oct&#160;22&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Root CA2 | 23&#160;Oct&#160;42&#160;23:59&#160;UTC | false | [view](CERTS/a26e04fc786ab70b8085236b2c53f8cfbf5d0c6a5c2c9c3e9f91669fbb8ea4d5/README.md) |


Generated: 22 Nov 23 03:38 UTC