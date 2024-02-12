# STIR/SHAKEN CA Ecosystem Compliance

## Peeringhub

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 43 certificates were included in the corpus being tested
- 9 certificates in the corpus were skipped because they are duplicates
- 14 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 20 certificates being tested against the remaining rules
- 1.00 issues on average found in unexpired, trusted, and non-compliant certificates
- 55.00% of certificates contain one or more Error level issue
- 0.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 220 days is the average remaining validity for the certificates in the corpus
- 230 days is the average initial validity for the certificates in the corpus
- 7 certificates expire in the next 30 days
- 1.00 average number of unexpired certificates per OCN observed
- 20 unique OCNs observed in unexpired and valid certificate corpus

| Instances | Test | Source |
|-----------|------|--------|
| 11 | [e_shaken_certificate_policies_id](ISSUES/e_shaken_certificate_policies_id/README.md) | US_SHAKEN_CP |

#### CA Certificates

- 2 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 0 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 2 certificates being tested against the remaining rules
- 1.00 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 0.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 100.00% of certificates are too old to be assessed against currently enforced expectations
- 5699 days is the average remaining validity for the certificates in the corpus
- 5475 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days

| Instances | Test | Source |
|-----------|------|--------|
| 1 | [e_atis_subject_cn_ca](ISSUES/e_atis_subject_cn_ca/README.md) | ATIS1000080 |
| 1 | [e_shaken_certificate_policies_id_ca](ISSUES/e_shaken_certificate_policies_id_ca/README.md) | US_SHAKEN_CP |

### Details

