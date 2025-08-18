# STIR/SHAKEN CA Ecosystem Compliance

## Telonium

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 125 certificates were included in the corpus being tested
- 22 certificates in the corpus were skipped because they are duplicates
- 31 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 72 certificates being tested against the remaining rules
- 1.00 issues on average found in unexpired, trusted, and non-compliant certificates
- 1.39% of certificates contain one or more Error level issue
- 0.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 427 days is the average remaining validity for the certificates in the corpus
- 431 days is the average initial validity for the certificates in the corpus
- 4 certificates expire in the next 30 days
- 1.01 average number of unexpired certificates per OCN observed
- 71 unique OCNs observed in unexpired and valid certificate corpus

| Instances | Test | Source |
|-----------|------|--------|
| 1 | [e_atis_ext_crl_distribution_struct](ISSUES/e_atis_ext_crl_distribution_struct/README.md) | ATIS1000080 |

#### CA Certificates

- 5 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 0 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 5 certificates being tested against the remaining rules
- 4.00 issues on average found in unexpired, trusted, and non-compliant certificates
- 60.00% of certificates contain one or more Error level issue
- 0.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 4195 days is the average remaining validity for the certificates in the corpus
- 4309 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days

| Instances | Test | Source |
|-----------|------|--------|
| 1 | [e_atis_ext_certificate_policies_ca](ISSUES/e_atis_ext_certificate_policies_ca/README.md) | ATIS1000080 |
| 1 | [e_atis_ext_crl_distribution_ca](ISSUES/e_atis_ext_crl_distribution_ca/README.md) | ATIS1000080 |
| 2 | [e_atis_ext_key_usage](ISSUES/e_atis_ext_key_usage/README.md) | ATIS1000080 |
| 1 | [e_atis_subject_c_iso_ca](ISSUES/e_atis_subject_c_iso_ca/README.md) | ATIS1000080 |
| 3 | [e_atis_subject_cn_ca](ISSUES/e_atis_subject_cn_ca/README.md) | ATIS1000080 |
| 2 | [e_atis_subject_cn_root](ISSUES/e_atis_subject_cn_root/README.md) | ATIS1000080 |
| 1 | [e_atis_subject_dn_ca](ISSUES/e_atis_subject_dn_ca/README.md) | ATIS1000080 |
| 1 | [e_shaken_certificate_policies_id_ca](ISSUES/e_shaken_certificate_policies_id_ca/README.md) | US_SHAKEN_CP |

### Details

