# STIR/SHAKEN CA Ecosystem Compliance

## Sansay

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 190 certificates were included in the corpus being tested
- 4 certificates in the corpus were skipped because they are duplicates
- 72 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 114 certificates being tested against the remaining rules
- 5.51 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 100.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 102 days is the average remaining validity for the certificates in the corpus
- 103 days is the average initial validity for the certificates in the corpus
- 88 certificates expire in the next 30 days
- 1.81 average number of unexpired certificates per OCN observed
- 63 unique OCNs observed in unexpired and valid certificate corpus

| Instances | Test | Source |
|-----------|------|--------|
| 114 | [e_atis_certificate_policies](ISSUES/e_atis_certificate_policies/README.md) | ATIS1000080 |
| 114 | [e_atis_subject_cn](ISSUES/e_atis_subject_cn/README.md) | ATIS1000080 |
| 114 | [e_us_cp_ambiguous_identifier](ISSUES/e_us_cp_ambiguous_identifier/README.md) | US_SHAKEN_CP |
| 114 | [e_us_cp_subject_sn](ISSUES/e_us_cp_subject_sn/README.md) | US_SHAKEN_CP |
| 58 | [w_cp_1_3_subject_email](ISSUES/w_cp_1_3_subject_email/README.md) | US_SHAKEN_CP |
| 114 | [w_shaken_subject_rdn_unknown](ISSUES/w_shaken_subject_rdn_unknown/README.md) | SHAKEN_PKI_BEST_PRACTICES |

#### CA Certificates

