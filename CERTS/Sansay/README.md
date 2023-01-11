# STIR/SHAKEN CA Ecosystem Compliance

## Sansay

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 284 certificates were included in the corpus being tested
- 6 certificates in the corpus were skipped because they are duplicates
- 172 certificates in the corpus were skipped because they are expired
- 1 certificates in the corpus were skipped because they are not currently trusted
- 105 certificates being tested against the remaining rules
- 5.46 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 100.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 122 days is the average remaining validity for the certificates in the corpus
- 123 days is the average initial validity for the certificates in the corpus
- 75 certificates expire in the next 30 days
- 1.67 average number of unexpired certificates per OCN observed
- 63 unique OCNs observed in unexpired and valid certificate corpus

| Instances | Test | Source |
|-----------|------|--------|
| 105 | [e_atis_certificate_policies](ISSUES/e_atis_certificate_policies/README.md) | ATIS1000080 |
| 105 | [e_atis_subject_cn](ISSUES/e_atis_subject_cn/README.md) | ATIS1000080 |
| 105 | [e_us_cp_ambiguous_identifier](ISSUES/e_us_cp_ambiguous_identifier/README.md) | US_SHAKEN_CP |
| 105 | [e_us_cp_subject_sn](ISSUES/e_us_cp_subject_sn/README.md) | US_SHAKEN_CP |
| 48 | [w_cp_1_3_subject_email](ISSUES/w_cp_1_3_subject_email/README.md) | US_SHAKEN_CP |
| 105 | [w_shaken_subject_rdn_unknown](ISSUES/w_shaken_subject_rdn_unknown/README.md) | SHAKEN_PKI_BEST_PRACTICES |

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
- 5427 days is the average remaining validity for the certificates in the corpus
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
| 11&#160;Oct&#160;22&#160;17:14&#160;UTC | SHAKEN Technology Innovation Lab 599J | 11&#160;Oct&#160;23&#160;17:14&#160;UTC | true | [view](CERTS/07d98b6eeb180548fa10e06aedbd69ce0816a1040344c91d25b8dcf29f68e7e6/README.md) |
| 11&#160;Oct&#160;22&#160;17:17&#160;UTC | SHAKEN Current Calls, LLC 746J | 11&#160;Oct&#160;23&#160;17:17&#160;UTC | true | [view](CERTS/52d6a93a1b72d2f2980699e759068dd9dbc8314c953e03613f18d9da1dcf156d/README.md) |
| 11&#160;Oct&#160;22&#160;17:19&#160;UTC | SHAKEN Inventive Labs Corp 649J | 09&#160;Apr&#160;23&#160;17:19&#160;UTC | true | [view](CERTS/a2f02cfef1eba726cf7dbd0f018a1119d40600aba568619f16b4c08b8d3a7c12/README.md) |
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
| 21&#160;Nov&#160;22&#160;21:15&#160;UTC | SHAKEN Star2Star Communications, LLC 590J | 21&#160;Nov&#160;23&#160;21:15&#160;UTC | true | [view](CERTS/9bc9dde8921387803d93036c7d2f8085af32b028fca8f17336d2e22ab51fd278/README.md) |
| 29&#160;Nov&#160;22&#160;22:04&#160;UTC | SHAKEN MagicJack 324E | 29&#160;Nov&#160;23&#160;22:04&#160;UTC | true | [view](CERTS/75b4b7b400b1252e48faa1d93f6a94f7bd4a6383c88ddf6baa167b85d9ac4ee8/README.md) |
| 05&#160;Dec&#160;22&#160;22:28&#160;UTC | SHAKEN Bulk Solutions, LLC 644J | 05&#160;Dec&#160;23&#160;22:28&#160;UTC | true | [view](CERTS/3cf0aa2a24845e3fe6b27605e223e8e0c73d6bd4f73279b8a1e5e16fd2feeb80/README.md) |
| 07&#160;Dec&#160;22&#160;17:49&#160;UTC | SHAKEN Connexum LLC 203K | 05&#160;Feb&#160;23&#160;17:49&#160;UTC | true | [view](CERTS/f51c5a0a4f577f5322a53e5b0c369a450417a971a18214adad54978aa46ffefe/README.md) |
| 10&#160;Dec&#160;22&#160;02:11&#160;UTC | SHAKEN Drop Inc 258K | 10&#160;Dec&#160;23&#160;02:11&#160;UTC | true | [view](CERTS/fc457741017b89b9126882710d8fb44883d7603f79cec0a1989eaa2b08034ee5/README.md) |
| 12&#160;Dec&#160;22&#160;22:46&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 11&#160;Jan&#160;23&#160;22:46&#160;UTC | true | [view](CERTS/050db0bdfc54e966b6d6f0f4ee116ce80edb08bbb0055a3f69718a1354a900f6/README.md) |
| 13&#160;Dec&#160;22&#160;02:17&#160;UTC | SHAKEN NTC International, INC 016K | 12&#160;Jan&#160;23&#160;02:17&#160;UTC | true | [view](CERTS/2b422e3cc0de6a2e49eb785d9d6fb925117a8a9c41d6235f1b625f86244ff227/README.md) |
| 13&#160;Dec&#160;22&#160;12:00&#160;UTC | SHAKEN 1stPoint Communications, LLC 463G | 12&#160;Jan&#160;23&#160;12:00&#160;UTC | true | [view](CERTS/ad3a026d9469d4cf7ee1db57b29e760079f8e3a118efe195588967552a99ee1f/README.md) |
| 13&#160;Dec&#160;22&#160;12:15&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 12&#160;Jan&#160;23&#160;12:15&#160;UTC | true | [view](CERTS/5778e05160c778f148c31a58676b3a86da2c9f7ecc749cbe29f61705690bdc5e/README.md) |
| 13&#160;Dec&#160;22&#160;19:38&#160;UTC | SHAKEN ConvergeTel LLC 388K | 11&#160;Jun&#160;23&#160;19:38&#160;UTC | true | [view](CERTS/4d8d9a75c4778c757d2473fa21ca43dd53ab97b308aa3988a094314df250a67d/README.md) |
| 13&#160;Dec&#160;22&#160;22:41&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 12&#160;Jan&#160;23&#160;22:41&#160;UTC | true | [view](CERTS/d6ab94aca97e6b8e273417da18c3203c50da5e29de6813a81c20799fb157ba3b/README.md) |
| 14&#160;Dec&#160;22&#160;04:19&#160;UTC | SHAKEN Ace Innovative Networks, Inc. 040K | 13&#160;Jan&#160;23&#160;04:19&#160;UTC | true | [view](CERTS/c97d843d48fc3c9ba10c80bc3ccfc4ca9807d38a00384abaaaec57c1b539af5a/README.md) |
| 14&#160;Dec&#160;22&#160;12:10&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 13&#160;Jan&#160;23&#160;12:10&#160;UTC | true | [view](CERTS/7e5f80893ef45ea2801cbb5e0002dbd6c928de1ae5faca7415b78ff8ece3e590/README.md) |
| 14&#160;Dec&#160;22&#160;15:23&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 13&#160;Jan&#160;23&#160;15:23&#160;UTC | true | [view](CERTS/7e2e47dfa7d154dceaabca2aba05a3c7e818d7a70c62a458d664ba95fade83c7/README.md) |
| 14&#160;Dec&#160;22&#160;17:59&#160;UTC | SHAKEN IPSBS Managed Services LLC 828J | 13&#160;Jan&#160;23&#160;17:59&#160;UTC | true | [view](CERTS/e942b1926bb4c94fbd45c4c218cc5aa09a909b511da8c725ae5b9c46edcf8c5c/README.md) |
| 14&#160;Dec&#160;22&#160;22:36&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 13&#160;Jan&#160;23&#160;22:36&#160;UTC | true | [view](CERTS/b89b123301269457589812eae19eafd6e455c239e03291368d4f26de068e6568/README.md) |
| 15&#160;Dec&#160;22&#160;02:07&#160;UTC | SHAKEN NTC International, INC 016K | 14&#160;Jan&#160;23&#160;02:07&#160;UTC | true | [view](CERTS/9b0309682ea39951b1a040ae0c4676fbd3b2a16602a6fe3602ba959fb898d32a/README.md) |
| 15&#160;Dec&#160;22&#160;12:05&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 14&#160;Jan&#160;23&#160;12:05&#160;UTC | true | [view](CERTS/09914e74c5ad04841f48ce52169cdc4a6e3221a7dc7b38bcecc8f7394d745df0/README.md) |
| 15&#160;Dec&#160;22&#160;22:31&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 14&#160;Jan&#160;23&#160;22:31&#160;UTC | true | [view](CERTS/2e841d05e26aa40a3c667b4732649a4c321ba92227836aedff8e9311076dcda6/README.md) |
| 16&#160;Dec&#160;22&#160;02:02&#160;UTC | SHAKEN NTC International, INC 016K | 15&#160;Jan&#160;23&#160;02:02&#160;UTC | true | [view](CERTS/32b46fcbf043a37727256b5226c54156987a902024cf2986a094f8106b7c12bc/README.md) |
| 16&#160;Dec&#160;22&#160;12:00&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 15&#160;Jan&#160;23&#160;12:00&#160;UTC | true | [view](CERTS/2cdb416d1880e6d76725062e3a5f82b29b68d6371da9db0c70505bb69d561d2f/README.md) |
| 16&#160;Dec&#160;22&#160;17:49&#160;UTC | SHAKEN IPSBS Managed Services LLC 828J | 15&#160;Jan&#160;23&#160;17:49&#160;UTC | true | [view](CERTS/7df100baafa93bf508dd1f321609ddf32bd6b06d88fca514d33bb72c55094542/README.md) |
| 16&#160;Dec&#160;22&#160;22:26&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 15&#160;Jan&#160;23&#160;22:26&#160;UTC | true | [view](CERTS/998e2735d58770424adfc182a7e7599ff5379e6091f39e457c3a4ed5bb8fa51d/README.md) |
| 17&#160;Dec&#160;22&#160;09:55&#160;UTC | SHAKEN InteractiveTel, LLC 920J | 16&#160;Jan&#160;23&#160;09:55&#160;UTC | true | [view](CERTS/b5652ed675a53137c9f7a6322b2b463da5f126d15bc39dfcf41fcafef3156fa2/README.md) |
| 17&#160;Dec&#160;22&#160;11:55&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 16&#160;Jan&#160;23&#160;11:55&#160;UTC | true | [view](CERTS/91e899f3f7b441f7ed4a1c25a256395085d13f7dbe28d6d2e871a0e0318e0acf/README.md) |
| 20&#160;Dec&#160;22&#160;22:06&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 19&#160;Jan&#160;23&#160;22:06&#160;UTC | true | [view](CERTS/92a96b5668713e24ea5a8a75c97587b9ea19ee5794a8c38fed48d600d7c4aff1/README.md) |
| 21&#160;Dec&#160;22&#160;01:37&#160;UTC | SHAKEN NTC International, INC 016K | 20&#160;Jan&#160;23&#160;01:37&#160;UTC | true | [view](CERTS/4470db7f4419e0d3f341c63642e4be7dfbdd3d417dce0b20b69b92ac6960d080/README.md) |
| 21&#160;Dec&#160;22&#160;11:35&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 20&#160;Jan&#160;23&#160;11:35&#160;UTC | true | [view](CERTS/b8362ad5a842fbbd7c10276490a0f46e08251939277093d5d971bb45b65f6d49/README.md) |
| 21&#160;Dec&#160;22&#160;14:48&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 20&#160;Jan&#160;23&#160;14:48&#160;UTC | true | [view](CERTS/e8cd22ad675fcb689610f0bde10e5ee1eacdb384d93eadc2de37f118a0b20670/README.md) |
| 21&#160;Dec&#160;22&#160;22:01&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 20&#160;Jan&#160;23&#160;22:01&#160;UTC | true | [view](CERTS/8a46f77c064c92c85a3b65c181a66afffc0a41cf6418de10700fa0a00dbe4252/README.md) |
| 22&#160;Dec&#160;22&#160;01:32&#160;UTC | SHAKEN NTC International, INC 016K | 21&#160;Jan&#160;23&#160;01:32&#160;UTC | true | [view](CERTS/029650af21e9d25231b7989f6d3a8082ef7bf72d91753beca2e6960aafc5f27f/README.md) |
| 22&#160;Dec&#160;22&#160;11:15&#160;UTC | SHAKEN 1stPoint Communications, LLC 463G | 21&#160;Jan&#160;23&#160;11:15&#160;UTC | true | [view](CERTS/8464cd0fdd4f9a87fb8742a4ed7249bb5d6f6e84daefba6f2313cdfd3dcb3e20/README.md) |
| 22&#160;Dec&#160;22&#160;11:30&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 21&#160;Jan&#160;23&#160;11:30&#160;UTC | true | [view](CERTS/3de78d12f53dd04098737680274267d52caa13f5b7bde9334ee2ff399bb9b061/README.md) |
| 22&#160;Dec&#160;22&#160;21:56&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 21&#160;Jan&#160;23&#160;21:56&#160;UTC | true | [view](CERTS/c3dee3fcbce7362f6a312d82d090a1625495e874b04527adfe83f40a0023f966/README.md) |
| 23&#160;Dec&#160;22&#160;01:27&#160;UTC | SHAKEN NTC International, INC 016K | 22&#160;Jan&#160;23&#160;01:27&#160;UTC | true | [view](CERTS/4c5305b923c35b8f2b826de7ba2c3be14f72f823abbd54a1ec6bd804ac931418/README.md) |
| 23&#160;Dec&#160;22&#160;11:25&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 22&#160;Jan&#160;23&#160;11:25&#160;UTC | true | [view](CERTS/31141784b6e952930b2a4bab58615dadfba00f2a42f75b431b635cc991d40216/README.md) |
| 23&#160;Dec&#160;22&#160;16:36&#160;UTC | SHAKEN Global Net Holdings Inc 306K | 22&#160;Jan&#160;23&#160;16:36&#160;UTC | true | [view](CERTS/a459318457982cbd8ff785e8305d3935482b059cf82b9dd0e2aa2415ac267222/README.md) |
| 23&#160;Dec&#160;22&#160;21:51&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 22&#160;Jan&#160;23&#160;21:51&#160;UTC | true | [view](CERTS/8c094cc5cc2ea457c7f46cdcf4206aa02c9cdc153f3941f290201b69967f01c1/README.md) |
| 24&#160;Dec&#160;22&#160;11:20&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 23&#160;Jan&#160;23&#160;11:20&#160;UTC | true | [view](CERTS/d10cd662f863e933a20762d943a83b124c75d2d14b1094ba7109ab953264b7c9/README.md) |
| 27&#160;Dec&#160;22&#160;21:31&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 26&#160;Jan&#160;23&#160;21:31&#160;UTC | true | [view](CERTS/e69891d24449a7e273ac487b55019d5e6ae05da92e2bd057f6a2599320cd15bb/README.md) |
| 28&#160;Dec&#160;22&#160;01:02&#160;UTC | SHAKEN NTC International, INC 016K | 27&#160;Jan&#160;23&#160;01:02&#160;UTC | true | [view](CERTS/b120e379b470aa0c4384d56a6e05a8fb72284c8538f15b4de436a7605d5d860b/README.md) |
| 28&#160;Dec&#160;22&#160;11:00&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 27&#160;Jan&#160;23&#160;11:00&#160;UTC | true | [view](CERTS/e22890a448417d551df754826ed86c56ca03d8edfee99ab2c26da88912c59949/README.md) |
| 28&#160;Dec&#160;22&#160;21:26&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 27&#160;Jan&#160;23&#160;21:26&#160;UTC | true | [view](CERTS/2a1f04ab9511f538b3184d563cc957271b113c63d5359d9a23cfc4748e86b28c/README.md) |
| 29&#160;Dec&#160;22&#160;00:57&#160;UTC | SHAKEN NTC International, INC 016K | 28&#160;Jan&#160;23&#160;00:57&#160;UTC | true | [view](CERTS/1f46c68713b38387e621eaf1a3acbe604da82d1c85e4b4716e6ad8eb3f3d8ab6/README.md) |
| 29&#160;Dec&#160;22&#160;10:55&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 28&#160;Jan&#160;23&#160;10:55&#160;UTC | true | [view](CERTS/288132c938f750585880384f6b12cc818f27d5b9d5d42833e6677a2c518f42c6/README.md) |
| 29&#160;Dec&#160;22&#160;21:21&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 28&#160;Jan&#160;23&#160;21:21&#160;UTC | true | [view](CERTS/2b49f7aea93733177ceaf58dd27e245b305bab09e64feb83a020b92297171361/README.md) |
| 30&#160;Dec&#160;22&#160;00:52&#160;UTC | SHAKEN NTC International, INC 016K | 29&#160;Jan&#160;23&#160;00:52&#160;UTC | true | [view](CERTS/e486057ab4dec50e88d5ddd8c82d024be9ddcac65209fd2f06a145e786be91da/README.md) |
| 30&#160;Dec&#160;22&#160;10:50&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 29&#160;Jan&#160;23&#160;10:50&#160;UTC | true | [view](CERTS/f492b3ab74d255ac6c44b83c7ce6de02e011fa841e8d8eae4ea165a26fc3538c/README.md) |
| 30&#160;Dec&#160;22&#160;21:16&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 29&#160;Jan&#160;23&#160;21:16&#160;UTC | true | [view](CERTS/0d2d628ffbacd9bf7f920ff396ee09b7047699c5ca3064fea3d7ae85d409005b/README.md) |
| 31&#160;Dec&#160;22&#160;08:45&#160;UTC | SHAKEN InteractiveTel, LLC 920J | 30&#160;Jan&#160;23&#160;08:45&#160;UTC | true | [view](CERTS/23c170ade1a604c621de0e066cd05a2e2ad4ee263070185d07f21a6dc1a9bf05/README.md) |
| 31&#160;Dec&#160;22&#160;10:45&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 30&#160;Jan&#160;23&#160;10:45&#160;UTC | true | [view](CERTS/44b33ac7ad301c3100614d289e8df3627084214bcaca542649011e6cb55cb418/README.md) |
| 03&#160;Jan&#160;23&#160;20:28&#160;UTC | SHAKEN Threshold Communications Inc 563J | 02&#160;Feb&#160;23&#160;20:28&#160;UTC | true | [view](CERTS/f574056c800f35d91cc740076f5d7845985c5cbc3e11e15bca6aa099ef10db92/README.md) |
| 03&#160;Jan&#160;23&#160;20:29&#160;UTC | SHAKEN Broadband Dynamics LLC 583j | 02&#160;Feb&#160;23&#160;20:29&#160;UTC | true | [view](CERTS/08c007c2b1c62cfe5fd4f74f47791b31e854c94202cecad703d30dacd543b4ac/README.md) |
| 03&#160;Jan&#160;23&#160;20:32&#160;UTC | SHAKEN Consolidated Communications 5113 | 02&#160;Feb&#160;23&#160;20:32&#160;UTC | true | [view](CERTS/5da37e774444caa72cd542ea213f296b5168be6bec11105cc5ea1f0dafe77bd1/README.md) |
| 03&#160;Jan&#160;23&#160;20:38&#160;UTC | SHAKEN Touchtone 683A | 02&#160;Feb&#160;23&#160;20:38&#160;UTC | true | [view](CERTS/dbaaf13a67598243c626db4e6ab8fa61b807f3d275a01d840e2ec8e3c4e958e5/README.md) |
| 03&#160;Jan&#160;23&#160;20:39&#160;UTC | SHAKEN Apeiron Systems 012J | 02&#160;Feb&#160;23&#160;20:39&#160;UTC | true | [view](CERTS/23a47201a04255c1c8ffe79a9d68576aa6e701551341d9dade08a20152faaf6c/README.md) |
| 03&#160;Jan&#160;23&#160;20:46&#160;UTC | SHAKEN Fonative, Inc. 684J | 02&#160;Feb&#160;23&#160;20:46&#160;UTC | true | [view](CERTS/095ed02be5493deb6f9110330eda85e4187a1c93a124b59a074057d8f2e54c8e/README.md) |
| 03&#160;Jan&#160;23&#160;20:47&#160;UTC | SHAKEN IPitomy 652J | 02&#160;Feb&#160;23&#160;20:47&#160;UTC | true | [view](CERTS/a5eea25dd9d7e31c24b3c7924d241c4a35672298a25807dc1e713789bf1ea882/README.md) |
| 03&#160;Jan&#160;23&#160;20:47&#160;UTC | SHAKEN Phone.com, Inc. 633J | 02&#160;Feb&#160;23&#160;20:47&#160;UTC | true | [view](CERTS/65ca80a96a68f400e8f42b6a7d692a443aede7ee0835688fad3afda7c7f8fa1b/README.md) |
| 03&#160;Jan&#160;23&#160;20:50&#160;UTC | SHAKEN NETRIO LLC 020K | 02&#160;Feb&#160;23&#160;20:50&#160;UTC | true | [view](CERTS/74b1b3e8e3e985c0e22639ca557949cc511d07685687a9dfec14e5514cad8cf7/README.md) |
| 03&#160;Jan&#160;23&#160;20:52&#160;UTC | SHAKEN VoIP Innovations 597F | 02&#160;Feb&#160;23&#160;20:52&#160;UTC | true | [view](CERTS/423f721439854827df8788d87721db165d76792922cd0ad9dbcf2f1153213ee9/README.md) |
| 03&#160;Jan&#160;23&#160;20:53&#160;UTC | SHAKEN IDT America, Corp 363A | 02&#160;Feb&#160;23&#160;20:53&#160;UTC | true | [view](CERTS/4ebb7e18288f15379f58ee858abedbbaa4cb42193b8bf7fba15f4ee3a3a261e3/README.md) |
| 03&#160;Jan&#160;23&#160;20:55&#160;UTC | SHAKEN Noble Systems Communications LLC 187J | 02&#160;Feb&#160;23&#160;20:55&#160;UTC | true | [view](CERTS/76fa62b93ec4749ecd817eb9826147c6d4e1abcf77bf3aeae6bf989c668a711d/README.md) |
| 03&#160;Jan&#160;23&#160;20:58&#160;UTC | SHAKEN Airespring 996H | 02&#160;Feb&#160;23&#160;20:58&#160;UTC | true | [view](CERTS/1ef92d3b3b53e467676a52a88907705f4e8446d113b22263559cc430437879ec/README.md) |
| 03&#160;Jan&#160;23&#160;20:59&#160;UTC | SHAKEN Nobelbiz, Inc. 596J | 02&#160;Feb&#160;23&#160;20:59&#160;UTC | true | [view](CERTS/289dc83e4f23e39e1521c6f128c55eaed4770e449d487dc7b33a96739e624b56/README.md) |
| 03&#160;Jan&#160;23&#160;21:02&#160;UTC | SHAKEN Momentum Telecom 1417 | 02&#160;Feb&#160;23&#160;21:02&#160;UTC | true | [view](CERTS/f98a15166b0e6d7d1daf245bb6abe9af1a851fbc5f9e05b15b33ce5fc0c28108/README.md) |
| 03&#160;Jan&#160;23&#160;21:03&#160;UTC | SHAKEN Momentum Telecom 9157 | 02&#160;Feb&#160;23&#160;21:03&#160;UTC | true | [view](CERTS/14ffa96863cf41c4c854df2b3a90477e5f289ef4ca6fef586c9507aa9fe20467/README.md) |
| 03&#160;Jan&#160;23&#160;21:05&#160;UTC | SHAKEN Terra Nova Telecom 382G | 02&#160;Feb&#160;23&#160;21:05&#160;UTC | true | [view](CERTS/50e59895e184789307de039653cf28f0e6dce9c9fda0b70c9a96db32155bcf66/README.md) |
| 03&#160;Jan&#160;23&#160;21:06&#160;UTC | SHAKEN Matrix 3058 | 02&#160;Feb&#160;23&#160;21:06&#160;UTC | true | [view](CERTS/c656d782cf4539c6964490136fefebb854df893bcff4dd7c52ea2525cace9660/README.md) |
| 03&#160;Jan&#160;23&#160;21:07&#160;UTC | SHAKEN Matrix 9451 | 02&#160;Feb&#160;23&#160;21:07&#160;UTC | true | [view](CERTS/0f3c15daa08b5cd800da9a0b24f3fd006fadb46515566f0bb32ea04020d9165c/README.md) |
| 03&#160;Jan&#160;23&#160;21:09&#160;UTC | SHAKEN Matrix 7379 | 02&#160;Feb&#160;23&#160;21:09&#160;UTC | true | [view](CERTS/3091ad0e5632ae85b7822f9584fb73fc30935652ed1cead05438ef7955b689e1/README.md) |
| 03&#160;Jan&#160;23&#160;21:11&#160;UTC | SHAKEN PNG Telecommunications Inc 3395 | 02&#160;Feb&#160;23&#160;21:11&#160;UTC | true | [view](CERTS/28c97d7ebaf9424567a67e7b039711808b6a4771b74cc0d37212912691c04bf6/README.md) |
| 03&#160;Jan&#160;23&#160;21:17&#160;UTC | SHAKEN Magna5, LLC 3849 | 02&#160;Feb&#160;23&#160;21:17&#160;UTC | true | [view](CERTS/673e340df2034424383377fe26757696ce99d2395eea74a7efa6c57e7fd4c747/README.md) |
| 03&#160;Jan&#160;23&#160;21:20&#160;UTC | SHAKEN Magna5, LLC 8249 | 02&#160;Feb&#160;23&#160;21:20&#160;UTC | true | [view](CERTS/0af17f6ce2b0933b9cdc8c8b62e78e83215746f1de36379b1e6227cacb076bf0/README.md) |
| 04&#160;Jan&#160;23&#160;07:00&#160;UTC | SHAKEN Convoso 758J | 08&#160;Feb&#160;23&#160;07:00&#160;UTC | true | [view](CERTS/7dbe075e08216745a506c6874d295474143fe07e582a32081fab641ab41277e4/README.md) |
| 04&#160;Jan&#160;23&#160;08:24&#160;UTC | SHAKEN InteractiveTel, LLC 920J | 03&#160;Feb&#160;23&#160;08:24&#160;UTC | true | [view](CERTS/40f9bf01d0c26777216d7ba43012f8137f458ebaa968ac2de5b47d32631000d9/README.md) |
| 04&#160;Jan&#160;23&#160;10:25&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 03&#160;Feb&#160;23&#160;10:25&#160;UTC | true | [view](CERTS/8a8d29e20e2e0aa8f344a52f1052834af66c2b7d8193f4ce103a3c63bf1c1052/README.md) |
| 05&#160;Jan&#160;23&#160;13:09&#160;UTC | SHAKEN Primo Dialler LLC 249K | 09&#160;Feb&#160;23&#160;13:09&#160;UTC | true | [view](CERTS/a9ac5a87a408fda204e6fae41d1665972954365e287f4a507918192bb9ab16d3/README.md) |
| 06&#160;Jan&#160;23&#160;10:15&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 05&#160;Feb&#160;23&#160;10:15&#160;UTC | true | [view](CERTS/9c9d6bedb06785838d21abf1b4adeabfbc22a5aad1a214eeeee17ea12d580d83/README.md) |
| 06&#160;Jan&#160;23&#160;16:13&#160;UTC | SHAKEN Telcentris Inc. dba Voxox 696J | 05&#160;Feb&#160;23&#160;16:13&#160;UTC | true | [view](CERTS/3686ad01ca8cabbf3924695171a78d536c6742b72304f4f30203c1c91f9be6bf/README.md) |

#### CA Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 21&#160;Aug&#160;20&#160;01:22&#160;UTC | SHAKEN Sansay Root CA US | 16&#160;Aug&#160;40&#160;01:22&#160;UTC | true | [view](CERTS/8356d251b255f4ac3fd108bb79be4c02dcea2d3b13d138892bdb3a70cbe6a343/README.md) |
| 02&#160;Sep&#160;22&#160;20:53&#160;UTC | SHAKEN Sansay Intermediate CA US WEST 1 | 31&#160;Aug&#160;29&#160;20:53&#160;UTC | true | [view](CERTS/4b1dfdba2b1e4bbffbf900a20f1f6f7befbef0008b963e4922a64469cb97d24b/README.md) |


Generated: 11 Jan 23 21:04 UTC