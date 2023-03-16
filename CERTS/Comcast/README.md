# STIR/SHAKEN CA Ecosystem Compliance

## Comcast

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 132 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 102 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 30 certificates being tested against the remaining rules
- 6.00 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 100.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 29 days is the average remaining validity for the certificates in the corpus
- 30 days is the average initial validity for the certificates in the corpus
- 30 certificates expire in the next 30 days
- 30.00 average number of unexpired certificates per OCN observed
- 1 unique OCNs observed in unexpired and valid certificate corpus

| Instances | Test | Source |
|-----------|------|--------|
| 30 | [e_atis_certificate_policies](ISSUES/e_atis_certificate_policies/README.md) | ATIS1000080 |
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
- 6835 days is the average remaining validity for the certificates in the corpus
- 7150 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days

No error, warning, or notice level issues were found

### Details

#### Leaf Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
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
| 02&#160;Mar&#160;23&#160;11:12&#160;UTC | SHAKEN | 01&#160;Apr&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/1514f88613f65399a608cf702e47e47207759a5ff8ec5aa792bf3661c378c88f/README.md) |
| 03&#160;Mar&#160;23&#160;11:12&#160;UTC | SHAKEN | 02&#160;Apr&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/290814af5e30f5990cc2bbbdb77ff95dd380e6a5dfa7e1aaf7bc663c6b95583c/README.md) |
| 04&#160;Mar&#160;23&#160;11:12&#160;UTC | SHAKEN | 03&#160;Apr&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/52301d9812d4e7a6d9985d164f571388055ff7eeed805cc7be8e2a833701196c/README.md) |
| 05&#160;Mar&#160;23&#160;11:12&#160;UTC | SHAKEN | 04&#160;Apr&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/ee88d3da6db1bdd755c7405c7371b6069a7e4ca5f70477cc464338d63b201f5c/README.md) |
| 06&#160;Mar&#160;23&#160;11:12&#160;UTC | SHAKEN | 05&#160;Apr&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/d4304a132aa816e8cc6642a7519dcc4c12c02c6ba5a487eb89c19e438af45b6e/README.md) |
| 07&#160;Mar&#160;23&#160;11:12&#160;UTC | SHAKEN | 06&#160;Apr&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/c3f510fd95361320628c0ff09b4d3fa9348ddc4b0840ed0dc3e264ad4745d6b4/README.md) |
| 08&#160;Mar&#160;23&#160;11:12&#160;UTC | SHAKEN | 07&#160;Apr&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/af71bef161a1f5259e0f32175d2e12a4e6b131e37e68d7d5737bf535a36d1da3/README.md) |
| 09&#160;Mar&#160;23&#160;11:12&#160;UTC | SHAKEN | 08&#160;Apr&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/7e604187e85a1c89718b6dbc19b61d4a2c6b762e2a7b2507bb9e0f3bc51b889e/README.md) |
| 10&#160;Mar&#160;23&#160;11:12&#160;UTC | SHAKEN | 09&#160;Apr&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/455eba1402d25b7a69275679dab1707b531c8c6b597f191b7e8dbe2e15e9a684/README.md) |
| 11&#160;Mar&#160;23&#160;11:12&#160;UTC | SHAKEN | 10&#160;Apr&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/ca2672ecb3ba798dbae159fd4310598fd3e89b4ff8354535405b8568f67a5f18/README.md) |
| 12&#160;Mar&#160;23&#160;11:12&#160;UTC | SHAKEN | 11&#160;Apr&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/12efa23c06bec7e1edeee4c80c2fb878f727268df2d02815dd7b4df12391d7af/README.md) |
| 13&#160;Mar&#160;23&#160;11:12&#160;UTC | SHAKEN | 12&#160;Apr&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/841a488f6f5d25b2cfa66b862df20d87b81060f5f6716c2c0eeef964131f5be6/README.md) |
| 14&#160;Mar&#160;23&#160;11:12&#160;UTC | SHAKEN | 13&#160;Apr&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/e95cce039c3b097be89374269fbdf27ce849912047019c4713c8ba1923d942ec/README.md) |
| 15&#160;Mar&#160;23&#160;11:12&#160;UTC | SHAKEN | 14&#160;Apr&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/47063eb8ae5f8a0db43f57772ab2af0f89c90bf9b5f785b58c45ddd633e23cc0/README.md) |
| 16&#160;Mar&#160;23&#160;11:12&#160;UTC | SHAKEN | 15&#160;Apr&#160;23&#160;11:12&#160;UTC | true | [view](CERTS/b84e30afc6140b051ef1eaf464c2c35f4506dd12781940914fe6a6841a0228cb/README.md) |

#### CA Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 17&#160;Mar&#160;20&#160;14:45&#160;UTC | Comcast SHAKEN Root CA | 12&#160;Mar&#160;40&#160;14:45&#160;UTC | false | [view](CERTS/b1132c5f12c3ca4d2ff119f2df99544336eb1703512ac99cc42d596e25125bbd/README.md) |
| 06&#160;Apr&#160;20&#160;13:48&#160;UTC | Comcast SHAKEN Intermediate CA | 06&#160;Jun&#160;39&#160;13:48&#160;UTC | false | [view](CERTS/2f3abdfe711377f4d59f31d941962797c61f021c3924af31de99031f9ac54f77/README.md) |


Generated: 16 Mar 23 19:18 UTC