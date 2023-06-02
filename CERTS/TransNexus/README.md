# STIR/SHAKEN CA Ecosystem Compliance

## TransNexus

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 2755 certificates were included in the corpus being tested
- 1 certificates in the corpus were skipped because they are duplicates
- 2666 certificates in the corpus were skipped because they are expired
- 1 certificates in the corpus were skipped because they are not currently trusted
- 87 certificates being tested against the remaining rules
- 1.18 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 9.20% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 78 days is the average remaining validity for the certificates in the corpus
- 78 days is the average initial validity for the certificates in the corpus
- 74 certificates expire in the next 30 days
- 1.40 average number of unexpired certificates per OCN observed
- 62 unique OCNs observed in unexpired and valid certificate corpus

| Instances | Test | Source |
|-----------|------|--------|
| 8 | [e_incorrect_ku_encoding](ISSUES/e_incorrect_ku_encoding/README.md) | RFC5280 |
| 87 | [e_us_cp_subject_sn](ISSUES/e_us_cp_subject_sn/README.md) | US_SHAKEN_CP |
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
- 5442 days is the average remaining validity for the certificates in the corpus
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
| 10&#160;Jun&#160;22&#160;14:00&#160;UTC | SHAKEN 597J | 10&#160;Jun&#160;23&#160;14:00&#160;UTC | true | [view](CERTS/4baa838eb96ce358aa1565d0add8b75072258efd5d21a5bbb0eec878026a67be/README.md) |
| 20&#160;Jun&#160;22&#160;20:25&#160;UTC | SHAKEN 857J | 20&#160;Jun&#160;23&#160;20:25&#160;UTC | true | [view](CERTS/e1e0a4851fb7b51d617959c00fb6a17e7b23d34e1555ba70ab01c9182a2f85f8/README.md) |
| 22&#160;Jun&#160;22&#160;18:46&#160;UTC | SHAKEN 755J | 22&#160;Jun&#160;23&#160;18:46&#160;UTC | true | [view](CERTS/a8cac316fddf25b48b9350a325791be8b87e6484a5959ab605db1ae5e3e99a84/README.md) |
| 29&#160;Jun&#160;22&#160;20:24&#160;UTC | SHAKEN 736J | 29&#160;Jun&#160;23&#160;20:24&#160;UTC | true | [view](CERTS/71b25a480077816998e48e6a76f272e70aa91ecb19b0e773cd8777823cdead32/README.md) |
| 20&#160;Jul&#160;22&#160;17:00&#160;UTC | SHAKEN 737J | 20&#160;Jul&#160;23&#160;17:00&#160;UTC | true | [view](CERTS/23db796295ad1249fcc3c9a42417c22f38be134d3b3b30082d35eab7cc153e5d/README.md) |
| 25&#160;Jul&#160;22&#160;18:36&#160;UTC | SHAKEN 578J | 25&#160;Jul&#160;23&#160;18:36&#160;UTC | true | [view](CERTS/f4af5a28131d680ef2682ae86181b64f0c919864d7dbaef00c0598d314b537a1/README.md) |
| 10&#160;Aug&#160;22&#160;18:11&#160;UTC | SHAKEN 073H | 10&#160;Aug&#160;23&#160;18:11&#160;UTC | true | [view](CERTS/87dabb91a689a6ac540b5ffbf5eb89326fada83fb822d2bb9ec31b68251b579e/README.md) |
| 13&#160;Sep&#160;22&#160;17:24&#160;UTC | SHAKEN 706J | 13&#160;Sep&#160;23&#160;17:24&#160;UTC | true | [view](CERTS/1eea5e17ae8b36cb5b6d657178273af745cea015e63586e6d44d1ba2ee8f722a/README.md) |
| 07&#160;Dec&#160;22&#160;00:45&#160;UTC | SHAKEN 299K | 07&#160;Dec&#160;23&#160;00:45&#160;UTC | true | [view](CERTS/7b5e04ff01ace2ee384df3507886684204901bf23851866eb6b3891ecdd96a92/README.md) |
| 31&#160;Dec&#160;22&#160;03:33&#160;UTC | SHAKEN 8526 | 31&#160;Dec&#160;23&#160;03:33&#160;UTC | true | [view](CERTS/96672cb8cc7c71e38d9e6350eac6c4f5c430c1eefeabe00cc095527e20a20c9b/README.md) |
| 03&#160;Jan&#160;23&#160;16:07&#160;UTC | SHAKEN 6628 | 03&#160;Jun&#160;23&#160;16:07&#160;UTC | true | [view](CERTS/dec9a2b7e2fce8ca94e2b1313886772d8176d6adcac7c4cb2b295315fa79f5ab/README.md) |
| 16&#160;Jan&#160;23&#160;17:54&#160;UTC | SHAKEN 622J | 15&#160;Jul&#160;23&#160;17:54&#160;UTC | true | [view](CERTS/4c659921cd1ec39c239ab31d4c1ca813017c01d51512ea11a1fe0572f034c534/README.md) |
| 25&#160;Jan&#160;23&#160;14:28&#160;UTC | SHAKEN 873J | 25&#160;Jan&#160;24&#160;14:28&#160;UTC | true | [view](CERTS/36e9ba32d337ad13909bee0b62125c9769a049c4ae10d749687d75af4e49387d/README.md) |
| 20&#160;Mar&#160;23&#160;18:51&#160;UTC | SHAKEN 505J | 18&#160;Jun&#160;23&#160;18:51&#160;UTC | true | [view](CERTS/aa04536bc3bbc9914c2b27653d1cf862dfcb571b53c784e83942afff6598e11d/README.md) |
| 20&#160;Mar&#160;23&#160;18:53&#160;UTC | SHAKEN 505J | 18&#160;Jun&#160;23&#160;18:53&#160;UTC | true | [view](CERTS/9a4feceff25f99e7d04d1542ef56bba8f060987c5b6734239f185997c5e58cf1/README.md) |
| 27&#160;Mar&#160;23&#160;17:18&#160;UTC | SHAKEN 2720 | 26&#160;Mar&#160;24&#160;17:18&#160;UTC | true | [view](CERTS/7282c0683ebff27881b1ccd1f66664c3102340ff657d9363fcecb9c1b7924c12/README.md) |
| 08&#160;Apr&#160;23&#160;18:53&#160;UTC | SHAKEN 193E | 07&#160;Jun&#160;23&#160;18:53&#160;UTC | true | [view](CERTS/8cc7b08b6c144bae17ff71317ebea1877859c2bbdd387c6b9083a130bf7ab18c/README.md) |
| 10&#160;Apr&#160;23&#160;12:54&#160;UTC | SHAKEN 815G | 09&#160;Jul&#160;23&#160;12:54&#160;UTC | true | [view](CERTS/71206dbb241c8dbf9de187fca7b7e515fff61fbedcd1baacdc4a6a399ca2f887/README.md) |
| 27&#160;Apr&#160;23&#160;17:53&#160;UTC | SHAKEN 807J | 26&#160;Jun&#160;23&#160;17:53&#160;UTC | true | [view](CERTS/ebff6737c13e98d15dbabf1df2676994b4033eeaf0abbd64136795f8954d36f0/README.md) |
| 28&#160;Apr&#160;23&#160;17:50&#160;UTC | SHAKEN 663J | 27&#160;Apr&#160;24&#160;17:50&#160;UTC | true | [view](CERTS/bd944d8e1acfeaaa520a8c87826b19a8509b52b083525a940e9ee4db8af1a99b/README.md) |
| 04&#160;May&#160;23&#160;14:22&#160;UTC | SHAKEN 366G | 03&#160;Jun&#160;23&#160;14:22&#160;UTC | true | [view](CERTS/2cfa43117f29f6425f005ea2ef1da616405e64a02a499d985bfe60a588b38d16/README.md) |
| 05&#160;May&#160;23&#160;16:22&#160;UTC | SHAKEN 722J | 04&#160;Jun&#160;23&#160;16:22&#160;UTC | true | [view](CERTS/b964f6aa5b042cf76920981b31eaecbe4db21d55eebaea8a4441e18bf5fd7857/README.md) |
| 05&#160;May&#160;23&#160;18:54&#160;UTC | SHAKEN 952J | 04&#160;Jun&#160;23&#160;18:54&#160;UTC | true | [view](CERTS/1702e4fb7ef8657467e391ba6ef702ae74877371cd816e0d65eb98a2aa0c2626/README.md) |
| 08&#160;May&#160;23&#160;17:37&#160;UTC | SHAKEN 4036 | 04&#160;Nov&#160;23&#160;17:37&#160;UTC | true | [view](CERTS/1f7c5f89be3bc2a774c0b657a342d25a23ec7f07e38c8e3ed1d45b720641d919/README.md) |
| 10&#160;May&#160;23&#160;23:24&#160;UTC | SHAKEN 4632 | 09&#160;Jun&#160;23&#160;23:24&#160;UTC | true | [view](CERTS/bc967e0070f63cc0ce7848dfd7d6e726a600ef4758680c3422168a2201e1d48a/README.md) |
| 13&#160;May&#160;23&#160;22:12&#160;UTC | SHAKEN 186K | 12&#160;Jun&#160;23&#160;22:12&#160;UTC | true | [view](CERTS/eac0052184c5747a3c07ea1a48c019ba0c6d3ed3433a0409705f9903e0ca406c/README.md) |
| 16&#160;May&#160;23&#160;17:02&#160;UTC | SHAKEN 551G | 15&#160;Jun&#160;23&#160;17:02&#160;UTC | true | [view](CERTS/60a5f9798d01960d0025e3fd706bfc475bf43089053b49f4c7762042839af8df/README.md) |
| 20&#160;May&#160;23&#160;04:53&#160;UTC | SHAKEN 345J | 19&#160;Jun&#160;23&#160;04:53&#160;UTC | true | [view](CERTS/0a4f91830ffd51e03595b626d1adde9e3b162384ba0ed2279d32de2d4c65fc7c/README.md) |
| 20&#160;May&#160;23&#160;05:23&#160;UTC | SHAKEN 345J | 19&#160;Jun&#160;23&#160;05:23&#160;UTC | true | [view](CERTS/4cdac9e950272e8f498966ad69bdc1ccbda340b97a381d6a4745763205926481/README.md) |
| 20&#160;May&#160;23&#160;19:31&#160;UTC | SHAKEN 577F | 19&#160;Jun&#160;23&#160;19:31&#160;UTC | true | [view](CERTS/b5927ac34a8810ccc6f62c3d3aedec71ea2e5025850e9055c333d576c6fa7838/README.md) |
| 23&#160;May&#160;23&#160;19:01&#160;UTC | SHAKEN 952J | 22&#160;Jun&#160;23&#160;19:01&#160;UTC | true | [view](CERTS/c1a9a4851e79d3179cf3c96608ea854f532b1f825460ff62a188a34e3a467176/README.md) |
| 24&#160;May&#160;23&#160;08:31&#160;UTC | SHAKEN 736J | 23&#160;May&#160;24&#160;08:31&#160;UTC | true | [view](CERTS/d273815c57e441d2066144a02d029b73c3d839f5d31cbaeb72a7b6bded62da53/README.md) |
| 25&#160;May&#160;23&#160;14:30&#160;UTC | SHAKEN 366G | 24&#160;Jun&#160;23&#160;14:30&#160;UTC | true | [view](CERTS/34b446bfe8c1285257f0397a4fedf8dc0266694b4b3a773fbc892bb4cbd7d2b7/README.md) |
| 26&#160;May&#160;23&#160;01:56&#160;UTC | SHAKEN 278K | 02&#160;Jun&#160;23&#160;01:56&#160;UTC | true | [view](CERTS/0756f89b00467abc9783008a21297c1ae0e02651b202adcad86ff5b31acb9415/README.md) |
| 26&#160;May&#160;23&#160;06:04&#160;UTC | SHAKEN 0172 | 02&#160;Jun&#160;23&#160;06:04&#160;UTC | true | [view](CERTS/3e9d35fc61b030404d3547cce10ea7b878ee41b1b78e96ab66c43c3c1db6899f/README.md) |
| 26&#160;May&#160;23&#160;14:51&#160;UTC | SHAKEN 2550 | 02&#160;Jun&#160;23&#160;14:51&#160;UTC | true | [view](CERTS/85975ebb4b4e7f8632a2ef4139884dd003f28fa6428b2d43bbf27d7d34130409/README.md) |
| 26&#160;May&#160;23&#160;20:30&#160;UTC | SHAKEN 841J | 09&#160;Jun&#160;23&#160;20:30&#160;UTC | true | [view](CERTS/f067fcbe89bd3ed916b721c0d5b4332d45a84c8d2eb9c9f558049a48403a4532/README.md) |
| 27&#160;May&#160;23&#160;17:44&#160;UTC | SHAKEN 107K | 03&#160;Jun&#160;23&#160;17:44&#160;UTC | true | [view](CERTS/f5811dae9edc690efd4317c4777925b9d3241e963d37bf1473241685e47f5422/README.md) |
| 27&#160;May&#160;23&#160;18:55&#160;UTC | SHAKEN 089K | 03&#160;Jun&#160;23&#160;18:55&#160;UTC | true | [view](CERTS/d4511b666b9e67b0a9a7e2aee0d72a3b220cb004c55f98e037de130923ce36c4/README.md) |
| 27&#160;May&#160;23&#160;20:04&#160;UTC | SHAKEN 159H | 03&#160;Jun&#160;23&#160;20:04&#160;UTC | true | [view](CERTS/1fa3a4d6b08340a52ecb0161b10d45bba1806f0d31428a4022d1d9d848ddda97/README.md) |
| 27&#160;May&#160;23&#160;20:13&#160;UTC | SHAKEN 297K | 03&#160;Jun&#160;23&#160;20:13&#160;UTC | true | [view](CERTS/911247efcad937ce24b08f47d8a4abb3e62561f1340345b47eabafda48d79cba/README.md) |
| 27&#160;May&#160;23&#160;20:36&#160;UTC | SHAKEN 674J | 03&#160;Jun&#160;23&#160;20:36&#160;UTC | true | [view](CERTS/094c6e87f9d8c3e9ba56a8f33778529da52280c05dd82417011b311738d41565/README.md) |
| 27&#160;May&#160;23&#160;20:37&#160;UTC | SHAKEN 735J | 03&#160;Jun&#160;23&#160;20:37&#160;UTC | true | [view](CERTS/c8157778d5aad091e39f36e7ed443842ae4349eaf90d01e8346b57b139d92ae0/README.md) |
| 27&#160;May&#160;23&#160;20:39&#160;UTC | SHAKEN 738J | 03&#160;Jun&#160;23&#160;20:39&#160;UTC | true | [view](CERTS/7b719e24d334841f9ae2b4cb3f92b6322e8bb6692ee29bbdc592dd0a68904154/README.md) |
| 27&#160;May&#160;23&#160;20:43&#160;UTC | SHAKEN 053G | 03&#160;Jun&#160;23&#160;20:43&#160;UTC | true | [view](CERTS/e9cee3b229b692a95ecfb19ad973886904e37ac186b2f005ae02c668e47802a2/README.md) |
| 27&#160;May&#160;23&#160;20:43&#160;UTC | SHAKEN 769J | 03&#160;Jun&#160;23&#160;20:43&#160;UTC | true | [view](CERTS/ccb9ec238fe5b7a13de79a9ee2d7bd7ff5a258892e5b8cb2fbf82bd3f73429cf/README.md) |
| 27&#160;May&#160;23&#160;20:45&#160;UTC | SHAKEN 469A | 03&#160;Jun&#160;23&#160;20:45&#160;UTC | true | [view](CERTS/108cf3f4670a5db16ab3f168645ee5fd64faa8b7c1e3866e977dd1a37c7bc17c/README.md) |
| 27&#160;May&#160;23&#160;20:45&#160;UTC | SHAKEN 726J | 03&#160;Jun&#160;23&#160;20:45&#160;UTC | true | [view](CERTS/0380c6ea078c2e7a3bfb97be7a37b4d02e435bcbf9a2a2dc9fd518ff5249e646/README.md) |
| 27&#160;May&#160;23&#160;20:46&#160;UTC | SHAKEN 738J | 03&#160;Jun&#160;23&#160;20:46&#160;UTC | true | [view](CERTS/0b05901ccff6f9b47bb049fa6ae3b33868848877c43cb7cae372891dc2b40499/README.md) |
| 27&#160;May&#160;23&#160;20:46&#160;UTC | SHAKEN 459J | 03&#160;Jun&#160;23&#160;20:46&#160;UTC | true | [view](CERTS/28bc76cfb4fe960672d2a3d2edfded91755034c9160418b6367c8eb0e70e94fd/README.md) |
| 27&#160;May&#160;23&#160;20:47&#160;UTC | SHAKEN 366G | 03&#160;Jun&#160;23&#160;20:47&#160;UTC | true | [view](CERTS/150d4476979581bdd9e3db8456cbd18eacfd58e793ba059c2df0c74e6591b774/README.md) |
| 27&#160;May&#160;23&#160;20:47&#160;UTC | SHAKEN 0226 | 03&#160;Jun&#160;23&#160;20:47&#160;UTC | true | [view](CERTS/3df43a4015f471471e32fb8e929b928dc3a9b3247bbc03b8caafbe63b53301bb/README.md) |
| 27&#160;May&#160;23&#160;20:47&#160;UTC | SHAKEN 738J | 03&#160;Jun&#160;23&#160;20:47&#160;UTC | true | [view](CERTS/383e545b0cf0e5892e76e4cce9824a5063d2d3dcaa32236a7a939d10dadea78a/README.md) |
| 28&#160;May&#160;23&#160;03:37&#160;UTC | SHAKEN 9472 | 04&#160;Jun&#160;23&#160;03:37&#160;UTC | true | [view](CERTS/714a8be7cfd797acbd5534d6fb8e87cd8d52102fe320d365a9633ab8d3fefdd5/README.md) |
| 28&#160;May&#160;23&#160;18:21&#160;UTC | SHAKEN 406K | 04&#160;Jun&#160;23&#160;18:21&#160;UTC | true | [view](CERTS/03bc696c68604bd4506464fa7a54acca3609f71e09545916b8f64e77720081b2/README.md) |
| 28&#160;May&#160;23&#160;20:33&#160;UTC | SHAKEN 983J | 04&#160;Jun&#160;23&#160;20:33&#160;UTC | true | [view](CERTS/d550f645ef25e44323cac8b662fc53435061320dabb548f28f3ba8530b4395e5/README.md) |
| 28&#160;May&#160;23&#160;22:03&#160;UTC | SHAKEN 029K | 04&#160;Jun&#160;23&#160;22:02&#160;UTC | true | [view](CERTS/9aa8a74cacbe860d9c9ccfa05c301c08c5258f69af890a6165d7364a08e77de6/README.md) |
| 28&#160;May&#160;23&#160;22:12&#160;UTC | SHAKEN 2278 | 04&#160;Jun&#160;23&#160;22:12&#160;UTC | true | [view](CERTS/aa492735d65a97775cec6d461bea1e984ab8c6e209d1c146948cde599d47a89a/README.md) |
| 28&#160;May&#160;23&#160;22:42&#160;UTC | SHAKEN 120K | 04&#160;Jun&#160;23&#160;22:42&#160;UTC | true | [view](CERTS/3b0e1d2b2202c0fd3bc689b0d2a3d4fcf8ab8cbd6d14d60373fb10ddf896f1b1/README.md) |
| 28&#160;May&#160;23&#160;23:31&#160;UTC | SHAKEN 4632 | 27&#160;Jun&#160;23&#160;23:31&#160;UTC | true | [view](CERTS/06a35beaaecaa1289e36d54e20a2e2b7b73e23d10d0e1138e90d311544a4a247/README.md) |
| 29&#160;May&#160;23&#160;01:56&#160;UTC | SHAKEN 278K | 05&#160;Jun&#160;23&#160;01:56&#160;UTC | true | [view](CERTS/f1da3d9b567dd92b1f494c3436bfa2e2a38a4f55b405dff040ec48d42441b3bc/README.md) |
| 29&#160;May&#160;23&#160;14:52&#160;UTC | SHAKEN 2550 | 05&#160;Jun&#160;23&#160;14:52&#160;UTC | true | [view](CERTS/6d11d4bffef9b13e338ea0ef37458bf3c1d0b48a966c9e64f897b9bb03b8034e/README.md) |
| 29&#160;May&#160;23&#160;15:38&#160;UTC | SHAKEN 284K | 05&#160;Jun&#160;23&#160;15:38&#160;UTC | true | [view](CERTS/9f771c12d09097e4163e59187d7fbcc127ba6efee43cbc6f0c008ca57fd98d9f/README.md) |
| 29&#160;May&#160;23&#160;16:34&#160;UTC | SHAKEN 722J | 28&#160;Jun&#160;23&#160;16:34&#160;UTC | true | [view](CERTS/16719a4c0173079651642171f9d1c1daa69811d122cc78cb5af8aabe540dbd0f/README.md) |
| 29&#160;May&#160;23&#160;20:31&#160;UTC | SHAKEN 177K | 05&#160;Jun&#160;23&#160;20:31&#160;UTC | true | [view](CERTS/bc19b8d812e492f2bbc3d062548b9c5f7584630beddf31263b58683f28dbec84/README.md) |
| 29&#160;May&#160;23&#160;20:51&#160;UTC | SHAKEN 738J | 07&#160;Jun&#160;23&#160;20:51&#160;UTC | true | [view](CERTS/5cbd06ff7bef53fac02aa0b4e4506feda23e48e30da23aeefbf042667087c691/README.md) |
| 30&#160;May&#160;23&#160;10:53&#160;UTC | SHAKEN 0172 | 06&#160;Jun&#160;23&#160;10:53&#160;UTC | true | [view](CERTS/7915601d479e970aeafe82ebec168f47cbdfb005e6ad391a7bf9e2b80b707ad0/README.md) |
| 30&#160;May&#160;23&#160;12:25&#160;UTC | SHAKEN 1577 | 06&#160;Jun&#160;23&#160;12:25&#160;UTC | true | [view](CERTS/d4f26d63b375325b626d6be264c6cb982ba807c5b36fc057d3d5580d6ab81aa8/README.md) |
| 30&#160;May&#160;23&#160;14:57&#160;UTC | SHAKEN 012K | 06&#160;Jun&#160;23&#160;14:57&#160;UTC | true | [view](CERTS/8b72415be247536f4a3a002e4a9c0ee591ce956adc5526cd19d71ae8956f9e29/README.md) |
| 30&#160;May&#160;23&#160;17:44&#160;UTC | SHAKEN 107K | 06&#160;Jun&#160;23&#160;17:44&#160;UTC | true | [view](CERTS/2f63af94f77d9641f68b0170bf7284906300afcca7a9e576e0c0afb56799c2dc/README.md) |
| 30&#160;May&#160;23&#160;18:07&#160;UTC | SHAKEN 0360 | 06&#160;Jun&#160;23&#160;18:07&#160;UTC | true | [view](CERTS/d033ca3060d474199d71ab2a0c82636e08d246677418e1e4ceff834e40d3b89f/README.md) |
| 30&#160;May&#160;23&#160;18:55&#160;UTC | SHAKEN 056K | 06&#160;Jun&#160;23&#160;18:55&#160;UTC | true | [view](CERTS/4dd88905fad3402af24bc0fcdd552a3b1e7e2570cc6b297599ae32e74e5a5308/README.md) |
| 30&#160;May&#160;23&#160;19:23&#160;UTC | SHAKEN 982J | 06&#160;Jun&#160;23&#160;19:23&#160;UTC | true | [view](CERTS/7b5783440365718951f2e43e6b5d1ac93ccfb67e45853117363ec8aa55900574/README.md) |
| 30&#160;May&#160;23&#160;20:36&#160;UTC | SHAKEN 674J | 06&#160;Jun&#160;23&#160;20:36&#160;UTC | true | [view](CERTS/046e6b5469240c1481bfc5049d346895bcd4800e1206203a20f7f06d803edd49/README.md) |
| 30&#160;May&#160;23&#160;20:40&#160;UTC | SHAKEN 738J | 06&#160;Jun&#160;23&#160;20:40&#160;UTC | true | [view](CERTS/87b673d82f1f3b7109158dd0158ce67ec1b8acc20d180500ccb7886b04ad69f4/README.md) |
| 30&#160;May&#160;23&#160;20:41&#160;UTC | SHAKEN 700H | 06&#160;Jun&#160;23&#160;20:41&#160;UTC | true | [view](CERTS/c78fc4deb95cb2b5ab11b185271b80cd287228cf9a3758e19e81221bd0ce56cb/README.md) |
| 30&#160;May&#160;23&#160;20:42&#160;UTC | SHAKEN 819H | 06&#160;Jun&#160;23&#160;20:42&#160;UTC | true | [view](CERTS/11eb72fa305788da2ee2a8b6b5a349f561115149a3c381e31885e9a469e46225/README.md) |
| 30&#160;May&#160;23&#160;20:43&#160;UTC | SHAKEN 053G | 06&#160;Jun&#160;23&#160;20:43&#160;UTC | true | [view](CERTS/cf0e79e29604d37e60a2ad9830ca01ac17ab3140b4ab24cbc0d31a0c0f17a6e0/README.md) |
| 30&#160;May&#160;23&#160;20:45&#160;UTC | SHAKEN 849J | 06&#160;Jun&#160;23&#160;20:45&#160;UTC | true | [view](CERTS/71af65f58aec0112dd6e793c8dd2cd11754cbb84b29af2cfab4a71168b1a7d26/README.md) |
| 30&#160;May&#160;23&#160;20:45&#160;UTC | SHAKEN 469A | 06&#160;Jun&#160;23&#160;20:45&#160;UTC | true | [view](CERTS/9369c36bdf67e97309f62676661c967497a1dc8cbdaace72141ab23ee9a69c14/README.md) |
| 30&#160;May&#160;23&#160;20:45&#160;UTC | SHAKEN 790J | 06&#160;Jun&#160;23&#160;20:45&#160;UTC | true | [view](CERTS/7943809868c0b650e18e4497b0ac85cf6503c5b7533ce7c50d8b9c6f808f859e/README.md) |
| 30&#160;May&#160;23&#160;20:46&#160;UTC | SHAKEN 738J | 06&#160;Jun&#160;23&#160;20:46&#160;UTC | true | [view](CERTS/8eb5f59dfe5fe3657b1c9680d74bea5e9d5c440c6d5afde3d80d6712ed584e9f/README.md) |
| 30&#160;May&#160;23&#160;20:46&#160;UTC | SHAKEN 459J | 06&#160;Jun&#160;23&#160;20:46&#160;UTC | true | [view](CERTS/90a0cb640b6cfaf3155237f98ec2e136888401cd419d0d1bc1f33f239da39020/README.md) |
| 30&#160;May&#160;23&#160;20:48&#160;UTC | SHAKEN 366G | 06&#160;Jun&#160;23&#160;20:48&#160;UTC | true | [view](CERTS/4553a50dccdd1b56f249863cd94236d06fe1c4213b004c09898ad515541a8f93/README.md) |
| 30&#160;May&#160;23&#160;20:48&#160;UTC | SHAKEN 0226 | 06&#160;Jun&#160;23&#160;20:48&#160;UTC | true | [view](CERTS/b7cbef33560ef5c1dc077aecffccacbc6ad009295c456bdc457df57fc3c45d7b/README.md) |
| 30&#160;May&#160;23&#160;20:48&#160;UTC | SHAKEN 738J | 06&#160;Jun&#160;23&#160;20:48&#160;UTC | true | [view](CERTS/489c9c9fc5edb30f35eeab2f9bdac60568a976dd505c3d374b704166f98e8552/README.md) |
| 31&#160;May&#160;23&#160;18:21&#160;UTC | SHAKEN 406K | 07&#160;Jun&#160;23&#160;18:21&#160;UTC | true | [view](CERTS/c09bf056b8916e812a8fc046a2a360e9d93cfb83982adff87622ea2c36ae7714/README.md) |

