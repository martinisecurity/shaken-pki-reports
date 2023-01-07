# STIR/SHAKEN CA Ecosystem Compliance

## TransNexus

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 1004 certificates were included in the corpus being tested
- 1 certificates in the corpus were skipped because they are duplicates
- 916 certificates in the corpus were skipped because they are expired
- 1 certificates in the corpus were skipped because they are not currently trusted
- 86 certificates being tested against the remaining rules
- 1.37 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 15.12% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 1.16% of certificates are too old to be assessed against currently enforced expectations
- 63 days is the average remaining validity for the certificates in the corpus
- 63 days is the average initial validity for the certificates in the corpus
- 71 certificates expire in the next 30 days
- 1.48 average number of unexpired certificates per OCN observed
- 58 unique OCNs observed in unexpired and valid certificate corpus

| Instances | Test | Source |
|-----------|------|--------|
| 3 | [e_atis_certificate_policies](ISSUES/e_atis_certificate_policies/README.md) | ATIS1000080 |
| 13 | [e_incorrect_ku_encoding](ISSUES/e_incorrect_ku_encoding/README.md) | RFC5280 |
| 3 | [e_us_cp_ambiguous_identifier](ISSUES/e_us_cp_ambiguous_identifier/README.md) | US_SHAKEN_CP |
| 86 | [e_us_cp_subject_sn](ISSUES/e_us_cp_subject_sn/README.md) | US_SHAKEN_CP |
| 13 | [w_shaken_subject_rdn_unknown](ISSUES/w_shaken_subject_rdn_unknown/README.md) | SHAKEN_PKI_BEST_PRACTICES |

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
- 5466 days is the average remaining validity for the certificates in the corpus
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
| 13&#160;Sep&#160;22&#160;17:24&#160;UTC | SHAKEN 706J | 13&#160;Sep&#160;23&#160;17:24&#160;UTC | true | [view](CERTS/1eea5e17ae8b36cb5b6d657178273af745cea015e63586e6d44d1ba2ee8f722a/README.md) |
| 12&#160;Oct&#160;22&#160;19:39&#160;UTC | SHAKEN 815G | 10&#160;Jan&#160;23&#160;19:39&#160;UTC | true | [view](CERTS/a79f50d4fcd3e680518379f12a86a4b9a21c5e27dc5422a8439525edf350c4a7/README.md) |
| 20&#160;Oct&#160;22&#160;15:48&#160;UTC | SHAKEN 622J | 18&#160;Jan&#160;23&#160;15:48&#160;UTC | true | [view](CERTS/2e34c765f110b197f2ec1052b1960bc25ce320dfce7d7035cc532cd1eb6dbd57/README.md) |
| 23&#160;Nov&#160;22&#160;17:31&#160;UTC | SHAKEN 4036 | 10&#160;May&#160;23&#160;17:31&#160;UTC | true | [view](CERTS/e03ff4259e6b456f1bb19d011f005932e83b0a35b520730c97cd77ad8868071c/README.md) |
| 30&#160;Nov&#160;22&#160;12:32&#160;UTC | SHAKEN 193E | 29&#160;Jan&#160;23&#160;12:32&#160;UTC | true | [view](CERTS/b1085c4ed6a32a87dc773d65e4e7aec591f7625ad80f63ee690d56b74b7a1430/README.md) |
| 12&#160;Dec&#160;22&#160;15:43&#160;UTC | SHAKEN 722J | 11&#160;Jan&#160;23&#160;15:43&#160;UTC | true | [view](CERTS/234bb11af3db6a37bc8a1807ba5458331e17031ec7772a9d790809a0c2856761/README.md) |
| 12&#160;Dec&#160;22&#160;18:12&#160;UTC | SHAKEN 952J | 11&#160;Jan&#160;23&#160;18:12&#160;UTC | true | [view](CERTS/060316d5062c63a01850cffc57dd7d976e37d02e17d9f77d165ba3a16c1172e0/README.md) |
| 22&#160;Dec&#160;22&#160;17:30&#160;UTC | SHAKEN 505J | 22&#160;Mar&#160;23&#160;17:30&#160;UTC | true | [view](CERTS/7ff43923df80c3a7e2d854b42f33c69069af4625f5267df0d7ff7b2004ad1c1b/README.md) |
| 24&#160;Dec&#160;22&#160;04:21&#160;UTC | SHAKEN 345J | 23&#160;Jan&#160;23&#160;04:21&#160;UTC | true | [view](CERTS/4fef66289eb1ff86ee5182af55c3cb2a6d4050e31fd08e250807b47e7d2b45ea/README.md) |
| 24&#160;Dec&#160;22&#160;05:00&#160;UTC | SHAKEN 345J | 23&#160;Jan&#160;23&#160;05:00&#160;UTC | true | [view](CERTS/1046866f335f424f2678b093d3af73a2a75c9cf37302e75a94954054f092fc91/README.md) |
| 26&#160;Dec&#160;22&#160;14:30&#160;UTC | SHAKEN 841J | 09&#160;Jan&#160;23&#160;14:30&#160;UTC | true | [view](CERTS/b06c04c6087013d23450759d1cb3836b75f437037e0ac6ee2a32cd8e05ee33cb/README.md) |
| 30&#160;Dec&#160;22&#160;06:30&#160;UTC | SHAKEN 807J | 28&#160;Feb&#160;23&#160;06:30&#160;UTC | true | [view](CERTS/62da6686ffc785eb4f15528d7bf06acd4d4778bbe258808e662ee91ffb868ccf/README.md) |
| 30&#160;Dec&#160;22&#160;12:34&#160;UTC | SHAKEN 193E | 28&#160;Feb&#160;23&#160;12:34&#160;UTC | true | [view](CERTS/36248069600bf4dcf5c85d58b56e6960278411436d818c338226d070a778cf31/README.md) |
| 30&#160;Dec&#160;22&#160;18:15&#160;UTC | SHAKEN 952J | 29&#160;Jan&#160;23&#160;18:15&#160;UTC | true | [view](CERTS/044daf68cdb455f9fa3bb1b40256824bc3915431742054b6dcd66ef7e4516119/README.md) |
| 31&#160;Dec&#160;22&#160;03:33&#160;UTC | SHAKEN 8526 | 31&#160;Dec&#160;23&#160;03:33&#160;UTC | true | [view](CERTS/96672cb8cc7c71e38d9e6350eac6c4f5c430c1eefeabe00cc095527e20a20c9b/README.md) |
| 31&#160;Dec&#160;22&#160;20:23&#160;UTC | SHAKEN 674J | 07&#160;Jan&#160;23&#160;20:23&#160;UTC | true | [view](CERTS/eaca40abca7317973cadd7598241bf1b5a6d8950c6cbd07d6603cae64d02b68c/README.md) |
| 31&#160;Dec&#160;22&#160;20:24&#160;UTC | SHAKEN 738J | 07&#160;Jan&#160;23&#160;20:24&#160;UTC | true | [view](CERTS/ae29dab996036cfab6ed19a1a2b32c19787cb15266b811d14285024f7a2a1e63/README.md) |
| 31&#160;Dec&#160;22&#160;20:25&#160;UTC | SHAKEN 700H | 07&#160;Jan&#160;23&#160;20:25&#160;UTC | true | [view](CERTS/28b1b8764e1c61afba3b2246aae56b3020f86fb32b4602f2f7299c10cb31934e/README.md) |
| 31&#160;Dec&#160;22&#160;20:27&#160;UTC | SHAKEN 819H | 07&#160;Jan&#160;23&#160;20:27&#160;UTC | true | [view](CERTS/e736e30ea0bee7c8b47fb7559275b98133a4d8fdedb13531c375590078c78529/README.md) |
| 31&#160;Dec&#160;22&#160;20:30&#160;UTC | SHAKEN 0127 | 07&#160;Jan&#160;23&#160;20:30&#160;UTC | true | [view](CERTS/c0db2f1f7ce890d073ec5f92160ceef88fcf5a0fe4bb7065115a443a1f14f674/README.md) |
| 31&#160;Dec&#160;22&#160;20:30&#160;UTC | SHAKEN 849J | 07&#160;Jan&#160;23&#160;20:30&#160;UTC | true | [view](CERTS/b93e37153b859330238aa5b2cf14d8f5edb0faa6adc2e603920b41c0def0d80f/README.md) |
| 31&#160;Dec&#160;22&#160;20:30&#160;UTC | SHAKEN 495J | 07&#160;Jan&#160;23&#160;20:30&#160;UTC | true | [view](CERTS/b188b3413f3120529bd044a4bf515e4b3c6f089447bc22db2e5f8aa1e92fd014/README.md) |
| 31&#160;Dec&#160;22&#160;20:30&#160;UTC | SHAKEN 790J | 07&#160;Jan&#160;23&#160;20:30&#160;UTC | true | [view](CERTS/2f871c63e3f4d61821996918501c1af531526bdd0c4da793ad6a47760bd263b2/README.md) |
| 31&#160;Dec&#160;22&#160;20:31&#160;UTC | SHAKEN 738J | 07&#160;Jan&#160;23&#160;20:31&#160;UTC | true | [view](CERTS/28359791d8f2c4ad1beba0b2166c3b1abf4fc12a8bfe3f2bd9915971ab2406de/README.md) |
| 31&#160;Dec&#160;22&#160;20:32&#160;UTC | SHAKEN 366G | 07&#160;Jan&#160;23&#160;20:32&#160;UTC | true | [view](CERTS/0b45da5c4f8749ec665fb7ab1e87f19f029adbb4b52bf9688d81c8e460e8862b/README.md) |
| 31&#160;Dec&#160;22&#160;20:32&#160;UTC | SHAKEN 0226 | 07&#160;Jan&#160;23&#160;20:32&#160;UTC | true | [view](CERTS/1985c1d2cec4349c5fbdedaf5c79eb81184237ed5fbffdf63af5b5e8dbc023eb/README.md) |
| 31&#160;Dec&#160;22&#160;20:32&#160;UTC | SHAKEN 738J | 07&#160;Jan&#160;23&#160;20:32&#160;UTC | true | [view](CERTS/337eb08320ccc54b4e3347c5def21de091a20bad9dc37ba01d1fb2b04a58db48/README.md) |
| 01&#160;Jan&#160;23&#160;16:38&#160;UTC | SHAKEN 291K | 08&#160;Jan&#160;23&#160;16:38&#160;UTC | true | [view](CERTS/1727c20509b53406ceb32e28969840769bf2b8db386d2203560ad6703eb44a87/README.md) |
| 01&#160;Jan&#160;23&#160;18:08&#160;UTC | SHAKEN 406K | 08&#160;Jan&#160;23&#160;18:08&#160;UTC | true | [view](CERTS/2f264f8a36373dddc77af521a67d563737a19e08abf1f7ee73cc6b167264d31e/README.md) |
| 01&#160;Jan&#160;23&#160;20:20&#160;UTC | SHAKEN 983J | 08&#160;Jan&#160;23&#160;20:20&#160;UTC | true | [view](CERTS/b3c1644791338af25c2e950ba9b8738905e8da2998f38a11c8c528c866e35fc4/README.md) |
| 01&#160;Jan&#160;23&#160;21:47&#160;UTC | SHAKEN 606F | 08&#160;Jan&#160;23&#160;21:47&#160;UTC | true | [view](CERTS/bdc3d96df2da242494c2fd6f14a7ae7514596c433484740b36512eb34a96c100/README.md) |
| 02&#160;Jan&#160;23&#160;01:45&#160;UTC | SHAKEN 278K | 09&#160;Jan&#160;23&#160;01:45&#160;UTC | true | [view](CERTS/de9136817d57dbd018f960fe85769f2196e21525c04e09b20a50a47d46c6473e/README.md) |
| 02&#160;Jan&#160;23&#160;14:39&#160;UTC | SHAKEN 2550 | 09&#160;Jan&#160;23&#160;14:39&#160;UTC | true | [view](CERTS/86c2443ac69d8674998150273e9590eb3def129436b38db64d4dc56062353879/README.md) |
| 02&#160;Jan&#160;23&#160;16:21&#160;UTC | SHAKEN 753J | 09&#160;Jan&#160;23&#160;16:21&#160;UTC | true | [view](CERTS/62bdcbc00e3f1544012f3d85995db0d6335968dfcc16e06969313b4485246ed9/README.md) |
| 02&#160;Jan&#160;23&#160;18:36&#160;UTC | SHAKEN 242K | 09&#160;Jan&#160;23&#160;18:36&#160;UTC | true | [view](CERTS/74d809075217f56d0ef8d07afea90f294910581e1b06c8ec22ddd2feae32c0c2/README.md) |
| 02&#160;Jan&#160;23&#160;20:11&#160;UTC | SHAKEN 864J | 09&#160;Jan&#160;23&#160;20:11&#160;UTC | true | [view](CERTS/6e70aedd1bd0de6e5c5c58edaebb37ff500c412d8aff6b93b97dff6fd134d1b7/README.md) |
| 02&#160;Jan&#160;23&#160;20:17&#160;UTC | SHAKEN 177K | 09&#160;Jan&#160;23&#160;20:17&#160;UTC | true | [view](CERTS/723c42e429585cbf1fcd5e457b5f96fe6652d38e9b6d58603f574d4e0c877b25/README.md) |
| 02&#160;Jan&#160;23&#160;20:31&#160;UTC | SHAKEN 738J | 11&#160;Jan&#160;23&#160;20:31&#160;UTC | true | [view](CERTS/153c0049d2b46f904919fd1c9648a3f817e8a98d735b432f5adaa356bcfd6a14/README.md) |
| 03&#160;Jan&#160;23&#160;09:14&#160;UTC | SHAKEN 0172 | 10&#160;Jan&#160;23&#160;09:14&#160;UTC | true | [view](CERTS/8885bc04807ab3a1d3bfbe9d908088c88bf177dcd10975f8bb4217cca344a51d/README.md) |
| 03&#160;Jan&#160;23&#160;14:05&#160;UTC | SHAKEN 551G | 10&#160;Jan&#160;23&#160;14:05&#160;UTC | true | [view](CERTS/a05ef6a6c4f4f7ce85e17a35458440742f34b4eb8dd26ba28c56f02afb10dbd7/README.md) |
| 03&#160;Jan&#160;23&#160;14:46&#160;UTC | SHAKEN 012K | 10&#160;Jan&#160;23&#160;14:46&#160;UTC | true | [view](CERTS/423639fe3fbb386998640c82ea272ef06c563a301b87552a2024ac7f2fe999f2/README.md) |
| 03&#160;Jan&#160;23&#160;14:47&#160;UTC | SHAKEN 186K | 10&#160;Jan&#160;23&#160;14:47&#160;UTC | true | [view](CERTS/15c556d686447cefcebba11ae9fae7ae12b45535edb2d25352481fde237257a9/README.md) |
| 03&#160;Jan&#160;23&#160;14:51&#160;UTC | SHAKEN 982J | 10&#160;Jan&#160;23&#160;14:51&#160;UTC | true | [view](CERTS/ba0caa7d0d6b614d125cfb1686dd0530c879fc2527012519b416bcdf58889360/README.md) |
| 03&#160;Jan&#160;23&#160;17:29&#160;UTC | SHAKEN 107K | 10&#160;Jan&#160;23&#160;17:29&#160;UTC | true | [view](CERTS/ba0a490099ad9edc6333bb87c5631c1467babddddfe89d6f4cfb56e5ba18ee6d/README.md) |
| 03&#160;Jan&#160;23&#160;18:42&#160;UTC | SHAKEN 089K | 10&#160;Jan&#160;23&#160;18:42&#160;UTC | true | [view](CERTS/7fb9157697ec0ea4423cac70ec040e99feb85f33144e23e1b21ff52a33452854/README.md) |
| 03&#160;Jan&#160;23&#160;18:44&#160;UTC | SHAKEN 9714 | 10&#160;Jan&#160;23&#160;18:44&#160;UTC | true | [view](CERTS/4ff0251809cd75d4eeb230123fce799df07f72bfacff211f7555b627c1d80b89/README.md) |
| 03&#160;Jan&#160;23&#160;20:08&#160;UTC | SHAKEN 366G | 10&#160;Jan&#160;23&#160;20:08&#160;UTC | true | [view](CERTS/58a52365dad03fe4f6f8ac11118978d2f68a3570b90e282beaebe6bb575c19bc/README.md) |
| 03&#160;Jan&#160;23&#160;20:23&#160;UTC | SHAKEN 674J | 10&#160;Jan&#160;23&#160;20:23&#160;UTC | true | [view](CERTS/8079d5006dcddbe5a8e6526868abc5f60e89970fa5ce91952ac4ab2e3cbaa8e0/README.md) |
| 03&#160;Jan&#160;23&#160;20:24&#160;UTC | SHAKEN 738J | 10&#160;Jan&#160;23&#160;20:24&#160;UTC | true | [view](CERTS/87b4c077d8a37ab104e44cf3a2f7835449d709c3c4472074256d9d1c44e57d4b/README.md) |
| 03&#160;Jan&#160;23&#160;20:25&#160;UTC | SHAKEN 700H | 10&#160;Jan&#160;23&#160;20:25&#160;UTC | true | [view](CERTS/f81e5ec4e420951d51a6f39b2547ff1f3551c190e08c70265064a50b406ce552/README.md) |
| 03&#160;Jan&#160;23&#160;20:25&#160;UTC | SHAKEN 7453 | 10&#160;Jan&#160;23&#160;20:25&#160;UTC | true | [view](CERTS/b67f4cc84627b8a718d60ddb94258bc3569b8727da6a2608e835c7983e3f7ab0/README.md) |
| 03&#160;Jan&#160;23&#160;20:27&#160;UTC | SHAKEN 819H | 10&#160;Jan&#160;23&#160;20:27&#160;UTC | true | [view](CERTS/6cc88b752b76eecad7b63df8c420e2563fe125c63f312d81cee263baacb06cc7/README.md) |
| 03&#160;Jan&#160;23&#160;20:28&#160;UTC | SHAKEN 769J | 10&#160;Jan&#160;23&#160;20:28&#160;UTC | true | [view](CERTS/42c5a9d5554f5656e2b9b33348531d19580ee7f8c804fd0e76e5dc1897e1ff42/README.md) |
| 03&#160;Jan&#160;23&#160;20:30&#160;UTC | SHAKEN 691A | 10&#160;Jan&#160;23&#160;20:30&#160;UTC | true | [view](CERTS/32acfdf6737045ca759008827dd916472a6030f06005eeb5d3e31a5984c2551e/README.md) |
| 03&#160;Jan&#160;23&#160;20:30&#160;UTC | SHAKEN 849J | 10&#160;Jan&#160;23&#160;20:30&#160;UTC | true | [view](CERTS/f47314b281eb05cb513018e463e549c77bf9dde05cc81ce53c641a2df34eafc5/README.md) |
| 03&#160;Jan&#160;23&#160;20:30&#160;UTC | SHAKEN 469A | 10&#160;Jan&#160;23&#160;20:30&#160;UTC | true | [view](CERTS/ecb85eb5afb38beb5c4b1db731eb9067d4769458c9f456fa9528a71f5863df7e/README.md) |
| 03&#160;Jan&#160;23&#160;20:30&#160;UTC | SHAKEN 726J | 10&#160;Jan&#160;23&#160;20:30&#160;UTC | true | [view](CERTS/f4ce9a0f68a1a9bb129893474433fe5e4fc857b8b1f6e807ed26ef3c7bb6955f/README.md) |
| 03&#160;Jan&#160;23&#160;20:30&#160;UTC | SHAKEN 495J | 10&#160;Jan&#160;23&#160;20:30&#160;UTC | true | [view](CERTS/c54baba352a8ac71c237aecc3c904801e8f446c84536fb492ccc6cfe34bbcd7c/README.md) |
| 03&#160;Jan&#160;23&#160;20:30&#160;UTC | SHAKEN 790J | 10&#160;Jan&#160;23&#160;20:30&#160;UTC | true | [view](CERTS/ced8bf34ef851d0512559a4ac5367868f16cee031cd3e1de9aaeb0bc3304db94/README.md) |
| 03&#160;Jan&#160;23&#160;20:31&#160;UTC | SHAKEN 738J | 10&#160;Jan&#160;23&#160;20:31&#160;UTC | true | [view](CERTS/77e8796d3fb3fe265e0e711c3cbfedf66f6f076b0b5135af6e94f1ead3fb9438/README.md) |
| 03&#160;Jan&#160;23&#160;20:31&#160;UTC | SHAKEN 459J | 10&#160;Jan&#160;23&#160;20:31&#160;UTC | true | [view](CERTS/cfad08ffb7b9204c8fcef597d58eead2e30c07d094af14573a4c142b1759b3e9/README.md) |
| 03&#160;Jan&#160;23&#160;20:31&#160;UTC | SHAKEN 672B | 10&#160;Jan&#160;23&#160;20:31&#160;UTC | true | [view](CERTS/0752ce8c51da5d965d73b15a4a63a783e14f613801b5c762b501ed2c24ecc587/README.md) |
| 03&#160;Jan&#160;23&#160;20:32&#160;UTC | SHAKEN 366G | 10&#160;Jan&#160;23&#160;20:32&#160;UTC | true | [view](CERTS/d79a5fc468e95f180591d369f50806aba2da07cba24b7e79ea3b0ebb0cdc9fe0/README.md) |
| 03&#160;Jan&#160;23&#160;20:32&#160;UTC | SHAKEN 0226 | 10&#160;Jan&#160;23&#160;20:32&#160;UTC | true | [view](CERTS/36c7742eaf7b3483636468244fdce7f0ed729f1b72f9db2068f3a697b250a4a6/README.md) |
| 03&#160;Jan&#160;23&#160;20:32&#160;UTC | SHAKEN 738J | 10&#160;Jan&#160;23&#160;20:32&#160;UTC | true | [view](CERTS/54b4a67c136e8ea02ce2780a06ec840f89e9a8c91cd9f57b6e2aaa2b7d80b2d9/README.md) |
| 04&#160;Jan&#160;23&#160;00:13&#160;UTC | SHAKEN 841J | 18&#160;Jan&#160;23&#160;00:13&#160;UTC | true | [view](CERTS/404d2f2770d0c3eebdd8196d210b40fed9b2ab8165b097a71d4f8a59a76f4371/README.md) |
| 04&#160;Jan&#160;23&#160;16:39&#160;UTC | SHAKEN 291K | 11&#160;Jan&#160;23&#160;16:39&#160;UTC | true | [view](CERTS/025c190d6fde6ddf4492d6ea24dad320614fa864477c69c0eab46acd495f2ed1/README.md) |
| 04&#160;Jan&#160;23&#160;20:21&#160;UTC | SHAKEN 983J | 11&#160;Jan&#160;23&#160;20:21&#160;UTC | true | [view](CERTS/65f9c3bd2e32c514674b265137f7e5375b59dccb0139fa7ab74d416d0fa31c60/README.md) |
| 04&#160;Jan&#160;23&#160;21:48&#160;UTC | SHAKEN 606F | 11&#160;Jan&#160;23&#160;21:48&#160;UTC | true | [view](CERTS/da6fddb87fe7282a51f743e97cd631d8b49037a13e73ba19fa358c512d771074/README.md) |
| 05&#160;Jan&#160;23&#160;01:45&#160;UTC | SHAKEN 278K | 12&#160;Jan&#160;23&#160;01:45&#160;UTC | true | [view](CERTS/548e7087ebad0fdece5ac3b8068ec9d7de0cb41fbc8ff94a092a3db350d1eece/README.md) |
| 05&#160;Jan&#160;23&#160;14:39&#160;UTC | SHAKEN 2550 | 12&#160;Jan&#160;23&#160;14:39&#160;UTC | true | [view](CERTS/ff0dfd9578ca50491fbcd45e1609c7a4c7a3cd1de9e8410a779b1c04f525415c/README.md) |
| 05&#160;Jan&#160;23&#160;15:50&#160;UTC | SHAKEN 722J | 04&#160;Feb&#160;23&#160;15:50&#160;UTC | true | [view](CERTS/6707ea8caeed61aaa9c0489a5b1434622a9c938978f1bea4d2ea7fc5dbd81586/README.md) |
| 05&#160;Jan&#160;23&#160;16:02&#160;UTC | SHAKEN 311H | 12&#160;Jan&#160;23&#160;16:02&#160;UTC | true | [view](CERTS/c6cb968670633b1daa2a4dff8b3511647c12e9a13afeccc87a3496942320eb52/README.md) |
| 05&#160;Jan&#160;23&#160;16:21&#160;UTC | SHAKEN 753J | 12&#160;Jan&#160;23&#160;16:21&#160;UTC | true | [view](CERTS/6fc10f4cf376d908ca0fe901272b3d2dcbf2125ff88c86663e8ec3f54a9ef40a/README.md) |
| 05&#160;Jan&#160;23&#160;20:18&#160;UTC | SHAKEN 177K | 12&#160;Jan&#160;23&#160;20:18&#160;UTC | true | [view](CERTS/5f57d0a3ab4a74ce18851a50276f6331fdf6cea00dd0e262d37ba06600a901d1/README.md) |
| 06&#160;Jan&#160;23&#160;14:47&#160;UTC | SHAKEN 186K | 13&#160;Jan&#160;23&#160;14:47&#160;UTC | true | [view](CERTS/0401494275b7a3ba436d8669e461c131e87ad6a488d26efe9428261851cc04d4/README.md) |

