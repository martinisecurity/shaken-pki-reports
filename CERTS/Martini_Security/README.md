# STIR/SHAKEN CA Ecosystem Compliance

## Martini Security

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 146 certificates were included in the corpus being tested
- 45 certificates in the corpus were skipped because they are duplicates
- 55 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 46 certificates being tested against the remaining rules
- 1.00 issues on average found in unexpired, trusted, and non-compliant certificates
- 0.00% of certificates contain one or more Error level issue
- 0.00% of certificates contain one or more Warning level issue
- 2.17% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 104 days is the average remaining validity for the certificates in the corpus
- 105 days is the average initial validity for the certificates in the corpus
- 15 certificates expire in the next 30 days
- 3.54 average number of unexpired certificates per OCN observed
- 13 unique OCNs observed in unexpired and valid certificate corpus

| Instances | Test | Source |
|-----------|------|--------|
| 1 | [n_atis_certificate_policy_critical](ISSUES/n_atis_certificate_policy_critical/README.md) | ATIS1000080 |

#### CA Certificates

- 3 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 0 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 3 certificates being tested against the remaining rules
- 1.00 issues on average found in unexpired, trusted, and non-compliant certificates
- 0.00% of certificates contain one or more Error level issue
- 0.00% of certificates contain one or more Warning level issue
- 33.33% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 5380 days is the average remaining validity for the certificates in the corpus
- 4261 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days

| Instances | Test | Source |
|-----------|------|--------|
| 1 | [n_atis_ca_certificate_policy_critical](ISSUES/n_atis_ca_certificate_policy_critical/README.md) | ATIS1000080 |

### Details

