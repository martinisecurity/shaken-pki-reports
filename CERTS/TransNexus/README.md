# STIR/SHAKEN CA Ecosystem Compliance

## TransNexus

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 330 certificates were included in the corpus being tested
- 0 repositories in the corpus were skipped because they are duplicates
- 267 certificates in the corpus were skipped because they are expired
- 1 certificates in the corpus were skipped because they are not currently trusted
- 62 certificates being tested against the remaining rules
- 100.00% of certificates contain one or more Error level issue
- 100.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 3.23% of certificates are too old to be assessed against currently enforced expectations
- 81 days is the average remaining validity for the certificates in the corpus
- 83 days is the average initial validity for the certificates in the corpus
- 46 certificates expire in the next 30 days

| Instances | Test Status | Source |
|-----------|-------------|--------|
| 4 | [e_cp1_3_ambiguous_identifier](ISSUES/e_cp1_3_ambiguous_identifier/README.md) | United States SHAKEN CP |
| 62 | [e_cp1_3_subject_sn](ISSUES/e_cp1_3_subject_sn/README.md) | United States SHAKEN CP |
| 46 | [e_incorrect_ku_encoding](ISSUES/e_incorrect_ku_encoding/README.md) | RFC5280 |
| 3 | [e_sti_certificate_policies](ISSUES/e_sti_certificate_policies/README.md) | ATIS-1000080 |
| 2 | [e_sti_extension_unknown](ISSUES/e_sti_extension_unknown/README.md) | ATIS-1000080 |
| 2 | [e_sti_serial_number](ISSUES/e_sti_serial_number/README.md) | ATIS-1000080 |
| 2 | [e_sti_signature_algorithm](ISSUES/e_sti_signature_algorithm/README.md) | ATIS-1000080 |
| 2 | [e_sti_subject_cn](ISSUES/e_sti_subject_cn/README.md) | ATIS-1000080 |
| 1 | [w_cp_1_3_subject_email](ISSUES/w_cp_1_3_subject_email/README.md) | United States SHAKEN CP |
| 62 | [w_pki_subject_rdn_unknown](ISSUES/w_pki_subject_rdn_unknown/README.md) | SHAKEN PKI Best Practice |

#### CA Certificates

- 4 certificates were included in the corpus being tested
- 0 repositories in the corpus were skipped because they are duplicates
- 0 certificates in the corpus were skipped because they are expired
- 1 certificates in the corpus were skipped because they are not currently trusted
- 3 certificates being tested against the remaining rules
- 100.00% of certificates contain one or more Error level issue
- 100.00% of certificates contain one or more Warning level issue
- 100.00% of certificates contain one or more Notice level issue
- 100.00% of certificates are too old to be assessed against currently enforced expectations
- 5369 days is the average remaining validity for the certificates in the corpus
- 4870 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days

| Instances | Test Status | Source |
|-----------|-------------|--------|
| 3 | [e_cp1_3_ca_key_usage_crl_sign](ISSUES/e_cp1_3_ca_key_usage_crl_sign/README.md) | United States SHAKEN CP |
| 3 | [e_incorrect_ku_encoding](ISSUES/e_incorrect_ku_encoding/README.md) | RFC5280 |
| 2 | [e_sti_ca_certificate_policies](ISSUES/e_sti_ca_certificate_policies/README.md) | ATIS-1000080 |
| 2 | [e_sti_ca_crl_distribution](ISSUES/e_sti_ca_crl_distribution/README.md) | ATIS-1000080 |
| 2 | [e_sti_ca_extension_unknown](ISSUES/e_sti_ca_extension_unknown/README.md) | ATIS-1000080 |
| 3 | [e_sti_ca_serial_number](ISSUES/e_sti_ca_serial_number/README.md) | ATIS-1000080 |
| 3 | [e_sti_ca_subject_cn](ISSUES/e_sti_ca_subject_cn/README.md) | ATIS-1000080 |
| 1 | [e_sti_root_certificate_policies](ISSUES/e_sti_root_certificate_policies/README.md) | ATIS-1000080 |
| 1 | [e_sti_root_extension_unknown](ISSUES/e_sti_root_extension_unknown/README.md) | ATIS-1000080 |
| 3 | [n_pki_ca_key_usage](ISSUES/n_pki_ca_key_usage/README.md) | SHAKEN PKI Best Practice |
| 2 | [n_sti_ca_certificate_policy_critical](ISSUES/n_sti_ca_certificate_policy_critical/README.md) | ATIS-1000080 |
| 3 | [w_pki_ca_subject_rdn_unknown](ISSUES/w_pki_ca_subject_rdn_unknown/README.md) | SHAKEN PKI Best Practice |

