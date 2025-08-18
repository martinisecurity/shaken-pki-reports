# STIR/SHAKEN CA Ecosystem Compliance

## Peeringhub

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 150 certificates were included in the corpus being tested
- 23 certificates in the corpus were skipped because they are duplicates
- 76 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 51 certificates being tested against the remaining rules
- 0.00 issues on average found in unexpired, trusted, and non-compliant certificates
- 0.00% of certificates contain one or more Error level issue
- 0.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 309 days is the average remaining validity for the certificates in the corpus
- 314 days is the average initial validity for the certificates in the corpus
- 8 certificates expire in the next 30 days
- 1.02 average number of unexpired certificates per OCN observed
- 50 unique OCNs observed in unexpired and valid certificate corpus

No error, warning, or notice level issues were found

#### CA Certificates

- 2 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 0 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 2 certificates being tested against the remaining rules
- 1.00 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 0.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 100.00% of certificates are too old to be assessed against currently enforced expectations
- 5515 days is the average remaining validity for the certificates in the corpus
- 5475 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days

| Instances | Test | Source |
|-----------|------|--------|
| 1 | [e_atis_subject_cn_ca](ISSUES/e_atis_subject_cn_ca/README.md) | ATIS1000080 |
| 1 | [e_shaken_certificate_policies_id_ca](ISSUES/e_shaken_certificate_policies_id_ca/README.md) | US_SHAKEN_CP |

### Details

