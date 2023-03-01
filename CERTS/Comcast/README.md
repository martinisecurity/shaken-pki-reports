# STIR/SHAKEN CA Ecosystem Compliance

## Comcast

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 171 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 141 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 30 certificates being tested against the remaining rules
- 6.10 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 100.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 30 days is the average remaining validity for the certificates in the corpus
- 30 days is the average initial validity for the certificates in the corpus
- 30 certificates expire in the next 30 days
- 30.00 average number of unexpired certificates per OCN observed
- 1 unique OCNs observed in unexpired and valid certificate corpus

| Instances | Test | Source |
|-----------|------|--------|
| 30 | [e_atis_certificate_policies](ISSUES/e_atis_certificate_policies/README.md) | ATIS1000080 |
| 3 | [e_atis_serial_number](ISSUES/e_atis_serial_number/README.md) | ATIS1000080 |
| 30 | [e_atis_subject_cn](ISSUES/e_atis_subject_cn/README.md) | ATIS1000080 |
| 30 | [e_atis_subject_key_identifier](ISSUES/e_atis_subject_key_identifier/README.md) | ATIS1000080 |
| 30 | [e_us_cp_ambiguous_identifier](ISSUES/e_us_cp_ambiguous_identifier/README.md) | US_SHAKEN_CP |
| 30 | [e_us_cp_subject_sn](ISSUES/e_us_cp_subject_sn/README.md) | US_SHAKEN_CP |
| 30 | [w_ext_subject_key_identifier_missing_sub_cert](ISSUES/w_ext_subject_key_identifier_missing_sub_cert/README.md) | RFC5280 |

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
- 6840 days is the average remaining validity for the certificates in the corpus
- 7150 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days

No error, warning, or notice level issues were found

### Details

