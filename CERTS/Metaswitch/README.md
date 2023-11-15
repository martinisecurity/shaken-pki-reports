# STIR/SHAKEN CA Ecosystem Compliance

## Metaswitch

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 110 certificates were included in the corpus being tested
- 39 certificates in the corpus were skipped because they are duplicates
- 0 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 71 certificates being tested against the remaining rules
- 4.39 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 1.41% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 57.75% of certificates are too old to be assessed against currently enforced expectations
- 1081 days is the average remaining validity for the certificates in the corpus
- 1095 days is the average initial validity for the certificates in the corpus
- 1 certificates expire in the next 30 days
- 1.01 average number of unexpired certificates per OCN observed
- 70 unique OCNs observed in unexpired and valid certificate corpus

| Instances | Test | Source |
|-----------|------|--------|
| 71 | [e_atis_issuer_dn](ISSUES/e_atis_issuer_dn/README.md) | ATIS1000080 |
| 71 | [e_atis_key_usage](ISSUES/e_atis_key_usage/README.md) | ATIS1000080 |
| 1 | [e_atis_subject](ISSUES/e_atis_subject/README.md) | ATIS1000080 |
| 26 | [e_atis_subject_cn](ISSUES/e_atis_subject_cn/README.md) | ATIS1000080 |
| 71 | [e_us_cp_ambiguous_identifier](ISSUES/e_us_cp_ambiguous_identifier/README.md) | US_SHAKEN_CP |
| 71 | [e_us_cp_subject_sn](ISSUES/e_us_cp_subject_sn/README.md) | US_SHAKEN_CP |
| 1 | [w_shaken_subject_rdn_unknown](ISSUES/w_shaken_subject_rdn_unknown/README.md) | SHAKEN_PKI_BEST_PRACTICES |

#### CA Certificates

- 3 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 0 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 3 certificates being tested against the remaining rules
- 1.67 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 0.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 66.67% of certificates are too old to be assessed against currently enforced expectations
- 5569 days is the average remaining validity for the certificates in the corpus
- 5353 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days

| Instances | Test | Source |
|-----------|------|--------|
| 1 | [e_atis_ca_certificate_policies](ISSUES/e_atis_ca_certificate_policies/README.md) | ATIS1000080 |
| 1 | [e_atis_ca_issuer_dn](ISSUES/e_atis_ca_issuer_dn/README.md) | ATIS1000080 |
| 3 | [e_atis_ca_subject](ISSUES/e_atis_ca_subject/README.md) | ATIS1000080 |

### Details

