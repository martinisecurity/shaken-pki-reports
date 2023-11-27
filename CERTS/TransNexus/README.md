# STIR/SHAKEN CA Ecosystem Compliance

## TransNexus

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 5025 certificates were included in the corpus being tested
- 3 certificates in the corpus were skipped because they are duplicates
- 4533 certificates in the corpus were skipped because they are expired
- 458 certificates in the corpus were skipped because they are not currently trusted
- 31 certificates being tested against the remaining rules
- 1.00 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 0.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 9.68% of certificates are too old to be assessed against currently enforced expectations
- 193 days is the average remaining validity for the certificates in the corpus
- 192 days is the average initial validity for the certificates in the corpus
- 16 certificates expire in the next 30 days
- 1.19 average number of unexpired certificates per OCN observed
- 26 unique OCNs observed in unexpired and valid certificate corpus

| Instances | Test | Source |
|-----------|------|--------|
| 31 | [e_us_cp_subject_sn_shall](ISSUES/e_us_cp_subject_sn_shall/README.md) | US_SHAKEN_CP |

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
- 5955 days is the average remaining validity for the certificates in the corpus
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
| 13&#160;Nov&#160;23&#160;17:34&#160;UTC | SHAKEN 722J | 13&#160;Dec&#160;23&#160;17:34&#160;UTC | true | [view](CERTS/ffa6831454ed93bdbf626272b821ecf80e15570f68f94f9b1fe0065b5e5f6c54/README.md) |
| 15&#160;Nov&#160;23&#160;02:51&#160;UTC | SHAKEN 345J | 15&#160;Dec&#160;23&#160;02:51&#160;UTC | true | [view](CERTS/1a88a3c0462f90cd3362ad4dd4e8c33342e084d2f922abbf8c4deed275e347a8/README.md) |
| 19&#160;Nov&#160;23&#160;20:14&#160;UTC | SHAKEN 952J | 19&#160;Dec&#160;23&#160;20:14&#160;UTC | true | [view](CERTS/f8af5c3d8985594921c4658dd1889e7670d0cd729e292f8e2405e215860ab80e/README.md) |
| 20&#160;Nov&#160;23&#160;21:16&#160;UTC | SHAKEN 738J | 29&#160;Nov&#160;23&#160;21:16&#160;UTC | true | [view](CERTS/1484fa0270b00bc45d75eb840aa6388deb78d595556be04da305af5b76eb8dc8/README.md) |
| 21&#160;Nov&#160;23&#160;17:34&#160;UTC | SHAKEN 551G | 21&#160;Dec&#160;23&#160;17:34&#160;UTC | true | [view](CERTS/07666e96f8549a2f5852bd5dd8766670d08d66a332370a25619b15319b603811/README.md) |

#### CA Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 24&#160;Oct&#160;22&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Issuing CA4 | 23&#160;Oct&#160;32&#160;23:59&#160;UTC | false | [view](CERTS/80dd12d935f46c256c7c1289a0834b21ec8f2f8c35f2864928d1e75f3a280665/README.md) |
| 24&#160;Oct&#160;22&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Root CA2 | 23&#160;Oct&#160;42&#160;23:59&#160;UTC | false | [view](CERTS/a26e04fc786ab70b8085236b2c53f8cfbf5d0c6a5c2c9c3e9f91669fbb8ea4d5/README.md) |


Generated: 27 Nov 23 22:56 UTC