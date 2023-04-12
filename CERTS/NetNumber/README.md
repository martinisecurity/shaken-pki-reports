# STIR/SHAKEN CA Ecosystem Compliance

## NetNumber

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 15 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 10 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 5 certificates being tested against the remaining rules
- 5.40 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 0.00% of certificates contain one or more Warning level issue
- 100.00% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 84 days is the average remaining validity for the certificates in the corpus
- 99 days is the average initial validity for the certificates in the corpus
- 4 certificates expire in the next 30 days
- 1.00 average number of unexpired certificates per OCN observed
- 5 unique OCNs observed in unexpired and valid certificate corpus

| Instances | Test | Source |
|-----------|------|--------|
| 5 | [e_atis_certificate_policies](ISSUES/e_atis_certificate_policies/README.md) | ATIS1000080 |
| 3 | [e_atis_crl_distribution](ISSUES/e_atis_crl_distribution/README.md) | ATIS1000080 |
| 4 | [e_atis_subject_cn](ISSUES/e_atis_subject_cn/README.md) | ATIS1000080 |
| 5 | [e_us_cp_ambiguous_identifier](ISSUES/e_us_cp_ambiguous_identifier/README.md) | US_SHAKEN_CP |
| 5 | [e_us_cp_subject_sn](ISSUES/e_us_cp_subject_sn/README.md) | US_SHAKEN_CP |
| 5 | [n_atis_certificate_policy_critical](ISSUES/n_atis_certificate_policy_critical/README.md) | ATIS1000080 |

#### CA Certificates

- 3 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 0 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 3 certificates being tested against the remaining rules
- 3.00 issues on average found in unexpired, trusted, and non-compliant certificates
- 66.67% of certificates contain one or more Error level issue
- 66.67% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 100.00% of certificates are too old to be assessed against currently enforced expectations
- 7342 days is the average remaining validity for the certificates in the corpus
- 6935 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days

| Instances | Test | Source |
|-----------|------|--------|
| 2 | [e_atis_ca_signature_algorithm](ISSUES/e_atis_ca_signature_algorithm/README.md) | ATIS1000080 |
| 2 | [e_atis_ca_subject_public_key](ISSUES/e_atis_ca_subject_public_key/README.md) | ATIS1000080 |
| 2 | [w_shaken_ca_subject_rdn_unknown](ISSUES/w_shaken_ca_subject_rdn_unknown/README.md) | SHAKEN_PKI_BEST_PRACTICES |

### Details

#### Leaf Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 10&#160;Jun&#160;22&#160;00:00&#160;UTC | Plivo Inc | 09&#160;Jun&#160;23&#160;00:00&#160;UTC | true | [view](CERTS/7dc750fb7aa68d2b67b8dbc89f65217f92db54504685058be016638011adf8bf/README.md) |
| 04&#160;Mar&#160;23&#160;15:47&#160;UTC | Number Access LLC SHAKEN 343J | 15&#160;Apr&#160;23&#160;15:47&#160;UTC | true | [view](CERTS/5c2a8cbaa1d3792fc657a359988afde21a9881410e255b83c684e24e90c28c61/README.md) |
| 16&#160;Mar&#160;23&#160;00:00&#160;UTC | HD CARRIER LLC | 14&#160;Apr&#160;23&#160;23:59&#160;UTC | true | [view](CERTS/ef561f10277effd73ea38baf8b79d3860a5f6de4a7ffa440630ac0a8da4ac3ae/README.md) |
| 20&#160;Mar&#160;23&#160;17:58&#160;UTC | Google SHAKEN cert 969H | 19&#160;Apr&#160;23&#160;17:58&#160;UTC | true | [view](CERTS/6920f54a1ec0d3e2e114bc73efa64820b84b7bb75fe4c2e562184e1f46375a70/README.md) |
| 23&#160;Mar&#160;23&#160;15:00&#160;UTC | Baltimore-Washington Telephone Company SHAKEN cert 8697 | 23&#160;Apr&#160;23&#160;14:59&#160;UTC | true | [view](CERTS/5f9e936b92edadbabbb6778755b4af6b7765d07aa08f39463de490e0f4ed628e/README.md) |

#### CA Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 12&#160;Jul&#160;21&#160;23:25&#160;UTC | NetNumber SHAKEN Root CA | 07&#160;Jul&#160;41&#160;23:25&#160;UTC | false | [view](CERTS/7ac80e8481ecb019dc95484016842db78686069efbc0f703e7f39310217b6157/README.md) |
| 27&#160;Sep&#160;21&#160;19:45&#160;UTC | NetNumber SHAKEN Root CA 1 | 21&#160;Sep&#160;46&#160;19:45&#160;UTC | true | [view](CERTS/2dd1386ca717f31d550b35b9bce9daa9b02483bcdb98bdfcfca07202276136d7/README.md) |
| 29&#160;Sep&#160;21&#160;13:22&#160;UTC | NetNumber SHAKEN Root Intermediate CA 1 | 26&#160;Sep&#160;33&#160;13:22&#160;UTC | true | [view](CERTS/e449803766edf02ab50b034dd7e89e54efd332cce87688a032f89b340d039878/README.md) |


Generated: 12 Apr 23 01:46 UTC