#### CA Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 20&#160;Aug&#160;21&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Issuing CA3 | 19&#160;Aug&#160;31&#160;23:59&#160;UTC | true | [view](CERTS/51b51431ffd8055cd7e30c213540f4868258058d4c4c9f23f1ed8bea75101e37/README.md) |
| 20&#160;Aug&#160;21&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Issuing CA1 | 19&#160;Aug&#160;31&#160;23:59&#160;UTC | true | [view](CERTS/037a7c6621814bd5f01fa2080a77a245458694d9d566e0050e068fc853317857/README.md) |
| 20&#160;Aug&#160;21&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Root CA1 | 19&#160;Aug&#160;41&#160;23:59&#160;UTC | true | [view](CERTS/59ab943f4e8c3c17755cc8e4abb5ac65736ce74ec587c62cb744ab0babbbe2fe/README.md) |
| 24&#160;Oct&#160;22&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Issuing CA4 | 23&#160;Oct&#160;32&#160;23:59&#160;UTC | false | [view](CERTS/80dd12d935f46c256c7c1289a0834b21ec8f2f8c35f2864928d1e75f3a280665/README.md) |
| 24&#160;Oct&#160;22&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Root CA2 | 23&#160;Oct&#160;42&#160;23:59&#160;UTC | false | [view](CERTS/a26e04fc786ab70b8085236b2c53f8cfbf5d0c6a5c2c9c3e9f91669fbb8ea4d5/README.md) |


Generated: 02 Jun 23 01:12 UTC