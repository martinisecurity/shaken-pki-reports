# STIR/SHAKEN CA Ecosystem Compliance

## Sansay

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 1269 certificates were included in the corpus being tested
- 15 certificates in the corpus were skipped because they are duplicates
- 1060 certificates in the corpus were skipped because they are expired
- 1 certificates in the corpus were skipped because they are not currently trusted
- 193 certificates being tested against the remaining rules
- 4.87 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 100.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 126 days is the average remaining validity for the certificates in the corpus
- 127 days is the average initial validity for the certificates in the corpus
- 134 certificates expire in the next 30 days
- 2.14 average number of unexpired certificates per OCN observed
- 90 unique OCNs observed in unexpired and valid certificate corpus

| Instances | Test | Source |
|-----------|------|--------|
| 58 | [e_atis_certificate_policies](ISSUES/e_atis_certificate_policies/README.md) | ATIS1000080 |
| 193 | [e_atis_subject_cn](ISSUES/e_atis_subject_cn/README.md) | ATIS1000080 |
| 193 | [e_us_cp_ambiguous_identifier](ISSUES/e_us_cp_ambiguous_identifier/README.md) | US_SHAKEN_CP |
| 193 | [e_us_cp_subject_sn](ISSUES/e_us_cp_subject_sn/README.md) | US_SHAKEN_CP |
| 110 | [w_cp_1_3_subject_email](ISSUES/w_cp_1_3_subject_email/README.md) | US_SHAKEN_CP |
| 193 | [w_shaken_subject_rdn_unknown](ISSUES/w_shaken_subject_rdn_unknown/README.md) | SHAKEN_PKI_BEST_PRACTICES |

#### CA Certificates

- 3 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 0 certificates in the corpus were skipped because they are expired
- 1 certificates in the corpus were skipped because they are not currently trusted
- 2 certificates being tested against the remaining rules
- 1.00 issues on average found in unexpired, trusted, and non-compliant certificates
- 0.00% of certificates contain one or more Error level issue
- 100.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 50.00% of certificates are too old to be assessed against currently enforced expectations
- 5353 days is the average remaining validity for the certificates in the corpus
- 4928 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days

| Instances | Test | Source |
|-----------|------|--------|
| 2 | [w_shaken_ca_subject_rdn_unknown](ISSUES/w_shaken_ca_subject_rdn_unknown/README.md) | SHAKEN_PKI_BEST_PRACTICES |

### Details

