# STIR/SHAKEN CA Ecosystem Compliance

## TransNexus

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 1171 certificates were included in the corpus being tested
- 1 certificates in the corpus were skipped because they are duplicates
- 1146 certificates in the corpus were skipped because they are expired
- 1 certificates in the corpus were skipped because they are not currently trusted
- 23 certificates being tested against the remaining rules
- 2.04 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 43.48% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 210 days is the average remaining validity for the certificates in the corpus
- 217 days is the average initial validity for the certificates in the corpus
- 8 certificates expire in the next 30 days
- 1.05 average number of unexpired certificates per OCN observed
- 22 unique OCNs observed in unexpired and valid certificate corpus

| Instances | Test | Source |
|-----------|------|--------|
| 2 | [e_atis_certificate_policies](ISSUES/e_atis_certificate_policies/README.md) | ATIS1000080 |
| 10 | [e_incorrect_ku_encoding](ISSUES/e_incorrect_ku_encoding/README.md) | RFC5280 |
| 2 | [e_us_cp_ambiguous_identifier](ISSUES/e_us_cp_ambiguous_identifier/README.md) | US_SHAKEN_CP |
| 23 | [e_us_cp_subject_sn](ISSUES/e_us_cp_subject_sn/README.md) | US_SHAKEN_CP |
| 10 | [w_shaken_subject_rdn_unknown](ISSUES/w_shaken_subject_rdn_unknown/README.md) | SHAKEN_PKI_BEST_PRACTICES |

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
- 5462 days is the average remaining validity for the certificates in the corpus
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
| 23&#160;Nov&#160;22&#160;17:31&#160;UTC | SHAKEN 4036 | 10&#160;May&#160;23&#160;17:31&#160;UTC | true | [view](CERTS/e03ff4259e6b456f1bb19d011f005932e83b0a35b520730c97cd77ad8868071c/README.md) |
| 07&#160;Dec&#160;22&#160;00:45&#160;UTC | SHAKEN 299K | 07&#160;Dec&#160;23&#160;00:45&#160;UTC | true | [view](CERTS/7b5e04ff01ace2ee384df3507886684204901bf23851866eb6b3891ecdd96a92/README.md) |
| 22&#160;Dec&#160;22&#160;17:30&#160;UTC | SHAKEN 505J | 22&#160;Mar&#160;23&#160;17:30&#160;UTC | true | [view](CERTS/7ff43923df80c3a7e2d854b42f33c69069af4625f5267df0d7ff7b2004ad1c1b/README.md) |
| 30&#160;Dec&#160;22&#160;06:30&#160;UTC | SHAKEN 807J | 28&#160;Feb&#160;23&#160;06:30&#160;UTC | true | [view](CERTS/62da6686ffc785eb4f15528d7bf06acd4d4778bbe258808e662ee91ffb868ccf/README.md) |
| 30&#160;Dec&#160;22&#160;12:34&#160;UTC | SHAKEN 193E | 28&#160;Feb&#160;23&#160;12:34&#160;UTC | true | [view](CERTS/36248069600bf4dcf5c85d58b56e6960278411436d818c338226d070a778cf31/README.md) |
| 31&#160;Dec&#160;22&#160;03:33&#160;UTC | SHAKEN 8526 | 31&#160;Dec&#160;23&#160;03:33&#160;UTC | true | [view](CERTS/96672cb8cc7c71e38d9e6350eac6c4f5c430c1eefeabe00cc095527e20a20c9b/README.md) |
| 05&#160;Jan&#160;23&#160;15:50&#160;UTC | SHAKEN 722J | 04&#160;Feb&#160;23&#160;15:50&#160;UTC | true | [view](CERTS/6707ea8caeed61aaa9c0489a5b1434622a9c938978f1bea4d2ea7fc5dbd81586/README.md) |
| 10&#160;Jan&#160;23&#160;13:53&#160;UTC | SHAKEN 815G | 10&#160;Apr&#160;23&#160;13:53&#160;UTC | true | [view](CERTS/2b3c67da667aa641a6a6da726a363284a1267727596a620d8e8b2b68fc3a6319/README.md) |
| 10&#160;Jan&#160;23&#160;16:32&#160;UTC | SHAKEN 551G | 09&#160;Feb&#160;23&#160;16:32&#160;UTC | true | [view](CERTS/628611026de07060665254cc9e3821cf274d115b7a1c27c0f747972dbf9ed121/README.md) |
| 14&#160;Jan&#160;23&#160;04:23&#160;UTC | SHAKEN 345J | 13&#160;Feb&#160;23&#160;04:23&#160;UTC | true | [view](CERTS/ab63023fad304f5898fcd2921df1a0b78cf13aff37ce6a926370b243580920a6/README.md) |
| 14&#160;Jan&#160;23&#160;05:01&#160;UTC | SHAKEN 345J | 13&#160;Feb&#160;23&#160;05:01&#160;UTC | true | [view](CERTS/a1085082cf0a00fba83dcc43ff32d5aca855b418133feeebeb14728119047b3d/README.md) |
| 16&#160;Jan&#160;23&#160;17:54&#160;UTC | SHAKEN 622J | 15&#160;Jul&#160;23&#160;17:54&#160;UTC | true | [view](CERTS/4c659921cd1ec39c239ab31d4c1ca813017c01d51512ea11a1fe0572f034c534/README.md) |
| 17&#160;Jan&#160;23&#160;18:22&#160;UTC | SHAKEN 952J | 16&#160;Feb&#160;23&#160;18:21&#160;UTC | true | [view](CERTS/02c5d91ab286934436712c3537a6b58b3a1758b95a530644c10355fbfa9384a8/README.md) |

#### CA Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 20&#160;Aug&#160;21&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Issuing CA3 | 19&#160;Aug&#160;31&#160;23:59&#160;UTC | true | [view](CERTS/51b51431ffd8055cd7e30c213540f4868258058d4c4c9f23f1ed8bea75101e37/README.md) |
| 20&#160;Aug&#160;21&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Issuing CA1 | 19&#160;Aug&#160;31&#160;23:59&#160;UTC | true | [view](CERTS/037a7c6621814bd5f01fa2080a77a245458694d9d566e0050e068fc853317857/README.md) |
| 20&#160;Aug&#160;21&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Root CA1 | 19&#160;Aug&#160;41&#160;23:59&#160;UTC | true | [view](CERTS/59ab943f4e8c3c17755cc8e4abb5ac65736ce74ec587c62cb744ab0babbbe2fe/README.md) |
| 24&#160;Oct&#160;22&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Issuing CA4 | 23&#160;Oct&#160;32&#160;23:59&#160;UTC | false | [view](CERTS/80dd12d935f46c256c7c1289a0834b21ec8f2f8c35f2864928d1e75f3a280665/README.md) |
| 24&#160;Oct&#160;22&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Root CA2 | 23&#160;Oct&#160;42&#160;23:59&#160;UTC | false | [view](CERTS/a26e04fc786ab70b8085236b2c53f8cfbf5d0c6a5c2c9c3e9f91669fbb8ea4d5/README.md) |


Generated: 30 Jan 23 23:10 UTC