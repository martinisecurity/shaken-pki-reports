# STIR/SHAKEN CA Ecosystem Compliance

## Sansay

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 2768 certificates were included in the corpus being tested
- 19 certificates in the corpus were skipped because they are duplicates
- 2515 certificates in the corpus were skipped because they are expired
- 1 certificates in the corpus were skipped because they are not currently trusted
- 233 certificates being tested against the remaining rules
- 1.90 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 0.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 155 days is the average remaining validity for the certificates in the corpus
- 156 days is the average initial validity for the certificates in the corpus
- 156 certificates expire in the next 30 days
- 2.06 average number of unexpired certificates per OCN observed
- 113 unique OCNs observed in unexpired and valid certificate corpus

| Instances | Test | Source |
|-----------|------|--------|
| 146 | [e_atis_ext_crl_distribution_struct](ISSUES/e_atis_ext_crl_distribution_struct/README.md) | ATIS1000080 |
| 201 | [e_atis_subject_cn_spc](ISSUES/e_atis_subject_cn_spc/README.md) | ATIS1000080 |
| 8 | [e_atis_tn_auth_list_spc_format](ISSUES/e_atis_tn_auth_list_spc_format/README.md) | ATIS1000080 |
| 60 | [e_shaken_certificate_policies_id](ISSUES/e_shaken_certificate_policies_id/README.md) | US_SHAKEN_CP |
| 27 | [e_us_cp_subject_sn_shall](ISSUES/e_us_cp_subject_sn_shall/README.md) | US_SHAKEN_CP |

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
- 5217 days is the average remaining validity for the certificates in the corpus
- 4928 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days

No error, warning, or notice level issues were found

### Details

