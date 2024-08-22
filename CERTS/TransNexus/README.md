# STIR/SHAKEN CA Ecosystem Compliance

## TransNexus

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 6690 certificates were included in the corpus being tested
- 2 certificates in the corpus were skipped because they are duplicates
- 6174 certificates in the corpus were skipped because they are expired
- 457 certificates in the corpus were skipped because they are not currently trusted
- 57 certificates being tested against the remaining rules
- 1.00 issues on average found in unexpired, trusted, and non-compliant certificates
- 12.28% of certificates contain one or more Error level issue
- 0.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 106 days is the average remaining validity for the certificates in the corpus
- 104 days is the average initial validity for the certificates in the corpus
- 38 certificates expire in the next 30 days
- 1.21 average number of unexpired certificates per OCN observed
- 47 unique OCNs observed in unexpired and valid certificate corpus

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
- 5312 days is the average remaining validity for the certificates in the corpus
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
| 24&#160;Jul&#160;24&#160;04:01&#160;UTC | SHAKEN 345J | 23&#160;Aug&#160;24&#160;04:01&#160;UTC | false | [view](CERTS/99c0eda5fce20d3a71ee499b507667d1ea174e18808c81a861549f042bfadf0f/README.md) |
| 25&#160;Jul&#160;24&#160;22:13&#160;UTC | SHAKEN 577F | 24&#160;Aug&#160;24&#160;22:13&#160;UTC | false | [view](CERTS/4391a7de99de252b9ab3046f9bac1b9ed80a54df55705df27a899580ccbc40b8/README.md) |
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
| 15&#160;Aug&#160;24&#160;16:13&#160;UTC | SHAKEN 0172 | 22&#160;Aug&#160;24&#160;16:13&#160;UTC | false | [view](CERTS/3323d24f73b09fd8ff93a2d1a22ccf3b19f8fcb2634b2cd44bdd1301f616ffdc/README.md) |
| 15&#160;Aug&#160;24&#160;16:42&#160;UTC | SHAKEN 646K | 22&#160;Aug&#160;24&#160;16:42&#160;UTC | false | [view](CERTS/8a209ebad68ddd83d767e9fe98f80292c6950b19257491c45481176f8a154a7e/README.md) |
| 15&#160;Aug&#160;24&#160;17:06&#160;UTC | SHAKEN 738K | 22&#160;Aug&#160;24&#160;17:06&#160;UTC | false | [view](CERTS/2b651055781cd784a90e3595659106e8a0163c3bb151f259b68464e509ebe95d/README.md) |
| 15&#160;Aug&#160;24&#160;18:17&#160;UTC | SHAKEN 833J | 22&#160;Aug&#160;24&#160;18:17&#160;UTC | false | [view](CERTS/bffa277aa1b06f3547de0569e3a9ad498a5bd8def76c7b3430d2cf806a328e8c/README.md) |
| 15&#160;Aug&#160;24&#160;21:10&#160;UTC | SHAKEN 177K | 22&#160;Aug&#160;24&#160;21:10&#160;UTC | false | [view](CERTS/1a13bba9bf760b810530600ce9fcf5d59c1845594145ca57cb1c25737bf3e2f5/README.md) |
| 16&#160;Aug&#160;24&#160;03:42&#160;UTC | SHAKEN 952J | 15&#160;Sep&#160;24&#160;03:42&#160;UTC | false | [view](CERTS/0678862654a272680c27f8553872cfc8acf877e722e5bebfb16eb1d50aa050c8/README.md) |
| 16&#160;Aug&#160;24&#160;15:35&#160;UTC | SHAKEN 747J | 23&#160;Aug&#160;24&#160;15:35&#160;UTC | false | [view](CERTS/b095b0bf61f8b8d366622c2aed3182ff43a7677c8c224c93733ce55b33b8c5d3/README.md) |
| 16&#160;Aug&#160;24&#160;21:24&#160;UTC | SHAKEN 733J | 23&#160;Aug&#160;24&#160;21:24&#160;UTC | false | [view](CERTS/577442375cd51d8adcba43fd6bd38ba6a4e991c0ba52b982c502a780350242c2/README.md) |
| 16&#160;Aug&#160;24&#160;21:26&#160;UTC | SHAKEN 819H | 23&#160;Aug&#160;24&#160;21:26&#160;UTC | false | [view](CERTS/1564b556f87af2fa02b5a3cb36d74558349c421887703cda43a069d2b6e9ace7/README.md) |
| 16&#160;Aug&#160;24&#160;21:30&#160;UTC | SHAKEN 053G | 23&#160;Aug&#160;24&#160;21:30&#160;UTC | false | [view](CERTS/df82fa194ff36b1511d0abb9f09f267c237f10cddc749c3bd7edfcc66b66cbe4/README.md) |
| 16&#160;Aug&#160;24&#160;21:36&#160;UTC | SHAKEN 790J | 23&#160;Aug&#160;24&#160;21:36&#160;UTC | false | [view](CERTS/f6bb7c8d48c588ab34b5cea08633fef4629c33a2ad10e101904ed1e326c1ca9b/README.md) |
| 16&#160;Aug&#160;24&#160;21:40&#160;UTC | SHAKEN 459J | 23&#160;Aug&#160;24&#160;21:40&#160;UTC | false | [view](CERTS/95b5fce668b178087d43752e435f4730b120486c47e9f7d6ea20b05c99cf79bd/README.md) |
| 17&#160;Aug&#160;24&#160;15:39&#160;UTC | SHAKEN 208K | 24&#160;Aug&#160;24&#160;15:39&#160;UTC | false | [view](CERTS/b31af1d7e63a998a290010b9af01bf8fa8281473d6183c8e221f0e6e66c51d45/README.md) |
| 17&#160;Aug&#160;24&#160;16:03&#160;UTC | SHAKEN 190J | 24&#160;Aug&#160;24&#160;16:03&#160;UTC | false | [view](CERTS/088c7b36c8d306bf38b9f664fe10f8d2a868fc70ce0e8038ae35b033062e4b72/README.md) |
| 17&#160;Aug&#160;24&#160;18:35&#160;UTC | SHAKEN 927K | 24&#160;Aug&#160;24&#160;18:35&#160;UTC | false | [view](CERTS/3a1573f6e420b6d25c3387fafe0d13b9cefe4308560b663152c0d167b35f8a75/README.md) |
| 17&#160;Aug&#160;24&#160;20:38&#160;UTC | SHAKEN 0588 | 24&#160;Aug&#160;24&#160;20:38&#160;UTC | false | [view](CERTS/bf35d8b510a134b85f3448ed44f6f9ae51e3ca246d8d77662b640dd384043bcb/README.md) |
| 17&#160;Aug&#160;24&#160;21:11&#160;UTC | SHAKEN 983J | 24&#160;Aug&#160;24&#160;21:11&#160;UTC | false | [view](CERTS/7649fa6cd6fac7eb06b02febe7d3638c2a64c5832036c219b7b081e5096aa7c7/README.md) |
| 17&#160;Aug&#160;24&#160;21:40&#160;UTC | SHAKEN 663G | 24&#160;Aug&#160;24&#160;21:40&#160;UTC | false | [view](CERTS/ece91c890428bef3605092516ecf595ebb849e4fbd91f1f9a57f1bcf20dc267c/README.md) |
| 18&#160;Aug&#160;24&#160;06:00&#160;UTC | SHAKEN 278K | 25&#160;Aug&#160;24&#160;06:00&#160;UTC | false | [view](CERTS/7bcbb7557bc6d0512d8bd1d11f34c10c1c0a2ee68042fdfbdd667996650f2cda/README.md) |
| 18&#160;Aug&#160;24&#160;06:00&#160;UTC | SHAKEN 120K | 25&#160;Aug&#160;24&#160;06:00&#160;UTC | false | [view](CERTS/34bac5a6f45851707635f294c4de0f8007b667b2bf5622fc3824531786cedbcd/README.md) |
| 18&#160;Aug&#160;24&#160;12:14&#160;UTC | SHAKEN 1577 | 25&#160;Aug&#160;24&#160;12:14&#160;UTC | false | [view](CERTS/45e9e187a113c06aba0b92f35f2d234471c122b8c01700b2aeadd0c29a353ba3/README.md) |
| 18&#160;Aug&#160;24&#160;16:42&#160;UTC | SHAKEN 646K | 25&#160;Aug&#160;24&#160;16:42&#160;UTC | false | [view](CERTS/be98c3da59d432d393e54aad548026aa7ab010cdca0e149552bd8efd9b117c40/README.md) |
| 19&#160;Aug&#160;24&#160;15:32&#160;UTC | SHAKEN 878K | 26&#160;Aug&#160;24&#160;15:32&#160;UTC | false | [view](CERTS/6229db2b4ff58f07bea2726332a52094634862db67bb2d61e1dee15309c631ab/README.md) |
| 19&#160;Aug&#160;24&#160;15:36&#160;UTC | SHAKEN 747J | 26&#160;Aug&#160;24&#160;15:36&#160;UTC | false | [view](CERTS/2418eb2715467440c29dc41eb3ff2bbe9132ece037a27f02d64a282d30cd3ea6/README.md) |
| 19&#160;Aug&#160;24&#160;18:45&#160;UTC | SHAKEN 0360 | 26&#160;Aug&#160;24&#160;18:45&#160;UTC | false | [view](CERTS/053bba00482d4dde2ca3a525262ea430dd54168fa0259b89b2e98f41dc29748f/README.md) |
| 19&#160;Aug&#160;24&#160;21:01&#160;UTC | SHAKEN 0172 | 26&#160;Aug&#160;24&#160;21:01&#160;UTC | false | [view](CERTS/abaa564e15130f68385d01d2472a07b231b05a774ec2e70a86a3fd3014cad43f/README.md) |
| 19&#160;Aug&#160;24&#160;21:36&#160;UTC | SHAKEN 790J | 26&#160;Aug&#160;24&#160;21:36&#160;UTC | false | [view](CERTS/7d8f4491d9634121e00f2726f91274724a98184a568f2b5e17a770b397942a3f/README.md) |

#### CA Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 24&#160;Oct&#160;22&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Issuing CA4 | 23&#160;Oct&#160;32&#160;23:59&#160;UTC | false | [view](CERTS/80dd12d935f46c256c7c1289a0834b21ec8f2f8c35f2864928d1e75f3a280665/README.md) |
| 24&#160;Oct&#160;22&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Root CA2 | 23&#160;Oct&#160;42&#160;23:59&#160;UTC | false | [view](CERTS/a26e04fc786ab70b8085236b2c53f8cfbf5d0c6a5c2c9c3e9f91669fbb8ea4d5/README.md) |
| 21&#160;Mar&#160;24&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Issuing CA5 | 20&#160;Mar&#160;34&#160;23:59&#160;UTC | false | [view](CERTS/cd50eeb8c083af686a49964a10b030048b800530edbeee8f0991388c3a79e75a/README.md) |


Generated: 22 Aug 24 16:06 UTC