- 3 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 1 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 2 certificates being tested against the remaining rules
- 1.00 issues on average found in unexpired, trusted, and non-compliant certificates
- 0.00% of certificates contain one or more Error level issue
- 100.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 50.00% of certificates are too old to be assessed against currently enforced expectations
- 5441 days is the average remaining validity for the certificates in the corpus
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
| 01&#160;Oct&#160;22&#160;00:00&#160;UTC | SHAKEN IP Link Telecom Inc. 902J | 30&#160;Dec&#160;22&#160;00:00&#160;UTC | true | [view](CERTS/53d491627ef52772ffd75835f745e895f025bfc853523028b42c0bc1681e5f67/README.md) |
| 05&#160;Oct&#160;22&#160;19:08&#160;UTC | SHAKEN IP Link Telecom Inc. 902J | 03&#160;Jan&#160;23&#160;19:08&#160;UTC | true | [view](CERTS/fa453a929f14e705532a1974b216dedb71ba1f10d6af0df07f1084bfbb2038db/README.md) |
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
| 24&#160;Oct&#160;22&#160;20:23&#160;UTC | SHAKEN Arbeit 816J | 24&#160;Oct&#160;23&#160;20:23&#160;UTC | true | [view](CERTS/377e182a223e6cc8d7e9ce697e7a3e829b1c6b16c299c26f6d1f1e33aa29524b/README.md) |
| 24&#160;Oct&#160;22&#160;21:11&#160;UTC | SHAKEN Ytel Inc. 703J | 24&#160;Oct&#160;23&#160;21:11&#160;UTC | true | [view](CERTS/3d6a7a2ff23b90fba1674f600a108b8a11a110f8bb1723df86627001f7367d8d/README.md) |
| 25&#160;Oct&#160;22&#160;20:17&#160;UTC | SHAKEN Talk IT Pro 321K | 25&#160;Oct&#160;23&#160;20:17&#160;UTC | true | [view](CERTS/e45dada701a589e681d12207ebf16985abf6d62cf429b6e03bdcf8c0f97c3bf2/README.md) |
| 26&#160;Oct&#160;22&#160;19:34&#160;UTC | SHAKEN Vinculum Communications, Inc 787J | 26&#160;Oct&#160;23&#160;19:34&#160;UTC | true | [view](CERTS/22936e87ea3c45af88f1e501b88c6c6db3c271bd6ef73ab33c5d68198f9d4d66/README.md) |
| 26&#160;Oct&#160;22&#160;19:41&#160;UTC | SHAKEN DLS Internet Services 815J | 26&#160;Oct&#160;23&#160;19:41&#160;UTC | true | [view](CERTS/7cd8319bedd12f040e8bd7b522d981aabcd24dc5aef74614a67fb6fdc9b9823b/README.md) |
| 26&#160;Oct&#160;22&#160;19:43&#160;UTC | SHAKEN Systemverse, LLC. 194K | 26&#160;Oct&#160;23&#160;19:43&#160;UTC | true | [view](CERTS/edbe74f809b9e0e1ebea447df8bdbfb272144f9c8c18df81e397a374df61c4cd/README.md) |
| 27&#160;Oct&#160;22&#160;20:11&#160;UTC | SHAKEN Mitel Cloud Services, Inc. 670J | 27&#160;Oct&#160;23&#160;20:11&#160;UTC | true | [view](CERTS/e45c92abcfe2fe6d0863200900b66e835aa98712f974efe3837e34d787f2ad5e/README.md) |
| 02&#160;Nov&#160;22&#160;16:14&#160;UTC | SHAKEN InteractiveTel, LLC 920J | 02&#160;Dec&#160;22&#160;16:14&#160;UTC | true | [view](CERTS/0ced8150d7cbed20775669aafeecd05726f233c901d024bb4fa64e9ce75b9f7a/README.md) |
| 03&#160;Nov&#160;22&#160;15:35&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 03&#160;Dec&#160;22&#160;15:35&#160;UTC | true | [view](CERTS/e05e81f7e34f2088102df66b64b44c5549708f9182a79e6e7d1157333260d98b/README.md) |
| 04&#160;Nov&#160;22&#160;15:30&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 04&#160;Dec&#160;22&#160;15:30&#160;UTC | true | [view](CERTS/51862e962c0ec2667b8ee2862f0eee4d50fcdaf629bfaa3e7191ed7018f56db1/README.md) |
| 05&#160;Nov&#160;22&#160;06:00&#160;UTC | SHAKEN Convoso 758J | 10&#160;Dec&#160;22&#160;06:00&#160;UTC | true | [view](CERTS/a33008ec0c4ad56c980d02e36d07e33a42f8c2f203fea78024c5018d5cf04e8d/README.md) |
| 07&#160;Nov&#160;22&#160;00:42&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 07&#160;Dec&#160;22&#160;00:42&#160;UTC | true | [view](CERTS/d27a959a21dc295781c932032c24d6e53b61a5088ab3d2cc98ac93562c5ec6cd/README.md) |
| 07&#160;Nov&#160;22&#160;15:15&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 07&#160;Dec&#160;22&#160;15:15&#160;UTC | true | [view](CERTS/16593a1d36bc2265a80aa16a52e6936c5004486f6a84266e6b9d166c82509a8c/README.md) |
| 07&#160;Nov&#160;22&#160;18:27&#160;UTC | SHAKEN IPSBS Managed Services LLC 828J | 07&#160;Dec&#160;22&#160;18:27&#160;UTC | true | [view](CERTS/5c216d5eb90d5b09e51db9c9a1b3bad35f94f504a92261dc65dde90913fb410a/README.md) |
| 08&#160;Nov&#160;22&#160;14:55&#160;UTC | SHAKEN 1stPoint Communications, LLC 463G | 08&#160;Dec&#160;22&#160;14:55&#160;UTC | true | [view](CERTS/322096d99c6f6beb803c72f64ba15430bf34292da2e7736382405346d1920c2e/README.md) |
| 08&#160;Nov&#160;22&#160;15:10&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 08&#160;Dec&#160;22&#160;15:10&#160;UTC | true | [view](CERTS/0bab2fef8838a943dc8b17e3632776a4f4d8e74508ca22efbeb0dd282adde5d2/README.md) |
| 08&#160;Nov&#160;22&#160;18:18&#160;UTC | SHAKEN Primo Dialler LLC 249K | 08&#160;Dec&#160;22&#160;18:18&#160;UTC | true | [view](CERTS/1865568906d0385d9cdfa133305a3aeea8bb25f8630ffc630c654a8d47273168/README.md) |
| 09&#160;Nov&#160;22&#160;02:14&#160;UTC | SHAKEN Connexum LLC 203K | 09&#160;Dec&#160;22&#160;02:14&#160;UTC | true | [view](CERTS/b6f81e35a57af695e4d70f61d18bac7cbd8f823b687f6adbb91561da8ecaab81/README.md) |
| 09&#160;Nov&#160;22&#160;12:51&#160;UTC | SHAKEN InteractiveTel, LLC 920J | 09&#160;Dec&#160;22&#160;12:51&#160;UTC | true | [view](CERTS/b032d7edd0943eabb4b18af4e272cc93270997d66ce6e5051193454be2dd7369/README.md) |
| 09&#160;Nov&#160;22&#160;14:50&#160;UTC | SHAKEN 1stPoint Communications, LLC 463G | 09&#160;Dec&#160;22&#160;14:50&#160;UTC | true | [view](CERTS/b4e31ab7b8fc550e75a7b00886b00d69db00cd56a13c4dd2cf8f1911749a4dcb/README.md) |
| 09&#160;Nov&#160;22&#160;15:05&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 09&#160;Dec&#160;22&#160;15:05&#160;UTC | true | [view](CERTS/31721e97e825e9f221be10049da5735c86acdafc29714b581cf12e22acc72bb6/README.md) |
| 09&#160;Nov&#160;22&#160;18:49&#160;UTC | SHAKEN Primo Dialler LLC 249K | 09&#160;Dec&#160;22&#160;18:49&#160;UTC | true | [view](CERTS/bee9863e96afaa84fae21ba884a7370e6a0bc0fb5ac23eb94bf41983aa841a06/README.md) |
| 09&#160;Nov&#160;22&#160;18:56&#160;UTC | SHAKEN ConvergeTel LLC 388K | 09&#160;Dec&#160;22&#160;18:56&#160;UTC | true | [view](CERTS/c7fd0f44c8754d373dcca18b5ccaf082cdc1fe357f18b46803a105d3d2d6ec6b/README.md) |
| 10&#160;Nov&#160;22&#160;06:30&#160;UTC | SHAKEN  XCast Labs 689J | 11&#160;Dec&#160;22&#160;06:30&#160;UTC | true | [view](CERTS/402222a08d910e9aa8cd945c479d323cddbda54590a82d1ae2ef0023f33d9d80/README.md) |
| 10&#160;Nov&#160;22&#160;12:45&#160;UTC | SHAKEN InteractiveTel, LLC 920J | 10&#160;Dec&#160;22&#160;12:45&#160;UTC | true | [view](CERTS/a38b71eecf1a65b81996eb8b86989987aca9275bac11d583a6bfc2e31483ace8/README.md) |
| 10&#160;Nov&#160;22&#160;15:00&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 10&#160;Dec&#160;22&#160;15:00&#160;UTC | true | [view](CERTS/65e152732791c2c218f59531f0d67add39635289adb11ab77f3b356958f38ab0/README.md) |
| 10&#160;Nov&#160;22&#160;21:45&#160;UTC | SHAKEN Drop Inc 258K | 10&#160;Dec&#160;22&#160;21:45&#160;UTC | true | [view](CERTS/c4d7cbedcfefba961e4c61b9e7674c1a56dbfa3c036e6a89590edb0ece9019bb/README.md) |
| 11&#160;Nov&#160;22&#160;14:55&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 11&#160;Dec&#160;22&#160;14:55&#160;UTC | true | [view](CERTS/fc167560710eed74538e76603a2d40cde73816d8b8209490e0116f6e2dc0b5a9/README.md) |
| 14&#160;Nov&#160;22&#160;00:07&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 14&#160;Dec&#160;22&#160;00:07&#160;UTC | true | [view](CERTS/4097442cf48bd256c17ba8c6216a1628dd6173fa170b20c5044148ccc4f62de7/README.md) |
| 14&#160;Nov&#160;22&#160;14:40&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 14&#160;Dec&#160;22&#160;14:40&#160;UTC | true | [view](CERTS/a4d481efa9f4220995729cdb89ad8288cf550d625d73e2e36c02a63e5854ed25/README.md) |
| 15&#160;Nov&#160;22&#160;00:02&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 15&#160;Dec&#160;22&#160;00:02&#160;UTC | true | [view](CERTS/95eafe838e93ad32691c4a673afbedced0121bf5142d014aa9fbc4502f73df42/README.md) |
| 15&#160;Nov&#160;22&#160;19:11&#160;UTC | SHAKEN Momentum Telecom 1417 | 15&#160;Dec&#160;22&#160;19:11&#160;UTC | true | [view](CERTS/0ab37326807057d99a07b19947ef7f53dcfde3ed53492e528fd0572a46f9f9c1/README.md) |
| 15&#160;Nov&#160;22&#160;19:11&#160;UTC | SHAKEN Momentum Telecom 9157 | 15&#160;Dec&#160;22&#160;19:11&#160;UTC | true | [view](CERTS/32fc146d2eb7476b959e615591c31cebf8a2c3b09f630b58b0c8b32fdc55f959/README.md) |
| 15&#160;Nov&#160;22&#160;20:47&#160;UTC | SHAKEN Broadband Dynamics LLC 583j | 15&#160;Dec&#160;22&#160;20:47&#160;UTC | true | [view](CERTS/1314c42206284ec6487c7e86cc7e882a7b49e8d1b3bbd4be12a3e1407494a233/README.md) |
| 15&#160;Nov&#160;22&#160;20:54&#160;UTC | SHAKEN Threshold Communications Inc 563J | 15&#160;Dec&#160;22&#160;20:54&#160;UTC | true | [view](CERTS/27a0bd4e9c23c4a6e6cad5232e8c0812dbf0151ccfadea64bfb8e91883fd7d88/README.md) |
| 15&#160;Nov&#160;22&#160;20:56&#160;UTC | SHAKEN Consolidated Communications 5113 | 15&#160;Dec&#160;22&#160;20:56&#160;UTC | true | [view](CERTS/bd309b15066f2de4544ab567afe6db96b5be23ac91cef688b938ff58b8e3b091/README.md) |
| 15&#160;Nov&#160;22&#160;21:00&#160;UTC | SHAKEN Touchtone 683A | 15&#160;Dec&#160;22&#160;21:00&#160;UTC | true | [view](CERTS/37bd09282885e894569982f8ad1e592380285e2032767508974615af390fefd9/README.md) |
| 15&#160;Nov&#160;22&#160;21:02&#160;UTC | SHAKEN Apeiron Systems 012J | 15&#160;Dec&#160;22&#160;21:02&#160;UTC | true | [view](CERTS/89ad016561840539edaa6de902af87d7548e6a73b95ff47d4183b8fd7833c20b/README.md) |
| 15&#160;Nov&#160;22&#160;21:02&#160;UTC | SHAKEN Fonative, Inc. 684J | 15&#160;Dec&#160;22&#160;21:02&#160;UTC | true | [view](CERTS/c24083689456c887c751f152ef9aa7751bb8b6ed304ef1353494612069fdd4a4/README.md) |
| 15&#160;Nov&#160;22&#160;21:03&#160;UTC | SHAKEN IPitomy 652J | 15&#160;Dec&#160;22&#160;21:03&#160;UTC | true | [view](CERTS/bc8c3bbffbef28abab947c46fde46dd0795151de4b9486d52fb91d0e3ca23193/README.md) |
| 15&#160;Nov&#160;22&#160;21:04&#160;UTC | SHAKEN Phone.com, Inc. 633J | 15&#160;Dec&#160;22&#160;21:04&#160;UTC | true | [view](CERTS/2854defc6a8758dd27cf3655f48bce5e2db3f347538aefcdef67c49e74de627b/README.md) |
| 15&#160;Nov&#160;22&#160;21:04&#160;UTC | SHAKEN NETRIO LLC 020K | 15&#160;Dec&#160;22&#160;21:04&#160;UTC | true | [view](CERTS/05590af8e751e3bce1bec51db6a8b0a1e839cb2599138f3d89e984a5a4667179/README.md) |
| 15&#160;Nov&#160;22&#160;21:06&#160;UTC | SHAKEN VoIP Innovations 597F | 15&#160;Dec&#160;22&#160;21:06&#160;UTC | true | [view](CERTS/dd463b8fa534ffd3261de018c3a6a5080d10fb4457cf90807eea76b344752d2d/README.md) |
| 15&#160;Nov&#160;22&#160;21:06&#160;UTC | SHAKEN IDT America, Corp 363A | 15&#160;Dec&#160;22&#160;21:06&#160;UTC | true | [view](CERTS/d754834bfae4fbcfc91ca89e72cfa2ac6e8437ecbd70caae4e9971af37553f0b/README.md) |
| 15&#160;Nov&#160;22&#160;21:08&#160;UTC | SHAKEN Noble Systems Communications LLC 187J | 15&#160;Dec&#160;22&#160;21:08&#160;UTC | true | [view](CERTS/68596ea14ed54a4f373a20018c0cd17cc64c49569b0595b4a0bf06638f4d5445/README.md) |
| 15&#160;Nov&#160;22&#160;21:09&#160;UTC | SHAKEN Telcentris Inc. dba Voxox 696J | 15&#160;Dec&#160;22&#160;21:09&#160;UTC | true | [view](CERTS/f7e1211b2a732fb5ba09b47d98de86536f03a4fd7ab3f44b1c63aa6382a17a91/README.md) |
| 15&#160;Nov&#160;22&#160;21:10&#160;UTC | SHAKEN Airespring 996H | 15&#160;Dec&#160;22&#160;21:10&#160;UTC | true | [view](CERTS/7560a9d1569990b94ad8c978c95c9d61f745b15dce383433d34797548f57f395/README.md) |
| 15&#160;Nov&#160;22&#160;21:10&#160;UTC | SHAKEN Nobelbiz, Inc. 596J | 15&#160;Dec&#160;22&#160;21:10&#160;UTC | true | [view](CERTS/1a9ad822d22a48348a5ebc8692b2bbbc442172b28cf43e9ae28ba0cafed5cb20/README.md) |
| 15&#160;Nov&#160;22&#160;21:12&#160;UTC | SHAKEN Terra Nova Telecom 382G | 15&#160;Dec&#160;22&#160;21:12&#160;UTC | true | [view](CERTS/e3314c858960e9f66e4d11033b2ef1aeaa54794506dcd6ea8e01e9f4120bf976/README.md) |
| 15&#160;Nov&#160;22&#160;21:13&#160;UTC | SHAKEN Matrix 3058 | 15&#160;Dec&#160;22&#160;21:13&#160;UTC | true | [view](CERTS/9be01b3b668a930db37633e4f4ad35be67d1e14dabad3aa8510a0e2b9f2bb9cd/README.md) |
| 15&#160;Nov&#160;22&#160;21:14&#160;UTC | SHAKEN Matrix 9451 | 15&#160;Dec&#160;22&#160;21:14&#160;UTC | true | [view](CERTS/d0b669748645515139d27d551c74cd1a364beec2fcedfebb46621b770551a755/README.md) |
| 15&#160;Nov&#160;22&#160;21:14&#160;UTC | SHAKEN Matrix 7379 | 15&#160;Dec&#160;22&#160;21:14&#160;UTC | true | [view](CERTS/e41da9b1e48d8fdf606bfd088852fec6662504194a2622f9de0abe129f5509ff/README.md) |
| 15&#160;Nov&#160;22&#160;21:15&#160;UTC | SHAKEN PNG Telecommunications Inc 3395 | 15&#160;Dec&#160;22&#160;21:15&#160;UTC | true | [view](CERTS/147025b31a2b121cf511c20d46d9f1245de0a33ff56671ffc34939da2bb02fe1/README.md) |
| 15&#160;Nov&#160;22&#160;21:19&#160;UTC | SHAKEN Magna5, LLC 3849 | 15&#160;Dec&#160;22&#160;21:19&#160;UTC | true | [view](CERTS/c36de43ee0a9b4612d80bed1855eb699081a38df9f133fef41c1ffa459447b9c/README.md) |
| 15&#160;Nov&#160;22&#160;21:20&#160;UTC | SHAKEN Magna5, LLC 8249 | 15&#160;Dec&#160;22&#160;21:20&#160;UTC | true | [view](CERTS/7cc24bd53fa4555f91d216b251a55545b57070ac0c941e56e966562a4a4630b4/README.md) |
| 16&#160;Nov&#160;22&#160;14:30&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 16&#160;Dec&#160;22&#160;14:30&#160;UTC | true | [view](CERTS/0cfded3de974423502006a8ffa6de05a0576750a98a85052410ce78886149254/README.md) |
| 16&#160;Nov&#160;22&#160;23:52&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 16&#160;Dec&#160;22&#160;23:52&#160;UTC | true | [view](CERTS/bb44d024604a112054381618c76117b261b121ecae2dbabf308c56cc043e47dc/README.md) |
| 17&#160;Nov&#160;22&#160;14:10&#160;UTC | SHAKEN 1stPoint Communications, LLC 463G | 17&#160;Dec&#160;22&#160;14:10&#160;UTC | true | [view](CERTS/7966daa2575a676e64515d90b5037c1473eac821e2f55c5ff8ddd3e1d2453f05/README.md) |
| 17&#160;Nov&#160;22&#160;14:25&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 17&#160;Dec&#160;22&#160;14:25&#160;UTC | true | [view](CERTS/9da75410fb2ab55dd04bc69fe3c917b591909fc01af838f793348cd890720efa/README.md) |
| 17&#160;Nov&#160;22&#160;20:14&#160;UTC | SHAKEN IPSBS Managed Services LLC 828J | 17&#160;Dec&#160;22&#160;20:14&#160;UTC | true | [view](CERTS/273c6545be615c3dcce645179486c9e00ae37b1fbee1440ccd94cc28720e6fe1/README.md) |
| 17&#160;Nov&#160;22&#160;23:47&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 17&#160;Dec&#160;22&#160;23:47&#160;UTC | true | [view](CERTS/4fe1804a715815dfb130fbd80f4687bf31f2a4ea79233e48c1cd27c30fa6a2f8/README.md) |
| 18&#160;Nov&#160;22&#160;14:20&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 18&#160;Dec&#160;22&#160;14:20&#160;UTC | true | [view](CERTS/ef0e092d74d92c39305b38f4d157d4dd03149743c20fdf6ab2643b02e7f618bd/README.md) |
| 19&#160;Nov&#160;22&#160;12:00&#160;UTC | SHAKEN InteractiveTel, LLC 920J | 19&#160;Dec&#160;22&#160;12:00&#160;UTC | true | [view](CERTS/84096b06593e0ddb60a331132dc601eeb5387365b3ed898ea4c2607b58c0e63b/README.md) |
| 19&#160;Nov&#160;22&#160;14:15&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 19&#160;Dec&#160;22&#160;14:15&#160;UTC | true | [view](CERTS/126f5c685a3e684b10371399918f9d37580e0a9e45eed6c795ac380aacdfab4d/README.md) |
| 20&#160;Nov&#160;22&#160;23:32&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 20&#160;Dec&#160;22&#160;23:32&#160;UTC | true | [view](CERTS/1f8fa64ee3dec6393a22f3c7958b609ebba84daaa737011c3182f39d997f73c5/README.md) |
| 21&#160;Nov&#160;22&#160;04:07&#160;UTC | SHAKEN NTC International, INC 016K | 21&#160;Dec&#160;22&#160;04:07&#160;UTC | true | [view](CERTS/f87e145d0fea830dcab24a7a3c054aab4e142a0ff87db63aa2c6c4d9fd6b5c84/README.md) |
| 21&#160;Nov&#160;22&#160;11:50&#160;UTC | SHAKEN InteractiveTel, LLC 920J | 21&#160;Dec&#160;22&#160;11:50&#160;UTC | true | [view](CERTS/2322461ea6facdf55eff9dbf53b51b42a8d3c24dc179cbcba28e05a5da0600b6/README.md) |
| 21&#160;Nov&#160;22&#160;14:05&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 21&#160;Dec&#160;22&#160;14:05&#160;UTC | true | [view](CERTS/e38231cf718dbea66eda1d6e7034371478d390926df5b627f0ce9d9e4ce028da/README.md) |
| 21&#160;Nov&#160;22&#160;19:54&#160;UTC | SHAKEN IPSBS Managed Services LLC 828J | 21&#160;Dec&#160;22&#160;19:54&#160;UTC | true | [view](CERTS/d8fb28426197f31ebb43aa58286b50bf249430a1f3c1145b8260f33ee5b54c43/README.md) |
| 21&#160;Nov&#160;22&#160;21:15&#160;UTC | SHAKEN Star2Star Communications, LLC 590J | 21&#160;Nov&#160;23&#160;21:15&#160;UTC | true | [view](CERTS/9bc9dde8921387803d93036c7d2f8085af32b028fca8f17336d2e22ab51fd278/README.md) |
| 21&#160;Nov&#160;22&#160;22:15&#160;UTC | SHAKEN Global Net Holdings Inc 306K | 21&#160;Dec&#160;22&#160;22:15&#160;UTC | true | [view](CERTS/2da401f846b27ab52c1e3165eb3a05817cceeeb85473df066979c2af2d90b9f8/README.md) |
| 21&#160;Nov&#160;22&#160;23:27&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 21&#160;Dec&#160;22&#160;23:27&#160;UTC | true | [view](CERTS/e96c75c714a476ff6540d344291ae4f45b201ee3bb6fc1488be9299e48a9b443/README.md) |
| 22&#160;Nov&#160;22&#160;04:02&#160;UTC | SHAKEN NTC International, INC 016K | 22&#160;Dec&#160;22&#160;04:02&#160;UTC | true | [view](CERTS/23f33ad7a9251d08069b73258d15317cbdb903645adeb3f8e0db9147660895d4/README.md) |
| 22&#160;Nov&#160;22&#160;11:45&#160;UTC | SHAKEN InteractiveTel, LLC 920J | 22&#160;Dec&#160;22&#160;11:45&#160;UTC | true | [view](CERTS/9b65e5ddc85ff994ab983b4e2e98241ca756a0840237aedd031d34b83cc0894a/README.md) |
| 22&#160;Nov&#160;22&#160;13:45&#160;UTC | SHAKEN 1stPoint Communications, LLC 463G | 22&#160;Dec&#160;22&#160;13:45&#160;UTC | true | [view](CERTS/fd6d73b3c2895571a436b1d5301fe14cfcf8485ed70415118ef0a2c07e9e435e/README.md) |
| 22&#160;Nov&#160;22&#160;14:00&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 22&#160;Dec&#160;22&#160;14:00&#160;UTC | true | [view](CERTS/8505f217db21211d544c87d27f932ad737b4a44956ebef5c0b546169c2c5a4fb/README.md) |
| 22&#160;Nov&#160;22&#160;23:22&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 22&#160;Dec&#160;22&#160;23:22&#160;UTC | true | [view](CERTS/413854a61ea72db633497dd251bfb85f7dd9dfe3e0bfecd568198af1910879c6/README.md) |
| 23&#160;Nov&#160;22&#160;03:57&#160;UTC | SHAKEN NTC International, INC 016K | 23&#160;Dec&#160;22&#160;03:57&#160;UTC | true | [view](CERTS/09409ec33539419a0321c54eaad06baeccf5708bce2112dd1a8499731d7c8aee/README.md) |
| 23&#160;Nov&#160;22&#160;11:40&#160;UTC | SHAKEN InteractiveTel, LLC 920J | 23&#160;Dec&#160;22&#160;11:40&#160;UTC | true | [view](CERTS/86b6da6915cbf408b97937eca513d5494b79302b558310e3a4fdf4eacfb8df39/README.md) |
| 23&#160;Nov&#160;22&#160;13:55&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 23&#160;Dec&#160;22&#160;13:55&#160;UTC | true | [view](CERTS/54089cd5eaf505b050251f8201fdc1296b9b3ce3b0ebc0ff684ae5102dc9fa5d/README.md) |
| 24&#160;Nov&#160;22&#160;23:12&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 24&#160;Dec&#160;22&#160;23:12&#160;UTC | true | [view](CERTS/e6410ba9fa07944c517ccab0a073939d061571acd1a28a96232d630020797a91/README.md) |
| 25&#160;Nov&#160;22&#160;13:45&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 25&#160;Dec&#160;22&#160;13:45&#160;UTC | true | [view](CERTS/3291177e2e236b6f31b28a8387fc7bf27560a19acb825da2a835d8f006bfa549/README.md) |
| 26&#160;Nov&#160;22&#160;03:42&#160;UTC | SHAKEN NTC International, INC 016K | 26&#160;Dec&#160;22&#160;03:42&#160;UTC | true | [view](CERTS/b7fa79f7dc6eeef8b3c02a8b1db0bcd102d85777210e7f7e14eb2f3cecf27974/README.md) |
| 26&#160;Nov&#160;22&#160;13:40&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 26&#160;Dec&#160;22&#160;13:40&#160;UTC | true | [view](CERTS/73f0b8709ec9702c103d7bbd9d808ba041783cc35a7492fa6e47bc1672b2ba6b/README.md) |
| 27&#160;Nov&#160;22&#160;22:57&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 27&#160;Dec&#160;22&#160;22:57&#160;UTC | true | [view](CERTS/dfdb808f2f94929ae6847f80fddc6fbcf54d7cbf490c28090e4fea49012bbf1d/README.md) |
| 28&#160;Nov&#160;22&#160;13:30&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 28&#160;Dec&#160;22&#160;13:30&#160;UTC | true | [view](CERTS/e15f0ef42f7371f303897b998c174f2800d8d3ee2671ded8c0254716795ccafd/README.md) |
| 28&#160;Nov&#160;22&#160;22:52&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 28&#160;Dec&#160;22&#160;22:52&#160;UTC | true | [view](CERTS/597d2bfba3f5c1445f03792348e4cd98ab8cc8b44bd896be347489510679a137/README.md) |
| 29&#160;Nov&#160;22&#160;03:27&#160;UTC | SHAKEN NTC International, INC 016K | 29&#160;Dec&#160;22&#160;03:27&#160;UTC | true | [view](CERTS/24fe23bbad3fa34c0c9f7bcd9e36a45bf0974a82c2b574780a3d582f301f049e/README.md) |
| 29&#160;Nov&#160;22&#160;13:10&#160;UTC | SHAKEN 1stPoint Communications, LLC 463G | 29&#160;Dec&#160;22&#160;13:10&#160;UTC | true | [view](CERTS/cf147be7a87bd98ac78358181238d18bd46fe0bc992f25e5a875098cd27a0ba9/README.md) |
| 29&#160;Nov&#160;22&#160;13:25&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 29&#160;Dec&#160;22&#160;13:25&#160;UTC | true | [view](CERTS/91a428d4713a6d944a1c0509925d0df05e24d9e2490457099544333a10cd9fbf/README.md) |
| 29&#160;Nov&#160;22&#160;16:38&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 29&#160;Dec&#160;22&#160;16:38&#160;UTC | true | [view](CERTS/fa5d3ad33ea9b4bd7b4e0dffe89a7147e629942573cc303d1b404f096ec4de76/README.md) |
| 29&#160;Nov&#160;22&#160;22:04&#160;UTC | SHAKEN MagicJack 324E | 29&#160;Nov&#160;23&#160;22:04&#160;UTC | true | [view](CERTS/75b4b7b400b1252e48faa1d93f6a94f7bd4a6383c88ddf6baa167b85d9ac4ee8/README.md) |
| 29&#160;Nov&#160;22&#160;22:47&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 29&#160;Dec&#160;22&#160;22:47&#160;UTC | true | [view](CERTS/518f8e632dd4bb454ac599c8dfea1c3d51b289ae8623627dbef3abedca67ff8d/README.md) |
| 30&#160;Nov&#160;22&#160;03:22&#160;UTC | SHAKEN NTC International, INC 016K | 30&#160;Dec&#160;22&#160;03:22&#160;UTC | true | [view](CERTS/9c5bef68583aef1a5034f87aa532bc676fb1d40f81325b536dc6336f3e295020/README.md) |
| 30&#160;Nov&#160;22&#160;11:19&#160;UTC | SHAKEN InteractiveTel, LLC 920J | 30&#160;Dec&#160;22&#160;11:19&#160;UTC | true | [view](CERTS/0629f9b5eba88aefc3c032e28da28ac6467ae0e7bddbe2503db61e0e3b6ea28f/README.md) |
| 30&#160;Nov&#160;22&#160;13:20&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 30&#160;Dec&#160;22&#160;13:20&#160;UTC | true | [view](CERTS/408a130690cd123cd4bae17e157b93eb41d7355c82eef7bb36b7253ead7cf31a/README.md) |

#### CA Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 21&#160;Aug&#160;20&#160;01:22&#160;UTC | SHAKEN Sansay Root CA US | 16&#160;Aug&#160;40&#160;01:22&#160;UTC | true | [view](CERTS/8356d251b255f4ac3fd108bb79be4c02dcea2d3b13d138892bdb3a70cbe6a343/README.md) |
| 02&#160;Sep&#160;22&#160;20:53&#160;UTC | SHAKEN Sansay Intermediate CA US WEST 1 | 31&#160;Aug&#160;29&#160;20:53&#160;UTC | true | [view](CERTS/4b1dfdba2b1e4bbffbf900a20f1f6f7befbef0008b963e4922a64469cb97d24b/README.md) |


Generated: 02 Dec 22 07:30 UTC