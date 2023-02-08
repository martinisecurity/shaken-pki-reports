# STIR/SHAKEN CA Ecosystem Compliance

## TransNexus

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 1384 certificates were included in the corpus being tested
- 1 certificates in the corpus were skipped because they are duplicates
- 1297 certificates in the corpus were skipped because they are expired
- 1 certificates in the corpus were skipped because they are not currently trusted
- 85 certificates being tested against the remaining rules
- 1.28 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 11.76% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 70 days is the average remaining validity for the certificates in the corpus
- 71 days is the average initial validity for the certificates in the corpus
- 68 certificates expire in the next 30 days
- 1.39 average number of unexpired certificates per OCN observed
- 61 unique OCNs observed in unexpired and valid certificate corpus

| Instances | Test | Source |
|-----------|------|--------|
| 2 | [e_atis_certificate_policies](ISSUES/e_atis_certificate_policies/README.md) | ATIS1000080 |
| 10 | [e_incorrect_ku_encoding](ISSUES/e_incorrect_ku_encoding/README.md) | RFC5280 |
| 2 | [e_us_cp_ambiguous_identifier](ISSUES/e_us_cp_ambiguous_identifier/README.md) | US_SHAKEN_CP |
| 85 | [e_us_cp_subject_sn](ISSUES/e_us_cp_subject_sn/README.md) | US_SHAKEN_CP |
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
- 5461 days is the average remaining validity for the certificates in the corpus
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
| 10&#160;Jan&#160;23&#160;13:53&#160;UTC | SHAKEN 815G | 10&#160;Apr&#160;23&#160;13:53&#160;UTC | true | [view](CERTS/2b3c67da667aa641a6a6da726a363284a1267727596a620d8e8b2b68fc3a6319/README.md) |
| 10&#160;Jan&#160;23&#160;16:32&#160;UTC | SHAKEN 551G | 09&#160;Feb&#160;23&#160;16:32&#160;UTC | true | [view](CERTS/628611026de07060665254cc9e3821cf274d115b7a1c27c0f747972dbf9ed121/README.md) |
| 12&#160;Jan&#160;23&#160;05:08&#160;UTC | SHAKEN 621J | 12&#160;May&#160;23&#160;05:08&#160;UTC | true | [view](CERTS/017fbb31b0e0c0f3dbb7de67e5e3f95c4ce10e546aaaca32faa97fb61eafcfe6/README.md) |
| 14&#160;Jan&#160;23&#160;04:23&#160;UTC | SHAKEN 345J | 13&#160;Feb&#160;23&#160;04:23&#160;UTC | true | [view](CERTS/ab63023fad304f5898fcd2921df1a0b78cf13aff37ce6a926370b243580920a6/README.md) |
| 14&#160;Jan&#160;23&#160;05:01&#160;UTC | SHAKEN 345J | 13&#160;Feb&#160;23&#160;05:01&#160;UTC | true | [view](CERTS/a1085082cf0a00fba83dcc43ff32d5aca855b418133feeebeb14728119047b3d/README.md) |
| 16&#160;Jan&#160;23&#160;17:54&#160;UTC | SHAKEN 622J | 15&#160;Jul&#160;23&#160;17:54&#160;UTC | true | [view](CERTS/4c659921cd1ec39c239ab31d4c1ca813017c01d51512ea11a1fe0572f034c534/README.md) |
| 17&#160;Jan&#160;23&#160;18:22&#160;UTC | SHAKEN 952J | 16&#160;Feb&#160;23&#160;18:21&#160;UTC | true | [view](CERTS/02c5d91ab286934436712c3537a6b58b3a1758b95a530644c10355fbfa9384a8/README.md) |
| 25&#160;Jan&#160;23&#160;14:28&#160;UTC | SHAKEN 873J | 25&#160;Jan&#160;24&#160;14:28&#160;UTC | true | [view](CERTS/36e9ba32d337ad13909bee0b62125c9769a049c4ae10d749687d75af4e49387d/README.md) |
| 29&#160;Jan&#160;23&#160;05:10&#160;UTC | SHAKEN 841J | 12&#160;Feb&#160;23&#160;05:10&#160;UTC | true | [view](CERTS/b44d1c38bea2afab00ed90771ae66b232913e4cca7841ff6677d8276eec79482/README.md) |
| 29&#160;Jan&#160;23&#160;15:53&#160;UTC | SHAKEN 722J | 28&#160;Feb&#160;23&#160;15:53&#160;UTC | true | [view](CERTS/83dcf31f5c9ad810ac66c6d6357bf3807839c0b77cdc92033b8a4030f3a1e509/README.md) |
| 30&#160;Jan&#160;23&#160;20:35&#160;UTC | SHAKEN 738J | 08&#160;Feb&#160;23&#160;20:35&#160;UTC | true | [view](CERTS/cf832a21f53dc7eeace367df07cb6e7193c85ea15c7cfa3407e127412c907ccf/README.md) |
| 31&#160;Jan&#160;23&#160;16:33&#160;UTC | SHAKEN 551G | 02&#160;Mar&#160;23&#160;16:33&#160;UTC | true | [view](CERTS/a670b41080ada0e8ccc5fa5de6d0eecb7ffe7a37b272e63d9640b4f2f7c2435c/README.md) |
| 01&#160;Feb&#160;23&#160;20:20&#160;UTC | SHAKEN 177K | 08&#160;Feb&#160;23&#160;20:20&#160;UTC | true | [view](CERTS/4dd2d045ec3c5c57db48de45169a41a759436036bbd6aa2ec1526f1e5630127a/README.md) |
| 01&#160;Feb&#160;23&#160;20:31&#160;UTC | SHAKEN 848J | 08&#160;Feb&#160;23&#160;20:31&#160;UTC | true | [view](CERTS/32254328ee850d6b3fe72a616c5484c6f9bf5ac08677a95ea8ab1de69f6dcc1c/README.md) |
| 02&#160;Feb&#160;23&#160;14:50&#160;UTC | SHAKEN 186K | 09&#160;Feb&#160;23&#160;14:50&#160;UTC | true | [view](CERTS/e0919590c48f562d6c711ed2aeb86ab52ce749917d6c67326f69cde670a6d41c/README.md) |
| 02&#160;Feb&#160;23&#160;14:50&#160;UTC | SHAKEN 747J | 09&#160;Feb&#160;23&#160;14:50&#160;UTC | true | [view](CERTS/0ec0af685f7a48fb18741bffbfdd0e8f631a74163b790806f0a83f9b56a11e42/README.md) |
| 02&#160;Feb&#160;23&#160;17:33&#160;UTC | SHAKEN 107K | 09&#160;Feb&#160;23&#160;17:33&#160;UTC | true | [view](CERTS/8a15c00796ebaee2f38c76fb25187ed0f635710df1be099d95c3a37689f8d153/README.md) |
| 02&#160;Feb&#160;23&#160;20:26&#160;UTC | SHAKEN 674J | 09&#160;Feb&#160;23&#160;20:26&#160;UTC | true | [view](CERTS/a96e2d7b39dd6b0ef8d1e25da9e2fef861f691e9f453de5ccdb3115816c52fb6/README.md) |
| 02&#160;Feb&#160;23&#160;20:27&#160;UTC | SHAKEN 738J | 09&#160;Feb&#160;23&#160;20:27&#160;UTC | true | [view](CERTS/1e42a1bb58a5e201c9ba9e79d0c446e9bdc1f73f6c9081f8d5e02f257a1a0a59/README.md) |
| 02&#160;Feb&#160;23&#160;20:28&#160;UTC | SHAKEN 733J | 09&#160;Feb&#160;23&#160;20:28&#160;UTC | true | [view](CERTS/b30fc6169c94d0c73d30a774affc320785e8282ff1a6c42f40badc8d48de0e82/README.md) |
| 02&#160;Feb&#160;23&#160;20:33&#160;UTC | SHAKEN 691J | 09&#160;Feb&#160;23&#160;20:33&#160;UTC | true | [view](CERTS/16c543a044f5513ddedc35df558aaf350db21cc4a6dd376bd55f8fd0af42d5b0/README.md) |
| 02&#160;Feb&#160;23&#160;20:34&#160;UTC | SHAKEN 625J | 09&#160;Feb&#160;23&#160;20:34&#160;UTC | true | [view](CERTS/10fed67fe933a6ef3c6dc123d1ab5d9cf70eb63be105978f314c3e36c2ac3c76/README.md) |
| 02&#160;Feb&#160;23&#160;20:34&#160;UTC | SHAKEN 738J | 09&#160;Feb&#160;23&#160;20:34&#160;UTC | true | [view](CERTS/52e3364e820d93238bc6281e1cc002cfb89ffb42e913a0a06e8b3262c7176f7f/README.md) |
| 02&#160;Feb&#160;23&#160;20:34&#160;UTC | SHAKEN 459J | 09&#160;Feb&#160;23&#160;20:34&#160;UTC | true | [view](CERTS/7da66bba39b5c52c1c6bd5b66f652762f6540c88ad492288653d378082e08872/README.md) |
| 02&#160;Feb&#160;23&#160;20:35&#160;UTC | SHAKEN 366G | 09&#160;Feb&#160;23&#160;20:35&#160;UTC | true | [view](CERTS/2865c60df05221a9d8dfee945d980b8c4ab7e1bf26f4ba351fd84b61abd93979/README.md) |
| 02&#160;Feb&#160;23&#160;20:35&#160;UTC | SHAKEN 738J | 09&#160;Feb&#160;23&#160;20:35&#160;UTC | true | [view](CERTS/7ba100a525a1efc56c2a242d4370aaa7180119af56678fde9211aa8bf4f16a84/README.md) |
| 03&#160;Feb&#160;23&#160;16:49&#160;UTC | SHAKEN 967J | 10&#160;Feb&#160;23&#160;16:49&#160;UTC | true | [view](CERTS/806b1627950b79fb3f66ba0894d9566636a096e4e7d88a1c52a7281fff092303/README.md) |
| 03&#160;Feb&#160;23&#160;19:15&#160;UTC | SHAKEN 749J | 10&#160;Feb&#160;23&#160;19:15&#160;UTC | true | [view](CERTS/8e76fb9f64d17973626b80dd2a01c5486e281e33c9dc437e33995e20429a6405/README.md) |
| 03&#160;Feb&#160;23&#160;20:24&#160;UTC | SHAKEN 983J | 10&#160;Feb&#160;23&#160;20:23&#160;UTC | true | [view](CERTS/cddd54eda87431e658a7fbb3ab9cd818d77cdab0e1c10ca01c61f145718890d9/README.md) |
| 03&#160;Feb&#160;23&#160;21:49&#160;UTC | SHAKEN 606F | 10&#160;Feb&#160;23&#160;21:49&#160;UTC | true | [view](CERTS/d641ae547514e5b28c5abe16afa9e92b9ba8a2eefe19ac475f415cb4bff3043e/README.md) |
| 04&#160;Feb&#160;23&#160;01:47&#160;UTC | SHAKEN 278K | 11&#160;Feb&#160;23&#160;01:47&#160;UTC | true | [view](CERTS/1bd32c89b0baa94982b0a32d46bc4d4b2e899ac3a55dfe4e2dbe3937e854ac0d/README.md) |
| 04&#160;Feb&#160;23&#160;04:30&#160;UTC | SHAKEN 345J | 06&#160;Mar&#160;23&#160;04:30&#160;UTC | true | [view](CERTS/36db5d2ae222cbd48a977867b2d20674492d8f764f8f2c25f8a8366ac6c4e4df/README.md) |
| 04&#160;Feb&#160;23&#160;05:05&#160;UTC | SHAKEN 345J | 06&#160;Mar&#160;23&#160;05:05&#160;UTC | true | [view](CERTS/a288fed8ab89a89acc2e957e18dcbede0be84d111f2f3a2a75c3372ce021f082/README.md) |
| 04&#160;Feb&#160;23&#160;14:43&#160;UTC | SHAKEN 2550 | 11&#160;Feb&#160;23&#160;14:43&#160;UTC | true | [view](CERTS/53f1790a51421b50c91b2715c1cbaf561117c8d297eb04a09cb0f51ae3e07333/README.md) |
| 04&#160;Feb&#160;23&#160;18:30&#160;UTC | SHAKEN 952J | 06&#160;Mar&#160;23&#160;18:30&#160;UTC | true | [view](CERTS/ac30e0e0a8dc1da3eea39fb6a9a3cb840aeb4637df1beeb25ff452eab1970eb1/README.md) |
| 04&#160;Feb&#160;23&#160;20:20&#160;UTC | SHAKEN 177K | 11&#160;Feb&#160;23&#160;20:20&#160;UTC | true | [view](CERTS/2016d1373a93cc6205ee1c14e0054840740cadab3e5529b839fb9676d68fc072/README.md) |
| 05&#160;Feb&#160;23&#160;14:49&#160;UTC | SHAKEN 012K | 12&#160;Feb&#160;23&#160;14:49&#160;UTC | true | [view](CERTS/551f749446b1b891c4ef12f46d6ba4b4fd15ba8c272387a131b37ffc1ca78e20/README.md) |
| 05&#160;Feb&#160;23&#160;14:50&#160;UTC | SHAKEN 186K | 12&#160;Feb&#160;23&#160;14:50&#160;UTC | true | [view](CERTS/d21e5bd61bb9f1168575124da96ac48d6c884f4c669fb8ab075714ff1c02b0cc/README.md) |
| 05&#160;Feb&#160;23&#160;17:33&#160;UTC | SHAKEN 107K | 12&#160;Feb&#160;23&#160;17:33&#160;UTC | true | [view](CERTS/71102e6d2ffa8ea99f7e17baeed2bb9b0e0b30eceefcf2d291d74e128480b31c/README.md) |
| 05&#160;Feb&#160;23&#160;18:44&#160;UTC | SHAKEN 089K | 12&#160;Feb&#160;23&#160;18:44&#160;UTC | true | [view](CERTS/6d8244e3ba5c3415746a16f845f331dceaabe51856b59efe5faecd41d0b4d863/README.md) |
| 05&#160;Feb&#160;23&#160;18:47&#160;UTC | SHAKEN 9714 | 12&#160;Feb&#160;23&#160;18:47&#160;UTC | true | [view](CERTS/e5a988c729c94e539295c6b10c92029a63dda7bcf99d5fe1e538c8ed437e77f7/README.md) |
| 05&#160;Feb&#160;23&#160;20:04&#160;UTC | SHAKEN 297K | 12&#160;Feb&#160;23&#160;20:04&#160;UTC | true | [view](CERTS/1aedf433924f7f6178bd55027d75871794db67512d3a12466733706133209c25/README.md) |
| 05&#160;Feb&#160;23&#160;20:26&#160;UTC | SHAKEN 674J | 12&#160;Feb&#160;23&#160;20:26&#160;UTC | true | [view](CERTS/55298e793dfcae1077c8ec088ac4d564a8d8b42b9b2dfc248f177d35c10d2e80/README.md) |
| 05&#160;Feb&#160;23&#160;20:26&#160;UTC | SHAKEN 3013 | 12&#160;Feb&#160;23&#160;20:26&#160;UTC | true | [view](CERTS/1fb755b425644eb9e290b7c969954741c5cf6e61d09cf5f03693bf51bbeb7d0a/README.md) |
| 05&#160;Feb&#160;23&#160;20:27&#160;UTC | SHAKEN 735J | 12&#160;Feb&#160;23&#160;20:27&#160;UTC | true | [view](CERTS/7dd6c90864c8069625830645accd92f61cf105c48a3641d22292cb6bdaf7e5c1/README.md) |
| 05&#160;Feb&#160;23&#160;20:27&#160;UTC | SHAKEN 738J | 12&#160;Feb&#160;23&#160;20:27&#160;UTC | true | [view](CERTS/7e4fa785d57626fb89d4e523763eca2bc2cbefd628f6fbc9345302ea2e0ca3d1/README.md) |
| 05&#160;Feb&#160;23&#160;20:28&#160;UTC | SHAKEN 700H | 12&#160;Feb&#160;23&#160;20:28&#160;UTC | true | [view](CERTS/d8555a49fa8be257c8a6204c4f6af75bb2ad2d431c5d01a68ae5c35a20e76ca4/README.md) |
| 05&#160;Feb&#160;23&#160;20:28&#160;UTC | SHAKEN 733J | 12&#160;Feb&#160;23&#160;20:28&#160;UTC | true | [view](CERTS/1d127b1d8666e0c945f13e546b996daa96c67b5485e921140492eec837030392/README.md) |
| 05&#160;Feb&#160;23&#160;20:28&#160;UTC | SHAKEN 7453 | 12&#160;Feb&#160;23&#160;20:28&#160;UTC | true | [view](CERTS/35f229ecbc36e3d54c73684cc8445ce8877bb94e96091e84855a316ca167b16b/README.md) |
| 05&#160;Feb&#160;23&#160;20:30&#160;UTC | SHAKEN 819H | 12&#160;Feb&#160;23&#160;20:30&#160;UTC | true | [view](CERTS/a07ab37343ca0c894a3bec5ca512f0b12f7a0dde2ac084a6323e44f13d10ef86/README.md) |
| 05&#160;Feb&#160;23&#160;20:33&#160;UTC | SHAKEN 849J | 12&#160;Feb&#160;23&#160;20:33&#160;UTC | true | [view](CERTS/b0c633787851fae0464b78e0403779a2e8bffb3fdccdaeea91a99f7a0d9cafa7/README.md) |
| 05&#160;Feb&#160;23&#160;20:33&#160;UTC | SHAKEN 469A | 12&#160;Feb&#160;23&#160;20:33&#160;UTC | true | [view](CERTS/0430b0f343c2bb2535022e961685ff106660ec62fa49749f1daea0b654b50319/README.md) |
| 05&#160;Feb&#160;23&#160;20:33&#160;UTC | SHAKEN 495J | 12&#160;Feb&#160;23&#160;20:33&#160;UTC | true | [view](CERTS/f90f55e56d92683562e9604afb5f7457a1a9b20d92720b9911337c8528a34530/README.md) |
| 05&#160;Feb&#160;23&#160;20:33&#160;UTC | SHAKEN 790J | 12&#160;Feb&#160;23&#160;20:33&#160;UTC | true | [view](CERTS/09af0077a40ac614530d6d57c9ca54a10a807e277250c34deab662faedab4a5f/README.md) |
| 05&#160;Feb&#160;23&#160;20:34&#160;UTC | SHAKEN 738J | 12&#160;Feb&#160;23&#160;20:34&#160;UTC | true | [view](CERTS/3bc7ca5f08d3b3b5f60ce49a8096e018eafb453fdf28ae91b41a13ed393060a9/README.md) |
| 05&#160;Feb&#160;23&#160;20:34&#160;UTC | SHAKEN 459J | 12&#160;Feb&#160;23&#160;20:34&#160;UTC | true | [view](CERTS/4e18fd1050bbd6528f9c8b00d5b89be375c893fa65cf8001002156bc1e0e245f/README.md) |
| 05&#160;Feb&#160;23&#160;20:34&#160;UTC | SHAKEN 672B | 12&#160;Feb&#160;23&#160;20:34&#160;UTC | true | [view](CERTS/a3c82e3dd0ff46fc0c76dbb849406c2902e2654cb043e69201905510a73cdc7f/README.md) |
| 05&#160;Feb&#160;23&#160;20:35&#160;UTC | SHAKEN 366G | 12&#160;Feb&#160;23&#160;20:35&#160;UTC | true | [view](CERTS/24a1f0456619350a1bcee8d5cc6aa515b3a856bfeecdd6e03d2f173acfa11b01/README.md) |
| 05&#160;Feb&#160;23&#160;20:35&#160;UTC | SHAKEN 0226 | 12&#160;Feb&#160;23&#160;20:35&#160;UTC | true | [view](CERTS/80014a7679c62baeacd39aa754aaa9edfc8841f56c75772092fedf2bfca491a9/README.md) |
| 05&#160;Feb&#160;23&#160;20:35&#160;UTC | SHAKEN 738J | 12&#160;Feb&#160;23&#160;20:35&#160;UTC | true | [view](CERTS/db30162de95ed0130b7c17e90a0d823faf8784d163cb5fb56468e4305efbff54/README.md) |
| 05&#160;Feb&#160;23&#160;23:53&#160;UTC | SHAKEN 0172 | 12&#160;Feb&#160;23&#160;23:53&#160;UTC | true | [view](CERTS/8f6a61f29bb8a32eca602c77ce39f1396568f8d997fd87bb3dbc6b5d97e97127/README.md) |
| 06&#160;Feb&#160;23&#160;14:54&#160;UTC | SHAKEN 841J | 20&#160;Feb&#160;23&#160;14:54&#160;UTC | true | [view](CERTS/5b773298ba88c03e3f148edd7eacd0cd578a782e42f6e97ff6ea63ab7bac6fef/README.md) |
| 06&#160;Feb&#160;23&#160;18:12&#160;UTC | SHAKEN 406K | 13&#160;Feb&#160;23&#160;18:12&#160;UTC | true | [view](CERTS/ad6fab5fbb8e344f577e676b5553b34a0138bf2f5e1407a674971cdc76c0841f/README.md) |
| 06&#160;Feb&#160;23&#160;20:24&#160;UTC | SHAKEN 983J | 13&#160;Feb&#160;23&#160;20:24&#160;UTC | true | [view](CERTS/635135ab4dc1aa297841a9503b8162081138e5f89e4cfb5326aaa72b65fac3a6/README.md) |
| 06&#160;Feb&#160;23&#160;20:36&#160;UTC | SHAKEN 738J | 15&#160;Feb&#160;23&#160;20:36&#160;UTC | true | [view](CERTS/5993786fa76005e81a0fbbb8a234d738be7f8c437948d84da38ed2a92ff9ff94/README.md) |
| 06&#160;Feb&#160;23&#160;21:49&#160;UTC | SHAKEN 606F | 13&#160;Feb&#160;23&#160;21:49&#160;UTC | true | [view](CERTS/6521257b58ae068612f2bcc52d8be2474ff9301bd8d06e0371d5dba736a9ebea/README.md) |
| 07&#160;Feb&#160;23&#160;01:47&#160;UTC | SHAKEN 278K | 14&#160;Feb&#160;23&#160;01:47&#160;UTC | true | [view](CERTS/26d86341ae45edce51d7e07d6fb3aa88f219e39b0ce0c34931780bec5c7a781c/README.md) |
| 07&#160;Feb&#160;23&#160;14:43&#160;UTC | SHAKEN 2550 | 14&#160;Feb&#160;23&#160;14:43&#160;UTC | true | [view](CERTS/77e5d84a6294dd3c159d943a645a8d4066786f5158da8f9634bd2c38e0fb8dd3/README.md) |
| 07&#160;Feb&#160;23&#160;16:05&#160;UTC | SHAKEN 311H | 14&#160;Feb&#160;23&#160;16:05&#160;UTC | true | [view](CERTS/9394eed2e68ec6016c58762101365f28992f9c48a909d04768ea62541785cbe7/README.md) |

