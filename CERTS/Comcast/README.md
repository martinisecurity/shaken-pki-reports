# STIR/SHAKEN CA Ecosystem Compliance

## Comcast

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 53 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 10 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 43 certificates being tested against the remaining rules
- 6.02 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 100.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 29 days is the average remaining validity for the certificates in the corpus
- 30 days is the average initial validity for the certificates in the corpus
- 43 certificates expire in the next 30 days

| Instances | Test | Source |
|-----------|------|--------|
| 43 | [e_atis_certificate_policies](ISSUES/e_atis_certificate_policies/README.md) | ATIS1000080 |
| 1 | [e_atis_serial_number](ISSUES/e_atis_serial_number/README.md) | ATIS1000080 |
| 43 | [e_atis_subject_cn](ISSUES/e_atis_subject_cn/README.md) | ATIS1000080 |
| 43 | [e_atis_subject_key_identifier](ISSUES/e_atis_subject_key_identifier/README.md) | ATIS1000080 |
| 43 | [e_us_cp_ambiguous_identifier](ISSUES/e_us_cp_ambiguous_identifier/README.md) | US_SHAKEN_CP |
| 43 | [e_us_cp_subject_sn](ISSUES/e_us_cp_subject_sn/README.md) | US_SHAKEN_CP |
| 43 | [w_ext_subject_key_identifier_missing_sub_cert](ISSUES/w_ext_subject_key_identifier_missing_sub_cert/README.md) | RFC5280 |

#### CA Certificates

- 2 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 0 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 2 certificates being tested against the remaining rules
- 0.00 issues on average found in unexpired, trusted, and non-compliant certificates
- 0.00% of certificates contain one or more Error level issue
- 0.00% of certificates contain one or more Warning level issue
- 100.00% of certificates contain one or more Notice level issue
- 100.00% of certificates are too old to be assessed against currently enforced expectations
- 6880 days is the average remaining validity for the certificates in the corpus
- 7150 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days

No error, warning, or notice level issues were found

### Details

#### Leaf Certificates

