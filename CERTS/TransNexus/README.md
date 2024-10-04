# STIR/SHAKEN CA Ecosystem Compliance

## TransNexus

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 7224 certificates were included in the corpus being tested
- 2 certificates in the corpus were skipped because they are duplicates
- 6679 certificates in the corpus were skipped because they are expired
- 457 certificates in the corpus were skipped because they are not currently trusted
- 86 certificates being tested against the remaining rules
- 1.00 issues on average found in unexpired, trusted, and non-compliant certificates
- 5.81% of certificates contain one or more Error level issue
- 0.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 75 days is the average remaining validity for the certificates in the corpus
- 74 days is the average initial validity for the certificates in the corpus
- 67 certificates expire in the next 30 days
- 1.34 average number of unexpired certificates per OCN observed
- 64 unique OCNs observed in unexpired and valid certificate corpus

| Instances | Test | Source |
|-----------|------|--------|
| 5 | [e_us_cp_subject_sn_shall](ISSUES/e_us_cp_subject_sn_shall/README.md) | US_SHAKEN_CP |

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
- 5301 days is the average remaining validity for the certificates in the corpus
- 4870 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days

No error, warning, or notice level issues were found

### Details

#### Leaf Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 06&#160;Nov&#160;23&#160;09:42&#160;UTC | SHAKEN 8526 | 05&#160;Nov&#160;24&#160;09:42&#160;UTC | true | [view](CERTS/4c86bf33be8b4189a469827d24c257723b4e5e3236981d9666d04de493b5cb6a/README.md) |
| 08&#160;Dec&#160;23&#160;21:57&#160;UTC | SHAKEN 299K | 07&#160;Dec&#160;24&#160;21:57&#160;UTC | true | [view](CERTS/7b677c8ad6481aa908931a3bec7ec5645e51770d15afca1e706b99f09203eca5/README.md) |
| 31&#160;Jan&#160;24&#160;18:59&#160;UTC | SHAKEN 873J | 30&#160;Jan&#160;25&#160;18:59&#160;UTC | true | [view](CERTS/8342db52ef1e9b25620a252b82f1378cff688a8cc0e594a6669f50ac17b34d03/README.md) |
| 01&#160;Mar&#160;24&#160;06:44&#160;UTC | SHAKEN 736J | 01&#160;Mar&#160;25&#160;06:44&#160;UTC | true | [view](CERTS/006173a1cbc2ce2a785eecfce090df4cc92990833dcae0bb0486f3f2dbf1e9c9/README.md) |
| 18&#160;Mar&#160;24&#160;16:29&#160;UTC | SHAKEN 663J | 18&#160;Mar&#160;25&#160;16:29&#160;UTC | true | [view](CERTS/3eba87a4329cf652b47f2f727feb9b01b5b3465ed41208db0e4e1190bafe9036/README.md) |
| 02&#160;Apr&#160;24&#160;15:50&#160;UTC | SHAKEN 2720 | 02&#160;Apr&#160;25&#160;15:50&#160;UTC | false | [view](CERTS/95a275a51dc842f102c43cd6b11f37b85a911ba727fca4cc6e3f9b859e05070d/README.md) |
| 30&#160;May&#160;24&#160;17:57&#160;UTC | SHAKEN 597J | 30&#160;May&#160;25&#160;17:57&#160;UTC | false | [view](CERTS/71feb09973117fb5d4aef0cfa3de26c3bfabfcb28a658609059777ae3af2d10c/README.md) |
| 06&#160;Jun&#160;24&#160;15:31&#160;UTC | SHAKEN 622J | 03&#160;Dec&#160;24&#160;15:31&#160;UTC | false | [view](CERTS/a9ccde055631fc083e64ef93fcd1baf25af1d2dfc7cedb8633e485f7699064c1/README.md) |
| 19&#160;Jun&#160;24&#160;22:07&#160;UTC | SHAKEN 505J | 16&#160;Dec&#160;24&#160;22:06&#160;UTC | false | [view](CERTS/c60b1f474c0d2900182c895719a34ade2a616a7cde7a3a3016c2d438579e8084/README.md) |
| 05&#160;Jul&#160;24&#160;16:00&#160;UTC | SHAKEN 578J | 05&#160;Jul&#160;25&#160;16:00&#160;UTC | false | [view](CERTS/b48322c4b531e0140e731d1af1acdfcf3535c2f41842655af5fb9f672fb164ee/README.md) |
| 15&#160;Jul&#160;24&#160;22:33&#160;UTC | SHAKEN 6628 | 11&#160;Jan&#160;25&#160;22:33&#160;UTC | false | [view](CERTS/c2bbe61b47c13c983b1c78857c023fbbbf0a0278ae08f58539b816814009cc60/README.md) |
| 16&#160;Jul&#160;24&#160;20:09&#160;UTC | SHAKEN 158K | 16&#160;Jul&#160;25&#160;20:08&#160;UTC | false | [view](CERTS/c6b5bf2bb2de6fa53997627a7050a90ec0f6b9786721642366f4ed5392037d64/README.md) |
| 21&#160;Jul&#160;24&#160;18:14&#160;UTC | SHAKEN 807J | 21&#160;Jul&#160;25&#160;18:14&#160;UTC | false | [view](CERTS/8e674cdfd9e11313d090b12d33deee35dceed42b64e1c751aba1a605ac1265f8/README.md) |
| 07&#160;Aug&#160;24&#160;19:30&#160;UTC | SHAKEN 193E | 06&#160;Oct&#160;24&#160;19:30&#160;UTC | false | [view](CERTS/2fa7a6597084a5ee798eab7c3d9d07f598427dca64e33004e7678e1560d448ee/README.md) |
| 08&#160;Aug&#160;24&#160;21:49&#160;UTC | SHAKEN 073H | 08&#160;Aug&#160;25&#160;21:49&#160;UTC | false | [view](CERTS/618b83ffb285c2c03f430be1e7783fc04767222fb618affd492c968e1e834035/README.md) |
| 06&#160;Sep&#160;24&#160;19:33&#160;UTC | SHAKEN 193E | 05&#160;Nov&#160;24&#160;19:33&#160;UTC | false | [view](CERTS/723b4c8404f36d63551aaf21371f5ef5cc32ef29771ba6cdfc3443e74e43dedc/README.md) |
| 11&#160;Sep&#160;24&#160;02:33&#160;UTC | SHAKEN 551G | 11&#160;Oct&#160;24&#160;02:33&#160;UTC | false | [view](CERTS/ddc9cb545e49a2988f0f40696b1131514900c8978a1653372688fcad937ad2bf/README.md) |
| 11&#160;Sep&#160;24&#160;15:51&#160;UTC | SHAKEN 8526 | 11&#160;Sep&#160;25&#160;15:51&#160;UTC | false | [view](CERTS/d6a1f4a1d2f8ce03e6dba18e0bc71b409cac0908d422f832e2571d23abfda212/README.md) |
| 14&#160;Sep&#160;24&#160;06:52&#160;UTC | SHAKEN 345J | 14&#160;Oct&#160;24&#160;06:52&#160;UTC | false | [view](CERTS/5e44d51572fa49a9d68060fb16b561d605a72db6a06ca7c5969ce23ee9d48dff/README.md) |
| 16&#160;Sep&#160;24&#160;19:01&#160;UTC | SHAKEN 706J | 16&#160;Sep&#160;25&#160;19:01&#160;UTC | false | [view](CERTS/3caad270afdbb509ba024f3862f34fa12de1f441e22bfd8f9935505f568feb0d/README.md) |
| 17&#160;Sep&#160;24&#160;22:33&#160;UTC | SHAKEN 577F | 17&#160;Oct&#160;24&#160;22:33&#160;UTC | false | [view](CERTS/3c5ba27099885e969eaf996878f91ba492d68956a3b828abadc0d087902a842f/README.md) |
| 19&#160;Sep&#160;24&#160;16:30&#160;UTC | SHAKEN 366G | 19&#160;Oct&#160;24&#160;16:30&#160;UTC | false | [view](CERTS/9774a72a8582faaddacc81b2a04b29bbc876375b9941c71e2042fa7d23ef90c3/README.md) |
| 20&#160;Sep&#160;24&#160;18:40&#160;UTC | SHAKEN 722J | 20&#160;Oct&#160;24&#160;18:40&#160;UTC | false | [view](CERTS/1bceba7fdc2643043cdbf7dca702c75b3515d972cd75136aa4295db1c56db16e/README.md) |
| 21&#160;Sep&#160;24&#160;03:53&#160;UTC | SHAKEN 952J | 21&#160;Oct&#160;24&#160;03:53&#160;UTC | false | [view](CERTS/dee6db2c5e08c61be79e9ca90ff6778b7f7e112ff85bd8743188e5e51c094a69/README.md) |
| 25&#160;Sep&#160;24&#160;04:14&#160;UTC | SHAKEN 345J | 25&#160;Oct&#160;24&#160;04:14&#160;UTC | false | [view](CERTS/70443242fd31e8c28e7239dcb6942b4079b65d7f1cd8c50df0b0ea1ac050b9a3/README.md) |
| 25&#160;Sep&#160;24&#160;05:54&#160;UTC | SHAKEN 841J | 09&#160;Oct&#160;24&#160;05:54&#160;UTC | false | [view](CERTS/6e1dc81a41886a7884f1c92ae6714881b5ee9b337cc5d01b60e9ad1e106f64af/README.md) |
| 26&#160;Sep&#160;24&#160;17:15&#160;UTC | SHAKEN 4036 | 25&#160;Mar&#160;25&#160;17:15&#160;UTC | false | [view](CERTS/a53604a9349ef1dc2fbf2f550a6fd2122790c71bcec9ce6f0f96914cd4804ded/README.md) |
| 27&#160;Sep&#160;24&#160;20:30&#160;UTC | SHAKEN 604K | 04&#160;Oct&#160;24&#160;20:30&#160;UTC | false | [view](CERTS/f602f0b32fd131d6ca8e9baf4d6b6ed9809cff0e4d04fb3ea2d3900e17e384e9/README.md) |
| 27&#160;Sep&#160;24&#160;21:28&#160;UTC | SHAKEN 733J | 04&#160;Oct&#160;24&#160;21:28&#160;UTC | false | [view](CERTS/73e1d3b3d74afb958f2e326dd0abc937d4576fcc926b9ce2e1075a7e5884a2d8/README.md) |
| 27&#160;Sep&#160;24&#160;21:28&#160;UTC | SHAKEN 856H | 04&#160;Oct&#160;24&#160;21:28&#160;UTC | false | [view](CERTS/65e47fef341c195d6f50a77167f5d21e892efc78737e59c5150910ca9c36491a/README.md) |
| 27&#160;Sep&#160;24&#160;21:30&#160;UTC | SHAKEN 819H | 04&#160;Oct&#160;24&#160;21:30&#160;UTC | false | [view](CERTS/00de0777fa221d11b070e2a325b44e40c622fd8f5f1d567b80abaad6a41377af/README.md) |
| 27&#160;Sep&#160;24&#160;21:39&#160;UTC | SHAKEN 769J | 04&#160;Oct&#160;24&#160;21:39&#160;UTC | false | [view](CERTS/b7e2d82a1d41e59d19ad0cb5b4b330819d9707866aa2441e65f5204343c30699/README.md) |
| 27&#160;Sep&#160;24&#160;21:40&#160;UTC | SHAKEN 691A | 04&#160;Oct&#160;24&#160;21:40&#160;UTC | false | [view](CERTS/bf60b09ea241a73e4e2342061e40916befacc9e6a1840738562b5bace74941b3/README.md) |
| 27&#160;Sep&#160;24&#160;21:41&#160;UTC | SHAKEN 790J | 04&#160;Oct&#160;24&#160;21:41&#160;UTC | false | [view](CERTS/d3bdab8d4a635c73f00b3b1f5c3ae5236c780ef6a379f9733f237b7b4cf3bf8d/README.md) |
| 27&#160;Sep&#160;24&#160;21:45&#160;UTC | SHAKEN 672B | 04&#160;Oct&#160;24&#160;21:45&#160;UTC | false | [view](CERTS/f6571eec7c46ca3c3d64722e38d4c7f3bb12a4b07386e00f95acb57ea05656c9/README.md) |
| 28&#160;Sep&#160;24&#160;02:33&#160;UTC | SHAKEN 625J | 05&#160;Oct&#160;24&#160;02:33&#160;UTC | false | [view](CERTS/63f915f49508f6b3190d704b82d451cabee784f3b037ea53e92918f1be3c0284/README.md) |
| 28&#160;Sep&#160;24&#160;18:49&#160;UTC | SHAKEN 6233 | 05&#160;Oct&#160;24&#160;18:49&#160;UTC | false | [view](CERTS/a926b94b9ef7dc0cee49cfe1df6e656158233013aee93a1562b79b18619b9a40/README.md) |
| 28&#160;Sep&#160;24&#160;18:57&#160;UTC | SHAKEN 927K | 05&#160;Oct&#160;24&#160;18:57&#160;UTC | false | [view](CERTS/456ce22445671df13fc3444c41a50b5153e77b0f0a249d7c55ad35ee5f6a5ab5/README.md) |
| 28&#160;Sep&#160;24&#160;19:01&#160;UTC | SHAKEN 406K | 05&#160;Oct&#160;24&#160;19:01&#160;UTC | false | [view](CERTS/469fe391186f7ba82ed038e09de10c37092432079a2dd34b9dfb22d28159da2a/README.md) |
| 28&#160;Sep&#160;24&#160;19:12&#160;UTC | SHAKEN 979K | 05&#160;Oct&#160;24&#160;19:12&#160;UTC | false | [view](CERTS/e6c11fdd9cf5a476b0c1496afc77bf9570cdd4f9d6fe9963d81ca8c04d9246c8/README.md) |
| 28&#160;Sep&#160;24&#160;19:16&#160;UTC | SHAKEN 1577 | 05&#160;Oct&#160;24&#160;19:16&#160;UTC | false | [view](CERTS/3f7ae52a8f4280c8ec0686824f1a41658a0c7bf8a92981e76647d18be8674e50/README.md) |
| 28&#160;Sep&#160;24&#160;19:16&#160;UTC | SHAKEN 390B | 05&#160;Oct&#160;24&#160;19:16&#160;UTC | false | [view](CERTS/f6a375ca943d6986bbf439f2f1c463a4167d55f04eef2770219cb1e6cc7e1ac5/README.md) |
| 28&#160;Sep&#160;24&#160;19:28&#160;UTC | SHAKEN 7820 | 05&#160;Oct&#160;24&#160;19:27&#160;UTC | false | [view](CERTS/f40f6e7a289ff8a2b785919dc132a4460e9ee247cc32b1f36c8305bda11803c0/README.md) |
| 28&#160;Sep&#160;24&#160;21:14&#160;UTC | SHAKEN 983J | 05&#160;Oct&#160;24&#160;21:14&#160;UTC | false | [view](CERTS/fa4cc4a3666cbb7a688abf1bf4cef78da90d6d7e430652c80a9772d33c704120/README.md) |
| 28&#160;Sep&#160;24&#160;21:42&#160;UTC | SHAKEN 663G | 05&#160;Oct&#160;24&#160;21:42&#160;UTC | false | [view](CERTS/a70813def27497ca82e76db024205b48bc987c4c79f4379cd5e58b09253d94a9/README.md) |
| 29&#160;Sep&#160;24&#160;12:17&#160;UTC | SHAKEN 1577 | 06&#160;Oct&#160;24&#160;12:17&#160;UTC | false | [view](CERTS/513f3dcbc86017bc4d01bcfe5ab38b568954dab33df0a3988b3afcf3ea691472/README.md) |
| 29&#160;Sep&#160;24&#160;16:46&#160;UTC | SHAKEN 646K | 06&#160;Oct&#160;24&#160;16:46&#160;UTC | false | [view](CERTS/6f9dd24a115d01958c9899f13a1cf0d0d0bfd55d36272b5b9a1fcbb3a1736906/README.md) |
| 29&#160;Sep&#160;24&#160;17:08&#160;UTC | SHAKEN 738K | 06&#160;Oct&#160;24&#160;17:08&#160;UTC | false | [view](CERTS/7b583bee8d393680a8ebb503ac2b5b6ae0133ac909206b380a71a946b980ec88/README.md) |
| 29&#160;Sep&#160;24&#160;19:19&#160;UTC | SHAKEN 0654 | 06&#160;Oct&#160;24&#160;19:19&#160;UTC | false | [view](CERTS/6c68bc1f34b84d3c6baf2aef08ee4c759b8a11f7c2643b88a24c934ac06bf360/README.md) |
| 29&#160;Sep&#160;24&#160;21:15&#160;UTC | SHAKEN 177K | 06&#160;Oct&#160;24&#160;21:15&#160;UTC | false | [view](CERTS/b644d8bb93f752f5a21da0dc2f91a766b448c304ece715e5e285253ab985a9d7/README.md) |
| 30&#160;Sep&#160;24&#160;12:02&#160;UTC | SHAKEN 690K | 07&#160;Oct&#160;24&#160;12:02&#160;UTC | false | [view](CERTS/3633f73807be95c7c07543305bef1c03404f76859adac8db7b01e6d36422275c/README.md) |
| 30&#160;Sep&#160;24&#160;21:02&#160;UTC | SHAKEN 297K | 07&#160;Oct&#160;24&#160;21:02&#160;UTC | false | [view](CERTS/b210d3baa81fb98d0becd4eb6e41fed523df2caf623bb0ccef663f212e2e00c5/README.md) |
| 30&#160;Sep&#160;24&#160;21:23&#160;UTC | SHAKEN 674J | 07&#160;Oct&#160;24&#160;21:23&#160;UTC | false | [view](CERTS/a4711fc20b60848e14fbea91796ab71b287f9b09180fe9f77340c37046653601/README.md) |
| 30&#160;Sep&#160;24&#160;21:28&#160;UTC | SHAKEN 733J | 07&#160;Oct&#160;24&#160;21:28&#160;UTC | false | [view](CERTS/0790da0bfb6a294701603ccb611b14395194a25a19386d58f04bb825880c1bf1/README.md) |
| 30&#160;Sep&#160;24&#160;21:31&#160;UTC | SHAKEN 819H | 07&#160;Oct&#160;24&#160;21:31&#160;UTC | false | [view](CERTS/5b79c3dbdc064bb5d0345bf0738372e6dde17a2317d9967eccb8f177f1288095/README.md) |
| 30&#160;Sep&#160;24&#160;21:39&#160;UTC | SHAKEN 769J | 07&#160;Oct&#160;24&#160;21:39&#160;UTC | false | [view](CERTS/cba6d6e6f9356a483e2ef3149df1ffd9a1ed48e3430f4ff1850c26fdd86f94b8/README.md) |
| 30&#160;Sep&#160;24&#160;21:40&#160;UTC | SHAKEN 1676 | 07&#160;Oct&#160;24&#160;21:40&#160;UTC | false | [view](CERTS/e32fc3cc5ed379252a5fad10213336152990b880ec44e75d25ee99257bed351f/README.md) |
| 30&#160;Sep&#160;24&#160;21:40&#160;UTC | SHAKEN 726J | 07&#160;Oct&#160;24&#160;21:40&#160;UTC | false | [view](CERTS/001f070f1f02ebcf29cc6b6ad9dfb0cf5c144badd03f588fc8d83e9019e88a4f/README.md) |
| 30&#160;Sep&#160;24&#160;21:40&#160;UTC | SHAKEN 849J | 07&#160;Oct&#160;24&#160;21:40&#160;UTC | false | [view](CERTS/bd64afad7ea679083236bf30f7d510878731cbb06b674975248416c4f68c2ed3/README.md) |
| 30&#160;Sep&#160;24&#160;21:41&#160;UTC | SHAKEN 790J | 07&#160;Oct&#160;24&#160;21:41&#160;UTC | false | [view](CERTS/ef440118e2a6e83ea183ec4303bfaa597560776db44f2a6984f5f5d50ce1b2ea/README.md) |
| 30&#160;Sep&#160;24&#160;21:45&#160;UTC | SHAKEN 0172 | 07&#160;Oct&#160;24&#160;21:44&#160;UTC | false | [view](CERTS/3ec82a60ac6cdc05fc0dc948234849b6c920164de8125aab44956f50c599661d/README.md) |
| 30&#160;Sep&#160;24&#160;21:45&#160;UTC | SHAKEN 459J | 07&#160;Oct&#160;24&#160;21:45&#160;UTC | false | [view](CERTS/c4352de9865ce41ad05d529be458778d10513fe333eae86d4e6977056769ef69/README.md) |
| 30&#160;Sep&#160;24&#160;21:45&#160;UTC | SHAKEN 672B | 07&#160;Oct&#160;24&#160;21:45&#160;UTC | false | [view](CERTS/568d292e636156fd6bd5a1392acdf857ac41d45279d26254d4c1322cd6cdf542/README.md) |
| 01&#160;Oct&#160;24&#160;10:12&#160;UTC | SHAKEN 815G | 30&#160;Dec&#160;24&#160;10:12&#160;UTC | false | [view](CERTS/47d716e167310b235ba5d20c966eee2f717ed849d331f865fa68583ab8cf9761/README.md) |
| 01&#160;Oct&#160;24&#160;15:47&#160;UTC | SHAKEN 738J | 10&#160;Oct&#160;24&#160;15:47&#160;UTC | false | [view](CERTS/52cce561e5ced05877348b525c9ad1c7e13edf07618c50ec6c3e8415cab99953/README.md) |
| 01&#160;Oct&#160;24&#160;16:57&#160;UTC | SHAKEN 602D | 08&#160;Oct&#160;24&#160;16:57&#160;UTC | false | [view](CERTS/a26d5e44d41aa79e07a7a8a053a039b66d0da9286b9bbfdafce26a54ea1e9050/README.md) |
| 01&#160;Oct&#160;24&#160;18:49&#160;UTC | SHAKEN 6233 | 08&#160;Oct&#160;24&#160;18:49&#160;UTC | false | [view](CERTS/5258447242c4c1e12887c3e0802bda70606118e87808f02c6e442210e651ff5d/README.md) |
| 01&#160;Oct&#160;24&#160;18:57&#160;UTC | SHAKEN 927K | 08&#160;Oct&#160;24&#160;18:57&#160;UTC | false | [view](CERTS/a39f8468b04dbb43c6eab31df40ac30662bd81ffc27abd6282ac8e96a825120c/README.md) |
| 01&#160;Oct&#160;24&#160;19:02&#160;UTC | SHAKEN 406K | 08&#160;Oct&#160;24&#160;19:02&#160;UTC | false | [view](CERTS/ed8d635cea2d2349ae51d58075633043744e3955f8e91c98226661b2e0e10244/README.md) |
| 01&#160;Oct&#160;24&#160;21:14&#160;UTC | SHAKEN 983J | 08&#160;Oct&#160;24&#160;21:14&#160;UTC | false | [view](CERTS/79c8d3241f6d9bd43aa0d15810121f003f88d13a401dd3cc537145e67f9c978a/README.md) |
| 01&#160;Oct&#160;24&#160;21:43&#160;UTC | SHAKEN 663G | 08&#160;Oct&#160;24&#160;21:43&#160;UTC | false | [view](CERTS/c320d3624523998123822962a8f6d9d4ffba4e29f05a690c1dc22ad99fa01e05/README.md) |
| 02&#160;Oct&#160;24&#160;02:40&#160;UTC | SHAKEN 551G | 01&#160;Nov&#160;24&#160;02:40&#160;UTC | false | [view](CERTS/1024798d6ccadb560d1896d30fe803937039cdaebbe29fa5825851862d9bfca4/README.md) |
| 02&#160;Oct&#160;24&#160;06:04&#160;UTC | SHAKEN 278K | 09&#160;Oct&#160;24&#160;06:04&#160;UTC | false | [view](CERTS/d10e478be0d8528984af6ba2a0a582cc312bc310ad4167836bd1b8757812ebd4/README.md) |
| 02&#160;Oct&#160;24&#160;06:04&#160;UTC | SHAKEN 029K | 09&#160;Oct&#160;24&#160;06:04&#160;UTC | false | [view](CERTS/96b3632a2c58a18c20ab8f18a31e2c62b38b6af88d425fbdd0c5c1630332f1a0/README.md) |
| 02&#160;Oct&#160;24&#160;06:04&#160;UTC | SHAKEN 120K | 09&#160;Oct&#160;24&#160;06:04&#160;UTC | false | [view](CERTS/aa32c483dda5b57fd114b10339e834e9021f0da4b46d1a5d6bf8f244d302ce5d/README.md) |
| 02&#160;Oct&#160;24&#160;07:21&#160;UTC | SHAKEN 625J | 09&#160;Oct&#160;24&#160;07:21&#160;UTC | false | [view](CERTS/6c3ead9891bab35a553ff77aea8d7930b317ea980ce0f5852c0e4fe4dfaa11a3/README.md) |
| 02&#160;Oct&#160;24&#160;12:17&#160;UTC | SHAKEN 1577 | 09&#160;Oct&#160;24&#160;12:17&#160;UTC | false | [view](CERTS/b7daf125bae6239e228b5c30417d7903e1ba58876f1c9a7e7bcfcece0a33c50e/README.md) |
| 02&#160;Oct&#160;24&#160;13:21&#160;UTC | SHAKEN 691J | 09&#160;Oct&#160;24&#160;13:21&#160;UTC | false | [view](CERTS/47aeb73be2b5bfb72d12e2992b32afd355eb221b0e3b540cfa3210ea3cc0949b/README.md) |
| 02&#160;Oct&#160;24&#160;15:28&#160;UTC | SHAKEN 2550 | 09&#160;Oct&#160;24&#160;15:28&#160;UTC | false | [view](CERTS/32bfd7c434951bfef2b1774c94eafaa18f299cabe28175568612633171496b8a/README.md) |
| 02&#160;Oct&#160;24&#160;16:47&#160;UTC | SHAKEN 646K | 09&#160;Oct&#160;24&#160;16:47&#160;UTC | false | [view](CERTS/f2a1e2d73422e0ed1497d2181c10a809591a28cb830ab70f234da9f4fb62f235/README.md) |
| 02&#160;Oct&#160;24&#160;17:08&#160;UTC | SHAKEN 738K | 09&#160;Oct&#160;24&#160;17:08&#160;UTC | false | [view](CERTS/6bc5d4b04383e168d9625ed4569a0bfb0f7f66e7b6d82a96dd9d1f2349b1e446/README.md) |
| 02&#160;Oct&#160;24&#160;18:21&#160;UTC | SHAKEN 833J | 09&#160;Oct&#160;24&#160;18:21&#160;UTC | false | [view](CERTS/c3ac585dd9f476ea57924aa9d0dfed7b37fa3f13f0b472b137871e0ac1edd3db/README.md) |
| 02&#160;Oct&#160;24&#160;21:15&#160;UTC | SHAKEN 177K | 09&#160;Oct&#160;24&#160;21:15&#160;UTC | false | [view](CERTS/cf1712457c7a2e9fb6e16e3e071660cd68fd927677bac2e012025faefff85684/README.md) |
| 03&#160;Oct&#160;24&#160;15:35&#160;UTC | SHAKEN 878K | 10&#160;Oct&#160;24&#160;15:35&#160;UTC | false | [view](CERTS/ecfd745fbd137130509733efd3bca85292b4d3583ada2f1d3d83e34483cec190/README.md) |
| 03&#160;Oct&#160;24&#160;15:43&#160;UTC | SHAKEN 841J | 17&#160;Oct&#160;24&#160;15:43&#160;UTC | false | [view](CERTS/6ef50507901c55b4864bbf3374182b13e8d4bc84e568a50faadb5276e0f57873/README.md) |
| 03&#160;Oct&#160;24&#160;21:41&#160;UTC | SHAKEN 790J | 10&#160;Oct&#160;24&#160;21:41&#160;UTC | false | [view](CERTS/68d1e2ffa2c6a784414c6976aaed7045eb52f9d0067561b99a8dec4def90b0fb/README.md) |

#### CA Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 24&#160;Oct&#160;22&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Issuing CA4 | 23&#160;Oct&#160;32&#160;23:59&#160;UTC | false | [view](CERTS/80dd12d935f46c256c7c1289a0834b21ec8f2f8c35f2864928d1e75f3a280665/README.md) |
| 24&#160;Oct&#160;22&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Root CA2 | 23&#160;Oct&#160;42&#160;23:59&#160;UTC | false | [view](CERTS/a26e04fc786ab70b8085236b2c53f8cfbf5d0c6a5c2c9c3e9f91669fbb8ea4d5/README.md) |
| 21&#160;Mar&#160;24&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Issuing CA5 | 20&#160;Mar&#160;34&#160;23:59&#160;UTC | false | [view](CERTS/cd50eeb8c083af686a49964a10b030048b800530edbeee8f0991388c3a79e75a/README.md) |


Generated: 04 Oct 24 16:29 UTC