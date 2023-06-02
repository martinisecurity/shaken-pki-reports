# STIR/SHAKEN CA Ecosystem Compliance

## Neustar

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 363 certificates were included in the corpus being tested
- 200 certificates in the corpus were skipped because they are duplicates
- 22 certificates in the corpus were skipped because they are expired
- 3 certificates in the corpus were skipped because they are not currently trusted
- 138 certificates being tested against the remaining rules
- 2.04 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 18.84% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 18.12% of certificates are too old to be assessed against currently enforced expectations
- 496 days is the average remaining validity for the certificates in the corpus
- 497 days is the average initial validity for the certificates in the corpus
- 20 certificates expire in the next 30 days
- 1.03 average number of unexpired certificates per OCN observed
- 134 unique OCNs observed in unexpired and valid certificate corpus

| Instances | Test | Source |
|-----------|------|--------|
| 55 | [e_atis_certificate_policies](ISSUES/e_atis_certificate_policies/README.md) | ATIS1000080 |
| 36 | [e_atis_extension_unknown](ISSUES/e_atis_extension_unknown/README.md) | ATIS1000080 |
| 27 | [e_us_cp_ambiguous_identifier](ISSUES/e_us_cp_ambiguous_identifier/README.md) | US_SHAKEN_CP |
| 138 | [e_us_cp_subject_sn](ISSUES/e_us_cp_subject_sn/README.md) | US_SHAKEN_CP |
| 26 | [w_shaken_subject_rdn_unknown](ISSUES/w_shaken_subject_rdn_unknown/README.md) | SHAKEN_PKI_BEST_PRACTICES |

#### CA Certificates

- 8 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 0 certificates in the corpus were skipped because they are expired
- 2 certificates in the corpus were skipped because they are not currently trusted
- 6 certificates being tested against the remaining rules
- 1.83 issues on average found in unexpired, trusted, and non-compliant certificates
- 66.67% of certificates contain one or more Error level issue
- 100.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 66.67% of certificates are too old to be assessed against currently enforced expectations
- 5125 days is the average remaining validity for the certificates in the corpus
- 4870 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days

| Instances | Test | Source |
|-----------|------|--------|
| 1 | [e_atis_ca_extension_unknown](ISSUES/e_atis_ca_extension_unknown/README.md) | ATIS1000080 |
| 4 | [e_us_cp_ca_key_usage_crl_sign](ISSUES/e_us_cp_ca_key_usage_crl_sign/README.md) | US_SHAKEN_CP |
| 6 | [w_shaken_ca_subject_rdn_unknown](ISSUES/w_shaken_ca_subject_rdn_unknown/README.md) | SHAKEN_PKI_BEST_PRACTICES |

### Details

