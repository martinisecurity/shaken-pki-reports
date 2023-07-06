# STIR/SHAKEN CA Ecosystem Compliance

## Sansay

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 1036 certificates were included in the corpus being tested
- 14 certificates in the corpus were skipped because they are duplicates
- 793 certificates in the corpus were skipped because they are expired
- 1 certificates in the corpus were skipped because they are not currently trusted
- 228 certificates being tested against the remaining rules
- 4.92 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 100.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 98 days is the average remaining validity for the certificates in the corpus
- 98 days is the average initial validity for the certificates in the corpus
- 178 certificates expire in the next 30 days
- 2.75 average number of unexpired certificates per OCN observed
- 83 unique OCNs observed in unexpired and valid certificate corpus

| Instances | Test | Source |
|-----------|------|--------|
| 56 | [e_atis_certificate_policies](ISSUES/e_atis_certificate_policies/README.md) | ATIS1000080 |
| 228 | [e_atis_subject_cn](ISSUES/e_atis_subject_cn/README.md) | ATIS1000080 |
| 228 | [e_us_cp_ambiguous_identifier](ISSUES/e_us_cp_ambiguous_identifier/README.md) | US_SHAKEN_CP |
| 228 | [e_us_cp_subject_sn](ISSUES/e_us_cp_subject_sn/README.md) | US_SHAKEN_CP |
| 153 | [w_cp_1_3_subject_email](ISSUES/w_cp_1_3_subject_email/README.md) | US_SHAKEN_CP |
| 228 | [w_shaken_subject_rdn_unknown](ISSUES/w_shaken_subject_rdn_unknown/README.md) | SHAKEN_PKI_BEST_PRACTICES |

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
- 5369 days is the average remaining validity for the certificates in the corpus
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
| 04&#160;Apr&#160;23&#160;16:38&#160;UTC | SHAKEN Inventive Labs Corp 649J | 01&#160;Oct&#160;23&#160;16:38&#160;UTC | true | [view](CERTS/61bd59e51cd88d19ca695a362569930dee7725fefa39299b507bc3269fefbd54/README.md) |
| 05&#160;Apr&#160;23&#160;16:27&#160;UTC | SHAKEN Swift Telco LLC 452K | 04&#160;Apr&#160;24&#160;16:27&#160;UTC | true | [view](CERTS/947b46067a639b79ff82ab3f48c453e4af7cc6d6036f6d66a742cc935bc8a35e/README.md) |
| 11&#160;Apr&#160;23&#160;11:48&#160;UTC | SHAKEN IP Link Telecom Inc. 902J | 10&#160;Jul&#160;23&#160;11:48&#160;UTC | true | [view](CERTS/809b5296d9ca783497ea5d96de8caf84001ef4c3285b8447b0096462ef4e8aca/README.md) |
| 19&#160;Apr&#160;23&#160;06:07&#160;UTC | SHAKEN Telxio Networks 492K | 18&#160;Apr&#160;24&#160;06:07&#160;UTC | true | [view](CERTS/010c6a74330323c20ceb343b1de3a1e3248b4a3926c9ad2ed53f02b02399d241/README.md) |
| 19&#160;Apr&#160;23&#160;06:39&#160;UTC | SHAKEN Telxio Networks 492K | 18&#160;Apr&#160;24&#160;06:39&#160;UTC | true | [view](CERTS/1baf0e5fedd50fb55ff4e07366c5eb4f8d849b760739ffb8a0df4eb4828d7944/README.md) |
| 25&#160;Apr&#160;23&#160;16:43&#160;UTC | SHAKEN GIP Technology 434K | 24&#160;Apr&#160;24&#160;16:43&#160;UTC | true | [view](CERTS/bc5cd573f3ab46daf12994739622514f37ac8cd3275ff8d493595fee747e8a0b/README.md) |
| 18&#160;May&#160;23&#160;00:00&#160;UTC | SHAKEN Primo Dialler LLC 249K | 14&#160;Nov&#160;23&#160;00:00&#160;UTC | true | [view](CERTS/a80051aad3dce4fd90a1e1f6758bff57994a1b1abc2e176244344c3a99c3e071/README.md) |
| 18&#160;May&#160;23&#160;01:12&#160;UTC | SHAKEN Cloud Connect LLC 455K | 17&#160;May&#160;24&#160;01:12&#160;UTC | true | [view](CERTS/0aa593ccacc13e85c2ec381274e47b597989c2a57173e248e25a91bc306c5f2c/README.md) |
| 24&#160;May&#160;23&#160;14:49&#160;UTC | SHAKEN Nextiva, Inc 914H | 23&#160;May&#160;24&#160;14:49&#160;UTC | true | [view](CERTS/ef77b45b37fa412f05b53660a4b60f37962a9be4587ec39211e94289a0087c20/README.md) |
| 01&#160;Jun&#160;23&#160;19:33&#160;UTC | SHAKEN Medtel Communications 994J | 31&#160;May&#160;24&#160;19:33&#160;UTC | true | [view](CERTS/e4b9231f5ad017174e74400e163effa66767d53af72cf608ba3046a0e639a813/README.md) |
| 06&#160;Jun&#160;23&#160;17:15&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 06&#160;Jul&#160;23&#160;17:15&#160;UTC | true | [view](CERTS/a3f9b977fcaf65b6fc4de36f43870adf6bf5f021574ff01ec0acb72068dd9b25/README.md) |
| 07&#160;Jun&#160;23&#160;02:32&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 07&#160;Jul&#160;23&#160;02:32&#160;UTC | true | [view](CERTS/df4a3e2fd38c7d7d12260739dcc0b368f44f155aebb023481b150eec3b3d05bd/README.md) |
| 07&#160;Jun&#160;23&#160;03:57&#160;UTC | SHAKEN InteractiveTel, LLC 920J | 07&#160;Jul&#160;23&#160;03:57&#160;UTC | true | [view](CERTS/a09ffbc96ba0451d099f99ef29a8e0e5a6cf6cc90e384094575385697af4874e/README.md) |
| 07&#160;Jun&#160;23&#160;05:58&#160;UTC | SHAKEN BareTelecom 864J | 07&#160;Jul&#160;23&#160;05:58&#160;UTC | true | [view](CERTS/82e206fcc997ad1ed67173ddc171437ec13d696179105d3428c17ec4a84bb30e/README.md) |
| 07&#160;Jun&#160;23&#160;12:33&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 07&#160;Jul&#160;23&#160;12:33&#160;UTC | true | [view](CERTS/3248ec6caab91e9491a26fb9ae559db31d2d833bbce7735c32a305d8a23d91f6/README.md) |
| 07&#160;Jun&#160;23&#160;12:45&#160;UTC | SHAKEN IDT America, Corp 363A | 07&#160;Jul&#160;23&#160;12:45&#160;UTC | true | [view](CERTS/c9c8d5e5566e30b1e132d614eb2a92306441eb76b16945ee855ec98a7e731efb/README.md) |
| 07&#160;Jun&#160;23&#160;17:10&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 07&#160;Jul&#160;23&#160;17:10&#160;UTC | true | [view](CERTS/1e7d3588b50485c9f432d626399efd12ac01c464cb9594dc12a1fd31dfa90105/README.md) |
| 08&#160;Jun&#160;23&#160;02:27&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 08&#160;Jul&#160;23&#160;02:27&#160;UTC | true | [view](CERTS/60bd7997db8b726d7f5e47140ccb561927dd20b3f2c3f6d7a5b4e8415470ad3a/README.md) |
| 08&#160;Jun&#160;23&#160;03:52&#160;UTC | SHAKEN InteractiveTel, LLC 920J | 08&#160;Jul&#160;23&#160;03:52&#160;UTC | true | [view](CERTS/b543b0fb66048a153899781f2c2c846a3c215219aa6ac9239911824065472830/README.md) |
| 08&#160;Jun&#160;23&#160;05:53&#160;UTC | SHAKEN BareTelecom 864J | 08&#160;Jul&#160;23&#160;05:53&#160;UTC | true | [view](CERTS/403b7c719b71e5bfd9333a0be4e5bcb47189759067cfc93186f8602a52337a0a/README.md) |
| 08&#160;Jun&#160;23&#160;12:28&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 08&#160;Jul&#160;23&#160;12:28&#160;UTC | true | [view](CERTS/4a2bb1b488fc7efa5c923e594570310d860fc845b9d7a6056dd8ed4c3c97f7e4/README.md) |
| 08&#160;Jun&#160;23&#160;12:40&#160;UTC | SHAKEN IDT America, Corp 363A | 08&#160;Jul&#160;23&#160;12:40&#160;UTC | true | [view](CERTS/e6d2d882077cb330cc6ae359206027feed0db3b942597d5106ff904a8d545f56/README.md) |
| 08&#160;Jun&#160;23&#160;17:05&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 08&#160;Jul&#160;23&#160;17:05&#160;UTC | true | [view](CERTS/b2f655eafccfbc2fd9621a6318998d6edfcfc3d658b0943875de58e8eb8949b2/README.md) |
| 09&#160;Jun&#160;23&#160;02:22&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 09&#160;Jul&#160;23&#160;02:22&#160;UTC | true | [view](CERTS/025f3421a2eced3a8f67243245844b549f8aa4df027f2c6e232f87ff3009ecf2/README.md) |
| 09&#160;Jun&#160;23&#160;02:54&#160;UTC | SHAKEN IPSBS Managed Services LLC 828J | 09&#160;Jul&#160;23&#160;02:54&#160;UTC | true | [view](CERTS/e1acc11f0c63aeb48a2f447a741e7ec2a8c977426a3c317eca5f4af51036df3e/README.md) |
| 09&#160;Jun&#160;23&#160;05:48&#160;UTC | SHAKEN BareTelecom 864J | 09&#160;Jul&#160;23&#160;05:48&#160;UTC | true | [view](CERTS/b8d67a760ac55eab493d5eb8fe5eabdb80a04eb5816fbce5c6e319513055c8e1/README.md) |
| 09&#160;Jun&#160;23&#160;08:52&#160;UTC | SHAKEN Townes Telecommunications 0335 | 09&#160;Jul&#160;23&#160;08:52&#160;UTC | true | [view](CERTS/3d7486360ecd41826dc2973378e1265ce468a89f2bd55ba409ae8832778f146d/README.md) |
| 09&#160;Jun&#160;23&#160;17:00&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 09&#160;Jul&#160;23&#160;17:00&#160;UTC | true | [view](CERTS/1d7efcda82f2796bcd0389cd264de2f46f34187ee7754f1e8581778b55a05119/README.md) |
| 10&#160;Jun&#160;23&#160;02:17&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 10&#160;Jul&#160;23&#160;02:17&#160;UTC | true | [view](CERTS/1c3c12ec075ffcb215f25064239a352cd88067f9fe01830e1defee9de923eb80/README.md) |
| 10&#160;Jun&#160;23&#160;05:43&#160;UTC | SHAKEN BareTelecom 864J | 10&#160;Jul&#160;23&#160;05:43&#160;UTC | true | [view](CERTS/81afbcc18dc6fe39dafeacff14cf59ad1d360b48002c5c802a4678cfda09bbd6/README.md) |
| 10&#160;Jun&#160;23&#160;12:18&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 10&#160;Jul&#160;23&#160;12:18&#160;UTC | true | [view](CERTS/1930aef5ddc2781dd48f71dbf73b250593d96bc5bf82269adcaeee81a07be747/README.md) |
| 10&#160;Jun&#160;23&#160;16:55&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 10&#160;Jul&#160;23&#160;16:55&#160;UTC | true | [view](CERTS/8a2ab60a81cce3c36d5cfacebfcf1867083b5d298ed2019b1e3eec4c1860dba8/README.md) |
| 11&#160;Jun&#160;23&#160;05:38&#160;UTC | SHAKEN BareTelecom 864J | 11&#160;Jul&#160;23&#160;05:38&#160;UTC | true | [view](CERTS/ee2f33ae792d4dcca1e3a23955873cb0fdf10ac73c50645372ab220b4d6161b3/README.md) |
| 11&#160;Jun&#160;23&#160;11:12&#160;UTC | SHAKEN NTC International, INC 016K | 11&#160;Jul&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/7d48fec2e964e4e26468e4850dc6a4faf58df8d45460e874354c52f0aca95859/README.md) |
| 11&#160;Jun&#160;23&#160;16:50&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 11&#160;Jul&#160;23&#160;16:50&#160;UTC | true | [view](CERTS/819a5b2a42e1fe6ae2150e3bfb15eee41e3d5412ae596628454532d73b5d367f/README.md) |
| 12&#160;Jun&#160;23&#160;02:07&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 12&#160;Jul&#160;23&#160;02:07&#160;UTC | true | [view](CERTS/b13571675a967d3328e68d6e4218a30239d9a220c29207b230cb5681ad9118c8/README.md) |
| 12&#160;Jun&#160;23&#160;05:33&#160;UTC | SHAKEN BareTelecom 864J | 12&#160;Jul&#160;23&#160;05:33&#160;UTC | true | [view](CERTS/dd5e85f306206fb7e0fbdaf7063572e03631134bd2d697fe79ecf55aab90b746/README.md) |
| 12&#160;Jun&#160;23&#160;12:08&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 12&#160;Jul&#160;23&#160;12:08&#160;UTC | true | [view](CERTS/e598d1ce0408123416f7dea05535bc724c17fbff0a76f1657712c257e400400c/README.md) |
| 12&#160;Jun&#160;23&#160;16:45&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 12&#160;Jul&#160;23&#160;16:45&#160;UTC | true | [view](CERTS/5d600f94723e4be177e06ae4c6f534d906eb0db5b7ca3ebc717b7eb5e64f07c9/README.md) |
| 12&#160;Jun&#160;23&#160;21:38&#160;UTC | SHAKEN IDT America, Corp 363A | 12&#160;Jul&#160;23&#160;21:38&#160;UTC | true | [view](CERTS/8b815a23add434d6471a05c482f361ce81ee37fe6b190ccbd361f32ea9abb032/README.md) |
| 13&#160;Jun&#160;23&#160;02:02&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 13&#160;Jul&#160;23&#160;02:02&#160;UTC | true | [view](CERTS/0ce6e0da6a4eac15030a012aca608ac834acd6dc32482853e60cc5d6bf909937/README.md) |
| 13&#160;Jun&#160;23&#160;02:34&#160;UTC | SHAKEN IPSBS Managed Services LLC 828J | 13&#160;Jul&#160;23&#160;02:34&#160;UTC | true | [view](CERTS/ada2b12ef11fd8b5aa08f3f3f742098cf70c8fa603831e07b0e1ab2a80348ede/README.md) |
| 13&#160;Jun&#160;23&#160;05:28&#160;UTC | SHAKEN BareTelecom 864J | 13&#160;Jul&#160;23&#160;05:28&#160;UTC | true | [view](CERTS/3e406b7a54d275a9dcd5093fe5e03c535de00c9919ede7679c9f14c489dab8c7/README.md) |
| 13&#160;Jun&#160;23&#160;12:03&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 13&#160;Jul&#160;23&#160;12:03&#160;UTC | true | [view](CERTS/d19e414259085cf40050818a7c45a1e3778ea913f4aa3083f355638eeabecd51/README.md) |
| 13&#160;Jun&#160;23&#160;16:40&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 13&#160;Jul&#160;23&#160;16:40&#160;UTC | true | [view](CERTS/5ba18d2a4e9454ee9a97e19b195736d0065386b3debf0f2f3673721dbc358240/README.md) |
| 13&#160;Jun&#160;23&#160;21:33&#160;UTC | SHAKEN IDT America, Corp 363A | 13&#160;Jul&#160;23&#160;21:33&#160;UTC | true | [view](CERTS/74ff91f6a29bea8c88fcf1c0504de1394d1ff69fdfea0a2fd1a5f9f6ed9e2def/README.md) |
| 14&#160;Jun&#160;23&#160;01:57&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 14&#160;Jul&#160;23&#160;01:57&#160;UTC | true | [view](CERTS/aabeff22f140d2e5e95cefa976a2ac1aa38b680fa7235f68b1658e2ea0b0804d/README.md) |
| 14&#160;Jun&#160;23&#160;02:29&#160;UTC | SHAKEN IPSBS Managed Services LLC 828J | 14&#160;Jul&#160;23&#160;02:29&#160;UTC | true | [view](CERTS/bb347d925f0e573bd51f40c9fb3fc6b9c20e56f2a1c88a5e02e22981a0a9b52c/README.md) |
| 14&#160;Jun&#160;23&#160;05:23&#160;UTC | SHAKEN BareTelecom 864J | 14&#160;Jul&#160;23&#160;05:23&#160;UTC | true | [view](CERTS/93cef77760be2facc1ada16386058c2d998c11148c06d3c1750d39bcfe91071c/README.md) |
| 14&#160;Jun&#160;23&#160;11:59&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 14&#160;Jul&#160;23&#160;11:59&#160;UTC | true | [view](CERTS/85f64620014236f5e60a4c7fd402b4437cf58604cdee054324e569bfe01d8ae6/README.md) |
| 14&#160;Jun&#160;23&#160;13:44&#160;UTC | SHAKEN Socket Telecom LLC 554a | 14&#160;Jul&#160;23&#160;13:44&#160;UTC | true | [view](CERTS/206e4a375b0eea9e54586c4c99e86a280e5bf84733d01fb36e0a26324a9f0595/README.md) |
| 14&#160;Jun&#160;23&#160;16:35&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 14&#160;Jul&#160;23&#160;16:35&#160;UTC | true | [view](CERTS/c3127c81983ddd602fa370bcd1f0c7e1184232cc2cd45259d2e9f64f8d1f6c89/README.md) |
| 14&#160;Jun&#160;23&#160;21:29&#160;UTC | SHAKEN IDT America, Corp 363A | 14&#160;Jul&#160;23&#160;21:29&#160;UTC | true | [view](CERTS/c242078f7bcbfcde8fe79a823630179fea2c1a4e7450c8d7490377b24f784ec6/README.md) |
| 15&#160;Jun&#160;23&#160;02:24&#160;UTC | SHAKEN IPSBS Managed Services LLC 828J | 15&#160;Jul&#160;23&#160;02:24&#160;UTC | true | [view](CERTS/c62c77afc998f00943cc7b3fb03a61b9733277ccf4670be7eb2939142ddabef3/README.md) |
| 15&#160;Jun&#160;23&#160;03:17&#160;UTC | SHAKEN InteractiveTel, LLC 920J | 15&#160;Jul&#160;23&#160;03:17&#160;UTC | true | [view](CERTS/2d3405e3638377425eec627f760f3c4c98fcac8859d6ef594d0cb4bfd3d49da8/README.md) |
| 15&#160;Jun&#160;23&#160;05:18&#160;UTC | SHAKEN BareTelecom 864J | 15&#160;Jul&#160;23&#160;05:18&#160;UTC | true | [view](CERTS/a84c6c93c6635dca80c4cbcde9ea5d49ee6f26bdb57ce25d386256c3495aa2aa/README.md) |
| 15&#160;Jun&#160;23&#160;11:54&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 15&#160;Jul&#160;23&#160;11:54&#160;UTC | true | [view](CERTS/c74dd8297d286359c678721ac2ad7b053111b7221bec7e4df3a66de4b94ad1cd/README.md) |
| 15&#160;Jun&#160;23&#160;16:30&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 15&#160;Jul&#160;23&#160;16:30&#160;UTC | true | [view](CERTS/ce3a9af3a62857b46c20e9d20d8ed5ce18af8c06a0abd236d2e8fee60954ef5a/README.md) |
| 15&#160;Jun&#160;23&#160;18:48&#160;UTC | SHAKEN Zray Technologies Corporation 862J | 15&#160;Jul&#160;23&#160;18:48&#160;UTC | true | [view](CERTS/9915012e6a8142d6ff74561ec947f7ee84b3da85f02c47f6c088f6cba7d1769a/README.md) |
| 16&#160;Jun&#160;23&#160;03:12&#160;UTC | SHAKEN InteractiveTel, LLC 920J | 16&#160;Jul&#160;23&#160;03:12&#160;UTC | true | [view](CERTS/0818e26c567af30bbd67bb6576a758baa93798b7a013a5052d58a07601dc626b/README.md) |
| 16&#160;Jun&#160;23&#160;05:13&#160;UTC | SHAKEN BareTelecom 864J | 16&#160;Jul&#160;23&#160;05:13&#160;UTC | true | [view](CERTS/4a03e99e67c7b7e1fe850fd5edecaade15f2c904ccbc66a0ffca4ada7f04f50b/README.md) |
| 16&#160;Jun&#160;23&#160;13:00&#160;UTC | SHAKEN IDT America, Corp 363A | 16&#160;Jul&#160;23&#160;13:00&#160;UTC | true | [view](CERTS/00d79330f4867ff5d68fec65b030b86c6742fbad317f539643e1d25b769c9d07/README.md) |
| 16&#160;Jun&#160;23&#160;16:25&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 16&#160;Jul&#160;23&#160;16:25&#160;UTC | true | [view](CERTS/b58af7f86301499f77299d01e161f92cb3d9d8eebb8f38a87c00927045b99c4b/README.md) |
| 17&#160;Jun&#160;23&#160;05:08&#160;UTC | SHAKEN BareTelecom 864J | 17&#160;Jul&#160;23&#160;05:08&#160;UTC | true | [view](CERTS/2870690357c98a2881c9703199cbea52644afe5dea4e54a9cb556ee23313b806/README.md) |
| 17&#160;Jun&#160;23&#160;12:55&#160;UTC | SHAKEN IDT America, Corp 363A | 17&#160;Jul&#160;23&#160;12:55&#160;UTC | true | [view](CERTS/bc8e165e5803ffff06db022be3acc99750c4dfdbccc3058c9b69f0075f4db24b/README.md) |
| 17&#160;Jun&#160;23&#160;16:20&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 17&#160;Jul&#160;23&#160;16:20&#160;UTC | true | [view](CERTS/ff19544f1be779d382d0a9430bd3f1071972bb485018e89a064681c78f8512b8/README.md) |
| 18&#160;Jun&#160;23&#160;05:03&#160;UTC | SHAKEN BareTelecom 864J | 18&#160;Jul&#160;23&#160;05:03&#160;UTC | true | [view](CERTS/86b26f998d56433b62e075619ef794d01853f8a3d49a8f6b4c386ab9da8c457d/README.md) |
| 18&#160;Jun&#160;23&#160;16:15&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 18&#160;Jul&#160;23&#160;16:15&#160;UTC | true | [view](CERTS/8fd184b9dc44bd44ec8d586be9847ba94b87716994fcd831537e9b17b90d4e36/README.md) |
| 19&#160;Jun&#160;23&#160;00:59&#160;UTC | SHAKEN Doylestown Communications, Inc 0609 | 19&#160;Jul&#160;23&#160;00:59&#160;UTC | true | [view](CERTS/53fded94e17ee34c59f0d9ed3c36feac00a7c4c3a74fa2ca12b2cc4cb55f8158/README.md) |
| 19&#160;Jun&#160;23&#160;02:04&#160;UTC | SHAKEN IPSBS Managed Services LLC 828J | 19&#160;Jul&#160;23&#160;02:04&#160;UTC | true | [view](CERTS/1c9df2591cf0b6a4ebd77e46a30a9fe2038de4fcbcf04de129a3043b86e50d39/README.md) |
| 19&#160;Jun&#160;23&#160;04:58&#160;UTC | SHAKEN BareTelecom 864J | 19&#160;Jul&#160;23&#160;04:58&#160;UTC | true | [view](CERTS/215d4c0b89ea921da287abd7bed0183f24e783f2873a8d3cdc48d311fef246f6/README.md) |
| 19&#160;Jun&#160;23&#160;08:29&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 19&#160;Jul&#160;23&#160;08:29&#160;UTC | true | [view](CERTS/ca043df5dab8e4bd17903b9e96d77da1912d3467aa014f133fa2812b769f3e7c/README.md) |
| 19&#160;Jun&#160;23&#160;11:34&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 19&#160;Jul&#160;23&#160;11:34&#160;UTC | true | [view](CERTS/57e76ae050744ce76f111e67ba7bdfad181e2c6ff3a5df655fb14d697c7bd0c1/README.md) |
| 19&#160;Jun&#160;23&#160;12:45&#160;UTC | SHAKEN IDT America, Corp 363A | 19&#160;Jul&#160;23&#160;12:45&#160;UTC | true | [view](CERTS/8ef32b355adcb38b93d88a09ea2ac51dba5f12fb8aefa36c1d75941bb7ba383d/README.md) |
| 19&#160;Jun&#160;23&#160;13:19&#160;UTC | SHAKEN Socket Telecom LLC 554a | 19&#160;Jul&#160;23&#160;13:19&#160;UTC | true | [view](CERTS/a20a973710654cd4c5a27a7ed81db0958635d3fb26ad5563441609ff98202d91/README.md) |
| 19&#160;Jun&#160;23&#160;16:10&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 19&#160;Jul&#160;23&#160;16:10&#160;UTC | true | [view](CERTS/ce0eb0aaa2fd7830d2cef407d338913aa355d5d1f650542dbf7824c526eb6a43/README.md) |
| 20&#160;Jun&#160;23&#160;04:53&#160;UTC | SHAKEN BareTelecom 864J | 20&#160;Jul&#160;23&#160;04:53&#160;UTC | true | [view](CERTS/30f8ef82605a44d6443b7b96fd8c10ccab931eaaa16a5f23baf181162fc46575/README.md) |
| 20&#160;Jun&#160;23&#160;08:24&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 20&#160;Jul&#160;23&#160;08:24&#160;UTC | true | [view](CERTS/09d04ce28c6bfa5a43afb4629171a3ffd96045475b441732c307fd004f5e408d/README.md) |
| 20&#160;Jun&#160;23&#160;11:29&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 20&#160;Jul&#160;23&#160;11:29&#160;UTC | true | [view](CERTS/95e95a57182b6080b5f86517cb82fdcc93d6b96a84be6cc143b1acdb940a7980/README.md) |
| 20&#160;Jun&#160;23&#160;12:40&#160;UTC | SHAKEN IDT America, Corp 363A | 20&#160;Jul&#160;23&#160;12:40&#160;UTC | true | [view](CERTS/ff62e5dc6f152e20e4a3a3c82d9e81612ae4940740eed65860b724aa5fa90430/README.md) |
| 20&#160;Jun&#160;23&#160;15:40&#160;UTC | SHAKEN IDT America, Corp 363A | 20&#160;Jul&#160;23&#160;15:40&#160;UTC | true | [view](CERTS/2626168a9ab7d792e3d6b22dc356726ca9dfe6ef13a70477119f141d4ff136f5/README.md) |
| 20&#160;Jun&#160;23&#160;16:05&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 20&#160;Jul&#160;23&#160;16:05&#160;UTC | true | [view](CERTS/ee65cc276ddc18d68f56e26e3c5eeba05d95d9f61ddf5472ec1d14ad09d215ae/README.md) |
| 20&#160;Jun&#160;23&#160;22:42&#160;UTC | SHAKEN IDT America, Corp 363A | 20&#160;Jul&#160;23&#160;22:42&#160;UTC | true | [view](CERTS/3dc95d5df7fbf6a7edc6962235a678044144738c04d08cdc8c57333dd0bcbf7e/README.md) |
| 21&#160;Jun&#160;23&#160;01:54&#160;UTC | SHAKEN IPSBS Managed Services LLC 828J | 21&#160;Jul&#160;23&#160;01:54&#160;UTC | true | [view](CERTS/c029dc108da89ae3b67ac854f6a4c2454beca673c96ab6331086e81ec2ba33c7/README.md) |
| 21&#160;Jun&#160;23&#160;04:48&#160;UTC | SHAKEN BareTelecom 864J | 21&#160;Jul&#160;23&#160;04:48&#160;UTC | true | [view](CERTS/72214a89d2ba7ba27098d8183e6290aa77ff5754c354be7e56e43c2b354f80bc/README.md) |
| 21&#160;Jun&#160;23&#160;08:19&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 21&#160;Jul&#160;23&#160;08:19&#160;UTC | true | [view](CERTS/3352e1f29a9bbdc1737be046e507fffb7fa6b61c84fa785bbe3a3692ff3bf8a1/README.md) |
| 21&#160;Jun&#160;23&#160;11:24&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 21&#160;Jul&#160;23&#160;11:24&#160;UTC | true | [view](CERTS/53ad29de12c0b97590e203120c5a0a4889bc3e2dd0cb6bb63592bbaedfea10f4/README.md) |
| 21&#160;Jun&#160;23&#160;16:00&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 21&#160;Jul&#160;23&#160;16:00&#160;UTC | true | [view](CERTS/e51fc03a57d6b9a39503b22898930a9a52da1d62b3c2f31bf0f32488d595850b/README.md) |
| 21&#160;Jun&#160;23&#160;21:40&#160;UTC | SHAKEN ONE OWL TELECOM INC 412K | 21&#160;Jul&#160;23&#160;21:40&#160;UTC | true | [view](CERTS/940219eb974aa917999eeafca38337d8a24fd4ef911d451f22e804d558f64e86/README.md) |
| 21&#160;Jun&#160;23&#160;22:41&#160;UTC | SHAKEN IDT America, Corp 363A | 21&#160;Jul&#160;23&#160;22:41&#160;UTC | true | [view](CERTS/88be09df5d808ee5be71aa156f81b2eea4085d9f9f0699ad663d5e24e9414ad1/README.md) |
| 22&#160;Jun&#160;23&#160;01:49&#160;UTC | SHAKEN IPSBS Managed Services LLC 828J | 22&#160;Jul&#160;23&#160;01:49&#160;UTC | true | [view](CERTS/bfa12034869d5deb78951ad798f9d12dbfeb6b8242f7206d8a4e362b3ce5dc7a/README.md) |
| 22&#160;Jun&#160;23&#160;02:42&#160;UTC | SHAKEN InteractiveTel, LLC 920J | 22&#160;Jul&#160;23&#160;02:42&#160;UTC | true | [view](CERTS/140f4fef71fd6de628cea5e3f76d20c7522e7fad3d6337db1a992fa0d1106e5a/README.md) |
| 22&#160;Jun&#160;23&#160;08:14&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 22&#160;Jul&#160;23&#160;08:14&#160;UTC | true | [view](CERTS/3b1bf279b065cd40ebdc992076fe29941962f1d2664c87870e7ba8a25276b15c/README.md) |
| 22&#160;Jun&#160;23&#160;11:18&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 22&#160;Jul&#160;23&#160;11:18&#160;UTC | true | [view](CERTS/61524e6e9ef449aa26375ebcf8060fb9dc974a1927d41b11ed90ab1a17d37c21/README.md) |
| 22&#160;Jun&#160;23&#160;13:04&#160;UTC | SHAKEN Socket Telecom LLC 554a | 22&#160;Jul&#160;23&#160;13:04&#160;UTC | true | [view](CERTS/48b62cbb1c78ccd4af13ab1c29a8c8153f89506c80dd3c9ef6e117b3a55fc5fb/README.md) |
| 22&#160;Jun&#160;23&#160;14:31&#160;UTC | SHAKEN IDT America, Corp 363A | 22&#160;Jul&#160;23&#160;14:31&#160;UTC | true | [view](CERTS/b62fe4ac3f9346f39df3a7926215550b0f3790a3e131b5d8d112625739bbc6f4/README.md) |
| 22&#160;Jun&#160;23&#160;15:55&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 22&#160;Jul&#160;23&#160;15:55&#160;UTC | true | [view](CERTS/0891100fe2b21febddc7a35c0bbe8f0b8398e9df03936f8fcbc30b182ad098f8/README.md) |
| 22&#160;Jun&#160;23&#160;19:59&#160;UTC | SHAKEN Connexum LLC 203K | 21&#160;Jun&#160;24&#160;19:59&#160;UTC | true | [view](CERTS/d1f5789f3d44bf7545537a3539f0f5dbf43de9cca281266af09943e5635fecfd/README.md) |
| 23&#160;Jun&#160;23&#160;02:37&#160;UTC | SHAKEN InteractiveTel, LLC 920J | 23&#160;Jul&#160;23&#160;02:37&#160;UTC | true | [view](CERTS/7ce792dd10dcea723fe869e3fb9a78c8c2a194d98c484788b0940b4379f4a241/README.md) |
| 23&#160;Jun&#160;23&#160;08:09&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 23&#160;Jul&#160;23&#160;08:09&#160;UTC | true | [view](CERTS/91feb1222b13f7dca32add83933d99642e52f861d42b6bffd789d0f06955b4fd/README.md) |
| 23&#160;Jun&#160;23&#160;11:13&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 23&#160;Jul&#160;23&#160;11:13&#160;UTC | true | [view](CERTS/d3c931601950b5f09c9cd4ffba79a9ce8c9b438052a8fdfa7b4c2c0963e002f5/README.md) |
| 23&#160;Jun&#160;23&#160;14:56&#160;UTC | SHAKEN IDT America, Corp 363A | 23&#160;Jul&#160;23&#160;14:56&#160;UTC | true | [view](CERTS/ca25e73d97f09169749ddeb580b19a0d0391cd3568f6762a188959a92ce6cb8c/README.md) |
| 23&#160;Jun&#160;23&#160;15:50&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 23&#160;Jul&#160;23&#160;15:50&#160;UTC | true | [view](CERTS/22871926996c59dc68bf8c76b89d192e23247366069bc6e5f26d3355d16a3bab/README.md) |
| 24&#160;Jun&#160;23&#160;14:51&#160;UTC | SHAKEN IDT America, Corp 363A | 24&#160;Jul&#160;23&#160;14:51&#160;UTC | true | [view](CERTS/04d59c7dbecc004c1c42840388f178222a650dbd9544476ec903eef5eeeb86ae/README.md) |
| 25&#160;Jun&#160;23&#160;04:28&#160;UTC | SHAKEN BareTelecom 864J | 25&#160;Jul&#160;23&#160;04:28&#160;UTC | true | [view](CERTS/91d23ea552705e63c890f197ffd105e6a576f2bb573d7ff2cf72f178a3fbd08e/README.md) |
| 25&#160;Jun&#160;23&#160;11:03&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 25&#160;Jul&#160;23&#160;11:03&#160;UTC | true | [view](CERTS/6a84042e4ee57d0137bd8865dfee1b8f09d64d4534d3c6ff7055d30efa4a9ec3/README.md) |
| 25&#160;Jun&#160;23&#160;14:46&#160;UTC | SHAKEN IDT America, Corp 363A | 25&#160;Jul&#160;23&#160;14:46&#160;UTC | true | [view](CERTS/eed9df89b8e63ce26404d676b3340067e4c1e8c1b81e029943adcee686d98137/README.md) |
| 25&#160;Jun&#160;23&#160;15:40&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 25&#160;Jul&#160;23&#160;15:40&#160;UTC | true | [view](CERTS/cd1ce91991a6041ceb8ebc99462f5e798e3e469dcd8022fe0f04f691f67fb1e5/README.md) |
| 25&#160;Jun&#160;23&#160;17:58&#160;UTC | SHAKEN Zray Technologies Corporation 862J | 25&#160;Jul&#160;23&#160;17:58&#160;UTC | true | [view](CERTS/d935be7a3def21ef62a77a2a24a9122f1d4fc03573363ca666ec55ab7136a463/README.md) |
| 26&#160;Jun&#160;23&#160;02:22&#160;UTC | SHAKEN InteractiveTel, LLC 920J | 26&#160;Jul&#160;23&#160;02:22&#160;UTC | true | [view](CERTS/de04fab2d47b157738f56a995c8cd2b95046089a89cab061a2540d5175766f16/README.md) |
| 26&#160;Jun&#160;23&#160;04:23&#160;UTC | SHAKEN BareTelecom 864J | 26&#160;Jul&#160;23&#160;04:23&#160;UTC | true | [view](CERTS/30d7debf2614695ede7f6744d11c785625427189195525637e57b4d388feee5f/README.md) |
| 26&#160;Jun&#160;23&#160;07:54&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 26&#160;Jul&#160;23&#160;07:54&#160;UTC | true | [view](CERTS/c8d388c8c6cb8603b5aec0149e11179fdf42e01f1e6209a0372252e410162c32/README.md) |
| 26&#160;Jun&#160;23&#160;08:51&#160;UTC | SHAKEN IDT America, Corp 363A | 26&#160;Jul&#160;23&#160;08:51&#160;UTC | true | [view](CERTS/11fa66a6a3ff2f55749dfae25f9c6b9d78c317c9aa48b178dd3961a09da785ab/README.md) |
| 26&#160;Jun&#160;23&#160;15:35&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 26&#160;Jul&#160;23&#160;15:35&#160;UTC | true | [view](CERTS/02ac013a99b4c11e73e4b3d52c9476737f6d3bc80105bae4ad6dbbdf2d278d3a/README.md) |
| 26&#160;Jun&#160;23&#160;17:26&#160;UTC | SHAKEN IDT America, Corp 363A | 26&#160;Jul&#160;23&#160;17:26&#160;UTC | true | [view](CERTS/d9b0f7b161ecd5eb27cce4bcebf7c57a379d642a1a2d877948132003fa3c81c5/README.md) |
| 26&#160;Jun&#160;23&#160;22:40&#160;UTC | SHAKEN IDT America, Corp 363A | 26&#160;Jul&#160;23&#160;22:40&#160;UTC | true | [view](CERTS/9672979ccf0090213afd97c5c4164c1b3465ec698a48ad9ecd13c15f8752993c/README.md) |
| 27&#160;Jun&#160;23&#160;02:35&#160;UTC | SHAKEN Godaddy 463K | 27&#160;Jul&#160;23&#160;02:35&#160;UTC | true | [view](CERTS/60c59ae0f4e5e0fc085f04f108fd5c6c817bcca083d8f86318a49944a6695b5e/README.md) |
| 27&#160;Jun&#160;23&#160;04:18&#160;UTC | SHAKEN BareTelecom 864J | 27&#160;Jul&#160;23&#160;04:18&#160;UTC | true | [view](CERTS/25ede2920ac89f8fa878050fbde3a3b58e4fe423df2836ef32135f3ee611590c/README.md) |
| 27&#160;Jun&#160;23&#160;05:26&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 27&#160;Jul&#160;23&#160;05:26&#160;UTC | true | [view](CERTS/844a2c9da350d4da5238b67bfb202a514c9f8a482edc963e48d24b6249f2ddbf/README.md) |
| 27&#160;Jun&#160;23&#160;12:42&#160;UTC | SHAKEN Socket Telecom LLC 554a | 27&#160;Jul&#160;23&#160;12:42&#160;UTC | true | [view](CERTS/f5d967bbdd6dbe8a4b2e77faa9d7683dd2ad37a6642f1bd15ef7d0fac49cf79d/README.md) |
| 27&#160;Jun&#160;23&#160;15:30&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 27&#160;Jul&#160;23&#160;15:30&#160;UTC | true | [view](CERTS/6ac6f5b38b4196ffcb10c8f59731c0ed504a0b449b1130ef5e2dbf1026abc281/README.md) |
| 27&#160;Jun&#160;23&#160;17:58&#160;UTC | SHAKEN Zito Media Voice 624G | 25&#160;Sep&#160;23&#160;17:58&#160;UTC | true | [view](CERTS/d6a216dd047251b6976b020aa8d2e7986b2342459ba23bf2e81a270369166bb1/README.md) |
| 27&#160;Jun&#160;23&#160;22:42&#160;UTC | SHAKEN IDT America, Corp 363A | 27&#160;Jul&#160;23&#160;22:42&#160;UTC | true | [view](CERTS/4c304b9571468075d962e35190bb006ea27b1fc335b43bcc5baa6c975bfb416a/README.md) |
| 28&#160;Jun&#160;23&#160;02:12&#160;UTC | SHAKEN InteractiveTel, LLC 920J | 28&#160;Jul&#160;23&#160;02:12&#160;UTC | true | [view](CERTS/67fa1dbd6de02139a072cb0311637174a4356b4b2bbafd0d1820c869e9977ab4/README.md) |
| 28&#160;Jun&#160;23&#160;05:22&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 28&#160;Jul&#160;23&#160;05:22&#160;UTC | true | [view](CERTS/f1cc551bccdf296573e7c4f134e4447ccfb8edf11147e348e344f2fc41288219/README.md) |
| 28&#160;Jun&#160;23&#160;06:30&#160;UTC | SHAKEN  XCast Labs 689J | 29&#160;Jul&#160;23&#160;06:30&#160;UTC | true | [view](CERTS/2122c73fe69e0b7864e6049af7c6440ca3c54a6a17f96a8e260a023334c43775/README.md) |
| 28&#160;Jun&#160;23&#160;08:47&#160;UTC | SHAKEN BareTelecom 864J | 28&#160;Jul&#160;23&#160;08:47&#160;UTC | true | [view](CERTS/450b4e002e48c0273a12832b68555f00f609918e39abb34da6a7892c7d76ff7f/README.md) |
| 28&#160;Jun&#160;23&#160;10:48&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 28&#160;Jul&#160;23&#160;10:48&#160;UTC | true | [view](CERTS/ec5c2c882c079e5a58629531da2ed8ca98ec4f1350e8ef52b263dc87514fef28/README.md) |
| 28&#160;Jun&#160;23&#160;15:25&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 28&#160;Jul&#160;23&#160;15:25&#160;UTC | true | [view](CERTS/56e2ac8ddb6977ae9345a2782560a1445c969023ac355e72a1ba8b4372657294/README.md) |
| 28&#160;Jun&#160;23&#160;22:40&#160;UTC | SHAKEN IDT America, Corp 363A | 28&#160;Jul&#160;23&#160;22:40&#160;UTC | true | [view](CERTS/6f99d01b244372f135cb33162f3cd9d6af828a090adca1ccac0f17cd49aa4b89/README.md) |
| 29&#160;Jun&#160;23&#160;02:07&#160;UTC | SHAKEN InteractiveTel, LLC 920J | 29&#160;Jul&#160;23&#160;02:07&#160;UTC | true | [view](CERTS/9184985b2ab8cf914fbc407a4cbb97a0344f5fcf716bb17734d082e8cbc4f87d/README.md) |
| 29&#160;Jun&#160;23&#160;05:16&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 29&#160;Jul&#160;23&#160;05:16&#160;UTC | true | [view](CERTS/c134f463d0229732e9fb77f349f98f94add646aafdc39daa0f7256d9693dda9d/README.md) |
| 29&#160;Jun&#160;23&#160;08:42&#160;UTC | SHAKEN BareTelecom 864J | 29&#160;Jul&#160;23&#160;08:42&#160;UTC | true | [view](CERTS/2e5d10816cd1cf4c592b525803539777f5394a9ba3aac24aa256119d1dd13ac2/README.md) |
| 29&#160;Jun&#160;23&#160;10:43&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 29&#160;Jul&#160;23&#160;10:43&#160;UTC | true | [view](CERTS/9b114c9f2a67c5685a8806c0c3456a5a2fe337fed7b3d6f552d9548ec429345f/README.md) |
| 29&#160;Jun&#160;23&#160;15:20&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 29&#160;Jul&#160;23&#160;15:20&#160;UTC | true | [view](CERTS/6b6f50536720f7e44bf0d27cc8a8d5aece8d2c6f2c4971110fbdb0e02877ae0f/README.md) |
| 29&#160;Jun&#160;23&#160;18:03&#160;UTC | SHAKEN IPBTel 535K | 29&#160;Jul&#160;23&#160;18:03&#160;UTC | true | [view](CERTS/0be38bca737405d65c03cc332212d3729fc4dedda4fe22feafd6e70d7dafbbc0/README.md) |
| 29&#160;Jun&#160;23&#160;22:35&#160;UTC | SHAKEN IDT America, Corp 363A | 29&#160;Jul&#160;23&#160;22:35&#160;UTC | true | [view](CERTS/318459f46c5f507fea12c6e5dd9d7909034468da4627efbbb25e2a53d1e9f14d/README.md) |
| 29&#160;Jun&#160;23&#160;22:41&#160;UTC | SHAKEN IDT America, Corp 363A | 29&#160;Jul&#160;23&#160;22:41&#160;UTC | true | [view](CERTS/d0a9eda3577a5001f1d247eb80a17744c8e640d290b27c38e5d4a2f0053c239d/README.md) |
| 30&#160;Jun&#160;23&#160;02:02&#160;UTC | SHAKEN InteractiveTel, LLC 920J | 30&#160;Jul&#160;23&#160;02:02&#160;UTC | true | [view](CERTS/f977dad0299f43b701e2e6a451c34815e91960d3af17bf6a5fa7cf4dca27b638/README.md) |
| 30&#160;Jun&#160;23&#160;05:12&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 30&#160;Jul&#160;23&#160;05:12&#160;UTC | true | [view](CERTS/dcf4210752e5bc468e6f7afae29c511658c37c7383a9bb7d872a3e981d2de54f/README.md) |
| 30&#160;Jun&#160;23&#160;08:37&#160;UTC | SHAKEN BareTelecom 864J | 30&#160;Jul&#160;23&#160;08:37&#160;UTC | true | [view](CERTS/35cebae2518f954a0f3d2ac808e6d13c9a339f6bcf67c73782d6b5b3ebc153a5/README.md) |
| 30&#160;Jun&#160;23&#160;10:57&#160;UTC | SHAKEN IDT America, Corp 363A | 30&#160;Jul&#160;23&#160;10:57&#160;UTC | true | [view](CERTS/19dd1c4c538310fab50e27985e509c8e27230a967778f1cf4c7c1d9012ab62a0/README.md) |
| 01&#160;Jul&#160;23&#160;05:06&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 31&#160;Jul&#160;23&#160;05:06&#160;UTC | true | [view](CERTS/abb27dc2d35fdfb6ab03cc18a798467e1eaf0e0308e4feace01fdfaed19d8c5d/README.md) |
| 01&#160;Jul&#160;23&#160;08:32&#160;UTC | SHAKEN BareTelecom 864J | 31&#160;Jul&#160;23&#160;08:32&#160;UTC | true | [view](CERTS/5b3c77a2271109add39654e47e7543dde1da9f14f94ef32e3f37f2471dd479b8/README.md) |
| 01&#160;Jul&#160;23&#160;10:33&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 31&#160;Jul&#160;23&#160;10:33&#160;UTC | true | [view](CERTS/9fc9028f076f4dcc9ce704ce7d5d490a2c967e78b6750304ef1d318e7a7cc93c/README.md) |
| 01&#160;Jul&#160;23&#160;10:52&#160;UTC | SHAKEN IDT America, Corp 363A | 31&#160;Jul&#160;23&#160;10:52&#160;UTC | true | [view](CERTS/d00c17d0f21bc1b0ba13ea77005ee28c4e41a629c70029c824ec682fb8e0037a/README.md) |
| 02&#160;Jul&#160;23&#160;08:27&#160;UTC | SHAKEN BareTelecom 864J | 01&#160;Aug&#160;23&#160;08:27&#160;UTC | true | [view](CERTS/3b20085149ecf59e5a55ee0b8f4b87fffee4931c749118427a72da783cc3a8a3/README.md) |
| 02&#160;Jul&#160;23&#160;10:47&#160;UTC | SHAKEN IDT America, Corp 363A | 01&#160;Aug&#160;23&#160;10:47&#160;UTC | true | [view](CERTS/e22052a7073e51c33023c37ac209a85ae0dfd53865bf443bad76ff3e8606ae4e/README.md) |
| 02&#160;Jul&#160;23&#160;22:40&#160;UTC | SHAKEN IDT America, Corp 363A | 01&#160;Aug&#160;23&#160;22:40&#160;UTC | true | [view](CERTS/ea1f3fb881c3a67b64d0a4d4d923882ed542a3973608c761e5c1b2417859b319/README.md) |
| 03&#160;Jul&#160;23&#160;01:47&#160;UTC | SHAKEN InteractiveTel, LLC 920J | 02&#160;Aug&#160;23&#160;01:47&#160;UTC | true | [view](CERTS/f2a075006ca7034bc802a1474f385122d19f0a66404627595e339cb7c61c439b/README.md) |
| 03&#160;Jul&#160;23&#160;04:56&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 02&#160;Aug&#160;23&#160;04:56&#160;UTC | true | [view](CERTS/8b5f62098c548f6e645d44204a60868da107c0679ed3a75315bee15d5aa2e64a/README.md) |
| 03&#160;Jul&#160;23&#160;06:00&#160;UTC | SHAKEN Convoso 758J | 07&#160;Aug&#160;23&#160;06:00&#160;UTC | true | [view](CERTS/adeb01b337d7c99dc3dd7ced180401acc709cafd45e5cb4141856030c9cfeed9/README.md) |
| 03&#160;Jul&#160;23&#160;08:22&#160;UTC | SHAKEN BareTelecom 864J | 02&#160;Aug&#160;23&#160;08:22&#160;UTC | true | [view](CERTS/bb041eabfa513835a82511a38c5a4f7ebdfbbcd2cfa0aa9aa7cf9ea81a171dc7/README.md) |
| 03&#160;Jul&#160;23&#160;10:24&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 02&#160;Aug&#160;23&#160;10:24&#160;UTC | true | [view](CERTS/95af125ade79be7756140e0797f9636d41f78ae8fada66058c9b7a0ac348adf1/README.md) |
| 03&#160;Jul&#160;23&#160;22:40&#160;UTC | SHAKEN IDT America, Corp 363A | 02&#160;Aug&#160;23&#160;22:40&#160;UTC | true | [view](CERTS/cd2278ca461c82c8c3c09b48f640ef4bb9272b949f01b5b4ff498e432ec1fdce/README.md) |
| 04&#160;Jul&#160;23&#160;04:51&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 03&#160;Aug&#160;23&#160;04:51&#160;UTC | true | [view](CERTS/44991feba84e666024893b08d23eeb26612beddde88060636f3eab21b1f0199c/README.md) |
| 04&#160;Jul&#160;23&#160;08:17&#160;UTC | SHAKEN BareTelecom 864J | 03&#160;Aug&#160;23&#160;08:17&#160;UTC | true | [view](CERTS/e1c3b4cdcaf4fc4fdba2af2d06b04b8acd9073df953cde4e55fa9ab10335a27c/README.md) |
| 04&#160;Jul&#160;23&#160;13:23&#160;UTC | SHAKEN Threshold Communications Inc 563J | 03&#160;Aug&#160;23&#160;13:23&#160;UTC | true | [view](CERTS/755f065dea36267bbd30d6944c516b5c2876a07a066537317d8b953c7bedc42c/README.md) |
| 04&#160;Jul&#160;23&#160;13:23&#160;UTC | SHAKEN Broadband Dynamics LLC 583j | 03&#160;Aug&#160;23&#160;13:23&#160;UTC | true | [view](CERTS/cb6c1333ed85cb42565bd0460f7c83d911d7dc596d64196e744bb2f41692892e/README.md) |
| 04&#160;Jul&#160;23&#160;13:24&#160;UTC | SHAKEN Consolidated Communications 5113 | 03&#160;Aug&#160;23&#160;13:24&#160;UTC | true | [view](CERTS/3248f1cf1b83c570916f3ec46a89548a1876e95e18a7208604fea42f740f0706/README.md) |
| 04&#160;Jul&#160;23&#160;13:25&#160;UTC | SHAKEN Touchtone 683A | 03&#160;Aug&#160;23&#160;13:25&#160;UTC | true | [view](CERTS/115d9bf6dab7943aafc141e2440a41d3b645041175b17c96457c559db0a06e40/README.md) |
| 04&#160;Jul&#160;23&#160;13:26&#160;UTC | SHAKEN Apeiron Systems 012J | 03&#160;Aug&#160;23&#160;13:26&#160;UTC | true | [view](CERTS/533e049c49aa26ebf2ec638a3a008f9732f69e513141223a22efb4d5e5445d6f/README.md) |
| 04&#160;Jul&#160;23&#160;13:26&#160;UTC | SHAKEN Fonative, Inc. 684J | 03&#160;Aug&#160;23&#160;13:26&#160;UTC | true | [view](CERTS/897fbea948a78753338a11f1174cf2988eebd0723431fbc9303cb99827b8668b/README.md) |
| 04&#160;Jul&#160;23&#160;13:27&#160;UTC | SHAKEN IPitomy 652J | 03&#160;Aug&#160;23&#160;13:27&#160;UTC | true | [view](CERTS/4272e0885fd35aa1568f14b9400e837c8e66e2e8d2844d293f6b7a97a7674cc2/README.md) |
| 04&#160;Jul&#160;23&#160;13:27&#160;UTC | SHAKEN Phone.com, Inc. 633J | 03&#160;Aug&#160;23&#160;13:27&#160;UTC | true | [view](CERTS/7e8b20576c0783ed25009e97eab1e8806b8f5c91f018078c5845591380a9d8ff/README.md) |
| 04&#160;Jul&#160;23&#160;13:28&#160;UTC | SHAKEN NETRIO LLC 020K | 03&#160;Aug&#160;23&#160;13:28&#160;UTC | true | [view](CERTS/a03ec775e2d9da04caf8cf8946e2625c4485473ac76344c173cfdc3c2592c598/README.md) |
| 04&#160;Jul&#160;23&#160;13:29&#160;UTC | SHAKEN Noble Systems Communications LLC 187J | 03&#160;Aug&#160;23&#160;13:29&#160;UTC | true | [view](CERTS/e605ce890fcdea05c14670b2a65f2b0dc672fc8f4c9aa367437c0467d5a51c47/README.md) |
| 04&#160;Jul&#160;23&#160;13:29&#160;UTC | SHAKEN Telcentris Inc. dba Voxox 696J | 03&#160;Aug&#160;23&#160;13:29&#160;UTC | true | [view](CERTS/21a5b2695c70d75517c3874bd72447ad87aa4b51c025e8dca8c214f7d7384d5a/README.md) |
| 04&#160;Jul&#160;23&#160;13:30&#160;UTC | SHAKEN Airespring 996H | 03&#160;Aug&#160;23&#160;13:30&#160;UTC | true | [view](CERTS/efb5b66d7a3ba9f683421eaebc8d9aacbe583327d243d3c7b93fba9bf1d3a652/README.md) |
| 04&#160;Jul&#160;23&#160;13:30&#160;UTC | SHAKEN Nobelbiz, Inc. 596J | 03&#160;Aug&#160;23&#160;13:30&#160;UTC | true | [view](CERTS/44a0976066f1582e33dec9a7f563f13c47168330fb15984db94c6e4186839b80/README.md) |
| 04&#160;Jul&#160;23&#160;13:31&#160;UTC | SHAKEN Momentum Telecom 1417 | 03&#160;Aug&#160;23&#160;13:31&#160;UTC | true | [view](CERTS/ce1dbd396722c85078e9aec7ff30b8759d66895231ba9e398b641902f9837523/README.md) |
| 04&#160;Jul&#160;23&#160;13:32&#160;UTC | SHAKEN Momentum Telecom 9157 | 03&#160;Aug&#160;23&#160;13:32&#160;UTC | true | [view](CERTS/5d5f1208363e253fac7209bceac09405f31af6f456c1bf0f406e004065790b2f/README.md) |
| 04&#160;Jul&#160;23&#160;13:33&#160;UTC | SHAKEN Terra Nova Telecom 382G | 03&#160;Aug&#160;23&#160;13:33&#160;UTC | true | [view](CERTS/2aa497fb827f41c2debc5caca3427afac07848ad23999a9018266f8f95538c6f/README.md) |
| 04&#160;Jul&#160;23&#160;13:33&#160;UTC | SHAKEN Matrix 3058 | 03&#160;Aug&#160;23&#160;13:33&#160;UTC | true | [view](CERTS/ff3a8b700d56a9c597998af31bf7e37161021060efba84bda9e5bda8758e20e7/README.md) |
| 04&#160;Jul&#160;23&#160;13:33&#160;UTC | SHAKEN Matrix 9451 | 03&#160;Aug&#160;23&#160;13:33&#160;UTC | true | [view](CERTS/298960cc1a62fa270c1fefd3530e9c6ed32838c7dcfa511d54163118f3e15ec8/README.md) |
| 04&#160;Jul&#160;23&#160;13:34&#160;UTC | SHAKEN Matrix 7379 | 03&#160;Aug&#160;23&#160;13:34&#160;UTC | true | [view](CERTS/c2c54dab2ffc7e78d8b8837e8370ccdd8366ef3cd8e4e6bb0a3d31da3f672d88/README.md) |
| 04&#160;Jul&#160;23&#160;13:35&#160;UTC | SHAKEN PNG Telecommunications Inc 3395 | 03&#160;Aug&#160;23&#160;13:35&#160;UTC | true | [view](CERTS/8d8aab5257669db99fdac0cec05d2fa3582a6136d0272d2f8edec3fa5e3bd08d/README.md) |
| 04&#160;Jul&#160;23&#160;13:36&#160;UTC | SHAKEN Magna5, LLC 3849 | 03&#160;Aug&#160;23&#160;13:36&#160;UTC | true | [view](CERTS/81f722e9337fdfe38ac18868246127dbebc60139d0ce11beb2869197e763d549/README.md) |
| 04&#160;Jul&#160;23&#160;13:38&#160;UTC | SHAKEN Magna5, LLC 8249 | 03&#160;Aug&#160;23&#160;13:38&#160;UTC | true | [view](CERTS/c1011383cf0b5208dd5468614975ae724c87193ca9a9498369b465cfa8a19209/README.md) |
| 05&#160;Jul&#160;23&#160;01:37&#160;UTC | SHAKEN InteractiveTel, LLC 920J | 04&#160;Aug&#160;23&#160;01:37&#160;UTC | true | [view](CERTS/c547acbd4792848762e24fc378aeb32daef33fd1b49e83640ca09a22df3792ba/README.md) |
| 05&#160;Jul&#160;23&#160;04:46&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 04&#160;Aug&#160;23&#160;04:46&#160;UTC | true | [view](CERTS/fec7e74b0c0ac7bc37ca866f1cd4ca720906dd98fd9a4a7c87eaed27c56a6d59/README.md) |
| 05&#160;Jul&#160;23&#160;08:46&#160;UTC | SHAKEN BareTelecom 864J | 04&#160;Aug&#160;23&#160;08:46&#160;UTC | true | [view](CERTS/c360c3e8e3099450e09e577ab6fba236fdd01fc6ae930e1f6f9ede78a17dd452/README.md) |
| 05&#160;Jul&#160;23&#160;10:13&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 04&#160;Aug&#160;23&#160;10:13&#160;UTC | true | [view](CERTS/396090a433f8611478608e9c2da551c5102d4b1ef3b9bbabdc46f88303bea9d7/README.md) |
| 05&#160;Jul&#160;23&#160;13:02&#160;UTC | SHAKEN IDT America, Corp 363A | 04&#160;Aug&#160;23&#160;13:02&#160;UTC | true | [view](CERTS/ab9c3b357890a28895d00b4b689462e1e3a24da9b9e6ef20d40c1ead8e9a6a9e/README.md) |
| 05&#160;Jul&#160;23&#160;14:16&#160;UTC | SHAKEN IDT America, Corp 363A | 04&#160;Aug&#160;23&#160;14:16&#160;UTC | true | [view](CERTS/94f57ba07cd29c66daf0f03047e5f91e199673d0c2a73221f6b60f8b7f10f648/README.md) |
| 05&#160;Jul&#160;23&#160;16:47&#160;UTC | SHAKEN Socket Telecom LLC 554a | 04&#160;Aug&#160;23&#160;16:47&#160;UTC | true | [view](CERTS/546a03d3eb72491b6c7f696c9e5be99ffd40a5ad5d5b663d112ed32a076ff81f/README.md) |
| 05&#160;Jul&#160;23&#160;18:58&#160;UTC | SHAKEN IDT America, Corp 363A | 04&#160;Aug&#160;23&#160;18:58&#160;UTC | true | [view](CERTS/9e33e07fea1d9693c3d073a15930a39eaf74fe6081cfe1ed6b89a518a5523bc1/README.md) |
| 05&#160;Jul&#160;23&#160;19:27&#160;UTC | SHAKEN IDT America, Corp 363A | 04&#160;Aug&#160;23&#160;19:27&#160;UTC | true | [view](CERTS/bd8383cea818f0d58a9334db61d80f6a6b442eafe3b38600e063d09b70e8b916/README.md) |
| 05&#160;Jul&#160;23&#160;20:40&#160;UTC | SHAKEN IDT America, Corp 363A | 04&#160;Aug&#160;23&#160;20:40&#160;UTC | true | [view](CERTS/b743e5e3d7ba569eef0c7f273c7169586b6bb374bdb1071a14c593c64af78c5c/README.md) |
| 05&#160;Jul&#160;23&#160;20:53&#160;UTC | SHAKEN IDT America, Corp 363A | 04&#160;Aug&#160;23&#160;20:53&#160;UTC | true | [view](CERTS/3827571192e85b6f7d76641dfad906278a26116b491e58f0bed0cf512e7f4246/README.md) |

#### CA Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 21&#160;Aug&#160;20&#160;01:22&#160;UTC | SHAKEN Sansay Root CA US | 16&#160;Aug&#160;40&#160;01:22&#160;UTC | true | [view](CERTS/8356d251b255f4ac3fd108bb79be4c02dcea2d3b13d138892bdb3a70cbe6a343/README.md) |
| 02&#160;Sep&#160;22&#160;20:53&#160;UTC | SHAKEN Sansay Intermediate CA US WEST 1 | 31&#160;Aug&#160;29&#160;20:53&#160;UTC | true | [view](CERTS/4b1dfdba2b1e4bbffbf900a20f1f6f7befbef0008b963e4922a64469cb97d24b/README.md) |


Generated: 06 Jul 23 14:08 UTC