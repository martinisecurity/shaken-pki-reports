# STIR/SHAKEN CA Ecosystem Compliance

## TransNexus

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 5678 certificates were included in the corpus being tested
- 3 certificates in the corpus were skipped because they are duplicates
- 5184 certificates in the corpus were skipped because they are expired
- 458 certificates in the corpus were skipped because they are not currently trusted
- 33 certificates being tested against the remaining rules
- 1.00 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 0.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 174 days is the average remaining validity for the certificates in the corpus
- 174 days is the average initial validity for the certificates in the corpus
- 14 certificates expire in the next 30 days
- 1.06 average number of unexpired certificates per OCN observed
- 31 unique OCNs observed in unexpired and valid certificate corpus

| Instances | Test | Source |
|-----------|------|--------|
| 33 | [e_us_cp_subject_sn_shall](ISSUES/e_us_cp_subject_sn_shall/README.md) | US_SHAKEN_CP |

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
- 5929 days is the average remaining validity for the certificates in the corpus
- 5479 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days

No error, warning, or notice level issues were found

### Details

#### Leaf Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 27&#160;Mar&#160;23&#160;17:18&#160;UTC | SHAKEN 2720 | 26&#160;Mar&#160;24&#160;17:18&#160;UTC | true | [view](CERTS/7282c0683ebff27881b1ccd1f66664c3102340ff657d9363fcecb9c1b7924c12/README.md) |
| 28&#160;Apr&#160;23&#160;17:50&#160;UTC | SHAKEN 663J | 27&#160;Apr&#160;24&#160;17:50&#160;UTC | true | [view](CERTS/bd944d8e1acfeaaa520a8c87826b19a8509b52b083525a940e9ee4db8af1a99b/README.md) |
| 24&#160;May&#160;23&#160;08:31&#160;UTC | SHAKEN 736J | 23&#160;May&#160;24&#160;08:31&#160;UTC | true | [view](CERTS/d273815c57e441d2066144a02d029b73c3d839f5d31cbaeb72a7b6bded62da53/README.md) |
| 06&#160;Jun&#160;23&#160;19:10&#160;UTC | SHAKEN 597J | 05&#160;Jun&#160;24&#160;19:10&#160;UTC | true | [view](CERTS/7f15a14bb0b8ad72d450e673344a747703cdbe15741d2696c371c83c5ca8e148/README.md) |
| 16&#160;Jun&#160;23&#160;18:51&#160;UTC | SHAKEN 857J | 15&#160;Jun&#160;24&#160;18:51&#160;UTC | true | [view](CERTS/97834968950fa42accd7ab2d9f43d12eb3dd919b5d5a72a153275419104aa534/README.md) |
| 08&#160;Jul&#160;23&#160;21:54&#160;UTC | SHAKEN 578J | 07&#160;Jul&#160;24&#160;21:54&#160;UTC | true | [view](CERTS/5f864f10b33ee5194c0f67cd87af157b58755671a4534b5bd9570ca5eacd3f01/README.md) |
| 19&#160;Jul&#160;23&#160;15:19&#160;UTC | SHAKEN 737J | 18&#160;Jul&#160;24&#160;15:19&#160;UTC | true | [view](CERTS/471e22577c0adbd26a690e87b378c41df242a42d52a06773fc91feb8283a34c0/README.md) |
| 08&#160;Aug&#160;23&#160;15:15&#160;UTC | SHAKEN 073H | 07&#160;Aug&#160;24&#160;15:15&#160;UTC | true | [view](CERTS/a4ec097b80cea72a2d5356c16b28166d113944db1d1e6d23b6d1d1258a66d946/README.md) |
| 27&#160;Sep&#160;23&#160;01:49&#160;UTC | SHAKEN 706J | 26&#160;Sep&#160;24&#160;01:49&#160;UTC | true | [view](CERTS/456dbbfbbde51c8c677b0b27e2fce5752a7de5c3407b4760d6fe1a306f3d8b94/README.md) |
| 16&#160;Oct&#160;23&#160;15:00&#160;UTC | SHAKEN 4036 | 13&#160;Apr&#160;24&#160;15:00&#160;UTC | true | [view](CERTS/906907e82fc83e435a762a6274c36c5ea6c88d6f933ba0ae03fe167912cc8429/README.md) |
| 06&#160;Nov&#160;23&#160;09:42&#160;UTC | SHAKEN 8526 | 05&#160;Nov&#160;24&#160;09:42&#160;UTC | true | [view](CERTS/4c86bf33be8b4189a469827d24c257723b4e5e3236981d9666d04de493b5cb6a/README.md) |
| 09&#160;Nov&#160;23&#160;22:47&#160;UTC | SHAKEN 621J | 07&#160;May&#160;24&#160;22:47&#160;UTC | true | [view](CERTS/92f1c34e2f1207cf5a547a8a7749fddcacfa95b144df511274f8547638eba51e/README.md) |
| 07&#160;Dec&#160;23&#160;17:35&#160;UTC | SHAKEN 505J | 04&#160;Jun&#160;24&#160;17:35&#160;UTC | true | [view](CERTS/677465016a885da84672aa6ec13c865e68d398d93c88d6da1c4e1d5a6c9a2314/README.md) |
| 07&#160;Dec&#160;23&#160;17:37&#160;UTC | SHAKEN 505J | 04&#160;Jun&#160;24&#160;17:37&#160;UTC | true | [view](CERTS/eb94a1635c1b9983430793df3cfb4fd316be625def2796006a78629b04199d8a/README.md) |
| 08&#160;Dec&#160;23&#160;21:57&#160;UTC | SHAKEN 299K | 07&#160;Dec&#160;24&#160;21:57&#160;UTC | true | [view](CERTS/7b677c8ad6481aa908931a3bec7ec5645e51770d15afca1e706b99f09203eca5/README.md) |
| 14&#160;Dec&#160;23&#160;17:33&#160;UTC | SHAKEN 622J | 11&#160;Jun&#160;24&#160;17:33&#160;UTC | true | [view](CERTS/4c841c675f8eb395ebd222c847faed78e3351a1421bfb95db5ee2f1b53c76c88/README.md) |
| 05&#160;Jan&#160;24&#160;11:50&#160;UTC | SHAKEN 815G | 04&#160;Apr&#160;24&#160;11:50&#160;UTC | true | [view](CERTS/8e1252bbfdba4afa4d8b8a9d5680c48f866572e76bcb086136f89084a8dac1ae/README.md) |
| 10&#160;Jan&#160;24&#160;19:01&#160;UTC | SHAKEN 193E | 10&#160;Mar&#160;24&#160;19:01&#160;UTC | true | [view](CERTS/702ffeabbff3a9d2213e3af119005dc2e6d8edc48b149b61841f2dbe8cca3bbd/README.md) |
| 17&#160;Jan&#160;24&#160;03:20&#160;UTC | SHAKEN 345J | 16&#160;Feb&#160;24&#160;03:20&#160;UTC | true | [view](CERTS/54bb0f84701f4ef932dec1ca4b619c1dc2e1891070f0e3bb0b43f734e7da6cf9/README.md) |
| 19&#160;Jan&#160;24&#160;20:13&#160;UTC | SHAKEN 807J | 19&#160;Mar&#160;24&#160;20:13&#160;UTC | true | [view](CERTS/93ace8966059b5d32985e46f82abd59a2f44bc65752322125f58a98258ea3460/README.md) |
| 23&#160;Jan&#160;24&#160;17:52&#160;UTC | SHAKEN 551G | 22&#160;Feb&#160;24&#160;17:52&#160;UTC | true | [view](CERTS/a255637799d625af4a1ecf81198dd52bb65f724b6bfd37d5bd6fac37c8fcb380/README.md) |
| 24&#160;Jan&#160;24&#160;17:53&#160;UTC | SHAKEN 722J | 23&#160;Feb&#160;24&#160;17:53&#160;UTC | true | [view](CERTS/bb7cb7abfe59afef301fb627418f653f791ae0b57b44301b663bcec9140fef74/README.md) |
| 27&#160;Jan&#160;24&#160;05:53&#160;UTC | SHAKEN 345J | 26&#160;Feb&#160;24&#160;05:53&#160;UTC | true | [view](CERTS/d93d8dac607b66507bd26a4a4a72496d769a9de13bc9782e9440d880229bf8c5/README.md) |
| 27&#160;Jan&#160;24&#160;21:02&#160;UTC | SHAKEN 577F | 26&#160;Feb&#160;24&#160;21:02&#160;UTC | true | [view](CERTS/dac72f86a5b887682e608d864e6b448d4811b0b09a53802cb135aa073dcde60f/README.md) |
| 30&#160;Jan&#160;24&#160;20:44&#160;UTC | SHAKEN 952J | 29&#160;Feb&#160;24&#160;20:44&#160;UTC | true | [view](CERTS/803f096972d200c7a10b5caa14407a1eb1b6b93ad6b4098867d9a3598020b540/README.md) |
| 31&#160;Jan&#160;24&#160;18:59&#160;UTC | SHAKEN 873J | 30&#160;Jan&#160;25&#160;18:59&#160;UTC | true | [view](CERTS/8342db52ef1e9b25620a252b82f1378cff688a8cc0e594a6669f50ac17b34d03/README.md) |
| 01&#160;Feb&#160;24&#160;15:23&#160;UTC | SHAKEN 366G | 02&#160;Mar&#160;24&#160;15:23&#160;UTC | true | [view](CERTS/9f2f1f1beb51e5c436bf480d156687ce992711454d654b0aa710935089a10a79/README.md) |
| 02&#160;Feb&#160;24&#160;22:41&#160;UTC | SHAKEN 841J | 16&#160;Feb&#160;24&#160;22:41&#160;UTC | true | [view](CERTS/8d17ad3cfdc1edbe8d1ca89e32dc50ae264431f22c69140f819aaf300648c010/README.md) |
| 05&#160;Feb&#160;24&#160;18:02&#160;UTC | SHAKEN 833J | 12&#160;Feb&#160;24&#160;18:02&#160;UTC | true | [view](CERTS/a50b3cd1aafd1a893d274113dd1ad8c1fa37a8003b0fc1ae6e2fba1f4b66c96c/README.md) |
| 05&#160;Feb&#160;24&#160;20:54&#160;UTC | SHAKEN 177K | 12&#160;Feb&#160;24&#160;20:54&#160;UTC | true | [view](CERTS/3eb95d0e3191c3be8338594ac785624879fe5c61815e58025d750cd3f74ccfe7/README.md) |
| 06&#160;Feb&#160;24&#160;15:15&#160;UTC | SHAKEN 738J | 15&#160;Feb&#160;24&#160;15:15&#160;UTC | true | [view](CERTS/9b5a695bd8ecedf3bb732257088c3bd761f18200562dac433d55312e1caa7989/README.md) |
| 06&#160;Feb&#160;24&#160;20:59&#160;UTC | SHAKEN 674J | 13&#160;Feb&#160;24&#160;20:59&#160;UTC | true | [view](CERTS/150109060a4ca84245287e0fc05a6008f51a3db509b26872c345d302ad58bebe/README.md) |
| 06&#160;Feb&#160;24&#160;21:17&#160;UTC | SHAKEN 459J | 13&#160;Feb&#160;24&#160;21:17&#160;UTC | true | [view](CERTS/c6ac34005c06d8fc7a80108e795097e2485c5b2e2c71a53a37d462372ca8bb9e/README.md) |

#### CA Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 24&#160;Oct&#160;22&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Issuing CA4 | 23&#160;Oct&#160;32&#160;23:59&#160;UTC | false | [view](CERTS/80dd12d935f46c256c7c1289a0834b21ec8f2f8c35f2864928d1e75f3a280665/README.md) |
| 24&#160;Oct&#160;22&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Root CA2 | 23&#160;Oct&#160;42&#160;23:59&#160;UTC | false | [view](CERTS/a26e04fc786ab70b8085236b2c53f8cfbf5d0c6a5c2c9c3e9f91669fbb8ea4d5/README.md) |


Generated: 12 Feb 24 17:02 UTC