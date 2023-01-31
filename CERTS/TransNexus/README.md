# STIR/SHAKEN CA Ecosystem Compliance

## TransNexus

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 1298 certificates were included in the corpus being tested
- 1 certificates in the corpus were skipped because they are duplicates
- 1197 certificates in the corpus were skipped because they are expired
- 1 certificates in the corpus were skipped because they are not currently trusted
- 99 certificates being tested against the remaining rules
- 1.24 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 10.10% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 61 days is the average remaining validity for the certificates in the corpus
- 61 days is the average initial validity for the certificates in the corpus
- 82 certificates expire in the next 30 days
- 1.50 average number of unexpired certificates per OCN observed
- 66 unique OCNs observed in unexpired and valid certificate corpus

| Instances | Test | Source |
|-----------|------|--------|
| 2 | [e_atis_certificate_policies](ISSUES/e_atis_certificate_policies/README.md) | ATIS1000080 |
| 10 | [e_incorrect_ku_encoding](ISSUES/e_incorrect_ku_encoding/README.md) | RFC5280 |
| 2 | [e_us_cp_ambiguous_identifier](ISSUES/e_us_cp_ambiguous_identifier/README.md) | US_SHAKEN_CP |
| 99 | [e_us_cp_subject_sn](ISSUES/e_us_cp_subject_sn/README.md) | US_SHAKEN_CP |
| 10 | [w_shaken_subject_rdn_unknown](ISSUES/w_shaken_subject_rdn_unknown/README.md) | SHAKEN_PKI_BEST_PRACTICES |

#### CA Certificates

- 6 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 0 certificates in the corpus were skipped because they are expired
- 1 certificates in the corpus were skipped because they are not currently trusted
- 5 certificates being tested against the remaining rules
- 3.00 issues on average found in unexpired, trusted, and non-compliant certificates
- 60.00% of certificates contain one or more Error level issue
- 60.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 60.00% of certificates are too old to be assessed against currently enforced expectations
- 5462 days is the average remaining validity for the certificates in the corpus
- 5113 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days

| Instances | Test | Source |
|-----------|------|--------|
| 3 | [e_incorrect_ku_encoding](ISSUES/e_incorrect_ku_encoding/README.md) | RFC5280 |
| 3 | [e_us_cp_ca_key_usage_crl_sign](ISSUES/e_us_cp_ca_key_usage_crl_sign/README.md) | US_SHAKEN_CP |
| 3 | [w_shaken_ca_subject_rdn_unknown](ISSUES/w_shaken_ca_subject_rdn_unknown/README.md) | SHAKEN_PKI_BEST_PRACTICES |

### Details

