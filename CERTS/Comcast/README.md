# STIR/SHAKEN CA Ecosystem Compliance

## Comcast

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 189 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 134 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 55 certificates being tested against the remaining rules
- 6.00 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 100.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 30 days is the average remaining validity for the certificates in the corpus
- 30 days is the average initial validity for the certificates in the corpus
- 55 certificates expire in the next 30 days
- 55.00 average number of unexpired certificates per OCN observed
- 1 unique OCNs observed in unexpired and valid certificate corpus

| Instances | Test | Source |
|-----------|------|--------|
| 55 | [e_atis_certificate_policies](ISSUES/e_atis_certificate_policies/README.md) | ATIS1000080 |
| 55 | [e_atis_subject_cn](ISSUES/e_atis_subject_cn/README.md) | ATIS1000080 |
| 55 | [e_atis_subject_key_identifier](ISSUES/e_atis_subject_key_identifier/README.md) | ATIS1000080 |
| 55 | [e_us_cp_ambiguous_identifier](ISSUES/e_us_cp_ambiguous_identifier/README.md) | US_SHAKEN_CP |
| 55 | [e_us_cp_subject_sn](ISSUES/e_us_cp_subject_sn/README.md) | US_SHAKEN_CP |
| 55 | [w_ext_subject_key_identifier_missing_sub_cert](ISSUES/w_ext_subject_key_identifier_missing_sub_cert/README.md) | RFC5280 |

#### CA Certificates

- 2 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 0 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 2 certificates being tested against the remaining rules
- 0.00 issues on average found in unexpired, trusted, and non-compliant certificates
- 0.00% of certificates contain one or more Error level issue
- 0.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 100.00% of certificates are too old to be assessed against currently enforced expectations
- 6798 days is the average remaining validity for the certificates in the corpus
- 7150 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days

No error, warning, or notice level issues were found

### Details