#### Leaf Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 06&#160;Oct&#160;23&#160;13:25&#160;UTC | SHAKEN Kloud 7 LLC 214K | 05&#160;Oct&#160;24&#160;13:25&#160;UTC | true | [view](CERTS/7105ca55bae14fd5fc5952beff1664e0b13f3b560dc07c777a8293f5d06ab752/README.md) |
| 06&#160;Oct&#160;23&#160;13:32&#160;UTC | SHAKEN Central Telephone Sales & Service, Inc. 294K | 05&#160;Oct&#160;24&#160;13:32&#160;UTC | true | [view](CERTS/949612b1d08f4185ad77017673b2c78b02b3a445786141e6fbea0c7f5511e684/README.md) |
| 12&#160;Oct&#160;23&#160;20:53&#160;UTC | SHAKEN Zito Media Voice 624G | 11&#160;Oct&#160;24&#160;20:53&#160;UTC | true | [view](CERTS/139f5ea60b5e891b041824e82c242eb6191006f1283e9bad6cd052bf821a0fc2/README.md) |
| 17&#160;Oct&#160;23&#160;22:55&#160;UTC | SHAKEN ConnectMeVoice 719J | 16&#160;Oct&#160;24&#160;22:55&#160;UTC | true | [view](CERTS/5a48a80de9440625305bf4a9af18bcf158e03286df27646af231d75e0cec315d/README.md) |
| 18&#160;Oct&#160;23&#160;22:36&#160;UTC | SHAKEN VOICE1 LLC 265K | 17&#160;Oct&#160;24&#160;22:36&#160;UTC | true | [view](CERTS/f40966bfe5a54cc2e71f5942314743411e143934deeeb8b87647eb1b5ac03aa2/README.md) |
| 19&#160;Oct&#160;23&#160;11:43&#160;UTC | SHAKEN Talk IT Pro 321K | 18&#160;Oct&#160;24&#160;11:43&#160;UTC | true | [view](CERTS/ac46b1e519197ae9ce860b54ca7b6a7150d4be7de7e581e7481003737cf28f68/README.md) |
| 23&#160;Oct&#160;23&#160;14:32&#160;UTC | SHAKEN Uvoice USA, LLC 328K | 22&#160;Oct&#160;24&#160;14:32&#160;UTC | true | [view](CERTS/63faaac90f1c3f037b3c850ad0f6203068aa3d62235646df63f6b07c9aa94d05/README.md) |
| 23&#160;Oct&#160;23&#160;15:54&#160;UTC | SHAKEN Comtalk Telecom 705K | 22&#160;Oct&#160;24&#160;15:54&#160;UTC | true | [view](CERTS/05f3a08f601f71f6b02f390050da6acb922f75b5f9d9178b637e96de52847674/README.md) |
| 25&#160;Oct&#160;23&#160;20:09&#160;UTC | SHAKEN Ytel Inc. 703J | 24&#160;Oct&#160;24&#160;20:09&#160;UTC | true | [view](CERTS/1abf91916a7c83ff53e58fdaf8fab1e6f6c232bc2aa2b1903c76c37d7a6984a6/README.md) |
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
| 14&#160;Nov&#160;23&#160;13:36&#160;UTC | SHAKEN KW Corporation, inc. 206k | 13&#160;Nov&#160;24&#160;13:36&#160;UTC | true | [view](CERTS/b6441dfbf4073cb729806ab15547ac613968b2235ed989b9ae02c05d16aa1807/README.md) |
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
| 03&#160;Jun&#160;24&#160;14:49&#160;UTC | SHAKEN 932K Alliance Phones | 03&#160;Jun&#160;25&#160;14:49&#160;UTC | true | [view](CERTS/277ded041e5ec5cda4c56eb0061ac4262dad2939b7ff40ce6b76946ed82facdb/README.md) |
| 09&#160;Jun&#160;24&#160;08:04&#160;UTC | SHAKEN 776K ActoinVox | 06&#160;Dec&#160;24&#160;08:04&#160;UTC | true | [view](CERTS/87b5e4ebe2f75f3bb1c8efe86e8a26d08435b5672243786f8a7f96d349694219/README.md) |
| 21&#160;Jun&#160;24&#160;20:19&#160;UTC | SHAKEN Connexum LLC 203K | 21&#160;Jun&#160;25&#160;20:19&#160;UTC | true | [view](CERTS/c405face9873a72f5ec4f1bc395a4124d7cfbf080a704060ebe5e9bf7a07d62d/README.md) |
| 29&#160;Jun&#160;24&#160;00:30&#160;UTC | SHAKEN Bek Communications Cooperative 1604 | 28&#160;Jun&#160;25&#160;00:30&#160;UTC | true | [view](CERTS/17a28f83e865e75a97e9a5a4b777355b93960ccf3b65a1d13e06693a0d03e399/README.md) |
| 01&#160;Jul&#160;24&#160;16:21&#160;UTC | SHAKEN Mercury Access Solutions 634K | 01&#160;Jul&#160;25&#160;16:21&#160;UTC | true | [view](CERTS/70cc27d2e747142c3446b6133993989be4e1ff3fe4d25b137def2745fc173d19/README.md) |
| 02&#160;Jul&#160;24&#160;15:32&#160;UTC | SHAKEN Bulk Solutions, LLC 644J | 31&#160;Jan&#160;25&#160;15:32&#160;UTC | true | [view](CERTS/9da93d10ef266f23087e3f23bd41ba68e75db57c81dd5099c76ae27074a16d47/README.md) |
| 10&#160;Jul&#160;24&#160;19:13&#160;UTC | SHAKEN 958K SaveVoip LLC | 10&#160;Jul&#160;25&#160;19:13&#160;UTC | true | [view](CERTS/e78654b5bea29fa67df948244d36aefdad852322d5c375149c5ce147430f2b95/README.md) |
| 12&#160;Jul&#160;24&#160;07:20&#160;UTC | SHAKEN 521K Voice SY LLC | 28&#160;Jan&#160;25&#160;07:20&#160;UTC | true | [view](CERTS/4543540ab57936221819dbdef685a8a4b09c87e6b72dd21fcbb346696de31a3b/README.md) |
| 19&#160;Jul&#160;24&#160;17:35&#160;UTC | SHAKEN IPBTel 535K | 19&#160;Jul&#160;25&#160;17:35&#160;UTC | true | [view](CERTS/c777a6bdc52c45c7f7f66f40b765c70e167fe34b8cd36fea8bc287b20cb7d80e/README.md) |
| 10&#160;Aug&#160;24&#160;19:01&#160;UTC | SHAKEN 688K Call Hub Inc. | 10&#160;Aug&#160;25&#160;19:01&#160;UTC | true | [view](CERTS/d8fe53673498502aa06d1deb531144cdeac93dba56e9c7094fef9fe192b5b14c/README.md) |
| 13&#160;Aug&#160;24&#160;19:57&#160;UTC | SHAKEN 986K 7G Network, Inc | 09&#160;Feb&#160;25&#160;19:57&#160;UTC | true | [view](CERTS/a72429200b39e65918e1e74562a1040baf0ec10307218baf94ec7b0d718a9ad4/README.md) |
| 19&#160;Aug&#160;24&#160;18:17&#160;UTC | SHAKEN 203K Connexum LLC | 19&#160;Aug&#160;25&#160;18:17&#160;UTC | true | [view](CERTS/901fcdbe4946f6c9f2f338dd125112cf748dec475f4052db54963ef4e3c55d1f/README.md) |
| 22&#160;Aug&#160;24&#160;16:18&#160;UTC | SHAKEN 688K Call Hub Inc. | 22&#160;Aug&#160;25&#160;16:18&#160;UTC | true | [view](CERTS/c9514ae8afe29c81cba005d0e97ddd87c0588caca1e39a20368d4f872619984d/README.md) |
| 22&#160;Aug&#160;24&#160;17:42&#160;UTC | SHAKEN Contactivity Corp. 711K | 22&#160;Aug&#160;25&#160;17:42&#160;UTC | true | [view](CERTS/6f93189ff71fccb600f52ac36ac52ec2d7ff916d7bf9e9d3a8eb8ae6d4c7cd31/README.md) |
| 04&#160;Sep&#160;24&#160;23:27&#160;UTC | SHAKEN IPSBS Managed Services LLC 828J | 04&#160;Oct&#160;24&#160;23:27&#160;UTC | true | [view](CERTS/86ec0db4adad2c26011b72e0f5d9fcc6404d4a29186d9215a644814ac09559b9/README.md) |
| 05&#160;Sep&#160;24&#160;02:49&#160;UTC | SHAKEN BareTelecom 864J | 05&#160;Oct&#160;24&#160;02:49&#160;UTC | true | [view](CERTS/b47ca0d6d20f0aa953543948bf7341b054363762422caf1567da5ff285c7b92c/README.md) |
| 05&#160;Sep&#160;24&#160;04:10&#160;UTC | SHAKEN IDT America, Corp 363A | 05&#160;Oct&#160;24&#160;04:10&#160;UTC | true | [view](CERTS/f4f8682edceead8b45963dc0566c27dffb045a97af6413d9f48c5da05baf326f/README.md) |
| 05&#160;Sep&#160;24&#160;04:45&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 05&#160;Oct&#160;24&#160;04:45&#160;UTC | true | [view](CERTS/beb9d43efccbf0c0208c96667a4326f1bbf19f82f1a898a8e3815cc36f517c12/README.md) |
| 05&#160;Sep&#160;24&#160;15:14&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 05&#160;Oct&#160;24&#160;15:14&#160;UTC | true | [view](CERTS/2e1fdbb5879169523c048bb81e0166c9c049f024830a17099f0bc3f52968f8a9/README.md) |
| 05&#160;Sep&#160;24&#160;15:58&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 05&#160;Oct&#160;24&#160;15:58&#160;UTC | true | [view](CERTS/8fc08adc6e4b930b17faa8887f64974a69942fa99cdd786a833025f31ca1ce79/README.md) |
| 06&#160;Sep&#160;24&#160;02:44&#160;UTC | SHAKEN BareTelecom 864J | 06&#160;Oct&#160;24&#160;02:44&#160;UTC | true | [view](CERTS/94902f31b8ea2faa0d87f69dd529a3f593d2e6e8ea06898f5911044d3ce4db18/README.md) |
| 06&#160;Sep&#160;24&#160;03:39&#160;UTC | SHAKEN 1stPoint Communications, LLC 463G | 06&#160;Oct&#160;24&#160;03:39&#160;UTC | true | [view](CERTS/d2fb90a73d3352321e0aca7f839e1c38f8be6ed5d14f85b35f4ad5e6cb47b1b2/README.md) |
| 06&#160;Sep&#160;24&#160;04:05&#160;UTC | SHAKEN IDT America, Corp 363A | 06&#160;Oct&#160;24&#160;04:05&#160;UTC | true | [view](CERTS/658fd8e79c2db9c79908fee48cf32a8129a647cc08bd5d7c288cbbeec82d2ec6/README.md) |
| 06&#160;Sep&#160;24&#160;04:40&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 06&#160;Oct&#160;24&#160;04:40&#160;UTC | true | [view](CERTS/3bdc9894412afdc43cbe44b3490fcf1928685089fe4f39a50bd6188fb7624839/README.md) |
| 06&#160;Sep&#160;24&#160;15:09&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 06&#160;Oct&#160;24&#160;15:09&#160;UTC | true | [view](CERTS/9b0e517839794248139d754aa2f5f24572f4dc71f600148f8db26962df12e3f0/README.md) |
| 06&#160;Sep&#160;24&#160;15:53&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 06&#160;Oct&#160;24&#160;15:53&#160;UTC | true | [view](CERTS/4c186ea5136b3c27728572bddeed4ee4afef0124e1d09c005441a315d3dc2103/README.md) |
| 07&#160;Sep&#160;24&#160;02:39&#160;UTC | SHAKEN BareTelecom 864J | 07&#160;Oct&#160;24&#160;02:39&#160;UTC | true | [view](CERTS/fae4403ba2431e1a2fb31b4136baee2649d84dcda50c4fb51869f365fbc676e5/README.md) |
| 07&#160;Sep&#160;24&#160;04:00&#160;UTC | SHAKEN IDT America, Corp 363A | 07&#160;Oct&#160;24&#160;04:00&#160;UTC | true | [view](CERTS/c3baa851487357740488ea1de2a860b6b1a9c5f91bb82c6c1d744e7e457be943/README.md) |
| 07&#160;Sep&#160;24&#160;04:35&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 07&#160;Oct&#160;24&#160;04:35&#160;UTC | true | [view](CERTS/b175019118d2d7d821103d49e48cc6bcb432a168302da263c387865bbd3401d7/README.md) |
| 08&#160;Sep&#160;24&#160;04:30&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 08&#160;Oct&#160;24&#160;04:30&#160;UTC | true | [view](CERTS/3eef531caccfa57c8c689c44eb789b86e17908cb8c96b2dd4fc190176d3bc54a/README.md) |
| 08&#160;Sep&#160;24&#160;14:59&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 08&#160;Oct&#160;24&#160;14:59&#160;UTC | true | [view](CERTS/be0105a24acfbbd738a7e9e56257002e9d197401ab0014a7ab6cf0c86f4ac8cb/README.md) |
| 08&#160;Sep&#160;24&#160;15:43&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 08&#160;Oct&#160;24&#160;15:43&#160;UTC | true | [view](CERTS/b318637b105103209b33e22c7ab46c1efdc130baefbfec59d6853b13b5af2173/README.md) |
| 08&#160;Sep&#160;24&#160;20:08&#160;UTC | SHAKEN Socket Telecom LLC 554a | 08&#160;Oct&#160;24&#160;20:08&#160;UTC | true | [view](CERTS/ad2c53f8babb8cedb1b62bd9ff5f1daae244e3ac70d6e59ce873493856803f07/README.md) |
| 08&#160;Sep&#160;24&#160;23:07&#160;UTC | SHAKEN IPSBS Managed Services LLC 828J | 08&#160;Oct&#160;24&#160;23:07&#160;UTC | true | [view](CERTS/72ba3ef768b00f9552ccce66f283928e547d356d39e92c7334e04fa9f9e21f52/README.md) |
| 09&#160;Sep&#160;24&#160;03:38&#160;UTC | SHAKEN BareTelecom 864J | 09&#160;Oct&#160;24&#160;03:38&#160;UTC | true | [view](CERTS/5425f58de3d1207613ac7751012eca0919a4a2264e3794c2463fddb37c834210/README.md) |
| 09&#160;Sep&#160;24&#160;03:50&#160;UTC | SHAKEN IDT America, Corp 363A | 09&#160;Oct&#160;24&#160;03:50&#160;UTC | true | [view](CERTS/ab0df089a20367ac11761caafadb8f82bb800c9c79da7266aa75c40e17483764/README.md) |
| 09&#160;Sep&#160;24&#160;12:36&#160;UTC | SHAKEN Lightspeed Voice 557F | 09&#160;Oct&#160;24&#160;12:36&#160;UTC | true | [view](CERTS/651fb24ddeec142984c05eb1e6c281458eac43d326369deb23cee7e8054bdca9/README.md) |
| 09&#160;Sep&#160;24&#160;14:54&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 09&#160;Oct&#160;24&#160;14:54&#160;UTC | true | [view](CERTS/5c5c3c8a38a19bbe433de5774bed34be6a7f21dbf1ad4e557cac2812a128429b/README.md) |
| 09&#160;Sep&#160;24&#160;15:38&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 09&#160;Oct&#160;24&#160;15:38&#160;UTC | true | [view](CERTS/74c98bca870154211caa83ce35cf17bab028c4d0796218bbc5af0ec85295c804/README.md) |
| 09&#160;Sep&#160;24&#160;20:03&#160;UTC | SHAKEN Socket Telecom LLC 554a | 09&#160;Oct&#160;24&#160;20:03&#160;UTC | true | [view](CERTS/2215dd5c98b1a8e4ba64083a71174a475e68f34324d1a39113c76a759663a77c/README.md) |
| 10&#160;Sep&#160;24&#160;03:33&#160;UTC | SHAKEN BareTelecom 864J | 10&#160;Oct&#160;24&#160;03:33&#160;UTC | true | [view](CERTS/582b755dd09060131537c698559ea19e6d2b5dfb68feb37443369e3079c2610c/README.md) |
| 10&#160;Sep&#160;24&#160;03:45&#160;UTC | SHAKEN IDT America, Corp 363A | 10&#160;Oct&#160;24&#160;03:45&#160;UTC | true | [view](CERTS/941a73a92e34043fb444e121ba0f7cf9c2ba0de9901fb9fd2391f76c63c19a1e/README.md) |
| 10&#160;Sep&#160;24&#160;06:50&#160;UTC | SHAKEN DLS Internet Services 815J | 10&#160;Oct&#160;24&#160;06:50&#160;UTC | true | [view](CERTS/6d78ebd8ebdca5591222f2775051f49ad8ffffce5756cf104881773b401e9348/README.md) |
| 10&#160;Sep&#160;24&#160;13:47&#160;UTC | SHAKEN 455K Cloud Connect LLC | 08&#160;Jan&#160;25&#160;13:47&#160;UTC | true | [view](CERTS/e74d06224975e1d5c4b1a8a889b81da3ed4a0f9ed46a4a0b1a9179faed5b301e/README.md) |
| 10&#160;Sep&#160;24&#160;13:49&#160;UTC | SHAKEN Consolidated Communications 5113 | 10&#160;Oct&#160;24&#160;13:49&#160;UTC | true | [view](CERTS/64633ff6e2430217d2e9aaaf7d9c6990e8fe02e73af4a982b1f1709cab884602/README.md) |
| 10&#160;Sep&#160;24&#160;13:50&#160;UTC | SHAKEN Touchtone 683A | 10&#160;Oct&#160;24&#160;13:50&#160;UTC | true | [view](CERTS/913d1d1d79719b2cc922953ab9ced00186db3684193a3c1775f50a54b0f4daa4/README.md) |
| 10&#160;Sep&#160;24&#160;13:51&#160;UTC | SHAKEN Apeiron Systems 012J | 10&#160;Oct&#160;24&#160;13:51&#160;UTC | true | [view](CERTS/730a38a60565abf2145f69ed82f181c4441695512d0c56128885fabfe9e01980/README.md) |
| 10&#160;Sep&#160;24&#160;13:52&#160;UTC | SHAKEN Fonative, Inc. 684J | 10&#160;Oct&#160;24&#160;13:52&#160;UTC | true | [view](CERTS/65477bc063e53a0c2dc2e061f3cbd5b0cbb3ad8be1d781415b2e967951320225/README.md) |
| 10&#160;Sep&#160;24&#160;13:52&#160;UTC | SHAKEN IPitomy 652J | 10&#160;Oct&#160;24&#160;13:52&#160;UTC | true | [view](CERTS/1e884365c25c8f8d33b748b2597ba49c0aadbf861a9df1ca4e0c20047d1628dd/README.md) |
| 10&#160;Sep&#160;24&#160;13:53&#160;UTC | SHAKEN Phone.com, Inc. 633J | 10&#160;Oct&#160;24&#160;13:53&#160;UTC | true | [view](CERTS/3b894fcf2c1e22b28d47081caf3d3bb9403c2229ad61c548f95b6134d6aba559/README.md) |
| 10&#160;Sep&#160;24&#160;13:54&#160;UTC | SHAKEN NETRIO LLC 020K | 10&#160;Oct&#160;24&#160;13:54&#160;UTC | true | [view](CERTS/1a8f467c190513e64df116810273db36b2f9019ae4676f5ecb15ecc46466eaf4/README.md) |
| 10&#160;Sep&#160;24&#160;13:55&#160;UTC | SHAKEN Noble Systems Communications LLC 187J | 10&#160;Oct&#160;24&#160;13:55&#160;UTC | true | [view](CERTS/7c99912713d42ac3e60ffdf9523f304fe8f7310252eddbb8e08f20d1be505a15/README.md) |
| 10&#160;Sep&#160;24&#160;13:57&#160;UTC | SHAKEN Airespring 996H | 10&#160;Oct&#160;24&#160;13:57&#160;UTC | true | [view](CERTS/02a3edec64305388cefa02c52c462319f354df1db342c8d880767bbe85ab3379/README.md) |
| 10&#160;Sep&#160;24&#160;14:00&#160;UTC | SHAKEN Nobelbiz, Inc. 596J | 10&#160;Oct&#160;24&#160;14:00&#160;UTC | true | [view](CERTS/45545db9bed50dddf20f4d6951489cf2225b2aa5a3632bed0be0b75bda4155ec/README.md) |
| 10&#160;Sep&#160;24&#160;14:02&#160;UTC | SHAKEN Momentum Telecom 9157 | 10&#160;Oct&#160;24&#160;14:02&#160;UTC | true | [view](CERTS/e2720b2916adb15afc2ecfea3ab910953e0f5c249d0697c562be6fa5a2a19325/README.md) |
| 10&#160;Sep&#160;24&#160;14:03&#160;UTC | SHAKEN Matrix 7379 | 10&#160;Oct&#160;24&#160;14:03&#160;UTC | true | [view](CERTS/06a9fc17f423ebf96925cb46a76abaacbdc1f3577db75812bd98c3923addbe39/README.md) |
| 10&#160;Sep&#160;24&#160;14:03&#160;UTC | SHAKEN Matrix 3058 | 10&#160;Oct&#160;24&#160;14:03&#160;UTC | true | [view](CERTS/8ffd97ba37d1011b87f90f23e4bdd624b8348419c44e239cabc005e5b21a1540/README.md) |
| 10&#160;Sep&#160;24&#160;14:04&#160;UTC | SHAKEN Matrix 9451 | 10&#160;Oct&#160;24&#160;14:04&#160;UTC | true | [view](CERTS/381f2acfb60de8658fb290b3206d67dfb3499db35c70a4e75a310a62e9581a6f/README.md) |
| 10&#160;Sep&#160;24&#160;14:05&#160;UTC | SHAKEN PNG Telecommunications Inc 3395 | 10&#160;Oct&#160;24&#160;14:05&#160;UTC | true | [view](CERTS/20efe75641ff7d3334f9d17781d00f692119797ad477b33fedaa825c3f9a4d8e/README.md) |
| 10&#160;Sep&#160;24&#160;14:07&#160;UTC | SHAKEN Magna5, LLC 3849 | 10&#160;Oct&#160;24&#160;14:07&#160;UTC | true | [view](CERTS/22887c7024bedb3ef7b29c91fa7e145dcc8ae9f86386eef3e1ee3535bb969606/README.md) |
| 10&#160;Sep&#160;24&#160;14:08&#160;UTC | SHAKEN Magna5, LLC 8249 | 10&#160;Oct&#160;24&#160;14:08&#160;UTC | true | [view](CERTS/b811c1e2ba15e05761dfef31b2f72bcd86b688c7fddd1af12597d96891ea51a5/README.md) |
| 10&#160;Sep&#160;24&#160;14:49&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 10&#160;Oct&#160;24&#160;14:49&#160;UTC | true | [view](CERTS/15250c1bee54bb5a8879cfb70a2afc455971e4ebff6d4050df245882bf39d388/README.md) |
| 10&#160;Sep&#160;24&#160;15:33&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 10&#160;Oct&#160;24&#160;15:33&#160;UTC | true | [view](CERTS/050f29f6c847007e160772bf209654e5c1116833317fdb5ee59e786f29a22bbf/README.md) |
| 11&#160;Sep&#160;24&#160;03:28&#160;UTC | SHAKEN BareTelecom 864J | 11&#160;Oct&#160;24&#160;03:28&#160;UTC | true | [view](CERTS/c7ed2be5460c774cdc2f70d9a2a373656f834b671502a1c3bbeb57274fc9ffc0/README.md) |
| 11&#160;Sep&#160;24&#160;03:29&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 11&#160;Oct&#160;24&#160;03:29&#160;UTC | true | [view](CERTS/ed04ac2ac1a1648da80caea9be4fa3621415dc79114f1908f2fef0add1d4589b/README.md) |
| 11&#160;Sep&#160;24&#160;03:40&#160;UTC | SHAKEN IDT America, Corp 363A | 11&#160;Oct&#160;24&#160;03:40&#160;UTC | true | [view](CERTS/26395a2ade862acdeb8de14e628d24f4025f0d19c82a8b5df4e6d06871a26297/README.md) |
| 11&#160;Sep&#160;24&#160;14:44&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 11&#160;Oct&#160;24&#160;14:44&#160;UTC | true | [view](CERTS/617f67e1e64fead5d81732fc6dd9c9be99a04383def443ddc45cb416e218ec51/README.md) |
| 12&#160;Sep&#160;24&#160;03:23&#160;UTC | SHAKEN BareTelecom 864J | 12&#160;Oct&#160;24&#160;03:23&#160;UTC | true | [view](CERTS/b9d04c9ddf9efdef7cb8fac76f89bac999c3f651264cff31b14a764dc520bc48/README.md) |
| 12&#160;Sep&#160;24&#160;03:24&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 12&#160;Oct&#160;24&#160;03:24&#160;UTC | true | [view](CERTS/9aaf022bda3e63891f265a00cfa9fd0efd43c1aca06583c45742cb0ed5a4ab83/README.md) |
| 12&#160;Sep&#160;24&#160;03:35&#160;UTC | SHAKEN IDT America, Corp 363A | 12&#160;Oct&#160;24&#160;03:35&#160;UTC | true | [view](CERTS/f756956fe809acbc440e1baa7c8cd57af7ed3705cc0566186a39a3790aaa2a60/README.md) |
| 12&#160;Sep&#160;24&#160;12:20&#160;UTC | SHAKEN Lightspeed Voice 557F | 12&#160;Oct&#160;24&#160;12:20&#160;UTC | true | [view](CERTS/0f61b5358f34c41c787f5cd0554f2d2554de713c5abbfa7abbcd8c4cc4b18dc2/README.md) |
| 12&#160;Sep&#160;24&#160;14:39&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 12&#160;Oct&#160;24&#160;14:39&#160;UTC | true | [view](CERTS/72b19b9248f4fa0c73dd3c48ba8260a9a00f1e73b0d68708b8c2e5d5ccd11d24/README.md) |
| 13&#160;Sep&#160;24&#160;03:18&#160;UTC | SHAKEN BareTelecom 864J | 13&#160;Oct&#160;24&#160;03:18&#160;UTC | true | [view](CERTS/07cb0fc25dbfa644ca29ec860a336bb449a2cac75b2e035429a4f134347b9d36/README.md) |
| 13&#160;Sep&#160;24&#160;03:19&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 13&#160;Oct&#160;24&#160;03:19&#160;UTC | true | [view](CERTS/220710f11dfa2d3be6821e79f069c161f366fd235796259e8d1d303dd6beb726/README.md) |
| 13&#160;Sep&#160;24&#160;03:30&#160;UTC | SHAKEN IDT America, Corp 363A | 13&#160;Oct&#160;24&#160;03:30&#160;UTC | true | [view](CERTS/898b242fdebccaad67fb2f134854f11b075f6c205744e83a2a25f582afe82a80/README.md) |
| 13&#160;Sep&#160;24&#160;13:54&#160;UTC | SHAKEN Mango Voice LLC 579K | 13&#160;Oct&#160;24&#160;13:54&#160;UTC | true | [view](CERTS/55c155ae238824836735772503878ecb35143bfe670fd6a0a80d96c02237a86d/README.md) |
| 13&#160;Sep&#160;24&#160;14:34&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 13&#160;Oct&#160;24&#160;14:34&#160;UTC | true | [view](CERTS/ba894b7f5b591813ab14fe176e388ae6a1b9ee8df2084142c302e79794f657eb/README.md) |
| 14&#160;Sep&#160;24&#160;03:13&#160;UTC | SHAKEN BareTelecom 864J | 14&#160;Oct&#160;24&#160;03:13&#160;UTC | true | [view](CERTS/47fadda3c759cc08a9db15103faa8f68b7f16c16d01ec8cb0ed33b3a4b86e4df/README.md) |
| 14&#160;Sep&#160;24&#160;03:14&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 14&#160;Oct&#160;24&#160;03:14&#160;UTC | true | [view](CERTS/f03892464fcb859ee389b6fd103a606829c7792abad25dfae305a596aec5a2c6/README.md) |
| 14&#160;Sep&#160;24&#160;11:38&#160;UTC | SHAKEN SouthPoint Communications 205k | 14&#160;Oct&#160;24&#160;11:38&#160;UTC | true | [view](CERTS/c426391e7c14ba62f2e2317e4e6b2cf6b89ac248631ba321d276aabb32c9dc68/README.md) |
| 15&#160;Sep&#160;24&#160;03:09&#160;UTC | SHAKEN BareTelecom 864J | 15&#160;Oct&#160;24&#160;03:09&#160;UTC | true | [view](CERTS/de63ae7a2cb94597eb39ce4aa6c5738614809b138bb224560b172299b67ee8ce/README.md) |
| 15&#160;Sep&#160;24&#160;04:28&#160;UTC | SHAKEN CIMA Telecom, Inc 313K | 15&#160;Oct&#160;24&#160;04:28&#160;UTC | true | [view](CERTS/2de05d397c1de5537bb4cb6575d5b94ad281bf50c87342384aef3e03cb11e3fb/README.md) |
| 15&#160;Sep&#160;24&#160;12:05&#160;UTC | SHAKEN Lightspeed Voice 557F | 15&#160;Oct&#160;24&#160;12:05&#160;UTC | true | [view](CERTS/b3fad654286bcac09078ba70e17c18513a0d5da072540400e9c5ce10ad6077f6/README.md) |
| 15&#160;Sep&#160;24&#160;14:24&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 15&#160;Oct&#160;24&#160;14:24&#160;UTC | true | [view](CERTS/f9cafe80e39c5b0a99d67cf6a31ced88e1d4aabfb1180720af0f3450eb70bc90/README.md) |
| 15&#160;Sep&#160;24&#160;22:32&#160;UTC | SHAKEN IPSBS Managed Services LLC 828J | 15&#160;Oct&#160;24&#160;22:32&#160;UTC | true | [view](CERTS/830e5b8dd24686e536b9bea98a36886f492f1a05a32d379c01118a6397abdc5f/README.md) |
| 16&#160;Sep&#160;24&#160;03:04&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 16&#160;Oct&#160;24&#160;03:04&#160;UTC | true | [view](CERTS/9d236b3d8b4fda8f6253d8096a185078def8139c2677ef7bf86856301fb706f2/README.md) |
| 16&#160;Sep&#160;24&#160;03:15&#160;UTC | SHAKEN IDT America, Corp 363A | 16&#160;Oct&#160;24&#160;03:15&#160;UTC | true | [view](CERTS/63c3cc5c6fcbe9d6f373f95efcb5fb63767c32ffc2127d4149d9dd9a4e2024df/README.md) |
| 16&#160;Sep&#160;24&#160;14:19&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 16&#160;Oct&#160;24&#160;14:19&#160;UTC | true | [view](CERTS/d208a0a87ce1893aeeda82b23c0d5cfcca0ded0f8c7cd17d2204d714f6b1f201/README.md) |
| 17&#160;Sep&#160;24&#160;02:59&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 17&#160;Oct&#160;24&#160;02:59&#160;UTC | true | [view](CERTS/617fffbd4fc0478c8822b0145da269adc075861fef1e4eed567c4b74916d0206/README.md) |
| 17&#160;Sep&#160;24&#160;03:10&#160;UTC | SHAKEN IDT America, Corp 363A | 17&#160;Oct&#160;24&#160;03:10&#160;UTC | true | [view](CERTS/a70b53708e740466ad6e64155acf8f486fa4a7013fed39565447794379c4c96a/README.md) |
| 17&#160;Sep&#160;24&#160;14:14&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 17&#160;Oct&#160;24&#160;14:14&#160;UTC | true | [view](CERTS/f7e8528c79da623f2ae5368e54a59e922da189ee06adc24a4fa4b7359cf90f12/README.md) |
| 18&#160;Sep&#160;24&#160;02:54&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 18&#160;Oct&#160;24&#160;02:54&#160;UTC | true | [view](CERTS/ac68eed7bc584bef0f43625d7a614a46bd0c2c2e6384bb51c6a3a3274e103bf7/README.md) |
| 18&#160;Sep&#160;24&#160;03:05&#160;UTC | SHAKEN IDT America, Corp 363A | 18&#160;Oct&#160;24&#160;03:05&#160;UTC | true | [view](CERTS/1daf4465d895a11fb1e67205f018ead21411ee37b144882735ff13c99f20d4d2/README.md) |
| 18&#160;Sep&#160;24&#160;14:09&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 18&#160;Oct&#160;24&#160;14:09&#160;UTC | true | [view](CERTS/b9339e6e1a054a89ce2442cd149d625bf1161a2c3769976225b97f6f0a2c9576/README.md) |
| 18&#160;Sep&#160;24&#160;22:17&#160;UTC | SHAKEN IPSBS Managed Services LLC 828J | 18&#160;Oct&#160;24&#160;22:17&#160;UTC | true | [view](CERTS/ddfbcc8df581958f651cfe2606e81c3bcaa63d8391389044003e8b6b29c8c4d2/README.md) |
| 19&#160;Sep&#160;24&#160;02:49&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 19&#160;Oct&#160;24&#160;02:49&#160;UTC | true | [view](CERTS/0d14c81e1e0e52daf3de2bc75fe9796415c71f65e423c663e72ccf85ba5e56b6/README.md) |
| 19&#160;Sep&#160;24&#160;03:00&#160;UTC | SHAKEN IDT America, Corp 363A | 19&#160;Oct&#160;24&#160;03:00&#160;UTC | true | [view](CERTS/27f4ceb7f9005dc609abf7d9846910356dab2f87dd846e0b91c18df720fd856e/README.md) |
| 19&#160;Sep&#160;24&#160;14:04&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 19&#160;Oct&#160;24&#160;14:04&#160;UTC | true | [view](CERTS/1727c9625fe46b1070c19d2fbf19f0ac148f74ce9c3812cb412dd4c27e858f80/README.md) |
| 19&#160;Sep&#160;24&#160;21:17&#160;UTC | SHAKEN 696J Telcentris Inc. dba Voxox | 19&#160;Oct&#160;24&#160;21:17&#160;UTC | true | [view](CERTS/1e0ddc4c4b14c59473e582fcdebca549406328dcd40f47102ca4f4bf66641e16/README.md) |
| 20&#160;Sep&#160;24&#160;02:55&#160;UTC | SHAKEN IDT America, Corp 363A | 20&#160;Oct&#160;24&#160;02:55&#160;UTC | true | [view](CERTS/b20c17abca67aa164675fcc54faeeb1ddc3a66f15256c5d59d2be1fbcc4cafe5/README.md) |
| 20&#160;Sep&#160;24&#160;13:19&#160;UTC | SHAKEN Mango Voice LLC 579K | 20&#160;Oct&#160;24&#160;13:19&#160;UTC | true | [view](CERTS/1a21f23117f7130cdff97d1a1d56b1f4a0662ea6d5b39385be7d062b3facea8f/README.md) |
| 22&#160;Sep&#160;24&#160;13:49&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 22&#160;Oct&#160;24&#160;13:49&#160;UTC | true | [view](CERTS/9a5e5fb75985c7c71bf03edfd21bf344ada97077e0d235f811b6dcaed2e3b283/README.md) |
| 22&#160;Sep&#160;24&#160;18:58&#160;UTC | SHAKEN Socket Telecom LLC 554a | 22&#160;Oct&#160;24&#160;18:58&#160;UTC | true | [view](CERTS/8c06057286de1eaea0bc2679df21968e8c57a52845265095dbc45e246d150fef/README.md) |
| 22&#160;Sep&#160;24&#160;21:57&#160;UTC | SHAKEN IPSBS Managed Services LLC 828J | 22&#160;Oct&#160;24&#160;21:57&#160;UTC | true | [view](CERTS/8c2c2ddf2c74d771259bc800356c365e53e0ea5a1f1ec71aa2160ea5298fe396/README.md) |
| 23&#160;Sep&#160;24&#160;02:29&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 23&#160;Oct&#160;24&#160;02:29&#160;UTC | true | [view](CERTS/64b8a3f1b378c478d072c840cd714a2854f9e6d5f80ce8ea97510b72b8f6bb40/README.md) |
| 23&#160;Sep&#160;24&#160;02:40&#160;UTC | SHAKEN IDT America, Corp 363A | 23&#160;Oct&#160;24&#160;02:40&#160;UTC | true | [view](CERTS/009a6d2eea93d6baec19c88400d738ffc271d281dc3ec2916c4ff3c472746c62/README.md) |
| 23&#160;Sep&#160;24&#160;06:30&#160;UTC | SHAKEN 249K Primo Dialler LLC | 21&#160;Jan&#160;25&#160;06:30&#160;UTC | true | [view](CERTS/e37c2d2abb454090dc41a8a9b8f15661bde2d9f635106de57ea81bbd74dd3c3b/README.md) |
| 24&#160;Sep&#160;24&#160;02:24&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 24&#160;Oct&#160;24&#160;02:24&#160;UTC | true | [view](CERTS/1a0d8ed50070fb6ad3ecc0a654a0f59cefa00ee9c15dd5ba2c74b45cedd280f2/README.md) |
| 24&#160;Sep&#160;24&#160;02:35&#160;UTC | SHAKEN IDT America, Corp 363A | 24&#160;Oct&#160;24&#160;02:35&#160;UTC | true | [view](CERTS/726796b97497bcfd95222a14501fd7d25be5fc80fdc93dfd136bc4cd664a273f/README.md) |
| 24&#160;Sep&#160;24&#160;13:39&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 24&#160;Oct&#160;24&#160;13:39&#160;UTC | true | [view](CERTS/e3499643e399f8cfc68c24be4a09a59998c536f4e84f35c633638d8c4caa3ec7/README.md) |
| 24&#160;Sep&#160;24&#160;17:01&#160;UTC | SHAKEN Sangoma 777G | 24&#160;Sep&#160;25&#160;17:01&#160;UTC | true | [view](CERTS/793541492ac2787c6052d4a9cfb63c7f6aed20d149d3d5e173aa531a5d1bb71b/README.md) |
| 24&#160;Sep&#160;24&#160;21:47&#160;UTC | SHAKEN IPSBS Managed Services LLC 828J | 24&#160;Oct&#160;24&#160;21:47&#160;UTC | true | [view](CERTS/02c3101853bb03f12838bf8acb5a5921a342d0c45b5e3b90acb1b8c6ebb12340/README.md) |
| 25&#160;Sep&#160;24&#160;02:19&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 25&#160;Oct&#160;24&#160;02:19&#160;UTC | true | [view](CERTS/e231b8aa49175fec507c50c4611460d45c495479090a3c59cba101a92ee9a8d0/README.md) |
| 25&#160;Sep&#160;24&#160;02:30&#160;UTC | SHAKEN IDT America, Corp 363A | 25&#160;Oct&#160;24&#160;02:30&#160;UTC | true | [view](CERTS/01cb84b293823051362a4cb0f0e332a1bffc4530e4601d31a46c1301de97db80/README.md) |
| 25&#160;Sep&#160;24&#160;11:15&#160;UTC | SHAKEN Lightspeed Voice 557F | 25&#160;Oct&#160;24&#160;11:15&#160;UTC | true | [view](CERTS/6e173418b967ddd5bbdb0d5571c9899ac95930f58332c3676a39858567847fd1/README.md) |
| 25&#160;Sep&#160;24&#160;13:34&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 25&#160;Oct&#160;24&#160;13:34&#160;UTC | true | [view](CERTS/a969ac34e43599365853f53fd2f239b081980d7bfebe8589eb802cd80d2e40a8/README.md) |
| 26&#160;Sep&#160;24&#160;02:14&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 26&#160;Oct&#160;24&#160;02:14&#160;UTC | true | [view](CERTS/82255919587070102dc0ecbe01da726845c298d924a586284104ee3248563025/README.md) |
| 26&#160;Sep&#160;24&#160;13:29&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 26&#160;Oct&#160;24&#160;13:29&#160;UTC | true | [view](CERTS/d2fc701abdb990d283767b81c20cf75b2747f34b24a144cfd258f68fc1ee0822/README.md) |
| 26&#160;Sep&#160;24&#160;18:38&#160;UTC | SHAKEN Socket Telecom LLC 554a | 26&#160;Oct&#160;24&#160;18:38&#160;UTC | true | [view](CERTS/412781039d3611efe205c0be9c99a43029f6a842812ddad54cb554afa1aa3225/README.md) |
| 27&#160;Sep&#160;24&#160;02:09&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 27&#160;Oct&#160;24&#160;02:09&#160;UTC | true | [view](CERTS/8ee5434e8d84c00c80927c0070ebdfcf7e40634d87f6419c048ff35fd08ff6ad/README.md) |
| 27&#160;Sep&#160;24&#160;03:03&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 27&#160;Oct&#160;24&#160;03:03&#160;UTC | true | [view](CERTS/13f83b0db80afe5287330ed0de56f3950c446a24fd6a3e42a305a1ed7859d82a/README.md) |
| 27&#160;Sep&#160;24&#160;13:24&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 27&#160;Oct&#160;24&#160;13:24&#160;UTC | true | [view](CERTS/fe13916f88bda07398f210dc52377859fc89481f1003066aed1ed304b69c8fc4/README.md) |
| 29&#160;Sep&#160;24&#160;01:59&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 29&#160;Oct&#160;24&#160;01:59&#160;UTC | true | [view](CERTS/8b25d7e08575c4286ba92a5e1424f2afe9edcf2ad7db8e219b24b735e5d2ba86/README.md) |
| 29&#160;Sep&#160;24&#160;02:53&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 29&#160;Oct&#160;24&#160;02:53&#160;UTC | true | [view](CERTS/b09dff2dfae9213e0acc6f9f24739d38cc6344c34af9b949f44602caee4a6c89/README.md) |
| 29&#160;Sep&#160;24&#160;15:49&#160;UTC | SHAKEN IDT America, Corp 363A | 29&#160;Oct&#160;24&#160;15:49&#160;UTC | true | [view](CERTS/850efa8c5caf562a4bd16c50d4f2278560ebfc7e66926e302b5825d844d77582/README.md) |
| 29&#160;Sep&#160;24&#160;18:23&#160;UTC | SHAKEN Socket Telecom LLC 554a | 29&#160;Oct&#160;24&#160;18:23&#160;UTC | true | [view](CERTS/0095b29170325b74ea27e6e0cb1a022c6059c8c67cd2edba86433c29c528026b/README.md) |
| 30&#160;Sep&#160;24&#160;01:54&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 30&#160;Oct&#160;24&#160;01:54&#160;UTC | true | [view](CERTS/e8735ccf2fcd27cec7194468bcbc410abb2c02d4ef5ed5dbefc0ccc90069dfbb/README.md) |
| 30&#160;Sep&#160;24&#160;10:50&#160;UTC | SHAKEN Lightspeed Voice 557F | 30&#160;Oct&#160;24&#160;10:50&#160;UTC | true | [view](CERTS/c6a4a858caeb3aab0b02f737ef4f7299f5c3d475753467fe2ecdc27aac1be862/README.md) |
| 30&#160;Sep&#160;24&#160;11:08&#160;UTC | SHAKEN Convoso 758J | 04&#160;Nov&#160;24&#160;11:08&#160;UTC | true | [view](CERTS/b235507a20c5a09a9c35e66bae5b3cff472db6142a907737802fac158356fa19/README.md) |
| 30&#160;Sep&#160;24&#160;12:32&#160;UTC | SHAKEN Mango Voice LLC 579K | 30&#160;Oct&#160;24&#160;12:32&#160;UTC | true | [view](CERTS/e565754c5b6a196ecd569af8895720e319216e9a50a21bd34d0af46b498c0e30/README.md) |
| 30&#160;Sep&#160;24&#160;13:09&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 30&#160;Oct&#160;24&#160;13:09&#160;UTC | true | [view](CERTS/d791eb9757bf2291eab31f82c104bc3cd3307a62a79c3c4c7699d3b3257bbe4c/README.md) |
| 30&#160;Sep&#160;24&#160;15:44&#160;UTC | SHAKEN IDT America, Corp 363A | 30&#160;Oct&#160;24&#160;15:44&#160;UTC | true | [view](CERTS/9d811914cacec78e5671559aed4d96529086330c680067162a7d08559495c003/README.md) |
| 01&#160;Oct&#160;24&#160;01:49&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 31&#160;Oct&#160;24&#160;01:49&#160;UTC | true | [view](CERTS/26e252529d9aa6f254e61993fcbb59b6be2c47118bf01120131f9b5d8902ba12/README.md) |
| 01&#160;Oct&#160;24&#160;02:43&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 31&#160;Oct&#160;24&#160;02:43&#160;UTC | true | [view](CERTS/a375d39b795b30c4b2806a167219df605159a44c88aaf26115f21e171dafbe56/README.md) |
| 01&#160;Oct&#160;24&#160;10:45&#160;UTC | SHAKEN Lightspeed Voice 557F | 31&#160;Oct&#160;24&#160;10:45&#160;UTC | true | [view](CERTS/4b99587f1b9d77da9981dd0945112c6d8d09acee53f42ef7dbf33c1460e393f0/README.md) |
| 01&#160;Oct&#160;24&#160;13:04&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 31&#160;Oct&#160;24&#160;13:04&#160;UTC | true | [view](CERTS/fbcc22a3cd903716e6a1a686ee7d15374117114f94660ba0599e1de36bf12e25/README.md) |
| 01&#160;Oct&#160;24&#160;13:48&#160;UTC | SHAKEN Touchtone 683A | 31&#160;Oct&#160;24&#160;13:48&#160;UTC | true | [view](CERTS/e8b604edd91a8c99bce8dcc06352294e1bbf2471c1fb6f2ee19eaab4e14baee3/README.md) |
| 01&#160;Oct&#160;24&#160;13:49&#160;UTC | SHAKEN Apeiron Systems 012J | 31&#160;Oct&#160;24&#160;13:49&#160;UTC | true | [view](CERTS/e3de70898a4eea5cccc2a00dac7b93804105ed78ec7660aa7e5a6b14787ab049/README.md) |
| 01&#160;Oct&#160;24&#160;13:50&#160;UTC | SHAKEN Fonative, Inc. 684J | 31&#160;Oct&#160;24&#160;13:50&#160;UTC | true | [view](CERTS/ff2923ee01bd6fc5373f1c6b187d1363b0d616120cb3bb1ec6703e15d5183144/README.md) |
| 01&#160;Oct&#160;24&#160;13:51&#160;UTC | SHAKEN Phone.com, Inc. 633J | 31&#160;Oct&#160;24&#160;13:51&#160;UTC | true | [view](CERTS/2f09099a4758f22858d76bc009eb9c3aae76a4ecb0694c8f3d697909994b654b/README.md) |
| 01&#160;Oct&#160;24&#160;13:52&#160;UTC | SHAKEN NETRIO LLC 020K | 31&#160;Oct&#160;24&#160;13:52&#160;UTC | true | [view](CERTS/bee959158adf225d0ed5ecd76c75b94e1b00532571f4ef199d49d4d18513d871/README.md) |
| 01&#160;Oct&#160;24&#160;13:54&#160;UTC | SHAKEN Airespring 996H | 31&#160;Oct&#160;24&#160;13:54&#160;UTC | true | [view](CERTS/7fb72a770b18f25883e3c0d9d757613fcf7e3a01c68dccd69d5c9301554f62b9/README.md) |
| 01&#160;Oct&#160;24&#160;13:56&#160;UTC | SHAKEN Nobelbiz, Inc. 596J | 31&#160;Oct&#160;24&#160;13:56&#160;UTC | true | [view](CERTS/8132bb2a29324c03db889fcefc9d02a5471e64f84ff2c55667d9f066462d72d1/README.md) |
| 01&#160;Oct&#160;24&#160;13:58&#160;UTC | SHAKEN Momentum Telecom 9157 | 31&#160;Oct&#160;24&#160;13:58&#160;UTC | true | [view](CERTS/17b855c1546c12c4901001f0d15a16979c1e2a91d0b2d9a19450adfa8bc9296b/README.md) |
| 01&#160;Oct&#160;24&#160;14:03&#160;UTC | SHAKEN Magna5, LLC 3849 | 31&#160;Oct&#160;24&#160;14:03&#160;UTC | true | [view](CERTS/fc2c0e841ac5bf2dcdafc6a74aca50645650e2b118cade6af3375dd6efa4021b/README.md) |
| 01&#160;Oct&#160;24&#160;14:04&#160;UTC | SHAKEN Magna5, LLC 8249 | 31&#160;Oct&#160;24&#160;14:04&#160;UTC | true | [view](CERTS/1345de3a8a2ca9ad3368f984de21adce909330fdb6e2dc951c1a440f404dd81d/README.md) |
| 01&#160;Oct&#160;24&#160;15:06&#160;UTC | SHAKEN Kloud 7 LLC 214K | 01&#160;Oct&#160;25&#160;15:06&#160;UTC | true | [view](CERTS/67fe5061b5b9244cdb233b8efd82fffde4a42fc661cce8860b6717318181ca80/README.md) |
| 01&#160;Oct&#160;24&#160;15:39&#160;UTC | SHAKEN IDT America, Corp 363A | 31&#160;Oct&#160;24&#160;15:39&#160;UTC | true | [view](CERTS/46afbd3cbe62cadcd8e0b2e358364633031afdef1107e2bbb1db6d7e73c18e74/README.md) |
| 02&#160;Oct&#160;24&#160;01:03&#160;UTC | SHAKEN Townes Telecommunications 0335 | 01&#160;Nov&#160;24&#160;01:03&#160;UTC | true | [view](CERTS/33e15865325441feb4fe0cabae77222f0107e00b64de9b7b3d38b5554f4a248a/README.md) |
| 02&#160;Oct&#160;24&#160;01:44&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 01&#160;Nov&#160;24&#160;01:44&#160;UTC | true | [view](CERTS/f432d70f353dd43d8886f90664b468fadff555696998fa5da83a172b62876814/README.md) |
| 02&#160;Oct&#160;24&#160;12:22&#160;UTC | SHAKEN Mango Voice LLC 579K | 01&#160;Nov&#160;24&#160;12:22&#160;UTC | true | [view](CERTS/6dabaea104b986a2d1e6159391f3c606811d585af7cbd6f2a07def6d96d35f43/README.md) |
| 02&#160;Oct&#160;24&#160;12:59&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 01&#160;Nov&#160;24&#160;12:59&#160;UTC | true | [view](CERTS/16241228497cb74debe19464f431fd839a88abfbf8b2c974af4c8c81eb96e5a0/README.md) |
| 02&#160;Oct&#160;24&#160;15:34&#160;UTC | SHAKEN IDT America, Corp 363A | 01&#160;Nov&#160;24&#160;15:34&#160;UTC | true | [view](CERTS/00ad775b20353d97d9c808158c91c77a01a3529778cdc3dc1394d35a081ef2ad/README.md) |
| 03&#160;Oct&#160;24&#160;01:39&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 02&#160;Nov&#160;24&#160;01:39&#160;UTC | true | [view](CERTS/075b169220d147430a9eccf048ae9c61dfda08b60b3f9c82ce5c0d539aa6d689/README.md) |
| 03&#160;Oct&#160;24&#160;10:35&#160;UTC | SHAKEN Lightspeed Voice 557F | 02&#160;Nov&#160;24&#160;10:35&#160;UTC | true | [view](CERTS/e394e7ed3c5c36107e36ca6c098ffdf2279ebd277b5c434e0f567c7f8d83a57c/README.md) |
| 03&#160;Oct&#160;24&#160;12:54&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 02&#160;Nov&#160;24&#160;12:54&#160;UTC | true | [view](CERTS/89c28ee75f98cfa9876c0f4a1c1d42a06324eecae7d7a1f3a28106c647826abc/README.md) |
| 03&#160;Oct&#160;24&#160;15:29&#160;UTC | SHAKEN IDT America, Corp 363A | 02&#160;Nov&#160;24&#160;15:29&#160;UTC | true | [view](CERTS/e917cf4b97ed045860f4777330983956a9b230286a81f06d8bdd4c56b1a786a2/README.md) |
| 03&#160;Oct&#160;24&#160;18:06&#160;UTC | SHAKEN Zray Technologies Corporation 862J | 02&#160;Nov&#160;24&#160;18:06&#160;UTC | true | [view](CERTS/930b00d582cc83dffb2bdb45b4d9c5f513cb4db7e50947d30289d5f98a68d5f4/README.md) |
| 04&#160;Oct&#160;24&#160;01:34&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 03&#160;Nov&#160;24&#160;01:34&#160;UTC | true | [view](CERTS/8e9aaa5c0368d9faa5d1c2d91939a5ce7bd9ce1fbf06e69c24b041a8b9ac0628/README.md) |
| 04&#160;Oct&#160;24&#160;12:49&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 03&#160;Nov&#160;24&#160;12:49&#160;UTC | true | [view](CERTS/cc39b9af163b9ac535b980a6f90585bff0fa38d55a2dbc5736317455ab960feb/README.md) |

#### CA Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 21&#160;Aug&#160;20&#160;01:22&#160;UTC | SHAKEN Sansay Root CA US | 16&#160;Aug&#160;40&#160;01:22&#160;UTC | false | [view](CERTS/8356d251b255f4ac3fd108bb79be4c02dcea2d3b13d138892bdb3a70cbe6a343/README.md) |
| 02&#160;Sep&#160;22&#160;20:53&#160;UTC | SHAKEN Sansay Intermediate CA US WEST 1 | 31&#160;Aug&#160;29&#160;20:53&#160;UTC | false | [view](CERTS/4b1dfdba2b1e4bbffbf900a20f1f6f7befbef0008b963e4922a64469cb97d24b/README.md) |


Generated: 04 Oct 24 16:29 UTC