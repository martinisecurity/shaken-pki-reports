# STIR/SHAKEN CA Ecosystem Compliance

## TransNexus

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

52 certificates were included in the corpus being tested\
264 certificates in the corpus were skipped because they were expired\
1 certificates in the corpus were skipped because they are not currently trusted\
100.00% of certificates contain one or more Error level issue\
100.00% of certificates contain one or more Warning level issue\
0.00% of certificates contain one or more Notice level issue\
3.85% of certificates are too old to be assessed against currently enforced expectations\
94 days is the average remaining validity for the certificates in the corpus\
96 days is the average initial validity for the certificates in the corpus\
37 certificates expire in the next 30 days

| Instances | Test Status | Source |
|-----------|-------------|--------|
| 1 | [e_cp1_3_ambiguous_identifier](ISSUES/e_cp1_3_ambiguous_identifier/README.md) | United States SHAKEN CP |
| 1 | [e_cp1_3_subject_sn](ISSUES/e_cp1_3_subject_sn/README.md) | United States SHAKEN CP |
| 1 | [e_incorrect_ku_encoding](ISSUES/e_incorrect_ku_encoding/README.md) | RFC5280 |
| 1 | [e_sti_certificate_policies](ISSUES/e_sti_certificate_policies/README.md) | ATIS-1000080 |
| 1 | [e_sti_extension_unknown](ISSUES/e_sti_extension_unknown/README.md) | ATIS-1000080 |
| 1 | [e_sti_serial_number](ISSUES/e_sti_serial_number/README.md) | ATIS-1000080 |
| 1 | [e_sti_signature_algorithm](ISSUES/e_sti_signature_algorithm/README.md) | ATIS-1000080 |
| 1 | [e_sti_subject_cn](ISSUES/e_sti_subject_cn/README.md) | ATIS-1000080 |
| 1 | [w_cp_1_3_subject_email](ISSUES/w_cp_1_3_subject_email/README.md) | United States SHAKEN CP |
| 1 | [w_pki_subject_rdn_unknown](ISSUES/w_pki_subject_rdn_unknown/README.md) | SHAKEN PKI Best Practice |

#### CA Certificates

3 certificates were included in the corpus being tested\
0 certificates in the corpus were skipped because they were expired\
1 certificates in the corpus were skipped because they are not currently trusted\
100.00% of certificates contain one or more Error level issue\
100.00% of certificates contain one or more Warning level issue\
100.00% of certificates contain one or more Notice level issue\
100.00% of certificates are too old to be assessed against currently enforced expectations\
5369 days is the average remaining validity for the certificates in the corpus\
4870 days is the average initial validity for the certificates in the corpus\
0 certificates expire in the next 30 days

| Instances | Test Status | Source |
|-----------|-------------|--------|
| 1 | [e_cp1_3_ca_key_usage_crl_sign](ISSUES/e_cp1_3_ca_key_usage_crl_sign/README.md) | United States SHAKEN CP |
| 1 | [e_incorrect_ku_encoding](ISSUES/e_incorrect_ku_encoding/README.md) | RFC5280 |
| 1 | [e_sti_ca_certificate_policies](ISSUES/e_sti_ca_certificate_policies/README.md) | ATIS-1000080 |
| 1 | [e_sti_ca_crl_distribution](ISSUES/e_sti_ca_crl_distribution/README.md) | ATIS-1000080 |
| 1 | [e_sti_ca_extension_unknown](ISSUES/e_sti_ca_extension_unknown/README.md) | ATIS-1000080 |
| 1 | [e_sti_ca_issuer_dn](ISSUES/e_sti_ca_issuer_dn/README.md) | ATIS-1000080 |
| 1 | [e_sti_ca_serial_number](ISSUES/e_sti_ca_serial_number/README.md) | ATIS-1000080 |
| 1 | [e_sti_ca_subject_cn](ISSUES/e_sti_ca_subject_cn/README.md) | ATIS-1000080 |
| 1 | [e_sti_root_certificate_policies](ISSUES/e_sti_root_certificate_policies/README.md) | ATIS-1000080 |
| 1 | [e_sti_root_extension_unknown](ISSUES/e_sti_root_extension_unknown/README.md) | ATIS-1000080 |
| 1 | [n_pki_ca_key_usage](ISSUES/n_pki_ca_key_usage/README.md) | SHAKEN PKI Best Practice |
| 1 | [n_sti_ca_certificate_policy_critical](ISSUES/n_sti_ca_certificate_policy_critical/README.md) | ATIS-1000080 |
| 1 | [w_pki_ca_subject_rdn_unknown](ISSUES/w_pki_ca_subject_rdn_unknown/README.md) | SHAKEN PKI Best Practice |

### Details

#### Leaf Certificates

