# STIR/SHAKEN CA Ecosystem Compliance

## Sansay

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

85 certificates were included in the corpus being tested\
25 certificates in the corpus were skipped because they were expired\
0 certificates in the corpus were skipped because they are not currently trusted\
100.00% of certificates contain one or more Error level issue\
100.00% of certificates contain one or more Warning level issue\
0.00% of certificates contain one or more Notice level issue\
0.00% of certificates are too old to be assessed against currently enforced expectations\
99 days is the average remaining validity for the certificates in the corpus\
100 days is the average initial validity for the certificates in the corpus\
65 certificates expire in the next 30 days

| Instances | Test Status | Source |
|-----------|-------------|--------|
| 1 | [e_cp1_3_ambiguous_identifier](ISSUES/e_cp1_3_ambiguous_identifier/README.md) | United States SHAKEN CP |
| 1 | [e_cp1_3_subject_sn](ISSUES/e_cp1_3_subject_sn/README.md) | United States SHAKEN CP |
| 1 | [e_sti_certificate_policies](ISSUES/e_sti_certificate_policies/README.md) | ATIS-1000080 |
| 1 | [e_sti_extension_unknown](ISSUES/e_sti_extension_unknown/README.md) | ATIS-1000080 |
| 1 | [e_sti_serial_number](ISSUES/e_sti_serial_number/README.md) | ATIS-1000080 |
| 1 | [e_sti_signature_algorithm](ISSUES/e_sti_signature_algorithm/README.md) | ATIS-1000080 |
| 1 | [e_sti_subject_cn](ISSUES/e_sti_subject_cn/README.md) | ATIS-1000080 |
| 1 | [w_cp_1_3_subject_email](ISSUES/w_cp_1_3_subject_email/README.md) | United States SHAKEN CP |
| 1 | [w_pki_subject_rdn_unknown](ISSUES/w_pki_subject_rdn_unknown/README.md) | SHAKEN PKI Best Practice |

#### CA Certificates

2 certificates were included in the corpus being tested\
2 certificates in the corpus were skipped because they were expired\
0 certificates in the corpus were skipped because they are not currently trusted\
0.00% of certificates contain one or more Error level issue\
100.00% of certificates contain one or more Warning level issue\
50.00% of certificates contain one or more Notice level issue\
50.00% of certificates are too old to be assessed against currently enforced expectations\
5451 days is the average remaining validity for the certificates in the corpus\
4928 days is the average initial validity for the certificates in the corpus\
0 certificates expire in the next 30 days

| Instances | Test Status | Source |
|-----------|-------------|--------|
| 1 | [e_cp1_3_ca_key_usage_crl_sign](ISSUES/e_cp1_3_ca_key_usage_crl_sign/README.md) | United States SHAKEN CP |
| 1 | [e_sti_basic_constraints](ISSUES/e_sti_basic_constraints/README.md) | ATIS-1000080 |
| 1 | [e_sti_ca_certificate_policies](ISSUES/e_sti_ca_certificate_policies/README.md) | ATIS-1000080 |
| 1 | [e_sti_ca_crl_distribution](ISSUES/e_sti_ca_crl_distribution/README.md) | ATIS-1000080 |
| 1 | [e_sti_ca_extension_unknown](ISSUES/e_sti_ca_extension_unknown/README.md) | ATIS-1000080 |
| 1 | [e_sti_ca_issuer_dn](ISSUES/e_sti_ca_issuer_dn/README.md) | ATIS-1000080 |
| 1 | [e_sti_ca_key_usage](ISSUES/e_sti_ca_key_usage/README.md) | ATIS-1000080 |
| 1 | [e_sti_ca_serial_number](ISSUES/e_sti_ca_serial_number/README.md) | ATIS-1000080 |
| 1 | [e_sti_ca_signature_algorithm](ISSUES/e_sti_ca_signature_algorithm/README.md) | ATIS-1000080 |
| 1 | [e_sti_ca_subject](ISSUES/e_sti_ca_subject/README.md) | ATIS-1000080 |
| 1 | [e_sti_ca_subject_cn](ISSUES/e_sti_ca_subject_cn/README.md) | ATIS-1000080 |
| 1 | [e_sti_ca_subject_key_identifier](ISSUES/e_sti_ca_subject_key_identifier/README.md) | ATIS-1000080 |
| 1 | [e_sti_ca_subject_public_key](ISSUES/e_sti_ca_subject_public_key/README.md) | ATIS-1000080 |
| 1 | [e_sti_ca_version](ISSUES/e_sti_ca_version/README.md) | ATIS-1000080 |
| 1 | [e_sti_root_certificate_policies](ISSUES/e_sti_root_certificate_policies/README.md) | ATIS-1000080 |
| 1 | [e_sti_root_extension_unknown](ISSUES/e_sti_root_extension_unknown/README.md) | ATIS-1000080 |
| 1 | [n_pki_ca_key_usage](ISSUES/n_pki_ca_key_usage/README.md) | SHAKEN PKI Best Practice |
| 1 | [n_sti_ca_certificate_policy_critical](ISSUES/n_sti_ca_certificate_policy_critical/README.md) | ATIS-1000080 |
| 1 | [w_pki_ca_subject_rdn_unknown](ISSUES/w_pki_ca_subject_rdn_unknown/README.md) | SHAKEN PKI Best Practice |

