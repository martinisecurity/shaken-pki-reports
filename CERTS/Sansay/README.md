# STIR/SHAKEN CA Ecosystem Compliance

## Sansay

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 564 certificates were included in the corpus being tested
- 9 certificates in the corpus were skipped because they are duplicates
- 403 certificates in the corpus were skipped because they are expired
- 1 certificates in the corpus were skipped because they are not currently trusted
- 151 certificates being tested against the remaining rules
- 4.92 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 100.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 113 days is the average remaining validity for the certificates in the corpus
- 113 days is the average initial validity for the certificates in the corpus
- 111 certificates expire in the next 30 days
- 2.01 average number of unexpired certificates per OCN observed
- 75 unique OCNs observed in unexpired and valid certificate corpus

| Instances | Test | Source |
|-----------|------|--------|
| 58 | [e_atis_certificate_policies](ISSUES/e_atis_certificate_policies/README.md) | ATIS1000080 |
| 151 | [e_atis_subject_cn](ISSUES/e_atis_subject_cn/README.md) | ATIS1000080 |
| 151 | [e_us_cp_ambiguous_identifier](ISSUES/e_us_cp_ambiguous_identifier/README.md) | US_SHAKEN_CP |
| 151 | [e_us_cp_subject_sn](ISSUES/e_us_cp_subject_sn/README.md) | US_SHAKEN_CP |
| 81 | [w_cp_1_3_subject_email](ISSUES/w_cp_1_3_subject_email/README.md) | US_SHAKEN_CP |
| 151 | [w_shaken_subject_rdn_unknown](ISSUES/w_shaken_subject_rdn_unknown/README.md) | SHAKEN_PKI_BEST_PRACTICES |

#### CA Certificates

