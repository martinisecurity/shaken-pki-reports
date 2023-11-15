# STIR/SHAKEN CA Ecosystem Compliance

## Peeringhub

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 32 certificates were included in the corpus being tested
- 7 certificates in the corpus were skipped because they are duplicates
- 9 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 16 certificates being tested against the remaining rules
- 3.06 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 6.25% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 262 days is the average remaining validity for the certificates in the corpus
- 277 days is the average initial validity for the certificates in the corpus
- 3 certificates expire in the next 30 days
- 1.00 average number of unexpired certificates per OCN observed
- 16 unique OCNs observed in unexpired and valid certificate corpus

| Instances | Test | Source |
|-----------|------|--------|
| 16 | [e_atis_certificate_policies](ISSUES/e_atis_certificate_policies/README.md) | ATIS1000080 |
| 16 | [e_us_cp_ambiguous_identifier](ISSUES/e_us_cp_ambiguous_identifier/README.md) | US_SHAKEN_CP |
| 16 | [e_us_cp_subject_sn](ISSUES/e_us_cp_subject_sn/README.md) | US_SHAKEN_CP |
| 1 | [w_shaken_subject_rdn_unknown](ISSUES/w_shaken_subject_rdn_unknown/README.md) | SHAKEN_PKI_BEST_PRACTICES |

#### CA Certificates

- 2 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 0 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 2 certificates being tested against the remaining rules
- 2.50 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 100.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 50.00% of certificates are too old to be assessed against currently enforced expectations
- 5729 days is the average remaining validity for the certificates in the corpus
- 5475 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days

| Instances | Test | Source |
|-----------|------|--------|
| 1 | [e_atis_ca_certificate_policies](ISSUES/e_atis_ca_certificate_policies/README.md) | ATIS1000080 |
| 2 | [e_us_cp_ca_key_usage_crl_sign](ISSUES/e_us_cp_ca_key_usage_crl_sign/README.md) | US_SHAKEN_CP |
| 2 | [w_shaken_ca_subject_rdn_unknown](ISSUES/w_shaken_ca_subject_rdn_unknown/README.md) | SHAKEN_PKI_BEST_PRACTICES |

### Details