#### Leaf Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 15&#160;Jul&#160;20&#160;04:13&#160;UTC | SHAKEN | 16&#160;Jul&#160;23&#160;04:13&#160;UTC | true | [view](CERTS/a76f137e6ae23b3f27db8b6c2c339571ff7b5a106f61709d6d36b49852d0a070/README.md) |
| 20&#160;Aug&#160;20&#160;00:54&#160;UTC | Inteliquent.com | 21&#160;Aug&#160;23&#160;00:54&#160;UTC | true | [view](CERTS/f1c1fe53212d9bd9b211b83c698572a9019078dd441e9e78638e55714dcafaab/README.md) |
| 15&#160;Sep&#160;20&#160;13:16&#160;UTC | intrado.com | 16&#160;Sep&#160;23&#160;13:16&#160;UTC | true | [view](CERTS/e28fb52c91d5ec227c26e93f02d1b1412bfcb534bca76b69c0a0de93fe26222b/README.md) |
| 25&#160;Sep&#160;20&#160;00:25&#160;UTC | SHAKEN | 26&#160;Sep&#160;23&#160;00:25&#160;UTC | true | [view](CERTS/158c61d4da5ed8d0aaef7582d9a821a358a22ec2f241d449ecdab2e35ccf2fad/README.md) |
| 30&#160;Oct&#160;20&#160;14:39&#160;UTC | SHAKEN | 31&#160;Oct&#160;23&#160;14:39&#160;UTC | true | [view](CERTS/cc1f9403c74042b640c6d9b9727431d381822937a2b623238e1ebae19e7e4561/README.md) |
| 19&#160;Nov&#160;20&#160;11:15&#160;UTC | SHAKEN | 20&#160;Nov&#160;23&#160;11:15&#160;UTC | true | [view](CERTS/12c2c2592b30f8c808a43c947be48b24f59accedc97c6f78f687139d4ed8f70c/README.md) |
| 21&#160;Jan&#160;21&#160;18:49&#160;UTC | BigRiverTelephoneCompany | 22&#160;Jan&#160;24&#160;18:49&#160;UTC | true | [view](CERTS/9efdf79a4d93cef188ca8bc8dd4ce626a21813ed665bf108d74ff1f0158caf24/README.md) |
| 16&#160;Feb&#160;21&#160;15:12&#160;UTC | Logix-Keystore | 17&#160;Feb&#160;24&#160;15:12&#160;UTC | true | [view](CERTS/c53c156a26a6dc23f61a04e882c7643e9285b94374d9e82df16623f9519bdd65/README.md) |
| 21&#160;Feb&#160;21&#160;13:57&#160;UTC | Shaken | 22&#160;Feb&#160;24&#160;13:57&#160;UTC | true | [view](CERTS/58f3d6011c5d96eac83256a4c25ea4a085137b7bd6d4ba1acc246e4098c25ff9/README.md) |
| 11&#160;Mar&#160;21&#160;18:18&#160;UTC | SHAKEN-6744 | 11&#160;Mar&#160;24&#160;18:18&#160;UTC | true | [view](CERTS/414671d6f2e7beffdd958279b4cb2e705c5ee59f107aa1fb7b2a06008ae117b6/README.md) |
| 12&#160;Mar&#160;21&#160;14:58&#160;UTC | Flowroute | 12&#160;Mar&#160;24&#160;14:58&#160;UTC | true | [view](CERTS/ff098de1abec5c5569a38f6ad950c3e478009ddef15d3e456b0884ffaceb0cec/README.md) |
| 24&#160;Mar&#160;21&#160;14:02&#160;UTC | Firstcomm5917 | 24&#160;Mar&#160;24&#160;14:02&#160;UTC | true | [view](CERTS/394a79b142d1ac314cfabcd5009dc93fff81c1e19a5d14fe5175ae197ab5c66f/README.md) |
| 26&#160;Mar&#160;21&#160;12:05&#160;UTC | SHAKEN-925C | 26&#160;Mar&#160;24&#160;12:05&#160;UTC | true | [view](CERTS/497fc9dff6a178de5b6c33f6181ab092e273d03797a2026562808de76aafd623/README.md) |
| 08&#160;Apr&#160;21&#160;09:47&#160;UTC | Orange | 08&#160;Apr&#160;24&#160;09:47&#160;UTC | true | [view](CERTS/ff8031532cb64b24c610ce0c4887726fb4d2b915322e0dac4e004f9f51ac118a/README.md) |
| 13&#160;Apr&#160;21&#160;13:12&#160;UTC | SHAKEN-406H | 13&#160;Apr&#160;24&#160;13:12&#160;UTC | true | [view](CERTS/4d557f4e82049588ee17381bcff4a39f2212833ff5f310964aa060d9663d5bb7/README.md) |
| 16&#160;Apr&#160;21&#160;20:26&#160;UTC | Five9 | 16&#160;Apr&#160;24&#160;20:26&#160;UTC | true | [view](CERTS/08087af11d90b5b7bd3519ce065262e6ec44ea1dcfbc1db3995cc1b0a08fd48e/README.md) |
| 03&#160;May&#160;21&#160;21:11&#160;UTC | WindstreamCommunication | 03&#160;May&#160;24&#160;21:11&#160;UTC | true | [view](CERTS/8626a6fbbb1f434721fe9a149a77d13a6ace0da4578904229794070248323973/README.md) |
| 04&#160;May&#160;21&#160;11:46&#160;UTC | intelepeer.com | 04&#160;May&#160;24&#160;11:46&#160;UTC | true | [view](CERTS/49c3c8001a4b6cd71ae4cc41b0e36f2c724c5e024e19d99b71956f563d14ab01/README.md) |
| 06&#160;May&#160;21&#160;15:50&#160;UTC | SHAKEN-8468 | 06&#160;May&#160;24&#160;15:50&#160;UTC | true | [view](CERTS/369607701f40ce537f70ed798778773367baf6683606e11b81b8eae0eb85c6ad/README.md) |
| 07&#160;May&#160;21&#160;14:09&#160;UTC | SHAKEN_0377 | 07&#160;May&#160;24&#160;14:09&#160;UTC | true | [view](CERTS/a891efce3a4b35ed5111a2f8d9dbf2abef04a0e5c9ce4ea5313c0ede3e92ce26/README.md) |
| 07&#160;May&#160;21&#160;16:55&#160;UTC | PRD | 07&#160;May&#160;24&#160;16:55&#160;UTC | true | [view](CERTS/d81e85538a8dc39d06fa99ded608ea70df6631e79431ec10fc7b4173cee3b991/README.md) |
| 18&#160;May&#160;21&#160;15:07&#160;UTC | Granite | 18&#160;May&#160;24&#160;15:07&#160;UTC | true | [view](CERTS/42e6a55be823bf374996654d48501a9c98c1b37cdd523a496d9ed9a08044b7cf/README.md) |
| 19&#160;May&#160;21&#160;14:02&#160;UTC | 846B | 19&#160;May&#160;24&#160;14:02&#160;UTC | true | [view](CERTS/91646ad947e19e7b3501103e2d65a5e66ab7bc806c8bb7d4e87cce23ef668183/README.md) |
| 21&#160;May&#160;21&#160;16:38&#160;UTC | ReInvent | 21&#160;May&#160;24&#160;16:38&#160;UTC | true | [view](CERTS/ffb808b6e361219b3f20bc740eae385eee965cd38ad8b7b4e425a0ed25c611ee/README.md) |
| 22&#160;May&#160;21&#160;07:35&#160;UTC | SHAKEN 3201 | 22&#160;May&#160;24&#160;07:35&#160;UTC | true | [view](CERTS/a874f39ebd9b44c0c93f1f99e8faccef5aa0821d1cd0ff48c0a4bc2141b50c48/README.md) |
| 08&#160;Jun&#160;22&#160;13:35&#160;UTC | SHAKEN 763H | 08&#160;Jun&#160;23&#160;13:35&#160;UTC | true | [view](CERTS/139b7bad7bcb2d61b922fe1c7bcb9b21da4bf45df4334ab1e2658a45c4ab6962/README.md) |
| 08&#160;Jun&#160;22&#160;16:50&#160;UTC | SHAKEN 819J | 08&#160;Jun&#160;23&#160;16:50&#160;UTC | true | [view](CERTS/127f660276cb1085f74a7f9955b94ced3bd3452a57c26a8240c31fd6d29ee512/README.md) |
| 12&#160;Jun&#160;22&#160;15:12&#160;UTC | SHAKEN 261H | 12&#160;Jun&#160;23&#160;15:12&#160;UTC | true | [view](CERTS/6815b252b41b1d162968ea43df4ae48355a8b786816ec2434d65c703e2a5a8fc/README.md) |
| 12&#160;Jun&#160;22&#160;15:18&#160;UTC | SHAKEN 033H | 12&#160;Jun&#160;23&#160;15:18&#160;UTC | true | [view](CERTS/dbc2c40b64af4f892b4991d76c182bc85345235069d1cadf854bd2ef5d7fee63/README.md) |
| 14&#160;Jun&#160;22&#160;18:44&#160;UTC | SHAKEN 1980 | 14&#160;Jun&#160;23&#160;18:44&#160;UTC | true | [view](CERTS/8e0869be1ccc51bedfa80138e0e83d7d98ac458f7f444842da0de0c6567ce407/README.md) |
| 16&#160;Jun&#160;22&#160;18:26&#160;UTC | SHAKEN 939H | 16&#160;Jun&#160;23&#160;18:26&#160;UTC | true | [view](CERTS/a9ff363f03fba28c1972002e99200ecd750e4aaf59fa7302e39b8777be615e08/README.md) |
| 17&#160;Jun&#160;22&#160;21:25&#160;UTC | SHAKEN 573J | 17&#160;Jun&#160;23&#160;21:25&#160;UTC | true | [view](CERTS/235085bee8917aacad49680ef5f7e11ff56192cf5239041332dcb1d0f858cb2e/README.md) |
| 21&#160;Jun&#160;22&#160;18:39&#160;UTC | SHAKEN 712J | 21&#160;Jun&#160;23&#160;18:39&#160;UTC | true | [view](CERTS/869b4bd6e2405b41153c7a848f5585a914e488fab44fb7ebb556f91c2ab07df4/README.md) |
| 21&#160;Jun&#160;22&#160;21:04&#160;UTC | SHAKEN 611J | 21&#160;Jun&#160;23&#160;21:04&#160;UTC | true | [view](CERTS/3f87c16b475eb0fc6235d4e73fe5f990364d4c2f109cceac951e5e8b449fcd56/README.md) |
| 22&#160;Jun&#160;22&#160;20:48&#160;UTC | SHAKEN 128K | 22&#160;Jun&#160;23&#160;20:48&#160;UTC | true | [view](CERTS/b0bc8688e5e7156322c04d7a96cfa948e6e1598db33f9a978794a1a324c29b6c/README.md) |
| 22&#160;Jun&#160;22&#160;20:56&#160;UTC | SHAKEN 755B | 22&#160;Jun&#160;23&#160;20:56&#160;UTC | true | [view](CERTS/a3bc5e4345eb8b1fef269a76079585eac603aaec0941135cbd653b503bddd642/README.md) |
| 23&#160;Jun&#160;22&#160;15:04&#160;UTC | SHAKEN 558a | 23&#160;Jun&#160;23&#160;15:04&#160;UTC | true | [view](CERTS/9546e877a2d1dde91b3eb3c9aa90c7b5aa2b3ae28a18d48c0b24587449c4a9e2/README.md) |
| 23&#160;Jun&#160;22&#160;16:36&#160;UTC | SHAKEN 553J | 23&#160;Jun&#160;23&#160;16:36&#160;UTC | true | [view](CERTS/a804cf0f4ff52fb03a1d23618ecabb5eda1527649e3347ca6a146b22cd024d3b/README.md) |
| 23&#160;Jun&#160;22&#160;18:09&#160;UTC | SHAKEN 171K | 23&#160;Jun&#160;23&#160;18:09&#160;UTC | true | [view](CERTS/e9355468f55b838570f9c50ac0abbeaa64f565b5a37c9653fbb07b59689d1771/README.md) |
| 24&#160;Jun&#160;22&#160;14:05&#160;UTC | SHAKEN 739J | 24&#160;Jun&#160;23&#160;14:05&#160;UTC | true | [view](CERTS/61bb2e5f8faffe66e9d319704c906bb3c211b31d62ef605c738f231ee3b305aa/README.md) |
| 24&#160;Jun&#160;22&#160;19:58&#160;UTC | SHAKEN 067K | 24&#160;Jun&#160;23&#160;19:58&#160;UTC | true | [view](CERTS/aed8214f2128df08605bba6f6596ee4fc2874f44577591a6e254c7117b819843/README.md) |
| 28&#160;Jun&#160;22&#160;20:15&#160;UTC | SHAKEN 738J | 28&#160;Jun&#160;23&#160;20:15&#160;UTC | true | [view](CERTS/4fc19fc620e600a7a819b80c3cc596dc0139a567ae19484472728564c8423574/README.md) |
| 28&#160;Jun&#160;22&#160;21:58&#160;UTC | SHAKEN 743J | 28&#160;Jun&#160;23&#160;21:58&#160;UTC | true | [view](CERTS/afecedf9de456e4170f9f52742f42bdca9437dcabb6867c7c71372f8ba4feddf/README.md) |
| 29&#160;Jun&#160;22&#160;22:00&#160;UTC | SHAKEN 049K | 29&#160;Jun&#160;23&#160;22:00&#160;UTC | true | [view](CERTS/37d491d2756716824cbf4707b31a034fef6584e4e6d95d327a7428e68e92a02f/README.md) |
| 01&#160;Jul&#160;22&#160;16:21&#160;UTC | SHAKEN 139K | 01&#160;Jul&#160;23&#160;16:21&#160;UTC | true | [view](CERTS/eacdef99ac79c2ae3e713ec4c5972e785548050a35db9747f21eb2d4bc06c31b/README.md) |
| 05&#160;Jul&#160;22&#160;23:04&#160;UTC | SHAKEN 0348 | 05&#160;Jul&#160;23&#160;23:04&#160;UTC | true | [view](CERTS/d444d828608427c274287e7e6f5f1ccaa085e5310ab384e9fa153ff0f8fffaf1/README.md) |
| 08&#160;Jul&#160;22&#160;17:49&#160;UTC | SHAKEN 709J | 08&#160;Jul&#160;23&#160;17:49&#160;UTC | true | [view](CERTS/a371d70a1b505735b513c7488d811eb869c801aa246cfc1424db1e373b579f34/README.md) |
| 12&#160;Jul&#160;22&#160;19:50&#160;UTC | SHAKEN 938H | 12&#160;Jul&#160;23&#160;19:50&#160;UTC | true | [view](CERTS/a983cc83662a6d0d80eea559049b0f897f84d5398b7a5ea3480a431e89cf4f23/README.md) |
| 13&#160;Jul&#160;22&#160;17:30&#160;UTC | SHAKEN 962J | 13&#160;Jul&#160;23&#160;17:30&#160;UTC | true | [view](CERTS/cd04dcf7d5eb1acdb0c0c3e60fb3977707961e4240f1dfcf67eff0338bd75ffb/README.md) |
| 13&#160;Jul&#160;22&#160;19:10&#160;UTC | SHAKEN 771J | 13&#160;Jul&#160;23&#160;19:10&#160;UTC | true | [view](CERTS/0d7b51533df7b19ce84f11cf53ff9c2260e112501402635565d154d9e9741d63/README.md) |
| 14&#160;Jul&#160;22&#160;17:34&#160;UTC | SHAKEN 697J | 14&#160;Jul&#160;23&#160;17:34&#160;UTC | true | [view](CERTS/f420af1ed448cd7a6bc1f7caf9abafeac0ae73600acc0c5f8f32b5c0aac4d56a/README.md) |
| 18&#160;Jul&#160;22&#160;17:33&#160;UTC | SHAKEN 715J | 18&#160;Jul&#160;23&#160;17:33&#160;UTC | true | [view](CERTS/450172823b48d29255bde6a9ec9e921f8e275b76c1a33986f23fa9d410802eac/README.md) |
| 28&#160;Jul&#160;22&#160;20:56&#160;UTC | SHAKEN 074K | 28&#160;Jul&#160;23&#160;20:56&#160;UTC | true | [view](CERTS/ce2299b32a7f2749baeb8864b84a8b531591a05a0c79c3bff8af7b8da441d513/README.md) |
| 08&#160;Aug&#160;22&#160;12:58&#160;UTC | SHAKEN 150K | 08&#160;Aug&#160;23&#160;12:58&#160;UTC | true | [view](CERTS/5cf325af20a32e0691ff11a47485c841d9dd35771bbe2fe8027deb13e70b2a81/README.md) |
| 08&#160;Aug&#160;22&#160;14:30&#160;UTC | SHAKEN 710A | 08&#160;Aug&#160;23&#160;14:30&#160;UTC | true | [view](CERTS/e97dffcbf85f9183ac12212e1a9e342f76ffd96d6c6d280854fb31e9465d481a/README.md) |
| 08&#160;Aug&#160;22&#160;18:38&#160;UTC | SHAKEN 5493 | 08&#160;Aug&#160;23&#160;18:38&#160;UTC | true | [view](CERTS/98f0f154eace54f5c146b3c711c40b116a574ce3596ab011537989ed5d0076c0/README.md) |
| 16&#160;Aug&#160;22&#160;20:07&#160;UTC | SHAKEN 295K | 16&#160;Aug&#160;23&#160;20:07&#160;UTC | true | [view](CERTS/1a821ccda1ba38a8ed98fbffbd79e4158497c99c35ed220d0110fcb080f9a933/README.md) |
| 18&#160;Aug&#160;22&#160;18:07&#160;UTC | SHAKEN 219K | 18&#160;Aug&#160;23&#160;18:07&#160;UTC | true | [view](CERTS/fc1a6306ba8c8e009e014efbcf859e5740cae4ebfac67e9628ef65e16209c9b1/README.md) |
| 22&#160;Aug&#160;22&#160;20:06&#160;UTC | SHAKEN 0347 | 22&#160;Aug&#160;23&#160;20:06&#160;UTC | true | [view](CERTS/5177a968576a738d0c1a28c4fcd35b626d2d75b890609ab930dcd671f3263dc6/README.md) |
| 29&#160;Aug&#160;22&#160;19:07&#160;UTC | ATT SHAKEN 4036 | 29&#160;Aug&#160;23&#160;19:07&#160;UTC | true | [view](CERTS/3a8d4d5fe47e784f925bca30b21d44f3492ad0813fd074ef2d438cca1d4acc68/README.md) |
| 31&#160;Aug&#160;22&#160;14:47&#160;UTC | SHAKEN 500J | 31&#160;Aug&#160;23&#160;14:47&#160;UTC | true | [view](CERTS/25393ddb6f4b6a42df56e2cafbff0fa710e60bbdf94d37561290cdb6e66f28d1/README.md) |
| 01&#160;Sep&#160;22&#160;20:43&#160;UTC | SHAKEN 813J | 01&#160;Sep&#160;23&#160;20:43&#160;UTC | true | [view](CERTS/5ec454f90ed7dffae1782a1028542871a5273b075f6855a9a7d0171dbdb78750/README.md) |
| 07&#160;Sep&#160;22&#160;17:53&#160;UTC | SHAKEN 872J | 07&#160;Sep&#160;23&#160;17:53&#160;UTC | true | [view](CERTS/4fdfa813115c43e28c707b69143e7b1b52e053494886c9e95fa3bd6d933d6438/README.md) |
| 09&#160;Sep&#160;22&#160;14:44&#160;UTC | SHAKEN 5606 | 09&#160;Sep&#160;23&#160;14:44&#160;UTC | true | [view](CERTS/8fc7d03e6cd7ab01c7c8b9051bfef91cd5bceecefd08f74b6fc948b65b15eca4/README.md) |
| 12&#160;Sep&#160;22&#160;19:52&#160;UTC | SHAKEN 707J | 12&#160;Sep&#160;23&#160;19:52&#160;UTC | true | [view](CERTS/132d0f59c15814fc3f80760fdf109ee36d0ac19eaf03a74457523f2cf4dcf982/README.md) |
| 15&#160;Sep&#160;22&#160;16:20&#160;UTC | SHAKEN 292K | 15&#160;Sep&#160;23&#160;16:20&#160;UTC | true | [view](CERTS/e449581f068a3e747acc8dfd7c0707c9aed3deca9f572afa13bbb0ffbebffd27/README.md) |
| 20&#160;Oct&#160;22&#160;18:45&#160;UTC | SHAKEN 305K | 20&#160;Oct&#160;23&#160;18:45&#160;UTC | true | [view](CERTS/6aa8246878d796479404bf793c7c1172ccd764b2c057d5a5b87968e43adc05bc/README.md) |
| 26&#160;Oct&#160;22&#160;15:49&#160;UTC | SHAKEN 745J | 26&#160;Oct&#160;23&#160;15:49&#160;UTC | true | [view](CERTS/85bf5f426006bc4a831a744672f2e2f2a936f7c6cdda1104d3ee1e7dab7268f9/README.md) |
| 26&#160;Oct&#160;22&#160;16:36&#160;UTC | SHAKEN 770J | 26&#160;Oct&#160;23&#160;16:36&#160;UTC | true | [view](CERTS/c0d94a93ff7d9519ab97a0e8a44ecae02198e5df6ca89a8783163d00a80583ea/README.md) |
| 26&#160;Oct&#160;22&#160;16:40&#160;UTC | SHAKEN 7575 | 26&#160;Oct&#160;23&#160;16:40&#160;UTC | true | [view](CERTS/9166142f8c959390cf18b23d135eae9c02172c9af68943eb8c8e98dd63868be8/README.md) |
| 26&#160;Oct&#160;22&#160;20:35&#160;UTC | SHAKEN 545B | 26&#160;Oct&#160;23&#160;20:35&#160;UTC | true | [view](CERTS/d37563f111fb51b1871819e1ca4ebc9e239ebf6f4f8049b738f7cb655cb44ea2/README.md) |
| 28&#160;Oct&#160;22&#160;14:59&#160;UTC | prod SHAKEN 811J | 28&#160;Oct&#160;23&#160;14:59&#160;UTC | true | [view](CERTS/f6e3384410e423b0e0006618d97878154ef07b52fbdc31a7ae6f3bc7f102e6ed/README.md) |
| 07&#160;Nov&#160;22&#160;15:08&#160;UTC | SHAKEN 197D | 07&#160;Nov&#160;23&#160;15:08&#160;UTC | true | [view](CERTS/b6a79b74a53e5134cb3a141bacae54e828d7f113fcdd52a1a2f9da98091e0583/README.md) |
| 07&#160;Nov&#160;22&#160;17:05&#160;UTC | SHAKEN 7126 | 07&#160;Nov&#160;23&#160;17:05&#160;UTC | true | [view](CERTS/9d33c8d547510a84cadd2b8f3c3118687f90cda17539f9e502cf3e71c93fdcce/README.md) |
| 08&#160;Nov&#160;22&#160;22:26&#160;UTC | SHAKEN 775J | 08&#160;Nov&#160;23&#160;22:26&#160;UTC | true | [view](CERTS/d5eb605cc1d3b7bc055c96732f267e1a74d7d4221627e5b7855686dac442766c/README.md) |
| 14&#160;Nov&#160;22&#160;14:53&#160;UTC | SHAKEN 333K | 14&#160;Nov&#160;23&#160;14:53&#160;UTC | true | [view](CERTS/951e91548d95354bc7579ee5c7d6766bdb399e4f411ddbb2e734d96b114f231c/README.md) |
| 21&#160;Nov&#160;22&#160;16:32&#160;UTC | SHAKEN 506J | 21&#160;Nov&#160;23&#160;16:32&#160;UTC | true | [view](CERTS/46377f0dc1b28492db49c31a70f5473c128622706287d6a94236679b0cddde14/README.md) |
| 30&#160;Nov&#160;22&#160;21:59&#160;UTC | SHAKEN 403K | 30&#160;Nov&#160;23&#160;21:59&#160;UTC | true | [view](CERTS/814488f1c616fe053f7675faefaebf8cf5d258eb5373b518a6e80abdea6c9eac/README.md) |
| 02&#160;Dec&#160;22&#160;15:44&#160;UTC | SHAKEN 2116 | 02&#160;Dec&#160;23&#160;15:44&#160;UTC | true | [view](CERTS/481efee60504ed3a1517d90888e132068069ee829d3af086cdb4019da592f3f2/README.md) |
| 06&#160;Dec&#160;22&#160;19:43&#160;UTC | SHAKEN 063E | 06&#160;Dec&#160;23&#160;19:43&#160;UTC | true | [view](CERTS/ac16d79be6e3e9a2256c334441599660b115071c7dd86269d9dc3b6de72db9f4/README.md) |
| 09&#160;Dec&#160;22&#160;14:48&#160;UTC | SHAKEN 899J | 09&#160;Dec&#160;23&#160;14:48&#160;UTC | true | [view](CERTS/5a4ff0a70a41cd82a090a2153380b441113444ccc15c2312a25119f01f775b09/README.md) |
| 15&#160;Dec&#160;22&#160;16:00&#160;UTC | SHAKEN 2473 | 15&#160;Dec&#160;23&#160;16:00&#160;UTC | true | [view](CERTS/dfda336ef5f289a46d60436f1cfaa52f5699b9895353a632a58e72aa4249c2be/README.md) |
| 19&#160;Dec&#160;22&#160;19:06&#160;UTC | SHAKEN 750J | 19&#160;Dec&#160;23&#160;19:06&#160;UTC | true | [view](CERTS/86dd3c7321faf962cc02f0baed6726223fd49e92dc19ce123eab134fcd534fdf/README.md) |
| 19&#160;Dec&#160;22&#160;19:17&#160;UTC | SHAKEN 7076 | 19&#160;Dec&#160;23&#160;19:17&#160;UTC | true | [view](CERTS/97ad9a2a46b3df41969a7b8ed480d7e858811350a835fd4cbf3a9286e1cbea96/README.md) |
| 04&#160;Jan&#160;23&#160;15:48&#160;UTC | SHAKEN 847J | 04&#160;Jan&#160;24&#160;15:48&#160;UTC | true | [view](CERTS/55944a7a6a5eb8c12ffea55e8f49762d5b5e8c316465dc07cb8cd94eaacc595b/README.md) |
| 27&#160;Jan&#160;23&#160;15:34&#160;UTC | SHAKEN 0377 | 27&#160;Jan&#160;24&#160;15:34&#160;UTC | true | [view](CERTS/b498cc31dae3ec27255101496ba26349ac318df5db5d4301267106469a969c40/README.md) |
| 01&#160;Feb&#160;23&#160;17:56&#160;UTC | SHAKEN 701J | 01&#160;Feb&#160;24&#160;17:56&#160;UTC | true | [view](CERTS/1915a0da7d3d06216e52ea2fa8fe13bda8f474476121f3ec13e61505dbb81b14/README.md) |
| 01&#160;Feb&#160;23&#160;20:15&#160;UTC | SHAKEN 917J | 01&#160;Feb&#160;24&#160;20:15&#160;UTC | true | [view](CERTS/a0e94d69d78f415a65335a51c5346d9434f81fcfda9767af5a43317163366172/README.md) |
| 01&#160;Feb&#160;23&#160;20:36&#160;UTC | SHAKEN 804J | 01&#160;Feb&#160;24&#160;20:36&#160;UTC | true | [view](CERTS/1441f59ba3cc930468486aa57cf974990dcaa9818bb7ef6c4b0384de7cd1538e/README.md) |
| 01&#160;Feb&#160;23&#160;21:47&#160;UTC | SHAKEN 5447 | 01&#160;Feb&#160;24&#160;21:47&#160;UTC | true | [view](CERTS/408629c7d0b644e0b75c87f2929e40c3d8fcaeb2340acf542b79a5c261613793/README.md) |
| 02&#160;Feb&#160;23&#160;16:44&#160;UTC | SHAKEN 393J | 02&#160;Feb&#160;24&#160;16:44&#160;UTC | true | [view](CERTS/5dc06b4dae770d7bbc3dcb527265bd045a31800cad333205abe6d55231e52e5c/README.md) |
| 06&#160;Feb&#160;23&#160;15:28&#160;UTC | SHAKEN 863J | 06&#160;Feb&#160;24&#160;15:28&#160;UTC | true | [view](CERTS/d828c0e0f959db420a930264a3a30051f6a8a457ed717183fe3061af78ea8172/README.md) |
| 09&#160;Feb&#160;23&#160;19:36&#160;UTC | prod SHAKEN 811J | 09&#160;Feb&#160;24&#160;19:36&#160;UTC | true | [view](CERTS/f3a0e7715828bc30d8214990da4c1e53ec1e600cbbf90784ef5dd9cbe3f57dfd/README.md) |
| 09&#160;Feb&#160;23&#160;20:10&#160;UTC | SHAKEN 597F | 09&#160;Feb&#160;24&#160;20:10&#160;UTC | true | [view](CERTS/60f18673867347c3986fd3a6e45e07dc85f1fc2a0137b836d3633e3dc16c8bf0/README.md) |
| 14&#160;Feb&#160;23&#160;13:24&#160;UTC | SHAKEN 886G | 14&#160;Feb&#160;24&#160;13:24&#160;UTC | true | [view](CERTS/ba0f7420cb1ee3a870908f4c5e0d25b909850f97b6a45b948a220e375f728c45/README.md) |
| 15&#160;Feb&#160;23&#160;17:06&#160;UTC | SHAKEN 1591 | 15&#160;Feb&#160;24&#160;17:06&#160;UTC | true | [view](CERTS/561deee1741ce86b0515ba80b91a7e4737183b27ab3b96fe1a38e1de641d5f90/README.md) |
| 16&#160;Feb&#160;23&#160;03:36&#160;UTC | SHAKEN 280K | 16&#160;Feb&#160;24&#160;03:36&#160;UTC | true | [view](CERTS/90b3cb13d9c2d940a389b6352d835e5adf1351b4e2226f913eece640ce364e83/README.md) |
| 16&#160;Feb&#160;23&#160;22:05&#160;UTC | SHAKEN 8943 | 16&#160;Feb&#160;24&#160;22:05&#160;UTC | true | [view](CERTS/7f251a9bc5f0be8eff3bd6bcf7ea6dc76765a5a56d778b66badf291cad1456df/README.md) |
| 22&#160;Feb&#160;23&#160;21:42&#160;UTC | SHAKEN 844J | 22&#160;Feb&#160;24&#160;21:42&#160;UTC | true | [view](CERTS/a7ef3c96a24628283c98e8919901d3759cd519f93713c5442225c81bccbd60eb/README.md) |
| 27&#160;Feb&#160;23&#160;15:34&#160;UTC | SHAKEN 7661 | 27&#160;Feb&#160;24&#160;15:34&#160;UTC | true | [view](CERTS/a7a87eb4bad4575a2fc6f5ad1da53ad1523a42b9f64746627ad4ed1119c0ed71/README.md) |
| 03&#160;Mar&#160;23&#160;17:51&#160;UTC | SHAKEN 963J | 02&#160;Mar&#160;24&#160;17:51&#160;UTC | true | [view](CERTS/a8265b81ba231916107001a33ad513830f6da031bb4ebc416aa647da8465faf0/README.md) |
| 03&#160;Mar&#160;23&#160;19:30&#160;UTC | SHAKEN 966J | 02&#160;Mar&#160;24&#160;19:30&#160;UTC | true | [view](CERTS/2a97197c7b1a48a343b0d3d74b034e1c4ae13d5f3f9b1d6e0d9d073be9af9810/README.md) |
| 08&#160;Mar&#160;23&#160;22:03&#160;UTC | SHAKEN 1563 | 07&#160;Mar&#160;24&#160;22:03&#160;UTC | true | [view](CERTS/32dec65d5504434aabf351b6c4de564ea0fe6d09cdc57509b2e3969909266b85/README.md) |
| 09&#160;Mar&#160;23&#160;22:19&#160;UTC | SHAKEN 468B | 08&#160;Mar&#160;24&#160;22:19&#160;UTC | true | [view](CERTS/e238e78eef9a711db54ac07c749a2e1f7aeb0875742d3ef2c904eb3ae4f9932d/README.md) |
| 13&#160;Mar&#160;23&#160;01:25&#160;UTC | SHAKEN 418c | 12&#160;Mar&#160;24&#160;01:25&#160;UTC | true | [view](CERTS/9504d166a1cd058cdc2de368a6296780f8605c9bad6f435fe733269f56333e3b/README.md) |
| 15&#160;Mar&#160;23&#160;21:16&#160;UTC | SHAKEN 030J | 14&#160;Mar&#160;24&#160;21:16&#160;UTC | true | [view](CERTS/2054a0a6de6024f2441d09f8cf8435ec7956a836444341a21a8180cc3d9470e8/README.md) |
| 21&#160;Mar&#160;23&#160;14:48&#160;UTC | SHAKEN 973J | 20&#160;Mar&#160;24&#160;14:48&#160;UTC | true | [view](CERTS/a297461669b36df06235c621e0e8edb01134b482355a8b1d2ca65ae3f039f6cf/README.md) |
| 21&#160;Mar&#160;23&#160;17:01&#160;UTC | SHAKEN 951J | 20&#160;Mar&#160;24&#160;17:01&#160;UTC | true | [view](CERTS/364eb4b1c2247d7a676f8e833c74c1e023847f1c81b8db0135aee4b2b439a936/README.md) |
| 06&#160;Apr&#160;23&#160;16:04&#160;UTC | SHAKEN 854J | 05&#160;Apr&#160;24&#160;16:04&#160;UTC | true | [view](CERTS/fe65b2e6e8aa1910bbf70471e2a7c378afe31e6b8b13d73465713fdcae273b25/README.md) |
| 06&#160;Apr&#160;23&#160;17:17&#160;UTC | SHAKEN 732J | 05&#160;Apr&#160;24&#160;17:17&#160;UTC | true | [view](CERTS/887852d15a88ff20647d28f8a482a72fd62e8fb60762d234ac16d3789fbc7458/README.md) |
| 06&#160;Apr&#160;23&#160;19:13&#160;UTC | SHAKEN 037K | 05&#160;Apr&#160;24&#160;19:13&#160;UTC | true | [view](CERTS/9eda285c12d8be336d26f3bf6e3fa56437a792d4bf6de0cd8ecd13296548e627/README.md) |
| 07&#160;Apr&#160;23&#160;14:55&#160;UTC | SHAKEN 402E | 06&#160;Apr&#160;24&#160;14:55&#160;UTC | true | [view](CERTS/1bbe8d208a3e41488f4962ece833675968d3964febabff2ff87bda7b0d85b9bc/README.md) |
| 07&#160;Apr&#160;23&#160;20:40&#160;UTC | SHAKEN 502E | 06&#160;Apr&#160;24&#160;20:40&#160;UTC | true | [view](CERTS/1273ffd08a646d060eb8a642da40518417e565144f635f1065c21dc56ed64d3e/README.md) |
| 07&#160;Apr&#160;23&#160;20:57&#160;UTC | SHAKEN 4427 | 06&#160;Apr&#160;24&#160;20:57&#160;UTC | true | [view](CERTS/a7e2daed5a72b699ce9101720b3f59a7ebcd6e9a21adc86d260c7a8b1731544f/README.md) |
| 07&#160;Apr&#160;23&#160;21:23&#160;UTC | SHAKEN 473G | 06&#160;Apr&#160;24&#160;21:23&#160;UTC | true | [view](CERTS/a33f37c5acd4b1b1602c4fcac0d0fcef66020c0f88f5745e44ddcc202a002f11/README.md) |
| 14&#160;Apr&#160;23&#160;19:27&#160;UTC | SHAKEN 525K | 13&#160;Apr&#160;24&#160;19:27&#160;UTC | true | [view](CERTS/3e9c61f0e5e87d7f60638dfe2e50fa27220c1370d07cc1cdea317b1c18aec2d0/README.md) |
| 20&#160;Apr&#160;23&#160;17:29&#160;UTC | SHAKEN 772J | 19&#160;Apr&#160;24&#160;17:29&#160;UTC | true | [view](CERTS/ed40b11900cf117dd4a3be5497716f782100263f7bb1d8fd58e352532e01ba6e/README.md) |
| 24&#160;Apr&#160;23&#160;17:54&#160;UTC | SHAKEN 821J | 23&#160;Apr&#160;24&#160;17:54&#160;UTC | true | [view](CERTS/707474b972584be7301d9b1ca170fe2e37ea80c92e09451843f66f0df60a27b6/README.md) |
| 24&#160;Apr&#160;23&#160;22:08&#160;UTC | SHAKEN 1845 | 23&#160;Apr&#160;24&#160;22:08&#160;UTC | true | [view](CERTS/ca030d5bd1de5251e02f96fa14994a4590277ddd571e2b7c2d546b31cb6766cd/README.md) |
| 26&#160;Apr&#160;23&#160;21:17&#160;UTC | SHAKEN 704J | 25&#160;Apr&#160;24&#160;21:17&#160;UTC | true | [view](CERTS/daa06df334b71280bbef4d2f7b1a528b3bc38be2a7ef4919ae2ad15bc12d14bc/README.md) |
| 28&#160;Apr&#160;23&#160;00:57&#160;UTC | SHAKEN 030K | 27&#160;Apr&#160;24&#160;00:57&#160;UTC | true | [view](CERTS/be9fe29ce5f2f08cdf3c15e19a42b592c67e6dcd593de67e3012d041c1c1f2f2/README.md) |
| 28&#160;Apr&#160;23&#160;22:27&#160;UTC | SHAKEN 782J | 27&#160;Apr&#160;24&#160;22:27&#160;UTC | true | [view](CERTS/8408053ed34c756b214c00f5cdde40eee1cceaf48a094e277f4a17f338a7bec3/README.md) |
| 02&#160;May&#160;23&#160;01:29&#160;UTC | SHAKEN 0523 | 01&#160;May&#160;24&#160;01:29&#160;UTC | true | [view](CERTS/1a8ae29a6c889b27b08f018046cde356191926e2da2f4b0b877832d5230f4d86/README.md) |
| 05&#160;May&#160;23&#160;18:11&#160;UTC | SHAKEN 534J | 04&#160;May&#160;24&#160;18:11&#160;UTC | true | [view](CERTS/9787db455edca00a14e9b536630532d269bc45ef79a70c58139de716a2a06037/README.md) |
| 09&#160;May&#160;23&#160;21:14&#160;UTC | SHAKEN 235K | 08&#160;May&#160;24&#160;21:14&#160;UTC | true | [view](CERTS/5934081781ff92190a826b8bbe25a78439eed9b2c460cb3bd6c64a538a07849f/README.md) |
| 12&#160;May&#160;23&#160;14:08&#160;UTC | SHAKEN 3130 | 11&#160;May&#160;24&#160;14:08&#160;UTC | true | [view](CERTS/24810723d9acfe684ce2bbc282b47b39f788047d417548217d31034bc42509ef/README.md) |
| 12&#160;May&#160;23&#160;14:58&#160;UTC | SHAKEN 767J | 11&#160;May&#160;24&#160;14:58&#160;UTC | true | [view](CERTS/39d08bb3239916fa0f20776d3ce9ac3d4b2b3eb5a8a2da983c617dcdefdca68d/README.md) |
| 12&#160;May&#160;23&#160;18:43&#160;UTC | SHAKEN 113K | 11&#160;May&#160;24&#160;18:43&#160;UTC | true | [view](CERTS/116ac219ba17db6bc35b290e21320e280612cd74d359d5fd920749a6596fdee0/README.md) |
| 12&#160;May&#160;23&#160;18:57&#160;UTC | SHAKEN 743J | 11&#160;May&#160;24&#160;18:57&#160;UTC | true | [view](CERTS/0d177e2c26f5198a22e5d0ef59ab49aa741d7edfb2d5260b919ec93d753d1ec0/README.md) |
| 12&#160;May&#160;23&#160;19:09&#160;UTC | SHAKEN 254H | 11&#160;May&#160;24&#160;19:09&#160;UTC | true | [view](CERTS/d128e949b794dd790caf7b9f21f99226090c2031b2918d1d3fbd589fc99d8763/README.md) |
| 13&#160;May&#160;23&#160;12:51&#160;UTC | SHAKEN 669B | 12&#160;May&#160;24&#160;12:51&#160;UTC | true | [view](CERTS/05e88739872a0741ae16ed4b98a60d0063fb30a6d6809512737d5b9499cbfa44/README.md) |
| 17&#160;May&#160;23&#160;16:18&#160;UTC | SHAKEN 845J | 16&#160;May&#160;24&#160;16:18&#160;UTC | true | [view](CERTS/dd6bf31a74a21b57fb42303fadc10c7d4a6b87c15324c9f0589a4e2f3c5652be/README.md) |
| 17&#160;May&#160;23&#160;17:40&#160;UTC | SHAKEN 520F | 16&#160;May&#160;24&#160;17:40&#160;UTC | true | [view](CERTS/b6eaef4bce63e8c8069908790c34346614631f8b912d4bbfb2cc2f04f7fa641d/README.md) |
| 24&#160;May&#160;23&#160;15:36&#160;UTC | SHAKEN 625H | 23&#160;May&#160;24&#160;15:36&#160;UTC | true | [view](CERTS/54ddee80553a275fff824ef1a23b01d35e56bed5d450c37211046787146ead9e/README.md) |
| 27&#160;May&#160;23&#160;14:23&#160;UTC | SHAKEN 869J | 26&#160;May&#160;24&#160;14:23&#160;UTC | true | [view](CERTS/d827d78eee1794b02205387ff0e32270c9c38f3a1870cb691a9008b8f094683d/README.md) |
| 29&#160;May&#160;23&#160;12:58&#160;UTC | SHAKEN 700H | 28&#160;May&#160;24&#160;12:58&#160;UTC | true | [view](CERTS/ba3f12f3b39595223f248a0c0f5d9e7b87446ad80aac42a60cef62cd8aebfcb5/README.md) |
| 29&#160;May&#160;23&#160;13:14&#160;UTC | SHAKEN 7914 | 28&#160;May&#160;24&#160;13:14&#160;UTC | true | [view](CERTS/994ad1fdcf7b0400b848abf6ca8d399deeda8920ad4a79433f3f21a778da9d88/README.md) |
| 31&#160;May&#160;23&#160;20:56&#160;UTC | SHAKEN 554J | 30&#160;May&#160;24&#160;20:56&#160;UTC | true | [view](CERTS/b59e9d7ad3f57adc1c9c10094f8e7d1b032338eb86048e36a7cc41c9b8fb0be4/README.md) |

