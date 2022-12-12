# STIR/SHAKEN CA Ecosystem Compliance

## TransNexus

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 740 certificates were included in the corpus being tested
- 1 certificates in the corpus were skipped because they are duplicates
- 640 certificates in the corpus were skipped because they are expired
- 1 certificates in the corpus were skipped because they are not currently trusted
- 98 certificates being tested against the remaining rules
- 1.41 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 18.37% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 1.02% of certificates are too old to be assessed against currently enforced expectations
- 55 days is the average remaining validity for the certificates in the corpus
- 55 days is the average initial validity for the certificates in the corpus
- 84 certificates expire in the next 30 days
- 1.53 average number of unexpired certificates per OCN observed
- 64 unique OCNs observed in unexpired and valid certificate corpus

| Instances | Test | Source |
|-----------|------|--------|
| 3 | [e_atis_certificate_policies](ISSUES/e_atis_certificate_policies/README.md) | ATIS1000080 |
| 16 | [e_incorrect_ku_encoding](ISSUES/e_incorrect_ku_encoding/README.md) | RFC5280 |
| 3 | [e_us_cp_ambiguous_identifier](ISSUES/e_us_cp_ambiguous_identifier/README.md) | US_SHAKEN_CP |
| 98 | [e_us_cp_subject_sn](ISSUES/e_us_cp_subject_sn/README.md) | US_SHAKEN_CP |
| 18 | [w_shaken_subject_rdn_unknown](ISSUES/w_shaken_subject_rdn_unknown/README.md) | SHAKEN_PKI_BEST_PRACTICES |

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
- 5470 days is the average remaining validity for the certificates in the corpus
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
| 13&#160;Jan&#160;22&#160;14:17&#160;UTC | MobileSphere SHAKEN 873J | 13&#160;Jan&#160;23&#160;14:17&#160;UTC | true | [view](CERTS/8f74ddd5e6964042b119d5b3f17d24694ec78c8688e0ffc96d13fdb0e8df26d4/README.md) |
| 30&#160;Mar&#160;22&#160;12:54&#160;UTC | Fusion Connect SHAKEN 2720 | 30&#160;Mar&#160;23&#160;12:54&#160;UTC | true | [view](CERTS/03a010b294807e90cc1309d7466fd2ddc591158a69be6950ff4d0eab32de0799/README.md) |
| 27&#160;Apr&#160;22&#160;18:22&#160;UTC | CCI SHAKEN 663J | 27&#160;Apr&#160;23&#160;18:22&#160;UTC | true | [view](CERTS/9cfd0b9974acdc9b5fce3888cf2f6cf99173e7955adeebd37a1e60dfc444a0b3/README.md) |
| 10&#160;Jun&#160;22&#160;14:00&#160;UTC | SHAKEN 597J | 10&#160;Jun&#160;23&#160;14:00&#160;UTC | true | [view](CERTS/4baa838eb96ce358aa1565d0add8b75072258efd5d21a5bbb0eec878026a67be/README.md) |
| 20&#160;Jun&#160;22&#160;20:25&#160;UTC | SHAKEN 857J | 20&#160;Jun&#160;23&#160;20:25&#160;UTC | true | [view](CERTS/e1e0a4851fb7b51d617959c00fb6a17e7b23d34e1555ba70ab01c9182a2f85f8/README.md) |
| 22&#160;Jun&#160;22&#160;18:46&#160;UTC | SHAKEN 755J | 22&#160;Jun&#160;23&#160;18:46&#160;UTC | true | [view](CERTS/a8cac316fddf25b48b9350a325791be8b87e6484a5959ab605db1ae5e3e99a84/README.md) |
| 29&#160;Jun&#160;22&#160;20:24&#160;UTC | SHAKEN 736J | 29&#160;Jun&#160;23&#160;20:24&#160;UTC | true | [view](CERTS/71b25a480077816998e48e6a76f272e70aa91ecb19b0e773cd8777823cdead32/README.md) |
| 25&#160;Jul&#160;22&#160;18:36&#160;UTC | SHAKEN 578J | 25&#160;Jul&#160;23&#160;18:36&#160;UTC | true | [view](CERTS/f4af5a28131d680ef2682ae86181b64f0c919864d7dbaef00c0598d314b537a1/README.md) |
| 10&#160;Aug&#160;22&#160;18:11&#160;UTC | SHAKEN 073H | 10&#160;Aug&#160;23&#160;18:11&#160;UTC | true | [view](CERTS/87dabb91a689a6ac540b5ffbf5eb89326fada83fb822d2bb9ec31b68251b579e/README.md) |
| 23&#160;Aug&#160;22&#160;04:14&#160;UTC | SHAKEN 866J | 19&#160;Feb&#160;23&#160;04:14&#160;UTC | true | [view](CERTS/4c4fdb320c51c8582d595c2d307cd770d409daf533ff573bcc73afaed83f6b7d/README.md) |
| 01&#160;Sep&#160;22&#160;03:25&#160;UTC | SHAKEN 6628 | 03&#160;Jan&#160;23&#160;03:25&#160;UTC | true | [view](CERTS/37b9a67bb8e3272048330f269b5dbc285ac0278df4670ab7e291b92b6548fa2d/README.md) |
| 13&#160;Sep&#160;22&#160;17:24&#160;UTC | SHAKEN 706J | 13&#160;Sep&#160;23&#160;17:24&#160;UTC | true | [view](CERTS/1eea5e17ae8b36cb5b6d657178273af745cea015e63586e6d44d1ba2ee8f722a/README.md) |
| 26&#160;Sep&#160;22&#160;18:40&#160;UTC | SHAKEN 505J | 25&#160;Dec&#160;22&#160;18:40&#160;UTC | true | [view](CERTS/b186c959554e652b69824a2e45ac08ec135105b0e01c9d2eadeb3cf46c130670/README.md) |
| 26&#160;Sep&#160;22&#160;18:42&#160;UTC | SHAKEN 505J | 25&#160;Dec&#160;22&#160;18:42&#160;UTC | true | [view](CERTS/1c0003ac10eeaa04229e507c15f71ee018c3902c3e8c20fbe42e533b8682ba8c/README.md) |
| 12&#160;Oct&#160;22&#160;19:39&#160;UTC | SHAKEN 815G | 10&#160;Jan&#160;23&#160;19:39&#160;UTC | true | [view](CERTS/a79f50d4fcd3e680518379f12a86a4b9a21c5e27dc5422a8439525edf350c4a7/README.md) |
| 20&#160;Oct&#160;22&#160;15:48&#160;UTC | SHAKEN 622J | 18&#160;Jan&#160;23&#160;15:48&#160;UTC | true | [view](CERTS/2e34c765f110b197f2ec1052b1960bc25ce320dfce7d7035cc532cd1eb6dbd57/README.md) |
| 31&#160;Oct&#160;22&#160;12:32&#160;UTC | SHAKEN 193E | 30&#160;Dec&#160;22&#160;12:32&#160;UTC | true | [view](CERTS/0de5e0e6787b9aa19c02c014821dd7a2e62e438f4eb9fbbb761eb86df8c69ff7/README.md) |
| 01&#160;Nov&#160;22&#160;11:14&#160;UTC | SHAKEN 807J | 31&#160;Dec&#160;22&#160;11:14&#160;UTC | true | [view](CERTS/0d97796c10ef21ba02a0f7954d82ace5e41d7ea66e4ea9eb16801c3a1a1d9f03/README.md) |
| 18&#160;Nov&#160;22&#160;15:42&#160;UTC | SHAKEN 722J | 18&#160;Dec&#160;22&#160;15:42&#160;UTC | true | [view](CERTS/4111bb56ba547a2f493d1e9c2baefa33ce9bd117904133477d063d6fa0d6b9b2/README.md) |
| 18&#160;Nov&#160;22&#160;18:04&#160;UTC | SHAKEN 952J | 18&#160;Dec&#160;22&#160;18:04&#160;UTC | true | [view](CERTS/b05bb158c054b6e579fe20750c83ba5524896545d629b96b51d4c8d9ab38dfd3/README.md) |
| 23&#160;Nov&#160;22&#160;17:31&#160;UTC | SHAKEN 4036 | 10&#160;May&#160;23&#160;17:31&#160;UTC | true | [view](CERTS/e03ff4259e6b456f1bb19d011f005932e83b0a35b520730c97cd77ad8868071c/README.md) |
| 30&#160;Nov&#160;22&#160;12:32&#160;UTC | SHAKEN 193E | 29&#160;Jan&#160;23&#160;12:32&#160;UTC | true | [view](CERTS/b1085c4ed6a32a87dc773d65e4e7aec591f7625ad80f63ee690d56b74b7a1430/README.md) |
| 01&#160;Dec&#160;22&#160;09:33&#160;UTC | SHAKEN 841J | 15&#160;Dec&#160;22&#160;09:33&#160;UTC | true | [view](CERTS/f98bb6bf4f32b42eda9a2113135539ab9422a7cb2d958e3e8c3fce20e1cbbd6d/README.md) |
| 03&#160;Dec&#160;22&#160;04:11&#160;UTC | SHAKEN 345J | 02&#160;Jan&#160;23&#160;04:11&#160;UTC | true | [view](CERTS/8991e6ec3e951d1b91f0e0f63b5bc9bdc929b5c76a73fea6c5af7a6401d0aa6e/README.md) |
| 03&#160;Dec&#160;22&#160;05:00&#160;UTC | SHAKEN 345J | 02&#160;Jan&#160;23&#160;05:00&#160;UTC | true | [view](CERTS/df273ff8f12fbf2fda88e527381fba4ac91927b0be97ec38a2d67993a5e0c074/README.md) |
| 05&#160;Dec&#160;22&#160;15:23&#160;UTC | SHAKEN 8526 | 04&#160;Jan&#160;23&#160;15:23&#160;UTC | true | [view](CERTS/29453875467b9cbda7eccf1ee68baaa883140c03b133004d5f21864cec03a029/README.md) |
| 05&#160;Dec&#160;22&#160;20:27&#160;UTC | SHAKEN 738J | 14&#160;Dec&#160;22&#160;20:27&#160;UTC | true | [view](CERTS/652a1a6e9b711bc5eba31cca81b066a5a8de44758bcc4e4b860efd27d1b42b15/README.md) |
| 06&#160;Dec&#160;22&#160;01:43&#160;UTC | SHAKEN 278K | 13&#160;Dec&#160;22&#160;01:43&#160;UTC | true | [view](CERTS/c5ff7f48274a7a8f03851226d7dd566ae43e54c22f5884794e36d894b8f6cc59/README.md) |
| 06&#160;Dec&#160;22&#160;14:37&#160;UTC | SHAKEN 2550 | 13&#160;Dec&#160;22&#160;14:37&#160;UTC | true | [view](CERTS/a268d78ca7a2ba41cf5deb6dac12c5621f6f73c3f68416e4366be9fe5026e10a/README.md) |
| 06&#160;Dec&#160;22&#160;16:20&#160;UTC | SHAKEN 753J | 13&#160;Dec&#160;22&#160;16:20&#160;UTC | true | [view](CERTS/bb3357f22e091ef9b762df56b657d69111737580852e4d1f9d803d7cd6c97863/README.md) |
| 06&#160;Dec&#160;22&#160;18:48&#160;UTC | SHAKEN 140K | 13&#160;Dec&#160;22&#160;18:48&#160;UTC | true | [view](CERTS/0fd2fa51646825281f95779e5032774a2db27d3205fc63de78f1a9563af94ee2/README.md) |
| 06&#160;Dec&#160;22&#160;20:09&#160;UTC | SHAKEN 864J | 13&#160;Dec&#160;22&#160;20:09&#160;UTC | true | [view](CERTS/09baaed6752794dbeb4f94bf75d92f6b9ba1cb8ba1e2a9def93c0a8c586a6a61/README.md) |
| 06&#160;Dec&#160;22&#160;20:15&#160;UTC | SHAKEN 177K | 13&#160;Dec&#160;22&#160;20:15&#160;UTC | true | [view](CERTS/9dc1905a2c780092ec257c60ad6bd9069edf9d994791374eb9486d0378790492/README.md) |
| 07&#160;Dec&#160;22&#160;05:23&#160;UTC | SHAKEN 625J | 14&#160;Dec&#160;22&#160;05:23&#160;UTC | true | [view](CERTS/e0700384ffb8d036dbd8a26996a30d64440be141f3207c566c8c74bc1b16c414/README.md) |
| 07&#160;Dec&#160;22&#160;14:44&#160;UTC | SHAKEN 012K | 14&#160;Dec&#160;22&#160;14:44&#160;UTC | true | [view](CERTS/a630e99a911618f4c67255e8cad1008c25dc5216b6b34e98ce200938f4c527eb/README.md) |
| 07&#160;Dec&#160;22&#160;14:45&#160;UTC | SHAKEN 747J | 14&#160;Dec&#160;22&#160;14:45&#160;UTC | true | [view](CERTS/0593bddc3e8ec77c3b28c0855ea7a6d7dbb1dfbb55b1c282681d30fa209d7012/README.md) |
| 07&#160;Dec&#160;22&#160;17:27&#160;UTC | SHAKEN 107K | 14&#160;Dec&#160;22&#160;17:27&#160;UTC | true | [view](CERTS/e5f402562fff0606692fc4a87938074a65fc53c5dcf21f267a05f1e1fa650a6f/README.md) |
| 07&#160;Dec&#160;22&#160;18:15&#160;UTC | SHAKEN 060K | 14&#160;Dec&#160;22&#160;18:15&#160;UTC | true | [view](CERTS/4bdb139e8c228f9307a5818d54a2897037bed43ad9ccf59b29a6670ecd0f90e8/README.md) |
| 07&#160;Dec&#160;22&#160;18:40&#160;UTC | SHAKEN 089K | 14&#160;Dec&#160;22&#160;18:40&#160;UTC | true | [view](CERTS/d2ce041d8c052a7c022fb617b7ca22686a9b6582d3cb21f899bbf227f9d2dfec/README.md) |
| 07&#160;Dec&#160;22&#160;18:40&#160;UTC | SHAKEN 056K | 14&#160;Dec&#160;22&#160;18:40&#160;UTC | true | [view](CERTS/e070876dac456673a17f10cc18fead1471f32983bf51ddc09e92733e804e9bf5/README.md) |
| 07&#160;Dec&#160;22&#160;19:59&#160;UTC | SHAKEN 297K | 14&#160;Dec&#160;22&#160;19:59&#160;UTC | true | [view](CERTS/e938f680ba0d223710400ddc99b9c2ed863adeb06f40b840de1371dcb1e87e50/README.md) |
| 07&#160;Dec&#160;22&#160;20:05&#160;UTC | SHAKEN 366G | 14&#160;Dec&#160;22&#160;20:05&#160;UTC | true | [view](CERTS/b76a6b44cae150a200c10a2cd2e946227da0d215b212901a1fbe58ba296c763b/README.md) |
| 07&#160;Dec&#160;22&#160;20:20&#160;UTC | SHAKEN 674J | 14&#160;Dec&#160;22&#160;20:20&#160;UTC | true | [view](CERTS/94b56856bbf0662864c76bd0cfb0975771904bb69495242902469a3b29108869/README.md) |
| 07&#160;Dec&#160;22&#160;20:22&#160;UTC | SHAKEN 065J | 14&#160;Dec&#160;22&#160;20:21&#160;UTC | true | [view](CERTS/4e0b03906f292eb27c26a1c97f210f09164efb22fe334a1768605406b3cba206/README.md) |
| 07&#160;Dec&#160;22&#160;20:22&#160;UTC | SHAKEN 738J | 14&#160;Dec&#160;22&#160;20:22&#160;UTC | true | [view](CERTS/d2482e7430ca6b1a8245ff4f1d3b5707505ccecae3d8515e4521dd15b6ed243e/README.md) |
| 07&#160;Dec&#160;22&#160;20:24&#160;UTC | SHAKEN 819H | 14&#160;Dec&#160;22&#160;20:24&#160;UTC | true | [view](CERTS/0319fb73caf1ef5fe590bc8d93fdac8ebbb915259e47b69d8698648bdb1e5c8d/README.md) |
| 07&#160;Dec&#160;22&#160;20:25&#160;UTC | SHAKEN 769J | 14&#160;Dec&#160;22&#160;20:25&#160;UTC | true | [view](CERTS/3973f6361ff713642794a7bf16e7a6c653d011bd25a1f6a851589484786b41d2/README.md) |
| 07&#160;Dec&#160;22&#160;20:26&#160;UTC | SHAKEN 849J | 14&#160;Dec&#160;22&#160;20:26&#160;UTC | true | [view](CERTS/f6d0c7591da9aaa9c4d1a978bd029eaf5fe1f342c645726d02e967cf90b358b7/README.md) |
| 07&#160;Dec&#160;22&#160;20:26&#160;UTC | SHAKEN 469A | 14&#160;Dec&#160;22&#160;20:26&#160;UTC | true | [view](CERTS/afaca3fe2ae42ad7a27b5a7ff8acce3484547dc8feca7d1827e35c4062b6fdd0/README.md) |
| 07&#160;Dec&#160;22&#160;20:26&#160;UTC | SHAKEN 726J | 14&#160;Dec&#160;22&#160;20:26&#160;UTC | true | [view](CERTS/446b75de369bade8b448a7f66012bc6853b9479c3de14985a4aeb7c3ce144f31/README.md) |
| 07&#160;Dec&#160;22&#160;20:26&#160;UTC | SHAKEN 495J | 14&#160;Dec&#160;22&#160;20:26&#160;UTC | true | [view](CERTS/169bf2dea785f1172bfd7c02433daeed42549d656e422e7dd173e37dc7d9a959/README.md) |
| 07&#160;Dec&#160;22&#160;20:26&#160;UTC | SHAKEN 790J | 14&#160;Dec&#160;22&#160;20:26&#160;UTC | true | [view](CERTS/8b77faf4ac62562102e4f1e64698f3d3ed5a825350af59730c794f46bc38bb5e/README.md) |
| 07&#160;Dec&#160;22&#160;20:27&#160;UTC | SHAKEN 625J | 14&#160;Dec&#160;22&#160;20:27&#160;UTC | true | [view](CERTS/d8d1fb0da1425a78a8b098db16d4ce876b3cebae77a0294f13a30a66e80ec057/README.md) |
| 07&#160;Dec&#160;22&#160;20:27&#160;UTC | SHAKEN 738J | 14&#160;Dec&#160;22&#160;20:27&#160;UTC | true | [view](CERTS/a804f6060d734c3db2067c68cfa665ec6d38befd022247c0ade6dd313cff9d01/README.md) |
| 07&#160;Dec&#160;22&#160;20:27&#160;UTC | SHAKEN 459J | 14&#160;Dec&#160;22&#160;20:27&#160;UTC | true | [view](CERTS/c77275b3b817df60fba1999766c03d058bb5a13ce350197e2905f17f4785a02c/README.md) |
| 07&#160;Dec&#160;22&#160;20:27&#160;UTC | SHAKEN 672B | 14&#160;Dec&#160;22&#160;20:27&#160;UTC | true | [view](CERTS/8bc78a69dbe6ac89c5770f103f57ab619422c3541e203fcd1f9366ec8a199b6e/README.md) |
| 07&#160;Dec&#160;22&#160;20:29&#160;UTC | SHAKEN 366G | 14&#160;Dec&#160;22&#160;20:28&#160;UTC | true | [view](CERTS/a5581304da9702a8fe6b402fd9e33deb0fdd0bef9b2d85f73b44c892faa78dbf/README.md) |
| 07&#160;Dec&#160;22&#160;20:29&#160;UTC | SHAKEN 0226 | 14&#160;Dec&#160;22&#160;20:29&#160;UTC | true | [view](CERTS/97047cfe4dbc5971a0843ef46e61d4e36f4db32cef88687693a4eaa3fc97eb94/README.md) |
| 07&#160;Dec&#160;22&#160;20:29&#160;UTC | SHAKEN 738J | 14&#160;Dec&#160;22&#160;20:29&#160;UTC | true | [view](CERTS/05f9f2554d6279e53203a5bd4b6e3e44d1480456b5cb3f74b3ade3d90ef12b8d/README.md) |
| 07&#160;Dec&#160;22&#160;21:02&#160;UTC | SHAKEN 518J | 14&#160;Dec&#160;22&#160;21:02&#160;UTC | true | [view](CERTS/54d63611e3db1366826aac52ac2d1db923c15cd4cfd19900be037eaaad44d181/README.md) |
| 08&#160;Dec&#160;22&#160;16:36&#160;UTC | SHAKEN 324K | 15&#160;Dec&#160;22&#160;16:36&#160;UTC | true | [view](CERTS/984efcff2ae364a7dfc127187eee78babe15813860aad42a66554bf8a6cc9509/README.md) |
| 08&#160;Dec&#160;22&#160;17:06&#160;UTC | SHAKEN 226K | 15&#160;Dec&#160;22&#160;17:06&#160;UTC | true | [view](CERTS/9ed35140a2480d64dfa1c6e878ebd737e6380469130d5b604d2f0648d5da2c7f/README.md) |
| 08&#160;Dec&#160;22&#160;20:18&#160;UTC | SHAKEN 983J | 15&#160;Dec&#160;22&#160;20:18&#160;UTC | true | [view](CERTS/746d32ab6fd8430232327db5fbdb00d86c5edd95632165a8d9adc49a6701c285/README.md) |
| 09&#160;Dec&#160;22&#160;04:01&#160;UTC | SHAKEN 0172 | 16&#160;Dec&#160;22&#160;04:01&#160;UTC | true | [view](CERTS/2ad38e67a6bc36eee459e79c791ddb021db300ed96cd1179fcf31c80fe29b9db/README.md) |
| 09&#160;Dec&#160;22&#160;14:37&#160;UTC | SHAKEN 2550 | 16&#160;Dec&#160;22&#160;14:37&#160;UTC | true | [view](CERTS/73497da42b213da29401fb5db6be05213e17dc02d2b6d4034d0589110eaab7f2/README.md) |
| 09&#160;Dec&#160;22&#160;19:11&#160;UTC | SHAKEN 841J | 23&#160;Dec&#160;22&#160;19:11&#160;UTC | true | [view](CERTS/af11f73f26eda12928f6b5c61c8bcecc0467dd3924267f7fb0fb8d6bd9820cf9/README.md) |
| 09&#160;Dec&#160;22&#160;20:09&#160;UTC | SHAKEN 864J | 16&#160;Dec&#160;22&#160;20:09&#160;UTC | true | [view](CERTS/73e3ddad1030641438e6bd903d72f3d92b27965bec98bef5b0a6bcc40be5dd1b/README.md) |
| 09&#160;Dec&#160;22&#160;20:15&#160;UTC | SHAKEN 177K | 16&#160;Dec&#160;22&#160;20:15&#160;UTC | true | [view](CERTS/1a1e9444a369d9ab1beffc010262de7ee1fa401f6f2ded2e77eb2d1d24150075/README.md) |
| 09&#160;Dec&#160;22&#160;20:27&#160;UTC | SHAKEN 848J | 16&#160;Dec&#160;22&#160;20:27&#160;UTC | true | [view](CERTS/1e14c7d4a0c5ba20431f150e71a2dd8fcdaa65ee1eb5736ac9a08ac94d5861e2/README.md) |
| 10&#160;Dec&#160;22&#160;01:03&#160;UTC | SHAKEN 551G | 17&#160;Dec&#160;22&#160;01:03&#160;UTC | true | [view](CERTS/2d46f2c08cdfe97a591b413d038ca4049809c390b642813ba2561fbe0d75d1aa/README.md) |
| 10&#160;Dec&#160;22&#160;14:44&#160;UTC | SHAKEN 012K | 17&#160;Dec&#160;22&#160;14:44&#160;UTC | true | [view](CERTS/338c5a71fdd5e25775ab10aa90e8dd22d5691d2fe0c003169d7c5fce9e480258/README.md) |
| 10&#160;Dec&#160;22&#160;16:52&#160;UTC | SHAKEN 159H | 17&#160;Dec&#160;22&#160;16:52&#160;UTC | true | [view](CERTS/8039bc905bee1c9c2a25684f45bb8684331786c9df430e4e8af1576fb669038b/README.md) |
| 10&#160;Dec&#160;22&#160;17:27&#160;UTC | SHAKEN 107K | 17&#160;Dec&#160;22&#160;17:27&#160;UTC | true | [view](CERTS/126451a6a680727d9468ca324c603bff3aa2703f87c585dff68277f4025b5b66/README.md) |
| 10&#160;Dec&#160;22&#160;18:40&#160;UTC | SHAKEN 089K | 17&#160;Dec&#160;22&#160;18:40&#160;UTC | true | [view](CERTS/354e72a4eb4a9cc96af2f1664c75a4a9cf7bac770a30a42a4d600cb630b844e3/README.md) |
| 10&#160;Dec&#160;22&#160;19:59&#160;UTC | SHAKEN 297K | 17&#160;Dec&#160;22&#160;19:59&#160;UTC | true | [view](CERTS/74e4079abfce7562746b817a8db78a4622b49d5ba09cbc4356b25fe6c916c81f/README.md) |
| 10&#160;Dec&#160;22&#160;20:05&#160;UTC | SHAKEN 366G | 17&#160;Dec&#160;22&#160;20:05&#160;UTC | true | [view](CERTS/52ac4bca14ab304a804678a214df11c7ddebad029ed1fe1fc6dc23f019dc2922/README.md) |
| 10&#160;Dec&#160;22&#160;20:20&#160;UTC | SHAKEN 674J | 17&#160;Dec&#160;22&#160;20:20&#160;UTC | true | [view](CERTS/5c5ff79c800c3b25ba803b8ecb8b7a0cb764400fd298c414c31b438852b26a95/README.md) |
| 10&#160;Dec&#160;22&#160;20:22&#160;UTC | SHAKEN 594J | 17&#160;Dec&#160;22&#160;20:21&#160;UTC | true | [view](CERTS/1184f51d0289faabb98ec8843a83aa26658fa76a2a9b8317e264f165453b2e17/README.md) |
| 10&#160;Dec&#160;22&#160;20:22&#160;UTC | SHAKEN 738J | 17&#160;Dec&#160;22&#160;20:22&#160;UTC | true | [view](CERTS/24c1aa497de706a085015aef6ad44b9aa5da9d655fa83296d73e0bb8187d39bf/README.md) |
| 10&#160;Dec&#160;22&#160;20:26&#160;UTC | SHAKEN 854D | 17&#160;Dec&#160;22&#160;20:26&#160;UTC | true | [view](CERTS/84d39a066d7b8cbb4c52bfa2cd8845ed95a83805cb8b8f9ecd2d7df1bba9aea4/README.md) |
| 10&#160;Dec&#160;22&#160;20:27&#160;UTC | SHAKEN 469A | 17&#160;Dec&#160;22&#160;20:27&#160;UTC | true | [view](CERTS/8f671a9cd4f526e0159d7630c926b7ca6cb8c012d8609293e46db0b1be7f217d/README.md) |
| 10&#160;Dec&#160;22&#160;20:27&#160;UTC | SHAKEN 726J | 17&#160;Dec&#160;22&#160;20:27&#160;UTC | true | [view](CERTS/4124d728e694d86020281ec3d2cbc38df0a1fa0104eb1049abbc5f77283d0cf8/README.md) |
| 10&#160;Dec&#160;22&#160;20:27&#160;UTC | SHAKEN 790J | 17&#160;Dec&#160;22&#160;20:27&#160;UTC | true | [view](CERTS/7b1e4e99647893cb3d6774b5ff13fc15e30ed61a6c8af049f75b528778da6b5f/README.md) |
| 10&#160;Dec&#160;22&#160;20:28&#160;UTC | SHAKEN 738J | 17&#160;Dec&#160;22&#160;20:28&#160;UTC | true | [view](CERTS/6ff6ba3e2275be1b5e5fa4349a316ab80ca9d5b36c93b65114d1e8d47d0de1e2/README.md) |
| 10&#160;Dec&#160;22&#160;20:28&#160;UTC | SHAKEN 459J | 17&#160;Dec&#160;22&#160;20:28&#160;UTC | true | [view](CERTS/d3fcb07b828ae15c1c9d9ac00c44f2771cf72557635791dadc31cfe3d3f3ebb1/README.md) |
| 10&#160;Dec&#160;22&#160;20:29&#160;UTC | SHAKEN 366G | 17&#160;Dec&#160;22&#160;20:29&#160;UTC | true | [view](CERTS/1efd59ce1b57bccdacec422adb4d52d5d66184d8a8cb3086afceaaf92479b3f5/README.md) |
| 10&#160;Dec&#160;22&#160;20:29&#160;UTC | SHAKEN 0226 | 17&#160;Dec&#160;22&#160;20:29&#160;UTC | true | [view](CERTS/e162ab2e8de82f27b19772dfe4024cd59fa9884e61b1192602d6073040c8edb9/README.md) |
| 10&#160;Dec&#160;22&#160;20:29&#160;UTC | SHAKEN 738J | 17&#160;Dec&#160;22&#160;20:29&#160;UTC | true | [view](CERTS/d7c02e574eb1f64712ae7891b67cc3195cc51836619f6a7868f30b1e8cd95a92/README.md) |
| 10&#160;Dec&#160;22&#160;21:02&#160;UTC | SHAKEN 518J | 17&#160;Dec&#160;22&#160;21:02&#160;UTC | true | [view](CERTS/35f9e8af3f166a65df27c6eadaf5d65a7977958b3ee9176fabe10ab353863882/README.md) |
| 11&#160;Dec&#160;22&#160;20:18&#160;UTC | SHAKEN 983J | 18&#160;Dec&#160;22&#160;20:18&#160;UTC | true | [view](CERTS/3d8d1ab66f127d204d94dd17537d725a6b9d9924dfd2e2b6084d3cc40ea96cfa/README.md) |
| 11&#160;Dec&#160;22&#160;21:45&#160;UTC | SHAKEN 606F | 18&#160;Dec&#160;22&#160;21:45&#160;UTC | true | [view](CERTS/4807a347ac25d603fc1d4851813375602c4efde6ecef01676d09cda953b7cfb2/README.md) |
| 11&#160;Dec&#160;22&#160;22:28&#160;UTC | SHAKEN 120K | 18&#160;Dec&#160;22&#160;22:28&#160;UTC | true | [view](CERTS/608ee867cf61116ebe11c1c16f03629aa5283bf1ba2a7f860b07e2881cf4b255/README.md) |
| 12&#160;Dec&#160;22&#160;01:44&#160;UTC | SHAKEN 278K | 19&#160;Dec&#160;22&#160;01:44&#160;UTC | true | [view](CERTS/779318d9559d6d699e61eda9216d72cbbe47707983cf0ae5be1de534febf8dd4/README.md) |
| 12&#160;Dec&#160;22&#160;05:27&#160;UTC | SHAKEN 2894 | 15&#160;Dec&#160;22&#160;05:27&#160;UTC | true | [view](CERTS/9c6d16fb625fae9f60d49445932c1e58cad1f46f29bf69ce66e60d0c813ab81d/README.md) |
| 12&#160;Dec&#160;22&#160;15:43&#160;UTC | SHAKEN 722J | 11&#160;Jan&#160;23&#160;15:43&#160;UTC | true | [view](CERTS/234bb11af3db6a37bc8a1807ba5458331e17031ec7772a9d790809a0c2856761/README.md) |
| 12&#160;Dec&#160;22&#160;18:12&#160;UTC | SHAKEN 952J | 11&#160;Jan&#160;23&#160;18:12&#160;UTC | true | [view](CERTS/060316d5062c63a01850cffc57dd7d976e37d02e17d9f77d165ba3a16c1172e0/README.md) |
| 12&#160;Dec&#160;22&#160;20:27&#160;UTC | SHAKEN 848J | 19&#160;Dec&#160;22&#160;20:27&#160;UTC | true | [view](CERTS/0fa6767e8ab704048bcfb0fd46ae496e2a5fd201e42dead0eb3462f8edb1c8c9/README.md) |
| 12&#160;Dec&#160;22&#160;20:28&#160;UTC | SHAKEN 738J | 21&#160;Dec&#160;22&#160;20:28&#160;UTC | true | [view](CERTS/759ad048688565b86768369621e35e96a6ad187ed59bbfa9067ff73f7d6596e4/README.md) |

