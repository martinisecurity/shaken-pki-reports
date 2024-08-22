# STIR/SHAKEN CA Ecosystem Compliance

## GBSDTech

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 7 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 2 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 5 certificates being tested against the remaining rules
- 3.40 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 0.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 544 days is the average remaining validity for the certificates in the corpus
- 510 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days
- 1.00 average number of unexpired certificates per OCN observed
- 5 unique OCNs observed in unexpired and valid certificate corpus

| Instances | Test | Source |
|-----------|------|--------|
| 1 | [e_atis_ext_certificate_policies](ISSUES/e_atis_ext_certificate_policies/README.md) | ATIS1000080 |
| 1 | [e_atis_ext_crl_distribution](ISSUES/e_atis_ext_crl_distribution/README.md) | ATIS1000080 |
| 5 | [e_atis_serial_number_size](ISSUES/e_atis_serial_number_size/README.md) | ATIS1000080 |
| 2 | [e_atis_subject_cn](ISSUES/e_atis_subject_cn/README.md) | ATIS1000080 |
| 5 | [e_atis_subject_cn_spc](ISSUES/e_atis_subject_cn_spc/README.md) | ATIS1000080 |
| 1 | [e_atis_tn_auth_list](ISSUES/e_atis_tn_auth_list/README.md) | ATIS1000080 |
| 1 | [e_atis_tn_auth_list_spc_format](ISSUES/e_atis_tn_auth_list_spc_format/README.md) | ATIS1000080 |
| 1 | [e_shaken_certificate_policies_id](ISSUES/e_shaken_certificate_policies_id/README.md) | US_SHAKEN_CP |

#### CA Certificates

- 3 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 0 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 3 certificates being tested against the remaining rules
- 3.00 issues on average found in unexpired, trusted, and non-compliant certificates
- 66.67% of certificates contain one or more Error level issue
- 0.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 66.67% of certificates are too old to be assessed against currently enforced expectations
- 6999 days is the average remaining validity for the certificates in the corpus
- 7300 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days

| Instances | Test | Source |
|-----------|------|--------|
| 1 | [e_atis_ext_authority_key_identifier_ca](ISSUES/e_atis_ext_authority_key_identifier_ca/README.md) | ATIS1000080 |
| 1 | [e_atis_ext_certificate_policies_ca](ISSUES/e_atis_ext_certificate_policies_ca/README.md) | ATIS1000080 |
| 1 | [e_atis_ext_crl_distribution_ca](ISSUES/e_atis_ext_crl_distribution_ca/README.md) | ATIS1000080 |
| 1 | [e_atis_serial_number_size_ca](ISSUES/e_atis_serial_number_size_ca/README.md) | ATIS1000080 |
| 1 | [e_ext_authority_key_identifier_missing](ISSUES/e_ext_authority_key_identifier_missing/README.md) | RFC5280 |
| 1 | [e_ext_authority_key_identifier_no_key_identifier](ISSUES/e_ext_authority_key_identifier_no_key_identifier/README.md) | RFC5280 |

### Details

#### Leaf Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 03&#160;Aug&#160;23&#160;19:55&#160;UTC | MYPBXManager SHAKEN | 02&#160;Aug&#160;26&#160;19:55&#160;UTC | true | [view](CERTS/ea5813855308274fae05fdcae622a159efa47cde2ccf87a9cdf09d9ef43d93f2/README.md) |
| 30&#160;Jan&#160;24&#160;19:17&#160;UTC | ccctelecom.com | 28&#160;Jan&#160;25&#160;19:17&#160;UTC | true | [view](CERTS/e485e7069de2a669c6cd53d3eb10263d437b1c616c03e7e2997a531d020361ad/README.md) |
| 15&#160;Apr&#160;24&#160;22:26&#160;UTC | NovoLink SHAKEN cert | 14&#160;Apr&#160;25&#160;22:26&#160;UTC | true | [view](CERTS/2a55030811ea276fc007055026d66293c349f35c5eebc7dd7b305b65f3b40c56/README.md) |
| 22&#160;Apr&#160;24&#160;20:43&#160;UTC | alluretelecom.com | 21&#160;Apr&#160;25&#160;20:43&#160;UTC | true | [view](CERTS/5486360ac5b339b588547ee88d42ca59ee3cdd87a0911ce2d16c3a8dd376efe0/README.md) |
| 23&#160;May&#160;24&#160;16:51&#160;UTC | FracTEL LLC SHAKEN | 22&#160;May&#160;25&#160;16:51&#160;UTC | true | [view](CERTS/436ab7f1876a96dfebbc3cdb0880c6f4dc9432a1db1bec8745e39d13d4febc22/README.md) |

#### CA Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 05&#160;May&#160;21&#160;19:05&#160;UTC | GBSDTech SHAKEN Root CA | 30&#160;Apr&#160;41&#160;19:05&#160;UTC | false | [view](CERTS/6d2bee73a01c1c9fe92273ff8ba56e0c870b7b901cbebcc9e12226fc109e1af9/README.md) |
| 05&#160;May&#160;21&#160;20:22&#160;UTC | 1RouteGroup SHAKEN Intermediate CA | 29&#160;Apr&#160;41&#160;20:22&#160;UTC | true | [view](CERTS/99e9a67644a30ebc33ecc9aa8df6335524d49a4691164e357c5d2406b58a578e/README.md) |
| 27&#160;Dec&#160;23&#160;06:18&#160;UTC | SHAKEN 1RouteGroup Intermediate CA | 22&#160;Dec&#160;43&#160;06:18&#160;UTC | true | [view](CERTS/de6fcb6ba446b8b669a53ed338795ce54b2e3f5be20320fef1644e10fe3b9c3a/README.md) |


Generated: 22 Aug 24 15:44 UTC