#### Leaf Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 19&#160;Aug&#160;24&#160;13:09&#160;UTC | SHAKEN 788J 1724072973 | 19&#160;Aug&#160;25&#160;13:09&#160;UTC | false | [view](CERTS/a93853adcdf9f247570c397d0d5d2af29ba09e0f132ff8c140aa39b326f48b38/README.md) |
| 04&#160;Sep&#160;24&#160;20:59&#160;UTC | Jaintel LLC SHAKEN 586K | 04&#160;Sep&#160;25&#160;20:59&#160;UTC | false | [view](CERTS/2b291c419ae7f36b4e6626c61b973a82ce6c2f840b90491e142de89571f827a7/README.md) |
| 11&#160;Oct&#160;24&#160;18:50&#160;UTC | VoyageNetworks_1728672609930 SHAKEN 006L | 11&#160;Oct&#160;25&#160;18:50&#160;UTC | false | [view](CERTS/c96b4c6a543b1df3110d0769b4c64e096e5749fb1ae434b0e9f317eec8bf4e6a/README.md) |
| 21&#160;Oct&#160;24&#160;18:47&#160;UTC | Voipedia_1729536460984 SHAKEN 712K | 21&#160;Oct&#160;25&#160;18:47&#160;UTC | false | [view](CERTS/77b337bd7af475c810110836e6cbd153aa04051e61712de0caed187aa1268618/README.md) |
| 07&#160;Nov&#160;24&#160;19:46&#160;UTC | Telecom Business Network LLC_1731008761381 SHAKEN 824K | 07&#160;Nov&#160;25&#160;19:46&#160;UTC | false | [view](CERTS/ffe00a2d194b964738ec33f98f6d07320e79d3407f363dc602922ca88d5a8cdb/README.md) |
| 25&#160;Nov&#160;24&#160;06:12&#160;UTC | Apex Teleocm LLC_1732515132751 SHAKEN 288K | 25&#160;Nov&#160;25&#160;06:12&#160;UTC | false | [view](CERTS/cd88f32e1fa0c6ade2be7748cff95117f7b0287e0145351a2bfec0824d00aa15/README.md) |
| 26&#160;Nov&#160;24&#160;20:24&#160;UTC | Georgialina Networks_1732652650783 SHAKEN 665K | 26&#160;Nov&#160;25&#160;20:24&#160;UTC | false | [view](CERTS/5ab2a88210a694b5779baf53946c2a3f09e79c8d9afbbceec509e04def23d079/README.md) |
| 12&#160;Dec&#160;24&#160;00:17&#160;UTC | 1234Voip.com LLC_1733962630456 SHAKEN 056L | 30&#160;Nov&#160;25&#160;22:28&#160;UTC | false | [view](CERTS/2c9bef3333855dc0338ce5a78bb57f9385213ba84691e7ae26dec1bafdd1943f/README.md) |
| 22&#160;Jan&#160;25&#160;22:57&#160;UTC | INSTACALL LLC_1737586625863 SHAKEN 281K | 22&#160;Jan&#160;26&#160;19:11&#160;UTC | false | [view](CERTS/e6a042ea51e05367aaa83b2d0be471bec283bff51ec9c6106df1010b52987c7d/README.md) |
| 03&#160;Mar&#160;25&#160;12:10&#160;UTC | VaultCom Networks Incorporated_1741003832157 SHAKEN 836K | 03&#160;Mar&#160;26&#160;12:10&#160;UTC | false | [view](CERTS/5875a7e3f07726202fe32e4de6f9bdfbab1e5357275be93df0b51e302dbb0284/README.md) |
| 10&#160;Mar&#160;25&#160;15:19&#160;UTC | VaultTel Solutions Inc_1741619990792 SHAKEN 811K | 10&#160;Mar&#160;26&#160;15:19&#160;UTC | false | [view](CERTS/34682bcacfad81aaaf20be99ae3a92f5ea1b6207e11ac74315c9f4ec7a55bd19/README.md) |
| 21&#160;Mar&#160;25&#160;21:01&#160;UTC | VoyageNetworks_1742590882637 SHAKEN 006L | 18&#160;Oct&#160;25&#160;18:41&#160;UTC | false | [view](CERTS/3ad7b737c9351e7fcdb9cde34349e3b50a84a94bccbbc2b80f3ffcbd814c6463/README.md) |
| 26&#160;Mar&#160;25&#160;20:06&#160;UTC | Infinity Sip_1743019587569 SHAKEN 279K | 26&#160;Mar&#160;26&#160;20:06&#160;UTC | false | [view](CERTS/581a906690a2c6bb1f898acbcc6319cc5939cc4c047e1020a3886a8a561d4270/README.md) |
| 31&#160;Mar&#160;25&#160;20:14&#160;UTC | TheSchmied LLC_1743452080175 SHAKEN 165L | 31&#160;Mar&#160;26&#160;20:14&#160;UTC | false | [view](CERTS/d7bee017f61367bc45f3dbe3cd0f5231a03d166b3641b3096ce133216a72401e/README.md) |
| 03&#160;Apr&#160;25&#160;12:58&#160;UTC | TalkAsiaVoip LLC_1743685125410 SHAKEN 198K | 03&#160;Apr&#160;26&#160;12:58&#160;UTC | false | [view](CERTS/f4d6413626903a53c4dd878d80dba7eeed0ce67663b188e3bfbb03eee29a3981/README.md) |
| 03&#160;Apr&#160;25&#160;20:41&#160;UTC | SHAKEN 331K | 03&#160;Apr&#160;26&#160;20:41&#160;UTC | false | [view](CERTS/f2ea79c13fbb33f721d7d87b4ed95fca835e4caf00d363dedfa6617f0f725d8c/README.md) |
| 08&#160;Apr&#160;25&#160;05:01&#160;UTC | FlowVOIP LLC._1744088504457 SHAKEN 849K | 08&#160;Apr&#160;26&#160;05:01&#160;UTC | false | [view](CERTS/8b02a3c2a6f90aadabb62241d56afa144fb8934f91e055813dcd7a6f5b264aa7/README.md) |
| 17&#160;Apr&#160;25&#160;16:30&#160;UTC | CPX Networks_1744907411997 SHAKEN 202L | 17&#160;Apr&#160;26&#160;16:30&#160;UTC | false | [view](CERTS/630445dac94acae87a4f2af2ffff4ab0961b1c17e51289129841690fb01605f0/README.md) |
| 08&#160;May&#160;25&#160;20:00&#160;UTC | Red Telecom LLC_1746734442503 SHAKEN 051L | 08&#160;May&#160;26&#160;20:00&#160;UTC | false | [view](CERTS/b630f3bf0495b4b8825a1578c0330d248ffc7f30b76db5a154d78218fae7e0c8/README.md) |
| 13&#160;May&#160;25&#160;19:38&#160;UTC | SHAKEN 088K | 13&#160;May&#160;26&#160;19:38&#160;UTC | false | [view](CERTS/9efea51511ce7ae4d8f9df9bbc31489c89bfebc0c1df07148ad06de607b15e16/README.md) |
| 16&#160;May&#160;25&#160;12:54&#160;UTC | Ajoxi Limited_1747400050960 SHAKEN 441K | 16&#160;May&#160;26&#160;12:54&#160;UTC | false | [view](CERTS/36151fe6391e5c93bd7f0a4321c96aba41ddc01755238af19fc4385f99028520/README.md) |
| 19&#160;May&#160;25&#160;16:47&#160;UTC | Volf Communications LLC_1747673274314 SHAKEN 020L | 19&#160;May&#160;26&#160;16:47&#160;UTC | false | [view](CERTS/a3fe7a44d8b96d8084849ba53666bcef4cb6ea0f1a2e6c2ea39684ca236310c6/README.md) |
| 29&#160;May&#160;25&#160;19:45&#160;UTC | Litcomp LLC_1748547950825 SHAKEN 278L | 29&#160;May&#160;26&#160;19:45&#160;UTC | false | [view](CERTS/e8daa8db2acf432aad4ac57b9b7c2fdb7c5bb5f542b9215617d3826e1872db94/README.md) |
| 09&#160;Jun&#160;25&#160;12:59&#160;UTC | ARit services LLC_1749473942413 SHAKEN 827K | 09&#160;Jun&#160;26&#160;12:59&#160;UTC | false | [view](CERTS/c18711408c1cdd54588c04f16c9e13eb7604819446494a8ec6d14fbc66bf766d/README.md) |
| 10&#160;Jun&#160;25&#160;22:43&#160;UTC | Ahoi llc_1749595415585 SHAKEN 883K | 10&#160;Jun&#160;26&#160;22:43&#160;UTC | false | [view](CERTS/2f4b681c33bbe22b5bdb78abcf33c922c1486846d20def457512c46aea827d56/README.md) |
| 11&#160;Jun&#160;25&#160;18:49&#160;UTC | Telcast Networks_1749667756733 SHAKEN 611J | 11&#160;Jun&#160;26&#160;18:49&#160;UTC | false | [view](CERTS/441193721d7ae4b15326c05f83878489eab71d49cbc21c6bf626df629551516b/README.md) |
| 12&#160;Jun&#160;25&#160;16:47&#160;UTC | Conveytel_1749746820747 SHAKEN 892K | 12&#160;Jun&#160;26&#160;16:47&#160;UTC | false | [view](CERTS/b603fcae748a76a8c35f0b6ea7928f9bae6b87528de900b50d020bc1cabd5e1c/README.md) |
| 13&#160;Jun&#160;25&#160;17:37&#160;UTC | Thenvisage LLC_1749836259934 SHAKEN 180L | 13&#160;Jun&#160;26&#160;17:37&#160;UTC | false | [view](CERTS/0855acbfe5030c89b4e2d76e8816e2c6038103cf6b3608ec2980fd18cb6b5176/README.md) |
| 16&#160;Jun&#160;25&#160;14:21&#160;UTC | VOXTiX_1750083693916 SHAKEN 343L | 16&#160;Jun&#160;26&#160;14:21&#160;UTC | false | [view](CERTS/8a4ba0d482d65e9cd4f8ff9741b0eb8df3058d2565027a7e4c05e4a47dc64923/README.md) |
| 17&#160;Jun&#160;25&#160;00:00&#160;UTC | VOCALTRANSIT SHAKEN 783J | 17&#160;Dec&#160;25&#160;00:00&#160;UTC | false | [view](CERTS/6eddb4fadf5a409fde3a4f52b52a48edac9f9522cc6136bee8b2fbf03f16cab8/README.md) |
| 19&#160;Jun&#160;25&#160;21:39&#160;UTC | Voxology Carrier Services, Inc._1750369178545 SHAKEN 664K | 19&#160;Jun&#160;26&#160;21:39&#160;UTC | false | [view](CERTS/a809bc07d3033c1e4189faef71fc85fe30d370c2ca283a7d3db1a3dd4209c6d0/README.md) |
| 27&#160;Jun&#160;25&#160;22:32&#160;UTC | Nextricify_1751063571710 SHAKEN 529L | 27&#160;Jun&#160;26&#160;22:32&#160;UTC | false | [view](CERTS/62fc8acaa69cd1249186565c004c0c9304de39488f657172f37f6c3e3fb3f097/README.md) |
| 27&#160;Jun&#160;25&#160;22:47&#160;UTC | SIP CREW LLC_1751064433061 SHAKEN 974K | 27&#160;Jun&#160;26&#160;22:47&#160;UTC | false | [view](CERTS/bea48779779ab745175e0d6053b4abd88bf75afc7c362e09c6934c6f95d607fa/README.md) |
| 28&#160;Jun&#160;25&#160;00:26&#160;UTC | DURKETA, LLC_1751070374071 SHAKEN 284L | 28&#160;Jun&#160;26&#160;00:26&#160;UTC | false | [view](CERTS/4cc94296f12884c36ad4f1aaca4f337df8d9b3b595bd159c97cdaf34e52ebd2c/README.md) |
| 03&#160;Jul&#160;25&#160;22:40&#160;UTC | SHAKEN 338L XEN Telecom_1751582413468 | 26&#160;Jun&#160;26&#160;22:45&#160;UTC | false | [view](CERTS/43c6d307ecde8a3fd3065c41562415a347250584512e785aad3f67864558fa53/README.md) |
| 07&#160;Jul&#160;25&#160;18:53&#160;UTC | SHAKEN 551L 1751914399544 | 07&#160;Jul&#160;26&#160;18:53&#160;UTC | false | [view](CERTS/4da31b5f512a4654c0a8dee53f63bb5e168d20a69a3867d8774f1f71b42553ca/README.md) |
| 08&#160;Jul&#160;25&#160;15:58&#160;UTC | SHAKEN 588L 1751990302539 | 08&#160;Jul&#160;26&#160;15:58&#160;UTC | false | [view](CERTS/1dcee861ed74d24b46982875d83b2333fd22fcf01e739b045e6cb45b9424f338/README.md) |
| 08&#160;Jul&#160;25&#160;16:50&#160;UTC | SHAKEN 103L 1751993458160 | 08&#160;Jul&#160;26&#160;16:50&#160;UTC | false | [view](CERTS/fb087940428b4708a949d3aa4a0e00bab2a1139a647924eb3bc92bbeb3d423a6/README.md) |
| 16&#160;Jul&#160;25&#160;07:54&#160;UTC | SHAKEN 132L 1752652441523 | 23&#160;Aug&#160;25&#160;07:46&#160;UTC | false | [view](CERTS/03403a4aa6acdb65e5da0384977b7fda213d25f3bbcec48203b3d44e9368933c/README.md) |
| 16&#160;Jul&#160;25&#160;16:01&#160;UTC | SHAKEN 507K 1752681709200 | 21&#160;Aug&#160;25&#160;16:46&#160;UTC | false | [view](CERTS/fe2176846ce8b6a1c51b50ec5e0472da04ba3ac573bc7196e14054edf2e11cfd/README.md) |
| 21&#160;Jul&#160;25&#160;21:30&#160;UTC | SHAKEN 255K 1753133429807 | 28&#160;Aug&#160;25&#160;21:28&#160;UTC | false | [view](CERTS/1f8f81fd3083954b79b29cc2664b0893ac2080d66258842b9ecd58d5da478340/README.md) |
| 22&#160;Jul&#160;25&#160;20:18&#160;UTC | SHAKEN 861J 1753215520750 | 22&#160;Jul&#160;26&#160;20:18&#160;UTC | false | [view](CERTS/329bce77edb00a58fb5011d6dc7ebd7640a4baad2a02b838f4bbd3805a625d7b/README.md) |
| 22&#160;Jul&#160;25&#160;21:06&#160;UTC | SHAKEN 458K 1753218382009 | 22&#160;Jul&#160;26&#160;21:06&#160;UTC | false | [view](CERTS/e8794656d3a0640273260f79ac65e354ae5b241bb3ea4ca90c35873ca927ff90/README.md) |
| 23&#160;Jul&#160;25&#160;19:09&#160;UTC | SHAKEN 745K 1753297774749 | 23&#160;Sep&#160;25&#160;18:52&#160;UTC | false | [view](CERTS/5e858426875f6fef27509283e9cd51fe260139cea24157df7ab486d2fcd5f97c/README.md) |
| 23&#160;Jul&#160;25&#160;22:42&#160;UTC | SHAKEN 727K 1753310551386 | 23&#160;Jul&#160;26&#160;22:42&#160;UTC | false | [view](CERTS/e99feaa30505b76f2bfa48b275667984125e36e1d2082ebcd8608bafbfa3720e/README.md) |
| 24&#160;Jul&#160;25&#160;15:17&#160;UTC | SHAKEN 623L 1753370251926 | 24&#160;Jul&#160;26&#160;15:17&#160;UTC | false | [view](CERTS/3bdfcc5eeb8e9284b66cd257c1bebef238389beb07c0ee60529c3f0b1cba5d70/README.md) |
| 29&#160;Jul&#160;25&#160;17:59&#160;UTC | SHAKEN 469K 1753811973709 | 05&#160;Sep&#160;25&#160;17:59&#160;UTC | false | [view](CERTS/5e38d610d31c1aea55ba7782995f1676d96cca101c71404a17980461981a26fc/README.md) |
| 30&#160;Jul&#160;25&#160;06:50&#160;UTC | SHAKEN 645L 1753858238261 | 30&#160;Jul&#160;26&#160;06:50&#160;UTC | false | [view](CERTS/12ed19eacde3f35b6bb81c26175c1dfa513639bf6d1c3754996b7a9bd69fd175/README.md) |
| 31&#160;Jul&#160;25&#160;21:07&#160;UTC | SHAKEN 648L 1753996053266 | 31&#160;Jul&#160;26&#160;21:07&#160;UTC | false | [view](CERTS/6a027fcdc9b18b04f924527f6a74c9dd492cc70d0f143d5b03ea2b5418208721/README.md) |
| 01&#160;Aug&#160;25&#160;13:34&#160;UTC | SHAKEN 066L 1754055248211 | 03&#160;Sep&#160;25&#160;14:51&#160;UTC | false | [view](CERTS/6771b8bff8ddd57adcbf21f5b5e93c85ce6fb8b1d55aa0ba15646815909d939f/README.md) |
| 06&#160;Aug&#160;25&#160;17:13&#160;UTC | SHAKEN 074K 1754500391789 | 11&#160;Sep&#160;25&#160;15:27&#160;UTC | false | [view](CERTS/51fcd12876e27676d8ac12b36ff2aebec0464f539d191adefe88efa624d54778/README.md) |

#### CA Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 17&#160;Dec&#160;20&#160;15:31&#160;UTC | Peeringhub Inc Root CA | 12&#160;Dec&#160;40&#160;15:31&#160;UTC | true | [view](CERTS/0ad4adc0b4d93fdb0e628c577020c73b8a5caff750e7e499f80ee2ab362a3f6a/README.md) |
| 22&#160;Jun&#160;22&#160;22:45&#160;UTC | Peeringhub Inc SHAKEN Intermediate CA 2 | 19&#160;Jun&#160;32&#160;22:45&#160;UTC | true | [view](CERTS/f00871963a40b04269c4b019968e42f9f40964cbfb512ff5342307e9942874ce/README.md) |


Generated: 18 Aug 25 21:13 UTC