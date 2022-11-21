# STIR/SHAKEN CA Ecosystem Compliance

## Sansay

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 148 certificates were included in the corpus being tested
- 3 certificates in the corpus were skipped because they are duplicates
- 54 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 91 certificates being tested against the remaining rules
- 5.46 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 100.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 98 days is the average remaining validity for the certificates in the corpus
- 99 days is the average initial validity for the certificates in the corpus
- 70 certificates expire in the next 30 days
- 1.57 average number of unexpired certificates per OCN observed
- 58 unique OCNs observed in unexpired and valid certificate corpus

| Instances | Test | Source |
|-----------|------|--------|
| 91 | [e_atis_certificate_policies](ISSUES/e_atis_certificate_policies/README.md) | ATIS1000080 |
| 91 | [e_atis_subject_cn](ISSUES/e_atis_subject_cn/README.md) | ATIS1000080 |
| 91 | [e_us_cp_ambiguous_identifier](ISSUES/e_us_cp_ambiguous_identifier/README.md) | US_SHAKEN_CP |
| 91 | [e_us_cp_subject_sn](ISSUES/e_us_cp_subject_sn/README.md) | US_SHAKEN_CP |
| 42 | [w_cp_1_3_subject_email](ISSUES/w_cp_1_3_subject_email/README.md) | US_SHAKEN_CP |
| 91 | [w_shaken_subject_rdn_unknown](ISSUES/w_shaken_subject_rdn_unknown/README.md) | SHAKEN_PKI_BEST_PRACTICES |

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
- 5444 days is the average remaining validity for the certificates in the corpus
- 4928 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days

| Instances | Test | Source |
|-----------|------|--------|
| 2 | [w_shaken_ca_subject_rdn_unknown](ISSUES/w_shaken_ca_subject_rdn_unknown/README.md) | SHAKEN_PKI_BEST_PRACTICES |

### Details

#### Leaf Certificates

