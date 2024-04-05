# STIR/SHAKEN CA Ecosystem Compliance

## Ribbon Communications

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 31 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 12 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 19 certificates being tested against the remaining rules
- 2.05 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 0.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 350 days is the average remaining validity for the certificates in the corpus
- 365 days is the average initial validity for the certificates in the corpus
- 3 certificates expire in the next 30 days
- 1.00 average number of unexpired certificates per OCN observed
- 19 unique OCNs observed in unexpired and valid certificate corpus

| Instances | Test | Source |
|-----------|------|--------|
| 19 | [e_atis_ext_not_specified](ISSUES/e_atis_ext_not_specified/README.md) | ATIS1000080 |
| 1 | [e_atis_subject_cn_spc](ISSUES/e_atis_subject_cn_spc/README.md) | ATIS1000080 |
| 19 | [e_shaken_certificate_policies_id](ISSUES/e_shaken_certificate_policies_id/README.md) | US_SHAKEN_CP |

#### CA Certificates

- 2 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 0 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 2 certificates being tested against the remaining rules
- 1.00 issues on average found in unexpired, trusted, and non-compliant certificates
- 50.00% of certificates contain one or more Error level issue
- 0.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 100.00% of certificates are too old to be assessed against currently enforced expectations
- 7196 days is the average remaining validity for the certificates in the corpus
- 6757 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days

| Instances | Test | Source |
|-----------|------|--------|
| 1 | [e_atis_ext_crl_distribution_ca](ISSUES/e_atis_ext_crl_distribution_ca/README.md) | ATIS1000080 |

### Details

