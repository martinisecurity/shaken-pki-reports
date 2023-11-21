# STIR/SHAKEN CA Ecosystem Compliance

## Sansay

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 1752 certificates were included in the corpus being tested
- 17 certificates in the corpus were skipped because they are duplicates
- 1566 certificates in the corpus were skipped because they are expired
- 1 certificates in the corpus were skipped because they are not currently trusted
- 168 certificates being tested against the remaining rules
- 2.02 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 0.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 6.55% of certificates are too old to be assessed against currently enforced expectations
- 161 days is the average remaining validity for the certificates in the corpus
- 162 days is the average initial validity for the certificates in the corpus
- 103 certificates expire in the next 30 days
- 1.62 average number of unexpired certificates per OCN observed
- 104 unique OCNs observed in unexpired and valid certificate corpus

| Instances | Test | Source |
|-----------|------|--------|
| 168 | [e_atis_subject_cn_spc](ISSUES/e_atis_subject_cn_spc/README.md) | ATIS1000080 |
| 4 | [e_atis_tn_auth_list_spc_format](ISSUES/e_atis_tn_auth_list_spc_format/README.md) | ATIS1000080 |
| 56 | [e_shaken_certificate_policies_id](ISSUES/e_shaken_certificate_policies_id/README.md) | US_SHAKEN_CP |
| 112 | [e_us_cp_subject_sn_shall](ISSUES/e_us_cp_subject_sn_shall/README.md) | US_SHAKEN_CP |

#### CA Certificates

- 3 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 0 certificates in the corpus were skipped because they are expired
- 1 certificates in the corpus were skipped because they are not currently trusted
- 2 certificates being tested against the remaining rules
- 0.00 issues on average found in unexpired, trusted, and non-compliant certificates
- 0.00% of certificates contain one or more Error level issue
- 0.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 100.00% of certificates are too old to be assessed against currently enforced expectations
- 5323 days is the average remaining validity for the certificates in the corpus
- 4928 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days

No error, warning, or notice level issues were found

### Details