| Created at | Subject | Problems | Link |
|------------|---------|----------|------|
| 12 Sep 22 19:37 UTC | SHAKEN Bulk Solutions, LLC 644J | true | [view](CERTS/2b980444a4603ddf16248bee9dbdce112f593d4d5324443e641624a827af0cb2/README.md) |
| 23 Sep 22 01:10 UTC | SHAKEN Star2Star Communications, LLC 590J | true | [view](CERTS/b6c27ce63b22687fcd2f9f64ee9067dd3c19a4eb223f1aef3934f7ba95c54ba6/README.md) |
| 01 Oct 22 00:00 UTC | SHAKEN IP Link Telecom Inc. 902J | true | [view](CERTS/53d491627ef52772ffd75835f745e895f025bfc853523028b42c0bc1681e5f67/README.md) |
| 05 Oct 22 19:08 UTC | SHAKEN IP Link Telecom Inc. 902J | true | [view](CERTS/fa453a929f14e705532a1974b216dedb71ba1f10d6af0df07f1084bfbb2038db/README.md) |
| 11 Oct 22 16:48 UTC | SHAKEN TeleVoIPs 138K | true | [view](CERTS/c41b66127049dbae159f8d68ac714616b9e99640c407bcdc749f3d49037db487/README.md) |
| 11 Oct 22 17:08 UTC | SHAKEN ALD Telecom 780J | true | [view](CERTS/53a14081c994555770bb8c5f3d160f89cf427258c9598d569c388a74bde6ea8f/README.md) |
| 11 Oct 22 17:14 UTC | SHAKEN Technology Innovation Lab 599J | true | [view](CERTS/07d98b6eeb180548fa10e06aedbd69ce0816a1040344c91d25b8dcf29f68e7e6/README.md) |
| 11 Oct 22 17:17 UTC | SHAKEN Current Calls, LLC 746J | true | [view](CERTS/52d6a93a1b72d2f2980699e759068dd9dbc8314c953e03613f18d9da1dcf156d/README.md) |
| 11 Oct 22 17:18 UTC | SHAKEN Global Net Holdings Inc 306K | true | [view](CERTS/7a67400c89c424c8f378f841d333cdd9d94ecc2e681339608615d0aee51a959c/README.md) |
| 11 Oct 22 17:19 UTC | SHAKEN Inventive Labs Corp 649J | true | [view](CERTS/a2f02cfef1eba726cf7dbd0f018a1119d40600aba568619f16b4c08b8d3a7c12/README.md) |
| 11 Oct 22 17:20 UTC | SHAKEN Carrier One Inc. 705J | true | [view](CERTS/a7447339990a198aac3d84ed38d80706e16b7aac171e6d6bd1b28275fe7c337e/README.md) |
| 11 Oct 22 17:21 UTC | SHAKEN Asia Pacific Network 988J | true | [view](CERTS/0b191ba4d02eaa4b595b67a4d3e6f35a6d6c184e5b7e464d471cb904ea2d0638/README.md) |
| 11 Oct 22 17:22 UTC | SHAKEN OneStream Networks, LLC 630J | true | [view](CERTS/f18d0d387f4abfadaa336e2ff00c0f6b0509898b7d2d54feb99e1e0fb2042d3a/README.md) |
| 11 Oct 22 17:23 UTC | SHAKEN Ringfree Communications Inc 317K | true | [view](CERTS/cc75f739ba9e082e5324936f9c5c1df2d896cb259ed0dd51065b937a0fce25aa/README.md) |
| 11 Oct 22 17:24 UTC | SHAKEN Xchange Telecom LLC 325B | true | [view](CERTS/6bab691174d8e7b237a7fe1b00556840e2a5c28a1839f8e345dd9ba721ba23bb/README.md) |
| 11 Oct 22 17:24 UTC | SHAKEN Sangoma 777G | true | [view](CERTS/53d28ac1fa5253468c11b9e3baaa6ad5481e83a7ea2ee6d715594dc6d4561ad4/README.md) |
| 12 Oct 22 12:50 UTC | SHAKEN Lightspeed Voice 557F | true | [view](CERTS/ab19df868054cb3392aa295bff737bf919f8dc55c64a91247621375bad7fb7c0/README.md) |
| 13 Oct 22 20:27 UTC | SHAKEN ConnectMeVoice 719J | true | [view](CERTS/a5edeeacfcec8ad6584f5a0b505978c4b72907a2e3a6540bb01350397f86814e/README.md) |
| 24 Oct 22 01:52 UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | true | [view](CERTS/6b38e94aa75bbddcc98a5b6c62853dddd130b7e2dadffd014a723f8664ed6b16/README.md) |
| 24 Oct 22 16:25 UTC | SHAKEN Quality Voice & Data Inc. 548J | true | [view](CERTS/7308cc01648f3523bcc3834e2c41e9846c4004e96159c7d0a9a5b90008ab42c3/README.md) |
| 24 Oct 22 20:23 UTC | SHAKEN Arbeit 816J | true | [view](CERTS/377e182a223e6cc8d7e9ce697e7a3e829b1c6b16c299c26f6d1f1e33aa29524b/README.md) |
| 24 Oct 22 21:11 UTC | SHAKEN Ytel Inc. 703J | true | [view](CERTS/3d6a7a2ff23b90fba1674f600a108b8a11a110f8bb1723df86627001f7367d8d/README.md) |
| 25 Oct 22 01:47 UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | true | [view](CERTS/878ab560cdf2d47147ab378eb62e43f12052e41097ac724967824263891d4f1b/README.md) |
| 25 Oct 22 06:22 UTC | SHAKEN NTC International, INC 016K | true | [view](CERTS/2b6ef0741d09c6a0221ce9bfc7ca8edeb972333214bb5d2a970215876a32d4a0/README.md) |
| 26 Oct 22 16:15 UTC | SHAKEN Quality Voice & Data Inc. 548J | true | [view](CERTS/5ccc340e70f7e4fd89a1d3b14455e6c559f76968a24bf27ac935e5595ceed7cf/README.md) |
| 26 Oct 22 16:49 UTC | SHAKEN InteractiveTel, LLC 920J | true | [view](CERTS/b5416e5eefd0dc2b74de073c207a83076d39bd7b88f9bfee3fe21b7418668c59/README.md) |
| 27 Oct 22 01:37 UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | true | [view](CERTS/8b0bb57dd9a98814d02347a4713fe7a0c3580c5341063a8cc60da4e8bd1dd2bf/README.md) |
| 27 Oct 22 16:10 UTC | SHAKEN Quality Voice & Data Inc. 548J | true | [view](CERTS/c98d31acc1d88aa6b8f23c05c1a3946642950394fc684d816d0196e708e21dbe/README.md) |
| 27 Oct 22 19:21 UTC | SHAKEN IPSBS Managed Services LLC 828J | true | [view](CERTS/a2a30355ac1d9f82570fd570ed6d0a330f2a2d41d6509428d6be20566cd377da/README.md) |
| 27 Oct 22 19:32 UTC | SHAKEN Zray Technologies Corporation 862J | true | [view](CERTS/67f53328009e04f0f6e9412d26d0c77eeeac1e3106e5598dd5a5a21493b4f999/README.md) |
| 27 Oct 22 20:11 UTC | SHAKEN Mitel Cloud Services, Inc. 670J | true | [view](CERTS/e45c92abcfe2fe6d0863200900b66e835aa98712f974efe3837e34d787f2ad5e/README.md) |
| 28 Oct 22 01:32 UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | true | [view](CERTS/9ab18008bc722a7eefa9a015f462aa24ccedac3b667a5425ba9ec3b124ce5909/README.md) |
| 28 Oct 22 06:07 UTC | SHAKEN NTC International, INC 016K | true | [view](CERTS/fd05b359b2354a6ee27eef74636343f0d0e5ea19f1f64fc8e082a1a190cb34d9/README.md) |
| 28 Oct 22 16:05 UTC | SHAKEN Quality Voice & Data Inc. 548J | true | [view](CERTS/fc9997a29b7cfedf5a4d3c899d7ab2ef73e316e8c83ad2efa997d33e34e7d6f5/README.md) |
| 28 Oct 22 16:39 UTC | SHAKEN InteractiveTel, LLC 920J | true | [view](CERTS/65d9f8a0ca34ff759de8d65a7db8f81b4d66880b0c05969c777496942e309f12/README.md) |
| 30 Oct 22 15:55 UTC | SHAKEN Quality Voice & Data Inc. 548J | true | [view](CERTS/c171a9efe723a6e1c3e773cba35e79b6eec435d305017ff54e088ba63c3d0828/README.md) |
| 31 Oct 22 15:50 UTC | SHAKEN Quality Voice & Data Inc. 548J | true | [view](CERTS/a079e0fcf48c3f26620d8674a59ed2320f58dafee72d4984ef6548e7be9b000c/README.md) |
| 31 Oct 22 20:18 UTC | SHAKEN Cyberlynk Network, LLC 086K | true | [view](CERTS/081c6b3caf13a694f3dc2feaaa56193480879b726acc425e65d8fb37a3e28cbc/README.md) |
| 01 Nov 22 15:45 UTC | SHAKEN Quality Voice & Data Inc. 548J | true | [view](CERTS/8660f3e745ab090b45cf80648e48b2cc83449bb33a650bee29b9c4c498d268d4/README.md) |
| 01 Nov 22 19:07 UTC | SHAKEN Zray Technologies Corporation 862J | true | [view](CERTS/3a227acdee5e639036b62df263e132c42bb684a9d6273fe6611464dbe434b2db/README.md) |
| 02 Nov 22 16:14 UTC | SHAKEN InteractiveTel, LLC 920J | true | [view](CERTS/0ced8150d7cbed20775669aafeecd05726f233c901d024bb4fa64e9ce75b9f7a/README.md) |
| 03 Nov 22 15:35 UTC | SHAKEN Quality Voice & Data Inc. 548J | true | [view](CERTS/e05e81f7e34f2088102df66b64b44c5549708f9182a79e6e7d1157333260d98b/README.md) |
| 04 Nov 22 15:30 UTC | SHAKEN Quality Voice & Data Inc. 548J | true | [view](CERTS/51862e962c0ec2667b8ee2862f0eee4d50fcdaf629bfaa3e7191ed7018f56db1/README.md) |
| 05 Nov 22 06:00 UTC | SHAKEN Convoso 758J | true | [view](CERTS/a33008ec0c4ad56c980d02e36d07e33a42f8c2f203fea78024c5018d5cf04e8d/README.md) |
| 07 Nov 22 00:42 UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | true | [view](CERTS/d27a959a21dc295781c932032c24d6e53b61a5088ab3d2cc98ac93562c5ec6cd/README.md) |
| 07 Nov 22 15:15 UTC | SHAKEN Quality Voice & Data Inc. 548J | true | [view](CERTS/16593a1d36bc2265a80aa16a52e6936c5004486f6a84266e6b9d166c82509a8c/README.md) |
| 07 Nov 22 18:27 UTC | SHAKEN IPSBS Managed Services LLC 828J | true | [view](CERTS/5c216d5eb90d5b09e51db9c9a1b3bad35f94f504a92261dc65dde90913fb410a/README.md) |
| 08 Nov 22 14:55 UTC | SHAKEN 1stPoint Communications, LLC 463G | true | [view](CERTS/322096d99c6f6beb803c72f64ba15430bf34292da2e7736382405346d1920c2e/README.md) |
| 08 Nov 22 15:10 UTC | SHAKEN Quality Voice & Data Inc. 548J | true | [view](CERTS/0bab2fef8838a943dc8b17e3632776a4f4d8e74508ca22efbeb0dd282adde5d2/README.md) |
| 08 Nov 22 18:18 UTC | SHAKEN Primo Dialler LLC 249K | true | [view](CERTS/1865568906d0385d9cdfa133305a3aeea8bb25f8630ffc630c654a8d47273168/README.md) |
| 09 Nov 22 02:14 UTC | SHAKEN Connexum LLC 203K | true | [view](CERTS/b6f81e35a57af695e4d70f61d18bac7cbd8f823b687f6adbb91561da8ecaab81/README.md) |
| 09 Nov 22 12:51 UTC | SHAKEN InteractiveTel, LLC 920J | true | [view](CERTS/b032d7edd0943eabb4b18af4e272cc93270997d66ce6e5051193454be2dd7369/README.md) |
| 09 Nov 22 14:50 UTC | SHAKEN 1stPoint Communications, LLC 463G | true | [view](CERTS/b4e31ab7b8fc550e75a7b00886b00d69db00cd56a13c4dd2cf8f1911749a4dcb/README.md) |
| 09 Nov 22 15:05 UTC | SHAKEN Quality Voice & Data Inc. 548J | true | [view](CERTS/31721e97e825e9f221be10049da5735c86acdafc29714b581cf12e22acc72bb6/README.md) |
| 09 Nov 22 18:49 UTC | SHAKEN Primo Dialler LLC 249K | true | [view](CERTS/bee9863e96afaa84fae21ba884a7370e6a0bc0fb5ac23eb94bf41983aa841a06/README.md) |
| 10 Nov 22 06:30 UTC | SHAKEN  XCast Labs 689J | true | [view](CERTS/402222a08d910e9aa8cd945c479d323cddbda54590a82d1ae2ef0023f33d9d80/README.md) |
| 10 Nov 22 12:45 UTC | SHAKEN InteractiveTel, LLC 920J | true | [view](CERTS/a38b71eecf1a65b81996eb8b86989987aca9275bac11d583a6bfc2e31483ace8/README.md) |
| 10 Nov 22 15:00 UTC | SHAKEN Quality Voice & Data Inc. 548J | true | [view](CERTS/65e152732791c2c218f59531f0d67add39635289adb11ab77f3b356958f38ab0/README.md) |
| 10 Nov 22 21:45 UTC | SHAKEN Drop Inc 258K | true | [view](CERTS/c4d7cbedcfefba961e4c61b9e7674c1a56dbfa3c036e6a89590edb0ece9019bb/README.md) |
| 11 Nov 22 14:55 UTC | SHAKEN Quality Voice & Data Inc. 548J | true | [view](CERTS/fc167560710eed74538e76603a2d40cde73816d8b8209490e0116f6e2dc0b5a9/README.md) |
| 14 Nov 22 00:07 UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | true | [view](CERTS/4097442cf48bd256c17ba8c6216a1628dd6173fa170b20c5044148ccc4f62de7/README.md) |
| 14 Nov 22 14:40 UTC | SHAKEN Quality Voice & Data Inc. 548J | true | [view](CERTS/a4d481efa9f4220995729cdb89ad8288cf550d625d73e2e36c02a63e5854ed25/README.md) |
| 15 Nov 22 00:02 UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | true | [view](CERTS/95eafe838e93ad32691c4a673afbedced0121bf5142d014aa9fbc4502f73df42/README.md) |
| 15 Nov 22 19:11 UTC | SHAKEN Momentum Telecom 1417 | true | [view](CERTS/0ab37326807057d99a07b19947ef7f53dcfde3ed53492e528fd0572a46f9f9c1/README.md) |
| 15 Nov 22 19:11 UTC | SHAKEN Momentum Telecom 9157 | true | [view](CERTS/32fc146d2eb7476b959e615591c31cebf8a2c3b09f630b58b0c8b32fdc55f959/README.md) |
| 15 Nov 22 20:47 UTC | SHAKEN Broadband Dynamics LLC 583j | true | [view](CERTS/1314c42206284ec6487c7e86cc7e882a7b49e8d1b3bbd4be12a3e1407494a233/README.md) |
| 15 Nov 22 20:54 UTC | SHAKEN Threshold Communications Inc 563J | true | [view](CERTS/27a0bd4e9c23c4a6e6cad5232e8c0812dbf0151ccfadea64bfb8e91883fd7d88/README.md) |
| 15 Nov 22 20:56 UTC | SHAKEN Consolidated Communications 5113 | true | [view](CERTS/bd309b15066f2de4544ab567afe6db96b5be23ac91cef688b938ff58b8e3b091/README.md) |
| 15 Nov 22 21:00 UTC | SHAKEN Touchtone 683A | true | [view](CERTS/37bd09282885e894569982f8ad1e592380285e2032767508974615af390fefd9/README.md) |
| 15 Nov 22 21:02 UTC | SHAKEN Apeiron Systems 012J | true | [view](CERTS/89ad016561840539edaa6de902af87d7548e6a73b95ff47d4183b8fd7833c20b/README.md) |
| 15 Nov 22 21:02 UTC | SHAKEN Fonative, Inc. 684J | true | [view](CERTS/c24083689456c887c751f152ef9aa7751bb8b6ed304ef1353494612069fdd4a4/README.md) |
| 15 Nov 22 21:03 UTC | SHAKEN IPitomy 652J | true | [view](CERTS/bc8c3bbffbef28abab947c46fde46dd0795151de4b9486d52fb91d0e3ca23193/README.md) |
| 15 Nov 22 21:04 UTC | SHAKEN Phone.com, Inc. 633J | true | [view](CERTS/2854defc6a8758dd27cf3655f48bce5e2db3f347538aefcdef67c49e74de627b/README.md) |
| 15 Nov 22 21:04 UTC | SHAKEN NETRIO LLC 020K | true | [view](CERTS/05590af8e751e3bce1bec51db6a8b0a1e839cb2599138f3d89e984a5a4667179/README.md) |
| 15 Nov 22 21:06 UTC | SHAKEN VoIP Innovations 597F | true | [view](CERTS/dd463b8fa534ffd3261de018c3a6a5080d10fb4457cf90807eea76b344752d2d/README.md) |
| 15 Nov 22 21:06 UTC | SHAKEN IDT America, Corp 363A | true | [view](CERTS/d754834bfae4fbcfc91ca89e72cfa2ac6e8437ecbd70caae4e9971af37553f0b/README.md) |
| 15 Nov 22 21:08 UTC | SHAKEN Noble Systems Communications LLC 187J | true | [view](CERTS/68596ea14ed54a4f373a20018c0cd17cc64c49569b0595b4a0bf06638f4d5445/README.md) |
| 15 Nov 22 21:09 UTC | SHAKEN Telcentris Inc. dba Voxox 696J | true | [view](CERTS/f7e1211b2a732fb5ba09b47d98de86536f03a4fd7ab3f44b1c63aa6382a17a91/README.md) |
| 15 Nov 22 21:10 UTC | SHAKEN Airespring 996H | true | [view](CERTS/7560a9d1569990b94ad8c978c95c9d61f745b15dce383433d34797548f57f395/README.md) |
| 15 Nov 22 21:10 UTC | SHAKEN Nobelbiz, Inc. 596J | true | [view](CERTS/1a9ad822d22a48348a5ebc8692b2bbbc442172b28cf43e9ae28ba0cafed5cb20/README.md) |
| 15 Nov 22 21:12 UTC | SHAKEN Terra Nova Telecom 382G | true | [view](CERTS/e3314c858960e9f66e4d11033b2ef1aeaa54794506dcd6ea8e01e9f4120bf976/README.md) |
| 15 Nov 22 21:13 UTC | SHAKEN Matrix 3058 | true | [view](CERTS/9be01b3b668a930db37633e4f4ad35be67d1e14dabad3aa8510a0e2b9f2bb9cd/README.md) |
| 15 Nov 22 21:14 UTC | SHAKEN Matrix 9451 | true | [view](CERTS/d0b669748645515139d27d551c74cd1a364beec2fcedfebb46621b770551a755/README.md) |
| 15 Nov 22 21:14 UTC | SHAKEN Matrix 7379 | true | [view](CERTS/e41da9b1e48d8fdf606bfd088852fec6662504194a2622f9de0abe129f5509ff/README.md) |
| 15 Nov 22 21:15 UTC | SHAKEN PNG Telecommunications Inc 3395 | true | [view](CERTS/147025b31a2b121cf511c20d46d9f1245de0a33ff56671ffc34939da2bb02fe1/README.md) |
| 15 Nov 22 21:19 UTC | SHAKEN Magna5, LLC 3849 | true | [view](CERTS/c36de43ee0a9b4612d80bed1855eb699081a38df9f133fef41c1ffa459447b9c/README.md) |
| 15 Nov 22 21:20 UTC | SHAKEN Magna5, LLC 8249 | true | [view](CERTS/7cc24bd53fa4555f91d216b251a55545b57070ac0c941e56e966562a4a4630b4/README.md) |
| 16 Nov 22 14:30 UTC | SHAKEN Quality Voice & Data Inc. 548J | true | [view](CERTS/0cfded3de974423502006a8ffa6de05a0576750a98a85052410ce78886149254/README.md) |
| 17 Nov 22 23:47 UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | true | [view](CERTS/4fe1804a715815dfb130fbd80f4687bf31f2a4ea79233e48c1cd27c30fa6a2f8/README.md) |
| 18 Nov 22 22:18 UTC | SHAKEN MagicJack 324E | true | [view](CERTS/609a7a6da13ee8652614edaab60d1c03920914276462e66145b9c2438130eaf8/README.md) |
| 19 Nov 22 14:15 UTC | SHAKEN Quality Voice & Data Inc. 548J | true | [view](CERTS/126f5c685a3e684b10371399918f9d37580e0a9e45eed6c795ac380aacdfab4d/README.md) |

#### CA Certificates

| Created at | Subject | Problems | Link |
|------------|---------|----------|------|
| 21 Aug 20 01:22 UTC | SHAKEN Sansay Root CA US | true | [view](CERTS/8356d251b255f4ac3fd108bb79be4c02dcea2d3b13d138892bdb3a70cbe6a343/README.md) |
| 02 Sep 22 20:53 UTC | SHAKEN Sansay Intermediate CA US WEST 1 | true | [view](CERTS/4b1dfdba2b1e4bbffbf900a20f1f6f7befbef0008b963e4922a64469cb97d24b/README.md) |


Generated: 21 Nov 22 20:55 UTC