### Details

#### Leaf Certificates

| Created at | Subject | Problems | Link |
|------------|---------|----------|------|
| 02 Nov 21 16:40 UTC | Charter Communications Inc SHAKEN 5606 | true | [view](CERTS/6e6a467b634c931ad2b441fcdbeeec1d29fb3b3a9687018774c563fa866e00ed/README.md) |
| 13 Jan 22 14:17 UTC | MobileSphere SHAKEN 873J | true | [view](CERTS/8f74ddd5e6964042b119d5b3f17d24694ec78c8688e0ffc96d13fdb0e8df26d4/README.md) |
| 30 Mar 22 12:54 UTC | Fusion Connect SHAKEN 2720 | true | [view](CERTS/03a010b294807e90cc1309d7466fd2ddc591158a69be6950ff4d0eab32de0799/README.md) |
| 27 Apr 22 18:22 UTC | CCI SHAKEN 663J | true | [view](CERTS/9cfd0b9974acdc9b5fce3888cf2f6cf99173e7955adeebd37a1e60dfc444a0b3/README.md) |
| 10 Jun 22 14:00 UTC | SHAKEN 597J | true | [view](CERTS/4baa838eb96ce358aa1565d0add8b75072258efd5d21a5bbb0eec878026a67be/README.md) |
| 20 Jun 22 20:25 UTC | SHAKEN 857J | true | [view](CERTS/e1e0a4851fb7b51d617959c00fb6a17e7b23d34e1555ba70ab01c9182a2f85f8/README.md) |
| 22 Jun 22 18:46 UTC | SHAKEN 755J | true | [view](CERTS/a8cac316fddf25b48b9350a325791be8b87e6484a5959ab605db1ae5e3e99a84/README.md) |
| 29 Jun 22 20:24 UTC | SHAKEN 736J | true | [view](CERTS/71b25a480077816998e48e6a76f272e70aa91ecb19b0e773cd8777823cdead32/README.md) |
| 25 Jul 22 18:36 UTC | SHAKEN 578J | true | [view](CERTS/f4af5a28131d680ef2682ae86181b64f0c919864d7dbaef00c0598d314b537a1/README.md) |
| 10 Aug 22 18:11 UTC | SHAKEN 073H | true | [view](CERTS/87dabb91a689a6ac540b5ffbf5eb89326fada83fb822d2bb9ec31b68251b579e/README.md) |
| 01 Sep 22 03:25 UTC | SHAKEN 6628 | true | [view](CERTS/37b9a67bb8e3272048330f269b5dbc285ac0278df4670ab7e291b92b6548fa2d/README.md) |
| 13 Sep 22 17:24 UTC | SHAKEN 706J | true | [view](CERTS/1eea5e17ae8b36cb5b6d657178273af745cea015e63586e6d44d1ba2ee8f722a/README.md) |
| 26 Sep 22 18:40 UTC | SHAKEN 505J | true | [view](CERTS/b186c959554e652b69824a2e45ac08ec135105b0e01c9d2eadeb3cf46c130670/README.md) |
| 26 Sep 22 18:42 UTC | SHAKEN 505J | true | [view](CERTS/1c0003ac10eeaa04229e507c15f71ee018c3902c3e8c20fbe42e533b8682ba8c/README.md) |
| 01 Oct 22 12:31 UTC | SHAKEN 193E | true | [view](CERTS/2a2a297de4aa7620b8f38d5266dd4f76fbd6b5c79a87411b589c2226518ce475/README.md) |
| 10 Oct 22 18:30 UTC | SHAKEN 115K | true | [view](CERTS/631861bcfeb6aeaf83a56e5c1c376c8981927a52fc8da57964b1776b69e2bf6e/README.md) |
| 10 Oct 22 18:39 UTC | SHAKEN 066K | true | [view](CERTS/51c8cf0bbfb640f7ef4e7e47bedd1a9a1d4f757fe142ad05fa84422544794ccd/README.md) |
| 12 Oct 22 19:39 UTC | SHAKEN 815G | true | [view](CERTS/a79f50d4fcd3e680518379f12a86a4b9a21c5e27dc5422a8439525edf350c4a7/README.md) |
| 13 Oct 22 18:02 UTC | SHAKEN 952J | true | [view](CERTS/bbccd67a84881c78b37ecfb961dd7029498490975d91e15f872d5a684a225626/README.md) |
| 20 Oct 22 09:13 UTC | SHAKEN 841J | true | [view](CERTS/8c54ce0987e3860a3b0863c2b6947d4c38a04acf52ebeed106e0af6e5300abcb/README.md) |
| 20 Oct 22 15:48 UTC | SHAKEN 622J | true | [view](CERTS/2e34c765f110b197f2ec1052b1960bc25ce320dfce7d7035cc532cd1eb6dbd57/README.md) |
| 22 Oct 22 04:04 UTC | SHAKEN 345J | true | [view](CERTS/27c018db560aefbff83f32326db2cbf3f1260051386569d7775cf98fae145bae/README.md) |
| 22 Oct 22 04:44 UTC | SHAKEN 345J | true | [view](CERTS/109a4ac394ec4e6eb6ba42d2ad6bc799f2548629a8bbfb9420b9ec1ca5df6ff9/README.md) |
| 24 Oct 22 20:22 UTC | SHAKEN 738J | true | [view](CERTS/1f99a81d0b090f467a47da69b1549fbabc7d8277f67a1f31d05c53e92cceb54c/README.md) |
| 25 Oct 22 14:33 UTC | SHAKEN 2550 | true | [view](CERTS/6970b8cb5229451110fe03840a2927e3d4d4dc548b6d5b4658097a6d996e1ea0/README.md) |
| 25 Oct 22 15:33 UTC | SHAKEN 722J | true | [view](CERTS/50eb0c6a670f4122f8fbdb75582aa257fe3979f441fb396ff738372627104f9c/README.md) |
| 26 Oct 22 08:39 UTC | SHAKEN 012K | true | [view](CERTS/a4a8fa499de2764fbb7db0b76e98d31078179de61e21e96df01f8132a44c1fd4/README.md) |
| 26 Oct 22 10:43 UTC | SHAKEN 8526 | true | [view](CERTS/b68dadcfa6267b8e9dd012e42154b55cab6d3694f4543e46af44bab1e4ba971e/README.md) |
| 26 Oct 22 13:45 UTC | SHAKEN 056J | true | [view](CERTS/860b047bd23110db2919afd52c5e095799e564dfe075a3d67e77c5883a4885bf/README.md) |
| 26 Oct 22 17:24 UTC | SHAKEN 107K | true | [view](CERTS/ab2574c5a2c3ecdc234e72fdc62635cc97f596dc21a512a916d39aa17349719c/README.md) |
| 26 Oct 22 18:13 UTC | SHAKEN 060K | true | [view](CERTS/8fc41fc2e5615fc87a1a9c1fa8cb6ea1c9200d551f5e988047221c335be31d39/README.md) |
| 26 Oct 22 18:36 UTC | SHAKEN 089K | true | [view](CERTS/3e4203b5c20a0a59801040c172eed8f39d2a609fd2429e4a111f40bd0d51c7c0/README.md) |
| 26 Oct 22 18:38 UTC | SHAKEN 9714 | true | [view](CERTS/3c8ced9819143833a95bace3c72bebb882cef867f17509dc4628e18dd74341b2/README.md) |
| 26 Oct 22 20:17 UTC | SHAKEN 674J | true | [view](CERTS/6fd005ce139c83958c022da431a4a721569905d4b6b43de0375a1f7d4b98d275/README.md) |
| 26 Oct 22 20:18 UTC | SHAKEN 735J | true | [view](CERTS/9316e334abe6c964bac27767052062b6c0de27827a606d1cd746e95f97472411/README.md) |
| 26 Oct 22 20:19 UTC | SHAKEN 738J | true | [view](CERTS/3a7079d02db69f8133c4892893d6d94a5b893aca3578938a74ed360392488012/README.md) |
| 26 Oct 22 20:22 UTC | SHAKEN 469A | true | [view](CERTS/8d2fef64712c082b249ae7805b66ad21f2aef4fee890f786747f3bcf77051ed7/README.md) |
| 26 Oct 22 20:22 UTC | SHAKEN 790J | true | [view](CERTS/854b6402cb0b321c3d810426dbf87b6b7b98210f836a6162f00759a448e96d9a/README.md) |
| 26 Oct 22 20:22 UTC | SHAKEN 721J | true | [view](CERTS/b245669c56c66223f8e5f946e40a3e4a2fd17a126e19bc5149a3e310f5448fb2/README.md) |
| 26 Oct 22 20:23 UTC | SHAKEN 625J | true | [view](CERTS/436afa7fd59130c2cfb61515f1881bdb6b26df00104fc20a34813ec2837ae5bc/README.md) |
| 26 Oct 22 20:23 UTC | SHAKEN 738J | true | [view](CERTS/24b935c4e2f91a82ff346147d216ffcfbbb408b7d41886788c3e7158c4573825/README.md) |
| 26 Oct 22 20:23 UTC | SHAKEN 459J | true | [view](CERTS/b8c1d7beff3357f2ba432ff63529da29fd4871b1e11ae59e0f3d30eb6f7b4b51/README.md) |
| 26 Oct 22 20:25 UTC | SHAKEN 366G | true | [view](CERTS/e3f28ec5803737280fee53e08b1097a304fc1b27cee05c9edb7836ce26691681/README.md) |
| 26 Oct 22 20:25 UTC | SHAKEN 0226 | true | [view](CERTS/fd4e9f41bea50e0e2615d2a9686e0dc7d4363c67912c4d355df0e4ec9b3faa14/README.md) |
| 26 Oct 22 20:25 UTC | SHAKEN 738J | true | [view](CERTS/e227f5d1d545dc5289555dfb281c3b6e3642834bad66c25d997dedeee729f046/README.md) |
| 26 Oct 22 22:00 UTC | SHAKEN 551G | true | [view](CERTS/6052628e9622c0b9247ee1139aec68db77c73a3857882d712de67adf4cce6083/README.md) |
| 27 Oct 22 16:32 UTC | SHAKEN 291K | true | [view](CERTS/deefe1aa4ab2ec077f2a86adefc2664087be0c4392d579fc0ed49b04fe7345f6/README.md) |
| 27 Oct 22 19:07 UTC | SHAKEN 749J | true | [view](CERTS/1fc9e8d8a8e19300e114d8b68b9eed28cad0ceb35dc90bb4ca4d07c39116248d/README.md) |
| 27 Oct 22 20:14 UTC | SHAKEN 983J | true | [view](CERTS/61a162ef08cf3cf9690416ee7985d04e655f82a4e2c6bd0760c006ba76794a78/README.md) |
| 28 Oct 22 03:01 UTC | SHAKEN 0172 | true | [view](CERTS/8a54ee5e7989bd5a6012afb4d66f1098b5e9aeda34642705d95cfc2a10e0128e/README.md) |
| 28 Oct 22 18:52 UTC | SHAKEN 841J | true | [view](CERTS/fffe0d037c365187a0e42767ad336159820f4bede1c69529e1492bb85050ad17/README.md) |
| 29 Oct 22 20:23 UTC | SHAKEN 790J | true | [view](CERTS/ae76f8535588450356145362122cfba3a0cb6994abc1c127fe67fcb9e3e0609d/README.md) |
| 29 Oct 22 20:24 UTC | SHAKEN 738J | true | [view](CERTS/1f2fe5e7a7d77bbe16aaa9ec9b2b79c923df88399b729c1abfb58c5fed0ea53a/README.md) |
| 29 Oct 22 20:24 UTC | SHAKEN 459J | true | [view](CERTS/05922e5a7474d0a4341ce93d68289b13cd5a2dccb22bde5f63c7fa5c636329b4/README.md) |
| 29 Oct 22 20:25 UTC | SHAKEN 366G | true | [view](CERTS/9fa4b48801d0e1a2ef02989c6ff4c86935d700ecce028fd84ae0238a8b136483/README.md) |
| 29 Oct 22 20:25 UTC | SHAKEN 0226 | true | [view](CERTS/29dca9044c516595332f19bdcd9020b3d7793ee6703ca8509643e8945f5d364b/README.md) |
| 29 Oct 22 20:25 UTC | SHAKEN 738J | true | [view](CERTS/b851335cc2cfc55f8a63e4e73bd3085f4d4cc44ed0d2850adf9eeded84b568cd/README.md) |
| 30 Oct 22 20:14 UTC | SHAKEN 983J | true | [view](CERTS/285d5bb65f645b987e1ee7ce6c08dff1f1ac700a48ec240a808015cc56dc012d/README.md) |
| 31 Oct 22 12:32 UTC | SHAKEN 193E | true | [view](CERTS/0de5e0e6787b9aa19c02c014821dd7a2e62e438f4eb9fbbb761eb86df8c69ff7/README.md) |
| 31 Oct 22 14:33 UTC | SHAKEN 2550 | true | [view](CERTS/87e20c4fe066f45021a5353c35fc4d0cfe05367ea8470c8d6749deb64141d2d4/README.md) |
| 31 Oct 22 16:03 UTC | SHAKEN 2277 | true | [view](CERTS/35ce538300b59678c2159df9fc2d9eb909d055ed309094d3843fef76c093561e/README.md) |
| 31 Oct 22 18:04 UTC | SHAKEN 952J | true | [view](CERTS/93040a7e0cb4876c378dd10737f7dec5d054a6a744ebec63e17f8adda605a6ca/README.md) |

#### CA Certificates

| Created at | Subject | Problems | Link |
|------------|---------|----------|------|
| 20 Aug 21 00:00 UTC | TransNexus, Inc. SHAKEN Issuing CA3 | true | [view](CERTS/51b51431ffd8055cd7e30c213540f4868258058d4c4c9f23f1ed8bea75101e37/README.md) |
| 20 Aug 21 00:00 UTC | TransNexus, Inc. SHAKEN Issuing CA1 | true | [view](CERTS/037a7c6621814bd5f01fa2080a77a245458694d9d566e0050e068fc853317857/README.md) |
| 20 Aug 21 00:00 UTC | TransNexus, Inc. SHAKEN Root CA1 | true | [view](CERTS/59ab943f4e8c3c17755cc8e4abb5ac65736ce74ec587c62cb744ab0babbbe2fe/README.md) |


Generated: 01/11/2022 at 10:05:32