#### Leaf Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 12&#160;Sep&#160;22&#160;19:37&#160;UTC | SHAKEN Bulk Solutions, LLC 644J | 12&#160;Sep&#160;23&#160;19:37&#160;UTC | true | [view](CERTS/2b980444a4603ddf16248bee9dbdce112f593d4d5324443e641624a827af0cb2/README.md) |
| 23&#160;Sep&#160;22&#160;01:10&#160;UTC | SHAKEN Star2Star Communications, LLC 590J | 23&#160;Sep&#160;23&#160;01:10&#160;UTC | true | [view](CERTS/b6c27ce63b22687fcd2f9f64ee9067dd3c19a4eb223f1aef3934f7ba95c54ba6/README.md) |
| 11&#160;Oct&#160;22&#160;16:48&#160;UTC | SHAKEN TeleVoIPs 138K | 11&#160;Oct&#160;23&#160;16:48&#160;UTC | true | [view](CERTS/c41b66127049dbae159f8d68ac714616b9e99640c407bcdc749f3d49037db487/README.md) |
| 11&#160;Oct&#160;22&#160;17:08&#160;UTC | SHAKEN ALD Telecom 780J | 11&#160;Oct&#160;23&#160;17:08&#160;UTC | true | [view](CERTS/53a14081c994555770bb8c5f3d160f89cf427258c9598d569c388a74bde6ea8f/README.md) |
| 11&#160;Oct&#160;22&#160;17:17&#160;UTC | SHAKEN Current Calls, LLC 746J | 11&#160;Oct&#160;23&#160;17:17&#160;UTC | true | [view](CERTS/52d6a93a1b72d2f2980699e759068dd9dbc8314c953e03613f18d9da1dcf156d/README.md) |
| 11&#160;Oct&#160;22&#160;17:20&#160;UTC | SHAKEN Carrier One Inc. 705J | 11&#160;Oct&#160;23&#160;17:20&#160;UTC | true | [view](CERTS/a7447339990a198aac3d84ed38d80706e16b7aac171e6d6bd1b28275fe7c337e/README.md) |
| 11&#160;Oct&#160;22&#160;17:21&#160;UTC | SHAKEN Asia Pacific Network 988J | 11&#160;Oct&#160;23&#160;17:21&#160;UTC | true | [view](CERTS/0b191ba4d02eaa4b595b67a4d3e6f35a6d6c184e5b7e464d471cb904ea2d0638/README.md) |
| 11&#160;Oct&#160;22&#160;17:22&#160;UTC | SHAKEN Vumber LLC 225K | 11&#160;Oct&#160;23&#160;17:22&#160;UTC | true | [view](CERTS/68075fd5ebbd21a4ecc74ecd70c85bb47ebfa522353477429221c911e84d0256/README.md) |
| 11&#160;Oct&#160;22&#160;17:22&#160;UTC | SHAKEN OneStream Networks, LLC 630J | 11&#160;Oct&#160;23&#160;17:22&#160;UTC | true | [view](CERTS/f18d0d387f4abfadaa336e2ff00c0f6b0509898b7d2d54feb99e1e0fb2042d3a/README.md) |
| 11&#160;Oct&#160;22&#160;17:23&#160;UTC | SHAKEN Ringfree Communications Inc 317K | 11&#160;Oct&#160;23&#160;17:23&#160;UTC | true | [view](CERTS/cc75f739ba9e082e5324936f9c5c1df2d896cb259ed0dd51065b937a0fce25aa/README.md) |
| 11&#160;Oct&#160;22&#160;17:24&#160;UTC | SHAKEN Xchange Telecom LLC 325B | 11&#160;Oct&#160;23&#160;17:24&#160;UTC | true | [view](CERTS/6bab691174d8e7b237a7fe1b00556840e2a5c28a1839f8e345dd9ba721ba23bb/README.md) |
| 11&#160;Oct&#160;22&#160;17:24&#160;UTC | SHAKEN Sangoma 777G | 11&#160;Oct&#160;23&#160;17:24&#160;UTC | true | [view](CERTS/53d28ac1fa5253468c11b9e3baaa6ad5481e83a7ea2ee6d715594dc6d4561ad4/README.md) |
| 12&#160;Oct&#160;22&#160;12:50&#160;UTC | SHAKEN Lightspeed Voice 557F | 12&#160;Oct&#160;23&#160;12:50&#160;UTC | true | [view](CERTS/ab19df868054cb3392aa295bff737bf919f8dc55c64a91247621375bad7fb7c0/README.md) |
| 13&#160;Oct&#160;22&#160;20:27&#160;UTC | SHAKEN ConnectMeVoice 719J | 13&#160;Oct&#160;23&#160;20:27&#160;UTC | true | [view](CERTS/a5edeeacfcec8ad6584f5a0b505978c4b72907a2e3a6540bb01350397f86814e/README.md) |
| 24&#160;Oct&#160;22&#160;20:23&#160;UTC | SHAKEN Arbeit 816J | 24&#160;Oct&#160;23&#160;20:23&#160;UTC | true | [view](CERTS/377e182a223e6cc8d7e9ce697e7a3e829b1c6b16c299c26f6d1f1e33aa29524b/README.md) |
| 24&#160;Oct&#160;22&#160;21:11&#160;UTC | SHAKEN Ytel Inc. 703J | 24&#160;Oct&#160;23&#160;21:11&#160;UTC | true | [view](CERTS/3d6a7a2ff23b90fba1674f600a108b8a11a110f8bb1723df86627001f7367d8d/README.md) |
| 25&#160;Oct&#160;22&#160;20:17&#160;UTC | SHAKEN Talk IT Pro 321K | 25&#160;Oct&#160;23&#160;20:17&#160;UTC | true | [view](CERTS/e45dada701a589e681d12207ebf16985abf6d62cf429b6e03bdcf8c0f97c3bf2/README.md) |
| 25&#160;Oct&#160;22&#160;21:11&#160;UTC | SHAKEN Intelegrated, LLC 325K | 25&#160;Oct&#160;23&#160;21:11&#160;UTC | true | [view](CERTS/974ca998bb71c71a66acf7c677484b70cb2620214bf4966b7488ac62755e3655/README.md) |
| 26&#160;Oct&#160;22&#160;19:34&#160;UTC | SHAKEN Vinculum Communications, Inc 787J | 26&#160;Oct&#160;23&#160;19:34&#160;UTC | true | [view](CERTS/22936e87ea3c45af88f1e501b88c6c6db3c271bd6ef73ab33c5d68198f9d4d66/README.md) |
| 26&#160;Oct&#160;22&#160;19:40&#160;UTC | SHAKEN California Telecom 319K | 26&#160;Oct&#160;23&#160;19:40&#160;UTC | true | [view](CERTS/35608b0f6472e87f044ac265b9c59e6c735e7f8102cfef09a46efbacad21e265/README.md) |
| 26&#160;Oct&#160;22&#160;19:41&#160;UTC | SHAKEN DLS Internet Services 815J | 26&#160;Oct&#160;23&#160;19:41&#160;UTC | true | [view](CERTS/7cd8319bedd12f040e8bd7b522d981aabcd24dc5aef74614a67fb6fdc9b9823b/README.md) |
| 26&#160;Oct&#160;22&#160;19:43&#160;UTC | SHAKEN Systemverse, LLC. 194K | 26&#160;Oct&#160;23&#160;19:43&#160;UTC | true | [view](CERTS/edbe74f809b9e0e1ebea447df8bdbfb272144f9c8c18df81e397a374df61c4cd/README.md) |
| 26&#160;Oct&#160;22&#160;19:44&#160;UTC | SHAKEN Rayfield Communications, Inc. 006K | 26&#160;Oct&#160;23&#160;19:44&#160;UTC | true | [view](CERTS/5032969f5932ac46a17b86c38dc72d666be454d1c3f11918edfa8385d9fc65e6/README.md) |
| 26&#160;Oct&#160;22&#160;20:04&#160;UTC | SHAKEN Kloud 7 LLC 214K | 26&#160;Oct&#160;23&#160;20:04&#160;UTC | true | [view](CERTS/f735d149cb8a392345d3e8a21bd578a46e96b1d4e0926068fd81476208acacef/README.md) |
| 27&#160;Oct&#160;22&#160;20:11&#160;UTC | SHAKEN Mitel Cloud Services, Inc. 670J | 27&#160;Oct&#160;23&#160;20:11&#160;UTC | true | [view](CERTS/e45c92abcfe2fe6d0863200900b66e835aa98712f974efe3837e34d787f2ad5e/README.md) |
| 07&#160;Nov&#160;22&#160;21:53&#160;UTC | SHAKEN Starlinq PBX Inc. 267K | 07&#160;Nov&#160;23&#160;21:53&#160;UTC | true | [view](CERTS/556a31c75cded397a5564c6f3bce1bf50e44e68e13d8ba757b5e6c20ad997fdb/README.md) |
| 21&#160;Nov&#160;22&#160;21:15&#160;UTC | SHAKEN Star2Star Communications, LLC 590J | 21&#160;Nov&#160;23&#160;21:15&#160;UTC | true | [view](CERTS/9bc9dde8921387803d93036c7d2f8085af32b028fca8f17336d2e22ab51fd278/README.md) |
| 29&#160;Nov&#160;22&#160;22:04&#160;UTC | SHAKEN MagicJack 324E | 29&#160;Nov&#160;23&#160;22:04&#160;UTC | true | [view](CERTS/75b4b7b400b1252e48faa1d93f6a94f7bd4a6383c88ddf6baa167b85d9ac4ee8/README.md) |
| 05&#160;Dec&#160;22&#160;22:28&#160;UTC | SHAKEN Bulk Solutions, LLC 644J | 05&#160;Dec&#160;23&#160;22:28&#160;UTC | true | [view](CERTS/3cf0aa2a24845e3fe6b27605e223e8e0c73d6bd4f73279b8a1e5e16fd2feeb80/README.md) |
| 10&#160;Dec&#160;22&#160;02:11&#160;UTC | SHAKEN Drop Inc 258K | 10&#160;Dec&#160;23&#160;02:11&#160;UTC | true | [view](CERTS/fc457741017b89b9126882710d8fb44883d7603f79cec0a1989eaa2b08034ee5/README.md) |
| 19&#160;Jan&#160;23&#160;22:50&#160;UTC | SHAKEN Technology Innovation Lab 599J | 19&#160;Jan&#160;24&#160;22:50&#160;UTC | true | [view](CERTS/12acafcf01348d278955bb9276e7a4d22a65ccdc61a59d08100177711f21b430/README.md) |
| 23&#160;Jan&#160;23&#160;21:55&#160;UTC | SHAKEN Swift Telco LLC 452K | 23&#160;Jan&#160;24&#160;21:55&#160;UTC | true | [view](CERTS/613861829aae7927f05dbd5a7b9f28ae8c4f995bb8ed115f95fc4be6644ccde1/README.md) |
| 26&#160;Jan&#160;23&#160;14:26&#160;UTC | SHAKEN Bulk Solutions, LLC 644J | 11&#160;Jan&#160;24&#160;14:26&#160;UTC | true | [view](CERTS/7dcc6cd32bf3c4e6e2901468097a88ad42a72ce42df34edce87c84dbce3691d2/README.md) |
| 08&#160;Feb&#160;23&#160;19:07&#160;UTC | SHAKEN ConvergeTel LLC 388K | 08&#160;Feb&#160;24&#160;19:07&#160;UTC | true | [view](CERTS/80706b79565f875515eb32f8cf113093a2658148ece8440e76199e4004254c31/README.md) |
| 14&#160;Feb&#160;23&#160;17:12&#160;UTC | SHAKEN Ytel Inc. 703J | 14&#160;Feb&#160;24&#160;17:12&#160;UTC | true | [view](CERTS/14c9bef113cfebe60611b0c56c430518ff8d42e8b98dd7e4653bd9cf619d5641/README.md) |
| 01&#160;Mar&#160;23&#160;14:14&#160;UTC | SHAKEN Voneto 485K | 29&#160;Feb&#160;24&#160;14:14&#160;UTC | true | [view](CERTS/8ca062af72eeb0b2eec4aaa1e6f3eb5ee3f01eccbd2895f282bbf4b70015a5d5/README.md) |
| 14&#160;Mar&#160;23&#160;21:39&#160;UTC | SHAKEN Cherry Voice 506K | 13&#160;Mar&#160;24&#160;21:39&#160;UTC | true | [view](CERTS/3756d1fa94ab1105ceb4bc6e3dcc371871b6edd45ac350df018a35f2f93f1967/README.md) |
| 19&#160;Mar&#160;23&#160;00:31&#160;UTC | SHAKEN Every1 Telecom 486K | 18&#160;Mar&#160;24&#160;00:31&#160;UTC | true | [view](CERTS/5cdfbb1a416083096dfef10c75a2b26a08d8fd5593d8ea9ceae0d70d878a97d1/README.md) |
| 27&#160;Mar&#160;23&#160;21:48&#160;UTC | SHAKEN Global Net Holdings Inc 306K | 21&#160;Jan&#160;24&#160;21:48&#160;UTC | true | [view](CERTS/f8d913af64a11c0718c78426b747dba4ed4ccb19239db573626a46f29b671825/README.md) |
| 28&#160;Mar&#160;23&#160;15:31&#160;UTC | SHAKEN VoIP Innovations 597F | 27&#160;Mar&#160;24&#160;15:31&#160;UTC | true | [view](CERTS/105a7683d4b5fac2ea7c2383e95250715bc0460d2cfdbea0d220201f44ea5d0c/README.md) |
| 03&#160;Apr&#160;23&#160;21:02&#160;UTC | SHAKEN WWT INC dba VoIP Networks 053K | 02&#160;Apr&#160;24&#160;21:02&#160;UTC | true | [view](CERTS/6229273e97413ee4c01a1810885db9b48d7e1bf7fe8aa4e7be22076effe2cc8b/README.md) |
| 04&#160;Apr&#160;23&#160;16:38&#160;UTC | SHAKEN Inventive Labs Corp 649J | 01&#160;Oct&#160;23&#160;16:38&#160;UTC | true | [view](CERTS/61bd59e51cd88d19ca695a362569930dee7725fefa39299b507bc3269fefbd54/README.md) |
| 05&#160;Apr&#160;23&#160;16:27&#160;UTC | SHAKEN Swift Telco LLC 452K | 04&#160;Apr&#160;24&#160;16:27&#160;UTC | true | [view](CERTS/947b46067a639b79ff82ab3f48c453e4af7cc6d6036f6d66a742cc935bc8a35e/README.md) |
| 19&#160;Apr&#160;23&#160;06:07&#160;UTC | SHAKEN Telxio Networks 492K | 18&#160;Apr&#160;24&#160;06:07&#160;UTC | true | [view](CERTS/010c6a74330323c20ceb343b1de3a1e3248b4a3926c9ad2ed53f02b02399d241/README.md) |
| 19&#160;Apr&#160;23&#160;06:39&#160;UTC | SHAKEN Telxio Networks 492K | 18&#160;Apr&#160;24&#160;06:39&#160;UTC | true | [view](CERTS/1baf0e5fedd50fb55ff4e07366c5eb4f8d849b760739ffb8a0df4eb4828d7944/README.md) |
| 25&#160;Apr&#160;23&#160;16:43&#160;UTC | SHAKEN GIP Technology 434K | 24&#160;Apr&#160;24&#160;16:43&#160;UTC | true | [view](CERTS/bc5cd573f3ab46daf12994739622514f37ac8cd3275ff8d493595fee747e8a0b/README.md) |
| 18&#160;May&#160;23&#160;00:00&#160;UTC | SHAKEN Primo Dialler LLC 249K | 14&#160;Nov&#160;23&#160;00:00&#160;UTC | true | [view](CERTS/a80051aad3dce4fd90a1e1f6758bff57994a1b1abc2e176244344c3a99c3e071/README.md) |
| 18&#160;May&#160;23&#160;01:12&#160;UTC | SHAKEN Cloud Connect LLC 455K | 17&#160;May&#160;24&#160;01:12&#160;UTC | true | [view](CERTS/0aa593ccacc13e85c2ec381274e47b597989c2a57173e248e25a91bc306c5f2c/README.md) |
| 24&#160;May&#160;23&#160;14:49&#160;UTC | SHAKEN Nextiva, Inc 914H | 23&#160;May&#160;24&#160;14:49&#160;UTC | true | [view](CERTS/ef77b45b37fa412f05b53660a4b60f37962a9be4587ec39211e94289a0087c20/README.md) |
| 01&#160;Jun&#160;23&#160;19:33&#160;UTC | SHAKEN Medtel Communications 994J | 31&#160;May&#160;24&#160;19:33&#160;UTC | true | [view](CERTS/e4b9231f5ad017174e74400e163effa66767d53af72cf608ba3046a0e639a813/README.md) |
| 22&#160;Jun&#160;23&#160;14:29&#160;UTC | SHAKEN Peachnet LLC 616K | 21&#160;Jun&#160;24&#160;14:29&#160;UTC | true | [view](CERTS/ab9d350cec16b5212bc010e5637396f92c2089704a3f985541d175a1f4b2a02a/README.md) |
| 22&#160;Jun&#160;23&#160;18:42&#160;UTC | SHAKEN FastCast Networks 318K | 21&#160;Jun&#160;24&#160;18:42&#160;UTC | true | [view](CERTS/13723330e372f36b27ca99e1ae23cae3ff5745d92ca6174c3987aef15dc26351/README.md) |
| 22&#160;Jun&#160;23&#160;19:59&#160;UTC | SHAKEN Connexum LLC 203K | 21&#160;Jun&#160;24&#160;19:59&#160;UTC | true | [view](CERTS/d1f5789f3d44bf7545537a3539f0f5dbf43de9cca281266af09943e5635fecfd/README.md) |
| 27&#160;Jun&#160;23&#160;11:43&#160;UTC | SHAKEN Bek Communications Cooperative 1604 | 26&#160;Jun&#160;24&#160;11:43&#160;UTC | true | [view](CERTS/b2f1cb0fa492baaa28a3591bc818d9946da92558364d772b2ddce7ec0028b4d0/README.md) |
| 27&#160;Jun&#160;23&#160;17:58&#160;UTC | SHAKEN Zito Media Voice 624G | 25&#160;Sep&#160;23&#160;17:58&#160;UTC | true | [view](CERTS/d6a216dd047251b6976b020aa8d2e7986b2342459ba23bf2e81a270369166bb1/README.md) |
| 03&#160;Jul&#160;23&#160;16:39&#160;UTC | SHAKEN Mercury Access Solutions 634K | 02&#160;Jul&#160;24&#160;16:39&#160;UTC | true | [view](CERTS/bbfd747cc6eb37a02ef174b5ada462124ba7990aae403e63cf86516d1fdfb816/README.md) |
| 11&#160;Jul&#160;23&#160;21:11&#160;UTC | SHAKEN Dalton Utilities 3139 | 01&#160;Jan&#160;24&#160;21:11&#160;UTC | true | [view](CERTS/44ee90926a740c783c607a1afc3c73605c6347147f338fd0eca781ad16721246/README.md) |
| 12&#160;Jul&#160;23&#160;14:58&#160;UTC | SHAKEN IP Link Telecom Inc. 902J | 10&#160;Oct&#160;23&#160;14:58&#160;UTC | true | [view](CERTS/19c4e7af6b2504818f87c7369227fa70bb64049b573fef75b7fba644215f7236/README.md) |
| 12&#160;Jul&#160;23&#160;14:59&#160;UTC | SHAKEN IP Link Telecom Inc. 902J | 10&#160;Oct&#160;23&#160;14:59&#160;UTC | true | [view](CERTS/8473486fc37a130119da6cf8bc24ec9ee5c489f1d9bedca5ea703560f16b055f/README.md) |
| 20&#160;Jul&#160;23&#160;18:52&#160;UTC | SHAKEN IPBTel 535K | 19&#160;Jul&#160;24&#160;18:52&#160;UTC | true | [view](CERTS/297dc6000e9cc8107a9b6069afd9e8548ca640ba1c4119d9c6187e5d7258def4/README.md) |
| 23&#160;Jul&#160;23&#160;11:40&#160;UTC | SHAKEN IDT America, Corp 363A | 22&#160;Aug&#160;23&#160;11:40&#160;UTC | true | [view](CERTS/91a86f8c65d0e69cf8f1c7a2301660ede6af4a94bda5647cbfb9ff68b03fae89/README.md) |
| 23&#160;Jul&#160;23&#160;23:09&#160;UTC | SHAKEN IPSBS Managed Services LLC 828J | 22&#160;Aug&#160;23&#160;23:09&#160;UTC | true | [view](CERTS/7f6fa07715ae9fc9f21059bd63fd95c548c5d5539fb9411e7cfa935709f8eb78/README.md) |
| 24&#160;Jul&#160;23&#160;00:02&#160;UTC | SHAKEN InteractiveTel, LLC 920J | 23&#160;Aug&#160;23&#160;00:02&#160;UTC | true | [view](CERTS/c58a604dd5c39e100f23d5117b9aa6358c84d89d0124d1d1bb8116ec195705e4/README.md) |
| 24&#160;Jul&#160;23&#160;03:12&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 23&#160;Aug&#160;23&#160;03:12&#160;UTC | true | [view](CERTS/2573a612de2798322de01a4d73f8f9a8c423a6688dc226ab959521fcdd6e0ee4/README.md) |
| 24&#160;Jul&#160;23&#160;07:23&#160;UTC | SHAKEN BareTelecom 864J | 23&#160;Aug&#160;23&#160;07:23&#160;UTC | true | [view](CERTS/2183a14e438b8b8962b03d3ffffbf95a2bce5ca1f59ce9fd38a0ca0d29acd935/README.md) |
| 24&#160;Jul&#160;23&#160;11:35&#160;UTC | SHAKEN IDT America, Corp 363A | 23&#160;Aug&#160;23&#160;11:35&#160;UTC | true | [view](CERTS/6549f2ffe75677d8d82c027a19f2d7d1911c45e9c6d98edc6dc6f8c557390a48/README.md) |
| 24&#160;Jul&#160;23&#160;11:59&#160;UTC | SHAKEN ONE OWL TELECOM INC 412K | 23&#160;Aug&#160;23&#160;11:59&#160;UTC | true | [view](CERTS/37bdfdf1c040cb266771e78a449d5f3495e45d9c451db075cab0318e74b47f88/README.md) |
| 24&#160;Jul&#160;23&#160;13:17&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 23&#160;Aug&#160;23&#160;13:17&#160;UTC | true | [view](CERTS/47cc37d02b1cd0e79473da49a8c99896ded4fe80e8b80846c4f713c2ba353a76/README.md) |
| 24&#160;Jul&#160;23&#160;23:04&#160;UTC | SHAKEN IPSBS Managed Services LLC 828J | 23&#160;Aug&#160;23&#160;23:04&#160;UTC | true | [view](CERTS/5b85027c5cd17ab11cc449c1a67672fe68557ebee1028bd2f06f945b9dc0f1a9/README.md) |
| 25&#160;Jul&#160;23&#160;03:07&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 24&#160;Aug&#160;23&#160;03:07&#160;UTC | true | [view](CERTS/8daddb43a973486a17cfd3f76865c84dc351a22141494cff7972bd6740ceb0a6/README.md) |
| 25&#160;Jul&#160;23&#160;07:18&#160;UTC | SHAKEN BareTelecom 864J | 24&#160;Aug&#160;23&#160;07:18&#160;UTC | true | [view](CERTS/bf714df8af35d9c244a2420c63e593a1915ab037003c79b81a2af543859565ff/README.md) |
| 25&#160;Jul&#160;23&#160;11:28&#160;UTC | SHAKEN IDT America, Corp 363A | 24&#160;Aug&#160;23&#160;11:28&#160;UTC | true | [view](CERTS/d81058d3816c0c9ce1b97b19ca23f90a4fd2bfe2ba7a4ec051fe95fdc6013ad4/README.md) |
| 25&#160;Jul&#160;23&#160;11:30&#160;UTC | SHAKEN IDT America, Corp 363A | 24&#160;Aug&#160;23&#160;11:30&#160;UTC | true | [view](CERTS/44d06770e3864d5b3f4da2dda469c30522cf70fd18c6b5cfb2a4a9a3c418a29d/README.md) |
| 25&#160;Jul&#160;23&#160;13:12&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 24&#160;Aug&#160;23&#160;13:12&#160;UTC | true | [view](CERTS/0f55461033589618022edce24e63bb6a37974acad75dcccb890ed870c5eb55fa/README.md) |
| 25&#160;Jul&#160;23&#160;15:28&#160;UTC | SHAKEN Zray Technologies Corporation 862J | 24&#160;Aug&#160;23&#160;15:28&#160;UTC | true | [view](CERTS/1e2f01dcb0f8d8075ad210278c7f5462dd7c35409063c16ce039b1af1383f04b/README.md) |
| 25&#160;Jul&#160;23&#160;23:52&#160;UTC | SHAKEN InteractiveTel, LLC 920J | 24&#160;Aug&#160;23&#160;23:52&#160;UTC | true | [view](CERTS/3088cfc2ae18a42fe69817d6edd162463d08e54fb95e410b4ae63390fd9702ba/README.md) |
| 26&#160;Jul&#160;23&#160;03:02&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 25&#160;Aug&#160;23&#160;03:02&#160;UTC | true | [view](CERTS/d196efa7a384127d5598b3c01bc5e22f18f24086d927c15b0251d2301c8222cf/README.md) |
| 26&#160;Jul&#160;23&#160;07:13&#160;UTC | SHAKEN BareTelecom 864J | 25&#160;Aug&#160;23&#160;07:13&#160;UTC | true | [view](CERTS/02f98b7dddde409bbdcc966bdaf440c7356c957e7d35cfc4886ae9bbfd3d6948/README.md) |
| 26&#160;Jul&#160;23&#160;11:23&#160;UTC | SHAKEN IDT America, Corp 363A | 25&#160;Aug&#160;23&#160;11:23&#160;UTC | true | [view](CERTS/1e828ab5134aa6eb61b90c92d789a33e445966727b4cec7c34b9cb7c0775ba1b/README.md) |
| 26&#160;Jul&#160;23&#160;11:25&#160;UTC | SHAKEN IDT America, Corp 363A | 25&#160;Aug&#160;23&#160;11:25&#160;UTC | true | [view](CERTS/e6b003d789f435302e0f22512a19020dc8e887b0cab9a632c8b644bf895aba18/README.md) |
| 26&#160;Jul&#160;23&#160;13:07&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 25&#160;Aug&#160;23&#160;13:07&#160;UTC | true | [view](CERTS/3433d3ebd43662e56141fe783e645289004c1ff44bcaf4c00119c82b3ff36188/README.md) |
| 27&#160;Jul&#160;23&#160;02:56&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 26&#160;Aug&#160;23&#160;02:56&#160;UTC | true | [view](CERTS/54ff56377f724991303b0228ab7360fb4348ed5150d46c8f0771e06322179252/README.md) |
| 27&#160;Jul&#160;23&#160;06:30&#160;UTC | SHAKEN  XCast Labs 689J | 27&#160;Aug&#160;23&#160;06:30&#160;UTC | true | [view](CERTS/21e6651b3448f1b14d0625c0ee80f66110972d47c9c5eae632ceb6ed4a511633/README.md) |
| 27&#160;Jul&#160;23&#160;07:08&#160;UTC | SHAKEN BareTelecom 864J | 26&#160;Aug&#160;23&#160;07:08&#160;UTC | true | [view](CERTS/492d87762b51bd50f12a09d988af763eaaf64aae01fa80a01a7cba6e285bd7b2/README.md) |
| 27&#160;Jul&#160;23&#160;07:24&#160;UTC | SHAKEN IDT America, Corp 363A | 26&#160;Aug&#160;23&#160;07:24&#160;UTC | true | [view](CERTS/5b7541d9d1f0094c98ae65fc3fbc84ab801156cf880333d8713a3df254c6ffa6/README.md) |
| 27&#160;Jul&#160;23&#160;07:38&#160;UTC | SHAKEN IDT America, Corp 363A | 26&#160;Aug&#160;23&#160;07:38&#160;UTC | true | [view](CERTS/6facd462379432680d61b452962596ce26ce437b6d9a4a17479f379353204f3f/README.md) |
| 27&#160;Jul&#160;23&#160;11:18&#160;UTC | SHAKEN IDT America, Corp 363A | 26&#160;Aug&#160;23&#160;11:18&#160;UTC | true | [view](CERTS/74b6176287c37b0dcb82787ac50e05e7c155a6c637279ee366fb864c101561f7/README.md) |
| 27&#160;Jul&#160;23&#160;11:20&#160;UTC | SHAKEN IDT America, Corp 363A | 26&#160;Aug&#160;23&#160;11:20&#160;UTC | true | [view](CERTS/0ff3ec11ac9f617bcabbdbdb764ab3437a111b0acb1cbb7572dbc086771d0f54/README.md) |
| 27&#160;Jul&#160;23&#160;13:02&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 26&#160;Aug&#160;23&#160;13:02&#160;UTC | true | [view](CERTS/735330b505d037448c3adb7d29c634007721dea318bceba06b92a54e1cd66f12/README.md) |
| 28&#160;Jul&#160;23&#160;02:52&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 27&#160;Aug&#160;23&#160;02:52&#160;UTC | true | [view](CERTS/24429d7836d91596c7c009daef85bcd2971ec8377c99e39edab244472d8b241a/README.md) |
| 28&#160;Jul&#160;23&#160;07:03&#160;UTC | SHAKEN BareTelecom 864J | 27&#160;Aug&#160;23&#160;07:03&#160;UTC | true | [view](CERTS/d1914116a366108af0227905872b38624df0e55c35d71969c3b8de6f88679bd3/README.md) |
| 28&#160;Jul&#160;23&#160;07:19&#160;UTC | SHAKEN IDT America, Corp 363A | 27&#160;Aug&#160;23&#160;07:19&#160;UTC | true | [view](CERTS/b1d832cf539034be11f12d66a4d8e7fafd333ae601cf2be95277ad12e0418dac/README.md) |
| 28&#160;Jul&#160;23&#160;07:33&#160;UTC | SHAKEN IDT America, Corp 363A | 27&#160;Aug&#160;23&#160;07:33&#160;UTC | true | [view](CERTS/142448519522ad118cf4ecaed913289b686b54f4f2ac6aa5c2d7a43227bf7923/README.md) |
| 28&#160;Jul&#160;23&#160;11:13&#160;UTC | SHAKEN IDT America, Corp 363A | 27&#160;Aug&#160;23&#160;11:13&#160;UTC | true | [view](CERTS/2329185a6a31e73c46d84fd48ec76957b634ec1a826708cdded0314390e3ebf4/README.md) |
| 28&#160;Jul&#160;23&#160;11:15&#160;UTC | SHAKEN IDT America, Corp 363A | 27&#160;Aug&#160;23&#160;11:15&#160;UTC | true | [view](CERTS/1ccd8190b61b6bd5f6821acd3d6bb39eead966f20f619dddec097e6d9284ca9a/README.md) |
| 28&#160;Jul&#160;23&#160;12:57&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 27&#160;Aug&#160;23&#160;12:57&#160;UTC | true | [view](CERTS/3e65aad9857f73511d0863b7cee1e8e421dc74348b0d95d8ec788887874a354c/README.md) |
| 28&#160;Jul&#160;23&#160;14:52&#160;UTC | SHAKEN Socket Telecom LLC 554a | 27&#160;Aug&#160;23&#160;14:52&#160;UTC | true | [view](CERTS/b60688aa3ed21dc8cba0acefb75c3807b68433a3e14f64a352a36afafd97006c/README.md) |
| 29&#160;Jul&#160;23&#160;06:58&#160;UTC | SHAKEN BareTelecom 864J | 28&#160;Aug&#160;23&#160;06:58&#160;UTC | true | [view](CERTS/cc8eb4d2f031f848cda55aa8fb202bdf5785089254801619c9026d5b02b0f37f/README.md) |
| 30&#160;Jul&#160;23&#160;07:09&#160;UTC | SHAKEN IDT America, Corp 363A | 29&#160;Aug&#160;23&#160;07:09&#160;UTC | true | [view](CERTS/452d72975c83bb529855aacb3889237fe97f09efae363f01288fa2078442978d/README.md) |
| 30&#160;Jul&#160;23&#160;07:23&#160;UTC | SHAKEN IDT America, Corp 363A | 29&#160;Aug&#160;23&#160;07:23&#160;UTC | true | [view](CERTS/84dddd55c266aed2f3e914c2da1abaeb139ae7b163763c25a047384d2a527406/README.md) |
| 30&#160;Jul&#160;23&#160;11:03&#160;UTC | SHAKEN IDT America, Corp 363A | 29&#160;Aug&#160;23&#160;11:03&#160;UTC | true | [view](CERTS/8ffcb933f007b5b9e53e3a65e9723ffc4574b4afa199b97d2bf6d97e57608a7b/README.md) |
| 31&#160;Jul&#160;23&#160;02:36&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 30&#160;Aug&#160;23&#160;02:36&#160;UTC | true | [view](CERTS/dfda83bbe12aafb81eb8301e7e14266aca58e3e376fa9439d98c995829007cb6/README.md) |
| 31&#160;Jul&#160;23&#160;06:48&#160;UTC | SHAKEN BareTelecom 864J | 30&#160;Aug&#160;23&#160;06:48&#160;UTC | true | [view](CERTS/4fdf300441f31ab8a40c7ca74c5a1f1a9b59bb7a9de295e17c1d5c36e7b88f6a/README.md) |
| 31&#160;Jul&#160;23&#160;07:04&#160;UTC | SHAKEN IDT America, Corp 363A | 30&#160;Aug&#160;23&#160;07:04&#160;UTC | true | [view](CERTS/27c0dc5afccf28d4220f65b5c9914511b15d993ab30e712d22c74c5b129bc65e/README.md) |
| 31&#160;Jul&#160;23&#160;07:18&#160;UTC | SHAKEN IDT America, Corp 363A | 30&#160;Aug&#160;23&#160;07:18&#160;UTC | true | [view](CERTS/87174a545b58be96982f847c393fbdb48179bd24024805d89f2ed4a71f994ab4/README.md) |
| 31&#160;Jul&#160;23&#160;10:58&#160;UTC | SHAKEN IDT America, Corp 363A | 30&#160;Aug&#160;23&#160;10:58&#160;UTC | true | [view](CERTS/a194b098d513ece68733106fc3e259254f09f796d186b4a73be2b4f6e3781b03/README.md) |
| 31&#160;Jul&#160;23&#160;12:42&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 30&#160;Aug&#160;23&#160;12:42&#160;UTC | true | [view](CERTS/12556dc17d33930d85fe0ec4e7fb281ee65d5c6f4edcffc5e1381b84dd233da8/README.md) |
| 31&#160;Jul&#160;23&#160;22:32&#160;UTC | SHAKEN IPSBS Managed Services LLC 828J | 30&#160;Aug&#160;23&#160;22:32&#160;UTC | true | [view](CERTS/ab70078bf98c21a5e85a0b8dc768da67bade2a9b222eb0b4255d4c11bb94a63a/README.md) |
| 31&#160;Jul&#160;23&#160;23:22&#160;UTC | SHAKEN InteractiveTel, LLC 920J | 30&#160;Aug&#160;23&#160;23:22&#160;UTC | true | [view](CERTS/6788387a4000afe4a89203d407362a1547b13f9bc32deabbaedaa9a8fbb854b9/README.md) |
| 01&#160;Aug&#160;23&#160;02:32&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 31&#160;Aug&#160;23&#160;02:32&#160;UTC | true | [view](CERTS/9ca83e241529bf6606d33b1dea8e2918958733f5b9eb9832bbd3a856bd4e2c72/README.md) |
| 01&#160;Aug&#160;23&#160;06:43&#160;UTC | SHAKEN BareTelecom 864J | 31&#160;Aug&#160;23&#160;06:43&#160;UTC | true | [view](CERTS/86cfc2670b62e8e3bb49f0d331de317adb50200c33e4b8d2afc20bb2a8946793/README.md) |
| 01&#160;Aug&#160;23&#160;06:59&#160;UTC | SHAKEN IDT America, Corp 363A | 31&#160;Aug&#160;23&#160;06:59&#160;UTC | true | [view](CERTS/4d01466bf29c8ddd1b814f4e31db2a87c5c14c5a58a0bdba10c815d5e6a5ff9e/README.md) |
| 01&#160;Aug&#160;23&#160;07:13&#160;UTC | SHAKEN IDT America, Corp 363A | 31&#160;Aug&#160;23&#160;07:13&#160;UTC | true | [view](CERTS/90ccaa5d06eb126988a29bc2706358eb240f0277fbe987d435820df3f499e97e/README.md) |
| 01&#160;Aug&#160;23&#160;12:37&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 31&#160;Aug&#160;23&#160;12:37&#160;UTC | true | [view](CERTS/80c5d92ed9f9787f07077599fe665aaea7bd3ae06182fbabc4291b2702d7f763/README.md) |
| 01&#160;Aug&#160;23&#160;14:46&#160;UTC | SHAKEN Cyberlynk Network, LLC 086K | 31&#160;Aug&#160;23&#160;14:46&#160;UTC | true | [view](CERTS/dd843d4688f8d5b655b98aad39ca9337d5b0a27e2521854cd484a159d9cece1d/README.md) |
| 01&#160;Aug&#160;23&#160;21:21&#160;UTC | SHAKEN Doylestown Communications, Inc 849C | 31&#160;Aug&#160;23&#160;21:21&#160;UTC | true | [view](CERTS/6dbd4112c41267196faac0e733026505510637746c7608931f493b536b3f5e17/README.md) |
| 01&#160;Aug&#160;23&#160;22:27&#160;UTC | SHAKEN IPSBS Managed Services LLC 828J | 31&#160;Aug&#160;23&#160;22:27&#160;UTC | true | [view](CERTS/d3e3855000fe79c6c0eb761b43bd1348b52919b92cc272710bff8ebe8d370cbf/README.md) |
| 02&#160;Aug&#160;23&#160;02:27&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 01&#160;Sep&#160;23&#160;02:27&#160;UTC | true | [view](CERTS/a01b82847d5b73a1995ff512864c5d78af5576e6101d5e897cc9c3ecb00ba9e6/README.md) |
| 02&#160;Aug&#160;23&#160;06:38&#160;UTC | SHAKEN BareTelecom 864J | 01&#160;Sep&#160;23&#160;06:38&#160;UTC | true | [view](CERTS/ad3ae56953de3be6dcab9d8795e8ada1429d11c14d75dd45f55b0d8a43cfd73b/README.md) |
| 02&#160;Aug&#160;23&#160;06:54&#160;UTC | SHAKEN IDT America, Corp 363A | 01&#160;Sep&#160;23&#160;06:54&#160;UTC | true | [view](CERTS/ea116798ed62fb9fc270a1b0ff4b176519473c5f9a375f0bb2a41855a35a6a23/README.md) |
| 02&#160;Aug&#160;23&#160;07:08&#160;UTC | SHAKEN IDT America, Corp 363A | 01&#160;Sep&#160;23&#160;07:08&#160;UTC | true | [view](CERTS/507e04cce2a84163b8cd68617698a70058ef6eae1d06d790bdc82ea0d2604554/README.md) |
| 02&#160;Aug&#160;23&#160;08:58&#160;UTC | SHAKEN IDT America, Corp 363A | 01&#160;Sep&#160;23&#160;08:58&#160;UTC | true | [view](CERTS/f7b16dbd37a2b15fd210aafa80906bc525b69f5f3608b954a45e6fc5b87f77ab/README.md) |
| 02&#160;Aug&#160;23&#160;12:32&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 01&#160;Sep&#160;23&#160;12:32&#160;UTC | true | [view](CERTS/a659959cf2ad9fa43626527c3a870fe40d731a5ae658aa4a464c3a08c69a4262/README.md) |
| 02&#160;Aug&#160;23&#160;14:27&#160;UTC | SHAKEN Socket Telecom LLC 554a | 01&#160;Sep&#160;23&#160;14:27&#160;UTC | true | [view](CERTS/915aad35032d1c65c1d993fab35cf2a68eec33ce96ecef53f454ef00a183c618/README.md) |
| 02&#160;Aug&#160;23&#160;22:22&#160;UTC | SHAKEN IPSBS Managed Services LLC 828J | 01&#160;Sep&#160;23&#160;22:22&#160;UTC | true | [view](CERTS/c155fcb8e419cd626ca82087626e4d7ef20e5b4409eadc89a0f2178f6192e574/README.md) |
| 03&#160;Aug&#160;23&#160;02:22&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 02&#160;Sep&#160;23&#160;02:22&#160;UTC | true | [view](CERTS/ea18ef4f50baf783f7bc649c5c63b68f2624e623b293de5cb6755113690b992a/README.md) |
| 03&#160;Aug&#160;23&#160;06:33&#160;UTC | SHAKEN BareTelecom 864J | 02&#160;Sep&#160;23&#160;06:33&#160;UTC | true | [view](CERTS/0b163969c0a19c619d9c2a225adc308007e90fff192d633be8eeacb8e3597700/README.md) |
| 03&#160;Aug&#160;23&#160;06:49&#160;UTC | SHAKEN IDT America, Corp 363A | 02&#160;Sep&#160;23&#160;06:49&#160;UTC | true | [view](CERTS/0657b5ed87f11d284f915a7148755830b6ec3b4f7dd22ea2948618f80891993e/README.md) |
| 03&#160;Aug&#160;23&#160;07:03&#160;UTC | SHAKEN IDT America, Corp 363A | 02&#160;Sep&#160;23&#160;07:03&#160;UTC | true | [view](CERTS/ad38d6d364d014e81f8df445c46d4e3b32caeaba518ff77c9e6a2f2f2ea2cd51/README.md) |
| 03&#160;Aug&#160;23&#160;07:17&#160;UTC | SHAKEN IDT America, Corp 363A | 02&#160;Sep&#160;23&#160;07:17&#160;UTC | true | [view](CERTS/a826b1a8865bce9390aacea3eb95a31683b7fdcce7b277803a946fd5cd92275e/README.md) |
| 03&#160;Aug&#160;23&#160;08:53&#160;UTC | SHAKEN IDT America, Corp 363A | 02&#160;Sep&#160;23&#160;08:53&#160;UTC | true | [view](CERTS/407e3726d5dab56ce21252285af217bf6bf8936b6b4a3952a2608e63a0149dca/README.md) |
| 03&#160;Aug&#160;23&#160;12:28&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 02&#160;Sep&#160;23&#160;12:28&#160;UTC | true | [view](CERTS/3cbc7d2bd58d6aa603707bc5efd21bf88265d27f60d4c54568b019dd0e58db16/README.md) |
| 03&#160;Aug&#160;23&#160;14:22&#160;UTC | SHAKEN Socket Telecom LLC 554a | 02&#160;Sep&#160;23&#160;14:22&#160;UTC | true | [view](CERTS/d5a89088c6314468bccf0e266dda54b449392ed4ac1b6134a3ade322cf3ef667/README.md) |
| 03&#160;Aug&#160;23&#160;14:36&#160;UTC | SHAKEN Cyberlynk Network, LLC 086K | 02&#160;Sep&#160;23&#160;14:36&#160;UTC | true | [view](CERTS/b4e10105a2852f5be3bbcba032c58cf0ea64015a1556cdab7f92ee4dc8acb813/README.md) |
| 03&#160;Aug&#160;23&#160;21:26&#160;UTC | SHAKEN Zray Technologies Corporation 862J | 02&#160;Sep&#160;23&#160;21:26&#160;UTC | true | [view](CERTS/d9025569ce42f2e33604468201e5bf95087ab368559cc6d6b5ad7870b2ddb0a9/README.md) |
| 03&#160;Aug&#160;23&#160;23:26&#160;UTC | SHAKEN NTC International, INC 016K | 02&#160;Sep&#160;23&#160;23:26&#160;UTC | true | [view](CERTS/47d5d0135f96a8203533a744025bc562ad62cab16ff076aab619aca9d93ef028/README.md) |
| 04&#160;Aug&#160;23&#160;02:16&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 03&#160;Sep&#160;23&#160;02:16&#160;UTC | true | [view](CERTS/4a728362a09c89de708a6ee9027224fd62538c714715467c4b21c5e4ddf11287/README.md) |
| 04&#160;Aug&#160;23&#160;06:28&#160;UTC | SHAKEN BareTelecom 864J | 03&#160;Sep&#160;23&#160;06:28&#160;UTC | true | [view](CERTS/ccbdedd68e01fbca81d7d99e810e2847dec4d0b995c3df9a097a3aea31206210/README.md) |
| 04&#160;Aug&#160;23&#160;06:44&#160;UTC | SHAKEN IDT America, Corp 363A | 03&#160;Sep&#160;23&#160;06:44&#160;UTC | true | [view](CERTS/ba667d06da48864f91f1f8f9e475546f732c14cf84d3add31511434724c35da7/README.md) |
| 04&#160;Aug&#160;23&#160;07:13&#160;UTC | SHAKEN IDT America, Corp 363A | 03&#160;Sep&#160;23&#160;07:13&#160;UTC | true | [view](CERTS/8fc0ac92ffe621a7a635427e855206ae73c0a293d4b7496de8a1fde5a35c2e26/README.md) |
| 04&#160;Aug&#160;23&#160;08:48&#160;UTC | SHAKEN IDT America, Corp 363A | 03&#160;Sep&#160;23&#160;08:48&#160;UTC | true | [view](CERTS/2df3f675812818e2d19668f9007e457b3d8b954c8ee135b9267cab891815a6a7/README.md) |
| 04&#160;Aug&#160;23&#160;12:22&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 03&#160;Sep&#160;23&#160;12:22&#160;UTC | true | [view](CERTS/897d5c279ecea96a09586a231e4f2a4be6a321d041bd224c9c356a17c8f8d4d4/README.md) |
| 05&#160;Aug&#160;23&#160;06:23&#160;UTC | SHAKEN BareTelecom 864J | 04&#160;Sep&#160;23&#160;06:23&#160;UTC | true | [view](CERTS/9a5e3e18f3b96c7189b9e65d0e12cf2c46919c784230880558e8136887d3b961/README.md) |
| 05&#160;Aug&#160;23&#160;06:39&#160;UTC | SHAKEN IDT America, Corp 363A | 04&#160;Sep&#160;23&#160;06:39&#160;UTC | true | [view](CERTS/2016acb45a1c955b370bdeac633d3a76e0abb0472f61c972c094d49af8313521/README.md) |
| 05&#160;Aug&#160;23&#160;06:53&#160;UTC | SHAKEN IDT America, Corp 363A | 04&#160;Sep&#160;23&#160;06:53&#160;UTC | true | [view](CERTS/df44b8129c35a11689e615307587aa342c191bd60ae0bc06cf3560efe7df1b02/README.md) |
| 05&#160;Aug&#160;23&#160;07:07&#160;UTC | SHAKEN IDT America, Corp 363A | 04&#160;Sep&#160;23&#160;07:07&#160;UTC | true | [view](CERTS/7425a1343be1c001dd5a3b80f4855512f65e4b175594f479893411708fc99f3b/README.md) |
| 05&#160;Aug&#160;23&#160;12:17&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 04&#160;Sep&#160;23&#160;12:17&#160;UTC | true | [view](CERTS/6c8a3877d04496b16b9cce544ba72435f4898e79de9eb1584b93933898077a67/README.md) |
| 05&#160;Aug&#160;23&#160;20:18&#160;UTC | SHAKEN BareTelecom 864J | 04&#160;Sep&#160;23&#160;20:18&#160;UTC | true | [view](CERTS/6f769d84fbd3fd6dbb160ee369a4de4f0b0368f026a3109c598d11ed96e8c292/README.md) |
| 06&#160;Aug&#160;23&#160;08:38&#160;UTC | SHAKEN IDT America, Corp 363A | 05&#160;Sep&#160;23&#160;08:38&#160;UTC | true | [view](CERTS/9fb430da1dd816ab69d9b0e3e2ce45e2f0bc96f7a6e60b1bf7d682a10138ee22/README.md) |
| 06&#160;Aug&#160;23&#160;12:12&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 05&#160;Sep&#160;23&#160;12:12&#160;UTC | true | [view](CERTS/df9563c5cff2224fcd75d168f2993ad1e7eeb2953d521170ef90104fb87c2b47/README.md) |
| 06&#160;Aug&#160;23&#160;20:13&#160;UTC | SHAKEN BareTelecom 864J | 05&#160;Sep&#160;23&#160;20:13&#160;UTC | true | [view](CERTS/822a1852e557d24451271c92ed0963b9675750260f910eb6149673a928a4cdd0/README.md) |
| 06&#160;Aug&#160;23&#160;21:11&#160;UTC | SHAKEN Zray Technologies Corporation 862J | 05&#160;Sep&#160;23&#160;21:11&#160;UTC | true | [view](CERTS/97ce9700e8741ac295f9afe16e3a80fa0a2b95e0f3e794b2a2ce93ef134cf292/README.md) |
| 06&#160;Aug&#160;23&#160;22:02&#160;UTC | SHAKEN IPSBS Managed Services LLC 828J | 05&#160;Sep&#160;23&#160;22:02&#160;UTC | true | [view](CERTS/96bea1a12b4204610dfa635c75b73096867a0e8be8174cebe01e84d5d97e0767/README.md) |
| 06&#160;Aug&#160;23&#160;23:11&#160;UTC | SHAKEN NTC International, INC 016K | 05&#160;Sep&#160;23&#160;23:11&#160;UTC | true | [view](CERTS/e39efb25d9820dc52ffc7c34a4620095f036d27e3b1b67b4a174cb7bc225fd5b/README.md) |
| 07&#160;Aug&#160;23&#160;02:01&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 06&#160;Sep&#160;23&#160;02:01&#160;UTC | true | [view](CERTS/fbfa16ffa1f92c1ced5c31548ebcbbb422be44a3c12bae3ded2fbba6823b42c2/README.md) |
| 07&#160;Aug&#160;23&#160;06:00&#160;UTC | SHAKEN Convoso 758J | 11&#160;Sep&#160;23&#160;06:00&#160;UTC | true | [view](CERTS/f4befd8381b1fb4f492ebff36b91c77c8bacc8f759c285395f2b3eb8293074ac/README.md) |
| 07&#160;Aug&#160;23&#160;06:58&#160;UTC | SHAKEN IDT America, Corp 363A | 06&#160;Sep&#160;23&#160;06:58&#160;UTC | true | [view](CERTS/56cb8f01440795bc3a4a8e2ca6560a26684b589c8a503e464f8a5e0515b4d94d/README.md) |
| 07&#160;Aug&#160;23&#160;08:33&#160;UTC | SHAKEN IDT America, Corp 363A | 06&#160;Sep&#160;23&#160;08:33&#160;UTC | true | [view](CERTS/cbc4f3807261c7b94b8c4bce433da9227c25646407312e54fe2fd8bed1cced1a/README.md) |
| 07&#160;Aug&#160;23&#160;12:07&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 06&#160;Sep&#160;23&#160;12:07&#160;UTC | true | [view](CERTS/adb19dd2045915342e0569f5b00cabc3aefa3a5acafaf23802430df4c8987748/README.md) |
| 07&#160;Aug&#160;23&#160;20:08&#160;UTC | SHAKEN BareTelecom 864J | 06&#160;Sep&#160;23&#160;20:08&#160;UTC | true | [view](CERTS/5032190085a0e7fe8d238d756bb3947afa3f5ba32c34e394346eea9b84761045/README.md) |
| 07&#160;Aug&#160;23&#160;21:06&#160;UTC | SHAKEN Zray Technologies Corporation 862J | 06&#160;Sep&#160;23&#160;21:06&#160;UTC | true | [view](CERTS/76090f09ede478cf08bc69e825c609104465d8deaf7fa111ba79a0318917538d/README.md) |
| 07&#160;Aug&#160;23&#160;21:57&#160;UTC | SHAKEN IPSBS Managed Services LLC 828J | 06&#160;Sep&#160;23&#160;21:57&#160;UTC | true | [view](CERTS/2d6fd09cf2f1d2ddfc573195d575bf8bd2cfdf13aa6beeb06313b5d593250017/README.md) |
| 08&#160;Aug&#160;23&#160;06:52&#160;UTC | SHAKEN IDT America, Corp 363A | 07&#160;Sep&#160;23&#160;06:52&#160;UTC | true | [view](CERTS/60bafb60159d4587affa204deb4be83858f816949d00961e6b2632b00061f841/README.md) |
| 08&#160;Aug&#160;23&#160;08:28&#160;UTC | SHAKEN IDT America, Corp 363A | 07&#160;Sep&#160;23&#160;08:28&#160;UTC | true | [view](CERTS/2efb8b3adaf257995ecf7909754504ae499be29e99fefb16fef0ac8ed3954352/README.md) |
| 08&#160;Aug&#160;23&#160;12:02&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 07&#160;Sep&#160;23&#160;12:02&#160;UTC | true | [view](CERTS/de0a8520aaf2825b748c3cf6de8eee0cd6be07e7f24a32b33c283009d1194f4a/README.md) |
| 08&#160;Aug&#160;23&#160;20:03&#160;UTC | SHAKEN BareTelecom 864J | 07&#160;Sep&#160;23&#160;20:03&#160;UTC | true | [view](CERTS/70657cafba97fbb6338564720a00af108ba016b4cbb0fa7d735ba63fdddd2ba0/README.md) |
| 09&#160;Aug&#160;23&#160;01:52&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 08&#160;Sep&#160;23&#160;01:52&#160;UTC | true | [view](CERTS/84651bf36d052f20720339aa800daf3740f0ea3f564b24a38ed6f7ddc3b6edba/README.md) |
| 09&#160;Aug&#160;23&#160;06:48&#160;UTC | SHAKEN IDT America, Corp 363A | 08&#160;Sep&#160;23&#160;06:48&#160;UTC | true | [view](CERTS/c2be1915a3152638e91cb34dd28a42c7986d69a8418022b4bac63e9f2e96e1b8/README.md) |
| 09&#160;Aug&#160;23&#160;08:23&#160;UTC | SHAKEN IDT America, Corp 363A | 08&#160;Sep&#160;23&#160;08:23&#160;UTC | true | [view](CERTS/fb7515b3e7c018411a812eecc6bfa7837fd05e8b9b56b7a80dde544a167730a1/README.md) |
| 09&#160;Aug&#160;23&#160;11:57&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 08&#160;Sep&#160;23&#160;11:57&#160;UTC | true | [view](CERTS/2ab42299e3946f30f467b5422137b661198d6c02a1d1e32a59ca2b3cee26b8f0/README.md) |
| 15&#160;Aug&#160;23&#160;13:27&#160;UTC | SHAKEN Threshold Communications Inc 563J | 14&#160;Sep&#160;23&#160;13:27&#160;UTC | true | [view](CERTS/5ac0991bbbf851c5fbefc1db2ba900170fe465c7a5c299b0211f2ffe3232660d/README.md) |
| 15&#160;Aug&#160;23&#160;13:28&#160;UTC | SHAKEN Broadband Dynamics LLC 583j | 14&#160;Sep&#160;23&#160;13:28&#160;UTC | true | [view](CERTS/eb450c457fbe963371412ee1402d71d021ca1461dae9145ac53fc495a6e4988d/README.md) |
| 15&#160;Aug&#160;23&#160;13:28&#160;UTC | SHAKEN Consolidated Communications 5113 | 14&#160;Sep&#160;23&#160;13:28&#160;UTC | true | [view](CERTS/e211b03d642dc7c5b5085fc722a5c043950d6566e5ea7978342602a466a895ba/README.md) |
| 15&#160;Aug&#160;23&#160;13:29&#160;UTC | SHAKEN Touchtone 683A | 14&#160;Sep&#160;23&#160;13:29&#160;UTC | true | [view](CERTS/774ad286ee97f9ece0d07d8172a720c8dd35c4d61d94670c62464f436db2eff2/README.md) |
| 15&#160;Aug&#160;23&#160;13:30&#160;UTC | SHAKEN Apeiron Systems 012J | 14&#160;Sep&#160;23&#160;13:30&#160;UTC | true | [view](CERTS/ec173d6928cfd7c3c9b9f903161349f1c7102790f4b97b076b983e2e27322dcf/README.md) |
| 15&#160;Aug&#160;23&#160;13:31&#160;UTC | SHAKEN Fonative, Inc. 684J | 14&#160;Sep&#160;23&#160;13:31&#160;UTC | true | [view](CERTS/f9dde013ffe31ec2ce49ffc5b81c128fe245a7a78ad7dffaec2cacf571e151e0/README.md) |
| 15&#160;Aug&#160;23&#160;13:32&#160;UTC | SHAKEN IPitomy 652J | 14&#160;Sep&#160;23&#160;13:32&#160;UTC | true | [view](CERTS/ea6c240d58f50bf618bfa45dd3d2db5d20e53841e44d782603aba0d00188cdf0/README.md) |
| 15&#160;Aug&#160;23&#160;13:32&#160;UTC | SHAKEN Phone.com, Inc. 633J | 14&#160;Sep&#160;23&#160;13:32&#160;UTC | true | [view](CERTS/b4d6be9578eae535cab17fbb7fda4013b9aa59329aab0efb3b2796c983814860/README.md) |
| 15&#160;Aug&#160;23&#160;13:33&#160;UTC | SHAKEN NETRIO LLC 020K | 14&#160;Sep&#160;23&#160;13:33&#160;UTC | true | [view](CERTS/7f561742ed5b0da3df6b43b187cc4a250d83df5c1ddcb12b3a7b5df3a23994b3/README.md) |
| 15&#160;Aug&#160;23&#160;13:35&#160;UTC | SHAKEN Noble Systems Communications LLC 187J | 14&#160;Sep&#160;23&#160;13:35&#160;UTC | true | [view](CERTS/08917a7a206eb9b99c0a5e7d702c7ef48bfe3ac4d472d2d2abf047db5b1717f0/README.md) |
| 15&#160;Aug&#160;23&#160;13:35&#160;UTC | SHAKEN Telcentris Inc. dba Voxox 696J | 14&#160;Sep&#160;23&#160;13:35&#160;UTC | true | [view](CERTS/94c088c8ea7a848ce65c68deaf7b5d5fc891c0b4ce22ae9f84a0b7bea499ef05/README.md) |
| 15&#160;Aug&#160;23&#160;13:38&#160;UTC | SHAKEN Airespring 996H | 14&#160;Sep&#160;23&#160;13:38&#160;UTC | true | [view](CERTS/71886497f8ef526719d94ed57b0723343de985a4e0461d7f5d93fd8f43da8232/README.md) |
| 15&#160;Aug&#160;23&#160;13:39&#160;UTC | SHAKEN Nobelbiz, Inc. 596J | 14&#160;Sep&#160;23&#160;13:39&#160;UTC | true | [view](CERTS/85de622552e3932b44a194784b2eb7422005e143c2fd0fe9dfef4d44e5d934b5/README.md) |
| 15&#160;Aug&#160;23&#160;13:40&#160;UTC | SHAKEN Momentum Telecom 1417 | 14&#160;Sep&#160;23&#160;13:40&#160;UTC | true | [view](CERTS/bed5eaacde9ccaff378a31c8f91fdc6018659970120bec52eac0b16af703b056/README.md) |
| 15&#160;Aug&#160;23&#160;13:41&#160;UTC | SHAKEN Momentum Telecom 9157 | 14&#160;Sep&#160;23&#160;13:41&#160;UTC | true | [view](CERTS/b4864c4456793ec298d7e1b97be15c0bafb5da2dd780fa6d1f373eff1919fa28/README.md) |
| 15&#160;Aug&#160;23&#160;13:42&#160;UTC | SHAKEN Terra Nova Telecom 382G | 14&#160;Sep&#160;23&#160;13:42&#160;UTC | true | [view](CERTS/3c33885076ac0dec32523679099f5ad48d7b81425015a12e30aee88bcff433fd/README.md) |
| 15&#160;Aug&#160;23&#160;13:43&#160;UTC | SHAKEN Matrix 3058 | 14&#160;Sep&#160;23&#160;13:43&#160;UTC | true | [view](CERTS/5e2d23150f40aa5fb6152b6a75f5cf815fca49013192e4335500c4496d58b664/README.md) |
| 15&#160;Aug&#160;23&#160;13:44&#160;UTC | SHAKEN Matrix 7379 | 14&#160;Sep&#160;23&#160;13:44&#160;UTC | true | [view](CERTS/9ac2238b89762b8c9aa3c4c66fa72b918ec82890decaa99609645948bb012afa/README.md) |
| 15&#160;Aug&#160;23&#160;13:44&#160;UTC | SHAKEN Matrix 9451 | 14&#160;Sep&#160;23&#160;13:44&#160;UTC | true | [view](CERTS/4a90ed152b173276ae9a1df13c12a0c96764e4ab333f6dfc3f28bb13c0675650/README.md) |
| 15&#160;Aug&#160;23&#160;13:46&#160;UTC | SHAKEN PNG Telecommunications Inc 3395 | 14&#160;Sep&#160;23&#160;13:46&#160;UTC | true | [view](CERTS/47482ec42c38a85177234e01461042a41c8499f795bff8b884a01faf900df8f4/README.md) |
| 15&#160;Aug&#160;23&#160;14:47&#160;UTC | SHAKEN Magna5, LLC 3849 | 14&#160;Sep&#160;23&#160;14:47&#160;UTC | true | [view](CERTS/f5e7ebb5bb6bc135c9383196afcc6a00ce097544150cd4cec103819d0c4d5ebd/README.md) |
| 15&#160;Aug&#160;23&#160;14:51&#160;UTC | SHAKEN Magna5, LLC 8249 | 14&#160;Sep&#160;23&#160;14:51&#160;UTC | true | [view](CERTS/9367f56d7b0ce769fd09305ab788eb47f8c6c357424de8d7cb8f233e7ed8c888/README.md) |
| 15&#160;Aug&#160;23&#160;21:14&#160;UTC | SHAKEN Godaddy 463K | 14&#160;Sep&#160;23&#160;21:14&#160;UTC | true | [view](CERTS/7c1116439178d6b902e1f9004bf27cf90fe73d107899ca402ad1b7d41cfba3c8/README.md) |

#### CA Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 21&#160;Aug&#160;20&#160;01:22&#160;UTC | SHAKEN Sansay Root CA US | 16&#160;Aug&#160;40&#160;01:22&#160;UTC | true | [view](CERTS/8356d251b255f4ac3fd108bb79be4c02dcea2d3b13d138892bdb3a70cbe6a343/README.md) |
| 02&#160;Sep&#160;22&#160;20:53&#160;UTC | SHAKEN Sansay Intermediate CA US WEST 1 | 31&#160;Aug&#160;29&#160;20:53&#160;UTC | true | [view](CERTS/4b1dfdba2b1e4bbffbf900a20f1f6f7befbef0008b963e4922a64469cb97d24b/README.md) |


Generated: 21 Aug 23 20:18 UTC