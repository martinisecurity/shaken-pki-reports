# STIR/SHAKEN CA Ecosystem Compliance

## Sansay

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 473 certificates were included in the corpus being tested
- 9 certificates in the corpus were skipped because they are duplicates
- 307 certificates in the corpus were skipped because they are expired
- 1 certificates in the corpus were skipped because they are not currently trusted
- 156 certificates being tested against the remaining rules
- 4.94 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 100.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 103 days is the average remaining validity for the certificates in the corpus
- 103 days is the average initial validity for the certificates in the corpus
- 120 certificates expire in the next 30 days
- 2.11 average number of unexpired certificates per OCN observed
- 74 unique OCNs observed in unexpired and valid certificate corpus

| Instances | Test | Source |
|-----------|------|--------|
| 60 | [e_atis_certificate_policies](ISSUES/e_atis_certificate_policies/README.md) | ATIS1000080 |
| 156 | [e_atis_subject_cn](ISSUES/e_atis_subject_cn/README.md) | ATIS1000080 |
| 156 | [e_us_cp_ambiguous_identifier](ISSUES/e_us_cp_ambiguous_identifier/README.md) | US_SHAKEN_CP |
| 156 | [e_us_cp_subject_sn](ISSUES/e_us_cp_subject_sn/README.md) | US_SHAKEN_CP |
| 86 | [w_cp_1_3_subject_email](ISSUES/w_cp_1_3_subject_email/README.md) | US_SHAKEN_CP |
| 156 | [w_shaken_subject_rdn_unknown](ISSUES/w_shaken_subject_rdn_unknown/README.md) | SHAKEN_PKI_BEST_PRACTICES |

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
- 5408 days is the average remaining validity for the certificates in the corpus
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
| 07&#160;Nov&#160;22&#160;21:53&#160;UTC | SHAKEN Starlinq PBX Inc. 267K | 07&#160;Nov&#160;23&#160;21:53&#160;UTC | true | [view](CERTS/556a31c75cded397a5564c6f3bce1bf50e44e68e13d8ba757b5e6c20ad997fdb/README.md) |
| 14&#160;Nov&#160;22&#160;16:52&#160;UTC | SHAKEN WWT INC dba VoIP Networks 053K | 14&#160;Nov&#160;23&#160;16:52&#160;UTC | true | [view](CERTS/308b5be0a61e548df4275a61a356c74e0d640ca3fbf6f6d87c8bc06211e745fb/README.md) |
| 21&#160;Nov&#160;22&#160;21:15&#160;UTC | SHAKEN Star2Star Communications, LLC 590J | 21&#160;Nov&#160;23&#160;21:15&#160;UTC | true | [view](CERTS/9bc9dde8921387803d93036c7d2f8085af32b028fca8f17336d2e22ab51fd278/README.md) |
| 29&#160;Nov&#160;22&#160;22:04&#160;UTC | SHAKEN MagicJack 324E | 29&#160;Nov&#160;23&#160;22:04&#160;UTC | true | [view](CERTS/75b4b7b400b1252e48faa1d93f6a94f7bd4a6383c88ddf6baa167b85d9ac4ee8/README.md) |
| 05&#160;Dec&#160;22&#160;22:28&#160;UTC | SHAKEN Bulk Solutions, LLC 644J | 05&#160;Dec&#160;23&#160;22:28&#160;UTC | true | [view](CERTS/3cf0aa2a24845e3fe6b27605e223e8e0c73d6bd4f73279b8a1e5e16fd2feeb80/README.md) |
| 10&#160;Dec&#160;22&#160;02:11&#160;UTC | SHAKEN Drop Inc 258K | 10&#160;Dec&#160;23&#160;02:11&#160;UTC | true | [view](CERTS/fc457741017b89b9126882710d8fb44883d7603f79cec0a1989eaa2b08034ee5/README.md) |
| 05&#160;Jan&#160;23&#160;22:47&#160;UTC | SHAKEN Cloud Connect LLC 455K | 21&#160;Mar&#160;23&#160;22:47&#160;UTC | true | [view](CERTS/647b845b46546bc4b5e0cc9f4f9183b358936e2a1334108ed441f2f073212cb8/README.md) |
| 05&#160;Jan&#160;23&#160;22:48&#160;UTC | SHAKEN IP Link Telecom Inc. 902J | 05&#160;Apr&#160;23&#160;22:48&#160;UTC | true | [view](CERTS/f688a135b60b18b7001827646c93befb7178c25ae6e5f9f21439ff407e8e44e9/README.md) |
| 05&#160;Jan&#160;23&#160;22:49&#160;UTC | SHAKEN IP Link Telecom Inc. 902J | 05&#160;Apr&#160;23&#160;22:49&#160;UTC | true | [view](CERTS/c6a55dde18451282141cc1ea8218e76bbc65fbbd52ed381bc7651670b16de2d6/README.md) |
| 19&#160;Jan&#160;23&#160;22:50&#160;UTC | SHAKEN Technology Innovation Lab 599J | 19&#160;Jan&#160;24&#160;22:50&#160;UTC | true | [view](CERTS/12acafcf01348d278955bb9276e7a4d22a65ccdc61a59d08100177711f21b430/README.md) |
| 23&#160;Jan&#160;23&#160;21:55&#160;UTC | SHAKEN Swift Telco LLC 452K | 23&#160;Jan&#160;24&#160;21:55&#160;UTC | true | [view](CERTS/613861829aae7927f05dbd5a7b9f28ae8c4f995bb8ed115f95fc4be6644ccde1/README.md) |
| 26&#160;Jan&#160;23&#160;14:26&#160;UTC | SHAKEN Bulk Solutions, LLC 644J | 11&#160;Jan&#160;24&#160;14:26&#160;UTC | true | [view](CERTS/7dcc6cd32bf3c4e6e2901468097a88ad42a72ce42df34edce87c84dbce3691d2/README.md) |
| 31&#160;Jan&#160;23&#160;23:39&#160;UTC | SHAKEN Connexum LLC 203K | 01&#160;May&#160;23&#160;23:39&#160;UTC | true | [view](CERTS/d5409b0a1e255cfd6a3557acd45479a31641ee57e77ac96e1865499037f56c18/README.md) |
| 06&#160;Feb&#160;23&#160;12:55&#160;UTC | SHAKEN Primo Dialler LLC 249K | 18&#160;Mar&#160;23&#160;12:55&#160;UTC | true | [view](CERTS/1f0beb9e2f7c9df37652e2b75862d6f82880a7b93ece5346c065a29c35bed2c1/README.md) |
| 08&#160;Feb&#160;23&#160;07:15&#160;UTC | SHAKEN 1stPoint Communications, LLC 463G | 10&#160;Mar&#160;23&#160;07:15&#160;UTC | true | [view](CERTS/7706ed15fe5a91158c5776abcd1bd020d4f9b44edcda5cb82f039c72204c58e9/README.md) |
| 08&#160;Feb&#160;23&#160;07:30&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 10&#160;Mar&#160;23&#160;07:30&#160;UTC | true | [view](CERTS/df4ab806dca954cd69355670ccf10811ade8c88dc9935158180039d4107db80b/README.md) |
| 08&#160;Feb&#160;23&#160;13:21&#160;UTC | SHAKEN IPSBS Managed Services LLC 828J | 10&#160;Mar&#160;23&#160;13:21&#160;UTC | true | [view](CERTS/143816e2864b8674c703dd88814a927999034408e12d61f971930d792ac30b3f/README.md) |
| 08&#160;Feb&#160;23&#160;17:56&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 10&#160;Mar&#160;23&#160;17:56&#160;UTC | true | [view](CERTS/d74eaf69af48a61a416456e029d3750fcc7acd854d11bd6df5bee9df47a4e0e3/README.md) |
| 08&#160;Feb&#160;23&#160;19:33&#160;UTC | SHAKEN ConvergeTel LLC 388K | 10&#160;Mar&#160;23&#160;19:33&#160;UTC | true | [view](CERTS/daf05ca0e5279f62e2982a0657e9541286f0b0de87f6edb444719c43602b82af/README.md) |
| 09&#160;Feb&#160;23&#160;07:25&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 11&#160;Mar&#160;23&#160;07:25&#160;UTC | true | [view](CERTS/bab0a77ee7801d516662d40f5a251a87158dde92270118a8cb955cc99a87c349/README.md) |
| 09&#160;Feb&#160;23&#160;17:51&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 11&#160;Mar&#160;23&#160;17:51&#160;UTC | true | [view](CERTS/d9c1164dd7b4f5636e1d8f634456b0d9ae5f33ccf0f20f1d333ee9d47e49dc73/README.md) |
| 09&#160;Feb&#160;23&#160;21:22&#160;UTC | SHAKEN NTC International, INC 016K | 11&#160;Mar&#160;23&#160;21:22&#160;UTC | true | [view](CERTS/34a44344eefd5581266dd71af8e29349b4e5b6c2743e4b51f7f1a1021a6fdd59/README.md) |
| 10&#160;Feb&#160;23&#160;14:47&#160;UTC | SHAKEN InteractiveTel, LLC 920J | 12&#160;Mar&#160;23&#160;14:47&#160;UTC | true | [view](CERTS/0e4c7f1769a035fafdc366a10fe61b9c58c18044d06ce13fc7a50b2129041f1f/README.md) |
| 10&#160;Feb&#160;23&#160;17:46&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 12&#160;Mar&#160;23&#160;17:46&#160;UTC | true | [view](CERTS/deb3206c227d708e165f72c32135a751853fb23c6518ca33e2be0a6c20583442/README.md) |
| 12&#160;Feb&#160;23&#160;07:10&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 14&#160;Mar&#160;23&#160;07:10&#160;UTC | true | [view](CERTS/c4378b592eb597e2d6fc279d1e37a43a9c43ec35da90ca8fc366a9038be495b3/README.md) |
| 12&#160;Feb&#160;23&#160;17:36&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 14&#160;Mar&#160;23&#160;17:36&#160;UTC | true | [view](CERTS/65c7bb739f9ac25aa166f27203778ba0c8f17a504914f8d7f87d6f6cc31da249/README.md) |
| 12&#160;Feb&#160;23&#160;19:20&#160;UTC | SHAKEN Zray Technologies Corporation 862J | 14&#160;Mar&#160;23&#160;19:20&#160;UTC | true | [view](CERTS/88d48e76f3caf436b58bf2393c32d6e7384f6b4ead9e08517eed1acc93b09225/README.md) |
| 13&#160;Feb&#160;23&#160;08:29&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 15&#160;Mar&#160;23&#160;08:29&#160;UTC | true | [view](CERTS/16ff3b5ae120d55cadf25917c29adcd7ca50bed54f4608c0272fa754b5e79be8/README.md) |
| 13&#160;Feb&#160;23&#160;17:31&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 15&#160;Mar&#160;23&#160;17:31&#160;UTC | true | [view](CERTS/462497d35ce7de0acb0d93fb99bc26e8367c8c65e369948d0f6397c5fa677853/README.md) |
| 14&#160;Feb&#160;23&#160;12:51&#160;UTC | SHAKEN IPSBS Managed Services LLC 828J | 16&#160;Mar&#160;23&#160;12:51&#160;UTC | true | [view](CERTS/5caefd0cbd11262ce3478c492199a896bf313573d47954cbbb71837edd332971/README.md) |
| 14&#160;Feb&#160;23&#160;17:12&#160;UTC | SHAKEN Ytel Inc. 703J | 14&#160;Feb&#160;24&#160;17:12&#160;UTC | true | [view](CERTS/14c9bef113cfebe60611b0c56c430518ff8d42e8b98dd7e4653bd9cf619d5641/README.md) |
| 14&#160;Feb&#160;23&#160;17:26&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 16&#160;Mar&#160;23&#160;17:26&#160;UTC | true | [view](CERTS/c605dce3519db85e8568e484cd5f431203ee87ba4f91e969aac8b1000ff7facc/README.md) |
| 15&#160;Feb&#160;23&#160;08:19&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 17&#160;Mar&#160;23&#160;08:19&#160;UTC | true | [view](CERTS/4416ff457de65c8e119a84d5bb084198f04951cd64415cb7c02cce45be00851b/README.md) |
| 15&#160;Feb&#160;23&#160;17:21&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 17&#160;Mar&#160;23&#160;17:21&#160;UTC | true | [view](CERTS/aa035e9702d43249e02a0ed06238a9a49d41bb6e414c1b41190e83e0dc8699ae/README.md) |
| 16&#160;Feb&#160;23&#160;00:30&#160;UTC | SHAKEN Every1 Telecom 486K | 18&#160;Mar&#160;23&#160;00:30&#160;UTC | true | [view](CERTS/2cdf313d28bd90cbc688184eb58c4f52993cfef4b0b12bf8629d941318f3c430/README.md) |
| 16&#160;Feb&#160;23&#160;12:41&#160;UTC | SHAKEN IPSBS Managed Services LLC 828J | 18&#160;Mar&#160;23&#160;12:41&#160;UTC | true | [view](CERTS/977229e1a828d61c81ef2fd4426a7b3cdd708e0ea0aacfc63f7ba901bb9c014d/README.md) |
| 16&#160;Feb&#160;23&#160;17:16&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 18&#160;Mar&#160;23&#160;17:16&#160;UTC | true | [view](CERTS/fecd7e981146bd86ef2b62e17e0c6c58dca58ecc9fcbb2688e2b802d5cf8a7ac/README.md) |
| 16&#160;Feb&#160;23&#160;19:00&#160;UTC | SHAKEN Zray Technologies Corporation 862J | 18&#160;Mar&#160;23&#160;19:00&#160;UTC | true | [view](CERTS/f8556f5621800a1e8a12a6db763e1d895e5563ffe89ad49267a60525b99eef43/README.md) |
| 17&#160;Feb&#160;23&#160;10:00&#160;UTC | SHAKEN ONE OWL TELECOM INC 412K | 19&#160;Mar&#160;23&#160;10:00&#160;UTC | true | [view](CERTS/b00ea70628bf8d9fa087731ae534a9429aa4fa1f84284d57024c052b90588b22/README.md) |
| 17&#160;Feb&#160;23&#160;17:11&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 19&#160;Mar&#160;23&#160;17:11&#160;UTC | true | [view](CERTS/3dbc7435ce3ff62cddb27f106a6abe4e73e33e53e4e5f3d935b851372744f1ce/README.md) |
| 17&#160;Feb&#160;23&#160;18:40&#160;UTC | SHAKEN ONE OWL TELECOM INC 412K | 19&#160;Mar&#160;23&#160;18:40&#160;UTC | true | [view](CERTS/c8f034fa1cbee169b45fa38e54d5ec669456438061fa205e3c645b8d1ffb843c/README.md) |
| 17&#160;Feb&#160;23&#160;18:55&#160;UTC | SHAKEN Zray Technologies Corporation 862J | 19&#160;Mar&#160;23&#160;18:55&#160;UTC | true | [view](CERTS/e6394011ecb58c1e191762e6be609fcb8a0fbc2b8d1ee57e67fccc461e7aa6e4/README.md) |
| 19&#160;Feb&#160;23&#160;17:01&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 21&#160;Mar&#160;23&#160;17:01&#160;UTC | true | [view](CERTS/b96db5d650211c67ab3e0a2031dc81cb5718836169b5386c355bfb2f490b12b5/README.md) |
| 19&#160;Feb&#160;23&#160;20:32&#160;UTC | SHAKEN NTC International, INC 016K | 21&#160;Mar&#160;23&#160;20:32&#160;UTC | true | [view](CERTS/1c7a155f200e776e556aa79cf3a6a40fcd27bb5e4facc0937dee599e4895f47c/README.md) |
| 20&#160;Feb&#160;23&#160;12:21&#160;UTC | SHAKEN IPSBS Managed Services LLC 828J | 22&#160;Mar&#160;23&#160;12:21&#160;UTC | true | [view](CERTS/fb4235ada88731f1b3d6fe66f13d1bcfce4bcda72ec834839d036f361e096bba/README.md) |
| 20&#160;Feb&#160;23&#160;13:41&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 22&#160;Mar&#160;23&#160;13:41&#160;UTC | true | [view](CERTS/de33663c4afb572ca2fc94eb1d9135f3ec054808b75e69a8e27817d6bdb77ca6/README.md) |
| 20&#160;Feb&#160;23&#160;16:56&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 22&#160;Mar&#160;23&#160;16:56&#160;UTC | true | [view](CERTS/3e6881b4cf560ff2a5be8edcbb99b8a98034fcc078ad4b62c11ff41285ffd04e/README.md) |
| 20&#160;Feb&#160;23&#160;18:40&#160;UTC | SHAKEN Zray Technologies Corporation 862J | 22&#160;Mar&#160;23&#160;18:40&#160;UTC | true | [view](CERTS/3c84e7f3477cd05ba9bf3e3c76f637abd8ece4867b37ee8c0281dbf1c70d6a7d/README.md) |
| 21&#160;Feb&#160;23&#160;13:36&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 23&#160;Mar&#160;23&#160;13:36&#160;UTC | true | [view](CERTS/70ab935ba2094c5b40ecf0c025269468d148583c46ebf8d2440b60d1a2f4217b/README.md) |
| 21&#160;Feb&#160;23&#160;16:51&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 23&#160;Mar&#160;23&#160;16:51&#160;UTC | true | [view](CERTS/295e3502e6b861a6b5727fff060fa0bfeb8a646067d404af1deb863d6d7b3a73/README.md) |
| 21&#160;Feb&#160;23&#160;18:35&#160;UTC | SHAKEN Zray Technologies Corporation 862J | 23&#160;Mar&#160;23&#160;18:35&#160;UTC | true | [view](CERTS/f29004047be75004b74dda2e4abb5eb4bee2fbdedb261f128ca33201a268eeed/README.md) |
| 21&#160;Feb&#160;23&#160;18:38&#160;UTC | SHAKEN i3 Broadband 5800 | 23&#160;Mar&#160;23&#160;18:38&#160;UTC | true | [view](CERTS/55efdaadada92c896d1847feeb1b31b581f2944656ebdf2494ca7d15614ad286/README.md) |
| 21&#160;Feb&#160;23&#160;20:22&#160;UTC | SHAKEN NTC International, INC 016K | 23&#160;Mar&#160;23&#160;20:22&#160;UTC | true | [view](CERTS/75301785792ada4d9fd742fa05addfdac7c9b65ee5c0cf287414a7e55e7e5541/README.md) |
| 22&#160;Feb&#160;23&#160;13:31&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 24&#160;Mar&#160;23&#160;13:31&#160;UTC | true | [view](CERTS/1d02ca6cff0a095155ff39e9e42ec44facd70f9d6a8ba05b131bb0348fa19262/README.md) |
| 22&#160;Feb&#160;23&#160;16:46&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 24&#160;Mar&#160;23&#160;16:46&#160;UTC | true | [view](CERTS/d90320851ffe9ff18813d8b2f751e1387ff8493d71e3770e641e87a74bbb59d2/README.md) |
| 22&#160;Feb&#160;23&#160;20:17&#160;UTC | SHAKEN NTC International, INC 016K | 24&#160;Mar&#160;23&#160;20:17&#160;UTC | true | [view](CERTS/5baa466c00cf8c139afceee4d1942d2f378afb4476b0b40c5110f52b40631918/README.md) |
| 22&#160;Feb&#160;23&#160;21:00&#160;UTC | SHAKEN Global Net Holdings Inc 306K | 24&#160;Mar&#160;23&#160;21:00&#160;UTC | true | [view](CERTS/8b577a4ea078997ac8ea9bec328d16b65c76470debc119312c882e14ca364566/README.md) |
| 23&#160;Feb&#160;23&#160;13:26&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 25&#160;Mar&#160;23&#160;13:26&#160;UTC | true | [view](CERTS/5eaaaad8f2c2c80546a7f712797dfcd7fb43de49f4d139934172e21aa0a5bd15/README.md) |
| 23&#160;Feb&#160;23&#160;16:41&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 25&#160;Mar&#160;23&#160;16:41&#160;UTC | true | [view](CERTS/1d9a04f66313c433d906db4fd750541beba74c3adbcc72752e00a027d61c6927/README.md) |
| 24&#160;Feb&#160;23&#160;05:55&#160;UTC | SHAKEN 1stPoint Communications, LLC 463G | 26&#160;Mar&#160;23&#160;05:55&#160;UTC | true | [view](CERTS/6662ecaa092c9f5d701c2154eef3a623a10054c139ec97def76326519714695c/README.md) |
| 24&#160;Feb&#160;23&#160;13:21&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 26&#160;Mar&#160;23&#160;13:21&#160;UTC | true | [view](CERTS/9918bec39e45c4b9d9c50fb3985d18b284f5bf93d9f9ec560810434e8645e2a5/README.md) |
| 24&#160;Feb&#160;23&#160;13:37&#160;UTC | SHAKEN InteractiveTel, LLC 920J | 26&#160;Mar&#160;23&#160;13:37&#160;UTC | true | [view](CERTS/4318b92589f52aad7ca720d88da246e4bee3d20ed7c425db700e313b44df6c0f/README.md) |
| 24&#160;Feb&#160;23&#160;16:36&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 26&#160;Mar&#160;23&#160;16:36&#160;UTC | true | [view](CERTS/aeda390f13a3da010f95bebc876f7eae7ef220cd2015dff7d53a89d5eb1f0c2c/README.md) |
| 24&#160;Feb&#160;23&#160;20:07&#160;UTC | SHAKEN NTC International, INC 016K | 26&#160;Mar&#160;23&#160;20:07&#160;UTC | true | [view](CERTS/265c1b8550cc771d7238bebf4ef72a60f69799bcbae20b79732fa4f266df43e9/README.md) |
| 26&#160;Feb&#160;23&#160;13:27&#160;UTC | SHAKEN InteractiveTel, LLC 920J | 28&#160;Mar&#160;23&#160;13:27&#160;UTC | true | [view](CERTS/d718e1af3b99269c3f59acb96cf721ae40262fc575accd68ed50e1a5fdce322a/README.md) |
| 26&#160;Feb&#160;23&#160;16:26&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 28&#160;Mar&#160;23&#160;16:26&#160;UTC | true | [view](CERTS/3dfcba32b73a2b7dfd6162cd628b2eeb27a4b13d663b2a1dcaad5b0ac3331660/README.md) |
| 26&#160;Feb&#160;23&#160;19:57&#160;UTC | SHAKEN NTC International, INC 016K | 28&#160;Mar&#160;23&#160;19:57&#160;UTC | true | [view](CERTS/b2325581d61ec6ca86cfc95a355a1589c4594a0c8db5d02b9123bbb8f75f212f/README.md) |
| 27&#160;Feb&#160;23&#160;05:40&#160;UTC | SHAKEN 1stPoint Communications, LLC 463G | 29&#160;Mar&#160;23&#160;05:40&#160;UTC | true | [view](CERTS/ff062080a9b5a1994069e7c91fa82b8e5e2368951486b9b9b5465389d72fa803/README.md) |
| 27&#160;Feb&#160;23&#160;13:06&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 29&#160;Mar&#160;23&#160;13:06&#160;UTC | true | [view](CERTS/7d07e0ea819c6389d10bbf313478b836696983c0c546ffc60659d210442ffe32/README.md) |
| 27&#160;Feb&#160;23&#160;13:22&#160;UTC | SHAKEN InteractiveTel, LLC 920J | 29&#160;Mar&#160;23&#160;13:22&#160;UTC | true | [view](CERTS/890cac6351906bfbebbe456e7b740f705d279ac15676f4393e90b780a002a7e5/README.md) |
| 27&#160;Feb&#160;23&#160;16:21&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 29&#160;Mar&#160;23&#160;16:21&#160;UTC | true | [view](CERTS/d6a8064400d111cf03fece69b1113815e259c2eab584e349914f532d21ab6f11/README.md) |
| 27&#160;Feb&#160;23&#160;19:52&#160;UTC | SHAKEN NTC International, INC 016K | 29&#160;Mar&#160;23&#160;19:52&#160;UTC | true | [view](CERTS/cd85ff39646f3946e5ca4a2ba15dbad4fe719bf11ccc03291512c88e47e38d14/README.md) |
| 27&#160;Feb&#160;23&#160;21:04&#160;UTC | SHAKEN Swift Telco LLC 452K | 29&#160;Mar&#160;23&#160;21:04&#160;UTC | true | [view](CERTS/623a611e57e834178467d3af273715bbbd3d4a0d587ce56fa7a2419915117367/README.md) |
| 28&#160;Feb&#160;23&#160;11:41&#160;UTC | SHAKEN IPSBS Managed Services LLC 828J | 30&#160;Mar&#160;23&#160;11:41&#160;UTC | true | [view](CERTS/b675765e377d64b039b8fbd981c90ed55109bc3726e2c27aa9d55c4d5c48b3d5/README.md) |
| 28&#160;Feb&#160;23&#160;13:01&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 30&#160;Mar&#160;23&#160;13:01&#160;UTC | true | [view](CERTS/cf44eaa687564a7050414aa43ee88bbbee4bb9f870cdfb039ff3fcc7c8c81f6f/README.md) |
| 28&#160;Feb&#160;23&#160;13:17&#160;UTC | SHAKEN InteractiveTel, LLC 920J | 30&#160;Mar&#160;23&#160;13:17&#160;UTC | true | [view](CERTS/c4520f8d9987c1b46deab924db0974c9ba391bff00e61e95fa6fffe4554d9520/README.md) |
| 28&#160;Feb&#160;23&#160;16:16&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 30&#160;Mar&#160;23&#160;16:16&#160;UTC | true | [view](CERTS/c376b922b265815d1a97d5c205b199ed94709fe9d25120245740b80bf75d6a33/README.md) |
| 28&#160;Feb&#160;23&#160;17:33&#160;UTC | SHAKEN Threshold Communications Inc 563J | 30&#160;Mar&#160;23&#160;17:33&#160;UTC | true | [view](CERTS/ac4c1aae3df1f007f3511f0455b8d4b8b80243eef4b99ab39795066e3f69dd2e/README.md) |
| 28&#160;Feb&#160;23&#160;17:35&#160;UTC | SHAKEN Broadband Dynamics LLC 583j | 30&#160;Mar&#160;23&#160;17:35&#160;UTC | true | [view](CERTS/f0b913efb5401789aa874335fab5fe296cef3ec5bcc231e95639212a0bae4ce0/README.md) |
| 28&#160;Feb&#160;23&#160;17:36&#160;UTC | SHAKEN Consolidated Communications 5113 | 30&#160;Mar&#160;23&#160;17:36&#160;UTC | true | [view](CERTS/65a9699220ceb4b47d4b358225daaf25f6b94c1f58b3f0b68a6e6b1fd39ec4c2/README.md) |
| 28&#160;Feb&#160;23&#160;17:37&#160;UTC | SHAKEN Touchtone 683A | 30&#160;Mar&#160;23&#160;17:37&#160;UTC | true | [view](CERTS/64eaf7fcfac1496e50f587f4aaf6ab958cd44c1c3359d8425e27f03222a0d1d4/README.md) |
| 28&#160;Feb&#160;23&#160;18:00&#160;UTC | SHAKEN Zray Technologies Corporation 862J | 30&#160;Mar&#160;23&#160;18:00&#160;UTC | true | [view](CERTS/2922a982b66ef73435e09f019ac454f1236b84af3d129d9f6bbd37bc4953fd85/README.md) |
| 28&#160;Feb&#160;23&#160;19:30&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 30&#160;Mar&#160;23&#160;19:30&#160;UTC | true | [view](CERTS/a10d087befe540af9e9ab9970f560f29268ec5fac306f745e9d57b0c723ad6de/README.md) |
| 28&#160;Feb&#160;23&#160;19:47&#160;UTC | SHAKEN NTC International, INC 016K | 30&#160;Mar&#160;23&#160;19:47&#160;UTC | true | [view](CERTS/57c73e3963da1833a07e5d919ef4d2f91675a9d34bceae7d9494c2cce24c96a5/README.md) |
| 01&#160;Mar&#160;23&#160;11:36&#160;UTC | SHAKEN IPSBS Managed Services LLC 828J | 31&#160;Mar&#160;23&#160;11:36&#160;UTC | true | [view](CERTS/1990fbbd3df1f936848d969febe709e09e96a5ef58ff924a9d0cb2d761e320a9/README.md) |
| 01&#160;Mar&#160;23&#160;12:56&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 31&#160;Mar&#160;23&#160;12:56&#160;UTC | true | [view](CERTS/bf659bd0f467e3797c323d990976a5a6b8940999f3d29598c34569251406cad1/README.md) |
| 01&#160;Mar&#160;23&#160;14:49&#160;UTC | SHAKEN Apeiron Systems 012J | 31&#160;Mar&#160;23&#160;14:49&#160;UTC | true | [view](CERTS/38bbe6fe024a7af7e4230b53294202322a59c3cda969f9683d581de07c582e37/README.md) |
| 01&#160;Mar&#160;23&#160;14:50&#160;UTC | SHAKEN Fonative, Inc. 684J | 31&#160;Mar&#160;23&#160;14:50&#160;UTC | true | [view](CERTS/edcbfaf6e446c740d37bcf3298e609ce17a53fb37cf097c747fc1a94a0fe667e/README.md) |
| 01&#160;Mar&#160;23&#160;14:51&#160;UTC | SHAKEN IPitomy 652J | 31&#160;Mar&#160;23&#160;14:51&#160;UTC | true | [view](CERTS/7cb98727274e05bfb5648f18b6f5f7f94765930b8a2eabed9f4385ed98392966/README.md) |
| 01&#160;Mar&#160;23&#160;14:51&#160;UTC | SHAKEN Phone.com, Inc. 633J | 31&#160;Mar&#160;23&#160;14:51&#160;UTC | true | [view](CERTS/6307a546be4e0e440dee6d2b404374f7c32e8790e2acc3792baddd71cba40a39/README.md) |
| 01&#160;Mar&#160;23&#160;14:51&#160;UTC | SHAKEN NETRIO LLC 020K | 31&#160;Mar&#160;23&#160;14:51&#160;UTC | true | [view](CERTS/62a05350bf9fca4a98937f14ac5184c64b094664b7eb79733b60654acad58b6a/README.md) |
| 01&#160;Mar&#160;23&#160;14:52&#160;UTC | SHAKEN VoIP Innovations 597F | 31&#160;Mar&#160;23&#160;14:52&#160;UTC | true | [view](CERTS/cc89e2e33b0551fd4a6848a610d70e1320c35de95f27b998cafbc3b3eda0ac28/README.md) |
| 01&#160;Mar&#160;23&#160;14:52&#160;UTC | SHAKEN IDT America, Corp 363A | 31&#160;Mar&#160;23&#160;14:52&#160;UTC | true | [view](CERTS/448eda358dbd5f1b66bd88f0725376c4cca8cc471fcefe4341df4f97966f4cd7/README.md) |
| 01&#160;Mar&#160;23&#160;14:53&#160;UTC | SHAKEN Noble Systems Communications LLC 187J | 31&#160;Mar&#160;23&#160;14:53&#160;UTC | true | [view](CERTS/433e8721ca5f8e5e5d3782cc19e0a4857aed626ff6fe775dd841be0cd1c33965/README.md) |
| 01&#160;Mar&#160;23&#160;14:53&#160;UTC | SHAKEN Telcentris Inc. dba Voxox 696J | 31&#160;Mar&#160;23&#160;14:53&#160;UTC | true | [view](CERTS/b0aa4ed1b04002696e90a71bd8717fda888a5b71f2b1f96fae4dbde4876f2812/README.md) |
| 01&#160;Mar&#160;23&#160;14:54&#160;UTC | SHAKEN Airespring 996H | 31&#160;Mar&#160;23&#160;14:54&#160;UTC | true | [view](CERTS/4de92d1c5501824fd403c561e266ee53fc64285137d40f5e3c46d15389f29fad/README.md) |
| 01&#160;Mar&#160;23&#160;14:54&#160;UTC | SHAKEN Nobelbiz, Inc. 596J | 31&#160;Mar&#160;23&#160;14:54&#160;UTC | true | [view](CERTS/d15c87c284f34b311f0b3cf77aa329496dc2f7c118032f58d411817a2bff5c42/README.md) |
| 01&#160;Mar&#160;23&#160;14:55&#160;UTC | SHAKEN Momentum Telecom 1417 | 31&#160;Mar&#160;23&#160;14:55&#160;UTC | true | [view](CERTS/995f29ca98566a49fd525665cb29085c4987e120df2cef1bdae23187c48f8a98/README.md) |
| 01&#160;Mar&#160;23&#160;14:56&#160;UTC | SHAKEN Momentum Telecom 9157 | 31&#160;Mar&#160;23&#160;14:56&#160;UTC | true | [view](CERTS/d6fbe84709eb5b6cd0b66b55711e412318f4c970d38bf2c6043d7f4eecbfa092/README.md) |
| 01&#160;Mar&#160;23&#160;14:58&#160;UTC | SHAKEN Terra Nova Telecom 382G | 31&#160;Mar&#160;23&#160;14:58&#160;UTC | true | [view](CERTS/4a4188f3a8809dcd26241a9434efbbfc5ed40b7abb59bcc4521a804e29becfcf/README.md) |
| 01&#160;Mar&#160;23&#160;14:58&#160;UTC | SHAKEN Matrix 3058 | 31&#160;Mar&#160;23&#160;14:58&#160;UTC | true | [view](CERTS/2c9148ee0aeb9c010bd86dd6075b4ef1a9ab5e16d364388f5f1f89ed31f9d531/README.md) |
| 01&#160;Mar&#160;23&#160;14:59&#160;UTC | SHAKEN Matrix 9451 | 31&#160;Mar&#160;23&#160;14:59&#160;UTC | true | [view](CERTS/fb61cd0a81eab5b962741267cc99ec52b9841b8929f4b9e0af67e638a096bceb/README.md) |
| 01&#160;Mar&#160;23&#160;15:00&#160;UTC | SHAKEN Matrix 7379 | 31&#160;Mar&#160;23&#160;15:00&#160;UTC | true | [view](CERTS/e7797df71c91bd23a6c7195052a63096bdceb801d6677e22daa86b7a4490b13d/README.md) |
| 01&#160;Mar&#160;23&#160;15:01&#160;UTC | SHAKEN PNG Telecommunications Inc 3395 | 31&#160;Mar&#160;23&#160;15:01&#160;UTC | true | [view](CERTS/4372471dd39005f378fffb78d9d322861da6d96e568bb998539f3c3f2acaafa8/README.md) |
| 01&#160;Mar&#160;23&#160;15:06&#160;UTC | SHAKEN Magna5, LLC 3849 | 31&#160;Mar&#160;23&#160;15:06&#160;UTC | true | [view](CERTS/176e8425e224ade0f779a0a1d902ce08b90f36b021dad01244b3fadb8ad9d486/README.md) |
| 01&#160;Mar&#160;23&#160;15:16&#160;UTC | SHAKEN Magna5, LLC 8249 | 31&#160;Mar&#160;23&#160;15:16&#160;UTC | true | [view](CERTS/220829d53abc0d14e678c2dcc987a61f4cc277a95599bd48dbb9d11c630500b9/README.md) |
| 01&#160;Mar&#160;23&#160;19:25&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 31&#160;Mar&#160;23&#160;19:25&#160;UTC | true | [view](CERTS/3c6e4a44cf6ad839fe3c052af148fdaca476a29b3e6bee565a11f70dc00b849f/README.md) |
| 02&#160;Mar&#160;23&#160;05:25&#160;UTC | SHAKEN 1stPoint Communications, LLC 463G | 01&#160;Apr&#160;23&#160;05:25&#160;UTC | true | [view](CERTS/7b7b2fbb4780540bd4d81fff93e4c77211f817922a2c011fd8ecf35c62e60247/README.md) |
| 02&#160;Mar&#160;23&#160;12:51&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 01&#160;Apr&#160;23&#160;12:51&#160;UTC | true | [view](CERTS/d219a9f6e36825d36ead34d5de97464b22a312b8a23e22d0650cee55b2964b17/README.md) |
| 02&#160;Mar&#160;23&#160;17:41&#160;UTC | SHAKEN BareTelecom 864J | 01&#160;Apr&#160;23&#160;17:41&#160;UTC | true | [view](CERTS/9fe2d62bee9f5583ec3f50e1dcbe63444fd6f0e59d3c92ab9d6c8d20256d5223/README.md) |
| 02&#160;Mar&#160;23&#160;19:20&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 01&#160;Apr&#160;23&#160;19:20&#160;UTC | true | [view](CERTS/ef081e2c0898e8533ad9e4d26f7791831e5a45c87aca4d810f24254b668655b6/README.md) |
| 02&#160;Mar&#160;23&#160;19:37&#160;UTC | SHAKEN NTC International, INC 016K | 01&#160;Apr&#160;23&#160;19:37&#160;UTC | true | [view](CERTS/9db3489538d7dba2d3f79236834e639ba0a752e87f35baf29fc68e5a542424fe/README.md) |
| 03&#160;Mar&#160;23&#160;11:26&#160;UTC | SHAKEN IPSBS Managed Services LLC 828J | 02&#160;Apr&#160;23&#160;11:26&#160;UTC | true | [view](CERTS/3853da6073f231ba0c2dfba15d98c0dac0d73a54b027ef8a583116aafcde8217/README.md) |
| 03&#160;Mar&#160;23&#160;19:15&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 02&#160;Apr&#160;23&#160;19:15&#160;UTC | true | [view](CERTS/39e8a6cdb19ce511efb49a03076db4dcd6c612016dbc3adec25227b44faa22cc/README.md) |
| 05&#160;Mar&#160;23&#160;07:00&#160;UTC | SHAKEN Convoso 758J | 09&#160;Apr&#160;23&#160;07:00&#160;UTC | true | [view](CERTS/97e2c0b86e830df512ab950a4a611216b177512b92c2821e535551f65a88efba/README.md) |
| 05&#160;Mar&#160;23&#160;19:05&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 04&#160;Apr&#160;23&#160;19:05&#160;UTC | true | [view](CERTS/532ed49602fc93aa93bfb6cc89ccf63efc9ba52682c0edc716b09fb7df20957f/README.md) |
| 05&#160;Mar&#160;23&#160;19:22&#160;UTC | SHAKEN NTC International, INC 016K | 04&#160;Apr&#160;23&#160;19:22&#160;UTC | true | [view](CERTS/ab90b4262682d09a31970ddcc2ac195641cf40bc0e3693c782050bc0e747b316/README.md) |
| 06&#160;Mar&#160;23&#160;06:30&#160;UTC | SHAKEN  XCast Labs 689J | 06&#160;Apr&#160;23&#160;06:30&#160;UTC | true | [view](CERTS/e459bc12f490af4f1eec5506a85402c13e046ce415f5fc1b9caf8105d7f55f9f/README.md) |
| 06&#160;Mar&#160;23&#160;12:31&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 05&#160;Apr&#160;23&#160;12:31&#160;UTC | true | [view](CERTS/a2c14721bdb858e491cc6b9cf332f0119511ea0a4878ba26194bcf1edfa2e479/README.md) |
| 06&#160;Mar&#160;23&#160;19:00&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 05&#160;Apr&#160;23&#160;19:00&#160;UTC | true | [view](CERTS/13cddbf30d199104b5460bf4a7854b224eafc5e6f4da2afa085267f547be48c1/README.md) |
| 07&#160;Mar&#160;23&#160;11:06&#160;UTC | SHAKEN IPSBS Managed Services LLC 828J | 06&#160;Apr&#160;23&#160;11:06&#160;UTC | true | [view](CERTS/7d2de8f4c9d251fc1c10ea005753f712fe54673d63b6d8d0c462651035317a11/README.md) |
| 07&#160;Mar&#160;23&#160;12:26&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 06&#160;Apr&#160;23&#160;12:26&#160;UTC | true | [view](CERTS/9b5b472e95fa894ece49739bc61d6448a67d5fbf0fc8d40a94cc1e1120e5aa4a/README.md) |
| 07&#160;Mar&#160;23&#160;12:42&#160;UTC | SHAKEN InteractiveTel, LLC 920J | 06&#160;Apr&#160;23&#160;12:42&#160;UTC | true | [view](CERTS/6e8bcd32ddda5819af36793b81205b691694e625192deb83b11ffa57ba0b58a5/README.md) |
| 07&#160;Mar&#160;23&#160;18:55&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 06&#160;Apr&#160;23&#160;18:55&#160;UTC | true | [view](CERTS/bb0fb530d7878d2a8464d3bdb90614da614763010d555af4d7772fcb513ccfb1/README.md) |
| 08&#160;Mar&#160;23&#160;12:21&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 07&#160;Apr&#160;23&#160;12:21&#160;UTC | true | [view](CERTS/5e140f39ae5bf09b0baa2f245d57352f1e0f68e25c86e5d599a588451b9bbbcd/README.md) |
| 08&#160;Mar&#160;23&#160;12:37&#160;UTC | SHAKEN InteractiveTel, LLC 920J | 07&#160;Apr&#160;23&#160;12:37&#160;UTC | true | [view](CERTS/acb8851d873aa0c37be62eb59e0a858a5fccc54e1238d6a1c2ef423d3b29f24b/README.md) |
| 08&#160;Mar&#160;23&#160;13:38&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 07&#160;Apr&#160;23&#160;13:38&#160;UTC | true | [view](CERTS/79d95a52f4b87f15fd30907d7b7314b1c42af03ea4886801f4b514a3f12f1226/README.md) |
| 08&#160;Mar&#160;23&#160;19:07&#160;UTC | SHAKEN NTC International, INC 016K | 07&#160;Apr&#160;23&#160;19:07&#160;UTC | true | [view](CERTS/43c593b1b338d16b592e7c014ab5994fca88c16e8d4b74184f4753c7aff1d674/README.md) |
| 08&#160;Mar&#160;23&#160;21:17&#160;UTC | SHAKEN Ace Innovative Networks, Inc. 040K | 07&#160;Apr&#160;23&#160;21:17&#160;UTC | true | [view](CERTS/b54728a5b3dfe9e25d5ac4ce9e66589dded588f205c726699a6313949272abee/README.md) |
| 09&#160;Mar&#160;23&#160;12:16&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 08&#160;Apr&#160;23&#160;12:16&#160;UTC | true | [view](CERTS/eb62dccbccd5e11f1d4d461c787e1eec7cb6c079c85308b836433511ae62838a/README.md) |
| 09&#160;Mar&#160;23&#160;13:33&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 08&#160;Apr&#160;23&#160;13:33&#160;UTC | true | [view](CERTS/0b95d12f4c3d0b0aedce8f38c45e506fe5fc07cbcb62d2e293d2ff644cd27de8/README.md) |
| 09&#160;Mar&#160;23&#160;15:43&#160;UTC | SHAKEN Zray Technologies Corporation 862J | 08&#160;Apr&#160;23&#160;15:43&#160;UTC | true | [view](CERTS/bfbe0be12cdf50afe76fadb155cb2be3d196794ef23d5f3a875f58495fa8f545/README.md) |

#### CA Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 21&#160;Aug&#160;20&#160;01:22&#160;UTC | SHAKEN Sansay Root CA US | 16&#160;Aug&#160;40&#160;01:22&#160;UTC | true | [view](CERTS/8356d251b255f4ac3fd108bb79be4c02dcea2d3b13d138892bdb3a70cbe6a343/README.md) |
| 02&#160;Sep&#160;22&#160;20:53&#160;UTC | SHAKEN Sansay Intermediate CA US WEST 1 | 31&#160;Aug&#160;29&#160;20:53&#160;UTC | true | [view](CERTS/4b1dfdba2b1e4bbffbf900a20f1f6f7befbef0008b963e4922a64469cb97d24b/README.md) |


Generated: 10 Mar 23 02:25 UTC