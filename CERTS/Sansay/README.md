# STIR/SHAKEN CA Ecosystem Compliance

## Sansay

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 660 certificates were included in the corpus being tested
- 8 certificates in the corpus were skipped because they are duplicates
- 469 certificates in the corpus were skipped because they are expired
- 1 certificates in the corpus were skipped because they are not currently trusted
- 182 certificates being tested against the remaining rules
- 4.95 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 100.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 103 days is the average remaining validity for the certificates in the corpus
- 103 days is the average initial validity for the certificates in the corpus
- 139 certificates expire in the next 30 days
- 2.43 average number of unexpired certificates per OCN observed
- 75 unique OCNs observed in unexpired and valid certificate corpus

| Instances | Test | Source |
|-----------|------|--------|
| 59 | [e_atis_certificate_policies](ISSUES/e_atis_certificate_policies/README.md) | ATIS1000080 |
| 182 | [e_atis_subject_cn](ISSUES/e_atis_subject_cn/README.md) | ATIS1000080 |
| 182 | [e_us_cp_ambiguous_identifier](ISSUES/e_us_cp_ambiguous_identifier/README.md) | US_SHAKEN_CP |
| 182 | [e_us_cp_subject_sn](ISSUES/e_us_cp_subject_sn/README.md) | US_SHAKEN_CP |
| 113 | [w_cp_1_3_subject_email](ISSUES/w_cp_1_3_subject_email/README.md) | US_SHAKEN_CP |
| 182 | [w_shaken_subject_rdn_unknown](ISSUES/w_shaken_subject_rdn_unknown/README.md) | SHAKEN_PKI_BEST_PRACTICES |

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
- 5392 days is the average remaining validity for the certificates in the corpus
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
| 19&#160;Oct&#160;22&#160;19:08&#160;UTC | SHAKEN Ytel Inc. 703J | 19&#160;Oct&#160;23&#160;19:08&#160;UTC | true | [view](CERTS/e0c7a355b91ad947dd48fc5a84523293447c45eac28d955f230fb212a73e34c3/README.md) |
| 24&#160;Oct&#160;22&#160;20:23&#160;UTC | SHAKEN Arbeit 816J | 24&#160;Oct&#160;23&#160;20:23&#160;UTC | true | [view](CERTS/377e182a223e6cc8d7e9ce697e7a3e829b1c6b16c299c26f6d1f1e33aa29524b/README.md) |
| 24&#160;Oct&#160;22&#160;21:11&#160;UTC | SHAKEN Ytel Inc. 703J | 24&#160;Oct&#160;23&#160;21:11&#160;UTC | true | [view](CERTS/3d6a7a2ff23b90fba1674f600a108b8a11a110f8bb1723df86627001f7367d8d/README.md) |
| 25&#160;Oct&#160;22&#160;20:17&#160;UTC | SHAKEN Talk IT Pro 321K | 25&#160;Oct&#160;23&#160;20:17&#160;UTC | true | [view](CERTS/e45dada701a589e681d12207ebf16985abf6d62cf429b6e03bdcf8c0f97c3bf2/README.md) |
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
| 31&#160;Jan&#160;23&#160;23:39&#160;UTC | SHAKEN Connexum LLC 203K | 01&#160;May&#160;23&#160;23:39&#160;UTC | true | [view](CERTS/d5409b0a1e255cfd6a3557acd45479a31641ee57e77ac96e1865499037f56c18/README.md) |
| 08&#160;Feb&#160;23&#160;19:07&#160;UTC | SHAKEN ConvergeTel LLC 388K | 08&#160;Feb&#160;24&#160;19:07&#160;UTC | true | [view](CERTS/80706b79565f875515eb32f8cf113093a2658148ece8440e76199e4004254c31/README.md) |
| 14&#160;Feb&#160;23&#160;17:12&#160;UTC | SHAKEN Ytel Inc. 703J | 14&#160;Feb&#160;24&#160;17:12&#160;UTC | true | [view](CERTS/14c9bef113cfebe60611b0c56c430518ff8d42e8b98dd7e4653bd9cf619d5641/README.md) |
| 19&#160;Mar&#160;23&#160;00:31&#160;UTC | SHAKEN Every1 Telecom 486K | 18&#160;Mar&#160;24&#160;00:31&#160;UTC | true | [view](CERTS/5cdfbb1a416083096dfef10c75a2b26a08d8fd5593d8ea9ceae0d70d878a97d1/README.md) |
| 27&#160;Mar&#160;23&#160;21:48&#160;UTC | SHAKEN Global Net Holdings Inc 306K | 21&#160;Jan&#160;24&#160;21:48&#160;UTC | true | [view](CERTS/f8d913af64a11c0718c78426b747dba4ed4ccb19239db573626a46f29b671825/README.md) |
| 28&#160;Mar&#160;23&#160;15:31&#160;UTC | SHAKEN VoIP Innovations 597F | 27&#160;Mar&#160;24&#160;15:31&#160;UTC | true | [view](CERTS/105a7683d4b5fac2ea7c2383e95250715bc0460d2cfdbea0d220201f44ea5d0c/README.md) |
| 29&#160;Mar&#160;23&#160;05:52&#160;UTC | SHAKEN BareTelecom 864J | 28&#160;Apr&#160;23&#160;05:52&#160;UTC | true | [view](CERTS/84cef36f0282839c80028c6f4f4839af490e48b3970009feb9a393cb37946ab3/README.md) |
| 29&#160;Mar&#160;23&#160;09:16&#160;UTC | SHAKEN IPSBS Managed Services LLC 828J | 28&#160;Apr&#160;23&#160;09:16&#160;UTC | true | [view](CERTS/f587f4e05a6683c7c9950522da769b148b810012dc0211351307ca66314d88a2/README.md) |
| 29&#160;Mar&#160;23&#160;10:36&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 28&#160;Apr&#160;23&#160;10:36&#160;UTC | true | [view](CERTS/c2a70c6910ac71b06e20e24debb7b6aa20dcee5ab33e5e3e7e595955744eef82/README.md) |
| 29&#160;Mar&#160;23&#160;12:18&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 28&#160;Apr&#160;23&#160;12:18&#160;UTC | true | [view](CERTS/6db836ae29da64bb3b1d5dcea17ddcd29df0f80eb5a8e620dd305bdb7448cfc1/README.md) |
| 29&#160;Mar&#160;23&#160;12:53&#160;UTC | SHAKEN Zray Technologies Corporation 862J | 28&#160;Apr&#160;23&#160;12:53&#160;UTC | true | [view](CERTS/61cf21e1efc96e62c17ea29d823333eaa8e7ecaa18e291cd1f10490f50c04698/README.md) |
| 29&#160;Mar&#160;23&#160;17:22&#160;UTC | SHAKEN NTC International, INC 016K | 28&#160;Apr&#160;23&#160;17:22&#160;UTC | true | [view](CERTS/89bacde241e86626f947011f3d5c2f8185e8c273043fb2d7353cb64fc5f69c04/README.md) |
| 29&#160;Mar&#160;23&#160;21:08&#160;UTC | SHAKEN IDT America, Corp 363A | 28&#160;May&#160;23&#160;21:08&#160;UTC | true | [view](CERTS/eeed0eb15030ae3a643aac3b531aeef87806dc25296f5c49140de1e71ac88c00/README.md) |
| 30&#160;Mar&#160;23&#160;05:47&#160;UTC | SHAKEN BareTelecom 864J | 29&#160;Apr&#160;23&#160;05:47&#160;UTC | true | [view](CERTS/5af5730554ccc77f1af169e7b4ceeb6a09fe5bf8a6b2cda2dc9bae8340cd8f2e/README.md) |
| 30&#160;Mar&#160;23&#160;09:11&#160;UTC | SHAKEN IPSBS Managed Services LLC 828J | 29&#160;Apr&#160;23&#160;09:11&#160;UTC | true | [view](CERTS/5ea6e6af245e07aec9c26230f6218d1a4d5334ad3e04bdd1c62f9f95d5f8a3ec/README.md) |
| 30&#160;Mar&#160;23&#160;10:31&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 29&#160;Apr&#160;23&#160;10:31&#160;UTC | true | [view](CERTS/600d281873b42cbc0445aa8a9fbb95386172952ec10de375b290901db60ddb26/README.md) |
| 30&#160;Mar&#160;23&#160;10:47&#160;UTC | SHAKEN InteractiveTel, LLC 920J | 29&#160;Apr&#160;23&#160;10:47&#160;UTC | true | [view](CERTS/8bde600afa9f8df40b7ba5aa83522de993dc37e6749ceac3976029ee3e56286d/README.md) |
| 30&#160;Mar&#160;23&#160;12:13&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 29&#160;Apr&#160;23&#160;12:13&#160;UTC | true | [view](CERTS/029b3b1e7c3f24f0c7faa79544dccc0aa26b6f19a6f3e204222cfcb9e01daec5/README.md) |
| 31&#160;Mar&#160;23&#160;03:00&#160;UTC | SHAKEN 1stPoint Communications, LLC 463G | 30&#160;Apr&#160;23&#160;03:00&#160;UTC | true | [view](CERTS/4a18703f800785363129c7c347f81406d2c28f5c3890909cd6549c24d216ef44/README.md) |
| 31&#160;Mar&#160;23&#160;05:42&#160;UTC | SHAKEN BareTelecom 864J | 30&#160;Apr&#160;23&#160;05:42&#160;UTC | true | [view](CERTS/22581b2f099ee948a330c33a898ced08d6efa8f84adadbbc453c1289b754294e/README.md) |
| 31&#160;Mar&#160;23&#160;09:06&#160;UTC | SHAKEN IPSBS Managed Services LLC 828J | 30&#160;Apr&#160;23&#160;09:06&#160;UTC | true | [view](CERTS/c0d5fe63d3eb082c66852ac45f88a6d561fc7e6639bac012cf2c1d4c2829ce64/README.md) |
| 31&#160;Mar&#160;23&#160;10:26&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 30&#160;Apr&#160;23&#160;10:26&#160;UTC | true | [view](CERTS/8cbabd0259acd180bab539fdf81234170021c039f1ae570ba2ee42756d042bb1/README.md) |
| 31&#160;Mar&#160;23&#160;12:08&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 30&#160;Apr&#160;23&#160;12:08&#160;UTC | true | [view](CERTS/18740193f53a527d8aded6aeae60d001c5fb7fd1f605544942daadd1e2d77955/README.md) |
| 31&#160;Mar&#160;23&#160;18:14&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 30&#160;Apr&#160;23&#160;18:14&#160;UTC | true | [view](CERTS/cae39149dc65689067e4b94b42f1e8523ae449dd0b29555e047c34a636a22e64/README.md) |
| 01&#160;Apr&#160;23&#160;05:37&#160;UTC | SHAKEN BareTelecom 864J | 01&#160;May&#160;23&#160;05:37&#160;UTC | true | [view](CERTS/c601e1fd08ae48c874d425a1f3fe37aec10b3f6fb1d9aa4a5a1d07d1d030629d/README.md) |
| 01&#160;Apr&#160;23&#160;10:21&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 01&#160;May&#160;23&#160;10:21&#160;UTC | true | [view](CERTS/4ffa441025ceec939af7baf8e0d93af1fc24039733a66519aa6f43623b00d535/README.md) |
| 02&#160;Apr&#160;23&#160;05:32&#160;UTC | SHAKEN BareTelecom 864J | 02&#160;May&#160;23&#160;05:32&#160;UTC | true | [view](CERTS/985e899d40e3ae190048b092266083403854374185d52c3832c24c6cedf975f7/README.md) |
| 02&#160;Apr&#160;23&#160;17:02&#160;UTC | SHAKEN NTC International, INC 016K | 02&#160;May&#160;23&#160;17:02&#160;UTC | true | [view](CERTS/e2d2e7f2c6d06d70dd58548efd33c7b2a963eca135db6b1f8003746efa9d0411/README.md) |
| 02&#160;Apr&#160;23&#160;18:03&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 02&#160;May&#160;23&#160;18:03&#160;UTC | true | [view](CERTS/0f52e763c6fe670cb53d06e8185b47c354d11e1c9c2b3a0e7290624a2ade5f45/README.md) |
| 03&#160;Apr&#160;23&#160;05:27&#160;UTC | SHAKEN BareTelecom 864J | 03&#160;May&#160;23&#160;05:27&#160;UTC | true | [view](CERTS/964296a5fe0036c8673663809408c6be0d0a1ae7a319f8eb327e358d42b77940/README.md) |
| 03&#160;Apr&#160;23&#160;10:11&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 03&#160;May&#160;23&#160;10:11&#160;UTC | true | [view](CERTS/daebdec1468f3a59d06a6cc27fa7e5c9de79c641177823095a3352bbbd1c86ac/README.md) |
| 03&#160;Apr&#160;23&#160;11:53&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 03&#160;May&#160;23&#160;11:53&#160;UTC | true | [view](CERTS/f9a55676f9282ce517e9f3db9d891e0f7c3389ac09a6cc4301d0dcab283942a0/README.md) |
| 03&#160;Apr&#160;23&#160;14:55&#160;UTC | SHAKEN CIMA Telecom, Inc 313K | 03&#160;May&#160;23&#160;14:55&#160;UTC | true | [view](CERTS/109c2f9c177891c0a164ef13f97f0b84c80d3405e843ebdb4586be6e036dd38a/README.md) |
| 03&#160;Apr&#160;23&#160;21:02&#160;UTC | SHAKEN WWT INC dba VoIP Networks 053K | 02&#160;Apr&#160;24&#160;21:02&#160;UTC | true | [view](CERTS/6229273e97413ee4c01a1810885db9b48d7e1bf7fe8aa4e7be22076effe2cc8b/README.md) |
| 04&#160;Apr&#160;23&#160;05:22&#160;UTC | SHAKEN BareTelecom 864J | 04&#160;May&#160;23&#160;05:22&#160;UTC | true | [view](CERTS/6f40217f90766b47b3cc33bc972969eff9f7cc61722a32c1caa2173425587eaf/README.md) |
| 04&#160;Apr&#160;23&#160;06:00&#160;UTC | SHAKEN Convoso 758J | 09&#160;May&#160;23&#160;06:00&#160;UTC | true | [view](CERTS/b17a0addedeea480f7a605fe1516b63f4edd4dc39b84659f71868fbd86ac33e5/README.md) |
| 04&#160;Apr&#160;23&#160;06:30&#160;UTC | SHAKEN  XCast Labs 689J | 05&#160;May&#160;23&#160;06:30&#160;UTC | true | [view](CERTS/10484815d46dbd8a060cb2e928023ea57c6651cbc05524caab490fc2b7fa5dad/README.md) |
| 04&#160;Apr&#160;23&#160;08:46&#160;UTC | SHAKEN IPSBS Managed Services LLC 828J | 04&#160;May&#160;23&#160;08:46&#160;UTC | true | [view](CERTS/c0d877671b1db709f26bbe51ef3307e51c89cb47395219e9256046a1b83c2bdb/README.md) |
| 04&#160;Apr&#160;23&#160;10:06&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 04&#160;May&#160;23&#160;10:06&#160;UTC | true | [view](CERTS/933a297e7c51c4f7222d0100cde83ca9d3384172d86a96d48c045c911d5ffeae/README.md) |
| 04&#160;Apr&#160;23&#160;10:22&#160;UTC | SHAKEN InteractiveTel, LLC 920J | 04&#160;May&#160;23&#160;10:22&#160;UTC | true | [view](CERTS/f1cbf1850d2840a5c03c570c3c4ccd11ed90aebb4831b4c80558cfd49a337b60/README.md) |
| 04&#160;Apr&#160;23&#160;11:48&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 04&#160;May&#160;23&#160;11:48&#160;UTC | true | [view](CERTS/80f6076fe33e02b41f390d830c727ba5d4c180b92538be959df869583e244e4a/README.md) |
| 04&#160;Apr&#160;23&#160;16:23&#160;UTC | SHAKEN Cyberlynk Network, LLC 086K | 03&#160;Jun&#160;23&#160;16:23&#160;UTC | true | [view](CERTS/0b911cba5097116a1c6e0874d1b80f728e3877fd11c49d10ef0625d905b2dfee/README.md) |
| 04&#160;Apr&#160;23&#160;16:38&#160;UTC | SHAKEN Inventive Labs Corp 649J | 01&#160;Oct&#160;23&#160;16:38&#160;UTC | true | [view](CERTS/61bd59e51cd88d19ca695a362569930dee7725fefa39299b507bc3269fefbd54/README.md) |
| 04&#160;Apr&#160;23&#160;16:52&#160;UTC | SHAKEN NTC International, INC 016K | 04&#160;May&#160;23&#160;16:52&#160;UTC | true | [view](CERTS/e54318a114cf591fa3686de50524be5c066e624f4895e369bfc4f052e70ffc23/README.md) |
| 04&#160;Apr&#160;23&#160;17:53&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 04&#160;May&#160;23&#160;17:53&#160;UTC | true | [view](CERTS/715e82b90f644904cf7e3b8c0b68e30b143707b8223dfae92da3142fac6f1b26/README.md) |
| 05&#160;Apr&#160;23&#160;05:17&#160;UTC | SHAKEN BareTelecom 864J | 05&#160;May&#160;23&#160;05:17&#160;UTC | true | [view](CERTS/c6a011f42c10232450ab062e487e1f211503704ffcae1e1ab98110bfdb4d2405/README.md) |
| 05&#160;Apr&#160;23&#160;10:01&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 05&#160;May&#160;23&#160;10:01&#160;UTC | true | [view](CERTS/e3e764fc2461ff3b48c1cdda7ca62d7b8bf27577de26094f9c93258113bd9579/README.md) |
| 05&#160;Apr&#160;23&#160;10:17&#160;UTC | SHAKEN InteractiveTel, LLC 920J | 05&#160;May&#160;23&#160;10:17&#160;UTC | true | [view](CERTS/fa79bea150b40a3beef2a7a5adcb5a2371f783fceee6f931696d2b2e73bbf118/README.md) |
| 05&#160;Apr&#160;23&#160;11:43&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 05&#160;May&#160;23&#160;11:43&#160;UTC | true | [view](CERTS/7974bc253d1dd9d85fc48f0c8d0027e9b33622705d0cfad234c95d035282c2dd/README.md) |
| 05&#160;Apr&#160;23&#160;17:48&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 05&#160;May&#160;23&#160;17:48&#160;UTC | true | [view](CERTS/46379baedb5eca4a9fef03c2de2abf75c1ac376db44f5ea233aa031ed46ae6e5/README.md) |
| 06&#160;Apr&#160;23&#160;05:12&#160;UTC | SHAKEN BareTelecom 864J | 06&#160;May&#160;23&#160;05:12&#160;UTC | true | [view](CERTS/7f63bfd7017f742589027e636d2036c0bd472f4eab15a4673adfc1cfe5a3529c/README.md) |
| 06&#160;Apr&#160;23&#160;08:18&#160;UTC | SHAKEN IDT America, Corp 363A | 06&#160;May&#160;23&#160;08:18&#160;UTC | true | [view](CERTS/10efd272208f395892b5c2e5e9ca6292ecc6b73ddc4dfbbde242429b50214188/README.md) |
| 06&#160;Apr&#160;23&#160;09:56&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 06&#160;May&#160;23&#160;09:56&#160;UTC | true | [view](CERTS/88903a4c7fc5acd3f370c06c5f5a69f965c852f2b478e5771da3773aed1232f3/README.md) |
| 06&#160;Apr&#160;23&#160;11:38&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 06&#160;May&#160;23&#160;11:38&#160;UTC | true | [view](CERTS/c4a2752f3a999dfb6eb9d0e1cfbb372f509b65aa6af2256a79915131b642aab8/README.md) |
| 07&#160;Apr&#160;23&#160;05:07&#160;UTC | SHAKEN BareTelecom 864J | 07&#160;May&#160;23&#160;05:07&#160;UTC | true | [view](CERTS/6ef8f07849285ada85cc9448ee88e0e7a08a858573e36353407a2580e1b13672/README.md) |
| 07&#160;Apr&#160;23&#160;08:13&#160;UTC | SHAKEN IDT America, Corp 363A | 07&#160;May&#160;23&#160;08:13&#160;UTC | true | [view](CERTS/a1846bf23bb980a907721a2f85a1d425cf4602fd6f2990328a0a2e0b25af54ba/README.md) |
| 07&#160;Apr&#160;23&#160;09:51&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 07&#160;May&#160;23&#160;09:51&#160;UTC | true | [view](CERTS/1235e62d6830384821956010055d10d6e4218c344454dd0f95b2da61ab51e69f/README.md) |
| 07&#160;Apr&#160;23&#160;11:33&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 07&#160;May&#160;23&#160;11:33&#160;UTC | true | [view](CERTS/a2f7ef628b8f1723c31355f016ff0a506b41539e4d108f6732fe105d5a968738/README.md) |
| 08&#160;Apr&#160;23&#160;05:02&#160;UTC | SHAKEN BareTelecom 864J | 08&#160;May&#160;23&#160;05:02&#160;UTC | true | [view](CERTS/54b7a842add78bfd5d68781ecb2367a670a9b4442da42a88a7f8fa1852f14ec6/README.md) |
| 08&#160;Apr&#160;23&#160;08:08&#160;UTC | SHAKEN IDT America, Corp 363A | 08&#160;May&#160;23&#160;08:08&#160;UTC | true | [view](CERTS/02fbde7e96775baf71567ee2fd3d6121d1b42156a7134164871d2ac999541a65/README.md) |
| 09&#160;Apr&#160;23&#160;08:03&#160;UTC | SHAKEN IDT America, Corp 363A | 09&#160;May&#160;23&#160;08:03&#160;UTC | true | [view](CERTS/a0441756b0b988f2f78a8fa40ff6718197cc90cb431b37fdfed873027c233f4f/README.md) |
| 09&#160;Apr&#160;23&#160;16:27&#160;UTC | SHAKEN NTC International, INC 016K | 09&#160;May&#160;23&#160;16:27&#160;UTC | true | [view](CERTS/29322bf8e3d9d72db8d52a4cc2a0e55d1195a4ca4f6703056be44e9cc90ecdbb/README.md) |
| 10&#160;Apr&#160;23&#160;04:52&#160;UTC | SHAKEN BareTelecom 864J | 10&#160;May&#160;23&#160;04:52&#160;UTC | true | [view](CERTS/0b4856f9f0b729cd475d620b7ad37be3c5078bcd5e8f697fbcf533c7a281c4f6/README.md) |
| 10&#160;Apr&#160;23&#160;07:22&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 10&#160;May&#160;23&#160;07:22&#160;UTC | true | [view](CERTS/5d36d2ff4459da3cf3f65300b2e15066e93fb454a9c03eb4eb29b993d74b45ed/README.md) |
| 10&#160;Apr&#160;23&#160;07:58&#160;UTC | SHAKEN IDT America, Corp 363A | 10&#160;May&#160;23&#160;07:58&#160;UTC | true | [view](CERTS/9ee154b16d3ea79c37ad32b275c2d28a7924f5c304834b0a5adf9a6149ef5149/README.md) |
| 10&#160;Apr&#160;23&#160;09:52&#160;UTC | SHAKEN InteractiveTel, LLC 920J | 10&#160;May&#160;23&#160;09:52&#160;UTC | true | [view](CERTS/703993e758ab1a00e2a1df33ce08072fd5b65863e614f237ec32acc2ee0ca53a/README.md) |
| 10&#160;Apr&#160;23&#160;11:18&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 10&#160;May&#160;23&#160;11:18&#160;UTC | true | [view](CERTS/cbcd4d803fd9ac115d141980402c6881cb4ad0a472a787b17509dc2fd4ff8bd1/README.md) |
| 10&#160;Apr&#160;23&#160;16:22&#160;UTC | SHAKEN NTC International, INC 016K | 10&#160;May&#160;23&#160;16:22&#160;UTC | true | [view](CERTS/e14233c21912bbcecc0d91b9abeb6e369dfbb5dad0f038ce712b17da9e1cf493/README.md) |
| 10&#160;Apr&#160;23&#160;17:23&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 10&#160;May&#160;23&#160;17:23&#160;UTC | true | [view](CERTS/5d0b648264e4c66bd426bd72951d55f3189dd46104b624e59b10b1c8856e81ee/README.md) |
| 11&#160;Apr&#160;23&#160;04:47&#160;UTC | SHAKEN BareTelecom 864J | 11&#160;May&#160;23&#160;04:47&#160;UTC | true | [view](CERTS/45f723de08fedaf15af8c1fdeb3dfbcdbcbc613ac8720f910e4dd30efcebac78/README.md) |
| 11&#160;Apr&#160;23&#160;07:17&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 11&#160;May&#160;23&#160;07:17&#160;UTC | true | [view](CERTS/a49b98a40302636680164b79387a9331a3d5c5b67964944d7ebc4557a5cfa465/README.md) |
| 11&#160;Apr&#160;23&#160;07:53&#160;UTC | SHAKEN IDT America, Corp 363A | 11&#160;May&#160;23&#160;07:53&#160;UTC | true | [view](CERTS/e629dc749d41ebeaeaea1158ac4444e0658b879f867c40da60b7f36dd0fd4e41/README.md) |
| 11&#160;Apr&#160;23&#160;09:47&#160;UTC | SHAKEN InteractiveTel, LLC 920J | 11&#160;May&#160;23&#160;09:47&#160;UTC | true | [view](CERTS/d97da9e831cede1623a43b68db920dfb1949977dd454ce7965c960557a8250ea/README.md) |
| 11&#160;Apr&#160;23&#160;11:13&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 11&#160;May&#160;23&#160;11:13&#160;UTC | true | [view](CERTS/5a364ba553a405f442c533e2e1fe286e182f454beb54071e319d67020a37157b/README.md) |
| 11&#160;Apr&#160;23&#160;11:48&#160;UTC | SHAKEN IP Link Telecom Inc. 902J | 10&#160;Jul&#160;23&#160;11:48&#160;UTC | true | [view](CERTS/809b5296d9ca783497ea5d96de8caf84001ef4c3285b8447b0096462ef4e8aca/README.md) |
| 11&#160;Apr&#160;23&#160;13:29&#160;UTC | SHAKEN Threshold Communications Inc 563J | 11&#160;May&#160;23&#160;13:29&#160;UTC | true | [view](CERTS/463e99d8aca30949bad60e2608e5a755639d9af17ef4380f16845413f8e58506/README.md) |
| 11&#160;Apr&#160;23&#160;13:29&#160;UTC | SHAKEN Broadband Dynamics LLC 583j | 11&#160;May&#160;23&#160;13:29&#160;UTC | true | [view](CERTS/dc9c1cf5f7365bb2ce5060538a4645b845ecf3c60dcffa5dce78a4fe7e3996e3/README.md) |
| 11&#160;Apr&#160;23&#160;13:32&#160;UTC | SHAKEN Consolidated Communications 5113 | 11&#160;May&#160;23&#160;13:32&#160;UTC | true | [view](CERTS/88d6165d5ba3ada1d62023f75a054ab87b45edf090a363bb44606bcb0fab407d/README.md) |
| 11&#160;Apr&#160;23&#160;13:33&#160;UTC | SHAKEN Touchtone 683A | 11&#160;May&#160;23&#160;13:33&#160;UTC | true | [view](CERTS/320fba9380e4c13e99969a293eb8b10a9e964036a7d84365a1c89026783b907e/README.md) |
| 11&#160;Apr&#160;23&#160;13:33&#160;UTC | SHAKEN Apeiron Systems 012J | 11&#160;May&#160;23&#160;13:33&#160;UTC | true | [view](CERTS/2390cfc6f42ab452c011a5688a7d9c65638ea1a3822aba6aaf7b031a24e3e5c3/README.md) |
| 11&#160;Apr&#160;23&#160;13:34&#160;UTC | SHAKEN Fonative, Inc. 684J | 11&#160;May&#160;23&#160;13:34&#160;UTC | true | [view](CERTS/13c8e01d6bae66d8932646cfdd3924c3cdd8dc3ca87724a8d6aa0904ab56fb11/README.md) |
| 11&#160;Apr&#160;23&#160;13:36&#160;UTC | SHAKEN IPitomy 652J | 11&#160;May&#160;23&#160;13:36&#160;UTC | true | [view](CERTS/a2bbfb22d299bbb2eb6b52547ac10f305255227f79460b0d94ce08749a5ac669/README.md) |
| 11&#160;Apr&#160;23&#160;13:37&#160;UTC | SHAKEN Phone.com, Inc. 633J | 11&#160;May&#160;23&#160;13:37&#160;UTC | true | [view](CERTS/01f426d824293b5bd15b2872e0c3065b5e928ea27cc58f41a3b4144a9dd8c950/README.md) |
| 11&#160;Apr&#160;23&#160;13:38&#160;UTC | SHAKEN NETRIO LLC 020K | 11&#160;May&#160;23&#160;13:38&#160;UTC | true | [view](CERTS/f5763bb2ad2b95d44b90407f5943ad158f3b52b72cd017c3e74a37f9deb32e1f/README.md) |
| 11&#160;Apr&#160;23&#160;13:39&#160;UTC | SHAKEN Noble Systems Communications LLC 187J | 11&#160;May&#160;23&#160;13:39&#160;UTC | true | [view](CERTS/ed4466be0526c2ac43faacefad337ca7c3cd241707febdd56c17da0d4476f16a/README.md) |
| 11&#160;Apr&#160;23&#160;13:40&#160;UTC | SHAKEN Telcentris Inc. dba Voxox 696J | 11&#160;May&#160;23&#160;13:40&#160;UTC | true | [view](CERTS/2de864e013a1e5a1ba8c8bfc0cee08672f7205ee980345090dfabd50f596f88c/README.md) |
| 11&#160;Apr&#160;23&#160;13:42&#160;UTC | SHAKEN Airespring 996H | 11&#160;May&#160;23&#160;13:42&#160;UTC | true | [view](CERTS/caec316df6d5d933cbc1a2f5748b9793f01d8833d837e4719c814d144aaa2698/README.md) |
| 11&#160;Apr&#160;23&#160;13:42&#160;UTC | SHAKEN Nobelbiz, Inc. 596J | 11&#160;May&#160;23&#160;13:42&#160;UTC | true | [view](CERTS/8d5ef6778684bd26957ef12722b25e55c2a1cc7f9b0ad2480f314bd708fe0158/README.md) |
| 11&#160;Apr&#160;23&#160;13:46&#160;UTC | SHAKEN Momentum Telecom 1417 | 11&#160;May&#160;23&#160;13:46&#160;UTC | true | [view](CERTS/a4f794badafe954532dba52b87d6c339275407c720f11ce7757fd5a4be1fa3b1/README.md) |
| 11&#160;Apr&#160;23&#160;13:46&#160;UTC | SHAKEN Momentum Telecom 9157 | 11&#160;May&#160;23&#160;13:46&#160;UTC | true | [view](CERTS/58c96105cf8c966b829589941ea9941954397de3debc1dddb09961a27b57dc3b/README.md) |
| 11&#160;Apr&#160;23&#160;13:48&#160;UTC | SHAKEN Terra Nova Telecom 382G | 11&#160;May&#160;23&#160;13:48&#160;UTC | true | [view](CERTS/e39926635814d011128c5a7e1e887bc3ebb691e40e1c4136c0206202cc6828d8/README.md) |
| 11&#160;Apr&#160;23&#160;13:48&#160;UTC | SHAKEN Matrix 3058 | 11&#160;May&#160;23&#160;13:48&#160;UTC | true | [view](CERTS/27efce97c50f4b0a234a5e4585809ef66c03b58f22b82943ab30947889b45e69/README.md) |
| 11&#160;Apr&#160;23&#160;13:49&#160;UTC | SHAKEN Matrix 9451 | 11&#160;May&#160;23&#160;13:49&#160;UTC | true | [view](CERTS/8bab61657239d7c6f217e4bbb1ef249656e33f0638d1bebfadf4a377d2b98753/README.md) |
| 11&#160;Apr&#160;23&#160;13:51&#160;UTC | SHAKEN Matrix 7379 | 11&#160;May&#160;23&#160;13:51&#160;UTC | true | [view](CERTS/aa9d3786ead5411da0da1a9ba7493ff4de2353ca2d639f321041e4614a28acf1/README.md) |
| 11&#160;Apr&#160;23&#160;13:53&#160;UTC | SHAKEN PNG Telecommunications Inc 3395 | 11&#160;May&#160;23&#160;13:53&#160;UTC | true | [view](CERTS/9d5d32bbfd6cf0601bb2a7cddb9b788d84b3ea60c899ff46e0be71e6360822da/README.md) |
| 11&#160;Apr&#160;23&#160;13:55&#160;UTC | SHAKEN Magna5, LLC 3849 | 11&#160;May&#160;23&#160;13:55&#160;UTC | true | [view](CERTS/ee95764539dcdaf9774aad5ed8a6d5144085e1c4560e0ccdeee3c61fb0f12d89/README.md) |
| 11&#160;Apr&#160;23&#160;13:57&#160;UTC | SHAKEN Magna5, LLC 8249 | 11&#160;May&#160;23&#160;13:57&#160;UTC | true | [view](CERTS/33a629f0083f85f08ce912ef96c3fad2e7577af9ba693bb654fda534e5180f4b/README.md) |
| 11&#160;Apr&#160;23&#160;16:17&#160;UTC | SHAKEN NTC International, INC 016K | 11&#160;May&#160;23&#160;16:17&#160;UTC | true | [view](CERTS/d309e2491bf038bae327524b3c878e5e7963af84f2bff2d49d7873a8d1d75375/README.md) |
| 12&#160;Apr&#160;23&#160;04:42&#160;UTC | SHAKEN BareTelecom 864J | 12&#160;May&#160;23&#160;04:42&#160;UTC | true | [view](CERTS/5989b43f791f100afe63d5a83114d1a027c6f08809d0545fcda789542c7392e9/README.md) |
| 12&#160;Apr&#160;23&#160;05:02&#160;UTC | SHAKEN IDT America, Corp 363A | 12&#160;May&#160;23&#160;05:02&#160;UTC | true | [view](CERTS/cfe3fb2fc1d375ca5eb435a5f619c6367dad46f769d654ca712fad263ae86ac3/README.md) |
| 12&#160;Apr&#160;23&#160;09:42&#160;UTC | SHAKEN InteractiveTel, LLC 920J | 12&#160;May&#160;23&#160;09:42&#160;UTC | true | [view](CERTS/5989893e2e562a92e74bb32d0079935d0b5bebee46a5341560f6a47afde94f6f/README.md) |
| 12&#160;Apr&#160;23&#160;11:08&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 12&#160;May&#160;23&#160;11:08&#160;UTC | true | [view](CERTS/378e38c9e22e9139e880ab13147a1aa9b8d89e6a91e3ad8f942150faa47ff9dd/README.md) |
| 12&#160;Apr&#160;23&#160;17:13&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 12&#160;May&#160;23&#160;17:13&#160;UTC | true | [view](CERTS/3f143aae86c433e9cdf5c6c20071a556a8c498d2f32c03fc2cf2a36ee967064f/README.md) |
| 13&#160;Apr&#160;23&#160;04:37&#160;UTC | SHAKEN BareTelecom 864J | 13&#160;May&#160;23&#160;04:37&#160;UTC | true | [view](CERTS/b11644f7681ba881b275e373720bb8b2d585c666f4a67636b1130523ed84df8f/README.md) |
| 13&#160;Apr&#160;23&#160;07:07&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 13&#160;May&#160;23&#160;07:07&#160;UTC | true | [view](CERTS/18d48ab7c0811c9f49c6c13ab834e70610036b1b033edb90350847e0a82a7b3f/README.md) |
| 13&#160;Apr&#160;23&#160;08:01&#160;UTC | SHAKEN IPSBS Managed Services LLC 828J | 13&#160;May&#160;23&#160;08:01&#160;UTC | true | [view](CERTS/b57cdd841fb80aee1bf28c2b226e88c7727ae97ce4eba9a7b36e311a5f7353da/README.md) |
| 13&#160;Apr&#160;23&#160;11:03&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 13&#160;May&#160;23&#160;11:03&#160;UTC | true | [view](CERTS/68b97352c6473612667dbab676b87b6f0c513dd4878128c351fe8247e6c0eee4/README.md) |
| 13&#160;Apr&#160;23&#160;17:08&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 13&#160;May&#160;23&#160;17:08&#160;UTC | true | [view](CERTS/76c2b865cff94cfe23c74c286774c2e479510ef01fb73b5ab623e3f79a1c2742/README.md) |
| 14&#160;Apr&#160;23&#160;04:32&#160;UTC | SHAKEN BareTelecom 864J | 14&#160;May&#160;23&#160;04:32&#160;UTC | true | [view](CERTS/049b442e6a151c894bb97cea6331deb8b36c57b022581fa2306a48d138437e73/README.md) |
| 14&#160;Apr&#160;23&#160;04:54&#160;UTC | SHAKEN IDT America, Corp 363A | 14&#160;May&#160;23&#160;04:54&#160;UTC | true | [view](CERTS/56afdcc4059e7be5f31828cd8728136d1355453c666d2b679697ddc9f4b1a603/README.md) |
| 14&#160;Apr&#160;23&#160;07:02&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 14&#160;May&#160;23&#160;07:02&#160;UTC | true | [view](CERTS/467841c6cbeb607d371eba9597fa4e6738d9b2e3e15892858b33c97d139b728b/README.md) |
| 14&#160;Apr&#160;23&#160;07:56&#160;UTC | SHAKEN IPSBS Managed Services LLC 828J | 14&#160;May&#160;23&#160;07:56&#160;UTC | true | [view](CERTS/9de9d41d6658ebfab0c35c5c37fb9fead1f50d914e3fb12b4e15cd232892b7c6/README.md) |
| 14&#160;Apr&#160;23&#160;09:32&#160;UTC | SHAKEN InteractiveTel, LLC 920J | 14&#160;May&#160;23&#160;09:32&#160;UTC | true | [view](CERTS/08d647100a86ab7b39ffa91c5a9704f4f6c1b3315671f1a096ece150047706f1/README.md) |
| 14&#160;Apr&#160;23&#160;10:58&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 14&#160;May&#160;23&#160;10:58&#160;UTC | true | [view](CERTS/ef33331194f42f5075e7e8265047574d3662238c87937586385855aa15b01915/README.md) |
| 14&#160;Apr&#160;23&#160;17:03&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 14&#160;May&#160;23&#160;17:03&#160;UTC | true | [view](CERTS/c4af575d0a6c59d99189a9d2e8d71c0f70eec8ae172671ecf69192b68df882a8/README.md) |
| 14&#160;Apr&#160;23&#160;22:55&#160;UTC | SHAKEN i3 Broadband 5800 | 14&#160;May&#160;23&#160;22:55&#160;UTC | true | [view](CERTS/3678eafc6767d931af4fadeb6c466c0dc782bfe62c61f9bf7f5141480e0f231a/README.md) |
| 15&#160;Apr&#160;23&#160;04:27&#160;UTC | SHAKEN BareTelecom 864J | 15&#160;May&#160;23&#160;04:27&#160;UTC | true | [view](CERTS/d51b15b455cc0307e8d0f10fe958f67240e68b284eda35f47ee6a11a746f2153/README.md) |
| 15&#160;Apr&#160;23&#160;10:53&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 15&#160;May&#160;23&#160;10:53&#160;UTC | true | [view](CERTS/5fc140665bc922bf72bb44fd72b2ce8cb3a88811e67f3a505a71bb9c5e80501b/README.md) |
| 15&#160;Apr&#160;23&#160;16:58&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 15&#160;May&#160;23&#160;16:58&#160;UTC | true | [view](CERTS/dc217cb66dcf35292c320b70dbb36e316796fdbe359868b4bfcd5c55ca8f2b40/README.md) |
| 16&#160;Apr&#160;23&#160;04:44&#160;UTC | SHAKEN IDT America, Corp 363A | 16&#160;May&#160;23&#160;04:44&#160;UTC | true | [view](CERTS/455d6d513dac32d965ac7afcfaacf47da58ba97bfff05a7c7dd34bc9377422f1/README.md) |
| 16&#160;Apr&#160;23&#160;16:53&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 16&#160;May&#160;23&#160;16:53&#160;UTC | true | [view](CERTS/599af887200337a7ae884e060747641a30dd9cd7a30a6e6d728e6911f165ad2b/README.md) |
| 17&#160;Apr&#160;23&#160;04:39&#160;UTC | SHAKEN IDT America, Corp 363A | 17&#160;May&#160;23&#160;04:39&#160;UTC | true | [view](CERTS/44b17616d9477d1bef92eb03f20d29541cc5c4724cf3fccc3304d751b581b1f8/README.md) |
| 17&#160;Apr&#160;23&#160;06:47&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 17&#160;May&#160;23&#160;06:47&#160;UTC | true | [view](CERTS/e885b36cba61aca8877712c5bf5fb29fe07967b1b726966a161996a752c1073a/README.md) |
| 17&#160;Apr&#160;23&#160;07:41&#160;UTC | SHAKEN IPSBS Managed Services LLC 828J | 17&#160;May&#160;23&#160;07:41&#160;UTC | true | [view](CERTS/a0198e43a853bd950d95619bb79a8ad649be0c8a7558d3a16bd58a71b0ca8595/README.md) |
| 17&#160;Apr&#160;23&#160;09:17&#160;UTC | SHAKEN InteractiveTel, LLC 920J | 17&#160;May&#160;23&#160;09:17&#160;UTC | true | [view](CERTS/0d0820b75fc14b2f83526476430acf951d4e7aa13e53035ce76818835c0ae040/README.md) |
| 17&#160;Apr&#160;23&#160;10:43&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 17&#160;May&#160;23&#160;10:43&#160;UTC | true | [view](CERTS/299bc8fce2ff01106648d02f6b000e96e39680ceff05350d628552cf4b1f86b7/README.md) |
| 17&#160;Apr&#160;23&#160;14:47&#160;UTC | SHAKEN ONE OWL TELECOM INC 412K | 17&#160;May&#160;23&#160;14:47&#160;UTC | true | [view](CERTS/f2858f82b8d5b7685abeba4bb9d1e6dcfc71516c50575025e9924e9b6dec0102/README.md) |
| 17&#160;Apr&#160;23&#160;16:48&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 17&#160;May&#160;23&#160;16:48&#160;UTC | true | [view](CERTS/93ab0cfeaccfbb12d26d4c1330e9dd5f1d6b3713b3fee5343a52520f1059d522/README.md) |
| 18&#160;Apr&#160;23&#160;04:34&#160;UTC | SHAKEN IDT America, Corp 363A | 18&#160;May&#160;23&#160;04:34&#160;UTC | true | [view](CERTS/862efce3dd457a61b6f4a017c3d1bf964a15fd83e4f3ac8ce298cab82cee6222/README.md) |
| 18&#160;Apr&#160;23&#160;06:42&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 18&#160;May&#160;23&#160;06:42&#160;UTC | true | [view](CERTS/e7757e966b2ff6eabcd770b145bfd07c76fa4c20233f723081cd3e1ebdf5ddea/README.md) |
| 18&#160;Apr&#160;23&#160;07:36&#160;UTC | SHAKEN IPSBS Managed Services LLC 828J | 18&#160;May&#160;23&#160;07:36&#160;UTC | true | [view](CERTS/31820d7fcac185227c179a2e9b6a19c5b9edc5aca24d3ab0ed4ff3e59a9e7b83/README.md) |
| 18&#160;Apr&#160;23&#160;10:38&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 18&#160;May&#160;23&#160;10:38&#160;UTC | true | [view](CERTS/d82c392b54136155bafffd27408119abcd54d169fe673eb67a0db9f089152d9f/README.md) |
| 18&#160;Apr&#160;23&#160;11:13&#160;UTC | SHAKEN Zray Technologies Corporation 862J | 18&#160;May&#160;23&#160;11:13&#160;UTC | true | [view](CERTS/9e0f6032c4b9c16c29e77132f4542f73ff80301c928cad3433959469bfe9cfe4/README.md) |
| 18&#160;Apr&#160;23&#160;16:43&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 18&#160;May&#160;23&#160;16:43&#160;UTC | true | [view](CERTS/2ab0f23aa0c6d3ff3e42812b5ea28e7bc5c5d34ef5dea30201635276efedcc43/README.md) |
| 19&#160;Apr&#160;23&#160;04:07&#160;UTC | SHAKEN BareTelecom 864J | 19&#160;May&#160;23&#160;04:07&#160;UTC | true | [view](CERTS/d3abc704d2d7aae81318dac0e0021aed4c2bc9e595dab457fec72c6bca21b582/README.md) |
| 19&#160;Apr&#160;23&#160;04:29&#160;UTC | SHAKEN IDT America, Corp 363A | 19&#160;May&#160;23&#160;04:29&#160;UTC | true | [view](CERTS/fadec2fe5dd6a58f0af87b605a2559f14f79c28ae8eae72b88851edd0d5011ba/README.md) |
| 19&#160;Apr&#160;23&#160;06:07&#160;UTC | SHAKEN Telxio Networks 492K | 18&#160;Apr&#160;24&#160;06:07&#160;UTC | true | [view](CERTS/010c6a74330323c20ceb343b1de3a1e3248b4a3926c9ad2ed53f02b02399d241/README.md) |
| 19&#160;Apr&#160;23&#160;06:37&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 19&#160;May&#160;23&#160;06:37&#160;UTC | true | [view](CERTS/3a215575e75a6446fa11fb2ea841044a7370eb4472ec1c63dea5e483c3743805/README.md) |
| 19&#160;Apr&#160;23&#160;09:07&#160;UTC | SHAKEN InteractiveTel, LLC 920J | 19&#160;May&#160;23&#160;09:07&#160;UTC | true | [view](CERTS/577f482ae0f6f33396bc68ee1e908160361f971fc96339dfbd7afe5e895c0ff3/README.md) |
| 19&#160;Apr&#160;23&#160;13:17&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 19&#160;May&#160;23&#160;13:17&#160;UTC | true | [view](CERTS/de6b08d2f84960cc03f096bdfb1865114c0866630939d6351eb5d7e776d89965/README.md) |
| 20&#160;Apr&#160;23&#160;04:02&#160;UTC | SHAKEN BareTelecom 864J | 20&#160;May&#160;23&#160;04:02&#160;UTC | true | [view](CERTS/e23c9463c3aa8e085aa6daac4f091a91903ae705cd763e1106005f4e1cce79f3/README.md) |
| 20&#160;Apr&#160;23&#160;04:24&#160;UTC | SHAKEN IDT America, Corp 363A | 20&#160;May&#160;23&#160;04:24&#160;UTC | true | [view](CERTS/5c000d16835503afaece38c29f92413f6edfcde70c950051bdce199c3e2095a4/README.md) |
| 20&#160;Apr&#160;23&#160;06:32&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 20&#160;May&#160;23&#160;06:32&#160;UTC | true | [view](CERTS/57b9d89f4cf77e008e86a0b7c85e0ee2c9efbda0ab4d01e07f6588d540ed50ed/README.md) |
| 20&#160;Apr&#160;23&#160;11:03&#160;UTC | SHAKEN Zray Technologies Corporation 862J | 20&#160;May&#160;23&#160;11:03&#160;UTC | true | [view](CERTS/c3d404e0d115f8a8fbf0c4f6c3a4bc93b6029e0a239b514dcddf1dc8cb4439e9/README.md) |
| 20&#160;Apr&#160;23&#160;13:40&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 20&#160;May&#160;23&#160;13:40&#160;UTC | true | [view](CERTS/5abef4e6c3960680c3628196fbb973f19407f0b83bbb1518bcfd33a5f3a7b23a/README.md) |
| 20&#160;Apr&#160;23&#160;20:25&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 20&#160;May&#160;23&#160;20:25&#160;UTC | true | [view](CERTS/3dfb0f0e68fe3c089905e27ab8ffb1fef83fca8ae277f1efef8d82dcaeff6e5e/README.md) |
| 20&#160;Apr&#160;23&#160;21:10&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 20&#160;May&#160;23&#160;21:10&#160;UTC | true | [view](CERTS/b752550689a1e4e76fd48f219d27495c51989ebd123801ca65444342c086f993/README.md) |
| 21&#160;Apr&#160;23&#160;03:57&#160;UTC | SHAKEN BareTelecom 864J | 21&#160;May&#160;23&#160;03:57&#160;UTC | true | [view](CERTS/4483e54eadbfc4ed7c8ff24752df5fcda064771ae680b978a72074d7668db105/README.md) |
| 21&#160;Apr&#160;23&#160;04:19&#160;UTC | SHAKEN IDT America, Corp 363A | 21&#160;May&#160;23&#160;04:19&#160;UTC | true | [view](CERTS/bdd062edeef0982455040eb880f801e6bca5046c66e895352afb3d34b1d5ea60/README.md) |
| 21&#160;Apr&#160;23&#160;16:28&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 21&#160;May&#160;23&#160;16:28&#160;UTC | true | [view](CERTS/289f045068f0ad241bca9e6b6c433190e2fa7df68acf7f92f211ea49478a7f81/README.md) |
| 25&#160;Apr&#160;23&#160;00:34&#160;UTC | SHAKEN Cloud Connect LLC 455K | 24&#160;Apr&#160;24&#160;00:34&#160;UTC | true | [view](CERTS/4619e88c821112db7427efe1257b2c699c668d68bb5a46ceb2658f5157ad11b8/README.md) |

#### CA Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 21&#160;Aug&#160;20&#160;01:22&#160;UTC | SHAKEN Sansay Root CA US | 16&#160;Aug&#160;40&#160;01:22&#160;UTC | true | [view](CERTS/8356d251b255f4ac3fd108bb79be4c02dcea2d3b13d138892bdb3a70cbe6a343/README.md) |
| 02&#160;Sep&#160;22&#160;20:53&#160;UTC | SHAKEN Sansay Intermediate CA US WEST 1 | 31&#160;Aug&#160;29&#160;20:53&#160;UTC | true | [view](CERTS/4b1dfdba2b1e4bbffbf900a20f1f6f7befbef0008b963e4922a64469cb97d24b/README.md) |


Generated: 28 Apr 23 02:17 UTC