# STIR/SHAKEN CA Ecosystem Compliance

## GBSDTech

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 1 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 0 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 1 certificates being tested against the remaining rules
- 7.00 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 0.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 287 days is the average remaining validity for the certificates in the corpus
- 365 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days
- 1.00 average number of unexpired certificates per OCN observed
- 1 unique OCNs observed in unexpired and valid certificate corpus

| Instances | Test | Source |
|-----------|------|--------|
| 1 | [e_atis_certificate_policies](ISSUES/e_atis_certificate_policies/README.md) | ATIS1000080 |
| 1 | [e_atis_crl_distribution](ISSUES/e_atis_crl_distribution/README.md) | ATIS1000080 |
| 1 | [e_atis_serial_number](ISSUES/e_atis_serial_number/README.md) | ATIS1000080 |
| 1 | [e_atis_subject_cn](ISSUES/e_atis_subject_cn/README.md) | ATIS1000080 |
| 1 | [e_atis_tn_auth_list](ISSUES/e_atis_tn_auth_list/README.md) | ATIS1000080 |
| 1 | [e_us_cp_ambiguous_identifier](ISSUES/e_us_cp_ambiguous_identifier/README.md) | US_SHAKEN_CP |
| 1 | [e_us_cp_subject_sn](ISSUES/e_us_cp_subject_sn/README.md) | US_SHAKEN_CP |

#### CA Certificates

- 2 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 0 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 2 certificates being tested against the remaining rules
- 3.00 issues on average found in unexpired, trusted, and non-compliant certificates
- 50.00% of certificates contain one or more Error level issue
- 0.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 100.00% of certificates are too old to be assessed against currently enforced expectations
- 7108 days is the average remaining validity for the certificates in the corpus
- 7300 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days

| Instances | Test | Source |
|-----------|------|--------|
| 1 | [e_atis_ca_authority_key_identifier](ISSUES/e_atis_ca_authority_key_identifier/README.md) | ATIS1000080 |
| 1 | [e_ext_authority_key_identifier_missing](ISSUES/e_ext_authority_key_identifier_missing/README.md) | RFC5280 |
| 1 | [e_ext_authority_key_identifier_no_key_identifier](ISSUES/e_ext_authority_key_identifier_no_key_identifier/README.md) | RFC5280 |

### Details

#### Leaf Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 28&#160;Jun&#160;22&#160;18:13&#160;UTC | MYPBXManager SHAKEN | 28&#160;Jun&#160;23&#160;18:13&#160;UTC | true | [view](CERTS/0c7bf2cc1741b8036003876afadd109dfd5a6b0fb7af3662ae4d02e3340ad9ce/README.md) |

#### CA Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 05&#160;May&#160;21&#160;19:05&#160;UTC | GBSDTech SHAKEN Root CA | 30&#160;Apr&#160;41&#160;19:05&#160;UTC | false | [view](CERTS/6d2bee73a01c1c9fe92273ff8ba56e0c870b7b901cbebcc9e12226fc109e1af9/README.md) |
| 05&#160;May&#160;21&#160;20:22&#160;UTC | 1RouteGroup SHAKEN Intermediate CA | 29&#160;Apr&#160;41&#160;20:22&#160;UTC | true | [view](CERTS/99e9a67644a30ebc33ecc9aa8df6335524d49a4691164e357c5d2406b58a578e/README.md) |


Generated: 02 Dec 22 07:30 UTC