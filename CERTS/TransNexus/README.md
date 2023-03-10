# STIR/SHAKEN CA Ecosystem Compliance

## TransNexus

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 1748 certificates were included in the corpus being tested
- 1 certificates in the corpus were skipped because they are duplicates
- 1648 certificates in the corpus were skipped because they are expired
- 1 certificates in the corpus were skipped because they are not currently trusted
- 98 certificates being tested against the remaining rules
- 1.22 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 9.18% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 59 days is the average remaining validity for the certificates in the corpus
- 60 days is the average initial validity for the certificates in the corpus
- 81 certificates expire in the next 30 days
- 1.38 average number of unexpired certificates per OCN observed
- 71 unique OCNs observed in unexpired and valid certificate corpus

| Instances | Test | Source |
|-----------|------|--------|
| 2 | [e_atis_certificate_policies](ISSUES/e_atis_certificate_policies/README.md) | ATIS1000080 |
| 9 | [e_incorrect_ku_encoding](ISSUES/e_incorrect_ku_encoding/README.md) | RFC5280 |
| 2 | [e_us_cp_ambiguous_identifier](ISSUES/e_us_cp_ambiguous_identifier/README.md) | US_SHAKEN_CP |
| 98 | [e_us_cp_subject_sn](ISSUES/e_us_cp_subject_sn/README.md) | US_SHAKEN_CP |
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
- 5456 days is the average remaining validity for the certificates in the corpus
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
| 03&#160;Mar&#160;23&#160;04:50&#160;UTC | SHAKEN 0172 | 10&#160;Mar&#160;23&#160;04:50&#160;UTC | true | [view](CERTS/3dbe183420319a03a510afbc3290ab55592ecd587a5f03541c55a3e1e92988f2/README.md) |
| 03&#160;Mar&#160;23&#160;14:45&#160;UTC | SHAKEN 2550 | 10&#160;Mar&#160;23&#160;14:45&#160;UTC | true | [view](CERTS/570fe862536bba372df5a508a9dcfb8eae39b9c1df45484a675246396b81c945/README.md) |
| 03&#160;Mar&#160;23&#160;19:51&#160;UTC | SHAKEN 841J | 17&#160;Mar&#160;23&#160;19:51&#160;UTC | true | [view](CERTS/f1ebcdf94e15084f39ab437b14d37ebadee3067e70110c026ec7fee3f3e47ad5/README.md) |
| 03&#160;Mar&#160;23&#160;20:22&#160;UTC | SHAKEN 177K | 10&#160;Mar&#160;23&#160;20:22&#160;UTC | true | [view](CERTS/d2c4f95c292a7de4da00981caea1fc5eb4aa221b2d203deb7321b94fbdb148be/README.md) |
| 04&#160;Mar&#160;23&#160;14:51&#160;UTC | SHAKEN 012K | 11&#160;Mar&#160;23&#160;14:51&#160;UTC | true | [view](CERTS/ecb27cee419df27c3e735eba84bcb37dc14abc74c9209244c8cc150936f81407/README.md) |
| 04&#160;Mar&#160;23&#160;17:10&#160;UTC | SHAKEN 660C | 11&#160;Mar&#160;23&#160;17:10&#160;UTC | true | [view](CERTS/e5e215f7c77d6086918f93a2e45ec604f65d508b46083f8c19fea3d0e8132496/README.md) |
| 04&#160;Mar&#160;23&#160;17:35&#160;UTC | SHAKEN 107K | 11&#160;Mar&#160;23&#160;17:35&#160;UTC | true | [view](CERTS/ce223fa34112b827085cbd72ae443f1144948143fa9cd5598f906cb7e0fc75b4/README.md) |
| 04&#160;Mar&#160;23&#160;18:22&#160;UTC | SHAKEN 159H | 11&#160;Mar&#160;23&#160;18:22&#160;UTC | true | [view](CERTS/570c741a9b6934e36c3625bd0fb255a2d5f0c044033da4af1fc20e90849994e6/README.md) |
| 04&#160;Mar&#160;23&#160;18:47&#160;UTC | SHAKEN 089K | 11&#160;Mar&#160;23&#160;18:47&#160;UTC | true | [view](CERTS/438498b0b7f0577481c80204fd3a9e095e3a32d66a45bcbce652e4e6be7f2516/README.md) |
| 04&#160;Mar&#160;23&#160;18:47&#160;UTC | SHAKEN 056K | 11&#160;Mar&#160;23&#160;18:47&#160;UTC | true | [view](CERTS/a5e9255e0cdd58a59c03e5a0dc4143251a7083b7a1bbb535e67d4fd73fd9681d/README.md) |
| 04&#160;Mar&#160;23&#160;18:49&#160;UTC | SHAKEN 9714 | 11&#160;Mar&#160;23&#160;18:49&#160;UTC | true | [view](CERTS/2eeeeb378e98fd865e6c43c1caa3baf4ff8a438a9507992ce5d38fcea686dc3c/README.md) |
| 04&#160;Mar&#160;23&#160;20:29&#160;UTC | SHAKEN 735J | 11&#160;Mar&#160;23&#160;20:29&#160;UTC | true | [view](CERTS/0818fc5f3e540f4ad8e4acd064f828612023ac6b8f94697e7dbc9fe7bef55031/README.md) |
| 04&#160;Mar&#160;23&#160;20:29&#160;UTC | SHAKEN 014E | 11&#160;Mar&#160;23&#160;20:29&#160;UTC | true | [view](CERTS/a900286e67338f1c2fd58307c834f11a3a4ee91b12929915d54627e7ec27557d/README.md) |
| 04&#160;Mar&#160;23&#160;20:30&#160;UTC | SHAKEN 738J | 11&#160;Mar&#160;23&#160;20:30&#160;UTC | true | [view](CERTS/d6111827229d39269a4a3a3ece18ef0e81fec933bdf357d2f6911ecaf3aa48a7/README.md) |
| 04&#160;Mar&#160;23&#160;20:31&#160;UTC | SHAKEN 700H | 11&#160;Mar&#160;23&#160;20:31&#160;UTC | true | [view](CERTS/833a63f85aaf28db8847a53257a450b58d7dc7b17cc2265639ad3142fa1b6d89/README.md) |
| 04&#160;Mar&#160;23&#160;20:32&#160;UTC | SHAKEN 819H | 11&#160;Mar&#160;23&#160;20:32&#160;UTC | true | [view](CERTS/9a9fd75d5910c28f118b68cded4470ee147509f2cc418612a1b35f4f669feb80/README.md) |
| 04&#160;Mar&#160;23&#160;20:33&#160;UTC | SHAKEN 769J | 11&#160;Mar&#160;23&#160;20:33&#160;UTC | true | [view](CERTS/855644e07246dc7228c45c1a5ec8b32608b6e5b491a34f941d7f1811d267dfe4/README.md) |
| 04&#160;Mar&#160;23&#160;20:35&#160;UTC | SHAKEN 691J | 11&#160;Mar&#160;23&#160;20:35&#160;UTC | true | [view](CERTS/22a2d4f8ef3fed4c300780fb2b898ecf4087ae5d280b9b9888d201510cf57119/README.md) |
| 04&#160;Mar&#160;23&#160;20:35&#160;UTC | SHAKEN 849J | 11&#160;Mar&#160;23&#160;20:35&#160;UTC | true | [view](CERTS/b1ed66388f79331e3a201914a189a7ea71007b4866253550e973e9723e8f92dc/README.md) |
| 04&#160;Mar&#160;23&#160;20:35&#160;UTC | SHAKEN 726J | 11&#160;Mar&#160;23&#160;20:35&#160;UTC | true | [view](CERTS/58cb83b79a46927b590c562d027be56fbba7b3f0082d04ab7bcd7c6cb6fb5be1/README.md) |
| 04&#160;Mar&#160;23&#160;20:35&#160;UTC | SHAKEN 469A | 11&#160;Mar&#160;23&#160;20:35&#160;UTC | true | [view](CERTS/3c906601d5cd96122ad30b30cdbadf795586c98d5a97f3c4d603e943afda6a02/README.md) |
| 04&#160;Mar&#160;23&#160;20:35&#160;UTC | SHAKEN 790J | 11&#160;Mar&#160;23&#160;20:35&#160;UTC | true | [view](CERTS/50a3bc377fddea20d68e3c4bd78d0012330fc3f1ff20a7010c59a8132394bc2b/README.md) |
| 04&#160;Mar&#160;23&#160;20:37&#160;UTC | SHAKEN 738J | 11&#160;Mar&#160;23&#160;20:37&#160;UTC | true | [view](CERTS/2a722491104583507f295f1ff16ef4c9e2bde11f3da5c6478ffa8cc9ef18e52c/README.md) |
| 04&#160;Mar&#160;23&#160;20:37&#160;UTC | SHAKEN 459J | 11&#160;Mar&#160;23&#160;20:37&#160;UTC | true | [view](CERTS/5860edca0e4ac64c38512dc8a0951a21aaba7970e3fe32627f8755d575d88043/README.md) |
| 04&#160;Mar&#160;23&#160;20:37&#160;UTC | SHAKEN 672B | 11&#160;Mar&#160;23&#160;20:37&#160;UTC | true | [view](CERTS/a1d41df60c11ef91c0b6a1c80763295300d5f2d7447ba63bee8f22b8ffb3b923/README.md) |
| 04&#160;Mar&#160;23&#160;20:38&#160;UTC | SHAKEN 366G | 11&#160;Mar&#160;23&#160;20:38&#160;UTC | true | [view](CERTS/7ad75d328d1fe8aea6c303397b1acc87e90bd8407cda8c3d75c28523cea8d1d9/README.md) |
| 04&#160;Mar&#160;23&#160;20:38&#160;UTC | SHAKEN 0226 | 11&#160;Mar&#160;23&#160;20:38&#160;UTC | true | [view](CERTS/a63975b48d72b4d6b23b4858e04f7f05e33eafee444a64a348ffa459488068de/README.md) |
| 04&#160;Mar&#160;23&#160;20:39&#160;UTC | SHAKEN 738J | 11&#160;Mar&#160;23&#160;20:39&#160;UTC | true | [view](CERTS/f1bc8947242aed018dab15014931844d5499e8b930c134d2382277aeed590738/README.md) |
| 04&#160;Mar&#160;23&#160;22:11&#160;UTC | SHAKEN 079C | 11&#160;Mar&#160;23&#160;22:11&#160;UTC | true | [view](CERTS/85d3541e8f217992b7614ef0b6957ac52fb68dad67dddae300574968c6b7dfb8/README.md) |
| 05&#160;Mar&#160;23&#160;12:03&#160;UTC | SHAKEN 625J | 12&#160;Mar&#160;23&#160;12:03&#160;UTC | true | [view](CERTS/2f134a7952ba7082491d8b8e6d9a73496114749732cf9e25a12088f4d588ed04/README.md) |
| 05&#160;Mar&#160;23&#160;12:21&#160;UTC | SHAKEN 130B | 12&#160;Mar&#160;23&#160;12:21&#160;UTC | true | [view](CERTS/0795d15887fa22f555a7a77b5a2733261a5bdefddea03c82c9760572e17e1cfa/README.md) |
| 05&#160;Mar&#160;23&#160;15:53&#160;UTC | SHAKEN 793J | 12&#160;Mar&#160;23&#160;15:53&#160;UTC | true | [view](CERTS/179975579ff4262f7dbd4f8b083f551616bf0b0a4eaeb18fdc0632df66c0302c/README.md) |
| 05&#160;Mar&#160;23&#160;16:32&#160;UTC | SHAKEN 035K | 12&#160;Mar&#160;23&#160;16:32&#160;UTC | true | [view](CERTS/11656688f82c01aac0f19f349d58e29bb855a48cca0849c2a66e073d36b7a3ac/README.md) |
| 05&#160;Mar&#160;23&#160;18:13&#160;UTC | SHAKEN 406K | 12&#160;Mar&#160;23&#160;18:13&#160;UTC | true | [view](CERTS/9512e99ae95cea1fb4f606d9c9a62656eac577cbbfb15aa7d93483931e200ff3/README.md) |
| 05&#160;Mar&#160;23&#160;20:26&#160;UTC | SHAKEN 983J | 12&#160;Mar&#160;23&#160;20:26&#160;UTC | true | [view](CERTS/46c9406b4aebf32ff66169eb23adaddf493769955b23462f6e1194c42cc2f369/README.md) |
| 05&#160;Mar&#160;23&#160;21:52&#160;UTC | SHAKEN 606F | 12&#160;Mar&#160;23&#160;21:52&#160;UTC | true | [view](CERTS/e59bc083de884010350b43a08bd3e19c02b44b68275848d78b828ee8b2b6fcbe/README.md) |
| 05&#160;Mar&#160;23&#160;21:55&#160;UTC | SHAKEN 029K | 12&#160;Mar&#160;23&#160;21:55&#160;UTC | true | [view](CERTS/2b52174d2092ba47fb8602b4d68c39ffce8c8300f9dc1cda3a15ec0126a5f534/README.md) |
| 06&#160;Mar&#160;23&#160;01:48&#160;UTC | SHAKEN 278K | 13&#160;Mar&#160;23&#160;01:48&#160;UTC | true | [view](CERTS/9720fec188dada428f20ea44d7357b67a68a5897468f6260f449dc6c6b7a742e/README.md) |
| 06&#160;Mar&#160;23&#160;14:45&#160;UTC | SHAKEN 2550 | 13&#160;Mar&#160;23&#160;14:45&#160;UTC | true | [view](CERTS/81f192f88d17fe6f7512db2649855e6f9adf467d4e77f19966d5b0e707a69974/README.md) |
| 06&#160;Mar&#160;23&#160;20:22&#160;UTC | SHAKEN 177K | 13&#160;Mar&#160;23&#160;20:22&#160;UTC | true | [view](CERTS/915c0dcad972a48ddeb6498a4151eb7526315d19f9123c5f155472edc89100d3/README.md) |
| 06&#160;Mar&#160;23&#160;20:34&#160;UTC | SHAKEN 848J | 13&#160;Mar&#160;23&#160;20:34&#160;UTC | true | [view](CERTS/fcfb2b9ef05936e3cb9817d52c3f6d4fe6a2ad79ceceb7ad95b825a48988bfa7/README.md) |
| 06&#160;Mar&#160;23&#160;20:40&#160;UTC | SHAKEN 738J | 15&#160;Mar&#160;23&#160;20:40&#160;UTC | true | [view](CERTS/5cc7027cfdafc4026d2572c1feccb84e7bd6d708fb0326198182812abf3c7570/README.md) |
| 07&#160;Mar&#160;23&#160;09:44&#160;UTC | SHAKEN 0172 | 14&#160;Mar&#160;23&#160;09:44&#160;UTC | true | [view](CERTS/dd4e31828512a793c67a888aba825299251e8edc3114e3696aac22487ed0d331/README.md) |
| 07&#160;Mar&#160;23&#160;14:52&#160;UTC | SHAKEN 747J | 14&#160;Mar&#160;23&#160;14:52&#160;UTC | true | [view](CERTS/48fac99a33f3bfe0f3147dfc3a92b8f904196fefa43575865c64ff9dac4b4e0a/README.md) |
| 07&#160;Mar&#160;23&#160;17:35&#160;UTC | SHAKEN 107K | 14&#160;Mar&#160;23&#160;17:35&#160;UTC | true | [view](CERTS/9b54ec07a7c6b3a589e1a8e3d543b990846614508692cf15a600ec829078df32/README.md) |
| 07&#160;Mar&#160;23&#160;18:49&#160;UTC | SHAKEN 9714 | 14&#160;Mar&#160;23&#160;18:49&#160;UTC | true | [view](CERTS/432c23856be71cd67e190a44a94f058f2a9d4cee4a04a3136814cb43a0405253/README.md) |
| 07&#160;Mar&#160;23&#160;20:06&#160;UTC | SHAKEN 297K | 14&#160;Mar&#160;23&#160;20:06&#160;UTC | true | [view](CERTS/c3a58d0b15949e7d2b62ae1148692895ef872332f9a7b65985bd7644f14c2678/README.md) |
| 07&#160;Mar&#160;23&#160;20:28&#160;UTC | SHAKEN 674J | 14&#160;Mar&#160;23&#160;20:28&#160;UTC | true | [view](CERTS/453384e0bce7e9e1fd0639f2ff053b76c4e99f062b430a71cd389b390f695863/README.md) |
| 07&#160;Mar&#160;23&#160;20:30&#160;UTC | SHAKEN 738J | 14&#160;Mar&#160;23&#160;20:30&#160;UTC | true | [view](CERTS/4c485f580f561fa66b86790ff31d2802b4649a1014bc9cb854035f8c2698bde3/README.md) |
| 07&#160;Mar&#160;23&#160;20:31&#160;UTC | SHAKEN 700H | 14&#160;Mar&#160;23&#160;20:31&#160;UTC | true | [view](CERTS/1d7a34332cfcec3c3ba9d5cf698d50bb4ffdf10a817277cb6c540d606d1a858c/README.md) |
| 07&#160;Mar&#160;23&#160;20:31&#160;UTC | SHAKEN 733J | 14&#160;Mar&#160;23&#160;20:31&#160;UTC | true | [view](CERTS/c974f7c3d22901b316a1f6489fe949b3379d9afa3bc6b264b7b58ec2f3bdf272/README.md) |
| 07&#160;Mar&#160;23&#160;20:32&#160;UTC | SHAKEN 819H | 14&#160;Mar&#160;23&#160;20:32&#160;UTC | true | [view](CERTS/754cda829392a96cce17f0d2a9716bd770d7ecfad6cf06729a5abf16a4d17d75/README.md) |
| 07&#160;Mar&#160;23&#160;20:33&#160;UTC | SHAKEN 0753 | 14&#160;Mar&#160;23&#160;20:33&#160;UTC | true | [view](CERTS/a8f43eeda1e492584a00fbb4778bb54c06b5a2b7ab20fb5809df9af7bf09b715/README.md) |
| 07&#160;Mar&#160;23&#160;20:33&#160;UTC | SHAKEN 053G | 14&#160;Mar&#160;23&#160;20:33&#160;UTC | true | [view](CERTS/d5495228f11260662465ec2c78adbd9e3db86bc497fb8d8532e2941e3e965c07/README.md) |
| 07&#160;Mar&#160;23&#160;20:33&#160;UTC | SHAKEN 769J | 14&#160;Mar&#160;23&#160;20:33&#160;UTC | true | [view](CERTS/02bc7ea40d2c2a21fa62efa3817e884e97df7e6a4d0f7195885cd286a2b4bcc6/README.md) |
| 07&#160;Mar&#160;23&#160;20:33&#160;UTC | SHAKEN 589J | 14&#160;Mar&#160;23&#160;20:33&#160;UTC | true | [view](CERTS/2186bfd6843420891b11ac8240768018552101ad185dbd3449435cff87b275bc/README.md) |
| 07&#160;Mar&#160;23&#160;20:36&#160;UTC | SHAKEN 691J | 14&#160;Mar&#160;23&#160;20:36&#160;UTC | true | [view](CERTS/b2e0a85e2b90ed288d68dd423a0610eca162f1b57209c0ea5e347fc409367611/README.md) |
| 07&#160;Mar&#160;23&#160;20:36&#160;UTC | SHAKEN 849J | 14&#160;Mar&#160;23&#160;20:36&#160;UTC | true | [view](CERTS/ffc612f62f9eee23d64d4cf693256996ea713cc58f2e09f9202dd033e1f8f881/README.md) |
| 07&#160;Mar&#160;23&#160;20:36&#160;UTC | SHAKEN 469A | 14&#160;Mar&#160;23&#160;20:36&#160;UTC | true | [view](CERTS/9b76ca966f7c6a465d08437dd9f504f392496f4ab51984528244376d6dfe5620/README.md) |
| 07&#160;Mar&#160;23&#160;20:36&#160;UTC | SHAKEN 790J | 14&#160;Mar&#160;23&#160;20:36&#160;UTC | true | [view](CERTS/7fb636489167e9d81aa05fa859699d4cfa486d917102640ac5b17bad0ded1d26/README.md) |
| 07&#160;Mar&#160;23&#160;20:37&#160;UTC | SHAKEN 625J | 14&#160;Mar&#160;23&#160;20:37&#160;UTC | true | [view](CERTS/0087f3d1ab7d502475afd5c86a3eee30265bc5eabc026ff05dfc9ef4e68e56fc/README.md) |
| 07&#160;Mar&#160;23&#160;20:37&#160;UTC | SHAKEN 738J | 14&#160;Mar&#160;23&#160;20:37&#160;UTC | true | [view](CERTS/5566494e7be757f610788a0f84b8b193a7aafd549b825e93b5e7ba01507aff92/README.md) |
| 07&#160;Mar&#160;23&#160;20:37&#160;UTC | SHAKEN 459J | 14&#160;Mar&#160;23&#160;20:37&#160;UTC | true | [view](CERTS/b7c6a9638d303cdb5a769c9d3333c1c4bdff714b62fd6180c2d337976f756e3e/README.md) |
| 07&#160;Mar&#160;23&#160;20:38&#160;UTC | SHAKEN 366G | 14&#160;Mar&#160;23&#160;20:38&#160;UTC | true | [view](CERTS/1cecbb37e7f1714373f26c10b498fdc91e88e74e29f2f94d4880e95b740cb2da/README.md) |
| 07&#160;Mar&#160;23&#160;20:38&#160;UTC | SHAKEN 0226 | 14&#160;Mar&#160;23&#160;20:38&#160;UTC | true | [view](CERTS/a10c7e1807fed3d7a6c85017948d63df9e1ae7fe2af9ca2a4d62e668d66aa0c1/README.md) |
| 07&#160;Mar&#160;23&#160;20:39&#160;UTC | SHAKEN 738J | 14&#160;Mar&#160;23&#160;20:39&#160;UTC | true | [view](CERTS/fb6b72e89978c7fa4f1d91a1984fa359bc8bff0b5664eb0c9bc87e90a4f25e4b/README.md) |
| 08&#160;Mar&#160;23&#160;16:51&#160;UTC | SHAKEN 967J | 15&#160;Mar&#160;23&#160;16:51&#160;UTC | true | [view](CERTS/4e4247ce7f5d140626cac150381a62d4b4e4238f353e7db77fd292540d8156bd/README.md) |
| 08&#160;Mar&#160;23&#160;18:13&#160;UTC | SHAKEN 406K | 15&#160;Mar&#160;23&#160;18:13&#160;UTC | true | [view](CERTS/637f7108dcf77001240966ff0d4e25040091d7ed356acab5969d516c3de5a530/README.md) |
| 08&#160;Mar&#160;23&#160;20:26&#160;UTC | SHAKEN 983J | 15&#160;Mar&#160;23&#160;20:26&#160;UTC | true | [view](CERTS/eaf2bee351873e2095c3a49e4a09cca933803a2f075018f1e5cb943754aba6ce/README.md) |
| 09&#160;Mar&#160;23&#160;01:49&#160;UTC | SHAKEN 278K | 16&#160;Mar&#160;23&#160;01:49&#160;UTC | true | [view](CERTS/4230a79f7a482999d3fd30c9b4d1f2397a7828183b9430e885fdc5dd08b7adce/README.md) |
| 09&#160;Mar&#160;23&#160;10:01&#160;UTC | SHAKEN 186K | 10&#160;Mar&#160;23&#160;10:01&#160;UTC | true | [view](CERTS/356fe7c23c978912c165061aa50af905993760808b7aa92ff64a12d89b6f0be7/README.md) |
| 09&#160;Mar&#160;23&#160;14:45&#160;UTC | SHAKEN 2550 | 16&#160;Mar&#160;23&#160;14:45&#160;UTC | true | [view](CERTS/6705dead756d995e8a7a9a206d5e0714e3789a8c79810f83ac470984fd5c0e73/README.md) |
| 09&#160;Mar&#160;23&#160;18:46&#160;UTC | SHAKEN 193E | 08&#160;May&#160;23&#160;18:46&#160;UTC | true | [view](CERTS/34187f020bbbd3596eea83f5f2243b84920396d2c46eed5b8078a0b696c383e5/README.md) |