| Created at | Subject | Problems | Link |
|------------|---------|----------|------|
| 14 Sep 22 08:36 UTC | SHAKEN 012K | true | [view](CERTS/788e828c8fe34b9e0116e8f5d3a7e6d71a22a83803056a63b5ab35f2ec52f6cf/README.md) |
| 17 Oct 22 08:38 UTC | SHAKEN 012K | true | [view](CERTS/c0e92d8eb48b361521d97d58db9fe08b38de0d4d3172e66f583c2436d51390f6/README.md) |
| 02 Oct 22 08:37 UTC | SHAKEN 012K | true | [view](CERTS/edb4f38534689a7fb5606964f8257852d660969f4c185ca83e4bb31524df8655/README.md) |
| 05 Oct 22 08:37 UTC | SHAKEN 012K | true | [view](CERTS/59b280001b41751aa11afd61094e56d63f5fe6967400368495d8cc2226b8ec0e/README.md) |
| 26 Oct 22 08:39 UTC | SHAKEN 012K | true | [view](CERTS/a4a8fa499de2764fbb7db0b76e98d31078179de61e21e96df01f8132a44c1fd4/README.md) |
| 14 Oct 22 08:38 UTC | SHAKEN 012K | true | [view](CERTS/c3f32dd3cafcf1371c0f264602dcee9af6f0c66acbd2bfd040da981b5c2e9c62/README.md) |
| 11 Oct 22 08:37 UTC | SHAKEN 012K | true | [view](CERTS/93fea95c13f727aa429aeead995801108e51c07c8652300509835435b82f0a79/README.md) |
| 23 Oct 22 08:39 UTC | SHAKEN 012K | true | [view](CERTS/c8405fe8cda1c783a7b545455d1dfac11dddb4e010ce1e5404f01778df063bed/README.md) |
| 20 Sep 22 08:36 UTC | SHAKEN 012K | true | [view](CERTS/47c17df5f540211b6fe4ac43904b7f9047397642b2ed124d0ef9c00c5556044c/README.md) |
| 02 Oct 22 20:15 UTC | SHAKEN 674J | true | [view](CERTS/e15e5fbd39b072e415b1b862ad24e57a3689bee2a27232cfb874b13098bd307a/README.md) |
| 08 Oct 22 20:16 UTC | SHAKEN 674J | true | [view](CERTS/67f6e357e2fe3a1ee2ffc4f4ba0bdeb9899ec1663d932215b5f0ef6e8fb06ada/README.md) |
| 11 Oct 22 20:16 UTC | SHAKEN 674J | true | [view](CERTS/42b05052f886b1211962f9034055ba2cc2d00d47715b9df67adc49891cd28840/README.md) |
| 26 Oct 22 20:17 UTC | SHAKEN 674J | true | [view](CERTS/6fd005ce139c83958c022da431a4a721569905d4b6b43de0375a1f7d4b98d275/README.md) |
| 29 Sep 22 20:15 UTC | SHAKEN 674J | true | [view](CERTS/8092fc343024305acd92236303bfa2a4976aaac39d1c3a8ad436927b49c32e91/README.md) |
| 23 Oct 22 20:17 UTC | SHAKEN 674J | true | [view](CERTS/ee4b4b4e7ec360d46bd00e78fe553bb975fdbec3696f3f3f97bb97506cf38117/README.md) |
| 20 Oct 22 20:17 UTC | SHAKEN 674J | true | [view](CERTS/fbb5652ae7637537a167e82fdb83cfd8e066479ab9dd7d118ff118aef04dd285/README.md) |
| 17 Oct 22 20:16 UTC | SHAKEN 674J | true | [view](CERTS/5fb494c9fe8cda08f247a54d8fe102e27e4f7c2267f135ce27727b65c8698686/README.md) |
| 02 Oct 22 18:11 UTC | SHAKEN 060K | true | [view](CERTS/7301d2047fa36d6afda04f5ada6442629cbb33883b3a17549cfccfb018bdb2bc/README.md) |
| 17 Oct 22 18:13 UTC | SHAKEN 060K | true | [view](CERTS/dc001c8050fc07cc724fd8edd73ed470b6cb33bca8a83b25e9f0c2d64b7497d7/README.md) |
| 20 Oct 22 18:13 UTC | SHAKEN 060K | true | [view](CERTS/9e357687c4cad81b36a0b692a6470b9c9c3845b1237d717e231bf4658a898740/README.md) |
| 08 Oct 22 18:12 UTC | SHAKEN 060K | true | [view](CERTS/5866f940f9f8cbe54b789232c6d82c0e36ef4fc664b3eddd0cd55676c86fdac3/README.md) |
| 26 Oct 22 18:13 UTC | SHAKEN 060K | true | [view](CERTS/8fc41fc2e5615fc87a1a9c1fa8cb6ea1c9200d551f5e988047221c335be31d39/README.md) |
| 05 Oct 22 18:12 UTC | SHAKEN 060K | true | [view](CERTS/55c4ade7d7a08ee468b5d53c7eb80a19e9a7bbd0c90ef8f4f4d018e72a10b69e/README.md) |
| 14 Sep 22 18:11 UTC | SHAKEN 060K | true | [view](CERTS/26f6d8b695c16bb2a8a5b9706e0f4841fae6de0b440ce3b92c07cd7f271f3e9e/README.md) |
| 16 Oct 22 16:01 UTC | SHAKEN 2277 | true | [view](CERTS/9dafd3e83a67a226833b5fdf61a432cf9a272294367db1bcdcca9e7c7c3d1634/README.md) |
| 26 Oct 22 20:18 UTC | SHAKEN 735J | true | [view](CERTS/9316e334abe6c964bac27767052062b6c0de27827a606d1cd746e95f97472411/README.md) |
| 02 Oct 22 20:16 UTC | SHAKEN 735J | true | [view](CERTS/24c98b04f18f554997c8a43137d7368273ec49176c495588dd03207c1d9189ed/README.md) |
| 17 Oct 22 20:17 UTC | SHAKEN 735J | true | [view](CERTS/99190988a0ce7d6dc5e470f0042334bb30718ec1907d6220389bdc5328428245/README.md) |
| 17 Sep 22 20:16 UTC | SHAKEN 594J | true | [view](CERTS/0a0426a877a419ac9775ee13cbf70b9139b4d52a579e53a51686b29ff7f6b0e5/README.md) |
| 17 Oct 22 20:18 UTC | SHAKEN 5512 | true | [view](CERTS/84e84f40193dc5f0ef9aa0c197047b439793e914cd74d03700069a3c277e9b96/README.md) |
| 11 Sep 22 20:15 UTC | SHAKEN 5512 | true | [view](CERTS/e2560f93f0e7a267b06f22127e85d2b6e67a6fbcdc42b74d9341f2fbc5783b36/README.md) |
| 02 Oct 22 20:17 UTC | SHAKEN 738J | true | [view](CERTS/bcbb62609b4bc2aff8765f8d2c4f8580302603f6c8e5ee85ce9f53d419c32c8a/README.md) |
| 17 Oct 22 20:18 UTC | SHAKEN 738J | true | [view](CERTS/29f87b482d88eeecebb347a63a3c504334a4d6dd40b9754a6351243a0a2801f0/README.md) |
| 14 Sep 22 20:15 UTC | SHAKEN 738J | true | [view](CERTS/535cdb8dbf8c15fa705b3a6f3c0ead06a8dd5bcf5dcd2aacaff10ff8822cbd0d/README.md) |
| 11 Oct 22 20:17 UTC | SHAKEN 738J | true | [view](CERTS/dba389b557235bf2ce2b52a80a4b3ffc5419b68efe015aa0a6c792600599bd9a/README.md) |
| 29 Sep 22 20:16 UTC | SHAKEN 738J | true | [view](CERTS/a4808634e396cd51c1d7f3d4c70e1a6a1acd399ae225a73b175b8ec9900ee3b8/README.md) |
| 17 Sep 22 20:16 UTC | SHAKEN 738J | true | [view](CERTS/fd63e66b5d02974e369e4ea028ab9fce22a5339df62bd014f97723c3e01cd450/README.md) |
| 26 Oct 22 20:19 UTC | SHAKEN 738J | true | [view](CERTS/3a7079d02db69f8133c4892893d6d94a5b893aca3578938a74ed360392488012/README.md) |
| 23 Oct 22 20:18 UTC | SHAKEN 738J | true | [view](CERTS/43f8f904370657717078245413fa389ddee998a1fe95c17e66709b5151cdb7cb/README.md) |
| 08 Oct 22 20:17 UTC | SHAKEN 738J | true | [view](CERTS/de320798056f6c358d084381ddf58fd89c64100111f03767bfcaa1565e474b3e/README.md) |
| 14 Oct 22 20:18 UTC | SHAKEN 738J | true | [view](CERTS/24baae986f85f095ed2938d81e205161d37641f23df0ecc88bdf9903fefcfff9/README.md) |
| 05 Oct 22 20:17 UTC | SHAKEN 738J | true | [view](CERTS/a28b5d42f525e61b44261d9e914e2f31cb6fac58ef16e9d89bb4218cbf6e1fe8/README.md) |
| 20 Oct 22 20:18 UTC | SHAKEN 738J | true | [view](CERTS/ba068759dc3a51f84bf65708045a3552ab7434d4d7a58898f9f4a146a3990aac/README.md) |
| 11 Oct 22 20:17 UTC | SHAKEN 700H | true | [view](CERTS/ae6c1eab209819f114366ad6782cfb97c167a7f6b2301c7cc3a0f7bfb741d7b0/README.md) |
| 17 Oct 22 20:18 UTC | SHAKEN 700H | true | [view](CERTS/c17cec9331aab14957ec1c2c3d695a603024c048302a1377e5a52b6e9452a1df/README.md) |
| 05 Oct 22 20:17 UTC | SHAKEN 700H | true | [view](CERTS/74e97d2a9389be89275d20f99f64c9fad75a05eb44dbd48f90a75a5588288a1b/README.md) |
| 14 Oct 22 20:18 UTC | SHAKEN 700H | true | [view](CERTS/060dfcd88d3000d3043aff87829193421bc67c4666b3a0f043d6872db64f0288/README.md) |
| 02 Oct 22 20:17 UTC | SHAKEN 700H | true | [view](CERTS/62ed719ab766567581e27553e150e13d5336bc3ec9ea4173aeb5e70e22387e0c/README.md) |
| 14 Sep 22 20:16 UTC | SHAKEN 700H | true | [view](CERTS/fd22e72305dbbc6e79321a288006476ce0da6b2e7d54c5d693644ff693faaf72/README.md) |
| 17 Sep 22 20:16 UTC | SHAKEN 700H | true | [view](CERTS/d943d6f461ae6f7c00ba704b0f161b2a11de8fd9ea6aca2a8d913de4e392c2fc/README.md) |
| 17 Oct 22 20:18 UTC | SHAKEN 7453 | true | [view](CERTS/04c139db79790115d1949cd697493e8a5eed8c543674e69999f638cc978c3736/README.md) |
| 11 Oct 22 20:17 UTC | SHAKEN 7453 | true | [view](CERTS/1129aac58c9239771ce04cc057a30cf6e4e94d7d071c8b8c3959e8dea662e301/README.md) |
| 23 Oct 22 20:18 UTC | SHAKEN 7453 | true | [view](CERTS/f17d5125d8152a034cf15fbdc5e317c8990574e675693cd370132bc6c6b1d8b9/README.md) |
| 14 Oct 22 20:18 UTC | SHAKEN 7453 | true | [view](CERTS/f83e556f48e27e468cd74c870e7d79581099ab80c7fdf109f8f6e410679fd195/README.md) |
| 11 Oct 22 18:44 UTC | SHAKEN 179K | true | [view](CERTS/3e88f96f97af777e1ef990f348ac4b833280dd6eace5791841f41c87358f84ea/README.md) |
| 14 Oct 22 18:44 UTC | SHAKEN 179K | true | [view](CERTS/095d8cfae4b85954a0aefde024b8e32555977229fe9d23a8ad135821c49c0eac/README.md) |
| 19 Oct 22 14:33 UTC | SHAKEN 2550 | true | [view](CERTS/bd4a772d64e7a1abbca23cb3ca834578f75ca825cf12878ad170b6afdffac93a/README.md) |
| 04 Oct 22 14:30 UTC | SHAKEN 2550 | true | [view](CERTS/ed51db207040a9e47a7bcc940f9f3ea44da001040041f2d19c76bc133824c93e/README.md) |
| 19 Sep 22 14:29 UTC | SHAKEN 2550 | true | [view](CERTS/cec634c7f40f586ebf11bc9281f618ccd0c7d2582d2121a3cf866b3b581d0578/README.md) |
| 13 Oct 22 14:32 UTC | SHAKEN 2550 | true | [view](CERTS/6b5bee6a9964d167e48f86ac15b63a333e1871117e9ba77ab25f7828158be76b/README.md) |
| 25 Oct 22 14:33 UTC | SHAKEN 2550 | true | [view](CERTS/6970b8cb5229451110fe03840a2927e3d4d4dc548b6d5b4658097a6d996e1ea0/README.md) |
| 16 Oct 22 14:32 UTC | SHAKEN 2550 | true | [view](CERTS/e82b1229e2ac28031edbb7fedc4850532feab0b4b232fe45f9267b4cb53336b6/README.md) |
| 10 Oct 22 14:31 UTC | SHAKEN 2550 | true | [view](CERTS/cdb9332538e51481ed770800a0f30536f31e0a3b5edf4d1fbfc43c854d66c6cd/README.md) |
| 01 Oct 22 14:30 UTC | SHAKEN 2550 | true | [view](CERTS/cac9bcb351fe3efe557f98e97a0504ae8ad8aeb0a3c8884906e9729529c79573/README.md) |
| 13 Sep 22 14:28 UTC | SHAKEN 2550 | true | [view](CERTS/06d48b57180c15660a3f2ad259d64f4203fca7cba1b3955bd893da3b2103ded4/README.md) |
| 13 Oct 22 01:40 UTC | SHAKEN 278K | true | [view](CERTS/5b3c33d0a0e1abcfb381d075cf2c5a75a92e1b85e8d26430739e6ed5a38dbf5a/README.md) |
| 04 Oct 22 01:39 UTC | SHAKEN 278K | true | [view](CERTS/9b034de24a597d0f990397b33007836114c62f77573a7e394a21a1ae955924e2/README.md) |
| 25 Oct 22 01:41 UTC | SHAKEN 278K | true | [view](CERTS/14d17618059d6271f3f150cb33546796bca31eb548ad5bf659230ae1bd638585/README.md) |
| 16 Sep 22 01:38 UTC | SHAKEN 278K | true | [view](CERTS/1c7cea382374d5b45ff4402a1042072b5f811745a0e4509fa216fce82d5f68e5/README.md) |
| 13 Sep 22 01:38 UTC | SHAKEN 278K | true | [view](CERTS/a98d50d558d5dc64b587c1593a13f293efae91ffc34bf34f06b17013f5fcf87e/README.md) |
| 19 Oct 22 01:41 UTC | SHAKEN 278K | true | [view](CERTS/70f95294c22a41fca33657c0d4e051dad716d4f7157b8f36d1ba953e3cd81327/README.md) |
| 02 Oct 22 13:42 UTC | SHAKEN 056J | true | [view](CERTS/5049cf16383007760bef96944055a9467cd62329eb26e14ada3d75b331d1379e/README.md) |
| 26 Oct 22 13:45 UTC | SHAKEN 056J | true | [view](CERTS/860b047bd23110db2919afd52c5e095799e564dfe075a3d67e77c5883a4885bf/README.md) |
| 15 Oct 22 19:44 UTC | SHAKEN 918J | true | [view](CERTS/1c7d3957d11cefd886bf0bf510b239182d7ff8442277496c73645cc2413d769f/README.md) |
| 15 Sep 22 19:41 UTC | SHAKEN 918J | true | [view](CERTS/5d08106dbed56c07125e3d69f03f694d0be2060bb799ddd3cb93cc8249ab7b80/README.md) |
| 29 Sep 22 20:17 UTC | SHAKEN 819H | true | [view](CERTS/b0c0da7636ee6f902f35660d78ad2bf3e0bb2e251e77399a0abd04bc83374589/README.md) |
| 08 Oct 22 20:18 UTC | SHAKEN 819H | true | [view](CERTS/482fe4037b08417f2096fe07d4ef45c866906eccb039e442283fa69bc5898eed/README.md) |
| 14 Oct 22 20:19 UTC | SHAKEN 819H | true | [view](CERTS/cbe216a5a18d73714abf854cb6eb46f81a3a6db00055d1dba7893bb6f56cca5a/README.md) |
| 14 Sep 22 20:16 UTC | SHAKEN 819H | true | [view](CERTS/44620f766a38e0bb8802de7f0c9b6193150ebf3f16ff4a29a0da27dfd272cf0a/README.md) |
| 02 Oct 22 20:18 UTC | SHAKEN 819H | true | [view](CERTS/16e6efe46b4e721200cea32aa2a5834aa8522f9e2f46e8fe40ddf782cd97dcd0/README.md) |
| 17 Oct 22 20:19 UTC | SHAKEN 819H | true | [view](CERTS/4b7d4a521ee6e0728bd8813ebce6c88601c52dfd2438d5c4cdf92b373b03bc9c/README.md) |
| 11 Oct 22 20:20 UTC | SHAKEN 769J | true | [view](CERTS/30bc2111f19202690014f1f3cdbd77456cf94020c9bfa5462d5708f3b919fadf/README.md) |
| 17 Sep 22 20:17 UTC | SHAKEN 769J | true | [view](CERTS/4c0f33fc08e385bce7e4d7cfa09d5f650af286256bd79ec3df9f15aa60b75c3d/README.md) |
| 05 Oct 22 20:19 UTC | SHAKEN 769J | true | [view](CERTS/9389c01e1b1c4507a7744cc8d90a2beeabd1bb5b262b1cf3be3d56cf7ae917e7/README.md) |
| 11 Oct 22 20:20 UTC | SHAKEN 916J | true | [view](CERTS/beb4e0a62abf56052f1fec23eac2e9a4c36f0eb426328b4110ee21660fd55de7/README.md) |
| 26 Oct 22 17:24 UTC | SHAKEN 107K | true | [view](CERTS/ab2574c5a2c3ecdc234e72fdc62635cc97f596dc21a512a916d39aa17349719c/README.md) |
| 08 Oct 22 17:22 UTC | SHAKEN 107K | true | [view](CERTS/dda31ece96b874b7a3c72d31df45e9b638601c9613560ab7c3a4f0080f447a61/README.md) |
| 17 Sep 22 17:20 UTC | SHAKEN 107K | true | [view](CERTS/b362455d163dc9c25f9ff360d6c5fa220edcfc01339d79b367d2a3d50cae34ec/README.md) |
| 17 Sep 22 21:03 UTC | SHAKEN 1821 | true | [view](CERTS/6f088f6e72c4a41d29f6ca37d74c6640931d34c461f47386a6945a15f7ba00d3/README.md) |
| 11 Oct 22 21:06 UTC | SHAKEN 1821 | true | [view](CERTS/ef932371738ca4cc0312b3ddce613574cec16188888710a25a3d6eba04476804/README.md) |
| 12 Oct 22 12:10 UTC | SHAKEN 130B | true | [view](CERTS/99fa5a828eb9a1144e8e25d0d4f11bf0e0451ac7a1e1ae83c6573be1e6269d3a/README.md) |
| 23 Oct 22 18:37 UTC | SHAKEN 9714 | true | [view](CERTS/363bca4ca21ef4777859a333e63d5dd07a9bd4e8e8e323b12c1f025ed4602961/README.md) |
| 17 Sep 22 18:36 UTC | SHAKEN 9714 | true | [view](CERTS/5ea00c2574467b5f51c6a0ee8b9916b888c1807c16c1d671fa300ee16e96f711/README.md) |
| 26 Oct 22 18:38 UTC | SHAKEN 9714 | true | [view](CERTS/3c8ced9819143833a95bace3c72bebb882cef867f17509dc4628e18dd74341b2/README.md) |
| 02 Oct 22 18:36 UTC | SHAKEN 9714 | true | [view](CERTS/fd214ff1198dab5eba195a77650ff9702a00edf810625e84a2441087a07217ee/README.md) |
| 14 Sep 22 18:35 UTC | SHAKEN 9714 | true | [view](CERTS/de65167a175420827d884ad492cb17ae2137b8621b304eb60f0f39a1a4e3fa7d/README.md) |
| 11 Oct 22 18:36 UTC | SHAKEN 9714 | true | [view](CERTS/1a27a6b8a17ca7306e08c98b74621cd17005e81e8311be56985183d7fbf391cb/README.md) |
| 21 Oct 22 20:13 UTC | SHAKEN 983J | true | [view](CERTS/2f3b194e29f8ccbe580a06c7913aab6246e3062e8707cc91005b0f5b766ff166/README.md) |
| 27 Sep 22 20:10 UTC | SHAKEN 983J | true | [view](CERTS/58b1e17c819b1af13f8a37031039ada9e7f2e864aea86392b338bfebf33f388c/README.md) |
| 09 Oct 22 20:12 UTC | SHAKEN 983J | true | [view](CERTS/8424fa4027a9348dee9b7877e136657772f72ec325f11a7347dba191925aa6df/README.md) |
| 12 Sep 22 20:07 UTC | SHAKEN 983J | true | [view](CERTS/3c255615a5e72e5ae0984f97e906204bc444d51b8ea1b03309c8f8e4ac9a9e96/README.md) |
| 27 Oct 22 20:14 UTC | SHAKEN 983J | true | [view](CERTS/61a162ef08cf3cf9690416ee7985d04e655f82a4e2c6bd0760c006ba76794a78/README.md) |
| 03 Oct 22 20:11 UTC | SHAKEN 983J | true | [view](CERTS/d93c909ed7c013de70ecfdfa2ed0d1ef8583d8edabdf0e35cde8e6d2166d5480/README.md) |
| 18 Oct 22 20:13 UTC | SHAKEN 983J | true | [view](CERTS/c9e676271dfc725bb79881ed2d16a954ac16ede82c39344f35757f1ba7e77600/README.md) |
| 30 Sep 22 20:10 UTC | SHAKEN 983J | true | [view](CERTS/0e5eb850585cbf7d5b764650c7c94c0dc7d9f6d68b5fe812caa48a76d31dbdca/README.md) |
| 15 Oct 22 20:12 UTC | SHAKEN 983J | true | [view](CERTS/d51e59633cc904429f836f57fa629a6c74e6120156278f83d8671c2be21f729c/README.md) |
| 12 Oct 22 20:12 UTC | SHAKEN 983J | true | [view](CERTS/ded08a2e2fce9b57468f6b4d105dfd1f985a5cbabcba4aa5123e6073cb7528b8/README.md) |
| 18 Sep 22 20:09 UTC | SHAKEN 983J | true | [view](CERTS/ce960401ab108f8063ef5be5e622a7df30a7a1dc150ca7e0c6edda12555e80ad/README.md) |
| 24 Oct 22 20:14 UTC | SHAKEN 983J | true | [view](CERTS/705070c289a3188817e36d9f08ace1363c875e63b5f3355c846758e80a20fea6/README.md) |
| 02 Oct 22 19:59 UTC | SHAKEN 366G | true | [view](CERTS/aabb9e35010ea117db3b79fa34d6b458c469d16c89afd4ffdea70471cd106d31/README.md) |
| 05 Oct 22 19:59 UTC | SHAKEN 366G | true | [view](CERTS/4683361ca58a8e8b91abe092b0674f8de43625c51739933968b094b4fe6f0acf/README.md) |
| 08 Oct 22 19:59 UTC | SHAKEN 366G | true | [view](CERTS/d5c958bf1e11a705b51c376caf54211f85462f835535c77a3210993e0a1c9c12/README.md) |
| 11 Oct 22 19:59 UTC | SHAKEN 366G | true | [view](CERTS/f613fa4dfa2c4af1faa30fc70a62d445d2f2a7b70733bbef78582a7d0aaae0fb/README.md) |
| 17 Sep 22 19:57 UTC | SHAKEN 366G | true | [view](CERTS/2f68a1eb92fe90bb9dffb285e9e8cdbc09f304821bdb8ef9ed8b1014e4e6266a/README.md) |
| 26 Oct 22 18:36 UTC | SHAKEN 089K | true | [view](CERTS/3e4203b5c20a0a59801040c172eed8f39d2a609fd2429e4a111f40bd0d51c7c0/README.md) |
| 05 Oct 22 18:34 UTC | SHAKEN 089K | true | [view](CERTS/f15aadf1e23bff190c1f21210b89334f083e4553ade86273ab8b42403edcab68/README.md) |
| 11 Oct 22 18:34 UTC | SHAKEN 089K | true | [view](CERTS/4b68083ab81c9861d4150008ce6e7339c7fcba7bedb123862342f7f30d6c188f/README.md) |
| 23 Oct 22 18:35 UTC | SHAKEN 089K | true | [view](CERTS/00c8d016170b36ffe7dfb881ca62a497cfc9118984c76cfaa526be9695dccf4f/README.md) |
| 17 Oct 22 18:35 UTC | SHAKEN 089K | true | [view](CERTS/375815b2fe8302e528bba8f569930f6e63fc165aa9a2e7996fb69b41bdacc43b/README.md) |
| 02 Oct 22 18:34 UTC | SHAKEN 089K | true | [view](CERTS/9bc106e00df1596895d02e70727205ee761bedc6f40c9d3aee31a841aca7aba4/README.md) |
| 20 Oct 22 20:21 UTC | SHAKEN 854D | true | [view](CERTS/a68d25fdd3a24fd8bac42c9d392a3b4bd28791ed8ab0a66a192a7745cd758aae/README.md) |
| 24 Oct 22 21:41 UTC | SHAKEN 606F | true | [view](CERTS/688235ecbd4d2a93153e566a38033fc10a9a3518b238b12462c794e671cef0ef/README.md) |
| 09 Oct 22 21:39 UTC | SHAKEN 606F | true | [view](CERTS/2dbf24de2689b8941da087e1dbc2ccaf924a347b278f3cf2e1f4416e71fa65af/README.md) |
| 09 Oct 22 20:50 UTC | SHAKEN 157C | true | [view](CERTS/93cfebea96065456516cc415ee7714185e135eff7d1a0d0ad9f40985233029ec/README.md) |
| 24 Oct 22 20:51 UTC | SHAKEN 157C | true | [view](CERTS/be04d1f3ac988f74ec890dff28073678aed38da0c6971bcb2ef64479cb01c5b3/README.md) |
| 08 Oct 22 20:20 UTC | SHAKEN 691J | true | [view](CERTS/2eec45ae67c91ed8776a4af225bc69efbae9304a191ea46e4b030efcc367f3ea/README.md) |
| 23 Oct 22 19:56 UTC | SHAKEN 297K | true | [view](CERTS/e2e6f6d55f4f1817ebba820cb4366c2bad62eae895f85287391e4a1614907040/README.md) |
| 05 Oct 22 19:54 UTC | SHAKEN 297K | true | [view](CERTS/d22b15aaf061f3c0f2dd05a87c2957e8cc74ce0ea821ff0282963f7bfc591f77/README.md) |
| 17 Sep 22 19:52 UTC | SHAKEN 297K | true | [view](CERTS/ce5ea60a4642e84597a2f1a4590f17ff23a5c0ca39765844f6c01be934d6d891/README.md) |
| 23 Oct 22 20:21 UTC | SHAKEN 849J | true | [view](CERTS/21cb74ec5134456e2be377bf1fbb781e2402b510cbae2ab7e329e39c780e7067/README.md) |
| 14 Sep 22 20:17 UTC | SHAKEN 849J | true | [view](CERTS/d8ab9db60c1814cfd5418912465407914d8fffde70d176970e9d1cb51cae1618/README.md) |
| 17 Sep 22 20:17 UTC | SHAKEN 849J | true | [view](CERTS/89d707442f1ef8852658fbc1ec9fc0859afafbb1122a08d74f49fe18e4ac234e/README.md) |
| 14 Oct 22 20:20 UTC | SHAKEN 849J | true | [view](CERTS/df8fe34acfa40db9bb818150464ee61642e09cb4f2aadcd733f072a0a609eb31/README.md) |
| 11 Oct 22 20:20 UTC | SHAKEN 849J | true | [view](CERTS/c4721cfebe4a0c4e4420a229fcd03a1541e6cb3ab71330d4443376715c61ee1c/README.md) |
| 08 Oct 22 20:20 UTC | SHAKEN 849J | true | [view](CERTS/04f70d019af63defb8c92fa3e4c28a32fb5853a89d1a90ae110d8728c944ca3b/README.md) |
| 17 Oct 22 20:21 UTC | SHAKEN 849J | true | [view](CERTS/b222c99e36cacf3e5a2b5e2658aed255399b06f00e49bd92949a46540583770f/README.md) |
| 23 Oct 22 20:21 UTC | SHAKEN 469A | true | [view](CERTS/0bda01aefdb88d7db058c1023d7d8c7c414230edaac2382614cbe6bec7082a61/README.md) |
| 26 Oct 22 20:22 UTC | SHAKEN 469A | true | [view](CERTS/8d2fef64712c082b249ae7805b66ad21f2aef4fee890f786747f3bcf77051ed7/README.md) |
| 14 Sep 22 20:17 UTC | SHAKEN 469A | true | [view](CERTS/30a1d90e56f8e35a70a464b19475c9e6938b78a0100d717e3518fef373a31ba0/README.md) |
| 11 Sep 22 20:17 UTC | SHAKEN 469A | true | [view](CERTS/8cde12cc7cf8d6a7b649ff34a1090a399c1f9c51959456097640134337a3a996/README.md) |
| 05 Oct 22 20:20 UTC | SHAKEN 469A | true | [view](CERTS/8d23ef545cd3d1c1ad2a14f17bfe80c2029fae1b2b2758af683f8047ab485ca9/README.md) |
| 02 Oct 22 20:19 UTC | SHAKEN 469A | true | [view](CERTS/e44354f462a4eeeea70f60e84e82820988ae3b5143f7c6ec4ebb8e2246eff8f3/README.md) |
| 11 Oct 22 20:20 UTC | SHAKEN 726J | true | [view](CERTS/af8bce06149d79ea85c21923125515243952a2b3bea914f38088b7e3e3a92c8d/README.md) |
| 16 Sep 22 17:35 UTC | SHAKEN 177K | true | [view](CERTS/090db492c71684cf04f2dc05834e83f6459cfca2dd933b5241c7bf69a374aa8a/README.md) |
| 04 Oct 22 17:36 UTC | SHAKEN 177K | true | [view](CERTS/87c094e487860bdea20e828dc4839bd5170201dfa86a73d56f348a8b1583a3bf/README.md) |
| 16 Oct 22 17:38 UTC | SHAKEN 177K | true | [view](CERTS/639464440c8041617c3b4468254f7400a4a6d582bc0510eafeafaa2d9a9e3e57/README.md) |
| 28 Sep 22 17:36 UTC | SHAKEN 177K | true | [view](CERTS/88ec0edb8eb1cf5216c8fc2321300f0855136d2af270e4a259595aab9567874b/README.md) |
| 13 Sep 22 17:35 UTC | SHAKEN 177K | true | [view](CERTS/4354336a1145f337a2e589cd6f66cf4c2b7c44602d2e18e5854805eed552e0a3/README.md) |
| 01 Oct 22 17:36 UTC | SHAKEN 177K | true | [view](CERTS/b4e5e68daef0a733c9c37e33180500ee7177233cd00a706b2cdfc2d5791a6a25/README.md) |
| 19 Oct 22 17:38 UTC | SHAKEN 177K | true | [view](CERTS/e7cbd9ad7e28ad8503da0162a83b721025804b2f04d43144e6154a0d053baaf8/README.md) |
| 10 Oct 22 17:37 UTC | SHAKEN 177K | true | [view](CERTS/3c5b9e4d88ab0ae272e030e0890caa74ea1eb4291b59c6c024237ed1876f4d5c/README.md) |
| 11 Oct 22 14:26 UTC | SHAKEN 186K | true | [view](CERTS/6cb5ae1e0289b98a3c2581c84d8bb89b6b8f03b21929602fa8d809dd98c9d793/README.md) |
| 17 Sep 22 20:17 UTC | SHAKEN 495J | true | [view](CERTS/bb76a2fa4279f2c6351057a7f00994ebf4f480dac8c14bcc41e29559eee54080/README.md) |
| 08 Oct 22 20:20 UTC | SHAKEN 495J | true | [view](CERTS/fc49e2eb67b0f0727ef14485966719ea2ff58990d4426feec8260f7468669fbd/README.md) |
| 29 Sep 22 04:53 UTC | SHAKEN 747J | true | [view](CERTS/5443df111071af6a14b3de697459c9905d70f9b6b8c270d31bb2ae2c054602f1/README.md) |
| 11 Oct 22 04:54 UTC | SHAKEN 747J | true | [view](CERTS/ea61f8c5c5616d9d9b50a64c77e63dc2b3d792389fe6adcd55914be39fbedfd1/README.md) |
| 05 Oct 22 20:20 UTC | SHAKEN 721J | true | [view](CERTS/98fe8671fc2592236299d90bed3be05076290f76a55d2ca4369e7b9a9431820d/README.md) |
| 11 Oct 22 20:20 UTC | SHAKEN 790J | true | [view](CERTS/b9d2a595f2fb861374dda81e83844be049e5e241f998968ccab820accc7de538/README.md) |
| 26 Oct 22 20:22 UTC | SHAKEN 790J | true | [view](CERTS/854b6402cb0b321c3d810426dbf87b6b7b98210f836a6162f00759a448e96d9a/README.md) |
| 20 Oct 22 20:21 UTC | SHAKEN 790J | true | [view](CERTS/9f50e8e3dd1964fae0fcf3874a08fb68743b2bb7ab587d92ba9686a679bfe87c/README.md) |
| 17 Oct 22 20:21 UTC | SHAKEN 790J | true | [view](CERTS/034c042a844aad35b7b82e35657f2c74a7082576a713215b06821be99ecc85b2/README.md) |
| 14 Oct 22 20:20 UTC | SHAKEN 790J | true | [view](CERTS/14d42047446e5c374ed07a60ca501d56565ed1443b8a581304e3415d01fdd328/README.md) |
| 23 Oct 22 20:21 UTC | SHAKEN 790J | true | [view](CERTS/1390925b5b40d6068c0fd4dd9fff6d14a7f37920db0bcb51f4d73645e10beea1/README.md) |
| 14 Sep 22 20:17 UTC | SHAKEN 790J | true | [view](CERTS/723d6df813466a259ed8182e06d18b54ce239666d71ebfd9494b2b5b8cf31fb8/README.md) |
| 01 Oct 22 15:43 UTC | SHAKEN 111K | true | [view](CERTS/f2a0e4c22a50c1272647af9c27cc1629df90df7222028166209d29e9a60778ef/README.md) |
| 24 Oct 22 16:32 UTC | SHAKEN 291K | true | [view](CERTS/23cb935c39539c445b7e29922c31d6ae0934cd878ac72d173f7ab8f44b68c769/README.md) |
| 27 Oct 22 16:32 UTC | SHAKEN 291K | true | [view](CERTS/deefe1aa4ab2ec077f2a86adefc2664087be0c4392d579fc0ed49b04fe7345f6/README.md) |
| 09 Oct 22 16:30 UTC | SHAKEN 291K | true | [view](CERTS/4d88f9af77d39040c35346556a4f8d53b3b174cdeb51e3fd88efc218b6d69d15/README.md) |
| 26 Oct 22 20:23 UTC | SHAKEN 625J | true | [view](CERTS/436afa7fd59130c2cfb61515f1881bdb6b26df00104fc20a34813ec2837ae5bc/README.md) |
| 17 Oct 22 20:22 UTC | SHAKEN 625J | true | [view](CERTS/7c6dc385b76132cc9ea1054054b2bfc87ee4bdc302cb8a580e4ffc8420313543/README.md) |
| 14 Oct 22 20:21 UTC | SHAKEN 625J | true | [view](CERTS/28793d092f06155dda297977b8864e00b7bc5de7fcacab6af02fc29699d2eff4/README.md) |
| 23 Oct 22 20:22 UTC | SHAKEN 625J | true | [view](CERTS/46892b3113972786587e0551e964716f4d9f9ed7387b7b895d4237247f8b87a0/README.md) |
| 08 Oct 22 20:21 UTC | SHAKEN 738J | true | [view](CERTS/d4aa95acb150aa17e93399b56a346648269bb1bacf184fdf5ba1ae097aed70e5/README.md) |
| 17 Sep 22 20:18 UTC | SHAKEN 738J | true | [view](CERTS/a21040d088d56137aabaaae29fedf98f38eae7f21b3c6b2c51901a5112d59d04/README.md) |
| 29 Sep 22 20:20 UTC | SHAKEN 738J | true | [view](CERTS/626c8a126b630abafee9ce2f10e0c6d56e84970a804b12a3fc9d5b0a6efc5e92/README.md) |
| 02 Oct 22 20:20 UTC | SHAKEN 738J | true | [view](CERTS/ac493bd57bf9bcc24e22d17edef46a552b7f40eb3dbc68d83ca6f09893101018/README.md) |
| 23 Oct 22 20:22 UTC | SHAKEN 738J | true | [view](CERTS/805521e62bc664e5d33f357452aebe5ca36a088a7c23ad7e28a010bf33455dfc/README.md) |
| 14 Sep 22 20:18 UTC | SHAKEN 738J | true | [view](CERTS/e2f306df1acb052d08acb80a7cdcbc316d2ac8ab84ef16a08dcd520d02dc2337/README.md) |
| 14 Oct 22 20:21 UTC | SHAKEN 738J | true | [view](CERTS/2a2b74cf83a3188608342c4d0289a013e318a933e4471af06e7efab34efbad3a/README.md) |
| 26 Oct 22 20:23 UTC | SHAKEN 738J | true | [view](CERTS/24b935c4e2f91a82ff346147d216ffcfbbb408b7d41886788c3e7158c4573825/README.md) |
| 05 Oct 22 20:21 UTC | SHAKEN 738J | true | [view](CERTS/894d51d461ae687dc30a5e21f37c59e924e713737af5e99b912134ddedba5c09/README.md) |
| 17 Oct 22 20:22 UTC | SHAKEN 738J | true | [view](CERTS/3943e80ac38842f4ad7b096d59d95eabea5ef7a41cce8232c3960f3cc13a17a3/README.md) |
| 11 Oct 22 20:21 UTC | SHAKEN 738J | true | [view](CERTS/c29ac15321634a8219ce0c043c231a03e29e51b8e5e32bd871a76e62f83d194f/README.md) |
| 23 Oct 22 20:23 UTC | SHAKEN 459J | true | [view](CERTS/37e069ce4a66eaed7239ae6f5bb0fb163d93da2d8508be3c4cc52a610f92061f/README.md) |
| 08 Oct 22 20:21 UTC | SHAKEN 459J | true | [view](CERTS/1ad71ffcb3d1eaf2d0c52cb2e7fdced4ef6fdedf9179eef1f38654462ca5f5a2/README.md) |
| 14 Oct 22 20:21 UTC | SHAKEN 459J | true | [view](CERTS/47261a138370a0fc78fcc4537613a691499feaec718a7aea6a2d2dd3e2991c76/README.md) |
| 11 Oct 22 20:21 UTC | SHAKEN 459J | true | [view](CERTS/d6dd3187b91d14f2ebc67550b730b8a3b52026ced30e0a9687355f480f187d9f/README.md) |
| 14 Sep 22 20:18 UTC | SHAKEN 459J | true | [view](CERTS/e79b8d87c083113f995d12fb2f4940511b603eb9bb664ad647a7d0ea8d289128/README.md) |
| 20 Oct 22 20:22 UTC | SHAKEN 459J | true | [view](CERTS/5ae64855406e5f4c929fcf9da5dbdc6747d09079aa1750c368266022b0f2dc63/README.md) |
| 17 Oct 22 20:22 UTC | SHAKEN 459J | true | [view](CERTS/4c2285b24ab4293b97d78964e8e8766a8c62013b2d0c262a05548f9829a01a04/README.md) |
| 05 Oct 22 20:21 UTC | SHAKEN 459J | true | [view](CERTS/3aad834ca7555f9af90ebcfbf670bce7341749fb349ae1fa2a608f91e338fd40/README.md) |
| 26 Oct 22 20:23 UTC | SHAKEN 459J | true | [view](CERTS/b8c1d7beff3357f2ba432ff63529da29fd4871b1e11ae59e0f3d30eb6f7b4b51/README.md) |
| 17 Sep 22 20:18 UTC | SHAKEN 459J | true | [view](CERTS/b4697b04081f22439bb2912d81801ee35e551c70522ac483a1029f7cdd7a956a/README.md) |
| 02 Oct 22 20:21 UTC | SHAKEN 459J | true | [view](CERTS/a98a3300499ce3082cbf827588987e24b5e557f1b5a2c34d2941445f272c3614/README.md) |
| 29 Sep 22 20:20 UTC | SHAKEN 459J | true | [view](CERTS/3cafff7408f6653f7c24bc5990701caebc4476a60224e27177793b30c1d25494/README.md) |
| 17 Oct 22 18:35 UTC | SHAKEN 056K | true | [view](CERTS/72210f6d6661b52c1280643d61c2cc44c4a36529693635832b214e354358f322/README.md) |
| 27 Oct 22 19:07 UTC | SHAKEN 749J | true | [view](CERTS/1fc9e8d8a8e19300e114d8b68b9eed28cad0ceb35dc90bb4ca4d07c39116248d/README.md) |
| 17 Sep 22 20:20 UTC | SHAKEN 366G | true | [view](CERTS/1be939eb5386850a81fe6fd59e5dcc58d9017257e7b02aa963be3782cb8118cf/README.md) |
| 11 Oct 22 20:22 UTC | SHAKEN 366G | true | [view](CERTS/9aac33299bf80607c32d19ee4f6f195574a60bdf28397ca985f78b3f8872d729/README.md) |
| 11 Sep 22 20:19 UTC | SHAKEN 366G | true | [view](CERTS/debeec9b06e59490d7559598953e168664069d9eeee5b062412250cef1e2f46e/README.md) |
| 17 Oct 22 20:23 UTC | SHAKEN 366G | true | [view](CERTS/eef7c68ba075ac0f5c7fbd08794348e04158e2aac08c538fe42e0f8d001acd92/README.md) |
| 14 Oct 22 20:22 UTC | SHAKEN 366G | true | [view](CERTS/7676cb33ab883cc70a0570c302ac106963757fa8f7e881f06d1b5a995f62803c/README.md) |
| 23 Oct 22 20:24 UTC | SHAKEN 366G | true | [view](CERTS/edcb371dbe3e1e45fe42a9721f0504e8c9972904f3901dd4efacf7c08ed8c987/README.md) |
| 05 Oct 22 20:22 UTC | SHAKEN 366G | true | [view](CERTS/9b81a2e880cc60154f4e9fcc5db48387b54311c04c51758928d95f38faeb53c6/README.md) |
| 26 Oct 22 20:25 UTC | SHAKEN 366G | true | [view](CERTS/e3f28ec5803737280fee53e08b1097a304fc1b27cee05c9edb7836ce26691681/README.md) |
| 20 Oct 22 20:23 UTC | SHAKEN 366G | true | [view](CERTS/5aade4ffb8881607734fffef8ec90e38037807b55e7d8eab96f8f8eca6e1d8a7/README.md) |
| 29 Sep 22 20:21 UTC | SHAKEN 366G | true | [view](CERTS/28f09449ba27c9c0c38e1f34eadd0aa5bc11f3a245fc0d34cacc74163ab79d7e/README.md) |
| 08 Oct 22 20:22 UTC | SHAKEN 366G | true | [view](CERTS/cc0ccad347ff3466998bac1e592a0908d01af14bcad1882591335d50d9ea42a3/README.md) |
| 14 Sep 22 20:19 UTC | SHAKEN 366G | true | [view](CERTS/0428cbcc042aafb9bdd4e7b0df6147003b1463810efe578e8226a5241124e66c/README.md) |
| 02 Oct 22 20:22 UTC | SHAKEN 366G | true | [view](CERTS/5fa9e7c6bc69290d2340fbe1c98ac3de6876a05748478b1cbfaa747a1426c89d/README.md) |
| 17 Sep 22 20:20 UTC | SHAKEN 0226 | true | [view](CERTS/ac32dec9c9b45d840addd139e6056df0dcde243d2b9feba98e1eee82551def8f/README.md) |
| 23 Oct 22 20:24 UTC | SHAKEN 0226 | true | [view](CERTS/e0eaaabd8eb9f1f65823272068eb5c5e20f9b6d735a316743259a34635301d3a/README.md) |
| 14 Oct 22 20:23 UTC | SHAKEN 0226 | true | [view](CERTS/f9ffb12ca8fd13c1e2e659de346878eca525f3798a84e46f7debee05d84bd482/README.md) |
| 17 Oct 22 20:23 UTC | SHAKEN 0226 | true | [view](CERTS/e688f895705cee46e9acccbb37cde0acac7fff60dc106641cafda2936ba0c48b/README.md) |
| 05 Oct 22 20:22 UTC | SHAKEN 0226 | true | [view](CERTS/1a20d7424bfb464cf2bf63be43ac0a41520de8313c1462399e9cebab48b57294/README.md) |
| 14 Sep 22 20:19 UTC | SHAKEN 0226 | true | [view](CERTS/c02b42d1ee351811b281aed8390bf65d7842127d434b2a5fcf9fa14f7897d2e7/README.md) |
| 02 Oct 22 20:22 UTC | SHAKEN 0226 | true | [view](CERTS/3bc042aac6e90fa4f179cd26b0185138d1c6a75b2cb7c7decff882b4c94e2374/README.md) |
| 11 Oct 22 20:22 UTC | SHAKEN 0226 | true | [view](CERTS/5300ee7c6248c05a46316aa9bd4c8e140613a94eff851e68434f6ab1dc57139b/README.md) |
| 26 Oct 22 20:25 UTC | SHAKEN 0226 | true | [view](CERTS/fd4e9f41bea50e0e2615d2a9686e0dc7d4363c67912c4d355df0e4ec9b3faa14/README.md) |
| 08 Oct 22 20:22 UTC | SHAKEN 0226 | true | [view](CERTS/c829dc2ffe86a10d561d3c7a141f74047985f05529ee43fe7da60de70270efb6/README.md) |
| 08 Oct 22 20:22 UTC | SHAKEN 738J | true | [view](CERTS/6dc263b1393c6f5befcff77ea2754fc705f1db8e6c79d8abd6d2c6882a2df1b0/README.md) |
| 02 Oct 22 20:22 UTC | SHAKEN 738J | true | [view](CERTS/744cf9d06f47252440a9840d7ca36a0e26022d0aa9fa07266e37b10b133cff1d/README.md) |
| 14 Oct 22 20:23 UTC | SHAKEN 738J | true | [view](CERTS/481354b21938bab68d1a59074aea1cdf1324e34c5c1f0f6bcbe878bb306d79b2/README.md) |
| 29 Sep 22 20:21 UTC | SHAKEN 738J | true | [view](CERTS/fd0efdb081b6992a338c79959bc340244c1999d6e3c5f5006acc8ae64e43fdae/README.md) |
| 17 Oct 22 20:23 UTC | SHAKEN 738J | true | [view](CERTS/0643366e5695133ed60e1dca453c36984f21341f357bc19375d1936e3b230d72/README.md) |
| 17 Sep 22 20:20 UTC | SHAKEN 738J | true | [view](CERTS/6f6bcf4d575e0a0270e152d8acf3836c18da9571e649a138e8c55917ec571466/README.md) |
| 20 Oct 22 20:23 UTC | SHAKEN 738J | true | [view](CERTS/44045bad68814a0f22cf249d5cecd58cc07b6b92d79d7746860b411fb7a672d2/README.md) |
| 26 Oct 22 20:25 UTC | SHAKEN 738J | true | [view](CERTS/e227f5d1d545dc5289555dfb281c3b6e3642834bad66c25d997dedeee729f046/README.md) |
| 14 Sep 22 20:19 UTC | SHAKEN 738J | true | [view](CERTS/42ffdb54dfd082e0172dcaf2b7b10a8b051ef1e04a0332214d7b4c40b572b87e/README.md) |
| 11 Oct 22 20:22 UTC | SHAKEN 738J | true | [view](CERTS/59f769c2570b0aaf4234a4a020b69fe34deea1a96e5864d31db79e64026b518e/README.md) |
| 11 Sep 22 20:19 UTC | SHAKEN 738J | true | [view](CERTS/9377333893f452564ecadfcf0e08f04503ba6c9ff82761c78051e2d35572e5f6/README.md) |
| 23 Oct 22 20:24 UTC | SHAKEN 738J | true | [view](CERTS/412cb239f01180939fcdbe1e36f47197908f6f02e0c76110ed9fcdd0be8af307/README.md) |
| 05 Oct 22 20:22 UTC | SHAKEN 738J | true | [view](CERTS/957785e910b95c872c65f9dc918fd9da912c53190d3c0dfbfadab52888916bbf/README.md) |
| 14 Oct 22 20:48 UTC | SHAKEN 518J | true | [view](CERTS/1858ee3bdee363baa2677195b237ad135f5e74d0fb79adce8d4e3f198e525233/README.md) |
| 05 Oct 22 20:48 UTC | SHAKEN 518J | true | [view](CERTS/c462e1d0c297f5e002f696218debe98f7c723a5634f1e5808c9070e446ff4fab/README.md) |
| 20 Oct 22 20:48 UTC | SHAKEN 518J | true | [view](CERTS/4bbed7a278afcf1910ceebbd08330047b4f4472ce201ed1018e25de8d3169389/README.md) |
| 08 Oct 22 20:48 UTC | SHAKEN 518J | true | [view](CERTS/ff8690fa8747963cf6fa6cdf25d81048da96eb7aa7fab45e12453ff04e2b85af/README.md) |
| 11 Oct 22 20:48 UTC | SHAKEN 518J | true | [view](CERTS/176416a90dba031bcd9802c7c37cee5ec99a6af34f65ffca04e340b811bc9cde/README.md) |
| 02 Oct 22 20:48 UTC | SHAKEN 518J | true | [view](CERTS/6744eb9caee8a51ab3839993b39df773e60d6f64bb3217c6ca4578c920309645/README.md) |
| 17 Oct 22 20:48 UTC | SHAKEN 518J | true | [view](CERTS/30273b668938e69f0d1103db1be272cf4261f31132377424e34566827ce214b3/README.md) |
| 17 Sep 22 20:48 UTC | SHAKEN 518J | true | [view](CERTS/20859b22cbc2dc8117cca69315a7e9e21a2f5e81d4335a855659077078c7fae8/README.md) |
| 23 Oct 22 20:48 UTC | SHAKEN 518J | true | [view](CERTS/6eebf5881e379b5de3bf62878cf0ad4c79d3cedb49f1ca890c8b87e4919df0c5/README.md) |
| 20 Sep 22 07:11 UTC | SHAKEN 0172 | true | [view](CERTS/529e4238d06270ae8298b7339ff9b73ae629c8615e260541d4a5cb60ae5c44c1/README.md) |
| 11 Oct 22 07:34 UTC | SHAKEN 0172 | true | [view](CERTS/48ab34f5193e1edadb154ad6c1a750b37d1b3c0a5ecf4bde295f1fef544e9d4e/README.md) |
| 15 Oct 22 12:22 UTC | SHAKEN 0172 | true | [view](CERTS/22d6a8fbca807d7d94efb3068fcdb32f413202fb87ff5d82f8e6981692bd353d/README.md) |
| 07 Oct 22 02:42 UTC | SHAKEN 0172 | true | [view](CERTS/e94e41f3c53de56e420bd3b96944a452318738088dd609c14fd2f06d3d7cbcd6/README.md) |
| 02 Oct 22 21:53 UTC | SHAKEN 0172 | true | [view](CERTS/e7cf263a429feef70ae06c69e36b1d41bc8a2b0071c7c4a717f367980bd07cd1/README.md) |
| 11 Sep 22 21:31 UTC | SHAKEN 0172 | true | [view](CERTS/154c58ca83926bee0c437987b8239980654b9cc98b688489636ca85159300eb9/README.md) |
| 23 Oct 22 22:11 UTC | SHAKEN 0172 | true | [view](CERTS/1e17f2953a3d29c86b5b1e45a664b3a39ee924f864b266898ef69d613e747765/README.md) |
| 19 Oct 22 17:22 UTC | SHAKEN 0172 | true | [view](CERTS/8b637ec5d4c08cc6f0cac6b9d21a2347d4f574fd9e026e2d0598c48586abf681/README.md) |
| 10 Oct 22 18:39 UTC | SHAKEN 066K | true | [view](CERTS/51c8cf0bbfb640f7ef4e7e47bedd1a9a1d4f757fe142ad05fa84422544794ccd/README.md) |
| 10 Aug 22 18:11 UTC | SHAKEN 073H | true | [view](CERTS/87dabb91a689a6ac540b5ffbf5eb89326fada83fb822d2bb9ec31b68251b579e/README.md) |
| 01 Jul 22 15:39 UTC | SHAKEN 093K | true | [view](CERTS/4653cd0d85308722c74cbf97447900a9e9318556fcfeaf0113c435f45b867251/README.md) |
| 10 Oct 22 18:30 UTC | SHAKEN 115K | true | [view](CERTS/631861bcfeb6aeaf83a56e5c1c376c8981927a52fc8da57964b1776b69e2bf6e/README.md) |
| 01 Jul 22 15:30 UTC | SHAKEN 148K | true | [view](CERTS/c5c51376832cc1ddc1007cc8e8cb79973cf959cf85596f753bc1f0adc8961e9b/README.md) |
| 12 Sep 22 09:13 UTC | SHAKEN 159H | true | [view](CERTS/39eb9c5dbd7a485d82ed9cf5262515f86e35d66e89eca7f4781f102cd27df04b/README.md) |
| 03 Oct 22 09:34 UTC | SHAKEN 159H | true | [view](CERTS/d96c81ec513b16256b347248706ca1affab7e834682f5321c3a41f4816c6bd97/README.md) |
| 08 Oct 22 15:35 UTC | SHAKEN 159H | true | [view](CERTS/16021ca7a8fa17b9eef493e41dd2d927d619bd7c5889fd91c75bf43fe9692910/README.md) |
| 13 Oct 22 21:41 UTC | SHAKEN 159H | true | [view](CERTS/379b10564d859b3ee0ffe6d9e113b969d8ae09aa1c39754ce4d2cf332f841498/README.md) |
| 01 Sep 22 12:31 UTC | SHAKEN 193E | true | [view](CERTS/04de048b1d1e8aad1f55abd43cabed4f2398c0e486dab878e0585d5b7d857bbe/README.md) |
| 01 Oct 22 12:31 UTC | SHAKEN 193E | true | [view](CERTS/2a2a297de4aa7620b8f38d5266dd4f76fbd6b5c79a87411b589c2226518ce475/README.md) |
| 01 Oct 22 04:04 UTC | SHAKEN 345J | true | [view](CERTS/d8165a1ac88e14630152e270616496da1cf609b10b965466b5efdf477ca92a34/README.md) |
| 10 Sep 22 04:01 UTC | SHAKEN 345J | true | [view](CERTS/4d839c60c791950256e186267b1483643334f1a529954e0a6e61462b8d9cc4bd/README.md) |
| 22 Oct 22 04:04 UTC | SHAKEN 345J | true | [view](CERTS/27c018db560aefbff83f32326db2cbf3f1260051386569d7775cf98fae145bae/README.md) |
| 10 Sep 22 04:23 UTC | SHAKEN 345J | true | [view](CERTS/97351f4d58baf9f87cfe826be2b11c786fd1648721bbf2887156c37941c8a735/README.md) |
| 01 Oct 22 04:34 UTC | SHAKEN 345J | true | [view](CERTS/a25f95eb3066f7622549b3bb9bfa7e21056092208ca3db3451ed182bb5e4a2db/README.md) |
| 22 Oct 22 04:44 UTC | SHAKEN 345J | true | [view](CERTS/109a4ac394ec4e6eb6ba42d2ad6bc799f2548629a8bbfb9420b9ec1ca5df6ff9/README.md) |
| 26 Sep 22 18:40 UTC | SHAKEN 505J | true | [view](CERTS/b186c959554e652b69824a2e45ac08ec135105b0e01c9d2eadeb3cf46c130670/README.md) |
| 26 Sep 22 18:42 UTC | SHAKEN 505J | true | [view](CERTS/1c0003ac10eeaa04229e507c15f71ee018c3902c3e8c20fbe42e533b8682ba8c/README.md) |
| 12 Oct 22 04:55 UTC | SHAKEN 551G | true | [view](CERTS/894874357a30f38e290f1b44d4ab11fb735a5eace7897bdbffa9fd68792a31c8/README.md) |
| 17 Oct 22 02:42 UTC | SHAKEN 551G | true | [view](CERTS/2223b9c56768b8a379657a650964a92360da7bdbf0939a72da664edd6520704e/README.md) |
| 02 Oct 22 09:31 UTC | SHAKEN 551G | true | [view](CERTS/2bfddc65cec446ec2c673b3744fc9153104c6f6e59cb6ffde053eb4ce80f09db/README.md) |
| 22 Oct 22 00:21 UTC | SHAKEN 551G | true | [view](CERTS/b85f047a667023f692cdc92e6c0de3fa2ee15b49064bee39fe8f0ff9abef9bb6/README.md) |
| 27 Sep 22 11:52 UTC | SHAKEN 551G | true | [view](CERTS/bb3e3cd31cd462f8592c2cd781f0678bac0055f6278a1a549728f0f97ae8a09c/README.md) |
| 12 Sep 22 18:53 UTC | SHAKEN 551G | true | [view](CERTS/d97057a9e3ae4abf7546ff40b5c9be952c194455de33256871c8472423eb541e/README.md) |
| 26 Oct 22 22:00 UTC | SHAKEN 551G | true | [view](CERTS/6052628e9622c0b9247ee1139aec68db77c73a3857882d712de67adf4cce6083/README.md) |
| 17 Sep 22 16:32 UTC | SHAKEN 551G | true | [view](CERTS/27a83315430b43e45e2d7235880b96bf75e150a5196fe7fbc08cb6018289c348/README.md) |
| 07 Oct 22 07:10 UTC | SHAKEN 551G | true | [view](CERTS/d78d702c9e41c41427aa707e6dda750cfd9ee9fa70792841ef60cdeae7b4e2e4/README.md) |
| 25 Jul 22 18:36 UTC | SHAKEN 578J | true | [view](CERTS/f4af5a28131d680ef2682ae86181b64f0c919864d7dbaef00c0598d314b537a1/README.md) |
| 10 Jun 22 14:00 UTC | SHAKEN 597J | true | [view](CERTS/4baa838eb96ce358aa1565d0add8b75072258efd5d21a5bbb0eec878026a67be/README.md) |
| 25 Jul 22 17:43 UTC | SHAKEN 622J | true | [view](CERTS/14f21e82532fc2d45d4d8b09592d32c19aa3ee63343a81fca8564bb4d160c841/README.md) |
| 20 Oct 22 15:48 UTC | SHAKEN 622J | true | [view](CERTS/2e34c765f110b197f2ec1052b1960bc25ce320dfce7d7035cc532cd1eb6dbd57/README.md) |
| 01 Sep 22 03:25 UTC | SHAKEN 6628 | true | [view](CERTS/37b9a67bb8e3272048330f269b5dbc285ac0278df4670ab7e291b92b6548fa2d/README.md) |
| 27 Apr 22 18:22 UTC | CCI SHAKEN 663J | true | [view](CERTS/9cfd0b9974acdc9b5fce3888cf2f6cf99173e7955adeebd37a1e60dfc444a0b3/README.md) |
| 29 Jul 21 22:04 UTC | TCN SHAKEN 706J | true | [view](CERTS/221a0b74876597aaab3a56ac25b522c35127f049c58ecd6cfcb02aada95144d0/README.md) |
| 13 Sep 22 17:24 UTC | SHAKEN 706J | true | [view](CERTS/1eea5e17ae8b36cb5b6d657178273af745cea015e63586e6d44d1ba2ee8f722a/README.md) |
| 07 Sep 22 15:22 UTC | SHAKEN 722J | true | [view](CERTS/95828a1468dffcee0669039c3eac65fec7d9517a383693700d99e97f5c741a74/README.md) |
| 01 Oct 22 15:23 UTC | SHAKEN 722J | true | [view](CERTS/4f02c0369baabbf0d1b7fd0f2adbb8d7b80a0b5693c2dd53a9bc22280fd8cd3c/README.md) |
| 25 Oct 22 15:33 UTC | SHAKEN 722J | true | [view](CERTS/50eb0c6a670f4122f8fbdb75582aa257fe3979f441fb396ff738372627104f9c/README.md) |
| 29 Jun 22 20:24 UTC | SHAKEN 736J | true | [view](CERTS/71b25a480077816998e48e6a76f272e70aa91ecb19b0e773cd8777823cdead32/README.md) |
| 10 Oct 22 20:20 UTC | SHAKEN 738J | true | [view](CERTS/435b78adba44dc3d7d5e0be8a9369539b8f58cde526c7e78b09a1b02f2e7730a/README.md) |
| 03 Oct 22 20:19 UTC | SHAKEN 738J | true | [view](CERTS/335922452742aec7089c5bca38b4b053b8a03cd69afe03f03bcf42c5f0e5a878/README.md) |
| 19 Sep 22 20:17 UTC | SHAKEN 738J | true | [view](CERTS/96adc1278a56fbebd3b5c156ae55329919cb46543824fa6c4fbbc398e23ba178/README.md) |
| 26 Sep 22 20:18 UTC | SHAKEN 738J | true | [view](CERTS/f5826e258a867d322ff75001275716441c9b5c40a818293dccd5e2ad8bc308c0/README.md) |
| 12 Sep 22 20:16 UTC | SHAKEN 738J | true | [view](CERTS/f7b37f810dd36f526c1ec6faf7e89a927b8938eea12b23c73a2e7c45fe882dbd/README.md) |
| 24 Oct 22 20:22 UTC | SHAKEN 738J | true | [view](CERTS/1f99a81d0b090f467a47da69b1549fbabc7d8277f67a1f31d05c53e92cceb54c/README.md) |
| 17 Oct 22 20:21 UTC | SHAKEN 738J | true | [view](CERTS/6b3cccbf5f4070a3f7602699fb951fc32e3e8cd06e68e6dbef854525ed00354b/README.md) |
| 22 Jun 22 18:46 UTC | SHAKEN 755J | true | [view](CERTS/a8cac316fddf25b48b9350a325791be8b87e6484a5959ab605db1ae5e3e99a84/README.md) |
| 06 Jul 22 20:42 UTC | SHAKEN 807J | true | [view](CERTS/471313a1c0ac5cee09b7397ddc0fedbdb18122d13cc73b5596486af0732b44dc/README.md) |
| 12 Oct 22 19:39 UTC | SHAKEN 815G | true | [view](CERTS/a79f50d4fcd3e680518379f12a86a4b9a21c5e27dc5422a8439525edf350c4a7/README.md) |
| 13 Jul 22 15:16 UTC | SHAKEN 815G | true | [view](CERTS/c156b07494124e858b6b7848c0d97b843a4629f0acb9b42d808952a45524a230/README.md) |
| 03 Oct 22 13:54 UTC | SHAKEN 841J | true | [view](CERTS/6476653a99f62e0c8aebcbb119af12a12deb3640b8ce105eb21cc6dd7301adc1/README.md) |
| 16 Sep 22 18:33 UTC | SHAKEN 841J | true | [view](CERTS/d79a73488c528c2b6483d0c129461abb70897742514137334f925551702d0678/README.md) |
| 25 Sep 22 04:14 UTC | SHAKEN 841J | true | [view](CERTS/ff2f398c0b380bb657022561a2447ba63f092c132b87397e9f6bb4023df6881a/README.md) |
| 11 Oct 22 23:30 UTC | SHAKEN 841J | true | [view](CERTS/9f333ccba7c99ededf9e1602476aa3f48c5c806d001aef41e74a6bfad898207a/README.md) |
| 20 Oct 22 09:13 UTC | SHAKEN 841J | true | [view](CERTS/8c54ce0987e3860a3b0863c2b6947d4c38a04acf52ebeed106e0af6e5300abcb/README.md) |
| 05 Sep 22 10:32 UTC | SHAKEN 8526 | true | [view](CERTS/26c56c58d7583997444745e9e42123729816187a351dc08c841d03b17f4dc259/README.md) |
| 30 Sep 22 22:34 UTC | SHAKEN 8526 | true | [view](CERTS/e4c465211a272c98c4b2bee1170632cb055c08bf78f91778572652278ccba4c4/README.md) |
| 26 Oct 22 10:43 UTC | SHAKEN 8526 | true | [view](CERTS/b68dadcfa6267b8e9dd012e42154b55cab6d3694f4543e46af44bab1e4ba971e/README.md) |
| 13 Jan 22 14:17 UTC | MobileSphere SHAKEN 873J | true | [view](CERTS/8f74ddd5e6964042b119d5b3f17d24694ec78c8688e0ffc96d13fdb0e8df26d4/README.md) |
| 07 Sep 22 17:54 UTC | SHAKEN 952J | true | [view](CERTS/9e7bb57b92b580c347de45d6faf291a9b9d60d40150891039c33c1295382fdc5/README.md) |
| 25 Sep 22 17:54 UTC | SHAKEN 952J | true | [view](CERTS/9c190dcf26694878152028da41b4a461451f63bcec7f3dcbec8985282105505e/README.md) |
| 13 Oct 22 18:02 UTC | SHAKEN 952J | true | [view](CERTS/bbccd67a84881c78b37ecfb961dd7029498490975d91e15f872d5a684a225626/README.md) |
| 20 Jun 22 20:25 UTC | SHAKEN 857J | true | [view](CERTS/e1e0a4851fb7b51d617959c00fb6a17e7b23d34e1555ba70ab01c9182a2f85f8/README.md) |
| 30 Mar 22 12:54 UTC | Fusion Connect SHAKEN 2720 | true | [view](CERTS/03a010b294807e90cc1309d7466fd2ddc591158a69be6950ff4d0eab32de0799/README.md) |
| 02 Nov 21 16:40 UTC | Charter Communications Inc SHAKEN 5606 | true | [view](CERTS/6e6a467b634c931ad2b441fcdbeeec1d29fb3b3a9687018774c563fa866e00ed/README.md) |

#### CA Certificates

| Created at | Subject | Problems | Link |
|------------|---------|----------|------|
| 20 Aug 21 00:00 UTC | TransNexus, Inc. SHAKEN Issuing CA3 | true | [view](CERTS/51b51431ffd8055cd7e30c213540f4868258058d4c4c9f23f1ed8bea75101e37/README.md) |
| 20 Aug 21 00:00 UTC | TransNexus, Inc. SHAKEN Issuing CA1 | true | [view](CERTS/037a7c6621814bd5f01fa2080a77a245458694d9d566e0050e068fc853317857/README.md) |
| 04 Nov 20 00:00 UTC | TransNexus Issuing CA G2 | true | [view](CERTS/56ee2c4c4f58e02b74b75eb3476c2f16a3d0c94a0dcfa1f28c214cdd90b8a466/README.md) |
| 20 Aug 21 00:00 UTC | TransNexus, Inc. SHAKEN Root CA1 | true | [view](CERTS/59ab943f4e8c3c17755cc8e4abb5ac65736ce74ec587c62cb744ab0babbbe2fe/README.md) |


Generated: 31/10/2022 at 16:43:22