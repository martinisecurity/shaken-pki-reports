# STIR/SHAKEN CA Ecosystem Compliance

## Peeringhub

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 18 certificates were included in the corpus being tested
- 3 certificates in the corpus were skipped because they are duplicates
- 1 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 14 certificates being tested against the remaining rules
- 3.07 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 7.14% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 230 days is the average remaining validity for the certificates in the corpus
- 246 days is the average initial validity for the certificates in the corpus
- 5 certificates expire in the next 30 days
- 1.00 average number of unexpired certificates per OCN observed
- 14 unique OCNs observed in unexpired and valid certificate corpus

| Instances | Test | Source |
|-----------|------|--------|
| 14 | [e_atis_certificate_policies](ISSUES/e_atis_certificate_policies/README.md) | ATIS1000080 |
| 14 | [e_us_cp_ambiguous_identifier](ISSUES/e_us_cp_ambiguous_identifier/README.md) | US_SHAKEN_CP |
| 14 | [e_us_cp_subject_sn](ISSUES/e_us_cp_subject_sn/README.md) | US_SHAKEN_CP |
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
- 5758 days is the average remaining validity for the certificates in the corpus
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
| 26&#160;Aug&#160;22&#160;23:31&#160;UTC | Teleinx SHAKEN 744J | 26&#160;Aug&#160;23&#160;23:31&#160;UTC | true | [view](CERTS/743032377136fc18e443399c5fc57e36a5706188f141a522438a806143997925/README.md) |
| 02&#160;Nov&#160;22&#160;07:51&#160;UTC | Apex Telecom LLC SHAKEN 288K | 02&#160;Nov&#160;23&#160;00:54&#160;UTC | true | [view](CERTS/e64593f4d9a81236af33c40c227d728e49418be50e05be2a8b65ec549a3b9275/README.md) |
| 16&#160;Nov&#160;22&#160;00:03&#160;UTC | Televoip SHAKEN 438K | 11&#160;Nov&#160;23&#160;00:08&#160;UTC | true | [view](CERTS/2e22094a03b8d7e2d2f37cb2ea5d5778d122fdbc3d0c83e3f8f984c8cf03a2e3/README.md) |
| 25&#160;May&#160;23&#160;20:39&#160;UTC | SHAKEN 088K 2023-05 | 24&#160;May&#160;24&#160;20:39&#160;UTC | true | [view](CERTS/1d7fe7341654ec5e2f3f89be1d2e410b4b1d979acb0b0b6b3ff6db0c9382eccb/README.md) |
| 01&#160;Jun&#160;23&#160;00:00&#160;UTC | VOCALTRANSIT SHAKEN 783J v3 | 15&#160;Nov&#160;23&#160;00:00&#160;UTC | true | [view](CERTS/7dc3f9db0b51eed136cea1a1de68f558a7efbac78d7381a341b68e7838db3814/README.md) |
| 21&#160;Jun&#160;23&#160;00:42&#160;UTC | Meta-lynk Telecom SHAKEN 442K | 30&#160;Jan&#160;24&#160;00:00&#160;UTC | true | [view](CERTS/723af9321b9721ed8c7efcecfc7c6dbb59b3de3b957d707fb030fa6006f29b8c/README.md) |
| 06&#160;Jul&#160;23&#160;23:22&#160;UTC | On Air Telecom LLC SHAKEN 861J | 05&#160;Jul&#160;24&#160;23:22&#160;UTC | true | [view](CERTS/c339bf7c529ac881591c7f1564e98d85608719029f9a467ca3683b007041b677/README.md) |
| 25&#160;Jul&#160;23&#160;11:02&#160;UTC | TalkAsiaVoip LLC SHAKEN 198K | 26&#160;Aug&#160;23&#160;12:36&#160;UTC | true | [view](CERTS/6178b5372d816afdb4a44ab6c741edcf380673b717e62d95f6500af6dccb0165/README.md) |
| 26&#160;Jul&#160;23&#160;14:46&#160;UTC | INS72IPES SHAKEN 281K | 01&#160;Sep&#160;23&#160;18:24&#160;UTC | true | [view](CERTS/cdbe1296f50fc9bbedad6f7111cc6dfbd17e3b8cd2a365e607422494bc8a5d43/README.md) |
| 30&#160;Jul&#160;23&#160;22:14&#160;UTC | MAQSSolutions SHAKEN 255K | 02&#160;Sep&#160;23&#160;20:13&#160;UTC | true | [view](CERTS/9dbaed23f2f42b4965f10a32f36223fc23e14c32c66550925b692027cf48c04e/README.md) |
| 01&#160;Aug&#160;23&#160;21:07&#160;UTC | Celerity SHAKEN 469K | 01&#160;Sep&#160;23&#160;15:59&#160;UTC | true | [view](CERTS/018499a3ccef1a8c4efc335aa8f346f04fb8277a4511244052acefe8fe8a7b26/README.md) |
| 03&#160;Aug&#160;23&#160;20:55&#160;UTC | Perfect Network LLC SHAKEN 458K 01 | 02&#160;Aug&#160;24&#160;20:55&#160;UTC | true | [view](CERTS/7bac44209231f9e843e44d1f88644dae698a5841e993a4dd998230f141300b41/README.md) |
| 04&#160;Aug&#160;23&#160;10:28&#160;UTC | Access Tandem Inc SHAKEN 731J | 03&#160;Aug&#160;24&#160;10:28&#160;UTC | true | [view](CERTS/8377b918743929fe382e6275a43e155380da3d51102e75459d5abf447f89633d/README.md) |
| 04&#160;Aug&#160;23&#160;10:36&#160;UTC | Voiceterm Inc SHAKEN 240K | 03&#160;Aug&#160;24&#160;10:36&#160;UTC | true | [view](CERTS/c28277e11e1f684965db049437ab4e6ab53b132b9c65cbc521f2521558f6793a/README.md) |

#### CA Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 17&#160;Dec&#160;20&#160;15:31&#160;UTC | Peeringhub Inc Root CA | 12&#160;Dec&#160;40&#160;15:31&#160;UTC | true | [view](CERTS/0ad4adc0b4d93fdb0e628c577020c73b8a5caff750e7e499f80ee2ab362a3f6a/README.md) |
| 22&#160;Jun&#160;22&#160;22:45&#160;UTC | Peeringhub Inc SHAKEN Intermediate CA 2 | 19&#160;Jun&#160;32&#160;22:45&#160;UTC | true | [view](CERTS/f00871963a40b04269c4b019968e42f9f40964cbfb512ff5342307e9942874ce/README.md) |


Generated: 21 Aug 23 20:18 UTC