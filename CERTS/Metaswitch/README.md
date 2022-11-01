# STIR/SHAKEN CA Ecosystem Compliance

## Metaswitch

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 21 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 0 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 21 certificates being tested against the remaining rules
- 100.00% of certificates contain one or more Error level issue
- 28.57% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 95.24% of certificates are too old to be assessed against currently enforced expectations
- 1067 days is the average remaining validity for the certificates in the corpus
- 1095 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days

| Instances | Test Status | Source |
|-----------|-------------|--------|
| 20 | [e_atis_extension_unknown](ISSUES/e_atis_extension_unknown/README.md) | ATIS1000080 |
| 21 | [e_atis_issuer_dn](ISSUES/e_atis_issuer_dn/README.md) | ATIS1000080 |
| 21 | [e_atis_key_usage](ISSUES/e_atis_key_usage/README.md) | ATIS1000080 |
| 20 | [e_atis_serial_number](ISSUES/e_atis_serial_number/README.md) | ATIS1000080 |
| 20 | [e_atis_signature_algorithm](ISSUES/e_atis_signature_algorithm/README.md) | ATIS1000080 |
| 1 | [e_atis_subject](ISSUES/e_atis_subject/README.md) | ATIS1000080 |
| 21 | [e_atis_subject_cn](ISSUES/e_atis_subject_cn/README.md) | ATIS1000080 |
| 21 | [e_us_cp_ambiguous_identifier](ISSUES/e_us_cp_ambiguous_identifier/README.md) | US_SHAKEN_CP |
| 21 | [e_us_cp_subject_sn](ISSUES/e_us_cp_subject_sn/README.md) | US_SHAKEN_CP |
| 6 | [w_shaken_subject_rdn_unknown](ISSUES/w_shaken_subject_rdn_unknown/README.md) | SHAKEN_PKI_BEST_PRACTICES |

#### CA Certificates

- 2 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 0 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 2 certificates being tested against the remaining rules
- 100.00% of certificates contain one or more Error level issue
- 0.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 100.00% of certificates are too old to be assessed against currently enforced expectations
- 6091 days is the average remaining validity for the certificates in the corpus
- 5840 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days

| Instances | Test Status | Source |
|-----------|-------------|--------|
| 1 | [e_atis_ca_certificate_policies](ISSUES/e_atis_ca_certificate_policies/README.md) | ATIS1000080 |
| 1 | [e_atis_ca_crl_distribution](ISSUES/e_atis_ca_crl_distribution/README.md) | ATIS1000080 |
| 1 | [e_atis_ca_extension_unknown](ISSUES/e_atis_ca_extension_unknown/README.md) | ATIS1000080 |
| 2 | [e_atis_ca_issuer_dn](ISSUES/e_atis_ca_issuer_dn/README.md) | ATIS1000080 |
| 2 | [e_atis_ca_serial_number](ISSUES/e_atis_ca_serial_number/README.md) | ATIS1000080 |
| 2 | [e_atis_ca_subject](ISSUES/e_atis_ca_subject/README.md) | ATIS1000080 |
| 2 | [e_atis_ca_subject_cn](ISSUES/e_atis_ca_subject_cn/README.md) | ATIS1000080 |
| 1 | [e_atis_root_certificate_policies](ISSUES/e_atis_root_certificate_policies/README.md) | ATIS1000080 |
| 1 | [e_atis_root_extension_unknown](ISSUES/e_atis_root_extension_unknown/README.md) | ATIS1000080 |

### Details

#### Leaf Certificates

