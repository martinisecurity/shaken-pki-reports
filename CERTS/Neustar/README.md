# STIR/SHAKEN CA Ecosystem Compliance

## Neustar

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 207 certificates were included in the corpus being tested
- 108 certificates in the corpus were skipped because they are duplicates
- 2 certificates in the corpus were skipped because they are expired
- 1 certificates in the corpus were skipped because they are not currently trusted
- 96 certificates being tested against the remaining rules
- 2.96 issues on average found in unexpired, trusted, and non-compliant certificates
- 93.75% of certificates contain one or more Error level issue
- 33.33% of certificates contain one or more Warning level issue
- 33.33% of certificates contain one or more Notice level issue
- 33.33% of certificates are too old to be assessed against currently enforced expectations
- 589 days is the average remaining validity for the certificates in the corpus
- 593 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days

| Instances | Test Status | Source |
|-----------|-------------|--------|
| 78 | [e_atis_certificate_policies](ISSUES/e_atis_certificate_policies/README.md) | ATIS1000080 |
| 59 | [e_atis_extension_unknown](ISSUES/e_atis_extension_unknown/README.md) | ATIS1000080 |
| 26 | [e_us_cp_ambiguous_identifier](ISSUES/e_us_cp_ambiguous_identifier/README.md) | US_SHAKEN_CP |
| 89 | [e_us_cp_subject_sn](ISSUES/e_us_cp_subject_sn/README.md) | US_SHAKEN_CP |
| 32 | [w_shaken_subject_rdn_unknown](ISSUES/w_shaken_subject_rdn_unknown/README.md) | SHAKEN_PKI_BEST_PRACTICES |

#### CA Certificates

- 7 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 0 certificates in the corpus were skipped because they are expired
- 1 certificates in the corpus were skipped because they are not currently trusted
- 6 certificates being tested against the remaining rules
- 1.83 issues on average found in unexpired, trusted, and non-compliant certificates
- 66.67% of certificates contain one or more Error level issue
- 100.00% of certificates contain one or more Warning level issue
- 66.67% of certificates contain one or more Notice level issue
- 66.67% of certificates are too old to be assessed against currently enforced expectations
- 5155 days is the average remaining validity for the certificates in the corpus
- 4870 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days

| Instances | Test Status | Source |
|-----------|-------------|--------|
| 1 | [e_atis_ca_extension_unknown](ISSUES/e_atis_ca_extension_unknown/README.md) | ATIS1000080 |
| 4 | [e_us_cp_ca_key_usage_crl_sign](ISSUES/e_us_cp_ca_key_usage_crl_sign/README.md) | US_SHAKEN_CP |
| 6 | [w_shaken_ca_subject_rdn_unknown](ISSUES/w_shaken_ca_subject_rdn_unknown/README.md) | SHAKEN_PKI_BEST_PRACTICES |

### Details

#### Leaf Certificates

