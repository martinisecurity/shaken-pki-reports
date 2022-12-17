# STIR/SHAKEN CA Ecosystem Compliance

## Sansay

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 248 certificates were included in the corpus being tested
- 6 certificates in the corpus were skipped because they are duplicates
- 91 certificates in the corpus were skipped because they are expired
- 1 certificates in the corpus were skipped because they are not currently trusted
- 150 certificates being tested against the remaining rules
- 5.61 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 100.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 95 days is the average remaining validity for the certificates in the corpus
- 96 days is the average initial validity for the certificates in the corpus
- 119 certificates expire in the next 30 days
- 2.27 average number of unexpired certificates per OCN observed
- 66 unique OCNs observed in unexpired and valid certificate corpus

| Instances | Test | Source |
|-----------|------|--------|
| 150 | [e_atis_certificate_policies](ISSUES/e_atis_certificate_policies/README.md) | ATIS1000080 |
| 150 | [e_atis_subject_cn](ISSUES/e_atis_subject_cn/README.md) | ATIS1000080 |
| 150 | [e_us_cp_ambiguous_identifier](ISSUES/e_us_cp_ambiguous_identifier/README.md) | US_SHAKEN_CP |
| 150 | [e_us_cp_subject_sn](ISSUES/e_us_cp_subject_sn/README.md) | US_SHAKEN_CP |
| 92 | [w_cp_1_3_subject_email](ISSUES/w_cp_1_3_subject_email/README.md) | US_SHAKEN_CP |
| 150 | [w_shaken_subject_rdn_unknown](ISSUES/w_shaken_subject_rdn_unknown/README.md) | SHAKEN_PKI_BEST_PRACTICES |

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
- 5436 days is the average remaining validity for the certificates in the corpus
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
| 19&#160;Oct&#160;22&#160;19:08&#160;UTC | SHAKEN Ytel Inc. 703J | 19&#160;Oct&#160;23&#160;19:08&#160;UTC | true | [view](CERTS/e0c7a355b91ad947dd48fc5a84523293447c45eac28d955f230fb212a73e34c3/README.md) |
| 24&#160;Oct&#160;22&#160;20:23&#160;UTC | SHAKEN Arbeit 816J | 24&#160;Oct&#160;23&#160;20:23&#160;UTC | true | [view](CERTS/377e182a223e6cc8d7e9ce697e7a3e829b1c6b16c299c26f6d1f1e33aa29524b/README.md) |
| 24&#160;Oct&#160;22&#160;21:11&#160;UTC | SHAKEN Ytel Inc. 703J | 24&#160;Oct&#160;23&#160;21:11&#160;UTC | true | [view](CERTS/3d6a7a2ff23b90fba1674f600a108b8a11a110f8bb1723df86627001f7367d8d/README.md) |
| 25&#160;Oct&#160;22&#160;20:17&#160;UTC | SHAKEN Talk IT Pro 321K | 25&#160;Oct&#160;23&#160;20:17&#160;UTC | true | [view](CERTS/e45dada701a589e681d12207ebf16985abf6d62cf429b6e03bdcf8c0f97c3bf2/README.md) |
| 26&#160;Oct&#160;22&#160;19:34&#160;UTC | SHAKEN Vinculum Communications, Inc 787J | 26&#160;Oct&#160;23&#160;19:34&#160;UTC | true | [view](CERTS/22936e87ea3c45af88f1e501b88c6c6db3c271bd6ef73ab33c5d68198f9d4d66/README.md) |
| 26&#160;Oct&#160;22&#160;19:41&#160;UTC | SHAKEN DLS Internet Services 815J | 26&#160;Oct&#160;23&#160;19:41&#160;UTC | true | [view](CERTS/7cd8319bedd12f040e8bd7b522d981aabcd24dc5aef74614a67fb6fdc9b9823b/README.md) |
| 26&#160;Oct&#160;22&#160;19:43&#160;UTC | SHAKEN Systemverse, LLC. 194K | 26&#160;Oct&#160;23&#160;19:43&#160;UTC | true | [view](CERTS/edbe74f809b9e0e1ebea447df8bdbfb272144f9c8c18df81e397a374df61c4cd/README.md) |
| 26&#160;Oct&#160;22&#160;19:44&#160;UTC | SHAKEN Rayfield Communications, Inc. 006K | 26&#160;Oct&#160;23&#160;19:44&#160;UTC | true | [view](CERTS/5032969f5932ac46a17b86c38dc72d666be454d1c3f11918edfa8385d9fc65e6/README.md) |
| 27&#160;Oct&#160;22&#160;20:11&#160;UTC | SHAKEN Mitel Cloud Services, Inc. 670J | 27&#160;Oct&#160;23&#160;20:11&#160;UTC | true | [view](CERTS/e45c92abcfe2fe6d0863200900b66e835aa98712f974efe3837e34d787f2ad5e/README.md) |
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
| 30&#160;Nov&#160;22&#160;22:42&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 30&#160;Dec&#160;22&#160;22:42&#160;UTC | true | [view](CERTS/cd9b1e842ea829e64f74b0c45657527590e2cf34473c6ac44146eff55dab012c/README.md) |
| 01&#160;Dec&#160;22&#160;03:17&#160;UTC | SHAKEN NTC International, INC 016K | 31&#160;Dec&#160;22&#160;03:17&#160;UTC | true | [view](CERTS/5f7bd941d77b9b4f3f9df7f1f1035757420be13aa9efadae73c825d3d06fd262/README.md) |
| 01&#160;Dec&#160;22&#160;11:14&#160;UTC | SHAKEN InteractiveTel, LLC 920J | 31&#160;Dec&#160;22&#160;11:14&#160;UTC | true | [view](CERTS/424c2a0b3387219e0385235fb34056e9e9f7d5e7a1e1e7efa371705db4463f1f/README.md) |
| 01&#160;Dec&#160;22&#160;13:00&#160;UTC | SHAKEN 1stPoint Communications, LLC 463G | 31&#160;Dec&#160;22&#160;13:00&#160;UTC | true | [view](CERTS/48b98ee3889f3c8e4c4c114281b366b919082b51b38bc078609b01af2c06eded/README.md) |
| 01&#160;Dec&#160;22&#160;13:15&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 31&#160;Dec&#160;22&#160;13:15&#160;UTC | true | [view](CERTS/27d516581e786c37e97b44143980b95f8e5a40f2863901ae6775621277b6d581/README.md) |
| 01&#160;Dec&#160;22&#160;22:37&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 31&#160;Dec&#160;22&#160;22:37&#160;UTC | true | [view](CERTS/19fdf5349eecd12295be792adbe8ddb7390214d276e34cde75897a6e591015a3/README.md) |
| 02&#160;Dec&#160;22&#160;03:12&#160;UTC | SHAKEN NTC International, INC 016K | 01&#160;Jan&#160;23&#160;03:12&#160;UTC | true | [view](CERTS/af1f7ad0bf05e4151ccdfa39217bbf0360d35cac9f1a6617ee7b9ddafca704cc/README.md) |
| 02&#160;Dec&#160;22&#160;11:10&#160;UTC | SHAKEN InteractiveTel, LLC 920J | 01&#160;Jan&#160;23&#160;11:10&#160;UTC | true | [view](CERTS/fa74168404912b89408df85dc58e6f00d39c8706773ca47b545768ada73ed383/README.md) |
| 02&#160;Dec&#160;22&#160;12:55&#160;UTC | SHAKEN 1stPoint Communications, LLC 463G | 01&#160;Jan&#160;23&#160;12:55&#160;UTC | true | [view](CERTS/f63b2b6591b43cc8bd4908a329e28ec7ee03860e068a2bb1b0445e028ecbb618/README.md) |
| 02&#160;Dec&#160;22&#160;13:10&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 01&#160;Jan&#160;23&#160;13:10&#160;UTC | true | [view](CERTS/df6084c1ec0684f3a6478c01a187c06089efaa8e36b66e5c32b3108717fd57bd/README.md) |
| 03&#160;Dec&#160;22&#160;03:07&#160;UTC | SHAKEN NTC International, INC 016K | 02&#160;Jan&#160;23&#160;03:07&#160;UTC | true | [view](CERTS/aed42fcd28e0bcd041d388ad0a0bdbc34d8fcb124e4ad6949b83ff090d3022ae/README.md) |
| 03&#160;Dec&#160;22&#160;16:18&#160;UTC | SHAKEN Televergence Solutions Inc 779J | 02&#160;Jan&#160;23&#160;16:18&#160;UTC | true | [view](CERTS/6559cf41e1abc7bf2dd92a993fee89a7ba6a85ab4588ac693ecaad8b45f1301b/README.md) |
| 04&#160;Dec&#160;22&#160;03:02&#160;UTC | SHAKEN NTC International, INC 016K | 03&#160;Jan&#160;23&#160;03:02&#160;UTC | true | [view](CERTS/e27604efd2c85ad99ef04bebcd3de3edd8be563f1acf05e3cbc3f65f92554ae7/README.md) |
| 04&#160;Dec&#160;22&#160;22:22&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 03&#160;Jan&#160;23&#160;22:22&#160;UTC | true | [view](CERTS/e9a574207a67ea953edf7412d098a0381bf671571fbb4b67a4dd9862e2720899/README.md) |
| 05&#160;Dec&#160;22&#160;02:57&#160;UTC | SHAKEN NTC International, INC 016K | 04&#160;Jan&#160;23&#160;02:57&#160;UTC | true | [view](CERTS/f1c5da1ee0a9178fe677b550b1af4538530b18c981452822c7789526a6c90e3f/README.md) |
| 05&#160;Dec&#160;22&#160;07:00&#160;UTC | SHAKEN Convoso 758J | 09&#160;Jan&#160;23&#160;07:00&#160;UTC | true | [view](CERTS/25407121a0f38b12a8ceda0b782b0512de9ef07a19a23d03cc09baa4989ed607/README.md) |
| 05&#160;Dec&#160;22&#160;12:40&#160;UTC | SHAKEN 1stPoint Communications, LLC 463G | 04&#160;Jan&#160;23&#160;12:40&#160;UTC | true | [view](CERTS/17b2120175007597c8a4de782271826764cdcc0d58e45e1e014d2033caa0a7ee/README.md) |
| 05&#160;Dec&#160;22&#160;12:55&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 04&#160;Jan&#160;23&#160;12:55&#160;UTC | true | [view](CERTS/f51dc2fc455bb27bcb756113c1e041b5c0c61b137609dbb61717cdae230934dd/README.md) |
| 05&#160;Dec&#160;22&#160;17:22&#160;UTC | SHAKEN Cyberlynk Network, LLC 086K | 04&#160;Jan&#160;23&#160;17:22&#160;UTC | true | [view](CERTS/0e6aa454cbce939bcf595664f67d567eb40ba223713b2dfbfa2da1658a4fa32d/README.md) |
| 05&#160;Dec&#160;22&#160;22:28&#160;UTC | SHAKEN Bulk Solutions, LLC 644J | 05&#160;Dec&#160;23&#160;22:28&#160;UTC | true | [view](CERTS/3cf0aa2a24845e3fe6b27605e223e8e0c73d6bd4f73279b8a1e5e16fd2feeb80/README.md) |
| 05&#160;Dec&#160;22&#160;23:21&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 04&#160;Jan&#160;23&#160;23:21&#160;UTC | true | [view](CERTS/144dbcccc4cdc7d74ab8b9906c815273ae740e6328ac9de339accb934c20005d/README.md) |
| 06&#160;Dec&#160;22&#160;02:52&#160;UTC | SHAKEN NTC International, INC 016K | 05&#160;Jan&#160;23&#160;02:52&#160;UTC | true | [view](CERTS/862329af5361e06a76e03c41951e923b7dcbd85adb43105761f35d9153e93aa8/README.md) |
| 06&#160;Dec&#160;22&#160;12:50&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 05&#160;Jan&#160;23&#160;12:50&#160;UTC | true | [view](CERTS/14ebea33889847de6d0787b5482d7f1fc567484fd508e896746f55a414097181/README.md) |
| 06&#160;Dec&#160;22&#160;18:39&#160;UTC | SHAKEN IPSBS Managed Services LLC 828J | 05&#160;Jan&#160;23&#160;18:39&#160;UTC | true | [view](CERTS/4f3b0ae63e9ca8a1d94d57287f6bed1e61c24b92f15ba0cf879c48ed3e3027e1/README.md) |
| 06&#160;Dec&#160;22&#160;23:16&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 05&#160;Jan&#160;23&#160;23:16&#160;UTC | true | [view](CERTS/300802d8087bd6e7cb1015c9865f35f1828119a0bd41be98c0b450b250e0d127/README.md) |
| 07&#160;Dec&#160;22&#160;02:47&#160;UTC | SHAKEN NTC International, INC 016K | 06&#160;Jan&#160;23&#160;02:47&#160;UTC | true | [view](CERTS/6fb49db6b185b693c7891b30b93f525e6883a25bcb8f3345369e6d3552c28d33/README.md) |
| 07&#160;Dec&#160;22&#160;12:30&#160;UTC | SHAKEN 1stPoint Communications, LLC 463G | 06&#160;Jan&#160;23&#160;12:30&#160;UTC | true | [view](CERTS/d705977eeefe2829819f9182adb234c177321eaa79a1887328f92a94b88ef879/README.md) |
| 07&#160;Dec&#160;22&#160;12:45&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 06&#160;Jan&#160;23&#160;12:45&#160;UTC | true | [view](CERTS/844b7dbcc3157584d5eaf843d51199a71a2d1ce3b2b474e6b889a05aa7a862e9/README.md) |
| 07&#160;Dec&#160;22&#160;13:32&#160;UTC | SHAKEN Threshold Communications Inc 563J | 06&#160;Jan&#160;23&#160;13:32&#160;UTC | true | [view](CERTS/1574436f964b80a25200da85ddd78eafed1b5b83e264fbe08ca0f87bff2ffede/README.md) |
| 07&#160;Dec&#160;22&#160;13:35&#160;UTC | SHAKEN Broadband Dynamics LLC 583j | 06&#160;Jan&#160;23&#160;13:35&#160;UTC | true | [view](CERTS/a48289433252222d5f6ce9a08b28eea8a045f12958617e1f830a744e07fe6df0/README.md) |
| 07&#160;Dec&#160;22&#160;13:38&#160;UTC | SHAKEN Consolidated Communications 5113 | 06&#160;Jan&#160;23&#160;13:38&#160;UTC | true | [view](CERTS/2e3fc61174f8cd6418114c916f46806c8e60697a49e33f779d8faaf377a24531/README.md) |
| 07&#160;Dec&#160;22&#160;13:41&#160;UTC | SHAKEN Touchtone 683A | 06&#160;Jan&#160;23&#160;13:41&#160;UTC | true | [view](CERTS/0a85df165a8c958d32e4bac0701e89b877ddf289e5283554fe4675ad0de1e451/README.md) |
| 07&#160;Dec&#160;22&#160;13:42&#160;UTC | SHAKEN Apeiron Systems 012J | 06&#160;Jan&#160;23&#160;13:42&#160;UTC | true | [view](CERTS/9bb9c28d9b397532ab2f55805058efca10702504172910ca4f2137b5de13c0dd/README.md) |
| 07&#160;Dec&#160;22&#160;13:43&#160;UTC | SHAKEN Fonative, Inc. 684J | 06&#160;Jan&#160;23&#160;13:43&#160;UTC | true | [view](CERTS/8baf31732c88f6c88de89d4fca46e314f9722b9b63f1b2a604cdabc9c2ab2a93/README.md) |
| 07&#160;Dec&#160;22&#160;13:43&#160;UTC | SHAKEN IPitomy 652J | 06&#160;Jan&#160;23&#160;13:43&#160;UTC | true | [view](CERTS/fa23816392d34740d0be85c940a769d447f806002cff4d0a712d090e010b6da1/README.md) |
| 07&#160;Dec&#160;22&#160;13:44&#160;UTC | SHAKEN Phone.com, Inc. 633J | 06&#160;Jan&#160;23&#160;13:44&#160;UTC | true | [view](CERTS/795152a916b566835cf210bc39ff2f4f4efa2243e86d83c3c56a9fa24ff523eb/README.md) |
| 07&#160;Dec&#160;22&#160;13:46&#160;UTC | SHAKEN NETRIO LLC 020K | 06&#160;Jan&#160;23&#160;13:46&#160;UTC | true | [view](CERTS/1b6e227b805347f6e6b68cf8f0cc4dfb0134b7c47e6c6eae4d7ebf450dbd5a55/README.md) |
| 07&#160;Dec&#160;22&#160;13:48&#160;UTC | SHAKEN VoIP Innovations 597F | 06&#160;Jan&#160;23&#160;13:48&#160;UTC | true | [view](CERTS/730a415d1afdf16dc2fe453f0e0e9695b9381854bf95c04e69e3a1ee75d6ef37/README.md) |
| 07&#160;Dec&#160;22&#160;13:48&#160;UTC | SHAKEN IDT America, Corp 363A | 06&#160;Jan&#160;23&#160;13:48&#160;UTC | true | [view](CERTS/cac6d944acbdcb026cfbad96e16b57a294def78ad2d0497fdef489c7ffd6f0c1/README.md) |
| 07&#160;Dec&#160;22&#160;13:51&#160;UTC | SHAKEN Noble Systems Communications LLC 187J | 06&#160;Jan&#160;23&#160;13:51&#160;UTC | true | [view](CERTS/b62490ac1eaf545f52354e4966be7b212ff5f61803bec71983f367c573af94ac/README.md) |
| 07&#160;Dec&#160;22&#160;13:53&#160;UTC | SHAKEN Telcentris Inc. dba Voxox 696J | 06&#160;Jan&#160;23&#160;13:53&#160;UTC | true | [view](CERTS/3ca8a945c3754fc603e7688b4054b39d48dc585017259d9dac27666b1387dff1/README.md) |
| 07&#160;Dec&#160;22&#160;13:55&#160;UTC | SHAKEN Airespring 996H | 06&#160;Jan&#160;23&#160;13:55&#160;UTC | true | [view](CERTS/99bc0e3a2ee69bae88035c3c3cae45f86ca4a81a9b47168eee93b6b5b531ce22/README.md) |
| 07&#160;Dec&#160;22&#160;13:57&#160;UTC | SHAKEN Nobelbiz, Inc. 596J | 06&#160;Jan&#160;23&#160;13:57&#160;UTC | true | [view](CERTS/b1c2c03743f90c520225c933f019a4a660f95f4b4951299c0931b3fab22e92a0/README.md) |
| 07&#160;Dec&#160;22&#160;14:00&#160;UTC | SHAKEN Momentum Telecom 1417 | 06&#160;Jan&#160;23&#160;14:00&#160;UTC | true | [view](CERTS/ea445c4926340df4a309a5def5c87f6c505f9b95c44a1163592cf64ad9a64a99/README.md) |
| 07&#160;Dec&#160;22&#160;14:00&#160;UTC | SHAKEN Momentum Telecom 9157 | 06&#160;Jan&#160;23&#160;14:00&#160;UTC | true | [view](CERTS/421cbc72131cecf9779ae86c4c25ced81f3f23258833881852ccc78250fe145e/README.md) |
| 07&#160;Dec&#160;22&#160;14:04&#160;UTC | SHAKEN Terra Nova Telecom 382G | 06&#160;Jan&#160;23&#160;14:04&#160;UTC | true | [view](CERTS/5a0592445a0a7b781a30d2d2d9b6358dbc2ba3679888ea51b6955d7e022403b2/README.md) |
| 07&#160;Dec&#160;22&#160;14:05&#160;UTC | SHAKEN Matrix 3058 | 06&#160;Jan&#160;23&#160;14:05&#160;UTC | true | [view](CERTS/5f676fa3dfbbffbfc5176245d48e49b3a4ecc449deb9ac998a54a4a961944423/README.md) |
| 07&#160;Dec&#160;22&#160;14:05&#160;UTC | SHAKEN Matrix 9451 | 06&#160;Jan&#160;23&#160;14:05&#160;UTC | true | [view](CERTS/775b09c8bc415614965a8872213bb6cee2b1d0130780c71f2be578c411f77c6b/README.md) |
| 07&#160;Dec&#160;22&#160;14:06&#160;UTC | SHAKEN Matrix 7379 | 06&#160;Jan&#160;23&#160;14:06&#160;UTC | true | [view](CERTS/9b009cbd408e5825cbf410414d505ebfd224274625a9f0c6899991b8bbc71fb4/README.md) |
| 07&#160;Dec&#160;22&#160;14:08&#160;UTC | SHAKEN PNG Telecommunications Inc 3395 | 06&#160;Jan&#160;23&#160;14:08&#160;UTC | true | [view](CERTS/85488e17097eb19f8823da83e03229abf7794901e76bbcb1490a7b5214b3e2d6/README.md) |
| 07&#160;Dec&#160;22&#160;14:16&#160;UTC | SHAKEN Magna5, LLC 3849 | 06&#160;Jan&#160;23&#160;14:16&#160;UTC | true | [view](CERTS/597dd9005c90e549779c8f696e1dbe792f76c760fb6d50d9b254fd1907165224/README.md) |
| 07&#160;Dec&#160;22&#160;14:18&#160;UTC | SHAKEN Magna5, LLC 8249 | 06&#160;Jan&#160;23&#160;14:18&#160;UTC | true | [view](CERTS/39d02173506632a74843809b7daf2198de267ec197ef2650b72546111025155a/README.md) |
| 07&#160;Dec&#160;22&#160;17:49&#160;UTC | SHAKEN Connexum LLC 203K | 05&#160;Feb&#160;23&#160;17:49&#160;UTC | true | [view](CERTS/f51c5a0a4f577f5322a53e5b0c369a450417a971a18214adad54978aa46ffefe/README.md) |
| 07&#160;Dec&#160;22&#160;18:34&#160;UTC | SHAKEN IPSBS Managed Services LLC 828J | 06&#160;Jan&#160;23&#160;18:34&#160;UTC | true | [view](CERTS/491a7178df7177f1b9706422f7fcdc4e03bc6a5c75e621e9d85fa762ebceb79a/README.md) |
| 07&#160;Dec&#160;22&#160;23:11&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 06&#160;Jan&#160;23&#160;23:11&#160;UTC | true | [view](CERTS/add3936a09829712ae735bbab25a6fa1bc119861a501defa9456195bf2d8f298/README.md) |
| 08&#160;Dec&#160;22&#160;02:42&#160;UTC | SHAKEN NTC International, INC 016K | 07&#160;Jan&#160;23&#160;02:42&#160;UTC | true | [view](CERTS/d4c8a2768014fda53ab2468692c46a3a9869ae517771920aba9e8d72d1924726/README.md) |
| 08&#160;Dec&#160;22&#160;10:40&#160;UTC | SHAKEN InteractiveTel, LLC 920J | 07&#160;Jan&#160;23&#160;10:40&#160;UTC | true | [view](CERTS/7a444e306ece73f178cc176196b14c32f64dc3759a09774c66cd00262405cac0/README.md) |
| 08&#160;Dec&#160;22&#160;12:25&#160;UTC | SHAKEN 1stPoint Communications, LLC 463G | 07&#160;Jan&#160;23&#160;12:25&#160;UTC | true | [view](CERTS/d25a0d5ad1db30b53a6d33de975655d9b8b2507401c1571c41f5cc4107e71a24/README.md) |
| 08&#160;Dec&#160;22&#160;12:40&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 07&#160;Jan&#160;23&#160;12:40&#160;UTC | true | [view](CERTS/7898310169108e245eb9406ffd2125eb0bf7a15772d70de651bd480ba07f2ea0/README.md) |
| 08&#160;Dec&#160;22&#160;23:06&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 07&#160;Jan&#160;23&#160;23:06&#160;UTC | true | [view](CERTS/d944dd27782504db666d3ca7d3a6774bcfce9390478a1c64ded9cd5d7ab5ded1/README.md) |
| 09&#160;Dec&#160;22&#160;02:37&#160;UTC | SHAKEN NTC International, INC 016K | 08&#160;Jan&#160;23&#160;02:37&#160;UTC | true | [view](CERTS/b527c9d0685bbeaf74cdd98f1344e4f6d890a5c6fc10cc8190899d8b87275ddd/README.md) |
| 09&#160;Dec&#160;22&#160;06:30&#160;UTC | SHAKEN  XCast Labs 689J | 09&#160;Jan&#160;23&#160;06:30&#160;UTC | true | [view](CERTS/5abc6e4997d766792f3809dc868f308392b1e338341ab18b72ffba7dd8d4ec5a/README.md) |
| 09&#160;Dec&#160;22&#160;12:35&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 08&#160;Jan&#160;23&#160;12:35&#160;UTC | true | [view](CERTS/3ddcdb62a9f2e7ec21c1dac733a1df38bf8a70cf55ca7e4a367428d63568ee77/README.md) |
| 09&#160;Dec&#160;22&#160;21:02&#160;UTC | SHAKEN Primo Dialler LLC 249K | 08&#160;Jan&#160;23&#160;21:02&#160;UTC | true | [view](CERTS/358d163c11a4c212e4f470e8823d9d9bf6ee6ec6a565957df41c4f3d4eb52a34/README.md) |
| 10&#160;Dec&#160;22&#160;02:11&#160;UTC | SHAKEN Drop Inc 258K | 10&#160;Dec&#160;23&#160;02:11&#160;UTC | true | [view](CERTS/fc457741017b89b9126882710d8fb44883d7603f79cec0a1989eaa2b08034ee5/README.md) |
| 11&#160;Dec&#160;22&#160;02:27&#160;UTC | SHAKEN NTC International, INC 016K | 10&#160;Jan&#160;23&#160;02:27&#160;UTC | true | [view](CERTS/a02719504c53126f8c169e516160f2802fbd402e966cf2f382f4e772023240ed/README.md) |
| 11&#160;Dec&#160;22&#160;22:51&#160;UTC | SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J | 10&#160;Jan&#160;23&#160;22:51&#160;UTC | true | [view](CERTS/58657e0f4cfaa1ea955b1186890bef2bb855aab4a123d83e2a1d6458fcab3a3e/README.md) |
| 12&#160;Dec&#160;22&#160;02:22&#160;UTC | SHAKEN NTC International, INC 016K | 11&#160;Jan&#160;23&#160;02:22&#160;UTC | true | [view](CERTS/304a8efb19acd3e64449987a2897bcd227da9d20321567b4aabb1ea4e43523e3/README.md) |
| 12&#160;Dec&#160;22&#160;10:19&#160;UTC | SHAKEN InteractiveTel, LLC 920J | 11&#160;Jan&#160;23&#160;10:19&#160;UTC | true | [view](CERTS/efdada9132eda0bbf96fed945eaea0cc73ef940549c1adfcf2378fe99f4af9bf/README.md) |
| 12&#160;Dec&#160;22&#160;12:05&#160;UTC | SHAKEN 1stPoint Communications, LLC 463G | 11&#160;Jan&#160;23&#160;12:05&#160;UTC | true | [view](CERTS/d498c87e4840c160ceaebbf72e279b855921f27b2187bf74624301ed506716c4/README.md) |
| 12&#160;Dec&#160;22&#160;12:20&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 11&#160;Jan&#160;23&#160;12:20&#160;UTC | true | [view](CERTS/fd183756ccba5d162fa7398f0ba8b57757aa606b0477231a741722f1959f22b8/README.md) |
| 12&#160;Dec&#160;22&#160;18:09&#160;UTC | SHAKEN IPSBS Managed Services LLC 828J | 11&#160;Jan&#160;23&#160;18:09&#160;UTC | true | [view](CERTS/ca8ccf4bdd084d83eaea04f52a28c6fe2dc932ed4cca933e42bd22dd26d89567/README.md) |
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
| 15&#160;Dec&#160;22&#160;12:05&#160;UTC | SHAKEN Quality Voice & Data Inc. 548J | 14&#160;Jan&#160;23&#160;12:05&#160;UTC | true | [view](CERTS/09914e74c5ad04841f48ce52169cdc4a6e3221a7dc7b38bcecc8f7394d745df0/README.md) |

#### CA Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 21&#160;Aug&#160;20&#160;01:22&#160;UTC | SHAKEN Sansay Root CA US | 16&#160;Aug&#160;40&#160;01:22&#160;UTC | true | [view](CERTS/8356d251b255f4ac3fd108bb79be4c02dcea2d3b13d138892bdb3a70cbe6a343/README.md) |
| 02&#160;Sep&#160;22&#160;20:53&#160;UTC | SHAKEN Sansay Intermediate CA US WEST 1 | 31&#160;Aug&#160;29&#160;20:53&#160;UTC | true | [view](CERTS/4b1dfdba2b1e4bbffbf900a20f1f6f7befbef0008b963e4922a64469cb97d24b/README.md) |


Generated: 17 Dec 22 12:22 UTC