#### CA Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 20&#160;Aug&#160;21&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Issuing CA3 | 19&#160;Aug&#160;31&#160;23:59&#160;UTC | true | [view](CERTS/51b51431ffd8055cd7e30c213540f4868258058d4c4c9f23f1ed8bea75101e37/README.md) |
| 20&#160;Aug&#160;21&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Issuing CA1 | 19&#160;Aug&#160;31&#160;23:59&#160;UTC | true | [view](CERTS/037a7c6621814bd5f01fa2080a77a245458694d9d566e0050e068fc853317857/README.md) |
| 20&#160;Aug&#160;21&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Root CA1 | 19&#160;Aug&#160;41&#160;23:59&#160;UTC | true | [view](CERTS/59ab943f4e8c3c17755cc8e4abb5ac65736ce74ec587c62cb744ab0babbbe2fe/README.md) |
| 24&#160;Oct&#160;22&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Issuing CA4 | 23&#160;Oct&#160;32&#160;23:59&#160;UTC | false | [view](CERTS/80dd12d935f46c256c7c1289a0834b21ec8f2f8c35f2864928d1e75f3a280665/README.md) |
| 24&#160;Oct&#160;22&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Root CA2 | 23&#160;Oct&#160;42&#160;23:59&#160;UTC | false | [view](CERTS/a26e04fc786ab70b8085236b2c53f8cfbf5d0c6a5c2c9c3e9f91669fbb8ea4d5/README.md) |


Generated: 12 Dec 22 23:45 UTC