# STIR/SHAKEN CA Ecosystem Compliance

## TransNexus

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 3678 certificates were included in the corpus being tested
- 2 certificates in the corpus were skipped because they are duplicates
- 3652 certificates in the corpus were skipped because they are expired
- 1 certificates in the corpus were skipped because they are not currently trusted
- 23 certificates being tested against the remaining rules
- 1.09 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 4.35% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 217 days is the average remaining validity for the certificates in the corpus
- 224 days is the average initial validity for the certificates in the corpus
- 8 certificates expire in the next 30 days
- 1.05 average number of unexpired certificates per OCN observed
- 22 unique OCNs observed in unexpired and valid certificate corpus

| Instances | Test | Source |
|-----------|------|--------|
| 1 | [e_incorrect_ku_encoding](ISSUES/e_incorrect_ku_encoding/README.md) | RFC5280 |
| 23 | [e_us_cp_subject_sn](ISSUES/e_us_cp_subject_sn/README.md) | US_SHAKEN_CP |
| 1 | [w_shaken_subject_rdn_unknown](ISSUES/w_shaken_subject_rdn_unknown/README.md) | SHAKEN_PKI_BEST_PRACTICES |

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
- 5428 days is the average remaining validity for the certificates in the corpus
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
| 13&#160;Sep&#160;22&#160;17:24&#160;UTC | SHAKEN 706J | 13&#160;Sep&#160;23&#160;17:24&#160;UTC | true | [view](CERTS/1eea5e17ae8b36cb5b6d657178273af745cea015e63586e6d44d1ba2ee8f722a/README.md) |
| 07&#160;Dec&#160;22&#160;00:45&#160;UTC | SHAKEN 299K | 07&#160;Dec&#160;23&#160;00:45&#160;UTC | true | [view](CERTS/7b5e04ff01ace2ee384df3507886684204901bf23851866eb6b3891ecdd96a92/README.md) |
| 31&#160;Dec&#160;22&#160;03:33&#160;UTC | SHAKEN 8526 | 31&#160;Dec&#160;23&#160;03:33&#160;UTC | true | [view](CERTS/96672cb8cc7c71e38d9e6350eac6c4f5c430c1eefeabe00cc095527e20a20c9b/README.md) |
| 25&#160;Jan&#160;23&#160;14:28&#160;UTC | SHAKEN 873J | 25&#160;Jan&#160;24&#160;14:28&#160;UTC | true | [view](CERTS/36e9ba32d337ad13909bee0b62125c9769a049c4ae10d749687d75af4e49387d/README.md) |
| 27&#160;Mar&#160;23&#160;17:18&#160;UTC | SHAKEN 2720 | 26&#160;Mar&#160;24&#160;17:18&#160;UTC | true | [view](CERTS/7282c0683ebff27881b1ccd1f66664c3102340ff657d9363fcecb9c1b7924c12/README.md) |
| 27&#160;Apr&#160;23&#160;17:52&#160;UTC | SHAKEN 807J | 24&#160;Oct&#160;23&#160;17:52&#160;UTC | true | [view](CERTS/b42f1c446cebfad758f79a67d3b0460b7c3d0670acc426dfa311ac86019cc289/README.md) |
| 28&#160;Apr&#160;23&#160;17:50&#160;UTC | SHAKEN 663J | 27&#160;Apr&#160;24&#160;17:50&#160;UTC | true | [view](CERTS/bd944d8e1acfeaaa520a8c87826b19a8509b52b083525a940e9ee4db8af1a99b/README.md) |
| 08&#160;May&#160;23&#160;17:37&#160;UTC | SHAKEN 4036 | 04&#160;Nov&#160;23&#160;17:37&#160;UTC | true | [view](CERTS/1f7c5f89be3bc2a774c0b657a342d25a23ec7f07e38c8e3ed1d45b720641d919/README.md) |
| 24&#160;May&#160;23&#160;08:31&#160;UTC | SHAKEN 736J | 23&#160;May&#160;24&#160;08:31&#160;UTC | true | [view](CERTS/d273815c57e441d2066144a02d029b73c3d839f5d31cbaeb72a7b6bded62da53/README.md) |
| 06&#160;Jun&#160;23&#160;13:41&#160;UTC | SHAKEN 6628 | 03&#160;Dec&#160;23&#160;13:41&#160;UTC | true | [view](CERTS/1892b1588772ce058cea6d06c2e82ea66fae4902994a60ab326d96de2635e22d/README.md) |
| 06&#160;Jun&#160;23&#160;19:10&#160;UTC | SHAKEN 597J | 05&#160;Jun&#160;24&#160;19:10&#160;UTC | true | [view](CERTS/7f15a14bb0b8ad72d450e673344a747703cdbe15741d2696c371c83c5ca8e148/README.md) |
| 15&#160;Jun&#160;23&#160;20:10&#160;UTC | SHAKEN 505J | 13&#160;Sep&#160;23&#160;20:10&#160;UTC | true | [view](CERTS/81d039b45fcc19c802cbe6760a256f6b17c2ce2ebecefbf0120e6bc14093bba4/README.md) |
| 16&#160;Jun&#160;23&#160;18:51&#160;UTC | SHAKEN 857J | 15&#160;Jun&#160;24&#160;18:51&#160;UTC | true | [view](CERTS/97834968950fa42accd7ab2d9f43d12eb3dd919b5d5a72a153275419104aa534/README.md) |
| 25&#160;Jun&#160;23&#160;13:15&#160;UTC | SHAKEN 807J | 24&#160;Aug&#160;23&#160;13:15&#160;UTC | true | [view](CERTS/e35d163bfe2822eb6b9914458ea238ebc9c114c1d69533a90a72b3e9a98832ba/README.md) |
| 08&#160;Jul&#160;23&#160;21:54&#160;UTC | SHAKEN 578J | 07&#160;Jul&#160;24&#160;21:54&#160;UTC | true | [view](CERTS/5f864f10b33ee5194c0f67cd87af157b58755671a4534b5bd9570ca5eacd3f01/README.md) |
| 10&#160;Jul&#160;23&#160;11:55&#160;UTC | SHAKEN 815G | 08&#160;Oct&#160;23&#160;11:55&#160;UTC | true | [view](CERTS/c615ac7b9ffe1f47887a2cb93575c3dbb331e5bc3d4083c1292f6bca105f939d/README.md) |
| 10&#160;Jul&#160;23&#160;21:14&#160;UTC | SHAKEN 193E | 08&#160;Sep&#160;23&#160;21:14&#160;UTC | true | [view](CERTS/f84da3e31a70291f91ecc4c9821539ad88b938952453933a6e42f2068fd4191c/README.md) |
| 13&#160;Jul&#160;23&#160;20:37&#160;UTC | SHAKEN 622J | 09&#160;Jan&#160;24&#160;20:37&#160;UTC | true | [view](CERTS/15fbab040bd63fddd486a1fb0981b328f24805edcb6ab3bd4ed3aa7100b8270a/README.md) |
| 27&#160;Jul&#160;23&#160;14:43&#160;UTC | SHAKEN 366G | 26&#160;Aug&#160;23&#160;14:43&#160;UTC | true | [view](CERTS/08d9922b51b4c3c5d2f411fbe206ad914f6adab93007611a1b13df9f14498bbf/README.md) |
| 31&#160;Jul&#160;23&#160;20:03&#160;UTC | SHAKEN 577F | 30&#160;Aug&#160;23&#160;20:03&#160;UTC | true | [view](CERTS/e97acb88cd7cdfc326bad1f89666b012cac1364f26f98b2f4dc7254163d8095a/README.md) |
| 03&#160;Aug&#160;23&#160;19:21&#160;UTC | SHAKEN 952J | 02&#160;Sep&#160;23&#160;19:21&#160;UTC | true | [view](CERTS/6600984a0e3b8176b84c5134c7385fca83c51e7ae533ba8e89146b8caf2c0bd9/README.md) |
| 08&#160;Aug&#160;23&#160;15:15&#160;UTC | SHAKEN 073H | 07&#160;Aug&#160;24&#160;15:15&#160;UTC | true | [view](CERTS/a4ec097b80cea72a2d5356c16b28166d113944db1d1e6d23b6d1d1258a66d946/README.md) |
| 08&#160;Aug&#160;23&#160;17:14&#160;UTC | SHAKEN 551G | 07&#160;Sep&#160;23&#160;17:14&#160;UTC | true | [view](CERTS/ee1cffec1cff4706db55e335624e4f2c4dd330fe60611e33eecc1f5ceec5f5cb/README.md) |

