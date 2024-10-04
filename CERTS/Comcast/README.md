# STIR/SHAKEN CA Ecosystem Compliance

## Comcast

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 48 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 17 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 31 certificates being tested against the remaining rules
- 2.00 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 0.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 29 days is the average remaining validity for the certificates in the corpus
- 30 days is the average initial validity for the certificates in the corpus
- 31 certificates expire in the next 30 days
- 31.00 average number of unexpired certificates per OCN observed
- 1 unique OCNs observed in unexpired and valid certificate corpus

| Instances | Test | Source |
|-----------|------|--------|
| 31 | [e_atis_ext_crl_distribution_struct](ISSUES/e_atis_ext_crl_distribution_struct/README.md) | ATIS1000080 |
| 31 | [e_atis_serial_number_size](ISSUES/e_atis_serial_number_size/README.md) | ATIS1000080 |

#### CA Certificates

- 2 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 0 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 2 certificates being tested against the remaining rules
- 2.00 issues on average found in unexpired, trusted, and non-compliant certificates
- 50.00% of certificates contain one or more Error level issue
- 0.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 50.00% of certificates are too old to be assessed against currently enforced expectations
- 6746 days is the average remaining validity for the certificates in the corpus
- 7300 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days

| Instances | Test | Source |
|-----------|------|--------|
| 1 | [e_atis_ext_crl_distribution_struct_ca](ISSUES/e_atis_ext_crl_distribution_struct_ca/README.md) | ATIS1000080 |
| 1 | [e_atis_serial_number_size_ca](ISSUES/e_atis_serial_number_size_ca/README.md) | ATIS1000080 |

### Details