#### Leaf Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 30&#160;Mar&#160;22&#160;12:54&#160;UTC | Fusion Connect SHAKEN 2720 | 30&#160;Mar&#160;23&#160;12:54&#160;UTC | true | [view](CERTS/03a010b294807e90cc1309d7466fd2ddc591158a69be6950ff4d0eab32de0799/README.md) |
| 27&#160;Apr&#160;22&#160;18:22&#160;UTC | CCI SHAKEN 663J | 27&#160;Apr&#160;23&#160;18:22&#160;UTC | true | [view](CERTS/9cfd0b9974acdc9b5fce3888cf2f6cf99173e7955adeebd37a1e60dfc444a0b3/README.md) |
| 10&#160;Jun&#160;22&#160;14:00&#160;UTC | SHAKEN 597J | 10&#160;Jun&#160;23&#160;14:00&#160;UTC | true | [view](CERTS/4baa838eb96ce358aa1565d0add8b75072258efd5d21a5bbb0eec878026a67be/README.md) |
| 20&#160;Jun&#160;22&#160;20:25&#160;UTC | SHAKEN 857J | 20&#160;Jun&#160;23&#160;20:25&#160;UTC | true | [view](CERTS/e1e0a4851fb7b51d617959c00fb6a17e7b23d34e1555ba70ab01c9182a2f85f8/README.md) |
| 22&#160;Jun&#160;22&#160;18:46&#160;UTC | SHAKEN 755J | 22&#160;Jun&#160;23&#160;18:46&#160;UTC | true | [view](CERTS/a8cac316fddf25b48b9350a325791be8b87e6484a5959ab605db1ae5e3e99a84/README.md) |
| 29&#160;Jun&#160;22&#160;20:24&#160;UTC | SHAKEN 736J | 29&#160;Jun&#160;23&#160;20:24&#160;UTC | true | [view](CERTS/71b25a480077816998e48e6a76f272e70aa91ecb19b0e773cd8777823cdead32/README.md) |
| 25&#160;Jul&#160;22&#160;18:36&#160;UTC | SHAKEN 578J | 25&#160;Jul&#160;23&#160;18:36&#160;UTC | true | [view](CERTS/f4af5a28131d680ef2682ae86181b64f0c919864d7dbaef00c0598d314b537a1/README.md) |
| 10&#160;Aug&#160;22&#160;18:11&#160;UTC | SHAKEN 073H | 10&#160;Aug&#160;23&#160;18:11&#160;UTC | true | [view](CERTS/87dabb91a689a6ac540b5ffbf5eb89326fada83fb822d2bb9ec31b68251b579e/README.md) |
| 23&#160;Aug&#160;22&#160;04:14&#160;UTC | SHAKEN 866J | 19&#160;Feb&#160;23&#160;04:14&#160;UTC | true | [view](CERTS/4c4fdb320c51c8582d595c2d307cd770d409daf533ff573bcc73afaed83f6b7d/README.md) |
| 13&#160;Sep&#160;22&#160;17:24&#160;UTC | SHAKEN 706J | 13&#160;Sep&#160;23&#160;17:24&#160;UTC | true | [view](CERTS/1eea5e17ae8b36cb5b6d657178273af745cea015e63586e6d44d1ba2ee8f722a/README.md) |
| 23&#160;Nov&#160;22&#160;17:31&#160;UTC | SHAKEN 4036 | 10&#160;May&#160;23&#160;17:31&#160;UTC | true | [view](CERTS/e03ff4259e6b456f1bb19d011f005932e83b0a35b520730c97cd77ad8868071c/README.md) |
| 07&#160;Dec&#160;22&#160;00:45&#160;UTC | SHAKEN 299K | 07&#160;Dec&#160;23&#160;00:45&#160;UTC | true | [view](CERTS/7b5e04ff01ace2ee384df3507886684204901bf23851866eb6b3891ecdd96a92/README.md) |
| 22&#160;Dec&#160;22&#160;17:30&#160;UTC | SHAKEN 505J | 22&#160;Mar&#160;23&#160;17:30&#160;UTC | true | [view](CERTS/7ff43923df80c3a7e2d854b42f33c69069af4625f5267df0d7ff7b2004ad1c1b/README.md) |
| 30&#160;Dec&#160;22&#160;06:30&#160;UTC | SHAKEN 807J | 28&#160;Feb&#160;23&#160;06:30&#160;UTC | true | [view](CERTS/62da6686ffc785eb4f15528d7bf06acd4d4778bbe258808e662ee91ffb868ccf/README.md) |
| 30&#160;Dec&#160;22&#160;12:34&#160;UTC | SHAKEN 193E | 28&#160;Feb&#160;23&#160;12:34&#160;UTC | true | [view](CERTS/36248069600bf4dcf5c85d58b56e6960278411436d818c338226d070a778cf31/README.md) |
| 31&#160;Dec&#160;22&#160;03:33&#160;UTC | SHAKEN 8526 | 31&#160;Dec&#160;23&#160;03:33&#160;UTC | true | [view](CERTS/96672cb8cc7c71e38d9e6350eac6c4f5c430c1eefeabe00cc095527e20a20c9b/README.md) |
| 05&#160;Jan&#160;23&#160;15:50&#160;UTC | SHAKEN 722J | 04&#160;Feb&#160;23&#160;15:50&#160;UTC | true | [view](CERTS/6707ea8caeed61aaa9c0489a5b1434622a9c938978f1bea4d2ea7fc5dbd81586/README.md) |
| 10&#160;Jan&#160;23&#160;13:53&#160;UTC | SHAKEN 815G | 10&#160;Apr&#160;23&#160;13:53&#160;UTC | true | [view](CERTS/2b3c67da667aa641a6a6da726a363284a1267727596a620d8e8b2b68fc3a6319/README.md) |
| 10&#160;Jan&#160;23&#160;16:32&#160;UTC | SHAKEN 551G | 09&#160;Feb&#160;23&#160;16:32&#160;UTC | true | [view](CERTS/628611026de07060665254cc9e3821cf274d115b7a1c27c0f747972dbf9ed121/README.md) |
| 12&#160;Jan&#160;23&#160;05:08&#160;UTC | SHAKEN 621J | 12&#160;May&#160;23&#160;05:08&#160;UTC | true | [view](CERTS/017fbb31b0e0c0f3dbb7de67e5e3f95c4ce10e546aaaca32faa97fb61eafcfe6/README.md) |
| 14&#160;Jan&#160;23&#160;04:23&#160;UTC | SHAKEN 345J | 13&#160;Feb&#160;23&#160;04:23&#160;UTC | true | [view](CERTS/ab63023fad304f5898fcd2921df1a0b78cf13aff37ce6a926370b243580920a6/README.md) |
| 14&#160;Jan&#160;23&#160;05:01&#160;UTC | SHAKEN 345J | 13&#160;Feb&#160;23&#160;05:01&#160;UTC | true | [view](CERTS/a1085082cf0a00fba83dcc43ff32d5aca855b418133feeebeb14728119047b3d/README.md) |
| 16&#160;Jan&#160;23&#160;17:54&#160;UTC | SHAKEN 622J | 15&#160;Jul&#160;23&#160;17:54&#160;UTC | true | [view](CERTS/4c659921cd1ec39c239ab31d4c1ca813017c01d51512ea11a1fe0572f034c534/README.md) |
| 17&#160;Jan&#160;23&#160;18:22&#160;UTC | SHAKEN 952J | 16&#160;Feb&#160;23&#160;18:21&#160;UTC | true | [view](CERTS/02c5d91ab286934436712c3537a6b58b3a1758b95a530644c10355fbfa9384a8/README.md) |
| 20&#160;Jan&#160;23&#160;19:31&#160;UTC | SHAKEN 841J | 03&#160;Feb&#160;23&#160;19:31&#160;UTC | true | [view](CERTS/96c578c5d7fd58be7968ca6a766b1b934d67ac653b3f3d26b77d494a3308b2c9/README.md) |
| 23&#160;Jan&#160;23&#160;20:34&#160;UTC | SHAKEN 738J | 01&#160;Feb&#160;23&#160;20:34&#160;UTC | true | [view](CERTS/340dec71f5239c8544ce1b145029a4c0c2a8c71383c02ec657c9f5ea941cacbc/README.md) |
| 24&#160;Jan&#160;23&#160;18:18&#160;UTC | SHAKEN 060K | 31&#160;Jan&#160;23&#160;18:18&#160;UTC | true | [view](CERTS/150a33a5020c2e7896817dff0ae5fd32cbf2b43c0e7f25f933c931a22f7f012f/README.md) |
| 24&#160;Jan&#160;23&#160;18:44&#160;UTC | SHAKEN 089K | 31&#160;Jan&#160;23&#160;18:44&#160;UTC | true | [view](CERTS/3ede36af3f8f7524776677f95a02d9a2b8316891855ca264c86dff21249226c1/README.md) |
| 24&#160;Jan&#160;23&#160;18:46&#160;UTC | SHAKEN 9714 | 31&#160;Jan&#160;23&#160;18:46&#160;UTC | true | [view](CERTS/8a2f4b095038ce4428131e93b6ac1a9c5330b9b46e370dd94597b728420d52e9/README.md) |
| 24&#160;Jan&#160;23&#160;20:03&#160;UTC | SHAKEN 297K | 31&#160;Jan&#160;23&#160;20:03&#160;UTC | true | [view](CERTS/791b5750c1c347f7d3fc60cdde1954881504baea498a2ee90c8a9bd99d8ea4b1/README.md) |
| 24&#160;Jan&#160;23&#160;20:10&#160;UTC | SHAKEN 366G | 31&#160;Jan&#160;23&#160;20:10&#160;UTC | true | [view](CERTS/13fc5be7dd6a771495fb65240a0609b1a85780b22fc9cc12f3941a21f44845b9/README.md) |
| 24&#160;Jan&#160;23&#160;20:25&#160;UTC | SHAKEN 674J | 31&#160;Jan&#160;23&#160;20:25&#160;UTC | true | [view](CERTS/db9ec8215606a125a901814f76c96f87bc60daf3d31a92643efc8f9c5a2366b9/README.md) |
| 24&#160;Jan&#160;23&#160;20:26&#160;UTC | SHAKEN 5512 | 31&#160;Jan&#160;23&#160;20:26&#160;UTC | true | [view](CERTS/b135c6a089896e41f00e4946c188e9a5f28d608212717b2cffe36af9d2f5bc66/README.md) |
| 24&#160;Jan&#160;23&#160;20:26&#160;UTC | SHAKEN 738J | 31&#160;Jan&#160;23&#160;20:26&#160;UTC | true | [view](CERTS/6237893c4e9197b5c01108389bf8b9665a5d7cced94f098de9d15953c9c3779a/README.md) |
| 24&#160;Jan&#160;23&#160;20:27&#160;UTC | SHAKEN 700H | 31&#160;Jan&#160;23&#160;20:27&#160;UTC | true | [view](CERTS/a74dbce570eefb0435073ac874c054b495d97654c4433fca74f672b6833ef3cf/README.md) |
| 24&#160;Jan&#160;23&#160;20:29&#160;UTC | SHAKEN 819H | 31&#160;Jan&#160;23&#160;20:29&#160;UTC | true | [view](CERTS/da2f079bdf594fd8bb060376ba75de1c0452e8ee7f338e07685e5efb94df58b0/README.md) |
| 24&#160;Jan&#160;23&#160;20:30&#160;UTC | SHAKEN 0753 | 31&#160;Jan&#160;23&#160;20:30&#160;UTC | true | [view](CERTS/3e3e5dff578225404c5e042cec83469e2164e370375e5b8f5fceaa6597e01a82/README.md) |
| 24&#160;Jan&#160;23&#160;20:30&#160;UTC | SHAKEN 769J | 31&#160;Jan&#160;23&#160;20:30&#160;UTC | true | [view](CERTS/44c6cdfd03b99e7eb34b3385d275b3e3d43e175aa8749ec8daf2800274aa9dc4/README.md) |
| 24&#160;Jan&#160;23&#160;20:32&#160;UTC | SHAKEN 691J | 31&#160;Jan&#160;23&#160;20:32&#160;UTC | true | [view](CERTS/8bab419e927e17def099d28f3c3146738f765abf62abc75d8970db664486407a/README.md) |
| 24&#160;Jan&#160;23&#160;20:32&#160;UTC | SHAKEN 849J | 31&#160;Jan&#160;23&#160;20:32&#160;UTC | true | [view](CERTS/75d72e05f44060d67af064f7a99374f1bfd0be0c38cb1226ad17f6408e098a75/README.md) |
| 24&#160;Jan&#160;23&#160;20:32&#160;UTC | SHAKEN 469A | 31&#160;Jan&#160;23&#160;20:32&#160;UTC | true | [view](CERTS/12917d7b086ca5324850944c4ed8b77925c3947cbdb0bb851f04e59ee6a11cc6/README.md) |
| 24&#160;Jan&#160;23&#160;20:32&#160;UTC | SHAKEN 726J | 31&#160;Jan&#160;23&#160;20:32&#160;UTC | true | [view](CERTS/53fd1ae3b9e168476dcf5d8278b137bc4ef82e2b5e3d71bf7fff4155e3a3ec04/README.md) |
| 24&#160;Jan&#160;23&#160;20:32&#160;UTC | SHAKEN 721J | 31&#160;Jan&#160;23&#160;20:32&#160;UTC | true | [view](CERTS/0fd3e48f7debd8062982ca5b70f335e442fedc10accfbb0822c3d7f50b8729b6/README.md) |
| 24&#160;Jan&#160;23&#160;20:32&#160;UTC | SHAKEN 790J | 31&#160;Jan&#160;23&#160;20:32&#160;UTC | true | [view](CERTS/eae6a62dd818921a98a143d9d77160aaa254049ae09819e9d732b4910b4fd71e/README.md) |
| 24&#160;Jan&#160;23&#160;20:33&#160;UTC | SHAKEN 738J | 31&#160;Jan&#160;23&#160;20:33&#160;UTC | true | [view](CERTS/630b774fe9efa172187912f367d5f40cb301b78b74fc1f898370aafe012ebcb9/README.md) |
| 24&#160;Jan&#160;23&#160;20:33&#160;UTC | SHAKEN 459J | 31&#160;Jan&#160;23&#160;20:33&#160;UTC | true | [view](CERTS/99a329dca09bb6d474d939951426820739cafc78efc8fc307ced813d29f45db0/README.md) |
| 24&#160;Jan&#160;23&#160;20:33&#160;UTC | SHAKEN 672B | 31&#160;Jan&#160;23&#160;20:33&#160;UTC | true | [view](CERTS/b2d2ffa1103306580162d3d07f17ffc8c0ff994c5cd0a70a9199b84f44f46a95/README.md) |
| 24&#160;Jan&#160;23&#160;20:34&#160;UTC | SHAKEN 366G | 31&#160;Jan&#160;23&#160;20:34&#160;UTC | true | [view](CERTS/2fb51706f96ba187c05f4b143c4acc91dc91a4cd13482552df760e8371d0ac76/README.md) |
| 24&#160;Jan&#160;23&#160;20:34&#160;UTC | SHAKEN 0226 | 31&#160;Jan&#160;23&#160;20:34&#160;UTC | true | [view](CERTS/00d3a4589eac39d1271441eeb3fdfb95321506b43ba8d7f1880622253be9d46b/README.md) |
| 24&#160;Jan&#160;23&#160;20:34&#160;UTC | SHAKEN 738J | 31&#160;Jan&#160;23&#160;20:34&#160;UTC | true | [view](CERTS/50d1fa37eb4e85ca5ab994c9fd2fc23cc1d036b731e3504c907c19855d253907/README.md) |
| 24&#160;Jan&#160;23&#160;21:06&#160;UTC | SHAKEN 518J | 31&#160;Jan&#160;23&#160;21:06&#160;UTC | true | [view](CERTS/57538d40a8474a090c6f6af845e45d39a4d74cda09c46c37553ff5357b2d6bbb/README.md) |
| 25&#160;Jan&#160;23&#160;14:28&#160;UTC | SHAKEN 873J | 25&#160;Jan&#160;24&#160;14:28&#160;UTC | true | [view](CERTS/36e9ba32d337ad13909bee0b62125c9769a049c4ae10d749687d75af4e49387d/README.md) |
| 25&#160;Jan&#160;23&#160;16:41&#160;UTC | SHAKEN 291K | 01&#160;Feb&#160;23&#160;16:41&#160;UTC | true | [view](CERTS/357f60aede892d2053b8f5d7de1702c18dbf8de2c97cd09d136706affc207b00/README.md) |
| 25&#160;Jan&#160;23&#160;18:11&#160;UTC | SHAKEN 406K | 01&#160;Feb&#160;23&#160;18:11&#160;UTC | true | [view](CERTS/223e0d18cb274b2f9d2f56a1e962b457862cf54b915cf8c817d0409a6eec3762/README.md) |
| 25&#160;Jan&#160;23&#160;20:22&#160;UTC | SHAKEN 983J | 01&#160;Feb&#160;23&#160;20:22&#160;UTC | true | [view](CERTS/5ed5f058cae002d64cf7e878ee4801df7a2f28328d1c3efadec013464c0c9b8f/README.md) |
| 25&#160;Jan&#160;23&#160;21:27&#160;UTC | SHAKEN 888J | 01&#160;Feb&#160;23&#160;21:27&#160;UTC | true | [view](CERTS/fc8fa7e4b53145c20dbbf269200560f3d3106d12bd7698c8c765a02ea8a91c54/README.md) |
| 25&#160;Jan&#160;23&#160;21:49&#160;UTC | SHAKEN 606F | 01&#160;Feb&#160;23&#160;21:49&#160;UTC | true | [view](CERTS/c16f0c9721343af60d1777a122630967762bc94b8663cf453598c391bf82db6a/README.md) |
| 26&#160;Jan&#160;23&#160;14:42&#160;UTC | SHAKEN 2550 | 02&#160;Feb&#160;23&#160;14:42&#160;UTC | true | [view](CERTS/2d24e7b2c542d5626f1f150e2d8ae721f5c8972bf3dfcb5fb6f80ee22dade58c/README.md) |
| 26&#160;Jan&#160;23&#160;16:03&#160;UTC | SHAKEN 625J | 02&#160;Feb&#160;23&#160;16:03&#160;UTC | true | [view](CERTS/4d52bd4a25557c0be9b9db4aef3ab392b6d0fbbd7af50131ae325e76427fcb9f/README.md) |
| 27&#160;Jan&#160;23&#160;14:49&#160;UTC | SHAKEN 186K | 03&#160;Feb&#160;23&#160;14:49&#160;UTC | true | [view](CERTS/d5f8c8046df9f33da26c926fb4cdf47c2a84892ae20c8dcc6f67af6b10315f85/README.md) |
| 27&#160;Jan&#160;23&#160;14:49&#160;UTC | SHAKEN 747J | 03&#160;Feb&#160;23&#160;14:49&#160;UTC | true | [view](CERTS/a5893561564f0232ad1ce0aeb52ca95cdf279fdd9c81e45aa7293902dfebce2c/README.md) |
| 27&#160;Jan&#160;23&#160;17:32&#160;UTC | SHAKEN 107K | 03&#160;Feb&#160;23&#160;17:32&#160;UTC | true | [view](CERTS/3b00b0f1b7227d2b10f57103a90b6b3e677ad14ecf47c7c820a04ee94072096e/README.md) |
| 27&#160;Jan&#160;23&#160;18:46&#160;UTC | SHAKEN 9714 | 03&#160;Feb&#160;23&#160;18:46&#160;UTC | true | [view](CERTS/1c709372131e845876d4deb120cf9422aff5f986a264adc7d512bfd569330184/README.md) |
| 27&#160;Jan&#160;23&#160;20:10&#160;UTC | SHAKEN 366G | 03&#160;Feb&#160;23&#160;20:10&#160;UTC | true | [view](CERTS/e0c3bbc2ac8e20a9fc2d2910d48b19e79f8a31740a3044f6b9e158ee9340804c/README.md) |
| 27&#160;Jan&#160;23&#160;20:25&#160;UTC | SHAKEN 674J | 03&#160;Feb&#160;23&#160;20:25&#160;UTC | true | [view](CERTS/797ae8f99abcf7f737a9714a87d82be90bb7cf48053d6dec9b33e3a67b0d26ea/README.md) |
| 27&#160;Jan&#160;23&#160;20:26&#160;UTC | SHAKEN 735J | 03&#160;Feb&#160;23&#160;20:26&#160;UTC | true | [view](CERTS/bdbb2cf7aa5c609eb25d015e817b5ad3ce0aeb0d347d18518c02fe82b054f0f4/README.md) |
| 27&#160;Jan&#160;23&#160;20:26&#160;UTC | SHAKEN 594J | 03&#160;Feb&#160;23&#160;20:26&#160;UTC | true | [view](CERTS/4f6440cd825453e07b362ecdb91d481ce47e52a9ec4240fb537119faa2b1d764/README.md) |
| 27&#160;Jan&#160;23&#160;20:27&#160;UTC | SHAKEN 738J | 03&#160;Feb&#160;23&#160;20:27&#160;UTC | true | [view](CERTS/6a6faab19afb561caa0c53c07597910d3cea3cb0910fb559a530a4c266c4c0fa/README.md) |
| 27&#160;Jan&#160;23&#160;20:28&#160;UTC | SHAKEN 700H | 03&#160;Feb&#160;23&#160;20:28&#160;UTC | true | [view](CERTS/cd37795af7f500abf1fbb558a85bf22749aef118354df67f73f30f31ac759c6d/README.md) |
| 27&#160;Jan&#160;23&#160;20:28&#160;UTC | SHAKEN 733J | 03&#160;Feb&#160;23&#160;20:28&#160;UTC | true | [view](CERTS/5c9e5f36c805b6769051f81be3ded35845efbdcd4467438ba7ee72d310695091/README.md) |
| 27&#160;Jan&#160;23&#160;20:31&#160;UTC | SHAKEN 727J | 03&#160;Feb&#160;23&#160;20:31&#160;UTC | true | [view](CERTS/363b4e9d947a16576791b4a55e8188c7865a311f3ba170f1511bbcfd6ba1ec74/README.md) |
| 27&#160;Jan&#160;23&#160;20:32&#160;UTC | SHAKEN 495J | 03&#160;Feb&#160;23&#160;20:32&#160;UTC | true | [view](CERTS/1474c306039dabad95ba4a45f1bfcae71b59a04dc387f7907bb70d70c1a2f840/README.md) |
| 27&#160;Jan&#160;23&#160;20:32&#160;UTC | SHAKEN 790J | 03&#160;Feb&#160;23&#160;20:32&#160;UTC | true | [view](CERTS/20dc6e293375e74407379ac68d9a2bd4c6f317878adc55c829cabaf5b83c649e/README.md) |
| 27&#160;Jan&#160;23&#160;20:33&#160;UTC | SHAKEN 738J | 03&#160;Feb&#160;23&#160;20:33&#160;UTC | true | [view](CERTS/be2544c80176f85963bf4f252ae7387e902762aad756cfc0bcb688d560af365b/README.md) |
| 27&#160;Jan&#160;23&#160;20:33&#160;UTC | SHAKEN 459J | 03&#160;Feb&#160;23&#160;20:33&#160;UTC | true | [view](CERTS/d24c7759146b2a1b7e98e50c581c87b27494c151ab93b7faf959005e4261dc0d/README.md) |
| 27&#160;Jan&#160;23&#160;20:34&#160;UTC | SHAKEN 672B | 03&#160;Feb&#160;23&#160;20:34&#160;UTC | true | [view](CERTS/4fbaec4168e89596738f0e898cfd7cd237acfc15b0ac553cdea03ecce37991d2/README.md) |
| 27&#160;Jan&#160;23&#160;20:35&#160;UTC | SHAKEN 366G | 03&#160;Feb&#160;23&#160;20:35&#160;UTC | true | [view](CERTS/5803a77c7137e46ec54982a338c380d0efe61e50668bf7d7424cfdcfb01402b0/README.md) |
| 27&#160;Jan&#160;23&#160;20:35&#160;UTC | SHAKEN 0226 | 03&#160;Feb&#160;23&#160;20:35&#160;UTC | true | [view](CERTS/405e5f856001ae649e579ec0fe0ee10154947453d8c919eedd7a921b9c536d3d/README.md) |
| 27&#160;Jan&#160;23&#160;20:35&#160;UTC | SHAKEN 738J | 03&#160;Feb&#160;23&#160;20:35&#160;UTC | true | [view](CERTS/a8cad6537699e072b5e1d042e25047b02556209f2da2d8d580f99f7404fa32bc/README.md) |
| 27&#160;Jan&#160;23&#160;21:06&#160;UTC | SHAKEN 518J | 03&#160;Feb&#160;23&#160;21:06&#160;UTC | true | [view](CERTS/1369cfb795bf78394a102bbfbebfc37b8463824eb70614c2a7124027fb1fee52/README.md) |
| 28&#160;Jan&#160;23&#160;14:13&#160;UTC | SHAKEN 0172 | 04&#160;Feb&#160;23&#160;14:13&#160;UTC | true | [view](CERTS/f55c927e328034978bc63cebb7c2232dde12c6235b916fdeb58fdba63e06325e/README.md) |
| 28&#160;Jan&#160;23&#160;16:42&#160;UTC | SHAKEN 291K | 04&#160;Feb&#160;23&#160;16:42&#160;UTC | true | [view](CERTS/958bb9cb29d715554be0957be333c155d1cf67ec000afc6961d274b8359ab0e3/README.md) |
| 28&#160;Jan&#160;23&#160;16:48&#160;UTC | SHAKEN 967J | 04&#160;Feb&#160;23&#160;16:48&#160;UTC | true | [view](CERTS/d22a63b0b3b8d2193802ef6964681c9322f4e2abba7fa8c9592da694e4ef7471/README.md) |
| 28&#160;Jan&#160;23&#160;20:23&#160;UTC | SHAKEN 983J | 04&#160;Feb&#160;23&#160;20:23&#160;UTC | true | [view](CERTS/d829f8f9d69340598a9ff8b1df789b5844a501967924832bd45ae3e55e9e3ea3/README.md) |
| 29&#160;Jan&#160;23&#160;01:46&#160;UTC | SHAKEN 278K | 05&#160;Feb&#160;23&#160;01:46&#160;UTC | true | [view](CERTS/6d1d07601a1ef96468d34711ab215889e9e16460fbf117ac540777b45d86d986/README.md) |
| 29&#160;Jan&#160;23&#160;05:10&#160;UTC | SHAKEN 841J | 12&#160;Feb&#160;23&#160;05:10&#160;UTC | true | [view](CERTS/b44d1c38bea2afab00ed90771ae66b232913e4cca7841ff6677d8276eec79482/README.md) |
| 29&#160;Jan&#160;23&#160;14:42&#160;UTC | SHAKEN 2550 | 05&#160;Feb&#160;23&#160;14:42&#160;UTC | true | [view](CERTS/8ebcdf19dc5bf7d633d0f76bec8cacd0426288fd52fc18b7cb122ce64844ad6a/README.md) |
| 29&#160;Jan&#160;23&#160;15:53&#160;UTC | SHAKEN 722J | 28&#160;Feb&#160;23&#160;15:53&#160;UTC | true | [view](CERTS/83dcf31f5c9ad810ac66c6d6357bf3807839c0b77cdc92033b8a4030f3a1e509/README.md) |
| 29&#160;Jan&#160;23&#160;20:20&#160;UTC | SHAKEN 177K | 05&#160;Feb&#160;23&#160;20:20&#160;UTC | true | [view](CERTS/9d8f207f4df7ae4d6ee4023da664c69879fd16a18e9022ca09629940f54bb009/README.md) |
| 30&#160;Jan&#160;23&#160;17:33&#160;UTC | SHAKEN 107K | 06&#160;Feb&#160;23&#160;17:33&#160;UTC | true | [view](CERTS/3f822413b91fa7bce223753af07c2c041821213838e71a8725bd6135ad8c853f/README.md) |
| 30&#160;Jan&#160;23&#160;20:10&#160;UTC | SHAKEN 366G | 06&#160;Feb&#160;23&#160;20:10&#160;UTC | true | [view](CERTS/8765e65e63054b439e947d4857f38230f17bc8e32b1b9ba0483fabae6e50858c/README.md) |
| 30&#160;Jan&#160;23&#160;20:26&#160;UTC | SHAKEN 674J | 06&#160;Feb&#160;23&#160;20:26&#160;UTC | true | [view](CERTS/1a1b82552e3e5d7916fece8d013e64b79bd2f34af019314eb245c6c2f56e49d5/README.md) |
| 30&#160;Jan&#160;23&#160;20:27&#160;UTC | SHAKEN 5512 | 06&#160;Feb&#160;23&#160;20:27&#160;UTC | true | [view](CERTS/c723b083bfa10ac731962a8007dcd98ef5aa2d540ad67bc1cfd4aecfd5026e3c/README.md) |
| 30&#160;Jan&#160;23&#160;20:30&#160;UTC | SHAKEN 769J | 06&#160;Feb&#160;23&#160;20:30&#160;UTC | true | [view](CERTS/f232f3f12815591f7aecc981ca9bc905f91ec7d989a29f89b48fc01c88cc091f/README.md) |
| 30&#160;Jan&#160;23&#160;20:34&#160;UTC | SHAKEN 738J | 06&#160;Feb&#160;23&#160;20:34&#160;UTC | true | [view](CERTS/efe0cd43307c3c9059142fce0564d40cd54dcafb33ed79ad437127cfe51e58cf/README.md) |
| 30&#160;Jan&#160;23&#160;20:35&#160;UTC | SHAKEN 738J | 08&#160;Feb&#160;23&#160;20:35&#160;UTC | true | [view](CERTS/cf832a21f53dc7eeace367df07cb6e7193c85ea15c7cfa3407e127412c907ccf/README.md) |
| 30&#160;Jan&#160;23&#160;20:35&#160;UTC | SHAKEN 366G | 06&#160;Feb&#160;23&#160;20:35&#160;UTC | true | [view](CERTS/7bfa690b2fd3c3710420e04ea43ed4fe051bbba65c1bba275ecb5d2831da4321/README.md) |
| 30&#160;Jan&#160;23&#160;20:35&#160;UTC | SHAKEN 0226 | 06&#160;Feb&#160;23&#160;20:35&#160;UTC | true | [view](CERTS/2ff237bf4f734cc87d6219874eab7af38bdc5ee5bba9b12c9e8ba1f6ac026924/README.md) |
| 30&#160;Jan&#160;23&#160;20:35&#160;UTC | SHAKEN 738J | 06&#160;Feb&#160;23&#160;20:35&#160;UTC | true | [view](CERTS/161e58ede036c61e4b3def63796ec61737a51bdea4c6277e1fdfa96c107576db/README.md) |