#### CA Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 23&#160;Sep&#160;19&#160;13:26&#160;UTC | Neustar Certified Caller ID Root CA | 23&#160;Sep&#160;39&#160;13:26&#160;UTC | true | [view](CERTS/4a77c17cd411cb0ff2984b97687f75ab1db451ac7b717ab81c931351c2d547a1/README.md) |
| 23&#160;Sep&#160;19&#160;13:32&#160;UTC | Neustar Certified Caller ID CA-1 | 23&#160;Sep&#160;29&#160;13:32&#160;UTC | true | [view](CERTS/dade1a52e76c29fc9af1e1221a2a6be02c9899a552d396580855935c9592733b/README.md) |
| 17&#160;Aug&#160;21&#160;17:19&#160;UTC | Neustar Certified Caller ID SHAKEN Root CA | 17&#160;Aug&#160;41&#160;17:19&#160;UTC | true | [view](CERTS/4c728d18b628cc67dda5490e0b4aa8ef4ba679f96d033f34f1680e219e0806c3/README.md) |
| 19&#160;Aug&#160;21&#160;03:25&#160;UTC | Neustar Certified Caller ID SHAKEN CA-1 | 20&#160;Aug&#160;31&#160;03:25&#160;UTC | true | [view](CERTS/b6dc9bf58a55979c78ad569a17c86a7f644721bd3ab2bcf99a27d13636900cf4/README.md) |
| 30&#160;Aug&#160;22&#160;06:39&#160;UTC | Neustar Certified Caller ID SHAKEN CA-2 | 30&#160;Aug&#160;32&#160;06:39&#160;UTC | true | [view](CERTS/3ea530838e9952fdda913a8bd669bf37f88f4ffdb39a34698f34a63915c9e404/README.md) |
| 05&#160;Oct&#160;22&#160;17:26&#160;UTC | Neustar Certified Caller ID SHAKEN CA-2 | 05&#160;Oct&#160;32&#160;17:26&#160;UTC | true | [view](CERTS/0bd95ecbb97c09de0df079ca41e10c360c4b5928ac56c496879a2c90c6bbffe4/README.md) |


Generated: 02 Jun 23 01:12 UTC