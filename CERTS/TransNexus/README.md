# STIR/SHAKEN CA Ecosystem Compliance

## TransNexus

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 2305 certificates were included in the corpus being tested
- 2 certificates in the corpus were skipped because they are duplicates
- 2272 certificates in the corpus were skipped because they are expired
- 1 certificates in the corpus were skipped because they are not currently trusted
- 30 certificates being tested against the remaining rules
- 1.47 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 23.33% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 182 days is the average remaining validity for the certificates in the corpus
- 177 days is the average initial validity for the certificates in the corpus
- 13 certificates expire in the next 30 days
- 1.15 average number of unexpired certificates per OCN observed
- 26 unique OCNs observed in unexpired and valid certificate corpus

| Instances | Test | Source |
|-----------|------|--------|
| 7 | [e_incorrect_ku_encoding](ISSUES/e_incorrect_ku_encoding/README.md) | RFC5280 |
| 30 | [e_us_cp_subject_sn](ISSUES/e_us_cp_subject_sn/README.md) | US_SHAKEN_CP |
| 7 | [w_shaken_subject_rdn_unknown](ISSUES/w_shaken_subject_rdn_unknown/README.md) | SHAKEN_PKI_BEST_PRACTICES |

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
- 5448 days is the average remaining validity for the certificates in the corpus
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
| 10&#160;Jun&#160;22&#160;14:00&#160;UTC | SHAKEN 597J | 10&#160;Jun&#160;23&#160;14:00&#160;UTC | true | [view](CERTS/4baa838eb96ce358aa1565d0add8b75072258efd5d21a5bbb0eec878026a67be/README.md) |
| 20&#160;Jun&#160;22&#160;20:25&#160;UTC | SHAKEN 857J | 20&#160;Jun&#160;23&#160;20:25&#160;UTC | true | [view](CERTS/e1e0a4851fb7b51d617959c00fb6a17e7b23d34e1555ba70ab01c9182a2f85f8/README.md) |
| 22&#160;Jun&#160;22&#160;18:46&#160;UTC | SHAKEN 755J | 22&#160;Jun&#160;23&#160;18:46&#160;UTC | true | [view](CERTS/a8cac316fddf25b48b9350a325791be8b87e6484a5959ab605db1ae5e3e99a84/README.md) |
| 29&#160;Jun&#160;22&#160;20:24&#160;UTC | SHAKEN 736J | 29&#160;Jun&#160;23&#160;20:24&#160;UTC | true | [view](CERTS/71b25a480077816998e48e6a76f272e70aa91ecb19b0e773cd8777823cdead32/README.md) |
| 25&#160;Jul&#160;22&#160;18:36&#160;UTC | SHAKEN 578J | 25&#160;Jul&#160;23&#160;18:36&#160;UTC | true | [view](CERTS/f4af5a28131d680ef2682ae86181b64f0c919864d7dbaef00c0598d314b537a1/README.md) |
| 10&#160;Aug&#160;22&#160;18:11&#160;UTC | SHAKEN 073H | 10&#160;Aug&#160;23&#160;18:11&#160;UTC | true | [view](CERTS/87dabb91a689a6ac540b5ffbf5eb89326fada83fb822d2bb9ec31b68251b579e/README.md) |
| 13&#160;Sep&#160;22&#160;17:24&#160;UTC | SHAKEN 706J | 13&#160;Sep&#160;23&#160;17:24&#160;UTC | true | [view](CERTS/1eea5e17ae8b36cb5b6d657178273af745cea015e63586e6d44d1ba2ee8f722a/README.md) |
| 23&#160;Nov&#160;22&#160;17:31&#160;UTC | SHAKEN 4036 | 10&#160;May&#160;23&#160;17:31&#160;UTC | true | [view](CERTS/e03ff4259e6b456f1bb19d011f005932e83b0a35b520730c97cd77ad8868071c/README.md) |
| 07&#160;Dec&#160;22&#160;00:45&#160;UTC | SHAKEN 299K | 07&#160;Dec&#160;23&#160;00:45&#160;UTC | true | [view](CERTS/7b5e04ff01ace2ee384df3507886684204901bf23851866eb6b3891ecdd96a92/README.md) |
| 31&#160;Dec&#160;22&#160;03:33&#160;UTC | SHAKEN 8526 | 31&#160;Dec&#160;23&#160;03:33&#160;UTC | true | [view](CERTS/96672cb8cc7c71e38d9e6350eac6c4f5c430c1eefeabe00cc095527e20a20c9b/README.md) |
| 03&#160;Jan&#160;23&#160;16:07&#160;UTC | SHAKEN 6628 | 03&#160;Jun&#160;23&#160;16:07&#160;UTC | true | [view](CERTS/dec9a2b7e2fce8ca94e2b1313886772d8176d6adcac7c4cb2b295315fa79f5ab/README.md) |
| 12&#160;Jan&#160;23&#160;05:08&#160;UTC | SHAKEN 621J | 12&#160;May&#160;23&#160;05:08&#160;UTC | true | [view](CERTS/017fbb31b0e0c0f3dbb7de67e5e3f95c4ce10e546aaaca32faa97fb61eafcfe6/README.md) |
| 16&#160;Jan&#160;23&#160;17:54&#160;UTC | SHAKEN 622J | 15&#160;Jul&#160;23&#160;17:54&#160;UTC | true | [view](CERTS/4c659921cd1ec39c239ab31d4c1ca813017c01d51512ea11a1fe0572f034c534/README.md) |
| 25&#160;Jan&#160;23&#160;14:28&#160;UTC | SHAKEN 873J | 25&#160;Jan&#160;24&#160;14:28&#160;UTC | true | [view](CERTS/36e9ba32d337ad13909bee0b62125c9769a049c4ae10d749687d75af4e49387d/README.md) |
| 09&#160;Mar&#160;23&#160;18:46&#160;UTC | SHAKEN 193E | 08&#160;May&#160;23&#160;18:46&#160;UTC | true | [view](CERTS/34187f020bbbd3596eea83f5f2243b84920396d2c46eed5b8078a0b696c383e5/README.md) |
| 20&#160;Mar&#160;23&#160;18:51&#160;UTC | SHAKEN 505J | 18&#160;Jun&#160;23&#160;18:51&#160;UTC | true | [view](CERTS/aa04536bc3bbc9914c2b27653d1cf862dfcb571b53c784e83942afff6598e11d/README.md) |
| 20&#160;Mar&#160;23&#160;18:53&#160;UTC | SHAKEN 505J | 18&#160;Jun&#160;23&#160;18:53&#160;UTC | true | [view](CERTS/9a4feceff25f99e7d04d1542ef56bba8f060987c5b6734239f185997c5e58cf1/README.md) |
| 27&#160;Mar&#160;23&#160;17:18&#160;UTC | SHAKEN 2720 | 26&#160;Mar&#160;24&#160;17:18&#160;UTC | true | [view](CERTS/7282c0683ebff27881b1ccd1f66664c3102340ff657d9363fcecb9c1b7924c12/README.md) |
| 30&#160;Mar&#160;23&#160;18:51&#160;UTC | SHAKEN 952J | 29&#160;Apr&#160;23&#160;18:51&#160;UTC | true | [view](CERTS/e21d4b5c92be0e6db90e3a1a986ca049c8edd8c4bad4290f0ca6f4ae42adc537/README.md) |
| 04&#160;Apr&#160;23&#160;16:51&#160;UTC | SHAKEN 551G | 04&#160;May&#160;23&#160;16:50&#160;UTC | true | [view](CERTS/5123718621bfacb494869d8ac857c3387021c70d5e838ee00fadbeaa399bdec2/README.md) |
| 04&#160;Apr&#160;23&#160;23:14&#160;UTC | SHAKEN 4632 | 04&#160;May&#160;23&#160;23:14&#160;UTC | true | [view](CERTS/a89b5761f81a706cc098bc2c966e51a203c45a76bf423d09f379210a17aeae44/README.md) |
| 08&#160;Apr&#160;23&#160;04:42&#160;UTC | SHAKEN 345J | 08&#160;May&#160;23&#160;04:42&#160;UTC | true | [view](CERTS/dd31135745befe8cbca897037627e600c7991cb797c5b6244a61783993e2f2cd/README.md) |
| 08&#160;Apr&#160;23&#160;05:13&#160;UTC | SHAKEN 345J | 08&#160;May&#160;23&#160;05:13&#160;UTC | true | [view](CERTS/ed628fcdb76ed3264990a17371588f8b30d06f282a040b2d612a3bcb7aac835b/README.md) |
| 08&#160;Apr&#160;23&#160;18:53&#160;UTC | SHAKEN 193E | 07&#160;Jun&#160;23&#160;18:53&#160;UTC | true | [view](CERTS/8cc7b08b6c144bae17ff71317ebea1877859c2bbdd387c6b9083a130bf7ab18c/README.md) |
| 10&#160;Apr&#160;23&#160;12:54&#160;UTC | SHAKEN 815G | 09&#160;Jul&#160;23&#160;12:54&#160;UTC | true | [view](CERTS/71206dbb241c8dbf9de187fca7b7e515fff61fbedcd1baacdc4a6a399ca2f887/README.md) |
| 11&#160;Apr&#160;23&#160;16:12&#160;UTC | SHAKEN 722J | 11&#160;May&#160;23&#160;16:12&#160;UTC | true | [view](CERTS/e9b82c67177482f81b00257a1bb3e819f7e5b91b9f3a5e0780411449a321ffb2/README.md) |
| 13&#160;Apr&#160;23&#160;14:12&#160;UTC | SHAKEN 366G | 13&#160;May&#160;23&#160;14:12&#160;UTC | true | [view](CERTS/c6909229d017c42518bc16a1e108ce2744ebfe792e043034072b4b1c4d2c1327/README.md) |
| 14&#160;Apr&#160;23&#160;19:23&#160;UTC | SHAKEN 577F | 14&#160;May&#160;23&#160;19:23&#160;UTC | true | [view](CERTS/e6894fb62b47a83fd94ab52c22f53a85ef46382c47ae1e76bef418c948daaa4a/README.md) |
| 14&#160;Apr&#160;23&#160;20:11&#160;UTC | SHAKEN 841J | 28&#160;Apr&#160;23&#160;20:11&#160;UTC | true | [view](CERTS/5beb9d77b1b2f46209097868fe384ac1a895617e95117d7e6e2a3b407350401c/README.md) |
| 17&#160;Apr&#160;23&#160;18:52&#160;UTC | SHAKEN 952J | 17&#160;May&#160;23&#160;18:52&#160;UTC | true | [view](CERTS/76d40e4c6db8915a623775f1d7c2bf960e2747a9b1ccea0ffc3c23f197c02e44/README.md) |

