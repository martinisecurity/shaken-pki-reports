# STIR/SHAKEN CA Ecosystem Compliance

## Sansay

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 2051 certificates were included in the corpus being tested
- 19 certificates in the corpus were skipped because they are duplicates
- 1826 certificates in the corpus were skipped because they are expired
- 1 certificates in the corpus were skipped because they are not currently trusted
- 205 certificates being tested against the remaining rules
- 2.02 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 0.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 0.98% of certificates are too old to be assessed against currently enforced expectations
- 148 days is the average remaining validity for the certificates in the corpus
- 149 days is the average initial validity for the certificates in the corpus
- 131 certificates expire in the next 30 days
- 2.05 average number of unexpired certificates per OCN observed
- 100 unique OCNs observed in unexpired and valid certificate corpus

| Instances | Test | Source |
|-----------|------|--------|
| 91 | [e_atis_ext_crl_distribution_struct](ISSUES/e_atis_ext_crl_distribution_struct/README.md) | ATIS1000080 |
| 203 | [e_atis_subject_cn_spc](ISSUES/e_atis_subject_cn_spc/README.md) | ATIS1000080 |
| 6 | [e_atis_tn_auth_list_spc_format](ISSUES/e_atis_tn_auth_list_spc_format/README.md) | ATIS1000080 |
| 49 | [e_shaken_certificate_policies_id](ISSUES/e_shaken_certificate_policies_id/README.md) | US_SHAKEN_CP |
| 65 | [e_us_cp_subject_sn_shall](ISSUES/e_us_cp_subject_sn_shall/README.md) | US_SHAKEN_CP |

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
- 5295 days is the average remaining validity for the certificates in the corpus
- 4928 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days

No error, warning, or notice level issues were found

### Details