#### CA Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 20&#160;Aug&#160;21&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Issuing CA3 | 19&#160;Aug&#160;31&#160;23:59&#160;UTC | true | [view](CERTS/51b51431ffd8055cd7e30c213540f4868258058d4c4c9f23f1ed8bea75101e37/README.md) |
| 20&#160;Aug&#160;21&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Issuing CA1 | 19&#160;Aug&#160;31&#160;23:59&#160;UTC | true | [view](CERTS/037a7c6621814bd5f01fa2080a77a245458694d9d566e0050e068fc853317857/README.md) |
| 20&#160;Aug&#160;21&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Root CA1 | 19&#160;Aug&#160;41&#160;23:59&#160;UTC | true | [view](CERTS/59ab943f4e8c3c17755cc8e4abb5ac65736ce74ec587c62cb744ab0babbbe2fe/README.md) |
| 24&#160;Oct&#160;22&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Issuing CA4 | 23&#160;Oct&#160;32&#160;23:59&#160;UTC | false | [view](CERTS/80dd12d935f46c256c7c1289a0834b21ec8f2f8c35f2864928d1e75f3a280665/README.md) |
| 24&#160;Oct&#160;22&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Root CA2 | 23&#160;Oct&#160;42&#160;23:59&#160;UTC | false | [view](CERTS/a26e04fc786ab70b8085236b2c53f8cfbf5d0c6a5c2c9c3e9f91669fbb8ea4d5/README.md) |


Generated: 21 Aug 23 20:18 UTC