#### Leaf Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 07&#160;Jun&#160;23&#160;12:30&#160;UTC | SHAKEN | 07&#160;Jul&#160;23&#160;12:30&#160;UTC | true | [view](CERTS/f98847bdfc80c8939d25f8cda2b3f29386c3d4c3975dc6bf0fe5071fc349bdab/README.md) |
| 08&#160;Jun&#160;23&#160;10:11&#160;UTC | SHAKEN | 08&#160;Jul&#160;23&#160;10:11&#160;UTC | true | [view](CERTS/b681018f43ba8be51bbacf3ce029962e5fe262b493a2aba43e5e439e6495c64d/README.md) |
| 08&#160;Jun&#160;23&#160;12:30&#160;UTC | SHAKEN | 08&#160;Jul&#160;23&#160;12:30&#160;UTC | true | [view](CERTS/a2c17ee56ad56204516728c36927857347cf16d2bff0ebd18edc6edb3d713473/README.md) |
| 09&#160;Jun&#160;23&#160;10:11&#160;UTC | SHAKEN | 09&#160;Jul&#160;23&#160;10:11&#160;UTC | true | [view](CERTS/778fd42ab6c2cde00492f75bed9eb9a166c50011dabe2632e5ee33f478b2739a/README.md) |
| 09&#160;Jun&#160;23&#160;12:30&#160;UTC | SHAKEN | 09&#160;Jul&#160;23&#160;12:30&#160;UTC | true | [view](CERTS/c3d11f8a9fd244d619c20f2a0548b9c48c692752bd78e80bc5fddcefdfe8c6b7/README.md) |
| 09&#160;Jun&#160;23&#160;13:49&#160;UTC | SHAKEN | 09&#160;Jul&#160;23&#160;13:49&#160;UTC | true | [view](CERTS/bb88c3b2afc7cb38e46aa9a216cea25788ac797db66eb44651e3684987819f36/README.md) |
| 09&#160;Jun&#160;23&#160;13:54&#160;UTC | SHAKEN | 09&#160;Jul&#160;23&#160;13:54&#160;UTC | true | [view](CERTS/aa88db360da1c026e71032f0e1d45582482fa19c7a8c559f040d5ac728bb28f1/README.md) |
| 10&#160;Jun&#160;23&#160;10:11&#160;UTC | SHAKEN | 10&#160;Jul&#160;23&#160;10:11&#160;UTC | true | [view](CERTS/6f773bfd32a8782fb6678cc93f9436b670ad61e17c0caadad4ba23a53f533342/README.md) |
| 10&#160;Jun&#160;23&#160;13:54&#160;UTC | SHAKEN | 10&#160;Jul&#160;23&#160;13:54&#160;UTC | true | [view](CERTS/842eaf12de7a18d270fcbab90d61c3b5945bbad9d0f19432793ed5ebe21abd3f/README.md) |
| 11&#160;Jun&#160;23&#160;13:55&#160;UTC | SHAKEN | 11&#160;Jul&#160;23&#160;13:55&#160;UTC | true | [view](CERTS/5e878af36e1aaf0317ad6edac899f0394547b3d5d69ce1a7bf20647f7688af22/README.md) |
| 12&#160;Jun&#160;23&#160;10:11&#160;UTC | SHAKEN | 12&#160;Jul&#160;23&#160;10:11&#160;UTC | true | [view](CERTS/982a5d3a9735195059004475b9e25f9e280984be1fb643551a7c014f925aa60b/README.md) |
| 12&#160;Jun&#160;23&#160;13:54&#160;UTC | SHAKEN | 12&#160;Jul&#160;23&#160;13:54&#160;UTC | true | [view](CERTS/c785e14d91cac7e4eb0f3260aa62077c6cbcfcb6ae520474206936bf57179478/README.md) |
| 13&#160;Jun&#160;23&#160;10:11&#160;UTC | SHAKEN | 13&#160;Jul&#160;23&#160;10:11&#160;UTC | true | [view](CERTS/996c378d5408fa98adf877c6b3d5aa3954539640246bf498bb25076d23a5efbc/README.md) |
| 13&#160;Jun&#160;23&#160;13:55&#160;UTC | SHAKEN | 13&#160;Jul&#160;23&#160;13:55&#160;UTC | true | [view](CERTS/6b4c19d53de9722d64b9ce61b56cdabca39ceb862c143ae4838253848eb3bc31/README.md) |
| 14&#160;Jun&#160;23&#160;10:11&#160;UTC | SHAKEN | 14&#160;Jul&#160;23&#160;10:11&#160;UTC | true | [view](CERTS/d29fdfba6b5a4a438ab9c03d042827cd5cbf1ab04d14b81c2716a96ed13bc527/README.md) |
| 14&#160;Jun&#160;23&#160;13:54&#160;UTC | SHAKEN | 14&#160;Jul&#160;23&#160;13:54&#160;UTC | true | [view](CERTS/f1c2c3cb5578c866e664402833a2138f87a837fa5a38c9f799939bc3cd39495a/README.md) |
| 15&#160;Jun&#160;23&#160;10:11&#160;UTC | SHAKEN | 15&#160;Jul&#160;23&#160;10:11&#160;UTC | true | [view](CERTS/e491fb5bdbe8bbfd4b9f6d7be6b9f1bd9b0eea1558b5a08f3a3108bc398e04e1/README.md) |
| 15&#160;Jun&#160;23&#160;13:55&#160;UTC | SHAKEN | 15&#160;Jul&#160;23&#160;13:55&#160;UTC | true | [view](CERTS/657204621eb39f3631c549faef27c6d92a261926e45cc3c7a069ef32a21b1078/README.md) |
| 16&#160;Jun&#160;23&#160;10:11&#160;UTC | SHAKEN | 16&#160;Jul&#160;23&#160;10:11&#160;UTC | true | [view](CERTS/4375160ac74531af55a5bed514e400b8c3b3e59c47180a566bfa12618b5ff184/README.md) |
| 16&#160;Jun&#160;23&#160;13:54&#160;UTC | SHAKEN | 16&#160;Jul&#160;23&#160;13:54&#160;UTC | true | [view](CERTS/11d355df337c41188bd716a830725f73157ee4867da4d9017c00ac79f23b203f/README.md) |
| 17&#160;Jun&#160;23&#160;10:11&#160;UTC | SHAKEN | 17&#160;Jul&#160;23&#160;10:11&#160;UTC | true | [view](CERTS/44fe2e8d6f22d6431097fec836f3d79f1d0acd857f91b31f29f6d0a6d2e1cb5c/README.md) |
| 17&#160;Jun&#160;23&#160;13:55&#160;UTC | SHAKEN | 17&#160;Jul&#160;23&#160;13:55&#160;UTC | true | [view](CERTS/8d0daaf246e6b39a011404495c53b39326ae2a4cb9be702901f178bd08d869de/README.md) |
| 18&#160;Jun&#160;23&#160;13:54&#160;UTC | SHAKEN | 18&#160;Jul&#160;23&#160;13:54&#160;UTC | true | [view](CERTS/bef83720d8f03273e570e9aacbb10de234e21d483ec46264b2af92ce9c5e73c4/README.md) |
| 19&#160;Jun&#160;23&#160;10:11&#160;UTC | SHAKEN | 19&#160;Jul&#160;23&#160;10:11&#160;UTC | true | [view](CERTS/94be8955234597d7336bf7cdf88651a1f0af24fd285bc6b4384f301778829848/README.md) |
| 19&#160;Jun&#160;23&#160;13:55&#160;UTC | SHAKEN | 19&#160;Jul&#160;23&#160;13:55&#160;UTC | true | [view](CERTS/e416ad2c14d8a340b2e283fc3b1e21134f67b191ba7e35c278cc3e67c36c08b3/README.md) |
| 20&#160;Jun&#160;23&#160;10:11&#160;UTC | SHAKEN | 20&#160;Jul&#160;23&#160;10:11&#160;UTC | true | [view](CERTS/350d11cdd0abb3352e4ec6a645c0ee8b622849ec8f50d7e16572039a826fca0f/README.md) |
| 20&#160;Jun&#160;23&#160;13:54&#160;UTC | SHAKEN | 20&#160;Jul&#160;23&#160;13:54&#160;UTC | true | [view](CERTS/ec8a4d6e7069b4896ddb619be5cc1b1ced0c5530c17f6bac94263367831b63ad/README.md) |
| 21&#160;Jun&#160;23&#160;10:11&#160;UTC | SHAKEN | 21&#160;Jul&#160;23&#160;10:11&#160;UTC | true | [view](CERTS/913d1535bbab951708986be551ea0dd362b43c621b3ce2fac8a7e7a100de2aaf/README.md) |
| 21&#160;Jun&#160;23&#160;13:54&#160;UTC | SHAKEN | 21&#160;Jul&#160;23&#160;13:54&#160;UTC | true | [view](CERTS/4391696e31b5159da35fa696d235e3e6a8e96685de9ed7a2c5ec1b81c5104dad/README.md) |
| 22&#160;Jun&#160;23&#160;10:11&#160;UTC | SHAKEN | 22&#160;Jul&#160;23&#160;10:11&#160;UTC | true | [view](CERTS/8eb372fcffbc57100b3b74a04b53bf51b3a179a10404e4d1848868445b22886d/README.md) |
| 22&#160;Jun&#160;23&#160;13:54&#160;UTC | SHAKEN | 22&#160;Jul&#160;23&#160;13:54&#160;UTC | true | [view](CERTS/e7b22b6c0b1e66c5926eb12ff78a5f0f58f666d35e090b470e849731bdcfb8e7/README.md) |
| 23&#160;Jun&#160;23&#160;10:11&#160;UTC | SHAKEN | 23&#160;Jul&#160;23&#160;10:11&#160;UTC | true | [view](CERTS/6ef38a6e4fdf005f08efe5872b2737683b6087465b8e1d25aa162812769b9697/README.md) |
| 23&#160;Jun&#160;23&#160;13:54&#160;UTC | SHAKEN | 23&#160;Jul&#160;23&#160;13:54&#160;UTC | true | [view](CERTS/387690827af3148409132f0bc1441aadf79ab67be4be5c33a7f16d18d9692fc0/README.md) |
| 24&#160;Jun&#160;23&#160;10:11&#160;UTC | SHAKEN | 24&#160;Jul&#160;23&#160;10:11&#160;UTC | true | [view](CERTS/f1e7fae6e562b50f619d70ff1c91b4194d513ae6177a84c43916616dd7e4bd5c/README.md) |
| 24&#160;Jun&#160;23&#160;13:54&#160;UTC | SHAKEN | 24&#160;Jul&#160;23&#160;13:54&#160;UTC | true | [view](CERTS/dd81008792338cb0e9200cb198cf3f565d348d8420a6ef3413d918bffc230148/README.md) |
| 25&#160;Jun&#160;23&#160;13:55&#160;UTC | SHAKEN | 25&#160;Jul&#160;23&#160;13:55&#160;UTC | true | [view](CERTS/eee8e9f1d84d8d45b3d21acae8c2e49b99deb6a22020a3c2a30c8f00699d28f0/README.md) |
| 26&#160;Jun&#160;23&#160;10:11&#160;UTC | SHAKEN | 26&#160;Jul&#160;23&#160;10:11&#160;UTC | true | [view](CERTS/cdd133f785d488d39ecb105da97b5e6415056c010e45b2c18c1adef410fee8f1/README.md) |
| 26&#160;Jun&#160;23&#160;13:54&#160;UTC | SHAKEN | 26&#160;Jul&#160;23&#160;13:54&#160;UTC | true | [view](CERTS/3f955219d8c7993d7e7122dba529ae1b49879a7874e3406001a221785eb7c54f/README.md) |
| 27&#160;Jun&#160;23&#160;10:11&#160;UTC | SHAKEN | 27&#160;Jul&#160;23&#160;10:11&#160;UTC | true | [view](CERTS/5b7b51cb79812ead7b4cf025a0bf7578c21e9894f715f2b1e741c6b78434dbac/README.md) |
| 27&#160;Jun&#160;23&#160;13:54&#160;UTC | SHAKEN | 27&#160;Jul&#160;23&#160;13:54&#160;UTC | true | [view](CERTS/3b2fd37712a4532accaff64811c67e947807bc6d5258fe6b98e25ba9c7e26107/README.md) |
| 28&#160;Jun&#160;23&#160;10:11&#160;UTC | SHAKEN | 28&#160;Jul&#160;23&#160;10:11&#160;UTC | true | [view](CERTS/3a07b64edfe6ef58b83f1f9d14702e89a653481562e63604282029c6a6a4b45f/README.md) |
| 28&#160;Jun&#160;23&#160;13:54&#160;UTC | SHAKEN | 28&#160;Jul&#160;23&#160;13:54&#160;UTC | true | [view](CERTS/da2e11a479178a5afe1ebe39513bf80345d7b6c8ada1950198df000adce45191/README.md) |
| 29&#160;Jun&#160;23&#160;10:11&#160;UTC | SHAKEN | 29&#160;Jul&#160;23&#160;10:11&#160;UTC | true | [view](CERTS/0e6aa72c7c4278945e9724af770c2d72515fde3a72a9055d2a6252613de6143b/README.md) |
| 29&#160;Jun&#160;23&#160;13:55&#160;UTC | SHAKEN | 29&#160;Jul&#160;23&#160;13:55&#160;UTC | true | [view](CERTS/da48a75d7461f942f43871068ba04493cb2c2e527b7134788fec4a72e08ccbb7/README.md) |
| 30&#160;Jun&#160;23&#160;10:11&#160;UTC | SHAKEN | 30&#160;Jul&#160;23&#160;10:11&#160;UTC | true | [view](CERTS/1ef4773a6d2bf8f05f954ef9c705c6ac2a1324d332b9478d83eafe7a3d8ca4b2/README.md) |
| 30&#160;Jun&#160;23&#160;13:54&#160;UTC | SHAKEN | 30&#160;Jul&#160;23&#160;13:54&#160;UTC | true | [view](CERTS/3f374b17894a63881e489dbd74e99ab79c7d78cfd0f45abced8d10d461c32179/README.md) |
| 01&#160;Jul&#160;23&#160;10:11&#160;UTC | SHAKEN | 31&#160;Jul&#160;23&#160;10:11&#160;UTC | true | [view](CERTS/ee436e030b4bbc72becaa7eb184272cb4a12e80d09814674ec0c999b52f6de2b/README.md) |
| 01&#160;Jul&#160;23&#160;13:55&#160;UTC | SHAKEN | 31&#160;Jul&#160;23&#160;13:55&#160;UTC | true | [view](CERTS/5e518fca1142004c3989ecc8b4d3b0d1be4ad68cc6da97211d2fca82f932b827/README.md) |
| 02&#160;Jul&#160;23&#160;13:54&#160;UTC | SHAKEN | 01&#160;Aug&#160;23&#160;13:54&#160;UTC | true | [view](CERTS/898ba9e516ff50145b4c539c5a136310eea03e62e1283eb7df8ca78f5ac21d5f/README.md) |
| 03&#160;Jul&#160;23&#160;10:11&#160;UTC | SHAKEN | 02&#160;Aug&#160;23&#160;10:11&#160;UTC | true | [view](CERTS/120cd66b7093d397996b7bf8127630fb24c916354f55283a26a95f6e0d3a9e72/README.md) |
| 03&#160;Jul&#160;23&#160;13:55&#160;UTC | SHAKEN | 02&#160;Aug&#160;23&#160;13:55&#160;UTC | true | [view](CERTS/604f8d9a8a9e1e9fee706ebf0391c6f1cff66ec3f96074c96aa692bf66d749a9/README.md) |
| 04&#160;Jul&#160;23&#160;10:11&#160;UTC | SHAKEN | 03&#160;Aug&#160;23&#160;10:11&#160;UTC | true | [view](CERTS/755dd72f075dced39e6e38b422b344cd291a6ea6beb29b6d8eff4fe406f022db/README.md) |
| 04&#160;Jul&#160;23&#160;13:54&#160;UTC | SHAKEN | 03&#160;Aug&#160;23&#160;13:54&#160;UTC | true | [view](CERTS/a69797c5c67cd24c6027de0622ef0e57501b2f8f810caae802d5396d3e26c808/README.md) |
| 05&#160;Jul&#160;23&#160;10:11&#160;UTC | SHAKEN | 04&#160;Aug&#160;23&#160;10:11&#160;UTC | true | [view](CERTS/d4e3c90168de183d32f909e42c04ec80ef632958399ebf79e2a1a5c0de86186a/README.md) |
| 05&#160;Jul&#160;23&#160;13:55&#160;UTC | SHAKEN | 04&#160;Aug&#160;23&#160;13:55&#160;UTC | true | [view](CERTS/a02cc610fb352d722462de03e4e5068376205a05c887f25ac0439f1e0ce38baa/README.md) |

#### CA Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 17&#160;Mar&#160;20&#160;14:45&#160;UTC | Comcast SHAKEN Root CA | 12&#160;Mar&#160;40&#160;14:45&#160;UTC | false | [view](CERTS/b1132c5f12c3ca4d2ff119f2df99544336eb1703512ac99cc42d596e25125bbd/README.md) |
| 06&#160;Apr&#160;20&#160;13:48&#160;UTC | Comcast SHAKEN Intermediate CA | 06&#160;Jun&#160;39&#160;13:48&#160;UTC | false | [view](CERTS/2f3abdfe711377f4d59f31d941962797c61f021c3924af31de99031f9ac54f77/README.md) |


Generated: 06 Jul 23 14:08 UTC