#### Leaf Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 14&#160;Feb&#160;23&#160;17:12&#160;UTC | SHAKEN Ytel Inc. 703J | 14&#160;Feb&#160;24&#160;17:12&#160;UTC | true | [view](CERTS/14c9bef113cfebe60611b0c56c430518ff8d42e8b98dd7e4653bd9cf619d5641/README.md) |
| 01&#160;Mar&#160;23&#160;14:14&#160;UTC | SHAKEN Voneto 485K | 29&#160;Feb&#160;24&#160;14:14&#160;UTC | true | [view](CERTS/8ca062af72eeb0b2eec4aaa1e6f3eb5ee3f01eccbd2895f282bbf4b70015a5d5/README.md) |
| 14&#160;Mar&#160;23&#160;21:39&#160;UTC | SHAKEN Cherry Voice 506K | 13&#160;Mar&#160;24&#160;21:39&#160;UTC | true | [view](CERTS/3756d1fa94ab1105ceb4bc6e3dcc371871b6edd45ac350df018a35f2f93f1967/README.md) |
| 19&#160;Mar&#160;23&#160;00:31&#160;UTC | SHAKEN Every1 Telecom 486K | 18&#160;Mar&#160;24&#160;00:31&#160;UTC | true | [view](CERTS/5cdfbb1a416083096dfef10c75a2b26a08d8fd5593d8ea9ceae0d70d878a97d1/README.md) |
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
| 02&#160;Oct&#160;23&#160;15:22&#160;UTC | SHAKEN TeleVoIPs 138K | 01&#160;Oct&#160;24&#160;15:22&#160;UTC | true | [view](CERTS/8df110d0b382bc6c2ffd95e4adde6fa10f679b4fdb578f0dc8163b7e06be7635/README.md) |
| 06&#160;Oct&#160;23&#160;13:25&#160;UTC | SHAKEN Kloud 7 LLC 214K | 05&#160;Oct&#160;24&#160;13:25&#160;UTC | true | [view](CERTS/7105ca55bae14fd5fc5952beff1664e0b13f3b560dc07c777a8293f5d06ab752/README.md) |
| 12&#160;Oct&#160;23&#160;20:53&#160;UTC | SHAKEN Zito Media Voice 624G | 11&#160;Oct&#160;24&#160;20:53&#160;UTC | true | [view](CERTS/139f5ea60b5e891b041824e82c242eb6191006f1283e9bad6cd052bf821a0fc2/README.md) |
| 17&#160;Oct&#160;23&#160;22:55&#160;UTC | SHAKEN ConnectMeVoice 719J | 16&#160;Oct&#160;24&#160;22:55&#160;UTC | true | [view](CERTS/5a48a80de9440625305bf4a9af18bcf158e03286df27646af231d75e0cec315d/README.md) |
| 18&#160;Oct&#160;23&#160;22:36&#160;UTC | SHAKEN VOICE1 LLC 265K | 17&#160;Oct&#160;24&#160;22:36&#160;UTC | true | [view](CERTS/f40966bfe5a54cc2e71f5942314743411e143934deeeb8b87647eb1b5ac03aa2/README.md) |
| 19&#160;Oct&#160;23&#160;11:43&#160;UTC | SHAKEN Talk IT Pro 321K | 18&#160;Oct&#160;24&#160;11:43&#160;UTC | true | [view](CERTS/ac46b1e519197ae9ce860b54ca7b6a7150d4be7de7e581e7481003737cf28f68/README.md) |
| 23&#160;Oct&#160;23&#160;15:54&#160;UTC | SHAKEN Comtalk Telecom 705K | 22&#160;Oct&#160;24&#160;15:54&#160;UTC | true | [view](CERTS/05f3a08f601f71f6b02f390050da6acb922f75b5f9d9178b637e96de52847674/README.md) |
| 25&#160;Oct&#160;23&#160;20:09&#160;UTC | SHAKEN Ytel Inc. 703J | 24&#160;Oct&#160;24&#160;20:09&#160;UTC | true | [view](CERTS/1abf91916a7c83ff53e58fdaf8fab1e6f6c232bc2aa2b1903c76c37d7a6984a6/README.md) |
| 27&#160;Oct&#160;23&#160;09:33&#160;UTC | SHAKEN Primo Dialler LLC 249K | 26&#160;Oct&#160;24&#160;09:33&#160;UTC | true | [view](CERTS/2345fd9200b5754c1d5fb353ea414d4be8e9eda729202bd84c0c3f7c6a0d1ad6/README.md) |
| 27&#160;Oct&#160;23&#160;22:04&#160;UTC | SHAKEN Mitel Cloud Services, Inc. 670J | 26&#160;Oct&#160;24&#160;22:04&#160;UTC | true | [view](CERTS/a3fecbe272ac59a623119a5b64a5e3cc79f9f880c2b82584a79f6f40e7f1562a/README.md) |
| 27&#160;Oct&#160;23&#160;22:26&#160;UTC | SHAKEN Mitel Cloud Services, Inc. 670J | 26&#160;Oct&#160;24&#160;22:26&#160;UTC | true | [view](CERTS/6ce3263c3dd426c9f8f57b7fabf533b30c5cf3c3582f7c5cd6f970b0ed0c01b1/README.md) |
| 30&#160;Oct&#160;23&#160;14:15&#160;UTC | SHAKEN Vinculum Communications, Inc 787J | 29&#160;Oct&#160;24&#160;14:15&#160;UTC | true | [view](CERTS/2d3cf73af0f77cf310d869e984084841f7b2ebb149a0d3d09694dfa5838b074b/README.md) |
| 31&#160;Oct&#160;23&#160;16:23&#160;UTC | SHAKEN ConvergeTel LLC 388K | 30&#160;Oct&#160;24&#160;16:23&#160;UTC | true | [view](CERTS/886577ab6ce0cf85d3ae954a4bf47c8ebf47246340a1775e0c7bca8b6e267244/README.md) |
| 02&#160;Nov&#160;23&#160;13:49&#160;UTC | SHAKEN Starlinq PBX Inc. 267K | 01&#160;Nov&#160;24&#160;13:49&#160;UTC | true | [view](CERTS/772495eb19f14de884567979521029ea8d0f85712aae22b5dea6ffadb6fd6aaf/README.md) |
| 03&#160;Nov&#160;23&#160;12:06&#160;UTC | SHAKEN Celtic Communications, LLC 836J | 02&#160;Nov&#160;24&#160;12:06&#160;UTC | true | [view](CERTS/0dae343c2561ca3b8f6c28949ee3800d2ca4b5f3fbffcc147407e74272912faf/README.md) |
| 07&#160;Nov&#160;23&#160;15:55&#160;UTC | SHAKEN Carrier One Inc. 705J | 06&#160;Nov&#160;24&#160;15:55&#160;UTC | true | [view](CERTS/89208e1662d350e58c5f5da80982aa62487d1cbf640a574d6d2ce7ebbf17a2fb/README.md) |
| 10&#160;Nov&#160;23&#160;02:10&#160;UTC | SHAKEN Star2Star Communications, LLC 590J | 09&#160;Nov&#160;24&#160;02:10&#160;UTC | true | [view](CERTS/d8197ed17be0c4415c77edbda2946e9c03f29dd67299927e8ecc654794c83951/README.md) |
| 13&#160;Nov&#160;23&#160;16:45&#160;UTC | SHAKEN ComData Solutions 451K | 12&#160;Nov&#160;24&#160;16:45&#160;UTC | true | [view](CERTS/dfe0f154e9a583eef556d6547d9d9d68ad605f05e867c9b8e7707811e2a9ed83/README.md) |
| 14&#160;Nov&#160;23&#160;13:29&#160;UTC | SHAKEN Asia Pacific Network 988J | 13&#160;Nov&#160;24&#160;13:29&#160;UTC | true | [view](CERTS/12b9440a58d22c821e164016a60deccd67ed5036e2244bd1694f39ce326e5811/README.md) |
| 14&#160;Nov&#160;23&#160;13:34&#160;UTC | SHAKEN Current Calls, LLC 746J | 13&#160;Nov&#160;24&#160;13:34&#160;UTC | true | [view](CERTS/a1db34e8f1daf7c5246c10508cf9dd97a3ab166f0ad7304ba124cdd53f199143/README.md) |
| 14&#160;Nov&#160;23&#160;13:36&#160;UTC | SHAKEN Godaddy 463K | 13&#160;Nov&#160;24&#160;13:36&#160;UTC | true | [view](CERTS/cc4ebe2a97d75e8560b603d277ce9e5bfb7a1e8dd5da0628e64e90576d270750/README.md) |
| 14&#160;Nov&#160;23&#160;13:38&#160;UTC | SHAKEN ONE OWL TELECOM INC 412K | 13&#160;Nov&#160;24&#160;13:38&#160;UTC | true | [view](CERTS/a8a6cbe884baa1d42768251df5233f7ffd0b33886539cf00df0d1b16cc48c3f3/README.md) |
| 14&#160;Nov&#160;23&#160;13:38&#160;UTC | SHAKEN OneStream Networks, LLC 630J | 13&#160;Nov&#160;24&#160;13:38&#160;UTC | true | [view](CERTS/d0a3700b7239c2f2aac6b3bbbac04b98db8bd5a8f725c2ebd1f8aed90126a4e1/README.md) |
| 14&#160;Nov&#160;23&#160;13:43&#160;UTC | SHAKEN Ringfree Communications Inc 317K | 13&#160;Nov&#160;24&#160;13:43&#160;UTC | true | [view](CERTS/066d5e0f2e11207f9a8aab5ae03e7bb3877292af0f325a690ce9cff6a9db42f9/README.md) |
| 14&#160;Nov&#160;23&#160;13:45&#160;UTC | SHAKEN Vumber LLC 225K | 13&#160;Nov&#160;24&#160;13:45&#160;UTC | true | [view](CERTS/4c72b9f5cbad420cbbb77aa0ca320baae202029d469d1953cb7e3b84ea42110d/README.md) |
| 15&#160;Nov&#160;23&#160;18:42&#160;UTC | SHAKEN Xchange Telecom LLC 325B | 14&#160;Nov&#160;24&#160;18:42&#160;UTC | true | [view](CERTS/25a16a8e688b969af60b0be46ce6b3d6b2525148559e5e60734bcf64f2afb635/README.md) |
| 27&#160;Nov&#160;23&#160;16:20&#160;UTC | SHAKEN Rayfield Communications, Inc. 006K | 26&#160;Nov&#160;24&#160;16:20&#160;UTC | true | [view](CERTS/b2ed7ac129b77320b8746043e76d523a27cfa3518f64cdffde32461737dae402/README.md) |
| 27&#160;Nov&#160;23&#160;17:58&#160;UTC | SHAKEN Drop Inc 258K | 26&#160;Nov&#160;24&#160;17:58&#160;UTC | true | [view](CERTS/3e07b9b82a23c5d6100f52a7d4532a6c3577b1a74a4e51850148bd1a44798a0b/README.md) |
| 27&#160;Nov&#160;23&#160;18:04&#160;UTC | SHAKEN MagicJack 324E | 26&#160;Nov&#160;24&#160;18:04&#160;UTC | true | [view](CERTS/09c017d4d76a9d8888d3aa630017a70534b727afdb22a0ad8031c4bf27e92454/README.md) |
| 27&#160;Nov&#160;23&#160;22:11&#160;UTC | SHAKEN Systemverse, LLC. 194K | 26&#160;Nov&#160;24&#160;22:11&#160;UTC | true | [view](CERTS/3c28ba328474fb61e883ba300c87a3881dc0ada933b30dc13f01209444ff5e55/README.md) |
| 01&#160;Dec&#160;23&#160;17:32&#160;UTC | SHAKEN ALD Telecom 780J | 29&#160;Feb&#160;24&#160;17:32&#160;UTC | true | [view](CERTS/798a1045417882dbc904706a6cbbff32020feb687897b659196803382124d9ae/README.md) |
| 06&#160;Dec&#160;23&#160;03:21&#160;UTC | SHAKEN Bulk Solutions, LLC 644J | 05&#160;Dec&#160;24&#160;03:21&#160;UTC | true | [view](CERTS/ab634090acf3c02b7dcd70995171b6522e39c5a72d73fe26e32e99b1acd716d7/README.md) |
| 08&#160;Dec&#160;23&#160;02:27&#160;UTC | SHAKEN Every1 Telecom 486K | 16&#160;Apr&#160;24&#160;02:27&#160;UTC | true | [view](CERTS/08a1c73bb18eaee3f90f85a99562a04a0b14827f7f8c663ae2fdbda74e0b5e29/README.md) |
| 18&#160;Dec&#160;23&#160;13:02&#160;UTC | SHAKEN Voice SY LLC 521K | 17&#160;Dec&#160;24&#160;13:02&#160;UTC | true | [view](CERTS/17cf3db291a3964a6e7d259ee034a5bf1ff6b0acfc4f2fee863c3984f67d37c2/README.md) |
| 19&#160;Dec&#160;23&#160;16:51&#160;UTC | SHAKEN Global Telecom Exchange LLC 270K | 17&#160;Apr&#160;24&#160;16:51&#160;UTC | true | [view](CERTS/6213824afa3022d4c424510204d4163db66f9f5c7644a7ab148928ba5616799a/README.md) |
| 27&#160;Dec&#160;23&#160;21:00&#160;UTC | SHAKEN Broadband Dynamics LLC 583j | 26&#160;Dec&#160;24&#160;21:00&#160;UTC | true | [view](CERTS/02c4ce9cf4dc82e32824fad2a5838b26279caa4b63db74fd1c40dabca7b9251b/README.md) |
| 28&#160;Dec&#160;23&#160;21:49&#160;UTC | SHAKEN California Telecom 319K | 27&#160;Dec&#160;24&#160;21:49&#160;UTC | true | [view](CERTS/28bc7840d450b8a6a1136187440b94d10a75465c409275151cd81f34ada8c490/README.md) |
| 29&#160;Dec&#160;23&#160;13:22&#160;UTC | SHAKEN Dalton Utilities 3139 | 26&#160;Jun&#160;24&#160;13:22&#160;UTC | true | [view](CERTS/12f9630b56a30698915ca2ebb4d7b8f650dd5ebcda09ec4c0d3aa0c360b5e690/README.md) |
| 29&#160;Dec&#160;23&#160;18:24&#160;UTC | SHAKEN IP Link Telecom Inc. 902J | 28&#160;Dec&#160;24&#160;18:24&#160;UTC | true | [view](CERTS/1f14385f4abd71de4dd72a9c1fd6699618abc7032794445bc8656c644a4a621d/README.md) |
| 10&#160;Jan&#160;24&#160;20:26&#160;UTC | SHAKEN Bulk Solutions, LLC 644J | 08&#160;Jul&#160;24&#160;20:26&#160;UTC | true | [view](CERTS/bd2ecfd7e5901a59f5e166982bdd69071b35aa4464ac54040aaed7f39680a77e/README.md) |
| 11&#160;Jan&#160;24&#160;19:30&#160;UTC | SHAKEN Arbeit 816J | 10&#160;Jan&#160;25&#160;19:30&#160;UTC | true | [view](CERTS/655f8cf330bddb83c0b081bc6438e4ba274013772961c17e2417630e1f8eb7f7/README.md) |
| 14&#160;Jan&#160;24&#160;19:29&#160;UTC | SHAKEN IDT America, Corp 363A | 13&#160;Feb&#160;24&#160;19:29&#160;UTC | true | [view](CERTS/6839680024f140b962f45b08e59e32931dc35fc9c274d8e8f576a9ed9e682689/README.md) |
| 14&#160;Jan&#160;24&#160;20:00&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 13&#160;Feb&#160;24&#160;20:00&#160;UTC | true | [view](CERTS/65498c417edb85b6f155d8508bdbb77002e0aa6a529f59d261dc47ea0988bf78/README.md) |
| 14&#160;Jan&#160;24&#160;22:12&#160;UTC | SHAKEN InteractiveTel, LLC 920J | 13&#160;Feb&#160;24&#160;22:12&#160;UTC | true | [view](CERTS/1276cd266f88ffa9bc6b9c2b070cc5805e087adcea005ced048caf7fc1fcb21e/README.md) |
| 15&#160;Jan&#160;24&#160;00:00&#160;UTC | SHAKEN Vitelglobal communications 698K | 14&#160;Feb&#160;24&#160;00:00&#160;UTC | true | [view](CERTS/5c5b94d0dc7b608bf7adce7a74bd3964015900f27aa792f6c93eb08788b4f608/README.md) |
| 15&#160;Jan&#160;24&#160;07:00&#160;UTC | SHAKEN Convoso 758J | 19&#160;Feb&#160;24&#160;07:00&#160;UTC | true | [view](CERTS/a0f46ba80831f05d96b2c2372a56011f38303b42f9786f81f37d7454462141d9/README.md) |
| 15&#160;Jan&#160;24&#160;08:12&#160;UTC | SHAKEN Lightspeed Voice 557F | 14&#160;Feb&#160;24&#160;08:12&#160;UTC | true | [view](CERTS/8d9493f79e82be18478105fb7cbe7e34c1891da2773cb8c3bb362a7b592ee1a3/README.md) |
| 15&#160;Jan&#160;24&#160;13:16&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 14&#160;Feb&#160;24&#160;13:16&#160;UTC | true | [view](CERTS/5e73fe43f15549522c1c5dd54873e2677d7fa765e496eee370492f10fc7eed29/README.md) |
| 15&#160;Jan&#160;24&#160;19:24&#160;UTC | SHAKEN IDT America, Corp 363A | 14&#160;Feb&#160;24&#160;19:24&#160;UTC | true | [view](CERTS/dec9bbf3bbbc11c6cbcdecd578e5e524be16c48b1685ee09d59f10a41535af1a/README.md) |
| 15&#160;Jan&#160;24&#160;19:55&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 14&#160;Feb&#160;24&#160;19:55&#160;UTC | true | [view](CERTS/7443c5ece8dfcf5f6142393ffdf81662ff96347d1dc5bccf51b5702d1671497a/README.md) |
| 15&#160;Jan&#160;24&#160;23:59&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 14&#160;Feb&#160;24&#160;23:59&#160;UTC | true | [view](CERTS/dcad41c194243496700c536647ae20fe4ed7149c438a249db88335d2dbf5deb4/README.md) |
| 16&#160;Jan&#160;24&#160;19:19&#160;UTC | SHAKEN IDT America, Corp 363A | 15&#160;Feb&#160;24&#160;19:19&#160;UTC | true | [view](CERTS/e4d217396f28380e0550dfa4c9b88cc08b01d2ae15c233bb5fd339be0645b5cf/README.md) |
| 16&#160;Jan&#160;24&#160;19:50&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 15&#160;Feb&#160;24&#160;19:50&#160;UTC | true | [view](CERTS/e7135b1d6e5709c523510efd03321c0d2720065886e29da8912ec0665ac24d2e/README.md) |
| 16&#160;Jan&#160;24&#160;23:54&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 15&#160;Feb&#160;24&#160;23:54&#160;UTC | true | [view](CERTS/c0f0f8ebc565aedcb53981cdf0c389b38ccbbc2c9161919ec1101d2b47f28b01/README.md) |
| 17&#160;Jan&#160;24&#160;08:02&#160;UTC | SHAKEN Lightspeed Voice 557F | 16&#160;Feb&#160;24&#160;08:02&#160;UTC | true | [view](CERTS/5d1b80a318d4b77e9b5fecd9eeddd81d7f01eac4a213e53e27ae0d0b024f16d1/README.md) |
| 17&#160;Jan&#160;24&#160;13:07&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 16&#160;Feb&#160;24&#160;13:07&#160;UTC | true | [view](CERTS/841460c2978d2356e9eb886ae7b93972ea445b24fcdc666f4ee9dbc1571acd4a/README.md) |
| 17&#160;Jan&#160;24&#160;19:14&#160;UTC | SHAKEN IDT America, Corp 363A | 16&#160;Feb&#160;24&#160;19:14&#160;UTC | true | [view](CERTS/be766d89d3a6de34feb1545c90add4d2ebb3a8f6cce601cc8484fed679ed696b/README.md) |
| 17&#160;Jan&#160;24&#160;19:45&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 16&#160;Feb&#160;24&#160;19:45&#160;UTC | true | [view](CERTS/ac7667f04ff7f7b6978faf920ac11f1dd76c51c0d67fd4a260c596500c406a66/README.md) |
| 17&#160;Jan&#160;24&#160;21:57&#160;UTC | SHAKEN InteractiveTel, LLC 920J | 16&#160;Feb&#160;24&#160;21:57&#160;UTC | true | [view](CERTS/4fa91ca189d005cf17e1f1048c41bde385425da65a4347e0d1f36f6cbd922466/README.md) |
| 18&#160;Jan&#160;24&#160;14:11&#160;UTC | SHAKEN BareTelecom 864J | 17&#160;Feb&#160;24&#160;14:11&#160;UTC | true | [view](CERTS/317e2a2fccfeba68b2a4e591bb1feb5adbc38bc893102d2a318443df4c5b2a21/README.md) |
| 18&#160;Jan&#160;24&#160;15:13&#160;UTC | SHAKEN IDT America, Corp 363A | 17&#160;Feb&#160;24&#160;15:13&#160;UTC | true | [view](CERTS/83a9bec210af02bceb5a7764648823db32d07952584c4761ab0cfc791db015b7/README.md) |
| 18&#160;Jan&#160;24&#160;18:50&#160;UTC | SHAKEN Socket Telecom LLC 554a | 17&#160;Feb&#160;24&#160;18:50&#160;UTC | true | [view](CERTS/ff720eb3e2229c368536377b52aa8dfecb403918daa1a02972ba347bfe105269/README.md) |
| 18&#160;Jan&#160;24&#160;19:40&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 17&#160;Feb&#160;24&#160;19:40&#160;UTC | true | [view](CERTS/d41815b70b299372a6df0d150f00b63a097b91e61d0d55f5fffabfe4602f579b/README.md) |
| 18&#160;Jan&#160;24&#160;21:41&#160;UTC | SHAKEN IDT America, Corp 363A | 17&#160;Feb&#160;24&#160;21:41&#160;UTC | true | [view](CERTS/82384e0fffa162fa19d212b261f59b7860542069434aacf97c67b1c62342dea9/README.md) |
| 18&#160;Jan&#160;24&#160;22:40&#160;UTC | SHAKEN Telcentris Inc. dba Voxox 696J | 17&#160;Feb&#160;24&#160;22:40&#160;UTC | true | [view](CERTS/66bdecd7fff03628beaa1c8b0d945047ab8feb32ad079d88bb5abd2f348d5725/README.md) |
| 18&#160;Jan&#160;24&#160;22:48&#160;UTC | SHAKEN Telcentris Inc. dba Voxox 696J | 17&#160;Feb&#160;24&#160;22:48&#160;UTC | true | [view](CERTS/ad992530df144b1734dc2af4bf73fa4b26b1fa504965e5c96bad53cb3d89ab52/README.md) |
| 18&#160;Jan&#160;24&#160;23:44&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 17&#160;Feb&#160;24&#160;23:44&#160;UTC | true | [view](CERTS/e9e322c9977fc7a36e1551fb345eca2160d9d5cfef9114f483bf77f9b53214ba/README.md) |
| 19&#160;Jan&#160;24&#160;06:30&#160;UTC | SHAKEN  XCast Labs 689J | 19&#160;Feb&#160;24&#160;06:30&#160;UTC | true | [view](CERTS/a3c3f08b59754feb33f9f6721fdd82b41897973c7e23efd35bd7641ea5c171af/README.md) |
| 19&#160;Jan&#160;24&#160;14:06&#160;UTC | SHAKEN BareTelecom 864J | 18&#160;Feb&#160;24&#160;14:06&#160;UTC | true | [view](CERTS/e5ca123a39613c506a76d678afd2b7a25fa0e1aff31a1cdb6166b772306bba25/README.md) |
| 19&#160;Jan&#160;24&#160;15:09&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 18&#160;Feb&#160;24&#160;15:09&#160;UTC | true | [view](CERTS/4d55cf1e983ef1eebc7c9dc1f761321f04ff290e581397459c63c285081e8dd5/README.md) |
| 19&#160;Jan&#160;24&#160;19:35&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 18&#160;Feb&#160;24&#160;19:35&#160;UTC | true | [view](CERTS/ac8d9bf1fe753976bd4a160f570b4a9ebf756d7a7c318a6783afbc14f29c4a5e/README.md) |
| 19&#160;Jan&#160;24&#160;23:26&#160;UTC | SHAKEN IDT America, Corp 363A | 18&#160;Feb&#160;24&#160;23:26&#160;UTC | true | [view](CERTS/67b7023b4b655cdc415ee150aa39237586a4677412257570ef562d300b72a2c3/README.md) |
| 20&#160;Jan&#160;24&#160;23:21&#160;UTC | SHAKEN IDT America, Corp 363A | 19&#160;Feb&#160;24&#160;23:21&#160;UTC | true | [view](CERTS/b17da06865b5a09ae26ff84a496ee2405ea777dd467f79bd6aaead4398717ce6/README.md) |
| 21&#160;Jan&#160;24&#160;07:42&#160;UTC | SHAKEN Lightspeed Voice 557F | 20&#160;Feb&#160;24&#160;07:42&#160;UTC | true | [view](CERTS/862b7e792faa3ba010933a88488cf08aaba97a38a808822d6cea52128ae4d84c/README.md) |
| 21&#160;Jan&#160;24&#160;14:59&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 20&#160;Feb&#160;24&#160;14:59&#160;UTC | true | [view](CERTS/17f3859ec71d1d18ddfcee0cc04358119b7557bade5b312564d04428c1f578ce/README.md) |
| 21&#160;Jan&#160;24&#160;19:25&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 20&#160;Feb&#160;24&#160;19:25&#160;UTC | true | [view](CERTS/9fdc729f4fbeb7418426e0f865c8b033d0eeac0c3b284a1f70a32c6b603e37e0/README.md) |
| 21&#160;Jan&#160;24&#160;23:16&#160;UTC | SHAKEN IDT America, Corp 363A | 20&#160;Feb&#160;24&#160;23:16&#160;UTC | true | [view](CERTS/25f56418c5d1ed8f87c91955ead7ce04297ecce35b536af54fd808a05c9b06a6/README.md) |
| 21&#160;Jan&#160;24&#160;23:29&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 20&#160;Feb&#160;24&#160;23:29&#160;UTC | true | [view](CERTS/6603399f3a62a5466cf447232380fd2788e49bb79e50a120e7a9a81f6fba38c9/README.md) |
| 22&#160;Jan&#160;24&#160;00:14&#160;UTC | SHAKEN BareTelecom 864J | 21&#160;Feb&#160;24&#160;00:14&#160;UTC | true | [view](CERTS/0e693c132317ac2d62f6a454990ce16e0c1dc084aa2b2035c9df7da34500229e/README.md) |
| 22&#160;Jan&#160;24&#160;14:33&#160;UTC | SHAKEN Global Net Holdings Inc 306K | 21&#160;Jan&#160;25&#160;14:33&#160;UTC | true | [view](CERTS/df27855ac3adb23e2757af24806db9d4a98f70446fff9606d9a8d4196c66f64e/README.md) |
| 22&#160;Jan&#160;24&#160;14:37&#160;UTC | SHAKEN Global Net Holdings Inc 306K | 21&#160;Jan&#160;25&#160;14:37&#160;UTC | true | [view](CERTS/5b37238808c628f29580dc7b328a570ffa010e1bf0a2a02a6b15e0341dfc7d93/README.md) |
| 22&#160;Jan&#160;24&#160;14:54&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 21&#160;Feb&#160;24&#160;14:54&#160;UTC | true | [view](CERTS/6daeff963cca6b32641d6594111b6712df239d6f1590fa0ae30514ae12cb13d6/README.md) |
| 22&#160;Jan&#160;24&#160;16:19&#160;UTC | SHAKEN Identidad Advertising Development LLC 617K | 21&#160;Feb&#160;24&#160;16:19&#160;UTC | true | [view](CERTS/8b19b553d0c880fc1c053a113717174cb17a5017456c895c486e24071bc13179/README.md) |
| 22&#160;Jan&#160;24&#160;16:45&#160;UTC | SHAKEN BareTelecom 864J | 21&#160;Feb&#160;24&#160;16:45&#160;UTC | true | [view](CERTS/9a55d36a4a7ce83bb68e1e4ae82ece70f3cebfb894393e56383dbbb46780f461/README.md) |
| 22&#160;Jan&#160;24&#160;19:09&#160;UTC | SHAKEN BareTelecom 864J | 21&#160;Feb&#160;24&#160;19:09&#160;UTC | true | [view](CERTS/7f00ba05ace860fd4b8bfe5d0ca881a29bcbe536e530c335a5c6ae2816279de4/README.md) |
| 22&#160;Jan&#160;24&#160;19:20&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 21&#160;Feb&#160;24&#160;19:20&#160;UTC | true | [view](CERTS/cd3fa37cfa280576d4e16062d0ad98cc3b6faa09326a22cce6c7f03e4dd4010e/README.md) |
| 22&#160;Jan&#160;24&#160;23:11&#160;UTC | SHAKEN IDT America, Corp 363A | 21&#160;Feb&#160;24&#160;23:11&#160;UTC | true | [view](CERTS/bbcb1ca0e3f2347118cd29d01e3862ed92d4c69c8a3815d2263dbed31c1c204c/README.md) |
| 22&#160;Jan&#160;24&#160;23:24&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 21&#160;Feb&#160;24&#160;23:24&#160;UTC | true | [view](CERTS/1ef2b0d493fa01e812010b6f8386f9613c86d6e1ce3ee51cfb5e3624cad860de/README.md) |
| 23&#160;Jan&#160;24&#160;14:49&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 22&#160;Feb&#160;24&#160;14:49&#160;UTC | true | [view](CERTS/b8f500454d5a699ef32fc8237d03d1b374bd16f159df64c6b3ef25506e3b50e4/README.md) |
| 23&#160;Jan&#160;24&#160;18:14&#160;UTC | SHAKEN 599J Technology Innovation Lab | 22&#160;Jan&#160;25&#160;18:14&#160;UTC | true | [view](CERTS/1f4f7ee438c7fb0441e8b67a27ee334e35a339112bfb63acb3e8cece2322c1ac/README.md) |
| 23&#160;Jan&#160;24&#160;19:04&#160;UTC | SHAKEN BareTelecom 864J | 22&#160;Feb&#160;24&#160;19:04&#160;UTC | true | [view](CERTS/9265476fd3e23aa51443ffac6380f1f6e90272e2f80c780415b7083d3335b40c/README.md) |
| 23&#160;Jan&#160;24&#160;19:15&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 22&#160;Feb&#160;24&#160;19:15&#160;UTC | true | [view](CERTS/bd74888a80ec972926d38cfa6ab7bdc3f29fb0e118a78864fa1c9a39e9862892/README.md) |
| 23&#160;Jan&#160;24&#160;23:06&#160;UTC | SHAKEN IDT America, Corp 363A | 22&#160;Feb&#160;24&#160;23:06&#160;UTC | true | [view](CERTS/64a0c66fbbc9e27f065bec301f6cb8cb5568f5241287f7396bb969d1ca3e03ce/README.md) |
| 23&#160;Jan&#160;24&#160;23:18&#160;UTC | SHAKEN IPSBS Managed Services LLC 828J | 22&#160;Feb&#160;24&#160;23:18&#160;UTC | true | [view](CERTS/513113885c83f9b9b8f22a534e220a679b5a96e772b5747c01d21dfe88f8afe5/README.md) |
| 23&#160;Jan&#160;24&#160;23:19&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 22&#160;Feb&#160;24&#160;23:19&#160;UTC | true | [view](CERTS/5a1cd4731c132cf8ec6a189c9143be22a179225cd859c277d2acde3956e96eb0/README.md) |
| 24&#160;Jan&#160;24&#160;14:44&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 23&#160;Feb&#160;24&#160;14:44&#160;UTC | true | [view](CERTS/95d2c639d34c4d9b13ac86fba525c95f6e301582d81284dff0d07df98e2a1982/README.md) |
| 24&#160;Jan&#160;24&#160;18:20&#160;UTC | SHAKEN Socket Telecom LLC 554a | 23&#160;Feb&#160;24&#160;18:20&#160;UTC | true | [view](CERTS/8a86e05d85a5bb6eca5124d901211640b23478e3a3b2044552334a207997e266/README.md) |
| 24&#160;Jan&#160;24&#160;18:59&#160;UTC | SHAKEN BareTelecom 864J | 23&#160;Feb&#160;24&#160;18:59&#160;UTC | true | [view](CERTS/03cd6ecec1292cc276cbd1da52a8fe11c4232f44ccf500020bcec72b8cbaf5f4/README.md) |
| 24&#160;Jan&#160;24&#160;19:10&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 23&#160;Feb&#160;24&#160;19:10&#160;UTC | true | [view](CERTS/ee36a815ff8d93348381ba2501d867f2ceeed26208c9c9d486855bd7bddd433c/README.md) |
| 24&#160;Jan&#160;24&#160;23:01&#160;UTC | SHAKEN IDT America, Corp 363A | 23&#160;Feb&#160;24&#160;23:01&#160;UTC | true | [view](CERTS/40f61502bdd0efcd1578871e26b56cf3a802ad60f5168417cefa50da13859438/README.md) |
| 24&#160;Jan&#160;24&#160;23:14&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 23&#160;Feb&#160;24&#160;23:14&#160;UTC | true | [view](CERTS/ed8271f4f91c80a7794b9e9f265c0e8c2d4c105a527a0c73c8615aacef0525b6/README.md) |
| 25&#160;Jan&#160;24&#160;00:00&#160;UTC | SHAKEN 732K Serius Network | 24&#160;Jan&#160;25&#160;00:00&#160;UTC | true | [view](CERTS/942253ee2461223f4489b0996e387e0a8b43f365cdbc33f791f3542c153c7dde/README.md) |
| 25&#160;Jan&#160;24&#160;14:39&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 24&#160;Feb&#160;24&#160;14:39&#160;UTC | true | [view](CERTS/ca0b3589482cc98d840a9ef10b6d41f884f1cbdaef984834697ff20dfee4c1ba/README.md) |
| 25&#160;Jan&#160;24&#160;18:15&#160;UTC | SHAKEN Socket Telecom LLC 554a | 24&#160;Feb&#160;24&#160;18:15&#160;UTC | true | [view](CERTS/63b90a5b65a4abf7fcc83b29dd80fcdc46613a65583e57d75c93d655d7a85ecb/README.md) |
| 25&#160;Jan&#160;24&#160;18:54&#160;UTC | SHAKEN BareTelecom 864J | 24&#160;Feb&#160;24&#160;18:54&#160;UTC | true | [view](CERTS/a63e86d7f6aa5f5cad48dd6b6974a760d96b3639fa9f6772d9109e19251e0ef8/README.md) |
| 25&#160;Jan&#160;24&#160;19:05&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 24&#160;Feb&#160;24&#160;19:05&#160;UTC | true | [view](CERTS/47efe3a992f6647a933dca7daac63b197bcb33aa85eaa01330b3de83b58bef53/README.md) |
| 25&#160;Jan&#160;24&#160;21:17&#160;UTC | SHAKEN InteractiveTel, LLC 920J | 24&#160;Feb&#160;24&#160;21:17&#160;UTC | true | [view](CERTS/80b86bb730a0476deef35769742b2a9e45aa36986b7cb6a6b0f27ad4a225356b/README.md) |
| 25&#160;Jan&#160;24&#160;22:56&#160;UTC | SHAKEN IDT America, Corp 363A | 24&#160;Feb&#160;24&#160;22:56&#160;UTC | true | [view](CERTS/142aa71bf1635d1fc942ed8095506e6d05da88a733bebe3ff1486c6b48657774/README.md) |
| 25&#160;Jan&#160;24&#160;23:09&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 24&#160;Feb&#160;24&#160;23:09&#160;UTC | true | [view](CERTS/2b9102b798cbc8eac1dc7dd5f197c5c68019c4eb2c6b0bfeef31b0af6d3deb6d/README.md) |
| 26&#160;Jan&#160;24&#160;18:49&#160;UTC | SHAKEN BareTelecom 864J | 25&#160;Feb&#160;24&#160;18:49&#160;UTC | true | [view](CERTS/301a2e26b2d1b8cb94cc1c49745c7f47268953f0669ac5d0e7e45fc4fcd6b660/README.md) |
| 26&#160;Jan&#160;24&#160;19:00&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 25&#160;Feb&#160;24&#160;19:00&#160;UTC | true | [view](CERTS/a1d62a400d9beb46ce4819a0954d5ae9d8ef0191db4345f95b7820fa730f29b3/README.md) |
| 26&#160;Jan&#160;24&#160;23:04&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 25&#160;Feb&#160;24&#160;23:04&#160;UTC | true | [view](CERTS/a6415891467c60d7a196c604c75815027feeeedca28ee65a1128e5c8ac07cba1/README.md) |
| 27&#160;Jan&#160;24&#160;22:46&#160;UTC | SHAKEN IDT America, Corp 363A | 26&#160;Feb&#160;24&#160;22:46&#160;UTC | true | [view](CERTS/2562abf32f61f1cdb284209df247ccb45c08c28af3aec883a52be1dee2d27a06/README.md) |
| 27&#160;Jan&#160;24&#160;22:59&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 26&#160;Feb&#160;24&#160;22:59&#160;UTC | true | [view](CERTS/648ee0f1e74bd3043296c939a70d9f619ed073a2e3fde6edd3977b7c6208a97f/README.md) |
| 28&#160;Jan&#160;24&#160;18:39&#160;UTC | SHAKEN BareTelecom 864J | 27&#160;Feb&#160;24&#160;18:39&#160;UTC | true | [view](CERTS/e33487e7fd2d6193e0409181209d6b53057cb278995b29c8f5329c6c897ee85a/README.md) |
| 28&#160;Jan&#160;24&#160;18:50&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 27&#160;Feb&#160;24&#160;18:50&#160;UTC | true | [view](CERTS/535c75fe3ea36e3a3f337708b5ebe77c88f5f551e17060eafdbfe91a16458678/README.md) |
| 28&#160;Jan&#160;24&#160;22:41&#160;UTC | SHAKEN IDT America, Corp 363A | 27&#160;Feb&#160;24&#160;22:41&#160;UTC | true | [view](CERTS/7795319a123f7418e910fe33a3c23cea07033fc090d3e620ea7afe667f3725c1/README.md) |
| 28&#160;Jan&#160;24&#160;22:53&#160;UTC | SHAKEN IPSBS Managed Services LLC 828J | 27&#160;Feb&#160;24&#160;22:53&#160;UTC | true | [view](CERTS/bd0858fbca6d369b6abb9ea8ae22cd062aad55c53506b490c58baac75e55a8fa/README.md) |
| 28&#160;Jan&#160;24&#160;22:54&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 27&#160;Feb&#160;24&#160;22:54&#160;UTC | true | [view](CERTS/4223d7f39e369516461e81d5982bc59ca5e2c7b01ab9000c586b335d18365b9d/README.md) |
| 29&#160;Jan&#160;24&#160;14:19&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 28&#160;Feb&#160;24&#160;14:19&#160;UTC | true | [view](CERTS/5feb64b5c3859d147638526859c8a0c3536e53cdebb768a125c1290835a9d7b8/README.md) |
| 29&#160;Jan&#160;24&#160;17:55&#160;UTC | SHAKEN Socket Telecom LLC 554a | 28&#160;Feb&#160;24&#160;17:55&#160;UTC | true | [view](CERTS/eee0c9816d3c01fb2084c84caae54054f9c62fd858934dc2a506aca1eef90d36/README.md) |
| 29&#160;Jan&#160;24&#160;18:34&#160;UTC | SHAKEN BareTelecom 864J | 28&#160;Feb&#160;24&#160;18:34&#160;UTC | true | [view](CERTS/3c0712482f416241b16367d85bd91a58e9a8f671dd3d175d5cfb8e0e49cdb496/README.md) |
| 29&#160;Jan&#160;24&#160;18:45&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 28&#160;Feb&#160;24&#160;18:45&#160;UTC | true | [view](CERTS/f2fd4a948b4320ff118288893505c1887e761beab799a9a2aeabd711482d57d8/README.md) |
| 29&#160;Jan&#160;24&#160;22:36&#160;UTC | SHAKEN IDT America, Corp 363A | 28&#160;Feb&#160;24&#160;22:36&#160;UTC | true | [view](CERTS/280e848bebddb9873524ff57027ae37cd2091065bb585e16e14dfe47c2fb6375/README.md) |
| 29&#160;Jan&#160;24&#160;22:49&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 28&#160;Feb&#160;24&#160;22:49&#160;UTC | true | [view](CERTS/38bf6d61551f680d3577c52d7acd795ac1ef82864c48d72ad83a1898a7e63bad/README.md) |
| 30&#160;Jan&#160;24&#160;14:14&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 29&#160;Feb&#160;24&#160;14:14&#160;UTC | true | [view](CERTS/4b705b44f4250bc05badb14b524c55b018ba52098f7c23892e2d272062d33782/README.md) |
| 30&#160;Jan&#160;24&#160;14:53&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 29&#160;Feb&#160;24&#160;14:53&#160;UTC | true | [view](CERTS/f8c55641708e33857672ed3c54bec498d4fbdb415ad148da6821bd03277a12a2/README.md) |
| 30&#160;Jan&#160;24&#160;17:50&#160;UTC | SHAKEN Socket Telecom LLC 554a | 29&#160;Feb&#160;24&#160;17:50&#160;UTC | true | [view](CERTS/163905c79810503cc0922a28d114c9103b1b1a522dda7cb9a61f4a567ee46ccf/README.md) |
| 30&#160;Jan&#160;24&#160;18:29&#160;UTC | SHAKEN BareTelecom 864J | 29&#160;Feb&#160;24&#160;18:29&#160;UTC | true | [view](CERTS/e5c9f23440e3293ad159dfee66f6fe21dd97c3a99a984cd8705e2286f5a42fda/README.md) |
| 30&#160;Jan&#160;24&#160;22:31&#160;UTC | SHAKEN IDT America, Corp 363A | 29&#160;Feb&#160;24&#160;22:31&#160;UTC | true | [view](CERTS/defc8733002453c8b25264cc0f2c608d700f1c5ab827a3c232f482342d01c149/README.md) |
| 31&#160;Jan&#160;24&#160;06:52&#160;UTC | SHAKEN Lightspeed Voice 557F | 01&#160;Mar&#160;24&#160;06:52&#160;UTC | true | [view](CERTS/477d2e0d46475df6afa0a82568ce75e91399fda6ff6bc53d4dbb04753f929798/README.md) |
| 31&#160;Jan&#160;24&#160;14:09&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 01&#160;Mar&#160;24&#160;14:09&#160;UTC | true | [view](CERTS/31ccb8367b0ea64c107054fd449ff2ebef1c3d23c8c16e0312b4b58a8f2ee87f/README.md) |
| 31&#160;Jan&#160;24&#160;14:48&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 01&#160;Mar&#160;24&#160;14:48&#160;UTC | true | [view](CERTS/e2a83dc8d90389cdd427640621bd37d555b4845da53f914abb7c8eed94aa81b7/README.md) |
| 31&#160;Jan&#160;24&#160;22:26&#160;UTC | SHAKEN IDT America, Corp 363A | 01&#160;Mar&#160;24&#160;22:26&#160;UTC | true | [view](CERTS/aaffa20d6472c41a5f7383e06379564c6b3c2c7dfe8e46b1d69589613d6d173d/README.md) |
| 31&#160;Jan&#160;24&#160;22:38&#160;UTC | SHAKEN IPSBS Managed Services LLC 828J | 01&#160;Mar&#160;24&#160;22:38&#160;UTC | true | [view](CERTS/257972d633380930805a01009f8d61e52b95e0d5f8cc879b3958e5f3a1fefc32/README.md) |
| 01&#160;Feb&#160;24&#160;01:13&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 02&#160;Mar&#160;24&#160;01:13&#160;UTC | true | [view](CERTS/85709f3f1badc11fbec8b056150b71de6c6f7fb0e048f4df134bc4905653cf56/README.md) |
| 01&#160;Feb&#160;24&#160;06:47&#160;UTC | SHAKEN Lightspeed Voice 557F | 02&#160;Mar&#160;24&#160;06:47&#160;UTC | true | [view](CERTS/18685e35b65d1ad362960ef56ff91709655b7484f4918b520c31da248c0bc754/README.md) |
| 01&#160;Feb&#160;24&#160;19:12&#160;UTC | SHAKEN IDT America, Corp 363A | 02&#160;Mar&#160;24&#160;19:12&#160;UTC | true | [view](CERTS/de9a0448da786b788a197fc8ab99c0f63d473b8dfc82733131d82a7c276426d3/README.md) |
| 01&#160;Feb&#160;24&#160;22:33&#160;UTC | SHAKEN IPSBS Managed Services LLC 828J | 02&#160;Mar&#160;24&#160;22:33&#160;UTC | true | [view](CERTS/a97621e393452ff32a535ef72752538f931668aa5cb5a96e5bb2460416eb4ee3/README.md) |
| 02&#160;Feb&#160;24&#160;01:08&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 03&#160;Mar&#160;24&#160;01:08&#160;UTC | true | [view](CERTS/fa40cf7ee68270e6ba87f7a4de73f2a290552c15adb0ffeba774301d30c0b7d3/README.md) |
| 02&#160;Feb&#160;24&#160;06:42&#160;UTC | SHAKEN Lightspeed Voice 557F | 03&#160;Mar&#160;24&#160;06:42&#160;UTC | true | [view](CERTS/4bb3351a22a790375348027ee4736b9d402d7a455153e44edb6c480dc3b2fb26/README.md) |
| 02&#160;Feb&#160;24&#160;16:37&#160;UTC | SHAKEN IDT America, Corp 363A | 03&#160;Mar&#160;24&#160;16:37&#160;UTC | true | [view](CERTS/2a5d4c474a990077fd489e9917f6e30c8741bf788b208bdd67bfbbaa1196ff03/README.md) |
| 02&#160;Feb&#160;24&#160;21:01&#160;UTC | SHAKEN IDT America, Corp 363A | 03&#160;Mar&#160;24&#160;21:01&#160;UTC | true | [view](CERTS/3497d3261fbd40c881ec663ced3ac3901c2dae3d52ee402e419cf44856f03ba4/README.md) |
| 03&#160;Feb&#160;24&#160;03:37&#160;UTC | SHAKEN IDT America, Corp 363A | 04&#160;Mar&#160;24&#160;03:37&#160;UTC | true | [view](CERTS/cff5f26341538a372a6b2c1d7130ccad2923049e4f143bebbce2ac6d0ff441ba/README.md) |
| 04&#160;Feb&#160;24&#160;20:27&#160;UTC | SHAKEN InteractiveTel, LLC 920J | 05&#160;Mar&#160;24&#160;20:27&#160;UTC | true | [view](CERTS/2acdce9ec2a2809baaa21e06a1f49c28b477f763b69483ba3b41026ed1ccfdc5/README.md) |
| 04&#160;Feb&#160;24&#160;22:18&#160;UTC | SHAKEN IPSBS Managed Services LLC 828J | 05&#160;Mar&#160;24&#160;22:18&#160;UTC | true | [view](CERTS/5ea5b6cf67dfd276c50abc220db5d6f9f95276decbb892d71f3c48442e65bff1/README.md) |
| 04&#160;Feb&#160;24&#160;22:19&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 05&#160;Mar&#160;24&#160;22:19&#160;UTC | true | [view](CERTS/0518b5180e30b8004f84bc39c97ec584d5473e68c7cc25a6eb2973e55e0ab639/README.md) |
| 05&#160;Feb&#160;24&#160;00:53&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 06&#160;Mar&#160;24&#160;00:53&#160;UTC | true | [view](CERTS/b44b75b77af217394edf3b4faf7c29d8224e1da127a91991eec1ae87b0ecb06b/README.md) |
| 05&#160;Feb&#160;24&#160;06:27&#160;UTC | SHAKEN Lightspeed Voice 557F | 06&#160;Mar&#160;24&#160;06:27&#160;UTC | true | [view](CERTS/ccac2371ce548e7dc592c226c000824ee29fe4c511ae39ba446fc81cd5f77e96/README.md) |
| 05&#160;Feb&#160;24&#160;13:42&#160;UTC | SHAKEN IDT America, Corp 363A | 06&#160;Mar&#160;24&#160;13:42&#160;UTC | true | [view](CERTS/ddf7506ebf804a6aa1ceff165e760ccf31d21b3c263b6c894010f4a5e82ba38e/README.md) |
| 05&#160;Feb&#160;24&#160;13:44&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 06&#160;Mar&#160;24&#160;13:44&#160;UTC | true | [view](CERTS/a678b5e36897011c7234ab584b67b535c16dc3abb7cade17dd489b0b468c78ea/README.md) |
| 05&#160;Feb&#160;24&#160;14:11&#160;UTC | SHAKEN Consolidated Communications 5113 | 06&#160;Mar&#160;24&#160;14:11&#160;UTC | true | [view](CERTS/681d85fe97a19b8601ce86d44d94cb84c01234058b1a88da8466c9286553abb9/README.md) |
| 05&#160;Feb&#160;24&#160;14:13&#160;UTC | SHAKEN Touchtone 683A | 06&#160;Mar&#160;24&#160;14:13&#160;UTC | true | [view](CERTS/4441d6806019a6e22df8ae04a2ac1208cd230701e63b1bb23f103f630021ae4c/README.md) |
| 05&#160;Feb&#160;24&#160;14:14&#160;UTC | SHAKEN Apeiron Systems 012J | 06&#160;Mar&#160;24&#160;14:14&#160;UTC | true | [view](CERTS/8397a14e0e020ac843ff741ca6d01eed3e87ac0e5de15c2c9bad6fcf73515be7/README.md) |
| 05&#160;Feb&#160;24&#160;14:15&#160;UTC | SHAKEN Fonative, Inc. 684J | 06&#160;Mar&#160;24&#160;14:15&#160;UTC | true | [view](CERTS/d2fded0ff42d2caa6b3f940b27f571aeca0a47eb38b46215966c494bc22d209a/README.md) |
| 05&#160;Feb&#160;24&#160;14:16&#160;UTC | SHAKEN IPitomy 652J | 06&#160;Mar&#160;24&#160;14:16&#160;UTC | true | [view](CERTS/0376631da08fe561c0b9f45b102b9f7dd3c8fa32c53ebf7a8da60b4f4552d083/README.md) |
| 05&#160;Feb&#160;24&#160;14:18&#160;UTC | SHAKEN Phone.com, Inc. 633J | 06&#160;Mar&#160;24&#160;14:18&#160;UTC | true | [view](CERTS/69d35058436e76d1480833ccaff4ba7461a8d119f07327e5569c1f3d4de77ac7/README.md) |
| 05&#160;Feb&#160;24&#160;14:19&#160;UTC | SHAKEN NETRIO LLC 020K | 06&#160;Mar&#160;24&#160;14:19&#160;UTC | true | [view](CERTS/a8d21f0d029a1e7863670b7343bbb5020e3d7821303eddd6d58e250002e7147d/README.md) |
| 05&#160;Feb&#160;24&#160;14:21&#160;UTC | SHAKEN Noble Systems Communications LLC 187J | 06&#160;Mar&#160;24&#160;14:21&#160;UTC | true | [view](CERTS/c1f736487f02e1e4b1932a8abce21d4ea712a3a69fcbecfd843e2e768adc8d62/README.md) |
| 05&#160;Feb&#160;24&#160;14:24&#160;UTC | SHAKEN Airespring 996H | 06&#160;Mar&#160;24&#160;14:24&#160;UTC | true | [view](CERTS/1eb165ad54b0a70c89bb0de8e6010d81a1b67ffd9601046fd4b5d8a2fa4bd425/README.md) |
| 05&#160;Feb&#160;24&#160;14:26&#160;UTC | SHAKEN Nobelbiz, Inc. 596J | 06&#160;Mar&#160;24&#160;14:26&#160;UTC | true | [view](CERTS/60d6534a31d147eef4db2a03a4c9c24bf820c9202b81cc734be19640c22e958e/README.md) |
| 05&#160;Feb&#160;24&#160;14:28&#160;UTC | SHAKEN Momentum Telecom 9157 | 06&#160;Mar&#160;24&#160;14:28&#160;UTC | true | [view](CERTS/06b24241ac971745de0625a44ae25d583ce6ef9ab8a0555c454098f8c99e021e/README.md) |
| 05&#160;Feb&#160;24&#160;14:32&#160;UTC | SHAKEN Matrix 7379 | 06&#160;Mar&#160;24&#160;14:32&#160;UTC | true | [view](CERTS/ee8066a211d2ad2f2ce8304cb568baf4272ffa96545d15ee91902082e065922c/README.md) |
| 05&#160;Feb&#160;24&#160;14:32&#160;UTC | SHAKEN Matrix 9451 | 06&#160;Mar&#160;24&#160;14:32&#160;UTC | true | [view](CERTS/85fdd7760ea550e598be498a111fe2e0d5544fbe9ae0d4d6f0d4ec517af73b73/README.md) |
| 05&#160;Feb&#160;24&#160;14:37&#160;UTC | SHAKEN Magna5, LLC 3849 | 06&#160;Mar&#160;24&#160;14:37&#160;UTC | true | [view](CERTS/a1ac149fe4f707f5c809ea1044c3953c4990279f9b4dfa1574a54a7dabc78eb5/README.md) |
| 05&#160;Feb&#160;24&#160;14:38&#160;UTC | SHAKEN Magna5, LLC 8249 | 06&#160;Mar&#160;24&#160;14:38&#160;UTC | true | [view](CERTS/abfa418f5e44f6469a2358e92b6c5bac6e4ca4f18b2afcfd0022f55515bb89eb/README.md) |
| 05&#160;Feb&#160;24&#160;18:03&#160;UTC | SHAKEN BareTelecom 864J | 06&#160;Mar&#160;24&#160;18:03&#160;UTC | true | [view](CERTS/6d11a9f62f298d564508be8a19757cbc542ac3cc82ee8ac3ae29e3b03bb56b0a/README.md) |
| 05&#160;Feb&#160;24&#160;20:56&#160;UTC | SHAKEN IDT America, Corp 363A | 06&#160;Mar&#160;24&#160;20:56&#160;UTC | true | [view](CERTS/e36273d0c059a12a8ab9a2c93fbd06ef06deb35427d918e3323a2845505dd7c4/README.md) |
| 05&#160;Feb&#160;24&#160;22:13&#160;UTC | SHAKEN IPSBS Managed Services LLC 828J | 06&#160;Mar&#160;24&#160;22:13&#160;UTC | true | [view](CERTS/6792b0b3f47968d1009324081feed1b3294004386e44a3e3b7bd81a2e9eddc61/README.md) |
| 05&#160;Feb&#160;24&#160;22:14&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 06&#160;Mar&#160;24&#160;22:14&#160;UTC | true | [view](CERTS/372dc873b603438305fb1452cb80819d5f1b29c9a80c1e8aaafad62322b2801a/README.md) |
| 06&#160;Feb&#160;24&#160;00:48&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 07&#160;Mar&#160;24&#160;00:48&#160;UTC | true | [view](CERTS/1a6bd409014fc56ed198fdee7b0e3c07a1e9a0914dc0c829a066171f5e3a622f/README.md) |
| 06&#160;Feb&#160;24&#160;17:58&#160;UTC | SHAKEN BareTelecom 864J | 07&#160;Mar&#160;24&#160;17:58&#160;UTC | true | [view](CERTS/7f2b3ce699a851ddab6be4e7cb3622b346cfaac4b6b3cf98259e7e0e4f6860dc/README.md) |
| 06&#160;Feb&#160;24&#160;20:51&#160;UTC | SHAKEN IDT America, Corp 363A | 07&#160;Mar&#160;24&#160;20:51&#160;UTC | true | [view](CERTS/f52b2b67232df268bc68d783a719dd14e1060457ca90e9064e198267288953a5/README.md) |

#### CA Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 21&#160;Aug&#160;20&#160;01:22&#160;UTC | SHAKEN Sansay Root CA US | 16&#160;Aug&#160;40&#160;01:22&#160;UTC | false | [view](CERTS/8356d251b255f4ac3fd108bb79be4c02dcea2d3b13d138892bdb3a70cbe6a343/README.md) |
| 02&#160;Sep&#160;22&#160;20:53&#160;UTC | SHAKEN Sansay Intermediate CA US WEST 1 | 31&#160;Aug&#160;29&#160;20:53&#160;UTC | false | [view](CERTS/4b1dfdba2b1e4bbffbf900a20f1f6f7befbef0008b963e4922a64469cb97d24b/README.md) |


Generated: 12 Feb 24 17:02 UTC