#### CA Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 20&#160;Aug&#160;21&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Issuing CA3 | 19&#160;Aug&#160;31&#160;23:59&#160;UTC | true | [view](CERTS/51b51431ffd8055cd7e30c213540f4868258058d4c4c9f23f1ed8bea75101e37/README.md) |
| 20&#160;Aug&#160;21&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Issuing CA1 | 19&#160;Aug&#160;31&#160;23:59&#160;UTC | true | [view](CERTS/037a7c6621814bd5f01fa2080a77a245458694d9d566e0050e068fc853317857/README.md) |
| 20&#160;Aug&#160;21&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Root CA1 | 19&#160;Aug&#160;41&#160;23:59&#160;UTC | true | [view](CERTS/59ab943f4e8c3c17755cc8e4abb5ac65736ce74ec587c62cb744ab0babbbe2fe/README.md) |
| 24&#160;Oct&#160;22&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Issuing CA4 | 23&#160;Oct&#160;32&#160;23:59&#160;UTC | false | [view](CERTS/80dd12d935f46c256c7c1289a0834b21ec8f2f8c35f2864928d1e75f3a280665/README.md) |
| 24&#160;Oct&#160;22&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Root CA2 | 23&#160;Oct&#160;42&#160;23:59&#160;UTC | false | [view](CERTS/a26e04fc786ab70b8085236b2c53f8cfbf5d0c6a5c2c9c3e9f91669fbb8ea4d5/README.md) |


Generated: 07 Jan 23 19:18 UTC