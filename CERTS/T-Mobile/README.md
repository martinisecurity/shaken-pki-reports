# STIR/SHAKEN CA Ecosystem Compliance

## T-Mobile

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 2 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 0 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 2 certificates being tested against the remaining rules
- 4.50 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 100.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 359 days is the average remaining validity for the certificates in the corpus
- 366 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days
- 2.00 average number of unexpired certificates per OCN observed
- 1 unique OCNs observed in unexpired and valid certificate corpus

| Instances | Test | Source |
|-----------|------|--------|
| 1 | [e_atis_basic_constraints](ISSUES/e_atis_basic_constraints/README.md) | ATIS1000080 |
| 1 | [e_atis_certificate_policies](ISSUES/e_atis_certificate_policies/README.md) | ATIS1000080 |
| 1 | [e_atis_extension_unknown](ISSUES/e_atis_extension_unknown/README.md) | ATIS1000080 |
| 1 | [e_atis_subject_cn](ISSUES/e_atis_subject_cn/README.md) | ATIS1000080 |
| 1 | [e_us_cp_ambiguous_identifier](ISSUES/e_us_cp_ambiguous_identifier/README.md) | US_SHAKEN_CP |
| 2 | [e_us_cp_subject_sn](ISSUES/e_us_cp_subject_sn/README.md) | US_SHAKEN_CP |
| 2 | [w_shaken_subject_rdn_unknown](ISSUES/w_shaken_subject_rdn_unknown/README.md) | SHAKEN_PKI_BEST_PRACTICES |

#### CA Certificates

- 4 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 0 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 4 certificates being tested against the remaining rules
- 2.00 issues on average found in unexpired, trusted, and non-compliant certificates
- 25.00% of certificates contain one or more Error level issue
- 0.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 75.00% of certificates are too old to be assessed against currently enforced expectations
- 5977 days is the average remaining validity for the certificates in the corpus
- 5478 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days

| Instances | Test | Source |
|-----------|------|--------|
| 1 | [e_atis_ca_serial_number](ISSUES/e_atis_ca_serial_number/README.md) | ATIS1000080 |
| 1 | [e_us_cp_ca_key_usage_crl_sign](ISSUES/e_us_cp_ca_key_usage_crl_sign/README.md) | US_SHAKEN_CP |

### Details

#### Leaf Certificates

| Created at | Subject | Problems | Link |
|------------|---------|----------|------|
| 14 Jul 22 21:05 UTC | cert.stir.t-mobile.com | true | [view](CERTS/7f653e15453416082531011acd1d7dad4f664ddf5124f73e27d841138f4a89f8/README.md) |
| 01 Nov 22 13:46 UTC | SHAKEN 6529 | true | [view](CERTS/a3bffabdf710ee8fd719f4bf09ec27341e7e9705b4ac4687b98e4d137222cf38/README.md) |

#### CA Certificates

| Created at | Subject | Problems | Link |
|------------|---------|----------|------|
| 19 Sep 19 20:12 UTC | TMOBILE-PROD-ROOT-STIRSHAKEN-EC | false | [view](CERTS/d54b8c44268da3eaee9c5483c289652d1bd7f82420891114475470adebf8bf1e/README.md) |
| 19 Sep 19 20:12 UTC | TMOBILE-PROD-ROOT-STIRSHAKEN-EC | false | [view](CERTS/c12a1d22654cc7c8bdcd4d502979b061750e11291d529a13a0b71bb7e96fcaff/README.md) |
| 27 Sep 19 17:12 UTC | TMOBILE-PROD-SUB-STIRSHAKEN-EC | false | [view](CERTS/4378a3f136510465232246b4d6a72d600db7e11117ac0e3d59c65528f47c9c4f/README.md) |
| 27 Oct 22 21:18 UTC | TMOBILE-PROD-SUB-STIRSHAKEN-EC | true | [view](CERTS/a22dda815630c32b2fa32fb3483ded024fe4d333b6865bf47dbb00a5194472ad/README.md) |


Generated: 21 Nov 22 20:55 UTC