# STIR/SHAKEN CA Ecosystem Compliance

## Peeringhub

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 10 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 1 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 9 certificates being tested against the remaining rules
- 3.44 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 44.44% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 166 days is the average remaining validity for the certificates in the corpus
- 183 days is the average initial validity for the certificates in the corpus
- 4 certificates expire in the next 30 days

| Instances | Test Status | Source |
|-----------|-------------|--------|
| 9 | [e_atis_certificate_policies](ISSUES/e_atis_certificate_policies/README.md) | ATIS1000080 |
| 9 | [e_us_cp_ambiguous_identifier](ISSUES/e_us_cp_ambiguous_identifier/README.md) | US_SHAKEN_CP |
| 9 | [e_us_cp_subject_sn](ISSUES/e_us_cp_subject_sn/README.md) | US_SHAKEN_CP |
| 4 | [w_shaken_subject_rdn_unknown](ISSUES/w_shaken_subject_rdn_unknown/README.md) | SHAKEN_PKI_BEST_PRACTICES |

#### CA Certificates

- 2 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 0 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 2 certificates being tested against the remaining rules
- 2.50 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 100.00% of certificates contain one or more Warning level issue
- 50.00% of certificates contain one or more Notice level issue
- 50.00% of certificates are too old to be assessed against currently enforced expectations
- 5855 days is the average remaining validity for the certificates in the corpus
- 5475 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days

| Instances | Test Status | Source |
|-----------|-------------|--------|
| 1 | [e_atis_ca_certificate_policies](ISSUES/e_atis_ca_certificate_policies/README.md) | ATIS1000080 |
| 2 | [e_us_cp_ca_key_usage_crl_sign](ISSUES/e_us_cp_ca_key_usage_crl_sign/README.md) | US_SHAKEN_CP |
| 2 | [w_shaken_ca_subject_rdn_unknown](ISSUES/w_shaken_ca_subject_rdn_unknown/README.md) | SHAKEN_PKI_BEST_PRACTICES |

### Details

#### Leaf Certificates

| Created at | Subject | Problems | Link |
|------------|---------|----------|------|
| 29 Jul 22 16:41 UTC | ATI SHAKEN 731J | true | [view](CERTS/bbc0a32743f1e659eff17172cbef96ee8cc3000aeb1dda85e1a47d47391ec3ab/README.md) |
| 31 Jul 22 09:09 UTC | Voiceterm SHAKEN 240K | true | [view](CERTS/84819934eb8b7f347d7365133dba4376d162d283fa9d09858aa1706ec88487c6/README.md) |
| 26 Aug 22 23:31 UTC | Teleinx SHAKEN 744J | true | [view](CERTS/743032377136fc18e443399c5fc57e36a5706188f141a522438a806143997925/README.md) |
| 15 Oct 22 00:00 UTC | VOCALTRANSIT SHAKEN 783J | true | [view](CERTS/81b78fff8a772249d72d4854d97672d7ac69a83c4900beaac699d28d220d8c13/README.md) |
| 19 Oct 22 19:24 UTC | TalkAsiaVoip LLC SHAKEN 198K | true | [view](CERTS/b49964274f962d87bf69c69b8c2efd07cd4de3d20155339e41ab2752477fd19b/README.md) |
| 01 Nov 22 19:16 UTC | Definitely SHAKEN 709J but not SHAKEN 0007 | true | [view](CERTS/8960e7cbd26618777e4018fdaf9853ff67f176b10954cb37596a75c2dd0a4ddb/README.md) |
| 01 Nov 22 19:26 UTC | Definitely SHAKEN 709J but not SHAKEN 0007 | true | [view](CERTS/148fa003df4d2ae06a49e002c44e3dd1a78feedd963cfb762b7a07829c496be8/README.md) |
| 01 Nov 22 19:59 UTC | Definitely SHAKEN 709J but not SHAKEN 0007 | true | [view](CERTS/d35cdea11102a752f34309c8ecae597b4b7c1caf7f41dea46415471eb5317cce/README.md) |
| 02 Nov 22 07:51 UTC | Apex Telecom LLC SHAKEN 288K | true | [view](CERTS/e64593f4d9a81236af33c40c227d728e49418be50e05be2a8b65ec549a3b9275/README.md) |

#### CA Certificates

| Created at | Subject | Problems | Link |
|------------|---------|----------|------|
| 17 Dec 20 15:31 UTC | Peeringhub Inc Root CA | true | [view](CERTS/0ad4adc0b4d93fdb0e628c577020c73b8a5caff750e7e499f80ee2ab362a3f6a/README.md) |
| 22 Jun 22 22:45 UTC | Peeringhub Inc SHAKEN Intermediate CA 2 | true | [view](CERTS/f00871963a40b04269c4b019968e42f9f40964cbfb512ff5342307e9942874ce/README.md) |


Generated: 02 Nov 22 15:10 UTC