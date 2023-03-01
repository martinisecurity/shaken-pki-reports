# STIR/SHAKEN CA Ecosystem Compliance

## TransNexus

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 1626 certificates were included in the corpus being tested
- 1 certificates in the corpus were skipped because they are duplicates
- 1532 certificates in the corpus were skipped because they are expired
- 1 certificates in the corpus were skipped because they are not currently trusted
- 92 certificates being tested against the remaining rules
- 1.24 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 9.78% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 63 days is the average remaining validity for the certificates in the corpus
- 64 days is the average initial validity for the certificates in the corpus
- 76 certificates expire in the next 30 days
- 1.48 average number of unexpired certificates per OCN observed
- 62 unique OCNs observed in unexpired and valid certificate corpus

| Instances | Test | Source |
|-----------|------|--------|
| 2 | [e_atis_certificate_policies](ISSUES/e_atis_certificate_policies/README.md) | ATIS1000080 |
| 9 | [e_incorrect_ku_encoding](ISSUES/e_incorrect_ku_encoding/README.md) | RFC5280 |
| 2 | [e_us_cp_ambiguous_identifier](ISSUES/e_us_cp_ambiguous_identifier/README.md) | US_SHAKEN_CP |
| 92 | [e_us_cp_subject_sn](ISSUES/e_us_cp_subject_sn/README.md) | US_SHAKEN_CP |
| 9 | [w_shaken_subject_rdn_unknown](ISSUES/w_shaken_subject_rdn_unknown/README.md) | SHAKEN_PKI_BEST_PRACTICES |

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
- 5457 days is the average remaining validity for the certificates in the corpus
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
| 13&#160;Sep&#160;22&#160;17:24&#160;UTC | SHAKEN 706J | 13&#160;Sep&#160;23&#160;17:24&#160;UTC | true | [view](CERTS/1eea5e17ae8b36cb5b6d657178273af745cea015e63586e6d44d1ba2ee8f722a/README.md) |
| 23&#160;Nov&#160;22&#160;17:31&#160;UTC | SHAKEN 4036 | 10&#160;May&#160;23&#160;17:31&#160;UTC | true | [view](CERTS/e03ff4259e6b456f1bb19d011f005932e83b0a35b520730c97cd77ad8868071c/README.md) |
| 07&#160;Dec&#160;22&#160;00:45&#160;UTC | SHAKEN 299K | 07&#160;Dec&#160;23&#160;00:45&#160;UTC | true | [view](CERTS/7b5e04ff01ace2ee384df3507886684204901bf23851866eb6b3891ecdd96a92/README.md) |
| 22&#160;Dec&#160;22&#160;17:30&#160;UTC | SHAKEN 505J | 22&#160;Mar&#160;23&#160;17:30&#160;UTC | true | [view](CERTS/7ff43923df80c3a7e2d854b42f33c69069af4625f5267df0d7ff7b2004ad1c1b/README.md) |
| 31&#160;Dec&#160;22&#160;03:33&#160;UTC | SHAKEN 8526 | 31&#160;Dec&#160;23&#160;03:33&#160;UTC | true | [view](CERTS/96672cb8cc7c71e38d9e6350eac6c4f5c430c1eefeabe00cc095527e20a20c9b/README.md) |
| 10&#160;Jan&#160;23&#160;13:53&#160;UTC | SHAKEN 815G | 10&#160;Apr&#160;23&#160;13:53&#160;UTC | true | [view](CERTS/2b3c67da667aa641a6a6da726a363284a1267727596a620d8e8b2b68fc3a6319/README.md) |
| 12&#160;Jan&#160;23&#160;05:08&#160;UTC | SHAKEN 621J | 12&#160;May&#160;23&#160;05:08&#160;UTC | true | [view](CERTS/017fbb31b0e0c0f3dbb7de67e5e3f95c4ce10e546aaaca32faa97fb61eafcfe6/README.md) |
| 16&#160;Jan&#160;23&#160;17:54&#160;UTC | SHAKEN 622J | 15&#160;Jul&#160;23&#160;17:54&#160;UTC | true | [view](CERTS/4c659921cd1ec39c239ab31d4c1ca813017c01d51512ea11a1fe0572f034c534/README.md) |
| 25&#160;Jan&#160;23&#160;14:28&#160;UTC | SHAKEN 873J | 25&#160;Jan&#160;24&#160;14:28&#160;UTC | true | [view](CERTS/36e9ba32d337ad13909bee0b62125c9769a049c4ae10d749687d75af4e49387d/README.md) |
| 31&#160;Jan&#160;23&#160;16:33&#160;UTC | SHAKEN 551G | 02&#160;Mar&#160;23&#160;16:33&#160;UTC | true | [view](CERTS/a670b41080ada0e8ccc5fa5de6d0eecb7ffe7a37b272e63d9640b4f2f7c2435c/README.md) |
| 01&#160;Feb&#160;23&#160;14:54&#160;UTC | SHAKEN 577F | 03&#160;Mar&#160;23&#160;14:54&#160;UTC | true | [view](CERTS/42c031e9c52fa1ffe78c41858e23a1870613c8a0f862013b3c664fc71475fdfe/README.md) |
| 04&#160;Feb&#160;23&#160;04:30&#160;UTC | SHAKEN 345J | 06&#160;Mar&#160;23&#160;04:30&#160;UTC | true | [view](CERTS/36db5d2ae222cbd48a977867b2d20674492d8f764f8f2c25f8a8366ac6c4e4df/README.md) |
| 04&#160;Feb&#160;23&#160;05:05&#160;UTC | SHAKEN 345J | 06&#160;Mar&#160;23&#160;05:05&#160;UTC | true | [view](CERTS/a288fed8ab89a89acc2e957e18dcbede0be84d111f2f3a2a75c3372ce021f082/README.md) |
| 04&#160;Feb&#160;23&#160;18:30&#160;UTC | SHAKEN 952J | 06&#160;Mar&#160;23&#160;18:30&#160;UTC | true | [view](CERTS/ac30e0e0a8dc1da3eea39fb6a9a3cb840aeb4637df1beeb25ff452eab1970eb1/README.md) |
| 19&#160;Feb&#160;23&#160;15:03&#160;UTC | SHAKEN 577F | 21&#160;Mar&#160;23&#160;15:03&#160;UTC | true | [view](CERTS/9d063726d179f92c7f9438afcf499083ea14b81851800001fa529ece9610253d/README.md) |
| 20&#160;Feb&#160;23&#160;20:38&#160;UTC | SHAKEN 738J | 01&#160;Mar&#160;23&#160;20:38&#160;UTC | true | [view](CERTS/1f1d6a4470a065cfb0ca8b3282135deffbe2547d7be195b53266e5fb1840d3c8/README.md) |
| 21&#160;Feb&#160;23&#160;16:41&#160;UTC | SHAKEN 551G | 23&#160;Mar&#160;23&#160;16:41&#160;UTC | true | [view](CERTS/108270b9fe753c5036472d046f78a2ba12a48bd6a1da0d55df207fdb907d9821/README.md) |
| 22&#160;Feb&#160;23&#160;15:54&#160;UTC | SHAKEN 722J | 24&#160;Mar&#160;23&#160;15:54&#160;UTC | true | [view](CERTS/cc3c678774a06d4a1cf0f216048f71627afaeb92d62800d19fed64c029a1a729/README.md) |
| 22&#160;Feb&#160;23&#160;18:40&#160;UTC | SHAKEN 952J | 24&#160;Mar&#160;23&#160;18:40&#160;UTC | true | [view](CERTS/52a3abc5046e0441a8fa5c7dbe5f04f79db1c75a20fec50f377b5cf45b5f4b4f/README.md) |
| 22&#160;Feb&#160;23&#160;19:12&#160;UTC | SHAKEN 0172 | 01&#160;Mar&#160;23&#160;19:12&#160;UTC | true | [view](CERTS/d7ebc07802bebd2555405575e00ac6ab71bcdac5bf31a929bc2111375f738840/README.md) |
| 22&#160;Feb&#160;23&#160;20:21&#160;UTC | SHAKEN 177K | 01&#160;Mar&#160;23&#160;20:21&#160;UTC | true | [view](CERTS/62c0a49bc4af90269c8c8b47d65940c63601fba906489ff30581563d55d1e892/README.md) |
| 23&#160;Feb&#160;23&#160;10:11&#160;UTC | SHAKEN 841J | 09&#160;Mar&#160;23&#160;10:11&#160;UTC | true | [view](CERTS/d899b2aa1d09dde6b28e656338de29018e4af4575197c7bb2d00183f1e9b43a8/README.md) |
| 23&#160;Feb&#160;23&#160;14:51&#160;UTC | SHAKEN 186K | 02&#160;Mar&#160;23&#160;14:51&#160;UTC | true | [view](CERTS/31b7f430577d5fed9cffbe626e6187fad963e984fc4e92a4511ac4d80374ae90/README.md) |
| 23&#160;Feb&#160;23&#160;17:10&#160;UTC | SHAKEN 660C | 02&#160;Mar&#160;23&#160;17:10&#160;UTC | true | [view](CERTS/a2d3e6b1956852107d2d9566e689aa3459dd5f3e2092eca0f505c2e0a5c09094/README.md) |
| 23&#160;Feb&#160;23&#160;17:34&#160;UTC | SHAKEN 107K | 02&#160;Mar&#160;23&#160;17:34&#160;UTC | true | [view](CERTS/107858a0e9dcc3dc7329ab90302a08febb1cd99875eb0a2c25ac2ba8dd5cf85a/README.md) |
| 23&#160;Feb&#160;23&#160;18:21&#160;UTC | SHAKEN 060K | 02&#160;Mar&#160;23&#160;18:21&#160;UTC | true | [view](CERTS/b9134765b897e6b790a7729835f7e2682e069cfb85c1ab1b1119f35600353689/README.md) |
| 23&#160;Feb&#160;23&#160;18:46&#160;UTC | SHAKEN 089K | 02&#160;Mar&#160;23&#160;18:46&#160;UTC | true | [view](CERTS/dfd22f8d92e9100d1fe2d9febe0158d0e2c51b5bad4a3b49dd4bccb7a7156cdf/README.md) |
| 23&#160;Feb&#160;23&#160;20:28&#160;UTC | SHAKEN 674J | 02&#160;Mar&#160;23&#160;20:28&#160;UTC | true | [view](CERTS/0feb4287179ab9d7f7cd9e9b84b2d221712763b9e248d593b38d56cbb7ec237d/README.md) |
| 23&#160;Feb&#160;23&#160;20:29&#160;UTC | SHAKEN 738J | 02&#160;Mar&#160;23&#160;20:29&#160;UTC | true | [view](CERTS/d7b61fdda8b58100d87ac5646f44eb852bf7fa24d8e88123e8586e900bf9dfee/README.md) |
| 23&#160;Feb&#160;23&#160;20:30&#160;UTC | SHAKEN 7453 | 02&#160;Mar&#160;23&#160;20:30&#160;UTC | true | [view](CERTS/d79cca8693c8f782f81d0509f3a9775e532c56a99dbc80bc0359dfb64d3b04b2/README.md) |
| 23&#160;Feb&#160;23&#160;20:31&#160;UTC | SHAKEN 819H | 02&#160;Mar&#160;23&#160;20:31&#160;UTC | true | [view](CERTS/be83c93e27a8ce85183e20453b75dc9c59089fd5248d7ccc4a52d881db36586e/README.md) |
| 23&#160;Feb&#160;23&#160;20:32&#160;UTC | SHAKEN 769J | 02&#160;Mar&#160;23&#160;20:32&#160;UTC | true | [view](CERTS/27473b227dc2a879ee34187802dc2dc67285e93f8f0f51d6568862b7a2c34eb2/README.md) |
| 23&#160;Feb&#160;23&#160;20:35&#160;UTC | SHAKEN 691A | 02&#160;Mar&#160;23&#160;20:35&#160;UTC | true | [view](CERTS/856977a265bed68e5e7cb7f60477d97cf43fb4aa401aad583ea69025d0b02c49/README.md) |
| 23&#160;Feb&#160;23&#160;20:35&#160;UTC | SHAKEN 849J | 02&#160;Mar&#160;23&#160;20:35&#160;UTC | true | [view](CERTS/62c7cc1ad0e60bd9f0185f10baabf2294310260f0764363f9c19dbebe83e519a/README.md) |
| 23&#160;Feb&#160;23&#160;20:35&#160;UTC | SHAKEN 726J | 02&#160;Mar&#160;23&#160;20:35&#160;UTC | true | [view](CERTS/987ce6696a80e1172ff3558153272f8b286579933089112208cfc5e9a31c27af/README.md) |
| 23&#160;Feb&#160;23&#160;20:35&#160;UTC | SHAKEN 469A | 02&#160;Mar&#160;23&#160;20:35&#160;UTC | true | [view](CERTS/ee86a6290814dfb347acd293cccbce1a421e83843a843dbe75784715202d9063/README.md) |
| 23&#160;Feb&#160;23&#160;20:35&#160;UTC | SHAKEN 790J | 02&#160;Mar&#160;23&#160;20:35&#160;UTC | true | [view](CERTS/f5fff9e6bac5acb888435212be1e25261bb20d13e1a102ce957a431fb111ede2/README.md) |
| 23&#160;Feb&#160;23&#160;20:36&#160;UTC | SHAKEN 738J | 02&#160;Mar&#160;23&#160;20:36&#160;UTC | true | [view](CERTS/12343acfa20a9b8a14c94ddf32cb91403e58942df0505241fc7f3e820e964fd1/README.md) |
| 23&#160;Feb&#160;23&#160;20:36&#160;UTC | SHAKEN 459J | 02&#160;Mar&#160;23&#160;20:36&#160;UTC | true | [view](CERTS/e6c652abbbb8730fc0e8a81619fa03030152c9303c094a832a483a486c0c80f0/README.md) |
| 23&#160;Feb&#160;23&#160;20:37&#160;UTC | SHAKEN 366G | 02&#160;Mar&#160;23&#160;20:37&#160;UTC | true | [view](CERTS/48dd42abb861c42e8b845e4fe21b887b336ca680a6d089dc04f19f4006ad25f6/README.md) |
| 23&#160;Feb&#160;23&#160;20:38&#160;UTC | SHAKEN 738J | 02&#160;Mar&#160;23&#160;20:38&#160;UTC | true | [view](CERTS/73c08a77b2a9bffd70329125ab30592fa04c979604d83e10de566af4257765f2/README.md) |
| 24&#160;Feb&#160;23&#160;20:25&#160;UTC | SHAKEN 983J | 03&#160;Mar&#160;23&#160;20:25&#160;UTC | true | [view](CERTS/40a47e7c11b390b6784e598d73053876fd7e6e222b76e216e7c1353c4f7f7872/README.md) |
| 24&#160;Feb&#160;23&#160;21:51&#160;UTC | SHAKEN 606F | 03&#160;Mar&#160;23&#160;21:51&#160;UTC | true | [view](CERTS/255308b0e2b5617ae948cd2e46e9ce6aefd860010c305528ff2d0281d661d83b/README.md) |
| 25&#160;Feb&#160;23&#160;01:48&#160;UTC | SHAKEN 278K | 04&#160;Mar&#160;23&#160;01:48&#160;UTC | true | [view](CERTS/c6bc5cd1fa45e21231106ad38e8b02d58ee57ceb5075b47dbdd01b678c60331c/README.md) |
| 25&#160;Feb&#160;23&#160;04:32&#160;UTC | SHAKEN 345J | 27&#160;Mar&#160;23&#160;04:32&#160;UTC | true | [view](CERTS/0b6e254205072f53c65dc8afdaa9ea170fd1d3ac3fa99c7b6636546b91b8cde6/README.md) |
| 25&#160;Feb&#160;23&#160;05:10&#160;UTC | SHAKEN 345J | 27&#160;Mar&#160;23&#160;05:10&#160;UTC | true | [view](CERTS/1bc9676b5e47ed33ba1ce088f65d28fa48fef5cf80621e52eef0f90148b9d337/README.md) |
| 25&#160;Feb&#160;23&#160;14:45&#160;UTC | SHAKEN 2550 | 04&#160;Mar&#160;23&#160;14:45&#160;UTC | true | [view](CERTS/002a8181afd31bb2ca078c74e10647f54fb655835f721f1173fffb8d89a20a5d/README.md) |
| 25&#160;Feb&#160;23&#160;20:21&#160;UTC | SHAKEN 177K | 04&#160;Mar&#160;23&#160;20:21&#160;UTC | true | [view](CERTS/69404cf714510293714925fe5c933081430471e50e3a336eb6dfe837468cc108/README.md) |
| 26&#160;Feb&#160;23&#160;14:51&#160;UTC | SHAKEN 012K | 05&#160;Mar&#160;23&#160;14:50&#160;UTC | true | [view](CERTS/b40b3931dbbf5be7df381b3abb1648d0c9518c0d2e4bbebb8a1c336dfd409e31/README.md) |
| 26&#160;Feb&#160;23&#160;14:52&#160;UTC | SHAKEN 186K | 05&#160;Mar&#160;23&#160;14:52&#160;UTC | true | [view](CERTS/98879e892c05aa60bafb8b7cce6a42eba2e65bd787345c6d421851153a9ffcab/README.md) |
| 26&#160;Feb&#160;23&#160;17:34&#160;UTC | SHAKEN 107K | 05&#160;Mar&#160;23&#160;17:34&#160;UTC | true | [view](CERTS/fd842368007fa55f4039040f67dee04ffb829ca241174e01096298bb877b76e3/README.md) |
| 26&#160;Feb&#160;23&#160;18:21&#160;UTC | SHAKEN 060K | 05&#160;Mar&#160;23&#160;18:21&#160;UTC | true | [view](CERTS/51e1de8f320f937800373d6256e8f9e504d7b0af2b61e9a01c217b9f9c75c7a0/README.md) |
| 26&#160;Feb&#160;23&#160;18:46&#160;UTC | SHAKEN 089K | 05&#160;Mar&#160;23&#160;18:46&#160;UTC | true | [view](CERTS/a0bc3ee9313b3077df01991997200a7d77e6e79bfa04e1707a07cf73b243748f/README.md) |
| 26&#160;Feb&#160;23&#160;18:49&#160;UTC | SHAKEN 9714 | 05&#160;Mar&#160;23&#160;18:49&#160;UTC | true | [view](CERTS/d83ac5d62c9e388d1f973a8d84c5bc09063e5afc65e4f70ec4147a773112d354/README.md) |
| 26&#160;Feb&#160;23&#160;20:05&#160;UTC | SHAKEN 297K | 05&#160;Mar&#160;23&#160;20:05&#160;UTC | true | [view](CERTS/fcb85d16b1e1ab6b46ccd301dc5de04d35fa683a5d2ec27d0ddaced9f48b4c37/README.md) |
| 26&#160;Feb&#160;23&#160;20:28&#160;UTC | SHAKEN 674J | 05&#160;Mar&#160;23&#160;20:28&#160;UTC | true | [view](CERTS/b68ceba25cf1d3180d65628bb5df751a261e517d81a7f771604584410f38c597/README.md) |
| 26&#160;Feb&#160;23&#160;20:29&#160;UTC | SHAKEN 738J | 05&#160;Mar&#160;23&#160;20:29&#160;UTC | true | [view](CERTS/40317a5eb09e15cadd0d6c065db8dab81272fff8f18a68f31daa49c37204b099/README.md) |
| 26&#160;Feb&#160;23&#160;20:30&#160;UTC | SHAKEN 700H | 05&#160;Mar&#160;23&#160;20:30&#160;UTC | true | [view](CERTS/29e8fd6b1abce1cd50a38f8392021c90bad912dfb9a213adccb069b978d5af95/README.md) |
| 26&#160;Feb&#160;23&#160;20:32&#160;UTC | SHAKEN 819H | 05&#160;Mar&#160;23&#160;20:32&#160;UTC | true | [view](CERTS/6b5dbe220c71a4c914752176c3e21c1dccaed60b53cbe0fdc61df01b6b6904cc/README.md) |
| 26&#160;Feb&#160;23&#160;20:33&#160;UTC | SHAKEN 053G | 05&#160;Mar&#160;23&#160;20:33&#160;UTC | true | [view](CERTS/e4393fdb061c03d4f9ce7a2ac2f36e21e91d3948597ada11632a4893ea2ef34d/README.md) |
| 26&#160;Feb&#160;23&#160;20:33&#160;UTC | SHAKEN 769J | 05&#160;Mar&#160;23&#160;20:33&#160;UTC | true | [view](CERTS/c7d3883edd8f6f4e1079d0fbd634124de91cf3d2296b4480fa1b3095e9e3ff3e/README.md) |
| 26&#160;Feb&#160;23&#160;20:35&#160;UTC | SHAKEN 849J | 05&#160;Mar&#160;23&#160;20:35&#160;UTC | true | [view](CERTS/af780d6e12f3311aac90a23246a8e1a4381f340b6d0698395556ff965008f6e8/README.md) |
| 26&#160;Feb&#160;23&#160;20:35&#160;UTC | SHAKEN 469A | 05&#160;Mar&#160;23&#160;20:35&#160;UTC | true | [view](CERTS/4948a7df93a415449d99f02f9403120b4311e6582b213a22f481daec45f0ad0a/README.md) |
| 26&#160;Feb&#160;23&#160;20:35&#160;UTC | SHAKEN 495J | 05&#160;Mar&#160;23&#160;20:35&#160;UTC | true | [view](CERTS/3a30212380f4f5b6274e6c048294770192379841343bd73a34c45d5067d6a337/README.md) |
| 26&#160;Feb&#160;23&#160;20:35&#160;UTC | SHAKEN 721J | 05&#160;Mar&#160;23&#160;20:35&#160;UTC | true | [view](CERTS/46666285b1f5791ad08231e5aabe7638eb0d49aad637eae93c407f4c6eaa2abe/README.md) |
| 26&#160;Feb&#160;23&#160;20:35&#160;UTC | SHAKEN 790J | 05&#160;Mar&#160;23&#160;20:35&#160;UTC | true | [view](CERTS/3e533669707740f51a447f4299125e16b63b77e8393856fbf91b210d65d43062/README.md) |
| 26&#160;Feb&#160;23&#160;20:36&#160;UTC | SHAKEN 625J | 05&#160;Mar&#160;23&#160;20:36&#160;UTC | true | [view](CERTS/29cde07873a76cb2fa531a3d77d9f0cb0683886a708299907f5e4beee739b85b/README.md) |
| 26&#160;Feb&#160;23&#160;20:36&#160;UTC | SHAKEN 738J | 05&#160;Mar&#160;23&#160;20:36&#160;UTC | true | [view](CERTS/b876c59a1e2e63e7330321dee43ec2efed0d38a5b5ab577448a3950b109bda9c/README.md) |
| 26&#160;Feb&#160;23&#160;20:36&#160;UTC | SHAKEN 459J | 05&#160;Mar&#160;23&#160;20:36&#160;UTC | true | [view](CERTS/67e082877dc917867a0a80b8f8ade5efd0960e8e479b7fa262fcecdb285790ce/README.md) |
| 26&#160;Feb&#160;23&#160;20:36&#160;UTC | SHAKEN 672B | 05&#160;Mar&#160;23&#160;20:36&#160;UTC | true | [view](CERTS/a654bf60e38a381b7a7086dd00638c1d5203ddd690727bbe26949f2b150bc827/README.md) |
| 26&#160;Feb&#160;23&#160;20:37&#160;UTC | SHAKEN 366G | 05&#160;Mar&#160;23&#160;20:37&#160;UTC | true | [view](CERTS/0f14106c6ea2d647a9a2df6f09e1c8b7098641eacb75bce00b34c7ac9f8dc0a4/README.md) |
| 26&#160;Feb&#160;23&#160;20:37&#160;UTC | SHAKEN 0226 | 05&#160;Mar&#160;23&#160;20:37&#160;UTC | true | [view](CERTS/e6d9fb626cc80e9c50a298faa90ddacc2e386ee3337d921ea566c9b8d8e923d8/README.md) |
| 26&#160;Feb&#160;23&#160;20:38&#160;UTC | SHAKEN 738J | 05&#160;Mar&#160;23&#160;20:38&#160;UTC | true | [view](CERTS/ce4e3b57b9d27a113d14551c300fd1c1ff48819279d92efd3eb9688a0d88301c/README.md) |
| 26&#160;Feb&#160;23&#160;22:10&#160;UTC | SHAKEN 079C | 05&#160;Mar&#160;23&#160;22:10&#160;UTC | true | [view](CERTS/1aa475c97978f033ef912b282db6fce6009d5779431b74c60e7b9185fd6a38a2/README.md) |
| 27&#160;Feb&#160;23&#160;00:01&#160;UTC | SHAKEN 0172 | 06&#160;Mar&#160;23&#160;00:01&#160;UTC | true | [view](CERTS/aa33605ed7ae80358576aebd26f5819d1c85d016494c5beeacb9c836f4eaf560/README.md) |
| 27&#160;Feb&#160;23&#160;01:43&#160;UTC | SHAKEN 807J | 28&#160;Apr&#160;23&#160;01:43&#160;UTC | true | [view](CERTS/ba83be2505aeab5690525902158ec1372b1b8917d4db12e9e43dceee579dadfb/README.md) |
| 27&#160;Feb&#160;23&#160;12:22&#160;UTC | SHAKEN 159H | 06&#160;Mar&#160;23&#160;12:22&#160;UTC | true | [view](CERTS/cf91c66727a7c0037c1ad545d73053168c57fd04ef1a24c957df3136ca643a97/README.md) |
| 27&#160;Feb&#160;23&#160;18:13&#160;UTC | SHAKEN 406K | 06&#160;Mar&#160;23&#160;18:13&#160;UTC | true | [view](CERTS/43a39853c82e045b6bd0d7740ecb8a782c1b700c7a7e188dd47a15bf5f8d2b7c/README.md) |
| 27&#160;Feb&#160;23&#160;19:17&#160;UTC | SHAKEN 749J | 06&#160;Mar&#160;23&#160;19:17&#160;UTC | true | [view](CERTS/ae376628ddc6445f6fb2223b66874a2cfb02514cece488475f59e459114a089d/README.md) |
| 27&#160;Feb&#160;23&#160;20:25&#160;UTC | SHAKEN 983J | 06&#160;Mar&#160;23&#160;20:25&#160;UTC | true | [view](CERTS/1412911a57dd3ec965dd51eeaeb526d912c88bd92dffab34c41e73df1345d262/README.md) |
| 27&#160;Feb&#160;23&#160;20:39&#160;UTC | SHAKEN 738J | 08&#160;Mar&#160;23&#160;20:39&#160;UTC | true | [view](CERTS/22e447ee1a278877f2e0f3956ceb04e71fa316aceb8ec68e8fb0300e21ca9ff8/README.md) |
| 28&#160;Feb&#160;23&#160;01:48&#160;UTC | SHAKEN 278K | 07&#160;Mar&#160;23&#160;01:48&#160;UTC | true | [view](CERTS/a182e17e759a565794b09f1237bdcc611159598568d6d8a6e60d73f6e3d9e28e/README.md) |
| 28&#160;Feb&#160;23&#160;14:45&#160;UTC | SHAKEN 2550 | 07&#160;Mar&#160;23&#160;14:45&#160;UTC | true | [view](CERTS/27e33531866ac0cee918c39a4b7702e87abbfd6543e50a31c0fa702f3fc5a225/README.md) |
| 28&#160;Feb&#160;23&#160;15:11&#160;UTC | SHAKEN 9595 | 07&#160;Mar&#160;23&#160;15:11&#160;UTC | true | [view](CERTS/85a84a9fa636c93cc582c501f4e3b208b03b8efa0fbe983636aac8e7a647ff7a/README.md) |