#### CA Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 20&#160;Aug&#160;21&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Issuing CA3 | 19&#160;Aug&#160;31&#160;23:59&#160;UTC | true | [view](CERTS/51b51431ffd8055cd7e30c213540f4868258058d4c4c9f23f1ed8bea75101e37/README.md) |
| 20&#160;Aug&#160;21&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Issuing CA1 | 19&#160;Aug&#160;31&#160;23:59&#160;UTC | true | [view](CERTS/037a7c6621814bd5f01fa2080a77a245458694d9d566e0050e068fc853317857/README.md) |
| 20&#160;Aug&#160;21&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Root CA1 | 19&#160;Aug&#160;41&#160;23:59&#160;UTC | true | [view](CERTS/59ab943f4e8c3c17755cc8e4abb5ac65736ce74ec587c62cb744ab0babbbe2fe/README.md) |
| 24&#160;Oct&#160;22&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Issuing CA4 | 23&#160;Oct&#160;32&#160;23:59&#160;UTC | false | [view](CERTS/80dd12d935f46c256c7c1289a0834b21ec8f2f8c35f2864928d1e75f3a280665/README.md) |
| 24&#160;Oct&#160;22&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Root CA2 | 23&#160;Oct&#160;42&#160;23:59&#160;UTC | false | [view](CERTS/a26e04fc786ab70b8085236b2c53f8cfbf5d0c6a5c2c9c3e9f91669fbb8ea4d5/README.md) |


Generated: 10 Mar 23 02:25 UTC