| Created at | Subject | Problems | Link |
|------------|---------|----------|------|
| 04 Oct 22 12:46 UTC | SHAKEN | true | [view](CERTS/baf4d6d5e2605e4bd43f049b12d1dc66a533ce0d58a07001734d835f633e3602/README.md) |
| 04 Oct 22 13:29 UTC | SHAKEN | true | [view](CERTS/e12cde948a53eda9a8c6820b3777c71e1b99c3712d80d36bd7f3fd312f76b341/README.md) |
| 05 Oct 22 13:29 UTC | SHAKEN | true | [view](CERTS/7fda3ccce7ec708448edc8f96c95b90bc375152feea1c026234beefc37e81999/README.md) |
| 06 Oct 22 10:05 UTC | SHAKEN | true | [view](CERTS/5c9719585c750cd739181b59450db89581791c376bbfc680ef2ae8fa293d75b6/README.md) |
| 06 Oct 22 12:46 UTC | SHAKEN | true | [view](CERTS/c90c48389bbe12c02325baedc377a16cad73a248cb694897540a40e02b878bc0/README.md) |
| 06 Oct 22 13:29 UTC | SHAKEN | true | [view](CERTS/6073c26c2c63fa39c8e02de6010827d8bd5509e6f110a6218841a19e813e86d5/README.md) |
| 07 Oct 22 12:46 UTC | SHAKEN | true | [view](CERTS/b96dae77063316253cbb42805030185deff677108a05e39b2405255316978fd8/README.md) |
| 07 Oct 22 13:30 UTC | SHAKEN | true | [view](CERTS/e7aad6234dba0ba07582d87288a9444c111ae6ad4451ea85dc26c966d307c02e/README.md) |
| 08 Oct 22 10:05 UTC | SHAKEN | true | [view](CERTS/cea5d6f242a1a4ba4b6cb6508c37e90bca07ab8c9296345cac84052670b45b79/README.md) |
| 08 Oct 22 13:29 UTC | SHAKEN | true | [view](CERTS/b08ca401227d50e851c405102296f6f2476c3583f0e622e94536011f74d5f23d/README.md) |
| 09 Oct 22 12:46 UTC | SHAKEN | true | [view](CERTS/ecadce07ce7ec53f7d525daa5f6379dd1ad5b2605ae82e533a13b1f0c3dba5db/README.md) |
| 10 Oct 22 12:46 UTC | SHAKEN | true | [view](CERTS/2750a59b74cbb4082d06dc5c8a297b7a211f318bef6a3f79a5cb194e8f5547fe/README.md) |
| 11 Oct 22 13:30 UTC | SHAKEN | true | [view](CERTS/b0c45f73954a4e38e300e41c38ce976628c64c2388298b5fdd227bc3d1e0e32e/README.md) |
| 12 Oct 22 12:46 UTC | SHAKEN | true | [view](CERTS/566e9572a4ea7da8554e2e98ef13e3ec5e4cf13d4d58fe14de464e57e470b352/README.md) |
| 12 Oct 22 13:29 UTC | SHAKEN | true | [view](CERTS/953e14280d248a2cfd95a650e81ed9094d23f05f5b410e01f66eb7377cdafded/README.md) |
| 13 Oct 22 12:46 UTC | SHAKEN | true | [view](CERTS/d355843a5156aa30667defed29052b11c79b3acc85123c6938d0475f3827826d/README.md) |
| 14 Oct 22 10:05 UTC | SHAKEN | true | [view](CERTS/5078ad95faed5786174370b1cddba8b278ff75faebecbe014ba3c5f81e49a03c/README.md) |
| 14 Oct 22 12:46 UTC | SHAKEN | true | [view](CERTS/0d091abfa7776698a16efd051c855907fae95fea780154634f1b988fce868b56/README.md) |
| 14 Oct 22 13:30 UTC | SHAKEN | true | [view](CERTS/229186c59640158d1f281af3617371c072f0663799dabb05a8e4d577b11fb83e/README.md) |
| 15 Oct 22 13:29 UTC | SHAKEN | true | [view](CERTS/b017ee25cc6232cd9591bc6df0cedbe2b95465b8123f651f7b23e71bba10d6f8/README.md) |
| 17 Oct 22 10:05 UTC | SHAKEN | true | [view](CERTS/f98af32ca082d74e5675ad0b61d5291332e78ac4d87b0968a41dd13e707d2493/README.md) |
| 17 Oct 22 12:46 UTC | SHAKEN | true | [view](CERTS/107165d254b42351095a42289be894289212c27f44c959eed5bfb209c91de6d0/README.md) |
| 17 Oct 22 13:29 UTC | SHAKEN | true | [view](CERTS/f43c128e22aaa4c560e4ecc67309f7e6c0aef21f62c76faa9d648109f572378c/README.md) |
| 18 Oct 22 12:46 UTC | SHAKEN | true | [view](CERTS/b28bbfa48eeb7401282c0af83ad52bf3214c96f77e791ee9d3c02a53ea25cc02/README.md) |
| 18 Oct 22 13:30 UTC | SHAKEN | true | [view](CERTS/fac21dba03b7a0483042818a126b8dc22298b595e36252e3ece22596f707de47/README.md) |
| 19 Oct 22 13:29 UTC | SHAKEN | true | [view](CERTS/735539bc495f743e25fea39ecef9566c17bedaf976cf0057d2bb3569ab448d02/README.md) |
| 20 Oct 22 11:12 UTC | SHAKEN | true | [view](CERTS/e01c63f891eca98375b1583ba29104c72618fbe3109cd7b558c37529af2cb29e/README.md) |
| 20 Oct 22 13:29 UTC | SHAKEN | true | [view](CERTS/56ac7c9635383e869ee9b816df361d00c4f17afe60188f234cc1e2906a913781/README.md) |
| 21 Oct 22 11:12 UTC | SHAKEN | true | [view](CERTS/8fb4a8e814f1c87cf2456b1ecf1e5ce491a5eb399b2d37de0d1a34b4ddb637db/README.md) |
| 21 Oct 22 13:29 UTC | SHAKEN | true | [view](CERTS/fc473a65f1226915a1bddec7b99fa1e53f0e1acc2f4e10106b07b18fc0054b4d/README.md) |
| 22 Oct 22 13:30 UTC | SHAKEN | true | [view](CERTS/8b259dd4c8adee7126f6d77d613a73d8e123a155ea3455a9531f86ef8572aa0e/README.md) |
| 24 Oct 22 11:12 UTC | SHAKEN | true | [view](CERTS/7dd1ed0b11d7d96b811a11562effd14cedc04552b643303507df78f89983e24c/README.md) |
| 25 Oct 22 11:12 UTC | SHAKEN | true | [view](CERTS/5eaa41cb4d2988f9518bc9d580a9c548dd41a6d7e12ba89c919fc302ca3ebaa4/README.md) |
| 26 Oct 22 11:12 UTC | SHAKEN | true | [view](CERTS/753a9492590251d2621c7d0ca1debbbb708d93027dca45ea69f3dfa251819eda/README.md) |
| 27 Oct 22 11:12 UTC | SHAKEN | true | [view](CERTS/4359cd2f3d758f495356d11669974761f5671c7512a2278340d0927a0a390fb4/README.md) |
| 27 Oct 22 13:29 UTC | SHAKEN | true | [view](CERTS/7fb833a2e3672fc5bcabcf05829f5d5336abba7e88577ad444e5179dd2df5ebb/README.md) |
| 28 Oct 22 11:12 UTC | SHAKEN | true | [view](CERTS/67a979476b3e4fb995be4075053f9a805b3df4b18dc09f532c2c239b3dadc13d/README.md) |
| 28 Oct 22 13:29 UTC | SHAKEN | true | [view](CERTS/748774b40fe43fdc9b29363b0bdf107884b3447e2110f8429cc6b893993c1bd2/README.md) |
| 29 Oct 22 13:29 UTC | SHAKEN | true | [view](CERTS/af6785002cd605e314824b90e6e06e3e39bb849078865200001a86382c0dd370/README.md) |
| 31 Oct 22 11:12 UTC | SHAKEN | true | [view](CERTS/cf1f5b80441df2551cfa967ecd684687eb9ea9dd5ce30a4035c8646cdb966c59/README.md) |
| 31 Oct 22 13:30 UTC | SHAKEN | true | [view](CERTS/3d44d0a8d9179445f0ac646cceb3f6a3566f63e2bed4ec08b58a017b063fa36d/README.md) |
| 01 Nov 22 11:12 UTC | SHAKEN | true | [view](CERTS/672077a75a58fb6c5d6bb8f067812b721dbad767c682832c2fa646e43ffd36b6/README.md) |
| 01 Nov 22 13:29 UTC | SHAKEN | true | [view](CERTS/a47b02b168a680b892b4ce263af30b19632ac8b8d62b1ef29c055c6f20f4ed01/README.md) |

#### CA Certificates

| Created at | Subject | Problems | Link |
|------------|---------|----------|------|
| 17 Mar 20 14:45 UTC | Comcast SHAKEN Root CA | false | [view](CERTS/b1132c5f12c3ca4d2ff119f2df99544336eb1703512ac99cc42d596e25125bbd/README.md) |
| 06 Apr 20 13:48 UTC | Comcast SHAKEN Intermediate CA | false | [view](CERTS/2f3abdfe711377f4d59f31d941962797c61f021c3924af31de99031f9ac54f77/README.md) |


Generated: 02 Nov 22 17:25 UTC