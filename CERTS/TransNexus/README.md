# STIR/SHAKEN CA Ecosystem Compliance

## TransNexus

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 1824 certificates were included in the corpus being tested
- 1 certificates in the corpus were skipped because they are duplicates
- 1729 certificates in the corpus were skipped because they are expired
- 1 certificates in the corpus were skipped because they are not currently trusted
- 93 certificates being tested against the remaining rules
- 1.24 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 9.68% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 63 days is the average remaining validity for the certificates in the corpus
- 63 days is the average initial validity for the certificates in the corpus
- 77 certificates expire in the next 30 days
- 1.45 average number of unexpired certificates per OCN observed
- 64 unique OCNs observed in unexpired and valid certificate corpus

| Instances | Test | Source |
|-----------|------|--------|
| 2 | [e_atis_certificate_policies](ISSUES/e_atis_certificate_policies/README.md) | ATIS1000080 |
| 9 | [e_incorrect_ku_encoding](ISSUES/e_incorrect_ku_encoding/README.md) | RFC5280 |
| 2 | [e_us_cp_ambiguous_identifier](ISSUES/e_us_cp_ambiguous_identifier/README.md) | US_SHAKEN_CP |
| 93 | [e_us_cp_subject_sn](ISSUES/e_us_cp_subject_sn/README.md) | US_SHAKEN_CP |
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
- 5455 days is the average remaining validity for the certificates in the corpus
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
| 19&#160;Feb&#160;23&#160;15:03&#160;UTC | SHAKEN 577F | 21&#160;Mar&#160;23&#160;15:03&#160;UTC | true | [view](CERTS/9d063726d179f92c7f9438afcf499083ea14b81851800001fa529ece9610253d/README.md) |
| 21&#160;Feb&#160;23&#160;16:41&#160;UTC | SHAKEN 551G | 23&#160;Mar&#160;23&#160;16:41&#160;UTC | true | [view](CERTS/108270b9fe753c5036472d046f78a2ba12a48bd6a1da0d55df207fdb907d9821/README.md) |
| 22&#160;Feb&#160;23&#160;15:54&#160;UTC | SHAKEN 722J | 24&#160;Mar&#160;23&#160;15:54&#160;UTC | true | [view](CERTS/cc3c678774a06d4a1cf0f216048f71627afaeb92d62800d19fed64c029a1a729/README.md) |
| 22&#160;Feb&#160;23&#160;18:40&#160;UTC | SHAKEN 952J | 24&#160;Mar&#160;23&#160;18:40&#160;UTC | true | [view](CERTS/52a3abc5046e0441a8fa5c7dbe5f04f79db1c75a20fec50f377b5cf45b5f4b4f/README.md) |
| 25&#160;Feb&#160;23&#160;04:32&#160;UTC | SHAKEN 345J | 27&#160;Mar&#160;23&#160;04:32&#160;UTC | true | [view](CERTS/0b6e254205072f53c65dc8afdaa9ea170fd1d3ac3fa99c7b6636546b91b8cde6/README.md) |
| 25&#160;Feb&#160;23&#160;05:10&#160;UTC | SHAKEN 345J | 27&#160;Mar&#160;23&#160;05:10&#160;UTC | true | [view](CERTS/1bc9676b5e47ed33ba1ce088f65d28fa48fef5cf80621e52eef0f90148b9d337/README.md) |
| 27&#160;Feb&#160;23&#160;01:43&#160;UTC | SHAKEN 807J | 28&#160;Apr&#160;23&#160;01:43&#160;UTC | true | [view](CERTS/ba83be2505aeab5690525902158ec1372b1b8917d4db12e9e43dceee579dadfb/README.md) |
| 27&#160;Feb&#160;23&#160;23:03&#160;UTC | SHAKEN 4632 | 29&#160;Mar&#160;23&#160;23:03&#160;UTC | true | [view](CERTS/9cd3b7456e12980f51974742f457f3f4864cb85818bb882a8d7a8411617add72/README.md) |
| 03&#160;Mar&#160;23&#160;19:51&#160;UTC | SHAKEN 841J | 17&#160;Mar&#160;23&#160;19:51&#160;UTC | true | [view](CERTS/f1ebcdf94e15084f39ab437b14d37ebadee3067e70110c026ec7fee3f3e47ad5/README.md) |
| 09&#160;Mar&#160;23&#160;18:46&#160;UTC | SHAKEN 193E | 08&#160;May&#160;23&#160;18:46&#160;UTC | true | [view](CERTS/34187f020bbbd3596eea83f5f2243b84920396d2c46eed5b8078a0b696c383e5/README.md) |
| 09&#160;Mar&#160;23&#160;20:22&#160;UTC | SHAKEN 177K | 16&#160;Mar&#160;23&#160;20:22&#160;UTC | true | [view](CERTS/e827bd734d605b4a766e8beefb628abd84ebaa9b1be8e6ad241711affcd8fac6/README.md) |
| 10&#160;Mar&#160;23&#160;14:52&#160;UTC | SHAKEN 056J | 17&#160;Mar&#160;23&#160;14:52&#160;UTC | true | [view](CERTS/288336d042d1367fca11bc5d94d80ef2bf959b2794766c9ea738ca712f3aba2d/README.md) |
| 10&#160;Mar&#160;23&#160;17:35&#160;UTC | SHAKEN 107K | 17&#160;Mar&#160;23&#160;17:35&#160;UTC | true | [view](CERTS/b71462cf77b5d7d12aae7dbeb1145acd53ed2aaa9e10de248feccb7b10c4f5e4/README.md) |
| 10&#160;Mar&#160;23&#160;18:49&#160;UTC | SHAKEN 9714 | 17&#160;Mar&#160;23&#160;18:49&#160;UTC | true | [view](CERTS/303d7d9fd337111f423db6034581a4128b32b373c8f1284cfea9d628eb37544a/README.md) |
| 10&#160;Mar&#160;23&#160;20:06&#160;UTC | SHAKEN 297K | 17&#160;Mar&#160;23&#160;20:06&#160;UTC | true | [view](CERTS/bc029a81d749f4796b4657ba9eddf854307933fc657ee2c6647eecfdb75efe17/README.md) |
| 10&#160;Mar&#160;23&#160;20:29&#160;UTC | SHAKEN 674J | 17&#160;Mar&#160;23&#160;20:29&#160;UTC | true | [view](CERTS/1f80ace0ef2ffb2a6ae02843d99331378751874e0a0a8e5ebd61a4db96e01a2f/README.md) |
| 10&#160;Mar&#160;23&#160;20:30&#160;UTC | SHAKEN 738J | 17&#160;Mar&#160;23&#160;20:30&#160;UTC | true | [view](CERTS/c4ea048d5493f74360d6db4559855479522416a57e8666962c71e54e671cf258/README.md) |
| 10&#160;Mar&#160;23&#160;20:31&#160;UTC | SHAKEN 700H | 17&#160;Mar&#160;23&#160;20:31&#160;UTC | true | [view](CERTS/4fe78123e89b197eec192d79d73e56a7293621dcd490c588a6b2e597a0991060/README.md) |
| 10&#160;Mar&#160;23&#160;20:31&#160;UTC | SHAKEN 733J | 17&#160;Mar&#160;23&#160;20:31&#160;UTC | true | [view](CERTS/b6b5f611c09e3052a0c2b65cfd91920352a6059d9ec4208a47d56d66374fa1dd/README.md) |
| 10&#160;Mar&#160;23&#160;20:31&#160;UTC | SHAKEN 7453 | 17&#160;Mar&#160;23&#160;20:31&#160;UTC | true | [view](CERTS/a6d7dabbd69e0821214ae85fa164100d3ec86af1af02c3623795a04f84b5a643/README.md) |
| 10&#160;Mar&#160;23&#160;20:33&#160;UTC | SHAKEN 769J | 17&#160;Mar&#160;23&#160;20:33&#160;UTC | true | [view](CERTS/94dce1f4c23a17836376fc533dfb0e2aa6bf29747069e5b69e783f7132b2b602/README.md) |
| 10&#160;Mar&#160;23&#160;20:36&#160;UTC | SHAKEN 691J | 17&#160;Mar&#160;23&#160;20:36&#160;UTC | true | [view](CERTS/1089a781a03f289341821e9404eff3156bedf5466c900956d332fe24b542ae78/README.md) |
| 10&#160;Mar&#160;23&#160;20:36&#160;UTC | SHAKEN 469A | 17&#160;Mar&#160;23&#160;20:36&#160;UTC | true | [view](CERTS/da7caabcb537f475599e51e084c4dfb4eb72a9b4819ccd6792a57d0635c58fc2/README.md) |
| 10&#160;Mar&#160;23&#160;20:37&#160;UTC | SHAKEN 625J | 17&#160;Mar&#160;23&#160;20:37&#160;UTC | true | [view](CERTS/399235641df038df0ea9d33af5b2de87a9f068be079db01329c644b859e532d3/README.md) |
| 10&#160;Mar&#160;23&#160;20:37&#160;UTC | SHAKEN 738J | 17&#160;Mar&#160;23&#160;20:37&#160;UTC | true | [view](CERTS/299a317ba8b7c8d8e41660c859040408c33c42c592da9f00c1c1082652092390/README.md) |
| 10&#160;Mar&#160;23&#160;20:37&#160;UTC | SHAKEN 459J | 17&#160;Mar&#160;23&#160;20:37&#160;UTC | true | [view](CERTS/7281f54c69ac0407d298a06018cfa87d72e14632e59a4366247002dc39cc5717/README.md) |
| 10&#160;Mar&#160;23&#160;20:38&#160;UTC | SHAKEN 366G | 17&#160;Mar&#160;23&#160;20:38&#160;UTC | true | [view](CERTS/edf4c2b2833c3a025bbe3570fe536f433f519f92fdd8929618cc908fc66d6f19/README.md) |
| 10&#160;Mar&#160;23&#160;20:40&#160;UTC | SHAKEN 738J | 17&#160;Mar&#160;23&#160;20:40&#160;UTC | true | [view](CERTS/85205524552f43027b321a9979e7c2b3bf7bff473e2c9a37cb0ee991ce7158f9/README.md) |
| 11&#160;Mar&#160;23&#160;05:01&#160;UTC | SHAKEN 982J | 18&#160;Mar&#160;23&#160;05:01&#160;UTC | true | [view](CERTS/4e44a39d04d8729b5fbd7847a7b5ef3dcd79f32097fffc105a1cf3e6b9671a63/README.md) |
| 11&#160;Mar&#160;23&#160;14:32&#160;UTC | SHAKEN 0172 | 18&#160;Mar&#160;23&#160;14:32&#160;UTC | true | [view](CERTS/e01e1c8e1ee045adc449b9b37cac65ce77f6753562209df0a2b084d5fa1598d8/README.md) |
| 11&#160;Mar&#160;23&#160;20:26&#160;UTC | SHAKEN 983J | 18&#160;Mar&#160;23&#160;20:26&#160;UTC | true | [view](CERTS/318d4b429a920a3ed45f34913720e34b31e7fd2378cb58dfb6d34162635f8c2d/README.md) |
| 11&#160;Mar&#160;23&#160;21:53&#160;UTC | SHAKEN 606F | 18&#160;Mar&#160;23&#160;21:52&#160;UTC | true | [view](CERTS/d4c80783815ba561afcf7423bba8c409c114ba6bf9a812d16c6042f5c41f0210/README.md) |
| 11&#160;Mar&#160;23&#160;22:35&#160;UTC | SHAKEN 120K | 18&#160;Mar&#160;23&#160;22:35&#160;UTC | true | [view](CERTS/9d8cfdd31fe7d107b61011e497bac3c138eb87aa802cd91cea774e3c9cde72df/README.md) |
| 12&#160;Mar&#160;23&#160;01:49&#160;UTC | SHAKEN 278K | 19&#160;Mar&#160;23&#160;01:49&#160;UTC | true | [view](CERTS/04d6e1c0abeeca0b677a46f9c30e00cda81313615cca5ce4583b9d7d1e58da9e/README.md) |
| 12&#160;Mar&#160;23&#160;05:32&#160;UTC | SHAKEN 841J | 26&#160;Mar&#160;23&#160;05:32&#160;UTC | true | [view](CERTS/455097278a1b024630523ffdc1dbed0ceaefff2a37459b827b5ed325b1a78ace/README.md) |
| 12&#160;Mar&#160;23&#160;14:45&#160;UTC | SHAKEN 2550 | 19&#160;Mar&#160;23&#160;14:45&#160;UTC | true | [view](CERTS/f482daca61529e2a55adf52300e5aedb8d9ddde3f71762948a35c44326628a91/README.md) |
| 12&#160;Mar&#160;23&#160;15:30&#160;UTC | SHAKEN 284K | 19&#160;Mar&#160;23&#160;15:30&#160;UTC | true | [view](CERTS/23bf3701cb0e8ee33e091bbd46e9d3e09a070bd1463b70ffb576bc2b9877a068/README.md) |
| 12&#160;Mar&#160;23&#160;17:38&#160;UTC | SHAKEN 833J | 19&#160;Mar&#160;23&#160;17:38&#160;UTC | true | [view](CERTS/a6121db28b321b83fe1753f33c49fe3bb26b26db6f3f40cc6dc2fbb1ea41f6dc/README.md) |
| 12&#160;Mar&#160;23&#160;18:44&#160;UTC | SHAKEN 952J | 11&#160;Apr&#160;23&#160;18:44&#160;UTC | true | [view](CERTS/8825c9d169a6a46d061e253597a7c52050715658538811ddbc13807df3415fac/README.md) |
| 12&#160;Mar&#160;23&#160;20:23&#160;UTC | SHAKEN 177K | 19&#160;Mar&#160;23&#160;20:23&#160;UTC | true | [view](CERTS/145c7b0f7b1e5b6031240f58e710b2d62380f42243e54522a36eac4a33586789/README.md) |
| 13&#160;Mar&#160;23&#160;14:52&#160;UTC | SHAKEN 012K | 20&#160;Mar&#160;23&#160;14:52&#160;UTC | true | [view](CERTS/aa0023330082f7a735bd7f01f5fb13bb782923080e738257685cb223d1b1410f/README.md) |
| 13&#160;Mar&#160;23&#160;17:35&#160;UTC | SHAKEN 107K | 20&#160;Mar&#160;23&#160;17:35&#160;UTC | true | [view](CERTS/4782e5d60013b0565dbf40835c9bdd2a28413ebdee02c9ce4c9e2a156a5fe031/README.md) |
| 13&#160;Mar&#160;23&#160;18:23&#160;UTC | SHAKEN 060K | 20&#160;Mar&#160;23&#160;18:23&#160;UTC | true | [view](CERTS/d82adeeb2cb8c7ccd55cf023aacc8349e27cb25cad251f45072f3e32bbdda9dd/README.md) |
| 13&#160;Mar&#160;23&#160;19:20&#160;UTC | SHAKEN 019K | 20&#160;Mar&#160;23&#160;19:20&#160;UTC | true | [view](CERTS/b4b524597421a87d5bc83cf702b603b27d8059293131ebea8d3cd48d655c8a3a/README.md) |
| 13&#160;Mar&#160;23&#160;20:06&#160;UTC | SHAKEN 297K | 20&#160;Mar&#160;23&#160;20:06&#160;UTC | true | [view](CERTS/6eb76e5d3163ee5bbb4395e187c0ff1ea0bf0b7002ac0f297c3eb7e693528d9b/README.md) |
| 13&#160;Mar&#160;23&#160;20:29&#160;UTC | SHAKEN 674J | 20&#160;Mar&#160;23&#160;20:29&#160;UTC | true | [view](CERTS/191f5081a07a0429f3cdfac3fe608a0cb2ddaab16fbbbaf67370209983aeeac0/README.md) |
| 13&#160;Mar&#160;23&#160;20:30&#160;UTC | SHAKEN 5512 | 20&#160;Mar&#160;23&#160;20:30&#160;UTC | true | [view](CERTS/30f6d9fba1b64c4fd460e8b2d710acfe9ac01e8491e2275aea625c4cae5278d0/README.md) |
| 13&#160;Mar&#160;23&#160;20:30&#160;UTC | SHAKEN 738J | 20&#160;Mar&#160;23&#160;20:30&#160;UTC | true | [view](CERTS/da02ce87d658ee890d159a7edfbd6df2726c47b97e41aba7c35745afdf0810b4/README.md) |
| 13&#160;Mar&#160;23&#160;20:31&#160;UTC | SHAKEN 700H | 20&#160;Mar&#160;23&#160;20:31&#160;UTC | true | [view](CERTS/772e0de7c783d3ee779efe35af73516e8c129681519af59cf6cebf3502235521/README.md) |
| 13&#160;Mar&#160;23&#160;20:32&#160;UTC | SHAKEN 733J | 20&#160;Mar&#160;23&#160;20:32&#160;UTC | true | [view](CERTS/53e7c8729420cba7851cec39c997f31d8d74b672aff38482f997e094c8eee4ab/README.md) |
| 13&#160;Mar&#160;23&#160;20:32&#160;UTC | SHAKEN 7453 | 20&#160;Mar&#160;23&#160;20:32&#160;UTC | true | [view](CERTS/73014fd8fc8f56d0b95e1f2b8ec22b417213bd63ad24297b303d7edf6b1940dc/README.md) |
| 13&#160;Mar&#160;23&#160;20:33&#160;UTC | SHAKEN 819H | 20&#160;Mar&#160;23&#160;20:33&#160;UTC | true | [view](CERTS/478678f92ef6fafb116908e700623a331527602f9ccd8e8bc5b8f85740571ceb/README.md) |
| 13&#160;Mar&#160;23&#160;20:34&#160;UTC | SHAKEN 0753 | 20&#160;Mar&#160;23&#160;20:34&#160;UTC | true | [view](CERTS/2cae42a8a119865957286380bbca8b995f9991bf04f5f9f0e0f512d995dbdb3d/README.md) |
| 13&#160;Mar&#160;23&#160;20:34&#160;UTC | SHAKEN 769J | 20&#160;Mar&#160;23&#160;20:34&#160;UTC | true | [view](CERTS/bea181b8350bbb4faf7a35eb3fba436a91f1f149e219a2a7e2afaa8672925d68/README.md) |
| 13&#160;Mar&#160;23&#160;20:36&#160;UTC | SHAKEN 691J | 20&#160;Mar&#160;23&#160;20:36&#160;UTC | true | [view](CERTS/efe59d71c9d990f5c928b0523334de512d6d7cb9fb36fe81465149279db9e8bb/README.md) |
| 13&#160;Mar&#160;23&#160;20:36&#160;UTC | SHAKEN 849J | 20&#160;Mar&#160;23&#160;20:36&#160;UTC | true | [view](CERTS/353d8e2ceba5b71773db7cbed77867cb942a4617b1a5294fb77d8a4ba6ef01e6/README.md) |
| 13&#160;Mar&#160;23&#160;20:36&#160;UTC | SHAKEN 469A | 20&#160;Mar&#160;23&#160;20:36&#160;UTC | true | [view](CERTS/df97d1d5c9c0cb133ea0b2984dfca69a42f3bf7f2bd02c17ad286184c9bd9deb/README.md) |
| 13&#160;Mar&#160;23&#160;20:36&#160;UTC | SHAKEN 495J | 20&#160;Mar&#160;23&#160;20:36&#160;UTC | true | [view](CERTS/d0dccbc7ee929ac683275d7c0d117ebf8baf9d1ec4fb620dab864af241e61bb9/README.md) |
| 13&#160;Mar&#160;23&#160;20:36&#160;UTC | SHAKEN 790J | 20&#160;Mar&#160;23&#160;20:36&#160;UTC | true | [view](CERTS/52d1f7bcfc5e405fa52fd7e4fbd0b1d7d38bea289229695b9d382a23f17e6f87/README.md) |
| 13&#160;Mar&#160;23&#160;20:37&#160;UTC | SHAKEN 625J | 20&#160;Mar&#160;23&#160;20:37&#160;UTC | true | [view](CERTS/56bb7a7dad0c1bed693bc2906c3b03ee26166f947689b1ec725fabf3f105052b/README.md) |
| 13&#160;Mar&#160;23&#160;20:37&#160;UTC | SHAKEN 738J | 20&#160;Mar&#160;23&#160;20:37&#160;UTC | true | [view](CERTS/37aa0a61f16199ff24a32291a657ec9643ae90a141bc794b2ecd019d06036598/README.md) |
| 13&#160;Mar&#160;23&#160;20:38&#160;UTC | SHAKEN 459J | 20&#160;Mar&#160;23&#160;20:38&#160;UTC | true | [view](CERTS/23c8ccb03b322b762b027a01cb9c44796f3a6be4cbe14e8680fcb2a052b94277/README.md) |
| 13&#160;Mar&#160;23&#160;20:38&#160;UTC | SHAKEN 672B | 20&#160;Mar&#160;23&#160;20:38&#160;UTC | true | [view](CERTS/722258d7d5cb6a6b648cd0e6ff28b1686ccab1bbc0a27b4d025dcaa79c859539/README.md) |
| 13&#160;Mar&#160;23&#160;20:39&#160;UTC | SHAKEN 366G | 20&#160;Mar&#160;23&#160;20:39&#160;UTC | true | [view](CERTS/fa87e0c83121f10986acf156cb0f51b32f99cb83f7b9fac219b5d4cd8a0b8b74/README.md) |
| 13&#160;Mar&#160;23&#160;20:39&#160;UTC | SHAKEN 0226 | 20&#160;Mar&#160;23&#160;20:39&#160;UTC | true | [view](CERTS/b654662a1e36dc9ae0c90de1a46a1ac184c3345c355f284087b2d963a10e1566/README.md) |
| 13&#160;Mar&#160;23&#160;20:40&#160;UTC | SHAKEN 738J | 20&#160;Mar&#160;23&#160;20:40&#160;UTC | true | [view](CERTS/d4f35d827da4934575f6dafbf3b641306e9b40d09b8fa572b2dd5d7dfe658078/README.md) |
| 13&#160;Mar&#160;23&#160;20:41&#160;UTC | SHAKEN 738J | 22&#160;Mar&#160;23&#160;20:41&#160;UTC | true | [view](CERTS/ae4a3df8dcb3ff869012cd43827f6cb0b8514c3a236b171609cbd7ea1e335f01/README.md) |
| 13&#160;Mar&#160;23&#160;21:41&#160;UTC | SHAKEN 625J | 20&#160;Mar&#160;23&#160;21:41&#160;UTC | true | [view](CERTS/2b641e0612219dbd17b71b7e026d88a760de04f749e6b0daea7b810bd29884cd/README.md) |
| 14&#160;Mar&#160;23&#160;16:44&#160;UTC | SHAKEN 551G | 13&#160;Apr&#160;23&#160;16:44&#160;UTC | true | [view](CERTS/59dedf353c817690bfe09fbbdb984cbecf13fd9193ea8c309d750e9113cba84b/README.md) |
| 14&#160;Mar&#160;23&#160;16:46&#160;UTC | SHAKEN 324K | 21&#160;Mar&#160;23&#160;16:46&#160;UTC | true | [view](CERTS/274e2b256b09c42d627b61e4659560c3e2864de149b672be91740c637b5f8917/README.md) |
| 14&#160;Mar&#160;23&#160;20:26&#160;UTC | SHAKEN 983J | 21&#160;Mar&#160;23&#160;20:26&#160;UTC | true | [view](CERTS/edebb08f9fcef7cd845ef3c1560667b16acfcb15d5d7b9b6e06f7cc590a0c5f8/README.md) |
| 14&#160;Mar&#160;23&#160;22:35&#160;UTC | SHAKEN 120K | 21&#160;Mar&#160;23&#160;22:35&#160;UTC | true | [view](CERTS/766f80c864d3b24f8bbcb663a97fac66d614b0d71dd636c21eeed070d1eed4ce/README.md) |
| 15&#160;Mar&#160;23&#160;14:46&#160;UTC | SHAKEN 2550 | 22&#160;Mar&#160;23&#160;14:46&#160;UTC | true | [view](CERTS/dadd983d7ce72544d17b423deb9b3bbf5f4faf8ae100ab52b4e94c8f3de83661/README.md) |
| 15&#160;Mar&#160;23&#160;19:21&#160;UTC | SHAKEN 0172 | 22&#160;Mar&#160;23&#160;19:21&#160;UTC | true | [view](CERTS/d21faa8976b349d72c26a6be89245643221652b93e3c95b26ad46ddffd2953a8/README.md) |
| 15&#160;Mar&#160;23&#160;20:23&#160;UTC | SHAKEN 177K | 22&#160;Mar&#160;23&#160;20:23&#160;UTC | true | [view](CERTS/bf2718ac36ec34f3aa8ea589abe33aebfa9014cee681ce5fe8770381a761325c/README.md) |
| 16&#160;Mar&#160;23&#160;01:32&#160;UTC | SHAKEN 186K | 17&#160;Mar&#160;23&#160;01:32&#160;UTC | true | [view](CERTS/16dda60cafb6dbc86e46c1da74526a2a197fa0a65450a7f81b08bdc164a71058/README.md) |