| Created at | Subject | Problems | Link |
|------------|---------|----------|------|
| 18 Dec 19 20:01 UTC | alticeusa.com | true | [view](CERTS/5254d4ab17f12340d76d15e07e4b3f881b4c6ffd1fd1c2147b26ede27098ad9d/README.md) |
| 18 Dec 19 20:08 UTC | vonage | true | [view](CERTS/e6d9c80f45bc76e1389d8f6213cf3d382c5b2f3839184e3be15f1e03b6c780e7/README.md) |
| 07 Jan 20 18:48 UTC | peerlessnetwork.com | true | [view](CERTS/58950c854b20b1343a688944b90a65f56f820a98192d396695f4525521085f84/README.md) |
| 10 Jan 20 20:09 UTC | CenturyLink.com | true | [view](CERTS/266e4919267dfd279237a8e32923b1484c606dd8eb3f5a07b5c676d7cf78c053/README.md) |
| 29 Jan 20 15:16 UTC | frontier.com | true | [view](CERTS/bd545bbe3cb81d574b5b7d358bbd704139506cde8a612b24c24bdf6d0eeacd8e/README.md) |
| 25 Mar 20 17:29 UTC | cox.com | true | [view](CERTS/21a9862f2a9b3b61084354f282f285d9a7b563d755d7e1a8ceb9669b6e6e71e3/README.md) |
| 04 Jun 20 18:39 UTC | AGOC | true | [view](CERTS/601d6d6df50a7b1431e92d7955fbacf9254cd20f171ba02b8918b475504ecec1/README.md) |
| 09 Jul 20 15:29 UTC | digitalipvoice.com | true | [view](CERTS/b95967027d535f36fc40bb91a16bc17d7fe58af8fa14d922a1c9daed9933443a/README.md) |
| 15 Jul 20 04:13 UTC | SHAKEN | true | [view](CERTS/a76f137e6ae23b3f27db8b6c2c339571ff7b5a106f61709d6d36b49852d0a070/README.md) |
| 20 Aug 20 00:54 UTC | Inteliquent.com | true | [view](CERTS/f1c1fe53212d9bd9b211b83c698572a9019078dd441e9e78638e55714dcafaab/README.md) |
| 15 Sep 20 13:16 UTC | intrado.com | true | [view](CERTS/e28fb52c91d5ec227c26e93f02d1b1412bfcb534bca76b69c0a0de93fe26222b/README.md) |
| 25 Sep 20 00:25 UTC | SHAKEN | true | [view](CERTS/158c61d4da5ed8d0aaef7582d9a821a358a22ec2f241d449ecdab2e35ccf2fad/README.md) |
| 21 Jan 21 18:49 UTC | BigRiverTelephoneCompany | true | [view](CERTS/9efdf79a4d93cef188ca8bc8dd4ce626a21813ed665bf108d74ff1f0158caf24/README.md) |
| 09 Feb 21 21:49 UTC | Ringcentral-ProdKeystore | true | [view](CERTS/043a925e40a83e722da0360f49488513cd8d8cfadbea2469800e80a673541fc4/README.md) |
| 16 Feb 21 15:12 UTC | Logix-Keystore | true | [view](CERTS/c53c156a26a6dc23f61a04e882c7643e9285b94374d9e82df16623f9519bdd65/README.md) |
| 21 Feb 21 13:57 UTC | Shaken | true | [view](CERTS/58f3d6011c5d96eac83256a4c25ea4a085137b7bd6d4ba1acc246e4098c25ff9/README.md) |
| 11 Mar 21 18:18 UTC | SHAKEN-6744 | true | [view](CERTS/414671d6f2e7beffdd958279b4cb2e705c5ee59f107aa1fb7b2a06008ae117b6/README.md) |
| 12 Mar 21 14:58 UTC | Flowroute | true | [view](CERTS/ff098de1abec5c5569a38f6ad950c3e478009ddef15d3e456b0884ffaceb0cec/README.md) |
| 24 Mar 21 14:02 UTC | Firstcomm5917 | true | [view](CERTS/394a79b142d1ac314cfabcd5009dc93fff81c1e19a5d14fe5175ae197ab5c66f/README.md) |
| 08 Apr 21 09:47 UTC | Orange | true | [view](CERTS/ff8031532cb64b24c610ce0c4887726fb4d2b915322e0dac4e004f9f51ac118a/README.md) |
| 16 Apr 21 20:26 UTC | Five9 | true | [view](CERTS/08087af11d90b5b7bd3519ce065262e6ec44ea1dcfbc1db3995cc1b0a08fd48e/README.md) |
| 19 Apr 21 23:44 UTC | Warmego | true | [view](CERTS/65bfa3f8b6aa342704363e4e0db5198f3f3583a9663f17901a7bb96fe2bfddf1/README.md) |
| 03 May 21 21:11 UTC | WindstreamCommunication | true | [view](CERTS/8626a6fbbb1f434721fe9a149a77d13a6ace0da4578904229794070248323973/README.md) |
| 04 May 21 11:46 UTC | intelepeer.com | true | [view](CERTS/49c3c8001a4b6cd71ae4cc41b0e36f2c724c5e024e19d99b71956f563d14ab01/README.md) |
| 06 May 21 15:50 UTC | SHAKEN-8468 | true | [view](CERTS/369607701f40ce537f70ed798778773367baf6683606e11b81b8eae0eb85c6ad/README.md) |
| 07 May 21 14:09 UTC | SHAKEN_0377 | true | [view](CERTS/a891efce3a4b35ed5111a2f8d9dbf2abef04a0e5c9ce4ea5313c0ede3e92ce26/README.md) |
| 07 May 21 16:55 UTC | PRD | true | [view](CERTS/d81e85538a8dc39d06fa99ded608ea70df6631e79431ec10fc7b4173cee3b991/README.md) |
| 18 May 21 15:07 UTC | Granite | true | [view](CERTS/42e6a55be823bf374996654d48501a9c98c1b37cdd523a496d9ed9a08044b7cf/README.md) |
| 19 May 21 14:02 UTC | 846B | true | [view](CERTS/91646ad947e19e7b3501103e2d65a5e66ab7bc806c8bb7d4e87cce23ef668183/README.md) |
| 21 May 21 16:38 UTC | ReInvent | true | [view](CERTS/ffb808b6e361219b3f20bc740eae385eee965cd38ad8b7b4e425a0ed25c611ee/README.md) |
| 15 Dec 21 19:36 UTC | SHAKEN 775J | true | [view](CERTS/aaea251bd6d9305ec47058014214117750c0b0059a1d221f4e584e9b680aa560/README.md) |
| 11 Jan 22 21:31 UTC | SHAKEN 506J | true | [view](CERTS/00635768d8b9c2e0fc7fd4d35c6c192107102b0f4ba22ddf9c0c5c76ee257d32/README.md) |
| 18 Jan 22 18:06 UTC | SHAKEN 899J | true | [view](CERTS/df5ff82d270dec47f3255a7415d68ae4348953383baedd15c367a34752e1c7ec/README.md) |
| 20 Jan 22 17:34 UTC | SHAKEN 750J | true | [view](CERTS/52cb96a36b3dac1475b04a10b545d0f9399e4dc1b3227aaccb7dc4a01278e7c3/README.md) |
| 02 Feb 22 20:45 UTC | SHAKEN 785J | true | [view](CERTS/433a795031c3e471766ce6bdfcff8da439d089a2209b44167af945b481151a18/README.md) |
| 04 Mar 22 20:34 UTC | SHAKEN 863J | true | [view](CERTS/a823137b09d7d9bc4b78bcef8048a37ceeb899121f4d227e0fc5e784faef7038/README.md) |
| 16 Mar 22 13:10 UTC | SHAKEN 393J | true | [view](CERTS/1f41d26a18d6dbb6f9a0710eea2b870dd939d4ac06b530761b055d215a12f14f/README.md) |
| 18 Mar 22 18:18 UTC | SHAKEN 804J | true | [view](CERTS/b6ee8f5ec7b9c2d085d1e950408477af3a3647a513dbfe0cbfcc7585fb7e1763/README.md) |
| 22 Mar 22 14:44 UTC | SHAKEN 701J | true | [view](CERTS/5cb4eb85a71e34614d1be9d02610289539c46e8e183fe36247b25a78d94c3377/README.md) |
| 22 Mar 22 18:39 UTC | SHAKEN 597F | true | [view](CERTS/09db47bffeb34d5524095972b6eb3ec83aeb917757c32fc0bf473437c8f1b42d/README.md) |
| 24 Mar 22 19:58 UTC | SHAKEN 917J | true | [view](CERTS/e20d3859c17ea71fc31a6c5fd81daa1ca0394335b59be10f44d0cb463a7b5a68/README.md) |
| 30 Mar 22 00:05 UTC | SHAKEN 5447 | true | [view](CERTS/2e3df7cab0124f9c13fa56485daf2d8f8d6bc98cd9b07c15ca02ffceccbc1072/README.md) |
| 06 Apr 22 12:06 UTC | SHAKEN 963J | true | [view](CERTS/89a285de49404991c7aef6ece2e9172f0e3ba219c221a49bbe039a2080f6244d/README.md) |
| 08 Apr 22 18:10 UTC | SHAKEN 973J | true | [view](CERTS/a5dbf6d412928fdb65a65a75f7df93724892c57efb4ae1e6e9161d740309df48/README.md) |
| 11 Apr 22 18:33 UTC | SHAKEN 951J | true | [view](CERTS/b1ba32c225a987ec5d2b5ecb2f3638488cbc4b08a71ce1f447af83ab6b67a30c/README.md) |
| 20 Apr 22 20:54 UTC | SHAKEN 966J | true | [view](CERTS/de066971a52eed6408d1d00c985cee5ab0554844ff694b8660bd627644264b7e/README.md) |
| 23 Apr 22 01:57 UTC | SHAKEN 502E | true | [view](CERTS/71b03e0429cbcc39d1a9b0194ea163a28c7a980d58675d8a53eefb632597bc18/README.md) |
| 27 Apr 22 19:37 UTC | SHAKEN 782J | true | [view](CERTS/4f42d1733ab75724e7ca941c3e6fad930113cb66e0fc80f8a3f59815b7bc6aff/README.md) |
| 02 May 22 13:10 UTC | SHAKEN 854J | true | [view](CERTS/4535b112c5d3933f3d71bc6e4cc98a71479c9618690d778c85c170cf41422e76/README.md) |
| 09 May 22 14:53 UTC | SHAKEN 821J | true | [view](CERTS/3e6046796b5cc7710e323960eff671ce5e0e2aef1b21e8657e0c3564e09877e1/README.md) |
| 10 May 22 10:42 UTC | SHAKEN 005K | true | [view](CERTS/639ec0e02a1c7bd0b564efb570da5cdd6e681facf141a9f696be31bc2103bb8e/README.md) |
| 10 May 22 18:17 UTC | SHAKEN 402E | true | [view](CERTS/28ad10381bce413d117b7eded64501ec8efb9a2587e41cc0d63ceae2f35ad351/README.md) |
| 13 May 22 14:37 UTC | SHAKEN 473G | true | [view](CERTS/12cde52b1a2b43aeeab573361ef36eeb5b4783bf72442aeb0737ff7c9fba0fa2/README.md) |
| 17 May 22 21:17 UTC | SHAKEN 030K | true | [view](CERTS/0996dbe49d9bef155744a54e00e3e0440a145ba18e1ca6629f067be906609e49/README.md) |
| 19 May 22 17:45 UTC | SHAKEN 618J | true | [view](CERTS/15cda6f5e0190b5dad65a3ed1ed1ac9459c30c269cf7b800417bdeb214832853/README.md) |
| 22 May 22 15:34 UTC | SHAKEN 772J | true | [view](CERTS/644bdc695198227f9a5bedb1c70aab96c11416e2ca996c0cbeb03a906cbc54d6/README.md) |
| 23 May 22 16:43 UTC | SHAKEN 3130 | true | [view](CERTS/9710435c3c5ba2df401e4295efd73c53830e1372dcf07dff0601125d129c2de5/README.md) |
| 24 May 22 15:44 UTC | SHAKEN 845J | true | [view](CERTS/5779b9ac57abada079f96ac7625dfe923e44f6993378e4291a5ed16b1ba65a64/README.md) |
| 24 May 22 16:18 UTC | SHAKEN 4427 | true | [view](CERTS/a9f8804c4ba7e8886318aece12a9eb6e4f755708e0bcac09cf1bb2a136880369/README.md) |
| 25 May 22 16:41 UTC | SHAKEN 672J | true | [view](CERTS/30bbe0bdbf6c302c8da0f6e23fb56df86d12bd6a784383dc6e475d2b6a172925/README.md) |
| 31 May 22 15:42 UTC | SHAKEN 704J | true | [view](CERTS/7d6ad3d1bfc3820ad65514ccd15606e93f9df5c8732fe3e374b3e3522c40cf57/README.md) |
| 08 Jun 22 15:19 UTC | SHAKEN 534J | true | [view](CERTS/eec71348d87837eafbc112c892fe6c362c79adb14ae697cf60522e03f19230c3/README.md) |
| 08 Jun 22 16:50 UTC | SHAKEN 819J | true | [view](CERTS/127f660276cb1085f74a7f9955b94ced3bd3452a57c26a8240c31fd6d29ee512/README.md) |
| 12 Jun 22 15:12 UTC | SHAKEN 261H | true | [view](CERTS/6815b252b41b1d162968ea43df4ae48355a8b786816ec2434d65c703e2a5a8fc/README.md) |
| 12 Jun 22 15:18 UTC | SHAKEN 033H | true | [view](CERTS/dbc2c40b64af4f892b4991d76c182bc85345235069d1cadf854bd2ef5d7fee63/README.md) |
| 13 Jun 22 18:51 UTC | SHAKEN 554J | true | [view](CERTS/48e98eb1d0597ad80afaba38213dc9da04352e1ca69b6e98b56d4d556c35b354/README.md) |
| 16 Jun 22 18:26 UTC | SHAKEN 939H | true | [view](CERTS/a9ff363f03fba28c1972002e99200ecd750e4aaf59fa7302e39b8777be615e08/README.md) |
| 17 Jun 22 21:25 UTC | SHAKEN 573J | true | [view](CERTS/235085bee8917aacad49680ef5f7e11ff56192cf5239041332dcb1d0f858cb2e/README.md) |
| 21 Jun 22 18:39 UTC | SHAKEN 712J | true | [view](CERTS/869b4bd6e2405b41153c7a848f5585a914e488fab44fb7ebb556f91c2ab07df4/README.md) |
| 21 Jun 22 19:56 UTC | SHAKEN 254H | true | [view](CERTS/8605c03e7a4468d46dd9b066a26a157499596a64c033f0e029d1a14dd756a56a/README.md) |
| 21 Jun 22 21:04 UTC | SHAKEN 611J | true | [view](CERTS/3f87c16b475eb0fc6235d4e73fe5f990364d4c2f109cceac951e5e8b449fcd56/README.md) |
| 23 Jun 22 16:36 UTC | SHAKEN 553J | true | [view](CERTS/a804cf0f4ff52fb03a1d23618ecabb5eda1527649e3347ca6a146b22cd024d3b/README.md) |
| 23 Jun 22 18:09 UTC | SHAKEN 171K | true | [view](CERTS/e9355468f55b838570f9c50ac0abbeaa64f565b5a37c9653fbb07b59689d1771/README.md) |
| 24 Jun 22 14:05 UTC | SHAKEN 739J | true | [view](CERTS/61bb2e5f8faffe66e9d319704c906bb3c211b31d62ef605c738f231ee3b305aa/README.md) |
| 24 Jun 22 19:58 UTC | SHAKEN 067K | true | [view](CERTS/aed8214f2128df08605bba6f6596ee4fc2874f44577591a6e254c7117b819843/README.md) |
| 28 Jun 22 20:15 UTC | SHAKEN 738J | true | [view](CERTS/4fc19fc620e600a7a819b80c3cc596dc0139a567ae19484472728564c8423574/README.md) |
| 28 Jun 22 21:58 UTC | SHAKEN 743J | true | [view](CERTS/afecedf9de456e4170f9f52742f42bdca9437dcabb6867c7c71372f8ba4feddf/README.md) |
| 29 Jun 22 22:00 UTC | SHAKEN 049K | true | [view](CERTS/37d491d2756716824cbf4707b31a034fef6584e4e6d95d327a7428e68e92a02f/README.md) |
| 30 Jun 22 01:51 UTC | SHAKEN 235K | true | [view](CERTS/cad17de4a8e16887b31ad2c62a00554e87f43ee1a45f477fa4193b3cfb1fe281/README.md) |
| 01 Jul 22 16:21 UTC | SHAKEN 139K | true | [view](CERTS/eacdef99ac79c2ae3e713ec4c5972e785548050a35db9747f21eb2d4bc06c31b/README.md) |
| 08 Jul 22 17:49 UTC | SHAKEN 709J | true | [view](CERTS/a371d70a1b505735b513c7488d811eb869c801aa246cfc1424db1e373b579f34/README.md) |
| 13 Jul 22 17:30 UTC | SHAKEN 962J | true | [view](CERTS/cd04dcf7d5eb1acdb0c0c3e60fb3977707961e4240f1dfcf67eff0338bd75ffb/README.md) |
| 13 Jul 22 19:10 UTC | SHAKEN 771J | true | [view](CERTS/0d7b51533df7b19ce84f11cf53ff9c2260e112501402635565d154d9e9741d63/README.md) |
| 18 Jul 22 17:33 UTC | SHAKEN 715J | true | [view](CERTS/450172823b48d29255bde6a9ec9e921f8e275b76c1a33986f23fa9d410802eac/README.md) |
| 19 Jul 22 17:18 UTC | SHAKEN 023K | true | [view](CERTS/930eea9676ea20bbbf31176053b4990fd6e912ab6c7bfe0fa2619b6f5c0ce864/README.md) |
| 28 Jul 22 20:56 UTC | SHAKEN 074K | true | [view](CERTS/ce2299b32a7f2749baeb8864b84a8b531591a05a0c79c3bff8af7b8da441d513/README.md) |
| 08 Aug 22 12:58 UTC | SHAKEN 150K | true | [view](CERTS/5cf325af20a32e0691ff11a47485c841d9dd35771bbe2fe8027deb13e70b2a81/README.md) |
| 08 Aug 22 14:30 UTC | SHAKEN 710A | true | [view](CERTS/e97dffcbf85f9183ac12212e1a9e342f76ffd96d6c6d280854fb31e9465d481a/README.md) |
| 18 Aug 22 18:07 UTC | SHAKEN 219K | true | [view](CERTS/fc1a6306ba8c8e009e014efbcf859e5740cae4ebfac67e9628ef65e16209c9b1/README.md) |
| 29 Aug 22 19:07 UTC | ATT SHAKEN 4036 | true | [view](CERTS/3a8d4d5fe47e784f925bca30b21d44f3492ad0813fd074ef2d438cca1d4acc68/README.md) |
| 31 Aug 22 14:47 UTC | SHAKEN 500J | true | [view](CERTS/25393ddb6f4b6a42df56e2cafbff0fa710e60bbdf94d37561290cdb6e66f28d1/README.md) |
| 09 Sep 22 14:44 UTC | SHAKEN 5606 | true | [view](CERTS/8fc7d03e6cd7ab01c7c8b9051bfef91cd5bceecefd08f74b6fc948b65b15eca4/README.md) |
| 12 Sep 22 19:52 UTC | SHAKEN 707J | true | [view](CERTS/132d0f59c15814fc3f80760fdf109ee36d0ac19eaf03a74457523f2cf4dcf982/README.md) |
| 15 Sep 22 16:20 UTC | SHAKEN 292K | true | [view](CERTS/e449581f068a3e747acc8dfd7c0707c9aed3deca9f572afa13bbb0ffbebffd27/README.md) |
| 26 Oct 22 16:36 UTC | SHAKEN 770J | true | [view](CERTS/c0d94a93ff7d9519ab97a0e8a44ecae02198e5df6ca89a8783163d00a80583ea/README.md) |
| 28 Oct 22 14:59 UTC | prod SHAKEN 811J | true | [view](CERTS/f6e3384410e423b0e0006618d97878154ef07b52fbdc31a7ae6f3bc7f102e6ed/README.md) |

