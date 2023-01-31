# STIR/SHAKEN CA Ecosystem Compliance

## NetNumber

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 10 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 6 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 4 certificates being tested against the remaining rules
- 5.25 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 0.00% of certificates contain one or more Warning level issue
- 100.00% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 95 days is the average remaining validity for the certificates in the corpus
- 117 days is the average initial validity for the certificates in the corpus
- 3 certificates expire in the next 30 days
- 1.00 average number of unexpired certificates per OCN observed
- 4 unique OCNs observed in unexpired and valid certificate corpus

| Instances | Test | Source |
|-----------|------|--------|
| 4 | [e_atis_certificate_policies](ISSUES/e_atis_certificate_policies/README.md) | ATIS1000080 |
| 2 | [e_atis_crl_distribution](ISSUES/e_atis_crl_distribution/README.md) | ATIS1000080 |
| 3 | [e_atis_subject_cn](ISSUES/e_atis_subject_cn/README.md) | ATIS1000080 |
| 4 | [e_us_cp_ambiguous_identifier](ISSUES/e_us_cp_ambiguous_identifier/README.md) | US_SHAKEN_CP |
| 4 | [e_us_cp_subject_sn](ISSUES/e_us_cp_subject_sn/README.md) | US_SHAKEN_CP |
| 4 | [n_atis_certificate_policy_critical](ISSUES/n_atis_certificate_policy_critical/README.md) | ATIS1000080 |

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
- 7360 days is the average remaining validity for the certificates in the corpus
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
| 02&#160;Jan&#160;23&#160;11:31&#160;UTC | Number Access LLC SHAKEN 343J | 13&#160;Feb&#160;23&#160;11:31&#160;UTC | true | [view](CERTS/95f6f8173b047c16711c37180aef55ed9542d62a9158c4ff9088fd648d95ade3/README.md) |
| 10&#160;Jan&#160;23&#160;15:08&#160;UTC | Google SHAKEN cert 969H | 09&#160;Feb&#160;23&#160;15:08&#160;UTC | true | [view](CERTS/4dc40d22f12a9db1bad2249a598a27c12049fe4eafe76583bce25d4ee8e6a01c/README.md) |
| 25&#160;Jan&#160;23&#160;15:00&#160;UTC | Baltimore-Washington Telephone Company SHAKEN cert 8697 | 25&#160;Feb&#160;23&#160;14:59&#160;UTC | true | [view](CERTS/699a1a5f7af369d88918c316d1a98a785ada5c369bdaeded06cb0536b5d1b217/README.md) |

#### CA Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 12&#160;Jul&#160;21&#160;23:25&#160;UTC | NetNumber SHAKEN Root CA | 07&#160;Jul&#160;41&#160;23:25&#160;UTC | false | [view](CERTS/7ac80e8481ecb019dc95484016842db78686069efbc0f703e7f39310217b6157/README.md) |
| 27&#160;Sep&#160;21&#160;19:45&#160;UTC | NetNumber SHAKEN Root CA 1 | 21&#160;Sep&#160;46&#160;19:45&#160;UTC | true | [view](CERTS/2dd1386ca717f31d550b35b9bce9daa9b02483bcdb98bdfcfca07202276136d7/README.md) |
| 29&#160;Sep&#160;21&#160;13:22&#160;UTC | NetNumber SHAKEN Root Intermediate CA 1 | 26&#160;Sep&#160;33&#160;13:22&#160;UTC | true | [view](CERTS/e449803766edf02ab50b034dd7e89e54efd332cce87688a032f89b340d039878/README.md) |


Generated: 31 Jan 23 17:51 UTC