#### CA Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 20&#160;Aug&#160;21&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Issuing CA3 | 19&#160;Aug&#160;31&#160;23:59&#160;UTC | true | [view](CERTS/51b51431ffd8055cd7e30c213540f4868258058d4c4c9f23f1ed8bea75101e37/README.md) |
| 20&#160;Aug&#160;21&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Issuing CA1 | 19&#160;Aug&#160;31&#160;23:59&#160;UTC | true | [view](CERTS/037a7c6621814bd5f01fa2080a77a245458694d9d566e0050e068fc853317857/README.md) |
| 20&#160;Aug&#160;21&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Root CA1 | 19&#160;Aug&#160;41&#160;23:59&#160;UTC | true | [view](CERTS/59ab943f4e8c3c17755cc8e4abb5ac65736ce74ec587c62cb744ab0babbbe2fe/README.md) |
| 24&#160;Oct&#160;22&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Issuing CA4 | 23&#160;Oct&#160;32&#160;23:59&#160;UTC | false | [view](CERTS/80dd12d935f46c256c7c1289a0834b21ec8f2f8c35f2864928d1e75f3a280665/README.md) |
| 24&#160;Oct&#160;22&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Root CA2 | 23&#160;Oct&#160;42&#160;23:59&#160;UTC | false | [view](CERTS/a26e04fc786ab70b8085236b2c53f8cfbf5d0c6a5c2c9c3e9f91669fbb8ea4d5/README.md) |


Generated: 16 Mar 23 19:18 UTC