| Created at | Subject | Problems | Link |
|------------|---------|----------|------|
| 17 Dec 20 15:19 UTC | TDS Telecom SHAKEN Cert 7804 | true | [view](CERTS/a04a669738b79ff55c2b2197f72a12a112b731906e2e6a925d37ccee2fa00a11/README.md) |
| 07 Jan 21 16:17 UTC | Call48 SHAKEN Cert 505J | true | [view](CERTS/5bb4516a62167e6c55e9704e4a49b39ce9ec480e678aa6c02fb6985fc9594997/README.md) |
| 27 Jan 21 17:16 UTC | Telesystem SHAKEN Cert 786E | true | [view](CERTS/2d9aca0895c94291596161363091718089a6e7c19dfa57329ae548432533860f/README.md) |
| 16 Feb 21 17:30 UTC | Verizon SHAKEN cert 5807 | true | [view](CERTS/d7b413267be2d050d516af4f4a864ffdc2feacc774a1a6264b9cfe68c796f43f/README.md) |
| 10 Mar 21 20:40 UTC | Allstream SHAKEN cert 4130 | true | [view](CERTS/a9eec584cf66eade89fa84e043e0eb05900dc40653ea963a96c079ef92d9e2ad/README.md) |
| 10 Mar 21 20:50 UTC | Northeast Communications of Wisconsin SHAKEN Cert 6692 | true | [view](CERTS/90042ab31d5de7abb64b89073a0d931ce22d864b89df13a9372887cb5db45f49/README.md) |
| 17 Mar 21 17:06 UTC | RCN SHAKEN Cert 7615 | true | [view](CERTS/bbdec20ad80f4a2a8ed097204a9299566beca170460fb648c81a51d195d9b6f1/README.md) |
| 17 Mar 21 17:12 UTC | GCI SHAKEN Cert 7785 | true | [view](CERTS/312e58dffa682b464f9867a7c373f9881d092b834767dcabe5baf8c7245e937c/README.md) |
| 17 Mar 21 17:19 UTC | South Central Rural Telecommunications Coop SHAKEN Cert 0418 | true | [view](CERTS/6e988e3a74dbb6c1862118841d78e438f345c1ef7e79a96c2087328b5de146c1/README.md) |
| 24 Mar 21 23:52 UTC | Cspire SHAKEN Cert 6581 | true | [view](CERTS/09ed5b3292b5bfc7ac80b1027a827138b9503aa8053a61431a8dc851ecad04f2/README.md) |
| 01 Apr 21 15:28 UTC | USCellular SHAKEN Cert 6349 | true | [view](CERTS/c75937e7c843b0054b1a51f6149f47661696b3cfcbcf6bab4dbc708c4453244f/README.md) |
| 07 Apr 21 16:44 UTC | U. S. Telepacific Corp SHAKEN 7453 | true | [view](CERTS/9ed03dac797a5a27d52aa5209a4caa6a3ec9c3943d55a2cbfb69416480787da0/README.md) |
| 16 Apr 21 15:27 UTC | CBTS Technology Solutions SHAKEN Cert 600F | true | [view](CERTS/3d02021a2da14f1ebfe588256a419be9ebc03c0d1fccc51cc29fa9d4a625c6bf/README.md) |
| 18 May 21 13:14 UTC | Segra SHAKEN Cert 1784 | true | [view](CERTS/f802b8b879d063b87665d8b2a67f6d3ba78a94aa35782f664a6a40afc7586f56/README.md) |
| 20 May 21 22:04 UTC | Appalachian Wireless SHAKEN Cert 6940 | true | [view](CERTS/e14170c681e75c37d0ca45e304c09cc0d148246bd7d72e96f91f7a8fe27339fa/README.md) |
| 20 May 21 22:20 UTC | Everstream SHAKEN Cert 472C  | true | [view](CERTS/8710bb38debebd39698fb1c273409b173951cca1fab53a6d4c4aca91e61e06df/README.md) |
| 26 May 21 15:37 UTC | WOW Internet Cable and Phone SHAKEN Cert 665E | true | [view](CERTS/9f86e038b4f6c3bf7fabfe8d4008b071f81bd33a80c55ea66a1aadd79eb9b989/README.md) |
| 27 May 21 13:28 UTC | Mediacom 846F | true | [view](CERTS/e6c9e9fd411d8174b3ffe1af4d569c6919f4b98a5d0c6e429cd3682d82284e7e/README.md) |
| 30 Jun 21 16:55 UTC | Sonic Telecom SHAKEN cert 433E | true | [view](CERTS/a5082b808b3bf200a634e273b7988a617099c9c897a9c596858e6d4ffe4fd352/README.md) |
| 20 Jul 21 19:54 UTC | New Horizon SHAKEN Cert 127E | true | [view](CERTS/5b2842e49ecc543187018171fa660e32bdba390a7977c0de97a00aef35b8ae01/README.md) |
| 07 Jun 22 12:24 UTC | Avid Communication SHAKEN Cert 742D | true | [view](CERTS/b63d54026dfcdfd16495ad6fdda8993de182c86b4aa870784177c38c53842cba/README.md) |

#### CA Certificates

| Created at | Subject | Problems | Link |
|------------|---------|----------|------|
| 25 Nov 20 11:21 UTC | Metaswitch STI-CA SHAKEN Root | true | [view](CERTS/c27184f75f81fc119b85d51d416477bf635040e5d66ce6051799b37b9aa17485/README.md) |
| 25 Nov 20 11:57 UTC | Metaswitch STI-CA SHAKEN Issuing 1 | true | [view](CERTS/b91a9874fbefc3feda9d5f9bd336e8b999c9b15b25aae7fe3c61d87373a5d1a1/README.md) |


Generated: 01/11/2022 at 20:31:14