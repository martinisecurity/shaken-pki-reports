# STIR/SHAKEN CA Ecosystem Compliance

## TransNexus

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 362 certificates were included in the corpus being tested
- 1 certificates in the corpus were skipped because they are duplicates
- 293 certificates in the corpus were skipped because they are expired
- 1 certificates in the corpus were skipped because they are not currently trusted
- 67 certificates being tested against the remaining rules
- 2.46 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 100.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 1.49% of certificates are too old to be assessed against currently enforced expectations
- 73 days is the average remaining validity for the certificates in the corpus
- 72 days is the average initial validity for the certificates in the corpus
- 51 certificates expire in the next 30 days

| Instances | Test | Source |
|-----------|------|--------|
| 3 | [e_atis_certificate_policies](ISSUES/e_atis_certificate_policies/README.md) | ATIS1000080 |
| 25 | [e_incorrect_ku_encoding](ISSUES/e_incorrect_ku_encoding/README.md) | RFC5280 |
| 3 | [e_us_cp_ambiguous_identifier](ISSUES/e_us_cp_ambiguous_identifier/README.md) | US_SHAKEN_CP |
| 67 | [e_us_cp_subject_sn](ISSUES/e_us_cp_subject_sn/README.md) | US_SHAKEN_CP |
| 67 | [w_shaken_subject_rdn_unknown](ISSUES/w_shaken_subject_rdn_unknown/README.md) | SHAKEN_PKI_BEST_PRACTICES |

#### CA Certificates

- 4 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 0 certificates in the corpus were skipped because they are expired
- 1 certificates in the corpus were skipped because they are not currently trusted
- 3 certificates being tested against the remaining rules
- 3.00 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 100.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 100.00% of certificates are too old to be assessed against currently enforced expectations
- 5368 days is the average remaining validity for the certificates in the corpus
- 4870 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days

| Instances | Test | Source |
|-----------|------|--------|
| 3 | [e_incorrect_ku_encoding](ISSUES/e_incorrect_ku_encoding/README.md) | RFC5280 |
| 3 | [e_us_cp_ca_key_usage_crl_sign](ISSUES/e_us_cp_ca_key_usage_crl_sign/README.md) | US_SHAKEN_CP |
| 3 | [w_shaken_ca_subject_rdn_unknown](ISSUES/w_shaken_ca_subject_rdn_unknown/README.md) | SHAKEN_PKI_BEST_PRACTICES |

### Details

#### Leaf Certificates

