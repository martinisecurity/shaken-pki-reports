# STIR/SHAKEN CA Ecosystem Compliance

## TransNexus

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 5676 certificates were included in the corpus being tested
- 2 certificates in the corpus were skipped because they are duplicates
- 5201 certificates in the corpus were skipped because they are expired
- 457 certificates in the corpus were skipped because they are not currently trusted
- 16 certificates being tested against the remaining rules
- 1.00 issues on average found in unexpired, trusted, and non-compliant certificates
- 93.75% of certificates contain one or more Error level issue
- 0.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 306 days is the average remaining validity for the certificates in the corpus
- 319 days is the average initial validity for the certificates in the corpus
- 1 certificates expire in the next 30 days
- 1.07 average number of unexpired certificates per OCN observed
- 15 unique OCNs observed in unexpired and valid certificate corpus

| Instances | Test | Source |
|-----------|------|--------|
| 15 | [e_us_cp_subject_sn_shall](ISSUES/e_us_cp_subject_sn_shall/README.md) | US_SHAKEN_CP |

#### CA Certificates

- 6 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 0 certificates in the corpus were skipped because they are expired
- 3 certificates in the corpus were skipped because they are not currently trusted
- 3 certificates being tested against the remaining rules
- 0.00 issues on average found in unexpired, trusted, and non-compliant certificates
- 0.00% of certificates contain one or more Error level issue
- 0.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 66.67% of certificates are too old to be assessed against currently enforced expectations
- 5346 days is the average remaining validity for the certificates in the corpus
- 4870 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days

No error, warning, or notice level issues were found

### Details

#### Leaf Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 28&#160;Apr&#160;23&#160;17:50&#160;UTC | SHAKEN 663J | 27&#160;Apr&#160;24&#160;17:50&#160;UTC | true | [view](CERTS/bd944d8e1acfeaaa520a8c87826b19a8509b52b083525a940e9ee4db8af1a99b/README.md) |
| 24&#160;May&#160;23&#160;08:31&#160;UTC | SHAKEN 736J | 23&#160;May&#160;24&#160;08:31&#160;UTC | true | [view](CERTS/d273815c57e441d2066144a02d029b73c3d839f5d31cbaeb72a7b6bded62da53/README.md) |
| 06&#160;Jun&#160;23&#160;19:10&#160;UTC | SHAKEN 597J | 05&#160;Jun&#160;24&#160;19:10&#160;UTC | true | [view](CERTS/7f15a14bb0b8ad72d450e673344a747703cdbe15741d2696c371c83c5ca8e148/README.md) |
| 16&#160;Jun&#160;23&#160;18:51&#160;UTC | SHAKEN 857J | 15&#160;Jun&#160;24&#160;18:51&#160;UTC | true | [view](CERTS/97834968950fa42accd7ab2d9f43d12eb3dd919b5d5a72a153275419104aa534/README.md) |
| 08&#160;Jul&#160;23&#160;21:54&#160;UTC | SHAKEN 578J | 07&#160;Jul&#160;24&#160;21:54&#160;UTC | true | [view](CERTS/5f864f10b33ee5194c0f67cd87af157b58755671a4534b5bd9570ca5eacd3f01/README.md) |
| 19&#160;Jul&#160;23&#160;15:19&#160;UTC | SHAKEN 737J | 18&#160;Jul&#160;24&#160;15:19&#160;UTC | true | [view](CERTS/471e22577c0adbd26a690e87b378c41df242a42d52a06773fc91feb8283a34c0/README.md) |
| 08&#160;Aug&#160;23&#160;15:15&#160;UTC | SHAKEN 073H | 07&#160;Aug&#160;24&#160;15:15&#160;UTC | true | [view](CERTS/a4ec097b80cea72a2d5356c16b28166d113944db1d1e6d23b6d1d1258a66d946/README.md) |
| 27&#160;Sep&#160;23&#160;01:49&#160;UTC | SHAKEN 706J | 26&#160;Sep&#160;24&#160;01:49&#160;UTC | true | [view](CERTS/456dbbfbbde51c8c677b0b27e2fce5752a7de5c3407b4760d6fe1a306f3d8b94/README.md) |
| 06&#160;Nov&#160;23&#160;09:42&#160;UTC | SHAKEN 8526 | 05&#160;Nov&#160;24&#160;09:42&#160;UTC | true | [view](CERTS/4c86bf33be8b4189a469827d24c257723b4e5e3236981d9666d04de493b5cb6a/README.md) |
| 09&#160;Nov&#160;23&#160;22:47&#160;UTC | SHAKEN 621J | 07&#160;May&#160;24&#160;22:47&#160;UTC | true | [view](CERTS/92f1c34e2f1207cf5a547a8a7749fddcacfa95b144df511274f8547638eba51e/README.md) |
| 07&#160;Dec&#160;23&#160;17:35&#160;UTC | SHAKEN 505J | 04&#160;Jun&#160;24&#160;17:35&#160;UTC | true | [view](CERTS/677465016a885da84672aa6ec13c865e68d398d93c88d6da1c4e1d5a6c9a2314/README.md) |
| 07&#160;Dec&#160;23&#160;17:37&#160;UTC | SHAKEN 505J | 04&#160;Jun&#160;24&#160;17:37&#160;UTC | true | [view](CERTS/eb94a1635c1b9983430793df3cfb4fd316be625def2796006a78629b04199d8a/README.md) |
| 08&#160;Dec&#160;23&#160;21:57&#160;UTC | SHAKEN 299K | 07&#160;Dec&#160;24&#160;21:57&#160;UTC | true | [view](CERTS/7b677c8ad6481aa908931a3bec7ec5645e51770d15afca1e706b99f09203eca5/README.md) |
| 14&#160;Dec&#160;23&#160;17:33&#160;UTC | SHAKEN 622J | 11&#160;Jun&#160;24&#160;17:33&#160;UTC | true | [view](CERTS/4c841c675f8eb395ebd222c847faed78e3351a1421bfb95db5ee2f1b53c76c88/README.md) |
| 31&#160;Jan&#160;24&#160;18:59&#160;UTC | SHAKEN 873J | 30&#160;Jan&#160;25&#160;18:59&#160;UTC | true | [view](CERTS/8342db52ef1e9b25620a252b82f1378cff688a8cc0e594a6669f50ac17b34d03/README.md) |
| 02&#160;Apr&#160;24&#160;15:50&#160;UTC | SHAKEN 2720 | 02&#160;Apr&#160;25&#160;15:50&#160;UTC | false | [view](CERTS/95a275a51dc842f102c43cd6b11f37b85a911ba727fca4cc6e3f9b859e05070d/README.md) |

#### CA Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 24&#160;Oct&#160;22&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Issuing CA4 | 23&#160;Oct&#160;32&#160;23:59&#160;UTC | false | [view](CERTS/80dd12d935f46c256c7c1289a0834b21ec8f2f8c35f2864928d1e75f3a280665/README.md) |
| 24&#160;Oct&#160;22&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Root CA2 | 23&#160;Oct&#160;42&#160;23:59&#160;UTC | false | [view](CERTS/a26e04fc786ab70b8085236b2c53f8cfbf5d0c6a5c2c9c3e9f91669fbb8ea4d5/README.md) |
| 21&#160;Mar&#160;24&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Issuing CA5 | 20&#160;Mar&#160;34&#160;23:59&#160;UTC | false | [view](CERTS/cd50eeb8c083af686a49964a10b030048b800530edbeee8f0991388c3a79e75a/README.md) |


Generated: 05 Apr 24 19:04 UTC