#### CA Certificates

| Created at | Subject | Problems | Link |
|------------|---------|----------|------|
| 23 Sep 19 13:26 UTC | Neustar Certified Caller ID Root CA | true | [view](CERTS/4a77c17cd411cb0ff2984b97687f75ab1db451ac7b717ab81c931351c2d547a1/README.md) |
| 23 Sep 19 13:32 UTC | Neustar Certified Caller ID CA-1 | true | [view](CERTS/dade1a52e76c29fc9af1e1221a2a6be02c9899a552d396580855935c9592733b/README.md) |
| 17 Aug 21 17:19 UTC | Neustar Certified Caller ID SHAKEN Root CA | true | [view](CERTS/4c728d18b628cc67dda5490e0b4aa8ef4ba679f96d033f34f1680e219e0806c3/README.md) |
| 19 Aug 21 03:25 UTC | Neustar Certified Caller ID SHAKEN CA-1 | true | [view](CERTS/b6dc9bf58a55979c78ad569a17c86a7f644721bd3ab2bcf99a27d13636900cf4/README.md) |
| 30 Aug 22 06:39 UTC | Neustar Certified Caller ID SHAKEN CA-2 | true | [view](CERTS/3ea530838e9952fdda913a8bd669bf37f88f4ffdb39a34698f34a63915c9e404/README.md) |
| 05 Oct 22 17:26 UTC | Neustar Certified Caller ID SHAKEN CA-2 | true | [view](CERTS/0bd95ecbb97c09de0df079ca41e10c360c4b5928ac56c496879a2c90c6bbffe4/README.md) |


Generated: 02 Nov 22 15:10 UTC