#### CA Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 20&#160;Aug&#160;21&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Issuing CA3 | 19&#160;Aug&#160;31&#160;23:59&#160;UTC | true | [view](CERTS/51b51431ffd8055cd7e30c213540f4868258058d4c4c9f23f1ed8bea75101e37/README.md) |
| 20&#160;Aug&#160;21&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Issuing CA1 | 19&#160;Aug&#160;31&#160;23:59&#160;UTC | true | [view](CERTS/037a7c6621814bd5f01fa2080a77a245458694d9d566e0050e068fc853317857/README.md) |
| 20&#160;Aug&#160;21&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Root CA1 | 19&#160;Aug&#160;41&#160;23:59&#160;UTC | true | [view](CERTS/59ab943f4e8c3c17755cc8e4abb5ac65736ce74ec587c62cb744ab0babbbe2fe/README.md) |
| 24&#160;Oct&#160;22&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Issuing CA4 | 23&#160;Oct&#160;32&#160;23:59&#160;UTC | false | [view](CERTS/80dd12d935f46c256c7c1289a0834b21ec8f2f8c35f2864928d1e75f3a280665/README.md) |
| 24&#160;Oct&#160;22&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Root CA2 | 23&#160;Oct&#160;42&#160;23:59&#160;UTC | false | [view](CERTS/a26e04fc786ab70b8085236b2c53f8cfbf5d0c6a5c2c9c3e9f91669fbb8ea4d5/README.md) |


Generated: 08 Feb 23 19:45 UTC