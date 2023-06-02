# STIR/SHAKEN CA Ecosystem Compliance

## Sansay

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 833 certificates were included in the corpus being tested
- 10 certificates in the corpus were skipped because they are duplicates
- 635 certificates in the corpus were skipped because they are expired
- 1 certificates in the corpus were skipped because they are not currently trusted
- 187 certificates being tested against the remaining rules
- 4.93 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 100.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 105 days is the average remaining validity for the certificates in the corpus
- 106 days is the average initial validity for the certificates in the corpus
- 142 certificates expire in the next 30 days
- 2.46 average number of unexpired certificates per OCN observed
- 76 unique OCNs observed in unexpired and valid certificate corpus

| Instances | Test | Source |
|-----------|------|--------|
| 55 | [e_atis_certificate_policies](ISSUES/e_atis_certificate_policies/README.md) | ATIS1000080 |
| 187 | [e_atis_subject_cn](ISSUES/e_atis_subject_cn/README.md) | ATIS1000080 |
| 187 | [e_us_cp_ambiguous_identifier](ISSUES/e_us_cp_ambiguous_identifier/README.md) | US_SHAKEN_CP |
| 187 | [e_us_cp_subject_sn](ISSUES/e_us_cp_subject_sn/README.md) | US_SHAKEN_CP |
| 118 | [w_cp_1_3_subject_email](ISSUES/w_cp_1_3_subject_email/README.md) | US_SHAKEN_CP |
| 187 | [w_shaken_subject_rdn_unknown](ISSUES/w_shaken_subject_rdn_unknown/README.md) | SHAKEN_PKI_BEST_PRACTICES |

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
- 5380 days is the average remaining validity for the certificates in the corpus
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
| 26&#160;Oct&#160;22&#160;19:41&#160;UTC | SHAKEN DLS Internet Services 815J | 26&#160;Oct&#160;23&#160;19:41&#160;UTC | true | [view](CERTS/7cd8319bedd12f040e8bd7b522d981aabcd24dc5aef74614a67fb6fdc9b9823b/README.md) |
| 26&#160;Oct&#160;22&#160;19:43&#160;UTC | SHAKEN Systemverse, LLC. 194K | 26&#160;Oct&#160;23&#160;19:43&#160;UTC | true | [view](CERTS/edbe74f809b9e0e1ebea447df8bdbfb272144f9c8c18df81e397a374df61c4cd/README.md) |
| 26&#160;Oct&#160;22&#160;19:44&#160;UTC | SHAKEN Rayfield Communications, Inc. 006K | 26&#160;Oct&#160;23&#160;19:44&#160;UTC | true | [view](CERTS/5032969f5932ac46a17b86c38dc72d666be454d1c3f11918edfa8385d9fc65e6/README.md) |
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
| 20&#160;Feb&#160;23&#160;17:39&#160;UTC | SHAKEN InPhonex 1821 | 19&#160;Aug&#160;23&#160;17:39&#160;UTC | true | [view](CERTS/eec4207313f9328228bbf09a134b862b705be6ebf77e3418d3800ca13b83cc9a/README.md) |
| 19&#160;Mar&#160;23&#160;00:31&#160;UTC | SHAKEN Every1 Telecom 486K | 18&#160;Mar&#160;24&#160;00:31&#160;UTC | true | [view](CERTS/5cdfbb1a416083096dfef10c75a2b26a08d8fd5593d8ea9ceae0d70d878a97d1/README.md) |
| 27&#160;Mar&#160;23&#160;21:48&#160;UTC | SHAKEN Global Net Holdings Inc 306K | 21&#160;Jan&#160;24&#160;21:48&#160;UTC | true | [view](CERTS/f8d913af64a11c0718c78426b747dba4ed4ccb19239db573626a46f29b671825/README.md) |
| 28&#160;Mar&#160;23&#160;15:31&#160;UTC | SHAKEN VoIP Innovations 597F | 27&#160;Mar&#160;24&#160;15:31&#160;UTC | true | [view](CERTS/105a7683d4b5fac2ea7c2383e95250715bc0460d2cfdbea0d220201f44ea5d0c/README.md) |
| 03&#160;Apr&#160;23&#160;21:02&#160;UTC | SHAKEN WWT INC dba VoIP Networks 053K | 02&#160;Apr&#160;24&#160;21:02&#160;UTC | true | [view](CERTS/6229273e97413ee4c01a1810885db9b48d7e1bf7fe8aa4e7be22076effe2cc8b/README.md) |
| 04&#160;Apr&#160;23&#160;16:23&#160;UTC | SHAKEN Cyberlynk Network, LLC 086K | 03&#160;Jun&#160;23&#160;16:23&#160;UTC | true | [view](CERTS/0b911cba5097116a1c6e0874d1b80f728e3877fd11c49d10ef0625d905b2dfee/README.md) |
| 04&#160;Apr&#160;23&#160;16:38&#160;UTC | SHAKEN Inventive Labs Corp 649J | 01&#160;Oct&#160;23&#160;16:38&#160;UTC | true | [view](CERTS/61bd59e51cd88d19ca695a362569930dee7725fefa39299b507bc3269fefbd54/README.md) |
| 05&#160;Apr&#160;23&#160;16:27&#160;UTC | SHAKEN Swift Telco LLC 452K | 04&#160;Apr&#160;24&#160;16:27&#160;UTC | true | [view](CERTS/947b46067a639b79ff82ab3f48c453e4af7cc6d6036f6d66a742cc935bc8a35e/README.md) |
| 11&#160;Apr&#160;23&#160;11:48&#160;UTC | SHAKEN IP Link Telecom Inc. 902J | 10&#160;Jul&#160;23&#160;11:48&#160;UTC | true | [view](CERTS/809b5296d9ca783497ea5d96de8caf84001ef4c3285b8447b0096462ef4e8aca/README.md) |
| 19&#160;Apr&#160;23&#160;06:07&#160;UTC | SHAKEN Telxio Networks 492K | 18&#160;Apr&#160;24&#160;06:07&#160;UTC | true | [view](CERTS/010c6a74330323c20ceb343b1de3a1e3248b4a3926c9ad2ed53f02b02399d241/README.md) |
| 19&#160;Apr&#160;23&#160;06:39&#160;UTC | SHAKEN Telxio Networks 492K | 18&#160;Apr&#160;24&#160;06:39&#160;UTC | true | [view](CERTS/1baf0e5fedd50fb55ff4e07366c5eb4f8d849b760739ffb8a0df4eb4828d7944/README.md) |
| 03&#160;May&#160;23&#160;03:19&#160;UTC | SHAKEN IDT America, Corp 363A | 02&#160;Jun&#160;23&#160;03:19&#160;UTC | true | [view](CERTS/660fdd210c02841fda06eb1b4305e1d961dc4bc77a720a37aa3fb7ce4a5ae89f/README.md) |
| 03&#160;May&#160;23&#160;05:27&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 02&#160;Jun&#160;23&#160;05:27&#160;UTC | true | [view](CERTS/49e78fa184e284fe92c8da5897495e93f14ed24b65df643fa33799841077575e/README.md) |
| 03&#160;May&#160;23&#160;07:57&#160;UTC | SHAKEN InteractiveTel, LLC 920J | 02&#160;Jun&#160;23&#160;07:57&#160;UTC | true | [view](CERTS/c64b672b455b96ec69413804ee4814c498595cd7a8cefaaf1906fc2824b12e02/README.md) |
| 03&#160;May&#160;23&#160;08:52&#160;UTC | SHAKEN BareTelecom 864J | 02&#160;Jun&#160;23&#160;08:52&#160;UTC | true | [view](CERTS/4c4e6fb3ea409fd3639d93e613060756cd67f9a16ecd3d6a1f5477cf6d5f9a43/README.md) |
| 03&#160;May&#160;23&#160;09:58&#160;UTC | SHAKEN Zray Technologies Corporation 862J | 02&#160;Jun&#160;23&#160;09:58&#160;UTC | true | [view](CERTS/39948f480a7087689f5898ab22ffd80dd8d96295e21a9a97e2e63a78688cee39/README.md) |
| 03&#160;May&#160;23&#160;20:05&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 02&#160;Jun&#160;23&#160;20:05&#160;UTC | true | [view](CERTS/7c0de139061e7f514abe257fb5d5d1821ea1bcab6034da7ca12ddf2e500c9393/README.md) |
| 04&#160;May&#160;23&#160;03:14&#160;UTC | SHAKEN IDT America, Corp 363A | 03&#160;Jun&#160;23&#160;03:14&#160;UTC | true | [view](CERTS/123c46f9f7fabb12d9b74a67368c9141e44f761620965827a603203d9c8dcdd6/README.md) |
| 04&#160;May&#160;23&#160;05:22&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 03&#160;Jun&#160;23&#160;05:22&#160;UTC | true | [view](CERTS/8cd686eeece552f024aa5f4a83aaf6c31e0e0dda6add967d4f002ed3f7332ed2/README.md) |
| 04&#160;May&#160;23&#160;06:00&#160;UTC | SHAKEN Convoso 758J | 08&#160;Jun&#160;23&#160;06:00&#160;UTC | true | [view](CERTS/76117f56c7f3978c79cf0b5fc5ad95d0e2059a067e491df9520965449f5e7a0d/README.md) |
| 04&#160;May&#160;23&#160;07:52&#160;UTC | SHAKEN InteractiveTel, LLC 920J | 03&#160;Jun&#160;23&#160;07:52&#160;UTC | true | [view](CERTS/bc6787ddc4e691419c4fdd6fede3cb00621f40c37470fb28b85e5b1601deaafb/README.md) |
| 04&#160;May&#160;23&#160;08:47&#160;UTC | SHAKEN BareTelecom 864J | 03&#160;Jun&#160;23&#160;08:47&#160;UTC | true | [view](CERTS/2316f8102561c30cc6b84d8fa0374fedf683b4e03040e056165a2232df94ba3c/README.md) |
| 04&#160;May&#160;23&#160;14:22&#160;UTC | SHAKEN NTC International, INC 016K | 03&#160;Jun&#160;23&#160;14:22&#160;UTC | true | [view](CERTS/45e1bf4e35e8a17d1f8b4d3c0b69a721089dbe285f4bc97ab1304e6116642d4f/README.md) |
| 04&#160;May&#160;23&#160;14:23&#160;UTC | SHAKEN i3 Broadband 5800 | 03&#160;Jun&#160;23&#160;14:23&#160;UTC | true | [view](CERTS/9f6c4670251ee9bc192d55616049bbf5ac79d7d269cef8888825be6a3dbb32eb/README.md) |
| 04&#160;May&#160;23&#160;20:00&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 03&#160;Jun&#160;23&#160;20:00&#160;UTC | true | [view](CERTS/c3d4c1c6b7f1b109e4a3167ecab224ba832f9529eeeb22c981a6fc8b59ac9e48/README.md) |
| 05&#160;May&#160;23&#160;03:09&#160;UTC | SHAKEN IDT America, Corp 363A | 04&#160;Jun&#160;23&#160;03:09&#160;UTC | true | [view](CERTS/2d5a41787720650ec579a7742199ef7b32ce8181ef82af77cb9e61d9c7dc27c4/README.md) |
| 05&#160;May&#160;23&#160;05:17&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 04&#160;Jun&#160;23&#160;05:17&#160;UTC | true | [view](CERTS/feca64b615ceb9e79e2b125cb4ef617c7b6038f26096736113b5eca23c53d601/README.md) |
| 05&#160;May&#160;23&#160;06:11&#160;UTC | SHAKEN IPSBS Managed Services LLC 828J | 04&#160;Jun&#160;23&#160;06:11&#160;UTC | true | [view](CERTS/9c62721a4ee633754c15dba7f498819a7bb1c315c82b01ebb8dfcaea0ebe5bb4/README.md) |
| 05&#160;May&#160;23&#160;08:42&#160;UTC | SHAKEN BareTelecom 864J | 04&#160;Jun&#160;23&#160;08:42&#160;UTC | true | [view](CERTS/e2a3ba8772e47c4e391f68ebb5747c6fc0fefde24545315e2ba9a45ced30ca91/README.md) |
| 06&#160;May&#160;23&#160;08:37&#160;UTC | SHAKEN BareTelecom 864J | 05&#160;Jun&#160;23&#160;08:37&#160;UTC | true | [view](CERTS/7cb2ddafd296f7dc5f2457db8dc2998ed5cdb8ce944deb21bd666a65374f6478/README.md) |
| 07&#160;May&#160;23&#160;02:59&#160;UTC | SHAKEN IDT America, Corp 363A | 06&#160;Jun&#160;23&#160;02:59&#160;UTC | true | [view](CERTS/58d51842cab1fd09084ec3f07c9b9d66a5810243c5251768c40e77f1aaf66947/README.md) |
| 07&#160;May&#160;23&#160;08:32&#160;UTC | SHAKEN BareTelecom 864J | 06&#160;Jun&#160;23&#160;08:32&#160;UTC | true | [view](CERTS/0d29f8af5bf271be2aa83670c9e986570f00cc890f97724f872597fdcbb3b772/README.md) |
| 07&#160;May&#160;23&#160;19:45&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 06&#160;Jun&#160;23&#160;19:45&#160;UTC | true | [view](CERTS/da9c991855490134073c9ce45df9a7b3b6b5bdf1a1602901e23dc15ebf6b2e5c/README.md) |
| 08&#160;May&#160;23&#160;02:54&#160;UTC | SHAKEN IDT America, Corp 363A | 07&#160;Jun&#160;23&#160;02:54&#160;UTC | true | [view](CERTS/882af374e30ea28b6fc9b8a1a0fb1cd7add2516fb1af441fb2b075f0cb0dd98f/README.md) |
| 08&#160;May&#160;23&#160;05:02&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 07&#160;Jun&#160;23&#160;05:02&#160;UTC | true | [view](CERTS/c225be6cc30962b2a01ed5721d5be7fb98af336892b331b002c54c9ffbea2381/README.md) |
| 08&#160;May&#160;23&#160;07:32&#160;UTC | SHAKEN InteractiveTel, LLC 920J | 07&#160;Jun&#160;23&#160;07:32&#160;UTC | true | [view](CERTS/760154f13fddd2193eae5f06818e9af4e67d8eaec3e59200b89f0c57a2b84b19/README.md) |
| 08&#160;May&#160;23&#160;08:27&#160;UTC | SHAKEN BareTelecom 864J | 07&#160;Jun&#160;23&#160;08:27&#160;UTC | true | [view](CERTS/c8d43fe7bc0c066be8d44386ef6132bc19320f8af46882a03149801fb61e14af/README.md) |
| 08&#160;May&#160;23&#160;19:40&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 07&#160;Jun&#160;23&#160;19:40&#160;UTC | true | [view](CERTS/a614d08817bfdf572935a9cc446279ab7f102dcb6e81df8f026dbdd293a2fc39/README.md) |
| 09&#160;May&#160;23&#160;02:49&#160;UTC | SHAKEN IDT America, Corp 363A | 08&#160;Jun&#160;23&#160;02:49&#160;UTC | true | [view](CERTS/e0c61bb12c93fd61af1bb88b33a29707132a04a2a68f0118a037def7d383c33d/README.md) |
| 09&#160;May&#160;23&#160;04:57&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 08&#160;Jun&#160;23&#160;04:57&#160;UTC | true | [view](CERTS/74a53ebcf921673292da165b0095d9a0db49550d497cf99630916d54570a30c5/README.md) |
| 09&#160;May&#160;23&#160;08:22&#160;UTC | SHAKEN BareTelecom 864J | 08&#160;Jun&#160;23&#160;08:22&#160;UTC | true | [view](CERTS/88080c0c81719b5271c4f30e4c1b309569c727ed7babd45a4e8065e86de4f4e9/README.md) |
| 09&#160;May&#160;23&#160;14:58&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 08&#160;Jun&#160;23&#160;14:58&#160;UTC | true | [view](CERTS/09af4439412a056358d649818a06656d55740b4fbab46afb7ea3843502ff4afe/README.md) |
| 09&#160;May&#160;23&#160;19:35&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 08&#160;Jun&#160;23&#160;19:35&#160;UTC | true | [view](CERTS/5fe13470e287df8a532644d688dc586f994f4bd9249574458fa15c1a3af8bb5b/README.md) |
| 10&#160;May&#160;23&#160;02:44&#160;UTC | SHAKEN IDT America, Corp 363A | 09&#160;Jun&#160;23&#160;02:44&#160;UTC | true | [view](CERTS/56b6d7727a6fc8cf76bcad853acdc3eb7edc7ec8d98e3f92ad34ea69a82c5601/README.md) |
| 10&#160;May&#160;23&#160;04:52&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 09&#160;Jun&#160;23&#160;04:52&#160;UTC | true | [view](CERTS/0be43527d39c08ba9c048fb415ced5b86b5247a73a6e59eac8aa0ea328a93c0b/README.md) |
| 10&#160;May&#160;23&#160;08:17&#160;UTC | SHAKEN BareTelecom 864J | 09&#160;Jun&#160;23&#160;08:17&#160;UTC | true | [view](CERTS/d78a55068c207a61deaf0340b73fe6be02fc3cad8f526c08e6e9457453d84b22/README.md) |
| 10&#160;May&#160;23&#160;14:53&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 09&#160;Jun&#160;23&#160;14:53&#160;UTC | true | [view](CERTS/622aa0b0cee3aa1479bea495744414965aa74412a8ed27af3ba8d64fde08d257/README.md) |
| 10&#160;May&#160;23&#160;19:30&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 09&#160;Jun&#160;23&#160;19:30&#160;UTC | true | [view](CERTS/0bda5f404d81006cd6e396aaf42dc5f6e83335dd8c47855f7ee21330b981798a/README.md) |
| 10&#160;May&#160;23&#160;23:35&#160;UTC | SHAKEN 1stPoint Communications, LLC 463G | 09&#160;Jun&#160;23&#160;23:35&#160;UTC | true | [view](CERTS/93c77e7e5ea502164b51520dcf90ea2649bf33d4d2dbc4d23c297780fc251831/README.md) |
| 11&#160;May&#160;23&#160;02:39&#160;UTC | SHAKEN IDT America, Corp 363A | 10&#160;Jun&#160;23&#160;02:39&#160;UTC | true | [view](CERTS/c389d0f56d910b8955b86bcad9cb7e3194a52236883140a2dea361b083a5cc27/README.md) |
| 11&#160;May&#160;23&#160;08:12&#160;UTC | SHAKEN BareTelecom 864J | 10&#160;Jun&#160;23&#160;08:12&#160;UTC | true | [view](CERTS/a8a6759f0ed508993bb4e4d726f05a0b43d59b7ca58501f67ae8acd236076845/README.md) |
| 11&#160;May&#160;23&#160;19:25&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 10&#160;Jun&#160;23&#160;19:25&#160;UTC | true | [view](CERTS/6ff4b1dea410e172be05c865cec361189c5da4b71896cde7cc53be875d34f6fb/README.md) |
| 12&#160;May&#160;23&#160;02:34&#160;UTC | SHAKEN IDT America, Corp 363A | 11&#160;Jun&#160;23&#160;02:34&#160;UTC | true | [view](CERTS/559c51ef6a729c33e7d356377c3ed79e7d36df580b22b4a34d945dec49b10880/README.md) |
| 12&#160;May&#160;23&#160;04:42&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 11&#160;Jun&#160;23&#160;04:42&#160;UTC | true | [view](CERTS/d6e6eb4f45bdb9ca32a87986aaaabcf70f5b8289e2385864ac100a90534af50e/README.md) |
| 12&#160;May&#160;23&#160;06:07&#160;UTC | SHAKEN InteractiveTel, LLC 920J | 11&#160;Jun&#160;23&#160;06:07&#160;UTC | true | [view](CERTS/3048eb4e5f8a5ad848eeccd4b1e40cd9af1b17652a0274ee40ecd96e823ff259/README.md) |
| 12&#160;May&#160;23&#160;08:07&#160;UTC | SHAKEN BareTelecom 864J | 11&#160;Jun&#160;23&#160;08:07&#160;UTC | true | [view](CERTS/39ab3c6c082c70658f2d1f3f8a1323ec188f5b00e89f9f8d695ebf2ed5fefd1f/README.md) |
| 12&#160;May&#160;23&#160;14:43&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 11&#160;Jun&#160;23&#160;14:43&#160;UTC | true | [view](CERTS/e079449f3150bc06252ea2b6d7ad336c03efbd6c75a0faaf0e76a5687bb41bea/README.md) |
| 12&#160;May&#160;23&#160;19:20&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 11&#160;Jun&#160;23&#160;19:20&#160;UTC | true | [view](CERTS/ad1123dfc180896dec26b47d990445a4d08e22c3bac82552cea07dd808df0f7f/README.md) |
| 13&#160;May&#160;23&#160;08:02&#160;UTC | SHAKEN BareTelecom 864J | 12&#160;Jun&#160;23&#160;08:02&#160;UTC | true | [view](CERTS/8c2d913c1bd98a0805b2b06e35bb7aa98f591b1b0a9a731d4c7369c85250438c/README.md) |
| 14&#160;May&#160;23&#160;07:57&#160;UTC | SHAKEN BareTelecom 864J | 13&#160;Jun&#160;23&#160;07:57&#160;UTC | true | [view](CERTS/db3559aaa11a3a981e6c754613bd9054303528966f47e53ce25fca6f578bb443/README.md) |
| 14&#160;May&#160;23&#160;14:33&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 13&#160;Jun&#160;23&#160;14:33&#160;UTC | true | [view](CERTS/8bd0352762c4cb1a562b60995e5de197935522247d1b6a82bfead7dcc172d45a/README.md) |
| 14&#160;May&#160;23&#160;19:10&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 13&#160;Jun&#160;23&#160;19:10&#160;UTC | true | [view](CERTS/54b62bcb30cb0101e7aa6ca2094bb76bc42ea4d890e2879da7d5c96bc33e748a/README.md) |
| 15&#160;May&#160;23&#160;04:27&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 14&#160;Jun&#160;23&#160;04:27&#160;UTC | true | [view](CERTS/7891857d1204d374e30eb9d472ea2a0098d2ef5c1b7efc7db664ca328cfe907b/README.md) |
| 15&#160;May&#160;23&#160;05:52&#160;UTC | SHAKEN InteractiveTel, LLC 920J | 14&#160;Jun&#160;23&#160;05:52&#160;UTC | true | [view](CERTS/c2c5d85dd5538cf714f010bc1d307121d58dd874f3e7229bf8396f3e7eb5f368/README.md) |
| 15&#160;May&#160;23&#160;07:52&#160;UTC | SHAKEN BareTelecom 864J | 14&#160;Jun&#160;23&#160;07:52&#160;UTC | true | [view](CERTS/91c7ac68e569fb9d6aa0858d056f34fade3e4b66bc6335448a464c0de3a9c6b0/README.md) |
| 15&#160;May&#160;23&#160;19:05&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 14&#160;Jun&#160;23&#160;19:05&#160;UTC | true | [view](CERTS/dd4e7457c588d844ee1afd934515ab9aa4cdcb52ee2ac04376ef8fafffc6afdf/README.md) |
| 16&#160;May&#160;23&#160;02:48&#160;UTC | SHAKEN Cyberlynk Network, LLC 086K | 15&#160;Jun&#160;23&#160;02:48&#160;UTC | true | [view](CERTS/a38c8f614fe3d1fe11ff71a5b32afd67d38cd43e71f0e9b5559e293ff8d67738/README.md) |
| 16&#160;May&#160;23&#160;04:22&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 15&#160;Jun&#160;23&#160;04:22&#160;UTC | true | [view](CERTS/644c97d37b010021cdd32396516de2614c4b720ae9b221358fbbd262103acbf4/README.md) |
| 16&#160;May&#160;23&#160;07:47&#160;UTC | SHAKEN BareTelecom 864J | 15&#160;Jun&#160;23&#160;07:47&#160;UTC | true | [view](CERTS/f54ee0727fa39e21493266005bd62e14b950d604ec2c52efca7514e5d1e4c3ba/README.md) |
| 17&#160;May&#160;23&#160;04:17&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 16&#160;Jun&#160;23&#160;04:17&#160;UTC | true | [view](CERTS/3b5962ac419ba9fcc823d1c8998333addc278c9126129ae90adaa30887daac33/README.md) |
| 17&#160;May&#160;23&#160;05:11&#160;UTC | SHAKEN IPSBS Managed Services LLC 828J | 16&#160;Jun&#160;23&#160;05:11&#160;UTC | true | [view](CERTS/d00cc55cdc988ec2551e0007552d1a708622252df216e466229c5876d50d75b8/README.md) |
| 17&#160;May&#160;23&#160;05:42&#160;UTC | SHAKEN InteractiveTel, LLC 920J | 16&#160;Jun&#160;23&#160;05:42&#160;UTC | true | [view](CERTS/b2ae0d03670d105ae8625b8dd03446ede2468bc96e24f23f5ea78fd9a45a1ccb/README.md) |
| 17&#160;May&#160;23&#160;07:42&#160;UTC | SHAKEN BareTelecom 864J | 16&#160;Jun&#160;23&#160;07:42&#160;UTC | true | [view](CERTS/854f797eb6c887fcda308deaf7673b2f636044e42509676adb15d6b4841ea450/README.md) |
| 17&#160;May&#160;23&#160;14:18&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 16&#160;Jun&#160;23&#160;14:18&#160;UTC | true | [view](CERTS/1dd04367a72884c639035dfa74b0b1f44a8530c223a8ca44703bcc3e1dd8fefd/README.md) |
| 17&#160;May&#160;23&#160;18:55&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 16&#160;Jun&#160;23&#160;18:55&#160;UTC | true | [view](CERTS/e41bdf20482d6d094ce604f1f20cea6e30a297258c2361e09fd151d4b50b0ec5/README.md) |
| 18&#160;May&#160;23&#160;00:00&#160;UTC | SHAKEN Primo Dialler LLC 249K | 14&#160;Nov&#160;23&#160;00:00&#160;UTC | true | [view](CERTS/a80051aad3dce4fd90a1e1f6758bff57994a1b1abc2e176244344c3a99c3e071/README.md) |
| 18&#160;May&#160;23&#160;01:12&#160;UTC | SHAKEN Cloud Connect LLC 455K | 17&#160;May&#160;24&#160;01:12&#160;UTC | true | [view](CERTS/0aa593ccacc13e85c2ec381274e47b597989c2a57173e248e25a91bc306c5f2c/README.md) |
| 18&#160;May&#160;23&#160;04:12&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 17&#160;Jun&#160;23&#160;04:12&#160;UTC | true | [view](CERTS/953000e8535871b4a28d8342d337889604359c1669f9a4fa212ba0abe9c1ec6c/README.md) |
| 18&#160;May&#160;23&#160;07:37&#160;UTC | SHAKEN BareTelecom 864J | 17&#160;Jun&#160;23&#160;07:37&#160;UTC | true | [view](CERTS/8eb9ec47008e905e476d8de68bf869d202f226006a6975dc59dcceba5c863ef7/README.md) |
| 18&#160;May&#160;23&#160;14:31&#160;UTC | SHAKEN ONE OWL TELECOM INC 412K | 17&#160;Jun&#160;23&#160;14:31&#160;UTC | true | [view](CERTS/e92fb7ca6b6bfed7ae054fca7b0aa907cea1a90ceebe787a3857cc095757c155/README.md) |
| 19&#160;May&#160;23&#160;07:32&#160;UTC | SHAKEN BareTelecom 864J | 18&#160;Jun&#160;23&#160;07:32&#160;UTC | true | [view](CERTS/67441a8c4e174b014fdee77b332bc561967481bb25be83170f476def5be1612d/README.md) |
| 20&#160;May&#160;23&#160;13:02&#160;UTC | SHAKEN NTC International, INC 016K | 19&#160;Jun&#160;23&#160;13:02&#160;UTC | true | [view](CERTS/9571006bce66d351104faf583c83d604826b1e4fcc1f3bf90c2b78d0020dfe24/README.md) |
| 22&#160;May&#160;23&#160;03:52&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 21&#160;Jun&#160;23&#160;03:52&#160;UTC | true | [view](CERTS/e54bd1590e008c2a5ff5800870f9522d210ecc4b955fa1d52fda9e80370a554b/README.md) |
| 23&#160;May&#160;23&#160;01:42&#160;UTC | SHAKEN IDT America, Corp 363A | 22&#160;Jun&#160;23&#160;01:42&#160;UTC | true | [view](CERTS/34e2360ef23f2dcb5f2ff7552493ab59d60f16807245c6c28e304df3074b09a1/README.md) |
| 23&#160;May&#160;23&#160;03:47&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 22&#160;Jun&#160;23&#160;03:47&#160;UTC | true | [view](CERTS/84e1043c7ec913e0f130cde768042b6ecdced5ba10fdc1d00ebe141319fe69ae/README.md) |
| 23&#160;May&#160;23&#160;04:41&#160;UTC | SHAKEN IPSBS Managed Services LLC 828J | 22&#160;Jun&#160;23&#160;04:41&#160;UTC | true | [view](CERTS/12242d570e21afa6d3f2347911cfb0e77de344003f7509dfba4d6060af88d43a/README.md) |
| 23&#160;May&#160;23&#160;13:48&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 22&#160;Jun&#160;23&#160;13:48&#160;UTC | true | [view](CERTS/77783d0c5aad2017ce5cfddce8265a4ca613debdebbdcdae9782369cd4752170/README.md) |
| 24&#160;May&#160;23&#160;01:37&#160;UTC | SHAKEN IDT America, Corp 363A | 23&#160;Jun&#160;23&#160;01:37&#160;UTC | true | [view](CERTS/abf7c60c99db16afa94b878cf0d48e7fbdf2ac5b291b5550cad2d75ea271a989/README.md) |
| 24&#160;May&#160;23&#160;03:42&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 23&#160;Jun&#160;23&#160;03:42&#160;UTC | true | [view](CERTS/2b2c3628dd2ebdf64d595f084040d99f9b15047fff8b9bf3a83b0b502795542e/README.md) |
| 24&#160;May&#160;23&#160;04:36&#160;UTC | SHAKEN IPSBS Managed Services LLC 828J | 23&#160;Jun&#160;23&#160;04:36&#160;UTC | true | [view](CERTS/1a5f75f6160eca7862252e8a53f46c5b1ceb92e40c294e77d84bd817fb837f94/README.md) |
| 24&#160;May&#160;23&#160;05:07&#160;UTC | SHAKEN InteractiveTel, LLC 920J | 23&#160;Jun&#160;23&#160;05:07&#160;UTC | true | [view](CERTS/e61e50dea7f9f1e43b703d277cc458899d28a6f876913bf65b60f513fdd4ee5e/README.md) |
| 24&#160;May&#160;23&#160;13:43&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 23&#160;Jun&#160;23&#160;13:43&#160;UTC | true | [view](CERTS/b05b690937b4e42941dd4120234ff34b498085e1769e69338c84d0871e0cc463/README.md) |
| 24&#160;May&#160;23&#160;14:40&#160;UTC | SHAKEN IDT America, Corp 363A | 23&#160;Jun&#160;23&#160;14:40&#160;UTC | true | [view](CERTS/e292adddacb3f692c705fcc7257771e0d223a90937ed1f2389fb4080a7c72d35/README.md) |
| 24&#160;May&#160;23&#160;14:42&#160;UTC | SHAKEN IDT America, Corp 363A | 23&#160;Jun&#160;23&#160;14:42&#160;UTC | true | [view](CERTS/81f24da3e67586ed46a80b0d2eabb0d98328f2a6e5a2838da1b6a4b9da7916dd/README.md) |
| 24&#160;May&#160;23&#160;15:09&#160;UTC | SHAKEN IDT America, Corp 363A | 23&#160;Jun&#160;23&#160;15:09&#160;UTC | true | [view](CERTS/a93b0793ed78f447138b66cc076cbabcaf9caddd999ca28d4cbd64a00bc44135/README.md) |
| 24&#160;May&#160;23&#160;15:31&#160;UTC | SHAKEN IDT America, Corp 363A | 23&#160;Jun&#160;23&#160;15:31&#160;UTC | true | [view](CERTS/3f5a4406d01b722d8a9ca9979d992eae1d3431a72b938e6515e2acdf438ec688/README.md) |
| 24&#160;May&#160;23&#160;15:40&#160;UTC | SHAKEN IDT America, Corp 363A | 23&#160;Jun&#160;23&#160;15:40&#160;UTC | true | [view](CERTS/581a2b525424ad67dce4c6777b5657be88c77ba0a57fc98f7d87228392307321/README.md) |
| 24&#160;May&#160;23&#160;15:42&#160;UTC | SHAKEN IDT America, Corp 363A | 23&#160;Jun&#160;23&#160;15:42&#160;UTC | true | [view](CERTS/49ef53d95315591cd25458bf50cc23fdb0ae1a7da5659bad85599f79877985e2/README.md) |
| 24&#160;May&#160;23&#160;18:20&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 23&#160;Jun&#160;23&#160;18:20&#160;UTC | true | [view](CERTS/f271fbb50c72239d644e6a4381d6f27c4dde9c97ae6be0ce6877c926ce371c87/README.md) |
| 25&#160;May&#160;23&#160;03:37&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 24&#160;Jun&#160;23&#160;03:37&#160;UTC | true | [view](CERTS/7a5e46a9bdc17a7a46af3d313b5c8684f3fcec25f0ff65b111e04a986fe73619/README.md) |
| 25&#160;May&#160;23&#160;05:02&#160;UTC | SHAKEN InteractiveTel, LLC 920J | 24&#160;Jun&#160;23&#160;05:02&#160;UTC | true | [view](CERTS/724c1b05a434efa0078e5225b92304bb9835cb73f1615216270f30fc968ff4df/README.md) |
| 25&#160;May&#160;23&#160;07:40&#160;UTC | SHAKEN IDT America, Corp 363A | 24&#160;Jun&#160;23&#160;07:40&#160;UTC | true | [view](CERTS/94a8b92406b0208298bed1b931f53d2bbdc8975cf930e0e46240b2930e2fe53c/README.md) |
| 25&#160;May&#160;23&#160;13:38&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 24&#160;Jun&#160;23&#160;13:38&#160;UTC | true | [view](CERTS/3030ad217ef39c6be8f2aebc0d8adc6f83a5499f239a8fc79c1c4c34650a6ab0/README.md) |
| 26&#160;May&#160;23&#160;03:32&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 25&#160;Jun&#160;23&#160;03:32&#160;UTC | true | [view](CERTS/569edfefeafcec92020510841c2f69112a7c2cf46865d9dd4946d25059f54a64/README.md) |
| 26&#160;May&#160;23&#160;04:57&#160;UTC | SHAKEN InteractiveTel, LLC 920J | 25&#160;Jun&#160;23&#160;04:57&#160;UTC | true | [view](CERTS/0df3308986fb3262eda853a8d6b6ebb5de7e85a8ff8937059d802e10120eb214/README.md) |
| 26&#160;May&#160;23&#160;12:32&#160;UTC | SHAKEN NTC International, INC 016K | 25&#160;Jun&#160;23&#160;12:32&#160;UTC | true | [view](CERTS/cfca0581d1e1461c864b4b8f2f0cd2c7361ce532fc598544b21e7fb95e6f7bf1/README.md) |
| 26&#160;May&#160;23&#160;13:21&#160;UTC | SHAKEN IDT America, Corp 363A | 25&#160;Jun&#160;23&#160;13:21&#160;UTC | true | [view](CERTS/cf8ec506966f455e0e0616101d7f1b5c73fc1369d3c2553872e0f520f365d899/README.md) |
| 26&#160;May&#160;23&#160;13:33&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 25&#160;Jun&#160;23&#160;13:33&#160;UTC | true | [view](CERTS/ffe2855eb987d67bd845ddf53abfde51c517d2669888bc4191e64c3776e0ed66/README.md) |
| 27&#160;May&#160;23&#160;06:53&#160;UTC | SHAKEN BareTelecom 864J | 26&#160;Jun&#160;23&#160;06:53&#160;UTC | true | [view](CERTS/edb9506791a734656fc207ec95f5f1f2e7b144c0a3f8885c61974d7ac962ca2b/README.md) |
| 27&#160;May&#160;23&#160;13:16&#160;UTC | SHAKEN IDT America, Corp 363A | 26&#160;Jun&#160;23&#160;13:16&#160;UTC | true | [view](CERTS/1004e3c98d8a6df49ede6e485c037fee94083d0817f2fc0cd1116985594d0864/README.md) |
| 27&#160;May&#160;23&#160;13:28&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 26&#160;Jun&#160;23&#160;13:28&#160;UTC | true | [view](CERTS/77489e18c17593372dc257c9305c725269cdce0066615dc5ef223c24e5cafedc/README.md) |
| 28&#160;May&#160;23&#160;06:48&#160;UTC | SHAKEN BareTelecom 864J | 27&#160;Jun&#160;23&#160;06:48&#160;UTC | true | [view](CERTS/15ce397519f5fb7a74f4e10a9edff0a9c250e797c740281de14ec233a563827a/README.md) |
| 28&#160;May&#160;23&#160;12:22&#160;UTC | SHAKEN NTC International, INC 016K | 27&#160;Jun&#160;23&#160;12:22&#160;UTC | true | [view](CERTS/8ca924637a76c7ebfe7b12d5204e8ac0f0f97300b3d38e7b309adebfea444871/README.md) |
| 29&#160;May&#160;23&#160;06:43&#160;UTC | SHAKEN BareTelecom 864J | 28&#160;Jun&#160;23&#160;06:43&#160;UTC | true | [view](CERTS/0d89e4d23216334192a591fe42b50f3bfed840f8648600e2dc7a468790cb6a19/README.md) |
| 29&#160;May&#160;23&#160;13:06&#160;UTC | SHAKEN IDT America, Corp 363A | 28&#160;Jun&#160;23&#160;13:06&#160;UTC | true | [view](CERTS/eb415b6ca53cfb8606cb95f331ae46e534ea8cb9600870c895569898c975ac30/README.md) |
| 30&#160;May&#160;23&#160;03:12&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 29&#160;Jun&#160;23&#160;03:12&#160;UTC | true | [view](CERTS/0eebdf403304dbc27ebab0760f66fdc45a1c0ebaefe6619d0b5a764977a0f178/README.md) |
| 30&#160;May&#160;23&#160;04:37&#160;UTC | SHAKEN InteractiveTel, LLC 920J | 29&#160;Jun&#160;23&#160;04:37&#160;UTC | true | [view](CERTS/c4266ca0c04111b40fb774dfa2af1e9b0ee04c30f87bd9bf736b8e0b3df153fc/README.md) |
| 30&#160;May&#160;23&#160;06:30&#160;UTC | SHAKEN  XCast Labs 689J | 30&#160;Jun&#160;23&#160;06:30&#160;UTC | true | [view](CERTS/7665c56a31f8ab444741351615ef1e00c96411f60829170f6db36a4a83f584ec/README.md) |
| 30&#160;May&#160;23&#160;06:38&#160;UTC | SHAKEN BareTelecom 864J | 29&#160;Jun&#160;23&#160;06:38&#160;UTC | true | [view](CERTS/a76efcbe7ad8cd2a9372051013d5d48d441aa97e43446c5dfabab18557548292/README.md) |
| 30&#160;May&#160;23&#160;13:01&#160;UTC | SHAKEN IDT America, Corp 363A | 29&#160;Jun&#160;23&#160;13:01&#160;UTC | true | [view](CERTS/d8118e61dfa496412464d1952a04db6185d072bca54edbe4b7ed2419b4741537/README.md) |
| 30&#160;May&#160;23&#160;14:03&#160;UTC | SHAKEN IDT America, Corp 363A | 29&#160;Jun&#160;23&#160;14:03&#160;UTC | true | [view](CERTS/c2bba3d9657c7f8a0b84316bf6f08a8998e6a36498b284cf25e733b093d5068f/README.md) |
| 30&#160;May&#160;23&#160;14:28&#160;UTC | SHAKEN IDT America, Corp 363A | 29&#160;Jun&#160;23&#160;14:28&#160;UTC | true | [view](CERTS/54dba873c2b0e92a5b228005f6ca9cdc9fcbcdf79502c26e903143a71d83eb33/README.md) |
| 30&#160;May&#160;23&#160;15:04&#160;UTC | SHAKEN IDT America, Corp 363A | 29&#160;Jun&#160;23&#160;15:04&#160;UTC | true | [view](CERTS/f1fa2d9a18faba86387efdf064b37b422d1d4098f90887cc52a5f0ea20b439d0/README.md) |
| 30&#160;May&#160;23&#160;15:26&#160;UTC | SHAKEN IDT America, Corp 363A | 29&#160;Jun&#160;23&#160;15:26&#160;UTC | true | [view](CERTS/c5a2945980d62778c01277d8ba9ee43efdd96c7feea8dcea13e3e3d6eeafa11e/README.md) |
| 30&#160;May&#160;23&#160;16:17&#160;UTC | SHAKEN IDT America, Corp 363A | 29&#160;Jun&#160;23&#160;16:17&#160;UTC | true | [view](CERTS/b7242ea1e60b6cd253d8e2668261723c6dd76839b6839c6855936b5f3a0cf496/README.md) |
| 30&#160;May&#160;23&#160;17:50&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 29&#160;Jun&#160;23&#160;17:50&#160;UTC | true | [view](CERTS/bd859eb198f90a60647e0608f6c6d8a16a4efc16bb45bd3aac296ea972388497/README.md) |
| 31&#160;May&#160;23&#160;03:07&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 30&#160;Jun&#160;23&#160;03:07&#160;UTC | true | [view](CERTS/17c330b4fbf350d72b6f28174da0d532a48b7aff12d2d599cded6c454327f5b1/README.md) |
| 31&#160;May&#160;23&#160;04:32&#160;UTC | SHAKEN InteractiveTel, LLC 920J | 30&#160;Jun&#160;23&#160;04:32&#160;UTC | true | [view](CERTS/68b0db51248945565213cdf77738ec7242eec7472f9c3fc3a26e759bbbda2ee5/README.md) |
| 31&#160;May&#160;23&#160;06:33&#160;UTC | SHAKEN BareTelecom 864J | 30&#160;Jun&#160;23&#160;06:33&#160;UTC | true | [view](CERTS/6c36bae41b433672163c43b67160be37f3e6031d65dcba8c8df8c80000e5e867/README.md) |
| 31&#160;May&#160;23&#160;11:51&#160;UTC | SHAKEN IDT America, Corp 363A | 30&#160;Jun&#160;23&#160;11:51&#160;UTC | true | [view](CERTS/ddd9dca339658b1ab563945a35e216c6be8b50f20ca3778844ca4125f7edebf6/README.md) |
| 31&#160;May&#160;23&#160;13:09&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 30&#160;Jun&#160;23&#160;13:09&#160;UTC | true | [view](CERTS/b2d4badd99100395a6cc96212739e5fff770eb5caf9039399e9eefab9b7a3b1b/README.md) |
| 31&#160;May&#160;23&#160;17:45&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 30&#160;Jun&#160;23&#160;17:45&#160;UTC | true | [view](CERTS/0ebabaea42201fab3ccfd0b27541bceb7c2410986154ab1153ea1ee8e704d772/README.md) |
| 01&#160;Jun&#160;23&#160;13:40&#160;UTC | SHAKEN Threshold Communications Inc 563J | 01&#160;Jul&#160;23&#160;13:40&#160;UTC | true | [view](CERTS/56fb6a85786144cf0c6236fb03ddfc9dd34ea46ae1164669a2f9baafa9d48b97/README.md) |
| 01&#160;Jun&#160;23&#160;13:40&#160;UTC | SHAKEN Broadband Dynamics LLC 583j | 01&#160;Jul&#160;23&#160;13:40&#160;UTC | true | [view](CERTS/43441993fdc3e8200987dd06ca7fca3c845b4c279ada5804009fcb839f5b7c47/README.md) |
| 01&#160;Jun&#160;23&#160;13:41&#160;UTC | SHAKEN Consolidated Communications 5113 | 01&#160;Jul&#160;23&#160;13:41&#160;UTC | true | [view](CERTS/aa1859b0753f7741cfd2ce51db84bf56c4814052ac92dc0b070c829674029a38/README.md) |
| 01&#160;Jun&#160;23&#160;13:41&#160;UTC | SHAKEN Touchtone 683A | 01&#160;Jul&#160;23&#160;13:41&#160;UTC | true | [view](CERTS/f22a28562017435da1da9126a37fde17570c0313973e8c2e0e166f13014b97dd/README.md) |
| 01&#160;Jun&#160;23&#160;13:41&#160;UTC | SHAKEN Apeiron Systems 012J | 01&#160;Jul&#160;23&#160;13:41&#160;UTC | true | [view](CERTS/2755187e92c5ba6eae1f3c06abf2a6a9f396ac9c78638ab3fa76a1221da6f786/README.md) |
| 01&#160;Jun&#160;23&#160;13:42&#160;UTC | SHAKEN Fonative, Inc. 684J | 01&#160;Jul&#160;23&#160;13:42&#160;UTC | true | [view](CERTS/0e6e5c193a279e8c5392416890ee5e7d1c7d6b3a2a4b4e9952d4564435171786/README.md) |
| 01&#160;Jun&#160;23&#160;13:42&#160;UTC | SHAKEN IPitomy 652J | 01&#160;Jul&#160;23&#160;13:42&#160;UTC | true | [view](CERTS/ba41c7dd1713ae95461a82fae955b87903c72992f86e4ee94dbe4e78e2a92895/README.md) |
| 01&#160;Jun&#160;23&#160;13:42&#160;UTC | SHAKEN Phone.com, Inc. 633J | 01&#160;Jul&#160;23&#160;13:42&#160;UTC | true | [view](CERTS/403ec7e2f77eac5c62e1ef1b1fcfa458e21c06b32b05d72212413512594aa9ce/README.md) |
| 01&#160;Jun&#160;23&#160;13:42&#160;UTC | SHAKEN NETRIO LLC 020K | 01&#160;Jul&#160;23&#160;13:42&#160;UTC | true | [view](CERTS/8b6bfce71811fce0d3056a0261a7b62ed269da58d84f25bf9b5dbf34b60daafa/README.md) |
| 01&#160;Jun&#160;23&#160;13:43&#160;UTC | SHAKEN Noble Systems Communications LLC 187J | 01&#160;Jul&#160;23&#160;13:43&#160;UTC | true | [view](CERTS/79c53deb7dac170bf1d626cb5f9bdbe42560358d1ca948b207797940db5f651d/README.md) |
| 01&#160;Jun&#160;23&#160;13:43&#160;UTC | SHAKEN Telcentris Inc. dba Voxox 696J | 01&#160;Jul&#160;23&#160;13:43&#160;UTC | true | [view](CERTS/c5f21a01ad0a32298e28e1947c5b0ab8a868baf674fa9874139fc5fffa795d82/README.md) |
| 01&#160;Jun&#160;23&#160;13:44&#160;UTC | SHAKEN Airespring 996H | 01&#160;Jul&#160;23&#160;13:44&#160;UTC | true | [view](CERTS/b45b3decc2ee724e86ccd716519fdf2f729d88072ea2d6f390a88d169bdc9b8b/README.md) |
| 01&#160;Jun&#160;23&#160;13:44&#160;UTC | SHAKEN Nobelbiz, Inc. 596J | 01&#160;Jul&#160;23&#160;13:44&#160;UTC | true | [view](CERTS/49113ba0f82e19643449d1a7f6f600e851e00e4f9fde8d7abea56ed7a3ab7128/README.md) |
| 01&#160;Jun&#160;23&#160;13:45&#160;UTC | SHAKEN Momentum Telecom 1417 | 01&#160;Jul&#160;23&#160;13:45&#160;UTC | true | [view](CERTS/e5b63b0c7d237325c046554e38204fc4e948fe2ebd7166475eab8967f0ae8566/README.md) |
| 01&#160;Jun&#160;23&#160;13:46&#160;UTC | SHAKEN Momentum Telecom 9157 | 01&#160;Jul&#160;23&#160;13:46&#160;UTC | true | [view](CERTS/c3d61affd2970fe8fa0a188265c9900a844ee9908f7a9f1263e8d70eb84ccdde/README.md) |
| 01&#160;Jun&#160;23&#160;13:48&#160;UTC | SHAKEN Terra Nova Telecom 382G | 01&#160;Jul&#160;23&#160;13:48&#160;UTC | true | [view](CERTS/0772e298d057d313a382a18c41614d58496b6b100029a5e81a420b7a40676843/README.md) |
| 01&#160;Jun&#160;23&#160;13:49&#160;UTC | SHAKEN Matrix 3058 | 01&#160;Jul&#160;23&#160;13:49&#160;UTC | true | [view](CERTS/3c5064ba3702315507ce3e9aa2b7931d900900b720a614ee4c93b45a1ab9c03d/README.md) |
| 01&#160;Jun&#160;23&#160;13:49&#160;UTC | SHAKEN Matrix 9451 | 01&#160;Jul&#160;23&#160;13:49&#160;UTC | true | [view](CERTS/6f7a865bde3de4e4c7dc86a2a7b6599a8e90f15810cdbb07b0b5b1b7690d2ed7/README.md) |
| 01&#160;Jun&#160;23&#160;13:50&#160;UTC | SHAKEN Matrix 7379 | 01&#160;Jul&#160;23&#160;13:50&#160;UTC | true | [view](CERTS/e1f97153a4754863515537607aa7f0ff33ac4330668ca40da12fa314c1a8442c/README.md) |
| 01&#160;Jun&#160;23&#160;13:51&#160;UTC | SHAKEN PNG Telecommunications Inc 3395 | 01&#160;Jul&#160;23&#160;13:51&#160;UTC | true | [view](CERTS/f720a2ca6997c52c42215cb56cd592d9caf4322b30cb39837d3c041164ce2b94/README.md) |
| 01&#160;Jun&#160;23&#160;13:54&#160;UTC | SHAKEN Magna5, LLC 3849 | 01&#160;Jul&#160;23&#160;13:54&#160;UTC | true | [view](CERTS/747bd838405635ea75c116785c82e95e1a32c92207d2d21f804d4823d326b198/README.md) |
| 01&#160;Jun&#160;23&#160;13:57&#160;UTC | SHAKEN Magna5, LLC 8249 | 01&#160;Jul&#160;23&#160;13:57&#160;UTC | true | [view](CERTS/e1176e75a53615195c0e9d442bb509ea623da87ce540d7e2df1ab0c287a5b8cb/README.md) |

#### CA Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 21&#160;Aug&#160;20&#160;01:22&#160;UTC | SHAKEN Sansay Root CA US | 16&#160;Aug&#160;40&#160;01:22&#160;UTC | true | [view](CERTS/8356d251b255f4ac3fd108bb79be4c02dcea2d3b13d138892bdb3a70cbe6a343/README.md) |
| 02&#160;Sep&#160;22&#160;20:53&#160;UTC | SHAKEN Sansay Intermediate CA US WEST 1 | 31&#160;Aug&#160;29&#160;20:53&#160;UTC | true | [view](CERTS/4b1dfdba2b1e4bbffbf900a20f1f6f7befbef0008b963e4922a64469cb97d24b/README.md) |


Generated: 02 Jun 23 01:12 UTC