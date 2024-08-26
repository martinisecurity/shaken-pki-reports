# STIR/SHAKEN CA Ecosystem Compliance

## TransNexus

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 6777 certificates were included in the corpus being tested
- 2 certificates in the corpus were skipped because they are duplicates
- 6211 certificates in the corpus were skipped because they are expired
- 457 certificates in the corpus were skipped because they are not currently trusted
- 107 certificates being tested against the remaining rules
- 1.00 issues on average found in unexpired, trusted, and non-compliant certificates
- 6.54% of certificates contain one or more Error level issue
- 0.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 60 days is the average remaining validity for the certificates in the corpus
- 59 days is the average initial validity for the certificates in the corpus
- 88 certificates expire in the next 30 days
- 1.37 average number of unexpired certificates per OCN observed
- 78 unique OCNs observed in unexpired and valid certificate corpus

| Instances | Test | Source |
|-----------|------|--------|
| 7 | [e_us_cp_subject_sn_shall](ISSUES/e_us_cp_subject_sn_shall/README.md) | US_SHAKEN_CP |

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
- 5311 days is the average remaining validity for the certificates in the corpus
- 4870 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days

No error, warning, or notice level issues were found

### Details

#### Leaf Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 27&#160;Sep&#160;23&#160;01:49&#160;UTC | SHAKEN 706J | 26&#160;Sep&#160;24&#160;01:49&#160;UTC | true | [view](CERTS/456dbbfbbde51c8c677b0b27e2fce5752a7de5c3407b4760d6fe1a306f3d8b94/README.md) |
| 05&#160;Oct&#160;23&#160;15:41&#160;UTC | SHAKEN 219H | 04&#160;Oct&#160;24&#160;15:41&#160;UTC | true | [view](CERTS/1db70b77e5c0c181106039f4e2e44543f730b6fea9d84d7beb33d913c8f40148/README.md) |
| 06&#160;Nov&#160;23&#160;09:42&#160;UTC | SHAKEN 8526 | 05&#160;Nov&#160;24&#160;09:42&#160;UTC | true | [view](CERTS/4c86bf33be8b4189a469827d24c257723b4e5e3236981d9666d04de493b5cb6a/README.md) |
| 08&#160;Dec&#160;23&#160;21:57&#160;UTC | SHAKEN 299K | 07&#160;Dec&#160;24&#160;21:57&#160;UTC | true | [view](CERTS/7b677c8ad6481aa908931a3bec7ec5645e51770d15afca1e706b99f09203eca5/README.md) |
| 31&#160;Jan&#160;24&#160;18:59&#160;UTC | SHAKEN 873J | 30&#160;Jan&#160;25&#160;18:59&#160;UTC | true | [view](CERTS/8342db52ef1e9b25620a252b82f1378cff688a8cc0e594a6669f50ac17b34d03/README.md) |
| 01&#160;Mar&#160;24&#160;06:44&#160;UTC | SHAKEN 736J | 01&#160;Mar&#160;25&#160;06:44&#160;UTC | true | [view](CERTS/006173a1cbc2ce2a785eecfce090df4cc92990833dcae0bb0486f3f2dbf1e9c9/README.md) |
| 18&#160;Mar&#160;24&#160;16:29&#160;UTC | SHAKEN 663J | 18&#160;Mar&#160;25&#160;16:29&#160;UTC | true | [view](CERTS/3eba87a4329cf652b47f2f727feb9b01b5b3465ed41208db0e4e1190bafe9036/README.md) |
| 02&#160;Apr&#160;24&#160;15:43&#160;UTC | SHAKEN 4036 | 29&#160;Sep&#160;24&#160;15:43&#160;UTC | false | [view](CERTS/cbc83bcb04416d64b8b1715ab74bbe8e9f9875bc4036039e8600a83f7e754945/README.md) |
| 02&#160;Apr&#160;24&#160;15:50&#160;UTC | SHAKEN 2720 | 02&#160;Apr&#160;25&#160;15:50&#160;UTC | false | [view](CERTS/95a275a51dc842f102c43cd6b11f37b85a911ba727fca4cc6e3f9b859e05070d/README.md) |
| 30&#160;May&#160;24&#160;17:57&#160;UTC | SHAKEN 597J | 30&#160;May&#160;25&#160;17:57&#160;UTC | false | [view](CERTS/71feb09973117fb5d4aef0cfa3de26c3bfabfcb28a658609059777ae3af2d10c/README.md) |
| 06&#160;Jun&#160;24&#160;15:31&#160;UTC | SHAKEN 622J | 03&#160;Dec&#160;24&#160;15:31&#160;UTC | false | [view](CERTS/a9ccde055631fc083e64ef93fcd1baf25af1d2dfc7cedb8633e485f7699064c1/README.md) |
| 12&#160;Jun&#160;24&#160;20:54&#160;UTC | SHAKEN 4036 | 30&#160;Sep&#160;24&#160;20:54&#160;UTC | false | [view](CERTS/2bb50df4cebf6a72babc41cf9b4a966d78df42d5e94fff4b8ed7d2cd43155f31/README.md) |
| 19&#160;Jun&#160;24&#160;22:07&#160;UTC | SHAKEN 505J | 16&#160;Dec&#160;24&#160;22:06&#160;UTC | false | [view](CERTS/c60b1f474c0d2900182c895719a34ade2a616a7cde7a3a3016c2d438579e8084/README.md) |
| 03&#160;Jul&#160;24&#160;10:58&#160;UTC | SHAKEN 815G | 01&#160;Oct&#160;24&#160;10:58&#160;UTC | false | [view](CERTS/5c2e209fed8ec5599d4d9990e58c36da4ea68a83a413f3966b219ccec9dfc16a/README.md) |
| 05&#160;Jul&#160;24&#160;16:00&#160;UTC | SHAKEN 578J | 05&#160;Jul&#160;25&#160;16:00&#160;UTC | false | [view](CERTS/b48322c4b531e0140e731d1af1acdfcf3535c2f41842655af5fb9f672fb164ee/README.md) |
| 08&#160;Jul&#160;24&#160;19:21&#160;UTC | SHAKEN 193E | 06&#160;Sep&#160;24&#160;19:21&#160;UTC | false | [view](CERTS/629f67b910ddebf2d7167d2bb6e2a79009062360c82bb127e9e4a89276675091/README.md) |
| 15&#160;Jul&#160;24&#160;22:33&#160;UTC | SHAKEN 6628 | 11&#160;Jan&#160;25&#160;22:33&#160;UTC | false | [view](CERTS/c2bbe61b47c13c983b1c78857c023fbbbf0a0278ae08f58539b816814009cc60/README.md) |
| 21&#160;Jul&#160;24&#160;18:14&#160;UTC | SHAKEN 807J | 21&#160;Jul&#160;25&#160;18:14&#160;UTC | false | [view](CERTS/8e674cdfd9e11313d090b12d33deee35dceed42b64e1c751aba1a605ac1265f8/README.md) |
| 29&#160;Jul&#160;24&#160;03:42&#160;UTC | SHAKEN 952J | 28&#160;Aug&#160;24&#160;03:42&#160;UTC | false | [view](CERTS/04f1d8050f99731d88f9101f4c411501c0bebbe832e431c7a311d43b68cf4b6f/README.md) |
| 31&#160;Jul&#160;24&#160;02:23&#160;UTC | SHAKEN 551G | 30&#160;Aug&#160;24&#160;02:23&#160;UTC | false | [view](CERTS/d4210a5c1f2b540cea1c5d70a53ea25698b8ab58197c6e771ed42f4c4d435a4c/README.md) |
| 03&#160;Aug&#160;24&#160;06:51&#160;UTC | SHAKEN 345J | 02&#160;Sep&#160;24&#160;06:51&#160;UTC | false | [view](CERTS/780c89a5a2128e9fe54a8bdac56f786eeec4d252ecaa4df1aa17c108a82f8920/README.md) |
| 03&#160;Aug&#160;24&#160;18:24&#160;UTC | SHAKEN 722J | 02&#160;Sep&#160;24&#160;18:24&#160;UTC | false | [view](CERTS/f5399dbfa8f5330d23e1f7052de58c32c224b6c346b974a8ca31cd376efd19fa/README.md) |
| 07&#160;Aug&#160;24&#160;19:30&#160;UTC | SHAKEN 193E | 06&#160;Oct&#160;24&#160;19:30&#160;UTC | false | [view](CERTS/2fa7a6597084a5ee798eab7c3d9d07f598427dca64e33004e7678e1560d448ee/README.md) |
| 08&#160;Aug&#160;24&#160;16:13&#160;UTC | SHAKEN 366G | 07&#160;Sep&#160;24&#160;16:13&#160;UTC | false | [view](CERTS/7843e095320e32655bfd2a3c06813c798df174fcdd9a9ce2e797276764614875/README.md) |
| 08&#160;Aug&#160;24&#160;21:49&#160;UTC | SHAKEN 073H | 08&#160;Aug&#160;25&#160;21:49&#160;UTC | false | [view](CERTS/618b83ffb285c2c03f430be1e7783fc04767222fb618affd492c968e1e834035/README.md) |
| 12&#160;Aug&#160;24&#160;22:24&#160;UTC | SHAKEN 577F | 11&#160;Sep&#160;24&#160;22:24&#160;UTC | false | [view](CERTS/7b2154825b993d0b0316fa8ad9d2c069eb08c33a41515c8810f164214cf714a2/README.md) |
| 14&#160;Aug&#160;24&#160;04:03&#160;UTC | SHAKEN 345J | 13&#160;Sep&#160;24&#160;04:03&#160;UTC | false | [view](CERTS/932436d23fa1052319c3921bc5f0a3775ec959dd75b153d8a6158878d63c93df/README.md) |
| 14&#160;Aug&#160;24&#160;05:30&#160;UTC | SHAKEN 841J | 28&#160;Aug&#160;24&#160;05:30&#160;UTC | false | [view](CERTS/d7d3b91be086f812937dbf7310bdb203bcb06eb2a59e02c581cea713bbf4b963/README.md) |
| 16&#160;Aug&#160;24&#160;03:42&#160;UTC | SHAKEN 952J | 15&#160;Sep&#160;24&#160;03:42&#160;UTC | false | [view](CERTS/0678862654a272680c27f8553872cfc8acf877e722e5bebfb16eb1d50aa050c8/README.md) |
| 19&#160;Aug&#160;24&#160;18:45&#160;UTC | SHAKEN 0360 | 26&#160;Aug&#160;24&#160;18:45&#160;UTC | false | [view](CERTS/053bba00482d4dde2ca3a525262ea430dd54168fa0259b89b2e98f41dc29748f/README.md) |
| 19&#160;Aug&#160;24&#160;19:06&#160;UTC | SHAKEN 060K | 26&#160;Aug&#160;24&#160;19:06&#160;UTC | false | [view](CERTS/eb30f15254e1c1e145ba6a3cd63e22dfef3d7c9c54ac92a3bf86787b2b919f48/README.md) |
| 19&#160;Aug&#160;24&#160;20:06&#160;UTC | SHAKEN 019K | 26&#160;Aug&#160;24&#160;20:06&#160;UTC | false | [view](CERTS/7b5c1e0d95887ed18f4b960d6596e8d0e0188ae4b0b3afa80131639f6fef4fad/README.md) |
| 19&#160;Aug&#160;24&#160;20:26&#160;UTC | SHAKEN 604K | 26&#160;Aug&#160;24&#160;20:26&#160;UTC | false | [view](CERTS/59eedbd312860eb6460120876a2b81727e041edd40f58b8a68d526e8ce5b0306/README.md) |
| 19&#160;Aug&#160;24&#160;21:01&#160;UTC | SHAKEN 0172 | 26&#160;Aug&#160;24&#160;21:01&#160;UTC | false | [view](CERTS/abaa564e15130f68385d01d2472a07b231b05a774ec2e70a86a3fd3014cad43f/README.md) |
| 19&#160;Aug&#160;24&#160;21:02&#160;UTC | SHAKEN 320K | 26&#160;Aug&#160;24&#160;21:02&#160;UTC | false | [view](CERTS/cad13002c6b7529e99ebbb472c87796e7f3c27054226171c77c6fad21d42017b/README.md) |
| 19&#160;Aug&#160;24&#160;21:20&#160;UTC | SHAKEN 674J | 26&#160;Aug&#160;24&#160;21:20&#160;UTC | false | [view](CERTS/616b23880e9e8e37a11c155a7f031d42468aa07c8d642267d41518f4a0ef952c/README.md) |
| 19&#160;Aug&#160;24&#160;21:22&#160;UTC | SHAKEN 1642 | 26&#160;Aug&#160;24&#160;21:22&#160;UTC | false | [view](CERTS/de06bdea1aa7dc26b0a0321b37fb6f309a3d22d25386dd086ad9e2570bfc3a0d/README.md) |
| 19&#160;Aug&#160;24&#160;21:23&#160;UTC | SHAKEN 014E | 26&#160;Aug&#160;24&#160;21:23&#160;UTC | false | [view](CERTS/bfdab90ff4afa814d502388a46897ee09b989a2b3e5aef6b3675e9cca6627636/README.md) |
| 19&#160;Aug&#160;24&#160;21:24&#160;UTC | SHAKEN 856H | 26&#160;Aug&#160;24&#160;21:24&#160;UTC | false | [view](CERTS/3bb3cd357da5efa90458a41805f99c6a12d99a4aa88d907bd85a4ec623cb493f/README.md) |
| 19&#160;Aug&#160;24&#160;21:26&#160;UTC | SHAKEN 819H | 26&#160;Aug&#160;24&#160;21:26&#160;UTC | false | [view](CERTS/b70d047b25b04a0866badd31c56e287a773a8823a5a4f28e4f3305ec5defca45/README.md) |
| 19&#160;Aug&#160;24&#160;21:29&#160;UTC | SHAKEN 0753 | 26&#160;Aug&#160;24&#160;21:29&#160;UTC | false | [view](CERTS/dfc5b1889190ce87502e3eb4462eaf92fd6ecb8e852a382865bdac43144b2b8d/README.md) |
| 19&#160;Aug&#160;24&#160;21:35&#160;UTC | SHAKEN 469A | 26&#160;Aug&#160;24&#160;21:35&#160;UTC | false | [view](CERTS/da97cf5c90a23bc69e7747a7efb62b20e60be36c458a2fc902f551106aaff554/README.md) |
| 19&#160;Aug&#160;24&#160;21:35&#160;UTC | SHAKEN 726J | 26&#160;Aug&#160;24&#160;21:35&#160;UTC | false | [view](CERTS/f17a4f8cd618f552a20af5a234a3be2b33270794d1617808335c7882f30a7a51/README.md) |
| 19&#160;Aug&#160;24&#160;21:36&#160;UTC | SHAKEN 790J | 26&#160;Aug&#160;24&#160;21:36&#160;UTC | false | [view](CERTS/7d8f4491d9634121e00f2726f91274724a98184a568f2b5e17a770b397942a3f/README.md) |
| 19&#160;Aug&#160;24&#160;21:41&#160;UTC | SHAKEN 459J | 26&#160;Aug&#160;24&#160;21:41&#160;UTC | false | [view](CERTS/5760a3ca4dd06a66f79b4e38f8972589f86d9b30ade51def296d7e677d68fd36/README.md) |
| 19&#160;Aug&#160;24&#160;21:41&#160;UTC | SHAKEN 672B | 26&#160;Aug&#160;24&#160;21:41&#160;UTC | false | [view](CERTS/01d5b8117d736764d0e0496b1ec7ae6035d6e831a0346196b3ab9500cb2d25fb/README.md) |
| 20&#160;Aug&#160;24&#160;14:25&#160;UTC | SHAKEN 912K | 27&#160;Aug&#160;24&#160;14:25&#160;UTC | false | [view](CERTS/d41ecf2105de52ce12af9d4f12891a3055d02e93b1ad7dcf94f6ead19da10665/README.md) |
| 20&#160;Aug&#160;24&#160;15:39&#160;UTC | SHAKEN 208K | 27&#160;Aug&#160;24&#160;15:39&#160;UTC | false | [view](CERTS/cdfaeef7117d9b605ae3121fa1cba831db0f59d5c7755802b1ebca33dfcef161/README.md) |
| 20&#160;Aug&#160;24&#160;15:41&#160;UTC | SHAKEN 738J | 29&#160;Aug&#160;24&#160;15:41&#160;UTC | false | [view](CERTS/c65aeff28907b5dcda306096a6d59e8d10cb94ec8dd4e040c63462f0f77a7b8b/README.md) |
| 20&#160;Aug&#160;24&#160;16:03&#160;UTC | SHAKEN 190J | 27&#160;Aug&#160;24&#160;16:03&#160;UTC | false | [view](CERTS/a279173867cfaf9e82c40f260012b3141d598102eeb7cc36994a20614438aa95/README.md) |
| 20&#160;Aug&#160;24&#160;16:07&#160;UTC | SHAKEN 444G | 27&#160;Aug&#160;24&#160;16:07&#160;UTC | false | [view](CERTS/478caf148d410029022cec415ad7967b6f36074d2d85d3fbc58a8e1e5aa3fde5/README.md) |
| 20&#160;Aug&#160;24&#160;17:53&#160;UTC | SHAKEN 1577 | 27&#160;Aug&#160;24&#160;17:53&#160;UTC | false | [view](CERTS/711144f9eaa56f6ca114457764e478938b5573b21f2db25207a4ea863958580d/README.md) |
| 20&#160;Aug&#160;24&#160;18:32&#160;UTC | SHAKEN 6233 | 27&#160;Aug&#160;24&#160;18:32&#160;UTC | false | [view](CERTS/5b8757f64828d10e6ba0d9d201616d2565b4d63af105c3991aec98597813b287/README.md) |
| 20&#160;Aug&#160;24&#160;18:35&#160;UTC | SHAKEN 927K | 27&#160;Aug&#160;24&#160;18:35&#160;UTC | false | [view](CERTS/48ee6498265e79527bf5eb7f2f7dfbee86c9ebc7176210458062e3aad69f09e6/README.md) |
| 20&#160;Aug&#160;24&#160;18:51&#160;UTC | SHAKEN 0860 | 27&#160;Aug&#160;24&#160;18:51&#160;UTC | false | [view](CERTS/ba3a6e22e90979cc75b73b7a92defe244d875aff344cc15ba6ad60f848113f35/README.md) |
| 20&#160;Aug&#160;24&#160;18:58&#160;UTC | SHAKEN 406K | 27&#160;Aug&#160;24&#160;18:58&#160;UTC | false | [view](CERTS/826d7d6fc459cd311292403f953ca8229bd4f9a2b41b5c84837b09bde1df4f3a/README.md) |
| 20&#160;Aug&#160;24&#160;19:09&#160;UTC | SHAKEN 979K | 27&#160;Aug&#160;24&#160;19:09&#160;UTC | false | [view](CERTS/eb8ee79190ee2b8691e4a2d4459a9221299f60f46850d2ff16799ba66e0786a5/README.md) |
| 20&#160;Aug&#160;24&#160;19:21&#160;UTC | SHAKEN 1917 | 27&#160;Aug&#160;24&#160;19:21&#160;UTC | false | [view](CERTS/e8f1e3be5262876899b2901efc0803bca5836968cc6aabd314734878ef4d77b8/README.md) |
| 20&#160;Aug&#160;24&#160;21:11&#160;UTC | SHAKEN 983J | 27&#160;Aug&#160;24&#160;21:11&#160;UTC | false | [view](CERTS/8347f21b68a7b662e7706d67434c5e8555dd0cc3ed46d6515a26165382115081/README.md) |
| 20&#160;Aug&#160;24&#160;21:40&#160;UTC | SHAKEN 663G | 27&#160;Aug&#160;24&#160;21:40&#160;UTC | false | [view](CERTS/4be9436b9d4bee87f78b2cec231251dbe6b6c91a565f135e8e9a8d4f5d8872a9/README.md) |
| 21&#160;Aug&#160;24&#160;02:31&#160;UTC | SHAKEN 551G | 20&#160;Sep&#160;24&#160;02:31&#160;UTC | false | [view](CERTS/e4e5a1389303848d0cafc4aa7e18164fbf2f4a5af802af85c2556b512cb2849e/README.md) |
| 21&#160;Aug&#160;24&#160;06:01&#160;UTC | SHAKEN 278K | 28&#160;Aug&#160;24&#160;06:01&#160;UTC | false | [view](CERTS/d03df7e6b8cd798a92b63ef748b37247485d3f172811211c6b0d2ecb80e95fc0/README.md) |
| 21&#160;Aug&#160;24&#160;06:01&#160;UTC | SHAKEN 606F | 28&#160;Aug&#160;24&#160;06:01&#160;UTC | false | [view](CERTS/54a124f67fd5571e8cf9157aa9fac90a4e9dcd2908fc1b0725e0174a7ba4b416/README.md) |
| 21&#160;Aug&#160;24&#160;06:01&#160;UTC | SHAKEN 653J | 28&#160;Aug&#160;24&#160;06:01&#160;UTC | false | [view](CERTS/fc1e50acf2ee8280740d7d51ed6af99f2c3e2df84c1992387d159ce8da78de95/README.md) |
| 21&#160;Aug&#160;24&#160;07:01&#160;UTC | SHAKEN 625J | 28&#160;Aug&#160;24&#160;07:01&#160;UTC | false | [view](CERTS/78a7977e2d20915bb7b8bc6b8d878bfce1976af9ccef35609a4488044b671a29/README.md) |
| 21&#160;Aug&#160;24&#160;10:32&#160;UTC | SHAKEN 982J | 28&#160;Aug&#160;24&#160;10:32&#160;UTC | false | [view](CERTS/d0cf5c10a3c82a91777b7aa67a914164e1017d7a6b4fdb4e862aa62009984c4f/README.md) |
| 21&#160;Aug&#160;24&#160;12:15&#160;UTC | SHAKEN 1577 | 28&#160;Aug&#160;24&#160;12:15&#160;UTC | false | [view](CERTS/6abfe5b090af1fcc4fefe8798696ae98d4982e55a8368b7e89799e525249aeb5/README.md) |
| 21&#160;Aug&#160;24&#160;13:18&#160;UTC | SHAKEN 691J | 28&#160;Aug&#160;24&#160;13:18&#160;UTC | false | [view](CERTS/2671db4d5b06bb6f854599c8207e0b95a5dd406bc9f8db540df74d18cfe8a652/README.md) |
| 21&#160;Aug&#160;24&#160;14:37&#160;UTC | SHAKEN 495K | 28&#160;Aug&#160;24&#160;14:37&#160;UTC | false | [view](CERTS/802f1d5d9ecb34c3ecf806ed9d63ccf78cf3d8a6fbd8bada2ee6787f5b3f8d6a/README.md) |
| 21&#160;Aug&#160;24&#160;15:25&#160;UTC | SHAKEN 2550 | 28&#160;Aug&#160;24&#160;15:25&#160;UTC | false | [view](CERTS/fd465f805a0df89b665e697de4f9f95daa6f480c997b99fecf65176cc007d106/README.md) |
| 21&#160;Aug&#160;24&#160;16:43&#160;UTC | SHAKEN 646K | 28&#160;Aug&#160;24&#160;16:43&#160;UTC | false | [view](CERTS/f2ea0b871219bf56443f801e47b2dcebbbcb9f006a4381f7b5f2d35f66e55e6d/README.md) |
| 21&#160;Aug&#160;24&#160;17:06&#160;UTC | SHAKEN 738K | 28&#160;Aug&#160;24&#160;17:06&#160;UTC | false | [view](CERTS/290c88e6790f9f6a9944f94313383fe36284c38ebef0908c5386045d349e3c61/README.md) |
| 21&#160;Aug&#160;24&#160;18:18&#160;UTC | SHAKEN 833J | 28&#160;Aug&#160;24&#160;18:18&#160;UTC | false | [view](CERTS/e14f88611aab7c3bceac7f57d918c6625c3398a8312305891c1eaa0c43d91058/README.md) |
| 21&#160;Aug&#160;24&#160;20:39&#160;UTC | SHAKEN 796K | 28&#160;Aug&#160;24&#160;20:39&#160;UTC | false | [view](CERTS/47f28b240e6ed301691b4a518caebf6a1b1e7674efb2288cf3e3efbaef700230/README.md) |
| 21&#160;Aug&#160;24&#160;21:11&#160;UTC | SHAKEN 177K | 28&#160;Aug&#160;24&#160;21:11&#160;UTC | false | [view](CERTS/db2dacb79e0cff0051fbcff24500b82eaccf65b6163ba28065adc3b9ec6ab6c6/README.md) |
| 22&#160;Aug&#160;24&#160;15:12&#160;UTC | SHAKEN 841J | 05&#160;Sep&#160;24&#160;15:12&#160;UTC | false | [view](CERTS/abcbc623a2ac5754d5cca3a42bc1b2286a1507847db5284a66c90f4ee1d37165/README.md) |
| 22&#160;Aug&#160;24&#160;15:34&#160;UTC | SHAKEN 012K | 29&#160;Aug&#160;24&#160;15:34&#160;UTC | false | [view](CERTS/73f485a4a8fb6c77f4128e8dc7b1b46e751b73933926f1f3edc8d465dc86aed5/README.md) |
| 22&#160;Aug&#160;24&#160;19:06&#160;UTC | SHAKEN 060K | 29&#160;Aug&#160;24&#160;19:06&#160;UTC | false | [view](CERTS/8ad274cdf33526b8b27c610c508658ea72d7e236ccb7128d9d9b3dd04f9f3d89/README.md) |
| 22&#160;Aug&#160;24&#160;19:35&#160;UTC | SHAKEN 9714 | 29&#160;Aug&#160;24&#160;19:35&#160;UTC | false | [view](CERTS/dc2fd2819dcf6b32a115ef22f64d3195e15da9184a4e01d01ca5ca16dbda5e95/README.md) |
| 22&#160;Aug&#160;24&#160;20:06&#160;UTC | SHAKEN 019K | 29&#160;Aug&#160;24&#160;20:06&#160;UTC | false | [view](CERTS/9f9b045919b457296a096b2ffe1b55b6eeb30112afb7939db9844fb284842b60/README.md) |
| 22&#160;Aug&#160;24&#160;20:27&#160;UTC | SHAKEN 604K | 29&#160;Aug&#160;24&#160;20:27&#160;UTC | false | [view](CERTS/f50e6684fa860fb633d79cb7386d1d5ab34e3f02dd7cebe66c82fd74e72e31fb/README.md) |
| 22&#160;Aug&#160;24&#160;21:00&#160;UTC | SHAKEN 297K | 29&#160;Aug&#160;24&#160;21:00&#160;UTC | false | [view](CERTS/bf72a17b97416695d3ccbe6ccb17cb03c4d6e5224529b601395fc97ecc09d8c2/README.md) |
| 22&#160;Aug&#160;24&#160;21:20&#160;UTC | SHAKEN 674J | 29&#160;Aug&#160;24&#160;21:20&#160;UTC | false | [view](CERTS/d0a6b996700912c3a1443785f40b83bfbeb4a52506403a34437570f94fa56eec/README.md) |
| 22&#160;Aug&#160;24&#160;21:21&#160;UTC | SHAKEN 735J | 29&#160;Aug&#160;24&#160;21:21&#160;UTC | false | [view](CERTS/8943f53535a1f8774b7f5e6801ccc89d51069d9623990247a3c9973f7be215a2/README.md) |
| 22&#160;Aug&#160;24&#160;21:27&#160;UTC | SHAKEN 819H | 29&#160;Aug&#160;24&#160;21:26&#160;UTC | false | [view](CERTS/d8d56f168fc72fa7b0aba73f11db2e2c8a23c494e4f6c1c497c78d19efe5b66f/README.md) |
| 22&#160;Aug&#160;24&#160;21:35&#160;UTC | SHAKEN 769J | 29&#160;Aug&#160;24&#160;21:35&#160;UTC | false | [view](CERTS/984ef22b60a451de14a1e6ece80059844aae3d6a8a7995d213780706e1c9020c/README.md) |
| 22&#160;Aug&#160;24&#160;21:36&#160;UTC | SHAKEN 849J | 29&#160;Aug&#160;24&#160;21:36&#160;UTC | false | [view](CERTS/877eafad7970e1e8592fc377d91435b0d363b3ba6008e9b9ff473dfed23a180f/README.md) |
| 22&#160;Aug&#160;24&#160;21:36&#160;UTC | SHAKEN 469A | 29&#160;Aug&#160;24&#160;21:36&#160;UTC | false | [view](CERTS/2840bd21d7f8350ec9a8eda0528d2ad43e7d16ef100b02f57db6e5d71efc7885/README.md) |
| 22&#160;Aug&#160;24&#160;21:37&#160;UTC | SHAKEN 790J | 29&#160;Aug&#160;24&#160;21:37&#160;UTC | false | [view](CERTS/408f7b5b456f3610bcd0a63e321ccddcb3d78f47134857c160716700bef87eef/README.md) |
| 22&#160;Aug&#160;24&#160;21:41&#160;UTC | SHAKEN 459J | 29&#160;Aug&#160;24&#160;21:41&#160;UTC | false | [view](CERTS/3a30e9b63c240c195072096084f528247c6b421bbce02ae85f463ffe94676c91/README.md) |
| 22&#160;Aug&#160;24&#160;21:41&#160;UTC | SHAKEN 672B | 29&#160;Aug&#160;24&#160;21:41&#160;UTC | false | [view](CERTS/61bd0f4cbe8e109ccd76faf9be8f939aab0a1e02392e78923cac9da9c15cbe89/README.md) |
| 23&#160;Aug&#160;24&#160;15:40&#160;UTC | SHAKEN 208K | 30&#160;Aug&#160;24&#160;15:40&#160;UTC | false | [view](CERTS/d061ed571f8c788824596e097ef166dd1bf789e6703caf4e5f95b4acf610ac02/README.md) |
| 23&#160;Aug&#160;24&#160;18:02&#160;UTC | SHAKEN 835K | 30&#160;Aug&#160;24&#160;18:02&#160;UTC | false | [view](CERTS/1b9e692ba61e818970dbaac862ec78819d6e6b4fe0b383cfd6c5dffafc57665a/README.md) |
| 23&#160;Aug&#160;24&#160;18:58&#160;UTC | SHAKEN 406K | 30&#160;Aug&#160;24&#160;18:58&#160;UTC | false | [view](CERTS/ba1357c6660cc1375636350b562038e71f93e8d4c5be23d26c8dc901002c1cac/README.md) |
| 23&#160;Aug&#160;24&#160;21:12&#160;UTC | SHAKEN 983J | 30&#160;Aug&#160;24&#160;21:12&#160;UTC | false | [view](CERTS/75a856d41167a256b0632f222855648bb834137dea623768a3833d54d2bf2f29/README.md) |
| 23&#160;Aug&#160;24&#160;21:40&#160;UTC | SHAKEN 663G | 30&#160;Aug&#160;24&#160;21:40&#160;UTC | false | [view](CERTS/f281e7107daf1f375c56afb87a1c027fe14c4cdc218414d0a4357af8a87080a4/README.md) |
| 24&#160;Aug&#160;24&#160;01:53&#160;UTC | SHAKEN 0172 | 31&#160;Aug&#160;24&#160;01:53&#160;UTC | false | [view](CERTS/efe561e048922e144ba0fec2422cc9351d42b0adfea3267fb9b4cd8781c1124d/README.md) |
| 24&#160;Aug&#160;24&#160;06:52&#160;UTC | SHAKEN 345J | 23&#160;Sep&#160;24&#160;06:51&#160;UTC | false | [view](CERTS/03a2c8388e295255cac1d2d88409fbceb57cc9d8688339e7c302a29c7a939fc3/README.md) |
| 24&#160;Aug&#160;24&#160;12:15&#160;UTC | SHAKEN 1577 | 31&#160;Aug&#160;24&#160;12:15&#160;UTC | false | [view](CERTS/0af3b33adfde865b82ee33e0281e331a30575b6f93999170f458601f2e210198/README.md) |
| 24&#160;Aug&#160;24&#160;16:43&#160;UTC | SHAKEN 646K | 31&#160;Aug&#160;24&#160;16:43&#160;UTC | false | [view](CERTS/34fea035aa1b3165d5f10f8fb307a00b91f8bb79d95050a0bfb9e671cef70850/README.md) |
| 24&#160;Aug&#160;24&#160;17:06&#160;UTC | SHAKEN 738K | 31&#160;Aug&#160;24&#160;17:06&#160;UTC | false | [view](CERTS/ab643f6f4b9a9d00d8bf5a65428f78f098631c8f8f3530eed9823410cab84821/README.md) |
| 25&#160;Aug&#160;24&#160;00:02&#160;UTC | SHAKEN 523H | 01&#160;Sep&#160;24&#160;00:02&#160;UTC | false | [view](CERTS/dd3fcc3f8541aad5a2b39324a8f2cf716c6ae87f8ce7d2bb5dbc66300aa0ef0d/README.md) |
| 25&#160;Aug&#160;24&#160;15:36&#160;UTC | SHAKEN 747J | 01&#160;Sep&#160;24&#160;15:36&#160;UTC | false | [view](CERTS/afd029621efbba4866d1a08c677e8f56084b6664135fa72d95eac897977bb43c/README.md) |
| 25&#160;Aug&#160;24&#160;21:24&#160;UTC | SHAKEN 856H | 01&#160;Sep&#160;24&#160;21:24&#160;UTC | false | [view](CERTS/e29f6b8caca92c75a262aaa6d18e9cc8923d1da986ef0acee39d8b84165dfeb7/README.md) |
| 25&#160;Aug&#160;24&#160;21:37&#160;UTC | SHAKEN 790J | 01&#160;Sep&#160;24&#160;21:37&#160;UTC | false | [view](CERTS/8824e923289eab3b1a32687fa6631c71efcd88913d0009718114a61ac7e4f8cb/README.md) |
| 25&#160;Aug&#160;24&#160;21:41&#160;UTC | SHAKEN 459J | 01&#160;Sep&#160;24&#160;21:41&#160;UTC | false | [view](CERTS/10185ea40f5724c010449d1279bf5cc44ea2dd8110734e64fc00374a8d86db73/README.md) |
| 25&#160;Aug&#160;24&#160;21:41&#160;UTC | SHAKEN 672B | 01&#160;Sep&#160;24&#160;21:41&#160;UTC | false | [view](CERTS/990ed0fa0aa3009a36f1dfc56e8f024b7962ba1799663f9caeae2b18a09549a1/README.md) |

#### CA Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 24&#160;Oct&#160;22&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Issuing CA4 | 23&#160;Oct&#160;32&#160;23:59&#160;UTC | false | [view](CERTS/80dd12d935f46c256c7c1289a0834b21ec8f2f8c35f2864928d1e75f3a280665/README.md) |
| 24&#160;Oct&#160;22&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Root CA2 | 23&#160;Oct&#160;42&#160;23:59&#160;UTC | false | [view](CERTS/a26e04fc786ab70b8085236b2c53f8cfbf5d0c6a5c2c9c3e9f91669fbb8ea4d5/README.md) |
| 21&#160;Mar&#160;24&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Issuing CA5 | 20&#160;Mar&#160;34&#160;23:59&#160;UTC | false | [view](CERTS/cd50eeb8c083af686a49964a10b030048b800530edbeee8f0991388c3a79e75a/README.md) |


Generated: 26 Aug 24 18:03 UTC