#### Leaf Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 28&#160;Sep&#160;22&#160;18:19&#160;UTC | SHAKEN 709J | 28&#160;Sep&#160;23&#160;17:45&#160;UTC | true | [view](CERTS/8862209bef596987c13ab19a89a9fc62018dc2a4e8c9cb927827aadf1c458eee/README.md) |
| 02&#160;Mar&#160;23&#160;20:12&#160;UTC | SHAKEN 709J | 21&#160;Oct&#160;23&#160;17:35&#160;UTC | false | [view](CERTS/99305af5923a22b7d885f5b666f926b69a26825a390e94a318f13ef60e4d465e/README.md) |
| 03&#160;Mar&#160;23&#160;15:17&#160;UTC | SHAKEN 473K | 02&#160;Mar&#160;24&#160;15:17&#160;UTC | false | [view](CERTS/0eb7d483fd17cb0f1c65b68302dcfe25acbfa599174977835e04fc37784795bb/README.md) |
| 08&#160;Mar&#160;23&#160;19:20&#160;UTC | SHAKEN 709J | 06&#160;Jun&#160;23&#160;19:20&#160;UTC | false | [view](CERTS/f384b0350255145848e221abe01dc77e00ee96e584cd02f60285f88e9bd4ad04/README.md) |
| 08&#160;Mar&#160;23&#160;19:24&#160;UTC | SHAKEN 709J | 05&#160;Jun&#160;23&#160;19:24&#160;UTC | false | [view](CERTS/ef8ab01efc6230b7dba6be2f0b1199651123deab527dfa69b1c352a473d7e3d4/README.md) |
| 31&#160;Mar&#160;23&#160;17:38&#160;UTC | SHAKEN 709J | 29&#160;Jun&#160;23&#160;17:38&#160;UTC | false | [view](CERTS/d81be92be009d229a0580c25d7c379fcc71a48bd83ed2725c95a33ed97818e77/README.md) |
| 01&#160;Apr&#160;23&#160;09:38&#160;UTC | SHAKEN 042K | 30&#160;Jun&#160;23&#160;04:00&#160;UTC | false | [view](CERTS/ff3e17cf9b64cae520378120341ca0900e4b59e38ca14c768cab69ee5a769c56/README.md) |
| 01&#160;Apr&#160;23&#160;20:59&#160;UTC | SHAKEN 048K | 30&#160;Jun&#160;23&#160;00:00&#160;UTC | false | [view](CERTS/9d90646df9685962a73f6312a36bba0687e0d053245a727a2ea6b9d75cd3ef16/README.md) |
| 01&#160;Apr&#160;23&#160;21:06&#160;UTC | SHAKEN 066K | 30&#160;Jun&#160;23&#160;05:00&#160;UTC | false | [view](CERTS/5e6f7d19b896754b8eb8abaee15608716cd1607ac256d97c2e8c836646d8f6f2/README.md) |
| 01&#160;Apr&#160;23&#160;21:08&#160;UTC | SHAKEN 031K | 30&#160;Jun&#160;23&#160;05:00&#160;UTC | false | [view](CERTS/5e68b4b128d8575a56e3d98429362fb2fc184497ee6e15d71e42f15ea7ddd88d/README.md) |
| 01&#160;Apr&#160;23&#160;21:08&#160;UTC | SHAKEN 073K | 30&#160;Jun&#160;23&#160;06:00&#160;UTC | false | [view](CERTS/00f6a96504326942f4f5ec4c2c9b44c62cbe0efb4771c7f6e1d41fc3ecfff8ab/README.md) |
| 01&#160;Apr&#160;23&#160;21:11&#160;UTC | SHAKEN 039K | 30&#160;Jun&#160;23&#160;05:00&#160;UTC | false | [view](CERTS/b63961c57c5c1b46f1821f45c39dffc4a78218c1e07a5acae9a4d588646460a1/README.md) |
| 01&#160;Apr&#160;23&#160;21:12&#160;UTC | SHAKEN 115K | 30&#160;Jun&#160;23&#160;04:00&#160;UTC | false | [view](CERTS/19d2269e83fb8a8df50965533fcb664ba1ef92fc94889c4bcf3902d044da1c69/README.md) |
| 01&#160;Apr&#160;23&#160;21:12&#160;UTC | SHAKEN 223K | 30&#160;Jun&#160;23&#160;07:00&#160;UTC | false | [view](CERTS/23945d092fe11e5ebc1c53475bb97dc42a54c08b216cbcca8c8d7225b7d9936c/README.md) |
| 01&#160;Apr&#160;23&#160;21:13&#160;UTC | SHAKEN 110K | 30&#160;Jun&#160;23&#160;06:00&#160;UTC | false | [view](CERTS/484610d2d39d30434f730dc6b4973f58ed36897b5f75cbc72038e710548c16e4/README.md) |
| 01&#160;Apr&#160;23&#160;21:17&#160;UTC | SHAKEN 148K | 30&#160;Jun&#160;23&#160;04:00&#160;UTC | false | [view](CERTS/fb0a02939d367fb6c702cd1468e48ec874dc31a6f2366d530fb7c112b011c6ec/README.md) |
| 01&#160;Apr&#160;23&#160;21:28&#160;UTC | SHAKEN 148K | 30&#160;Jun&#160;23&#160;04:00&#160;UTC | false | [view](CERTS/f2553306faf88ac70da8a084420b9fdd639e708ef48a9ade4e7b2f12ab5efc81/README.md) |
| 01&#160;Apr&#160;23&#160;21:36&#160;UTC | SHAKEN 709J | 30&#160;Jun&#160;23&#160;05:00&#160;UTC | false | [view](CERTS/7af39a6b83b65520b1ce0e386c64c27c40c1d762b06a36e63248906593eea8ab/README.md) |
| 03&#160;Apr&#160;23&#160;14:33&#160;UTC | SHAKEN 709J | 02&#160;Jul&#160;23&#160;14:33&#160;UTC | false | [view](CERTS/cc6980699aca4721c3fbd210bac9dcd83358d9eda817ffd8cabe9b52858c0fc9/README.md) |
| 03&#160;Apr&#160;23&#160;19:16&#160;UTC | SHAKEN 709J | 02&#160;Jul&#160;23&#160;19:16&#160;UTC | false | [view](CERTS/c83fbce3b5df845caa20379b86300f4629fa53765a1dae6526a7405731c9ebd1/README.md) |
| 04&#160;Apr&#160;23&#160;20:17&#160;UTC | SHAKEN 709J | 03&#160;Jul&#160;23&#160;20:17&#160;UTC | false | [view](CERTS/51640f16691a97153cb1db61a05476560de4ff1f05e142fe3210edb6651d2d7c/README.md) |
| 06&#160;Apr&#160;23&#160;03:48&#160;UTC | SHAKEN 709J | 05&#160;Jul&#160;23&#160;03:48&#160;UTC | false | [view](CERTS/7f6c7f3c59405d789dfb672b4a2d390b0a30cc6457d55eb39d1ba9899517674b/README.md) |
| 06&#160;Apr&#160;23&#160;18:11&#160;UTC | SHAKEN 709J | 05&#160;Jul&#160;23&#160;05:00&#160;UTC | false | [view](CERTS/e93b89a83fc7b9c051c8ad023d0aa0a1434a719310b3ed10d955ba8c2f5ce986/README.md) |
| 11&#160;Apr&#160;23&#160;21:05&#160;UTC | SHAKEN 709J | 10&#160;Jul&#160;23&#160;21:05&#160;UTC | false | [view](CERTS/a6f9e76e35777d37a241cb49118bb8a89a7e4661f311c02878aa17ee22196452/README.md) |
| 12&#160;Apr&#160;23&#160;20:24&#160;UTC | SHAKEN 709J | 11&#160;Jul&#160;23&#160;05:00&#160;UTC | false | [view](CERTS/ecc78184d11065338b9f698d81e8da92b2dc21251e414a01276b130862b5a5ab/README.md) |
| 12&#160;Apr&#160;23&#160;21:37&#160;UTC | SHAKEN 709J | 11&#160;Jul&#160;23&#160;05:00&#160;UTC | false | [view](CERTS/390ec6256c934f1af3c55af054de2c077db2698d3f3895d27e567094133c33d8/README.md) |
| 14&#160;Apr&#160;23&#160;16:55&#160;UTC | SHAKEN 683G | 13&#160;Jul&#160;23&#160;16:55&#160;UTC | false | [view](CERTS/2ba0ded53149fda9868572f31c430f49ec68abf3d03a6dc6a560d4de6c5e6b7d/README.md) |
| 14&#160;Apr&#160;23&#160;17:14&#160;UTC | SHAKEN 683G | 13&#160;Jul&#160;23&#160;17:14&#160;UTC | false | [view](CERTS/e16dffbe1d32b721aa736a056e30b96018f58fe03b7aa6058dae2dcb891cb56b/README.md) |
| 14&#160;Apr&#160;23&#160;17:18&#160;UTC | SHAKEN 683G | 13&#160;Jul&#160;23&#160;17:18&#160;UTC | false | [view](CERTS/264ffe0d52c568655f5e24b000a009441a4d8ee9e79162485c9a4d9972b8c997/README.md) |
| 17&#160;Apr&#160;23&#160;18:04&#160;UTC | SHAKEN 709J | 16&#160;Jul&#160;23&#160;05:00&#160;UTC | false | [view](CERTS/d263df10d2af09bfc4477a75a3e7cce9c9c175cd944a2e246035ea055861c8a6/README.md) |
| 01&#160;May&#160;23&#160;19:02&#160;UTC | SHAKEN 709J | 30&#160;Jul&#160;23&#160;19:02&#160;UTC | false | [view](CERTS/5c7488e9fcf0dd53f5f8bd1ffb6c53437367cc28cdc08d738d29f9c268dae4e3/README.md) |
| 01&#160;May&#160;23&#160;19:33&#160;UTC | SHAKEN 709J | 30&#160;Jul&#160;23&#160;19:33&#160;UTC | false | [view](CERTS/e9dfa99ca40bb4651ba594c4416682b31ae76a78a79a5c7b10d830ab23c1b3f4/README.md) |
| 01&#160;May&#160;23&#160;20:07&#160;UTC | SHAKEN 709J | 30&#160;Jul&#160;23&#160;20:07&#160;UTC | false | [view](CERTS/e8cbdc5ddbc3b315fe0bbb6ff31ab878778efb1cb1350262d6d2c4e9c20b7814/README.md) |
| 31&#160;May&#160;23&#160;09:38&#160;UTC | SHAKEN 042K | 29&#160;Aug&#160;23&#160;04:00&#160;UTC | false | [view](CERTS/258a0a421ecda377dabfcdd1ff0bb7a78f511a010933e96ad98caef068048b96/README.md) |
| 31&#160;May&#160;23&#160;20:12&#160;UTC | SHAKEN 048K | 29&#160;Aug&#160;23&#160;05:00&#160;UTC | false | [view](CERTS/0f41bd55c10d187e604511897636752352dce7a2d23c6e93b9a42ed3fa435e98/README.md) |
| 31&#160;May&#160;23&#160;20:59&#160;UTC | SHAKEN 048K | 29&#160;Aug&#160;23&#160;00:00&#160;UTC | false | [view](CERTS/10ffc87300eb5c07f00cc4a31c2137328cf4d518bd817a32a4ceaad93c32fc44/README.md) |
| 31&#160;May&#160;23&#160;21:06&#160;UTC | SHAKEN 066K | 29&#160;Aug&#160;23&#160;05:00&#160;UTC | false | [view](CERTS/bc36f9c42bf955662324a71c1fe297d4e25aa792f6600d85038872a95a26bc83/README.md) |
| 31&#160;May&#160;23&#160;21:08&#160;UTC | SHAKEN 031K | 29&#160;Aug&#160;23&#160;05:00&#160;UTC | false | [view](CERTS/263568a39c02254818ccc3b243a096aaa67e428512a6d1658d5be77bcf40d46c/README.md) |
| 31&#160;May&#160;23&#160;21:08&#160;UTC | SHAKEN 073K | 29&#160;Aug&#160;23&#160;06:00&#160;UTC | false | [view](CERTS/80719e892811da995951737f6f6f3596fa8de0f3ba124d94f04b891df52aab49/README.md) |
| 31&#160;May&#160;23&#160;21:11&#160;UTC | SHAKEN 039K | 29&#160;Aug&#160;23&#160;05:00&#160;UTC | false | [view](CERTS/39eb88adac142f3ee0f7ea50be124a6651e12ff155a908ae6a07a5e6120f50ac/README.md) |
| 31&#160;May&#160;23&#160;21:12&#160;UTC | SHAKEN 115K | 29&#160;Aug&#160;23&#160;04:00&#160;UTC | false | [view](CERTS/74275d0df578ea675a89ebc72e4303308d528281cd64b645828657a036c67f14/README.md) |
| 31&#160;May&#160;23&#160;21:12&#160;UTC | SHAKEN 223K | 29&#160;Aug&#160;23&#160;07:00&#160;UTC | false | [view](CERTS/7002c573274d9105533e7457d24f412767e49a4bd0d66ca1ca7c8fbb0a5062c3/README.md) |
| 31&#160;May&#160;23&#160;21:13&#160;UTC | SHAKEN 110K | 29&#160;Aug&#160;23&#160;06:00&#160;UTC | false | [view](CERTS/3c50d75f5e51879a00d00c6b99263cfdaf1ce400122f7b8588c349cf08d86ab9/README.md) |
| 31&#160;May&#160;23&#160;21:17&#160;UTC | SHAKEN 148K | 29&#160;Aug&#160;23&#160;04:00&#160;UTC | false | [view](CERTS/12f458f6182d69dad8f8807e7150312f0ae9cd23838528a5b4d39a5b842bacf6/README.md) |
| 31&#160;May&#160;23&#160;21:28&#160;UTC | SHAKEN 148K | 29&#160;Aug&#160;23&#160;04:00&#160;UTC | false | [view](CERTS/327cdc4972bda2eef6b6fd84b302faa2bd15e22a521102e230957b86fd1b1de4/README.md) |
| 31&#160;May&#160;23&#160;21:36&#160;UTC | SHAKEN 709J | 29&#160;Aug&#160;23&#160;05:00&#160;UTC | false | [view](CERTS/d0c978fad4fc5c1d2854cddfe5165bbe18509bb334bf1b45fb72dd25db5d7c06/README.md) |

#### CA Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 03&#160;May&#160;22&#160;15:31&#160;UTC | Martini Security SHAKEN R1 | 03&#160;May&#160;47&#160;21:31&#160;UTC | false | [view](CERTS/6077e368b0a0e4b6076eaa07ce67d6652ef310c4757776b76af84a6a9e003cdf/README.md) |
| 03&#160;May&#160;22&#160;15:41&#160;UTC | Martini Security SHAKEN G1 | 02&#160;May&#160;27&#160;15:41&#160;UTC | true | [view](CERTS/46e98b1599688c83928b66bcdd49e723ee4207128b5ba4488046edf2879970ef/README.md) |
| 02&#160;Oct&#160;22&#160;10:40&#160;UTC | Martini Security SHAKEN G2 | 01&#160;Oct&#160;27&#160;10:40&#160;UTC | false | [view](CERTS/bf818ddbd3ae492e4a85331b85b52f4d2cdef8287bf910b59e247b6c132fa7fd/README.md) |


Generated: 02 Jun 23 01:12 UTC