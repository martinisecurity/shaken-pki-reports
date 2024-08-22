# STIR/SHAKEN CA Ecosystem Compliance

## Sansay

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 2541 certificates were included in the corpus being tested
- 17 certificates in the corpus were skipped because they are duplicates
- 2303 certificates in the corpus were skipped because they are expired
- 1 certificates in the corpus were skipped because they are not currently trusted
- 220 certificates being tested against the remaining rules
- 1.90 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 0.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 157 days is the average remaining validity for the certificates in the corpus
- 158 days is the average initial validity for the certificates in the corpus
- 137 certificates expire in the next 30 days
- 2.06 average number of unexpired certificates per OCN observed
- 107 unique OCNs observed in unexpired and valid certificate corpus

| Instances | Test | Source |
|-----------|------|--------|
| 128 | [e_atis_ext_crl_distribution_struct](ISSUES/e_atis_ext_crl_distribution_struct/README.md) | ATIS1000080 |
| 194 | [e_atis_subject_cn_spc](ISSUES/e_atis_subject_cn_spc/README.md) | ATIS1000080 |
| 5 | [e_atis_tn_auth_list_spc_format](ISSUES/e_atis_tn_auth_list_spc_format/README.md) | ATIS1000080 |
| 63 | [e_shaken_certificate_policies_id](ISSUES/e_shaken_certificate_policies_id/README.md) | US_SHAKEN_CP |
| 29 | [e_us_cp_subject_sn_shall](ISSUES/e_us_cp_subject_sn_shall/README.md) | US_SHAKEN_CP |

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
- 5231 days is the average remaining validity for the certificates in the corpus
- 4928 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days

No error, warning, or notice level issues were found

### Details