### Details

#### Leaf Certificates

| Created at | Subject | Problems | Link |
|------------|---------|----------|------|
| 12 Oct 22 17:35 UTC | SHAKEN Drop Inc 258K | true | [view](CERTS/2bf75c611bf6cdf4708719ecd41aee2869ddc8a72443993535c4344fa16dc9a2/README.md) |
| 27 Oct 22 18:39 UTC | SHAKEN Primo Dialler LLC 249K | true | [view](CERTS/8a9aa96133219cd164299b0f44144621987365aeb1c7538b079ea4e088f1195b/README.md) |
| 12 Sep 22 19:37 UTC | SHAKEN Bulk Solutions, LLC 644J | true | [view](CERTS/2b980444a4603ddf16248bee9dbdce112f593d4d5324443e641624a827af0cb2/README.md) |
| 12 Sep 22 19:37 UTC | SHAKEN Bulk Solutions, LLC 644J | true | [view](CERTS/2b980444a4603ddf16248bee9dbdce112f593d4d5324443e641624a827af0cb2/README.md) |
| 27 Oct 22 18:39 UTC | SHAKEN Primo Dialler LLC 249K | true | [view](CERTS/8a9aa96133219cd164299b0f44144621987365aeb1c7538b079ea4e088f1195b/README.md) |
| 12 Oct 22 07:27 UTC | SHAKEN NTC International, INC 016K | true | [view](CERTS/c4d0451ea21e6315e0dcdcb35469bc7d3a923a39f7c533b7d7b9ec4b37cd85e5/README.md) |
| 25 Oct 22 06:22 UTC | SHAKEN NTC International, INC 016K | true | [view](CERTS/2b6ef0741d09c6a0221ce9bfc7ca8edeb972333214bb5d2a970215876a32d4a0/README.md) |
| 14 Sep 22 05:12 UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | true | [view](CERTS/97e964a519370823455ea0f37ed9ca95c343aedd8d64f9215450d4de41f7af46/README.md) |
| 15 Sep 22 05:07 UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | true | [view](CERTS/540f2fd3efaf05a013774fb20d76c33e3136e10aa998e871b7ee6bff088ef571/README.md) |
| 16 Sep 22 05:02 UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | true | [view](CERTS/06c343355c0d9bb3a3d1067214e876b5f75ba6eb60fb85222509fdba1e901111/README.md) |
| 19 Sep 22 04:47 UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | true | [view](CERTS/95d4b2c125666234e336dabb1418eb7b7d9defa144f581923c96ba94b9011eca/README.md) |
| 20 Sep 22 04:42 UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | true | [view](CERTS/938113f93985847c1822cfb759ecfefb2b68bf68cfa086564102c40014f21ed4/README.md) |
| 03 Oct 22 03:37 UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | true | [view](CERTS/f5246093b0f190b0bc61a1ed5ec350f1b273c0a2bb228818af1459181b616b07/README.md) |
| 04 Oct 22 03:32 UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | true | [view](CERTS/d4d182512d7a0a08d36bdc57a093093c9d76c4ba70e6d05976634507efbbf16d/README.md) |
| 05 Oct 22 03:27 UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | true | [view](CERTS/900378b0848e406f48795d3e65e4c6147eb5877ae8d8d28b778ac42a428f58df/README.md) |
| 06 Oct 22 03:22 UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | true | [view](CERTS/10df08b76b13d13b8335da031a88902f217a899b6f4ed1e0af0a153be065dd3b/README.md) |
| 07 Oct 22 03:17 UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | true | [view](CERTS/130d0562b3a4c2d2338a50598c3f0efd47475a20deafda7258d6b6fbc661bd2f/README.md) |
| 11 Oct 22 02:57 UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | true | [view](CERTS/bcb003436c28b00c58f6ac3a9e93d32a3b26887acec22a7bfd682bf5f16c5f9e/README.md) |
| 12 Oct 22 02:52 UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | true | [view](CERTS/49c5a9cbfa67f6b5fa3f697038739d480efc0d206a21111854a5284122481316/README.md) |
| 13 Oct 22 02:47 UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | true | [view](CERTS/0c2acb5a833668f880d8170fa2db587b4ccb032422a2ec57210c746bc00f038a/README.md) |
| 14 Oct 22 02:42 UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | true | [view](CERTS/3ea47a78715c8ca2275f160ffd8a89af77cf79637198d723c9a6dcbda08f8e43/README.md) |
| 17 Oct 22 02:27 UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | true | [view](CERTS/cb26178771888c67934d3899aee9a9510941ce3acdb855cac2c87fb25ed54048/README.md) |
| 18 Oct 22 02:22 UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | true | [view](CERTS/15cc1f291e1f89c5855bda5986f1220e03c3534557eb8a50491b1da3106b8215/README.md) |
| 19 Oct 22 02:17 UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | true | [view](CERTS/c1bbb28b6924f72b66a383d5ce6b85c8f99cd742a4d56fd2aa82323f783b84ff/README.md) |
| 20 Oct 22 02:12 UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | true | [view](CERTS/7bb6c220b51e7a432c818dd42b99da02e9c4a66622cb9c95d0ef5c4f510553f5/README.md) |
| 21 Oct 22 02:07 UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | true | [view](CERTS/c7f300ba8278360ca314ce2fbae3025a00ef3987ee901ac7c940026c8c4cd226/README.md) |
| 24 Oct 22 01:52 UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | true | [view](CERTS/6b38e94aa75bbddcc98a5b6c62853dddd130b7e2dadffd014a723f8664ed6b16/README.md) |
| 25 Oct 22 01:47 UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | true | [view](CERTS/878ab560cdf2d47147ab378eb62e43f12052e41097ac724967824263891d4f1b/README.md) |
| 27 Oct 22 01:37 UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | true | [view](CERTS/8b0bb57dd9a98814d02347a4713fe7a0c3580c5341063a8cc60da4e8bd1dd2bf/README.md) |
| 12 Oct 22 17:35 UTC | SHAKEN Drop Inc 258K | true | [view](CERTS/2bf75c611bf6cdf4708719ecd41aee2869ddc8a72443993535c4344fa16dc9a2/README.md) |
| 11 Oct 22 17:24 UTC | SHAKEN Xchange Telecom LLC 325B | true | [view](CERTS/6bab691174d8e7b237a7fe1b00556840e2a5c28a1839f8e345dd9ba721ba23bb/README.md) |
| 26 Oct 22 22:53 UTC | SHAKEN Terra Nova Telecom 382G | true | [view](CERTS/ee166bd50d1486842178e96cd2da9da3f5638cf7cebdd56d951405427b97ee5d/README.md) |
| 19 Oct 22 16:35 UTC | SHAKEN 1stPoint Communications, LLC 463G | true | [view](CERTS/541acb18fc151d7d7192dbb6f9875853da85e3d30fc60e888f6206b097383707/README.md) |
| 21 Aug 22 19:06 UTC | SHAKEN Quality Voice & Data Inc. 548J | true | [view](CERTS/4631ec75151d90bbc876b172b0ebab5f88e11fa2cc851a1688975a08d48712e3/README.md) |
| 29 Sep 22 18:30 UTC | SHAKEN Quality Voice & Data Inc. 548J | true | [view](CERTS/0213907ac27ea330f259ea650d6c6ac19575228c28ce0ab49d029f09c26d3dab/README.md) |
| 03 Oct 22 18:10 UTC | SHAKEN Quality Voice & Data Inc. 548J | true | [view](CERTS/d8b296dab57106b2dd9654daeec135bf974ac6a610a317961fa7e03ad28a2442/README.md) |
| 04 Oct 22 18:05 UTC | SHAKEN Quality Voice & Data Inc. 548J | true | [view](CERTS/7c7df9bbe4fd2cf4819a97d078efecd02b047be763bb1ce7af981f0365358722/README.md) |
| 05 Oct 22 18:00 UTC | SHAKEN Quality Voice & Data Inc. 548J | true | [view](CERTS/cfc4289177789bc08fa2eb7b82a70f5220b0be4caeb388460b9aa3e9d51c0e5f/README.md) |
| 06 Oct 22 17:55 UTC | SHAKEN Quality Voice & Data Inc. 548J | true | [view](CERTS/16e4173562e9332228018b23e16cba6f85f147903cbae0ed9c81ed0c4c1d4895/README.md) |
| 09 Oct 22 17:40 UTC | SHAKEN Quality Voice & Data Inc. 548J | true | [view](CERTS/3889da782b38236bf6fde75142de0e20fef3414a531f182b8ff4f17621c4341b/README.md) |
| 10 Oct 22 17:35 UTC | SHAKEN Quality Voice & Data Inc. 548J | true | [view](CERTS/a6a89d754f8d558b1ce6d16624821f500fe7900626a2866a895fbcd9f76c7018/README.md) |
| 11 Oct 22 17:30 UTC | SHAKEN Quality Voice & Data Inc. 548J | true | [view](CERTS/997fd7cbfb80de86f60477feec6067c4f9f44837d8835185e07d7f1d54c9713b/README.md) |
| 12 Oct 22 17:25 UTC | SHAKEN Quality Voice & Data Inc. 548J | true | [view](CERTS/bd9b1d4aa62584bd9171f66fc1b4eb4fa6a9813ffb9f72acc164f5564abc0365/README.md) |
| 13 Oct 22 17:20 UTC | SHAKEN Quality Voice & Data Inc. 548J | true | [view](CERTS/8b32ad4d74e8b02255b012304d87442cbfe072cdc061dd1818195e0fbe25ff9f/README.md) |
| 17 Oct 22 17:00 UTC | SHAKEN Quality Voice & Data Inc. 548J | true | [view](CERTS/1365aeaa9b1d3d34d88ac9465ee709f88d9d1b34e38ad3388501a0214ff2e376/README.md) |
| 19 Oct 22 16:50 UTC | SHAKEN Quality Voice & Data Inc. 548J | true | [view](CERTS/e5889097187bc85d466d5edf5e339ba2c74ba917e4845691b25e544c83ba958a/README.md) |
| 24 Oct 22 16:25 UTC | SHAKEN Quality Voice & Data Inc. 548J | true | [view](CERTS/7308cc01648f3523bcc3834e2c41e9846c4004e96159c7d0a9a5b90008ab42c3/README.md) |
| 26 Oct 22 16:15 UTC | SHAKEN Quality Voice & Data Inc. 548J | true | [view](CERTS/5ccc340e70f7e4fd89a1d3b14455e6c559f76968a24bf27ac935e5595ceed7cf/README.md) |
| 27 Oct 22 16:10 UTC | SHAKEN Quality Voice & Data Inc. 548J | true | [view](CERTS/c98d31acc1d88aa6b8f23c05c1a3946642950394fc684d816d0196e708e21dbe/README.md) |
| 16 Aug 22 11:43 UTC | SHAKEN Lightspeed Voice 557F | true | [view](CERTS/fdc1362656ad1955711140f9f59ed3634d73811cf565b7e57b340e33a1786c45/README.md) |
| 06 Aug 21 14:13 UTC | SHAKEN Threshold Communications Inc 563J | true | [view](CERTS/3f3fa6ce5ff65b2ac841382ffd8a54cbbeb773a7913f0122f0baa35837af0e9d/README.md) |
| 26 Oct 22 22:26 UTC | SHAKEN Broadband Dynamics LLC 583j | true | [view](CERTS/aa4e5011e0e5ab6aac5180965888256227ecea3944cd35e13cca22398aa27f76/README.md) |
| 09 Sep 22 13:29 UTC | SHAKEN IPSBS Managed Services LLC 828J | true | [view](CERTS/605df407175181811a26f9cd97990c1d868c5b1f5b5d206874fb6176740c5d21/README.md) |
| 10 Oct 22 20:46 UTC | SHAKEN IPSBS Managed Services LLC 828J | true | [view](CERTS/b0861748e982e85ba5be363b4740bd726f3ca416bfab59ebbb9bd7145705f53d/README.md) |
| 22 Aug 22 18:28 UTC | SHAKEN Zray Technologies Corporation 862J | true | [view](CERTS/f957156b73a532f82676dd4245488058c58a8cfbab86d9036acc4fb9d5be77b6/README.md) |
| 12 Oct 22 20:47 UTC | SHAKEN Zray Technologies Corporation 862J | true | [view](CERTS/646c2010d83e94541b14b276b629bdc63d2288dcbfa7e28ce289159115327e5f/README.md) |
| 13 Oct 22 20:43 UTC | SHAKEN Zray Technologies Corporation 862J | true | [view](CERTS/f974722756a9ab31ee2e25efc88903b8b8bc0bddc5c96ed2f143fd18fcaac72b/README.md) |
| 14 Oct 22 20:38 UTC | SHAKEN Zray Technologies Corporation 862J | true | [view](CERTS/b1e6000eb4c7a6cc19b74ebf6707e3cff819d5e116ccdfffa330b90217d556bc/README.md) |
| 16 Oct 22 20:27 UTC | SHAKEN Zray Technologies Corporation 862J | true | [view](CERTS/9a592f926f5059b2f97043c202a7bb62e9b8d563b7c57c1a835caffdd82426ce/README.md) |
| 18 Oct 22 20:17 UTC | SHAKEN Zray Technologies Corporation 862J | true | [view](CERTS/a0fd285994c443dafc97d40dcd3d9ce20d4ae4492085720daf92178e1fb9b43c/README.md) |
| 01 Jul 22 00:00 UTC | SHAKEN IP Link Telecom Inc. 902J | true | [view](CERTS/e679b86b0a2afabd6deeb947c600361bc98e9b1a061240923e0e0a7999dfca6e/README.md) |
| 01 Oct 22 00:00 UTC | SHAKEN IP Link Telecom Inc. 902J | true | [view](CERTS/53d491627ef52772ffd75835f745e895f025bfc853523028b42c0bc1681e5f67/README.md) |
| 05 Oct 22 19:08 UTC | SHAKEN IP Link Telecom Inc. 902J | true | [view](CERTS/fa453a929f14e705532a1974b216dedb71ba1f10d6af0df07f1084bfbb2038db/README.md) |
| 03 Oct 22 18:44 UTC | SHAKEN InteractiveTel, LLC 920J | true | [view](CERTS/e38a973a0f060fbc1455f49b74ba81e0732500642c7d6302e3f06f6dcb306456/README.md) |
| 13 Oct 22 17:54 UTC | SHAKEN InteractiveTel, LLC 920J | true | [view](CERTS/c9d50511f042f19d1ca15f837a50a61c6a568ab4e0c9439acfa1902713e4ca42/README.md) |
| 19 Oct 22 17:24 UTC | SHAKEN InteractiveTel, LLC 920J | true | [view](CERTS/b30a67134d89241639dc14773d4f44842d52d0b7369ff195161d1b8e61490a09/README.md) |
| 26 Oct 22 16:49 UTC | SHAKEN InteractiveTel, LLC 920J | true | [view](CERTS/b5416e5eefd0dc2b74de073c207a83076d39bd7b88f9bfee3fe21b7418668c59/README.md) |
| 16 Sep 22 03:00 UTC | SHAKEN InteractiveTel, LLC 920J | true | [view](CERTS/7ec21783675a70d50bc3e55adc887f2d9e0e0f54b70cad27d1fd26cdcfe7ef91/README.md) |
| 11 Oct 22 17:08 UTC | SHAKEN ALD Telecom 780J | true | [view](CERTS/53a14081c994555770bb8c5f3d160f89cf427258c9598d569c388a74bde6ea8f/README.md) |
| 27 Oct 22 00:40 UTC | SHAKEN Airespring 996H | true | [view](CERTS/58a5f9cfa6a93167794a1fbd2d98777150cb8e4633c6617b57ac189800562a8d/README.md) |
| 27 Oct 22 00:36 UTC | SHAKEN Apeiron Systems 012J | true | [view](CERTS/b2f0624c685d38bf45c1a28928b9fdc8efcf0b3d19fa107dd1721f0f6edb1468/README.md) |
| 24 Oct 22 20:23 UTC | SHAKEN Arbeit 816J | true | [view](CERTS/377e182a223e6cc8d7e9ce697e7a3e829b1c6b16c299c26f6d1f1e33aa29524b/README.md) |
| 11 Oct 22 17:21 UTC | SHAKEN Asia Pacific Network 988J | true | [view](CERTS/0b191ba4d02eaa4b595b67a4d3e6f35a6d6c184e5b7e464d471cb904ea2d0638/README.md) |
| 11 Oct 22 17:20 UTC | SHAKEN Carrier One Inc. 705J | true | [view](CERTS/a7447339990a198aac3d84ed38d80706e16b7aac171e6d6bd1b28275fe7c337e/README.md) |
| 13 Oct 22 20:27 UTC | SHAKEN ConnectMeVoice 719J | true | [view](CERTS/a5edeeacfcec8ad6584f5a0b505978c4b72907a2e3a6540bb01350397f86814e/README.md) |
| 11 Oct 22 17:19 UTC | SHAKEN Connexum LLC 203K | true | [view](CERTS/1f823041acaec58b2c428be26eb8240228b6fd2b4d4d85436959aeda41bfe6a7/README.md) |
| 26 Oct 22 22:40 UTC | SHAKEN Consolidated Communications 5113 | true | [view](CERTS/b43c49cb5dd8b81bfbd6e21269a04a95923de71b44bcd4c240d1a4c49367851c/README.md) |
| 11 Oct 22 17:17 UTC | SHAKEN Current Calls, LLC 746J | true | [view](CERTS/52d6a93a1b72d2f2980699e759068dd9dbc8314c953e03613f18d9da1dcf156d/README.md) |
| 27 Oct 22 00:36 UTC | SHAKEN Fonative, Inc. 684J | true | [view](CERTS/04a47fb7fae7ff86f64fe7404197587d55581d659246c8c5f4c13183b2e57de2/README.md) |
| 27 Oct 22 00:38 UTC | SHAKEN IDT America, Corp 363A | true | [view](CERTS/63fe65f5789a08d669b2fb87b56dab400f17d42cf8b9bc5c1f61a81269837ec4/README.md) |
| 27 Oct 22 00:37 UTC | SHAKEN IPitomy 652J | true | [view](CERTS/353917762cb287ae83217d7b6a81b1af355bf021f7ba44d95f9b0bc6b6852744/README.md) |
| 11 Oct 22 17:19 UTC | SHAKEN Inventive Labs Corp 649J | true | [view](CERTS/a2f02cfef1eba726cf7dbd0f018a1119d40600aba568619f16b4c08b8d3a7c12/README.md) |
| 27 Oct 22 00:42 UTC | SHAKEN Matrix 9451 | true | [view](CERTS/73fa16a4e09bfed6129cdce6a5f141486b993826071af3c52c6faa642d196f3d/README.md) |
| 24 Oct 22 19:50 UTC | SHAKEN MagicJack 324E | true | [view](CERTS/828f14b0b96ddef337c05ed75551cea08ba7f76092ec9c814ab8342b46f220f5/README.md) |
| 26 Oct 22 22:10 UTC | SHAKEN Magna5, LLC 3849 | true | [view](CERTS/0fd5c659735095358e2a470c5fa788d14c4a55a34d107deab41eb4a3d48131ec/README.md) |
| 27 Oct 22 00:43 UTC | SHAKEN Matrix 7379 | true | [view](CERTS/cc0e632876194043a2bcdbb903fbc8a6faa5060d06c7a8a8e0554f1680bbf871/README.md) |
| 11 Oct 22 17:21 UTC | SHAKEN Momentum Telecom 1417 | true | [view](CERTS/063625a9c4c3a4e40f664e5ae334b8af15a18e04f853bfe25711df5faccdcbe0/README.md) |
| 27 Oct 22 00:38 UTC | SHAKEN NETRIO LLC 020K | true | [view](CERTS/335fa30fc4ef664740fd56ce52a0e020028671a2ca3d63c973712b205cf637fd/README.md) |
| 27 Oct 22 00:41 UTC | SHAKEN Nobelbiz, Inc. 596J | true | [view](CERTS/66824b20074f7c2afd910b494bcf0433436c8bf781dc42dd6545845e8c5c8be5/README.md) |
| 27 Oct 22 00:39 UTC | SHAKEN Noble Systems Communications LLC 187J | true | [view](CERTS/c2d9b2a2581455056d7b9ec15209368bfc8a177acda9459b2c618fc99616f4f2/README.md) |
| 11 Oct 22 17:22 UTC | SHAKEN OneStream Networks, LLC 630J | true | [view](CERTS/f18d0d387f4abfadaa336e2ff00c0f6b0509898b7d2d54feb99e1e0fb2042d3a/README.md) |
| 27 Oct 22 00:37 UTC | SHAKEN Phone.com, Inc. 633J | true | [view](CERTS/4dbe667d905260ebde18d12354e15cd118ed4922bf7a6236106c8ae94eacef59/README.md) |
| 27 Oct 22 00:43 UTC | SHAKEN PNG Telecommunications Inc 3395 | true | [view](CERTS/fb43066c14061a9603565ed3aaa4d7b1f052222d8f85da46f09883de3f594b67/README.md) |
| 08 Sep 22 14:21 UTC | SHAKEN Primo Dialler LLC 249K | true | [view](CERTS/0abd5fdc57586328458ef4cdc7420b6c95ff063049f94320c8a47c6c448e694f/README.md) |
| 11 Oct 22 17:24 UTC | SHAKEN Sangoma 777G | true | [view](CERTS/53d28ac1fa5253468c11b9e3baaa6ad5481e83a7ea2ee6d715594dc6d4561ad4/README.md) |
| 11 Oct 22 17:14 UTC | SHAKEN Technology Innovation Lab 599J | true | [view](CERTS/07d98b6eeb180548fa10e06aedbd69ce0816a1040344c91d25b8dcf29f68e7e6/README.md) |
| 11 Oct 22 16:48 UTC | SHAKEN TeleVoIPs 138K | true | [view](CERTS/c41b66127049dbae159f8d68ac714616b9e99640c407bcdc749f3d49037db487/README.md) |
| 27 Oct 22 00:35 UTC | SHAKEN Touchtone 683A | true | [view](CERTS/c53fcad60304c62293820901b4b9f365026fdf07f5df0065abe1afcfbf080db0/README.md) |
| 27 Oct 22 00:38 UTC | SHAKEN VoIP Innovations 597F | true | [view](CERTS/f9262d5cb74f09d76db538abcedbc1f93dcf7a59a67ab4769a1d1661e798df35/README.md) |
| 27 Oct 22 00:39 UTC | SHAKEN Telcentris Inc. dba Voxox 696J | true | [view](CERTS/91959b045a62dabf47cc92cc36bd234c1d6a4cd71487538fe9b88f8209ad28ef/README.md) |
| 24 Oct 22 21:11 UTC | SHAKEN Ytel Inc. 703J | true | [view](CERTS/3d6a7a2ff23b90fba1674f600a108b8a11a110f8bb1723df86627001f7367d8d/README.md) |
| 12 Oct 22 12:50 UTC | SHAKEN Lightspeed Voice 557F | true | [view](CERTS/ab19df868054cb3392aa295bff737bf919f8dc55c64a91247621375bad7fb7c0/README.md) |
| 23 Sep 22 01:10 UTC | SHAKEN Star2Star Communications, LLC 590J | true | [view](CERTS/b6c27ce63b22687fcd2f9f64ee9067dd3c19a4eb223f1aef3934f7ba95c54ba6/README.md) |
| 11 Oct 22 19:28 UTC | SHAKEN Star2Star Communications, LLC 590J | true | [view](CERTS/d3b8c22f303d29663b925cee65862ea3c3dec5e29f2d6872ece7c6e09b2b0f6a/README.md) |
| 23 Aug 22 06:30 UTC | SHAKEN  XCast Labs 689J | true | [view](CERTS/d6b76422bc88de5a2775b4a87e79be30eabf873db20b8ca0bcd1ecbabc51f77b/README.md) |
| 21 Sep 22 06:30 UTC | SHAKEN  XCast Labs 689J | true | [view](CERTS/b8e2172abf6f9919fde79f300bd848d91722301ae3287275a02f9fe3ae88ff2f/README.md) |
| 12 Oct 22 06:30 UTC | SHAKEN  XCast Labs 689J | true | [view](CERTS/fc5d5185fa770bfb1041d4346da797e8064b96241ac3239aa30fc0b63fa8828a/README.md) |
| 03 Jun 22 22:24 UTC | SHAKEN Ytel Inc. 703J | true | [view](CERTS/59dddd5c7a0daea466e319945be2964eab99e5a57c1ab92b1ddb2e2fa73db367/README.md) |
| 19 Oct 22 16:16 UTC | SHAKEN MagicJack 324E | true | [view](CERTS/a1f403a54a623739f646170ed600b0af39d7afe02b24af439b151e5da68e03b5/README.md) |
| 06 Oct 22 06:00 UTC | SHAKEN Convoso 758J | true | [view](CERTS/26b0c746370814ebf975e56be0f0be72d4a84c19defe8b8e5bc37080da1c44f2/README.md) |

#### CA Certificates

| Created at | Subject | Problems | Link |
|------------|---------|----------|------|
| 02 Sep 22 20:53 UTC | SHAKEN Sansay Intermediate CA US WEST 1 | true | [view](CERTS/4b1dfdba2b1e4bbffbf900a20f1f6f7befbef0008b963e4922a64469cb97d24b/README.md) |
| 28 Oct 20 14:41 UTC | SHAKEN Sansay Intermediate CA US WEST 1 | true | [view](CERTS/52041b27f345f20093d35f09ac6ca8e6123622aca8d71d3b358e2a94b674e2c9/README.md) |
| 13 Oct 21 20:46 UTC | SHAKEN Sansay Intermediate CA US WEST 1 | true | [view](CERTS/270561fad27744354bcd306ad2ca62858118b1ad7ce5098cebbe5549b11b03f0/README.md) |
| 21 Aug 20 01:22 UTC | SHAKEN Sansay Root CA US | true | [view](CERTS/8356d251b255f4ac3fd108bb79be4c02dcea2d3b13d138892bdb3a70cbe6a343/README.md) |


Generated: 31/10/2022 at 16:43:22