#### Leaf Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 31&#160;Jan&#160;23&#160;11:12&#160;UTC | SHAKEN | 02&#160;Mar&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/59aa65bba8fcf131d7ffdcf5871e036b4e57daaa5e6bd4539fa658e3f10a5da7/README.md) |
| 01&#160;Feb&#160;23&#160;11:12&#160;UTC | SHAKEN | 03&#160;Mar&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/8a2d26d03d085aa11269119b3d03dde1da46343044b0b04fa8d71497ca5e4334/README.md) |
| 02&#160;Feb&#160;23&#160;11:12&#160;UTC | SHAKEN | 04&#160;Mar&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/4eafa96478e89933220c35bdbceb083d5c3721d7e76027aeb7461cf0103303a8/README.md) |
| 03&#160;Feb&#160;23&#160;11:12&#160;UTC | SHAKEN | 05&#160;Mar&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/457754bac432e826f3867cee3b27e6a7142cd7fe68fcc3d1af0de1d6f5e6dbaf/README.md) |
| 04&#160;Feb&#160;23&#160;11:12&#160;UTC | SHAKEN | 06&#160;Mar&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/64dc9ebfa1f96eef8e55764dbdb78850ef55e8636eaaa594fc0db3bcf057170e/README.md) |
| 05&#160;Feb&#160;23&#160;11:12&#160;UTC | SHAKEN | 07&#160;Mar&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/957d85c76db214cacf0045e0ea978c0926cf048f8d3b1b3964a38b08e68a6ce4/README.md) |
| 06&#160;Feb&#160;23&#160;11:12&#160;UTC | SHAKEN | 08&#160;Mar&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/b26cfca2ab38d21830eb6ff86284b231d70ae623d7c7586ce6256e10cb1b5e20/README.md) |
| 07&#160;Feb&#160;23&#160;11:12&#160;UTC | SHAKEN | 09&#160;Mar&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/f8b8b9dc673c93755897ef90880a82d1a64dfeaaa3a86104f31643b76dc99b43/README.md) |
| 08&#160;Feb&#160;23&#160;11:12&#160;UTC | SHAKEN | 10&#160;Mar&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/1d9e8aff500637e6087b96bd75176ed9d9d61d64699713cc611c59bb56804a7d/README.md) |
| 09&#160;Feb&#160;23&#160;11:12&#160;UTC | SHAKEN | 11&#160;Mar&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/85a0014e954ad63c09a148f471bf720976af67123b39f127d2b3220aed7e4add/README.md) |
| 10&#160;Feb&#160;23&#160;11:12&#160;UTC | SHAKEN | 12&#160;Mar&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/02ec67107ee584ed3feef0496bc25758212505e1d23acf17a8cb084bd33c8801/README.md) |
| 11&#160;Feb&#160;23&#160;11:12&#160;UTC | SHAKEN | 13&#160;Mar&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/6670d7f4687e8a66cb2a7a864cda72d91e442968dd6fba49904ea0d71de7aff3/README.md) |
| 12&#160;Feb&#160;23&#160;11:12&#160;UTC | SHAKEN | 14&#160;Mar&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/911426ac28ce1651ca3043f3643a722aca0cd197a4a54981c4f4bdcd4956e80a/README.md) |
| 13&#160;Feb&#160;23&#160;11:12&#160;UTC | SHAKEN | 15&#160;Mar&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/7ab2e97af439c0ee0c74424feb9a7437b92f870cbba0e2c78b49ae879bd3072c/README.md) |
| 14&#160;Feb&#160;23&#160;11:12&#160;UTC | SHAKEN | 16&#160;Mar&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/947d71fe19dfb214fa9f46c253c159bb072807e6ed0dd096ad2553ecd51f4600/README.md) |
| 15&#160;Feb&#160;23&#160;11:12&#160;UTC | SHAKEN | 17&#160;Mar&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/62868767c873186de497f2a6ecbc3b35a9c70d34fbb651d51cb8a189d371bb25/README.md) |
| 16&#160;Feb&#160;23&#160;11:12&#160;UTC | SHAKEN | 18&#160;Mar&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/8e4c84991a1aab89d3f8dba823fd04b279e716e69a6ef8b7a6d753a6accf16ec/README.md) |
| 17&#160;Feb&#160;23&#160;11:12&#160;UTC | SHAKEN | 19&#160;Mar&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/a15ffe939e244f3e88899ea9ca7631de0b309cf7ef2b2b19c7b85c30b577b3bb/README.md) |
| 18&#160;Feb&#160;23&#160;11:12&#160;UTC | SHAKEN | 20&#160;Mar&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/a80a75bd927015c3ef5bfde32a834594694c058870cdf2eaf9ac9faacd9e2be8/README.md) |
| 19&#160;Feb&#160;23&#160;11:12&#160;UTC | SHAKEN | 21&#160;Mar&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/310db73d59112c413676507e5aa2985b9fd996b0b05afd7d1c6fdd7d2f9c77f4/README.md) |
| 20&#160;Feb&#160;23&#160;11:12&#160;UTC | SHAKEN | 22&#160;Mar&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/b17c72457dfe0668bfa84c0161c0e50cca422b8b24c0db1376f04ce21bd504b6/README.md) |
| 21&#160;Feb&#160;23&#160;11:12&#160;UTC | SHAKEN | 23&#160;Mar&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/2fe5dfb3de7a791be5bf6290a828f1bd1aca388a2c7d508e8708499648d51939/README.md) |
| 22&#160;Feb&#160;23&#160;11:12&#160;UTC | SHAKEN | 24&#160;Mar&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/3df6c667af1ac28d421817258ed7475d49e3c7dd17e22a5dfc5aa4012b7f4b40/README.md) |
| 23&#160;Feb&#160;23&#160;11:12&#160;UTC | SHAKEN | 25&#160;Mar&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/56967fb8aea0dab933c759effe3244347968d53e5f277340620bb031beda571e/README.md) |
| 24&#160;Feb&#160;23&#160;11:12&#160;UTC | SHAKEN | 26&#160;Mar&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/47981b91863b9cff404e60c4dd227145d34185291ee494a197bb44236db652f0/README.md) |
| 25&#160;Feb&#160;23&#160;11:12&#160;UTC | SHAKEN | 27&#160;Mar&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/ab07c118beaaf667717ea284616f620e2d7b3260135fd02e859de5dda802a07b/README.md) |
| 26&#160;Feb&#160;23&#160;11:12&#160;UTC | SHAKEN | 28&#160;Mar&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/f1a38964c2ac928d68c6fb35ff8661b8ce67fd564e627e0a75dc9c761770912f/README.md) |
| 27&#160;Feb&#160;23&#160;11:12&#160;UTC | SHAKEN | 29&#160;Mar&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/6133e368fdef712289059b6da5d9c7a1b5e5e1bac4a996c1278c18d6a9361243/README.md) |
| 28&#160;Feb&#160;23&#160;11:12&#160;UTC | SHAKEN | 30&#160;Mar&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/e12d09394c2662bfa2008de03881e8ca6d4bcdd49599a46a3ac82dedc226619d/README.md) |
| 01&#160;Mar&#160;23&#160;11:12&#160;UTC | SHAKEN | 31&#160;Mar&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/0227f5b1b4f5218c5e1d6edd2dd7307faaa156c2a727fdcacbe15145b64ebab2/README.md) |

#### CA Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 17&#160;Mar&#160;20&#160;14:45&#160;UTC | Comcast SHAKEN Root CA | 12&#160;Mar&#160;40&#160;14:45&#160;UTC | false | [view](CERTS/b1132c5f12c3ca4d2ff119f2df99544336eb1703512ac99cc42d596e25125bbd/README.md) |
| 06&#160;Apr&#160;20&#160;13:48&#160;UTC | Comcast SHAKEN Intermediate CA | 06&#160;Jun&#160;39&#160;13:48&#160;UTC | false | [view](CERTS/2f3abdfe711377f4d59f31d941962797c61f021c3924af31de99031f9ac54f77/README.md) |


Generated: 01 Mar 23 18:22 UTC