#### CA Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 20&#160;Aug&#160;21&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Issuing CA3 | 19&#160;Aug&#160;31&#160;23:59&#160;UTC | true | [view](CERTS/51b51431ffd8055cd7e30c213540f4868258058d4c4c9f23f1ed8bea75101e37/README.md) |
| 20&#160;Aug&#160;21&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Issuing CA1 | 19&#160;Aug&#160;31&#160;23:59&#160;UTC | true | [view](CERTS/037a7c6621814bd5f01fa2080a77a245458694d9d566e0050e068fc853317857/README.md) |
| 20&#160;Aug&#160;21&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Root CA1 | 19&#160;Aug&#160;41&#160;23:59&#160;UTC | true | [view](CERTS/59ab943f4e8c3c17755cc8e4abb5ac65736ce74ec587c62cb744ab0babbbe2fe/README.md) |
| 24&#160;Oct&#160;22&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Issuing CA4 | 23&#160;Oct&#160;32&#160;23:59&#160;UTC | false | [view](CERTS/80dd12d935f46c256c7c1289a0834b21ec8f2f8c35f2864928d1e75f3a280665/README.md) |
| 24&#160;Oct&#160;22&#160;00:00&#160;UTC | TransNexus, Inc. SHAKEN Root CA2 | 23&#160;Oct&#160;42&#160;23:59&#160;UTC | false | [view](CERTS/a26e04fc786ab70b8085236b2c53f8cfbf5d0c6a5c2c9c3e9f91669fbb8ea4d5/README.md) |


Generated: 28 Apr 23 02:17 UTC