| Created at | Subject | Problems | Link |
|------------|---------|----------|------|
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
| 10 Oct 22 18:27 UTC | SHAKEN 166K | true | [view](CERTS/d4d723c554cc42a2cd242010ffd2499aab6436955e8031d4f3109db3b66a3b64/README.md) |
| 10 Oct 22 18:30 UTC | SHAKEN 115K | true | [view](CERTS/631861bcfeb6aeaf83a56e5c1c376c8981927a52fc8da57964b1776b69e2bf6e/README.md) |
| 10 Oct 22 18:36 UTC | SHAKEN 031K | true | [view](CERTS/a6b43e7cf487f1fd10f20b7197a556d8aecab5c71e91ab87a634894ca4fb00f6/README.md) |
| 10 Oct 22 18:39 UTC | SHAKEN 066K | true | [view](CERTS/51c8cf0bbfb640f7ef4e7e47bedd1a9a1d4f757fe142ad05fa84422544794ccd/README.md) |
| 12 Oct 22 19:39 UTC | SHAKEN 815G | true | [view](CERTS/a79f50d4fcd3e680518379f12a86a4b9a21c5e27dc5422a8439525edf350c4a7/README.md) |
| 13 Oct 22 18:02 UTC | SHAKEN 952J | true | [view](CERTS/bbccd67a84881c78b37ecfb961dd7029498490975d91e15f872d5a684a225626/README.md) |
| 20 Oct 22 15:48 UTC | SHAKEN 622J | true | [view](CERTS/2e34c765f110b197f2ec1052b1960bc25ce320dfce7d7035cc532cd1eb6dbd57/README.md) |
| 22 Oct 22 04:04 UTC | SHAKEN 345J | true | [view](CERTS/27c018db560aefbff83f32326db2cbf3f1260051386569d7775cf98fae145bae/README.md) |
| 22 Oct 22 04:44 UTC | SHAKEN 345J | true | [view](CERTS/109a4ac394ec4e6eb6ba42d2ad6bc799f2548629a8bbfb9420b9ec1ca5df6ff9/README.md) |
| 25 Oct 22 15:33 UTC | SHAKEN 722J | true | [view](CERTS/50eb0c6a670f4122f8fbdb75582aa257fe3979f441fb396ff738372627104f9c/README.md) |
| 26 Oct 22 10:43 UTC | SHAKEN 8526 | true | [view](CERTS/b68dadcfa6267b8e9dd012e42154b55cab6d3694f4543e46af44bab1e4ba971e/README.md) |
| 28 Oct 22 03:01 UTC | SHAKEN 0172 | true | [view](CERTS/8a54ee5e7989bd5a6012afb4d66f1098b5e9aeda34642705d95cfc2a10e0128e/README.md) |
| 28 Oct 22 18:52 UTC | SHAKEN 841J | true | [view](CERTS/fffe0d037c365187a0e42767ad336159820f4bede1c69529e1492bb85050ad17/README.md) |
| 29 Oct 22 20:19 UTC | SHAKEN 738J | true | [view](CERTS/e8b37b41b0774a63135275e04ea75a1ce6754e1cf3efc56e4625f7ad6d467cbc/README.md) |
| 29 Oct 22 20:19 UTC | SHAKEN 700H | true | [view](CERTS/d8142d12123966456ea28d898006b88990bebad9ae5e466b7246b25d8d5ff69f/README.md) |
| 29 Oct 22 20:21 UTC | SHAKEN 854D | true | [view](CERTS/393859cb08a16390bb1c01bbd326b02d7d7e7b1482fb5a0ee14ec4da0af451c7/README.md) |
| 29 Oct 22 20:23 UTC | SHAKEN 726J | true | [view](CERTS/af138aa33ad44d44f16c4615aac8299b01edfa6b4a9a0be9b082d07003952bcf/README.md) |
| 29 Oct 22 20:23 UTC | SHAKEN 790J | true | [view](CERTS/ae76f8535588450356145362122cfba3a0cb6994abc1c127fe67fcb9e3e0609d/README.md) |
| 29 Oct 22 20:24 UTC | SHAKEN 738J | true | [view](CERTS/1f2fe5e7a7d77bbe16aaa9ec9b2b79c923df88399b729c1abfb58c5fed0ea53a/README.md) |
| 29 Oct 22 20:24 UTC | SHAKEN 459J | true | [view](CERTS/05922e5a7474d0a4341ce93d68289b13cd5a2dccb22bde5f63c7fa5c636329b4/README.md) |
| 29 Oct 22 20:25 UTC | SHAKEN 366G | true | [view](CERTS/9fa4b48801d0e1a2ef02989c6ff4c86935d700ecce028fd84ae0238a8b136483/README.md) |
| 29 Oct 22 20:25 UTC | SHAKEN 0226 | true | [view](CERTS/29dca9044c516595332f19bdcd9020b3d7793ee6703ca8509643e8945f5d364b/README.md) |
| 29 Oct 22 20:25 UTC | SHAKEN 738J | true | [view](CERTS/b851335cc2cfc55f8a63e4e73bd3085f4d4cc44ed0d2850adf9eeded84b568cd/README.md) |
| 29 Oct 22 20:49 UTC | SHAKEN 518J | true | [view](CERTS/10c03645d46583869470a4ebebbcc3285d717f931128d773def089ea19b4c349/README.md) |
| 30 Oct 22 15:44 UTC | SHAKEN 793J | true | [view](CERTS/8e0c8e856d277087efe6c7445f34d96106b947494f9ce5f8362dbb1cf922e636/README.md) |
| 30 Oct 22 20:14 UTC | SHAKEN 983J | true | [view](CERTS/285d5bb65f645b987e1ee7ce6c08dff1f1ac700a48ec240a808015cc56dc012d/README.md) |
| 30 Oct 22 21:41 UTC | SHAKEN 606F | true | [view](CERTS/40341cbe06af452ba8ae3a3c611a55a5aeef513ac74c57354d1f109b05223b50/README.md) |
| 30 Oct 22 23:51 UTC | SHAKEN 970J | true | [view](CERTS/4cef965f31ef863571fda568d639c7ad03047a072e14f91b3f67b16c3bd7075e/README.md) |
| 31 Oct 22 12:32 UTC | SHAKEN 193E | true | [view](CERTS/0de5e0e6787b9aa19c02c014821dd7a2e62e438f4eb9fbbb761eb86df8c69ff7/README.md) |
| 31 Oct 22 14:33 UTC | SHAKEN 2550 | true | [view](CERTS/87e20c4fe066f45021a5353c35fc4d0cfe05367ea8470c8d6749deb64141d2d4/README.md) |
| 31 Oct 22 16:03 UTC | SHAKEN 2277 | true | [view](CERTS/35ce538300b59678c2159df9fc2d9eb909d055ed309094d3843fef76c093561e/README.md) |
| 31 Oct 22 18:04 UTC | SHAKEN 952J | true | [view](CERTS/93040a7e0cb4876c378dd10737f7dec5d054a6a744ebec63e17f8adda605a6ca/README.md) |
| 31 Oct 22 19:42 UTC | SHAKEN 551G | true | [view](CERTS/2365797797b5dbcbb9ba052f2b000d8ec74a4bcc179dc804a319f7ee6bdaba92/README.md) |
| 31 Oct 22 20:23 UTC | SHAKEN 738J | true | [view](CERTS/a0263a15ca1cfa43364234e6efd2ebbf63c0de798f997b821037b886a57d401f/README.md) |
| 01 Nov 22 07:54 UTC | SHAKEN 0172 | true | [view](CERTS/da115d82c8349414f0301412e7dd0b15a62ee5139be1c44f9ab3919d98892c81/README.md) |
| 01 Nov 22 08:39 UTC | SHAKEN 012K | true | [view](CERTS/0109d51cebdf401540978141570a5a1aadeb0fd6002c21160c6fd8969250b976/README.md) |
| 01 Nov 22 17:24 UTC | SHAKEN 107K | true | [view](CERTS/13f374a0908277e55e93d2fea7850ed060e747e7fa0504e6abda05b1b8897b5d/README.md) |
| 01 Nov 22 18:14 UTC | SHAKEN 060K | true | [view](CERTS/b030c024fa6dce4bdc3396abea1a91832059477d42f5da8349bc08d8bc6b2e0e/README.md) |
| 01 Nov 22 18:36 UTC | SHAKEN 089K | true | [view](CERTS/652d8939c282b5aa9b4802f241bf24b5aac41c47d9ce45da9fe3bb7d12995da9/README.md) |
| 01 Nov 22 20:19 UTC | SHAKEN 738J | true | [view](CERTS/79c9cd15d9fee3269e22f0d42fe0664aa30fb6ff0f0c695faabb92bdfdd67b7e/README.md) |
| 01 Nov 22 20:22 UTC | SHAKEN 854D | true | [view](CERTS/78812b33fd7f91b7afac409193b563bc217dbb6ca632e9c70ff53d4d726faa8b/README.md) |
| 01 Nov 22 20:23 UTC | SHAKEN 849J | true | [view](CERTS/1fa87fba37044222ac41151a8767f7192808a3486317a32eaf57779ba9ca8465/README.md) |
| 01 Nov 22 20:23 UTC | SHAKEN 469A | true | [view](CERTS/87b3fcee6dac3936816284b8a5f6fa2a93394a74df9ee72212833b13f391b14a/README.md) |
| 01 Nov 22 20:23 UTC | SHAKEN 726J | true | [view](CERTS/56d48711f96e5dd79100d5c7faf0028bafa0b3dc933328820ab07d8e3344583c/README.md) |
| 01 Nov 22 20:23 UTC | SHAKEN 790J | true | [view](CERTS/99c39661465e0c3c4e2960d8a128f28927fbca7b93a4354a196841211ca331d5/README.md) |
| 01 Nov 22 20:24 UTC | SHAKEN 738J | true | [view](CERTS/235555acef2b1aba6d597b5b6a737376c2a1d1c8b0c3c441e79f677ed6a818a9/README.md) |
| 01 Nov 22 20:24 UTC | SHAKEN 459J | true | [view](CERTS/fd41ee1d9aff5c387cb37d109eeec217ad940da95326bb8ab506e28de6e4147b/README.md) |
| 01 Nov 22 20:24 UTC | SHAKEN 672B | true | [view](CERTS/e9ecf9074efa7ab923863795d8b78db29bd73f0b4ecce7348792c8f14ef42fd4/README.md) |
| 01 Nov 22 20:25 UTC | SHAKEN 366G | true | [view](CERTS/399102dcf72656b8eef0879d40fa055c6857e3db4598d412414dbd0d47cb6b94/README.md) |
| 01 Nov 22 20:25 UTC | SHAKEN 0226 | true | [view](CERTS/31ae1b04edd9f97bd6e1964edc094700225e0c6577fdf5913d115a2c60f67560/README.md) |
| 01 Nov 22 20:25 UTC | SHAKEN 738J | true | [view](CERTS/b41dc3697994708aeb98d7cbf3e7cc5e2266d2fd8dc2f1978163d3223d46e774/README.md) |
| 01 Nov 22 20:49 UTC | SHAKEN 518J | true | [view](CERTS/9ed9a0f7491e6567aaeace1c4455b7640dba8da8a719c9a2f26ed7be3c27cb0b/README.md) |
| 02 Nov 22 20:15 UTC | SHAKEN 983J | true | [view](CERTS/57d221a8a3c82e338ea3b856d945e44628c17894b976a8b74a1965ad69ef5aa4/README.md) |

#### CA Certificates

| Created at | Subject | Problems | Link |
|------------|---------|----------|------|
| 20 Aug 21 00:00 UTC | TransNexus, Inc. SHAKEN Issuing CA3 | true | [view](CERTS/51b51431ffd8055cd7e30c213540f4868258058d4c4c9f23f1ed8bea75101e37/README.md) |
| 20 Aug 21 00:00 UTC | TransNexus, Inc. SHAKEN Issuing CA1 | true | [view](CERTS/037a7c6621814bd5f01fa2080a77a245458694d9d566e0050e068fc853317857/README.md) |
| 20 Aug 21 00:00 UTC | TransNexus, Inc. SHAKEN Root CA1 | true | [view](CERTS/59ab943f4e8c3c17755cc8e4abb5ac65736ce74ec587c62cb744ab0babbbe2fe/README.md) |


Generated: 04 Nov 22 01:11 UTC