- 4 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 1 certificates in the corpus were skipped because they are expired
- 1 certificates in the corpus were skipped because they are not currently trusted
- 2 certificates being tested against the remaining rules
- 1.00 issues on average found in unexpired, trusted, and non-compliant certificates
- 0.00% of certificates contain one or more Error level issue
- 100.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 50.00% of certificates are too old to be assessed against currently enforced expectations
- 5397 days is the average remaining validity for the certificates in the corpus
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
| 13&#160;Mar&#160;23&#160;04:30&#160;UTC | SHAKEN 1stPoint Communications, LLC 463G | 12&#160;Apr&#160;23&#160;04:30&#160;UTC | true | [view](CERTS/615a103854056f8ae5dd7794ac2f621c081637ca1e6d1877ae9c1413e08e9f60/README.md) |
| 13&#160;Mar&#160;23&#160;06:39&#160;UTC | SHAKEN Primo Dialler LLC 249K | 27&#160;Apr&#160;23&#160;06:39&#160;UTC | true | [view](CERTS/8b881d01de770031b18500ca0a7b8f5ed3490281d68f14b993c7315ddddac4d9/README.md) |
| 13&#160;Mar&#160;23&#160;11:56&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 12&#160;Apr&#160;23&#160;11:56&#160;UTC | true | [view](CERTS/0fab85834a79c866ee836d27a69c6bdd841fa5af9e9c8a239b7d4a3e4877ba0e/README.md) |
| 13&#160;Mar&#160;23&#160;12:02&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 12&#160;Apr&#160;23&#160;12:02&#160;UTC | true | [view](CERTS/c1f727cfd1ec0fa19885890370ee62a0b1e75b54386f814a983a64b1c78ec3e9/README.md) |
| 13&#160;Mar&#160;23&#160;12:12&#160;UTC | SHAKEN InteractiveTel, LLC 920J | 12&#160;Apr&#160;23&#160;12:12&#160;UTC | true | [view](CERTS/5d6bedbf7edde2a3017283545f1ee97c55938df2cbf556e99742abe88373fdf5/README.md) |
| 13&#160;Mar&#160;23&#160;19:36&#160;UTC | SHAKEN i3 Broadband 5800 | 12&#160;Apr&#160;23&#160;19:36&#160;UTC | true | [view](CERTS/112940a1e0a8559b4620eaae17dd7d518b6c8d1003164f7e5f62ab64d058efe4/README.md) |
| 14&#160;Mar&#160;23&#160;07:07&#160;UTC | SHAKEN BareTelecom 864J | 13&#160;Apr&#160;23&#160;07:07&#160;UTC | true | [view](CERTS/655d3a655d526164359d2867a16938a2d62c40a37d4d7400a67e0033a11b8f2c/README.md) |
| 14&#160;Mar&#160;23&#160;10:31&#160;UTC | SHAKEN IPSBS Managed Services LLC 828J | 13&#160;Apr&#160;23&#160;10:31&#160;UTC | true | [view](CERTS/12e16a0e54b9d3ed9ed448d4522e98024308cf4ec4506e597960536c05e335ba/README.md) |
| 14&#160;Mar&#160;23&#160;11:51&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 13&#160;Apr&#160;23&#160;11:51&#160;UTC | true | [view](CERTS/585ee1cf6cc92a0d29d543206d72848835f3947ea360079cc6a82b73e3cad300/README.md) |
| 14&#160;Mar&#160;23&#160;11:57&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 13&#160;Apr&#160;23&#160;11:57&#160;UTC | true | [view](CERTS/be6748ddcc775e6adfa4b17208473356fdc68d4a404119333e13b09e1e3f1225/README.md) |
| 15&#160;Mar&#160;23&#160;04:20&#160;UTC | SHAKEN 1stPoint Communications, LLC 463G | 14&#160;Apr&#160;23&#160;04:20&#160;UTC | true | [view](CERTS/007c9e70e3658bd731ec14a1095176b2b1b0a6691cd405552b94aadea91415e5/README.md) |
| 15&#160;Mar&#160;23&#160;07:02&#160;UTC | SHAKEN BareTelecom 864J | 14&#160;Apr&#160;23&#160;07:02&#160;UTC | true | [view](CERTS/9028747a825477319810035997b9ddbcc62b3039f5e87f04a015665f66df0a57/README.md) |
| 15&#160;Mar&#160;23&#160;11:52&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 14&#160;Apr&#160;23&#160;11:52&#160;UTC | true | [view](CERTS/d57cfaf295c96b4454b5b6e2c10821da2e8a69eb1894129ba725fb52dd1aa150/README.md) |
| 15&#160;Mar&#160;23&#160;18:32&#160;UTC | SHAKEN NTC International, INC 016K | 14&#160;Apr&#160;23&#160;18:32&#160;UTC | true | [view](CERTS/b08628c607378de238f8e52e569040ef465e5ae04efd19e0ca7ad7a301149ca5/README.md) |
| 16&#160;Mar&#160;23&#160;06:57&#160;UTC | SHAKEN BareTelecom 864J | 15&#160;Apr&#160;23&#160;06:57&#160;UTC | true | [view](CERTS/50d980fcfeda5367337ad960a7e70097ed80b1c5bf50705f822c09d8f76cc0fe/README.md) |
| 16&#160;Mar&#160;23&#160;10:21&#160;UTC | SHAKEN IPSBS Managed Services LLC 828J | 15&#160;Apr&#160;23&#160;10:21&#160;UTC | true | [view](CERTS/d3a84294c32712b3c0eee32820295ab7b037ee9d93f163705f94770f75b94802/README.md) |
| 16&#160;Mar&#160;23&#160;11:41&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 15&#160;Apr&#160;23&#160;11:41&#160;UTC | true | [view](CERTS/a73d6d4986649582ac833bd9b812750a3648487e7ef3acb4a34425c287929b8f/README.md) |
| 17&#160;Mar&#160;23&#160;06:52&#160;UTC | SHAKEN BareTelecom 864J | 16&#160;Apr&#160;23&#160;06:52&#160;UTC | true | [view](CERTS/8f015e9883111a35a4a46f5a6ad25c45320f1e8020ee4116d69beda6290d118d/README.md) |
| 17&#160;Mar&#160;23&#160;08:46&#160;UTC | SHAKEN Doylestown Communications, Inc 849C | 16&#160;Apr&#160;23&#160;08:46&#160;UTC | true | [view](CERTS/addacc85cf35be470aa7ef54fe602ec79eb6fb6132d898f69490c1b19638a6fa/README.md) |
| 17&#160;Mar&#160;23&#160;11:36&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 16&#160;Apr&#160;23&#160;11:36&#160;UTC | true | [view](CERTS/65f35e1dc4031857191a80585d21d32d9f43c9543966394daf796c6a31cb11f2/README.md) |
| 17&#160;Mar&#160;23&#160;18:22&#160;UTC | SHAKEN NTC International, INC 016K | 16&#160;Apr&#160;23&#160;18:22&#160;UTC | true | [view](CERTS/3e39bc80eb28fbd556bf37b4d4f0c002978fc52ac04f266f4374c4fb8c07eb1a/README.md) |
| 18&#160;Mar&#160;23&#160;06:04&#160;UTC | SHAKEN ONE OWL TELECOM INC 412K | 17&#160;Apr&#160;23&#160;06:04&#160;UTC | true | [view](CERTS/b0efcbe3cad70bc6c1f988feada7740e5c1edfc5ace698828297b383257f426a/README.md) |
| 18&#160;Mar&#160;23&#160;06:47&#160;UTC | SHAKEN BareTelecom 864J | 17&#160;Apr&#160;23&#160;06:47&#160;UTC | true | [view](CERTS/50ae08ab835ff359907f9cbcb721eeeb62c3e746bf291806120adccff84616ca/README.md) |
| 18&#160;Mar&#160;23&#160;11:47&#160;UTC | SHAKEN InteractiveTel, LLC 920J | 17&#160;Apr&#160;23&#160;11:47&#160;UTC | true | [view](CERTS/5ba7d717d049a5b32aef4e3034ea5c517535529fe30f111e09663b6142a4a658/README.md) |
| 18&#160;Mar&#160;23&#160;13:13&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 17&#160;Apr&#160;23&#160;13:13&#160;UTC | true | [view](CERTS/ab5801f8842c40a14a3c355ee89fcb508c82fe1c26abfe4d99aa5d795bece680/README.md) |
| 18&#160;Mar&#160;23&#160;18:17&#160;UTC | SHAKEN NTC International, INC 016K | 17&#160;Apr&#160;23&#160;18:17&#160;UTC | true | [view](CERTS/4e5fd575827f53757f94aba5eba6c4692ea95864cb2170a3615edddd42a24d62/README.md) |
| 19&#160;Mar&#160;23&#160;00:31&#160;UTC | SHAKEN Every1 Telecom 486K | 18&#160;Mar&#160;24&#160;00:31&#160;UTC | true | [view](CERTS/5cdfbb1a416083096dfef10c75a2b26a08d8fd5593d8ea9ceae0d70d878a97d1/README.md) |
| 19&#160;Mar&#160;23&#160;06:42&#160;UTC | SHAKEN BareTelecom 864J | 18&#160;Apr&#160;23&#160;06:42&#160;UTC | true | [view](CERTS/68ceca7b1b845a549c44a681f30e6eddab24f3f428f4fadd9645b480dde5e44b/README.md) |
| 19&#160;Mar&#160;23&#160;13:08&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 18&#160;Apr&#160;23&#160;13:08&#160;UTC | true | [view](CERTS/31c834f828d38cf423f8dd5dbfcc61a548fa6224babcb20c21be9f1b55aa8c1f/README.md) |
| 20&#160;Mar&#160;23&#160;04:53&#160;UTC | SHAKEN Telxio Networks 492K | 19&#160;Apr&#160;23&#160;04:53&#160;UTC | true | [view](CERTS/b251c83f426c9d8062d824f377fb0fee29b3368040f6a41015b73d268bc78bcd/README.md) |
| 20&#160;Mar&#160;23&#160;06:37&#160;UTC | SHAKEN BareTelecom 864J | 19&#160;Apr&#160;23&#160;06:37&#160;UTC | true | [view](CERTS/3f658be27381a5988ec45d5c24d65c9b67c5617979198421b4c18df6f5ffd6fa/README.md) |
| 20&#160;Mar&#160;23&#160;11:21&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 19&#160;Apr&#160;23&#160;11:21&#160;UTC | true | [view](CERTS/9357cc586e9c5f90984d11b6603e1e9bdf3ffc38bb542110eeb87ba52c1d516e/README.md) |
| 20&#160;Mar&#160;23&#160;13:03&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 19&#160;Apr&#160;23&#160;13:03&#160;UTC | true | [view](CERTS/23fff3a8d61d6d1a39ebfed7e8fb9f7be441e64b17dd733af05970fa2e99d897/README.md) |
| 20&#160;Mar&#160;23&#160;18:07&#160;UTC | SHAKEN NTC International, INC 016K | 19&#160;Apr&#160;23&#160;18:07&#160;UTC | true | [view](CERTS/23c173c0257bd0d8813326a9d2441ce48c5ab626c17a7537b182283c2a2ec001/README.md) |
| 21&#160;Mar&#160;23&#160;06:32&#160;UTC | SHAKEN BareTelecom 864J | 20&#160;Apr&#160;23&#160;06:32&#160;UTC | true | [view](CERTS/e3f0d81bfbc176a40bcad65038b55518bd110280cc02c0ec6294e4e51b27cd6e/README.md) |
| 21&#160;Mar&#160;23&#160;09:56&#160;UTC | SHAKEN IPSBS Managed Services LLC 828J | 20&#160;Apr&#160;23&#160;09:56&#160;UTC | true | [view](CERTS/e31a5c9c1c5ee72d1738275f28fc1113320b798b8206a3f1126e76750034f4c6/README.md) |
| 21&#160;Mar&#160;23&#160;11:16&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 20&#160;Apr&#160;23&#160;11:16&#160;UTC | true | [view](CERTS/e4e8578ebb523c9aa6ceb565f12ab5c7f09eb05726378a8296aa1fd834008fd2/README.md) |
| 21&#160;Mar&#160;23&#160;12:58&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 20&#160;Apr&#160;23&#160;12:58&#160;UTC | true | [view](CERTS/5fd5419ee38e3a75a47424c3acbfbee0c87f0e24c1639a40197a31900fd15dd8/README.md) |
| 21&#160;Mar&#160;23&#160;18:02&#160;UTC | SHAKEN NTC International, INC 016K | 20&#160;Apr&#160;23&#160;18:02&#160;UTC | true | [view](CERTS/85a448d0f10d0061438b4f9ce1b2d86f05ca679839bb1fc5fe3e01d9110fbe93/README.md) |
| 22&#160;Mar&#160;23&#160;06:27&#160;UTC | SHAKEN BareTelecom 864J | 21&#160;Apr&#160;23&#160;06:27&#160;UTC | true | [view](CERTS/0f650b71b323e16240f7e08ce76fcb75ac32b02e64bd30b8d8cf3ca8db50647c/README.md) |
| 22&#160;Mar&#160;23&#160;09:51&#160;UTC | SHAKEN IPSBS Managed Services LLC 828J | 21&#160;Apr&#160;23&#160;09:51&#160;UTC | true | [view](CERTS/e1922e73908400a86041d102b900c0d599454742250b7171247f109481fa9f1c/README.md) |
| 22&#160;Mar&#160;23&#160;11:11&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 21&#160;Apr&#160;23&#160;11:11&#160;UTC | true | [view](CERTS/0040a0ef8f5a4d56f2dbf90e9b202b419997ce01296f4b24a20b4f6cd95a9efb/README.md) |
| 22&#160;Mar&#160;23&#160;11:27&#160;UTC | SHAKEN InteractiveTel, LLC 920J | 21&#160;Apr&#160;23&#160;11:27&#160;UTC | true | [view](CERTS/1fdccdf383a21d683eac69aa592786fbeb3c297d8b7831fa9970d903cfb01905/README.md) |
| 22&#160;Mar&#160;23&#160;12:53&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 21&#160;Apr&#160;23&#160;12:53&#160;UTC | true | [view](CERTS/305f172b6bfdb908b5ee29dd5e280f92d689f022b8cca22ad17bad1fbfa3ad1f/README.md) |
| 23&#160;Mar&#160;23&#160;06:22&#160;UTC | SHAKEN BareTelecom 864J | 22&#160;Apr&#160;23&#160;06:22&#160;UTC | true | [view](CERTS/82e8e073f9d25ac50a212164a0e13d24ce1446becfb24641bbd35b617578265e/README.md) |
| 23&#160;Mar&#160;23&#160;09:46&#160;UTC | SHAKEN IPSBS Managed Services LLC 828J | 22&#160;Apr&#160;23&#160;09:46&#160;UTC | true | [view](CERTS/38fc052d37c56a0a27ac6011c49f2a94cd5137e2bb455eb069dcf7dbb128c9c5/README.md) |
| 23&#160;Mar&#160;23&#160;11:06&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 22&#160;Apr&#160;23&#160;11:06&#160;UTC | true | [view](CERTS/b62f321bcc75e6defb19f3afec8c29643ec6843db5b8e0323f2a323439f01d91/README.md) |
| 23&#160;Mar&#160;23&#160;11:22&#160;UTC | SHAKEN InteractiveTel, LLC 920J | 22&#160;Apr&#160;23&#160;11:22&#160;UTC | true | [view](CERTS/63f43998c5b808cfe047b1aa7f09405baad9f287cf00b6d7a927593a6c5cb692/README.md) |
| 23&#160;Mar&#160;23&#160;12:48&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 22&#160;Apr&#160;23&#160;12:48&#160;UTC | true | [view](CERTS/30979a50c48bbace3a85cacbb82b6c468549ac53d84e780ab24441ba60bc2bb6/README.md) |
| 23&#160;Mar&#160;23&#160;13:23&#160;UTC | SHAKEN Zray Technologies Corporation 862J | 22&#160;Apr&#160;23&#160;13:23&#160;UTC | true | [view](CERTS/3cdc4fd869aa1b8fd4da6708de8acab414f4ccb9b1a0ca3bc221173cd7e2313a/README.md) |
| 24&#160;Mar&#160;23&#160;06:17&#160;UTC | SHAKEN BareTelecom 864J | 23&#160;Apr&#160;23&#160;06:17&#160;UTC | true | [view](CERTS/78143915082e4f6ab37b77b2b1cf79d58b4c0a91ab363bf8bca0cf9b3ec716c0/README.md) |
| 24&#160;Mar&#160;23&#160;11:01&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 23&#160;Apr&#160;23&#160;11:01&#160;UTC | true | [view](CERTS/30b1d17afa15e4c9070f1eb10aef8db306c806cad67f293817e7ba281029c4f8/README.md) |
| 24&#160;Mar&#160;23&#160;16:01&#160;UTC | SHAKEN Global Net Holdings Inc 306K | 23&#160;Apr&#160;23&#160;16:01&#160;UTC | true | [view](CERTS/e6b5ef6ca2810a29de5b2e32df2e79ddedabd3429f12d61fe2e7bb17d3814b41/README.md) |
| 24&#160;Mar&#160;23&#160;17:47&#160;UTC | SHAKEN NTC International, INC 016K | 23&#160;Apr&#160;23&#160;17:47&#160;UTC | true | [view](CERTS/937643484ee3e763749f606c0db9234e43e7055b14e0137ae777cb07519c8849/README.md) |
| 25&#160;Mar&#160;23&#160;06:12&#160;UTC | SHAKEN BareTelecom 864J | 24&#160;Apr&#160;23&#160;06:12&#160;UTC | true | [view](CERTS/ea7d6c7815c3cedeb11aaa2423ad3334c3cd4092c3fe19a5a333b1f324c48729/README.md) |
| 25&#160;Mar&#160;23&#160;11:12&#160;UTC | SHAKEN InteractiveTel, LLC 920J | 24&#160;Apr&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/0a7d4c0ec5db8542234645c008be7be0bf39fb04ee4855958cc39459c97d5cca/README.md) |
| 26&#160;Mar&#160;23&#160;06:07&#160;UTC | SHAKEN BareTelecom 864J | 25&#160;Apr&#160;23&#160;06:07&#160;UTC | true | [view](CERTS/46346421a266535c1ca36512526fceb06dd7d359b571acdc6b45fa8f9c80f1c7/README.md) |
| 26&#160;Mar&#160;23&#160;12:33&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 25&#160;Apr&#160;23&#160;12:33&#160;UTC | true | [view](CERTS/5fd5b2c4b66526673543f04cb05675ce5941516e0b53c7ce7a82c12d271bfbc6/README.md) |
| 26&#160;Mar&#160;23&#160;17:37&#160;UTC | SHAKEN NTC International, INC 016K | 25&#160;Apr&#160;23&#160;17:37&#160;UTC | true | [view](CERTS/a52dddf322a3f1ef09b52af1cde6e6829810008b5877b79dcd2385945a9e607c/README.md) |
| 27&#160;Mar&#160;23&#160;06:02&#160;UTC | SHAKEN BareTelecom 864J | 26&#160;Apr&#160;23&#160;06:02&#160;UTC | true | [view](CERTS/b59d9a92adcf3d1cce41d230053ac4b2c53dca55696bd96259fe00006eeef4f9/README.md) |
| 27&#160;Mar&#160;23&#160;09:26&#160;UTC | SHAKEN IPSBS Managed Services LLC 828J | 26&#160;Apr&#160;23&#160;09:26&#160;UTC | true | [view](CERTS/1d2165c03f04ca88c384b138fce1b0fd6652fad18c5ee027b07a3c5c93f8536f/README.md) |
| 27&#160;Mar&#160;23&#160;10:46&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 26&#160;Apr&#160;23&#160;10:46&#160;UTC | true | [view](CERTS/73c8ecae53fe9ef9a5d6c581388a25911fb3a2537a5a2e290be62f6d3ef1bd48/README.md) |
| 27&#160;Mar&#160;23&#160;12:28&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 26&#160;Apr&#160;23&#160;12:28&#160;UTC | true | [view](CERTS/332a8a97f5595f49f05be7852b97c789cb82b6371870de5b3704f2653b823cd0/README.md) |
| 27&#160;Mar&#160;23&#160;13:03&#160;UTC | SHAKEN Zray Technologies Corporation 862J | 26&#160;Apr&#160;23&#160;13:03&#160;UTC | true | [view](CERTS/b6eaa50ee78e158b6710e8cff3a7e923adda90f7313c9155f9254f2f05aa7f97/README.md) |
| 27&#160;Mar&#160;23&#160;17:32&#160;UTC | SHAKEN NTC International, INC 016K | 26&#160;Apr&#160;23&#160;17:32&#160;UTC | true | [view](CERTS/915747295c5d7dab12d6f9de6d10014d2c376e90003cd1dbce9371f839b68b1e/README.md) |
| 27&#160;Mar&#160;23&#160;18:33&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 26&#160;Apr&#160;23&#160;18:33&#160;UTC | true | [view](CERTS/039c99c29cd3039c1cf33ee670f536bd0a6e043004e8ccf4a079f629e97160b2/README.md) |
| 27&#160;Mar&#160;23&#160;21:48&#160;UTC | SHAKEN Global Net Holdings Inc 306K | 21&#160;Jan&#160;24&#160;21:48&#160;UTC | true | [view](CERTS/f8d913af64a11c0718c78426b747dba4ed4ccb19239db573626a46f29b671825/README.md) |
| 27&#160;Mar&#160;23&#160;23:11&#160;UTC | SHAKEN Swift Telco LLC 452K | 26&#160;Apr&#160;23&#160;23:11&#160;UTC | true | [view](CERTS/04fd890a8b928a7797150b2ebaf60e0d05b23b7a55471af338d3d4fb2ca1beec/README.md) |
| 28&#160;Mar&#160;23&#160;05:57&#160;UTC | SHAKEN BareTelecom 864J | 27&#160;Apr&#160;23&#160;05:57&#160;UTC | true | [view](CERTS/329bfb4035fa85c121c7a7958e311ab1cc947a8b3f42b376d3642ea16b2f7056/README.md) |
| 28&#160;Mar&#160;23&#160;10:41&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 27&#160;Apr&#160;23&#160;10:41&#160;UTC | true | [view](CERTS/0bd447384fc3f46f7e67b4ad426e5243de56d725eee7568783542d2564f3b4f7/README.md) |
| 28&#160;Mar&#160;23&#160;12:23&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 27&#160;Apr&#160;23&#160;12:23&#160;UTC | true | [view](CERTS/12ae40fe758649792c3c0825d4a4ad233dd0f852089b198cde9b6a396ade2d30/README.md) |
| 28&#160;Mar&#160;23&#160;15:31&#160;UTC | SHAKEN VoIP Innovations 597F | 27&#160;Mar&#160;24&#160;15:31&#160;UTC | true | [view](CERTS/105a7683d4b5fac2ea7c2383e95250715bc0460d2cfdbea0d220201f44ea5d0c/README.md) |
| 28&#160;Mar&#160;23&#160;17:08&#160;UTC | SHAKEN SouthPoint Communications 205k | 27&#160;Apr&#160;23&#160;17:08&#160;UTC | true | [view](CERTS/64b0e036e7ca432550ad3ecd293f7f725a2bd08d2d49ccea610f450127dfccd8/README.md) |
| 28&#160;Mar&#160;23&#160;17:27&#160;UTC | SHAKEN NTC International, INC 016K | 27&#160;Apr&#160;23&#160;17:27&#160;UTC | true | [view](CERTS/87ed81806e92cd33eb451320e34b1a70947a03aa78c066fd7580b877bfb3b033/README.md) |
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
| 31&#160;Mar&#160;23&#160;05:42&#160;UTC | SHAKEN BareTelecom 864J | 30&#160;Apr&#160;23&#160;05:42&#160;UTC | true | [view](CERTS/22581b2f099ee948a330c33a898ced08d6efa8f84adadbbc453c1289b754294e/README.md) |
| 31&#160;Mar&#160;23&#160;09:06&#160;UTC | SHAKEN IPSBS Managed Services LLC 828J | 30&#160;Apr&#160;23&#160;09:06&#160;UTC | true | [view](CERTS/c0d5fe63d3eb082c66852ac45f88a6d561fc7e6639bac012cf2c1d4c2829ce64/README.md) |
| 31&#160;Mar&#160;23&#160;10:26&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 30&#160;Apr&#160;23&#160;10:26&#160;UTC | true | [view](CERTS/8cbabd0259acd180bab539fdf81234170021c039f1ae570ba2ee42756d042bb1/README.md) |
| 31&#160;Mar&#160;23&#160;12:08&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 30&#160;Apr&#160;23&#160;12:08&#160;UTC | true | [view](CERTS/18740193f53a527d8aded6aeae60d001c5fb7fd1f605544942daadd1e2d77955/README.md) |
| 03&#160;Apr&#160;23&#160;14:55&#160;UTC | SHAKEN CIMA Telecom, Inc 313K | 03&#160;May&#160;23&#160;14:55&#160;UTC | true | [view](CERTS/109c2f9c177891c0a164ef13f97f0b84c80d3405e843ebdb4586be6e036dd38a/README.md) |
| 03&#160;Apr&#160;23&#160;21:02&#160;UTC | SHAKEN WWT INC dba VoIP Networks 053K | 02&#160;Apr&#160;24&#160;21:02&#160;UTC | true | [view](CERTS/6229273e97413ee4c01a1810885db9b48d7e1bf7fe8aa4e7be22076effe2cc8b/README.md) |
| 04&#160;Apr&#160;23&#160;06:00&#160;UTC | SHAKEN Convoso 758J | 09&#160;May&#160;23&#160;06:00&#160;UTC | true | [view](CERTS/b17a0addedeea480f7a605fe1516b63f4edd4dc39b84659f71868fbd86ac33e5/README.md) |
| 04&#160;Apr&#160;23&#160;16:23&#160;UTC | SHAKEN Cyberlynk Network, LLC 086K | 03&#160;Jun&#160;23&#160;16:23&#160;UTC | true | [view](CERTS/0b911cba5097116a1c6e0874d1b80f728e3877fd11c49d10ef0625d905b2dfee/README.md) |
| 04&#160;Apr&#160;23&#160;16:38&#160;UTC | SHAKEN Inventive Labs Corp 649J | 01&#160;Oct&#160;23&#160;16:38&#160;UTC | true | [view](CERTS/61bd59e51cd88d19ca695a362569930dee7725fefa39299b507bc3269fefbd54/README.md) |
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

#### CA Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 21&#160;Aug&#160;20&#160;01:22&#160;UTC | SHAKEN Sansay Root CA US | 16&#160;Aug&#160;40&#160;01:22&#160;UTC | true | [view](CERTS/8356d251b255f4ac3fd108bb79be4c02dcea2d3b13d138892bdb3a70cbe6a343/README.md) |
| 02&#160;Sep&#160;22&#160;20:53&#160;UTC | SHAKEN Sansay Intermediate CA US WEST 1 | 31&#160;Aug&#160;29&#160;20:53&#160;UTC | true | [view](CERTS/4b1dfdba2b1e4bbffbf900a20f1f6f7befbef0008b963e4922a64469cb97d24b/README.md) |


Generated: 12 Apr 23 01:46 UTC