#### CA Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 20&#160;Aug&#160;21&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Issuing CA3 | 19&#160;Aug&#160;31&#160;23:59&#160;UTC | true | [view](CERTS/51b51431ffd8055cd7e30c213540f4868258058d4c4c9f23f1ed8bea75101e37/README.md) |
| 20&#160;Aug&#160;21&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Issuing CA1 | 19&#160;Aug&#160;31&#160;23:59&#160;UTC | true | [view](CERTS/037a7c6621814bd5f01fa2080a77a245458694d9d566e0050e068fc853317857/README.md) |
| 20&#160;Aug&#160;21&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Root CA1 | 19&#160;Aug&#160;41&#160;23:59&#160;UTC | true | [view](CERTS/59ab943f4e8c3c17755cc8e4abb5ac65736ce74ec587c62cb744ab0babbbe2fe/README.md) |
| 24&#160;Oct&#160;22&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Issuing CA4 | 23&#160;Oct&#160;32&#160;23:59&#160;UTC | false | [view](CERTS/80dd12d935f46c256c7c1289a0834b21ec8f2f8c35f2864928d1e75f3a280665/README.md) |
| 24&#160;Oct&#160;22&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Root CA2 | 23&#160;Oct&#160;42&#160;23:59&#160;UTC | false | [view](CERTS/a26e04fc786ab70b8085236b2c53f8cfbf5d0c6a5c2c9c3e9f91669fbb8ea4d5/README.md) |


Generated: 31 Jan 23 17:51 UTC