#### Leaf Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 21&#160;Nov&#160;22&#160;21:15&#160;UTC | SHAKEN Star2Star Communications, LLC 590J | 21&#160;Nov&#160;23&#160;21:15&#160;UTC | true | [view](CERTS/9bc9dde8921387803d93036c7d2f8085af32b028fca8f17336d2e22ab51fd278/README.md) |
| 29&#160;Nov&#160;22&#160;22:04&#160;UTC | SHAKEN MagicJack 324E | 29&#160;Nov&#160;23&#160;22:04&#160;UTC | true | [view](CERTS/75b4b7b400b1252e48faa1d93f6a94f7bd4a6383c88ddf6baa167b85d9ac4ee8/README.md) |
| 05&#160;Dec&#160;22&#160;22:28&#160;UTC | SHAKEN Bulk Solutions, LLC 644J | 05&#160;Dec&#160;23&#160;22:28&#160;UTC | true | [view](CERTS/3cf0aa2a24845e3fe6b27605e223e8e0c73d6bd4f73279b8a1e5e16fd2feeb80/README.md) |
| 10&#160;Dec&#160;22&#160;02:11&#160;UTC | SHAKEN Drop Inc 258K | 10&#160;Dec&#160;23&#160;02:11&#160;UTC | true | [view](CERTS/fc457741017b89b9126882710d8fb44883d7603f79cec0a1989eaa2b08034ee5/README.md) |
| 19&#160;Jan&#160;23&#160;22:50&#160;UTC | SHAKEN Technology Innovation Lab 599J | 19&#160;Jan&#160;24&#160;22:50&#160;UTC | true | [view](CERTS/12acafcf01348d278955bb9276e7a4d22a65ccdc61a59d08100177711f21b430/README.md) |
| 23&#160;Jan&#160;23&#160;21:55&#160;UTC | SHAKEN Swift Telco LLC 452K | 23&#160;Jan&#160;24&#160;21:55&#160;UTC | true | [view](CERTS/613861829aae7927f05dbd5a7b9f28ae8c4f995bb8ed115f95fc4be6644ccde1/README.md) |
| 26&#160;Jan&#160;23&#160;14:26&#160;UTC | SHAKEN Bulk Solutions, LLC 644J | 11&#160;Jan&#160;24&#160;14:26&#160;UTC | true | [view](CERTS/7dcc6cd32bf3c4e6e2901468097a88ad42a72ce42df34edce87c84dbce3691d2/README.md) |
| 01&#160;Feb&#160;23&#160;20:51&#160;UTC | SHAKEN ACS Technologies 488K | 01&#160;Feb&#160;24&#160;20:51&#160;UTC | true | [view](CERTS/2692a8727b74b38da7c0c4e3fb446dd123a70ea8c7e6b9af7e6560890d3ab076/README.md) |
| 08&#160;Feb&#160;23&#160;19:07&#160;UTC | SHAKEN ConvergeTel LLC 388K | 08&#160;Feb&#160;24&#160;19:07&#160;UTC | true | [view](CERTS/80706b79565f875515eb32f8cf113093a2658148ece8440e76199e4004254c31/README.md) |
| 14&#160;Feb&#160;23&#160;17:12&#160;UTC | SHAKEN Ytel Inc. 703J | 14&#160;Feb&#160;24&#160;17:12&#160;UTC | true | [view](CERTS/14c9bef113cfebe60611b0c56c430518ff8d42e8b98dd7e4653bd9cf619d5641/README.md) |
| 01&#160;Mar&#160;23&#160;14:14&#160;UTC | SHAKEN Voneto 485K | 29&#160;Feb&#160;24&#160;14:14&#160;UTC | true | [view](CERTS/8ca062af72eeb0b2eec4aaa1e6f3eb5ee3f01eccbd2895f282bbf4b70015a5d5/README.md) |
| 14&#160;Mar&#160;23&#160;21:39&#160;UTC | SHAKEN Cherry Voice 506K | 13&#160;Mar&#160;24&#160;21:39&#160;UTC | true | [view](CERTS/3756d1fa94ab1105ceb4bc6e3dcc371871b6edd45ac350df018a35f2f93f1967/README.md) |
| 19&#160;Mar&#160;23&#160;00:31&#160;UTC | SHAKEN Every1 Telecom 486K | 18&#160;Mar&#160;24&#160;00:31&#160;UTC | true | [view](CERTS/5cdfbb1a416083096dfef10c75a2b26a08d8fd5593d8ea9ceae0d70d878a97d1/README.md) |
| 27&#160;Mar&#160;23&#160;21:48&#160;UTC | SHAKEN Global Net Holdings Inc 306K | 21&#160;Jan&#160;24&#160;21:48&#160;UTC | true | [view](CERTS/f8d913af64a11c0718c78426b747dba4ed4ccb19239db573626a46f29b671825/README.md) |
| 28&#160;Mar&#160;23&#160;15:31&#160;UTC | SHAKEN VoIP Innovations 597F | 27&#160;Mar&#160;24&#160;15:31&#160;UTC | true | [view](CERTS/105a7683d4b5fac2ea7c2383e95250715bc0460d2cfdbea0d220201f44ea5d0c/README.md) |
| 28&#160;Mar&#160;23&#160;23:59&#160;UTC | SHAKEN Nemerald 326K | 27&#160;Mar&#160;24&#160;23:59&#160;UTC | true | [view](CERTS/bb3a45c0c72542b6ecdd1570bfbc446273c0d7f109dfdd0097fd8875339682e3/README.md) |
| 03&#160;Apr&#160;23&#160;21:02&#160;UTC | SHAKEN WWT INC dba VoIP Networks 053K | 02&#160;Apr&#160;24&#160;21:02&#160;UTC | true | [view](CERTS/6229273e97413ee4c01a1810885db9b48d7e1bf7fe8aa4e7be22076effe2cc8b/README.md) |
| 05&#160;Apr&#160;23&#160;16:27&#160;UTC | SHAKEN Swift Telco LLC 452K | 04&#160;Apr&#160;24&#160;16:27&#160;UTC | true | [view](CERTS/947b46067a639b79ff82ab3f48c453e4af7cc6d6036f6d66a742cc935bc8a35e/README.md) |
| 19&#160;Apr&#160;23&#160;06:07&#160;UTC | SHAKEN Telxio Networks 492K | 18&#160;Apr&#160;24&#160;06:07&#160;UTC | true | [view](CERTS/010c6a74330323c20ceb343b1de3a1e3248b4a3926c9ad2ed53f02b02399d241/README.md) |
| 19&#160;Apr&#160;23&#160;06:39&#160;UTC | SHAKEN Telxio Networks 492K | 18&#160;Apr&#160;24&#160;06:39&#160;UTC | true | [view](CERTS/1baf0e5fedd50fb55ff4e07366c5eb4f8d849b760739ffb8a0df4eb4828d7944/README.md) |
| 25&#160;Apr&#160;23&#160;16:43&#160;UTC | SHAKEN GIP Technology 434K | 24&#160;Apr&#160;24&#160;16:43&#160;UTC | true | [view](CERTS/bc5cd573f3ab46daf12994739622514f37ac8cd3275ff8d493595fee747e8a0b/README.md) |
| 18&#160;May&#160;23&#160;01:12&#160;UTC | SHAKEN Cloud Connect LLC 455K | 17&#160;May&#160;24&#160;01:12&#160;UTC | true | [view](CERTS/0aa593ccacc13e85c2ec381274e47b597989c2a57173e248e25a91bc306c5f2c/README.md) |
| 24&#160;May&#160;23&#160;14:49&#160;UTC | SHAKEN Nextiva, Inc 914H | 23&#160;May&#160;24&#160;14:49&#160;UTC | true | [view](CERTS/ef77b45b37fa412f05b53660a4b60f37962a9be4587ec39211e94289a0087c20/README.md) |
| 25&#160;May&#160;23&#160;15:00&#160;UTC | SHAKEN Midwest Telecom of America 919A | 24&#160;May&#160;24&#160;15:00&#160;UTC | true | [view](CERTS/75298e172e11d29eaefdfbc7758fbe988c29e3380097086cc8214d6f8d09636c/README.md) |
| 01&#160;Jun&#160;23&#160;19:33&#160;UTC | SHAKEN Medtel Communications 994J | 31&#160;May&#160;24&#160;19:33&#160;UTC | true | [view](CERTS/e4b9231f5ad017174e74400e163effa66767d53af72cf608ba3046a0e639a813/README.md) |
| 20&#160;Jun&#160;23&#160;20:56&#160;UTC | SHAKEN Dynalink Communications Inc 991D | 19&#160;Jun&#160;24&#160;20:56&#160;UTC | true | [view](CERTS/f78e86abb18b515b725f81ef27a39158f381acfdaa05f352572611cbfec7e93f/README.md) |
| 22&#160;Jun&#160;23&#160;14:29&#160;UTC | SHAKEN Peachnet LLC 616K | 21&#160;Jun&#160;24&#160;14:29&#160;UTC | true | [view](CERTS/ab9d350cec16b5212bc010e5637396f92c2089704a3f985541d175a1f4b2a02a/README.md) |
| 22&#160;Jun&#160;23&#160;18:42&#160;UTC | SHAKEN FastCast Networks 318K | 21&#160;Jun&#160;24&#160;18:42&#160;UTC | true | [view](CERTS/13723330e372f36b27ca99e1ae23cae3ff5745d92ca6174c3987aef15dc26351/README.md) |
| 22&#160;Jun&#160;23&#160;19:59&#160;UTC | SHAKEN Connexum LLC 203K | 21&#160;Jun&#160;24&#160;19:59&#160;UTC | true | [view](CERTS/d1f5789f3d44bf7545537a3539f0f5dbf43de9cca281266af09943e5635fecfd/README.md) |
| 27&#160;Jun&#160;23&#160;11:43&#160;UTC | SHAKEN Bek Communications Cooperative 1604 | 26&#160;Jun&#160;24&#160;11:43&#160;UTC | true | [view](CERTS/b2f1cb0fa492baaa28a3591bc818d9946da92558364d772b2ddce7ec0028b4d0/README.md) |
| 29&#160;Jun&#160;23&#160;00:15&#160;UTC | SHAKEN Cascabel Networks 536K | 28&#160;Jun&#160;24&#160;00:15&#160;UTC | true | [view](CERTS/d49bf3ecfefc71259bf47bfcc24a2fc5abc0e50fe3c9727b6e920d284c98b53f/README.md) |
| 03&#160;Jul&#160;23&#160;16:39&#160;UTC | SHAKEN Mercury Access Solutions 634K | 02&#160;Jul&#160;24&#160;16:39&#160;UTC | true | [view](CERTS/bbfd747cc6eb37a02ef174b5ada462124ba7990aae403e63cf86516d1fdfb816/README.md) |
| 11&#160;Jul&#160;23&#160;21:11&#160;UTC | SHAKEN Dalton Utilities 3139 | 01&#160;Jan&#160;24&#160;21:11&#160;UTC | true | [view](CERTS/44ee90926a740c783c607a1afc3c73605c6347147f338fd0eca781ad16721246/README.md) |
| 17&#160;Jul&#160;23&#160;14:20&#160;UTC | SHAKEN Identidad Advertising Development LLC 617K | 13&#160;Jan&#160;24&#160;14:20&#160;UTC | true | [view](CERTS/405a42b26de76c3ca814dd2fb735a3abbd4bf028c0937384be59d4a80c4bff44/README.md) |
| 20&#160;Jul&#160;23&#160;18:52&#160;UTC | SHAKEN IPBTel 535K | 19&#160;Jul&#160;24&#160;18:52&#160;UTC | true | [view](CERTS/297dc6000e9cc8107a9b6069afd9e8548ca640ba1c4119d9c6187e5d7258def4/README.md) |
| 11&#160;Aug&#160;23&#160;11:45&#160;UTC | SHAKEN Call Hub Inc. 688K | 10&#160;Aug&#160;24&#160;11:45&#160;UTC | true | [view](CERTS/598f41e0ff84df174314d76b406d85a0f3875aa68f1266c268b98c22f4ee912b/README.md) |
| 22&#160;Aug&#160;23&#160;00:06&#160;UTC | SHAKEN InPhonex 1821 | 21&#160;Aug&#160;24&#160;00:06&#160;UTC | true | [view](CERTS/c5b74544beeefab58f10794e04e2602f84bbb6f397cb81700706983e8e1e128c/README.md) |
| 05&#160;Sep&#160;23&#160;15:54&#160;UTC | SHAKEN Go Voip Dialing LLC 704K | 04&#160;Sep&#160;24&#160;15:54&#160;UTC | true | [view](CERTS/645767c01de0509deb545953141dbf2e136665480917cc8c4ecc73a45608af68/README.md) |
| 08&#160;Sep&#160;23&#160;12:28&#160;UTC | SHAKEN Contactivity Corp. 711K | 07&#160;Sep&#160;24&#160;12:28&#160;UTC | true | [view](CERTS/673f3091743237809463817a98f68fb9fd95c3b112030e4fcbe84096f54d5038/README.md) |
| 15&#160;Sep&#160;23&#160;14:40&#160;UTC | SHAKEN Miracle Telecom 496K | 14&#160;Sep&#160;24&#160;14:40&#160;UTC | true | [view](CERTS/f7312e8a32e80db109fe012d1e00c252afc4eec07f6292fa6f714e926910d14e/README.md) |
| 26&#160;Sep&#160;23&#160;13:26&#160;UTC | SHAKEN Sangoma 777G | 25&#160;Sep&#160;24&#160;13:26&#160;UTC | true | [view](CERTS/2db8f92946ee707e88748612d6ea1b786c7137a8c56e53c80d9484911a2ed51d/README.md) |
| 26&#160;Sep&#160;23&#160;15:59&#160;UTC | SHAKEN Inventive Labs Corp 649J | 24&#160;Mar&#160;24&#160;15:59&#160;UTC | true | [view](CERTS/5eea3beab5e9c7e63c845ab1e34e0ab46e8e8f3fdd61dcaa5d8e9db000aad0fe/README.md) |
| 26&#160;Sep&#160;23&#160;20:51&#160;UTC | SHAKEN Mercury Network Corporation 046K | 12&#160;Jun&#160;24&#160;20:51&#160;UTC | true | [view](CERTS/97be10182ad3307bac843d16be3dbc14d65a71d7f236fa65e802f748330eaa9a/README.md) |
| 27&#160;Sep&#160;23&#160;17:50&#160;UTC | SHAKEN Double A Solutions 644K | 26&#160;Sep&#160;24&#160;17:50&#160;UTC | true | [view](CERTS/fb18701e25fb45d49e3966dd64fb3a518f6f2a9150059b39ad61632ac9317922/README.md) |
| 01&#160;Oct&#160;23&#160;00:00&#160;UTC | SHAKEN IP Link Telecom Inc. 902J | 30&#160;Dec&#160;23&#160;00:00&#160;UTC | true | [view](CERTS/d1baabc43f3d24a4230d8e1d2cbe0da86db660ebc469c93bda23e5d1b5bfcb38/README.md) |
| 02&#160;Oct&#160;23&#160;15:22&#160;UTC | SHAKEN TeleVoIPs 138K | 01&#160;Oct&#160;24&#160;15:22&#160;UTC | true | [view](CERTS/8df110d0b382bc6c2ffd95e4adde6fa10f679b4fdb578f0dc8163b7e06be7635/README.md) |
| 06&#160;Oct&#160;23&#160;13:25&#160;UTC | SHAKEN Kloud 7 LLC 214K | 05&#160;Oct&#160;24&#160;13:25&#160;UTC | true | [view](CERTS/7105ca55bae14fd5fc5952beff1664e0b13f3b560dc07c777a8293f5d06ab752/README.md) |
| 17&#160;Oct&#160;23&#160;22:55&#160;UTC | SHAKEN ConnectMeVoice 719J | 16&#160;Oct&#160;24&#160;22:55&#160;UTC | true | [view](CERTS/5a48a80de9440625305bf4a9af18bcf158e03286df27646af231d75e0cec315d/README.md) |
| 18&#160;Oct&#160;23&#160;22:36&#160;UTC | SHAKEN VOICE1 LLC 265K | 17&#160;Oct&#160;24&#160;22:36&#160;UTC | true | [view](CERTS/f40966bfe5a54cc2e71f5942314743411e143934deeeb8b87647eb1b5ac03aa2/README.md) |
| 19&#160;Oct&#160;23&#160;11:43&#160;UTC | SHAKEN Talk IT Pro 321K | 18&#160;Oct&#160;24&#160;11:43&#160;UTC | true | [view](CERTS/ac46b1e519197ae9ce860b54ca7b6a7150d4be7de7e581e7481003737cf28f68/README.md) |
| 23&#160;Oct&#160;23&#160;00:00&#160;UTC | SHAKEN XCast Labs 689J | 24&#160;Nov&#160;23&#160;00:00&#160;UTC | true | [view](CERTS/f364b4d3290509430aa72f92e06f47b2e4b92112d9eecde416c184bab8c3abbc/README.md) |
| 23&#160;Oct&#160;23&#160;03:56&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 22&#160;Nov&#160;23&#160;03:56&#160;UTC | true | [view](CERTS/9dce58e2b6d1f0368222d0893686bf70d5b0d8581286cbfc949be925af0d1899/README.md) |
| 23&#160;Oct&#160;23&#160;04:10&#160;UTC | SHAKEN BareTelecom 864J | 22&#160;Nov&#160;23&#160;04:10&#160;UTC | true | [view](CERTS/6b8bef3a80a2ab446b186f1400d52ef58bc0051707e2951eb06b4d20c8d7d4ed/README.md) |
| 23&#160;Oct&#160;23&#160;10:44&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 22&#160;Nov&#160;23&#160;10:44&#160;UTC | true | [view](CERTS/fd0a26403e17587862f88ae178a6c83bec359df99dc5eb126f9656a14d79a408/README.md) |
| 23&#160;Oct&#160;23&#160;14:41&#160;UTC | SHAKEN Zray Technologies Corporation 862J | 22&#160;Nov&#160;23&#160;14:41&#160;UTC | true | [view](CERTS/5b1596da8d4551f732745d471df62d958c169ce05ef34c881fa574faca8cb439/README.md) |
| 23&#160;Oct&#160;23&#160;15:12&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 22&#160;Nov&#160;23&#160;15:12&#160;UTC | true | [view](CERTS/929a5fb532e31549f5c31fe01f9a943e8afe19fcc974b01d97152427830e3cd0/README.md) |
| 23&#160;Oct&#160;23&#160;15:32&#160;UTC | SHAKEN IPSBS Managed Services LLC 828J | 22&#160;Nov&#160;23&#160;15:32&#160;UTC | true | [view](CERTS/bd0a213bb641559e345c8bcb73b061d9549f4b25b01de6dc9f1c249ab4ed4842/README.md) |
| 23&#160;Oct&#160;23&#160;15:54&#160;UTC | SHAKEN Comtalk Telecom 705K | 22&#160;Oct&#160;24&#160;15:54&#160;UTC | true | [view](CERTS/05f3a08f601f71f6b02f390050da6acb922f75b5f9d9178b637e96de52847674/README.md) |
| 23&#160;Oct&#160;23&#160;17:32&#160;UTC | SHAKEN Socket Telecom LLC 554a | 22&#160;Nov&#160;23&#160;17:32&#160;UTC | true | [view](CERTS/0f17d8ec6f7f333946fbbc1059f015d9788bd70b211e5a787db9900f0be7fb19/README.md) |
| 23&#160;Oct&#160;23&#160;18:02&#160;UTC | SHAKEN IDT America, Corp 363A | 22&#160;Nov&#160;23&#160;18:02&#160;UTC | true | [view](CERTS/af46c3c1fd71609ed88b4035e5de849abb47c74376bf2ced5b077f1631e4ec1f/README.md) |
| 23&#160;Oct&#160;23&#160;18:58&#160;UTC | SHAKEN IDT America, Corp 363A | 22&#160;Nov&#160;23&#160;18:58&#160;UTC | true | [view](CERTS/9ad76b8cbd15f1762e22209a4bb60e8d34d53f263199b5e53ac7bfa45a70d73c/README.md) |
| 24&#160;Oct&#160;23&#160;04:05&#160;UTC | SHAKEN BareTelecom 864J | 23&#160;Nov&#160;23&#160;04:05&#160;UTC | true | [view](CERTS/1693b44189ee56e4a178fbfef92325b84e0fe17942c573a01ea154278e97e6e4/README.md) |
| 24&#160;Oct&#160;23&#160;10:16&#160;UTC | SHAKEN InteractiveTel, LLC 920J | 23&#160;Nov&#160;23&#160;10:16&#160;UTC | true | [view](CERTS/459df7745429fb1b52a5ae28d6231bb9cc41a0474d77c80a2b7a619072b960be/README.md) |
| 24&#160;Oct&#160;23&#160;10:40&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 23&#160;Nov&#160;23&#160;10:40&#160;UTC | true | [view](CERTS/ac98afb0780dd6ee0b0cfc64a0af071e6c5170b1b8d180979ccb21c7b0b44eda/README.md) |
| 24&#160;Oct&#160;23&#160;15:07&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 23&#160;Nov&#160;23&#160;15:07&#160;UTC | true | [view](CERTS/837dc9311e312b893c77a2f949ef1e72cad0699bfb1e2e025bcb0e97536197d9/README.md) |
| 24&#160;Oct&#160;23&#160;16:35&#160;UTC | SHAKEN NTC International, INC 016K | 23&#160;Nov&#160;23&#160;16:35&#160;UTC | true | [view](CERTS/f49c4b72f138aa526c175cae6c715031679ad0f7fac0c2fbe5fbb7ec9ed594f8/README.md) |
| 24&#160;Oct&#160;23&#160;17:27&#160;UTC | SHAKEN Socket Telecom LLC 554a | 23&#160;Nov&#160;23&#160;17:27&#160;UTC | true | [view](CERTS/26a97afe6b64491ec50f879372acc0a8ef170d1088b6adae7a85d6991a6586f6/README.md) |
| 24&#160;Oct&#160;23&#160;17:57&#160;UTC | SHAKEN IDT America, Corp 363A | 23&#160;Nov&#160;23&#160;17:57&#160;UTC | true | [view](CERTS/6e5db8c30856104444214b0403dc0fbf3de86319d5ed7a50edb2b74c4a0b1591/README.md) |
| 24&#160;Oct&#160;23&#160;18:52&#160;UTC | SHAKEN IDT America, Corp 363A | 23&#160;Nov&#160;23&#160;18:52&#160;UTC | true | [view](CERTS/540f210f00e252eca6e3a49de1a5f4d34b7ccff5e9bb7b4805ce16832d71cb67/README.md) |
| 24&#160;Oct&#160;23&#160;20:55&#160;UTC | SHAKEN IDT America, Corp 363A | 23&#160;Nov&#160;23&#160;20:55&#160;UTC | true | [view](CERTS/5e1b03fd689eed4c093b29b9b7d0f4487eea7137d110490ff2f0a65af7671aff/README.md) |
| 25&#160;Oct&#160;23&#160;04:00&#160;UTC | SHAKEN BareTelecom 864J | 24&#160;Nov&#160;23&#160;04:00&#160;UTC | true | [view](CERTS/cdb58c58d263486a39b73bd257ac7b8f180947b8417efdc75b872aac1309d394/README.md) |
| 25&#160;Oct&#160;23&#160;10:12&#160;UTC | SHAKEN InteractiveTel, LLC 920J | 24&#160;Nov&#160;23&#160;10:12&#160;UTC | true | [view](CERTS/1a0e980a5f480efc33b53bdb6084b18f7d4f15622b5e10a1e89bf8c747f46565/README.md) |
| 25&#160;Oct&#160;23&#160;10:47&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 24&#160;Nov&#160;23&#160;10:47&#160;UTC | true | [view](CERTS/0e313207df2047abe7c048b6c5f7d7c93767a4df067957ab126832c16405b158/README.md) |
| 25&#160;Oct&#160;23&#160;14:31&#160;UTC | SHAKEN Zray Technologies Corporation 862J | 24&#160;Nov&#160;23&#160;14:31&#160;UTC | true | [view](CERTS/bb38abfcc9773da1bdd4fb6aa42b2a47f6386adb49d890185f771aeab15bd00e/README.md) |
| 25&#160;Oct&#160;23&#160;15:02&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 24&#160;Nov&#160;23&#160;15:02&#160;UTC | true | [view](CERTS/af8621167b0e6693260a7adf568302846030a7972521441517e81d882a02c05c/README.md) |
| 25&#160;Oct&#160;23&#160;15:22&#160;UTC | SHAKEN IPSBS Managed Services LLC 828J | 24&#160;Nov&#160;23&#160;15:22&#160;UTC | true | [view](CERTS/c723e8ccf5376a107dc5d0f5160fd8ac1c096770d993ae18638a7c09aa964f95/README.md) |
| 25&#160;Oct&#160;23&#160;15:38&#160;UTC | SHAKEN 1stPoint Communications, LLC 463G | 24&#160;Nov&#160;23&#160;15:38&#160;UTC | true | [view](CERTS/4cbfa0cc3351291afd9816b6fcd5da0692155c486c2de4e24ca25e1ad5ed84ef/README.md) |
| 25&#160;Oct&#160;23&#160;20:09&#160;UTC | SHAKEN Ytel Inc. 703J | 24&#160;Oct&#160;24&#160;20:09&#160;UTC | true | [view](CERTS/1abf91916a7c83ff53e58fdaf8fab1e6f6c232bc2aa2b1903c76c37d7a6984a6/README.md) |
| 26&#160;Oct&#160;23&#160;14:57&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 25&#160;Nov&#160;23&#160;14:57&#160;UTC | true | [view](CERTS/c9bde356922d775f1bf8edbcf64e276170f9f75f83588620fb4a4f56f8e8ed23/README.md) |
| 26&#160;Oct&#160;23&#160;15:17&#160;UTC | SHAKEN IPSBS Managed Services LLC 828J | 25&#160;Nov&#160;23&#160;15:17&#160;UTC | true | [view](CERTS/8a4373d79e08a1bb6f8e2eef12d70948767eb095db88c67aa0b14f2e9143daa9/README.md) |
| 26&#160;Oct&#160;23&#160;15:33&#160;UTC | SHAKEN 1stPoint Communications, LLC 463G | 25&#160;Nov&#160;23&#160;15:33&#160;UTC | true | [view](CERTS/87e4ae71ef0ab8c07a693189a2fe0d03a66a9c24de356366b92f5e6fa823c57b/README.md) |
| 26&#160;Oct&#160;23&#160;15:42&#160;UTC | SHAKEN i3 Broadband 5800 | 25&#160;Nov&#160;23&#160;15:42&#160;UTC | true | [view](CERTS/695c3e6c9855eb302a174d778797ec481154f90e19baeabfb2225d8d948e1a17/README.md) |
| 26&#160;Oct&#160;23&#160;16:26&#160;UTC | SHAKEN NTC International, INC 016K | 25&#160;Nov&#160;23&#160;16:26&#160;UTC | true | [view](CERTS/d4156f86d1b8f160b903ca0deb73475638ff8a103997aeb17493ca9e2a480c50/README.md) |
| 26&#160;Oct&#160;23&#160;18:43&#160;UTC | SHAKEN IDT America, Corp 363A | 25&#160;Nov&#160;23&#160;18:43&#160;UTC | true | [view](CERTS/bc4bd0b3c293ed05c2fb10a48deab001a1bb2ba7cdef79d0a86736fcc3d3e8f7/README.md) |
| 26&#160;Oct&#160;23&#160;20:45&#160;UTC | SHAKEN IDT America, Corp 363A | 25&#160;Nov&#160;23&#160;20:45&#160;UTC | true | [view](CERTS/9b771873abc7b3e88d09f242f8a20d0567255353ea6366f51e72fbb845896526/README.md) |
| 27&#160;Oct&#160;23&#160;03:50&#160;UTC | SHAKEN BareTelecom 864J | 26&#160;Nov&#160;23&#160;03:50&#160;UTC | true | [view](CERTS/22a5dfe630e666aeb755e866ffb374f8bdc6196d8b12934eaf8b3a4cd05a64b7/README.md) |
| 27&#160;Oct&#160;23&#160;05:40&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 26&#160;Nov&#160;23&#160;05:40&#160;UTC | true | [view](CERTS/dc62ac4ef175b71500590db41da19765f1664e81ec7c30d3f07487ca887b0aa4/README.md) |
| 27&#160;Oct&#160;23&#160;09:33&#160;UTC | SHAKEN Primo Dialler LLC 249K | 26&#160;Oct&#160;24&#160;09:33&#160;UTC | true | [view](CERTS/2345fd9200b5754c1d5fb353ea414d4be8e9eda729202bd84c0c3f7c6a0d1ad6/README.md) |
| 27&#160;Oct&#160;23&#160;10:46&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 26&#160;Nov&#160;23&#160;10:46&#160;UTC | true | [view](CERTS/1d6937ddefc9d2f654a6b0dfcc9eb08d6292571c2e0868104b66e3368766ee2d/README.md) |
| 27&#160;Oct&#160;23&#160;14:52&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 26&#160;Nov&#160;23&#160;14:52&#160;UTC | true | [view](CERTS/767fae03d5d29b031aa21315d75e062ebe9295d647bffe3ccd42951cc7b077a8/README.md) |
| 27&#160;Oct&#160;23&#160;15:28&#160;UTC | SHAKEN 1stPoint Communications, LLC 463G | 26&#160;Nov&#160;23&#160;15:28&#160;UTC | true | [view](CERTS/df0604ecec1f90bfec3c1860eed738adbad378ddf8e2e79230132e5fe2290bb6/README.md) |
| 27&#160;Oct&#160;23&#160;16:20&#160;UTC | SHAKEN NTC International, INC 016K | 26&#160;Nov&#160;23&#160;16:20&#160;UTC | true | [view](CERTS/9066994ddb110cd743fe907d2a63137a736c01c974054d278e26943470c19980/README.md) |
| 27&#160;Oct&#160;23&#160;17:12&#160;UTC | SHAKEN Socket Telecom LLC 554a | 26&#160;Nov&#160;23&#160;17:12&#160;UTC | true | [view](CERTS/0fdb0bfa6c4aa43de1164f148f8f8aaee0aac0c1bc4820255498ca8cce922fe0/README.md) |
| 27&#160;Oct&#160;23&#160;18:38&#160;UTC | SHAKEN IDT America, Corp 363A | 26&#160;Nov&#160;23&#160;18:38&#160;UTC | true | [view](CERTS/391ad73514796d9a14b3ff5168d779d21064378bfe3227ae6c6b820abb9ce76c/README.md) |
| 27&#160;Oct&#160;23&#160;20:37&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 26&#160;Nov&#160;23&#160;20:37&#160;UTC | true | [view](CERTS/4378e509b0ab683797c13bcd29bfc3bc0b3caced45ad1fb0ba29a6d052d9a593/README.md) |
| 27&#160;Oct&#160;23&#160;20:40&#160;UTC | SHAKEN IDT America, Corp 363A | 26&#160;Nov&#160;23&#160;20:40&#160;UTC | true | [view](CERTS/b34cc6beb60813f748c599cc6b801b971b431221389af066337a2c4116a04237/README.md) |
| 27&#160;Oct&#160;23&#160;22:04&#160;UTC | SHAKEN Mitel Cloud Services, Inc. 670J | 26&#160;Oct&#160;24&#160;22:04&#160;UTC | true | [view](CERTS/a3fecbe272ac59a623119a5b64a5e3cc79f9f880c2b82584a79f6f40e7f1562a/README.md) |
| 27&#160;Oct&#160;23&#160;22:26&#160;UTC | SHAKEN Mitel Cloud Services, Inc. 670J | 26&#160;Oct&#160;24&#160;22:26&#160;UTC | true | [view](CERTS/6ce3263c3dd426c9f8f57b7fabf533b30c5cf3c3582f7c5cd6f970b0ed0c01b1/README.md) |
| 30&#160;Oct&#160;23&#160;14:00&#160;UTC | SHAKEN Rayfield Communications, Inc. 006K | 29&#160;Nov&#160;23&#160;14:00&#160;UTC | true | [view](CERTS/6bd7eb6c9f2bc8622d5741b8d177f4ef7d6b0d2601f7c639d2bbf97d3ad27ad4/README.md) |
| 30&#160;Oct&#160;23&#160;14:03&#160;UTC | SHAKEN Arbeit 816J | 29&#160;Nov&#160;23&#160;14:03&#160;UTC | true | [view](CERTS/19302fe2683c1e9d476d7c73442b6a8be47400527a27f604e29247473bba6801/README.md) |
| 30&#160;Oct&#160;23&#160;14:05&#160;UTC | SHAKEN California Telecom 319K | 29&#160;Nov&#160;23&#160;14:05&#160;UTC | true | [view](CERTS/12f75904f75221e0207d53a08b97e2752726e5c63787eedf9235a246cd325648/README.md) |
| 30&#160;Oct&#160;23&#160;14:07&#160;UTC | SHAKEN DLS Internet Services 815J | 29&#160;Nov&#160;23&#160;14:07&#160;UTC | true | [view](CERTS/c22004d54103b1b05b537e938c2a606c4f2c7ebc382bf3743cb8eb2117600295/README.md) |
| 30&#160;Oct&#160;23&#160;14:11&#160;UTC | SHAKEN Systemverse, LLC. 194K | 29&#160;Nov&#160;23&#160;14:11&#160;UTC | true | [view](CERTS/6a11eb269c46b9ab489356f3a10985d7191cc8193c8600d336df77d8bac82f12/README.md) |
| 30&#160;Oct&#160;23&#160;14:15&#160;UTC | SHAKEN Vinculum Communications, Inc 787J | 29&#160;Oct&#160;24&#160;14:15&#160;UTC | true | [view](CERTS/2d3cf73af0f77cf310d869e984084841f7b2ebb149a0d3d09694dfa5838b074b/README.md) |
| 31&#160;Oct&#160;23&#160;16:23&#160;UTC | SHAKEN ConvergeTel LLC 388K | 30&#160;Oct&#160;24&#160;16:23&#160;UTC | true | [view](CERTS/886577ab6ce0cf85d3ae954a4bf47c8ebf47246340a1775e0c7bca8b6e267244/README.md) |
| 01&#160;Nov&#160;23&#160;02:29&#160;UTC | SHAKEN Lightspeed Voice 557F | 01&#160;Dec&#160;23&#160;02:29&#160;UTC | true | [view](CERTS/61d51d9b0e3fdc8335ef1ba24833a3296e17722060bb34dc4c549aaa936e2b56/README.md) |
| 02&#160;Nov&#160;23&#160;13:49&#160;UTC | SHAKEN Starlinq PBX Inc. 267K | 01&#160;Nov&#160;24&#160;13:49&#160;UTC | true | [view](CERTS/772495eb19f14de884567979521029ea8d0f85712aae22b5dea6ffadb6fd6aaf/README.md) |
| 03&#160;Nov&#160;23&#160;04:43&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 03&#160;Dec&#160;23&#160;04:43&#160;UTC | true | [view](CERTS/7d217ffbb8f2a4ea5ab39134ce550139f63bb54de724454eadc405ad81f679b4/README.md) |
| 06&#160;Nov&#160;23&#160;04:28&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 06&#160;Dec&#160;23&#160;04:28&#160;UTC | true | [view](CERTS/38a1069f3140d48fbd9eb3a1bf613d797d8568a4fe6b7ec4332e3a7f07de061f/README.md) |
| 07&#160;Nov&#160;23&#160;15:55&#160;UTC | SHAKEN Carrier One Inc. 705J | 06&#160;Nov&#160;24&#160;15:55&#160;UTC | true | [view](CERTS/89208e1662d350e58c5f5da80982aa62487d1cbf640a574d6d2ce7ebbf17a2fb/README.md) |
| 07&#160;Nov&#160;23&#160;16:22&#160;UTC | SHAKEN Threshold Communications Inc 563J | 07&#160;Dec&#160;23&#160;16:22&#160;UTC | true | [view](CERTS/3fd34d37c264c4d6d83c3dae5d3e522d76092dc85c469957ad342bb2790fc0e6/README.md) |
| 07&#160;Nov&#160;23&#160;16:23&#160;UTC | SHAKEN Broadband Dynamics LLC 583j | 07&#160;Dec&#160;23&#160;16:23&#160;UTC | true | [view](CERTS/e72cdb3f7f157a4c1c2f077ee961b97059b68167af770679273defefaf137cc7/README.md) |
| 07&#160;Nov&#160;23&#160;16:26&#160;UTC | SHAKEN Consolidated Communications 5113 | 07&#160;Dec&#160;23&#160;16:26&#160;UTC | true | [view](CERTS/7aefd45174ca06235229e9cf8bf32a9f912fa71b843a18f31cf288d60ec5f592/README.md) |
| 07&#160;Nov&#160;23&#160;16:29&#160;UTC | SHAKEN Touchtone 683A | 07&#160;Dec&#160;23&#160;16:29&#160;UTC | true | [view](CERTS/2e3d0ce1f0cb48d2eac17ed0c0a73a05494e7fa56ee764d97eb7ede51db428c1/README.md) |
| 07&#160;Nov&#160;23&#160;16:30&#160;UTC | SHAKEN Apeiron Systems 012J | 07&#160;Dec&#160;23&#160;16:30&#160;UTC | true | [view](CERTS/312d7c315dc3d17147304d28655b0ac5e09aa986053646e0fce0f66c729e5c3d/README.md) |
| 07&#160;Nov&#160;23&#160;16:33&#160;UTC | SHAKEN Fonative, Inc. 684J | 07&#160;Dec&#160;23&#160;16:33&#160;UTC | true | [view](CERTS/8d7f84322d0694ecf6277b89f6bc7ffe3a487f25baa395a41b637d22d23fa00b/README.md) |
| 07&#160;Nov&#160;23&#160;16:34&#160;UTC | SHAKEN IPitomy 652J | 07&#160;Dec&#160;23&#160;16:34&#160;UTC | true | [view](CERTS/671c0fc0d64d09e6ad78b5468c73d4a61e5e027f759f23e5f9cbcd41a985dc76/README.md) |
| 07&#160;Nov&#160;23&#160;16:35&#160;UTC | SHAKEN Phone.com, Inc. 633J | 07&#160;Dec&#160;23&#160;16:35&#160;UTC | true | [view](CERTS/5907798637f4201c99dcbd9552bd5dc1913ccd904c7556f13e53d0e6965dac5b/README.md) |
| 07&#160;Nov&#160;23&#160;16:37&#160;UTC | SHAKEN NETRIO LLC 020K | 07&#160;Dec&#160;23&#160;16:37&#160;UTC | true | [view](CERTS/d02b5c19f21e04b8d9ccb2a62d1d40e6128aa42fb571df0263d005383c634022/README.md) |
| 07&#160;Nov&#160;23&#160;16:38&#160;UTC | SHAKEN Impulse 250G | 07&#160;Dec&#160;23&#160;16:38&#160;UTC | true | [view](CERTS/cfdf3f24ce7d6c32fe786fa0e08215ef2ebac156dd7588a647335b827f2e64de/README.md) |
| 07&#160;Nov&#160;23&#160;16:39&#160;UTC | SHAKEN Noble Systems Communications LLC 187J | 07&#160;Dec&#160;23&#160;16:39&#160;UTC | true | [view](CERTS/8c4918ac5b52df996b4458fe3df4b956ec193529b26235022822987fcf4b35ae/README.md) |
| 07&#160;Nov&#160;23&#160;16:45&#160;UTC | SHAKEN Airespring 996H | 07&#160;Dec&#160;23&#160;16:45&#160;UTC | true | [view](CERTS/f454b995db0b3d751d6eee89f4837e23badb4eb609e442e9a47b490e7e72ee3f/README.md) |
| 07&#160;Nov&#160;23&#160;16:46&#160;UTC | SHAKEN Nobelbiz, Inc. 596J | 07&#160;Dec&#160;23&#160;16:46&#160;UTC | true | [view](CERTS/4d0d3c8d42b4964318c967b69c7fcec182e56898491e1fde666f4f31984c43ba/README.md) |
| 07&#160;Nov&#160;23&#160;16:47&#160;UTC | SHAKEN Momentum Telecom 1417 | 07&#160;Dec&#160;23&#160;16:47&#160;UTC | true | [view](CERTS/a70ed94542005801728ab4c62335f820e099fb9636f7060384398f3c944e29b9/README.md) |
| 07&#160;Nov&#160;23&#160;16:47&#160;UTC | SHAKEN Momentum Telecom 9157 | 07&#160;Dec&#160;23&#160;16:47&#160;UTC | true | [view](CERTS/4ec895268e04cebf2bfcd2b398ee50a5d8102b8a060ca6ddaec264f478457559/README.md) |
| 07&#160;Nov&#160;23&#160;16:49&#160;UTC | SHAKEN Terra Nova Telecom 382G | 07&#160;Dec&#160;23&#160;16:49&#160;UTC | true | [view](CERTS/759eac995630578c46c23b9e7dfff54f374c91b715257fc43f511cb4eaf1b5a3/README.md) |
| 07&#160;Nov&#160;23&#160;16:49&#160;UTC | SHAKEN Matrix 3058 | 07&#160;Dec&#160;23&#160;16:49&#160;UTC | true | [view](CERTS/18e02c5101799dba40ccfabb942583674333e5914f075427badd54cabc7f469c/README.md) |
| 07&#160;Nov&#160;23&#160;16:50&#160;UTC | SHAKEN Matrix 7379 | 07&#160;Dec&#160;23&#160;16:50&#160;UTC | true | [view](CERTS/306dd68622498cab3d28f02ef0fbda6d9b98fd33b15526ee4f8e13cbc0dc0e93/README.md) |
| 07&#160;Nov&#160;23&#160;16:50&#160;UTC | SHAKEN Matrix 9451 | 07&#160;Dec&#160;23&#160;16:50&#160;UTC | true | [view](CERTS/f09132d6c0128508fcc6b3dad3e0d8087e915212de13b5ba3098df63531be282/README.md) |
| 07&#160;Nov&#160;23&#160;17:26&#160;UTC | SHAKEN PNG Telecommunications Inc 3395 | 07&#160;Dec&#160;23&#160;17:26&#160;UTC | true | [view](CERTS/737272d371de73a464eaf25c76ff1598a00e273118acd0047c12d365476bbd50/README.md) |
| 07&#160;Nov&#160;23&#160;17:28&#160;UTC | SHAKEN Magna5, LLC 3849 | 07&#160;Dec&#160;23&#160;17:28&#160;UTC | true | [view](CERTS/e20ed45847c806a56bff1297cd0243b28afc90409f2dcd821f25a79ed187d6d5/README.md) |
| 07&#160;Nov&#160;23&#160;17:30&#160;UTC | SHAKEN Magna5, LLC 8249 | 07&#160;Dec&#160;23&#160;17:30&#160;UTC | true | [view](CERTS/afe721747365649605537bae6f329efe498b39154f44cb48b95ec3e12478c7e1/README.md) |
| 08&#160;Nov&#160;23&#160;01:54&#160;UTC | SHAKEN Lightspeed Voice 557F | 08&#160;Dec&#160;23&#160;01:54&#160;UTC | true | [view](CERTS/4884a2d92c2d7297b680e41065e1e3528eb767bf8e6f121a603d0ea68a4ad80c/README.md) |
| 08&#160;Nov&#160;23&#160;05:30&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 08&#160;Dec&#160;23&#160;05:30&#160;UTC | true | [view](CERTS/d152636c4682fd55f00a014d4ddbea2e85f032b5d5656dce591122020cec56e0/README.md) |
| 09&#160;Nov&#160;23&#160;01:49&#160;UTC | SHAKEN Lightspeed Voice 557F | 09&#160;Dec&#160;23&#160;01:49&#160;UTC | true | [view](CERTS/7b6692d2a0817d90cc280d96ea69cfe24fcb001bb70786b0469738fcef65dfc1/README.md) |
| 09&#160;Nov&#160;23&#160;05:25&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 09&#160;Dec&#160;23&#160;05:25&#160;UTC | true | [view](CERTS/e5ba7957dbf771424d3b1ec5546af1e4317169bda8bcfaedbd93d38ac8252cc1/README.md) |
| 10&#160;Nov&#160;23&#160;02:10&#160;UTC | SHAKEN Star2Star Communications, LLC 590J | 09&#160;Nov&#160;24&#160;02:10&#160;UTC | true | [view](CERTS/d8197ed17be0c4415c77edbda2946e9c03f29dd67299927e8ecc654794c83951/README.md) |
| 10&#160;Nov&#160;23&#160;05:20&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 10&#160;Dec&#160;23&#160;05:20&#160;UTC | true | [view](CERTS/9b5560a5b2ab109dec26eab5da79461f5a4ed1aed7cfcf3edc0a07d6aee5f9a5/README.md) |
| 11&#160;Nov&#160;23&#160;04:03&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 11&#160;Dec&#160;23&#160;04:03&#160;UTC | true | [view](CERTS/1c9ffa257dcb95ede977bd380ed7f496e7434233507e8dce78e07ebee7564fa1/README.md) |
| 13&#160;Nov&#160;23&#160;05:05&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 13&#160;Dec&#160;23&#160;05:05&#160;UTC | true | [view](CERTS/bb10f16805f129b5927464e92c6edbf61035b181ee44406adf6fb7dcbadf6819/README.md) |
| 13&#160;Nov&#160;23&#160;16:45&#160;UTC | SHAKEN Technology Innovation Lab 599J | 12&#160;Nov&#160;24&#160;16:45&#160;UTC | true | [view](CERTS/c71e376f24ac69d9010c19ffbcdacc9b1b0286b63dc216f70f624d1269c3719d/README.md) |
| 14&#160;Nov&#160;23&#160;05:00&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 14&#160;Dec&#160;23&#160;05:00&#160;UTC | true | [view](CERTS/eeeda779549bf70c1e368c2b5a408c93427c7c75fcb61e8727b2f2affc8df4b7/README.md) |
| 14&#160;Nov&#160;23&#160;13:29&#160;UTC | SHAKEN Asia Pacific Network 988J | 13&#160;Nov&#160;24&#160;13:29&#160;UTC | true | [view](CERTS/12b9440a58d22c821e164016a60deccd67ed5036e2244bd1694f39ce326e5811/README.md) |
| 14&#160;Nov&#160;23&#160;13:34&#160;UTC | SHAKEN Current Calls, LLC 746J | 13&#160;Nov&#160;24&#160;13:34&#160;UTC | true | [view](CERTS/a1db34e8f1daf7c5246c10508cf9dd97a3ab166f0ad7304ba124cdd53f199143/README.md) |
| 14&#160;Nov&#160;23&#160;13:36&#160;UTC | SHAKEN Godaddy 463K | 13&#160;Nov&#160;24&#160;13:36&#160;UTC | true | [view](CERTS/cc4ebe2a97d75e8560b603d277ce9e5bfb7a1e8dd5da0628e64e90576d270750/README.md) |
| 14&#160;Nov&#160;23&#160;13:38&#160;UTC | SHAKEN ONE OWL TELECOM INC 412K | 13&#160;Nov&#160;24&#160;13:38&#160;UTC | true | [view](CERTS/a8a6cbe884baa1d42768251df5233f7ffd0b33886539cf00df0d1b16cc48c3f3/README.md) |
| 14&#160;Nov&#160;23&#160;13:38&#160;UTC | SHAKEN OneStream Networks, LLC 630J | 13&#160;Nov&#160;24&#160;13:38&#160;UTC | true | [view](CERTS/d0a3700b7239c2f2aac6b3bbbac04b98db8bd5a8f725c2ebd1f8aed90126a4e1/README.md) |
| 14&#160;Nov&#160;23&#160;13:43&#160;UTC | SHAKEN Ringfree Communications Inc 317K | 13&#160;Nov&#160;24&#160;13:43&#160;UTC | true | [view](CERTS/066d5e0f2e11207f9a8aab5ae03e7bb3877292af0f325a690ce9cff6a9db42f9/README.md) |
| 14&#160;Nov&#160;23&#160;13:45&#160;UTC | SHAKEN Vumber LLC 225K | 13&#160;Nov&#160;24&#160;13:45&#160;UTC | true | [view](CERTS/4c72b9f5cbad420cbbb77aa0ca320baae202029d469d1953cb7e3b84ea42110d/README.md) |
| 15&#160;Nov&#160;23&#160;04:55&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 15&#160;Dec&#160;23&#160;04:55&#160;UTC | true | [view](CERTS/21d053255499360069ba61e07806bb91185485e839faf2d4d2bf0c8ab43eb7fc/README.md) |
| 15&#160;Nov&#160;23&#160;07:03&#160;UTC | SHAKEN BareTelecom 864J | 15&#160;Dec&#160;23&#160;07:03&#160;UTC | true | [view](CERTS/e40d2f4848c64d6c5a92244e07beba256987363e920936c0f60a9374e03efd80/README.md) |
| 15&#160;Nov&#160;23&#160;18:42&#160;UTC | SHAKEN Xchange Telecom LLC 325B | 14&#160;Nov&#160;24&#160;18:42&#160;UTC | true | [view](CERTS/25a16a8e688b969af60b0be46ce6b3d6b2525148559e5e60734bcf64f2afb635/README.md) |
| 16&#160;Nov&#160;23&#160;04:50&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 16&#160;Dec&#160;23&#160;04:50&#160;UTC | true | [view](CERTS/c2cf8d2b2646d593471df3f582130e38b8992face569e4db1ec8a66f7c3c2111/README.md) |
| 16&#160;Nov&#160;23&#160;06:58&#160;UTC | SHAKEN BareTelecom 864J | 16&#160;Dec&#160;23&#160;06:58&#160;UTC | true | [view](CERTS/2b044fd24aaef56d124df614a9e7c40871eccc53bd7061b48b4f3793e362f65e/README.md) |
| 16&#160;Nov&#160;23&#160;18:15&#160;UTC | SHAKEN Convoso 758J | 21&#160;Dec&#160;23&#160;18:15&#160;UTC | true | [view](CERTS/f0417daec81596a69a76756ca06e1247f2199913b94e6f1199f3c38a6d1f1ce1/README.md) |
| 17&#160;Nov&#160;23&#160;03:33&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 17&#160;Dec&#160;23&#160;03:33&#160;UTC | true | [view](CERTS/b70b8f359a8a37e25fffb0b9eb06f7e4aa6083eeb191b08235a5c0e31dbb0a6b/README.md) |
| 17&#160;Nov&#160;23&#160;04:45&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 17&#160;Dec&#160;23&#160;04:45&#160;UTC | true | [view](CERTS/d6e2ac1a7b13f4c80036edd08ad91fe82a0f50dbd5ce6c6c5e036b1be9ccb43e/README.md) |
| 17&#160;Nov&#160;23&#160;04:56&#160;UTC | SHAKEN IPSBS Managed Services LLC 828J | 17&#160;Dec&#160;23&#160;04:56&#160;UTC | true | [view](CERTS/5c67dddbb909a7683d6673c0fb15514d62e6ee3e7e719e774aa4ebefbaffb5d9/README.md) |
| 17&#160;Nov&#160;23&#160;06:53&#160;UTC | SHAKEN BareTelecom 864J | 17&#160;Dec&#160;23&#160;06:53&#160;UTC | true | [view](CERTS/c6c5bb6403975d415cd8649c947c9ff507f7212814d4a3b86f96bdd8fe1fcff1/README.md) |
| 18&#160;Nov&#160;23&#160;01:51&#160;UTC | SHAKEN Cyberlynk Network, LLC 086K | 18&#160;Dec&#160;23&#160;01:51&#160;UTC | true | [view](CERTS/990854ddf2a19d4d29136b0850753e6a44d6d84033558467f429b6343df2b51d/README.md) |
| 18&#160;Nov&#160;23&#160;03:28&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 18&#160;Dec&#160;23&#160;03:28&#160;UTC | true | [view](CERTS/4bc3c297433c03e0d8f1e814d5df40ea9113967997f67146c70013af578c1fa1/README.md) |
| 20&#160;Nov&#160;23&#160;02:52&#160;UTC | SHAKEN InteractiveTel, LLC 920J | 20&#160;Dec&#160;23&#160;02:52&#160;UTC | true | [view](CERTS/9a33758c3c419eae24501c0f590270f68e00a5efd5348caa83006b35e3aeeeac/README.md) |
| 20&#160;Nov&#160;23&#160;04:30&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 20&#160;Dec&#160;23&#160;04:30&#160;UTC | true | [view](CERTS/337c7ce30c1ec9550e481a19297a6c9a6b5682e4b549a8e982565163940ad6e0/README.md) |
| 20&#160;Nov&#160;23&#160;04:41&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 20&#160;Dec&#160;23&#160;04:41&#160;UTC | true | [view](CERTS/e9f5d4cbd7195889e2ccfbf1e5c37e013a7a783ec98cc724a5bdda1bec499ffc/README.md) |
| 20&#160;Nov&#160;23&#160;06:38&#160;UTC | SHAKEN BareTelecom 864J | 20&#160;Dec&#160;23&#160;06:38&#160;UTC | true | [view](CERTS/02695a30a82b9532871f9bff03110262397badfd7b0db5bb57d86573318250d4/README.md) |
| 21&#160;Nov&#160;23&#160;04:25&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 21&#160;Dec&#160;23&#160;04:25&#160;UTC | true | [view](CERTS/41be7ca4ef3f0468a08d3d1048b605ee32a87d3b38887dcddede3ae020ca0521/README.md) |
| 21&#160;Nov&#160;23&#160;06:33&#160;UTC | SHAKEN BareTelecom 864J | 21&#160;Dec&#160;23&#160;06:33&#160;UTC | true | [view](CERTS/a3f20e79da71d7e4b117eaf2013b9f8cd1ace375df14f4eeee7e8ec400c8e85e/README.md) |
| 21&#160;Nov&#160;23&#160;14:25&#160;UTC | SHAKEN IDT America, Corp 363A | 21&#160;Dec&#160;23&#160;14:25&#160;UTC | true | [view](CERTS/5ea10877a903060a3055f8b35c28d277b36159bbf04de6354a0c14a3803dc2fd/README.md) |

#### CA Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 21&#160;Aug&#160;20&#160;01:22&#160;UTC | SHAKEN Sansay Root CA US | 16&#160;Aug&#160;40&#160;01:22&#160;UTC | false | [view](CERTS/8356d251b255f4ac3fd108bb79be4c02dcea2d3b13d138892bdb3a70cbe6a343/README.md) |
| 02&#160;Sep&#160;22&#160;20:53&#160;UTC | SHAKEN Sansay Intermediate CA US WEST 1 | 31&#160;Aug&#160;29&#160;20:53&#160;UTC | false | [view](CERTS/4b1dfdba2b1e4bbffbf900a20f1f6f7befbef0008b963e4922a64469cb97d24b/README.md) |


Generated: 21 Nov 23 19:18 UTC