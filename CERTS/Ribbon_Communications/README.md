# STIR/SHAKEN CA Ecosystem Compliance

## Ribbon Communications

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 5 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 0 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 5 certificates being tested against the remaining rules
- 5.00 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 100.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 334 days is the average remaining validity for the certificates in the corpus
- 365 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days
- 1.00 average number of unexpired certificates per OCN observed
- 5 unique OCNs observed in unexpired and valid certificate corpus

| Instances | Test | Source |
|-----------|------|--------|
| 5 | [e_atis_certificate_policies](ISSUES/e_atis_certificate_policies/README.md) | ATIS1000080 |
| 5 | [e_atis_extension_unknown](ISSUES/e_atis_extension_unknown/README.md) | ATIS1000080 |
| 5 | [e_us_cp_ambiguous_identifier](ISSUES/e_us_cp_ambiguous_identifier/README.md) | US_SHAKEN_CP |
| 5 | [e_us_cp_subject_sn](ISSUES/e_us_cp_subject_sn/README.md) | US_SHAKEN_CP |
| 5 | [w_shaken_subject_rdn_unknown](ISSUES/w_shaken_subject_rdn_unknown/README.md) | SHAKEN_PKI_BEST_PRACTICES |

#### CA Certificates

- 2 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 0 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 2 certificates being tested against the remaining rules
- 2.00 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 100.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 100.00% of certificates are too old to be assessed against currently enforced expectations
- 7354 days is the average remaining validity for the certificates in the corpus
- 6757 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days

| Instances | Test | Source |
|-----------|------|--------|
| 2 | [e_us_cp_ca_key_usage_crl_sign](ISSUES/e_us_cp_ca_key_usage_crl_sign/README.md) | US_SHAKEN_CP |
| 2 | [w_shaken_ca_subject_rdn_unknown](ISSUES/w_shaken_ca_subject_rdn_unknown/README.md) | SHAKEN_PKI_BEST_PRACTICES |

### Details

#### Leaf Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 29&#160;Apr&#160;22&#160;17:43&#160;UTC | Nuwave Communications SHAKEN 620J | 29&#160;Apr&#160;23&#160;17:42&#160;UTC | true | [view](CERTS/885b7fea5177d66a28bd5dea525144479428e27d2ed1ff7c493d767c73173fbc/README.md) |
| 20&#160;May&#160;22&#160;14:32&#160;UTC | Veracity SHAKEN 716D | 20&#160;May&#160;23&#160;14:32&#160;UTC | true | [view](CERTS/fccfbb719ccc9513231e9ea6f38906f0271f4640f253e1be0780da1e7b5f03ff/README.md) |
| 10&#160;Jun&#160;22&#160;19:00&#160;UTC | Netfortris SHAKEN 8886 | 10&#160;Jun&#160;23&#160;19:00&#160;UTC | true | [view](CERTS/64b7a3eed364e863b36e050a95b35799c594ea1c15c17562611907a2f0dd8bbe/README.md) |
| 16&#160;Jun&#160;22&#160;18:50&#160;UTC | Peoples SHAKEN 2130 | 16&#160;Jun&#160;23&#160;18:50&#160;UTC | true | [view](CERTS/5c8b37f7c0235bb218b1f90787953048164805b88f222f7391c7b319b42c1d14/README.md) |
| 22&#160;Aug&#160;22&#160;16:04&#160;UTC | Siskiyou SHAKEN 2339 | 22&#160;Aug&#160;23&#160;16:04&#160;UTC | true | [view](CERTS/dc3afdbb361402b02a75c92bb00cca9316fcbec50ad3c3093eaff09911ed3220/README.md) |

#### CA Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 13&#160;May&#160;21&#160;00:00&#160;UTC | SHAKEN Ribbon Issuing CA | 12&#160;May&#160;33&#160;23:59&#160;UTC | true | [view](CERTS/05a71a04eaedbdf4b0534f40768616d7c19c8deb5a3aefd1f4a04b3aab55a48f/README.md) |
| 13&#160;May&#160;21&#160;00:00&#160;UTC | SHAKEN Ribbon Root CA | 12&#160;May&#160;46&#160;23:59&#160;UTC | true | [view](CERTS/7c2799d3642d04f04fe667c3ab251c18689af323acdc43b2fa5f3dc89e3a0f14/README.md) |


Generated: 17 Dec 22 00:12 UTC