#### Leaf Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 27&#160;Feb&#160;24&#160;16:46&#160;UTC | SHAKEN 963J | 27&#160;Feb&#160;26&#160;16:47&#160;UTC | true | [view](CERTS/ceef9f6d88a1efc4c056ecaf0f5d42c5f106fb5fc895c8ccc1f7ea97b6ac1093/README.md) |
| 20&#160;Jun&#160;24&#160;19:05&#160;UTC | SHAKEN 633K | 27&#160;Dec&#160;25&#160;16:31&#160;UTC | false | [view](CERTS/5c0d32d70b614b08cec3502a7ef34663daa44eb814894c021def70d56c113d62/README.md) |
| 12&#160;Aug&#160;24&#160;14:42&#160;UTC | SHAKEN 709K | 25&#160;Aug&#160;25&#160;20:59&#160;UTC | false | [view](CERTS/284ede2987dda2f3409b88ce30ee10e32787908a6cfbfb5e8abbaf305b33c7a4/README.md) |
| 12&#160;Aug&#160;24&#160;16:23&#160;UTC | SHAKEN 715K | 11&#160;Sep&#160;26&#160;18:30&#160;UTC | false | [view](CERTS/396ffc900acd18b268a476122bfaae0835afda50c8648220980093051eb2bbfb/README.md) |
| 20&#160;Aug&#160;24&#160;17:07&#160;UTC | SHAKEN 974K | 20&#160;Aug&#160;25&#160;17:08&#160;UTC | false | [view](CERTS/cc64adfd746943da601c5f1fe4572b2124f6ac35a134d79a0e1df3981d31d518/README.md) |
| 27&#160;Sep&#160;24&#160;18:32&#160;UTC | SHAKEN 025L | 27&#160;Sep&#160;25&#160;18:33&#160;UTC | false | [view](CERTS/8f5f429f932d898d06db9e4c2cdce15ddbef77a6b3a460c26167d161198632e4/README.md) |
| 04&#160;Oct&#160;24&#160;20:26&#160;UTC | SHAKEN 721K | 04&#160;Oct&#160;25&#160;20:27&#160;UTC | false | [view](CERTS/fcea8461820fd32d651d21bd56cfe89d49c4baf3838c5980defe495954371da8/README.md) |
| 12&#160;Oct&#160;24&#160;00:08&#160;UTC | SHAKEN 016L | 12&#160;Oct&#160;25&#160;00:09&#160;UTC | false | [view](CERTS/89a8ebcd6931cc9cbef9496b87e12100c865d92738d91eef202b9c5dfbd37389/README.md) |
| 24&#160;Oct&#160;24&#160;19:26&#160;UTC | SHAKEN 161K | 24&#160;Oct&#160;25&#160;19:27&#160;UTC | false | [view](CERTS/400f3284e4959f3f5340073e43c43d521b6b305cf2b1d3774fc182bcd1870653/README.md) |
| 13&#160;Nov&#160;24&#160;02:00&#160;UTC | SHAKEN 018L | 13&#160;Nov&#160;25&#160;02:01&#160;UTC | false | [view](CERTS/b8c98d15da79b182ecf38a430f0bbbfe9b2972f1671c9201163ea6ade4780643/README.md) |
| 04&#160;Dec&#160;24&#160;16:57&#160;UTC | SHAKEN 062L | 04&#160;Dec&#160;25&#160;16:58&#160;UTC | false | [view](CERTS/f816c771e07ff0acc479deac83178d94808ebe44d64e081e276ea599ccd25cac/README.md) |
| 17&#160;Dec&#160;24&#160;15:21&#160;UTC | SHAKEN 979K | 17&#160;Dec&#160;25&#160;15:22&#160;UTC | false | [view](CERTS/1a4bbce4d44b3de984f058ce18f11566508d6121e05d9956cee18421ce003fba/README.md) |
| 27&#160;Dec&#160;24&#160;18:01&#160;UTC | SHAKEN 033L | 27&#160;Dec&#160;25&#160;18:02&#160;UTC | false | [view](CERTS/fd08f0121922847730d6c92d0c49e66da96a6276dc9fe70235fecd728c70c661/README.md) |
| 07&#160;Jan&#160;25&#160;18:55&#160;UTC | SHAKEN 830K | 07&#160;Feb&#160;27&#160;17:05&#160;UTC | false | [view](CERTS/a2342336f01f23a09d73c1dfd195bbbc1c5889e50bde8b88dde4f3fedaebf4aa/README.md) |
| 07&#160;Jan&#160;25&#160;19:07&#160;UTC | SHAKEN 503K | 07&#160;Jan&#160;26&#160;19:08&#160;UTC | false | [view](CERTS/4cedbd042ae0557033b2310d79f69ef7d4dc5e53f06e0048cfa80c1f9d8fd3ae/README.md) |
| 07&#160;Jan&#160;25&#160;19:37&#160;UTC | SHAKEN 785K | 09&#160;Jan&#160;27&#160;14:51&#160;UTC | false | [view](CERTS/17b3fae323f14b2b2a097443ebf81c70f6ac4576456487bb78b90a70178791dd/README.md) |
| 08&#160;Jan&#160;25&#160;20:25&#160;UTC | SHAKEN 520F | 08&#160;Jan&#160;26&#160;20:26&#160;UTC | false | [view](CERTS/057554502eb6a8d0e6dd41a4a51c1e149459d48e646dafa43a961bc4e1eeae24/README.md) |
| 21&#160;Jan&#160;25&#160;18:35&#160;UTC | SHAKEN 850J | 21&#160;Jan&#160;26&#160;18:36&#160;UTC | false | [view](CERTS/2b71a54aac8322484e0654761b1eb04c63d9f35658a9142ae505981263090eb7/README.md) |
| 21&#160;Jan&#160;25&#160;20:57&#160;UTC | SHAKEN 772J | 04&#160;Apr&#160;26&#160;19:31&#160;UTC | false | [view](CERTS/701b836f7ac8a52b75b394f23dedf984b6b7ea747f25dcfcf439fbfd3f794b97/README.md) |
| 05&#160;Feb&#160;25&#160;15:42&#160;UTC | SHAKEN 847K | 04&#160;Mar&#160;26&#160;16:13&#160;UTC | false | [view](CERTS/d93fb236210dd6292ca8ac07fce7b55bb857ed234d49cbc2c51e5d20315f23ff/README.md) |
| 06&#160;Feb&#160;25&#160;18:23&#160;UTC | SHAKEN 122L | 06&#160;Feb&#160;26&#160;18:24&#160;UTC | false | [view](CERTS/f3117c32bc70624628e299c1963c1ce9feada35bb4e3f1749e0d8317753f6825/README.md) |
| 14&#160;Feb&#160;25&#160;22:24&#160;UTC | SHAKEN 869K | 12&#160;Mar&#160;26&#160;18:23&#160;UTC | false | [view](CERTS/d77e8f5e5f3da13e0c65244159f6eaf4e6666f23942bdb8dd0de38d383540b01/README.md) |
| 19&#160;Feb&#160;25&#160;16:02&#160;UTC | SHAKEN 034L | 07&#160;Oct&#160;25&#160;17:13&#160;UTC | false | [view](CERTS/10d385f563843e1536adb4a1426ac3c1ec323359c2d9ec2475f10ec9901420ca/README.md) |
| 19&#160;Feb&#160;25&#160;16:19&#160;UTC | SHAKEN 037L | 08&#160;Oct&#160;25&#160;20:56&#160;UTC | false | [view](CERTS/2118f6710c44b0249c1cc49d3f81179265f02300455e68ef227215036dd87d91/README.md) |
| 20&#160;Feb&#160;25&#160;16:55&#160;UTC | SHAKEN 014L | 27&#160;Sep&#160;25&#160;20:15&#160;UTC | false | [view](CERTS/62d31253cf6851ade341783eb31b8d6e901e11e100f81795935d22a43a84fe95/README.md) |
| 26&#160;Feb&#160;25&#160;16:30&#160;UTC | SHAKEN 854K | 08&#160;Mar&#160;26&#160;18:12&#160;UTC | false | [view](CERTS/f730cae1df3042f247b301603ef2952b82fd30367107bb654bd4a63e3db0a80b/README.md) |
| 06&#160;Mar&#160;25&#160;14:29&#160;UTC | SHAKEN 141L | 06&#160;Mar&#160;26&#160;14:30&#160;UTC | false | [view](CERTS/65821fb6778ef9d7f382a3da0c2b36a5601ae6757a9a4183d5a31345cbcf013a/README.md) |
| 13&#160;Mar&#160;25&#160;14:39&#160;UTC | SHAKEN 086L | 13&#160;Mar&#160;26&#160;14:40&#160;UTC | false | [view](CERTS/95c0032daa6ee1d27993df8eafab6bc987c645154f78dffe4e7a8c87ab4f2890/README.md) |
| 21&#160;Mar&#160;25&#160;15:08&#160;UTC | SHAKEN 976J | 02&#160;Apr&#160;26&#160;18:05&#160;UTC | false | [view](CERTS/b3c7ffa68bb46d7a3cb11c0f7229d0c8fbb39841aa948aafe3d6ff888e803fcd/README.md) |
| 25&#160;Mar&#160;25&#160;19:17&#160;UTC | SHAKEN 123L | 26&#160;Apr&#160;26&#160;15:52&#160;UTC | false | [view](CERTS/11925bc269c017e724332fcdfa849b2567cb3fab7297b45a068f2399d59e186f/README.md) |
| 26&#160;Mar&#160;25&#160;13:44&#160;UTC | SHAKEN 991K | 26&#160;Mar&#160;26&#160;13:45&#160;UTC | false | [view](CERTS/4068d2696207160d510a5b336ed6605a8f3c3914d252a52653e062d8ac5960ea/README.md) |
| 03&#160;Apr&#160;25&#160;14:06&#160;UTC | SHAKEN 155L | 03&#160;Apr&#160;26&#160;14:07&#160;UTC | false | [view](CERTS/52bbec30845d0ae93203eaef74da92345e1a93ab8612360898d584bb0ec3a7e2/README.md) |
| 03&#160;Apr&#160;25&#160;14:19&#160;UTC | SHAKEN 902K | 30&#160;Apr&#160;26&#160;15:59&#160;UTC | false | [view](CERTS/e60602f293ba4cc681635c88bf749352789afc78685435f2ba359f21c931dde0/README.md) |
| 15&#160;Apr&#160;25&#160;17:31&#160;UTC | SHAKEN 778K | 15&#160;Apr&#160;27&#160;17:32&#160;UTC | false | [view](CERTS/5c6a797e6df35d08e3435ea002cff2b24ca9002a53f48e8056710785ce95dbe3/README.md) |
| 21&#160;Apr&#160;25&#160;15:37&#160;UTC | SHAKEN 698K | 22&#160;Apr&#160;26&#160;18:39&#160;UTC | false | [view](CERTS/5579504d87539e03d7d19900820655eb5b8baacae7535e7a5038aab8e7ac3f00/README.md) |
| 03&#160;May&#160;25&#160;01:27&#160;UTC | SHAKEN 142L | 03&#160;May&#160;27&#160;01:28&#160;UTC | false | [view](CERTS/2093ae4e3a20fe19baffea6cf5b4838cd208525e1226e23ebbbd21b8a393354b/README.md) |
| 14&#160;May&#160;25&#160;18:47&#160;UTC | SHAKEN 242L | 14&#160;May&#160;26&#160;18:48&#160;UTC | false | [view](CERTS/4f0e992395004eb24a9c1fd525258ed6097bdd211a51170b0101b8edc8af6737/README.md) |
| 15&#160;May&#160;25&#160;20:02&#160;UTC | SHAKEN 865K | 20&#160;May&#160;26&#160;13:50&#160;UTC | false | [view](CERTS/2c8997d580c5d54a6b52b322329573ea88192c6e8b102534b8573ac2faad76b1/README.md) |
| 16&#160;May&#160;25&#160;22:56&#160;UTC | SHAKEN 259L | 16&#160;May&#160;26&#160;22:57&#160;UTC | false | [view](CERTS/8b1b583f024cff265283dfed260d779e62b61655c1dfb83e4c0e27a5a06c8f62/README.md) |
| 16&#160;May&#160;25&#160;23:06&#160;UTC | SHAKEN 194L | 16&#160;May&#160;27&#160;23:07&#160;UTC | false | [view](CERTS/f4226d2455292bb4a71afb8d3db5bdc107e8175e104954b6a70030f4ff5093aa/README.md) |
| 19&#160;May&#160;25&#160;21:14&#160;UTC | SHAKEN 796J | 19&#160;May&#160;27&#160;21:15&#160;UTC | false | [view](CERTS/30f88f22cb8f58597bfdce7f29e7491ce5985ead58afd4ed58562ff371dbcb56/README.md) |
| 28&#160;May&#160;25&#160;00:20&#160;UTC | SHAKEN 421K | 21&#160;Jun&#160;27&#160;03:53&#160;UTC | false | [view](CERTS/3b357603ce5157433f095eba5370d708a739ee893eff32def8f7a5dc69d6b8f8/README.md) |
| 28&#160;May&#160;25&#160;15:57&#160;UTC | SHAKEN 850K | 28&#160;May&#160;26&#160;15:58&#160;UTC | false | [view](CERTS/ebdcb190fbf3fd9d2f4d1580edb71c9badc8950023d652588efe0bbd6dda69ad/README.md) |
| 29&#160;May&#160;25&#160;00:18&#160;UTC | SHAKEN 885K | 29&#160;May&#160;27&#160;00:19&#160;UTC | false | [view](CERTS/cabf3508f382561b62c6a6d9dd662196609d5d9841c7850b2cc1a74a7b89c136/README.md) |
| 30&#160;May&#160;25&#160;16:14&#160;UTC | SHAKEN 442K | 30&#160;Jan&#160;26&#160;20:39&#160;UTC | false | [view](CERTS/7bfbc29ebb789c6f68d4830531f906ee250ad3ba96a2130b350b4980ec2b1fef/README.md) |
| 02&#160;Jun&#160;25&#160;19:56&#160;UTC | SHAKEN 279L | 02&#160;Jun&#160;26&#160;19:57&#160;UTC | false | [view](CERTS/f8b9dd590d5c388e6f32f983f94847c53c8ef1c0d8e41ddbcf878f3df9a3f6d9/README.md) |
| 03&#160;Jun&#160;25&#160;22:37&#160;UTC | SHAKEN 294L | 29&#160;May&#160;26&#160;13:13&#160;UTC | false | [view](CERTS/2200731ab9d76c73651188c3d4342271d2916f0520f85c45787994f1e365ca4c/README.md) |
| 03&#160;Jun&#160;25&#160;22:46&#160;UTC | SHAKEN 347L | 03&#160;Jun&#160;26&#160;22:47&#160;UTC | false | [view](CERTS/d53fcc47b11156668e805de002939c0636ca65c3b3fd54fe741cac3b33b7e74a/README.md) |
| 04&#160;Jun&#160;25&#160;21:46&#160;UTC | SHAKEN LUP01 | 04&#160;Jun&#160;26&#160;21:47&#160;UTC | false | [view](CERTS/cc8edc6e1bd4406fcddc877473ca384884f848a1a9b564c376e4664587a83258/README.md) |
| 05&#160;Jun&#160;25&#160;14:35&#160;UTC | SHAKEN 999K | 05&#160;Jun&#160;26&#160;14:36&#160;UTC | false | [view](CERTS/4166ee80117704133fd35316c22e71df65d375a8a333f58b37ff29512f546588/README.md) |
| 06&#160;Jun&#160;25&#160;16:31&#160;UTC | SHAKEN 179L | 06&#160;Jun&#160;26&#160;16:32&#160;UTC | false | [view](CERTS/d445d9492e9dbb5d687dceee35fd52ffa693fabf4c018ffafeba7eb6217a0552/README.md) |
| 09&#160;Jun&#160;25&#160;21:24&#160;UTC | SHAKEN 436L | 09&#160;Jun&#160;26&#160;21:25&#160;UTC | false | [view](CERTS/5e24f3b1e53ca684b47c2c2904c942e0a4429e2345bf5bf6eac9279d3f96f0d3/README.md) |
| 09&#160;Jun&#160;25&#160;21:31&#160;UTC | SHAKEN 864J | 09&#160;Jun&#160;27&#160;21:32&#160;UTC | false | [view](CERTS/1d247788de649767a33f08d128e52eb3ec37ecc0308cba97b1613d03bdc46ce3/README.md) |
| 10&#160;Jun&#160;25&#160;19:24&#160;UTC | SHAKEN 239L | 10&#160;Jun&#160;27&#160;19:25&#160;UTC | false | [view](CERTS/bdd90f740fcbfcc6a762e37bcc5cff6e7e6ba32d67981bf12811225701fd658d/README.md) |
| 11&#160;Jun&#160;25&#160;14:48&#160;UTC | SHAKEN 838J | 11&#160;Jun&#160;26&#160;14:49&#160;UTC | false | [view](CERTS/9b1a412254071bda01e0850df79208e1330c1d606d660c6b9f16e6c59fce42cf/README.md) |
| 12&#160;Jun&#160;25&#160;18:11&#160;UTC | SHAKEN 229K | 12&#160;Jun&#160;26&#160;18:12&#160;UTC | false | [view](CERTS/582bf4d452cf3b010e1be9853a5a14cb3b874bef27058233aec785e0bb62ba27/README.md) |
| 16&#160;Jun&#160;25&#160;20:21&#160;UTC | SHAKEN 324L | 16&#160;Jun&#160;26&#160;20:22&#160;UTC | false | [view](CERTS/97e98548e99f13ddf579ebebc00d1cf6ecf3484e5a2962871f37f183e3f39ee3/README.md) |
| 17&#160;Jun&#160;25&#160;19:52&#160;UTC | SHAKEN 280L | 17&#160;Jun&#160;26&#160;19:53&#160;UTC | false | [view](CERTS/3d808036615d3757a4dcdc005a4d3fff85e1f34cf275816b0e94f10665d4333f/README.md) |
| 17&#160;Jun&#160;25&#160;19:55&#160;UTC | SHAKEN 651K | 13&#160;Jul&#160;26&#160;17:57&#160;UTC | false | [view](CERTS/b25a86c1d7363b0930fa9982e85c87d504725224356db3ae628a8d46807df4f7/README.md) |
| 20&#160;Jun&#160;25&#160;01:33&#160;UTC | SHAKEN 863K | 20&#160;Jun&#160;26&#160;01:34&#160;UTC | false | [view](CERTS/26621d78d106855fd0e8492cbcdcc09632fdbb509187eb154e884a6e76d0dc92/README.md) |
| 24&#160;Jun&#160;25&#160;21:20&#160;UTC | SHAKEN 469L | 24&#160;Jun&#160;26&#160;21:21&#160;UTC | false | [view](CERTS/61156392b813d8d9c56e5fdc1ecac4ecbf4c460141bbfbf9cb3d0e21d67da235/README.md) |
| 25&#160;Jun&#160;25&#160;18:56&#160;UTC | SHAKEN 547L | 25&#160;Jun&#160;26&#160;18:57&#160;UTC | false | [view](CERTS/ca01c190a5b575222ef92238da397797bc575b05f2ddba7202b8b3c0aeb018c2/README.md) |
| 27&#160;Jun&#160;25&#160;16:24&#160;UTC | SHAKEN 386K | 27&#160;Jun&#160;27&#160;16:25&#160;UTC | false | [view](CERTS/68577bc579e72b6e09dd2a276029794591199b702847157f782f95b109b844a6/README.md) |
| 30&#160;Jun&#160;25&#160;21:27&#160;UTC | SHAKEN 317L | 30&#160;Jun&#160;27&#160;21:28&#160;UTC | false | [view](CERTS/b8ed5925be53f3f850b3d35f439ffc8256a88c96a29cd398d75810e404f3e55c/README.md) |
| 01&#160;Jul&#160;25&#160;16:41&#160;UTC | SHAKEN 457L | 01&#160;Jul&#160;26&#160;16:42&#160;UTC | false | [view](CERTS/a642d148265ee677a43b946c42e4ca8895e7fb7c0f39738173850310810a8b7f/README.md) |
| 07&#160;Jul&#160;25&#160;21:39&#160;UTC | SHAKEN 564L | 07&#160;Jul&#160;26&#160;21:40&#160;UTC | false | [view](CERTS/69c80ac8f619344813028870b8aaca4057e44ddeb33312330b36da0de8d7ffae/README.md) |
| 08&#160;Jul&#160;25&#160;20:30&#160;UTC | SHAKEN 189L | 08&#160;Jul&#160;27&#160;20:31&#160;UTC | false | [view](CERTS/b1d4583df00c272e0ce5f212c525ef1acefd9f2f95bb96baee9e5947e15c655f/README.md) |
| 18&#160;Jul&#160;25&#160;18:07&#160;UTC | SHAKEN 451L | 18&#160;Jul&#160;26&#160;18:08&#160;UTC | false | [view](CERTS/d3c187c9f14729579ac4b1007ca2a64ccf2203e2971be465047ccab6acfd645a/README.md) |
| 23&#160;Jul&#160;25&#160;14:53&#160;UTC | SHAKEN 930K | 23&#160;Jul&#160;26&#160;14:54&#160;UTC | false | [view](CERTS/c0a3b0985d9fab800f6032824b257c3fb8fd5fbfa8b6e6a8c016efb757b0e987/README.md) |
| 31&#160;Jul&#160;25&#160;21:15&#160;UTC | SHAKEN 223L | 28&#160;Aug&#160;25&#160;23:00&#160;UTC | false | [view](CERTS/a8033b2fc47332a2a4e9f06cdde19e62f5fb90a5edcce353820528514c91c708/README.md) |
| 01&#160;Aug&#160;25&#160;21:23&#160;UTC | SHAKEN 991K | 26&#160;Mar&#160;26&#160;13:45&#160;UTC | false | [view](CERTS/f555b31b01d3ff69e3c073b6328c7e76b3f32a54b9f7863852031cdc800dc707/README.md) |
| 12&#160;Aug&#160;25&#160;21:08&#160;UTC | SHAKEN 144L | 13&#160;Sep&#160;25&#160;23:13&#160;UTC | false | [view](CERTS/0c81a6f209e58a83bd550e6f2b0267d1133baab4ff626398735945734241bda0/README.md) |