#### Leaf Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 05&#160;Sep&#160;24&#160;10:51&#160;UTC | SHAKEN 318J | 05&#160;Oct&#160;24&#160;10:51&#160;UTC | true | [view](CERTS/0c9682a36a7f55aca73664314021c6d707460962a46e1c1a1e48c1f9f6605d34/README.md) |
| 06&#160;Sep&#160;24&#160;10:51&#160;UTC | SHAKEN 318J | 06&#160;Oct&#160;24&#160;10:51&#160;UTC | true | [view](CERTS/98f85d7cb7ee9aa5dfe7efb49c20aea5c45a54afb911c86f083f7dcffd10bda1/README.md) |
| 07&#160;Sep&#160;24&#160;10:51&#160;UTC | SHAKEN 318J | 07&#160;Oct&#160;24&#160;10:51&#160;UTC | true | [view](CERTS/2cca6e9729499867a4cd163639f947102bd0fd538eaf84fc305c1cfbe5957bf7/README.md) |
| 08&#160;Sep&#160;24&#160;10:51&#160;UTC | SHAKEN 318J | 08&#160;Oct&#160;24&#160;10:51&#160;UTC | true | [view](CERTS/0fc4b2708c12cebd9dfc1de32cc5c2b0e394c583497ea3fe267253107c7da3ee/README.md) |
| 09&#160;Sep&#160;24&#160;10:51&#160;UTC | SHAKEN 318J | 09&#160;Oct&#160;24&#160;10:51&#160;UTC | true | [view](CERTS/dcfb9efda1ed810a85d5090d751db09bfa3f83397d115c454fbe6232f0145c20/README.md) |
| 10&#160;Sep&#160;24&#160;10:51&#160;UTC | SHAKEN 318J | 10&#160;Oct&#160;24&#160;10:51&#160;UTC | true | [view](CERTS/089839bc1ef1e2977cde7e9ca477fe8c44ea85126bad345297676da98468eb07/README.md) |
| 11&#160;Sep&#160;24&#160;10:51&#160;UTC | SHAKEN 318J | 11&#160;Oct&#160;24&#160;10:51&#160;UTC | true | [view](CERTS/ef8c68219490a43926e2d9f1e6059d357f3d9dc472646a0c200f9b69a506092c/README.md) |
| 12&#160;Sep&#160;24&#160;10:51&#160;UTC | SHAKEN 318J | 12&#160;Oct&#160;24&#160;10:51&#160;UTC | true | [view](CERTS/d171be8bb640d1b68aa2468de53cc56dab43d4dc1361c8d733bd205e2b0dbbb3/README.md) |
| 13&#160;Sep&#160;24&#160;10:51&#160;UTC | SHAKEN 318J | 13&#160;Oct&#160;24&#160;10:51&#160;UTC | true | [view](CERTS/d79de41532d2be7ddee01da88ee00a2857179924a519afd362090211c45d1f53/README.md) |
| 14&#160;Sep&#160;24&#160;10:51&#160;UTC | SHAKEN 318J | 14&#160;Oct&#160;24&#160;10:51&#160;UTC | true | [view](CERTS/7259a95dd63a3dbd7382e1b8b8205df089475e7309edc75e52b92cbb0abab97b/README.md) |
| 15&#160;Sep&#160;24&#160;10:51&#160;UTC | SHAKEN 318J | 15&#160;Oct&#160;24&#160;10:51&#160;UTC | true | [view](CERTS/d27f624df30a354c58a821e67006615662d179b51d30c7aba3f748fcaf2d0e6d/README.md) |
| 16&#160;Sep&#160;24&#160;10:51&#160;UTC | SHAKEN 318J | 16&#160;Oct&#160;24&#160;10:51&#160;UTC | true | [view](CERTS/c511dd9ea9c346d2d9932224a44510e3052393794b441cf348b8de2b337121ff/README.md) |
| 17&#160;Sep&#160;24&#160;10:51&#160;UTC | SHAKEN 318J | 17&#160;Oct&#160;24&#160;10:51&#160;UTC | true | [view](CERTS/85d43e88f49e3d93f427a542edd1ea75f569a27d6f4ba0d04f101b92c9c6bdc0/README.md) |
| 18&#160;Sep&#160;24&#160;10:51&#160;UTC | SHAKEN 318J | 18&#160;Oct&#160;24&#160;10:51&#160;UTC | true | [view](CERTS/df19c71644e2911d72db646e57aaa620d117b7c8183c8771b323dbee2956cb71/README.md) |
| 19&#160;Sep&#160;24&#160;10:51&#160;UTC | SHAKEN 318J | 19&#160;Oct&#160;24&#160;10:51&#160;UTC | true | [view](CERTS/723a4c2b7d3d89e7112b63cc999cabce9af10c88497cdf36ee8d50d379ad06d8/README.md) |
| 20&#160;Sep&#160;24&#160;10:51&#160;UTC | SHAKEN 318J | 20&#160;Oct&#160;24&#160;10:51&#160;UTC | true | [view](CERTS/a1bb9e984f775aef538fb005bb0d0b8c7f1c49a484783a1844f9ca836d9215c9/README.md) |
| 21&#160;Sep&#160;24&#160;10:51&#160;UTC | SHAKEN 318J | 21&#160;Oct&#160;24&#160;10:51&#160;UTC | true | [view](CERTS/c4f386c176fb5e17760fa1105585f776711fc4882063e91fef14968c05ba5c96/README.md) |
| 22&#160;Sep&#160;24&#160;10:51&#160;UTC | SHAKEN 318J | 22&#160;Oct&#160;24&#160;10:51&#160;UTC | true | [view](CERTS/2b314b69e696faa6019f975e44d11ff6c9faca922094040a2ed0e77cadcddc25/README.md) |
| 23&#160;Sep&#160;24&#160;10:51&#160;UTC | SHAKEN 318J | 23&#160;Oct&#160;24&#160;10:51&#160;UTC | true | [view](CERTS/c4f5fa49eb0181eb323b0561ff3759c44cbbb2ecb60508b91c9c35434837c119/README.md) |
| 24&#160;Sep&#160;24&#160;10:51&#160;UTC | SHAKEN 318J | 24&#160;Oct&#160;24&#160;10:51&#160;UTC | true | [view](CERTS/49aea36fbfee73862ef5c21e9966d14ea0e23b3455bb29592c0f48e7ba0bf23c/README.md) |
| 25&#160;Sep&#160;24&#160;10:51&#160;UTC | SHAKEN 318J | 25&#160;Oct&#160;24&#160;10:51&#160;UTC | true | [view](CERTS/d4342a26bf246de4a703481b28ee1102df419d3fd2fb6764098aa97851076e37/README.md) |
| 25&#160;Sep&#160;24&#160;21:43&#160;UTC | SHAKEN 318J | 25&#160;Oct&#160;24&#160;21:43&#160;UTC | true | [view](CERTS/942f229c15165d09a83f96d5b1e54b6c125844cc870b3fb4bc5051b5ad759803/README.md) |
| 26&#160;Sep&#160;24&#160;21:43&#160;UTC | SHAKEN 318J | 26&#160;Oct&#160;24&#160;21:43&#160;UTC | true | [view](CERTS/f15e34fa248cb946896d1f5d0ec4a0da6792664ec3985ed4fbcb319f26b7c278/README.md) |
| 27&#160;Sep&#160;24&#160;21:43&#160;UTC | SHAKEN 318J | 27&#160;Oct&#160;24&#160;21:43&#160;UTC | true | [view](CERTS/255e99ac26c069fcb7eac289a1fcefde2d1e20a29457f2066b041fcfab15103b/README.md) |
| 28&#160;Sep&#160;24&#160;21:43&#160;UTC | SHAKEN 318J | 28&#160;Oct&#160;24&#160;21:43&#160;UTC | true | [view](CERTS/c6972417e5b55a5f21927139eafbf8a96f9f1184ee723c4507943cf605a800aa/README.md) |
| 29&#160;Sep&#160;24&#160;21:43&#160;UTC | SHAKEN 318J | 29&#160;Oct&#160;24&#160;21:43&#160;UTC | true | [view](CERTS/572d1ddac5468a8b19e89d8730b0db99308a3b99dae1df43c2f193cd67d16933/README.md) |
| 30&#160;Sep&#160;24&#160;21:43&#160;UTC | SHAKEN 318J | 30&#160;Oct&#160;24&#160;21:43&#160;UTC | true | [view](CERTS/156730893bd9b8045036512c3cd40721ffeee159e2c312833f3c8fbb5ea3a3eb/README.md) |
| 01&#160;Oct&#160;24&#160;06:36&#160;UTC | SHAKEN 318J | 31&#160;Oct&#160;24&#160;06:36&#160;UTC | true | [view](CERTS/94a8ecbc64ea99bdad388ac35a3435a487b47ae2dc3ee9fdbb8eb8e6375b5c6d/README.md) |
| 02&#160;Oct&#160;24&#160;06:36&#160;UTC | SHAKEN 318J | 01&#160;Nov&#160;24&#160;06:36&#160;UTC | true | [view](CERTS/12dcf1a5b032fa0037e6d1184df49b35441d40abf80a13bc25826264d436664b/README.md) |
| 03&#160;Oct&#160;24&#160;06:36&#160;UTC | SHAKEN 318J | 02&#160;Nov&#160;24&#160;06:36&#160;UTC | true | [view](CERTS/59d04b4782a838fe32db40716704d0ca8a3bf84bc8f7906e64fb984bb648f917/README.md) |
| 04&#160;Oct&#160;24&#160;06:36&#160;UTC | SHAKEN 318J | 03&#160;Nov&#160;24&#160;06:36&#160;UTC | true | [view](CERTS/ce37efb8d48bcb24fdddb7f7ec185e80d4269c376b2ba68fa8077936998af0a6/README.md) |

#### CA Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 17&#160;Mar&#160;20&#160;14:45&#160;UTC | Comcast SHAKEN Root CA | 12&#160;Mar&#160;40&#160;14:45&#160;UTC | false | [view](CERTS/b1132c5f12c3ca4d2ff119f2df99544336eb1703512ac99cc42d596e25125bbd/README.md) |
| 11&#160;Jan&#160;24&#160;02:59&#160;UTC | Comcast SHAKEN Intermediate CA | 06&#160;Jan&#160;44&#160;02:59&#160;UTC | true | [view](CERTS/25af737667ed8b05cb8b8e7f44b2d7b5861551bea95ec48a73306ec75a92a662/README.md) |


Generated: 04 Oct 24 16:29 UTC