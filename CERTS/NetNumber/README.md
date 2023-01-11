# STIR/SHAKEN CA Ecosystem Compliance

## NetNumber

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 8 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 5 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 3 certificates being tested against the remaining rules
- 5.33 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 0.00% of certificates contain one or more Warning level issue
- 100.00% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 108 days is the average remaining validity for the certificates in the corpus
- 142 days is the average initial validity for the certificates in the corpus
- 2 certificates expire in the next 30 days
- 1.00 average number of unexpired certificates per OCN observed
- 3 unique OCNs observed in unexpired and valid certificate corpus

| Instances | Test | Source |
|-----------|------|--------|
| 3 | [e_atis_certificate_policies](ISSUES/e_atis_certificate_policies/README.md) | ATIS1000080 |
| 1 | [e_atis_crl_distribution](ISSUES/e_atis_crl_distribution/README.md) | ATIS1000080 |
| 3 | [e_atis_subject_cn](ISSUES/e_atis_subject_cn/README.md) | ATIS1000080 |
| 3 | [e_us_cp_ambiguous_identifier](ISSUES/e_us_cp_ambiguous_identifier/README.md) | US_SHAKEN_CP |
| 3 | [e_us_cp_subject_sn](ISSUES/e_us_cp_subject_sn/README.md) | US_SHAKEN_CP |
| 3 | [n_atis_certificate_policy_critical](ISSUES/n_atis_certificate_policy_critical/README.md) | ATIS1000080 |

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
- 7365 days is the average remaining validity for the certificates in the corpus
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
| 18&#160;Dec&#160;22&#160;14:59&#160;UTC | Google SHAKEN cert 969H | 17&#160;Jan&#160;23&#160;14:59&#160;UTC | true | [view](CERTS/6f97b41760bbb78a17309694eb34e5a0b896c78bfea22fae23731f3b2bc7bcea/README.md) |
| 26&#160;Dec&#160;22&#160;15:00&#160;UTC | Baltimore-Washington Telephone Company SHAKEN cert 8697 | 26&#160;Jan&#160;23&#160;14:59&#160;UTC | true | [view](CERTS/06404c88b7c84beeea562e1978a3763e7c0242cdd065f2577d6477c225bc061a/README.md) |

#### CA Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 12&#160;Jul&#160;21&#160;23:25&#160;UTC | NetNumber SHAKEN Root CA | 07&#160;Jul&#160;41&#160;23:25&#160;UTC | false | [view](CERTS/7ac80e8481ecb019dc95484016842db78686069efbc0f703e7f39310217b6157/README.md) |
| 27&#160;Sep&#160;21&#160;19:45&#160;UTC | NetNumber SHAKEN Root CA 1 | 21&#160;Sep&#160;46&#160;19:45&#160;UTC | true | [view](CERTS/2dd1386ca717f31d550b35b9bce9daa9b02483bcdb98bdfcfca07202276136d7/README.md) |
| 29&#160;Sep&#160;21&#160;13:22&#160;UTC | NetNumber SHAKEN Root Intermediate CA 1 | 26&#160;Sep&#160;33&#160;13:22&#160;UTC | true | [view](CERTS/e449803766edf02ab50b034dd7e89e54efd332cce87688a032f89b340d039878/README.md) |


Generated: 11 Jan 23 21:04 UTC