#### Leaf Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 05&#160;Sep&#160;23&#160;15:54&#160;UTC | SHAKEN Go Voip Dialing LLC 704K | 04&#160;Sep&#160;24&#160;15:54&#160;UTC | true | [view](CERTS/645767c01de0509deb545953141dbf2e136665480917cc8c4ecc73a45608af68/README.md) |
| 08&#160;Sep&#160;23&#160;12:28&#160;UTC | SHAKEN Contactivity Corp. 711K | 07&#160;Sep&#160;24&#160;12:28&#160;UTC | true | [view](CERTS/673f3091743237809463817a98f68fb9fd95c3b112030e4fcbe84096f54d5038/README.md) |
| 15&#160;Sep&#160;23&#160;14:40&#160;UTC | SHAKEN Miracle Telecom 496K | 14&#160;Sep&#160;24&#160;14:40&#160;UTC | true | [view](CERTS/f7312e8a32e80db109fe012d1e00c252afc4eec07f6292fa6f714e926910d14e/README.md) |
| 26&#160;Sep&#160;23&#160;13:26&#160;UTC | SHAKEN Sangoma 777G | 25&#160;Sep&#160;24&#160;13:26&#160;UTC | true | [view](CERTS/2db8f92946ee707e88748612d6ea1b786c7137a8c56e53c80d9484911a2ed51d/README.md) |
| 27&#160;Sep&#160;23&#160;17:50&#160;UTC | SHAKEN Double A Solutions 644K | 26&#160;Sep&#160;24&#160;17:50&#160;UTC | true | [view](CERTS/fb18701e25fb45d49e3966dd64fb3a518f6f2a9150059b39ad61632ac9317922/README.md) |
| 02&#160;Oct&#160;23&#160;15:22&#160;UTC | SHAKEN TeleVoIPs 138K | 01&#160;Oct&#160;24&#160;15:22&#160;UTC | true | [view](CERTS/8df110d0b382bc6c2ffd95e4adde6fa10f679b4fdb578f0dc8163b7e06be7635/README.md) |
| 06&#160;Oct&#160;23&#160;13:25&#160;UTC | SHAKEN Kloud 7 LLC 214K | 05&#160;Oct&#160;24&#160;13:25&#160;UTC | true | [view](CERTS/7105ca55bae14fd5fc5952beff1664e0b13f3b560dc07c777a8293f5d06ab752/README.md) |
| 12&#160;Oct&#160;23&#160;20:53&#160;UTC | SHAKEN Zito Media Voice 624G | 11&#160;Oct&#160;24&#160;20:53&#160;UTC | true | [view](CERTS/139f5ea60b5e891b041824e82c242eb6191006f1283e9bad6cd052bf821a0fc2/README.md) |
| 17&#160;Oct&#160;23&#160;22:55&#160;UTC | SHAKEN ConnectMeVoice 719J | 16&#160;Oct&#160;24&#160;22:55&#160;UTC | true | [view](CERTS/5a48a80de9440625305bf4a9af18bcf158e03286df27646af231d75e0cec315d/README.md) |
| 18&#160;Oct&#160;23&#160;22:36&#160;UTC | SHAKEN VOICE1 LLC 265K | 17&#160;Oct&#160;24&#160;22:36&#160;UTC | true | [view](CERTS/f40966bfe5a54cc2e71f5942314743411e143934deeeb8b87647eb1b5ac03aa2/README.md) |
| 19&#160;Oct&#160;23&#160;11:43&#160;UTC | SHAKEN Talk IT Pro 321K | 18&#160;Oct&#160;24&#160;11:43&#160;UTC | true | [view](CERTS/ac46b1e519197ae9ce860b54ca7b6a7150d4be7de7e581e7481003737cf28f68/README.md) |
| 23&#160;Oct&#160;23&#160;14:32&#160;UTC | SHAKEN Uvoice USA, LLC 328K | 22&#160;Oct&#160;24&#160;14:32&#160;UTC | true | [view](CERTS/63faaac90f1c3f037b3c850ad0f6203068aa3d62235646df63f6b07c9aa94d05/README.md) |
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
| 18&#160;Dec&#160;23&#160;13:02&#160;UTC | SHAKEN Voice SY LLC 521K | 17&#160;Dec&#160;24&#160;13:02&#160;UTC | true | [view](CERTS/17cf3db291a3964a6e7d259ee034a5bf1ff6b0acfc4f2fee863c3984f67d37c2/README.md) |
| 27&#160;Dec&#160;23&#160;21:00&#160;UTC | SHAKEN Broadband Dynamics LLC 583j | 26&#160;Dec&#160;24&#160;21:00&#160;UTC | true | [view](CERTS/02c4ce9cf4dc82e32824fad2a5838b26279caa4b63db74fd1c40dabca7b9251b/README.md) |
| 28&#160;Dec&#160;23&#160;21:49&#160;UTC | SHAKEN California Telecom 319K | 27&#160;Dec&#160;24&#160;21:49&#160;UTC | true | [view](CERTS/28bc7840d450b8a6a1136187440b94d10a75465c409275151cd81f34ada8c490/README.md) |
| 29&#160;Dec&#160;23&#160;18:24&#160;UTC | SHAKEN IP Link Telecom Inc. 902J | 28&#160;Dec&#160;24&#160;18:24&#160;UTC | true | [view](CERTS/1f14385f4abd71de4dd72a9c1fd6699618abc7032794445bc8656c644a4a621d/README.md) |
| 29&#160;Dec&#160;23&#160;18:27&#160;UTC | SHAKEN IP Link Telecom Inc. 902J | 28&#160;Dec&#160;24&#160;18:27&#160;UTC | true | [view](CERTS/0927057b9fc6b184e8be7a88161eb4e88cb7ef138b26bc83d953db5d68f8fcca/README.md) |
| 11&#160;Jan&#160;24&#160;19:30&#160;UTC | SHAKEN Arbeit 816J | 10&#160;Jan&#160;25&#160;19:30&#160;UTC | true | [view](CERTS/655f8cf330bddb83c0b081bc6438e4ba274013772961c17e2417630e1f8eb7f7/README.md) |
| 22&#160;Jan&#160;24&#160;14:33&#160;UTC | SHAKEN Global Net Holdings Inc 306K | 21&#160;Jan&#160;25&#160;14:33&#160;UTC | true | [view](CERTS/df27855ac3adb23e2757af24806db9d4a98f70446fff9606d9a8d4196c66f64e/README.md) |
| 22&#160;Jan&#160;24&#160;14:37&#160;UTC | SHAKEN Global Net Holdings Inc 306K | 21&#160;Jan&#160;25&#160;14:37&#160;UTC | true | [view](CERTS/5b37238808c628f29580dc7b328a570ffa010e1bf0a2a02a6b15e0341dfc7d93/README.md) |
| 23&#160;Jan&#160;24&#160;18:14&#160;UTC | SHAKEN 599J Technology Innovation Lab | 22&#160;Jan&#160;25&#160;18:14&#160;UTC | true | [view](CERTS/1f4f7ee438c7fb0441e8b67a27ee334e35a339112bfb63acb3e8cece2322c1ac/README.md) |
| 25&#160;Jan&#160;24&#160;00:00&#160;UTC | SHAKEN 732K Serius Network | 24&#160;Jan&#160;25&#160;00:00&#160;UTC | true | [view](CERTS/942253ee2461223f4489b0996e387e0a8b43f365cdbc33f791f3542c153c7dde/README.md) |
| 29&#160;Jan&#160;24&#160;13:23&#160;UTC | SHAKEN ACS Technologies 488K | 28&#160;Jan&#160;25&#160;13:23&#160;UTC | true | [view](CERTS/0218a28abeeeb238045a8ddffa3593779dcf82a26f219585c3af62900a27ac51/README.md) |
| 13&#160;Feb&#160;24&#160;17:15&#160;UTC | SHAKEN 826K Infinity Communications LLC | 12&#160;Feb&#160;25&#160;17:15&#160;UTC | true | [view](CERTS/68b1119315e40d8b5e803ae4a01fa6de1b7fbef58d02dea80f372988da6987a0/README.md) |
| 21&#160;Feb&#160;24&#160;13:34&#160;UTC | SHAKEN Watchcomm  0590 | 20&#160;Feb&#160;25&#160;13:34&#160;UTC | true | [view](CERTS/f37c46fe22ca2348db842b6bef04535d9be4d6b413c2282fdca2969d292bcdcd/README.md) |
| 23&#160;Feb&#160;24&#160;16:46&#160;UTC | SHAKEN 781J iNet Communications | 22&#160;Feb&#160;25&#160;16:46&#160;UTC | true | [view](CERTS/e46cb2d0ae92726b11a121cee6fa0726a9dfabc72cfa74b5f7547c79838c503b/README.md) |
| 28&#160;Feb&#160;24&#160;17:13&#160;UTC | SHAKEN 839K Sigma Broadband | 27&#160;Feb&#160;25&#160;17:13&#160;UTC | true | [view](CERTS/d6f9b064110046c7c42acd971a566fb99b97c3740df54afb4b903a39179966b1/README.md) |
| 29&#160;Feb&#160;24&#160;23:52&#160;UTC | SHAKEN Ytel Inc. 703J | 28&#160;Feb&#160;25&#160;23:52&#160;UTC | true | [view](CERTS/118ca5f6b2653c823022d22375f900d88e845681ca442938142a44da593c4bb5/README.md) |
| 04&#160;Mar&#160;24&#160;13:49&#160;UTC | SHAKEN Terra Nova Telecom 382G | 04&#160;Mar&#160;25&#160;13:49&#160;UTC | true | [view](CERTS/c114f1214f94fc4ab21e74e0ad29a2dd5e0ef40080a75abc428178ecab55c3e4/README.md) |
| 18&#160;Mar&#160;24&#160;17:10&#160;UTC | SHAKEN Identidad Advertising Development LLC 617K | 18&#160;Mar&#160;25&#160;17:10&#160;UTC | true | [view](CERTS/2f9861d9339c6094cd1c5acf955e0472ffd5534c329f4b69c90caae6f70d61a3/README.md) |
| 20&#160;Mar&#160;24&#160;17:08&#160;UTC | SHAKEN Inventive Labs Corp 649J | 26&#160;Sep&#160;24&#160;17:08&#160;UTC | true | [view](CERTS/3973d56f7972cd4b773b8f2da99604225d584cba98fa42e37f89320e939e3ff6/README.md) |
| 26&#160;Mar&#160;24&#160;13:55&#160;UTC | SHAKEN 597F VoIP Innovations | 26&#160;Mar&#160;25&#160;13:55&#160;UTC | true | [view](CERTS/80f9a4656d0d9e150183b3c4c10c11ef69148ab725647d650dd81be489d897cf/README.md) |
| 26&#160;Mar&#160;24&#160;16:02&#160;UTC | SHAKEN 867K Alpine Valley Consulting LLC | 26&#160;Mar&#160;25&#160;16:02&#160;UTC | true | [view](CERTS/d8e42ff7151f25922813aae6e3e9d214f0157ac39fa1132017259254a6fe40d2/README.md) |
| 26&#160;Mar&#160;24&#160;18:12&#160;UTC | SHAKEN 493K ComTelus Inc.  | 26&#160;Mar&#160;25&#160;18:12&#160;UTC | true | [view](CERTS/af4cf171b2ca99abd62e41f7649ac1e22395fda011d173fdfdbd31891f743971/README.md) |
| 28&#160;Mar&#160;24&#160;10:51&#160;UTC | SHAKEN Global Net Holdings Inc 306K | 28&#160;Mar&#160;25&#160;10:51&#160;UTC | true | [view](CERTS/fab84634747b289457158fa4984de05cadcd77c62855912a90ee98eb672b714f/README.md) |
| 29&#160;Mar&#160;24&#160;17:55&#160;UTC | SHAKEN 597F VoIP Innovations | 29&#160;Mar&#160;25&#160;17:55&#160;UTC | true | [view](CERTS/daa381020f162344a666c69b83e4adff78c7438c494521ce9b084dee47c37f49/README.md) |
| 01&#160;Apr&#160;24&#160;16:18&#160;UTC | SHAKEN 731K ChaseData Corp | 01&#160;Apr&#160;25&#160;16:18&#160;UTC | true | [view](CERTS/5e54f96ef86b3bfa4d20d0d4338366e809bc1cb7dd402fadc718bcecdd20de08/README.md) |
| 02&#160;Apr&#160;24&#160;22:06&#160;UTC | SHAKEN 056K Ascension Technologies | 02&#160;Apr&#160;25&#160;22:06&#160;UTC | true | [view](CERTS/b6dc8e79b26eaaaa7e9a1419681b884b85797c027ade6d4153821f46b18bb924/README.md) |
| 16&#160;Apr&#160;24&#160;00:42&#160;UTC | SHAKEN 525K AU Telecom | 16&#160;Apr&#160;25&#160;00:42&#160;UTC | true | [view](CERTS/694c1778f23e11db0dd3bdc9b48d3c96b62bba793f6883faa7d36776960ddc94/README.md) |
| 16&#160;Apr&#160;24&#160;01:44&#160;UTC | SHAKEN WWT INC dba VoIP Networks 053K | 16&#160;Apr&#160;25&#160;01:44&#160;UTC | true | [view](CERTS/06aa18d1b35b3d8560011f9b40f33801099a772c836ee30cf0c6677d08931cf5/README.md) |
| 16&#160;Apr&#160;24&#160;18:47&#160;UTC | SHAKEN 839J EON Telecom Inc. | 16&#160;Apr&#160;25&#160;18:47&#160;UTC | true | [view](CERTS/e732c0ce34c1ac96d34e6489654b580c6c96a076ff3c082bfd5d18946967c4fb/README.md) |
| 16&#160;Apr&#160;24&#160;20:38&#160;UTC | SHAKEN 894K SIPVOICE, LLC | 16&#160;Apr&#160;25&#160;20:38&#160;UTC | true | [view](CERTS/8f9b473e3a832636d10e3e83d66fc6c2298211e646bc7fa926fdfcc40875a177/README.md) |
| 17&#160;Apr&#160;24&#160;00:18&#160;UTC | SHAKEN 879K COMMUNICATIONS AND TELEPHONE SYSTEMS | 17&#160;Apr&#160;25&#160;00:18&#160;UTC | true | [view](CERTS/4fa5e10f0f209f8ea7352ff6ac522324adeb1ba6c2fc07b9bc0a5fddd586ecbd/README.md) |
| 17&#160;Apr&#160;24&#160;18:03&#160;UTC | SHAKEN 782J AM Communication Labs Inc. | 17&#160;Apr&#160;25&#160;18:03&#160;UTC | true | [view](CERTS/3abea89f7a0c9752a0af294e642070be6c7edd3fcf1f23c239708c1d2f798234/README.md) |
| 23&#160;Apr&#160;24&#160;20:24&#160;UTC | SHAKEN 089K Logista | 23&#160;Apr&#160;25&#160;20:24&#160;UTC | true | [view](CERTS/494eca4dc32b870f9dd3a0de68c1e36cc5dd601645447dcd19e871b74d4ce761/README.md) |
| 23&#160;Apr&#160;24&#160;20:32&#160;UTC | SHAKEN GIP Technology 434K | 23&#160;Apr&#160;25&#160;20:32&#160;UTC | true | [view](CERTS/8b89b7753c70f6c5e7642831b901ae0f37ce3be14e2a9bc4ec4b05646b23b330/README.md) |
| 23&#160;Apr&#160;24&#160;21:28&#160;UTC | SHAKEN 895K 6x6 Termination | 23&#160;Apr&#160;25&#160;21:28&#160;UTC | true | [view](CERTS/bef1d8df3e931bfba5331105c5375726011d60158f2463b25466bfd537ef1243/README.md) |
| 24&#160;Apr&#160;24&#160;05:56&#160;UTC | SHAKEN 492K Telxio Networks | 24&#160;Apr&#160;25&#160;05:56&#160;UTC | true | [view](CERTS/cf2eeaf7f14dcbcfac58ed276f8695116b260ed946dc7be13d5a66c07d2caf83/README.md) |
| 08&#160;May&#160;24&#160;00:00&#160;UTC | SHAKEN Nextiva, Inc 914H | 08&#160;May&#160;25&#160;00:00&#160;UTC | true | [view](CERTS/b1e659bd56d018982cb5db1667f7fdd898a2679f37cd2e75adec0b4fda02df8c/README.md) |
| 08&#160;May&#160;24&#160;16:51&#160;UTC | SHAKEN 267K Starlinq PBX Inc. | 08&#160;May&#160;25&#160;16:51&#160;UTC | true | [view](CERTS/f4180665065d7d42522d2ff11a94fb0b18be7c242bf316f57325219909d93828/README.md) |
| 22&#160;May&#160;24&#160;16:04&#160;UTC | SHAKEN ALD Telecom 780J | 22&#160;May&#160;25&#160;16:04&#160;UTC | true | [view](CERTS/af5a628bdfb896413578e9aae6e50c94fd8fb932a9535775b16af741f70f1708/README.md) |
| 23&#160;May&#160;24&#160;12:05&#160;UTC | SHAKEN Midwest Telecom of America 919A | 23&#160;May&#160;25&#160;12:05&#160;UTC | true | [view](CERTS/ec4710a2981dc9a9be4de30da579880ce1ddbdf830a53a2e790217bf5c13c6d4/README.md) |
| 28&#160;May&#160;24&#160;12:23&#160;UTC | SHAKEN Medtel Communications 994J | 28&#160;May&#160;25&#160;12:23&#160;UTC | true | [view](CERTS/6f77234f6369fa640a847f415b2c58f16044fc86453a2869d5d8f72d6dfd71a6/README.md) |
| 31&#160;May&#160;24&#160;19:33&#160;UTC | SHAKEN 938K Lynktel | 31&#160;May&#160;25&#160;19:33&#160;UTC | true | [view](CERTS/b2aa0de697866f9e9dec031d059f9a3d478ee21f540ff87ebe0a8c429527b220/README.md) |
| 09&#160;Jun&#160;24&#160;08:04&#160;UTC | SHAKEN 776K ActoinVox | 06&#160;Dec&#160;24&#160;08:04&#160;UTC | true | [view](CERTS/87b5e4ebe2f75f3bb1c8efe86e8a26d08435b5672243786f8a7f96d349694219/README.md) |
| 21&#160;Jun&#160;24&#160;20:19&#160;UTC | SHAKEN Connexum LLC 203K | 21&#160;Jun&#160;25&#160;20:19&#160;UTC | true | [view](CERTS/c405face9873a72f5ec4f1bc395a4124d7cfbf080a704060ebe5e9bf7a07d62d/README.md) |
| 01&#160;Jul&#160;24&#160;16:21&#160;UTC | SHAKEN Mercury Access Solutions 634K | 01&#160;Jul&#160;25&#160;16:21&#160;UTC | true | [view](CERTS/70cc27d2e747142c3446b6133993989be4e1ff3fe4d25b137def2745fc173d19/README.md) |
| 02&#160;Jul&#160;24&#160;15:32&#160;UTC | SHAKEN Bulk Solutions, LLC 644J | 31&#160;Jan&#160;25&#160;15:32&#160;UTC | true | [view](CERTS/9da93d10ef266f23087e3f23bd41ba68e75db57c81dd5099c76ae27074a16d47/README.md) |
| 12&#160;Jul&#160;24&#160;07:20&#160;UTC | SHAKEN 521K Voice SY LLC | 28&#160;Jan&#160;25&#160;07:20&#160;UTC | true | [view](CERTS/4543540ab57936221819dbdef685a8a4b09c87e6b72dd21fcbb346696de31a3b/README.md) |
| 23&#160;Jul&#160;24&#160;16:35&#160;UTC | SHAKEN Lightspeed Voice 557F | 22&#160;Aug&#160;24&#160;16:35&#160;UTC | true | [view](CERTS/22e5ec5478d0200e3911c65f2577f4996acfcef70d54872af51360f537fbd6f6/README.md) |
| 24&#160;Jul&#160;24&#160;00:13&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 23&#160;Aug&#160;24&#160;00:13&#160;UTC | true | [view](CERTS/71ef03308aa96f26f0e7430ea56b7d523d2cebe2c02fe79c2c141e899fbaa88f/README.md) |
| 24&#160;Jul&#160;24&#160;10:03&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 23&#160;Aug&#160;24&#160;10:03&#160;UTC | true | [view](CERTS/7ee1b0c7f07110553f906c97338cbf195c51e239249265d48d3f7bfa451d6de1/README.md) |
| 24&#160;Jul&#160;24&#160;14:04&#160;UTC | SHAKEN BareTelecom 864J | 23&#160;Aug&#160;24&#160;14:04&#160;UTC | true | [view](CERTS/7bacb5f3d5e909279a7489e366ebaeed5625a8f64a6bd7a5b164db555679583e/README.md) |
| 24&#160;Jul&#160;24&#160;14:11&#160;UTC | SHAKEN IDT America, Corp 363A | 23&#160;Aug&#160;24&#160;14:11&#160;UTC | true | [view](CERTS/2459b7dca5d4af419cb38891df59b181461e4fe620182f26b809e054efb6cc90/README.md) |
| 24&#160;Jul&#160;24&#160;23:27&#160;UTC | SHAKEN Consolidated Communications 5113 | 23&#160;Aug&#160;24&#160;23:27&#160;UTC | true | [view](CERTS/c81db59b18b9754badad15474d7dd3d9f2891007610e5e551a5fa1e2d809f34f/README.md) |
| 24&#160;Jul&#160;24&#160;23:28&#160;UTC | SHAKEN Touchtone 683A | 23&#160;Aug&#160;24&#160;23:28&#160;UTC | true | [view](CERTS/f8b3c1ddf87468ef14e93fac52c29dff88f1e6b55ac4e3d398ef7019184e5bb5/README.md) |
| 24&#160;Jul&#160;24&#160;23:29&#160;UTC | SHAKEN Apeiron Systems 012J | 23&#160;Aug&#160;24&#160;23:29&#160;UTC | true | [view](CERTS/54c553646fd836f295f22516dd5054cfd1c6eeb7b53517835cc25c1135978288/README.md) |
| 24&#160;Jul&#160;24&#160;23:31&#160;UTC | SHAKEN Fonative, Inc. 684J | 23&#160;Aug&#160;24&#160;23:31&#160;UTC | true | [view](CERTS/702a5c0afd220f39319d9568f1a242839317666a1aae485b851602726d597064/README.md) |
| 24&#160;Jul&#160;24&#160;23:32&#160;UTC | SHAKEN IPitomy 652J | 23&#160;Aug&#160;24&#160;23:32&#160;UTC | true | [view](CERTS/9f7e4f18a9cd3871dad5b9f2eef509c6627f0bb73d586a3745dc79f28c01c6af/README.md) |
| 24&#160;Jul&#160;24&#160;23:37&#160;UTC | SHAKEN Phone.com, Inc. 633J | 23&#160;Aug&#160;24&#160;23:37&#160;UTC | true | [view](CERTS/0fdfb4346c51ef76c91bebaef82a14ddd24237d842a89ee5cbf458a4aa0e36e9/README.md) |
| 24&#160;Jul&#160;24&#160;23:38&#160;UTC | SHAKEN NETRIO LLC 020K | 23&#160;Aug&#160;24&#160;23:38&#160;UTC | true | [view](CERTS/2140e0dae694d425088d35e4d74c149e9ad5f8d1f281ea7ea99798f46dda5210/README.md) |
| 24&#160;Jul&#160;24&#160;23:39&#160;UTC | SHAKEN Noble Systems Communications LLC 187J | 23&#160;Aug&#160;24&#160;23:39&#160;UTC | true | [view](CERTS/7359718d30407597b3f8782681f9f1f4dc3a1c8a7e3b1a610d9fbd9f77112a26/README.md) |
| 24&#160;Jul&#160;24&#160;23:40&#160;UTC | SHAKEN Airespring 996H | 23&#160;Aug&#160;24&#160;23:40&#160;UTC | true | [view](CERTS/03707d59dbf9a5370f8e0f30173531d5798ba858aa0955085e2daf11c4a6736b/README.md) |
| 24&#160;Jul&#160;24&#160;23:42&#160;UTC | SHAKEN Nobelbiz, Inc. 596J | 23&#160;Aug&#160;24&#160;23:42&#160;UTC | true | [view](CERTS/c9bb1a09eed2e99ab3d2a2f0b57fb6b0b5bfd5cd1f8802ae6d79cede58445469/README.md) |
| 24&#160;Jul&#160;24&#160;23:43&#160;UTC | SHAKEN Momentum Telecom 9157 | 23&#160;Aug&#160;24&#160;23:43&#160;UTC | true | [view](CERTS/5aee60c9d1a1a0df486118536c3f7bd0222147e744f8f76c515d28314d475e9d/README.md) |
| 24&#160;Jul&#160;24&#160;23:45&#160;UTC | SHAKEN Matrix 7379 | 23&#160;Aug&#160;24&#160;23:45&#160;UTC | true | [view](CERTS/ddef91ee458e1fae35d302090f0813849481f121cd001c93d1b402d21d6d0c03/README.md) |
| 24&#160;Jul&#160;24&#160;23:45&#160;UTC | SHAKEN Matrix 3058 | 23&#160;Aug&#160;24&#160;23:45&#160;UTC | true | [view](CERTS/acf86fb4519ffd201ff1a20c64a6bfc3fec4fa271a5138d0b04ef450b4e8dcf8/README.md) |
| 24&#160;Jul&#160;24&#160;23:45&#160;UTC | SHAKEN Matrix 9451 | 23&#160;Aug&#160;24&#160;23:45&#160;UTC | true | [view](CERTS/e4987b521c31012f660ba6f8f9a84cfb8f1a47c4271e15c08ff893ec4d5bddb5/README.md) |
| 24&#160;Jul&#160;24&#160;23:47&#160;UTC | SHAKEN PNG Telecommunications Inc 3395 | 23&#160;Aug&#160;24&#160;23:47&#160;UTC | true | [view](CERTS/cb9fb076e5187d06e1baa502c4160cbe0cbd73b1e362212967ba7e268fd22319/README.md) |
| 24&#160;Jul&#160;24&#160;23:49&#160;UTC | SHAKEN Magna5, LLC 3849 | 23&#160;Aug&#160;24&#160;23:49&#160;UTC | true | [view](CERTS/d8c2341090c0e4a2eb980a22d382289d557b7cbc47b24b468bd36152aa7f0318/README.md) |
| 24&#160;Jul&#160;24&#160;23:50&#160;UTC | SHAKEN Magna5, LLC 8249 | 23&#160;Aug&#160;24&#160;23:50&#160;UTC | true | [view](CERTS/069b14c2ceb1990a7e1ec186565405c5e8eeb068ca27238264605af6c1ad9d16/README.md) |
| 25&#160;Jul&#160;24&#160;05:34&#160;UTC | SHAKEN IDT America, Corp 363A | 24&#160;Aug&#160;24&#160;05:34&#160;UTC | true | [view](CERTS/92f0a0265a4bb146c97120e0c5c9997843008985e3c85ad63a2435da4010ba07/README.md) |
| 25&#160;Jul&#160;24&#160;06:29&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 24&#160;Aug&#160;24&#160;06:29&#160;UTC | true | [view](CERTS/57a0df899aac56244a352d92ba50db374f2e779889679df86c442ec8a61d9ca5/README.md) |
| 25&#160;Jul&#160;24&#160;13:59&#160;UTC | SHAKEN BareTelecom 864J | 24&#160;Aug&#160;24&#160;13:59&#160;UTC | true | [view](CERTS/50c1bac87db6d47fd07d32a8e5d34d73f8c21cbe04d37c7837458292b2fce063/README.md) |
| 26&#160;Jul&#160;24&#160;05:29&#160;UTC | SHAKEN IDT America, Corp 363A | 25&#160;Aug&#160;24&#160;05:29&#160;UTC | true | [view](CERTS/35dc589e519d4fdef080bc75b3ad78a963b649482d565ea241f9b061b7e18f42/README.md) |
| 26&#160;Jul&#160;24&#160;06:24&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 25&#160;Aug&#160;24&#160;06:24&#160;UTC | true | [view](CERTS/ece64cba5d222fb611cdfcc232929222f126c63dbf9cc5ce51f2940e23e04863/README.md) |
| 26&#160;Jul&#160;24&#160;13:54&#160;UTC | SHAKEN BareTelecom 864J | 25&#160;Aug&#160;24&#160;13:54&#160;UTC | true | [view](CERTS/308107bb8fe47bbcc700c4729599fa05f4a99d530e6d4b16598bb2e1334854b3/README.md) |
| 27&#160;Jul&#160;24&#160;13:49&#160;UTC | SHAKEN BareTelecom 864J | 26&#160;Aug&#160;24&#160;13:49&#160;UTC | true | [view](CERTS/2f9eb3442c7a66476ef5a5e72814070c73a12eda0715e274702de1a301c7737c/README.md) |
| 29&#160;Jul&#160;24&#160;02:37&#160;UTC | SHAKEN IPSBS Managed Services LLC 828J | 28&#160;Aug&#160;24&#160;02:37&#160;UTC | true | [view](CERTS/4966af38a9c4ce5abf595562142e8eb2e4de9da561cce06ccec93fda3b12e1fe/README.md) |
| 29&#160;Jul&#160;24&#160;03:59&#160;UTC | SHAKEN BareTelecom 864J | 28&#160;Aug&#160;24&#160;03:59&#160;UTC | true | [view](CERTS/33a4ce4fd1a250d8cc06e50cc67ec4e6666d5cc6a635ec2bacfd928ecd74332c/README.md) |
| 29&#160;Jul&#160;24&#160;05:14&#160;UTC | SHAKEN IDT America, Corp 363A | 28&#160;Aug&#160;24&#160;05:14&#160;UTC | true | [view](CERTS/c1c239d92eba75ea7af9ae4d5f18f7bcc2ef123002d7066dc895cc684abbb3c7/README.md) |
| 29&#160;Jul&#160;24&#160;06:09&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 28&#160;Aug&#160;24&#160;06:09&#160;UTC | true | [view](CERTS/bc2c328d273239dec21c2e8c036aa912a5bbc9b8ed4ed12f84a4bf457663c4a5/README.md) |
| 29&#160;Jul&#160;24&#160;09:38&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 28&#160;Aug&#160;24&#160;09:38&#160;UTC | true | [view](CERTS/1ca60ca643b00ebe8cef6af9b966155ecf431fb7d6cf723e41de3c347cadbbbb/README.md) |
| 29&#160;Jul&#160;24&#160;17:44&#160;UTC | SHAKEN Mango Voice LLC 579K | 28&#160;Aug&#160;24&#160;17:44&#160;UTC | true | [view](CERTS/e9df8eb80e4bfadc87682b847068187c00238a9c0cf49e1814e4eac09d299f5d/README.md) |
| 30&#160;Jul&#160;24&#160;02:32&#160;UTC | SHAKEN IPSBS Managed Services LLC 828J | 29&#160;Aug&#160;24&#160;02:32&#160;UTC | true | [view](CERTS/81b4f702b03a456491b86bd49f524445233ed9a84c2ad77f529dad2d4332d771/README.md) |
| 30&#160;Jul&#160;24&#160;05:09&#160;UTC | SHAKEN IDT America, Corp 363A | 29&#160;Aug&#160;24&#160;05:09&#160;UTC | true | [view](CERTS/c5ab43c0c203011c239efb5de90fc3960328adb873cd5c24ba31d3417c863da4/README.md) |
| 30&#160;Jul&#160;24&#160;06:04&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 29&#160;Aug&#160;24&#160;06:04&#160;UTC | true | [view](CERTS/45d269a3018e077c77f16b165768c81dd5a427e6bd884669d9f8b3e8aeb6dae0/README.md) |
| 30&#160;Jul&#160;24&#160;09:33&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 29&#160;Aug&#160;24&#160;09:33&#160;UTC | true | [view](CERTS/7e4c39d5be1b76a883c54b1c82838a76e47946e89eb7ee21ac2ec9d2d67ada3d/README.md) |
| 30&#160;Jul&#160;24&#160;16:00&#160;UTC | SHAKEN Lightspeed Voice 557F | 29&#160;Aug&#160;24&#160;16:00&#160;UTC | true | [view](CERTS/774780dc99f048e65326c4cbe5d2ef99344f82502e0cc122fb2a114e51529cdd/README.md) |
| 31&#160;Jul&#160;24&#160;02:27&#160;UTC | SHAKEN IPSBS Managed Services LLC 828J | 30&#160;Aug&#160;24&#160;02:27&#160;UTC | true | [view](CERTS/fb0172b6edaaa0e79af7d538b2a6b6d7f963110b7f7a29a3b2d107de0ea3ae02/README.md) |
| 31&#160;Jul&#160;24&#160;05:04&#160;UTC | SHAKEN IDT America, Corp 363A | 30&#160;Aug&#160;24&#160;05:04&#160;UTC | true | [view](CERTS/ec9e7bd575d805fc2f88804b85ba51277e0e4d7a85a097f1a74404ea27534a90/README.md) |
| 31&#160;Jul&#160;24&#160;05:59&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 30&#160;Aug&#160;24&#160;05:59&#160;UTC | true | [view](CERTS/1285c852fc5725fb80343a01b4609486df1c0f465a979c15945e06cd8a54f71d/README.md) |
| 31&#160;Jul&#160;24&#160;09:28&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 30&#160;Aug&#160;24&#160;09:28&#160;UTC | true | [view](CERTS/806f14c7aa6f029f167b75adf184fa5760ad34dbbfb3e10e23bb50f97d1ca029/README.md) |
| 31&#160;Jul&#160;24&#160;15:55&#160;UTC | SHAKEN Lightspeed Voice 557F | 30&#160;Aug&#160;24&#160;15:55&#160;UTC | true | [view](CERTS/775cbdb5fe9ff76db1c3615c26f006a66fab0cdcca7b976c84647fa64739ca9a/README.md) |
| 01&#160;Aug&#160;24&#160;03:43&#160;UTC | SHAKEN BareTelecom 864J | 31&#160;Aug&#160;24&#160;03:43&#160;UTC | true | [view](CERTS/1864ee4b26099a7b12a08454509479ddb0815963abe18de3ded13eba18611581/README.md) |
| 01&#160;Aug&#160;24&#160;05:00&#160;UTC | SHAKEN IDT America, Corp 363A | 31&#160;Aug&#160;24&#160;05:00&#160;UTC | true | [view](CERTS/61303349cdd1c6b1dcd34675d557eabad8766a38c71e619fd2c6f1794f8119ac/README.md) |
| 01&#160;Aug&#160;24&#160;09:23&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 31&#160;Aug&#160;24&#160;09:23&#160;UTC | true | [view](CERTS/d9e556ea1952af7995e40690a4681f688db65259c733b50a4fdcb5bb127bc65b/README.md) |
| 02&#160;Aug&#160;24&#160;02:17&#160;UTC | SHAKEN IPSBS Managed Services LLC 828J | 01&#160;Sep&#160;24&#160;02:17&#160;UTC | true | [view](CERTS/648ef6a285ac85d465f1af128d676f05e28124b26ea19cdee5b5c516d42a9d41/README.md) |
| 02&#160;Aug&#160;24&#160;03:38&#160;UTC | SHAKEN BareTelecom 864J | 01&#160;Sep&#160;24&#160;03:38&#160;UTC | true | [view](CERTS/81c51697772bcf39f46cabdd8891e2143263626c8cd05db4a78dfc61d9b8d5e0/README.md) |
| 02&#160;Aug&#160;24&#160;05:49&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 01&#160;Sep&#160;24&#160;05:49&#160;UTC | true | [view](CERTS/916d92cd712c8a57f4c495da0d9103d3782e19ebf70dcb84867fa15c9868daab/README.md) |
| 02&#160;Aug&#160;24&#160;09:18&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 01&#160;Sep&#160;24&#160;09:18&#160;UTC | true | [view](CERTS/90eaf67b11eaecab2914cde3821cadddc20d3b7186bc53a6bda02b1dcf0a6f7c/README.md) |
| 03&#160;Aug&#160;24&#160;03:33&#160;UTC | SHAKEN BareTelecom 864J | 02&#160;Sep&#160;24&#160;03:33&#160;UTC | true | [view](CERTS/d3f3168304ee2ea72dca6c32931e8f17232d4cb476d4c5d443a029635bc8b1af/README.md) |
| 03&#160;Aug&#160;24&#160;07:30&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 02&#160;Sep&#160;24&#160;07:30&#160;UTC | true | [view](CERTS/f462eb021877ef981267bae64c4f4415204e8a50bb3783a517c2d3f504076917/README.md) |
| 04&#160;Aug&#160;24&#160;04:44&#160;UTC | SHAKEN IDT America, Corp 363A | 03&#160;Sep&#160;24&#160;04:44&#160;UTC | true | [view](CERTS/2ed20264e4f0f0135143b1b037cc1c29926ba6f66d14e87f924d1a37fb99f164/README.md) |
| 04&#160;Aug&#160;24&#160;17:14&#160;UTC | SHAKEN Mango Voice LLC 579K | 03&#160;Sep&#160;24&#160;17:14&#160;UTC | true | [view](CERTS/c53193817c659f4d5e913b6b97bf501ecf14156566150105a0b64bd96b761219/README.md) |
| 05&#160;Aug&#160;24&#160;03:23&#160;UTC | SHAKEN BareTelecom 864J | 04&#160;Sep&#160;24&#160;03:23&#160;UTC | true | [view](CERTS/c9c7ea0d0edd80c46fa7ef1173b520565a9b54a82daab6d611a046ff56f4b906/README.md) |
| 05&#160;Aug&#160;24&#160;04:39&#160;UTC | SHAKEN IDT America, Corp 363A | 04&#160;Sep&#160;24&#160;04:39&#160;UTC | true | [view](CERTS/f1f0d72001e65fd1baa2b2f4f2fa9efd792b56a03605585738eb32a4493fb0da/README.md) |
| 05&#160;Aug&#160;24&#160;09:03&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 04&#160;Sep&#160;24&#160;09:03&#160;UTC | true | [view](CERTS/844e49538f2e3e59494651f83150657a9a84f24e6e2a1a9944053a01d234421d/README.md) |
| 05&#160;Aug&#160;24&#160;17:09&#160;UTC | SHAKEN Mango Voice LLC 579K | 04&#160;Sep&#160;24&#160;17:09&#160;UTC | true | [view](CERTS/a8d8aedc5e222dd0098ca4b965ac440440507ba1e48aed4778100253ff25acb8/README.md) |
| 05&#160;Aug&#160;24&#160;21:36&#160;UTC | SHAKEN Socket Telecom LLC 554a | 04&#160;Sep&#160;24&#160;21:36&#160;UTC | true | [view](CERTS/e5ec59732426cc9ebed1174ff061cc20fbbf897979a51576fda020f8e2903340/README.md) |
| 05&#160;Aug&#160;24&#160;22:32&#160;UTC | SHAKEN 696J Telcentris Inc. dba Voxox | 19&#160;Sep&#160;24&#160;22:32&#160;UTC | true | [view](CERTS/f599d892f9162a55fda61fc833f229a0ea6e3c7a2b4006731caeb0664611fbb9/README.md) |
| 06&#160;Aug&#160;24&#160;03:18&#160;UTC | SHAKEN BareTelecom 864J | 05&#160;Sep&#160;24&#160;03:18&#160;UTC | true | [view](CERTS/542de0476fa56bb6d01f14096918912a6549fd6e9883ab1df44a19a080c2aea9/README.md) |
| 06&#160;Aug&#160;24&#160;04:34&#160;UTC | SHAKEN IDT America, Corp 363A | 05&#160;Sep&#160;24&#160;04:34&#160;UTC | true | [view](CERTS/eeaf059cbbdf9955e5af0dbf2c013209c9d1f9fed2fb90529a0d9b85cdd6b112/README.md) |
| 06&#160;Aug&#160;24&#160;08:58&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 05&#160;Sep&#160;24&#160;08:58&#160;UTC | true | [view](CERTS/30899b6f0c4b43df0380c21080abba23f4e53a9d08891c14e77fc2c205e1226b/README.md) |
| 06&#160;Aug&#160;24&#160;15:25&#160;UTC | SHAKEN Lightspeed Voice 557F | 05&#160;Sep&#160;24&#160;15:25&#160;UTC | true | [view](CERTS/8ed8325b10fd23cd697b725e2d812a9d48ae9c0433ec0180004c7544e4dc4c60/README.md) |
| 06&#160;Aug&#160;24&#160;17:04&#160;UTC | SHAKEN Mango Voice LLC 579K | 05&#160;Sep&#160;24&#160;17:04&#160;UTC | true | [view](CERTS/23b1b2ce8f4ab6add2c834b788dc7bec60a773510a989321f8cd575ee7ef6a13/README.md) |
| 06&#160;Aug&#160;24&#160;21:31&#160;UTC | SHAKEN Socket Telecom LLC 554a | 05&#160;Sep&#160;24&#160;21:31&#160;UTC | true | [view](CERTS/b8b35a5b470142c1828ad8353726375cf6d79aef8db1e80918421a7fb9c990da/README.md) |
| 07&#160;Aug&#160;24&#160;03:13&#160;UTC | SHAKEN BareTelecom 864J | 06&#160;Sep&#160;24&#160;03:13&#160;UTC | true | [view](CERTS/cb11ee3834edfbf96206100a72d94eb9c50cf7edd519a5da71fdc6effa9012dc/README.md) |
| 07&#160;Aug&#160;24&#160;04:29&#160;UTC | SHAKEN IDT America, Corp 363A | 06&#160;Sep&#160;24&#160;04:29&#160;UTC | true | [view](CERTS/f25ed41ed449942c1a4d72971d8927c2325aaf25601241c096a489dd423be24b/README.md) |
| 07&#160;Aug&#160;24&#160;05:24&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 06&#160;Sep&#160;24&#160;05:24&#160;UTC | true | [view](CERTS/86ebc116991f197ba37ab6ed7150c5ecd041e216b4dc20c4900c0ab72033b03e/README.md) |
| 07&#160;Aug&#160;24&#160;08:29&#160;UTC | SHAKEN Townes Telecommunications 0335 | 06&#160;Sep&#160;24&#160;08:29&#160;UTC | true | [view](CERTS/8d5cad6058d5023a8281ef67754594dc90933b798ae280e0ee677c54be28c615/README.md) |
| 07&#160;Aug&#160;24&#160;08:53&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 06&#160;Sep&#160;24&#160;08:53&#160;UTC | true | [view](CERTS/e48b4d96f4e28d2b3be4fa6d2a8cd9203fe3f2de2b457a2ba1fa95744182ae71/README.md) |
| 07&#160;Aug&#160;24&#160;10:16&#160;UTC | SHAKEN Convoso 758J | 11&#160;Sep&#160;24&#160;10:16&#160;UTC | true | [view](CERTS/5ba3cb27852e0ba6f5b298682288518b4641306a664358825955e23a6a5f47fd/README.md) |
| 07&#160;Aug&#160;24&#160;16:59&#160;UTC | SHAKEN Mango Voice LLC 579K | 06&#160;Sep&#160;24&#160;16:59&#160;UTC | true | [view](CERTS/0833146cfb8fd48f3bf370df1c366a73c8cf03a28dbfd3f78df25ba4d1751c13/README.md) |
| 07&#160;Aug&#160;24&#160;21:26&#160;UTC | SHAKEN Socket Telecom LLC 554a | 06&#160;Sep&#160;24&#160;21:26&#160;UTC | true | [view](CERTS/a05b0edc3b290df4bc6b918216ed13af96c0ab9b62d671080dd89c6e84310e23/README.md) |
| 08&#160;Aug&#160;24&#160;03:08&#160;UTC | SHAKEN BareTelecom 864J | 07&#160;Sep&#160;24&#160;03:08&#160;UTC | true | [view](CERTS/d94eaf738b11c56cd97ec05bfdb6b8d8e0c6a23bd3b57da3d3732efff54cb1ec/README.md) |
| 08&#160;Aug&#160;24&#160;04:24&#160;UTC | SHAKEN IDT America, Corp 363A | 07&#160;Sep&#160;24&#160;04:24&#160;UTC | true | [view](CERTS/f25ff6dfe9836c6b786acacd2a4e35c08cae0c3b121feefc0a9b6123f37e57b6/README.md) |
| 08&#160;Aug&#160;24&#160;08:48&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 07&#160;Sep&#160;24&#160;08:48&#160;UTC | true | [view](CERTS/7c0037d93b91aab645b1385c23a6f9ecc90bcb7feb1093735d4996a888557d89/README.md) |
| 08&#160;Aug&#160;24&#160;18:15&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 07&#160;Sep&#160;24&#160;18:15&#160;UTC | true | [view](CERTS/0400b2d478591d3a53dc6d483d8bb7f934ce5e4a9121708da64eb1b87ea9c79f/README.md) |
| 09&#160;Aug&#160;24&#160;04:19&#160;UTC | SHAKEN IDT America, Corp 363A | 08&#160;Sep&#160;24&#160;04:19&#160;UTC | true | [view](CERTS/b60cfed1abda5ca51812543b4afc60e63ed70a9e29c028f180db0850131443aa/README.md) |
| 09&#160;Aug&#160;24&#160;05:04&#160;UTC | SHAKEN BareTelecom 864J | 08&#160;Sep&#160;24&#160;05:04&#160;UTC | true | [view](CERTS/1de0926e9a8dea55364b51a58196040bbb78a69c6aed94180a965478d1c5fd88/README.md) |
| 09&#160;Aug&#160;24&#160;05:14&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 08&#160;Sep&#160;24&#160;05:14&#160;UTC | true | [view](CERTS/594aadad75018973ae9750204481323070168bdfb4e4f11685f76c0aea7d9c44/README.md) |
| 09&#160;Aug&#160;24&#160;07:00&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 08&#160;Sep&#160;24&#160;07:00&#160;UTC | true | [view](CERTS/01d7fe4fe886c53e2f27fe24ecd926d5b024338554a503e11ca4fb4716aca8fd/README.md) |
| 09&#160;Aug&#160;24&#160;18:10&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 08&#160;Sep&#160;24&#160;18:10&#160;UTC | true | [view](CERTS/890d4d82e3636b672502c2cca5191ae5fef76273548d106364d4a30ea954b40a/README.md) |
| 10&#160;Aug&#160;24&#160;13:36&#160;UTC | SHAKEN 688K Call Hub Inc. | 10&#160;Aug&#160;25&#160;13:36&#160;UTC | true | [view](CERTS/a96a9718468b7536ed8286847382a28e9aae5b645c44621b234321ccddcf5c2a/README.md) |
| 10&#160;Aug&#160;24&#160;19:01&#160;UTC | SHAKEN 688K Call Hub Inc. | 10&#160;Aug&#160;25&#160;19:01&#160;UTC | true | [view](CERTS/d8fe53673498502aa06d1deb531144cdeac93dba56e9c7094fef9fe192b5b14c/README.md) |
| 11&#160;Aug&#160;24&#160;18:00&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 10&#160;Sep&#160;24&#160;18:00&#160;UTC | true | [view](CERTS/fcfc2129427f07df2811a569a5ee13234ba1c56da272cccefd4a6a902eccdd58/README.md) |
| 12&#160;Aug&#160;24&#160;04:04&#160;UTC | SHAKEN IDT America, Corp 363A | 11&#160;Sep&#160;24&#160;04:04&#160;UTC | true | [view](CERTS/65bfed7e481912ae4d67b5548e02db9cef1e048cdb337168859a1cf5afaed0e0/README.md) |
| 12&#160;Aug&#160;24&#160;04:49&#160;UTC | SHAKEN BareTelecom 864J | 11&#160;Sep&#160;24&#160;04:49&#160;UTC | true | [view](CERTS/e48e2b79a6af1b2d95423fe4afaebc7966a38ad79c786cf7e7d6f1633a14d7e2/README.md) |
| 12&#160;Aug&#160;24&#160;06:45&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 11&#160;Sep&#160;24&#160;06:45&#160;UTC | true | [view](CERTS/2d54077b37b7baeefa30b1bc7584e5b947468a326b0990f3b32bc36728960aac/README.md) |
| 12&#160;Aug&#160;24&#160;17:55&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 11&#160;Sep&#160;24&#160;17:55&#160;UTC | true | [view](CERTS/7a3f5b989705dc3eaf856afd60f0a7e75a30e6d29fc0253e16b560b935920f27/README.md) |
| 13&#160;Aug&#160;24&#160;03:59&#160;UTC | SHAKEN IDT America, Corp 363A | 12&#160;Sep&#160;24&#160;03:59&#160;UTC | true | [view](CERTS/a4ae4a5ca8221149d6147d1d54abcf2a8f9fa643015e6bb47c1a31a321d863b6/README.md) |
| 13&#160;Aug&#160;24&#160;04:44&#160;UTC | SHAKEN BareTelecom 864J | 12&#160;Sep&#160;24&#160;04:44&#160;UTC | true | [view](CERTS/db5dd7ca88e557f6851d1cbf9f83f179645f6d3fe088ced0c788ad56be300a38/README.md) |
| 13&#160;Aug&#160;24&#160;17:50&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 12&#160;Sep&#160;24&#160;17:50&#160;UTC | true | [view](CERTS/4b23998fbde1eec5f48fce45c654d4b3042eee11a74adec86802a5dbd318c84e/README.md) |
| 14&#160;Aug&#160;24&#160;03:54&#160;UTC | SHAKEN IDT America, Corp 363A | 13&#160;Sep&#160;24&#160;03:54&#160;UTC | true | [view](CERTS/9fb0b443b19c3bce7f4b641ae0ae1ab3c743876c8177ba619e4859f60568b2c8/README.md) |
| 14&#160;Aug&#160;24&#160;04:39&#160;UTC | SHAKEN BareTelecom 864J | 13&#160;Sep&#160;24&#160;04:39&#160;UTC | true | [view](CERTS/c2e57271dd3d79eb1b466c94ae4fd2be822384721f86ca363b660aee19171860/README.md) |
| 14&#160;Aug&#160;24&#160;06:35&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 13&#160;Sep&#160;24&#160;06:35&#160;UTC | true | [view](CERTS/eec0d793ea6d9fae9f4b756e3520f2468a97d523d33aa1b2f0ea828fde2107dc/README.md) |
| 14&#160;Aug&#160;24&#160;13:48&#160;UTC | SHAKEN Consolidated Communications 5113 | 13&#160;Sep&#160;24&#160;13:48&#160;UTC | true | [view](CERTS/dfffad982775d643f453a7fedcae6f6202f121c5e32c8175082a9f51e1548878/README.md) |
| 14&#160;Aug&#160;24&#160;13:49&#160;UTC | SHAKEN Touchtone 683A | 13&#160;Sep&#160;24&#160;13:49&#160;UTC | true | [view](CERTS/ea1c2bd35ccf27e2365bb748ea729cc355b5a6d453f828a71cab841c406ebe32/README.md) |
| 14&#160;Aug&#160;24&#160;13:51&#160;UTC | SHAKEN Apeiron Systems 012J | 13&#160;Sep&#160;24&#160;13:51&#160;UTC | true | [view](CERTS/e39287bee4eeff8f23f4f4b98255aeed511b9505e7e2ee6b3befe6468fc2aa56/README.md) |
| 14&#160;Aug&#160;24&#160;13:51&#160;UTC | SHAKEN Fonative, Inc. 684J | 13&#160;Sep&#160;24&#160;13:51&#160;UTC | true | [view](CERTS/0ef3f3968b30cbc8df7c14b8c736e04b0839fe747572d3c15af80fa69f743980/README.md) |
| 14&#160;Aug&#160;24&#160;13:53&#160;UTC | SHAKEN IPitomy 652J | 13&#160;Sep&#160;24&#160;13:53&#160;UTC | true | [view](CERTS/ec80026f58091b31ffa3741aa8744eabb89be465e4a8582a84ac564d87af8017/README.md) |
| 14&#160;Aug&#160;24&#160;13:53&#160;UTC | SHAKEN Phone.com, Inc. 633J | 13&#160;Sep&#160;24&#160;13:53&#160;UTC | true | [view](CERTS/093a717eccd59c0d07ba45c14205a65746999321cfef3049d75a461c2103eeab/README.md) |
| 14&#160;Aug&#160;24&#160;13:54&#160;UTC | SHAKEN NETRIO LLC 020K | 13&#160;Sep&#160;24&#160;13:54&#160;UTC | true | [view](CERTS/6f749da437a66f8af80062de0b80e0fc8ed1c946bee31dbb55abf68a8b1b635e/README.md) |
| 14&#160;Aug&#160;24&#160;13:57&#160;UTC | SHAKEN Noble Systems Communications LLC 187J | 13&#160;Sep&#160;24&#160;13:57&#160;UTC | true | [view](CERTS/2dcc77596e28b1217639d34ce10382ff87ef7efe56e8f0f64f83be3d25a0cc42/README.md) |
| 14&#160;Aug&#160;24&#160;14:18&#160;UTC | SHAKEN Airespring 996H | 13&#160;Sep&#160;24&#160;14:18&#160;UTC | true | [view](CERTS/6717d421d01d9f5a0dd820b6bb46f1d232647ce49d6389b72a0ce31a275afdd3/README.md) |
| 14&#160;Aug&#160;24&#160;14:19&#160;UTC | SHAKEN Nobelbiz, Inc. 596J | 13&#160;Sep&#160;24&#160;14:19&#160;UTC | true | [view](CERTS/b3125bade8576d8311a804f83435167b9fd4b448a973886b15d13a681551816b/README.md) |
| 14&#160;Aug&#160;24&#160;14:20&#160;UTC | SHAKEN Momentum Telecom 9157 | 13&#160;Sep&#160;24&#160;14:20&#160;UTC | true | [view](CERTS/a19aa34f144a9e8d351d3cf5c9e3572015bbf137e172a1ce8246637878fe938a/README.md) |
| 14&#160;Aug&#160;24&#160;14:22&#160;UTC | SHAKEN Matrix 7379 | 13&#160;Sep&#160;24&#160;14:22&#160;UTC | true | [view](CERTS/e68154c81586f934445c58e7713d7630b99e988eb49ca70eaa796d0798a15d46/README.md) |
| 14&#160;Aug&#160;24&#160;14:29&#160;UTC | SHAKEN Magna5, LLC 3849 | 13&#160;Sep&#160;24&#160;14:29&#160;UTC | true | [view](CERTS/7be694846d9339a8de5c3640b14503ac903877210ec8343620fc0cd5f5eb04ca/README.md) |
| 14&#160;Aug&#160;24&#160;14:30&#160;UTC | SHAKEN Magna5, LLC 8249 | 13&#160;Sep&#160;24&#160;14:30&#160;UTC | true | [view](CERTS/8cf85da9854d623d4c3a4a86ce5f345b4d2fc65b732ba8689f47e2ca6ce3f1b7/README.md) |
| 14&#160;Aug&#160;24&#160;14:45&#160;UTC | SHAKEN Lightspeed Voice 557F | 13&#160;Sep&#160;24&#160;14:45&#160;UTC | true | [view](CERTS/8bedf0648ae135bd5f9ca11217cacb3daf68e0de494ba2a9b68014ebd7a2c08e/README.md) |
| 14&#160;Aug&#160;24&#160;16:24&#160;UTC | SHAKEN Mango Voice LLC 579K | 13&#160;Sep&#160;24&#160;16:24&#160;UTC | true | [view](CERTS/e1e962a766c4690320dc5eda4f3b237aa59c0436fb032b513fd9003a3a04ffe6/README.md) |
| 14&#160;Aug&#160;24&#160;17:45&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 13&#160;Sep&#160;24&#160;17:45&#160;UTC | true | [view](CERTS/d29772d16baaddf00fbd642a4f2e5e9cc83fda476543b94fff40a15293dbdddb/README.md) |
| 14&#160;Aug&#160;24&#160;20:54&#160;UTC | SHAKEN Socket Telecom LLC 554a | 13&#160;Sep&#160;24&#160;20:54&#160;UTC | true | [view](CERTS/e4ae56c15ca63c2d79841003c7c26c3903224220b5855adba888bcda9c761bf6/README.md) |
| 14&#160;Aug&#160;24&#160;22:14&#160;UTC | SHAKEN Zray Technologies Corporation 862J | 13&#160;Sep&#160;24&#160;22:14&#160;UTC | true | [view](CERTS/bfca112440bdbc27acd9d83576c96fc627d3924bc85404200f3be7245e66164c/README.md) |
| 15&#160;Aug&#160;24&#160;03:49&#160;UTC | SHAKEN IDT America, Corp 363A | 14&#160;Sep&#160;24&#160;03:49&#160;UTC | true | [view](CERTS/9634abc71503a8764370028cc0729e43c8824a756400a46074b0bbd6e3359e67/README.md) |
| 15&#160;Aug&#160;24&#160;04:34&#160;UTC | SHAKEN BareTelecom 864J | 14&#160;Sep&#160;24&#160;04:34&#160;UTC | true | [view](CERTS/2683963999843a9e81409d8262f30c8b34e3aa173060652e72d4956642f56f08/README.md) |
| 15&#160;Aug&#160;24&#160;04:44&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 14&#160;Sep&#160;24&#160;04:44&#160;UTC | true | [view](CERTS/59889a8e83e061fbbdde52e72142f183bae794e706ecbf5c291b9d065ee4f885/README.md) |
| 15&#160;Aug&#160;24&#160;06:30&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 14&#160;Sep&#160;24&#160;06:30&#160;UTC | true | [view](CERTS/d759606d8cc65132c5fd474290591088ba206b0d7644dd0dcedd61a4977e9fa0/README.md) |
| 15&#160;Aug&#160;24&#160;14:40&#160;UTC | SHAKEN Lightspeed Voice 557F | 14&#160;Sep&#160;24&#160;14:40&#160;UTC | true | [view](CERTS/7b6153524300cf5dce460a6a696360d0c86f72cf932d0d01a980aa3676fe189a/README.md) |
| 15&#160;Aug&#160;24&#160;17:40&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 14&#160;Sep&#160;24&#160;17:40&#160;UTC | true | [view](CERTS/818c7a3e95fb08d101cce8bfac64d8bca10b021a490659b6e6a6977e13e779c8/README.md) |
| 16&#160;Aug&#160;24&#160;03:44&#160;UTC | SHAKEN IDT America, Corp 363A | 15&#160;Sep&#160;24&#160;03:44&#160;UTC | true | [view](CERTS/7af16a7d69f7a9241997575bf27add9bbdc0693d495e0a9aff101c270a8fd36c/README.md) |
| 16&#160;Aug&#160;24&#160;04:29&#160;UTC | SHAKEN BareTelecom 864J | 15&#160;Sep&#160;24&#160;04:29&#160;UTC | true | [view](CERTS/3473662bf5d684e9d5638c959a73b7549df1beaadc35205663c6c73b3e7f1393/README.md) |
| 16&#160;Aug&#160;24&#160;04:39&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 15&#160;Sep&#160;24&#160;04:39&#160;UTC | true | [view](CERTS/496bac45da6aeb0921551b18a5ab3078c8805f1aa2e664c31dc453b0c59bcbed/README.md) |
| 16&#160;Aug&#160;24&#160;17:35&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 15&#160;Sep&#160;24&#160;17:35&#160;UTC | true | [view](CERTS/10c93b8cedf586462465de00802bc1b56655a410a8697901048ea367c444988f/README.md) |
| 17&#160;Aug&#160;24&#160;04:24&#160;UTC | SHAKEN BareTelecom 864J | 16&#160;Sep&#160;24&#160;04:24&#160;UTC | true | [view](CERTS/c08eec973be775f2c0ba5eaf487209f36196110d1532d3447cc4f602929b04d1/README.md) |
| 17&#160;Aug&#160;24&#160;16:09&#160;UTC | SHAKEN Mango Voice LLC 579K | 16&#160;Sep&#160;24&#160;16:09&#160;UTC | true | [view](CERTS/4be19d39259aae48015762e3b9de5eb29e988ebf614f418aa042e117c568f556/README.md) |
| 18&#160;Aug&#160;24&#160;17:25&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 17&#160;Sep&#160;24&#160;17:25&#160;UTC | true | [view](CERTS/c1979419286165ab884f3c8a7451da95357f0109f00d7741f916b1a56db5b801/README.md) |
| 19&#160;Aug&#160;24&#160;03:29&#160;UTC | SHAKEN IDT America, Corp 363A | 18&#160;Sep&#160;24&#160;03:29&#160;UTC | true | [view](CERTS/96bb638cd3d4953219d99b105119831c9ebe429e86ad2639a7479ceec9ed7aa0/README.md) |
| 19&#160;Aug&#160;24&#160;04:14&#160;UTC | SHAKEN BareTelecom 864J | 18&#160;Sep&#160;24&#160;04:14&#160;UTC | true | [view](CERTS/38f7bd60efed574eb91fb68005de1a0f5774e47c1be297cc16ad8dd449162d1c/README.md) |
| 19&#160;Aug&#160;24&#160;10:05&#160;UTC | SHAKEN Zella Technologies LLC 647K | 18&#160;Sep&#160;24&#160;10:05&#160;UTC | true | [view](CERTS/b28b968776af86d7c3b7e4638d3464694fafffe9e8cd480b507edd54497a96af/README.md) |
| 19&#160;Aug&#160;24&#160;17:20&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 18&#160;Sep&#160;24&#160;17:20&#160;UTC | true | [view](CERTS/019dd950fe6d5e7682ccd9333b5d6007eb1c8af7c89588dd1c9ef33f65820ab4/README.md) |
| 20&#160;Aug&#160;24&#160;03:24&#160;UTC | SHAKEN IDT America, Corp 363A | 19&#160;Sep&#160;24&#160;03:24&#160;UTC | true | [view](CERTS/8ac228b9b7f7b5d431fdc17b9f80b4b29e4cc67cb0c2bd918bc73b9e880506dd/README.md) |
| 20&#160;Aug&#160;24&#160;04:09&#160;UTC | SHAKEN BareTelecom 864J | 19&#160;Sep&#160;24&#160;04:09&#160;UTC | true | [view](CERTS/a8ea42ea98f5158ccd075e74d9abfb8eb65c02f3070d0c399f9fb55935b6e7f4/README.md) |

#### CA Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 21&#160;Aug&#160;20&#160;01:22&#160;UTC | SHAKEN Sansay Root CA US | 16&#160;Aug&#160;40&#160;01:22&#160;UTC | false | [view](CERTS/8356d251b255f4ac3fd108bb79be4c02dcea2d3b13d138892bdb3a70cbe6a343/README.md) |
| 02&#160;Sep&#160;22&#160;20:53&#160;UTC | SHAKEN Sansay Intermediate CA US WEST 1 | 31&#160;Aug&#160;29&#160;20:53&#160;UTC | false | [view](CERTS/4b1dfdba2b1e4bbffbf900a20f1f6f7befbef0008b963e4922a64469cb97d24b/README.md) |


Generated: 22 Aug 24 15:44 UTC