#### Leaf Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 11&#160;Dec&#160;20&#160;00:14&#160;UTC | Fidelity Communications SHAKEN Cert 1882 | 11&#160;Dec&#160;23&#160;00:14&#160;UTC | true | [view](CERTS/baaf8e58db0f02327fc5b9b614a5633c7b505ca0b291b606d32a191ee73a05e5/README.md) |
| 17&#160;Dec&#160;20&#160;15:19&#160;UTC | TDS Telecom SHAKEN Cert 7804 | 17&#160;Dec&#160;23&#160;15:19&#160;UTC | true | [view](CERTS/a04a669738b79ff55c2b2197f72a12a112b731906e2e6a925d37ccee2fa00a11/README.md) |
| 27&#160;Jan&#160;21&#160;17:16&#160;UTC | Telesystem SHAKEN Cert 786E | 27&#160;Jan&#160;24&#160;17:16&#160;UTC | true | [view](CERTS/2d9aca0895c94291596161363091718089a6e7c19dfa57329ae548432533860f/README.md) |
| 01&#160;Feb&#160;21&#160;16:46&#160;UTC | Rainbow Communications SHAKEN Cert 1820 | 01&#160;Feb&#160;24&#160;16:46&#160;UTC | true | [view](CERTS/8dddc19bd59ffef92b8e34656d16d27f50cd7d2889c19a7691e00e7ba80f6acb/README.md) |
| 16&#160;Feb&#160;21&#160;17:30&#160;UTC | Verizon SHAKEN cert 5807 | 16&#160;Feb&#160;24&#160;17:30&#160;UTC | true | [view](CERTS/d7b413267be2d050d516af4f4a864ffdc2feacc774a1a6264b9cfe68c796f43f/README.md) |
| 10&#160;Mar&#160;21&#160;20:40&#160;UTC | Allstream SHAKEN cert 4130 | 09&#160;Mar&#160;24&#160;20:40&#160;UTC | true | [view](CERTS/a9eec584cf66eade89fa84e043e0eb05900dc40653ea963a96c079ef92d9e2ad/README.md) |
| 10&#160;Mar&#160;21&#160;20:50&#160;UTC | Northeast Communications of Wisconsin SHAKEN Cert 6692 | 09&#160;Mar&#160;24&#160;20:50&#160;UTC | true | [view](CERTS/90042ab31d5de7abb64b89073a0d931ce22d864b89df13a9372887cb5db45f49/README.md) |
| 17&#160;Mar&#160;21&#160;17:06&#160;UTC | RCN SHAKEN Cert 7615 | 16&#160;Mar&#160;24&#160;17:06&#160;UTC | true | [view](CERTS/bbdec20ad80f4a2a8ed097204a9299566beca170460fb648c81a51d195d9b6f1/README.md) |
| 17&#160;Mar&#160;21&#160;17:12&#160;UTC | GCI SHAKEN Cert 7785 | 16&#160;Mar&#160;24&#160;17:12&#160;UTC | true | [view](CERTS/312e58dffa682b464f9867a7c373f9881d092b834767dcabe5baf8c7245e937c/README.md) |
| 17&#160;Mar&#160;21&#160;17:19&#160;UTC | South Central Rural Telecommunications Coop SHAKEN Cert 0418 | 16&#160;Mar&#160;24&#160;17:19&#160;UTC | true | [view](CERTS/6e988e3a74dbb6c1862118841d78e438f345c1ef7e79a96c2087328b5de146c1/README.md) |
| 17&#160;Mar&#160;21&#160;17:25&#160;UTC | Yelcot SHAKEN Cert 1733 | 16&#160;Mar&#160;24&#160;17:25&#160;UTC | true | [view](CERTS/f7e9897313ee276a419725d0aa81886e8f3636ad1cd1e9aad623166f56e6b141/README.md) |
| 24&#160;Mar&#160;21&#160;17:57&#160;UTC | Duo Broadband SHAKEN Cert 0401 | 23&#160;Mar&#160;24&#160;17:57&#160;UTC | true | [view](CERTS/166507f7a0232396c751d4600be71de8a6e5d9312a92b1208b2954ca04236e74/README.md) |
| 24&#160;Mar&#160;21&#160;23:52&#160;UTC | Cspire SHAKEN Cert 6581 | 23&#160;Mar&#160;24&#160;23:52&#160;UTC | true | [view](CERTS/09ed5b3292b5bfc7ac80b1027a827138b9503aa8053a61431a8dc851ecad04f2/README.md) |
| 01&#160;Apr&#160;21&#160;15:28&#160;UTC | USCellular SHAKEN Cert 6349 | 31&#160;Mar&#160;24&#160;15:28&#160;UTC | true | [view](CERTS/c75937e7c843b0054b1a51f6149f47661696b3cfcbcf6bab4dbc708c4453244f/README.md) |
| 07&#160;Apr&#160;21&#160;16:44&#160;UTC | U. S. Telepacific Corp SHAKEN 7453 | 06&#160;Apr&#160;24&#160;16:44&#160;UTC | true | [view](CERTS/9ed03dac797a5a27d52aa5209a4caa6a3ec9c3943d55a2cbfb69416480787da0/README.md) |
| 16&#160;Apr&#160;21&#160;15:27&#160;UTC | CBTS Technology Solutions SHAKEN Cert 600F | 15&#160;Apr&#160;24&#160;15:27&#160;UTC | true | [view](CERTS/3d02021a2da14f1ebfe588256a419be9ebc03c0d1fccc51cc29fa9d4a625c6bf/README.md) |
| 30&#160;Apr&#160;21&#160;16:51&#160;UTC | 702 Communications SHAKEN Cert 2737 | 29&#160;Apr&#160;24&#160;16:51&#160;UTC | true | [view](CERTS/c5f150937d7e08266fb699af1cacc1767ea6249bddc2ce99a9f468553d4486ee/README.md) |
| 30&#160;Apr&#160;21&#160;16:59&#160;UTC | Nex-Tech Wireless SHAKEN Cert 122D | 29&#160;Apr&#160;24&#160;16:59&#160;UTC | true | [view](CERTS/efe569be452e63d91da204eeda5030f1b08846836bf9bf036f51b1da5411acdc/README.md) |
| 30&#160;Apr&#160;21&#160;17:05&#160;UTC | Hunter Communications Shaken Cert 660C | 29&#160;Apr&#160;24&#160;17:05&#160;UTC | true | [view](CERTS/0d2022504ffa5407f662990a785786cb0da72ce014838e2cffdcd95cb70c6f64/README.md) |
| 05&#160;May&#160;21&#160;21:02&#160;UTC | American Broadband SHAKEN Cert 355D | 04&#160;May&#160;24&#160;21:02&#160;UTC | true | [view](CERTS/f2fe275fba183def77918369f90d81bc080f1ef58c5422620c3dc45140b2ae75/README.md) |
| 18&#160;May&#160;21&#160;13:14&#160;UTC | Segra SHAKEN Cert 1784 | 17&#160;May&#160;24&#160;13:14&#160;UTC | true | [view](CERTS/f802b8b879d063b87665d8b2a67f6d3ba78a94aa35782f664a6a40afc7586f56/README.md) |
| 18&#160;May&#160;21&#160;13:22&#160;UTC | Viaero Wireless SHAKEN Cert 6874 | 17&#160;May&#160;24&#160;13:22&#160;UTC | true | [view](CERTS/fe826427006707276b684368c81c7b2039d4ce72b25c5cf0a0c039993b026573/README.md) |
| 20&#160;May&#160;21&#160;22:04&#160;UTC | Appalachian Wireless SHAKEN Cert 6940 | 19&#160;May&#160;24&#160;22:04&#160;UTC | true | [view](CERTS/e14170c681e75c37d0ca45e304c09cc0d148246bd7d72e96f91f7a8fe27339fa/README.md) |
| 20&#160;May&#160;21&#160;22:13&#160;UTC | Nemont SHAKEN Cert 2247 | 19&#160;May&#160;24&#160;22:13&#160;UTC | true | [view](CERTS/cd4fc6aff73ae9a3e16063ce1bda6bc6e265584b0b930d24ef615174fa6bed20/README.md) |
| 20&#160;May&#160;21&#160;22:20&#160;UTC | Everstream SHAKEN Cert 472C  | 19&#160;May&#160;24&#160;22:20&#160;UTC | true | [view](CERTS/8710bb38debebd39698fb1c273409b173951cca1fab53a6d4c4aca91e61e06df/README.md) |
| 26&#160;May&#160;21&#160;15:37&#160;UTC | WOW Internet Cable and Phone SHAKEN Cert 665E | 25&#160;May&#160;24&#160;15:37&#160;UTC | true | [view](CERTS/9f86e038b4f6c3bf7fabfe8d4008b071f81bd33a80c55ea66a1aadd79eb9b989/README.md) |
| 27&#160;May&#160;21&#160;13:28&#160;UTC | Mediacom 846F | 26&#160;May&#160;24&#160;13:28&#160;UTC | true | [view](CERTS/e6c9e9fd411d8174b3ffe1af4d569c6919f4b98a5d0c6e429cd3682d82284e7e/README.md) |
| 04&#160;Jun&#160;21&#160;17:29&#160;UTC | Union Telephone Company SHAKEN Cert 2297 | 03&#160;Jun&#160;24&#160;17:29&#160;UTC | true | [view](CERTS/e5d49f06964a0d852a706d02db3742482fada1f2c550cef4119ce0e25f80590e/README.md) |
| 11&#160;Jun&#160;21&#160;16:01&#160;UTC | Clearwave SHAKEN Cert 9915 | 10&#160;Jun&#160;24&#160;16:01&#160;UTC | true | [view](CERTS/de0c13b76c24e2ef34c5de58e88fbfeef9c8adb44175b42f3c69771db6586491/README.md) |
| 11&#160;Jun&#160;21&#160;16:14&#160;UTC | Carolina West Wireless SHAKEN Cert 5932 | 10&#160;Jun&#160;24&#160;16:14&#160;UTC | true | [view](CERTS/e85dec068c9fa23af68b345257d58269303c925d10fc9101b226d1e0e7d62a9d/README.md) |
| 15&#160;Jun&#160;21&#160;13:56&#160;UTC | ENA SHAKEN cert 521F | 14&#160;Jun&#160;24&#160;13:56&#160;UTC | true | [view](CERTS/a5293c77ed92fdeef1aeff04f6bed7932a6083abc55d03d74a3ec5b817fdba2b/README.md) |
| 21&#160;Jun&#160;21&#160;15:05&#160;UTC | Buckeye SHAKEN Cert 7608 | 20&#160;Jun&#160;24&#160;15:05&#160;UTC | true | [view](CERTS/61a18f942b4df978a21d6e9cf1df2f0cea0dbd9609b1fcfec8c5f97d5c6dcd7e/README.md) |
| 24&#160;Jun&#160;21&#160;22:33&#160;UTC | Hawaiian Telcom SHAKEN Cert 009G | 23&#160;Jun&#160;24&#160;22:33&#160;UTC | true | [view](CERTS/51136b93b3f0cdaef7856044881e5abc0747fe2b4d9a3202f78266ab5228a41d/README.md) |
| 30&#160;Jun&#160;21&#160;16:55&#160;UTC | Sonic Telecom SHAKEN cert 433E | 29&#160;Jun&#160;24&#160;16:55&#160;UTC | true | [view](CERTS/a5082b808b3bf200a634e273b7988a617099c9c897a9c596858e6d4ffe4fd352/README.md) |
| 30&#160;Jun&#160;21&#160;16:59&#160;UTC | Utility SHAKEN Cert 9262 | 29&#160;Jun&#160;24&#160;16:59&#160;UTC | true | [view](CERTS/97acbe1f0c7b2ceb99bec93a9da9ff2c1fd55e9a4cd78a32160237751f720236/README.md) |
| 20&#160;Jul&#160;21&#160;19:54&#160;UTC | New Horizon SHAKEN Cert 127E | 19&#160;Jul&#160;24&#160;19:54&#160;UTC | true | [view](CERTS/5b2842e49ecc543187018171fa660e32bdba390a7977c0de97a00aef35b8ae01/README.md) |
| 23&#160;Jul&#160;21&#160;16:59&#160;UTC | Tri-County Telephone Association Inc SHAKEN Cert 2296 | 22&#160;Jul&#160;24&#160;16:59&#160;UTC | true | [view](CERTS/059194bbdc7ecfd6fc3d30b071e4d73565e14360c7a9d19c60f9274ec1025d0f/README.md) |
| 30&#160;Jul&#160;21&#160;17:04&#160;UTC | Foothills Rural Tel SHAKEN Cert 0406 | 29&#160;Jul&#160;24&#160;17:04&#160;UTC | true | [view](CERTS/89a593de82dbaebcffd8d3d577faf5d7f4a6a898608bc75f9cbde11a9cf4f179/README.md) |
| 30&#160;Jul&#160;21&#160;17:16&#160;UTC | Syringa Networks SHAKEN Cert 5869 | 29&#160;Jul&#160;24&#160;17:16&#160;UTC | true | [view](CERTS/5b9b091228790e797adf24925cd1f2508b8372b744611f843a8026a703993d5a/README.md) |
| 03&#160;Aug&#160;21&#160;10:57&#160;UTC | KEPS Technologies INC SHAKEN Cert 3535 | 02&#160;Aug&#160;24&#160;10:57&#160;UTC | true | [view](CERTS/30e0ccb8a573d7422799716c4253e8bee0898d934de3ff555738175c9c80cbea/README.md) |
| 16&#160;Sep&#160;21&#160;13:09&#160;UTC | CTS Telecom, Inc SHAKEN Cert 8331 | 15&#160;Sep&#160;24&#160;13:09&#160;UTC | true | [view](CERTS/5c9f6132a96a05a58d02d845ee70e914659682623fb98acbbecd9f8e383dcbec/README.md) |
| 02&#160;Mar&#160;22&#160;17:52&#160;UTC | Valley Telephone Cooperative Inc / VTX1 SHAKEN 2159 | 01&#160;Mar&#160;25&#160;17:52&#160;UTC | true | [view](CERTS/f0883d130f625c1b8157e54e090e4df461fd9afc05aff09d5a08cf50f9cd37aa/README.md) |
| 05&#160;May&#160;22&#160;14:14&#160;UTC | Kaplan Telephone SHAKEN cert 0432 | 04&#160;May&#160;25&#160;14:14&#160;UTC | true | [view](CERTS/24c1b7c4dc4aeda21b5d1c5ff7f059903693c4a8b67ecc054ef05542a2ff4c35/README.md) |
| 07&#160;Jun&#160;22&#160;12:24&#160;UTC | Avid Communication SHAKEN Cert 742D | 06&#160;Jun&#160;25&#160;12:24&#160;UTC | true | [view](CERTS/b63d54026dfcdfd16495ad6fdda8993de182c86b4aa870784177c38c53842cba/README.md) |
| 02&#160;Aug&#160;22&#160;10:43&#160;UTC | Encore Communications SHAKEN Cert 956H | 01&#160;Aug&#160;25&#160;10:43&#160;UTC | true | [view](CERTS/cc63709890efdf1d938ab01b88b7d4f8368361ffd8572f46cd707e5b3807c95f/README.md) |
| 18&#160;Aug&#160;22&#160;10:39&#160;UTC | Pembroke Telephone Company, Inc SHAKEN Cert 0376 | 17&#160;Aug&#160;25&#160;10:39&#160;UTC | true | [view](CERTS/69ab198b52280bf3d729e1427f62d521e91796e2bae19f99fa0ccd8d024b32de/README.md) |
| 26&#160;Oct&#160;22&#160;16:12&#160;UTC | Northland Networks SHAKEN Cert 7556 | 25&#160;Oct&#160;25&#160;16:12&#160;UTC | true | [view](CERTS/7245e57bd5695bfa52c8ab554bf83a3a7dd21a0dacef76a9aee21b55db7f9963/README.md) |
| 02&#160;Dec&#160;22&#160;17:33&#160;UTC | Ritter Communications SHAKEN cert 095A | 01&#160;Dec&#160;25&#160;17:33&#160;UTC | true | [view](CERTS/28f471821d4f6e3c04d12ffc72c29afe143a2998efbd6fb4f4f55a42d8d594c2/README.md) |
| 22&#160;Feb&#160;23&#160;20:45&#160;UTC | GVTC SHAKEN Cert 2083 | 21&#160;Feb&#160;26&#160;20:45&#160;UTC | true | [view](CERTS/7d497f315d42dc54f854e8711b1873ac16c8966916434b58c8c02c003fb0c1f5/README.md) |
| 02&#160;Mar&#160;23&#160;17:57&#160;UTC | Summit Broadband SHAKEN Cert 7857 | 01&#160;Mar&#160;26&#160;17:57&#160;UTC | true | [view](CERTS/ed3e9cefca8e2fd0c60362a76618f89faafe7d12ea8979620b402eb887a53531/README.md) |
| 14&#160;Mar&#160;23&#160;23:37&#160;UTC | Consolidated Telcom ND SHAKEN Cert 7008 | 13&#160;Mar&#160;26&#160;23:37&#160;UTC | true | [view](CERTS/47618375275c0752509a0796d400c1a874033575cde6333b6f8dc730d496a253/README.md) |
| 26&#160;Mar&#160;23&#160;22:56&#160;UTC | Eatel SHAKEN Cert 8839 | 25&#160;Mar&#160;26&#160;22:56&#160;UTC | true | [view](CERTS/a28d2c246508e2e6cff13fe388c7553b5490afb448d03437142aaea9ddc8d439/README.md) |
| 29&#160;Mar&#160;23&#160;18:43&#160;UTC | Highland Telephone Cooperative SHAKEN Cert 0565 | 28&#160;Mar&#160;26&#160;18:43&#160;UTC | true | [view](CERTS/5d5c734ab29a426bd302cc7389cbad919187cdea1e6eb332af380dc9e1990045/README.md) |
| 31&#160;Mar&#160;23&#160;16:08&#160;UTC | Twin Lakes SHAKEN Cert 0579 | 30&#160;Mar&#160;26&#160;16:08&#160;UTC | true | [view](CERTS/863baf457cc2e6cbf29ba1bee6b324fb5546c206a8f279094b4ab189ec2a1b5f/README.md) |
| 06&#160;Apr&#160;23&#160;10:18&#160;UTC | West Kentucky Rural SHAKEN Cert 0421 | 05&#160;Apr&#160;26&#160;10:18&#160;UTC | true | [view](CERTS/62a29b3d0b7af5b22427dc5dee9d649652664d6a455a601ed01b85592ad9a0a6/README.md) |
| 19&#160;Apr&#160;23&#160;14:08&#160;UTC | Aeneas Communications SHAKEN Cert 2891 | 18&#160;Apr&#160;26&#160;14:08&#160;UTC | true | [view](CERTS/b93371d6de0afcb0e63f67eeca399e4ab95836083e70c573f10ace5c3c9d49de/README.md) |
| 19&#160;Apr&#160;23&#160;14:19&#160;UTC | Blackfoot Communications SHAKEN Cert 2235 | 18&#160;Apr&#160;26&#160;14:19&#160;UTC | true | [view](CERTS/1ae2c0be152e4cef8d41bcfc9809e5d8c71251f93f845e84304ca5a3603019a4/README.md) |
| 26&#160;Apr&#160;23&#160;23:44&#160;UTC | Point Broadband Inc Bristol SHAKEN Cert 9809 | 25&#160;Apr&#160;26&#160;23:44&#160;UTC | true | [view](CERTS/27b46292c626954e89bb6e92e5f752fd88b1efe5a0507ed0d1f72ab9bcdee40e/README.md) |
| 03&#160;May&#160;23&#160;16:53&#160;UTC | 3 Rivers Communications SHAKEN Cert 2255 | 02&#160;May&#160;26&#160;16:53&#160;UTC | true | [view](CERTS/8b53182b31f279a02c6fd03f50762f1d9e3412dc8bd0d2535380627545e28a2c/README.md) |
| 03&#160;May&#160;23&#160;16:55&#160;UTC | Chariton Valley SHAKEN 250A | 02&#160;May&#160;26&#160;16:55&#160;UTC | true | [view](CERTS/377c6fcbcdf110f0b2474415f03d1fb592deb2fcd7394e0404799ede3bb74163/README.md) |
| 05&#160;May&#160;23&#160;15:39&#160;UTC | Plant Telephone Company SHAKEN 0379 | 04&#160;May&#160;26&#160;15:39&#160;UTC | true | [view](CERTS/d3ebd5b2606561fe866139e30c52b829e0bc65c602e77467c8d070987f89df14/README.md) |
| 05&#160;May&#160;23&#160;15:48&#160;UTC | Crosstel Tandem Inc SEPB SHAKEN Cert 357H | 04&#160;May&#160;26&#160;15:48&#160;UTC | true | [view](CERTS/9543f4ce9c1bfae926f9f2103ed6366cb203a6f30586c4b496c24684aa1bec6e/README.md) |
| 09&#160;May&#160;23&#160;13:42&#160;UTC | Crosstel Tandem Inc Holston Shaken Cert 308H  | 08&#160;May&#160;26&#160;13:42&#160;UTC | true | [view](CERTS/c3ca2954b6b625e08636377012a5b92ef553eb182f77872b9a072b8644fe4096/README.md) |
| 12&#160;May&#160;23&#160;11:11&#160;UTC | Fastwyre Broadband SHAKEN Cert 0425 | 11&#160;May&#160;26&#160;11:11&#160;UTC | true | [view](CERTS/2da263aa90b8ba1747cc000ba40674f5d534903d8ee0c5cd38ea276bc95f02f3/README.md) |
| 12&#160;May&#160;23&#160;11:17&#160;UTC | Farmers Telecommunications Inc SHAKEN Cert 2188 | 11&#160;May&#160;26&#160;11:17&#160;UTC | true | [view](CERTS/4fd081c1f3f8f1b17d92df3b9eff24b52a2adcb75df8e8292defbf2cac0e035d/README.md) |
| 25&#160;May&#160;23&#160;22:17&#160;UTC | Alma Tel SHAKEN 0344 | 24&#160;May&#160;26&#160;22:17&#160;UTC | true | [view](CERTS/1008e70b562dafabfb601eb7e7706c597172afcd0728a9e7602454dd0f365a7b/README.md) |
| 08&#160;Jun&#160;23&#160;16:02&#160;UTC | Whidbey Telephone Company SHAKEN Cert 2452 | 07&#160;Jun&#160;26&#160;16:02&#160;UTC | true | [view](CERTS/11a8d77aa46349b365d5740e537d242f906b2729b3a16dcc5008297a8a7e3a3c/README.md) |
| 19&#160;Jun&#160;23&#160;09:22&#160;UTC | EPB Telecom SHAKEN Cert 4645 | 18&#160;Jun&#160;26&#160;09:22&#160;UTC | true | [view](CERTS/579c03c51c504eaf40865095ffa96130d4c195391294d6933ffe58dc3d673ece/README.md) |
| 20&#160;Jun&#160;23&#160;09:29&#160;UTC | Ben Lomand SHAKEN Cert 0553 | 19&#160;Jun&#160;26&#160;09:29&#160;UTC | true | [view](CERTS/ce469ff39f245721d8c523af0a6fbd26fc0742e3c18b0b5adc627794eaad6231/README.md) |
| 22&#160;Jun&#160;23&#160;10:20&#160;UTC | CAS Cable SHAKEN Cert 875F | 21&#160;Jun&#160;26&#160;10:20&#160;UTC | true | [view](CERTS/62fdef45ba8997cccb0bec68cb10afee549d72357286ccacfaa94f4040a9ed42/README.md) |
| 05&#160;Oct&#160;23&#160;13:44&#160;UTC | Fidelity Communications SHAKEN Cert 1882 | 04&#160;Oct&#160;26&#160;13:44&#160;UTC | true | [view](CERTS/af56b81c67b9b517530e053f153af3c7d6a72620c04f84b1ff648e7c137e1376/README.md) |

#### CA Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 25&#160;Nov&#160;20&#160;11:21&#160;UTC | Metaswitch STI-CA SHAKEN Root | 20&#160;Nov&#160;40&#160;11:21&#160;UTC | true | [view](CERTS/c27184f75f81fc119b85d51d416477bf635040e5d66ce6051799b37b9aa17485/README.md) |
| 25&#160;Nov&#160;20&#160;11:57&#160;UTC | Metaswitch STI-CA SHAKEN Issuing 1 | 22&#160;Nov&#160;32&#160;11:57&#160;UTC | true | [view](CERTS/b91a9874fbefc3feda9d5f9bd336e8b999c9b15b25aae7fe3c61d87373a5d1a1/README.md) |
| 10&#160;Feb&#160;23&#160;14:38&#160;UTC | Metaswitch STI-CA SHAKEN Issuing 1 | 07&#160;Feb&#160;35&#160;14:38&#160;UTC | true | [view](CERTS/8a7fb50e95b8c43a63d19e2f279de565fa611ae3f24a14f82394e3208782be7a/README.md) |


Generated: 15 Nov 23 17:17 UTC