#### Leaf Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 25&#160;May&#160;23&#160;20:39&#160;UTC | SHAKEN 088K 2023-05 | 24&#160;May&#160;24&#160;20:39&#160;UTC | true | [view](CERTS/1d7fe7341654ec5e2f3f89be1d2e410b4b1d979acb0b0b6b3ff6db0c9382eccb/README.md) |
| 06&#160;Jul&#160;23&#160;23:22&#160;UTC | On Air Telecom LLC SHAKEN 861J | 05&#160;Jul&#160;24&#160;23:22&#160;UTC | true | [view](CERTS/c339bf7c529ac881591c7f1564e98d85608719029f9a467ca3683b007041b677/README.md) |
| 03&#160;Aug&#160;23&#160;20:55&#160;UTC | Perfect Network LLC SHAKEN 458K 01 | 02&#160;Aug&#160;24&#160;20:55&#160;UTC | true | [view](CERTS/7bac44209231f9e843e44d1f88644dae698a5841e993a4dd998230f141300b41/README.md) |
| 04&#160;Aug&#160;23&#160;10:28&#160;UTC | Access Tandem Inc SHAKEN 731J | 03&#160;Aug&#160;24&#160;10:28&#160;UTC | true | [view](CERTS/8377b918743929fe382e6275a43e155380da3d51102e75459d5abf447f89633d/README.md) |
| 04&#160;Aug&#160;23&#160;10:36&#160;UTC | Voiceterm Inc SHAKEN 240K | 03&#160;Aug&#160;24&#160;10:36&#160;UTC | true | [view](CERTS/c28277e11e1f684965db049437ab4e6ab53b132b9c65cbc521f2521558f6793a/README.md) |
| 17&#160;Aug&#160;23&#160;22:05&#160;UTC | SHAKEN 788J 1692309910 | 16&#160;Aug&#160;24&#160;22:05&#160;UTC | true | [view](CERTS/6ff13b878806b2584400b1ecc88a909d07fd7480d4c72db64c430f52d8c9909c/README.md) |
| 24&#160;Aug&#160;23&#160;18:23&#160;UTC | Phonetime Inc SHAKEN 602K | 12&#160;Jul&#160;24&#160;14:12&#160;UTC | true | [view](CERTS/309c590e746ab032801e90a501b6a34d310d6435c77d90fec9ee5299c5c8e5ce/README.md) |
| 12&#160;Sep&#160;23&#160;18:41&#160;UTC | Jaintel SHAKEN 586K | 11&#160;Sep&#160;24&#160;18:41&#160;UTC | true | [view](CERTS/35aa37027485bf13ef7ad46d4c344e563d22a765975a62e193ecfc94841a3d27/README.md) |
| 13&#160;Sep&#160;23&#160;13:40&#160;UTC | Teleinx LLC SHAKEN 744J | 12&#160;Sep&#160;24&#160;13:40&#160;UTC | true | [view](CERTS/21f551ad9783b07beb40b5096a12e5fc37845b5f354b81ed89bde51f070404d6/README.md) |
| 24&#160;Oct&#160;23&#160;22:46&#160;UTC | Voipedia SHAKEN 712K | 23&#160;Oct&#160;24&#160;22:46&#160;UTC | true | [view](CERTS/b6029bec3835bab5b57926877ad0ad08198b25153da52b8c7bb40a43ee323b9f/README.md) |
| 30&#160;Oct&#160;23&#160;00:00&#160;UTC | VOCALTRANSIT SHAKEN 783J | 22&#160;May&#160;24&#160;00:00&#160;UTC | true | [view](CERTS/e8c92f75c60632193af34b988570281cbbcbafce7a9c3ae534b8406d62aab179/README.md) |
| 22&#160;Jan&#160;24&#160;16:17&#160;UTC | Celerity Telecom SHAKEN 469K | 23&#160;Feb&#160;24&#160;15:39&#160;UTC | false | [view](CERTS/eb5226368fc5bc4dcc0d5cbe0fd0cc16113f92acd5ce711d3a39ab3cd9e99992/README.md) |
| 24&#160;Jan&#160;24&#160;05:40&#160;UTC | INTENEXT TELECOM LLC SHAKEN 650K 4 | 20&#160;Feb&#160;24&#160;18:40&#160;UTC | false | [view](CERTS/b96eb0b940c7d0f228f7a0177d9427bf708cc2defe2ff2f47ffdc156a7835cc8/README.md) |
| 25&#160;Jan&#160;24&#160;01:10&#160;UTC | DiDCentral SHAKEN 756J | 29&#160;Feb&#160;24&#160;00:09&#160;UTC | false | [view](CERTS/255ce42aac1c7eea55b57bcd76bb7b42d6a06090449dc9e61c4d8b75f24194db/README.md) |
| 26&#160;Jan&#160;24&#160;21:15&#160;UTC | ARit services LLC_1706303744312 SHAKEN 827K | 04&#160;Mar&#160;24&#160;14:16&#160;UTC | false | [view](CERTS/89d1eea7a07c4606b1449185bdd9d70959917321987c774db4230839ee76e4a6/README.md) |
| 29&#160;Jan&#160;24&#160;05:50&#160;UTC | MAQS SHAKEN 255K | 04&#160;Mar&#160;24&#160;20:13&#160;UTC | false | [view](CERTS/1ac9e949c222ec086fcb7cc1c69f10ae55f2ba1e677e5830c3b12d3e2ba6cf77/README.md) |
| 01&#160;Feb&#160;24&#160;16:34&#160;UTC | Tanya David LLC SHAKEN 774K | 07&#160;Mar&#160;24&#160;19:23&#160;UTC | false | [view](CERTS/69207e6deef4272f49a7b321bfa697627492d7e43ce79434ed1a69c75ed2e52f/README.md) |
| 06&#160;Feb&#160;24&#160;05:15&#160;UTC | Georgialina Networks 020624 SHAKEN 665K | 08&#160;Mar&#160;24&#160;15:23&#160;UTC | false | [view](CERTS/4b450264860bff045f56e2cd84b0b9dee4a1820b2dba22cc8f7a44140a484f7f/README.md) |
| 08&#160;Feb&#160;24&#160;20:24&#160;UTC | Instacall SHAKEN 281K | 22&#160;Jan&#160;25&#160;19:11&#160;UTC | false | [view](CERTS/e280f42a746af003040e83e1e381a1c89a76e7160f3a4649b14543c5e2a1db0d/README.md) |
| 12&#160;Feb&#160;24&#160;00:00&#160;UTC | Losh Communications, Inc SHAKEN 149K 2024-02-12_000001 | 24&#160;Aug&#160;24&#160;19:56&#160;UTC | false | [view](CERTS/f59b2d9ab7057ce6d98cc07144bbdbfeda4bdd69efd245b84b86cac9d36e0bbe/README.md) |

#### CA Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 17&#160;Dec&#160;20&#160;15:31&#160;UTC | Peeringhub Inc Root CA | 12&#160;Dec&#160;40&#160;15:31&#160;UTC | true | [view](CERTS/0ad4adc0b4d93fdb0e628c577020c73b8a5caff750e7e499f80ee2ab362a3f6a/README.md) |
| 22&#160;Jun&#160;22&#160;22:45&#160;UTC | Peeringhub Inc SHAKEN Intermediate CA 2 | 19&#160;Jun&#160;32&#160;22:45&#160;UTC | true | [view](CERTS/f00871963a40b04269c4b019968e42f9f40964cbfb512ff5342307e9942874ce/README.md) |


Generated: 12 Feb 24 19:26 UTC