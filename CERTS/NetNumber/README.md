# STIR/SHAKEN CA Ecosystem Compliance

## NetNumber

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 6 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 2 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 4 certificates being tested against the remaining rules
- 5.25 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 0.00% of certificates contain one or more Warning level issue
- 100.00% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 92 days is the average remaining validity for the certificates in the corpus
- 114 days is the average initial validity for the certificates in the corpus
- 3 certificates expire in the next 30 days

| Instances | Test | Source |
|-----------|------|--------|
| 4 | [e_atis_certificate_policies](ISSUES/e_atis_certificate_policies/README.md) | ATIS1000080 |
| 1 | [e_atis_crl_distribution](ISSUES/e_atis_crl_distribution/README.md) | ATIS1000080 |
| 4 | [e_atis_subject_cn](ISSUES/e_atis_subject_cn/README.md) | ATIS1000080 |
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
- 7379 days is the average remaining validity for the certificates in the corpus
- 6935 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days

| Instances | Test | Source |
|-----------|------|--------|
| 2 | [e_atis_ca_signature_algorithm](ISSUES/e_atis_ca_signature_algorithm/README.md) | ATIS1000080 |
| 2 | [e_atis_ca_subject_public_key](ISSUES/e_atis_ca_subject_public_key/README.md) | ATIS1000080 |
| 2 | [w_shaken_ca_subject_rdn_unknown](ISSUES/w_shaken_ca_subject_rdn_unknown/README.md) | SHAKEN_PKI_BEST_PRACTICES |

### Details

#### Leaf Certificates

| Created at | Subject | Problems | Link |
|------------|---------|----------|------|
| 10 Jun 22 00:00 UTC | Plivo Inc | true | [view](CERTS/7dc750fb7aa68d2b67b8dbc89f65217f92db54504685058be016638011adf8bf/README.md) |
| 25 Oct 22 01:08 UTC | Google SHAKEN cert 969H | true | [view](CERTS/ea9be023aa06e4b6606c9048a6ccace9a16100695fe5b5747654a521d7acfc56/README.md) |
| 26 Oct 22 15:00 UTC | Baltimore-Washington Telephone Company SHAKEN cert 8697 | true | [view](CERTS/0541bcf0e982d7e6cb1c0bbbcce11b835523895f5816aab422fd0fbd9d8f2ffb/README.md) |
| 02 Nov 22 13:58 UTC | Google SHAKEN cert 969H | true | [view](CERTS/ac89ca8ee8dd2dad340232968a8cdcc5a613a485c108f8f14f14bd8cbbc1433e/README.md) |

#### CA Certificates

| Created at | Subject | Problems | Link |
|------------|---------|----------|------|
| 12 Jul 21 23:25 UTC | NetNumber SHAKEN Root CA | false | [view](CERTS/7ac80e8481ecb019dc95484016842db78686069efbc0f703e7f39310217b6157/README.md) |
| 27 Sep 21 19:45 UTC | NetNumber SHAKEN Root CA 1 | true | [view](CERTS/2dd1386ca717f31d550b35b9bce9daa9b02483bcdb98bdfcfca07202276136d7/README.md) |
| 29 Sep 21 13:22 UTC | NetNumber SHAKEN Root Intermediate CA 1 | true | [view](CERTS/e449803766edf02ab50b034dd7e89e54efd332cce87688a032f89b340d039878/README.md) |


Generated: 17 Nov 22 19:20 UTC