# STIR/SHAKEN CA Ecosystem Compliance

## Sansay

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 381 certificates were included in the corpus being tested
- 9 certificates in the corpus were skipped because they are duplicates
- 226 certificates in the corpus were skipped because they are expired
- 1 certificates in the corpus were skipped because they are not currently trusted
- 145 certificates being tested against the remaining rules
- 5.19 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 100.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 104 days is the average remaining validity for the certificates in the corpus
- 104 days is the average initial validity for the certificates in the corpus
- 108 certificates expire in the next 30 days
- 2.07 average number of unexpired certificates per OCN observed
- 70 unique OCNs observed in unexpired and valid certificate corpus

| Instances | Test | Source |
|-----------|------|--------|
| 92 | [e_atis_certificate_policies](ISSUES/e_atis_certificate_policies/README.md) | ATIS1000080 |
| 145 | [e_atis_subject_cn](ISSUES/e_atis_subject_cn/README.md) | ATIS1000080 |
| 145 | [e_us_cp_ambiguous_identifier](ISSUES/e_us_cp_ambiguous_identifier/README.md) | US_SHAKEN_CP |
| 145 | [e_us_cp_subject_sn](ISSUES/e_us_cp_subject_sn/README.md) | US_SHAKEN_CP |
| 81 | [w_cp_1_3_subject_email](ISSUES/w_cp_1_3_subject_email/README.md) | US_SHAKEN_CP |
| 145 | [w_shaken_subject_rdn_unknown](ISSUES/w_shaken_subject_rdn_unknown/README.md) | SHAKEN_PKI_BEST_PRACTICES |

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
- 5418 days is the average remaining validity for the certificates in the corpus
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
| 09&#160;Jan&#160;23&#160;23:57&#160;UTC | SHAKEN NTC International, INC 016K | 08&#160;Feb&#160;23&#160;23:57&#160;UTC | true | [view](CERTS/da2c521d460ad6de2288e70e92e1df8610a84d558ac6a7c758a38ab9c573826a/README.md) |
| 10&#160;Jan&#160;23&#160;07:55&#160;UTC | SHAKEN InteractiveTel, LLC 920J | 09&#160;Feb&#160;23&#160;07:55&#160;UTC | true | [view](CERTS/71a42f9b9dc4822bbac817e04708a502f2d54dfc74eee74ce4a055597b6aa086/README.md) |
| 10&#160;Jan&#160;23&#160;09:55&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 09&#160;Feb&#160;23&#160;09:55&#160;UTC | true | [view](CERTS/abc3e32bfbf0d003a8fca9dd4281bb4042717cb6c63bd87fb353d3b3b853f4ae/README.md) |
| 10&#160;Jan&#160;23&#160;11:19&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 09&#160;Feb&#160;23&#160;11:19&#160;UTC | true | [view](CERTS/c01056babc1f74136536f334fa3ed9b8c49a3ee1cb10fb04dcfa4ad1250a2364/README.md) |
| 10&#160;Jan&#160;23&#160;15:44&#160;UTC | SHAKEN IPSBS Managed Services LLC 828J | 09&#160;Feb&#160;23&#160;15:44&#160;UTC | true | [view](CERTS/cc87dfcad14b636266a6d2b8e597223548b3260b80280f0208b1b15ad4853cea/README.md) |
| 11&#160;Jan&#160;23&#160;09:50&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 10&#160;Feb&#160;23&#160;09:50&#160;UTC | true | [view](CERTS/1164d5f7020974f23a0a4e46ace2163ecdde49ae13e34f745de4314660d5badd/README.md) |
| 11&#160;Jan&#160;23&#160;11:14&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 10&#160;Feb&#160;23&#160;11:14&#160;UTC | true | [view](CERTS/4a5cef4518396875a172ca345ab2e5e524c67198f7b4b7672c4015dc4c51956b/README.md) |
| 12&#160;Jan&#160;23&#160;23:42&#160;UTC | SHAKEN NTC International, INC 016K | 11&#160;Feb&#160;23&#160;23:42&#160;UTC | true | [view](CERTS/4785c5cc7966400967976e33e89b3cb109855184e3f0a6215ca0e368b048f82e/README.md) |
| 13&#160;Jan&#160;23&#160;07:40&#160;UTC | SHAKEN InteractiveTel, LLC 920J | 12&#160;Feb&#160;23&#160;07:40&#160;UTC | true | [view](CERTS/f27bbd84a7fb66896e099d1c37b92ef025f61ce4275bd93576f1406c2043ed99/README.md) |
| 13&#160;Jan&#160;23&#160;09:40&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 12&#160;Feb&#160;23&#160;09:40&#160;UTC | true | [view](CERTS/c0241dcfc2983e2b6b4e57397c7d8c90d4d16da933e5e2df992459147fb5055a/README.md) |
| 13&#160;Jan&#160;23&#160;20:06&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 12&#160;Feb&#160;23&#160;20:06&#160;UTC | true | [view](CERTS/4b2da4a0d6b7e16ae744a697395e8ba946fc0070c85e75702f7f4577157cdecb/README.md) |
| 14&#160;Jan&#160;23&#160;09:35&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 13&#160;Feb&#160;23&#160;09:35&#160;UTC | true | [view](CERTS/91bcaa5ac3ae5a3aa7c8d36e50d7c83687ab7d871e2c036ed4fc98c66bcb8c41/README.md) |
| 17&#160;Jan&#160;23&#160;19:46&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 16&#160;Feb&#160;23&#160;19:46&#160;UTC | true | [view](CERTS/e248c1ac5fc708fe4894f8d9d0502c260d8bbf5d96fd36fea2bd645de2fd577f/README.md) |
| 17&#160;Jan&#160;23&#160;21:31&#160;UTC | SHAKEN Threshold Communications Inc 563J | 16&#160;Feb&#160;23&#160;21:31&#160;UTC | true | [view](CERTS/ab5dd3ebfe3882c8a685ed6bce46486fdcab519f3446bf1713e2f292d1d429eb/README.md) |
| 17&#160;Jan&#160;23&#160;21:32&#160;UTC | SHAKEN Broadband Dynamics LLC 583j | 16&#160;Feb&#160;23&#160;21:32&#160;UTC | true | [view](CERTS/17d556b64bac61cda21ac0b33e63c599ebb597e19f585ab7dec6de1ef5d9b821/README.md) |
| 17&#160;Jan&#160;23&#160;21:35&#160;UTC | SHAKEN Consolidated Communications 5113 | 16&#160;Feb&#160;23&#160;21:35&#160;UTC | true | [view](CERTS/566dc6b0c211c35f2388d1e9a8198440029c5b2845ff57fb3e3f0442bc1acd46/README.md) |
| 17&#160;Jan&#160;23&#160;21:37&#160;UTC | SHAKEN Touchtone 683A | 16&#160;Feb&#160;23&#160;21:37&#160;UTC | true | [view](CERTS/2a0a2422d9f52a5b45859c62be9ea582763229a948e46272a09244e5d7dcb242/README.md) |
| 17&#160;Jan&#160;23&#160;21:38&#160;UTC | SHAKEN Apeiron Systems 012J | 16&#160;Feb&#160;23&#160;21:38&#160;UTC | true | [view](CERTS/a29e9a25d7e2b1859aabe1f42c91d232c7710cd06d5d254db7460ac2e2252438/README.md) |
| 17&#160;Jan&#160;23&#160;21:40&#160;UTC | SHAKEN Fonative, Inc. 684J | 16&#160;Feb&#160;23&#160;21:40&#160;UTC | true | [view](CERTS/5ec73cc2cdfbaff0a331905ac7d205d114847589646e17d335645d35ae12cd06/README.md) |
| 17&#160;Jan&#160;23&#160;21:43&#160;UTC | SHAKEN IPitomy 652J | 16&#160;Feb&#160;23&#160;21:43&#160;UTC | true | [view](CERTS/58f7e612949f08c27727c991cc25b2870b7be88612da76d80c08cbf3b89f57f5/README.md) |
| 17&#160;Jan&#160;23&#160;21:45&#160;UTC | SHAKEN Phone.com, Inc. 633J | 16&#160;Feb&#160;23&#160;21:45&#160;UTC | true | [view](CERTS/28352d192e708f6e2fb757478e558de6b01f74378f611b70fe73d667562acf36/README.md) |
| 17&#160;Jan&#160;23&#160;21:46&#160;UTC | SHAKEN NETRIO LLC 020K | 16&#160;Feb&#160;23&#160;21:46&#160;UTC | true | [view](CERTS/bb8d33e19d4e815f312d4ea3e0529571ee624366442d45fc9286e20dd84348a7/README.md) |
| 17&#160;Jan&#160;23&#160;21:47&#160;UTC | SHAKEN VoIP Innovations 597F | 16&#160;Feb&#160;23&#160;21:47&#160;UTC | true | [view](CERTS/a3e9bc75383a8ecf49cd85550d2d2977dfc9007d64fd9573fab8e1cf77f77f53/README.md) |
| 17&#160;Jan&#160;23&#160;21:47&#160;UTC | SHAKEN IDT America, Corp 363A | 16&#160;Feb&#160;23&#160;21:47&#160;UTC | true | [view](CERTS/7093b5aac1ba935a3d520677a76cb73ea1e3b67329e02558868255429363b8c8/README.md) |
| 17&#160;Jan&#160;23&#160;21:48&#160;UTC | SHAKEN Noble Systems Communications LLC 187J | 16&#160;Feb&#160;23&#160;21:48&#160;UTC | true | [view](CERTS/71ed334e3d81a7c14be431b0fb1c26f0509ea3d62c1d2461fe15b7fb4155daff/README.md) |
| 17&#160;Jan&#160;23&#160;21:49&#160;UTC | SHAKEN Telcentris Inc. dba Voxox 696J | 16&#160;Feb&#160;23&#160;21:49&#160;UTC | true | [view](CERTS/c24c843568c4a9c93aed669df3cbf70dadf800ffe9b74a060d97d237e7c5e8ee/README.md) |
| 17&#160;Jan&#160;23&#160;21:50&#160;UTC | SHAKEN Airespring 996H | 16&#160;Feb&#160;23&#160;21:50&#160;UTC | true | [view](CERTS/4aecf5faaabd8b5d43c395dd587803ed61fa2b3c6258a2097d1f25a2c462fc76/README.md) |
| 17&#160;Jan&#160;23&#160;21:51&#160;UTC | SHAKEN Nobelbiz, Inc. 596J | 16&#160;Feb&#160;23&#160;21:51&#160;UTC | true | [view](CERTS/defb619b3a22f5cdce1fbdd841a2b9e6143d50285d980658a2b2f79dd7039729/README.md) |
| 17&#160;Jan&#160;23&#160;21:53&#160;UTC | SHAKEN Momentum Telecom 1417 | 16&#160;Feb&#160;23&#160;21:53&#160;UTC | true | [view](CERTS/cbf1bf713c958bebf0e470b36c43bb784f97f6a05171321aaa3b46d85ba48794/README.md) |
| 17&#160;Jan&#160;23&#160;21:53&#160;UTC | SHAKEN Momentum Telecom 9157 | 16&#160;Feb&#160;23&#160;21:53&#160;UTC | true | [view](CERTS/afb99cd0554a51045e825ce4352e8e50eb43d6a54820318ada11521067f72158/README.md) |
| 17&#160;Jan&#160;23&#160;21:54&#160;UTC | SHAKEN Terra Nova Telecom 382G | 16&#160;Feb&#160;23&#160;21:54&#160;UTC | true | [view](CERTS/c36042abf9d855a8b4e378f360c6a426f955fa18716e86b358918824a2d0f10e/README.md) |
| 17&#160;Jan&#160;23&#160;21:55&#160;UTC | SHAKEN Matrix 3058 | 16&#160;Feb&#160;23&#160;21:55&#160;UTC | true | [view](CERTS/fc8ce1925d30e9ae0d0e52586f972b3b37655d58188ff80e1b931bd27ffd63db/README.md) |
| 17&#160;Jan&#160;23&#160;21:56&#160;UTC | SHAKEN Matrix 9451 | 16&#160;Feb&#160;23&#160;21:56&#160;UTC | true | [view](CERTS/6745c656a94537f3bab030c4f9b329454dafa0bd3288d66f91271a0a3cbd1a8c/README.md) |
| 17&#160;Jan&#160;23&#160;21:56&#160;UTC | SHAKEN Matrix 7379 | 16&#160;Feb&#160;23&#160;21:56&#160;UTC | true | [view](CERTS/6453d26572c47e2357affef257b0fe7c34eb600c8b18486b78c5513d88416e6a/README.md) |
| 17&#160;Jan&#160;23&#160;21:57&#160;UTC | SHAKEN PNG Telecommunications Inc 3395 | 16&#160;Feb&#160;23&#160;21:57&#160;UTC | true | [view](CERTS/2605a8d7a1f83d8e410f2eb5df2d721392dc2dfedb65e2082c0add28b84dc473/README.md) |
| 17&#160;Jan&#160;23&#160;22:02&#160;UTC | SHAKEN Magna5, LLC 3849 | 16&#160;Feb&#160;23&#160;22:02&#160;UTC | true | [view](CERTS/536229a562d8ef73e679a82b46e1541217c5c4cfc2fee27bebc6612d4a002b3d/README.md) |
| 17&#160;Jan&#160;23&#160;22:05&#160;UTC | SHAKEN Magna5, LLC 8249 | 16&#160;Feb&#160;23&#160;22:05&#160;UTC | true | [view](CERTS/c32918368f2e258aa18657b34a1f5598590623301150f7d8a9308f99aeb3ccd9/README.md) |
| 18&#160;Jan&#160;23&#160;09:15&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 17&#160;Feb&#160;23&#160;09:15&#160;UTC | true | [view](CERTS/9bcbfb892244a557550152b43f2cf420a4b2318f58263e4ccf37795e661a36bb/README.md) |
| 18&#160;Jan&#160;23&#160;10:39&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 17&#160;Feb&#160;23&#160;10:39&#160;UTC | true | [view](CERTS/d141a4003020445a1fb43205b080c268a313698f014c673ee4125ae32635a8e4/README.md) |
| 18&#160;Jan&#160;23&#160;19:41&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 17&#160;Feb&#160;23&#160;19:41&#160;UTC | true | [view](CERTS/9cba9862280adb517058375366a03a1801eee2f1ba1c16beb7af94678876875e/README.md) |
| 18&#160;Jan&#160;23&#160;23:12&#160;UTC | SHAKEN NTC International, INC 016K | 17&#160;Feb&#160;23&#160;23:12&#160;UTC | true | [view](CERTS/4a82cbd851064adf7a25a5d19ad0025d617667e1986fb87c5152948dc4e050b5/README.md) |
| 19&#160;Jan&#160;23&#160;09:10&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 18&#160;Feb&#160;23&#160;09:10&#160;UTC | true | [view](CERTS/59ab916c6eb419914326f5a1a87ac198adfb81f67ac80747d864ce1135e854ff/README.md) |
| 19&#160;Jan&#160;23&#160;10:34&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 18&#160;Feb&#160;23&#160;10:34&#160;UTC | true | [view](CERTS/741aeca5b026ae5c6f5f074f43b52a7446362396d18ae093a7d2b28ab1c14c2e/README.md) |
| 19&#160;Jan&#160;23&#160;19:36&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 18&#160;Feb&#160;23&#160;19:36&#160;UTC | true | [view](CERTS/99cfc1b63c2f1febc340a11fd2ee50ce021f5da01114f250b181d2de98729a3c/README.md) |
| 19&#160;Jan&#160;23&#160;22:50&#160;UTC | SHAKEN Technology Innovation Lab 599J | 19&#160;Jan&#160;24&#160;22:50&#160;UTC | true | [view](CERTS/12acafcf01348d278955bb9276e7a4d22a65ccdc61a59d08100177711f21b430/README.md) |
| 19&#160;Jan&#160;23&#160;23:07&#160;UTC | SHAKEN NTC International, INC 016K | 18&#160;Feb&#160;23&#160;23:07&#160;UTC | true | [view](CERTS/5660a76ffc031d2c799d22f8c54824f5f5392dc5f2effbbc6290e7a5856d6e17/README.md) |
| 20&#160;Jan&#160;23&#160;07:05&#160;UTC | SHAKEN InteractiveTel, LLC 920J | 19&#160;Feb&#160;23&#160;07:05&#160;UTC | true | [view](CERTS/10d8e925057023b896346a3eeba52872e00d57ffc065266d3daa308f5dfe826c/README.md) |
| 20&#160;Jan&#160;23&#160;09:05&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 19&#160;Feb&#160;23&#160;09:05&#160;UTC | true | [view](CERTS/24afd223322a82d7f1cc5ffc1493ee903a8eda81a4cb3ab51f21b55cfb727032/README.md) |
| 20&#160;Jan&#160;23&#160;19:31&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 19&#160;Feb&#160;23&#160;19:31&#160;UTC | true | [view](CERTS/f188a68895e65a0b3a4424317978453f56709d30bef20a7214a8f47deb2c73be/README.md) |
| 20&#160;Jan&#160;23&#160;23:02&#160;UTC | SHAKEN NTC International, INC 016K | 19&#160;Feb&#160;23&#160;23:02&#160;UTC | true | [view](CERTS/36d50765eabdd867afd65d7d10e549979acc01b91e3646690a2cb6b5e862a4b3/README.md) |
| 21&#160;Jan&#160;23&#160;09:00&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 20&#160;Feb&#160;23&#160;09:00&#160;UTC | true | [view](CERTS/223fbe9b227137049ac02513c3f4270740630bdcdd9c82ecbe6b689cde4e4787/README.md) |
| 22&#160;Jan&#160;23&#160;19:21&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 21&#160;Feb&#160;23&#160;19:21&#160;UTC | true | [view](CERTS/44ff1c12add1f49ebda94f63e87b6f558aa5bfc334b586a567f6275ae0af501d/README.md) |
| 22&#160;Jan&#160;23&#160;22:52&#160;UTC | SHAKEN NTC International, INC 016K | 21&#160;Feb&#160;23&#160;22:52&#160;UTC | true | [view](CERTS/257191396787e93e5c8d756699cf6921a852685d21ffd6ec1ff9688cdb62e5b8/README.md) |
| 23&#160;Jan&#160;23&#160;06:49&#160;UTC | SHAKEN InteractiveTel, LLC 920J | 22&#160;Feb&#160;23&#160;06:49&#160;UTC | true | [view](CERTS/a5741804b44a01874768152580bb8c3f1e37e09079e543fd166b5ed6ec5dfc2d/README.md) |
| 23&#160;Jan&#160;23&#160;09:30&#160;UTC | SHAKEN Global Net Holdings Inc 306K | 22&#160;Feb&#160;23&#160;09:30&#160;UTC | true | [view](CERTS/ce5a08770726694f66c96ec566ccf7fcaa8731ac652b478f245469e5b0d28cb5/README.md) |
| 23&#160;Jan&#160;23&#160;19:16&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 22&#160;Feb&#160;23&#160;19:16&#160;UTC | true | [view](CERTS/c17e38aa1be632ff984e4c5276c5957b0e6d1ca80c1fef6973d5b6bb9f3c63bb/README.md) |
| 23&#160;Jan&#160;23&#160;21:55&#160;UTC | SHAKEN Swift Telco LLC 452K | 23&#160;Jan&#160;24&#160;21:55&#160;UTC | true | [view](CERTS/613861829aae7927f05dbd5a7b9f28ae8c4f995bb8ed115f95fc4be6644ccde1/README.md) |
| 24&#160;Jan&#160;23&#160;04:29&#160;UTC | SHAKEN Primo Dialler LLC 249K | 23&#160;Feb&#160;23&#160;04:29&#160;UTC | true | [view](CERTS/390981e1a1eaa5241867e99800532b455a6ac8564f5fc9d35d8692bd64819381/README.md) |
| 24&#160;Jan&#160;23&#160;08:45&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 23&#160;Feb&#160;23&#160;08:45&#160;UTC | true | [view](CERTS/cd1fa1b9e46b4c7b1e7ae124c187bcd83a7b17f5ada5ff41ff0abc70c46e9887/README.md) |
| 24&#160;Jan&#160;23&#160;10:09&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 23&#160;Feb&#160;23&#160;10:09&#160;UTC | true | [view](CERTS/88c8f0dcfec0870e189f8b57b77582f67d0fc42c215733b7d407c369689927ac/README.md) |
| 24&#160;Jan&#160;23&#160;19:11&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 23&#160;Feb&#160;23&#160;19:11&#160;UTC | true | [view](CERTS/a44f5ad970bbbb10fdd7879219ebcc9aa31b294974cf7994c1595995c03341a2/README.md) |
| 25&#160;Jan&#160;23&#160;06:40&#160;UTC | SHAKEN InteractiveTel, LLC 920J | 24&#160;Feb&#160;23&#160;06:40&#160;UTC | true | [view](CERTS/643c6eeacf5acbec8a6d719186c5c1b69a183dd23812a42aa056524f12270191/README.md) |
| 25&#160;Jan&#160;23&#160;08:40&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 24&#160;Feb&#160;23&#160;08:40&#160;UTC | true | [view](CERTS/5067423cca3c44619a2aa687760e3e22354dcc4dd755ad7c270524f0cca638a3/README.md) |
| 25&#160;Jan&#160;23&#160;10:04&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 24&#160;Feb&#160;23&#160;10:04&#160;UTC | true | [view](CERTS/33decdfcf406fb5d3edc8dea906e60fbcd1d625e78925946bd7fa62768c52e78/README.md) |
| 25&#160;Jan&#160;23&#160;19:06&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 24&#160;Feb&#160;23&#160;19:06&#160;UTC | true | [view](CERTS/3e7ede4a2a6594fb54d128047645ad8891f4b73dab62d07d2e82d252ff40d8c0/README.md) |
| 26&#160;Jan&#160;23&#160;06:34&#160;UTC | SHAKEN InteractiveTel, LLC 920J | 25&#160;Feb&#160;23&#160;06:34&#160;UTC | true | [view](CERTS/307580055f0b5279e89c8f720a0e934a4c1bafed3d42480de426e0803d83803c/README.md) |
| 26&#160;Jan&#160;23&#160;08:35&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 25&#160;Feb&#160;23&#160;08:35&#160;UTC | true | [view](CERTS/d3af02470bb8d59ac7af4d3abd298c4ffb250469cb120acfafa9e2180c4d923f/README.md) |
| 26&#160;Jan&#160;23&#160;09:59&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 25&#160;Feb&#160;23&#160;09:59&#160;UTC | true | [view](CERTS/4661c0270e57e4c3fe14a256aebd431966690c827a123a01fbf86eb7a037c859/README.md) |
| 26&#160;Jan&#160;23&#160;14:26&#160;UTC | SHAKEN IPSBS Managed Services LLC 828J | 25&#160;Feb&#160;23&#160;14:26&#160;UTC | true | [view](CERTS/85ce56a579fce3eea81d88c027b3fde867c106d7bed5d0b8cf06840039f3b50a/README.md) |
| 26&#160;Jan&#160;23&#160;19:01&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 25&#160;Feb&#160;23&#160;19:01&#160;UTC | true | [view](CERTS/9a39d734754b0df47685ff26f4e500e826a3f462abbdc2b831b4232451594ba2/README.md) |
| 26&#160;Jan&#160;23&#160;22:32&#160;UTC | SHAKEN NTC International, INC 016K | 25&#160;Feb&#160;23&#160;22:32&#160;UTC | true | [view](CERTS/ac12a922bd3a7f679f3f74cfadb2059b2b430f5366206006c29b35cff95dd826/README.md) |
| 27&#160;Jan&#160;23&#160;08:30&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 26&#160;Feb&#160;23&#160;08:30&#160;UTC | true | [view](CERTS/00b2562b12bade998cfbd8de4760d514209d4256b1e05742df94348e6bfd29df/README.md) |
| 27&#160;Jan&#160;23&#160;18:56&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 26&#160;Feb&#160;23&#160;18:56&#160;UTC | true | [view](CERTS/f6bfc75e6b2a38581b2bd5b00b945bbeee764fc655134a9aeb4519bc7080ce31/README.md) |
| 28&#160;Jan&#160;23&#160;08:25&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 27&#160;Feb&#160;23&#160;08:25&#160;UTC | true | [view](CERTS/b7a61d971cf1be4e415f3d322b8951e9f7b71ff03144d27371772fe84e3e5eb7/README.md) |
| 29&#160;Jan&#160;23&#160;18:46&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 28&#160;Feb&#160;23&#160;18:46&#160;UTC | true | [view](CERTS/d801be1d9b4e522a969f5cc03fcaddebea535c868055a45a231d833f2a55b290/README.md) |
| 29&#160;Jan&#160;23&#160;22:17&#160;UTC | SHAKEN NTC International, INC 016K | 28&#160;Feb&#160;23&#160;22:17&#160;UTC | true | [view](CERTS/5985a0adc174d99d06ac811d59f0a9f86cf1c869efc2a9f75bc66251fddd531b/README.md) |
| 30&#160;Jan&#160;23&#160;02:55&#160;UTC | SHAKEN BareTelecom 864J | 01&#160;Mar&#160;23&#160;02:55&#160;UTC | true | [view](CERTS/34be37c53279d1f3146f1cc4ad76fe323c88b92662282749a539628fe19e2941/README.md) |
| 30&#160;Jan&#160;23&#160;08:15&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 01&#160;Mar&#160;23&#160;08:15&#160;UTC | true | [view](CERTS/a78d79062dcd4315604c872e3607a1f22b38d7062307654d323071824e6aa1f2/README.md) |
| 30&#160;Jan&#160;23&#160;18:41&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 01&#160;Mar&#160;23&#160;18:41&#160;UTC | true | [view](CERTS/2ba31f0c8ddec231e5ebe2b67aadf6357459165833439331f07cc7ee198d0576/README.md) |
| 31&#160;Jan&#160;23&#160;06:10&#160;UTC | SHAKEN InteractiveTel, LLC 920J | 02&#160;Mar&#160;23&#160;06:10&#160;UTC | true | [view](CERTS/b156c90418c950fda197f3b8aa1cd095cc2067476c5bbd1041503ac6093f7ffe/README.md) |
| 31&#160;Jan&#160;23&#160;18:36&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 02&#160;Mar&#160;23&#160;18:36&#160;UTC | true | [view](CERTS/e258085bfa2688f629f3c8ae4ce63eb02f144da6256e9ea79436225c386461e1/README.md) |
| 31&#160;Jan&#160;23&#160;23:39&#160;UTC | SHAKEN Connexum LLC 203K | 01&#160;May&#160;23&#160;23:39&#160;UTC | true | [view](CERTS/d5409b0a1e255cfd6a3557acd45479a31641ee57e77ac96e1865499037f56c18/README.md) |
| 01&#160;Feb&#160;23&#160;06:04&#160;UTC | SHAKEN InteractiveTel, LLC 920J | 03&#160;Mar&#160;23&#160;06:04&#160;UTC | true | [view](CERTS/a9af2c2ee9f2ae2ba1247ec865f9df26f9ea65df9ae71a74da2d4d214478305a/README.md) |
| 01&#160;Feb&#160;23&#160;08:05&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 03&#160;Mar&#160;23&#160;08:05&#160;UTC | true | [view](CERTS/8867c2146558dc850c7735539c6c8c56680951a4209af5c6fe1be8de22519a86/README.md) |
| 01&#160;Feb&#160;23&#160;18:31&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 03&#160;Mar&#160;23&#160;18:31&#160;UTC | true | [view](CERTS/c124743827f158a4f85032f0bed5ce17fdc4b014daa2cb36241dd2e1aa210ae3/README.md) |
| 01&#160;Feb&#160;23&#160;20:15&#160;UTC | SHAKEN Zray Technologies Corporation 862J | 03&#160;Mar&#160;23&#160;20:15&#160;UTC | true | [view](CERTS/73181fe4791f578e7831027c45e472182465c2166256c4ee5f9fee069b366711/README.md) |
| 01&#160;Feb&#160;23&#160;22:02&#160;UTC | SHAKEN NTC International, INC 016K | 03&#160;Mar&#160;23&#160;22:02&#160;UTC | true | [view](CERTS/0527283ed0b4d18091b217766af9e880f3f0597140a3dc10172fc343ebd5ded9/README.md) |
| 02&#160;Feb&#160;23&#160;05:59&#160;UTC | SHAKEN InteractiveTel, LLC 920J | 04&#160;Mar&#160;23&#160;05:59&#160;UTC | true | [view](CERTS/3f2ad670fb733a1fd2cb90e3373ea8e1ea10056e38c51355faca1bcfbb60057c/README.md) |
| 02&#160;Feb&#160;23&#160;08:00&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 04&#160;Mar&#160;23&#160;08:00&#160;UTC | true | [view](CERTS/731ac0d52bdf232de2315ae9416517aa8cc81b3784d628e640e713a3e5322280/README.md) |
| 02&#160;Feb&#160;23&#160;13:51&#160;UTC | SHAKEN IPSBS Managed Services LLC 828J | 04&#160;Mar&#160;23&#160;13:51&#160;UTC | true | [view](CERTS/36f8fad217335efb3ae37a987f52c1949f01af8c399086765cb9db5ad4cfd7cc/README.md) |
| 02&#160;Feb&#160;23&#160;18:26&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 04&#160;Mar&#160;23&#160;18:26&#160;UTC | true | [view](CERTS/7ef3918b89dfee31240d4eea311b3f738be101b6a411e05473a4d9dd486e1741/README.md) |
| 02&#160;Feb&#160;23&#160;21:57&#160;UTC | SHAKEN NTC International, INC 016K | 04&#160;Mar&#160;23&#160;21:57&#160;UTC | true | [view](CERTS/de5e43e5a5c8e96da91e0904c790126213fdf2d58a8cd6fefcf2f2dc741d409c/README.md) |
| 03&#160;Feb&#160;23&#160;05:54&#160;UTC | SHAKEN InteractiveTel, LLC 920J | 05&#160;Mar&#160;23&#160;05:54&#160;UTC | true | [view](CERTS/6f5e9e66301757c9b51c7a46668c1bb4969c5692538ba304c810d694689e9672/README.md) |
| 03&#160;Feb&#160;23&#160;07:00&#160;UTC | SHAKEN Convoso 758J | 10&#160;Mar&#160;23&#160;07:00&#160;UTC | true | [view](CERTS/1b9234559702533a8cb320591bec3b154cb94885b9f991839248b04302073372/README.md) |
| 03&#160;Feb&#160;23&#160;18:21&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 05&#160;Mar&#160;23&#160;18:21&#160;UTC | true | [view](CERTS/7d55a687b5a1ea0a88aabff7c52113ff594fa303c34dae13ca321c921cacdfea/README.md) |
| 04&#160;Feb&#160;23&#160;18:16&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 06&#160;Mar&#160;23&#160;18:16&#160;UTC | true | [view](CERTS/2ae024c0827b10d2e742e1070f780bf21b3d80a5ea94ee35d295e0897b20e46b/README.md) |
| 05&#160;Feb&#160;23&#160;06:30&#160;UTC | SHAKEN  XCast Labs 689J | 08&#160;Mar&#160;23&#160;06:30&#160;UTC | true | [view](CERTS/0c2ce1f252ea7336f5905a17aa7920c76307b3af1a6e0bfcc0ed1f95669f39ad/README.md) |
| 05&#160;Feb&#160;23&#160;18:11&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 07&#160;Mar&#160;23&#160;18:11&#160;UTC | true | [view](CERTS/e876187fa338e08d92db02cab50d27bd177be48d17f8658146091868a795f8aa/README.md) |
| 05&#160;Feb&#160;23&#160;21:42&#160;UTC | SHAKEN NTC International, INC 016K | 07&#160;Mar&#160;23&#160;21:42&#160;UTC | true | [view](CERTS/5750a659208e66fcd0256e72671ba39cd81a8ab701f158eb02e7db760444f13f/README.md) |
| 06&#160;Feb&#160;23&#160;12:55&#160;UTC | SHAKEN Primo Dialler LLC 249K | 18&#160;Mar&#160;23&#160;12:55&#160;UTC | true | [view](CERTS/1f0beb9e2f7c9df37652e2b75862d6f82880a7b93ece5346c065a29c35bed2c1/README.md) |
| 06&#160;Feb&#160;23&#160;18:06&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 08&#160;Mar&#160;23&#160;18:06&#160;UTC | true | [view](CERTS/308916b04c3fbfa86fe339884c9f3030df9fe3dfd7380a3a764752a13f8bc95f/README.md) |
| 06&#160;Feb&#160;23&#160;19:50&#160;UTC | SHAKEN Zray Technologies Corporation 862J | 08&#160;Mar&#160;23&#160;19:50&#160;UTC | true | [view](CERTS/c1ec80576c4fcc8126f61fc3f37e660f6ab4501b0c887b2d00f6eee4f7171a6d/README.md) |
| 06&#160;Feb&#160;23&#160;21:37&#160;UTC | SHAKEN NTC International, INC 016K | 08&#160;Mar&#160;23&#160;21:37&#160;UTC | true | [view](CERTS/dd8537cc71060686a2f5a031b7ebc7d8c530c67b9774e4ccb671d101422a0f5d/README.md) |
| 07&#160;Feb&#160;23&#160;07:20&#160;UTC | SHAKEN 1stPoint Communications, LLC 463G | 09&#160;Mar&#160;23&#160;07:20&#160;UTC | true | [view](CERTS/28b46e2a9ce442f56ea4a4c82670940148817fe580de23fc967109e275d93b3e/README.md) |
| 07&#160;Feb&#160;23&#160;07:35&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 09&#160;Mar&#160;23&#160;07:35&#160;UTC | true | [view](CERTS/58723b73147a55d61f072388477f6332b3bf2d3051f0edd24962f044f01d2521/README.md) |
| 07&#160;Feb&#160;23&#160;08:59&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 09&#160;Mar&#160;23&#160;08:59&#160;UTC | true | [view](CERTS/6f4b1dbff2c22d15ac330c890d9175a7aa3bf3b29bd50085810c980144848afb/README.md) |
| 07&#160;Feb&#160;23&#160;13:26&#160;UTC | SHAKEN IPSBS Managed Services LLC 828J | 09&#160;Mar&#160;23&#160;13:26&#160;UTC | true | [view](CERTS/7e26581cd3852703a8d92fd6cdf90879123a888b3bd1d82c35371be9d7287457/README.md) |
| 07&#160;Feb&#160;23&#160;18:01&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 09&#160;Mar&#160;23&#160;18:01&#160;UTC | true | [view](CERTS/305562d2ca1cb410fd6a91e7a0ab5a20219c36d7f80b4bc2bf1989deeb7fcc2c/README.md) |
| 08&#160;Feb&#160;23&#160;07:15&#160;UTC | SHAKEN 1stPoint Communications, LLC 463G | 10&#160;Mar&#160;23&#160;07:15&#160;UTC | true | [view](CERTS/7706ed15fe5a91158c5776abcd1bd020d4f9b44edcda5cb82f039c72204c58e9/README.md) |
| 08&#160;Feb&#160;23&#160;07:30&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 10&#160;Mar&#160;23&#160;07:30&#160;UTC | true | [view](CERTS/df4ab806dca954cd69355670ccf10811ade8c88dc9935158180039d4107db80b/README.md) |
| 08&#160;Feb&#160;23&#160;13:21&#160;UTC | SHAKEN IPSBS Managed Services LLC 828J | 10&#160;Mar&#160;23&#160;13:21&#160;UTC | true | [view](CERTS/143816e2864b8674c703dd88814a927999034408e12d61f971930d792ac30b3f/README.md) |
| 08&#160;Feb&#160;23&#160;19:33&#160;UTC | SHAKEN ConvergeTel LLC 388K | 10&#160;Mar&#160;23&#160;19:33&#160;UTC | true | [view](CERTS/daf05ca0e5279f62e2982a0657e9541286f0b0de87f6edb444719c43602b82af/README.md) |

#### CA Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 21&#160;Aug&#160;20&#160;01:22&#160;UTC | SHAKEN Sansay Root CA US | 16&#160;Aug&#160;40&#160;01:22&#160;UTC | true | [view](CERTS/8356d251b255f4ac3fd108bb79be4c02dcea2d3b13d138892bdb3a70cbe6a343/README.md) |
| 02&#160;Sep&#160;22&#160;20:53&#160;UTC | SHAKEN Sansay Intermediate CA US WEST 1 | 31&#160;Aug&#160;29&#160;20:53&#160;UTC | true | [view](CERTS/4b1dfdba2b1e4bbffbf900a20f1f6f7befbef0008b963e4922a64469cb97d24b/README.md) |


Generated: 08 Feb 23 19:45 UTC