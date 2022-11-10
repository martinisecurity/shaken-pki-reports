# STIR/SHAKEN CA Ecosystem Compliance

## TransNexus

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 406 certificates were included in the corpus being tested
- 1 certificates in the corpus were skipped because they are duplicates
- 337 certificates in the corpus were skipped because they are expired
- 1 certificates in the corpus were skipped because they are not currently trusted
- 67 certificates being tested against the remaining rules
- 2.42 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 100.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 1.49% of certificates are too old to be assessed against currently enforced expectations
- 75 days is the average remaining validity for the certificates in the corpus
- 74 days is the average initial validity for the certificates in the corpus
- 51 certificates expire in the next 30 days

| Instances | Test | Source |
|-----------|------|--------|
| 3 | [e_atis_certificate_policies](ISSUES/e_atis_certificate_policies/README.md) | ATIS1000080 |
| 22 | [e_incorrect_ku_encoding](ISSUES/e_incorrect_ku_encoding/README.md) | RFC5280 |
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
- 5367 days is the average remaining validity for the certificates in the corpus
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
| 17 May 22 18:47 UTC | SHAKEN 4036 | true | [view](CERTS/c957b55c4ccec7e5ddb1673a1c9838bcf8e627bb72c9af9c18442e9489145a4d/README.md) |
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
| 12 Oct 22 19:39 UTC | SHAKEN 815G | true | [view](CERTS/a79f50d4fcd3e680518379f12a86a4b9a21c5e27dc5422a8439525edf350c4a7/README.md) |
| 13 Oct 22 18:02 UTC | SHAKEN 952J | true | [view](CERTS/bbccd67a84881c78b37ecfb961dd7029498490975d91e15f872d5a684a225626/README.md) |
| 20 Oct 22 15:48 UTC | SHAKEN 622J | true | [view](CERTS/2e34c765f110b197f2ec1052b1960bc25ce320dfce7d7035cc532cd1eb6dbd57/README.md) |
| 22 Oct 22 04:04 UTC | SHAKEN 345J | true | [view](CERTS/27c018db560aefbff83f32326db2cbf3f1260051386569d7775cf98fae145bae/README.md) |
| 22 Oct 22 04:44 UTC | SHAKEN 345J | true | [view](CERTS/109a4ac394ec4e6eb6ba42d2ad6bc799f2548629a8bbfb9420b9ec1ca5df6ff9/README.md) |
| 25 Oct 22 15:33 UTC | SHAKEN 722J | true | [view](CERTS/50eb0c6a670f4122f8fbdb75582aa257fe3979f441fb396ff738372627104f9c/README.md) |
| 26 Oct 22 10:43 UTC | SHAKEN 8526 | true | [view](CERTS/b68dadcfa6267b8e9dd012e42154b55cab6d3694f4543e46af44bab1e4ba971e/README.md) |
| 28 Oct 22 18:52 UTC | SHAKEN 841J | true | [view](CERTS/fffe0d037c365187a0e42767ad336159820f4bede1c69529e1492bb85050ad17/README.md) |
| 31 Oct 22 12:32 UTC | SHAKEN 193E | true | [view](CERTS/0de5e0e6787b9aa19c02c014821dd7a2e62e438f4eb9fbbb761eb86df8c69ff7/README.md) |
| 31 Oct 22 18:04 UTC | SHAKEN 952J | true | [view](CERTS/93040a7e0cb4876c378dd10737f7dec5d054a6a744ebec63e17f8adda605a6ca/README.md) |
| 03 Nov 22 14:34 UTC | SHAKEN 2550 | true | [view](CERTS/501b9a063f1a2bdca29969dd18608496a76956e777796c1250222cfa8f8221f8/README.md) |
| 03 Nov 22 17:38 UTC | SHAKEN 177K | true | [view](CERTS/d7a6216497a41058d09aecf76229fa1e61b41450152710c81d49a719fc51207f/README.md) |
| 03 Nov 22 21:53 UTC | SHAKEN 159H | true | [view](CERTS/e31657badaa9f1f7aeeb81b5eda1fbd39649c5783548fd4af66ee2439593e926/README.md) |
| 04 Nov 22 20:18 UTC | SHAKEN 674J | true | [view](CERTS/977c78ccdbfabbc7bef1702d35ec3fc0a1080d8b53e0c78b0e41f06fe18d85ea/README.md) |
| 04 Nov 22 20:19 UTC | SHAKEN 738J | true | [view](CERTS/2b1074ef0b03245bedc267759b511ebf99310da71382c6b0fe0ee15bdb676f77/README.md) |
| 04 Nov 22 20:24 UTC | SHAKEN 738J | true | [view](CERTS/20721ac69d5518fd4297527027f0f7be5b85eee78eea783804ef94fd17ea9003/README.md) |
| 04 Nov 22 20:24 UTC | SHAKEN 459J | true | [view](CERTS/c5bce1300b38df933fe75b3dff26cf5b019a662d38859e2197cd1befd58b9ad3/README.md) |
| 04 Nov 22 20:24 UTC | SHAKEN 672B | true | [view](CERTS/f585d59a376a89665ee8ace7b2dbd591162b61f63feb77a6ad05b32e6126c32a/README.md) |
| 04 Nov 22 20:25 UTC | SHAKEN 366G | true | [view](CERTS/256eebbfca5b7da1a1e66e19f6fd5376627e04c19fb6f37f14e87b96d1b3983c/README.md) |
| 04 Nov 22 20:25 UTC | SHAKEN 0226 | true | [view](CERTS/28a0caef1f4943a8024d11c3b0b6c86c584a0b86105387de89c6085f98dd523b/README.md) |
| 04 Nov 22 20:25 UTC | SHAKEN 738J | true | [view](CERTS/a048311e24949b2dc27dda18181e0a0189d98ca966c1750a457d6c2cde5b42e6/README.md) |
| 04 Nov 22 20:49 UTC | SHAKEN 518J | true | [view](CERTS/a43b22b9f194fc0c709895dfc132cf4fb5a3daee17baadb94a3523f75811b87d/README.md) |
| 05 Nov 22 12:43 UTC | SHAKEN 0172 | true | [view](CERTS/18fc313c2280c303a9b04d3674444612dbff56997327b176e7a885af8bf58f78/README.md) |
| 05 Nov 22 15:45 UTC | SHAKEN 793J | true | [view](CERTS/0dbec1afa193dc20a041304d5d5253aa994ee5722bd8e0a2dc221bfa8e166326/README.md) |
| 05 Nov 22 17:20 UTC | SHAKEN 551G | true | [view](CERTS/bb9ab1aa7cb2a8422f697ea23f6c1feb4e068c1305e087402e5f4a727041e85b/README.md) |
| 05 Nov 22 20:15 UTC | SHAKEN 983J | true | [view](CERTS/9a0da2bb5deeede99c1a0b84e38e223605aa2d2e022284bba8919f374e9ca01e/README.md) |
| 05 Nov 22 22:24 UTC | SHAKEN 120K | true | [view](CERTS/f684f7852d79afa16677830882c6ef2607572f14471df10ede3e9773fe3e6b43/README.md) |
| 06 Nov 22 04:33 UTC | SHAKEN 841J | true | [view](CERTS/83b7a11263e7321e0e1a41808db6ca873aa3c54db4732703196e9fabaadadbcc/README.md) |
| 06 Nov 22 14:34 UTC | SHAKEN 2550 | true | [view](CERTS/29429795816e5203bdee17e7978b5b06e345e8a4b2243f780783af6f04aeefb1/README.md) |
| 06 Nov 22 17:38 UTC | SHAKEN 177K | true | [view](CERTS/6b5aaa57bc81de249e6f6c82c49b65b23051863e9498dd932712d409f31ba8ed/README.md) |
| 07 Nov 22 08:39 UTC | SHAKEN 012K | true | [view](CERTS/fe63f5e2751e5c542e9e125db355abffeabf80c007e45f42ae7d42ab3c36fa5c/README.md) |
| 07 Nov 22 18:14 UTC | SHAKEN 060K | true | [view](CERTS/48bdddbe263b87843824debb51101f02a71a54b54f43bccbd05e8792da4b8050/README.md) |
| 07 Nov 22 20:18 UTC | SHAKEN 674J | true | [view](CERTS/774e3cc7b80cb5d0587c71dfa937ff2afd54edc214cd376cf892d14e97864a9d/README.md) |
| 07 Nov 22 20:19 UTC | SHAKEN 594J | true | [view](CERTS/c39c3f3c1d3022d18ed408051bf5d44b4ef6bdf907a6d85c236aaca005225809/README.md) |
| 07 Nov 22 20:19 UTC | SHAKEN 738J | true | [view](CERTS/893a422bde5665bfa03d25cea02082582885690be7cbedad12cded50455b3768/README.md) |
| 07 Nov 22 20:19 UTC | SHAKEN 700H | true | [view](CERTS/f2822f4c49f0fb3cab3a2e1d8c853fe5b2517eb70fabaf55706de2e996ae5a9e/README.md) |
| 07 Nov 22 20:21 UTC | SHAKEN 819H | true | [view](CERTS/c7ce135a1bb3bcf106b54a05ed88968939344c0130d9900e72654f2049c42428/README.md) |
| 07 Nov 22 20:22 UTC | SHAKEN 727J | true | [view](CERTS/1db6370e04e8fe961f02ccfb1ce821f7c14657871a46682d48a3775542909a3a/README.md) |
| 07 Nov 22 20:23 UTC | SHAKEN 849J | true | [view](CERTS/6ee8cd1092226fd264fc4148dcc9d52ac0e7eb1d0b474edee1a5f84d43a04e08/README.md) |
| 07 Nov 22 20:23 UTC | SHAKEN 726J | true | [view](CERTS/580621096f0c08b33849b0058c4ff41414550150c481ff5432295af8ad2654ab/README.md) |
| 07 Nov 22 20:24 UTC | SHAKEN 738J | true | [view](CERTS/a1a5d33195deb9ac73cfc79e46bab466be179dd0d7b8e0c195b6b3b1a07ff833/README.md) |
| 07 Nov 22 20:24 UTC | SHAKEN 738J | true | [view](CERTS/0eaf78e7598a9627e5e237cac616329576d25df4957737778dfeb600782df809/README.md) |
| 07 Nov 22 20:24 UTC | SHAKEN 459J | true | [view](CERTS/f9d0f94c8a8fc976b3c1cc72ebdf08a7e2f3b6fdeb40f8365ee082553ba2e1b5/README.md) |
| 07 Nov 22 20:24 UTC | SHAKEN 672B | true | [view](CERTS/10a697908e29b643668b811eee2f55b03a7a5f6b8dd11fa1464a572287ad3aaa/README.md) |
| 07 Nov 22 20:25 UTC | SHAKEN 366G | true | [view](CERTS/2a1b0aa5593c0b5f824a6d8d8dc1ab1f37de920fad58298ed5c294adceb8e7fe/README.md) |
| 07 Nov 22 20:26 UTC | SHAKEN 0226 | true | [view](CERTS/e2aa987add842ce641a9323244c2f8da382924191c49135397fc14a8c02df61d/README.md) |
| 07 Nov 22 20:26 UTC | SHAKEN 738J | true | [view](CERTS/799bd1bd91b56bf7249e862e7f2ef394f6424e5da51c73cd12f9b12b3080761b/README.md) |
| 07 Nov 22 20:49 UTC | SHAKEN 518J | true | [view](CERTS/0d679fe61e65d624ffe43d54bbd954d9e387830ef20bae2bf8bc7464f072f76e/README.md) |
| 08 Nov 22 15:45 UTC | SHAKEN 793J | true | [view](CERTS/0407b55a51f1c5473bf9f7877490e1831b458330f4fc53c73012b19e272eea52/README.md) |
| 08 Nov 22 20:15 UTC | SHAKEN 983J | true | [view](CERTS/0e54f14f59f8a37cdafe3c8fbe3cafb1eea1f0e24e0aacb9265705b58356b486/README.md) |
| 08 Nov 22 21:42 UTC | SHAKEN 606F | true | [view](CERTS/57a739a44d59dabc6e1737ef8c54447994d1d612c83711214c9374f530fde7d3/README.md) |
| 09 Nov 22 01:42 UTC | SHAKEN 278K | true | [view](CERTS/02a65a5016113a1cb2de1aa14b686e4b88cf3059a31abb7b3fb0cf1e802d0bd3/README.md) |

#### CA Certificates

| Created at | Subject | Problems | Link |
|------------|---------|----------|------|
| 20 Aug 21 00:00 UTC | TransNexus, Inc. SHAKEN Issuing CA3 | true | [view](CERTS/51b51431ffd8055cd7e30c213540f4868258058d4c4c9f23f1ed8bea75101e37/README.md) |
| 20 Aug 21 00:00 UTC | TransNexus, Inc. SHAKEN Issuing CA1 | true | [view](CERTS/037a7c6621814bd5f01fa2080a77a245458694d9d566e0050e068fc853317857/README.md) |
| 20 Aug 21 00:00 UTC | TransNexus, Inc. SHAKEN Root CA1 | true | [view](CERTS/59ab943f4e8c3c17755cc8e4abb5ac65736ce74ec587c62cb744ab0babbbe2fe/README.md) |


Generated: 10 Nov 22 06:43 UTC