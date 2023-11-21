# STIR/SHAKEN CA Ecosystem Compliance

## CTIA

### Summary

\* The percent of errors, warnings and notices is calculated against total observed unexpired and trusted certificates from the specified issuer.\
\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran.

#### CA Certificates

- 1 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 0 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 1 certificates being tested against the remaining rules
- 3.00 issues on average found in unexpired, trusted, and non-compliant certificates
- 100.00% of certificates contain one or more Error level issue
- 0.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 0.00% of certificates are too old to be assessed against currently enforced expectations
- 9048 days is the average remaining validity for the certificates in the corpus
- 9125 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days

| Instances | Test | Source |
|-----------|------|--------|
| 1 | [e_atis_ca_signature_algorithm](ISSUES/e_atis_ca_signature_algorithm/README.md) | ATIS1000080 |
| 1 | [e_atis_ca_subject_public_key](ISSUES/e_atis_ca_subject_public_key/README.md) | ATIS1000080 |
| 1 | [e_atis_subject_cn_root](ISSUES/e_atis_subject_cn_root/README.md) | ATIS1000080 |

### Details

#### CA Certificates

| Created At | Subject | Not After | Problems | Link |
|------------|---------|-----------|----------|------|
| 21&#160;Jun&#160;23&#160;13:15&#160;UTC | SHAKEN Root CA | 14&#160;Jun&#160;48&#160;13:15&#160;UTC | true | [view](CERTS/ee1cf83becad4777dcf250170efecc7fc7498d85097e9a570dfac542151ebf53/README.md) |


Generated: 21 Nov 23 19:18 UTC