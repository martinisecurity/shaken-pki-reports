# STIR/SHAKEN CA Ecosystem Compliance

## T-Mobile

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### Leaf Certificates

- 3 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 2 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 1 certificates being tested against the remaining rules
- 1.00 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 0.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 328 days is the average remaining validity for the certificates in the corpus
- 366 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days
- 1.00 average number of unexpired certificates per OCN observed
- 1 unique OCNs observed in unexpired and valid certificate corpus

| Instances | Test | Source |
|-----------|------|--------|
| 1 | [e_us_cp_subject_sn_shall](ISSUES/e_us_cp_subject_sn_shall/README.md) | US_SHAKEN_CP |

#### CA Certificates

- 4 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 0 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 4 certificates being tested against the remaining rules
- 1.00 issues on average found in unexpired, trusted, and non-compliant certificates
- 25.00% of certificates contain one or more Error level issue
- 0.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 100.00% of certificates are too old to be assessed against currently enforced expectations
- 5903 days is the average remaining validity for the certificates in the corpus
- 5478 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days

| Instances | Test | Source |
|-----------|------|--------|
| 1 | [e_atis_subject_cn_ca](ISSUES/e_atis_subject_cn_ca/README.md) | ATIS1000080 |

### Details

#### Leaf Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 12&#160;Sep&#160;23&#160;14:17&#160;UTC | SHAKEN 6529 | 11&#160;Sep&#160;24&#160;14:47&#160;UTC | true | [view](CERTS/ecd9d4ee9cf6d3fa4727df3dbf725fff3b1a0928545039fc00f0a5cc84d65f13/README.md) |

#### CA Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 19&#160;Sep&#160;19&#160;20:12&#160;UTC | TMOBILE-PROD-ROOT-STIRSHAKEN-EC | 18&#160;Sep&#160;44&#160;20:12&#160;UTC | false | [view](CERTS/d54b8c44268da3eaee9c5483c289652d1bd7f82420891114475470adebf8bf1e/README.md) |
| 19&#160;Sep&#160;19&#160;20:12&#160;UTC | TMOBILE-PROD-ROOT-STIRSHAKEN-EC | 18&#160;Sep&#160;44&#160;20:12&#160;UTC | false | [view](CERTS/c12a1d22654cc7c8bdcd4d502979b061750e11291d529a13a0b71bb7e96fcaff/README.md) |
| 27&#160;Sep&#160;19&#160;17:12&#160;UTC | TMOBILE-PROD-SUB-STIRSHAKEN-EC | 25&#160;Sep&#160;24&#160;17:42&#160;UTC | false | [view](CERTS/4378a3f136510465232246b4d6a72d600db7e11117ac0e3d59c65528f47c9c4f/README.md) |
| 27&#160;Oct&#160;22&#160;21:18&#160;UTC | TMOBILE-PROD-SUB-STIRSHAKEN-EC | 26&#160;Oct&#160;27&#160;21:48&#160;UTC | true | [view](CERTS/a22dda815630c32b2fa32fb3483ded024fe4d333b6865bf47dbb00a5194472ad/README.md) |


Generated: 27 Nov 23 22:56 UTC