#### Leaf Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 25&#160;May&#160;23&#160;20:39&#160;UTC | SHAKEN 088K 2023-05 | 24&#160;May&#160;24&#160;20:39&#160;UTC | true | [view](CERTS/1d7fe7341654ec5e2f3f89be1d2e410b4b1d979acb0b0b6b3ff6db0c9382eccb/README.md) |
| 21&#160;Jun&#160;23&#160;00:42&#160;UTC | Meta-lynk Telecom SHAKEN 442K | 30&#160;Jan&#160;24&#160;00:00&#160;UTC | true | [view](CERTS/723af9321b9721ed8c7efcecfc7c6dbb59b3de3b957d707fb030fa6006f29b8c/README.md) |
| 06&#160;Jul&#160;23&#160;23:22&#160;UTC | On Air Telecom LLC SHAKEN 861J | 05&#160;Jul&#160;24&#160;23:22&#160;UTC | true | [view](CERTS/c339bf7c529ac881591c7f1564e98d85608719029f9a467ca3683b007041b677/README.md) |
| 03&#160;Aug&#160;23&#160;20:55&#160;UTC | Perfect Network LLC SHAKEN 458K 01 | 02&#160;Aug&#160;24&#160;20:55&#160;UTC | true | [view](CERTS/7bac44209231f9e843e44d1f88644dae698a5841e993a4dd998230f141300b41/README.md) |
| 04&#160;Aug&#160;23&#160;10:28&#160;UTC | Access Tandem Inc SHAKEN 731J | 03&#160;Aug&#160;24&#160;10:28&#160;UTC | true | [view](CERTS/8377b918743929fe382e6275a43e155380da3d51102e75459d5abf447f89633d/README.md) |
| 04&#160;Aug&#160;23&#160;10:36&#160;UTC | Voiceterm Inc SHAKEN 240K | 03&#160;Aug&#160;24&#160;10:36&#160;UTC | true | [view](CERTS/c28277e11e1f684965db049437ab4e6ab53b132b9c65cbc521f2521558f6793a/README.md) |
| 17&#160;Aug&#160;23&#160;22:05&#160;UTC | SHAKEN 788J 1692309910 | 16&#160;Aug&#160;24&#160;22:05&#160;UTC | true | [view](CERTS/6ff13b878806b2584400b1ecc88a909d07fd7480d4c72db64c430f52d8c9909c/README.md) |
| 24&#160;Aug&#160;23&#160;18:23&#160;UTC | Phonetime Inc SHAKEN 602K | 12&#160;Jul&#160;24&#160;14:12&#160;UTC | true | [view](CERTS/309c590e746ab032801e90a501b6a34d310d6435c77d90fec9ee5299c5c8e5ce/README.md) |
| 12&#160;Sep&#160;23&#160;18:41&#160;UTC | Jaintel SHAKEN 586K | 11&#160;Sep&#160;24&#160;18:41&#160;UTC | true | [view](CERTS/35aa37027485bf13ef7ad46d4c344e563d22a765975a62e193ecfc94841a3d27/README.md) |
| 13&#160;Sep&#160;23&#160;13:40&#160;UTC | Teleinx LLC SHAKEN 744J | 12&#160;Sep&#160;24&#160;13:40&#160;UTC | true | [view](CERTS/21f551ad9783b07beb40b5096a12e5fc37845b5f354b81ed89bde51f070404d6/README.md) |
| 24&#160;Oct&#160;23&#160;22:46&#160;UTC | Voipedia SHAKEN 712K | 23&#160;Oct&#160;24&#160;22:46&#160;UTC | true | [view](CERTS/b6029bec3835bab5b57926877ad0ad08198b25153da52b8c7bb40a43ee323b9f/README.md) |
| 27&#160;Oct&#160;23&#160;10:01&#160;UTC | Instacall SHAKEN 281K 1 | 02&#160;Dec&#160;23&#160;18:24&#160;UTC | true | [view](CERTS/b49dc4d6b6a1700f6d2731dbd7b3eff9cb75cacaf5780ae08ac90ddcb32fac20/README.md) |
| 30&#160;Oct&#160;23&#160;00:00&#160;UTC | VOCALTRANSIT SHAKEN 783J | 22&#160;May&#160;24&#160;00:00&#160;UTC | true | [view](CERTS/e8c92f75c60632193af34b988570281cbbcbafce7a9c3ae534b8406d62aab179/README.md) |
| 30&#160;Oct&#160;23&#160;10:16&#160;UTC | TalkAsiaVoip LLC SHAKEN 198K | 07&#160;Dec&#160;23&#160;07:31&#160;UTC | true | [view](CERTS/7e7a5de2d0a8c98a58e98c631283e52e48b18bae23449d74e3e8f3f3807bb24b/README.md) |
| 01&#160;Nov&#160;23&#160;01:05&#160;UTC | MAQSSolutions SHAKEN 255K | 03&#160;Dec&#160;23&#160;20:13&#160;UTC | true | [view](CERTS/39eedca0fda366ccd88fc4140394b807ea5dbae5a5bd4fc130958d6855f93948/README.md) |
| 15&#160;Nov&#160;23&#160;00:00&#160;UTC | Losh Communications, Inc SHAKEN 149K 2023-11-15_000001 | 24&#160;Aug&#160;24&#160;19:56&#160;UTC | true | [view](CERTS/ce2ddf16eeeb3c1584777248d4c63e21300349cb62410a9b99ecfc0bf60dc315/README.md) |

#### CA Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 17&#160;Dec&#160;20&#160;15:31&#160;UTC | Peeringhub Inc Root CA | 12&#160;Dec&#160;40&#160;15:31&#160;UTC | true | [view](CERTS/0ad4adc0b4d93fdb0e628c577020c73b8a5caff750e7e499f80ee2ab362a3f6a/README.md) |
| 22&#160;Jun&#160;22&#160;22:45&#160;UTC | Peeringhub Inc SHAKEN Intermediate CA 2 | 19&#160;Jun&#160;32&#160;22:45&#160;UTC | true | [view](CERTS/f00871963a40b04269c4b019968e42f9f40964cbfb512ff5342307e9942874ce/README.md) |


Generated: 15 Nov 23 17:17 UTC