#### CA Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 20&#160;Aug&#160;21&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Issuing CA3 | 19&#160;Aug&#160;31&#160;23:59&#160;UTC | true | [view](CERTS/51b51431ffd8055cd7e30c213540f4868258058d4c4c9f23f1ed8bea75101e37/README.md) |
| 20&#160;Aug&#160;21&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Issuing CA1 | 19&#160;Aug&#160;31&#160;23:59&#160;UTC | true | [view](CERTS/037a7c6621814bd5f01fa2080a77a245458694d9d566e0050e068fc853317857/README.md) |
| 20&#160;Aug&#160;21&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Root CA1 | 19&#160;Aug&#160;41&#160;23:59&#160;UTC | true | [view](CERTS/59ab943f4e8c3c17755cc8e4abb5ac65736ce74ec587c62cb744ab0babbbe2fe/README.md) |
| 24&#160;Oct&#160;22&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Issuing CA4 | 23&#160;Oct&#160;32&#160;23:59&#160;UTC | false | [view](CERTS/80dd12d935f46c256c7c1289a0834b21ec8f2f8c35f2864928d1e75f3a280665/README.md) |
| 24&#160;Oct&#160;22&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Root CA2 | 23&#160;Oct&#160;42&#160;23:59&#160;UTC | false | [view](CERTS/a26e04fc786ab70b8085236b2c53f8cfbf5d0c6a5c2c9c3e9f91669fbb8ea4d5/README.md) |


Generated: 01 Mar 23 18:22 UTC