#### Leaf Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 27&#160;Apr&#160;23&#160;16:22&#160;UTC | Gearheart SHAKEN 0408 | 26&#160;Apr&#160;24&#160;16:22&#160;UTC | true | [view](CERTS/82f85175421b1ee40b8a1d53b83d1ebbd8de91ca5022eddb84176e0be736edda/README.md) |
| 02&#160;May&#160;23&#160;13:30&#160;UTC | Surrytel STI SHAKEN 0503 | 01&#160;May&#160;24&#160;13:30&#160;UTC | true | [view](CERTS/e83aa6bcda724b091675c017f7ab9b3bbf1d39ba19fc3e00fcfa5542efc0a1d8/README.md) |
| 03&#160;May&#160;23&#160;15:18&#160;UTC | ThreeriverTelco SHAKEN 1525 | 02&#160;May&#160;24&#160;15:18&#160;UTC | true | [view](CERTS/b3e31d4de30e8a1fa6926123f646630349aa26590a364c67f2882866207cd2cf/README.md) |
| 08&#160;May&#160;23&#160;17:13&#160;UTC | Randolph SHAKEN 0496 | 07&#160;May&#160;24&#160;17:13&#160;UTC | true | [view](CERTS/4104fcb493a3117fa60e0d6a906a076a5292d8614c334754c8a453cda1c49583/README.md) |
| 10&#160;May&#160;23&#160;18:27&#160;UTC | Polar SHAKEN 1630 | 09&#160;May&#160;24&#160;18:27&#160;UTC | true | [view](CERTS/cb8f7326228bf19dac9d8f6f6e5369c163bf71a62561175e5e0cdf069287797a/README.md) |
| 15&#160;May&#160;23&#160;20:13&#160;UTC | Thacker SHAKEN 0419 | 14&#160;May&#160;24&#160;20:13&#160;UTC | true | [view](CERTS/b2d676e38a61cbc0ff5f46c21c1440204da25d30817d711ab5c6d294b07c8e8f/README.md) |
| 16&#160;May&#160;23&#160;19:35&#160;UTC | Ironton SHAKEN 1234 0175 | 15&#160;May&#160;24&#160;19:35&#160;UTC | true | [view](CERTS/07fa407eefdb8964682deef6cbc734636b12b9e97acbc4268fef06f0182436d2/README.md) |
| 23&#160;May&#160;23&#160;06:03&#160;UTC | Veracity SHAKEN 716D | 22&#160;May&#160;24&#160;06:03&#160;UTC | true | [view](CERTS/e4aefc8dd53731a5551cd37bb287abb065dbc343e7e71113832134aee28a8c1a/README.md) |
| 07&#160;Jun&#160;23&#160;20:11&#160;UTC | Ribbon SHAKEN 2086 | 06&#160;Jun&#160;24&#160;20:11&#160;UTC | true | [view](CERTS/3b68b514e31a136d42c36b131614bf640f110cfadf0bc75d3bfcb638b411cfe5/README.md) |
| 07&#160;Jun&#160;23&#160;20:25&#160;UTC | Netfortris SHAKEN 8886 | 06&#160;Jun&#160;24&#160;20:25&#160;UTC | true | [view](CERTS/8c0640e8957388ff63bf60c6c9bef298b41d795259bfeb2f3b30985ead768806/README.md) |
| 14&#160;Jun&#160;23&#160;17:02&#160;UTC | Peoples SHAKEN 2130 | 13&#160;Jun&#160;24&#160;17:02&#160;UTC | true | [view](CERTS/dba8b610ca22ffb1c622ce60ca29d0c0ce70a08315d0eea73f6bd1cb6993a705/README.md) |
| 18&#160;Jun&#160;23&#160;04:33&#160;UTC | Solarus SHAKEN 0974 | 17&#160;Jun&#160;24&#160;04:33&#160;UTC | true | [view](CERTS/6869f2cedb9c2ac3446f1686697459c398d85985fbbffa33536b13ae17ad2855/README.md) |
| 19&#160;Jun&#160;23&#160;20:13&#160;UTC | Longlines SHAKEN 1260 | 18&#160;Jun&#160;24&#160;20:13&#160;UTC | true | [view](CERTS/ccf64de6a66f15d4f31b85f1e3d1331657f6258a7e4432dde4fe1fc689d3e0dd/README.md) |
| 22&#160;Jun&#160;23&#160;19:11&#160;UTC | Fulton SHAKEN 0455 | 21&#160;Jun&#160;24&#160;19:11&#160;UTC | true | [view](CERTS/bae3690bab160297e3f4328ff4de8067f4e4b1c36bdb97d0235e0d2df72090e0/README.md) |
| 03&#160;Jul&#160;23&#160;05:29&#160;UTC | SMU SHAKEN 3390 | 02&#160;Jul&#160;24&#160;05:29&#160;UTC | true | [view](CERTS/3af759966803ee6daf2e803a763ff85bc64bb51bc1cbcd1c99bf7ea09c1204b4/README.md) |
| 10&#160;Jul&#160;23&#160;16:18&#160;UTC | Data SHAKEN 852B | 09&#160;Jul&#160;24&#160;16:18&#160;UTC | true | [view](CERTS/d4f2bf344a7d269d73c368e2ab321520ee1c07acdc2477ca287b4411432af42a/README.md) |
| 27&#160;Jul&#160;23&#160;15:50&#160;UTC | Bascom SHAKEN 0589 | 26&#160;Jul&#160;24&#160;15:50&#160;UTC | true | [view](CERTS/b15800790c2115b033607a21ec526276d4b29615683a8a1745206fa9e164cb23/README.md) |
| 04&#160;Jan&#160;24&#160;17:59&#160;UTC | Empire SHAKEN 087H | 03&#160;Jan&#160;25&#160;17:59&#160;UTC | true | [view](CERTS/ff4a083ed7527187359cb180383afae989c9a1144ad498e79ad3e3e05f5091fb/README.md) |
| 22&#160;Jan&#160;24&#160;14:57&#160;UTC | Eastex SHAKEN 2068 | 21&#160;Jan&#160;25&#160;14:57&#160;UTC | true | [view](CERTS/edc37467b2db082bbed1e00e885cf81f7bb179ad1ac3a75fe5f82f7c284da595/README.md) |

#### CA Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 13&#160;May&#160;21&#160;00:00&#160;UTC | SHAKEN Ribbon Issuing CA | 12&#160;May&#160;33&#160;23:59&#160;UTC | true | [view](CERTS/05a71a04eaedbdf4b0534f40768616d7c19c8deb5a3aefd1f4a04b3aab55a48f/README.md) |
| 13&#160;May&#160;21&#160;00:00&#160;UTC | SHAKEN Ribbon Root CA | 12&#160;May&#160;46&#160;23:59&#160;UTC | false | [view](CERTS/7c2799d3642d04f04fe667c3ab251c18689af323acdc43b2fa5f3dc89e3a0f14/README.md) |


Generated: 05 Apr 24 19:04 UTC