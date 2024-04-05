# STIR/SHAKEN CA Ecosystem Compliance

## Sansay

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 2048 certificates were included in the corpus being tested
- 18 certificates in the corpus were skipped because they are duplicates
- 1960 certificates in the corpus were skipped because they are expired
- 1 certificates in the corpus were skipped because they are not currently trusted
- 69 certificates being tested against the remaining rules
- 1.97 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 0.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 337 days is the average remaining validity for the certificates in the corpus
- 342 days is the average initial validity for the certificates in the corpus
- 7 certificates expire in the next 30 days
- 1.08 average number of unexpired certificates per OCN observed
- 64 unique OCNs observed in unexpired and valid certificate corpus

| Instances | Test | Source |
|-----------|------|--------|
| 4 | [e_atis_ext_crl_distribution_struct](ISSUES/e_atis_ext_crl_distribution_struct/README.md) | ATIS1000080 |
| 66 | [e_atis_subject_cn_spc](ISSUES/e_atis_subject_cn_spc/README.md) | ATIS1000080 |
| 1 | [e_atis_tn_auth_list_spc_format](ISSUES/e_atis_tn_auth_list_spc_format/README.md) | ATIS1000080 |
| 28 | [e_shaken_certificate_policies_id](ISSUES/e_shaken_certificate_policies_id/README.md) | US_SHAKEN_CP |
| 37 | [e_us_cp_subject_sn_shall](ISSUES/e_us_cp_subject_sn_shall/README.md) | US_SHAKEN_CP |

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
- 5277 days is the average remaining validity for the certificates in the corpus
- 4928 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days

No error, warning, or notice level issues were found

### Details

#### Leaf Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
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
| 22&#160;Jan&#160;24&#160;14:33&#160;UTC | SHAKEN Global Net Holdings Inc 306K | 21&#160;Jan&#160;25&#160;14:33&#160;UTC | true | [view](CERTS/df27855ac3adb23e2757af24806db9d4a98f70446fff9606d9a8d4196c66f64e/README.md) |
| 22&#160;Jan&#160;24&#160;14:37&#160;UTC | SHAKEN Global Net Holdings Inc 306K | 21&#160;Jan&#160;25&#160;14:37&#160;UTC | true | [view](CERTS/5b37238808c628f29580dc7b328a570ffa010e1bf0a2a02a6b15e0341dfc7d93/README.md) |
| 23&#160;Jan&#160;24&#160;18:14&#160;UTC | SHAKEN 599J Technology Innovation Lab | 22&#160;Jan&#160;25&#160;18:14&#160;UTC | true | [view](CERTS/1f4f7ee438c7fb0441e8b67a27ee334e35a339112bfb63acb3e8cece2322c1ac/README.md) |
| 25&#160;Jan&#160;24&#160;00:00&#160;UTC | SHAKEN 732K Serius Network | 24&#160;Jan&#160;25&#160;00:00&#160;UTC | true | [view](CERTS/942253ee2461223f4489b0996e387e0a8b43f365cdbc33f791f3542c153c7dde/README.md) |
| 29&#160;Feb&#160;24&#160;17:02&#160;UTC | SHAKEN 521K Voice SY LLC | 14&#160;Apr&#160;24&#160;17:02&#160;UTC | true | [view](CERTS/d4f2a58cc39d7cb3d3f083c5637c392dc2972a0653aca6218501df4337a25891/README.md) |
| 13&#160;Mar&#160;24&#160;06:00&#160;UTC | SHAKEN Convoso 758J | 17&#160;Apr&#160;24&#160;06:00&#160;UTC | true | [view](CERTS/92c738a06590f4700891af189907fb42b4194ea6abcffc4207393c58efe3ccbe/README.md) |

#### CA Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 21&#160;Aug&#160;20&#160;01:22&#160;UTC | SHAKEN Sansay Root CA US | 16&#160;Aug&#160;40&#160;01:22&#160;UTC | false | [view](CERTS/8356d251b255f4ac3fd108bb79be4c02dcea2d3b13d138892bdb3a70cbe6a343/README.md) |
| 02&#160;Sep&#160;22&#160;20:53&#160;UTC | SHAKEN Sansay Intermediate CA US WEST 1 | 31&#160;Aug&#160;29&#160;20:53&#160;UTC | false | [view](CERTS/4b1dfdba2b1e4bbffbf900a20f1f6f7befbef0008b963e4922a64469cb97d24b/README.md) |


Generated: 05 Apr 24 19:04 UTC