#### CA Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 08&#160;Mar&#160;23&#160;18:40&#160;UTC | Telonium STI-CA Root1 | 05&#160;Mar&#160;35&#160;18:40&#160;UTC | true | [view](CERTS/96c66865ce5558c2ce3723c0b414538fcacadcd0f3286108fef57dc447f122f9/README.md) |
| 08&#160;Mar&#160;23&#160;18:40&#160;UTC | Telonium STI-CA Root2 | 07&#160;Mar&#160;38&#160;18:40&#160;UTC | true | [view](CERTS/a58b27999411d3d54121d4eadc82aa128be1fef96cda3029b2015677188ea40b/README.md) |
| 09&#160;Mar&#160;23&#160;15:18&#160;UTC | Telonium STI-CA Intermediate CA | 06&#160;Mar&#160;33&#160;15:18&#160;UTC | true | [view](CERTS/7c701216e591c9a3b84550ff46566dd420c7f182eb3cfc5abe5739cdbe271169/README.md) |
| 21&#160;Jul&#160;23&#160;00:49&#160;UTC | Telonium SHAKEN ROOT G1 | 21&#160;Jul&#160;35&#160;00:49&#160;UTC | false | [view](CERTS/37e1a126fc5d84ff59f332b2fe8196205bd0e4f7353be497ad17770d9ca6cea5/README.md) |
| 01&#160;Aug&#160;23&#160;16:36&#160;UTC | Telonium SHAKEN Intermediate G1 | 01&#160;Aug&#160;33&#160;16:36&#160;UTC | false | [view](CERTS/8c0ef8682826bec79a8c64881899f6a5a4a1d52dfebe28ae419c23f85df96ea0/README.md) |


Generated: 18 Aug 25 21:13 UTC